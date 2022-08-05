package main

import (
	"context"
	"flag"
	"log"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/lockbox/v1"
	ycsdk "github.com/yandex-cloud/go-sdk"
)

func main() {
	iamToken := flag.String("iam-token", "", "")
	folderID := flag.String("folder-id", "", "Your Yandex.Cloud Lockbox ID of the folder")
	sdk, err := ycsdk.Build(context.Background(), ycsdk.Config{
		Credentials: ycsdk.NewIAMTokenCredentials(*iamToken),
	})
	if err != nil {
		log.Fatal(err)
	}
	p, err := sdk.LockboxSecret().Secret().List(context.Background(), &lockbox.ListSecretsRequest{
		FolderId: *folderID,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(p)
}
