// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azuredevops/sdk/v3/go/azuredevops/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing Build Definition.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-azuredevops/sdk/v3/go/azuredevops"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			example, err := azuredevops.LookupProject(ctx, &azuredevops.LookupProjectArgs{
//				Name: pulumi.StringRef("Example Project"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			exampleGetBuildDefinition, err := azuredevops.LookupBuildDefinition(ctx, &azuredevops.LookupBuildDefinitionArgs{
//				ProjectId: example.Id,
//				Name:      "existing",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("id", exampleGetBuildDefinition.Id)
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func LookupBuildDefinition(ctx *pulumi.Context, args *LookupBuildDefinitionArgs, opts ...pulumi.InvokeOption) (*LookupBuildDefinitionResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupBuildDefinitionResult
	err := ctx.Invoke("azuredevops:index/getBuildDefinition:getBuildDefinition", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getBuildDefinition.
type LookupBuildDefinitionArgs struct {
	// The name of this Build Definition.
	Name string `pulumi:"name"`
	// The path of the build definition. Default to `\`.
	Path *string `pulumi:"path"`
	// The ID of the project.
	ProjectId string `pulumi:"projectId"`
}

// A collection of values returned by getBuildDefinition.
type LookupBuildDefinitionResult struct {
	// The agent pool that should execute the build.
	AgentPoolName string `pulumi:"agentPoolName"`
	// A `ciTrigger` block as defined below.
	CiTriggers []GetBuildDefinitionCiTrigger `pulumi:"ciTriggers"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The name of the variable.
	Name      string  `pulumi:"name"`
	Path      *string `pulumi:"path"`
	ProjectId string  `pulumi:"projectId"`
	// A `pullRequestTrigger` block as defined below.
	PullRequestTriggers []GetBuildDefinitionPullRequestTrigger `pulumi:"pullRequestTriggers"`
	// The queue status of the build definition.
	QueueStatus string `pulumi:"queueStatus"`
	// A `repository` block as defined below.
	Repositories []GetBuildDefinitionRepository `pulumi:"repositories"`
	// The revision of the build definition.
	Revision int `pulumi:"revision"`
	// A `schedules` block as defined below.
	Schedules []GetBuildDefinitionSchedule `pulumi:"schedules"`
	// A list of variable group IDs.
	VariableGroups []int `pulumi:"variableGroups"`
	// A `variable` block as defined below.
	Variables []GetBuildDefinitionVariable `pulumi:"variables"`
}

func LookupBuildDefinitionOutput(ctx *pulumi.Context, args LookupBuildDefinitionOutputArgs, opts ...pulumi.InvokeOption) LookupBuildDefinitionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBuildDefinitionResult, error) {
			args := v.(LookupBuildDefinitionArgs)
			r, err := LookupBuildDefinition(ctx, &args, opts...)
			var s LookupBuildDefinitionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBuildDefinitionResultOutput)
}

// A collection of arguments for invoking getBuildDefinition.
type LookupBuildDefinitionOutputArgs struct {
	// The name of this Build Definition.
	Name pulumi.StringInput `pulumi:"name"`
	// The path of the build definition. Default to `\`.
	Path pulumi.StringPtrInput `pulumi:"path"`
	// The ID of the project.
	ProjectId pulumi.StringInput `pulumi:"projectId"`
}

func (LookupBuildDefinitionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBuildDefinitionArgs)(nil)).Elem()
}

// A collection of values returned by getBuildDefinition.
type LookupBuildDefinitionResultOutput struct{ *pulumi.OutputState }

func (LookupBuildDefinitionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBuildDefinitionResult)(nil)).Elem()
}

func (o LookupBuildDefinitionResultOutput) ToLookupBuildDefinitionResultOutput() LookupBuildDefinitionResultOutput {
	return o
}

func (o LookupBuildDefinitionResultOutput) ToLookupBuildDefinitionResultOutputWithContext(ctx context.Context) LookupBuildDefinitionResultOutput {
	return o
}

// The agent pool that should execute the build.
func (o LookupBuildDefinitionResultOutput) AgentPoolName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBuildDefinitionResult) string { return v.AgentPoolName }).(pulumi.StringOutput)
}

// A `ciTrigger` block as defined below.
func (o LookupBuildDefinitionResultOutput) CiTriggers() GetBuildDefinitionCiTriggerArrayOutput {
	return o.ApplyT(func(v LookupBuildDefinitionResult) []GetBuildDefinitionCiTrigger { return v.CiTriggers }).(GetBuildDefinitionCiTriggerArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupBuildDefinitionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBuildDefinitionResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the variable.
func (o LookupBuildDefinitionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBuildDefinitionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupBuildDefinitionResultOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBuildDefinitionResult) *string { return v.Path }).(pulumi.StringPtrOutput)
}

func (o LookupBuildDefinitionResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBuildDefinitionResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

// A `pullRequestTrigger` block as defined below.
func (o LookupBuildDefinitionResultOutput) PullRequestTriggers() GetBuildDefinitionPullRequestTriggerArrayOutput {
	return o.ApplyT(func(v LookupBuildDefinitionResult) []GetBuildDefinitionPullRequestTrigger {
		return v.PullRequestTriggers
	}).(GetBuildDefinitionPullRequestTriggerArrayOutput)
}

// The queue status of the build definition.
func (o LookupBuildDefinitionResultOutput) QueueStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBuildDefinitionResult) string { return v.QueueStatus }).(pulumi.StringOutput)
}

// A `repository` block as defined below.
func (o LookupBuildDefinitionResultOutput) Repositories() GetBuildDefinitionRepositoryArrayOutput {
	return o.ApplyT(func(v LookupBuildDefinitionResult) []GetBuildDefinitionRepository { return v.Repositories }).(GetBuildDefinitionRepositoryArrayOutput)
}

// The revision of the build definition.
func (o LookupBuildDefinitionResultOutput) Revision() pulumi.IntOutput {
	return o.ApplyT(func(v LookupBuildDefinitionResult) int { return v.Revision }).(pulumi.IntOutput)
}

// A `schedules` block as defined below.
func (o LookupBuildDefinitionResultOutput) Schedules() GetBuildDefinitionScheduleArrayOutput {
	return o.ApplyT(func(v LookupBuildDefinitionResult) []GetBuildDefinitionSchedule { return v.Schedules }).(GetBuildDefinitionScheduleArrayOutput)
}

// A list of variable group IDs.
func (o LookupBuildDefinitionResultOutput) VariableGroups() pulumi.IntArrayOutput {
	return o.ApplyT(func(v LookupBuildDefinitionResult) []int { return v.VariableGroups }).(pulumi.IntArrayOutput)
}

// A `variable` block as defined below.
func (o LookupBuildDefinitionResultOutput) Variables() GetBuildDefinitionVariableArrayOutput {
	return o.ApplyT(func(v LookupBuildDefinitionResult) []GetBuildDefinitionVariable { return v.Variables }).(GetBuildDefinitionVariableArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBuildDefinitionResultOutput{})
}
