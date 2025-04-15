package ycsdk

import (
	"github.com/yandex-cloud/go-sdk/gen/spark"
)

const SparkServiceID = "managed-spark"

func (sdk *SDK) Spark() *spark.Spark {
	return spark.NewSpark(sdk.getConn(SparkServiceID))
}
