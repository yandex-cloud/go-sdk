// Code generated by sdkgen. DO NOT EDIT.

// nolint
package ocr

import (
	"context"

	"google.golang.org/grpc"

	ocr "github.com/yandex-cloud/go-genproto/yandex/cloud/ai/ocr/v1"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
)

//revive:disable

// TextRecognitionAsyncServiceClient is a ocr.TextRecognitionAsyncServiceClient with
// lazy GRPC connection initialization.
type TextRecognitionAsyncServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// GetRecognition implements ocr.TextRecognitionAsyncServiceClient
func (c *TextRecognitionAsyncServiceClient) GetRecognition(ctx context.Context, in *ocr.GetRecognitionRequest, opts ...grpc.CallOption) (ocr.TextRecognitionAsyncService_GetRecognitionClient, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return ocr.NewTextRecognitionAsyncServiceClient(conn).GetRecognition(ctx, in, opts...)
}

// Recognize implements ocr.TextRecognitionAsyncServiceClient
func (c *TextRecognitionAsyncServiceClient) Recognize(ctx context.Context, in *ocr.RecognizeTextRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return ocr.NewTextRecognitionAsyncServiceClient(conn).Recognize(ctx, in, opts...)
}
