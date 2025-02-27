package retry

import "time"

func defaultNameConfig() nameConfig {
	return nameConfig{}
}

func defaultRetryConfig() *RetryConfig {
	return &RetryConfig{
		mc: map[nameConfig]*methodConfig{
			// default value for all services and methods
			defaultNameConfig(): defaultMethodConfig(),
		},
	}
}

func defaultMethodConfig() *methodConfig {
	return &methodConfig{
		NameConfig: []nameConfig{{}},
		RetryPolicy: retryPolicy{
			MaxAttempts:          4,
			InitialBackoff:       Duration(time.Millisecond * 100),
			MaxBackoff:           Duration(time.Second * 20),
			BackoffMultiplier:    2,
			RetryableStatusCodes: []string{"UNAVAILABLE"},
		},
		RetryThrottling: retryThrottling{
			MaxTokens:  100,
			TokenRatio: 0.1,
		},
		WaitForReady: true,
	}
}
