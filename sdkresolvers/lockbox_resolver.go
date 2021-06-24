// Copyright (c) 2021 Yandex LLC. All rights reserved.
// Author: Andrey Zamyslov <zamysel@yandex-team.ru>

package sdkresolvers

import (
	"context"

	"google.golang.org/grpc"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/lockbox/v1"
	ycsdk "github.com/yandex-cloud/go-sdk"
)

type secretResolver struct {
	BaseNameResolver
}

func SecretResolver(name string, opts ...ResolveOption) ycsdk.Resolver {
	return &secretResolver{
		BaseNameResolver: NewBaseNameResolver(name, "secret", opts...),
	}
}

func (r *secretResolver) Run(ctx context.Context, sdk *ycsdk.SDK, opts ...grpc.CallOption) error {
	err := r.ensureFolderID()
	if err != nil {
		return err
	}
	res := []*lockbox.Secret{}
	nextPageToken := ""
	for ok := true; ok; ok = len(nextPageToken) > 0 {
		resp, err := sdk.LockboxSecret().Secret().List(ctx, &lockbox.ListSecretsRequest{
			FolderId:  r.FolderID(),
			PageSize:  DefaultResolverPageSize,
			PageToken: nextPageToken,
		}, opts...)
		if err != nil {
			return err
		}
		nextPageToken = resp.GetNextPageToken()
		res = append(res, resp.GetSecrets()...)
	}
	return r.findName(res, err)
}
