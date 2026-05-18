package grpcdebug

import (
	"context"
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	iampb "github.com/yandex-cloud/go-genproto/yandex/cloud/iam/v1"
)

func newObserverLogger(level zapcore.Level) (*zap.Logger, *observer.ObservedLogs) {
	core, logs := observer.New(level)
	return zap.New(core), logs
}

func findField(fields []zapcore.Field, key string) (zapcore.Field, bool) {
	for _, f := range fields {
		if f.Key == key {
			return f, true
		}
	}
	return zapcore.Field{}, false
}

// envelopeFor extracts the zap.Any field with the given key and re-encodes it
// to JSON, then returns the resulting string. Matches the actual rendered form
// users see in stderr / trace file output.
func envelopeFor(t *testing.T, ctx []zapcore.Field, key string) string {
	t.Helper()
	field, ok := findField(ctx, key)
	require.True(t, ok, "expected %q field in log entry", key)
	require.Equal(t, zapcore.ReflectType, field.Type, "field %q must be a structured object", key)
	data, err := json.Marshal(field.Interface)
	require.NoError(t, err)
	return string(data)
}

func TestUnaryClientInterceptor_NoOpAboveDebugLevel(t *testing.T) {
	logger, logs := newObserverLogger(zapcore.InfoLevel)
	interceptor := UnaryClientInterceptor(logger)

	called := false
	err := interceptor(
		context.Background(),
		"/svc/Method",
		&iampb.CreateIamTokenRequest{},
		&iampb.CreateIamTokenResponse{},
		nil,
		func(_ context.Context, _ string, _, _ any, _ *grpc.ClientConn, _ ...grpc.CallOption) error {
			called = true
			return nil
		},
	)

	require.NoError(t, err)
	assert.True(t, called, "invoker must still be called when logger is above Debug")
	assert.Equal(t, 0, logs.Len(), "no log entries when level is above Debug")
}

func TestUnaryClientInterceptor_LogsRequestResponseEnvelope(t *testing.T) {
	logger, logs := newObserverLogger(zapcore.DebugLevel)
	interceptor := UnaryClientInterceptor(logger)

	ctx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(
		"x-request-id", "req-id-1",
		"x-trace-id", "trace-id-1",
	))

	req := &iampb.CreateIamTokenForServiceAccountRequest{ServiceAccountId: "sa-marker-xyz"}
	resp := &iampb.CreateIamTokenResponse{}

	err := interceptor(
		ctx,
		"/yandex.cloud.iam.v1.IamTokenService/CreateForServiceAccount",
		req,
		resp,
		nil,
		func(_ context.Context, _ string, _, _ any, _ *grpc.ClientConn, _ ...grpc.CallOption) error {
			return nil
		},
	)
	require.NoError(t, err)

	entries := logs.AllUntimed()
	require.Len(t, entries, 2)
	assert.Equal(t, "Request IamTokenService/CreateForServiceAccount", entries[0].Message)
	assert.Equal(t, "Response IamTokenService/CreateForServiceAccount", entries[1].Message)

	reqIDField, ok := findField(entries[0].Context, "request_id")
	require.True(t, ok)
	assert.Equal(t, "req-id-1", reqIDField.String)

	reqJSON := envelopeFor(t, entries[0].Context, "request")
	assert.Contains(t, reqJSON, `"method":"/yandex.cloud.iam.v1.IamTokenService/CreateForServiceAccount"`)
	assert.Contains(t, reqJSON, `"x-request-id":["req-id-1"]`)
	assert.Contains(t, reqJSON, `"x-trace-id":["trace-id-1"]`)
	assert.Contains(t, reqJSON, `"service_account_id":"sa-marker-xyz"`)

	respJSON := envelopeFor(t, entries[1].Context, "response")
	assert.Contains(t, respJSON, `"method":"/yandex.cloud.iam.v1.IamTokenService/CreateForServiceAccount"`)
	assert.NotContains(t, respJSON, `"status_code"`, "no status_code on success path")
	assert.NotContains(t, respJSON, `"error"`, "no error on success path")
}

func TestUnaryClientInterceptor_MasksSensitivePayloadFields(t *testing.T) {
	logger, logs := newObserverLogger(zapcore.DebugLevel)
	interceptor := UnaryClientInterceptor(logger)

	oauthTokenValue := "y0_AgAAAABabcdefghijklmnopqrstuvwxyz0123456789ABCDEFXYZ"
	iamTokenValue := "t1.AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA.BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB"

	req := &iampb.CreateIamTokenRequest{
		Identity: &iampb.CreateIamTokenRequest_YandexPassportOauthToken{
			YandexPassportOauthToken: oauthTokenValue,
		},
	}
	resp := &iampb.CreateIamTokenResponse{IamToken: iamTokenValue}

	err := interceptor(
		context.Background(),
		"/yandex.cloud.iam.v1.IamTokenService/Create",
		req,
		resp,
		nil,
		func(_ context.Context, _ string, _, _ any, _ *grpc.ClientConn, _ ...grpc.CallOption) error {
			return nil
		},
	)
	require.NoError(t, err)

	entries := logs.AllUntimed()
	require.Len(t, entries, 2)

	reqJSON := envelopeFor(t, entries[0].Context, "request")
	assert.NotContains(t, reqJSON, oauthTokenValue, "OAuth token must not leak into log")
	assert.Contains(t, reqJSON, `"yandex_passport_oauth_token":"`+hiddenPlaceholder+`"`,
		"sensitive payload field must be replaced with hidden placeholder")

	respJSON := envelopeFor(t, entries[1].Context, "response")
	assert.NotContains(t, respJSON, iamTokenValue, "issued IAM token must not leak into log")
	assert.Contains(t, respJSON, `"iam_token":"`+hiddenPlaceholder+`"`,
		"iam_token response field must be replaced with hidden placeholder")
}

func TestUnaryClientInterceptor_FiltersAuthorizationHeader(t *testing.T) {
	logger, logs := newObserverLogger(zapcore.DebugLevel)
	interceptor := UnaryClientInterceptor(logger)

	ctx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(
		"authorization", "Bearer secret-token-value",
		"x-request-id", "req-1",
	))

	err := interceptor(
		ctx,
		"/svc/Method",
		&iampb.RevokeIamTokenRequest{},
		&iampb.RevokeIamTokenResponse{},
		nil,
		func(_ context.Context, _ string, _, _ any, _ *grpc.ClientConn, _ ...grpc.CallOption) error {
			return nil
		},
	)
	require.NoError(t, err)

	reqJSON := envelopeFor(t, logs.AllUntimed()[0].Context, "request")
	assert.NotContains(t, reqJSON, "secret-token-value", "authorization header must not be logged")
	assert.NotContains(t, reqJSON, "authorization", "authorization key must be stripped from header map")
	assert.Contains(t, reqJSON, `"x-request-id":["req-1"]`)
}

func TestSanitizeSecrets_ShortPassportToken(t *testing.T) {
	// Regression: upstream antisecret requires {50,} chars after the y[0-6]_
	// prefix. Anonymized / slightly-shorter tokens (~49 chars) used to slip
	// through. Our local pattern uses a lower threshold; verify it catches a
	// token that is too short for the upstream regex.
	short := "y0__xCeuI08GMHdEyD3t4WCF7XXXXXXXXXXXXXXXXXXXXXXXXXXX"
	got := sanitizeSecrets(short)
	assert.NotContains(t, got, short, "short y0_ token must be masked")
	assert.True(t, strings.HasPrefix(got, "y*"), "sanitized form should keep the y prefix and mask body, got %q", got)
}

func TestUnaryClientInterceptor_LogsError(t *testing.T) {
	logger, logs := newObserverLogger(zapcore.DebugLevel)
	interceptor := UnaryClientInterceptor(logger)

	rpcErr := status.Error(codes.Unauthenticated, "OAuth token is invalid or expired")
	err := interceptor(
		context.Background(),
		"/yandex.cloud.iam.v1.IamTokenService/Create",
		&iampb.CreateIamTokenRequest{},
		&iampb.CreateIamTokenResponse{},
		nil,
		func(_ context.Context, _ string, _, _ any, _ *grpc.ClientConn, _ ...grpc.CallOption) error {
			return rpcErr
		},
	)
	assert.ErrorIs(t, err, rpcErr)

	entries := logs.AllUntimed()
	require.Len(t, entries, 2)
	assert.Equal(t, "Response IamTokenService/Create", entries[1].Message)

	respJSON := envelopeFor(t, entries[1].Context, "response")
	assert.Contains(t, respJSON, `"status_code":"UNAUTHENTICATED"`)
	assert.Contains(t, respJSON, `"code":16`)
	assert.Contains(t, respJSON, `"message":"OAuth token is invalid or expired"`)
	assert.NotContains(t, respJSON, `"payload"`, "no payload on failure path")
}
