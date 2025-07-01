package sdkresolvers

import (
	"context"

	"google.golang.org/grpc"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/cloudregistry/v1"
	ycsdk "github.com/yandex-cloud/go-sdk"
)

type cloudRegistryResolver struct {
	BaseNameResolver
}

func CloudRegistryResolver(name string, opts ...ResolveOption) ycsdk.Resolver {
	return &cloudRegistryResolver{
		BaseNameResolver: NewBaseNameResolver(name, "registry", opts...), // todo device registry resolver полный перебор
	}
}

func (r *cloudRegistryResolver) Run(ctx context.Context, sdk *ycsdk.SDK, opts ...grpc.CallOption) error {
	err := r.ensureFolderID()
	if err != nil {
		return err
	}

	nextPageToken := ""
	var registries []*cloudregistry.Registry

	for ok := true; ok; ok = len(nextPageToken) > 0 {
		resp, err := sdk.CloudRegistry().Registry().List(ctx, &cloudregistry.ListRegistriesRequest{
			FolderId:  r.FolderID(),
			PageSize:  DefaultResolverPageSize,
			PageToken: nextPageToken,
		}, opts...)
		if err != nil {
			return err
		}
		nextPageToken = resp.GetNextPageToken()
		registries = append(registries, resp.GetRegistries()...)
	}

	return r.findName(registries, err)
}
