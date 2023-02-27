package sdkresolvers

import (
	"context"

	"google.golang.org/grpc"

	compute "github.com/yandex-cloud/go-genproto/yandex/cloud/compute/v1"
	ycsdk "github.com/yandex-cloud/go-sdk"
)

type gpuClusterResolver struct {
	BaseNameResolver
}

func GpuClusterResolver(name string, opts ...ResolveOption) ycsdk.Resolver {
	return &gpuClusterResolver{
		BaseNameResolver: NewBaseNameResolver(name, "GPU cluster", opts...),
	}
}

func (r *gpuClusterResolver) Run(ctx context.Context, sdk *ycsdk.SDK, opts ...grpc.CallOption) error {
	err := r.ensureFolderID()
	if err != nil {
		return err
	}

	resp, err := sdk.Compute().GpuCluster().List(ctx, &compute.ListGpuClustersRequest{
		FolderId: r.FolderID(),
		Filter:   CreateResolverFilter("name", r.Name),
		PageSize: DefaultResolverPageSize,
	}, opts...)
	return r.findName(resp.GetGpuClusters(), err)
}
