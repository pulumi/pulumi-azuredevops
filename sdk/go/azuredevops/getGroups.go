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

// Use this data source to access information about existing Groups within Azure DevOps
//
// ## Relevant Links
//
// - [Azure DevOps Service REST API 7.0 - Groups - List](https://docs.microsoft.com/en-us/rest/api/azure/devops/graph/groups/list?view=azure-devops-rest-7.0)
func GetGroups(ctx *pulumi.Context, args *GetGroupsArgs, opts ...pulumi.InvokeOption) (*GetGroupsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetGroupsResult
	err := ctx.Invoke("azuredevops:index/getGroups:getGroups", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGroups.
type GetGroupsArgs struct {
	// The Project ID. If no project ID is specified all groups of an organization will be returned
	ProjectId *string `pulumi:"projectId"`
}

// A collection of values returned by getGroups.
type GetGroupsResult struct {
	// A set of existing groups in your Azure DevOps Organization or project with details about every single group which includes:
	Groups []GetGroupsGroup `pulumi:"groups"`
	// The provider-assigned unique ID for this managed resource.
	Id        string  `pulumi:"id"`
	ProjectId *string `pulumi:"projectId"`
}

func GetGroupsOutput(ctx *pulumi.Context, args GetGroupsOutputArgs, opts ...pulumi.InvokeOption) GetGroupsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetGroupsResult, error) {
			args := v.(GetGroupsArgs)
			r, err := GetGroups(ctx, &args, opts...)
			var s GetGroupsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetGroupsResultOutput)
}

// A collection of arguments for invoking getGroups.
type GetGroupsOutputArgs struct {
	// The Project ID. If no project ID is specified all groups of an organization will be returned
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
}

func (GetGroupsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGroupsArgs)(nil)).Elem()
}

// A collection of values returned by getGroups.
type GetGroupsResultOutput struct{ *pulumi.OutputState }

func (GetGroupsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGroupsResult)(nil)).Elem()
}

func (o GetGroupsResultOutput) ToGetGroupsResultOutput() GetGroupsResultOutput {
	return o
}

func (o GetGroupsResultOutput) ToGetGroupsResultOutputWithContext(ctx context.Context) GetGroupsResultOutput {
	return o
}

func (o GetGroupsResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetGroupsResult] {
	return pulumix.Output[GetGroupsResult]{
		OutputState: o.OutputState,
	}
}

// A set of existing groups in your Azure DevOps Organization or project with details about every single group which includes:
func (o GetGroupsResultOutput) Groups() GetGroupsGroupArrayOutput {
	return o.ApplyT(func(v GetGroupsResult) []GetGroupsGroup { return v.Groups }).(GetGroupsGroupArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetGroupsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetGroupsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetGroupsResultOutput) ProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGroupsResult) *string { return v.ProjectId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetGroupsResultOutput{})
}
