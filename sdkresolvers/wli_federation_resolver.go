package sdkresolvers

import (
	"context"

	"google.golang.org/grpc"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/iam/v1/workload/oidc"
	ycsdk "github.com/yandex-cloud/go-sdk"
)

type wliFederationResolver struct {
	BaseNameResolver
}

func WliFederationResolver(name string, opts ...ResolveOption) ycsdk.Resolver {
	return &wliFederationResolver{
		BaseNameResolver: NewBaseNameResolver(name, "wli federation", opts...),
	}
}

func (r *wliFederationResolver) Run(ctx context.Context, sdk *ycsdk.SDK, opts ...grpc.CallOption) error {
	err := r.ensureFolderID()
	if err != nil {
		return err
	}
	res := []*oidc.Federation{}
	nextPageToken := ""
	for {
		resp, err := sdk.WorkloadOidc().Federation().List(ctx, &oidc.ListFederationsRequest{
			FolderId:  r.FolderID(),
			PageSize:  DefaultResolverPageSize,
			PageToken: nextPageToken,
		}, opts...)
		if err != nil {
			return err
		}
		nextPageToken = resp.GetNextPageToken()
		res = append(res, resp.GetFederations()...)

		if len(nextPageToken) == 0 {
			break
		}
	}
	return r.findName(res, err)
}
