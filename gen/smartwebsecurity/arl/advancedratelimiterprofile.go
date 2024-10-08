// Code generated by sdkgen. DO NOT EDIT.

// nolint
package advanced_rate_limiter

import (
	"context"

	"google.golang.org/grpc"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	advanced_rate_limiter "github.com/yandex-cloud/go-genproto/yandex/cloud/smartwebsecurity/v1/advanced_rate_limiter"
)

//revive:disable

// AdvancedRateLimiterProfileServiceClient is a advanced_rate_limiter.AdvancedRateLimiterProfileServiceClient with
// lazy GRPC connection initialization.
type AdvancedRateLimiterProfileServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// Create implements advanced_rate_limiter.AdvancedRateLimiterProfileServiceClient
func (c *AdvancedRateLimiterProfileServiceClient) Create(ctx context.Context, in *advanced_rate_limiter.CreateAdvancedRateLimiterProfileRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return advanced_rate_limiter.NewAdvancedRateLimiterProfileServiceClient(conn).Create(ctx, in, opts...)
}

// Delete implements advanced_rate_limiter.AdvancedRateLimiterProfileServiceClient
func (c *AdvancedRateLimiterProfileServiceClient) Delete(ctx context.Context, in *advanced_rate_limiter.DeleteAdvancedRateLimiterProfileRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return advanced_rate_limiter.NewAdvancedRateLimiterProfileServiceClient(conn).Delete(ctx, in, opts...)
}

// Get implements advanced_rate_limiter.AdvancedRateLimiterProfileServiceClient
func (c *AdvancedRateLimiterProfileServiceClient) Get(ctx context.Context, in *advanced_rate_limiter.GetAdvancedRateLimiterProfileRequest, opts ...grpc.CallOption) (*advanced_rate_limiter.AdvancedRateLimiterProfile, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return advanced_rate_limiter.NewAdvancedRateLimiterProfileServiceClient(conn).Get(ctx, in, opts...)
}

// List implements advanced_rate_limiter.AdvancedRateLimiterProfileServiceClient
func (c *AdvancedRateLimiterProfileServiceClient) List(ctx context.Context, in *advanced_rate_limiter.ListAdvancedRateLimiterProfilesRequest, opts ...grpc.CallOption) (*advanced_rate_limiter.ListAdvancedRateLimiterProfilesResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return advanced_rate_limiter.NewAdvancedRateLimiterProfileServiceClient(conn).List(ctx, in, opts...)
}

type AdvancedRateLimiterProfileIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *AdvancedRateLimiterProfileServiceClient
	request *advanced_rate_limiter.ListAdvancedRateLimiterProfilesRequest

	items []*advanced_rate_limiter.AdvancedRateLimiterProfile
}

func (c *AdvancedRateLimiterProfileServiceClient) AdvancedRateLimiterProfileIterator(ctx context.Context, req *advanced_rate_limiter.ListAdvancedRateLimiterProfilesRequest, opts ...grpc.CallOption) *AdvancedRateLimiterProfileIterator {
	var pageSize int64
	const defaultPageSize = 1000

	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &AdvancedRateLimiterProfileIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *AdvancedRateLimiterProfileIterator) Next() bool {
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

	it.items = response.AdvancedRateLimiterProfiles
	return len(it.items) > 0
}

func (it *AdvancedRateLimiterProfileIterator) Take(size int64) ([]*advanced_rate_limiter.AdvancedRateLimiterProfile, error) {
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

	var result []*advanced_rate_limiter.AdvancedRateLimiterProfile

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *AdvancedRateLimiterProfileIterator) TakeAll() ([]*advanced_rate_limiter.AdvancedRateLimiterProfile, error) {
	return it.Take(0)
}

func (it *AdvancedRateLimiterProfileIterator) Value() *advanced_rate_limiter.AdvancedRateLimiterProfile {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *AdvancedRateLimiterProfileIterator) Error() error {
	return it.err
}

// Update implements advanced_rate_limiter.AdvancedRateLimiterProfileServiceClient
func (c *AdvancedRateLimiterProfileServiceClient) Update(ctx context.Context, in *advanced_rate_limiter.UpdateAdvancedRateLimiterProfileRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return advanced_rate_limiter.NewAdvancedRateLimiterProfileServiceClient(conn).Update(ctx, in, opts...)
}
