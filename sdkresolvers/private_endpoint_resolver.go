package sdkresolvers

import (
	"context"

	"google.golang.org/grpc"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/vpc/v1/privatelink"
	ycsdk "github.com/yandex-cloud/go-sdk"
)

type privateEndpointResolver struct {
	BaseNameResolver
}

func PrivateEndpointResolver(name string, opts ...ResolveOption) ycsdk.Resolver {
	return &privateEndpointResolver{
		BaseNameResolver: NewBaseNameResolver(name, "private-endpoint", opts...),
	}
}

func (r *privateEndpointResolver) Run(ctx context.Context, sdk *ycsdk.SDK, opts ...grpc.CallOption) error {
	err := r.ensureFolderID()
	if err != nil {
		return err
	}

	resp, err := sdk.VPCPrivateLink().PrivateEndpoint().List(ctx, &privatelink.ListPrivateEndpointsRequest{
		Container: &privatelink.ListPrivateEndpointsRequest_FolderId{
			FolderId: r.FolderID(),
		},
		Filter:   CreateResolverFilter("name", r.Name),
		PageSize: DefaultResolverPageSize,
	}, opts...)
	return r.findName(resp.GetPrivateEndpoints(), err)
}
