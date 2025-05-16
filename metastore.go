package ycsdk

import (
	"github.com/yandex-cloud/go-sdk/gen/metastore"
)

const MetastoreServiceID = "managed-metastore"

func (sdk *SDK) Metastore() *metastore.Metastore {
	return metastore.NewMetastore(sdk.getConn(MetastoreServiceID))
}
