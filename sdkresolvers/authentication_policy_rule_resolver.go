package sdkresolvers

import (
	"context"

	"google.golang.org/grpc"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/organizationmanager/v1/policy"
	ycsdk "github.com/yandex-cloud/go-sdk"
)

type authenticationPolicyRuleResolver struct {
	BaseNameResolver
}

func AuthenticationPolicyRuleResolver(name string, opts ...ResolveOption) ycsdk.Resolver {
	return &authenticationPolicyRuleResolver{
		BaseNameResolver: NewBaseNameResolver(name, "authentication-policy-rule", opts...),
	}
}

func (r *authenticationPolicyRuleResolver) Run(ctx context.Context, sdk *ycsdk.SDK, opts ...grpc.CallOption) error {
	err := r.ensureOrganizationID()
	if err != nil {
		return err
	}

	resp, err := sdk.OrganizationManagerPolicy().AuthenticationPolicyRule().List(ctx, &policy.ListRulesRequest{
		OrganizationId: r.OrganizationID(),
		Filter:         CreateResolverFilter("name", r.Name),
		PageSize:       DefaultResolverPageSize,
	}, opts...)
	return r.findName(resp.GetAuthPolicyRules(), err)
}
