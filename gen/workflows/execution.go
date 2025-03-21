// Code generated by sdkgen. DO NOT EDIT.

// nolint
package workflows

import (
	"context"

	"google.golang.org/grpc"

	workflows "github.com/yandex-cloud/go-genproto/yandex/cloud/serverless/workflows/v1"
)

//revive:disable

// ExecutionServiceClient is a workflows.ExecutionServiceClient with
// lazy GRPC connection initialization.
type ExecutionServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// Get implements workflows.ExecutionServiceClient
func (c *ExecutionServiceClient) Get(ctx context.Context, in *workflows.GetExecutionRequest, opts ...grpc.CallOption) (*workflows.GetExecutionResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return workflows.NewExecutionServiceClient(conn).Get(ctx, in, opts...)
}

// GetHistory implements workflows.ExecutionServiceClient
func (c *ExecutionServiceClient) GetHistory(ctx context.Context, in *workflows.GetExecutionHistoryRequest, opts ...grpc.CallOption) (*workflows.GetExecutionHistoryResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return workflows.NewExecutionServiceClient(conn).GetHistory(ctx, in, opts...)
}

// List implements workflows.ExecutionServiceClient
func (c *ExecutionServiceClient) List(ctx context.Context, in *workflows.ListExecutionsRequest, opts ...grpc.CallOption) (*workflows.ListExecutionsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return workflows.NewExecutionServiceClient(conn).List(ctx, in, opts...)
}

type ExecutionIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *ExecutionServiceClient
	request *workflows.ListExecutionsRequest

	items []*workflows.ExecutionPreview
}

func (c *ExecutionServiceClient) ExecutionIterator(ctx context.Context, req *workflows.ListExecutionsRequest, opts ...grpc.CallOption) *ExecutionIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &ExecutionIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *ExecutionIterator) Next() bool {
	if it.err != nil {
		return false
	}
	if len(it.items) > 1 {
		it.items[0] = nil
		it.items = it.items[1:]
		return true
	}
	it.items = nil // consume last item, if any

	if it.started && it.request.PageToken == "" {
		return false
	}
	it.started = true

	if it.requestedSize == 0 || it.requestedSize > it.pageSize {
		it.request.PageSize = it.pageSize
	} else {
		it.request.PageSize = it.requestedSize
	}

	response, err := it.client.List(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.Executions
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *ExecutionIterator) Take(size int64) ([]*workflows.ExecutionPreview, error) {
	if it.err != nil {
		return nil, it.err
	}

	if size == 0 {
		size = 1 << 32 // something insanely large
	}
	it.requestedSize = size
	defer func() {
		// reset iterator for future calls.
		it.requestedSize = 0
	}()

	var result []*workflows.ExecutionPreview

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *ExecutionIterator) TakeAll() ([]*workflows.ExecutionPreview, error) {
	return it.Take(0)
}

func (it *ExecutionIterator) Value() *workflows.ExecutionPreview {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *ExecutionIterator) Error() error {
	return it.err
}

// Start implements workflows.ExecutionServiceClient
func (c *ExecutionServiceClient) Start(ctx context.Context, in *workflows.StartExecutionRequest, opts ...grpc.CallOption) (*workflows.StartExecutionResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return workflows.NewExecutionServiceClient(conn).Start(ctx, in, opts...)
}

// Stop implements workflows.ExecutionServiceClient
func (c *ExecutionServiceClient) Stop(ctx context.Context, in *workflows.StopExecutionRequest, opts ...grpc.CallOption) (*workflows.StopExecutionResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return workflows.NewExecutionServiceClient(conn).Stop(ctx, in, opts...)
}

// Terminate implements workflows.ExecutionServiceClient
func (c *ExecutionServiceClient) Terminate(ctx context.Context, in *workflows.TerminateExecutionRequest, opts ...grpc.CallOption) (*workflows.TerminateExecutionResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return workflows.NewExecutionServiceClient(conn).Terminate(ctx, in, opts...)
}
