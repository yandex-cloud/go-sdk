// Code generated by sdkgen. DO NOT EDIT.

package waf

import (
	"context"

	"google.golang.org/grpc"
)

// SmartWebSecurityWaf provides access to "waf" component of Yandex.Cloud
type SmartWebSecurityWaf struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// NewSmartWebSecurityWaf creates instance of SmartWebSecurityWaf
func NewSmartWebSecurityWaf(g func(ctx context.Context) (*grpc.ClientConn, error)) *SmartWebSecurityWaf {
	return &SmartWebSecurityWaf{g}
}

// WafProfile gets WafProfileService client
func (s *SmartWebSecurityWaf) WafProfile() *WafProfileServiceClient {
	return &WafProfileServiceClient{getConn: s.getConn}
}

// RuleSetDescriptor gets RuleSetDescriptorService client
func (s *SmartWebSecurityWaf) RuleSetDescriptor() *RuleSetDescriptorServiceClient {
	return &RuleSetDescriptorServiceClient{getConn: s.getConn}
}
