package ycsdk

import (
	"github.com/yandex-cloud/go-sdk/gen/trino"
)

const TrinoServiceID = "managed-trino"

func (sdk *SDK) Trino() *trino.Trino {
	return trino.NewTrino(sdk.getConn(TrinoServiceID))
}
