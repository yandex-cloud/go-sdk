package idempotency

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/metadata"
)

const idempotencyKey = "idempotency-key"

func TestAddIdempotencyKey(t *testing.T) {
	t.Run("have no idempotency key before", func(t *testing.T) {
		ctx := context.Background()
		actualContext := addIdempotencyToken(ctx)
		actualMetadata, ok := metadata.FromOutgoingContext(actualContext)
		require.True(t, ok)
		require.NotNil(t, actualMetadata.Get(idempotencyKey))

		_, ok = metadata.FromOutgoingContext(ctx)
		require.False(t, ok)
	})

	t.Run("have idempotency key before", func(t *testing.T) {
		expectedCtx := context.Background()
		expectedCtx = metadata.AppendToOutgoingContext(expectedCtx, idempotencyKey, "somevalue")
		actualContext := addIdempotencyToken(expectedCtx)

		expectedMetadata, ok := metadata.FromOutgoingContext(expectedCtx)
		require.True(t, ok)

		actualMetadata, ok := metadata.FromOutgoingContext(actualContext)
		require.True(t, ok)
		require.EqualValues(
			t,
			expectedMetadata.Get(idempotencyKey),
			actualMetadata.Get(idempotencyKey),
		)
	})
}
