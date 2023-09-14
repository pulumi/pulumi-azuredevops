// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azuredevops/sdk/v2/go/azuredevops/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Use this data source to access information about existing Projects within Azure DevOps.
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
//			example, err := azuredevops.GetProjects(ctx, &azuredevops.GetProjectsArgs{
//				Name:  pulumi.StringRef("Example Project"),
//				State: pulumi.StringRef("wellFormed"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			var splat0 []*string
//			for _, val0 := range example.Projects {
//				splat0 = append(splat0, val0.ProjectId)
//			}
//			ctx.Export("projectId", splat0)
//			var splat1 []*string
//			for _, val0 := range example.Projects {
//				splat1 = append(splat1, val0.Name)
//			}
//			ctx.Export("name", splat1)
//			var splat2 []*string
//			for _, val0 := range example.Projects {
//				splat2 = append(splat2, val0.ProjectUrl)
//			}
//			ctx.Export("projectUrl", splat2)
//			var splat3 []*string
//			for _, val0 := range example.Projects {
//				splat3 = append(splat3, val0.State)
//			}
//			ctx.Export("state", splat3)
//			return nil
//		})
//	}
//
// ```
// ## Relevant Links
//
// - [Azure DevOps Service REST API 7.0 - Projects - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/core/projects/get?view=azure-devops-rest-7.0)
func GetProjects(ctx *pulumi.Context, args *GetProjectsArgs, opts ...pulumi.InvokeOption) (*GetProjectsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetProjectsResult
	err := ctx.Invoke("azuredevops:index/getProjects:getProjects", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProjects.
type GetProjectsArgs struct {
	// Name of the Project, if not specified all projects will be returned.
	Name *string `pulumi:"name"`
	// State of the Project, if not specified all projects will be returned. Valid values are `all`, `deleting`, `new`, `wellFormed`, `createPending`, `unchanged`,`deleted`.
	//
	// DataSource without specifying any arguments will return all projects.
	State *string `pulumi:"state"`
}

// A collection of values returned by getProjects.
type GetProjectsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The name of the Project.
	Name *string `pulumi:"name"`
	// A list of existing projects in your Azure DevOps Organization with details about every project which includes:
	Projects []GetProjectsProject `pulumi:"projects"`
	// Project state.
	State *string `pulumi:"state"`
}

func GetProjectsOutput(ctx *pulumi.Context, args GetProjectsOutputArgs, opts ...pulumi.InvokeOption) GetProjectsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetProjectsResult, error) {
			args := v.(GetProjectsArgs)
			r, err := GetProjects(ctx, &args, opts...)
			var s GetProjectsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetProjectsResultOutput)
}

// A collection of arguments for invoking getProjects.
type GetProjectsOutputArgs struct {
	// Name of the Project, if not specified all projects will be returned.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// State of the Project, if not specified all projects will be returned. Valid values are `all`, `deleting`, `new`, `wellFormed`, `createPending`, `unchanged`,`deleted`.
	//
	// DataSource without specifying any arguments will return all projects.
	State pulumi.StringPtrInput `pulumi:"state"`
}

func (GetProjectsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProjectsArgs)(nil)).Elem()
}

// A collection of values returned by getProjects.
type GetProjectsResultOutput struct{ *pulumi.OutputState }

func (GetProjectsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProjectsResult)(nil)).Elem()
}

func (o GetProjectsResultOutput) ToGetProjectsResultOutput() GetProjectsResultOutput {
	return o
}

func (o GetProjectsResultOutput) ToGetProjectsResultOutputWithContext(ctx context.Context) GetProjectsResultOutput {
	return o
}

func (o GetProjectsResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetProjectsResult] {
	return pulumix.Output[GetProjectsResult]{
		OutputState: o.OutputState,
	}
}

// The provider-assigned unique ID for this managed resource.
func (o GetProjectsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectsResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the Project.
func (o GetProjectsResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProjectsResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// A list of existing projects in your Azure DevOps Organization with details about every project which includes:
func (o GetProjectsResultOutput) Projects() GetProjectsProjectArrayOutput {
	return o.ApplyT(func(v GetProjectsResult) []GetProjectsProject { return v.Projects }).(GetProjectsProjectArrayOutput)
}

// Project state.
func (o GetProjectsResultOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProjectsResult) *string { return v.State }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetProjectsResultOutput{})
}
