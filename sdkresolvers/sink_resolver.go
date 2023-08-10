package sdkresolvers

import (
	"context"

	"google.golang.org/grpc"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/logging/v1"
	ycsdk "github.com/yandex-cloud/go-sdk"
)

type sinkResolver struct {
	BaseNameResolver
}

func (r *sinkResolver) Run(ctx context.Context, sdk *ycsdk.SDK, opts ...grpc.CallOption) error {
	err := r.ensureFolderID()
	if err != nil {
		return err
	}
	resp, err := sdk.Logging().Sink().List(ctx, &logging.ListSinksRequest{
		FolderId: r.FolderID(),
		Filter:   CreateResolverFilter("name", r.Name),
		PageSize: DefaultResolverPageSize,
	}, opts...)
	if err != nil {
		return err
	}
	return r.findName(resp.Sinks, err)
}

func SinkResolver(name string, opts ...ResolveOption) ycsdk.Resolver {
	return &sinkResolver{
		BaseNameResolver: NewBaseNameResolver(name, "sink", opts...),
	}
}
