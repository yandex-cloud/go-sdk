package sdkresolvers

import (
	"context"

	"google.golang.org/grpc"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/smartwebsecurity/v1"
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
		// TODO: better to use Filter("name"), but now it's not supported now
	}, opts...)
	return r.findName(resp.GetSecurityProfiles(), err)
}
