package sdkresolvers

import (
	"context"

	"google.golang.org/grpc"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/serverless/apigateway/v1"
	ycsdk "github.com/yandex-cloud/go-sdk"
)

type apigatewayResolver struct {
	BaseNameResolver
}

func APIGatewayResolver(name string, opts ...ResolveOption) ycsdk.Resolver {
	return &apigatewayResolver{
		BaseNameResolver: NewBaseNameResolver(name, "api-gateway", opts...),
	}
}

func (r *apigatewayResolver) Run(ctx context.Context, sdk *ycsdk.SDK, opts ...grpc.CallOption) error {
	err := r.ensureFolderID()
	if err != nil {
		return err
	}

	resp, err := sdk.Serverless().APIGateway().ApiGateway().List(ctx, &apigateway.ListApiGatewayRequest{
		FolderId: r.FolderID(),
		Filter:   CreateResolverFilter("name", r.Name),
		PageSize: DefaultResolverPageSize,
	}, opts...)
	return r.findName(resp.GetApiGateways(), err)
}
