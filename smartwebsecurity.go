package ycsdk

import (
	"github.com/yandex-cloud/go-sdk/gen/smartwebsecurity"
)

const (
	SmartWebSecurityServiceID Endpoint = "smart-web-security"
)

// SmartWebSecurity returns SmartWebSecurity object that is used to operate with security profiles
func (sdk *SDK) SmartWebSecurity() *smartwebsecurity.SmartWebSecurity {
	return smartwebsecurity.NewSmartWebSecurity(sdk.getConn(SmartWebSecurityServiceID))
}
