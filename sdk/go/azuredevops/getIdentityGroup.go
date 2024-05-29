// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azuredevops/sdk/v3/go/azuredevops/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing Group within Azure DevOps On-Premise(Azure DevOps Server).
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
//			// load existing group with specific name
//			_, err := azuredevops.GetIdentityGroup(ctx, &azuredevops.GetIdentityGroupArgs{
//				ProjectId: example.Id,
//				Name:      "[Project-Name]\\Group-Name",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Relevant Links
//
// - [Azure DevOps Service REST API 7.0 - Identities](https://docs.microsoft.com/en-us/rest/api/azure/devops/ims/?view=azure-devops-rest-7.2)
func GetIdentityGroup(ctx *pulumi.Context, args *GetIdentityGroupArgs, opts ...pulumi.InvokeOption) (*GetIdentityGroupResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetIdentityGroupResult
	err := ctx.Invoke("azuredevops:index/getIdentityGroup:getIdentityGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIdentityGroup.
type GetIdentityGroupArgs struct {
	// The name of the group.
	Name string `pulumi:"name"`
	// The Project ID.
	ProjectId string `pulumi:"projectId"`
}

// A collection of values returned by getIdentityGroup.
type GetIdentityGroupResult struct {
	Descriptor string `pulumi:"descriptor"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// This is the non-unique display name of the identity subject. To change this field, you must alter its value in the source provider.
	Name      string `pulumi:"name"`
	ProjectId string `pulumi:"projectId"`
}

func GetIdentityGroupOutput(ctx *pulumi.Context, args GetIdentityGroupOutputArgs, opts ...pulumi.InvokeOption) GetIdentityGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetIdentityGroupResult, error) {
			args := v.(GetIdentityGroupArgs)
			r, err := GetIdentityGroup(ctx, &args, opts...)
			var s GetIdentityGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetIdentityGroupResultOutput)
}

// A collection of arguments for invoking getIdentityGroup.
type GetIdentityGroupOutputArgs struct {
	// The name of the group.
	Name pulumi.StringInput `pulumi:"name"`
	// The Project ID.
	ProjectId pulumi.StringInput `pulumi:"projectId"`
}

func (GetIdentityGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIdentityGroupArgs)(nil)).Elem()
}

// A collection of values returned by getIdentityGroup.
type GetIdentityGroupResultOutput struct{ *pulumi.OutputState }

func (GetIdentityGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIdentityGroupResult)(nil)).Elem()
}

func (o GetIdentityGroupResultOutput) ToGetIdentityGroupResultOutput() GetIdentityGroupResultOutput {
	return o
}

func (o GetIdentityGroupResultOutput) ToGetIdentityGroupResultOutputWithContext(ctx context.Context) GetIdentityGroupResultOutput {
	return o
}

func (o GetIdentityGroupResultOutput) Descriptor() pulumi.StringOutput {
	return o.ApplyT(func(v GetIdentityGroupResult) string { return v.Descriptor }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetIdentityGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetIdentityGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

// This is the non-unique display name of the identity subject. To change this field, you must alter its value in the source provider.
func (o GetIdentityGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetIdentityGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetIdentityGroupResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v GetIdentityGroupResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetIdentityGroupResultOutput{})
}