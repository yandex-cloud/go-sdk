package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	computeapi "github.com/yandex-cloud/go-genproto/yandex/cloud/compute/v1"
	vpcapi "github.com/yandex-cloud/go-genproto/yandex/cloud/vpc/v1"
	ycsdk "github.com/yandex-cloud/go-sdk/v2"
	"github.com/yandex-cloud/go-sdk/v2/credentials"
	"github.com/yandex-cloud/go-sdk/v2/pkg/options"
	computesdk "github.com/yandex-cloud/go-sdk/services/compute/v1"
	vpcsdk "github.com/yandex-cloud/go-sdk/services/vpc/v1"
)

const (
	defaultZone       = "ru-central1-b"
	defaultPlatformID = "standard-v1"
	defaultFamily     = "debian-9"
	imageFolderID     = "standard-images"
	createTimeout     = 5 * time.Minute
	deleteTimeout     = 2 * time.Minute
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("fatal: %v", err)
	}
}

func run() error {
	iamToken := flag.String("token", os.Getenv("YC_IAM_TOKEN"), "IAM token for Yandex.Cloud (env YC_IAM_TOKEN)")
	folderID := flag.String("folder-id", os.Getenv("YC_FOLDER_ID"), "Yandex.Cloud Folder ID (env YC_FOLDER_ID)")
	zone := flag.String("zone", defaultZone, "Compute Engine zone")
	name := flag.String("name", "demo-instance", "Name of the instance to create")
	subnetID := flag.String("subnet-id", "", "Subnet ID (опционально)")
	flag.Parse()

	if *iamToken == "" {
		return fmt.Errorf("token is required (use -token or set YC_IAM_TOKEN)")
	}
	if *folderID == "" {
		return fmt.Errorf("folder-id is required (use -folder-id or set FOLDER_ID)")
	}

	ctx := context.Background()

	sdk, err := ycsdk.Build(ctx,
		options.WithCredentials(credentials.IAMToken(*iamToken)),
	)
	if err != nil {
		return fmt.Errorf("failed to build SDK: %w", err)
	}

	computeClient := computesdk.NewInstanceClient(sdk)
	imageClient := computesdk.NewImageClient(sdk)
	subnetClient := vpcsdk.NewSubnetClient(sdk)

	if *subnetID == "" {
		log.Printf("No subnet-id provided, looking up in folder %s, zone %s…", *folderID, *zone)
		id, err := findSubnet(ctx, subnetClient, *folderID, *zone)
		if err != nil {
			return fmt.Errorf("find subnet: %w", err)
		}
		*subnetID = id
		log.Printf("→ Selected subnet-id: %s", *subnetID)
	}

	log.Printf("Fetching latest image for family %q…", defaultFamily)
	imageID, err := getLatestImage(ctx, imageClient, imageFolderID, defaultFamily)
	if err != nil {
		return fmt.Errorf("get latest image: %w", err)
	}
	log.Printf("→ Image ID: %s", imageID)

	op, err := createInstance(ctx, computeClient, *folderID, *zone, *name, *subnetID, imageID)
	if err != nil {
		return fmt.Errorf("create instance: %w", err)
	}

	ctxCreate, cancelCreate := context.WithTimeout(ctx, createTimeout)
	defer cancelCreate()

	log.Printf("Waiting for creation of instance %q…", *name)
	createdInst, err := op.Wait(ctxCreate)
	if err != nil {
		return fmt.Errorf("wait create op: %w", err)
	}
	log.Printf("Instance %q created, ID=%s", *name, createdInst.Id)

	log.Printf("Deleting instance ID=%s…", createdInst.Id)
	delOp, err := computeClient.Delete(ctx, &computeapi.DeleteInstanceRequest{
		InstanceId: createdInst.Id,
	})
	if err != nil {
		return fmt.Errorf("delete instance: %w", err)
	}

	ctxDel, cancelDel := context.WithTimeout(ctx, deleteTimeout)
	defer cancelDel()

	if _, err := delOp.Wait(ctxDel); err != nil {
		return fmt.Errorf("wait delete op: %w", err)
	}
	log.Printf("Instance ID=%s deleted", createdInst.Id)

	return nil
}

func createInstance(
	ctx context.Context,
	client computesdk.InstanceClient,
	folderID, zone, name, subnetID, imageID string,
) (*computesdk.InstanceCreateOperation, error) {
	req := &computeapi.CreateInstanceRequest{
		FolderId:   folderID,
		Name:       name,
		ZoneId:     zone,
		PlatformId: defaultPlatformID,
		ResourcesSpec: &computeapi.ResourcesSpec{
			Cores:  2,
			Memory: 2 * 1024 * 1024 * 1024, // 2 GiB
		},
		BootDiskSpec: &computeapi.AttachedDiskSpec{
			AutoDelete: true,
			Disk: &computeapi.AttachedDiskSpec_DiskSpec_{
				DiskSpec: &computeapi.AttachedDiskSpec_DiskSpec{
					TypeId: "network-hdd",
					Size:   20 * 1024 * 1024 * 1024, // 20 GiB
					Source: &computeapi.AttachedDiskSpec_DiskSpec_ImageId{
						ImageId: imageID,
					},
				},
			},
		},
		NetworkInterfaceSpecs: []*computeapi.NetworkInterfaceSpec{
			{
				SubnetId: subnetID,
				PrimaryV4AddressSpec: &computeapi.PrimaryAddressSpec{
					OneToOneNatSpec: &computeapi.OneToOneNatSpec{
						IpVersion: computeapi.IpVersion_IPV4,
					},
				},
			},
		},
	}
	return client.Create(ctx, req)
}

func getLatestImage(
	ctx context.Context,
	client computesdk.ImageClient,
	folderID, family string,
) (string, error) {
	resp, err := client.GetLatestByFamily(ctx, &computeapi.GetImageLatestByFamilyRequest{
		FolderId: folderID,
		Family:   family,
	})
	if err != nil {
		return "", err
	}
	return resp.Id, nil
}

func findSubnet(
	ctx context.Context,
	client vpcsdk.SubnetClient,
	folderID, zone string,
) (string, error) {
	resp, err := client.List(ctx, &vpcapi.ListSubnetsRequest{
		FolderId: folderID,
		PageSize: 100,
	})
	if err != nil {
		return "", err
	}
	for _, s := range resp.Subnets {
		if s.ZoneId == zone {
			return s.Id, nil
		}
	}
	return "", fmt.Errorf("no subnet found in zone %s", zone)
}
