// Code generated by sdkgen. DO NOT EDIT.

package cloudrouter

import (
	"context"

	"google.golang.org/grpc"
)

// CloudRouter provides access to "cloudrouter" component of Yandex.Cloud
type CloudRouter struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// NewCloudRouter creates instance of CloudRouter
func NewCloudRouter(g func(ctx context.Context) (*grpc.ClientConn, error)) *CloudRouter {
	return &CloudRouter{g}
}

// RoutingInstance gets RoutingInstanceService client
func (c *CloudRouter) RoutingInstance() *RoutingInstanceServiceClient {
	return &RoutingInstanceServiceClient{getConn: c.getConn}
}