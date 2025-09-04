package dial

import (
	"fmt"
	"runtime/debug"
)

func UserAgent() string {
	cloudUserAgent := "yandex-cloud/go-sdk"

	version := "unknown"
	if build, ok := debug.ReadBuildInfo(); ok {
		if build.Main.Version != "" {
			version = build.Main.Version
		}
	}

	return fmt.Sprintf("%s/%s", cloudUserAgent, version)
}
