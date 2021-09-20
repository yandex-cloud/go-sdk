package sdkresolvers

import (
	"context"

	"google.golang.org/grpc"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/compute/v1"
	ycsdk "github.com/yandex-cloud/go-sdk"
)

type filesystemResolver struct {
	BaseNameResolver
}

func FilesystemResolver(name string, opts ...ResolveOption) ycsdk.Resolver {
	return &filesystemResolver{
		BaseNameResolver: NewBaseNameResolver(name, "filesystem", opts...),
	}
}

func (r *filesystemResolver) Run(ctx context.Context, sdk *ycsdk.SDK, opts ...grpc.CallOption) error {
	if err := r.ensureFolderID(); err != nil {
		return err
	}

	resp, err := sdk.Compute().Filesystem().List(ctx, &compute.ListFilesystemsRequest{
		FolderId: r.FolderID(),
		Filter:   CreateResolverFilter("name", r.Name),
		PageSize: DefaultResolverPageSize,
	}, opts...)

	return r.findName(resp.GetFilesystems(), err)
}
