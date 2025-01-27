// Copyright (c) 2019 Yandex LLC. All rights reserved.
// Author: Dmitry Novikov <novikoff@yandex-team.ru>

package ycsdk

import (
	"github.com/yandex-cloud/go-sdk/gen/marketplace/licensemanager"
	"github.com/yandex-cloud/go-sdk/gen/marketplace/metering"
	pim "github.com/yandex-cloud/go-sdk/gen/marketplace/pim"
)

const (
	MarketplaceServiceID               Endpoint = "marketplace"
	MarketplaceMeteringServiceID                = MarketplaceServiceID
	MarketplaceLicenseManagerServiceID          = MarketplaceServiceID
	MarketplacePIMServiceID                     = MarketplaceServiceID
)

type Marketplace struct {
	sdk *SDK
}

func (m *Marketplace) Metering() *metering.Metering {
	return metering.NewMetering(m.sdk.getConn(MarketplaceMeteringServiceID))
}

func (m *Marketplace) LicenseManager() *licensemanager.LicenseManager {
	return licensemanager.NewLicenseManager(m.sdk.getConn(MarketplaceLicenseManagerServiceID))
}

func (m *Marketplace) PIM() *pim.PIM {
	return pim.NewPIM(m.sdk.getConn(MarketplacePIMServiceID))
}
