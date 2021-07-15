package sdkresolvers

import (
	"context"

	"google.golang.org/grpc"

	mdbproxy "github.com/yandex-cloud/go-genproto/yandex/cloud/serverless/mdbproxy/v1"
	ycsdk "github.com/yandex-cloud/go-sdk"
)

type mdbproxyResolver struct {
	BaseNameResolver
}

func MDBProxyResolver(name string, opts ...ResolveOption) ycsdk.Resolver {
	return &mdbproxyResolver{
		BaseNameResolver: NewBaseNameResolver(name, "proxy", opts...),
	}
}

func (r *mdbproxyResolver) Run(ctx context.Context, sdk *ycsdk.SDK, opts ...grpc.CallOption) error {
	err := r.ensureFolderID()
	if err != nil {
		return err
	}

	resp, err := sdk.Serverless().MDBProxy().Proxy().List(ctx, &mdbproxy.ListProxyRequest{
		FolderId: r.FolderID(),
		Filter:   CreateResolverFilter("name", r.Name),
		PageSize: DefaultResolverPageSize,
	}, opts...)
	return r.findName(resp.GetProxies(), err)
}
