package retry

import (
	"strconv"
	"time"

	"google.golang.org/grpc/codes"
)

//revive:enable:var-naming

type RetryConfig struct {
	mc              map[nameConfig]*methodConfig
	retryThrottling *retryThrottling
	waitForReady    bool
}

type RetryOption func(c *RetryConfig)

type grpcRetryPolicy struct {
	MethodConfig    []*methodConfig `json:"methodConfig"`
	RetryThrottling retryThrottling `json:"retryThrottling"`
	WaitForReady    bool            `json:"waitForReady"`
}

type methodConfig struct {
	NameConfig  []nameConfig `json:"name"`
	RetryPolicy retryPolicy  `json:"retryPolicy"`
}

type nameConfig struct {
	Service string `json:"service,omitempty"`
	Method  string `json:"method,omitempty"`
}

type retryPolicy struct {
	MaxAttempts          int      `json:"maxAttempts"`
	InitialBackoff       Duration `json:"initialBackoff"`
	MaxBackoff           Duration `json:"maxBackoff"`
	BackoffMultiplier    float64  `json:"backoffMultiplier"`
	RetryableStatusCodes []string `json:"retryableStatusCodes"`
}

type retryThrottling struct {
	MaxTokens  int     `json:"maxTokens"`
	TokenRatio float64 `json:"tokenRatio"`
}

type GRPCKeepAliveConfig struct {
	Time                time.Duration `yaml:"time" validate:"required"`
	Timeout             time.Duration `yaml:"timeout" validate:"required"`
	PermitWithoutStream bool          `yaml:"permit_without_stream"`
}

func DefaultRetryConfig() *RetryConfig {
	return &RetryConfig{}
}

func DefaultNameConfig() nameConfig {
	return nameConfig{}
}

func NewNameConfig(service, method string) nameConfig {
	return nameConfig{
		Service: service,
		Method:  method,
	}
}

func WithDefaultRetryConfig() RetryOption {
	return func(c *RetryConfig) { // nolint:staticcheck,ineffassign
		c = defaultRetryConfig() // nolint:ineffassign,staticcheck
	}
}

func WithRetries(nm nameConfig, n int) RetryOption {
	return func(c *RetryConfig) {
		if c == nil {
			c = defaultRetryConfig()
		}

		if _, ok := c.mc[nm]; !ok {
			c.mc[nm] = defaultMethodConfig()
			c.mc[nm].NameConfig = []nameConfig{nm}
		}

		c.mc[nm].RetryPolicy.MaxAttempts = n
	}
}

func WithRetryableStatusCodes(nm nameConfig, codes ...codes.Code) RetryOption {
	return func(c *RetryConfig) {
		if c == nil {
			c = defaultRetryConfig()
		}

		names := make([]string, len(codes))
		for i, code := range codes {
			names[i] = canonicalString(code)
		}

		if _, ok := c.mc[nm]; !ok {
			c.mc[nm] = defaultMethodConfig()
			c.mc[nm].NameConfig = []nameConfig{nm}
		}

		c.mc[nm].RetryPolicy.RetryableStatusCodes = names
	}
}

func canonicalString(c codes.Code) string {
	switch c {
	case codes.OK:
		return "OK"
	case codes.Canceled:
		return "CANCELLED"
	case codes.Unknown:
		return "UNKNOWN"
	case codes.InvalidArgument:
		return "INVALID_ARGUMENT"
	case codes.DeadlineExceeded:
		return "DEADLINE_EXCEEDED"
	case codes.NotFound:
		return "NOT_FOUND"
	case codes.AlreadyExists:
		return "ALREADY_EXISTS"
	case codes.PermissionDenied:
		return "PERMISSION_DENIED"
	case codes.ResourceExhausted:
		return "RESOURCE_EXHAUSTED"
	case codes.FailedPrecondition:
		return "FAILED_PRECONDITION"
	case codes.Aborted:
		return "ABORTED"
	case codes.OutOfRange:
		return "OUT_OF_RANGE"
	case codes.Unimplemented:
		return "UNIMPLEMENTED"
	case codes.Internal:
		return "INTERNAL"
	case codes.Unavailable:
		return "UNAVAILABLE"
	case codes.DataLoss:
		return "DATA_LOSS"
	case codes.Unauthenticated:
		return "UNAUTHENTICATED"
	default:
		return "CODE(" + strconv.FormatInt(int64(c), 10) + ")"
	}
}
