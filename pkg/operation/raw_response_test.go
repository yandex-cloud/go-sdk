package operation

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
)

func TestRawOperationAndResponse(t *testing.T) {
	response := wrapperspb.String("historical response")
	initial := &operation.Operation{Id: "raw-operation"}
	completed := &operation.Operation{Id: initial.Id, Done: true, Description: "completed", Result: &operation.Operation_Response{Response: Any(t, response)}}
	op, err := NewOperation(initial, &Concretization{ResponseType: (*emptypb.Empty)(nil), Poll: func(context.Context, string, ...grpc.CallOption) (YCOperation, error) { return completed, nil }})
	require.NoError(t, err)
	require.Same(t, initial, op.Proto())
	_, err = op.Wait(context.Background())
	require.NoError(t, err)
	require.Same(t, completed, op.Proto())
	raw, err := op.RawResponse()
	require.NoError(t, err)
	Equal(t, response, raw)
	// Typed SDK callers retain the declared Empty response.
	require.IsType(t, &emptypb.Empty{}, op.Response())
}
