package sdkresolvers

import (
	"context"

	"google.golang.org/grpc"

	compute "github.com/yandex-cloud/go-genproto/yandex/cloud/compute/v1"
	ycsdk "github.com/yandex-cloud/go-sdk"
)

type snapshotScheduleResolver struct {
	BaseNameResolver
}

func SnapshotScheduleResolver(name string, opts ...ResolveOption) ycsdk.Resolver {
	return &snapshotScheduleResolver{
		BaseNameResolver: NewBaseNameResolver(name, "snapshot schedule", opts...),
	}
}

func (r *snapshotScheduleResolver) Run(ctx context.Context, sdk *ycsdk.SDK, opts ...grpc.CallOption) error {
	err := r.ensureFolderID()
	if err != nil {
		return err
	}

	resp, err := sdk.Compute().SnapshotSchedule().List(ctx, &compute.ListSnapshotSchedulesRequest{
		FolderId: r.FolderID(),
		Filter:   CreateResolverFilter("name", r.Name),
		PageSize: DefaultResolverPageSize,
	}, opts...)
	return r.findName(resp.GetSnapshotSchedules(), err)
}
