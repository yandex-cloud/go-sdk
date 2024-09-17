package sdkresolvers

import (
	"context"

	"google.golang.org/grpc"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/serverless/eventrouter/v1"
	ycsdk "github.com/yandex-cloud/go-sdk"
)

type eventrouterConnectorResolver struct {
	BaseNameResolver
}

func EventrouterConnectorResolver(name string, opts ...ResolveOption) ycsdk.Resolver {
	return &eventrouterConnectorResolver{
		BaseNameResolver: NewBaseNameResolver(name, "connector", opts...),
	}
}

func (r *eventrouterConnectorResolver) Run(ctx context.Context, sdk *ycsdk.SDK, opts ...grpc.CallOption) error {
	if err := r.ensureFolderID(); err != nil {
		return err
	}
	request := &eventrouter.ListConnectorsRequest{
		ContainerId: &eventrouter.ListConnectorsRequest_FolderId{FolderId: r.FolderID()},
		Filter:      CreateResolverFilter("name", r.Name),
		PageSize:    DefaultResolverPageSize,
	}
	resp, err := sdk.Serverless().Eventrouter().Connector().List(ctx, request, opts...)
	return r.findName(resp.GetConnectors(), err)
}
