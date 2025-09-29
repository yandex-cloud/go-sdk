package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	ycsdk "github.com/yandex-cloud/go-sdk/v2"
	"github.com/yandex-cloud/go-sdk/v2/credentials"
	"github.com/yandex-cloud/go-sdk/v2/pkg/options"
	"github.com/yandex-cloud/go-sdk/v2/pkg/sdkresolvers"
	computesdk "github.com/yandex-cloud/go-sdk/v2/sdkresolvers/compute/v1"
)

func main() {
	iamToken := flag.String("token", os.Getenv("YC_IAM_TOKEN"), "IAM token for Yandex.Cloud (env YC_IAM_TOKEN)")
	folderID := flag.String("folder-id", os.Getenv("YC_FOLDER_ID"), "Yandex.Cloud Folder ID (env YC_FOLDER_ID)")

	ctx := context.Background()
	sdk, err := ycsdk.Build(ctx,
		options.WithCredentials(credentials.IAMToken(*iamToken)),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "test_name"
	r := computesdk.DiskResolver(name, sdkresolvers.FolderID(*folderID))
	if err = r.Run(ctx, sdk); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("id of object %s is %s", name, r.ID())
}
