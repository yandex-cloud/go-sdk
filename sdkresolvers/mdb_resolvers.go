package sdkresolvers

import (
	"context"

	"google.golang.org/grpc"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/clickhouse/v1"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/elasticsearch/v1"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/greenplum/v1"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/kafka/v1"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/mongodb/v1"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/mysql/v1"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/opensearch/v1"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/postgresql/v1"
	redis "github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/redis/v1"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/sqlserver/v1"
	ycsdk "github.com/yandex-cloud/go-sdk"
)

func PostgreSQLClusterResolver(name string, opts ...ResolveOption) ycsdk.Resolver {
	return &postgreSQLClusterResolver{
		BaseNameResolver: NewBaseNameResolver(name, "cluster", opts...),
	}
}

type postgreSQLClusterResolver struct {
	BaseNameResolver
}

func (r *postgreSQLClusterResolver) Run(ctx context.Context, sdk *ycsdk.SDK, opts ...grpc.CallOption) error {
	err := r.ensureFolderID()
	if err != nil {
		return err
	}

	resp, err := sdk.MDB().PostgreSQL().Cluster().List(ctx, &postgresql.ListClustersRequest{
		FolderId: r.FolderID(),
		Filter:   CreateResolverFilter("name", r.Name),
		PageSize: DefaultResolverPageSize,
	})
	return r.findName(resp.GetClusters(), err)
}

func MongoDBClusterResolver(name string, opts ...ResolveOption) ycsdk.Resolver {
	return &mongodbClusterResolver{
		BaseNameResolver: NewBaseNameResolver(name, "cluster", opts...),
	}
}

type mongodbClusterResolver struct {
	BaseNameResolver
}

func (r *mongodbClusterResolver) Run(ctx context.Context, sdk *ycsdk.SDK, opts ...grpc.CallOption) error {
	err := r.ensureFolderID()
	if err != nil {
		return err
	}

	resp, err := sdk.MDB().MongoDB().Cluster().List(ctx, &mongodb.ListClustersRequest{
		FolderId: r.FolderID(),
		Filter:   CreateResolverFilter("name", r.Name),
		PageSize: DefaultResolverPageSize,
	})
	return r.findName(resp.GetClusters(), err)
}

func ClickhouseClusterResolver(name string, opts ...ResolveOption) ycsdk.Resolver {
	return &clickhouseClusterResolver{
		BaseNameResolver: NewBaseNameResolver(name, "cluster", opts...),
	}
}

type clickhouseClusterResolver struct {
	BaseNameResolver
}

func (r *clickhouseClusterResolver) Run(ctx context.Context, sdk *ycsdk.SDK, opts ...grpc.CallOption) error {
	err := r.ensureFolderID()
	if err != nil {
		return err
	}

	resp, err := sdk.MDB().Clickhouse().Cluster().List(ctx, &clickhouse.ListClustersRequest{
		FolderId: r.FolderID(),
		Filter:   CreateResolverFilter("name", r.Name),
		PageSize: DefaultResolverPageSize,
	})
	return r.findName(resp.GetClusters(), err)
}

func RedisClusterResolver(name string, opts ...ResolveOption) ycsdk.Resolver {
	return &redisClusterResolver{
		BaseNameResolver: NewBaseNameResolver(name, "cluster", opts...),
	}
}

type redisClusterResolver struct {
	BaseNameResolver
}

func (r *redisClusterResolver) Run(ctx context.Context, sdk *ycsdk.SDK, opts ...grpc.CallOption) error {
	err := r.ensureFolderID()
	if err != nil {
		return err
	}

	resp, err := sdk.MDB().Redis().Cluster().List(ctx, &redis.ListClustersRequest{
		FolderId: r.FolderID(),
		Filter:   CreateResolverFilter("name", r.Name),
		PageSize: DefaultResolverPageSize,
	})
	return r.findName(resp.GetClusters(), err)
}

func MySQLClusterResolver(name string, opts ...ResolveOption) ycsdk.Resolver {
	return &mySQLClusterResolver{
		BaseNameResolver: NewBaseNameResolver(name, "cluster", opts...),
	}
}

type mySQLClusterResolver struct {
	BaseNameResolver
}

func (r *mySQLClusterResolver) Run(ctx context.Context, sdk *ycsdk.SDK, opts ...grpc.CallOption) error {
	err := r.ensureFolderID()
	if err != nil {
		return err
	}

	resp, err := sdk.MDB().MySQL().Cluster().List(ctx, &mysql.ListClustersRequest{
		FolderId: r.FolderID(),
		Filter:   CreateResolverFilter("name", r.Name),
		PageSize: DefaultResolverPageSize,
	})
	return r.findName(resp.GetClusters(), err)
}

func SQLServerClusterResolver(name string, opts ...ResolveOption) ycsdk.Resolver {
	return &sqlServerClusterResolver{
		BaseNameResolver: NewBaseNameResolver(name, "cluster", opts...),
	}
}

type sqlServerClusterResolver struct {
	BaseNameResolver
}

func (r *sqlServerClusterResolver) Run(ctx context.Context, sdk *ycsdk.SDK, opts ...grpc.CallOption) error {
	err := r.ensureFolderID()
	if err != nil {
		return err
	}

	resp, err := sdk.MDB().SQLServer().Cluster().List(ctx, &sqlserver.ListClustersRequest{
		FolderId: r.FolderID(),
		Filter:   CreateResolverFilter("name", r.Name),
		PageSize: DefaultResolverPageSize,
	})
	return r.findName(resp.GetClusters(), err)
}

func GreenplumClusterResolver(name string, opts ...ResolveOption) ycsdk.Resolver {
	return &greenplumClusterResolver{
		BaseNameResolver: NewBaseNameResolver(name, "cluster", opts...),
	}
}

type greenplumClusterResolver struct {
	BaseNameResolver
}

func (r *greenplumClusterResolver) Run(ctx context.Context, sdk *ycsdk.SDK, opts ...grpc.CallOption) error {
	err := r.ensureFolderID()
	if err != nil {
		return err
	}

	resp, err := sdk.MDB().Greenplum().Cluster().List(ctx, &greenplum.ListClustersRequest{
		FolderId: r.FolderID(),
		Filter:   CreateResolverFilter("name", r.Name),
		PageSize: DefaultResolverPageSize,
	})
	return r.findName(resp.GetClusters(), err)
}

func KafkaClusterResolver(name string, opts ...ResolveOption) ycsdk.Resolver {
	return &kafkaClusterResolver{
		BaseNameResolver: NewBaseNameResolver(name, "cluster", opts...),
	}
}

type kafkaClusterResolver struct {
	BaseNameResolver
}

func (r *kafkaClusterResolver) Run(ctx context.Context, sdk *ycsdk.SDK, opts ...grpc.CallOption) error {
	err := r.ensureFolderID()
	if err != nil {
		return err
	}

	resp, err := sdk.MDB().Kafka().Cluster().List(ctx, &kafka.ListClustersRequest{
		FolderId: r.FolderID(),
		Filter:   CreateResolverFilter("name", r.Name),
		PageSize: DefaultResolverPageSize,
	})
	return r.findName(resp.GetClusters(), err)
}

func ElasticSearchClusterResolver(name string, opts ...ResolveOption) ycsdk.Resolver {
	return &elasticSearchClusterResolver{
		BaseNameResolver: NewBaseNameResolver(name, "cluster", opts...),
	}
}

type elasticSearchClusterResolver struct {
	BaseNameResolver
}

func (r *elasticSearchClusterResolver) Run(ctx context.Context, sdk *ycsdk.SDK, opts ...grpc.CallOption) error {
	err := r.ensureFolderID()
	if err != nil {
		return err
	}

	resp, err := sdk.MDB().ElasticSearch().Cluster().List(ctx, &elasticsearch.ListClustersRequest{
		FolderId: r.FolderID(),
		Filter:   CreateResolverFilter("name", r.Name),
		PageSize: DefaultResolverPageSize,
	})
	return r.findName(resp.GetClusters(), err)
}

func OpenSearchClusterResolver(name string, opts ...ResolveOption) ycsdk.Resolver {
	return &openSearchClusterResolver{
		BaseNameResolver: NewBaseNameResolver(name, "cluster", opts...),
	}
}

type openSearchClusterResolver struct {
	BaseNameResolver
}

func (r *openSearchClusterResolver) Run(ctx context.Context, sdk *ycsdk.SDK, opts ...grpc.CallOption) error {
	err := r.ensureFolderID()
	if err != nil {
		return err
	}

	resp, err := sdk.MDB().OpenSearch().Cluster().List(ctx, &opensearch.ListClustersRequest{
		FolderId: r.FolderID(),
		Filter:   CreateResolverFilter("name", r.Name),
		PageSize: DefaultResolverPageSize,
	})
	return r.findName(resp.GetClusters(), err)
}
