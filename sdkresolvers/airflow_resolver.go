package sdkresolvers

import (
	"context"

	"google.golang.org/grpc"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/airflow/v1"
	ycsdk "github.com/yandex-cloud/go-sdk"
)

func AirflowClusterResolver(name string, opts ...ResolveOption) ycsdk.Resolver {
	return &airflowClusterResolver{
		BaseNameResolver: NewBaseNameResolver(name, "cluster", opts...),
	}
}

type airflowClusterResolver struct {
	BaseNameResolver
}

func (r *airflowClusterResolver) Run(ctx context.Context, sdk *ycsdk.SDK, opts ...grpc.CallOption) error {
	err := r.ensureFolderID()
	if err != nil {
		return err
	}

	resp, err := sdk.Airflow().Cluster().List(ctx, &airflow.ListClustersRequest{
		FolderId: r.FolderID(),
		Filter:   CreateResolverFilter("name", r.Name),
		PageSize: DefaultResolverPageSize,
	})
	return r.findName(resp.GetClusters(), err)
}
