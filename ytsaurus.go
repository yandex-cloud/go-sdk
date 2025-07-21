package ycsdk

import (
	"github.com/yandex-cloud/go-sdk/gen/ytsaurus"
)

const (
	YTsaurusServiceID Endpoint = "managed-ytsaurus"
)

func (sdk *SDK) YTsaurus() *ytsaurus.YTsaurus {
	return ytsaurus.NewYTsaurus(sdk.getConn(YTsaurusServiceID))
}
