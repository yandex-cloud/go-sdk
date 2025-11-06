package requestid

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// clientTraceIDHeader is the header key used for identifying the client trace ID.
// clientRequestIDHeader is the header key used for identifying the client request ID.
// serverRequestIDHeader is the header key used for identifying the server request ID.
// serverTraceIDHeader is the header key used for identifying the server trace ID.
const (
	clientTraceIDHeader   = "x-client-trace-id"
	clientRequestIDHeader = "x-client-request-id"
	serverRequestIDHeader = "x-request-id"
	serverTraceIDHeader   = "x-server-trace-id"
)

// Interceptor is a gRPC unary client interceptor for propagating trace and request IDs across service boundaries.
// It injects client-side headers and captures server-side headers to track the request lifecycle.
// The interceptor appends metadata and wraps errors with correlated request and trace information.
func Interceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req interface{}, reply interface{}, conn *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		clientTraceID := uuid.New().String()
		clientRequestID := uuid.New().String()

		md, ok := metadata.FromOutgoingContext(ctx)
		if ok && len(md.Get(clientTraceIDHeader)) > 0 {
			clientTraceID = md.Get(clientTraceIDHeader)[0]
		}

		ctx = withMetadata(ctx, map[string]string{
			clientRequestIDHeader: clientRequestID,
			clientTraceIDHeader:   clientTraceID,
		})

		var responseHeader metadata.MD
		opts = append(opts, grpc.Header(&responseHeader))
		err := invoker(ctx, method, req, reply, conn, opts...)

		return wrapError(err, clientTraceID, clientRequestID, responseHeader)
	}
}

// RequestIDs represents a collection of identifiers for tracking client and server request and trace information.
// ClientTraceID is a unique identifier for tracing client requests.
// ClientRequestID is a unique identifier for individual client requests.
// ServerRequestID is a unique identifier for server requests, typically derived from server response headers.
// ServerTraceID is a unique trace identifier for server-side operations, typically derived from server response headers.
type RequestIDs struct {
	ClientTraceID   string
	ClientRequestID string
	ServerRequestID string
	ServerTraceID   string
}

// ErrorWithRequestIDs wraps an original error and associates it with client and server request/trace IDs.
type ErrorWithRequestIDs struct {
	OrigErr error
	IDs     RequestIDs
}

// Error returns a formatted error message including server and client request/trace IDs, followed by the original error message.
func (e *ErrorWithRequestIDs) Error() (msg string) {
	if e.IDs.ServerRequestID != "" {
		msg += fmt.Sprintf("server-request-id = %s ", e.IDs.ServerRequestID)
	}

	if e.IDs.ServerTraceID != "" {
		msg += fmt.Sprintf("server-trace-id = %s ", e.IDs.ServerTraceID)
	}

	if e.IDs.ClientRequestID != "" {
		msg += fmt.Sprintf("client-request-id = %s ", e.IDs.ClientRequestID)
	}

	if e.IDs.ClientTraceID != "" {
		msg += fmt.Sprintf("client-trace-id = %s ", e.IDs.ClientTraceID)
	}

	return msg + e.OrigErr.Error()
}

// GRPCStatus converts the original error into a gRPC status representation and returns the associated status.
func (e ErrorWithRequestIDs) GRPCStatus() *status.Status {
	return status.Convert(e.OrigErr)
}

// RequestIDsFromError extracts RequestIDs from an error if it contains request-related context. Returns ok as true if successful.
func RequestIDsFromError(err error) (*RequestIDs, bool) {
	if withID, ok := err.(*ErrorWithRequestIDs); ok {
		return &withID.IDs, ok
	}

	return nil, false
}

// ContextWithClientTraceID returns a new context embedding the provided clientTraceID for outgoing requests.
func ContextWithClientTraceID(ctx context.Context, clientTraceID string) context.Context {
	return withMetadata(ctx, map[string]string{
		clientTraceIDHeader: clientTraceID,
	})
}

// wrapError wraps an error with client and server request/trace IDs, enriching error context with additional metadata.
// If the error is already wrapped with request IDs, it returns the original error. Returns nil if the input error is nil.
func wrapError(err error, clientTraceID, clientRequestID string, responseHeader metadata.MD) error {
	if err == nil {
		return nil
	}

	if _, ok := err.(*ErrorWithRequestIDs); ok {
		return err
	}

	serverRequestID := getServerHeader(responseHeader, serverRequestIDHeader)
	serverTraceID := getServerHeader(responseHeader, serverTraceIDHeader)

	return &ErrorWithRequestIDs{
		err,
		RequestIDs{
			ClientTraceID:   clientTraceID,
			ClientRequestID: clientRequestID,
			ServerRequestID: serverRequestID,
			ServerTraceID:   serverTraceID,
		},
	}
}

// getServerHeader retrieves the first value associated with the given key from the metadata.MD headers.
// Returns an empty string if the key is not present.
func getServerHeader(responseHeader metadata.MD, key string) string {
	serverHeaderIDRaw := responseHeader.Get(key)
	if len(serverHeaderIDRaw) == 0 {
		return ""
	}

	return serverHeaderIDRaw[0]
}

// withMetadata adds the specified metadata to the outgoing gRPC context and returns the updated context.
func withMetadata(ctx context.Context, meta map[string]string) context.Context {
	md, ok := metadata.FromOutgoingContext(ctx)
	if !ok {
		md = metadata.MD{}
	} else {
		md = md.Copy()
	}

	for k, v := range meta {
		md.Set(k, v)
	}

	return metadata.NewOutgoingContext(ctx, md)
}
