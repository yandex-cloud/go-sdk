package idempotency

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// Interceptor add idempotency key to context.
func Interceptor() grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply any,
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		ctx = addIdempotencyToken(ctx)
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

func addIdempotencyToken(ctx context.Context) context.Context {
	const idempotencyTokenMetadataKey = "idempotency-key"

	idempotencyTokenPresent := false
	md, ok := metadata.FromOutgoingContext(ctx)
	if ok {
		_, idempotencyTokenPresent = md[idempotencyTokenMetadataKey]
	}

	if !idempotencyTokenPresent {
		ctx = metadata.AppendToOutgoingContext(ctx, idempotencyTokenMetadataKey, uuid.New().String())
	}

	return ctx
}
