// Copyright (c) 2023 Yandex LLC. All rights reserved.
// Author: Ratbek Nurlanbekuulu <ratbek@yandex-team.ru>

package sdkresolvers

import (
	"context"

	"google.golang.org/grpc"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/datasphere/v2"
	ycsdk "github.com/yandex-cloud/go-sdk"
)

type communityResolver struct {
	BaseNameResolver
}

func NewDatasphereCommunityResolver(name string, opts ...ResolveOption) ycsdk.Resolver {
	return &communityResolver{
		BaseNameResolver: NewBaseNameResolver(name, "community", opts...),
	}
}

func (r *communityResolver) Run(ctx context.Context, sdk *ycsdk.SDK, opts ...grpc.CallOption) error {
	err := r.ensureOrganizationID()
	if err != nil {
		return err
	}

	resp, err := sdk.Datasphere().Community().List(ctx, &datasphere.ListCommunitiesRequest{
		OrganizationId: r.OrganizationID(),
		PageSize:       DefaultResolverPageSize,
	}, opts...)

	return r.findName(resp.GetCommunities(), err)
}

type projectResolver struct {
	BaseNameResolver
}

func NewDatasphereProjectResolver(name string, opts ...ResolveOption) ycsdk.Resolver {
	return &projectResolver{
		BaseNameResolver: NewBaseNameResolver(name, "project", opts...),
	}
}

func (r *projectResolver) Run(ctx context.Context, sdk *ycsdk.SDK, opts ...grpc.CallOption) error {
	err := r.ensureCommunityID()
	if err != nil {
		return err
	}

	resp, err := sdk.Datasphere().Project().List(ctx, &datasphere.ListProjectsRequest{
		CommunityId: r.CommunityID(),
		PageSize:    DefaultResolverPageSize,
	}, opts...)

	return r.findName(resp.GetProjects(), err)
}
