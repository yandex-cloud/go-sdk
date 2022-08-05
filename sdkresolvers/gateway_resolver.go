// Copyright (c) 2018 Yandex LLC. All rights reserved.
// Author: Maxim Kolganov <manykey@yandex-team.ru>

package sdkresolvers

import (
	"context"

	"google.golang.org/grpc"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/vpc/v1"
	ycsdk "github.com/yandex-cloud/go-sdk"
)

type gatewayResolver struct {
	BaseNameResolver
}

func GatewayResolver(name string, opts ...ResolveOption) ycsdk.Resolver {
	return &gatewayResolver{
		BaseNameResolver: NewBaseNameResolver(name, "gateway", opts...),
	}
}

func (r *gatewayResolver) Run(ctx context.Context, sdk *ycsdk.SDK, opts ...grpc.CallOption) error {
	err := r.ensureFolderID()
	if err != nil {
		return err
	}

	resp, err := sdk.VPC().Gateway().List(ctx, &vpc.ListGatewaysRequest{
		FolderId: r.FolderID(),
		Filter:   CreateResolverFilter("name", r.Name),
		PageSize: DefaultResolverPageSize,
	}, opts...)
	return r.findName(resp.GetGateways(), err)
}
