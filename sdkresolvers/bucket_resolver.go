package sdkresolvers

import (
	"context"

	"google.golang.org/grpc"

	storage "github.com/yandex-cloud/go-genproto/yandex/cloud/storage/v1"
	ycsdk "github.com/yandex-cloud/go-sdk"
	"github.com/yandex-cloud/go-sdk/pkg/sdkerrors"
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
	resp, err := sdk.StorageAPI().Bucket().Get(ctx, &storage.GetBucketRequest{
		Name: r.Name,
		View: storage.GetBucketRequest_VIEW_BASIC,
	}, opts...)

	if err != nil {
		return sdkerrors.WithMessagef(err, "failed to find %v with name \"%v\" %v", r.resolvingObjectType, r.Name, r.coordinates())
	}

	r.SetID(resp.ResourceId)

	return nil
}
