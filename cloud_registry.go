package ycsdk

import (
	"github.com/yandex-cloud/go-sdk/gen/cloudregistry"
)

const (
	CloudRegistryServiceID Endpoint = "cloud-registry"
)

func (sdk *SDK) CloudRegistry() *cloudregistry.CloudRegistry {
	return cloudregistry.NewCloudRegistry(sdk.getConn(CloudRegistryServiceID))
}
