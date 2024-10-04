// Code generated by sdkgen. DO NOT EDIT.

// nolint
package cic

import (
	"context"

	"google.golang.org/grpc"

	cic "github.com/yandex-cloud/go-genproto/yandex/cloud/cic/v1"
)

//revive:disable

// PointOfPresenceServiceClient is a cic.PointOfPresenceServiceClient with
// lazy GRPC connection initialization.
type PointOfPresenceServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// Get implements cic.PointOfPresenceServiceClient
func (c *PointOfPresenceServiceClient) Get(ctx context.Context, in *cic.GetPointOfPresenceRequest, opts ...grpc.CallOption) (*cic.PointOfPresence, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return cic.NewPointOfPresenceServiceClient(conn).Get(ctx, in, opts...)
}

// List implements cic.PointOfPresenceServiceClient
func (c *PointOfPresenceServiceClient) List(ctx context.Context, in *cic.ListPointOfPresencesRequest, opts ...grpc.CallOption) (*cic.ListPointOfPresencesResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return cic.NewPointOfPresenceServiceClient(conn).List(ctx, in, opts...)
}

type PointOfPresenceIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *PointOfPresenceServiceClient
	request *cic.ListPointOfPresencesRequest

	items []*cic.PointOfPresence
}

func (c *PointOfPresenceServiceClient) PointOfPresenceIterator(ctx context.Context, req *cic.ListPointOfPresencesRequest, opts ...grpc.CallOption) *PointOfPresenceIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &PointOfPresenceIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *PointOfPresenceIterator) Next() bool {
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

	it.items = response.PointOfPresences
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *PointOfPresenceIterator) Take(size int64) ([]*cic.PointOfPresence, error) {
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

	var result []*cic.PointOfPresence

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *PointOfPresenceIterator) TakeAll() ([]*cic.PointOfPresence, error) {
	return it.Take(0)
}

func (it *PointOfPresenceIterator) Value() *cic.PointOfPresence {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *PointOfPresenceIterator) Error() error {
	return it.err
}
