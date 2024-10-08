// Code generated by sdkgen. DO NOT EDIT.

// nolint
package cic

import (
	"context"

	"google.golang.org/grpc"

	cic "github.com/yandex-cloud/go-genproto/yandex/cloud/cic/v1"
)

//revive:disable

// PublicConnectionServiceClient is a cic.PublicConnectionServiceClient with
// lazy GRPC connection initialization.
type PublicConnectionServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// Get implements cic.PublicConnectionServiceClient
func (c *PublicConnectionServiceClient) Get(ctx context.Context, in *cic.GetPublicConnectionRequest, opts ...grpc.CallOption) (*cic.PublicConnection, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return cic.NewPublicConnectionServiceClient(conn).Get(ctx, in, opts...)
}

// List implements cic.PublicConnectionServiceClient
func (c *PublicConnectionServiceClient) List(ctx context.Context, in *cic.ListPublicConnectionsRequest, opts ...grpc.CallOption) (*cic.ListPublicConnectionsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return cic.NewPublicConnectionServiceClient(conn).List(ctx, in, opts...)
}

type PublicConnectionIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *PublicConnectionServiceClient
	request *cic.ListPublicConnectionsRequest

	items []*cic.PublicConnection
}

func (c *PublicConnectionServiceClient) PublicConnectionIterator(ctx context.Context, req *cic.ListPublicConnectionsRequest, opts ...grpc.CallOption) *PublicConnectionIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &PublicConnectionIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *PublicConnectionIterator) Next() bool {
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

	it.items = response.PublicConnections
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *PublicConnectionIterator) Take(size int64) ([]*cic.PublicConnection, error) {
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

	var result []*cic.PublicConnection

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *PublicConnectionIterator) TakeAll() ([]*cic.PublicConnection, error) {
	return it.Take(0)
}

func (it *PublicConnectionIterator) Value() *cic.PublicConnection {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *PublicConnectionIterator) Error() error {
	return it.err
}
