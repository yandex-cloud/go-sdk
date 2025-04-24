package dial

import (
	"fmt"
	"runtime/debug"
)

func UserAgent() string {
	cloudUserAgent := "yandex-cloud/go-sdk"

	build, _ := debug.ReadBuildInfo()
	version := "unknown"
	if build.Main.Version != "" {
		version = build.Main.Version
	}

	return fmt.Sprintf("%s/%s", cloudUserAgent, version)
}
