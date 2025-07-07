package credentials

// OAuthToken returns API credentials for user Yandex Passport OAuth token, that can be received
// on page https://oauth.yandex.ru/authorize?response_type=token&client_id=1a6990aa636648e9b2ef855fa7bec2fb
// See https://cloud.yandex.ru/docs/iam/concepts/authorization/oauth-token for details.
func OAuthToken(token string) ExchangeableCredentials {
	return exchangeableCredentialsFunc(func() (*CredentialsTokenRequest, error) {
		return &CredentialsTokenRequest{
			Identity: CredentialsIdentityYandexPassportOauthToken,
			Token:    token,
		}, nil
	})
}
