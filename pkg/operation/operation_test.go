package operation

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"
	
	"github.com/google/go-cmp/cmp"
	testing_interface "github.com/mitchellh/go-testing-interface"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	grpc_status "google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/protocmp"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/durationpb"
	
	"github.com/yandex-cloud/go-genproto/yandex/cloud/compute/v1"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
)

var _ PollIntervalFunc = defaultPollIntervalFunc

func TestWaitReturnsPollError(t *testing.T) {
	pollErr := errors.New("poll error")
	poll := func(context.Context, string, ...grpc.CallOption) (YCOperation, error) {
		return nil, pollErr
	}
	
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	
	_, err := (&Operation{&operation.Operation{}, &Concretization{poll, nil, nil, nil}, nil, nil, nil}).Wait(ctx)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), pollErr.Error())
}

func TestWaitReturnsNotFound(t *testing.T) {
	poll := func(context.Context, string, ...grpc.CallOption) (YCOperation, error) {
		return nil, grpc_status.Error(codes.NotFound, "NotFound")
	}
	
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	
	_, err := (&Operation{&operation.Operation{}, &Concretization{poll, nil, nil, nil}, nil, nil, nil}).Wait(ctx)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), context.DeadlineExceeded.Error())
}

func TestWaitReturnsPollTimeout(t *testing.T) {
	poll := func(context.Context, string, ...grpc.CallOption) (YCOperation, error) {
		return &operation.Operation{Done: false}, nil
	}
	
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	
	_, err := (&Operation{&operation.Operation{}, &Concretization{poll, nil, nil, nil}, nil, nil, nil}).Wait(ctx)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), context.DeadlineExceeded.Error())
}

func TestReturnsOperationError(t *testing.T) {
	errorDesc := "error description"
	poll := func(context.Context, string, ...grpc.CallOption) (YCOperation, error) {
		return &operation.Operation{
			Done: true,
			Result: &operation.Operation_Error{
				Error: &status.Status{
					Code:    int32(codes.Internal),
					Message: errorDesc,
				},
			},
		}, nil
	}
	
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	
	_, err := (&Operation{&operation.Operation{}, &Concretization{poll, nil, nil, nil}, nil, nil, nil}).Wait(ctx)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), errors.New(errorDesc).Error())
}

func TestWaitReturnsErrorOnIncorrectResult(t *testing.T) {
	poll := func(context.Context, string, ...grpc.CallOption) (YCOperation, error) {
		return &operation.Operation{
			Done: true,
			Result: &operation.Operation_Response{
				Response: &anypb.Any{},
			},
		}, nil
	}
	
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	
	_, err := (&Operation{&operation.Operation{}, &Concretization{poll, nil, nil, nil}, nil, nil, nil}).Wait(ctx)
	assert.NotNil(t, err)
}

func TestPollUpdatesMetadata(t *testing.T) {
	initialMeta := &durationpb.Duration{Seconds: 33}
	updatedMeta := &durationpb.Duration{Seconds: 100}
	poll := func(context.Context, string, ...grpc.CallOption) (YCOperation, error) {
		return &operation.Operation{
			Metadata: Any(t, updatedMeta),
		}, nil
	}
	op, err := NewOperation(
		&operation.Operation{
			Metadata: Any(t, initialMeta),
		},
		&Concretization{poll, &durationpb.Duration{}, &durationpb.Duration{}, nil},
	)
	require.NoError(t, err)
	
	Equal(t, initialMeta, op.Metadata())
	
	err = op.PollOnce(context.Background())
	require.NoError(t, err)
	Equal(t, updatedMeta, op.Metadata())
}

func TestWaitUpdatesMetadata(t *testing.T) {
	initialMeta := &durationpb.Duration{Seconds: 33}
	updatedMeta := &durationpb.Duration{Seconds: 100}
	expectedResp := &durationpb.Duration{Seconds: 10}
	poll := func(context.Context, string, ...grpc.CallOption) (YCOperation, error) {
		return &operation.Operation{
			Done:     true,
			Metadata: Any(t, updatedMeta),
			Result: &operation.Operation_Response{
				Response: Any(t, expectedResp),
			},
		}, nil
	}
	op, err := NewOperation(
		&operation.Operation{
			Metadata: Any(t, initialMeta),
		},
		&Concretization{poll, &durationpb.Duration{}, &durationpb.Duration{}, nil},
	)
	require.NoError(t, err)
	
	Equal(t, initialMeta, op.Metadata())
	
	resp, err := op.Wait(context.Background())
	require.NoError(t, err)
	Equal(t, expectedResp, resp)
	Equal(t, updatedMeta, op.Metadata())
}

func TestFillNilResponse(t *testing.T) {
	op := &Operation{concretization: &Concretization{ResponseType: (*compute.Instance)(nil)}}
	err := op.fillResponse(nil)
	require.Error(t, err)
}

func Any(t testing_interface.T, msg proto.Message) *anypb.Any {
	any, err := anypb.New(msg)
	require.NoError(t, err)
	return any
}

func Equal(t testing_interface.T, expected, actual interface{}, msgAndArgs ...interface{}) bool {
	t.Helper()
	expProto, expOK := expected.(proto.Message)
	actProto, actOK := actual.(proto.Message)
	if expOK && actOK {
		if proto.Equal(expProto, actProto) {
			return true
		}
	}
	
	diff := Diff(expected, actual)
	if diff == "" {
		return true
	}
	if len(msgAndArgs) != 0 {
		t.Errorf(messageFromMsgAndArgs(msgAndArgs)+"\nDiff:\n%s", diff)
	} else {
		t.Errorf("Not equal. Diff:\n%s", diff)
	}
	return false
}

func messageFromMsgAndArgs(msgAndArgs ...interface{}) string {
	if len(msgAndArgs) == 0 {
		return ""
	}
	if len(msgAndArgs) == 1 {
		msg := msgAndArgs[0]
		if msgAsStr, ok := msg.(string); ok {
			return msgAsStr
		}
		return fmt.Sprintf("%+v", msg)
	}
	if len(msgAndArgs) > 1 {
		return fmt.Sprintf(msgAndArgs[0].(string), msgAndArgs[1:]...)
	}
	return ""
}

func Diff(x, y interface{}, os ...cmp.Option) string {
	return cmp.Diff(x, y, append(os, protocmp.Transform())...)
}
