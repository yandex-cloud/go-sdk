package authentication

import (
	"context"
	"fmt"

	"go.uber.org/zap"

	iampb "bb.yandex-team.ru/cloud/cloud-go/genproto/privateapi/yandex/cloud/priv/iam/v1"
	"github.com/yandex-cloud/go-sdk/v2/credentials"
	"github.com/yandex-cloud/go-sdk/v2/pkg/endpoints"
	"github.com/yandex-cloud/go-sdk/v2/pkg/errors"
	"github.com/yandex-cloud/go-sdk/v2/pkg/transport"
	iamsdk "github.com/yandex-cloud/go-sdk/v2/private/services/iam/v1"
)

// PrivateAuthenticatorImpl provides authentication functionality for private APIs
type PrivateAuthenticatorImpl struct {
	creds          credentials.Credentials
	iamTokenClient iamsdk.IamTokenClient
	logger         *zap.Logger
}

var _ Authenticator = &PrivateAuthenticatorImpl{}

// NewPrivateAuthenticatorFromEndpoint creates a new PrivateAuthenticatorImpl using the provided credentials and endpoint
func NewPrivateAuthenticatorFromEndpoint(logger *zap.Logger, creds credentials.Credentials, endpoint *endpoints.Endpoint) (*PrivateAuthenticatorImpl, error) {
	return NewPrivateAuthenticator(logger, creds, iamsdk.NewIamTokenClient(transport.NewSingleConnector(endpoint.Addr, endpoint.DialOptions...))), nil
}

// NewAuthenticator creates and returns a new instance of PrivateAuthenticatorImpl
func NewPrivateAuthenticator(logger *zap.Logger, creds credentials.Credentials, iamTokenClient iamsdk.IamTokenClient) *PrivateAuthenticatorImpl {
	return &PrivateAuthenticatorImpl{
		creds:          creds,
		iamTokenClient: iamTokenClient,
		logger:         logger,
	}
}

// CreateIAMToken generates an IAM token using the provided credentials
func (a *PrivateAuthenticatorImpl) CreateIAMToken(ctx context.Context) (IamToken, error) {
	a.logger.Debug("Creating IAM token")
	switch c := a.creds.(type) {
	case credentials.ExchangeableCredentials:
		a.logger.Debug("Using exchangeable credentials")
		return a.createTokenFromExchangeable(ctx, c)
	case credentials.NonExchangeableCredentials:
		a.logger.Debug("Using non-exchangeable credentials")
		return a.createTokenFromNonExchangeable(ctx, c)
	default:
		err := fmt.Errorf("unsupported credentials type %T", c)
		a.logger.Error("Failed to create IAM token", zap.Error(err))
		return nil, &errors.AuthError{Err: err}
	}
}

// createTokenFromExchangeable generates an IAM token using exchangeable credentials
func (a *PrivateAuthenticatorImpl) createTokenFromExchangeable(ctx context.Context, creds credentials.ExchangeableCredentials) (IamToken, error) {
	a.logger.Debug("Generating IAM token request")
	req, err := creds.IAMTokenRequest()
	if err != nil {
		a.logger.Error("Failed to create IAM token request", zap.Error(err))
		return nil, &errors.AuthError{Err: err}
	}

	tokenReq := createIamTokenRequestFromCredentialToken(req)
	if tokenReq == nil {
		err := fmt.Errorf("invalid identity type")
		a.logger.Error("Failed to create token request from credentials", zap.Error(err))
		return nil, &errors.AuthError{Err: err}
	}

	a.logger.Debug("Creating IAM token from request")
	tokenResp, err := a.iamTokenClient.Create(ctx, tokenReq)
	if err != nil {
		a.logger.Error("Failed to create IAM token", zap.Error(err))
		return nil, &errors.AuthError{Err: err}
	}

	a.logger.Debug("Successfully created IAM token")
	return NewIamToken(tokenResp.IamToken, tokenResp.ExpiresAt.AsTime()), nil
}

// createTokenFromNonExchangeable generates an IAM token using non-exchangeable credentials
func (a *PrivateAuthenticatorImpl) createTokenFromNonExchangeable(ctx context.Context, creds credentials.NonExchangeableCredentials) (IamToken, error) {
	a.logger.Debug("Retrieving IAM token from non-exchangeable credentials")
	tokenResp, err := creds.IAMToken(ctx)
	if err != nil {
		a.logger.Error("Failed to get IAM token from non-exchangeable credentials", zap.Error(err))
		return nil, &errors.AuthError{Err: err}
	}
	a.logger.Debug("Successfully retrieved IAM token")
	return NewIamToken(tokenResp.Token, tokenResp.ExpiresAt), nil
}

// CreateIAMTokenForServiceAccount generates a new IAM token for the provided service account ID
func (a *PrivateAuthenticatorImpl) CreateIAMTokenForServiceAccount(ctx context.Context, serviceAccountID string) (IamToken, error) {
	a.logger.Debug("Creating IAM token for service account", zap.String("serviceAccountID", serviceAccountID))
	tokenResp, err := a.iamTokenClient.CreateForServiceAccount(ctx,
		&iampb.CreateIamTokenForServiceAccountRequest{
			ServiceAccountId: serviceAccountID,
		},
	)
	if err != nil {
		a.logger.Error("Failed to create IAM token for service account",
			zap.String("serviceAccountID", serviceAccountID),
			zap.Error(err))
		return nil, &errors.AuthError{Err: err}
	}
	a.logger.Debug("Successfully created IAM token for service account",
		zap.String("serviceAccountID", serviceAccountID))
	return NewIamToken(tokenResp.IamToken, tokenResp.ExpiresAt.AsTime()), nil
}

func (a *PrivateAuthenticatorImpl) InjectLogger(logger *zap.Logger) {
	a.logger = logger
}

// createIamTokenRequestFromCredentialToken converts a CredentialsTokenRequest into an iampb.CreateIamTokenRequest
func createIamTokenRequestFromCredentialToken(req *credentials.CredentialsTokenRequest) *iampb.CreateIamTokenRequest {
	switch req.Identity {
	case credentials.CredentialsIdentityYandexPassportOauthToken:
		return &iampb.CreateIamTokenRequest{
			Identity: &iampb.CreateIamTokenRequest_YandexPassportOauthToken{
				YandexPassportOauthToken: req.Token},
		}
	case credentials.CredentialsIdentityJWT:
		return &iampb.CreateIamTokenRequest{
			Identity: &iampb.CreateIamTokenRequest_Jwt{Jwt: req.Token}}
	}
	return nil
}
