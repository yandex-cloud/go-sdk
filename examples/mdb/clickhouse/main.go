package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/c2h5oh/datasize"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/clickhouse/v1"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/vpc/v1"
	ycsdk "github.com/yandex-cloud/go-sdk/v2"
	"github.com/yandex-cloud/go-sdk/v2/credentials"
	"github.com/yandex-cloud/go-sdk/v2/pkg/options"
	clickhousesdk "github.com/yandex-cloud/go-sdk/services/mdb/clickhouse/v1"
	vpcsdk "github.com/yandex-cloud/go-sdk/services/vpc/v1"
)

func main() {
	flags := parseCmd()
	ctx := context.Background()

	sdk, err := ycsdk.Build(ctx,
		options.WithCredentials(credentials.IAMToken(*flags.token)),
	)

	if err != nil {
		log.Fatal(err)
	}

	fillMissingFlags(ctx, sdk, flags)

	req := createClusterRequest(flags)
	cluster := createCluster(ctx, sdk, req)
	defer deleteCluster(ctx, sdk, cluster)
	changeClusterDescription(ctx, sdk, cluster)
	addClusterHost(ctx, sdk, cluster, flags)
}

type cmdFlags struct {
	token        *string
	folderID     *string
	zoneID       *string
	networkID    *string
	subnetID     *string
	clusterName  *string
	clusterDesc  *string
	dbName       *string
	userName     *string
	userPassword *string
}

func parseCmd() (ret *cmdFlags) {
	ret = &cmdFlags{}
	ret.token = flag.String("token", os.Getenv("YC_IAM_TOKEN"), "")
	ret.folderID = flag.String("folder-id", os.Getenv("YC_FOLDER_ID"), "Your Yandex.Cloud folder id")
	ret.zoneID = flag.String("zone", "ru-central1-b", "Compute Engine zone to deploy to.")
	ret.networkID = flag.String("network-id", "", "Your Yandex.Cloud network id")
	ret.subnetID = flag.String("subnet-id", "", "Subnet of the instance")
	ret.clusterName = flag.String("cluster-name", "clickhouse666", "")
	ret.clusterDesc = flag.String("cluster-desc", "", "")
	ret.dbName = flag.String("db-name", "db1", "")
	ret.userName = flag.String("user-name", "user1", "")
	ret.userPassword = flag.String("user-password", "password123", "")

	flag.Parse()
	return
}

func fillMissingFlags(ctx context.Context, sdk *ycsdk.SDK, flags *cmdFlags) {
	if *flags.networkID == "" {
		flags.networkID = findNetwork(ctx, sdk, *flags.folderID)
	}

	if *flags.subnetID == "" {
		flags.subnetID = findSubnet(ctx, sdk, *flags.folderID, *flags.networkID, *flags.zoneID)
	}
}

func findNetwork(ctx context.Context, sdk *ycsdk.SDK, folderID string) *string {
	resp, err := vpcsdk.NewNetworkClient(sdk).List(ctx, &vpc.ListNetworksRequest{
		FolderId: folderID,
		PageSize: 100,
	})

	if err != nil {
		log.Fatal(err)
	}
	networkID := ""
	for _, network := range resp.Networks {
		if network.FolderId != folderID {
			continue
		}
		networkID = network.Id
		break
	}
	if networkID == "" {
		log.Fatalf("no networks in folder: %s", folderID)
	}
	return &networkID
}

func findSubnet(ctx context.Context, sdk *ycsdk.SDK, folderID string, networkID string, zone string) *string {
	resp, err := vpcsdk.NewSubnetClient(sdk).List(ctx, &vpc.ListSubnetsRequest{
		FolderId: folderID,
		PageSize: 100,
	})

	if err != nil {
		log.Fatal(err)
	}
	subnetID := ""
	for _, subnet := range resp.Subnets {
		if subnet.ZoneId != zone || subnet.NetworkId != networkID {
			continue
		}
		subnetID = subnet.Id
		break
	}
	if subnetID == "" {
		log.Fatalf("no subnets in zone: %s", zone)
	}
	return &subnetID
}

func createCluster(ctx context.Context, sdk *ycsdk.SDK, req *clickhouse.CreateClusterRequest) *clickhouse.Cluster {
	op, err := clickhousesdk.NewClusterClient(sdk).Create(ctx, req)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Creating cluster %s\n", op.Metadata())

	cluster, err := op.Wait(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return cluster
}

func addClusterHost(ctx context.Context, sdk *ycsdk.SDK, cluster *clickhouse.Cluster, params *cmdFlags) {
	fmt.Printf("Adding host to cluster %s\n", cluster.Id)
	hostSpec := clickhouse.HostSpec{
		ZoneId: *params.zoneID, Type: clickhouse.Host_CLICKHOUSE,
		SubnetId:       *params.subnetID,
		AssignPublicIp: false}

	hostSpecs := []*clickhouse.HostSpec{&hostSpec}
	req := clickhouse.AddClusterHostsRequest{ClusterId: cluster.Id, HostSpecs: hostSpecs}

	op, err := clickhousesdk.NewClusterClient(sdk).AddHosts(ctx, &req)
	if err != nil {
		log.Panic(err)
	}
	_, err = op.Wait(ctx)
	if err != nil {
		log.Panic(err)
	}
}

func changeClusterDescription(ctx context.Context, sdk *ycsdk.SDK, cluster *clickhouse.Cluster) {
	fmt.Printf("Updating cluster %s\n", cluster.Id)
	mask := &fieldmaskpb.FieldMask{
		Paths: []string{
			"description",
		},
	}
	updateReq := clickhouse.UpdateClusterRequest{ClusterId: cluster.Id, UpdateMask: mask, Description: "New Description!!!"}

	op, err := clickhousesdk.NewClusterClient(sdk).Update(ctx, &updateReq)
	if err != nil {
		log.Panic(err)
	}
	_, err = op.Wait(ctx)
	if err != nil {
		log.Panic(err)
	}
}

func deleteCluster(ctx context.Context, sdk *ycsdk.SDK, cluster *clickhouse.Cluster) {
	fmt.Printf("Deleting cluster %s\n", cluster.Id)
	op, err := clickhousesdk.NewClusterClient(sdk).Delete(ctx, &clickhouse.DeleteClusterRequest{ClusterId: cluster.Id})

	if err != nil {
		log.Fatal(err)
	}
	_, err = op.Wait(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func createClusterRequest(params *cmdFlags) *clickhouse.CreateClusterRequest {
	dbSpec := clickhouse.DatabaseSpec{Name: *params.dbName}
	dbSpecs := []*clickhouse.DatabaseSpec{&dbSpec}

	permission := clickhouse.Permission{DatabaseName: *params.dbName}
	permissions := []*clickhouse.Permission{&permission}
	userSpec := clickhouse.UserSpec{Name: *params.userName, Password: *params.userPassword, Permissions: permissions}
	userSpecs := []*clickhouse.UserSpec{&userSpec}

	hostCKSpec := clickhouse.HostSpec{
		ZoneId: *params.zoneID, Type: clickhouse.Host_CLICKHOUSE,
		SubnetId:       *params.subnetID,
		AssignPublicIp: false}

	hostZooSpec := clickhouse.HostSpec{
		ZoneId: *params.zoneID, Type: clickhouse.Host_ZOOKEEPER,
		SubnetId:       *params.subnetID,
		AssignPublicIp: false}

	hostSpecs := []*clickhouse.HostSpec{&hostCKSpec, &hostCKSpec, &hostZooSpec, &hostZooSpec, &hostZooSpec}

	zres := &clickhouse.Resources{ResourcePresetId: "s1.micro", DiskSize: int64(10 * datasize.GB.Bytes()), DiskTypeId: "network-ssd"}
	cres := &clickhouse.Resources{ResourcePresetId: "s1.micro", DiskSize: int64(10 * datasize.GB.Bytes()), DiskTypeId: "network-ssd"}

	configSpec := &clickhouse.ConfigSpec{
		Clickhouse: &clickhouse.ConfigSpec_Clickhouse{Resources: cres},
		Zookeeper:  &clickhouse.ConfigSpec_Zookeeper{Resources: zres},
	}

	req := clickhouse.CreateClusterRequest{
		FolderId:      *params.folderID,
		Name:          *params.clusterName,
		Description:   *params.clusterDesc,
		Environment:   clickhouse.Cluster_PRODUCTION,
		ConfigSpec:    configSpec,
		DatabaseSpecs: dbSpecs,
		UserSpecs:     userSpecs,
		HostSpecs:     hostSpecs,
		NetworkId:     *params.networkID,
	}
	return &req
}
