# Yandex.Cloud Go SDK

[![GoDoc](https://godoc.org/github.com/yandex-cloud/go-sdk?status.svg)](https://godoc.org/github.com/yandex-cloud/go-sdk)
[![CircleCI](https://circleci.com/gh/yandex-cloud/go-sdk.svg?style=shield)](https://circleci.com/gh/yandex-cloud/go-sdk)

Go SDK for Yandex.Cloud services.

**NOTE:** SDK is under development, and may make
backwards-incompatible changes.

## Installation

```bash
go get github.com/yandex-cloud/go-sdk
```

## Example usages

### Initializing SDK

```go
sdk, err := ycsdk.Build(ctx, ycsdk.Config{
	Credentials: ycsdk.OAuthToken(token),
})
if err != nil {
	log.Fatal(err)
}
```

### Retries

SDK provide built-in retry policy, that supports [exponential backoff and jitter](https://aws.amazon.com/ru/blogs/architecture/exponential-backoff-and-jitter/), and also [retry budget](https://github.com/grpc/proposal/blob/master/A6-client-retries.md#throttling-retry-attempts-and-hedged-rpcs). 
It's necessary to avoid retry amplification.

```go
import (
	...
	ycsdk "github.com/yandex-cloud/go-sdk"
	"github.com/yandex-cloud/go-sdk/pkg/retry/v1"
)

...

retriesDialOption, err := retry.DefaultRetryDialOption()
if err != nil {
	log.Fatal(err)
}

_, err = ycsdk.Build(
	ctx,
	ycsdk.Config{
		Credentials: ycsdk.OAuthToken(*token),
	},
	retriesDialOption,
)
```

SDK provide different modes for retry throttling policy:

* `persistent` is suitable when you use SDK in any long-lived application, when SDK instance will live long enough for manage budget;
* `temporary` is suitable when you use SDK in any short-lived application, e.g. scripts or CI/CD.

By default, SDK will use temporary mode, but you can change it through functional option.


### More examples

More examples can be found in [examples dir](examples).
