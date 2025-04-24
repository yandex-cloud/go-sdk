package sdkresolvers

import (
	"context"

	"google.golang.org/grpc"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/spark/v1"
	ycsdk "github.com/yandex-cloud/go-sdk"
)

func SparkClusterResolver(name string, opts ...ResolveOption) ycsdk.Resolver {
	return &sparkClusterResolver{
		BaseNameResolver: NewBaseNameResolver(name, "cluster", opts...),
	}
}

type sparkClusterResolver struct {
	BaseNameResolver
}

func (r *sparkClusterResolver) Run(ctx context.Context, sdk *ycsdk.SDK, opts ...grpc.CallOption) error {
	err := r.ensureFolderID()
	if err != nil {
		return err
	}

	resp, err := sdk.Spark().Cluster().List(ctx, &spark.ListClustersRequest{
		FolderId: r.FolderID(),
		Filter:   CreateResolverFilter("name", r.Name),
		PageSize: DefaultResolverPageSize,
	})
	return r.findName(resp.GetClusters(), err)
}
