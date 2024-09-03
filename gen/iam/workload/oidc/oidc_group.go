// Code generated by sdkgen. DO NOT EDIT.

package oidc

import (
	"context"

	"google.golang.org/grpc"
)

// WorkloadOidc provides access to "oidc" component of Yandex.Cloud
type WorkloadOidc struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// NewWorkloadOidc creates instance of WorkloadOidc
func NewWorkloadOidc(g func(ctx context.Context) (*grpc.ClientConn, error)) *WorkloadOidc {
	return &WorkloadOidc{g}
}

// Federation gets FederationService client
func (w *WorkloadOidc) Federation() *FederationServiceClient {
	return &FederationServiceClient{getConn: w.getConn}
}