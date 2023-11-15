// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package core

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azuredevops/sdk/v2/go/azuredevops/internal"
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
//	"github.com/pulumi/pulumi-azuredevops/sdk/v2/go/azuredevops"
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
// ## Relevant Links
//
// - [Azure DevOps Service REST API 7.0 - Projects - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/core/projects/get?view=azure-devops-rest-7.0)
//
// Deprecated: azuredevops.core.getProject has been deprecated in favor of azuredevops.getProject
func LookupProject(ctx *pulumi.Context, args *LookupProjectArgs, opts ...pulumi.InvokeOption) (*LookupProjectResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupProjectResult
	err := ctx.Invoke("azuredevops:Core/getProject:getProject", args, &rv, opts...)
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
	ProjectId *string `pulumi:"projectId"`
}

// A collection of values returned by getProject.
type LookupProjectResult struct {
	Description string                 `pulumi:"description"`
	Features    map[string]interface{} `pulumi:"features"`
	// The provider-assigned unique ID for this managed resource.
	Id                string  `pulumi:"id"`
	Name              *string `pulumi:"name"`
	ProcessTemplateId string  `pulumi:"processTemplateId"`
	ProjectId         *string `pulumi:"projectId"`
	VersionControl    string  `pulumi:"versionControl"`
	Visibility        string  `pulumi:"visibility"`
	WorkItemTemplate  string  `pulumi:"workItemTemplate"`
}

func LookupProjectOutput(ctx *pulumi.Context, args LookupProjectOutputArgs, opts ...pulumi.InvokeOption) LookupProjectResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupProjectResult, error) {
			args := v.(LookupProjectArgs)
			r, err := LookupProject(ctx, &args, opts...)
			var s LookupProjectResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupProjectResultOutput)
}

// A collection of arguments for invoking getProject.
type LookupProjectOutputArgs struct {
	// Name of the Project.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// ID of the Project.
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

func (o LookupProjectResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupProjectResultOutput) Features() pulumi.MapOutput {
	return o.ApplyT(func(v LookupProjectResult) map[string]interface{} { return v.Features }).(pulumi.MapOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupProjectResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupProjectResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProjectResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupProjectResultOutput) ProcessTemplateId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.ProcessTemplateId }).(pulumi.StringOutput)
}

func (o LookupProjectResultOutput) ProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProjectResult) *string { return v.ProjectId }).(pulumi.StringPtrOutput)
}

func (o LookupProjectResultOutput) VersionControl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.VersionControl }).(pulumi.StringOutput)
}

func (o LookupProjectResultOutput) Visibility() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.Visibility }).(pulumi.StringOutput)
}

func (o LookupProjectResultOutput) WorkItemTemplate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.WorkItemTemplate }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProjectResultOutput{})
}
