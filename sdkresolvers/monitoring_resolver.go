package sdkresolvers

import (
	"context"

	"google.golang.org/grpc"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/monitoring/v3"
	ycsdk "github.com/yandex-cloud/go-sdk"
)

type monitoringDashboardResolver struct {
	BaseNameResolver
}

func MonitoringDashboardResolver(name string, opts ...ResolveOption) ycsdk.Resolver {
	return &monitoringDashboardResolver{
		BaseNameResolver: NewBaseNameResolver(name, "dashboard", opts...),
	}
}

func (r *monitoringDashboardResolver) Run(ctx context.Context, sdk *ycsdk.SDK, opts ...grpc.CallOption) error {
	err := r.ensureFolderID()
	if err != nil {
		return err
	}

	req := &monitoring.ListDashboardsRequest{
		Container: &monitoring.ListDashboardsRequest_FolderId{
			FolderId: r.FolderID(),
		},
		Filter:   CreateResolverFilter("name", r.Name),
		PageSize: DefaultResolverPageSize,
	}
	resp, err := sdk.Monitoring().Dashboard().List(ctx, req, opts...)
	return r.findName(resp.GetDashboards(), err)
}
