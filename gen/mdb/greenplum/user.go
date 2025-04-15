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

// UserServiceClient is a greenplum.UserServiceClient with
// lazy GRPC connection initialization.
type UserServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// Create implements greenplum.UserServiceClient
func (c *UserServiceClient) Create(ctx context.Context, in *greenplum.CreateUserRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return greenplum.NewUserServiceClient(conn).Create(ctx, in, opts...)
}

// Delete implements greenplum.UserServiceClient
func (c *UserServiceClient) Delete(ctx context.Context, in *greenplum.DeleteUserRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return greenplum.NewUserServiceClient(conn).Delete(ctx, in, opts...)
}

// List implements greenplum.UserServiceClient
func (c *UserServiceClient) List(ctx context.Context, in *greenplum.ListUsersRequest, opts ...grpc.CallOption) (*greenplum.ListUsersResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return greenplum.NewUserServiceClient(conn).List(ctx, in, opts...)
}

type UserIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *UserServiceClient
	request *greenplum.ListUsersRequest

	items []*greenplum.User
}

func (c *UserServiceClient) UserIterator(ctx context.Context, req *greenplum.ListUsersRequest, opts ...grpc.CallOption) *UserIterator {
	var pageSize int64
	const defaultPageSize = 1000

	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &UserIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *UserIterator) Next() bool {
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

	it.items = response.Users
	return len(it.items) > 0
}

func (it *UserIterator) Take(size int64) ([]*greenplum.User, error) {
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

	var result []*greenplum.User

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *UserIterator) TakeAll() ([]*greenplum.User, error) {
	return it.Take(0)
}

func (it *UserIterator) Value() *greenplum.User {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *UserIterator) Error() error {
	return it.err
}

// Update implements greenplum.UserServiceClient
func (c *UserServiceClient) Update(ctx context.Context, in *greenplum.UpdateUserRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return greenplum.NewUserServiceClient(conn).Update(ctx, in, opts...)
}
