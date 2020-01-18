package translate

import (
	"context"

	translate "github.com/yandex-cloud/go-genproto/yandex/cloud/ai/translate/v2"
	"google.golang.org/grpc"
)

// TranslateServiceClient -
type TranslateServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// Translate - Translates the text to the specified language.
func (c *TranslateServiceClient) Translate(ctx context.Context, in *translate.TranslateRequest, opts ...grpc.CallOption) (*translate.TranslateResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return translate.NewTranslationServiceClient(conn).Translate(ctx, in, opts...)
}

// DetectLanguage - Detects the language of the text.
func (c *TranslateServiceClient) DetectLanguage(ctx context.Context, in *translate.DetectLanguageRequest, opts ...grpc.CallOption) (*translate.DetectLanguageResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return translate.NewTranslationServiceClient(conn).DetectLanguage(ctx, in, opts...)
}

// ListLanguages - Retrieves the list of supported languages.
func (c *TranslateServiceClient) ListLanguages(ctx context.Context, in *translate.ListLanguagesRequest, opts ...grpc.CallOption) (*translate.ListLanguagesResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return translate.NewTranslationServiceClient(conn).ListLanguages(ctx, in, opts...)
}
