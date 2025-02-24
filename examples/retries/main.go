package main

import (
	"context"
	"flag"
	"log"

	"google.golang.org/grpc/codes"

	ycsdk "github.com/yandex-cloud/go-sdk"
	"github.com/yandex-cloud/go-sdk/pkg/retry/v1"
)

func main() {
	token := flag.String("token", "", "")
	flag.Parse()

	ctx := context.Background()

	// retriesDialOption, err := retry.DefaultRetryDialOption()
	retriesDialOption, err := retry.RetryDialOption(
		retry.WithRetries(retry.DefaultNameConfig(), 2),
		retry.WithRetryableStatusCodes(retry.DefaultNameConfig(), codes.AlreadyExists, codes.Unavailable),
	)
	if err != nil {
		log.Fatal(err)
	}

	_, err = ycsdk.Build(
		ctx,
		ycsdk.Config{
			Credentials: ycsdk.OAuthToken(*token),
		},
		retriesDialOption,
	)
	if err != nil {
		log.Fatal(err)
	}
}
