package sdkresolvers

import (
	"context"

	"google.golang.org/grpc"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/logging/v1"
	ycsdk "github.com/yandex-cloud/go-sdk"
)

type exportResolver struct {
	BaseNameResolver
}

func (r *exportResolver) Run(ctx context.Context, sdk *ycsdk.SDK, opts ...grpc.CallOption) error {
	err := r.ensureFolderID()
	if err != nil {
		return err
	}
	resp, err := sdk.Logging().Export().List(ctx, &logging.ListExportsRequest{
		FolderId: r.FolderID(),
		Filter:   CreateResolverFilter("name", r.Name),
		PageSize: DefaultResolverPageSize,
	}, opts...)
	if err != nil {
		return err
	}
	return r.findName(resp.Exports, err)
}

func ExportResolver(name string, opts ...ResolveOption) ycsdk.Resolver {
	return &exportResolver{
		BaseNameResolver: NewBaseNameResolver(name, "export", opts...),
	}
}
