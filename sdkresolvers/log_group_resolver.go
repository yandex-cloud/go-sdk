package sdkresolvers

import (
	"context"

	"google.golang.org/grpc"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/logging/v1"
	ycsdk "github.com/yandex-cloud/go-sdk"
)

type logGroupResolver struct {
	BaseNameResolver
}

func (r *logGroupResolver) Run(ctx context.Context, sdk *ycsdk.SDK, opts ...grpc.CallOption) error {
	err := r.ensureFolderID()
	if err != nil {
		return err
	}
	resp, err := sdk.Logging().LogGroup().List(ctx, &logging.ListLogGroupsRequest{
		FolderId: r.FolderID(),
		Filter:   CreateResolverFilter("name", r.Name),
		PageSize: DefaultResolverPageSize,
	}, opts...)
	if err != nil {
		return err
	}
	return r.findName(resp.Groups, err)
}

func LogGroupResolver(name string, opts ...ResolveOption) ycsdk.Resolver {
	return &logGroupResolver{
		BaseNameResolver: NewBaseNameResolver(name, "log-group", opts...),
	}
}
