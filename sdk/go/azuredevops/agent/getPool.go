// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package agent

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to access information about an existing Agent Pool within Azure DevOps.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azuredevops/sdk/go/azuredevops"
// 	"github.com/pulumi/pulumi-azuredevops/sdk/go/azuredevops/"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		pool, err := azuredevops.LookupPool(ctx, &azuredevops.LookupPoolArgs{
// 			Name: "Sample Agent Pool",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("name", pool.Name)
// 		ctx.Export("poolType", pool.PoolType)
// 		ctx.Export("autoProvision", pool.AutoProvision)
// 		return nil
// 	})
// }
// ```
// ## Relevant Links
//
// - [Azure DevOps Service REST API 5.1 - Agent Pools - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/distributedtask/pools/get?view=azure-devops-rest-5.1)
//
// Deprecated: azuredevops.agent.getPool has been deprecated in favor of azuredevops.getPool
func LookupPool(ctx *pulumi.Context, args *LookupPoolArgs, opts ...pulumi.InvokeOption) (*LookupPoolResult, error) {
	var rv LookupPoolResult
	err := ctx.Invoke("azuredevops:Agent/getPool:getPool", args, &rv, opts...)
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
	AutoProvision bool `pulumi:"autoProvision"`
	// The provider-assigned unique ID for this managed resource.
	Id       string `pulumi:"id"`
	Name     string `pulumi:"name"`
	PoolType string `pulumi:"poolType"`
}
