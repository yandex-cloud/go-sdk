// Code generated by sdkgen. DO NOT EDIT.

// nolint
package api

import (
	"context"

	"google.golang.org/grpc"

	api "github.com/yandex-cloud/go-genproto/yandex/cloud/loadtesting/api/v1"
	agent "github.com/yandex-cloud/go-genproto/yandex/cloud/loadtesting/api/v1/agent"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
)

//revive:disable

// AgentServiceClient is a api.AgentServiceClient with
// lazy GRPC connection initialization.
type AgentServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// Create implements api.AgentServiceClient
func (c *AgentServiceClient) Create(ctx context.Context, in *api.CreateAgentRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return api.NewAgentServiceClient(conn).Create(ctx, in, opts...)
}

// Delete implements api.AgentServiceClient
func (c *AgentServiceClient) Delete(ctx context.Context, in *api.DeleteAgentRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return api.NewAgentServiceClient(conn).Delete(ctx, in, opts...)
}

// Get implements api.AgentServiceClient
func (c *AgentServiceClient) Get(ctx context.Context, in *api.GetAgentRequest, opts ...grpc.CallOption) (*agent.Agent, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return api.NewAgentServiceClient(conn).Get(ctx, in, opts...)
}

// List implements api.AgentServiceClient
func (c *AgentServiceClient) List(ctx context.Context, in *api.ListAgentsRequest, opts ...grpc.CallOption) (*api.ListAgentsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return api.NewAgentServiceClient(conn).List(ctx, in, opts...)
}

type AgentIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *AgentServiceClient
	request *api.ListAgentsRequest

	items []*agent.Agent
}

func (c *AgentServiceClient) AgentIterator(ctx context.Context, req *api.ListAgentsRequest, opts ...grpc.CallOption) *AgentIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &AgentIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *AgentIterator) Next() bool {
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

	it.items = response.Agents
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *AgentIterator) Take(size int64) ([]*agent.Agent, error) {
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

	var result []*agent.Agent

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *AgentIterator) TakeAll() ([]*agent.Agent, error) {
	return it.Take(0)
}

func (it *AgentIterator) Value() *agent.Agent {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *AgentIterator) Error() error {
	return it.err
}

// Update implements api.AgentServiceClient
func (c *AgentServiceClient) Update(ctx context.Context, in *api.UpdateAgentRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return api.NewAgentServiceClient(conn).Update(ctx, in, opts...)
}
