// Copyright (c) 2025 Yandex LLC. All rights reserved.
// Author: Dmitry Rusanov <dmitryrusanov@yandex-team.ru>

package ycsdk

import (
	"github.com/yandex-cloud/go-sdk/gen/quotamanager"
)

const (
	QuotaManagementServiceID Endpoint = "quota-manager"
)

// QuotaManager returns QuotaManager object that is used to operate on QuotaLimits
func (sdk *SDK) QuotaManager() *quotamanager.QuotaManager {
	return quotamanager.NewQuotaManager(sdk.getConn(QuotaManagementServiceID))
}
