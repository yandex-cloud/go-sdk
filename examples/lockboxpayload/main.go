package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	lockboxapi "github.com/yandex-cloud/go-genproto/yandex/cloud/lockbox/v1"
	ycsdk "github.com/yandex-cloud/go-sdk/v2"
	"github.com/yandex-cloud/go-sdk/v2/credentials"
	"github.com/yandex-cloud/go-sdk/v2/pkg/options"
	lockboxsdk "github.com/yandex-cloud/go-sdk/v2/services/lockbox/v1"
)

const (
	payloadTimeout = 30 * time.Second
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("fatal: %v", err)
	}
}

func run() error {
	iamToken := flag.String("iam-token", os.Getenv("YC_IAM_TOKEN"), "IAM token for Yandex.Cloud (env YC_IAM_TOKEN)")
	secretID := flag.String("secret-id", os.Getenv("YC_SECRET_ID"), "Lockbox Secret ID (env YC_SECRET_ID)")
	flag.Parse()

	if *iamToken == "" {
		return fmt.Errorf("parameter -iam-token is required (or set YC_IAM_TOKEN)")
	}
	if *secretID == "" {
		return fmt.Errorf("parameter -secret-id is required (or set YC_SECRET_ID)")
	}

	sdk, err := ycsdk.Build(context.Background(), options.WithCredentials(credentials.IAMToken(*iamToken)))
	if err != nil {
		return fmt.Errorf("build SDK: %w", err)
	}

	payloadClient := lockboxsdk.NewPayloadClient(sdk)

	ctx, cancel := context.WithTimeout(context.Background(), payloadTimeout)
	defer cancel()

	log.Printf("Fetching payload for SecretID=%s…", *secretID)
	payload, err := getPayload(ctx, payloadClient, *secretID)
	if err != nil {
		return fmt.Errorf("get payload: %w", err)
	}

	log.Printf("Payload retrieved: %v", payload)

	return nil
}

func getPayload(
	ctx context.Context,
	client lockboxsdk.PayloadClient,
	secretID string,
) (*lockboxapi.Payload, error) {
	resp, err := client.Get(ctx, &lockboxapi.GetPayloadRequest{
		SecretId: secretID,
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}
