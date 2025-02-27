package retry

import (
	"strings"
	"time"

	"google.golang.org/grpc/codes"
)

//revive:enable:var-naming

type RetryConfig struct {
	mc map[nameConfig]*methodConfig
}

type RetryOption func(c *RetryConfig)

type grpcRetryPolicy struct {
	MethodConfig []*methodConfig `json:"methodConfig"`
}

type methodConfig struct {
	NameConfig      []nameConfig    `json:"name"`
	RetryPolicy     retryPolicy     `json:"retryPolicy"`
	RetryThrottling retryThrottling `json:"retryThrottling"`
	WaitForReady    bool            `json:"waitForReady"`
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
	return func(c *RetryConfig) {
		c = defaultRetryConfig()
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
			names[i] = strings.ToUpper(code.String())
		}

		if _, ok := c.mc[nm]; !ok {
			c.mc[nm] = defaultMethodConfig()
			c.mc[nm].NameConfig = []nameConfig{nm}
		}

		c.mc[nm].RetryPolicy.RetryableStatusCodes = names
	}
}
