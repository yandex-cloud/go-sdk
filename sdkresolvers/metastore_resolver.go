package sdkresolvers

import (
	"context"

	"google.golang.org/grpc"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/metastore/v1"
	ycsdk "github.com/yandex-cloud/go-sdk"
)

func MetastoreClusterResolver(name string, opts ...ResolveOption) ycsdk.Resolver {
	return &metastoreClusterResolver{
		BaseNameResolver: NewBaseNameResolver(name, "cluster", opts...),
	}
}

type metastoreClusterResolver struct {
	BaseNameResolver
}

func (r *metastoreClusterResolver) Run(ctx context.Context, sdk *ycsdk.SDK, opts ...grpc.CallOption) error {
	err := r.ensureFolderID()
	if err != nil {
		return err
	}

	resp, err := sdk.Metastore().Cluster().List(ctx, &metastore.ListClustersRequest{
		FolderId: r.FolderID(),
		Filter:   CreateResolverFilter("name", r.Name),
		PageSize: DefaultResolverPageSize,
	})
	return r.findName(resp.GetClusters(), err)
}
