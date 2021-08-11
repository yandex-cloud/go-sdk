package sdkresolvers

import (
	"context"

	"google.golang.org/grpc"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/organizationmanager/v1"
	ycsdk "github.com/yandex-cloud/go-sdk"
)

type organizationResolver struct {
	BaseNameResolver
}

func OrganizationResolver(name string, opts ...ResolveOption) ycsdk.Resolver {
	return &organizationResolver{
		BaseNameResolver: NewBaseNameResolver(name, "organization", opts...),
	}
}

func (r *organizationResolver) Run(ctx context.Context, sdk *ycsdk.SDK, opts ...grpc.CallOption) error {
	resp, err := sdk.OrganizationManager().Organization().List(ctx, &organizationmanager.ListOrganizationsRequest{
		Filter:   CreateResolverFilter("name", r.Name),
		PageSize: DefaultResolverPageSize,
	}, opts...)
	return r.findName(resp.GetOrganizations(), err)
}
