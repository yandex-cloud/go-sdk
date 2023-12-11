// Code generated by sdkgen. DO NOT EDIT.

// nolint
package compute

import (
	"context"

	"google.golang.org/grpc"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/access"
	compute "github.com/yandex-cloud/go-genproto/yandex/cloud/compute/v1"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
)

//revive:disable

// InstanceServiceClient is a compute.InstanceServiceClient with
// lazy GRPC connection initialization.
type InstanceServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// AddOneToOneNat implements compute.InstanceServiceClient
func (c *InstanceServiceClient) AddOneToOneNat(ctx context.Context, in *compute.AddInstanceOneToOneNatRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return compute.NewInstanceServiceClient(conn).AddOneToOneNat(ctx, in, opts...)
}

// AttachDisk implements compute.InstanceServiceClient
func (c *InstanceServiceClient) AttachDisk(ctx context.Context, in *compute.AttachInstanceDiskRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return compute.NewInstanceServiceClient(conn).AttachDisk(ctx, in, opts...)
}

// AttachFilesystem implements compute.InstanceServiceClient
func (c *InstanceServiceClient) AttachFilesystem(ctx context.Context, in *compute.AttachInstanceFilesystemRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return compute.NewInstanceServiceClient(conn).AttachFilesystem(ctx, in, opts...)
}

// Create implements compute.InstanceServiceClient
func (c *InstanceServiceClient) Create(ctx context.Context, in *compute.CreateInstanceRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return compute.NewInstanceServiceClient(conn).Create(ctx, in, opts...)
}

// Delete implements compute.InstanceServiceClient
func (c *InstanceServiceClient) Delete(ctx context.Context, in *compute.DeleteInstanceRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return compute.NewInstanceServiceClient(conn).Delete(ctx, in, opts...)
}

// DetachDisk implements compute.InstanceServiceClient
func (c *InstanceServiceClient) DetachDisk(ctx context.Context, in *compute.DetachInstanceDiskRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return compute.NewInstanceServiceClient(conn).DetachDisk(ctx, in, opts...)
}

// DetachFilesystem implements compute.InstanceServiceClient
func (c *InstanceServiceClient) DetachFilesystem(ctx context.Context, in *compute.DetachInstanceFilesystemRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return compute.NewInstanceServiceClient(conn).DetachFilesystem(ctx, in, opts...)
}

// Get implements compute.InstanceServiceClient
func (c *InstanceServiceClient) Get(ctx context.Context, in *compute.GetInstanceRequest, opts ...grpc.CallOption) (*compute.Instance, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return compute.NewInstanceServiceClient(conn).Get(ctx, in, opts...)
}

// GetSerialPortOutput implements compute.InstanceServiceClient
func (c *InstanceServiceClient) GetSerialPortOutput(ctx context.Context, in *compute.GetInstanceSerialPortOutputRequest, opts ...grpc.CallOption) (*compute.GetInstanceSerialPortOutputResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return compute.NewInstanceServiceClient(conn).GetSerialPortOutput(ctx, in, opts...)
}

// List implements compute.InstanceServiceClient
func (c *InstanceServiceClient) List(ctx context.Context, in *compute.ListInstancesRequest, opts ...grpc.CallOption) (*compute.ListInstancesResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return compute.NewInstanceServiceClient(conn).List(ctx, in, opts...)
}

type InstanceIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *InstanceServiceClient
	request *compute.ListInstancesRequest

	items []*compute.Instance
}

func (c *InstanceServiceClient) InstanceIterator(ctx context.Context, req *compute.ListInstancesRequest, opts ...grpc.CallOption) *InstanceIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &InstanceIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *InstanceIterator) Next() bool {
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

	it.items = response.Instances
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *InstanceIterator) Take(size int64) ([]*compute.Instance, error) {
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

	var result []*compute.Instance

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *InstanceIterator) TakeAll() ([]*compute.Instance, error) {
	return it.Take(0)
}

func (it *InstanceIterator) Value() *compute.Instance {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *InstanceIterator) Error() error {
	return it.err
}

// ListAccessBindings implements compute.InstanceServiceClient
func (c *InstanceServiceClient) ListAccessBindings(ctx context.Context, in *access.ListAccessBindingsRequest, opts ...grpc.CallOption) (*access.ListAccessBindingsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return compute.NewInstanceServiceClient(conn).ListAccessBindings(ctx, in, opts...)
}

type InstanceAccessBindingsIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *InstanceServiceClient
	request *access.ListAccessBindingsRequest

	items []*access.AccessBinding
}

func (c *InstanceServiceClient) InstanceAccessBindingsIterator(ctx context.Context, req *access.ListAccessBindingsRequest, opts ...grpc.CallOption) *InstanceAccessBindingsIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &InstanceAccessBindingsIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *InstanceAccessBindingsIterator) Next() bool {
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

	response, err := it.client.ListAccessBindings(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.AccessBindings
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *InstanceAccessBindingsIterator) Take(size int64) ([]*access.AccessBinding, error) {
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

	var result []*access.AccessBinding

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *InstanceAccessBindingsIterator) TakeAll() ([]*access.AccessBinding, error) {
	return it.Take(0)
}

func (it *InstanceAccessBindingsIterator) Value() *access.AccessBinding {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *InstanceAccessBindingsIterator) Error() error {
	return it.err
}

// ListOperations implements compute.InstanceServiceClient
func (c *InstanceServiceClient) ListOperations(ctx context.Context, in *compute.ListInstanceOperationsRequest, opts ...grpc.CallOption) (*compute.ListInstanceOperationsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return compute.NewInstanceServiceClient(conn).ListOperations(ctx, in, opts...)
}

type InstanceOperationsIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *InstanceServiceClient
	request *compute.ListInstanceOperationsRequest

	items []*operation.Operation
}

func (c *InstanceServiceClient) InstanceOperationsIterator(ctx context.Context, req *compute.ListInstanceOperationsRequest, opts ...grpc.CallOption) *InstanceOperationsIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &InstanceOperationsIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *InstanceOperationsIterator) Next() bool {
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

func (it *InstanceOperationsIterator) Take(size int64) ([]*operation.Operation, error) {
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

func (it *InstanceOperationsIterator) TakeAll() ([]*operation.Operation, error) {
	return it.Take(0)
}

func (it *InstanceOperationsIterator) Value() *operation.Operation {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *InstanceOperationsIterator) Error() error {
	return it.err
}

// Move implements compute.InstanceServiceClient
func (c *InstanceServiceClient) Move(ctx context.Context, in *compute.MoveInstanceRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return compute.NewInstanceServiceClient(conn).Move(ctx, in, opts...)
}

// Relocate implements compute.InstanceServiceClient
func (c *InstanceServiceClient) Relocate(ctx context.Context, in *compute.RelocateInstanceRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return compute.NewInstanceServiceClient(conn).Relocate(ctx, in, opts...)
}

// RemoveOneToOneNat implements compute.InstanceServiceClient
func (c *InstanceServiceClient) RemoveOneToOneNat(ctx context.Context, in *compute.RemoveInstanceOneToOneNatRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return compute.NewInstanceServiceClient(conn).RemoveOneToOneNat(ctx, in, opts...)
}

// Restart implements compute.InstanceServiceClient
func (c *InstanceServiceClient) Restart(ctx context.Context, in *compute.RestartInstanceRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return compute.NewInstanceServiceClient(conn).Restart(ctx, in, opts...)
}

// SetAccessBindings implements compute.InstanceServiceClient
func (c *InstanceServiceClient) SetAccessBindings(ctx context.Context, in *access.SetAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return compute.NewInstanceServiceClient(conn).SetAccessBindings(ctx, in, opts...)
}

// SimulateMaintenanceEvent implements compute.InstanceServiceClient
func (c *InstanceServiceClient) SimulateMaintenanceEvent(ctx context.Context, in *compute.SimulateInstanceMaintenanceEventRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return compute.NewInstanceServiceClient(conn).SimulateMaintenanceEvent(ctx, in, opts...)
}

// Start implements compute.InstanceServiceClient
func (c *InstanceServiceClient) Start(ctx context.Context, in *compute.StartInstanceRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return compute.NewInstanceServiceClient(conn).Start(ctx, in, opts...)
}

// Stop implements compute.InstanceServiceClient
func (c *InstanceServiceClient) Stop(ctx context.Context, in *compute.StopInstanceRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return compute.NewInstanceServiceClient(conn).Stop(ctx, in, opts...)
}

// Update implements compute.InstanceServiceClient
func (c *InstanceServiceClient) Update(ctx context.Context, in *compute.UpdateInstanceRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return compute.NewInstanceServiceClient(conn).Update(ctx, in, opts...)
}

// UpdateAccessBindings implements compute.InstanceServiceClient
func (c *InstanceServiceClient) UpdateAccessBindings(ctx context.Context, in *access.UpdateAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return compute.NewInstanceServiceClient(conn).UpdateAccessBindings(ctx, in, opts...)
}

// UpdateMetadata implements compute.InstanceServiceClient
func (c *InstanceServiceClient) UpdateMetadata(ctx context.Context, in *compute.UpdateInstanceMetadataRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return compute.NewInstanceServiceClient(conn).UpdateMetadata(ctx, in, opts...)
}

// UpdateNetworkInterface implements compute.InstanceServiceClient
func (c *InstanceServiceClient) UpdateNetworkInterface(ctx context.Context, in *compute.UpdateInstanceNetworkInterfaceRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return compute.NewInstanceServiceClient(conn).UpdateNetworkInterface(ctx, in, opts...)
}
