package ycsdk

import "github.com/yandex-cloud/go-sdk/gen/baremetal"

const (
	BaremetalServiceID Endpoint = "baremetal"
)

func (sdk *SDK) Baremetal() *baremetal.Baremetal {
	return baremetal.NewBaremetal(sdk.getConn(BaremetalServiceID))
}
