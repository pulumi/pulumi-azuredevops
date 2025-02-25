// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azuredevops/sdk/v3/go/azuredevops/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about the Azure DevOps organization configured for the provider.
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
//			example, err := azuredevops.GetClientConfig(ctx, map[string]interface{}{}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("orgUrl", example.OrganizationUrl)
//			return nil
//		})
//	}
//
// ```
func GetClientConfig(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetClientConfigResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetClientConfigResult
	err := ctx.Invoke("azuredevops:index/getClientConfig:getClientConfig", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getClientConfig.
type GetClientConfigResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The organization configured for the provider.
	OrganizationUrl string `pulumi:"organizationUrl"`
}

func GetClientConfigOutput(ctx *pulumi.Context, opts ...pulumi.InvokeOption) GetClientConfigResultOutput {
	return pulumi.ToOutput(0).ApplyT(func(int) (GetClientConfigResultOutput, error) {
		options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
		return ctx.InvokeOutput("azuredevops:index/getClientConfig:getClientConfig", nil, GetClientConfigResultOutput{}, options).(GetClientConfigResultOutput), nil
	}).(GetClientConfigResultOutput)
}

// A collection of values returned by getClientConfig.
type GetClientConfigResultOutput struct{ *pulumi.OutputState }

func (GetClientConfigResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetClientConfigResult)(nil)).Elem()
}

func (o GetClientConfigResultOutput) ToGetClientConfigResultOutput() GetClientConfigResultOutput {
	return o
}

func (o GetClientConfigResultOutput) ToGetClientConfigResultOutputWithContext(ctx context.Context) GetClientConfigResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetClientConfigResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetClientConfigResult) string { return v.Id }).(pulumi.StringOutput)
}

// The organization configured for the provider.
func (o GetClientConfigResultOutput) OrganizationUrl() pulumi.StringOutput {
	return o.ApplyT(func(v GetClientConfigResult) string { return v.OrganizationUrl }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetClientConfigResultOutput{})
}
