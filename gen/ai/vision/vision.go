package vision

import (
	"context"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/ai/vision/v1"
	vis "github.com/yandex-cloud/go-genproto/yandex/cloud/ai/vision/v1"
	"google.golang.org/grpc"
)

// Vision is a functions.Vision with
// lazy GRPC connection initialization.
type Vision struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// NewVision creates instance of vision
func NewVision(g func(ctx context.Context) (*grpc.ClientConn, error)) *Vision {
	return &Vision{g}
}

// BatchAnalyze implements vision.Translate
func (c *Vision) BatchAnalyze(ctx context.Context, in *vis.BatchAnalyzeRequest, opts ...grpc.CallOption) (*vis.BatchAnalyzeResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return vision.NewVisionServiceClient(conn).BatchAnalyze(ctx, in, opts...)
}
