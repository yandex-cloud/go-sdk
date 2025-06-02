package sdkresolvers

import (
	"context"

	"google.golang.org/grpc"

	baremetal "github.com/yandex-cloud/go-genproto/yandex/cloud/baremetal/v1alpha"
	ycsdk "github.com/yandex-cloud/go-sdk"
)

type baremetalServerResolver struct {
	BaseNameResolver
}

func BaremetalServerResolver(name string, opts ...ResolveOption) ycsdk.Resolver {
	return &baremetalServerResolver{
		BaseNameResolver: NewBaseNameResolver(name, "server", opts...),
	}
}

func (r *baremetalServerResolver) Run(ctx context.Context, sdk *ycsdk.SDK, opts ...grpc.CallOption) error {
	err := r.ensureFolderID()
	if err != nil {
		return err
	}

	resp, err := sdk.Baremetal().Server().List(ctx, &baremetal.ListServerRequest{
		FolderId: r.FolderID(),
		Filter:   CreateResolverFilter("name", r.Name),
		PageSize: DefaultResolverPageSize,
	}, opts...)
	return r.findName(resp.GetServers(), err)
}
