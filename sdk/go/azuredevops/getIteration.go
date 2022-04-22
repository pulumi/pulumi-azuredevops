// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing Iteration (Sprint) within Azure DevOps.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azuredevops/sdk/v2/go/azuredevops"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := azuredevops.NewProject(ctx, "example", &azuredevops.ProjectArgs{
// 			WorkItemTemplate: pulumi.String("Agile"),
// 			VersionControl:   pulumi.String("Git"),
// 			Visibility:       pulumi.String("private"),
// 			Description:      pulumi.String("Managed by Terraform"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_ = azuredevops.GetIterationOutput(ctx, GetIterationOutputArgs{
// 			ProjectId:     example.ID(),
// 			Path:          pulumi.String("/"),
// 			FetchChildren: pulumi.Bool(true),
// 		}, nil)
// 		_ = azuredevops.GetIterationOutput(ctx, GetIterationOutputArgs{
// 			ProjectId:     example.ID(),
// 			Path:          pulumi.String("/Iteration 1"),
// 			FetchChildren: pulumi.Bool(true),
// 		}, nil)
// 		return nil
// 	})
// }
// ```
// ## Relevant Links
//
// - [Azure DevOps Service REST API 6.0 - Classification Nodes - Get Classification Nodes](https://docs.microsoft.com/en-us/rest/api/azure/devops/wit/classification-nodes/get-classification-nodes?view=azure-devops-rest-6.0)
//
// ## PAT Permissions Required
//
// - **Project & Team**: vso.work - Grants the ability to read work items, queries, boards, area and iterations paths, and other work item tracking related metadata. Also grants the ability to execute queries, search work items and to receive notifications about work item events via service hooks.
func GetIteration(ctx *pulumi.Context, args *GetIterationArgs, opts ...pulumi.InvokeOption) (*GetIterationResult, error) {
	var rv GetIterationResult
	err := ctx.Invoke("azuredevops:index/getIteration:getIteration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIteration.
type GetIterationArgs struct {
	// Read children nodes, _Depth_: 1, _Default_: `true`
	FetchChildren *bool `pulumi:"fetchChildren"`
	// The path to the Iteration, _Format_: URL relative; if omitted, or value `"/"` is used, the root Iteration will be returned
	Path *string `pulumi:"path"`
	// The project ID.
	ProjectId string `pulumi:"projectId"`
}

// A collection of values returned by getIteration.
type GetIterationResult struct {
	// A list of `children` blocks as defined below, empty if `hasChildren == false`
	Childrens     []GetIterationChildren `pulumi:"childrens"`
	FetchChildren *bool                  `pulumi:"fetchChildren"`
	// Indicator if the child Iteration node has child nodes
	HasChildren bool `pulumi:"hasChildren"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The name of the child Iteration node
	Name string `pulumi:"name"`
	// The complete path (in relative URL format) of the child Iteration
	Path string `pulumi:"path"`
	// The project ID of the child Iteration node
	ProjectId string `pulumi:"projectId"`
}

func GetIterationOutput(ctx *pulumi.Context, args GetIterationOutputArgs, opts ...pulumi.InvokeOption) GetIterationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetIterationResult, error) {
			args := v.(GetIterationArgs)
			r, err := GetIteration(ctx, &args, opts...)
			var s GetIterationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetIterationResultOutput)
}

// A collection of arguments for invoking getIteration.
type GetIterationOutputArgs struct {
	// Read children nodes, _Depth_: 1, _Default_: `true`
	FetchChildren pulumi.BoolPtrInput `pulumi:"fetchChildren"`
	// The path to the Iteration, _Format_: URL relative; if omitted, or value `"/"` is used, the root Iteration will be returned
	Path pulumi.StringPtrInput `pulumi:"path"`
	// The project ID.
	ProjectId pulumi.StringInput `pulumi:"projectId"`
}

func (GetIterationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIterationArgs)(nil)).Elem()
}

// A collection of values returned by getIteration.
type GetIterationResultOutput struct{ *pulumi.OutputState }

func (GetIterationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIterationResult)(nil)).Elem()
}

func (o GetIterationResultOutput) ToGetIterationResultOutput() GetIterationResultOutput {
	return o
}

func (o GetIterationResultOutput) ToGetIterationResultOutputWithContext(ctx context.Context) GetIterationResultOutput {
	return o
}

// A list of `children` blocks as defined below, empty if `hasChildren == false`
func (o GetIterationResultOutput) Childrens() GetIterationChildrenArrayOutput {
	return o.ApplyT(func(v GetIterationResult) []GetIterationChildren { return v.Childrens }).(GetIterationChildrenArrayOutput)
}

func (o GetIterationResultOutput) FetchChildren() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetIterationResult) *bool { return v.FetchChildren }).(pulumi.BoolPtrOutput)
}

// Indicator if the child Iteration node has child nodes
func (o GetIterationResultOutput) HasChildren() pulumi.BoolOutput {
	return o.ApplyT(func(v GetIterationResult) bool { return v.HasChildren }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetIterationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetIterationResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the child Iteration node
func (o GetIterationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetIterationResult) string { return v.Name }).(pulumi.StringOutput)
}

// The complete path (in relative URL format) of the child Iteration
func (o GetIterationResultOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v GetIterationResult) string { return v.Path }).(pulumi.StringOutput)
}

// The project ID of the child Iteration node
func (o GetIterationResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v GetIterationResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetIterationResultOutput{})
}
