// Code generated by sdkgen. DO NOT EDIT.

package api

import (
	"context"

	"google.golang.org/grpc"
)

// Loadtesting provides access to "api" component of Yandex.Cloud
type Loadtesting struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// NewLoadtesting creates instance of Loadtesting
func NewLoadtesting(g func(ctx context.Context) (*grpc.ClientConn, error)) *Loadtesting {
	return &Loadtesting{g}
}

// Agent gets AgentService client
func (l *Loadtesting) Agent() *AgentServiceClient {
	return &AgentServiceClient{getConn: l.getConn}
}
