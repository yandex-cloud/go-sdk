package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/c2h5oh/datasize"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	postgresql "github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/postgresql/v1"
	vpc "github.com/yandex-cloud/go-genproto/yandex/cloud/vpc/v1"
	ycsdk "github.com/yandex-cloud/go-sdk/v2"
	"github.com/yandex-cloud/go-sdk/v2/credentials"
	"github.com/yandex-cloud/go-sdk/v2/pkg/options"
	pgsdk "github.com/yandex-cloud/go-sdk/v2/services/mdb/postgresql/v1"
	vpcsdk "github.com/yandex-cloud/go-sdk/v2/services/vpc/v1"
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

func parseCmd() *cmdFlags {
	ret := &cmdFlags{}
	ret.token = flag.String("token", os.Getenv("YC_IAM_TOKEN"), "")
	ret.folderID = flag.String("folder-id", os.Getenv("YC_FOLDER_ID"), "Yandex.Cloud folder ID")
	ret.zoneID = flag.String("zone", "ru-central1-b", "Zone to deploy to")
	ret.networkID = flag.String("network-id", "", "VPC network ID")
	ret.subnetID = flag.String("subnet-id", "", "VPC subnet ID")
	ret.clusterName = flag.String("cluster-name", "postgresql666", "Cluster name")
	ret.clusterDesc = flag.String("cluster-desc", "", "Cluster description")
	ret.dbName = flag.String("db-name", "db1", "Database name")
	ret.userName = flag.String("user-name", "user1", "Username")
	ret.userPassword = flag.String("user-password", "password123", "User password")

	flag.Parse()
	return ret
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
	var networkID string
	for _, net := range resp.Networks {
		if net.FolderId == folderID {
			networkID = net.Id
			break
		}
	}
	if networkID == "" {
		log.Fatalf("no networks in folder: %s", folderID)
	}
	return &networkID
}

func findSubnet(ctx context.Context, sdk *ycsdk.SDK, folderID, networkID, zone string) *string {
	resp, err := vpcsdk.NewSubnetClient(sdk).List(ctx, &vpc.ListSubnetsRequest{
		FolderId: folderID,
		PageSize: 100,
	})
	if err != nil {
		log.Fatal(err)
	}
	var subnetID string
	for _, sn := range resp.Subnets {
		if sn.NetworkId == networkID && sn.ZoneId == zone {
			subnetID = sn.Id
			break
		}
	}
	if subnetID == "" {
		log.Fatalf("no subnets in zone %s for network %s", zone, networkID)
	}
	return &subnetID
}

func createCluster(ctx context.Context, sdk *ycsdk.SDK, req *postgresql.CreateClusterRequest) *postgresql.Cluster {
	op, err := pgsdk.NewClusterClient(sdk).Create(ctx, req)
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

func addClusterHost(ctx context.Context, sdk *ycsdk.SDK, cluster *postgresql.Cluster, flags *cmdFlags) {
	fmt.Printf("Adding host to cluster %s\n", cluster.Id)
	host := postgresql.HostSpec{
		ZoneId:         *flags.zoneID,
		SubnetId:       *flags.subnetID,
		AssignPublicIp: false,
	}
	req := &postgresql.AddClusterHostsRequest{
		ClusterId: cluster.Id,
		HostSpecs: []*postgresql.HostSpec{&host},
	}
	op, err := pgsdk.NewClusterClient(sdk).AddHosts(ctx, req)
	if err != nil {
		log.Panic(err)
	}
	if _, err := op.Wait(ctx); err != nil {
		log.Panic(err)
	}
}

func changeClusterDescription(ctx context.Context, sdk *ycsdk.SDK, cluster *postgresql.Cluster) {
	fmt.Printf("Updating cluster %s\n", cluster.Id)
	mask := &fieldmaskpb.FieldMask{Paths: []string{"description"}}
	req := &postgresql.UpdateClusterRequest{
		ClusterId:   cluster.Id,
		UpdateMask:  mask,
		Description: "New Description!!!",
	}
	op, err := pgsdk.NewClusterClient(sdk).Update(ctx, req)
	if err != nil {
		log.Panic(err)
	}
	if _, err := op.Wait(ctx); err != nil {
		log.Panic(err)
	}
}

func deleteCluster(ctx context.Context, sdk *ycsdk.SDK, cluster *postgresql.Cluster) {
	fmt.Printf("Deleting cluster %s\n", cluster.Id)
	op, err := pgsdk.NewClusterClient(sdk).Delete(ctx, &postgresql.DeleteClusterRequest{
		ClusterId: cluster.Id,
	})
	if err != nil {
		log.Fatal(err)
	}
	if _, err := op.Wait(ctx); err != nil {
		log.Fatal(err)
	}
}

func createClusterRequest(flags *cmdFlags) *postgresql.CreateClusterRequest {
	dbSpec := &postgresql.DatabaseSpec{
		Name:  *flags.dbName,
		Owner: *flags.userName,
	}

	perm := &postgresql.Permission{DatabaseName: *flags.dbName}
	user := &postgresql.UserSpec{
		Name:        *flags.userName,
		Password:    *flags.userPassword,
		Permissions: []*postgresql.Permission{perm},
	}

	host := &postgresql.HostSpec{
		ZoneId:         *flags.zoneID,
		SubnetId:       *flags.subnetID,
		AssignPublicIp: false,
	}

	res := &postgresql.Resources{
		ResourcePresetId: "s1.micro",
		DiskSize:         int64(10 * datasize.GB.Bytes()),
		DiskTypeId:       "network-ssd",
	}

	config := &postgresql.ConfigSpec{
		Version:   "13",
		Resources: res,
	}

	return &postgresql.CreateClusterRequest{
		FolderId:      *flags.folderID,
		Name:          *flags.clusterName,
		Description:   *flags.clusterDesc,
		Environment:   postgresql.Cluster_PRODUCTION,
		ConfigSpec:    config,
		DatabaseSpecs: []*postgresql.DatabaseSpec{dbSpec},
		UserSpecs:     []*postgresql.UserSpec{user},
		HostSpecs:     []*postgresql.HostSpec{host},
		NetworkId:     *flags.networkID,
	}
}
