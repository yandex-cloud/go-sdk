// Code generated by sdkgen. DO NOT EDIT.

// nolint
package greenplum

import (
	"context"

	"google.golang.org/grpc"

	greenplum "github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/greenplum/v1"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
)

//revive:disable

// HBARuleServiceClient is a greenplum.HBARuleServiceClient with
// lazy GRPC connection initialization.
type HBARuleServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// BatchUpdate implements greenplum.HBARuleServiceClient
func (c *HBARuleServiceClient) BatchUpdate(ctx context.Context, in *greenplum.BatchUpdateHBARulesRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return greenplum.NewHBARuleServiceClient(conn).BatchUpdate(ctx, in, opts...)
}

// Create implements greenplum.HBARuleServiceClient
func (c *HBARuleServiceClient) Create(ctx context.Context, in *greenplum.CreateHBARuleRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return greenplum.NewHBARuleServiceClient(conn).Create(ctx, in, opts...)
}

// Delete implements greenplum.HBARuleServiceClient
func (c *HBARuleServiceClient) Delete(ctx context.Context, in *greenplum.DeleteHBARuleRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return greenplum.NewHBARuleServiceClient(conn).Delete(ctx, in, opts...)
}

// List implements greenplum.HBARuleServiceClient
func (c *HBARuleServiceClient) List(ctx context.Context, in *greenplum.ListHBARulesRequest, opts ...grpc.CallOption) (*greenplum.ListHBARulesResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return greenplum.NewHBARuleServiceClient(conn).List(ctx, in, opts...)
}

type HBARuleIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *HBARuleServiceClient
	request *greenplum.ListHBARulesRequest

	items []*greenplum.HBARule
}

func (c *HBARuleServiceClient) HBARuleIterator(ctx context.Context, req *greenplum.ListHBARulesRequest, opts ...grpc.CallOption) *HBARuleIterator {
	var pageSize int64
	const defaultPageSize = 1000

	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &HBARuleIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *HBARuleIterator) Next() bool {
	if it.err != nil {
		return false
	}
	if len(it.items) > 1 {
		it.items[0] = nil
		it.items = it.items[1:]
		return true
	}
	it.items = nil // consume last item, if any

	if it.started {
		return false
	}
	it.started = true

	response, err := it.client.List(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.HbaRules
	return len(it.items) > 0
}

func (it *HBARuleIterator) Take(size int64) ([]*greenplum.HBARule, error) {
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

	var result []*greenplum.HBARule

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *HBARuleIterator) TakeAll() ([]*greenplum.HBARule, error) {
	return it.Take(0)
}

func (it *HBARuleIterator) Value() *greenplum.HBARule {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *HBARuleIterator) Error() error {
	return it.err
}

// ListAtRevision implements greenplum.HBARuleServiceClient
func (c *HBARuleServiceClient) ListAtRevision(ctx context.Context, in *greenplum.ListHBARulesAtRevisionRequest, opts ...grpc.CallOption) (*greenplum.ListHBARulesResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return greenplum.NewHBARuleServiceClient(conn).ListAtRevision(ctx, in, opts...)
}

type HBARuleAtRevisionIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *HBARuleServiceClient
	request *greenplum.ListHBARulesAtRevisionRequest

	items []*greenplum.HBARule
}

func (c *HBARuleServiceClient) HBARuleAtRevisionIterator(ctx context.Context, req *greenplum.ListHBARulesAtRevisionRequest, opts ...grpc.CallOption) *HBARuleAtRevisionIterator {
	var pageSize int64
	const defaultPageSize = 1000

	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &HBARuleAtRevisionIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *HBARuleAtRevisionIterator) Next() bool {
	if it.err != nil {
		return false
	}
	if len(it.items) > 1 {
		it.items[0] = nil
		it.items = it.items[1:]
		return true
	}
	it.items = nil // consume last item, if any

	if it.started {
		return false
	}
	it.started = true

	response, err := it.client.ListAtRevision(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.HbaRules
	return len(it.items) > 0
}

func (it *HBARuleAtRevisionIterator) Take(size int64) ([]*greenplum.HBARule, error) {
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

	var result []*greenplum.HBARule

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *HBARuleAtRevisionIterator) TakeAll() ([]*greenplum.HBARule, error) {
	return it.Take(0)
}

func (it *HBARuleAtRevisionIterator) Value() *greenplum.HBARule {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *HBARuleAtRevisionIterator) Error() error {
	return it.err
}

// Update implements greenplum.HBARuleServiceClient
func (c *HBARuleServiceClient) Update(ctx context.Context, in *greenplum.UpdateHBARuleRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return greenplum.NewHBARuleServiceClient(conn).Update(ctx, in, opts...)
}
