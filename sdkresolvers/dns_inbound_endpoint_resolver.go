package sdkresolvers

import (
	"context"

	"google.golang.org/grpc"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/dns/v1"
	ycsdk "github.com/yandex-cloud/go-sdk"
)

type dnsInboundEndpointResolver struct {
	BaseNameResolver
}

func DNSInboundEndpointResolver(name string, opts ...ResolveOption) ycsdk.Resolver {
	return &dnsInboundEndpointResolver{
		BaseNameResolver: NewBaseNameResolver(name, "dns inbound endpoint", opts...),
	}
}

func (r *dnsInboundEndpointResolver) Run(ctx context.Context, sdk *ycsdk.SDK, opts ...grpc.CallOption) error {
	err := r.ensureFolderID()
	if err != nil {
		return err
	}
	resp, err := sdk.DNS().DnsInboundEndpoint().List(ctx, &dns.ListDnsInboundEndpointsRequest{
		FolderId: r.FolderID(),
		Filter:   CreateResolverFilter("name", r.Name),
		PageSize: DefaultResolverPageSize,
	}, opts...)
	return r.findName(resp.GetEndpoints(), err)
}
