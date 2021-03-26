package sdkresolvers

import (
	"context"

	"google.golang.org/grpc"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/ydb/v1"
	ycsdk "github.com/yandex-cloud/go-sdk"
)

func YDBDatabaseResolver(name string, opts ...ResolveOption) ycsdk.Resolver {
	return &ydbDatabaseResolver{
		BaseNameResolver: NewBaseNameResolver(name, "ydb_database", opts...),
	}
}

type ydbDatabaseResolver struct {
	BaseNameResolver
}

func (r *ydbDatabaseResolver) Run(ctx context.Context, sdk *ycsdk.SDK, opts ...grpc.CallOption) error {
	err := r.ensureFolderID()
	if err != nil {
		return err
	}

	resp, err := sdk.YDB().Database().List(ctx, &ydb.ListDatabasesRequest{
		FolderId: r.FolderID(),
		PageSize: DefaultResolverPageSize,
	}, opts...)
	return r.findName(resp.GetDatabases(), err)
}
