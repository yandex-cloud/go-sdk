package sdkresolvers

import (
	"context"

	"google.golang.org/grpc"

	compute "github.com/yandex-cloud/go-genproto/yandex/cloud/compute/v1"
	ycsdk "github.com/yandex-cloud/go-sdk"
)

type reservedInstancePoolResolver struct {
	BaseNameResolver
}

func ReservedInstancePoolResolver(name string, opts ...ResolveOption) ycsdk.Resolver {
	return &reservedInstancePoolResolver{
		BaseNameResolver: NewBaseNameResolver(name, "reserved_instance_pool", opts...),
	}
}

func (r *reservedInstancePoolResolver) Run(ctx context.Context, sdk *ycsdk.SDK, opts ...grpc.CallOption) error {
	err := r.ensureFolderID()
	if err != nil {
		return err
	}

	resp, err := sdk.Compute().ReservedInstancePool().List(ctx, &compute.ListReservedInstancePoolsRequest{
		FolderId: r.FolderID(),
		Filter:   CreateResolverFilter("name", r.Name),
		PageSize: DefaultResolverPageSize,
	}, opts...)
	return r.findName(resp.GetReservedInstancePools(), err)
}
