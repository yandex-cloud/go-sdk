// Capturing audio from the mic
// Alternatively, gst-launch can be used to capture audio from the mic. For example:

// gst-launch-1.0 -v pulsesrc ! audioconvert ! audioresample ! audio/x-raw,channels=1,rate=16000 ! filesink location=/dev/stdout | livecaption

package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	recog "github.com/yandex-cloud/go-genproto/yandex/cloud/ai/stt/v2"
	ycsdk "github.com/yandex-cloud/go-sdk"
)

// Test speech recognition
func main() {
	var err error

	// get environments
	token := os.Getenv("token")
	folderID := os.Getenv("folderid")

	ctx := context.Background()

	credentials := ycsdk.OAuthToken(token)

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

	// Config
	err = sp.StreamingRecognitionConfig(ctx, &recog.RecognitionConfig{
		Specification: &recog.RecognitionSpec{
			AudioEncoding:   recog.RecognitionSpec_LINEAR16_PCM,
			SampleRateHertz: 16000,
			LanguageCode:    "ru-RU",
		},
		FolderId: folderID,
	})
	if err != nil {
		panic(err.Error())
	}

	go func() {
		// Pipe stdin to the API.
		buf := make([]byte, 1024)
		for {
			n, err := os.Stdin.Read(buf)
			if n > 0 {
				if err := sp.StreamingRecognitionSend(ctx, &recog.StreamingRecognitionRequest{
					StreamingRequest: &recog.StreamingRecognitionRequest_AudioContent{
						AudioContent: buf[:n],
					},
				}); err != nil {
					log.Printf("Could not send audio: %v", err)
				}
			}
			if err == io.EOF {
				// Nothing else to pipe, close the stream.
				log.Fatalf("Could not close stream: %v", err)
				return
			}
			if err != nil {
				log.Printf("Could not read from stdin: %v", err)
				continue
			}
		}
	}()

	for {
		resp, err := sp.StreamingRecognitionReceive(ctx)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Cannot stream results: %v", err)
		}
		for _, itm := range resp.GetChunks() {
			aitm := itm.GetAlternatives()
			if len(aitm) < 1 {
				continue
			}

			fmt.Println(aitm[0].GetText())
		}
	}
}
