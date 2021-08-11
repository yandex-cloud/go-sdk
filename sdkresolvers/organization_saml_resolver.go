// Copyright (c) 2019 YANDEX LLC.

package sdkresolvers

import (
	"context"

	"google.golang.org/grpc"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/organizationmanager/v1/saml"
	ycsdk "github.com/yandex-cloud/go-sdk"
)

type organizationSamlFederationResolver struct {
	BaseNameResolver
}

func OrganizationSamlFederationResolver(name string, opts ...ResolveOption) ycsdk.Resolver {
	return &organizationSamlFederationResolver{
		BaseNameResolver: NewBaseNameResolver(name, "federation", opts...),
	}
}

func (r *organizationSamlFederationResolver) Run(ctx context.Context, sdk *ycsdk.SDK, opts ...grpc.CallOption) error {
	err := r.ensureOrganizationID()
	if err != nil {
		return err
	}

	resp, err := sdk.OrganizationManagerSAML().Federation().List(ctx, &saml.ListFederationsRequest{
		OrganizationId: r.OrganizationID(),
		Filter:         CreateResolverFilter("name", r.Name),
		PageSize:       DefaultResolverPageSize,
	}, opts...)
	return r.findName(resp.GetFederations(), err)
}

type organizationSamlCertificateResolver struct {
	BaseNameResolver
}

func OrganizationSamlCertificateResolver(name string, opts ...ResolveOption) ycsdk.Resolver {
	return &organizationSamlCertificateResolver{
		BaseNameResolver: NewBaseNameResolver(name, "certificate", opts...),
	}
}

func (r *organizationSamlCertificateResolver) Run(ctx context.Context, sdk *ycsdk.SDK, opts ...grpc.CallOption) error {
	resp, err := sdk.OrganizationManagerSAML().Certificate().List(ctx, &saml.ListCertificatesRequest{
		FederationId: r.opts.federationID,
		Filter:       CreateResolverFilter("name", r.Name),
		PageSize:     DefaultResolverPageSize,
	}, opts...)
	return r.findName(resp.GetCertificates(), err)
}
