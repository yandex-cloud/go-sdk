package ycsdk

import (
	"context"
	"crypto/tls"
	"fmt"
	"runtime/debug"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	grpccreds "google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/reflect/protoreflect"

	endpointpb "github.com/yandex-cloud/go-genproto/yandex/cloud/endpoint"
	"github.com/yandex-cloud/go-sdk/v2/credentials"
	"github.com/yandex-cloud/go-sdk/v2/pkg/authentication"
	"github.com/yandex-cloud/go-sdk/v2/pkg/endpoints"
	"github.com/yandex-cloud/go-sdk/v2/pkg/log"
	"github.com/yandex-cloud/go-sdk/v2/pkg/options"
	"github.com/yandex-cloud/go-sdk/v2/pkg/options/retry"
	"github.com/yandex-cloud/go-sdk/v2/pkg/transport"
	transportgrpc "github.com/yandex-cloud/go-sdk/v2/pkg/transport/grpc"
	transportauth "github.com/yandex-cloud/go-sdk/v2/pkg/transport/middleware/authentication"
	endpointsdk "github.com/yandex-cloud/go-sdk/v2/services/endpoint"
	endpointssdk "github.com/yandex-cloud/go-sdk/v2/services/endpoints"
	iamsdk "github.com/yandex-cloud/go-sdk/v2/services/iam/v1"
)

var _ transport.Connector = (*SDK)(nil)

// SDK provides a client connection wrapper managing connection pooling and endpoint resolution for gRPC services.
type SDK struct {
	connPool *transportgrpc.ConnPool
	conn     transport.Connector
}

// Build initializes and configures an SDK instance with the provided options and context.
// It applies default configurations and validates necessary parameters like credentials and endpoints.
// Returns an SDK instance or an error if initialization fails.
func Build(ctx context.Context, opts ...options.Option) (*SDK, error) {
	buildOptions := options.DefaultOptions()
	for _, opt := range opts {
		opt(buildOptions)
	}
	if buildOptions.Credentials == nil {
		return nil, fmt.Errorf("credentials must be provided")
	}

	logger := zap.NewNop()
	if buildOptions.Logger != nil {
		logger = buildOptions.Logger
	}

	if injector, ok := buildOptions.Credentials.(log.LogInjector); ok {
		injector.InjectLogger(logger)
	}

	var err error
	if buildOptions.EndpointsResolver == nil {
		buildOptions.EndpointsResolver, err = buildEndpointsResolver(ctx, buildOptions.DiscoveryEndpoint)
		if err != nil {
			return nil, fmt.Errorf("failed to get endpointsResolver: %w", err)
		}
	}

	if buildOptions.Authenticator == nil {
		buildOptions.Authenticator, err = defaultAuthenticator(ctx, logger, buildOptions.Credentials, buildOptions.EndpointsResolver)
		if err != nil {
			return nil, fmt.Errorf("failed to create authenticator: %w", err)
		}
	}

	if injector, ok := buildOptions.Authenticator.(log.LogInjector); ok {
		injector.InjectLogger(
			logger,
		)
	}

	dialOpts := []grpc.DialOption{
		grpc.WithUserAgent(userAgent()),
	}

	if _, ok := buildOptions.Credentials.(*credentials.NoCredentials); !ok {
		tokenMiddleware := transportauth.NewIAMTokenMiddleware(buildOptions.Authenticator, transportauth.WithLogger(logger))
		dialOpts = append(dialOpts,
			grpc.WithUnaryInterceptor(tokenMiddleware.InterceptUnary),
			grpc.WithStreamInterceptor(tokenMiddleware.InterceptStream),
		)
	}

	if buildOptions.Plaintext {
		dialOpts = append(dialOpts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	} else {
		tlsConfig := buildOptions.TlsConfig
		if tlsConfig == nil {
			tlsConfig = &tls.Config{}
		}
		dialOpts = append(dialOpts, grpc.WithTransportCredentials(grpccreds.NewTLS(tlsConfig)))
	}

	if buildOptions.DefaultRetryOptions {
		retryOpt, err := retry.DefaultRetryDialOption()
		if err != nil {
			return nil, fmt.Errorf("failed to apply default retry options: %w", err)
		}
		dialOpts = append(dialOpts, retryOpt)
	}

	if len(buildOptions.RetryOptions) > 0 {
		retryOpt, err := retry.RetryDialOption(buildOptions.RetryOptions...)
		if err != nil {
			return nil, fmt.Errorf("failed to apply retry options: %w", err)
		}
		dialOpts = append(dialOpts, retryOpt)
	}

	dialOpts = append(dialOpts, buildOptions.CustomDialOpts...)

	connectionPool := transportgrpc.NewConnPool(dialOpts)

	return &SDK{
		conn:     transport.NewConnector(buildOptions.EndpointsResolver, connectionPool),
		connPool: connectionPool,
	}, nil
}

// GetConnection retrieves a gRPC client connection for the specified method with optional call options.
func (sdk *SDK) GetConnection(ctx context.Context, method protoreflect.FullName, opts ...grpc.CallOption) (grpc.ClientConnInterface, error) {
	return sdk.conn.GetConnection(ctx, method, opts...)
}

// Shutdown gracefully terminates the SDK by closing all active gRPC connections in the connection pool.
func (sdk *SDK) Shutdown(ctx context.Context) error {
	return sdk.connPool.Shutdown(ctx)
}

// userAgent returns the User-Agent string that includes the SDK name and its version, read from the build info.
func userAgent() string {
	cloudUserAgent := "yandex-cloud/go-sdk-v2"

	build, _ := debug.ReadBuildInfo()
	version := "unknown"

	if build.Main.Version != "" {
		version = build.Main.Version
	}

	return fmt.Sprintf("%s/%s", cloudUserAgent, version)
}

// defaultAuthenticator initializes an Authenticator using provided credentials and an endpoint resolver in the given context.
func defaultAuthenticator(ctx context.Context, logger *zap.Logger, creds credentials.Credentials, resolver endpoints.EndpointsResolver) (authentication.Authenticator, error) {
	authEndpoint, err := resolver.Endpoint(ctx, iamsdk.IamTokenCreate)
	if err != nil {
		return nil, fmt.Errorf("failed to get auth endpoint: %w", err)
	}

	authernticator, err := authentication.NewAuthenticatorFromEndpoint(logger, creds, authEndpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to create authenticator: %w", err)
	}
	return authernticator, nil
}

// BuildEndpointsResolver creates an EndpointsResolver using a discovery endpoint to dynamically map service prefixes.
func buildEndpointsResolver(ctx context.Context, discoveryEndpoint string) (endpoints.EndpointsResolver, error) {
	conn := transport.NewSingleConnector(discoveryEndpoint, grpc.WithTransportCredentials(grpccreds.NewTLS(&tls.Config{})))
	client := endpointsdk.NewApiEndpointClient(conn)

	resp, err := client.List(ctx, &endpointpb.ListApiEndpointsRequest{})
	if err != nil {
		return nil, fmt.Errorf("failed to list endpoints: %w", err)
	}

	// Map endpoint IDs to addresses
	endpointMap := make(map[string]string, len(resp.Endpoints))
	for _, ep := range resp.Endpoints {
		endpointMap[ep.Id] = ep.Address
	}

	// Build resolver from dynamic endpoints
	p2e := make(endpoints.PrefixToEndpoint, len(endpointssdk.DynamicEndpoints))
	for prefix, id := range endpointssdk.DynamicEndpoints {
		if addr, ok := endpointMap[id]; ok {
			p2e[prefix] = endpoints.NewEndpointParams(addr)
		}
	}

	return endpoints.NewPrefixEndpointsResolver(p2e), nil
}
