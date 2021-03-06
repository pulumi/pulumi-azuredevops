// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an agent pool within Azure DevOps.
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
// - [Azure DevOps Service REST API 5.1 - Agent Pools](https://docs.microsoft.com/en-us/rest/api/azure/devops/distributedtask/pools?view=azure-devops-rest-5.1)
//
// ## Import
//
// Azure DevOps Agent Pools can be imported using the agent pool ID, e.g.
//
// ```sh
//  $ pulumi import azuredevops:index/pool:Pool pool 42
// ```
type Pool struct {
	pulumi.CustomResourceState

	// Specifies whether or not a queue should be automatically provisioned for each project collection. Defaults to `false`.
	AutoProvision pulumi.BoolPtrOutput `pulumi:"autoProvision"`
	// The name of the agent pool.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies whether the agent pool type is Automation or Deployment. Defaults to `automation`.
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
	// Specifies whether the agent pool type is Automation or Deployment. Defaults to `automation`.
	PoolType *string `pulumi:"poolType"`
}

type PoolState struct {
	// Specifies whether or not a queue should be automatically provisioned for each project collection. Defaults to `false`.
	AutoProvision pulumi.BoolPtrInput
	// The name of the agent pool.
	Name pulumi.StringPtrInput
	// Specifies whether the agent pool type is Automation or Deployment. Defaults to `automation`.
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
	// Specifies whether the agent pool type is Automation or Deployment. Defaults to `automation`.
	PoolType *string `pulumi:"poolType"`
}

// The set of arguments for constructing a Pool resource.
type PoolArgs struct {
	// Specifies whether or not a queue should be automatically provisioned for each project collection. Defaults to `false`.
	AutoProvision pulumi.BoolPtrInput
	// The name of the agent pool.
	Name pulumi.StringPtrInput
	// Specifies whether the agent pool type is Automation or Deployment. Defaults to `automation`.
	PoolType pulumi.StringPtrInput
}

func (PoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*poolArgs)(nil)).Elem()
}

type PoolInput interface {
	pulumi.Input

	ToPoolOutput() PoolOutput
	ToPoolOutputWithContext(ctx context.Context) PoolOutput
}

func (*Pool) ElementType() reflect.Type {
	return reflect.TypeOf((*Pool)(nil))
}

func (i *Pool) ToPoolOutput() PoolOutput {
	return i.ToPoolOutputWithContext(context.Background())
}

func (i *Pool) ToPoolOutputWithContext(ctx context.Context) PoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PoolOutput)
}

func (i *Pool) ToPoolPtrOutput() PoolPtrOutput {
	return i.ToPoolPtrOutputWithContext(context.Background())
}

func (i *Pool) ToPoolPtrOutputWithContext(ctx context.Context) PoolPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PoolPtrOutput)
}

type PoolPtrInput interface {
	pulumi.Input

	ToPoolPtrOutput() PoolPtrOutput
	ToPoolPtrOutputWithContext(ctx context.Context) PoolPtrOutput
}

type poolPtrType PoolArgs

func (*poolPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Pool)(nil))
}

func (i *poolPtrType) ToPoolPtrOutput() PoolPtrOutput {
	return i.ToPoolPtrOutputWithContext(context.Background())
}

func (i *poolPtrType) ToPoolPtrOutputWithContext(ctx context.Context) PoolPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PoolPtrOutput)
}

// PoolArrayInput is an input type that accepts PoolArray and PoolArrayOutput values.
// You can construct a concrete instance of `PoolArrayInput` via:
//
//          PoolArray{ PoolArgs{...} }
type PoolArrayInput interface {
	pulumi.Input

	ToPoolArrayOutput() PoolArrayOutput
	ToPoolArrayOutputWithContext(context.Context) PoolArrayOutput
}

type PoolArray []PoolInput

func (PoolArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Pool)(nil))
}

func (i PoolArray) ToPoolArrayOutput() PoolArrayOutput {
	return i.ToPoolArrayOutputWithContext(context.Background())
}

func (i PoolArray) ToPoolArrayOutputWithContext(ctx context.Context) PoolArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PoolArrayOutput)
}

// PoolMapInput is an input type that accepts PoolMap and PoolMapOutput values.
// You can construct a concrete instance of `PoolMapInput` via:
//
//          PoolMap{ "key": PoolArgs{...} }
type PoolMapInput interface {
	pulumi.Input

	ToPoolMapOutput() PoolMapOutput
	ToPoolMapOutputWithContext(context.Context) PoolMapOutput
}

type PoolMap map[string]PoolInput

func (PoolMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Pool)(nil))
}

func (i PoolMap) ToPoolMapOutput() PoolMapOutput {
	return i.ToPoolMapOutputWithContext(context.Background())
}

func (i PoolMap) ToPoolMapOutputWithContext(ctx context.Context) PoolMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PoolMapOutput)
}

type PoolOutput struct {
	*pulumi.OutputState
}

func (PoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Pool)(nil))
}

func (o PoolOutput) ToPoolOutput() PoolOutput {
	return o
}

func (o PoolOutput) ToPoolOutputWithContext(ctx context.Context) PoolOutput {
	return o
}

func (o PoolOutput) ToPoolPtrOutput() PoolPtrOutput {
	return o.ToPoolPtrOutputWithContext(context.Background())
}

func (o PoolOutput) ToPoolPtrOutputWithContext(ctx context.Context) PoolPtrOutput {
	return o.ApplyT(func(v Pool) *Pool {
		return &v
	}).(PoolPtrOutput)
}

type PoolPtrOutput struct {
	*pulumi.OutputState
}

func (PoolPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Pool)(nil))
}

func (o PoolPtrOutput) ToPoolPtrOutput() PoolPtrOutput {
	return o
}

func (o PoolPtrOutput) ToPoolPtrOutputWithContext(ctx context.Context) PoolPtrOutput {
	return o
}

type PoolArrayOutput struct{ *pulumi.OutputState }

func (PoolArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Pool)(nil))
}

func (o PoolArrayOutput) ToPoolArrayOutput() PoolArrayOutput {
	return o
}

func (o PoolArrayOutput) ToPoolArrayOutputWithContext(ctx context.Context) PoolArrayOutput {
	return o
}

func (o PoolArrayOutput) Index(i pulumi.IntInput) PoolOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Pool {
		return vs[0].([]Pool)[vs[1].(int)]
	}).(PoolOutput)
}

type PoolMapOutput struct{ *pulumi.OutputState }

func (PoolMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Pool)(nil))
}

func (o PoolMapOutput) ToPoolMapOutput() PoolMapOutput {
	return o
}

func (o PoolMapOutput) ToPoolMapOutputWithContext(ctx context.Context) PoolMapOutput {
	return o
}

func (o PoolMapOutput) MapIndex(k pulumi.StringInput) PoolOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Pool {
		return vs[0].(map[string]Pool)[vs[1].(string)]
	}).(PoolOutput)
}

func init() {
	pulumi.RegisterOutputType(PoolOutput{})
	pulumi.RegisterOutputType(PoolPtrOutput{})
	pulumi.RegisterOutputType(PoolArrayOutput{})
	pulumi.RegisterOutputType(PoolMapOutput{})
}
