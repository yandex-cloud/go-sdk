package main

import (
	"context"
	"flag"
	"log"

	"google.golang.org/grpc/codes"

	ycsdk "github.com/yandex-cloud/go-sdk/v2"
	"github.com/yandex-cloud/go-sdk/v2/credentials"
	"github.com/yandex-cloud/go-sdk/v2/pkg/options"
	"github.com/yandex-cloud/go-sdk/v2/pkg/options/retry"
)

func main() {
	token := flag.String("token", "", "")
	flag.Parse()

	ctx := context.Background()

	_, err := ycsdk.Build(
		ctx,
		options.WithCredentials(credentials.IAMToken(*token)),
		options.WithRetryOptions(
			retry.WithRetries(retry.DefaultNameConfig(), 2),
			retry.WithRetryableStatusCodes(retry.DefaultNameConfig(), codes.AlreadyExists, codes.Unavailable),
		),
		options.WithDefaultRetryOptions(),
	)
	if err != nil {
		log.Fatal(err)
	}
}
