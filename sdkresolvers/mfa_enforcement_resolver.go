package sdkresolvers

import (
	"context"

	"google.golang.org/grpc"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/organizationmanager/v1"
	ycsdk "github.com/yandex-cloud/go-sdk"
)

type mfaEnforcementResolver struct {
	BaseNameResolver
}

func MfaEnforcementResolver(name string, opts ...ResolveOption) ycsdk.Resolver {
	return &mfaEnforcementResolver{
		BaseNameResolver: NewBaseNameResolver(name, "mfa-enforcement", opts...),
	}
}

func (r *mfaEnforcementResolver) Run(ctx context.Context, sdk *ycsdk.SDK, opts ...grpc.CallOption) error {
	err := r.ensureOrganizationID()
	if err != nil {
		return err
	}
	var enforcements []*organizationmanager.MfaEnforcement
	var pageToken string
	for {
		resp, err := sdk.OrganizationManager().MfaEnforcement().List(ctx, &organizationmanager.ListMfaEnforcementsRequest{
			OrganizationId: r.OrganizationID(),
			PageSize:       DefaultResolverPageSize,
			PageToken:      pageToken,
		}, opts...)
		if err != nil {
			return err
		}
		enforcements = append(enforcements, resp.GetMfaEnforcements()...)
		if resp.GetNextPageToken() == "" {
			break
		}
		pageToken = resp.GetNextPageToken()
	}
	return r.findName(enforcements, err)
}
