package ycsdk

import (
	"github.com/yandex-cloud/go-sdk/gen/cloudregistry"
	cloudregistry_domain "github.com/yandex-cloud/go-sdk/gen/cloudregistry/domain"
)

const (
	CloudRegistryServiceID Endpoint = "cloud-registry"
)

func (sdk *SDK) CloudRegistry() *cloudregistry.CloudRegistry {
	return cloudregistry.NewCloudRegistry(sdk.getConn(CloudRegistryServiceID))
}

func (sdk *SDK) CloudRegistryDomain() *cloudregistry_domain.Domain {
	return cloudregistry_domain.NewDomain(sdk.getConn(CloudRegistryServiceID))
}
