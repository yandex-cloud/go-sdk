// Copyright (c) 2018 Yandex LLC. All rights reserved.
// Author: Vladimir Skipor <skipor@yandex-team.ru>

package ycsdk

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOAuthToken(t *testing.T) {
	const token = "AAAA00000000000000000000000000000000000"
	creds := OAuthToken(token)
	iamTokenReq, expiration, err := creds.(ExchangeableCredentials).IAMTokenRequest()
	require.NoError(t, err)
	assert.Equal(t, DefaultIAMTokenRefreshInterval, expiration)
	assert.Equal(t, token, iamTokenReq.GetYandexPassportOauthToken())
}
