package ycsdk

import (
	"github.com/yandex-cloud/go-sdk/gen/cloudrouter"
)

const (
	CloudRouterServiceID Endpoint = "cloud-router"
)

func (sdk *SDK) CloudRouter() *cloudrouter.CloudRouter {
	return cloudrouter.NewCloudRouter(sdk.getConn(CloudRegistryServiceID))
}
