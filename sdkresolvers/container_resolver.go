// Copyright (c) 2018 Yandex LLC. All rights reserved.
// Author: Dmitry Novikov <novikoff@yandex-team.ru>

package sdkresolvers

import (
	"context"

	"google.golang.org/grpc"

	containers "github.com/yandex-cloud/go-genproto/yandex/cloud/serverless/containers/v1"
	ycsdk "github.com/yandex-cloud/go-sdk"
)

type containerResolver struct {
	BaseNameResolver
}

func ContainerResolver(name string, opts ...ResolveOption) ycsdk.Resolver {
	return &containerResolver{
		BaseNameResolver: NewBaseNameResolver(name, "container", opts...),
	}
}

func (r *containerResolver) Run(ctx context.Context, sdk *ycsdk.SDK, opts ...grpc.CallOption) error {
	err := r.ensureFolderID()
	if err != nil {
		return err
	}

	resp, err := sdk.Serverless().Containers().Container().List(ctx, &containers.ListContainersRequest{
		FolderId: r.FolderID(),
		Filter:   CreateResolverFilter("name", r.Name),
		PageSize: DefaultResolverPageSize,
	}, opts...)
	return r.findName(resp.GetContainers(), err)
}
