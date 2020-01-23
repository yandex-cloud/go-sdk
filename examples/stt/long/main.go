package main

import (
	"context"
	"fmt"
	"time"

	"github.com/golang/protobuf/ptypes"
	recog "github.com/yandex-cloud/go-genproto/yandex/cloud/ai/stt/v2"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	ycsdk "github.com/yandex-cloud/go-sdk"
	"github.com/yandex-cloud/go-sdk/iamkey"
)

// Test speech recognition
func main() {
	var err error

	// Credentials
	iam, err := iamkey.ReadFromJSONFile("../../../test_data/service_account_key.json")
	if err != nil {
		panic(err.Error())
	}

	var credentials ycsdk.Credentials
	credentials, err = ycsdk.ServiceAccountKey(iam)
	if err != nil {
		panic(err.Error())
	}

	// Connect
	var sdk *ycsdk.SDK
	sdk, err = ycsdk.Build(context.Background(), ycsdk.Config{
		Credentials: credentials,
	})
	if err != nil {
		panic(err.Error())
	}
	defer sdk.Shutdown(context.Background())

	//=== Stream speech recognition
	sp := sdk.AI().STT()

	// Request
	fmt.Println("Request")

	res, err := sp.LongRunningRecognition(context.Background(), &recog.LongRunningRecognitionRequest{
		Audio: &recog.RecognitionAudio{
			AudioSource: &recog.RecognitionAudio_Uri{
				Uri: `https://storage.yandexcloud.net/speechkit/speech.ogg`,
			},
		},
		Config: &recog.RecognitionConfig{
			Specification: &recog.RecognitionSpec{
				LanguageCode: "ru-RU",
			},
		},
	})
	if err != nil {
		panic(err.Error())
	}

	OpID := res.GetId()
	fmt.Print("Wait response: ")
	for !res.GetDone() {
		fmt.Print("*")

		time.Sleep(1 * time.Second)

		// Response
		res, err = sdk.Operation().Get(context.Background(), &operation.GetOperationRequest{
			OperationId: OpID,
		})
		if err != nil {
			panic(err.Error())
		}

		if res.GetDone() {
			break
		}
	}

	var rez recog.LongRunningRecognitionResponse
	err = ptypes.UnmarshalAny(res.GetResponse(), &rez)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("\nResponse done: ")

	for _, itm := range rez.GetChunks() {

		aitm := itm.GetAlternatives()
		if len(aitm) < 1 {
			continue
		}

		fmt.Println(aitm[0].GetText())
	}
}
