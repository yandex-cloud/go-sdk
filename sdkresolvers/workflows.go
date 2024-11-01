package sdkresolvers

import (
	"context"

	"google.golang.org/grpc"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/serverless/workflows/v1"
	ycsdk "github.com/yandex-cloud/go-sdk"
)

type workflowsResolver struct {
	BaseNameResolver
}

func WorkflowsResolver(name string, opts ...ResolveOption) ycsdk.Resolver {
	return &workflowsResolver{
		BaseNameResolver: NewBaseNameResolver(name, "workflow", opts...),
	}
}

func (r *workflowsResolver) Run(ctx context.Context, sdk *ycsdk.SDK, opts ...grpc.CallOption) error {
	err := r.ensureFolderID()
	if err != nil {
		return err
	}

	resp, err := sdk.Serverless().Workflow().Workflow().List(ctx, &workflows.ListWorkflowsRequest{
		FolderId: r.FolderID(),
		Filter:   CreateResolverFilter("name", r.Name),
		PageSize: DefaultResolverPageSize,
	}, opts...)
	return r.findName(resp.GetWorkflows(), err)
}
