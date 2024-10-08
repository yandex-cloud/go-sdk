// Code generated by sdkgen. DO NOT EDIT.

package advanced_rate_limiter

import (
	"context"

	"google.golang.org/grpc"
)

// SmartWebSecurityArl provides access to "advanced_rate_limiter" component of Yandex.Cloud
type SmartWebSecurityArl struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// NewSmartWebSecurityArl creates instance of SmartWebSecurityArl
func NewSmartWebSecurityArl(g func(ctx context.Context) (*grpc.ClientConn, error)) *SmartWebSecurityArl {
	return &SmartWebSecurityArl{g}
}

// AdvancedRateLimiterProfile gets AdvancedRateLimiterProfileService client
func (s *SmartWebSecurityArl) AdvancedRateLimiterProfile() *AdvancedRateLimiterProfileServiceClient {
	return &AdvancedRateLimiterProfileServiceClient{getConn: s.getConn}
}
