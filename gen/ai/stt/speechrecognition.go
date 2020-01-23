package stt

import (
	"context"

	recog "github.com/yandex-cloud/go-genproto/yandex/cloud/ai/stt/v2"
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	"google.golang.org/grpc"
)

// STT is a speechkit.STT
type STT struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// NewSTT creates instance of speech recognition
func NewSTT(g func(ctx context.Context) (*grpc.ClientConn, error)) *STT {
	return &STT{g}
}

// LongRunningRecognition -
func (t *STT) LongRunningRecognition(ctx context.Context, in *recog.LongRunningRecognitionRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := t.getConn(ctx)
	if err != nil {
		return nil, err
	}

	return recog.NewSttServiceClient(conn).LongRunningRecognize(ctx, in, opts...)
}

// StreamingRecognitionConfig implements speechkit.STT
func (t *STT) StreamingRecognitionConfig(ctx context.Context, cfg *recog.RecognitionConfig, opts ...grpc.CallOption) error {
	conn, err := t.getConn(ctx)
	if err != nil {
		return err
	}

	// Get cleint
	client, err := recog.NewSttServiceClient(conn).StreamingRecognize(ctx, opts...)
	if err != nil {
		return err
	}

	// Send
	return client.Send(&recog.StreamingRecognitionRequest{
		StreamingRequest: &recog.StreamingRecognitionRequest_Config{
			Config: cfg,
		},
	})
}

// StreamingRecognitionSend implements speechkit.StreamingRecognitionRequest
func (t *STT) StreamingRecognitionSend(ctx context.Context, in *recog.StreamingRecognitionRequest, opts ...grpc.CallOption) error {
	conn, err := t.getConn(ctx)
	if err != nil {
		return err
	}

	// Get cleint
	client, err := recog.NewSttServiceClient(conn).StreamingRecognize(ctx, opts...)
	if err != nil {
		return err
	}

	return client.Send(in)
}

// StreamingRecognitionReceive implements speechkit.StreamingRecognitionResponse
func (t *STT) StreamingRecognitionReceive(ctx context.Context, opts ...grpc.CallOption) (*recog.StreamingRecognitionResponse, error) {
	conn, err := t.getConn(ctx)
	if err != nil {
		return nil, err
	}

	// Get cleint
	client, err := recog.NewSttServiceClient(conn).StreamingRecognize(ctx, opts...)
	if err != nil {
		return nil, err
	}

	return client.Recv()
}
