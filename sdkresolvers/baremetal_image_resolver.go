package sdkresolvers

import (
	"context"

	"google.golang.org/grpc"

	baremetal "github.com/yandex-cloud/go-genproto/yandex/cloud/baremetal/v1alpha"
	ycsdk "github.com/yandex-cloud/go-sdk"
)

type baremetalImageResolver struct {
	BaseNameResolver
}

func BaremetalImageResolver(name string, opts ...ResolveOption) ycsdk.Resolver {
	return &baremetalImageResolver{
		BaseNameResolver: NewBaseNameResolver(name, "boot-image", opts...),
	}
}

func (r *baremetalImageResolver) Run(ctx context.Context, sdk *ycsdk.SDK, opts ...grpc.CallOption) error {
	err := r.ensureFolderID()
	if err != nil {
		return err
	}

	resp, err := sdk.Baremetal().Image().List(ctx, &baremetal.ListImagesRequest{
		FolderId: r.FolderID(),
		Filter:   CreateResolverFilter("name", r.Name),
		PageSize: DefaultResolverPageSize,
	}, opts...)
	return r.findName(resp.GetImages(), err)
}
