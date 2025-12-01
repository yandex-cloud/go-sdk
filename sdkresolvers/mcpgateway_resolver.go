package sdkresolvers

import (
	"context"

	"google.golang.org/grpc"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/serverless/mcpgateway/v1"
	ycsdk "github.com/yandex-cloud/go-sdk"
)

type mcpGatewayResolver struct {
	BaseNameResolver
}

func McpGatewayResolver(name string, opts ...ResolveOption) ycsdk.Resolver {
	return &mcpGatewayResolver{
		BaseNameResolver: NewBaseNameResolver(name, "mcpgateway", opts...),
	}
}

func (r *mcpGatewayResolver) Run(ctx context.Context, sdk *ycsdk.SDK, opts ...grpc.CallOption) error {
	err := r.ensureFolderID()
	if err != nil {
		return err
	}

	resp, err := sdk.Serverless().McpGateway().McpGateway().List(ctx, &mcpgateway.ListMcpGatewayRequest{
		FolderId: r.FolderID(),
		Filter:   CreateResolverFilter("name", r.Name),
		PageSize: DefaultResolverPageSize,
	}, opts...)
	return r.findName(resp.GetGateways(), err)
}
