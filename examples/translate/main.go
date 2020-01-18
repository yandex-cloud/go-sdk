// Author: Ivan Zaytsev  <ivan@jad.ru>

package main

import (
	"context"
	"fmt"

	translate "github.com/yandex-cloud/go-genproto/yandex/cloud/ai/translate/v2"
	ycsdk "github.com/yandex-cloud/go-sdk"
)

var Token = "<OAuth-token>"
var FolderID = "<FolderID>"

// Test translate
func main() {
	var err error

	Credentials := ycsdk.OAuthToken(Token)

	// Connect
	var sdk *ycsdk.SDK
	sdk, err = ycsdk.Build(context.Background(), ycsdk.Config{
		Credentials: Credentials,
	})
	if err != nil {
		panic(err.Error())
	}
	defer sdk.Shutdown(context.Background())

	// Translate
	tr := sdk.Translate().Translate()

	// Get list of langs
	var ListRes *translate.ListLanguagesResponse
	ListRes, err = tr.ListLanguages(context.Background(), &translate.ListLanguagesRequest{
		FolderId: FolderID,
	})
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("=== List of languages")
	for _, itm := range ListRes.GetLanguages() {
		fmt.Printf("Code: %v; Name:%v\n", itm.GetCode(), itm.GetName())
	}

	// Test detect language
	var DetRes *translate.DetectLanguageResponse
	DetRes, err = tr.DetectLanguage(context.Background(), &translate.DetectLanguageRequest{
		Text:     "Привет мир",
		FolderId: FolderID,
	})
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("=== Detect results")
	fmt.Println("Source lng:", DetRes.GetLanguageCode())

	// Test translate
	var TrRes *translate.TranslateResponse
	TrRes, err = tr.Translate(context.Background(), &translate.TranslateRequest{
		//SourceLanguageCode: "ru",
		TargetLanguageCode: "en",
		Texts:              []string{"Привет мир"},
		FolderId:           FolderID,
	})
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("=== Translate results")
	for _, itm := range TrRes.GetTranslations() {
		fmt.Println("Source lng:", itm.GetDetectedLanguageCode(), "; Result:", itm.GetText())
	}
}
