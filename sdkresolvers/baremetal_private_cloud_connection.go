package sdkresolvers

import (
	"context"

	"google.golang.org/grpc"

	baremetal "github.com/yandex-cloud/go-genproto/yandex/cloud/baremetal/v1alpha"
	ycsdk "github.com/yandex-cloud/go-sdk"
)

type baremetalPrivateCloudConnectionResolver struct {
	BaseNameResolver
}

func BaremetalPrivateCloudConnectionResolver(name string, opts ...ResolveOption) ycsdk.Resolver {
	return &baremetalPrivateCloudConnectionResolver{
		BaseNameResolver: NewBaseNameResolver(name, "private-subnet", opts...),
	}
}

func (r *baremetalPrivateCloudConnectionResolver) Run(ctx context.Context, sdk *ycsdk.SDK, opts ...grpc.CallOption) error {
	err := r.ensureFolderID()
	if err != nil {
		return err
	}

	resp, err := sdk.Baremetal().PrivateCloudConnection().List(ctx, &baremetal.ListPrivateCloudConnectionRequest{
		FolderId: r.FolderID(),
		Filter:   CreateResolverFilter("name", r.Name),
		PageSize: DefaultResolverPageSize,
	}, opts...)
	return r.findName(resp.GetPrivateCloudConnections(), err)
}
