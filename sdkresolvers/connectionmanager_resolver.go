package sdkresolvers

import (
	"context"

	"google.golang.org/grpc"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/connectionmanager/v1"
	ycsdk "github.com/yandex-cloud/go-sdk"
)

type connectionResolver struct {
	BaseNameResolver
}

func NewConnectioManagerConnectionResolver(name string, opts ...ResolveOption) ycsdk.Resolver {
	return &connectionResolver{
		BaseNameResolver: NewBaseNameResolver(name, "connection", opts...),
	}
}

func (r *connectionResolver) Run(ctx context.Context, sdk *ycsdk.SDK, opts ...grpc.CallOption) error {
	err := r.ensureFolderID()
	if err != nil {
		return err
	}

	resp, err := sdk.ConnectionManager().Connection().List(ctx, &connectionmanager.ListConnectionRequest{
		FolderId: r.FolderID(),
		PageSize: DefaultResolverPageSize,
	}, opts...)

	return r.findName(resp.GetConnection(), err)
}
