// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azuredevops/sdk/v3/go/azuredevops/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing Agent Pool within Azure DevOps.
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
//			example, err := azuredevops.LookupPool(ctx, &azuredevops.LookupPoolArgs{
//				Name: "Example Agent Pool",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("name", example.Name)
//			ctx.Export("poolType", example.PoolType)
//			ctx.Export("autoProvision", example.AutoProvision)
//			ctx.Export("autoUpdate", example.AutoUpdate)
//			return nil
//		})
//	}
//
// ```
//
// ## Relevant Links
//
// - [Azure DevOps Service REST API 7.1 - Agent Pools - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/distributedtask/pools/get?view=azure-devops-rest-7.1)
func LookupPool(ctx *pulumi.Context, args *LookupPoolArgs, opts ...pulumi.InvokeOption) (*LookupPoolResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupPoolResult
	err := ctx.Invoke("azuredevops:index/getPool:getPool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPool.
type LookupPoolArgs struct {
	// Name of the Agent Pool.
	Name string `pulumi:"name"`
}

// A collection of values returned by getPool.
type LookupPoolResult struct {
	// Specifies whether a queue should be automatically provisioned for each project collection.
	AutoProvision bool `pulumi:"autoProvision"`
	// Specifies whether or not agents within the pool should be automatically updated.
	AutoUpdate bool `pulumi:"autoUpdate"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The name of the agent pool
	Name string `pulumi:"name"`
	// Specifies whether the agent pool type is Automation or Deployment.
	PoolType string `pulumi:"poolType"`
}

func LookupPoolOutput(ctx *pulumi.Context, args LookupPoolOutputArgs, opts ...pulumi.InvokeOption) LookupPoolResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupPoolResultOutput, error) {
			args := v.(LookupPoolArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azuredevops:index/getPool:getPool", args, LookupPoolResultOutput{}, options).(LookupPoolResultOutput), nil
		}).(LookupPoolResultOutput)
}

// A collection of arguments for invoking getPool.
type LookupPoolOutputArgs struct {
	// Name of the Agent Pool.
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupPoolOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPoolArgs)(nil)).Elem()
}

// A collection of values returned by getPool.
type LookupPoolResultOutput struct{ *pulumi.OutputState }

func (LookupPoolResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPoolResult)(nil)).Elem()
}

func (o LookupPoolResultOutput) ToLookupPoolResultOutput() LookupPoolResultOutput {
	return o
}

func (o LookupPoolResultOutput) ToLookupPoolResultOutputWithContext(ctx context.Context) LookupPoolResultOutput {
	return o
}

// Specifies whether a queue should be automatically provisioned for each project collection.
func (o LookupPoolResultOutput) AutoProvision() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupPoolResult) bool { return v.AutoProvision }).(pulumi.BoolOutput)
}

// Specifies whether or not agents within the pool should be automatically updated.
func (o LookupPoolResultOutput) AutoUpdate() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupPoolResult) bool { return v.AutoUpdate }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupPoolResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPoolResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the agent pool
func (o LookupPoolResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPoolResult) string { return v.Name }).(pulumi.StringOutput)
}

// Specifies whether the agent pool type is Automation or Deployment.
func (o LookupPoolResultOutput) PoolType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPoolResult) string { return v.PoolType }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPoolResultOutput{})
}
