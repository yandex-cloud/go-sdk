package sdkresolvers

import (
	"context"

	"google.golang.org/grpc"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/dns/v1"
	ycsdk "github.com/yandex-cloud/go-sdk"
)

type dnsZoneResolver struct {
	BaseNameResolver
}

func DNSZoneResolver(name string, opts ...ResolveOption) ycsdk.Resolver {
	return &dnsZoneResolver{
		BaseNameResolver: NewBaseNameResolver(name, "dns zone", opts...),
	}
}

func (r *dnsZoneResolver) Run(ctx context.Context, sdk *ycsdk.SDK, opts ...grpc.CallOption) error {
	err := r.ensureFolderID()
	if err != nil {
		return err
	}

	resp, err := sdk.DNS().DnsZone().List(ctx, &dns.ListDnsZonesRequest{
		FolderId: r.FolderID(),
		Filter:   CreateResolverFilter("name", r.Name),
		PageSize: DefaultResolverPageSize,
	}, opts...)
	return r.findName(resp.GetDnsZones(), err)
}
