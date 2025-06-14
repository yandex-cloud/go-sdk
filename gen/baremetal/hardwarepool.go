// Code generated by sdkgen. DO NOT EDIT.

// nolint
package baremetal

import (
	"context"

	"google.golang.org/grpc"

	baremetal "github.com/yandex-cloud/go-genproto/yandex/cloud/baremetal/v1alpha"
)

//revive:disable

// HardwarePoolServiceClient is a baremetal.HardwarePoolServiceClient with
// lazy GRPC connection initialization.
type HardwarePoolServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// Get implements baremetal.HardwarePoolServiceClient
func (c *HardwarePoolServiceClient) Get(ctx context.Context, in *baremetal.GetHardwarePoolRequest, opts ...grpc.CallOption) (*baremetal.HardwarePool, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return baremetal.NewHardwarePoolServiceClient(conn).Get(ctx, in, opts...)
}

// List implements baremetal.HardwarePoolServiceClient
func (c *HardwarePoolServiceClient) List(ctx context.Context, in *baremetal.ListHardwarePoolsRequest, opts ...grpc.CallOption) (*baremetal.ListHardwarePoolsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return baremetal.NewHardwarePoolServiceClient(conn).List(ctx, in, opts...)
}

type HardwarePoolIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *HardwarePoolServiceClient
	request *baremetal.ListHardwarePoolsRequest

	items []*baremetal.HardwarePool
}

func (c *HardwarePoolServiceClient) HardwarePoolIterator(ctx context.Context, req *baremetal.ListHardwarePoolsRequest, opts ...grpc.CallOption) *HardwarePoolIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &HardwarePoolIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *HardwarePoolIterator) Next() bool {
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

	it.items = response.HardwarePools
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *HardwarePoolIterator) Take(size int64) ([]*baremetal.HardwarePool, error) {
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

	var result []*baremetal.HardwarePool

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *HardwarePoolIterator) TakeAll() ([]*baremetal.HardwarePool, error) {
	return it.Take(0)
}

func (it *HardwarePoolIterator) Value() *baremetal.HardwarePool {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *HardwarePoolIterator) Error() error {
	return it.err
}
