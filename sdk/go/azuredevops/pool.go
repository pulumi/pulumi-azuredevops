// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an agent pool within Azure DevOps.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azuredevops/sdk/go/azuredevops"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := azuredevops.NewPool(ctx, "pool", &azuredevops.PoolArgs{
// 			AutoProvision: pulumi.Bool(false),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## Relevant Links
//
// * [Azure DevOps Service REST API 5.1 - Agent Pools](https://docs.microsoft.com/en-us/rest/api/azure/devops/distributedtask/pools?view=azure-devops-rest-5.1)
type Pool struct {
	pulumi.CustomResourceState

	// Specifies whether or not a queue should be automatically provisioned for each project collection. Defaults to `false`.
	AutoProvision pulumi.BoolPtrOutput `pulumi:"autoProvision"`
	// The name of the agent pool.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies whether the agent pool type is Automation or Deployment.  Defaults to `automation`.
	PoolType pulumi.StringPtrOutput `pulumi:"poolType"`
}

// NewPool registers a new resource with the given unique name, arguments, and options.
func NewPool(ctx *pulumi.Context,
	name string, args *PoolArgs, opts ...pulumi.ResourceOption) (*Pool, error) {
	if args == nil {
		args = &PoolArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azuredevops:Agent/pool:Pool"),
		},
	})
	opts = append(opts, aliases)
	var resource Pool
	err := ctx.RegisterResource("azuredevops:index/pool:Pool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPool gets an existing Pool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PoolState, opts ...pulumi.ResourceOption) (*Pool, error) {
	var resource Pool
	err := ctx.ReadResource("azuredevops:index/pool:Pool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Pool resources.
type poolState struct {
	// Specifies whether or not a queue should be automatically provisioned for each project collection. Defaults to `false`.
	AutoProvision *bool `pulumi:"autoProvision"`
	// The name of the agent pool.
	Name *string `pulumi:"name"`
	// Specifies whether the agent pool type is Automation or Deployment.  Defaults to `automation`.
	PoolType *string `pulumi:"poolType"`
}

type PoolState struct {
	// Specifies whether or not a queue should be automatically provisioned for each project collection. Defaults to `false`.
	AutoProvision pulumi.BoolPtrInput
	// The name of the agent pool.
	Name pulumi.StringPtrInput
	// Specifies whether the agent pool type is Automation or Deployment.  Defaults to `automation`.
	PoolType pulumi.StringPtrInput
}

func (PoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*poolState)(nil)).Elem()
}

type poolArgs struct {
	// Specifies whether or not a queue should be automatically provisioned for each project collection. Defaults to `false`.
	AutoProvision *bool `pulumi:"autoProvision"`
	// The name of the agent pool.
	Name *string `pulumi:"name"`
	// Specifies whether the agent pool type is Automation or Deployment.  Defaults to `automation`.
	PoolType *string `pulumi:"poolType"`
}

// The set of arguments for constructing a Pool resource.
type PoolArgs struct {
	// Specifies whether or not a queue should be automatically provisioned for each project collection. Defaults to `false`.
	AutoProvision pulumi.BoolPtrInput
	// The name of the agent pool.
	Name pulumi.StringPtrInput
	// Specifies whether the agent pool type is Automation or Deployment.  Defaults to `automation`.
	PoolType pulumi.StringPtrInput
}

func (PoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*poolArgs)(nil)).Elem()
}
