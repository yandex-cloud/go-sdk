package operation

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	grpc_status "google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/durationpb"

	"bb.yandex-team.ru/cloud/cloud-go/api/pkg/protoutil/prototest"
	"bb.yandex-team.ru/cloud/cloud-go/genproto/privateapi/yandex/cloud/priv/compute/v1"
	"bb.yandex-team.ru/cloud/cloud-go/genproto/privateapi/yandex/cloud/priv/operation"
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
			Metadata: prototest.Any(t, updatedMeta),
		}, nil
	}
	op, err := NewOperation(
		&operation.Operation{
			Metadata: prototest.Any(t, initialMeta),
		},
		&Concretization{poll, &durationpb.Duration{}, &durationpb.Duration{}, nil},
	)
	require.NoError(t, err)

	prototest.Equal(t, initialMeta, op.Metadata())

	err = op.PollOnce(context.Background())
	require.NoError(t, err)
	prototest.Equal(t, updatedMeta, op.Metadata())
}

func TestWaitUpdatesMetadata(t *testing.T) {
	initialMeta := &durationpb.Duration{Seconds: 33}
	updatedMeta := &durationpb.Duration{Seconds: 100}
	expectedResp := &durationpb.Duration{Seconds: 10}
	poll := func(context.Context, string, ...grpc.CallOption) (YCOperation, error) {
		return &operation.Operation{
			Done:     true,
			Metadata: prototest.Any(t, updatedMeta),
			Result: &operation.Operation_Response{
				Response: prototest.Any(t, expectedResp),
			},
		}, nil
	}
	op, err := NewOperation(
		&operation.Operation{
			Metadata: prototest.Any(t, initialMeta),
		},
		&Concretization{poll, &durationpb.Duration{}, &durationpb.Duration{}, nil},
	)
	require.NoError(t, err)

	prototest.Equal(t, initialMeta, op.Metadata())

	resp, err := op.Wait(context.Background())
	require.NoError(t, err)
	prototest.Equal(t, expectedResp, resp)
	prototest.Equal(t, updatedMeta, op.Metadata())
}

func TestFillNilResponse(t *testing.T) {
	op := &Operation{concretization: &Concretization{ResponseType: (*compute.Instance)(nil)}}
	err := op.fillResponse(nil)
	require.Error(t, err)
}
