package sdkresolvers

import (
	"context"

	"google.golang.org/grpc"

	smartwebsecurity "github.com/yandex-cloud/go-genproto/yandex/cloud/smartwebsecurity/v1"
	advanced_rate_limiter "github.com/yandex-cloud/go-genproto/yandex/cloud/smartwebsecurity/v1/advanced_rate_limiter"
	waf "github.com/yandex-cloud/go-genproto/yandex/cloud/smartwebsecurity/v1/waf"
	ycsdk "github.com/yandex-cloud/go-sdk"
)

type securityProfileResolver struct {
	BaseNameResolver
}

func SecurityProfileResolver(name string, opts ...ResolveOption) ycsdk.Resolver {
	return &securityProfileResolver{
		BaseNameResolver: NewBaseNameResolver(name, "security profile", opts...),
	}
}

func (r *securityProfileResolver) Run(ctx context.Context, sdk *ycsdk.SDK, opts ...grpc.CallOption) error {
	err := r.ensureFolderID()
	if err != nil {
		return err
	}

	resp, err := sdk.SmartWebSecurity().SecurityProfile().List(ctx, &smartwebsecurity.ListSecurityProfilesRequest{
		FolderId: r.FolderID(),
		// TODO: better to use Filter("name"), but it's not supported now
	}, opts...)
	return r.findName(resp.GetSecurityProfiles(), err)
}

type swsWafProfileResolver struct {
	BaseNameResolver
}

func SWSWafProfileResolver(name string, opts ...ResolveOption) ycsdk.Resolver {
	return &swsWafProfileResolver{
		BaseNameResolver: NewBaseNameResolver(name, "waf profile", opts...),
	}
}

func (r *swsWafProfileResolver) Run(ctx context.Context, sdk *ycsdk.SDK, opts ...grpc.CallOption) error {
	err := r.ensureFolderID()
	if err != nil {
		return err
	}

	resp, err := sdk.SmartWebSecurityWaf().WafProfile().List(ctx, &waf.ListWafProfilesRequest{
		FolderId: r.FolderID(),
		// TODO: better to use Filter("name"), but it's not supported now
	}, opts...)
	return r.findName(resp.GetWafProfiles(), err)
}

type swsAdvancedRateLimiterProfileResolver struct {
	BaseNameResolver
}

func SWSAdvancedRateLimiterProfileResolver(name string, opts ...ResolveOption) ycsdk.Resolver {
	return &swsAdvancedRateLimiterProfileResolver{
		BaseNameResolver: NewBaseNameResolver(name, "advanced rate limiter profile", opts...),
	}
}

func (r *swsAdvancedRateLimiterProfileResolver) Run(ctx context.Context, sdk *ycsdk.SDK, opts ...grpc.CallOption) error {
	err := r.ensureFolderID()
	if err != nil {
		return err
	}

	resp, err := sdk.SmartWebSecurityArl().AdvancedRateLimiterProfile().List(ctx, &advanced_rate_limiter.ListAdvancedRateLimiterProfilesRequest{
		FolderId: r.FolderID(),
		// TODO: better to use Filter("name"), but it's not supported now
	}, opts...)
	return r.findName(resp.GetAdvancedRateLimiterProfiles(), err)
}
