// Code generated by sdkgen. DO NOT EDIT.

// nolint
package clickhouse

import (
	"context"

	"google.golang.org/grpc"

	clickhouse "github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/clickhouse/v1"
)

//revive:disable

// ExtensionServiceClient is a clickhouse.ExtensionServiceClient with
// lazy GRPC connection initialization.
type ExtensionServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// Get implements clickhouse.ExtensionServiceClient
func (c *ExtensionServiceClient) Get(ctx context.Context, in *clickhouse.GetExtensionRequest, opts ...grpc.CallOption) (*clickhouse.Extension, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return clickhouse.NewExtensionServiceClient(conn).Get(ctx, in, opts...)
}

// List implements clickhouse.ExtensionServiceClient
func (c *ExtensionServiceClient) List(ctx context.Context, in *clickhouse.ListExtensionsRequest, opts ...grpc.CallOption) (*clickhouse.ListExtensionsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return clickhouse.NewExtensionServiceClient(conn).List(ctx, in, opts...)
}

type ExtensionIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *ExtensionServiceClient
	request *clickhouse.ListExtensionsRequest

	items []*clickhouse.Extension
}

func (c *ExtensionServiceClient) ExtensionIterator(ctx context.Context, req *clickhouse.ListExtensionsRequest, opts ...grpc.CallOption) *ExtensionIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &ExtensionIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *ExtensionIterator) Next() bool {
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

	it.items = response.Extensions
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *ExtensionIterator) Take(size int64) ([]*clickhouse.Extension, error) {
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

	var result []*clickhouse.Extension

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *ExtensionIterator) TakeAll() ([]*clickhouse.Extension, error) {
	return it.Take(0)
}

func (it *ExtensionIterator) Value() *clickhouse.Extension {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *ExtensionIterator) Error() error {
	return it.err
}
