// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package agent

import (
	"github.com/pulumi/pulumi-azuredevops/sdk/v2/go/azuredevops/internal"
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
//	"github.com/pulumi/pulumi-azuredevops/sdk/v2/go/azuredevops"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			example, err := azuredevops.GetPools(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			var splat0 []*string
//			for _, val0 := range example.AgentPools {
//				splat0 = append(splat0, val0.Name)
//			}
//			ctx.Export("agentPoolName", splat0)
//			var splat1 []*bool
//			for _, val0 := range example.AgentPools {
//				splat1 = append(splat1, val0.AutoProvision)
//			}
//			ctx.Export("autoProvision", splat1)
//			var splat2 []*bool
//			for _, val0 := range example.AgentPools {
//				splat2 = append(splat2, val0.AutoUpdate)
//			}
//			ctx.Export("autoUpdate", splat2)
//			var splat3 []*string
//			for _, val0 := range example.AgentPools {
//				splat3 = append(splat3, val0.PoolType)
//			}
//			ctx.Export("poolType", splat3)
//			return nil
//		})
//	}
//
// ```
// ## Relevant Links
//
// - [Azure DevOps Service REST API 7.0 - Agent Pools - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/distributedtask/pools/get?view=azure-devops-rest-7.0)
//
// Deprecated: azuredevops.agent.getPools has been deprecated in favor of azuredevops.getPools
func GetPools(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetPoolsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetPoolsResult
	err := ctx.Invoke("azuredevops:Agent/getPools:getPools", nil, &rv, opts...)
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
