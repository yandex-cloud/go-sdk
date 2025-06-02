package ycsdk

import (
	"github.com/yandex-cloud/go-sdk/gen/gitlab"
)

const GitlabServiceID = "gitlab"

func (sdk *SDK) Gitlab() *gitlab.Gitlab {
	return gitlab.NewGitlab(sdk.getConn(GitlabServiceID))
}
