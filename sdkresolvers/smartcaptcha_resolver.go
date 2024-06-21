package sdkresolvers

import (
	"context"

	"google.golang.org/grpc"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/smartcaptcha/v1"
	ycsdk "github.com/yandex-cloud/go-sdk"
)

type captchaResolver struct {
	BaseNameResolver
}

func CaptchaResolver(name string, opts ...ResolveOption) ycsdk.Resolver {
	return &captchaResolver{
		BaseNameResolver: NewBaseNameResolver(name, "captcha", opts...),
	}
}

func (r *captchaResolver) Run(ctx context.Context, sdk *ycsdk.SDK, opts ...grpc.CallOption) error {
	err := r.ensureFolderID()
	if err != nil {
		return err
	}

	resp, err := sdk.SmartCaptcha().Captcha().List(ctx, &smartcaptcha.ListCaptchasRequest{
		FolderId: r.FolderID(),
		// TODO: better to use Filter("name"), but now it's not supported now
	}, opts...)
	return r.findName(resp.GetResources(), err)
}
