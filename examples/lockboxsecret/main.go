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
	lockboxsdk "github.com/yandex-cloud/go-sdk/services/lockbox/v1"
)

const (
	// Таймаут на операцию ListSecrets
	listTimeout = 30 * time.Second
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("fatal: %v", err)
	}
}

func run() error {
	// Флаги CLI
	iamToken := flag.String(
		"iam-token",
		os.Getenv("YC_IAM_TOKEN"),
		"IAM token for Yandex.Cloud (env YC_IAM_TOKEN)",
	)
	folderID := flag.String(
		"folder-id",
		os.Getenv("YC_FOLDER_ID"),
		"Yandex.Cloud Folder ID for Lockbox (env YC_FOLDER_ID)",
	)
	flag.Parse()

	if *iamToken == "" {
		return fmt.Errorf("parameter -iam-token is required (or set YC_IAM_TOKEN)")
	}
	if *folderID == "" {
		return fmt.Errorf("parameter -folder-id is required (or set YC_FOLDER_ID)")
	}

	sdk, err := ycsdk.Build(context.Background(), options.WithCredentials(credentials.IAMToken(*iamToken)))

	if err != nil {
		return fmt.Errorf("build SDK: %w", err)
	}

	// Создаём клиента Secret
	secretClient := lockboxsdk.NewSecretClient(sdk)

	// Контекст с таймаутом для листинга
	ctx, cancel := context.WithTimeout(context.Background(), listTimeout)
	defer cancel()

	log.Printf("Listing secrets in folder %q…", *folderID)
	resp, err := secretClient.List(ctx, &lockboxapi.ListSecretsRequest{
		FolderId: *folderID,
	})
	if err != nil {
		return fmt.Errorf("list secrets: %w", err)
	}

	// Выводим результат
	for _, secret := range resp.Secrets {
		log.Printf("- ID: %s, Name: %s", secret.Id, secret.Name)
	}

	return nil
}
