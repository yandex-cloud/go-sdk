package translate

import (
	"context"

	trans "github.com/yandex-cloud/go-genproto/yandex/cloud/ai/translate/v2"
	"google.golang.org/grpc"
)

// Translate is a functions.Translate with
// lazy GRPC connection initialization.
type Translate struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// NewTranslate creates instance of translate
func NewTranslate(g func(ctx context.Context) (*grpc.ClientConn, error)) *Translate {
	return &Translate{g}
}

// Translate implements trans.Translate
func (t *Translate) Translate(ctx context.Context, in *trans.TranslateRequest, opts ...grpc.CallOption) (*trans.TranslateResponse, error) {
	conn, err := t.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return trans.NewTranslationServiceClient(conn).Translate(ctx, in, opts...)
}

// DetectLanguage implements trans.DetectLanguage
func (t *Translate) DetectLanguage(ctx context.Context, in *trans.DetectLanguageRequest, opts ...grpc.CallOption) (*trans.DetectLanguageResponse, error) {
	conn, err := t.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return trans.NewTranslationServiceClient(conn).DetectLanguage(ctx, in, opts...)
}

// ListLanguages implements trans.ListLanguages
func (t *Translate) ListLanguages(ctx context.Context, in *trans.ListLanguagesRequest, opts ...grpc.CallOption) (*trans.ListLanguagesResponse, error) {
	conn, err := t.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return trans.NewTranslationServiceClient(conn).ListLanguages(ctx, in, opts...)
}
