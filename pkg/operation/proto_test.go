package operation

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/durationpb"

	pb "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
)

func TestProtoTracksPolledOperation(t *testing.T) {
	running := &pb.Operation{Id: "operation", Description: "running"}
	done := &pb.Operation{Id: "operation", Description: "completed", Done: true, Result: &pb.Operation_Response{Response: Any(t, &durationpb.Duration{Seconds: 10})}}
	op, err := NewOperation(running, &Concretization{ResponseType: &durationpb.Duration{}, Poll: func(context.Context, string, ...grpc.CallOption) (YCOperation, error) { return done, nil }})
	require.NoError(t, err)
	require.Same(t, running, op.Proto())
	require.NoError(t, op.PollOnce(context.Background()))
	require.Same(t, done, op.Proto())
	require.True(t, op.Proto().GetDone())
}
