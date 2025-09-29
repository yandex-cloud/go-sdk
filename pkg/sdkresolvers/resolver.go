package sdkresolvers

import (
	"context"

	"google.golang.org/grpc"

	ycsdk "github.com/yandex-cloud/go-sdk/v2"
)

type Resolver interface {
	ID() string
	Err() error

	Run(context.Context, *ycsdk.SDK, ...grpc.CallOption) error
}
