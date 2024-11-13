// Code generated by sdkgen. DO NOT EDIT.

// nolint
package stt

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	stt "github.com/yandex-cloud/go-genproto/yandex/cloud/ai/stt/v3"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
)

//revive:disable

// AsyncRecognizerClient is a stt.AsyncRecognizerClient with
// lazy GRPC connection initialization.
type AsyncRecognizerClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// DeleteRecognition implements stt.AsyncRecognizerClient
func (c *AsyncRecognizerClient) DeleteRecognition(ctx context.Context, in *stt.DeleteRecognitionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return stt.NewAsyncRecognizerClient(conn).DeleteRecognition(ctx, in, opts...)
}

// GetRecognition implements stt.AsyncRecognizerClient
func (c *AsyncRecognizerClient) GetRecognition(ctx context.Context, in *stt.GetRecognitionRequest, opts ...grpc.CallOption) (stt.AsyncRecognizer_GetRecognitionClient, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return stt.NewAsyncRecognizerClient(conn).GetRecognition(ctx, in, opts...)
}

// RecognizeFile implements stt.AsyncRecognizerClient
func (c *AsyncRecognizerClient) RecognizeFile(ctx context.Context, in *stt.RecognizeFileRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return stt.NewAsyncRecognizerClient(conn).RecognizeFile(ctx, in, opts...)
}
