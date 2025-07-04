// Code generated by sdkgen. DO NOT EDIT.

// nolint
package loadbalancer

import (
	"context"

	"google.golang.org/grpc"

	loadbalancer "github.com/yandex-cloud/go-genproto/yandex/cloud/loadbalancer/v1"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
)

//revive:disable

// NetworkLoadBalancerServiceClient is a loadbalancer.NetworkLoadBalancerServiceClient with
// lazy GRPC connection initialization.
type NetworkLoadBalancerServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// AddListener implements loadbalancer.NetworkLoadBalancerServiceClient
func (c *NetworkLoadBalancerServiceClient) AddListener(ctx context.Context, in *loadbalancer.AddNetworkLoadBalancerListenerRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return loadbalancer.NewNetworkLoadBalancerServiceClient(conn).AddListener(ctx, in, opts...)
}

// AttachTargetGroup implements loadbalancer.NetworkLoadBalancerServiceClient
func (c *NetworkLoadBalancerServiceClient) AttachTargetGroup(ctx context.Context, in *loadbalancer.AttachNetworkLoadBalancerTargetGroupRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return loadbalancer.NewNetworkLoadBalancerServiceClient(conn).AttachTargetGroup(ctx, in, opts...)
}

// CancelZonalShift implements loadbalancer.NetworkLoadBalancerServiceClient
func (c *NetworkLoadBalancerServiceClient) CancelZonalShift(ctx context.Context, in *loadbalancer.CancelZonalShiftRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return loadbalancer.NewNetworkLoadBalancerServiceClient(conn).CancelZonalShift(ctx, in, opts...)
}

// Create implements loadbalancer.NetworkLoadBalancerServiceClient
func (c *NetworkLoadBalancerServiceClient) Create(ctx context.Context, in *loadbalancer.CreateNetworkLoadBalancerRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return loadbalancer.NewNetworkLoadBalancerServiceClient(conn).Create(ctx, in, opts...)
}

// Delete implements loadbalancer.NetworkLoadBalancerServiceClient
func (c *NetworkLoadBalancerServiceClient) Delete(ctx context.Context, in *loadbalancer.DeleteNetworkLoadBalancerRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return loadbalancer.NewNetworkLoadBalancerServiceClient(conn).Delete(ctx, in, opts...)
}

// DetachTargetGroup implements loadbalancer.NetworkLoadBalancerServiceClient
func (c *NetworkLoadBalancerServiceClient) DetachTargetGroup(ctx context.Context, in *loadbalancer.DetachNetworkLoadBalancerTargetGroupRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return loadbalancer.NewNetworkLoadBalancerServiceClient(conn).DetachTargetGroup(ctx, in, opts...)
}

// Get implements loadbalancer.NetworkLoadBalancerServiceClient
func (c *NetworkLoadBalancerServiceClient) Get(ctx context.Context, in *loadbalancer.GetNetworkLoadBalancerRequest, opts ...grpc.CallOption) (*loadbalancer.NetworkLoadBalancer, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return loadbalancer.NewNetworkLoadBalancerServiceClient(conn).Get(ctx, in, opts...)
}

// GetTargetStates implements loadbalancer.NetworkLoadBalancerServiceClient
func (c *NetworkLoadBalancerServiceClient) GetTargetStates(ctx context.Context, in *loadbalancer.GetTargetStatesRequest, opts ...grpc.CallOption) (*loadbalancer.GetTargetStatesResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return loadbalancer.NewNetworkLoadBalancerServiceClient(conn).GetTargetStates(ctx, in, opts...)
}

// List implements loadbalancer.NetworkLoadBalancerServiceClient
func (c *NetworkLoadBalancerServiceClient) List(ctx context.Context, in *loadbalancer.ListNetworkLoadBalancersRequest, opts ...grpc.CallOption) (*loadbalancer.ListNetworkLoadBalancersResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return loadbalancer.NewNetworkLoadBalancerServiceClient(conn).List(ctx, in, opts...)
}

type NetworkLoadBalancerIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *NetworkLoadBalancerServiceClient
	request *loadbalancer.ListNetworkLoadBalancersRequest

	items []*loadbalancer.NetworkLoadBalancer
}

func (c *NetworkLoadBalancerServiceClient) NetworkLoadBalancerIterator(ctx context.Context, req *loadbalancer.ListNetworkLoadBalancersRequest, opts ...grpc.CallOption) *NetworkLoadBalancerIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &NetworkLoadBalancerIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *NetworkLoadBalancerIterator) Next() bool {
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

	it.items = response.NetworkLoadBalancers
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *NetworkLoadBalancerIterator) Take(size int64) ([]*loadbalancer.NetworkLoadBalancer, error) {
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

	var result []*loadbalancer.NetworkLoadBalancer

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *NetworkLoadBalancerIterator) TakeAll() ([]*loadbalancer.NetworkLoadBalancer, error) {
	return it.Take(0)
}

func (it *NetworkLoadBalancerIterator) Value() *loadbalancer.NetworkLoadBalancer {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *NetworkLoadBalancerIterator) Error() error {
	return it.err
}

// ListOperations implements loadbalancer.NetworkLoadBalancerServiceClient
func (c *NetworkLoadBalancerServiceClient) ListOperations(ctx context.Context, in *loadbalancer.ListNetworkLoadBalancerOperationsRequest, opts ...grpc.CallOption) (*loadbalancer.ListNetworkLoadBalancerOperationsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return loadbalancer.NewNetworkLoadBalancerServiceClient(conn).ListOperations(ctx, in, opts...)
}

type NetworkLoadBalancerOperationsIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *NetworkLoadBalancerServiceClient
	request *loadbalancer.ListNetworkLoadBalancerOperationsRequest

	items []*operation.Operation
}

func (c *NetworkLoadBalancerServiceClient) NetworkLoadBalancerOperationsIterator(ctx context.Context, req *loadbalancer.ListNetworkLoadBalancerOperationsRequest, opts ...grpc.CallOption) *NetworkLoadBalancerOperationsIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &NetworkLoadBalancerOperationsIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *NetworkLoadBalancerOperationsIterator) Next() bool {
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

	response, err := it.client.ListOperations(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.Operations
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *NetworkLoadBalancerOperationsIterator) Take(size int64) ([]*operation.Operation, error) {
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

	var result []*operation.Operation

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *NetworkLoadBalancerOperationsIterator) TakeAll() ([]*operation.Operation, error) {
	return it.Take(0)
}

func (it *NetworkLoadBalancerOperationsIterator) Value() *operation.Operation {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *NetworkLoadBalancerOperationsIterator) Error() error {
	return it.err
}

// RemoveListener implements loadbalancer.NetworkLoadBalancerServiceClient
func (c *NetworkLoadBalancerServiceClient) RemoveListener(ctx context.Context, in *loadbalancer.RemoveNetworkLoadBalancerListenerRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return loadbalancer.NewNetworkLoadBalancerServiceClient(conn).RemoveListener(ctx, in, opts...)
}

// Start implements loadbalancer.NetworkLoadBalancerServiceClient
func (c *NetworkLoadBalancerServiceClient) Start(ctx context.Context, in *loadbalancer.StartNetworkLoadBalancerRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return loadbalancer.NewNetworkLoadBalancerServiceClient(conn).Start(ctx, in, opts...)
}

// StartZonalShift implements loadbalancer.NetworkLoadBalancerServiceClient
func (c *NetworkLoadBalancerServiceClient) StartZonalShift(ctx context.Context, in *loadbalancer.StartZonalShiftRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return loadbalancer.NewNetworkLoadBalancerServiceClient(conn).StartZonalShift(ctx, in, opts...)
}

// Stop implements loadbalancer.NetworkLoadBalancerServiceClient
func (c *NetworkLoadBalancerServiceClient) Stop(ctx context.Context, in *loadbalancer.StopNetworkLoadBalancerRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return loadbalancer.NewNetworkLoadBalancerServiceClient(conn).Stop(ctx, in, opts...)
}

// Update implements loadbalancer.NetworkLoadBalancerServiceClient
func (c *NetworkLoadBalancerServiceClient) Update(ctx context.Context, in *loadbalancer.UpdateNetworkLoadBalancerRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return loadbalancer.NewNetworkLoadBalancerServiceClient(conn).Update(ctx, in, opts...)
}
