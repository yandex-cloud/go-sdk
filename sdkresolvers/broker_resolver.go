package sdkresolvers

import (
	"context"

	"google.golang.org/grpc"

	iot "github.com/yandex-cloud/go-genproto/yandex/cloud/iot/broker/v1"
	ycsdk "github.com/yandex-cloud/go-sdk"
)

type brokerResolver struct {
	BaseNameResolver
}

func BrokerResolver(name string, opts ...ResolveOption) ycsdk.Resolver {
	return &brokerResolver{
		BaseNameResolver: NewBaseNameResolver(name, "broker", opts...),
	}
}

func (r *brokerResolver) Run(ctx context.Context, sdk *ycsdk.SDK, opts ...grpc.CallOption) error {
	err := r.ensureFolderID()
	if err != nil {
		return err
	}

	nextPageToken := ""
	var brks []*iot.Broker
	for ok := true; ok; ok = len(nextPageToken) > 0 {
		resp, err := sdk.IoT().Broker().Broker().List(ctx, &iot.ListBrokersRequest{
			FolderId:  r.FolderID(),
			PageSize:  DefaultResolverPageSize,
			PageToken: nextPageToken,
		}, opts...)
		if err != nil {
			return err
		}
		nextPageToken = resp.GetNextPageToken()
		brks = append(brks, resp.GetBrokers()...)
	}

	return r.findName(brks, err)
}
