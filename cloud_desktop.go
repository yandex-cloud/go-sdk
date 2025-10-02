package ycsdk

import (
	"github.com/yandex-cloud/go-sdk/gen/clouddesktop"
)

const (
	VDIServiceID Endpoint = "clouddesktops"
)

func (sdk *SDK) CloudDesktop() *clouddesktop.CloudDesktop {
	return clouddesktop.NewCloudDesktop(sdk.getConn(VDIServiceID))
}
