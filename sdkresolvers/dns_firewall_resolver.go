package sdkresolvers

import (
	"context"

	"google.golang.org/grpc"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/dns/v1"
	ycsdk "github.com/yandex-cloud/go-sdk"
)

type dnsFirewallResolver struct {
	BaseNameResolver
}

func DNSFirewallResolver(name string, opts ...ResolveOption) ycsdk.Resolver {
	return &dnsFirewallResolver{
		BaseNameResolver: NewBaseNameResolver(name, "dns firewall", opts...),
	}
}
func (r *dnsFirewallResolver) Run(ctx context.Context, sdk *ycsdk.SDK, opts ...grpc.CallOption) error {
	err := r.ensureFolderID()
	if err != nil {
		return err
	}
	resp, err := sdk.DNS().DnsFirewall().List(ctx, &dns.ListDnsFirewallsRequest{
		FolderId: r.FolderID(),
		Filter:   CreateResolverFilter("name", r.Name),
		PageSize: DefaultResolverPageSize,
	}, opts...)
	return r.findName(resp.GetDnsFirewalls(), err)
}
