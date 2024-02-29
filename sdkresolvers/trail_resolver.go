package sdkresolvers

import (
	"context"

	"google.golang.org/grpc"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/audittrails/v1"
	ycsdk "github.com/yandex-cloud/go-sdk"
)

type trailResolver struct {
	BaseNameResolver
}

func TrailResolver(name string, opts ...ResolveOption) ycsdk.Resolver {
	return &trailResolver{
		BaseNameResolver: NewBaseNameResolver(name, "trail", opts...),
	}
}

func (r *trailResolver) Run(ctx context.Context, sdk *ycsdk.SDK, opts ...grpc.CallOption) error {
	err := r.ensureFolderID()
	if err != nil {
		return err
	}

	resp, err := sdk.AuditTrails().Trail().List(ctx, &audittrails.ListTrailsRequest{
		FolderId: r.FolderID(),
		Filter:   CreateResolverFilter("name", r.Name),
		PageSize: DefaultResolverPageSize,
	}, opts...)

	return r.findName(resp.GetTrails(), err)
}
