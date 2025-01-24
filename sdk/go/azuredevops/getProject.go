// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azuredevops/sdk/v3/go/azuredevops/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing Project within Azure DevOps.
//
// ## Example Usage
//
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
//			ctx.Export("project", example)
//			return nil
//		})
//	}
//
// ```
//
// ## Relevant Links
//
// - [Azure DevOps Service REST API 7.0 - Projects - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/core/projects/get?view=azure-devops-rest-7.0)
//
// ## PAT Permissions Required
//
// - **Project & Team**: Read
// - **Work Items**: Read
func LookupProject(ctx *pulumi.Context, args *LookupProjectArgs, opts ...pulumi.InvokeOption) (*LookupProjectResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupProjectResult
	err := ctx.Invoke("azuredevops:index/getProject:getProject", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProject.
type LookupProjectArgs struct {
	// Name of the Project.
	Name *string `pulumi:"name"`
	// ID of the Project.
	//
	// > **NOTE:** One of either `projectId` or `name` must be specified.
	ProjectId *string `pulumi:"projectId"`
}

// A collection of values returned by getProject.
type LookupProjectResult struct {
	// The description of the project.
	Description string            `pulumi:"description"`
	Features    map[string]string `pulumi:"features"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The name of the project.
	Name *string `pulumi:"name"`
	// The process template ID for the project.
	ProcessTemplateId string  `pulumi:"processTemplateId"`
	ProjectId         *string `pulumi:"projectId"`
	// The version control of the project.
	VersionControl string `pulumi:"versionControl"`
	// The visibility of the project.
	Visibility string `pulumi:"visibility"`
	// The work item template for the project.
	WorkItemTemplate string `pulumi:"workItemTemplate"`
}

func LookupProjectOutput(ctx *pulumi.Context, args LookupProjectOutputArgs, opts ...pulumi.InvokeOption) LookupProjectResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupProjectResultOutput, error) {
			args := v.(LookupProjectArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azuredevops:index/getProject:getProject", args, LookupProjectResultOutput{}, options).(LookupProjectResultOutput), nil
		}).(LookupProjectResultOutput)
}

// A collection of arguments for invoking getProject.
type LookupProjectOutputArgs struct {
	// Name of the Project.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// ID of the Project.
	//
	// > **NOTE:** One of either `projectId` or `name` must be specified.
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
}

func (LookupProjectOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectArgs)(nil)).Elem()
}

// A collection of values returned by getProject.
type LookupProjectResultOutput struct{ *pulumi.OutputState }

func (LookupProjectResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectResult)(nil)).Elem()
}

func (o LookupProjectResultOutput) ToLookupProjectResultOutput() LookupProjectResultOutput {
	return o
}

func (o LookupProjectResultOutput) ToLookupProjectResultOutputWithContext(ctx context.Context) LookupProjectResultOutput {
	return o
}

// The description of the project.
func (o LookupProjectResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupProjectResultOutput) Features() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupProjectResult) map[string]string { return v.Features }).(pulumi.StringMapOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupProjectResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the project.
func (o LookupProjectResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProjectResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The process template ID for the project.
func (o LookupProjectResultOutput) ProcessTemplateId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.ProcessTemplateId }).(pulumi.StringOutput)
}

func (o LookupProjectResultOutput) ProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProjectResult) *string { return v.ProjectId }).(pulumi.StringPtrOutput)
}

// The version control of the project.
func (o LookupProjectResultOutput) VersionControl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.VersionControl }).(pulumi.StringOutput)
}

// The visibility of the project.
func (o LookupProjectResultOutput) Visibility() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.Visibility }).(pulumi.StringOutput)
}

// The work item template for the project.
func (o LookupProjectResultOutput) WorkItemTemplate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.WorkItemTemplate }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProjectResultOutput{})
}
