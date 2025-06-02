package sdkresolvers

import (
	"context"

	"google.golang.org/grpc"

	baremetal "github.com/yandex-cloud/go-genproto/yandex/cloud/baremetal/v1alpha"
	ycsdk "github.com/yandex-cloud/go-sdk"
)

type baremetalPrivateSubnetResolver struct {
	BaseNameResolver
}

func BaremetalPrivateSubnetResolver(name string, opts ...ResolveOption) ycsdk.Resolver {
	return &baremetalPrivateSubnetResolver{
		BaseNameResolver: NewBaseNameResolver(name, "private-subnet", opts...),
	}
}

func (r *baremetalPrivateSubnetResolver) Run(ctx context.Context, sdk *ycsdk.SDK, opts ...grpc.CallOption) error {
	err := r.ensureFolderID()
	if err != nil {
		return err
	}

	resp, err := sdk.Baremetal().PrivateSubnet().List(ctx, &baremetal.ListPrivateSubnetRequest{
		FolderId: r.FolderID(),
		Filter:   CreateResolverFilter("name", r.Name),
		PageSize: DefaultResolverPageSize,
	}, opts...)
	return r.findName(resp.GetPrivateSubnets(), err)
}

type baremetalPublicSubnetResolver struct {
	BaseNameResolver
}

func BaremetalPublicSubnetResolver(name string, opts ...ResolveOption) ycsdk.Resolver {
	return &baremetalPublicSubnetResolver{
		BaseNameResolver: NewBaseNameResolver(name, "public_subnet", opts...),
	}
}

func (r *baremetalPublicSubnetResolver) Run(ctx context.Context, sdk *ycsdk.SDK, opts ...grpc.CallOption) error {
	err := r.ensureFolderID()
	if err != nil {
		return err
	}

	resp, err := sdk.Baremetal().PublicSubnet().List(ctx, &baremetal.ListPublicSubnetRequest{
		FolderId: r.FolderID(),
		Filter:   CreateResolverFilter("name", r.Name),
		PageSize: DefaultResolverPageSize,
	}, opts...)
	return r.findName(resp.GetPublicSubnets(), err)
}
