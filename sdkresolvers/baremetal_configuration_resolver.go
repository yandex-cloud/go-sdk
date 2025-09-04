package sdkresolvers

import (
	"context"

	"google.golang.org/grpc"

	baremetal "github.com/yandex-cloud/go-genproto/yandex/cloud/baremetal/v1alpha"
	ycsdk "github.com/yandex-cloud/go-sdk"
)

type baremetalConfigurationResolver struct {
	BaseNameResolver
}

func BaremetalConfigurationResolver(name string, opts ...ResolveOption) ycsdk.Resolver {
	return &baremetalConfigurationResolver{
		BaseNameResolver: NewBaseNameResolver(name, "configuration", opts...),
	}
}

func (r *baremetalConfigurationResolver) Run(ctx context.Context, sdk *ycsdk.SDK, opts ...grpc.CallOption) error {
	err := r.ensureFolderID()
	if err != nil {
		return err
	}

	resp, err := sdk.Baremetal().Configuration().List(ctx, &baremetal.ListConfigurationsRequest{
		Filter:   CreateResolverFilter("name", r.Name),
		PageSize: DefaultResolverPageSize,
		FolderId: r.FolderID(),
	}, opts...)
	return r.findName(resp.GetConfigurations(), err)
}
