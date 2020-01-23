package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/golang/protobuf/jsonpb"
	vision "github.com/yandex-cloud/go-genproto/yandex/cloud/ai/vision/v1"
	ycsdk "github.com/yandex-cloud/go-sdk"
)

const ImgText = "https://cloud.yandex.ru/docs-assets/ecb7431a-17b5-4633-80ca-ea6b461c947f/ru/_assets/vision/text-detection-line.png"
const ImgFace = "https://cloud.yandex.com/docs-assets/ecb7431a-17b5-4633-80ca-ea6b461c947f/en/_assets/vision/face-detection.jpg"

func getImg(url string) []byte {

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

	marshaler := jsonpb.Marshaler{
		Indent: " ",
	}

	// Connect
	sdk, err := ycsdk.Build(context.Background(), ycsdk.Config{
		Credentials: ycsdk.OAuthToken(token),
	})
	if err != nil {
		panic(err.Error())
	}
	defer sdk.Shutdown(context.Background())
	vi := sdk.AI().Vision()

	// Text recognition
	rez, err := vi.BatchAnalyze(context.Background(), &vision.BatchAnalyzeRequest{
		FolderId: folderID,
		AnalyzeSpecs: []*vision.AnalyzeSpec{
			&vision.AnalyzeSpec{
				Source: &vision.AnalyzeSpec_Content{
					Content: getImg(ImgText),
				},
				Features: []*vision.Feature{
					&vision.Feature{
						Type: vision.Feature_TEXT_DETECTION,
						Config: &vision.Feature_TextDetectionConfig{
							TextDetectionConfig: &vision.FeatureTextDetectionConfig{
								LanguageCodes: []string{"*"},
								Model:         "line",
							},
						},
					},
				},
			},
		},
	})
	if err != nil {
		panic(err.Error())
	}
	out, _ := marshaler.MarshalToString(rez)
	fmt.Println("TEXT DETECTION:", out)

	// Image classification
	rez, err = vi.BatchAnalyze(context.Background(), &vision.BatchAnalyzeRequest{
		FolderId: folderID,
		AnalyzeSpecs: []*vision.AnalyzeSpec{
			&vision.AnalyzeSpec{
				Source: &vision.AnalyzeSpec_Content{
					Content: getImg(ImgText),
				},
				Features: []*vision.Feature{
					&vision.Feature{
						Type: vision.Feature_CLASSIFICATION,
						Config: &vision.Feature_ClassificationConfig{
							ClassificationConfig: &vision.FeatureClassificationConfig{
								Model: "quality",
							},
						},
					},
					&vision.Feature{
						Type: vision.Feature_CLASSIFICATION,
						Config: &vision.Feature_ClassificationConfig{
							ClassificationConfig: &vision.FeatureClassificationConfig{
								Model: "moderation",
							},
						},
					},
				},
			},
		},
	})
	if err != nil {
		panic(err.Error())
	}
	out, _ = marshaler.MarshalToString(rez)
	fmt.Println("\nCLASSIFICATION: ", out)

	// Face detection
	rez, err = vi.BatchAnalyze(context.Background(), &vision.BatchAnalyzeRequest{
		FolderId: folderID,
		AnalyzeSpecs: []*vision.AnalyzeSpec{
			&vision.AnalyzeSpec{
				Source: &vision.AnalyzeSpec_Content{
					Content: getImg(ImgFace),
				},
				Features: []*vision.Feature{
					&vision.Feature{
						Type: vision.Feature_FACE_DETECTION,
					},
				},
			},
		},
	})
	if err != nil {
		panic(err.Error())
	}
	out, _ = marshaler.MarshalToString(rez)
	fmt.Println("\nFACE DETECTION: ", out)
}
