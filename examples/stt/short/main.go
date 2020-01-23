package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	recog "github.com/yandex-cloud/go-genproto/yandex/cloud/ai/stt/v2"
	ycsdk "github.com/yandex-cloud/go-sdk"
)

const cBufSize = 4000

func getData(url string) []byte {

	resp, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		panic(fmt.Errorf("Invalid return status: %v", resp.StatusCode))
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}

	return data
}

func main() {
	var err error

	// get environments
	token := os.Getenv("token")
	folderID := os.Getenv("folderid")

	ctx := context.Background()

	credentials := ycsdk.OAuthToken(token)

	data := getData("https://storage.yandexcloud.net/speechkit/speech.pcm")

	// Connect
	var sdk *ycsdk.SDK
	sdk, err = ycsdk.Build(ctx, ycsdk.Config{
		Credentials: credentials,
	})
	if err != nil {
		panic(err.Error())
	}
	defer sdk.Shutdown(ctx)

	//=== Stream speech recognition
	sp := sdk.AI().STT()

	err = sp.StreamingRecognitionConfig(ctx, &recog.RecognitionConfig{
		Specification: &recog.RecognitionSpec{
			AudioEncoding:   recog.RecognitionSpec_LINEAR16_PCM,
			SampleRateHertz: 8000,
			LanguageCode:    "ru-RU",
			ProfanityFilter: true,
			PartialResults:  true,
		},
		FolderId: folderID,
	})
	if err != nil {
		panic(err.Error())
	}

	go func() {
		for {
			res, err := sp.StreamingRecognitionReceive(ctx)
			if err != nil {
				panic(err.Error())
			}

			fmt.Println(res.String())
		}
	}()

	i := 0
	for {
		if i >= len(data) {
			break
		}

		e := i + cBufSize
		if e >= len(data) {
			e = len(data)
		}

		fmt.Print("+")

		err = sp.StreamingRecognitionSend(ctx, &recog.StreamingRecognitionRequest{
			StreamingRequest: &recog.StreamingRecognitionRequest_AudioContent{
				AudioContent: data[i:e],
			},
		})
		if err != nil {
			panic(err.Error())
		}

		time.Sleep(100 * time.Millisecond)
		i = i + cBufSize
	}

	for {
		time.Sleep(1000 * time.Millisecond)
		fmt.Print("*")
	}
}
