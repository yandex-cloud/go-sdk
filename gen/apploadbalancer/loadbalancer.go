// Code generated by sdkgen. DO NOT EDIT.

// nolint
package apploadbalancer

import (
	"context"

	"google.golang.org/grpc"

	apploadbalancer "github.com/yandex-cloud/go-genproto/yandex/cloud/apploadbalancer/v1"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
)

//revive:disable

// LoadBalancerServiceClient is a apploadbalancer.LoadBalancerServiceClient with
// lazy GRPC connection initialization.
type LoadBalancerServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// AddListener implements apploadbalancer.LoadBalancerServiceClient
func (c *LoadBalancerServiceClient) AddListener(ctx context.Context, in *apploadbalancer.AddListenerRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return apploadbalancer.NewLoadBalancerServiceClient(conn).AddListener(ctx, in, opts...)
}

// AddSniMatch implements apploadbalancer.LoadBalancerServiceClient
func (c *LoadBalancerServiceClient) AddSniMatch(ctx context.Context, in *apploadbalancer.AddSniMatchRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return apploadbalancer.NewLoadBalancerServiceClient(conn).AddSniMatch(ctx, in, opts...)
}

// CancelZonalShift implements apploadbalancer.LoadBalancerServiceClient
func (c *LoadBalancerServiceClient) CancelZonalShift(ctx context.Context, in *apploadbalancer.CancelZonalShiftRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return apploadbalancer.NewLoadBalancerServiceClient(conn).CancelZonalShift(ctx, in, opts...)
}

// Create implements apploadbalancer.LoadBalancerServiceClient
func (c *LoadBalancerServiceClient) Create(ctx context.Context, in *apploadbalancer.CreateLoadBalancerRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return apploadbalancer.NewLoadBalancerServiceClient(conn).Create(ctx, in, opts...)
}

// Delete implements apploadbalancer.LoadBalancerServiceClient
func (c *LoadBalancerServiceClient) Delete(ctx context.Context, in *apploadbalancer.DeleteLoadBalancerRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return apploadbalancer.NewLoadBalancerServiceClient(conn).Delete(ctx, in, opts...)
}

// Get implements apploadbalancer.LoadBalancerServiceClient
func (c *LoadBalancerServiceClient) Get(ctx context.Context, in *apploadbalancer.GetLoadBalancerRequest, opts ...grpc.CallOption) (*apploadbalancer.LoadBalancer, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return apploadbalancer.NewLoadBalancerServiceClient(conn).Get(ctx, in, opts...)
}

// GetTargetStates implements apploadbalancer.LoadBalancerServiceClient
func (c *LoadBalancerServiceClient) GetTargetStates(ctx context.Context, in *apploadbalancer.GetTargetStatesRequest, opts ...grpc.CallOption) (*apploadbalancer.GetTargetStatesResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return apploadbalancer.NewLoadBalancerServiceClient(conn).GetTargetStates(ctx, in, opts...)
}

// List implements apploadbalancer.LoadBalancerServiceClient
func (c *LoadBalancerServiceClient) List(ctx context.Context, in *apploadbalancer.ListLoadBalancersRequest, opts ...grpc.CallOption) (*apploadbalancer.ListLoadBalancersResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return apploadbalancer.NewLoadBalancerServiceClient(conn).List(ctx, in, opts...)
}

type LoadBalancerIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *LoadBalancerServiceClient
	request *apploadbalancer.ListLoadBalancersRequest

	items []*apploadbalancer.LoadBalancer
}

func (c *LoadBalancerServiceClient) LoadBalancerIterator(ctx context.Context, req *apploadbalancer.ListLoadBalancersRequest, opts ...grpc.CallOption) *LoadBalancerIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &LoadBalancerIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *LoadBalancerIterator) Next() bool {
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

	it.items = response.LoadBalancers
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *LoadBalancerIterator) Take(size int64) ([]*apploadbalancer.LoadBalancer, error) {
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

	var result []*apploadbalancer.LoadBalancer

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *LoadBalancerIterator) TakeAll() ([]*apploadbalancer.LoadBalancer, error) {
	return it.Take(0)
}

func (it *LoadBalancerIterator) Value() *apploadbalancer.LoadBalancer {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *LoadBalancerIterator) Error() error {
	return it.err
}

// ListOperations implements apploadbalancer.LoadBalancerServiceClient
func (c *LoadBalancerServiceClient) ListOperations(ctx context.Context, in *apploadbalancer.ListLoadBalancerOperationsRequest, opts ...grpc.CallOption) (*apploadbalancer.ListLoadBalancerOperationsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return apploadbalancer.NewLoadBalancerServiceClient(conn).ListOperations(ctx, in, opts...)
}

type LoadBalancerOperationsIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *LoadBalancerServiceClient
	request *apploadbalancer.ListLoadBalancerOperationsRequest

	items []*operation.Operation
}

func (c *LoadBalancerServiceClient) LoadBalancerOperationsIterator(ctx context.Context, req *apploadbalancer.ListLoadBalancerOperationsRequest, opts ...grpc.CallOption) *LoadBalancerOperationsIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &LoadBalancerOperationsIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *LoadBalancerOperationsIterator) Next() bool {
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

func (it *LoadBalancerOperationsIterator) Take(size int64) ([]*operation.Operation, error) {
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

func (it *LoadBalancerOperationsIterator) TakeAll() ([]*operation.Operation, error) {
	return it.Take(0)
}

func (it *LoadBalancerOperationsIterator) Value() *operation.Operation {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *LoadBalancerOperationsIterator) Error() error {
	return it.err
}

// RemoveListener implements apploadbalancer.LoadBalancerServiceClient
func (c *LoadBalancerServiceClient) RemoveListener(ctx context.Context, in *apploadbalancer.RemoveListenerRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return apploadbalancer.NewLoadBalancerServiceClient(conn).RemoveListener(ctx, in, opts...)
}

// RemoveSniMatch implements apploadbalancer.LoadBalancerServiceClient
func (c *LoadBalancerServiceClient) RemoveSniMatch(ctx context.Context, in *apploadbalancer.RemoveSniMatchRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return apploadbalancer.NewLoadBalancerServiceClient(conn).RemoveSniMatch(ctx, in, opts...)
}

// Start implements apploadbalancer.LoadBalancerServiceClient
func (c *LoadBalancerServiceClient) Start(ctx context.Context, in *apploadbalancer.StartLoadBalancerRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return apploadbalancer.NewLoadBalancerServiceClient(conn).Start(ctx, in, opts...)
}

// StartZonalShift implements apploadbalancer.LoadBalancerServiceClient
func (c *LoadBalancerServiceClient) StartZonalShift(ctx context.Context, in *apploadbalancer.StartZonalShiftRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return apploadbalancer.NewLoadBalancerServiceClient(conn).StartZonalShift(ctx, in, opts...)
}

// Stop implements apploadbalancer.LoadBalancerServiceClient
func (c *LoadBalancerServiceClient) Stop(ctx context.Context, in *apploadbalancer.StopLoadBalancerRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return apploadbalancer.NewLoadBalancerServiceClient(conn).Stop(ctx, in, opts...)
}

// Update implements apploadbalancer.LoadBalancerServiceClient
func (c *LoadBalancerServiceClient) Update(ctx context.Context, in *apploadbalancer.UpdateLoadBalancerRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return apploadbalancer.NewLoadBalancerServiceClient(conn).Update(ctx, in, opts...)
}

// UpdateListener implements apploadbalancer.LoadBalancerServiceClient
func (c *LoadBalancerServiceClient) UpdateListener(ctx context.Context, in *apploadbalancer.UpdateListenerRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return apploadbalancer.NewLoadBalancerServiceClient(conn).UpdateListener(ctx, in, opts...)
}

// UpdateSniMatch implements apploadbalancer.LoadBalancerServiceClient
func (c *LoadBalancerServiceClient) UpdateSniMatch(ctx context.Context, in *apploadbalancer.UpdateSniMatchRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return apploadbalancer.NewLoadBalancerServiceClient(conn).UpdateSniMatch(ctx, in, opts...)
}
