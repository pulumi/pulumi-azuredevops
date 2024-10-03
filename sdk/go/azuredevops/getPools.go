// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azuredevops/sdk/v3/go/azuredevops/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about existing Agent Pools within Azure DevOps.
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
// func main() {
// pulumi.Run(func(ctx *pulumi.Context) error {
// example, err := azuredevops.GetPools(ctx, map[string]interface{}{
// }, nil);
// if err != nil {
// return err
// }
// ctx.Export("agentPoolName", pulumi.StringArray(%!v(PANIC=Format method: fatal: A failure has occurred: unlowered splat expression @ example.pp:3,11-37)))
// ctx.Export("autoProvision", pulumi.BoolArray(%!v(PANIC=Format method: fatal: A failure has occurred: unlowered splat expression @ example.pp:7,11-46)))
// ctx.Export("autoUpdate", pulumi.BoolArray(%!v(PANIC=Format method: fatal: A failure has occurred: unlowered splat expression @ example.pp:11,11-43)))
// ctx.Export("poolType", pulumi.StringArray(%!v(PANIC=Format method: fatal: A failure has occurred: unlowered splat expression @ example.pp:15,11-41)))
// return nil
// })
// }
// ```
//
// ## Relevant Links
//
// - [Azure DevOps Service REST API 7.0 - Agent Pools - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/distributedtask/pools/get?view=azure-devops-rest-7.0)
func GetPools(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetPoolsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetPoolsResult
	err := ctx.Invoke("azuredevops:index/getPools:getPools", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getPools.
type GetPoolsResult struct {
	// A list of existing agent pools in your Azure DevOps Organization with the following details about every agent pool:
	AgentPools []GetPoolsAgentPool `pulumi:"agentPools"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}

func GetPoolsOutput(ctx *pulumi.Context, opts ...pulumi.InvokeOption) GetPoolsResultOutput {
	return pulumi.ToOutput(0).ApplyT(func(int) (GetPoolsResultOutput, error) {
		opts = internal.PkgInvokeDefaultOpts(opts)
		var rv GetPoolsResult
		secret, err := ctx.InvokePackageRaw("azuredevops:index/getPools:getPools", nil, &rv, "", opts...)
		if err != nil {
			return GetPoolsResultOutput{}, err
		}

		output := pulumi.ToOutput(rv).(GetPoolsResultOutput)
		if secret {
			return pulumi.ToSecret(output).(GetPoolsResultOutput), nil
		}
		return output, nil
	}).(GetPoolsResultOutput)
}

// A collection of values returned by getPools.
type GetPoolsResultOutput struct{ *pulumi.OutputState }

func (GetPoolsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPoolsResult)(nil)).Elem()
}

func (o GetPoolsResultOutput) ToGetPoolsResultOutput() GetPoolsResultOutput {
	return o
}

func (o GetPoolsResultOutput) ToGetPoolsResultOutputWithContext(ctx context.Context) GetPoolsResultOutput {
	return o
}

// A list of existing agent pools in your Azure DevOps Organization with the following details about every agent pool:
func (o GetPoolsResultOutput) AgentPools() GetPoolsAgentPoolArrayOutput {
	return o.ApplyT(func(v GetPoolsResult) []GetPoolsAgentPool { return v.AgentPools }).(GetPoolsAgentPoolArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetPoolsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetPoolsResult) string { return v.Id }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetPoolsResultOutput{})
}
