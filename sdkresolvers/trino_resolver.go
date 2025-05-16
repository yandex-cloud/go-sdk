package sdkresolvers

import (
	"context"

	"google.golang.org/grpc"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/trino/v1"
	ycsdk "github.com/yandex-cloud/go-sdk"
)

func TrinoClusterResolver(name string, opts ...ResolveOption) ycsdk.Resolver {
	return &trinoClusterResolver{
		BaseNameResolver: NewBaseNameResolver(name, "cluster", opts...),
	}
}

type trinoClusterResolver struct {
	BaseNameResolver
}

func (r *trinoClusterResolver) Run(ctx context.Context, sdk *ycsdk.SDK, opts ...grpc.CallOption) error {
	err := r.ensureFolderID()
	if err != nil {
		return err
	}

	resp, err := sdk.Trino().Cluster().List(ctx, &trino.ListClustersRequest{
		FolderId: r.FolderID(),
		Filter:   CreateResolverFilter("name", r.Name),
		PageSize: DefaultResolverPageSize,
	})
	return r.findName(resp.GetClusters(), err)
}
