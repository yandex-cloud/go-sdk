package ycsdk

import "github.com/yandex-cloud/go-sdk/gen/kubernetes/marketplace"

func (sdk *SDK) KubernetesMarketplace() *marketplace.KubernetesMarketplace {
	return marketplace.NewKubernetesMarketplace(sdk.getConn(KubernetesServiceID))
}
