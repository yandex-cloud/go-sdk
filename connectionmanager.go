package ycsdk

import (
	"github.com/yandex-cloud/go-sdk/gen/connectionmanager"
)

const (
	ConnectionManagerID Endpoint = "connection-manager"
)

func (sdk *SDK) ConnectionManager() *connectionmanager.ConnectionManager {
	return connectionmanager.NewConnectionManager(sdk.getConn(ConnectionManagerID))
}
