// Code generated by sdkgen. DO NOT EDIT.

package postgresql

import (
	"context"

	"google.golang.org/grpc"
)

// PostgreSQL provides access to "postgresql" component of Yandex.Cloud
type PostgreSQL struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// NewPostgreSQL creates instance of PostgreSQL
func NewPostgreSQL(g func(ctx context.Context) (*grpc.ClientConn, error)) *PostgreSQL {
	return &PostgreSQL{g}
}

// Backup gets BackupService client
func (p *PostgreSQL) Backup() *BackupServiceClient {
	return &BackupServiceClient{getConn: p.getConn}
}

// BackupRetentionPolicy gets BackupRetentionPolicyService client
func (p *PostgreSQL) BackupRetentionPolicy() *BackupRetentionPolicyServiceClient {
	return &BackupRetentionPolicyServiceClient{getConn: p.getConn}
}

// Cluster gets ClusterService client
func (p *PostgreSQL) Cluster() *ClusterServiceClient {
	return &ClusterServiceClient{getConn: p.getConn}
}

// Database gets DatabaseService client
func (p *PostgreSQL) Database() *DatabaseServiceClient {
	return &DatabaseServiceClient{getConn: p.getConn}
}

// ResourcePreset gets ResourcePresetService client
func (p *PostgreSQL) ResourcePreset() *ResourcePresetServiceClient {
	return &ResourcePresetServiceClient{getConn: p.getConn}
}

// User gets UserService client
func (p *PostgreSQL) User() *UserServiceClient {
	return &UserServiceClient{getConn: p.getConn}
}
