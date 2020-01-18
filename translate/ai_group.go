package translate

import (
	"context"
	"google.golang.org/grpc"
)

// Translate -
type Translate struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// NewTranslate creates instance of Translate
func NewTranslate(g func(ctx context.Context) (*grpc.ClientConn, error)) *Translate {
	return &Translate{g}
}

// Translate gets Translate client
func (c *Translate) Translate() *TranslateServiceClient {
	return &TranslateServiceClient{getConn: c.getConn}
}
