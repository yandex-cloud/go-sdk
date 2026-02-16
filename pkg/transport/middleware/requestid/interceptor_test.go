package requestid

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	clientTraceID = "x-client-trace-id"
	requestID     = "x-request-id"
	traceID       = "x-trace-id"
)

type StatusError interface {
	GRPCStatus() *status.Status
}

func requestIDFromError(err error) string {
	if info, ok := RequestIDsFromError(err); ok {
		return info.RequestID
	}

	return ""
}

func clientTraceIDFromError(err error) string {
	if info, ok := RequestIDsFromError(err); ok {
		return info.ClientTraceID
	}

	return ""
}

func responseHeader(serverRequestID, serverTraceID string) metadata.MD {
	return metadata.New(map[string]string{
		requestIDHeader:     serverRequestID,
		serverTraceIDHeader: serverTraceID,
	})
}

func TestWrappedRequestIDs(t *testing.T) {
	t.Run("unwrap normal error", func(t *testing.T) {
		expected := fmt.Errorf("some error")
		errorInfo, ok := RequestIDsFromError(expected)
		assert.False(t, ok)
		assert.Nil(t, errorInfo)
	})
	t.Run("unwrap nil error", func(t *testing.T) {
		errorInfo, ok := RequestIDsFromError(nil)
		assert.False(t, ok)
		assert.Nil(t, errorInfo)
	})
	t.Run("wrap nil error", func(t *testing.T) {
		actual := wrapError(nil, clientTraceID, requestID, nil, nil)
		assert.Nil(t, actual)
	})
	t.Run("wrap err with client request id and nil header", func(t *testing.T) {
		err := fmt.Errorf("some error")
		actual := wrapError(err, clientTraceID, requestID, nil, nil)
		assert.Equal(t, &ErrorWithRequestIDs{err, RequestIDs{ClientTraceID: clientTraceID, RequestID: requestID, ServerTraceID: ""}}, actual)

		errorInfo, ok := RequestIDsFromError(actual)
		assert.True(t, ok)
		assert.Equal(t, requestID, requestIDFromError(actual))
		assert.Equal(t, clientTraceID, clientTraceIDFromError(actual))
		assert.Equal(t, requestID, errorInfo.RequestID)
		assert.Equal(t, "", errorInfo.ServerTraceID)
	})
	t.Run("wrap err with client and server request id", func(t *testing.T) {
		err := fmt.Errorf("some error")
		actual := wrapError(err, clientTraceID, requestID, responseHeader(requestID, traceID), nil)
		assert.Equal(t, &ErrorWithRequestIDs{err, RequestIDs{ClientTraceID: clientTraceID, RequestID: requestID, ServerTraceID: traceID}}, actual)

		errorInfo, ok := RequestIDsFromError(actual)
		assert.True(t, ok)
		assert.Equal(t, requestID, errorInfo.RequestID)
		assert.Equal(t, clientTraceID, errorInfo.ClientTraceID)
		assert.Equal(t, traceID, errorInfo.ServerTraceID)
	})
	t.Run("wrap err with empty header", func(t *testing.T) {
		err := fmt.Errorf("some error")
		actual := wrapError(err, clientTraceID, requestID, metadata.New(map[string]string{}), nil)
		assert.Equal(t, &ErrorWithRequestIDs{err, RequestIDs{ClientTraceID: clientTraceID, RequestID: requestID, ServerTraceID: ""}}, actual)

		errorInfo, ok := RequestIDsFromError(actual)
		assert.True(t, ok)
		assert.Equal(t, requestID, errorInfo.RequestID)
		assert.Equal(t, clientTraceID, errorInfo.ClientTraceID)
		assert.Equal(t, "", errorInfo.ServerTraceID)
	})
	t.Run("wrap wrapped", func(t *testing.T) {
		err := fmt.Errorf("some error")
		wrap1 := wrapError(err, "trace1", "id1", nil, nil)
		wrap2 := wrapError(wrap1, "trace1", "id2", nil, nil)
		// Should keep first requestID set
		assert.Equal(t, "id1", requestIDFromError(wrap1))
		assert.Equal(t, "trace1", clientTraceIDFromError(wrap1))
		assert.Equal(t, "id1", requestIDFromError(wrap2))
		assert.Equal(t, "trace1", clientTraceIDFromError(wrap2))
	})
	t.Run("server request id takes precedence over client", func(t *testing.T) {
		err := fmt.Errorf("some error")
		serverReqID := "server-generated-id"
		actual := wrapError(err, clientTraceID, requestID, responseHeader(serverReqID, traceID), nil)
		errorInfo, ok := RequestIDsFromError(actual)
		assert.True(t, ok)
		assert.Equal(t, serverReqID, errorInfo.RequestID)
		assert.Equal(t, traceID, errorInfo.ServerTraceID)
	})
	t.Run("RequestIDsFromError returns copy", func(t *testing.T) {
		err := fmt.Errorf("some error")
		actual := wrapError(err, clientTraceID, requestID, nil, nil)
		errorInfo, ok := RequestIDsFromError(actual)
		require.True(t, ok)
		errorInfo.RequestID = "mutated"
		// Original error should be unchanged
		errorInfo2, ok := RequestIDsFromError(actual)
		require.True(t, ok)
		assert.Equal(t, requestID, errorInfo2.RequestID)
	})
}

func TestErrorUnwrap(t *testing.T) {
	t.Run("errors.Unwrap returns original error", func(t *testing.T) {
		origErr := fmt.Errorf("original error")
		wrapped := wrapError(origErr, clientTraceID, requestID, nil, nil)
		assert.True(t, errors.Is(wrapped, origErr))
		unwrapped := errors.Unwrap(wrapped)
		assert.Equal(t, origErr, unwrapped)
	})
	t.Run("errors.As extracts ErrorWithRequestIDs", func(t *testing.T) {
		origErr := fmt.Errorf("original")
		wrapped := wrapError(origErr, clientTraceID, requestID, nil, nil)
		var withIDs *ErrorWithRequestIDs
		assert.True(t, errors.As(wrapped, &withIDs))
		assert.Equal(t, requestID, withIDs.IDs.RequestID)
		assert.Equal(t, origErr, withIDs.OrigErr)
	})
}

func TestAddRequestID(t *testing.T) {
	t.Run("no outgoing context", func(t *testing.T) {
		ctx := withMetadata(context.Background(), map[string]string{
			requestIDHeader:     requestID,
			clientTraceIDHeader: clientTraceID,
		})
		md, ok := metadata.FromOutgoingContext(ctx)
		require.True(t, ok)
		assert.Equal(t, metadata.New(map[string]string{
			clientTraceIDHeader: clientTraceID,
			requestIDHeader:     requestID,
		}), md)
	})
	t.Run("with outgoing context", func(t *testing.T) {
		ctx := context.Background()
		ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("it-very-long-header", "foobar"))
		ctx = withMetadata(ctx, map[string]string{
			requestIDHeader:     requestID,
			clientTraceIDHeader: clientTraceID,
		})
		md, ok := metadata.FromOutgoingContext(ctx)
		require.True(t, ok)
		assert.Equal(t, metadata.New(map[string]string{
			clientTraceIDHeader:   clientTraceID,
			requestIDHeader:       requestID,
			"it-very-long-header": "foobar",
		}), md)
	})
	t.Run("with old request-id", func(t *testing.T) {
		ctx := context.Background()
		ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs(requestIDHeader, "old"))
		ctx = withMetadata(ctx, map[string]string{
			requestIDHeader:     requestID,
			clientTraceIDHeader: clientTraceID,
		})
		md, ok := metadata.FromOutgoingContext(ctx)
		require.True(t, ok)
		assert.Equal(t, metadata.New(map[string]string{
			clientTraceIDHeader: clientTraceID,
			requestIDHeader:     requestID,
		}), md)
	})
	t.Run("several ids", func(t *testing.T) {
		ctx := context.Background()
		ctx1 := withMetadata(ctx, map[string]string{
			requestIDHeader:     "id1",
			clientTraceIDHeader: "trace1",
		})
		md, ok := metadata.FromOutgoingContext(ctx1)
		require.True(t, ok)
		assert.Equal(t, metadata.New(map[string]string{
			clientTraceIDHeader: "trace1",
			requestIDHeader:     "id1",
		}), md)

		ctx2 := withMetadata(ctx1, map[string]string{
			requestIDHeader:     "id2",
			clientTraceIDHeader: "trace1",
		})
		md, ok = metadata.FromOutgoingContext(ctx2)
		require.True(t, ok)
		assert.Equal(t, metadata.New(map[string]string{
			clientTraceIDHeader: "trace1",
			requestIDHeader:     "id2",
		}), md)
		// Original context not damaged
		md, ok = metadata.FromOutgoingContext(ctx1)
		require.True(t, ok)
		assert.Equal(t, metadata.New(map[string]string{
			clientTraceIDHeader: "trace1",
			requestIDHeader:     "id1",
		}), md)
	})
}

func TestWrappedErrorImplGRPCStatus(t *testing.T) {
	t.Run("wrapped error impl StatusError interface", func(t *testing.T) {
		err := fmt.Errorf("some error")
		actual := wrapError(err, clientTraceID, requestID, nil, nil)
		assert.Equal(t, &ErrorWithRequestIDs{err, RequestIDs{ClientTraceID: clientTraceID, RequestID: requestID, ServerTraceID: ""}}, actual)
		assert.Implements(t, (*StatusError)(nil), actual)
	})
	t.Run("get status by status.FromError method", func(t *testing.T) {
		err := fmt.Errorf("some error")
		actual := wrapError(err, clientTraceID, requestID, nil, nil)
		st, ok := status.FromError(actual)
		assert.True(t, ok)
		assert.Equal(t, codes.Unknown, st.Code())
	})
	t.Run("wrap status error", func(t *testing.T) {
		sErr := status.Error(codes.Aborted, "request aborted")
		actual := wrapError(sErr, clientTraceID, requestID, nil, nil)
		st, ok := status.FromError(actual)
		assert.True(t, ok)
		assert.Equal(t, "request aborted", st.Message())
		assert.Equal(t, codes.Aborted, st.Code())
	})
}
