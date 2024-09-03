// Code generated by sdkgen. DO NOT EDIT.

// nolint
package smartcaptcha

import (
	"context"

	"google.golang.org/grpc"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	smartcaptcha "github.com/yandex-cloud/go-genproto/yandex/cloud/smartcaptcha/v1"
)

//revive:disable

// CaptchaServiceClient is a smartcaptcha.CaptchaServiceClient with
// lazy GRPC connection initialization.
type CaptchaServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// Create implements smartcaptcha.CaptchaServiceClient
func (c *CaptchaServiceClient) Create(ctx context.Context, in *smartcaptcha.CreateCaptchaRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return smartcaptcha.NewCaptchaServiceClient(conn).Create(ctx, in, opts...)
}

// Delete implements smartcaptcha.CaptchaServiceClient
func (c *CaptchaServiceClient) Delete(ctx context.Context, in *smartcaptcha.DeleteCaptchaRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return smartcaptcha.NewCaptchaServiceClient(conn).Delete(ctx, in, opts...)
}

// Get implements smartcaptcha.CaptchaServiceClient
func (c *CaptchaServiceClient) Get(ctx context.Context, in *smartcaptcha.GetCaptchaRequest, opts ...grpc.CallOption) (*smartcaptcha.Captcha, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return smartcaptcha.NewCaptchaServiceClient(conn).Get(ctx, in, opts...)
}

// GetSecretKey implements smartcaptcha.CaptchaServiceClient
func (c *CaptchaServiceClient) GetSecretKey(ctx context.Context, in *smartcaptcha.GetCaptchaRequest, opts ...grpc.CallOption) (*smartcaptcha.CaptchaSecretKey, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return smartcaptcha.NewCaptchaServiceClient(conn).GetSecretKey(ctx, in, opts...)
}

// List implements smartcaptcha.CaptchaServiceClient
func (c *CaptchaServiceClient) List(ctx context.Context, in *smartcaptcha.ListCaptchasRequest, opts ...grpc.CallOption) (*smartcaptcha.ListCaptchasResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return smartcaptcha.NewCaptchaServiceClient(conn).List(ctx, in, opts...)
}

type CaptchaIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *CaptchaServiceClient
	request *smartcaptcha.ListCaptchasRequest

	items []*smartcaptcha.Captcha
}

func (c *CaptchaServiceClient) CaptchaIterator(ctx context.Context, req *smartcaptcha.ListCaptchasRequest, opts ...grpc.CallOption) *CaptchaIterator {
	var pageSize int64
	const defaultPageSize = 1000

	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &CaptchaIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *CaptchaIterator) Next() bool {
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

	it.items = response.Resources
	return len(it.items) > 0
}

func (it *CaptchaIterator) Take(size int64) ([]*smartcaptcha.Captcha, error) {
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

	var result []*smartcaptcha.Captcha

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *CaptchaIterator) TakeAll() ([]*smartcaptcha.Captcha, error) {
	return it.Take(0)
}

func (it *CaptchaIterator) Value() *smartcaptcha.Captcha {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *CaptchaIterator) Error() error {
	return it.err
}

// Update implements smartcaptcha.CaptchaServiceClient
func (c *CaptchaServiceClient) Update(ctx context.Context, in *smartcaptcha.UpdateCaptchaRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return smartcaptcha.NewCaptchaServiceClient(conn).Update(ctx, in, opts...)
}