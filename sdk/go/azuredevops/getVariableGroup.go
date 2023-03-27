// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about existing Variable Groups within Azure DevOps.
//
// > **Note:** Secret values are masked by service and cannot be obtained through API. [Set secret variables](https://docs.microsoft.com/en-us/azure/devops/pipelines/process/variables?view=azure-devops&tabs=yaml%2Cbatch#secret-variables)
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
//			exampleProject, err := azuredevops.LookupProject(ctx, &azuredevops.LookupProjectArgs{
//				Name: pulumi.StringRef("Example Project"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			exampleVariableGroup, err := azuredevops.LookupVariableGroup(ctx, &azuredevops.LookupVariableGroupArgs{
//				ProjectId: exampleProject.Id,
//				Name:      "Example Variable Group",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("id", exampleVariableGroup.Id)
//			return nil
//		})
//	}
//
// ```
// ## Relevant Links
//
// - [Azure DevOps Service REST API 6.0 - Variable Groups](https://docs.microsoft.com/en-us/rest/api/azure/devops/distributedtask/variablegroups?view=azure-devops-rest-6.0)
func LookupVariableGroup(ctx *pulumi.Context, args *LookupVariableGroupArgs, opts ...pulumi.InvokeOption) (*LookupVariableGroupResult, error) {
	var rv LookupVariableGroupResult
	err := ctx.Invoke("azuredevops:index/getVariableGroup:getVariableGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVariableGroup.
type LookupVariableGroupArgs struct {
	// The name of the Variable Group to retrieve.
	Name string `pulumi:"name"`
	// The project ID.
	ProjectId string `pulumi:"projectId"`
}

// A collection of values returned by getVariableGroup.
type LookupVariableGroupResult struct {
	// Boolean that indicate if this Variable Group is shared by all pipelines of this project.
	AllowAccess bool `pulumi:"allowAccess"`
	// The description of the Variable Group.
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of `keyVault` blocks as documented below.
	KeyVaults []GetVariableGroupKeyVault `pulumi:"keyVaults"`
	// The name of the Azure key vault to link secrets from as variables.
	Name      string `pulumi:"name"`
	ProjectId string `pulumi:"projectId"`
	// One or more `variable` blocks as documented below.
	Variables []GetVariableGroupVariable `pulumi:"variables"`
}

func LookupVariableGroupOutput(ctx *pulumi.Context, args LookupVariableGroupOutputArgs, opts ...pulumi.InvokeOption) LookupVariableGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVariableGroupResult, error) {
			args := v.(LookupVariableGroupArgs)
			r, err := LookupVariableGroup(ctx, &args, opts...)
			var s LookupVariableGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVariableGroupResultOutput)
}

// A collection of arguments for invoking getVariableGroup.
type LookupVariableGroupOutputArgs struct {
	// The name of the Variable Group to retrieve.
	Name pulumi.StringInput `pulumi:"name"`
	// The project ID.
	ProjectId pulumi.StringInput `pulumi:"projectId"`
}

func (LookupVariableGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVariableGroupArgs)(nil)).Elem()
}

// A collection of values returned by getVariableGroup.
type LookupVariableGroupResultOutput struct{ *pulumi.OutputState }

func (LookupVariableGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVariableGroupResult)(nil)).Elem()
}

func (o LookupVariableGroupResultOutput) ToLookupVariableGroupResultOutput() LookupVariableGroupResultOutput {
	return o
}

func (o LookupVariableGroupResultOutput) ToLookupVariableGroupResultOutputWithContext(ctx context.Context) LookupVariableGroupResultOutput {
	return o
}

// Boolean that indicate if this Variable Group is shared by all pipelines of this project.
func (o LookupVariableGroupResultOutput) AllowAccess() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupVariableGroupResult) bool { return v.AllowAccess }).(pulumi.BoolOutput)
}

// The description of the Variable Group.
func (o LookupVariableGroupResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVariableGroupResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupVariableGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVariableGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of `keyVault` blocks as documented below.
func (o LookupVariableGroupResultOutput) KeyVaults() GetVariableGroupKeyVaultArrayOutput {
	return o.ApplyT(func(v LookupVariableGroupResult) []GetVariableGroupKeyVault { return v.KeyVaults }).(GetVariableGroupKeyVaultArrayOutput)
}

// The name of the Azure key vault to link secrets from as variables.
func (o LookupVariableGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVariableGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupVariableGroupResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVariableGroupResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

// One or more `variable` blocks as documented below.
func (o LookupVariableGroupResultOutput) Variables() GetVariableGroupVariableArrayOutput {
	return o.ApplyT(func(v LookupVariableGroupResult) []GetVariableGroupVariable { return v.Variables }).(GetVariableGroupVariableArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVariableGroupResultOutput{})
}
