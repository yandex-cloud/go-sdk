// Code generated by sdkgen. DO NOT EDIT.

package text_generation

import (
	"context"

	"google.golang.org/grpc"
)

// FoundationModelsTextGeneration provides access to "text_generation" component of Yandex.Cloud
type FoundationModelsTextGeneration struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// NewFoundationModelsTextGeneration creates instance of FoundationModelsTextGeneration
func NewFoundationModelsTextGeneration(g func(ctx context.Context) (*grpc.ClientConn, error)) *FoundationModelsTextGeneration {
	return &FoundationModelsTextGeneration{g}
}

// TextGeneration gets TextGenerationService client
func (f *FoundationModelsTextGeneration) TextGeneration() *TextGenerationServiceClient {
	return &TextGenerationServiceClient{getConn: f.getConn}
}

// TextGenerationAsync gets TextGenerationAsyncService client
func (f *FoundationModelsTextGeneration) TextGenerationAsync() *TextGenerationAsyncServiceClient {
	return &TextGenerationAsyncServiceClient{getConn: f.getConn}
}

// Tokenizer gets TokenizerService client
func (f *FoundationModelsTextGeneration) Tokenizer() *TokenizerServiceClient {
	return &TokenizerServiceClient{getConn: f.getConn}
}
