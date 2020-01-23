package main

import (
	"context"
	"fmt"
	"os"

	translate "github.com/yandex-cloud/go-genproto/yandex/cloud/ai/translate/v2"
	ycsdk "github.com/yandex-cloud/go-sdk"
)

// Test translate
func main() {
	var err error

	// get environments
	token := os.Getenv("token")
	folderID := os.Getenv("folderid")

	// Connect
	sdk, err := ycsdk.Build(context.Background(), ycsdk.Config{
		Credentials: ycsdk.OAuthToken(token),
	})
	if err != nil {
		panic(err.Error())
	}
	defer sdk.Shutdown(context.Background())

	// Translate
	tr := sdk.AI().Translate()

	// Get list of langs
	var listRes *translate.ListLanguagesResponse
	listRes, err = tr.ListLanguages(context.Background(), &translate.ListLanguagesRequest{
		FolderId: folderID,
	})
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("=== List of languages")
	for _, itm := range listRes.GetLanguages() {
		fmt.Printf("Code: %v; Name:%v\n", itm.GetCode(), itm.GetName())
	}

	// Test detect language
	var detRes *translate.DetectLanguageResponse
	detRes, err = tr.DetectLanguage(context.Background(), &translate.DetectLanguageRequest{
		Text:     "Привет мир",
		FolderId: folderID,
	})
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("=== Detect results")
	fmt.Println("Source lng:", detRes.GetLanguageCode())

	// Test translate
	var trRes *translate.TranslateResponse
	trRes, err = tr.Translate(context.Background(), &translate.TranslateRequest{
		TargetLanguageCode: "en",
		Texts:              []string{"Привет мир"},
		FolderId:           folderID,
	})
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("=== Translate results")
	for _, itm := range trRes.GetTranslations() {
		fmt.Println("Source lng:", itm.GetDetectedLanguageCode(), "; Result:", itm.GetText())
	}
}
