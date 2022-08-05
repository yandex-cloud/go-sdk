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
	secretID := flag.String("secret-id", "", "Your Yandex.Cloud Lockbox ID of the secret")
	sdk, err := ycsdk.Build(context.Background(), ycsdk.Config{
		Credentials: ycsdk.NewIAMTokenCredentials(*iamToken),
	})
	if err != nil {
		log.Fatal(err)
	}
	p, err := sdk.LockboxPayload().Payload().Get(context.Background(), &lockbox.GetPayloadRequest{
		SecretId: *secretID,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(p)
}
