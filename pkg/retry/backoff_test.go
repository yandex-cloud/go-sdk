package retry

import (
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestBackoffLinearWithJitter(t *testing.T) {
	dto := float64(defaultLinearBackoffTimeout)
	toMin := time.Duration(dto * float64(1-defaultLinearBackoffJitter))
	toMax := time.Duration(dto * float64(1+defaultLinearBackoffJitter))

	backoff := BackoffLinearWithJitter(defaultLinearBackoffTimeout, defaultLinearBackoffJitter)

	for attempt := 0; attempt < 1000; attempt++ {
		to := backoff(attempt)
		res := to <= toMax && to >= toMin
		require.True(t, res)
	}
}

func TestBackoffExponentialWithJitter(t *testing.T) {
	maxBackoffTo := time.Duration(math.Pow(2, float64(20)) * float64(defaultExponentialBackoffBase))
	backoff := BackoffExponentialWithJitter(defaultExponentialBackoffBase, maxBackoffTo)

	for attempt := 0; attempt <= 20; attempt++ {
		to := backoff(attempt)
		maxTo := time.Duration(math.Pow(2, float64(attempt)) * float64(defaultExponentialBackoffBase))
		require.True(t, to <= maxTo)
	}

	for attempt := 20; attempt < 1000; attempt++ {
		to := backoff(attempt)
		require.True(t, to <= maxBackoffTo)
	}
}
