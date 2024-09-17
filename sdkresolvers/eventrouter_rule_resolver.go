package sdkresolvers

import (
	"context"

	"google.golang.org/grpc"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/serverless/eventrouter/v1"
	ycsdk "github.com/yandex-cloud/go-sdk"
)

type eventrouterRuleResolver struct {
	BaseNameResolver
}

func EventrouterRuleResolver(name string, opts ...ResolveOption) ycsdk.Resolver {
	return &eventrouterRuleResolver{
		BaseNameResolver: NewBaseNameResolver(name, "rule", opts...),
	}
}

func (r *eventrouterRuleResolver) Run(ctx context.Context, sdk *ycsdk.SDK, opts ...grpc.CallOption) error {
	if err := r.ensureFolderID(); err != nil {
		return err
	}
	request := &eventrouter.ListRulesRequest{
		ContainerId: &eventrouter.ListRulesRequest_FolderId{FolderId: r.FolderID()},
		Filter:      CreateResolverFilter("name", r.Name),
		PageSize:    DefaultResolverPageSize,
	}
	resp, err := sdk.Serverless().Eventrouter().Rule().List(ctx, request, opts...)
	return r.findName(resp.GetRules(), err)
}
