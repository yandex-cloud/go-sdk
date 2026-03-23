package sdkresolvers

import (
	"context"

	"google.golang.org/grpc"

	baremetal "github.com/yandex-cloud/go-genproto/yandex/cloud/baremetal/v1alpha"
	ycsdk "github.com/yandex-cloud/go-sdk"
)

type baremetalPublicPrefixPoolResolver struct {
	BaseNameResolver
}

func BaremetalPublicPrefixPoolResolver(name string, opts ...ResolveOption) ycsdk.Resolver {
	return &baremetalPublicPrefixPoolResolver{
		BaseNameResolver: NewBaseNameResolver(name, "public-prefix-pool", opts...),
	}
}

func (r *baremetalPublicPrefixPoolResolver) Run(ctx context.Context, sdk *ycsdk.SDK, opts ...grpc.CallOption) error {
	err := r.ensureFolderID()
	if err != nil {
		return err
	}

	resp, err := sdk.Baremetal().PublicPrefixPool().List(ctx, &baremetal.ListPublicPrefixPoolRequest{
		FolderId: r.FolderID(),
		Filter:   CreateResolverFilter("name", r.Name),
		PageSize: DefaultResolverPageSize,
	}, opts...)
	return r.findName(resp.GetPublicPrefixPools(), err)
}
