// Copyright (c) 2022 YANDEX LLC.

package sdkresolvers

import (
	"context"

	"google.golang.org/grpc"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/organizationmanager/v1"
	ycsdk "github.com/yandex-cloud/go-sdk"
)

type organizationGroupResolver struct {
	BaseNameResolver
}

func OrganizationGroupResolver(name string, opts ...ResolveOption) ycsdk.Resolver {
	return &organizationGroupResolver{
		BaseNameResolver: NewBaseNameResolver(name, "group", opts...),
	}
}

func (r *organizationGroupResolver) Run(ctx context.Context, sdk *ycsdk.SDK, opts ...grpc.CallOption) error {
	err := r.ensureOrganizationID()
	if err != nil {
		return err
	}

	resp, err := sdk.OrganizationManager().Group().List(ctx, &organizationmanager.ListGroupsRequest{
		OrganizationId: r.OrganizationID(),
		Filter:         CreateResolverFilter("name", r.Name),
		PageSize:       DefaultResolverPageSize,
	}, opts...)
	return r.findName(resp.GetGroups(), err)
}
