// Code generated by sdkgen. DO NOT EDIT.

package organizationmanager

import (
	"context"

	"google.golang.org/grpc"
)

// OrganizationManager provides access to "organizationmanager" component of Yandex.Cloud
type OrganizationManager struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// NewOrganizationManager creates instance of OrganizationManager
func NewOrganizationManager(g func(ctx context.Context) (*grpc.ClientConn, error)) *OrganizationManager {
	return &OrganizationManager{g}
}

// Organization gets OrganizationService client
func (o *OrganizationManager) Organization() *OrganizationServiceClient {
	return &OrganizationServiceClient{getConn: o.getConn}
}

// User gets UserService client
func (o *OrganizationManager) User() *UserServiceClient {
	return &UserServiceClient{getConn: o.getConn}
}

// Group gets GroupService client
func (o *OrganizationManager) Group() *GroupServiceClient {
	return &GroupServiceClient{getConn: o.getConn}
}

// GroupMapping gets GroupMappingService client
func (o *OrganizationManager) GroupMapping() *GroupMappingServiceClient {
	return &GroupMappingServiceClient{getConn: o.getConn}
}

// OsLogin gets OsLoginService client
func (o *OrganizationManager) OsLogin() *OsLoginServiceClient {
	return &OsLoginServiceClient{getConn: o.getConn}
}
