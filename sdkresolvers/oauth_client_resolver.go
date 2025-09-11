package sdkresolvers

import (
	"context"

	"google.golang.org/grpc"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/iam/v1"
	ycsdk "github.com/yandex-cloud/go-sdk"
)

type oauthclientResolver struct {
	BaseNameResolver
}

func OAuthClientResolver(name string, opts ...ResolveOption) ycsdk.Resolver {
	return &oauthclientResolver{
		BaseNameResolver: NewBaseNameResolver(name, "oauth-client", opts...),
	}
}

func (r *oauthclientResolver) Run(ctx context.Context, sdk *ycsdk.SDK, opts ...grpc.CallOption) error {
	err := r.ensureFolderID()
	if err != nil {
		return err
	}

	var clients []*iam.OAuthClientListView
	var pageToken string
	for {
		resp, err := sdk.IAM().OAuthClient().List(ctx, &iam.ListOAuthClientsRequest{
			FolderId:  r.FolderID(),
			PageSize:  DefaultResolverPageSize,
			PageToken: pageToken,
		}, opts...)
		if err != nil {
			return err
		}
		clients = append(clients, resp.GetOauthClients()...)
		if resp.GetNextPageToken() == "" {
			break
		}
		pageToken = resp.GetNextPageToken()
	}

	return r.findName(clients, err)
}
