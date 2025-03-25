// Code generated by sdkgen. DO NOT EDIT.

// nolint
package compute

import (
	"context"

	"google.golang.org/grpc"

	compute "github.com/yandex-cloud/go-genproto/yandex/cloud/compute/v1"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
)

//revive:disable

// ReservedInstancePoolServiceClient is a compute.ReservedInstancePoolServiceClient with
// lazy GRPC connection initialization.
type ReservedInstancePoolServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// Create implements compute.ReservedInstancePoolServiceClient
func (c *ReservedInstancePoolServiceClient) Create(ctx context.Context, in *compute.CreateReservedInstancePoolRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return compute.NewReservedInstancePoolServiceClient(conn).Create(ctx, in, opts...)
}

// Delete implements compute.ReservedInstancePoolServiceClient
func (c *ReservedInstancePoolServiceClient) Delete(ctx context.Context, in *compute.DeleteReservedInstancePoolRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return compute.NewReservedInstancePoolServiceClient(conn).Delete(ctx, in, opts...)
}

// Get implements compute.ReservedInstancePoolServiceClient
func (c *ReservedInstancePoolServiceClient) Get(ctx context.Context, in *compute.GetReservedInstancePoolRequest, opts ...grpc.CallOption) (*compute.ReservedInstancePool, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return compute.NewReservedInstancePoolServiceClient(conn).Get(ctx, in, opts...)
}

// List implements compute.ReservedInstancePoolServiceClient
func (c *ReservedInstancePoolServiceClient) List(ctx context.Context, in *compute.ListReservedInstancePoolsRequest, opts ...grpc.CallOption) (*compute.ListReservedInstancePoolsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return compute.NewReservedInstancePoolServiceClient(conn).List(ctx, in, opts...)
}

type ReservedInstancePoolIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *ReservedInstancePoolServiceClient
	request *compute.ListReservedInstancePoolsRequest

	items []*compute.ReservedInstancePool
}

func (c *ReservedInstancePoolServiceClient) ReservedInstancePoolIterator(ctx context.Context, req *compute.ListReservedInstancePoolsRequest, opts ...grpc.CallOption) *ReservedInstancePoolIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &ReservedInstancePoolIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *ReservedInstancePoolIterator) Next() bool {
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

	it.items = response.ReservedInstancePools
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *ReservedInstancePoolIterator) Take(size int64) ([]*compute.ReservedInstancePool, error) {
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

	var result []*compute.ReservedInstancePool

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *ReservedInstancePoolIterator) TakeAll() ([]*compute.ReservedInstancePool, error) {
	return it.Take(0)
}

func (it *ReservedInstancePoolIterator) Value() *compute.ReservedInstancePool {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *ReservedInstancePoolIterator) Error() error {
	return it.err
}

// Update implements compute.ReservedInstancePoolServiceClient
func (c *ReservedInstancePoolServiceClient) Update(ctx context.Context, in *compute.UpdateReservedInstancePoolRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return compute.NewReservedInstancePoolServiceClient(conn).Update(ctx, in, opts...)
}
