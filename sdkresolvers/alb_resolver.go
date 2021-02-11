package sdkresolvers

import (
	"context"

	"google.golang.org/grpc"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/apploadbalancer/v1"
	ycsdk "github.com/yandex-cloud/go-sdk"
)

type albLoadBalancerResolver struct {
	BaseNameResolver
}

func ApplicationLoadBalancerResolver(name string, opts ...ResolveOption) ycsdk.Resolver {
	return &albLoadBalancerResolver{
		BaseNameResolver: NewBaseNameResolver(name, "application load balancer", opts...),
	}
}

func (r *albLoadBalancerResolver) Run(ctx context.Context, sdk *ycsdk.SDK, opts ...grpc.CallOption) error {
	err := r.ensureFolderID()
	if err != nil {
		return err
	}

	resp, err := sdk.ApplicationLoadBalancer().LoadBalancer().List(ctx, &apploadbalancer.ListLoadBalancersRequest{
		FolderId: r.FolderID(),
		Filter:   CreateResolverFilter("name", r.Name),
		PageSize: DefaultResolverPageSize,
	}, opts...)
	return r.findName(resp.GetLoadBalancers(), err)
}

type albTargetGroupResolver struct {
	BaseNameResolver
}

func ALBTargetGroupResolver(name string, opts ...ResolveOption) ycsdk.Resolver {
	return &albTargetGroupResolver{
		BaseNameResolver: NewBaseNameResolver(name, "alb target group", opts...),
	}
}

func (r *albTargetGroupResolver) Run(ctx context.Context, sdk *ycsdk.SDK, opts ...grpc.CallOption) error {
	err := r.ensureFolderID()
	if err != nil {
		return err
	}

	resp, err := sdk.ApplicationLoadBalancer().TargetGroup().List(ctx, &apploadbalancer.ListTargetGroupsRequest{
		FolderId: r.FolderID(),
		Filter:   CreateResolverFilter("name", r.Name),
		PageSize: DefaultResolverPageSize,
	}, opts...)
	return r.findName(resp.GetTargetGroups(), err)
}

type albBackendGroupResolver struct {
	BaseNameResolver
}

func ALBBackendGroupResolver(name string, opts ...ResolveOption) ycsdk.Resolver {
	return &albBackendGroupResolver{
		BaseNameResolver: NewBaseNameResolver(name, "alb backend group", opts...),
	}
}

func (r *albBackendGroupResolver) Run(ctx context.Context, sdk *ycsdk.SDK, opts ...grpc.CallOption) error {
	err := r.ensureFolderID()
	if err != nil {
		return err
	}

	resp, err := sdk.ApplicationLoadBalancer().BackendGroup().List(ctx, &apploadbalancer.ListBackendGroupsRequest{
		FolderId: r.FolderID(),
		Filter:   CreateResolverFilter("name", r.Name),
		PageSize: DefaultResolverPageSize,
	}, opts...)
	return r.findName(resp.GetBackendGroups(), err)
}

type albHTTPRouterResolver struct {
	BaseNameResolver
}

func ALBHTTPRouterResolver(name string, opts ...ResolveOption) ycsdk.Resolver {
	return &albHTTPRouterResolver{
		BaseNameResolver: NewBaseNameResolver(name, "alb http router", opts...),
	}
}

func (r *albHTTPRouterResolver) Run(ctx context.Context, sdk *ycsdk.SDK, opts ...grpc.CallOption) error {
	err := r.ensureFolderID()
	if err != nil {
		return err
	}

	resp, err := sdk.ApplicationLoadBalancer().HttpRouter().List(ctx, &apploadbalancer.ListHttpRoutersRequest{
		FolderId: r.FolderID(),
		Filter:   CreateResolverFilter("name", r.Name),
		PageSize: DefaultResolverPageSize,
	}, opts...)
	return r.findName(resp.GetHttpRouters(), err)
}
