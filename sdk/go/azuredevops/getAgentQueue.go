// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azuredevops/sdk/v2/go/azuredevops/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing Agent Queue within Azure DevOps.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-azuredevops/sdk/v2/go/azuredevops"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleProject, err := azuredevops.NewProject(ctx, "exampleProject", &azuredevops.ProjectArgs{
//				WorkItemTemplate: pulumi.String("Agile"),
//				VersionControl:   pulumi.String("Git"),
//				Visibility:       pulumi.String("private"),
//				Description:      pulumi.String("Managed by Terraform"),
//			})
//			if err != nil {
//				return err
//			}
//			exampleAgentQueue := azuredevops.GetAgentQueueOutput(ctx, azuredevops.GetAgentQueueOutputArgs{
//				ProjectId: exampleProject.ID(),
//				Name:      pulumi.String("Example Agent Queue"),
//			}, nil)
//			ctx.Export("name", exampleAgentQueue.ApplyT(func(exampleAgentQueue azuredevops.GetAgentQueueResult) (*string, error) {
//				return &exampleAgentQueue.Name, nil
//			}).(pulumi.StringPtrOutput))
//			ctx.Export("poolId", exampleAgentQueue.ApplyT(func(exampleAgentQueue azuredevops.GetAgentQueueResult) (*int, error) {
//				return &exampleAgentQueue.AgentPoolId, nil
//			}).(pulumi.IntPtrOutput))
//			return nil
//		})
//	}
//
// ```
// ## Relevant Links
//
// - [Azure DevOps Service REST API 7.0 - Agent Queues - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/distributedtask/queues/get?view=azure-devops-rest-7.0)
func GetAgentQueue(ctx *pulumi.Context, args *GetAgentQueueArgs, opts ...pulumi.InvokeOption) (*GetAgentQueueResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetAgentQueueResult
	err := ctx.Invoke("azuredevops:index/getAgentQueue:getAgentQueue", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAgentQueue.
type GetAgentQueueArgs struct {
	// Name of the Agent Queue.
	Name string `pulumi:"name"`
	// The Project Id.
	ProjectId string `pulumi:"projectId"`
}

// A collection of values returned by getAgentQueue.
type GetAgentQueueResult struct {
	// Agent pool identifier to which the agent queue belongs.
	AgentPoolId int `pulumi:"agentPoolId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The name of the agent queue.
	Name string `pulumi:"name"`
	// Project identifier to which the agent queue belongs.
	ProjectId string `pulumi:"projectId"`
}

func GetAgentQueueOutput(ctx *pulumi.Context, args GetAgentQueueOutputArgs, opts ...pulumi.InvokeOption) GetAgentQueueResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAgentQueueResult, error) {
			args := v.(GetAgentQueueArgs)
			r, err := GetAgentQueue(ctx, &args, opts...)
			var s GetAgentQueueResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAgentQueueResultOutput)
}

// A collection of arguments for invoking getAgentQueue.
type GetAgentQueueOutputArgs struct {
	// Name of the Agent Queue.
	Name pulumi.StringInput `pulumi:"name"`
	// The Project Id.
	ProjectId pulumi.StringInput `pulumi:"projectId"`
}

func (GetAgentQueueOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAgentQueueArgs)(nil)).Elem()
}

// A collection of values returned by getAgentQueue.
type GetAgentQueueResultOutput struct{ *pulumi.OutputState }

func (GetAgentQueueResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAgentQueueResult)(nil)).Elem()
}

func (o GetAgentQueueResultOutput) ToGetAgentQueueResultOutput() GetAgentQueueResultOutput {
	return o
}

func (o GetAgentQueueResultOutput) ToGetAgentQueueResultOutputWithContext(ctx context.Context) GetAgentQueueResultOutput {
	return o
}

// Agent pool identifier to which the agent queue belongs.
func (o GetAgentQueueResultOutput) AgentPoolId() pulumi.IntOutput {
	return o.ApplyT(func(v GetAgentQueueResult) int { return v.AgentPoolId }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAgentQueueResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAgentQueueResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the agent queue.
func (o GetAgentQueueResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetAgentQueueResult) string { return v.Name }).(pulumi.StringOutput)
}

// Project identifier to which the agent queue belongs.
func (o GetAgentQueueResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v GetAgentQueueResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAgentQueueResultOutput{})
}
