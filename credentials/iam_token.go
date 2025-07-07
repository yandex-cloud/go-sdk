package credentials

import (
	"context"
)

var _ NonExchangeableCredentials = (*IAMTokenCredentials)(nil)

// IAMTokenCredentials implements Credentials with IAM token as-is
// Read more on https://yandex.cloud/en-ru/docs/iam/concepts/authorization/iam-token
type IAMTokenCredentials struct {
	iamToken string
}

func (creds *IAMTokenCredentials) YandexCloudAPICredentials() {}

func (creds *IAMTokenCredentials) IAMToken(ctx context.Context) (*CredentialsToken, error) {
	return &CredentialsToken{Token: creds.iamToken}, nil
}

func IAMToken(iamToken string) NonExchangeableCredentials {
	return &IAMTokenCredentials{
		iamToken: iamToken,
	}
}
