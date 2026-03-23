package ycsdk

import (
	"github.com/yandex-cloud/go-sdk/gen/datacatalog"
)

const (
	DatacatalogID Endpoint = "datacatalog"
)

func (sdk *SDK) Datacatalog() *datacatalog.Datacatalog {
	return datacatalog.NewDatacatalog(sdk.getConn(DatacatalogID))
}
