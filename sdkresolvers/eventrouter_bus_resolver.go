package sdkresolvers

import (
	"context"

	"google.golang.org/grpc"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/serverless/eventrouter/v1"
	ycsdk "github.com/yandex-cloud/go-sdk"
)

type eventrouterBusResolver struct {
	BaseNameResolver
}

func EventrouterBusResolver(name string, opts ...ResolveOption) ycsdk.Resolver {
	return &eventrouterBusResolver{
		BaseNameResolver: NewBaseNameResolver(name, "bus", opts...),
	}
}

func (r *eventrouterBusResolver) Run(ctx context.Context, sdk *ycsdk.SDK, opts ...grpc.CallOption) error {
	err := r.ensureFolderID()
	if err != nil {
		return err
	}
	resp, err := sdk.Serverless().Eventrouter().Bus().List(ctx, &eventrouter.ListBusesRequest{
		FolderId: r.FolderID(),
		Filter:   CreateResolverFilter("name", r.Name),
		PageSize: DefaultResolverPageSize,
	}, opts...)
	return r.findName(resp.GetBuses(), err)
}
