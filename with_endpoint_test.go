package ycsdk

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/yandex-cloud/go-sdk/v2/pkg/endpoints"
	"github.com/yandex-cloud/go-sdk/v2/pkg/transport"
	transportgrpc "github.com/yandex-cloud/go-sdk/v2/pkg/transport/grpc"
)

func newTestSDK(resolver *endpoints.PrefixEndpointsResolver) *SDK {
	pool := transportgrpc.NewConnPool(nil)
	return &SDK{
		conn:             transport.NewConnector(resolver, pool),
		endpointResolver: resolver,
		connPool:         pool,
	}
}

func TestWithEndpoint_OverridesOnlyTargetPrefix(t *testing.T) {
	base := endpoints.NewPrefixEndpointsResolver(endpoints.PrefixToEndpoint{
		"yandex.cloud.iam":     endpoints.NewEndpointParams("iam.base:443"),
		"yandex.cloud.compute": endpoints.NewEndpointParams("compute.base:443"),
	})
	sdk := newTestSDK(base)

	overridden := sdk.WithEndpoint(
		"yandex.cloud.compute",
		endpoints.NewEndpointParams("compute.override:443"),
	)

	const computeMethod = protoreflect.FullName("yandex.cloud.compute.v1.ImageService.List")
	const iamMethod = protoreflect.FullName("yandex.cloud.iam.v1.IamTokenService.Create")

	cep, err := overridden.GetEndpoint(computeMethod)
	require.NoError(t, err)
	require.Equal(t, "compute.override:443", cep.Addr)

	iep, err := overridden.GetEndpoint(iamMethod)
	require.NoError(t, err)
	require.Equal(t, "iam.base:443", iep.Addr, "non-overridden endpoints must be preserved")
}

func TestWithEndpoint_DoesNotMutateOriginal(t *testing.T) {
	base := endpoints.NewPrefixEndpointsResolver(endpoints.PrefixToEndpoint{
		"yandex.cloud.compute": endpoints.NewEndpointParams("compute.base:443"),
	})
	sdk := newTestSDK(base)

	_ = sdk.WithEndpoint(
		"yandex.cloud.compute",
		endpoints.NewEndpointParams("compute.override:443"),
	)

	const computeMethod = protoreflect.FullName("yandex.cloud.compute.v1.ImageService.List")
	ep, err := sdk.GetEndpoint(computeMethod)
	require.NoError(t, err)
	require.Equal(t, "compute.base:443", ep.Addr, "original sdk resolver must stay untouched")
}

func TestWithEndpoint_RebuildsInternalConnector(t *testing.T) {
	base := endpoints.NewPrefixEndpointsResolver(endpoints.PrefixToEndpoint{
		"yandex.cloud.iam": endpoints.NewEndpointParams("iam.base:443"),
	})
	sdk := newTestSDK(base)

	overridden := sdk.WithEndpoint(
		"yandex.cloud.compute",
		endpoints.NewEndpointParams("compute.override:443"),
	)

	require.NotSame(t, sdk.conn, overridden.conn,
		"WithEndpoint must rebuild the internal connector so gRPC routing picks up the override")
	require.Same(t, sdk.connPool, overridden.connPool,
		"WithEndpoint must reuse the connection pool to preserve auth interceptors and connection cache")
}
