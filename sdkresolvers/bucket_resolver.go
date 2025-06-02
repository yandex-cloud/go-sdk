package sdkresolvers

import (
	"context"

	"google.golang.org/grpc"

	storage "github.com/yandex-cloud/go-genproto/yandex/cloud/storage/v1"
	ycsdk "github.com/yandex-cloud/go-sdk"
)

type bucketResolver struct {
	BaseNameResolver
}

func BucketResolver(name string, opts ...ResolveOption) ycsdk.Resolver {
	return &bucketResolver{
		BaseNameResolver: NewBaseNameResolver(name, "bucket", opts...),
	}
}

func (r *bucketResolver) Run(ctx context.Context, sdk *ycsdk.SDK, opts ...grpc.CallOption) error {
	err := r.ensureFolderID()
	if err != nil {
		return err
	}

	resp, err := sdk.StorageAPI().Bucket().List(ctx, &storage.ListBucketsRequest{
		FolderId: r.FolderID(),
	}, opts...)
	return r.findName(resp.GetBuckets(), err)
}
