package ycsdk

import (
	"github.com/yandex-cloud/go-sdk/gen/ai/stt"
	"github.com/yandex-cloud/go-sdk/gen/ai/translate"
	"github.com/yandex-cloud/go-sdk/gen/ai/vision"
)

const (
	STTServiceID       Endpoint = "ai-stt"
	//SpeechKitServiceID Endpoint = "ai-speechkit"
	TranslateServiceID Endpoint = "ai-translate"
	VisionServiceID    Endpoint = "ai-vision"
)

// AI wraps AI services
func (sdk *SDK) AI() *AI {
	return &AI{sdk: sdk}
}

// AI is a Yandex.Cloud AI service
type AI struct {
	sdk *SDK
}

// Translate gets Translate client
func (ai *AI) Translate() *translate.Translate {
	return translate.NewTranslate(ai.sdk.getConn(TranslateServiceID))
}

// STT gets Speech recognition client
func (ai *AI) STT() *stt.STT {
	return stt.NewSTT(ai.sdk.getConn(STTServiceID))
}

// SpeechKit gets Speech synthesis client
//TODO: Speech Synthesis is not implemented
// func (ai *AI) SpeechKit() *stt.SpeechKit {
// 	return stt.NewSpeechKit(ai.sdk.getConn(SpeechKitServiceID))
// }

// Vision gets Vision client
func (ai *AI) Vision() *vision.Vision {
	return vision.NewVision(ai.sdk.getConn(VisionServiceID))
}
