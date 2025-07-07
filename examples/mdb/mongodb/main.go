package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/c2h5oh/datasize"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	mdb "github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/mongodb/v1"
	vpc "github.com/yandex-cloud/go-genproto/yandex/cloud/vpc/v1"
	ycsdk "github.com/yandex-cloud/go-sdk/v2"
	"github.com/yandex-cloud/go-sdk/v2/credentials"
	"github.com/yandex-cloud/go-sdk/v2/pkg/options"
	mdbsdk "github.com/yandex-cloud/go-sdk/v2/services/mdb/mongodb/v1"
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
	ret.folderID = flag.String("folder-id", os.Getenv("YC_FOLDER_ID"), "Your Yandex.Cloud folder id")
	ret.zoneID = flag.String("zone", "ru-central1-b", "Compute Engine zone to deploy to")
	ret.networkID = flag.String("network-id", "", "Your Yandex.Cloud network id")
	ret.subnetID = flag.String("subnet-id", "", "Subnet of the instance")
	ret.clusterName = flag.String("cluster-name", "mongodb666", "Cluster name")
	ret.clusterDesc = flag.String("cluster-desc", "", "Cluster description")
	ret.dbName = flag.String("db-name", "db1", "Database name")
	ret.userName = flag.String("user-name", "user1", "User name")
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

func findSubnet(ctx context.Context, sdk *ycsdk.SDK, folderID, networkID, zone string) *string {
	resp, err := vpcsdk.NewSubnetClient(sdk).List(ctx, &vpc.ListSubnetsRequest{
		FolderId: folderID,
		PageSize: 100,
	})
	if err != nil {
		log.Fatal(err)
	}
	var subnetID string
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

func createCluster(ctx context.Context, sdk *ycsdk.SDK, req *mdb.CreateClusterRequest) *mdb.Cluster {
	op, err := mdbsdk.NewClusterClient(sdk).Create(ctx, req)
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

func addClusterHost(ctx context.Context, sdk *ycsdk.SDK, cluster *mdb.Cluster, params *cmdFlags) {
	fmt.Printf("Adding host to cluster %s\n", cluster.Id)
	hostSpec := mdb.HostSpec{
		ZoneId:         *params.zoneID,
		SubnetId:       *params.subnetID,
		AssignPublicIp: false,
	}
	hostSpecs := []*mdb.HostSpec{&hostSpec}
	req := mdb.AddClusterHostsRequest{
		ClusterId: cluster.Id,
		HostSpecs: hostSpecs,
	}

	op, err := mdbsdk.NewClusterClient(sdk).AddHosts(ctx, &req)
	if err != nil {
		log.Panic(err)
	}
	if _, err := op.Wait(ctx); err != nil {
		log.Panic(err)
	}
}

func changeClusterDescription(ctx context.Context, sdk *ycsdk.SDK, cluster *mdb.Cluster) {
	fmt.Printf("Updating cluster %s\n", cluster.Id)
	mask := &fieldmaskpb.FieldMask{Paths: []string{"description"}}
	req := mdb.UpdateClusterRequest{
		ClusterId:   cluster.Id,
		UpdateMask:  mask,
		Description: "New Description!!!",
	}

	op, err := mdbsdk.NewClusterClient(sdk).Update(ctx, &req)
	if err != nil {
		log.Panic(err)
	}
	if _, err := op.Wait(ctx); err != nil {
		log.Panic(err)
	}
}

func deleteCluster(ctx context.Context, sdk *ycsdk.SDK, cluster *mdb.Cluster) {
	fmt.Printf("Deleting cluster %s\n", cluster.Id)
	op, err := mdbsdk.NewClusterClient(sdk).Delete(ctx, &mdb.DeleteClusterRequest{
		ClusterId: cluster.Id,
	})
	if err != nil {
		log.Fatal(err)
	}
	if _, err := op.Wait(ctx); err != nil {
		log.Fatal(err)
	}
}

func createClusterRequest(params *cmdFlags) *mdb.CreateClusterRequest {
	dbSpec := mdb.DatabaseSpec{Name: *params.dbName}
	dbSpecs := []*mdb.DatabaseSpec{&dbSpec}

	perm := mdb.Permission{DatabaseName: *params.dbName}
	perms := []*mdb.Permission{&perm}

	userSpec := mdb.UserSpec{
		Name:        *params.userName,
		Password:    *params.userPassword,
		Permissions: perms,
	}
	userSpecs := []*mdb.UserSpec{&userSpec}

	hostSpec := mdb.HostSpec{
		ZoneId:         *params.zoneID,
		SubnetId:       *params.subnetID,
		AssignPublicIp: false,
	}
	hostSpecs := []*mdb.HostSpec{&hostSpec}

	res := &mdb.Resources{
		ResourcePresetId: "s1.micro",
		DiskSize:         int64(10 * datasize.GB.Bytes()),
		DiskTypeId:       "network-ssd",
	}

	configSpec := &mdb.ConfigSpec{
		Version: "6.0",
		MongodbSpec: &mdb.ConfigSpec_MongodbSpec_6_0{
			MongodbSpec_6_0: &mdb.MongodbSpec6_0{
				Mongod: &mdb.MongodbSpec6_0_Mongod{Resources: res},
			},
		},
	}

	return &mdb.CreateClusterRequest{
		FolderId:      *params.folderID,
		Name:          *params.clusterName,
		Description:   *params.clusterDesc,
		Environment:   mdb.Cluster_PRODUCTION,
		ConfigSpec:    configSpec,
		DatabaseSpecs: dbSpecs,
		UserSpecs:     userSpecs,
		HostSpecs:     hostSpecs,
		NetworkId:     *params.networkID,
	}
}
