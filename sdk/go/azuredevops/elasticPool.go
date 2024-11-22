// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azuredevops/sdk/v3/go/azuredevops/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages Elastic pool within Azure DevOps.
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
//			example, err := azuredevops.NewProject(ctx, "example", &azuredevops.ProjectArgs{
//				Name:             pulumi.String("Example Project"),
//				Visibility:       pulumi.String("private"),
//				VersionControl:   pulumi.String("Git"),
//				WorkItemTemplate: pulumi.String("Agile"),
//				Description:      pulumi.String("Managed by Pulumi"),
//			})
//			if err != nil {
//				return err
//			}
//			exampleServiceEndpointAzureRM, err := azuredevops.NewServiceEndpointAzureRM(ctx, "example", &azuredevops.ServiceEndpointAzureRMArgs{
//				ProjectId:                           example.ID(),
//				ServiceEndpointName:                 pulumi.String("Example Azure Connection"),
//				Description:                         pulumi.String("Managed by Pulumi"),
//				ServiceEndpointAuthenticationScheme: pulumi.String("ServicePrincipal"),
//				Credentials: &azuredevops.ServiceEndpointAzureRMCredentialsArgs{
//					Serviceprincipalid:  pulumi.String("00000000-0000-0000-0000-000000000000"),
//					Serviceprincipalkey: pulumi.String("00000000-0000-0000-0000-000000000000"),
//				},
//				AzurermSpnTenantid:      pulumi.String("00000000-0000-0000-0000-000000000000"),
//				AzurermSubscriptionId:   pulumi.String("00000000-0000-0000-0000-000000000000"),
//				AzurermSubscriptionName: pulumi.String("Subscription Name"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuredevops.NewElasticPool(ctx, "example", &azuredevops.ElasticPoolArgs{
//				Name:                 pulumi.String("Example Elastic Pool"),
//				ServiceEndpointId:    exampleServiceEndpointAzureRM.ID(),
//				ServiceEndpointScope: example.ID(),
//				DesiredIdle:          pulumi.Int(2),
//				MaxCapacity:          pulumi.Int(3),
//				AzureResourceId:      pulumi.String("/subscriptions/<Subscription Id>/resourceGroups/<Resource Name>/providers/Microsoft.Compute/virtualMachineScaleSets/<VMSS Name>"),
//			})
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
// - [Azure DevOps Service REST API 7.0 - Elastic Pools](https://learn.microsoft.com/en-us/rest/api/azure/devops/distributedtask/elasticpools/create?view=azure-devops-rest-7.0)
//
// ## Import
//
// Azure DevOps Agent Pools can be imported using the Elastic pool ID, e.g.
//
// ```sh
// $ pulumi import azuredevops:index/elasticPool:ElasticPool example 0
// ```
type ElasticPool struct {
	pulumi.CustomResourceState

	// Set whether agents should be configured to run with interactive UI. Defaults to `false`.
	AgentInteractiveUi pulumi.BoolPtrOutput `pulumi:"agentInteractiveUi"`
	// Specifies whether a queue should be automatically provisioned for each project collection. Defaults to `false`.
	AutoProvision pulumi.BoolPtrOutput `pulumi:"autoProvision"`
	// Specifies whether or not agents within the pool should be automatically updated. Defaults to `true`.
	AutoUpdate pulumi.BoolPtrOutput `pulumi:"autoUpdate"`
	// The ID of the Azure resource.
	AzureResourceId pulumi.StringOutput `pulumi:"azureResourceId"`
	// Number of agents to keep on standby.
	DesiredIdle pulumi.IntOutput `pulumi:"desiredIdle"`
	// Maximum number of virtual machines in the scale set.
	MaxCapacity pulumi.IntOutput `pulumi:"maxCapacity"`
	// The name of the Elastic pool.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project where a new Elastic Pool will be created.
	ProjectId pulumi.StringPtrOutput `pulumi:"projectId"`
	// Tear down virtual machines after every use. Defaults to `false`.
	RecycleAfterEachUse pulumi.BoolPtrOutput `pulumi:"recycleAfterEachUse"`
	// The ID of Service Endpoint used to connect to Azure.
	ServiceEndpointId pulumi.StringOutput `pulumi:"serviceEndpointId"`
	// The Project ID of Service Endpoint belongs to.
	ServiceEndpointScope pulumi.StringOutput `pulumi:"serviceEndpointScope"`
	// Delay in minutes before deleting excess idle agents. Defaults to `30`.
	TimeToLiveMinutes pulumi.IntPtrOutput `pulumi:"timeToLiveMinutes"`
}

// NewElasticPool registers a new resource with the given unique name, arguments, and options.
func NewElasticPool(ctx *pulumi.Context,
	name string, args *ElasticPoolArgs, opts ...pulumi.ResourceOption) (*ElasticPool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AzureResourceId == nil {
		return nil, errors.New("invalid value for required argument 'AzureResourceId'")
	}
	if args.DesiredIdle == nil {
		return nil, errors.New("invalid value for required argument 'DesiredIdle'")
	}
	if args.MaxCapacity == nil {
		return nil, errors.New("invalid value for required argument 'MaxCapacity'")
	}
	if args.ServiceEndpointId == nil {
		return nil, errors.New("invalid value for required argument 'ServiceEndpointId'")
	}
	if args.ServiceEndpointScope == nil {
		return nil, errors.New("invalid value for required argument 'ServiceEndpointScope'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ElasticPool
	err := ctx.RegisterResource("azuredevops:index/elasticPool:ElasticPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetElasticPool gets an existing ElasticPool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetElasticPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ElasticPoolState, opts ...pulumi.ResourceOption) (*ElasticPool, error) {
	var resource ElasticPool
	err := ctx.ReadResource("azuredevops:index/elasticPool:ElasticPool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ElasticPool resources.
type elasticPoolState struct {
	// Set whether agents should be configured to run with interactive UI. Defaults to `false`.
	AgentInteractiveUi *bool `pulumi:"agentInteractiveUi"`
	// Specifies whether a queue should be automatically provisioned for each project collection. Defaults to `false`.
	AutoProvision *bool `pulumi:"autoProvision"`
	// Specifies whether or not agents within the pool should be automatically updated. Defaults to `true`.
	AutoUpdate *bool `pulumi:"autoUpdate"`
	// The ID of the Azure resource.
	AzureResourceId *string `pulumi:"azureResourceId"`
	// Number of agents to keep on standby.
	DesiredIdle *int `pulumi:"desiredIdle"`
	// Maximum number of virtual machines in the scale set.
	MaxCapacity *int `pulumi:"maxCapacity"`
	// The name of the Elastic pool.
	Name *string `pulumi:"name"`
	// The ID of the project where a new Elastic Pool will be created.
	ProjectId *string `pulumi:"projectId"`
	// Tear down virtual machines after every use. Defaults to `false`.
	RecycleAfterEachUse *bool `pulumi:"recycleAfterEachUse"`
	// The ID of Service Endpoint used to connect to Azure.
	ServiceEndpointId *string `pulumi:"serviceEndpointId"`
	// The Project ID of Service Endpoint belongs to.
	ServiceEndpointScope *string `pulumi:"serviceEndpointScope"`
	// Delay in minutes before deleting excess idle agents. Defaults to `30`.
	TimeToLiveMinutes *int `pulumi:"timeToLiveMinutes"`
}

type ElasticPoolState struct {
	// Set whether agents should be configured to run with interactive UI. Defaults to `false`.
	AgentInteractiveUi pulumi.BoolPtrInput
	// Specifies whether a queue should be automatically provisioned for each project collection. Defaults to `false`.
	AutoProvision pulumi.BoolPtrInput
	// Specifies whether or not agents within the pool should be automatically updated. Defaults to `true`.
	AutoUpdate pulumi.BoolPtrInput
	// The ID of the Azure resource.
	AzureResourceId pulumi.StringPtrInput
	// Number of agents to keep on standby.
	DesiredIdle pulumi.IntPtrInput
	// Maximum number of virtual machines in the scale set.
	MaxCapacity pulumi.IntPtrInput
	// The name of the Elastic pool.
	Name pulumi.StringPtrInput
	// The ID of the project where a new Elastic Pool will be created.
	ProjectId pulumi.StringPtrInput
	// Tear down virtual machines after every use. Defaults to `false`.
	RecycleAfterEachUse pulumi.BoolPtrInput
	// The ID of Service Endpoint used to connect to Azure.
	ServiceEndpointId pulumi.StringPtrInput
	// The Project ID of Service Endpoint belongs to.
	ServiceEndpointScope pulumi.StringPtrInput
	// Delay in minutes before deleting excess idle agents. Defaults to `30`.
	TimeToLiveMinutes pulumi.IntPtrInput
}

func (ElasticPoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*elasticPoolState)(nil)).Elem()
}

type elasticPoolArgs struct {
	// Set whether agents should be configured to run with interactive UI. Defaults to `false`.
	AgentInteractiveUi *bool `pulumi:"agentInteractiveUi"`
	// Specifies whether a queue should be automatically provisioned for each project collection. Defaults to `false`.
	AutoProvision *bool `pulumi:"autoProvision"`
	// Specifies whether or not agents within the pool should be automatically updated. Defaults to `true`.
	AutoUpdate *bool `pulumi:"autoUpdate"`
	// The ID of the Azure resource.
	AzureResourceId string `pulumi:"azureResourceId"`
	// Number of agents to keep on standby.
	DesiredIdle int `pulumi:"desiredIdle"`
	// Maximum number of virtual machines in the scale set.
	MaxCapacity int `pulumi:"maxCapacity"`
	// The name of the Elastic pool.
	Name *string `pulumi:"name"`
	// The ID of the project where a new Elastic Pool will be created.
	ProjectId *string `pulumi:"projectId"`
	// Tear down virtual machines after every use. Defaults to `false`.
	RecycleAfterEachUse *bool `pulumi:"recycleAfterEachUse"`
	// The ID of Service Endpoint used to connect to Azure.
	ServiceEndpointId string `pulumi:"serviceEndpointId"`
	// The Project ID of Service Endpoint belongs to.
	ServiceEndpointScope string `pulumi:"serviceEndpointScope"`
	// Delay in minutes before deleting excess idle agents. Defaults to `30`.
	TimeToLiveMinutes *int `pulumi:"timeToLiveMinutes"`
}

// The set of arguments for constructing a ElasticPool resource.
type ElasticPoolArgs struct {
	// Set whether agents should be configured to run with interactive UI. Defaults to `false`.
	AgentInteractiveUi pulumi.BoolPtrInput
	// Specifies whether a queue should be automatically provisioned for each project collection. Defaults to `false`.
	AutoProvision pulumi.BoolPtrInput
	// Specifies whether or not agents within the pool should be automatically updated. Defaults to `true`.
	AutoUpdate pulumi.BoolPtrInput
	// The ID of the Azure resource.
	AzureResourceId pulumi.StringInput
	// Number of agents to keep on standby.
	DesiredIdle pulumi.IntInput
	// Maximum number of virtual machines in the scale set.
	MaxCapacity pulumi.IntInput
	// The name of the Elastic pool.
	Name pulumi.StringPtrInput
	// The ID of the project where a new Elastic Pool will be created.
	ProjectId pulumi.StringPtrInput
	// Tear down virtual machines after every use. Defaults to `false`.
	RecycleAfterEachUse pulumi.BoolPtrInput
	// The ID of Service Endpoint used to connect to Azure.
	ServiceEndpointId pulumi.StringInput
	// The Project ID of Service Endpoint belongs to.
	ServiceEndpointScope pulumi.StringInput
	// Delay in minutes before deleting excess idle agents. Defaults to `30`.
	TimeToLiveMinutes pulumi.IntPtrInput
}

func (ElasticPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*elasticPoolArgs)(nil)).Elem()
}

type ElasticPoolInput interface {
	pulumi.Input

	ToElasticPoolOutput() ElasticPoolOutput
	ToElasticPoolOutputWithContext(ctx context.Context) ElasticPoolOutput
}

func (*ElasticPool) ElementType() reflect.Type {
	return reflect.TypeOf((**ElasticPool)(nil)).Elem()
}

func (i *ElasticPool) ToElasticPoolOutput() ElasticPoolOutput {
	return i.ToElasticPoolOutputWithContext(context.Background())
}

func (i *ElasticPool) ToElasticPoolOutputWithContext(ctx context.Context) ElasticPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ElasticPoolOutput)
}

// ElasticPoolArrayInput is an input type that accepts ElasticPoolArray and ElasticPoolArrayOutput values.
// You can construct a concrete instance of `ElasticPoolArrayInput` via:
//
//	ElasticPoolArray{ ElasticPoolArgs{...} }
type ElasticPoolArrayInput interface {
	pulumi.Input

	ToElasticPoolArrayOutput() ElasticPoolArrayOutput
	ToElasticPoolArrayOutputWithContext(context.Context) ElasticPoolArrayOutput
}

type ElasticPoolArray []ElasticPoolInput

func (ElasticPoolArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ElasticPool)(nil)).Elem()
}

func (i ElasticPoolArray) ToElasticPoolArrayOutput() ElasticPoolArrayOutput {
	return i.ToElasticPoolArrayOutputWithContext(context.Background())
}

func (i ElasticPoolArray) ToElasticPoolArrayOutputWithContext(ctx context.Context) ElasticPoolArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ElasticPoolArrayOutput)
}

// ElasticPoolMapInput is an input type that accepts ElasticPoolMap and ElasticPoolMapOutput values.
// You can construct a concrete instance of `ElasticPoolMapInput` via:
//
//	ElasticPoolMap{ "key": ElasticPoolArgs{...} }
type ElasticPoolMapInput interface {
	pulumi.Input

	ToElasticPoolMapOutput() ElasticPoolMapOutput
	ToElasticPoolMapOutputWithContext(context.Context) ElasticPoolMapOutput
}

type ElasticPoolMap map[string]ElasticPoolInput

func (ElasticPoolMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ElasticPool)(nil)).Elem()
}

func (i ElasticPoolMap) ToElasticPoolMapOutput() ElasticPoolMapOutput {
	return i.ToElasticPoolMapOutputWithContext(context.Background())
}

func (i ElasticPoolMap) ToElasticPoolMapOutputWithContext(ctx context.Context) ElasticPoolMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ElasticPoolMapOutput)
}

type ElasticPoolOutput struct{ *pulumi.OutputState }

func (ElasticPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ElasticPool)(nil)).Elem()
}

func (o ElasticPoolOutput) ToElasticPoolOutput() ElasticPoolOutput {
	return o
}

func (o ElasticPoolOutput) ToElasticPoolOutputWithContext(ctx context.Context) ElasticPoolOutput {
	return o
}

// Set whether agents should be configured to run with interactive UI. Defaults to `false`.
func (o ElasticPoolOutput) AgentInteractiveUi() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ElasticPool) pulumi.BoolPtrOutput { return v.AgentInteractiveUi }).(pulumi.BoolPtrOutput)
}

// Specifies whether a queue should be automatically provisioned for each project collection. Defaults to `false`.
func (o ElasticPoolOutput) AutoProvision() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ElasticPool) pulumi.BoolPtrOutput { return v.AutoProvision }).(pulumi.BoolPtrOutput)
}

// Specifies whether or not agents within the pool should be automatically updated. Defaults to `true`.
func (o ElasticPoolOutput) AutoUpdate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ElasticPool) pulumi.BoolPtrOutput { return v.AutoUpdate }).(pulumi.BoolPtrOutput)
}

// The ID of the Azure resource.
func (o ElasticPoolOutput) AzureResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *ElasticPool) pulumi.StringOutput { return v.AzureResourceId }).(pulumi.StringOutput)
}

// Number of agents to keep on standby.
func (o ElasticPoolOutput) DesiredIdle() pulumi.IntOutput {
	return o.ApplyT(func(v *ElasticPool) pulumi.IntOutput { return v.DesiredIdle }).(pulumi.IntOutput)
}

// Maximum number of virtual machines in the scale set.
func (o ElasticPoolOutput) MaxCapacity() pulumi.IntOutput {
	return o.ApplyT(func(v *ElasticPool) pulumi.IntOutput { return v.MaxCapacity }).(pulumi.IntOutput)
}

// The name of the Elastic pool.
func (o ElasticPoolOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ElasticPool) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the project where a new Elastic Pool will be created.
func (o ElasticPoolOutput) ProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ElasticPool) pulumi.StringPtrOutput { return v.ProjectId }).(pulumi.StringPtrOutput)
}

// Tear down virtual machines after every use. Defaults to `false`.
func (o ElasticPoolOutput) RecycleAfterEachUse() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ElasticPool) pulumi.BoolPtrOutput { return v.RecycleAfterEachUse }).(pulumi.BoolPtrOutput)
}

// The ID of Service Endpoint used to connect to Azure.
func (o ElasticPoolOutput) ServiceEndpointId() pulumi.StringOutput {
	return o.ApplyT(func(v *ElasticPool) pulumi.StringOutput { return v.ServiceEndpointId }).(pulumi.StringOutput)
}

// The Project ID of Service Endpoint belongs to.
func (o ElasticPoolOutput) ServiceEndpointScope() pulumi.StringOutput {
	return o.ApplyT(func(v *ElasticPool) pulumi.StringOutput { return v.ServiceEndpointScope }).(pulumi.StringOutput)
}

// Delay in minutes before deleting excess idle agents. Defaults to `30`.
func (o ElasticPoolOutput) TimeToLiveMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ElasticPool) pulumi.IntPtrOutput { return v.TimeToLiveMinutes }).(pulumi.IntPtrOutput)
}

type ElasticPoolArrayOutput struct{ *pulumi.OutputState }

func (ElasticPoolArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ElasticPool)(nil)).Elem()
}

func (o ElasticPoolArrayOutput) ToElasticPoolArrayOutput() ElasticPoolArrayOutput {
	return o
}

func (o ElasticPoolArrayOutput) ToElasticPoolArrayOutputWithContext(ctx context.Context) ElasticPoolArrayOutput {
	return o
}

func (o ElasticPoolArrayOutput) Index(i pulumi.IntInput) ElasticPoolOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ElasticPool {
		return vs[0].([]*ElasticPool)[vs[1].(int)]
	}).(ElasticPoolOutput)
}

type ElasticPoolMapOutput struct{ *pulumi.OutputState }

func (ElasticPoolMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ElasticPool)(nil)).Elem()
}

func (o ElasticPoolMapOutput) ToElasticPoolMapOutput() ElasticPoolMapOutput {
	return o
}

func (o ElasticPoolMapOutput) ToElasticPoolMapOutputWithContext(ctx context.Context) ElasticPoolMapOutput {
	return o
}

func (o ElasticPoolMapOutput) MapIndex(k pulumi.StringInput) ElasticPoolOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ElasticPool {
		return vs[0].(map[string]*ElasticPool)[vs[1].(string)]
	}).(ElasticPoolOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ElasticPoolInput)(nil)).Elem(), &ElasticPool{})
	pulumi.RegisterInputType(reflect.TypeOf((*ElasticPoolArrayInput)(nil)).Elem(), ElasticPoolArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ElasticPoolMapInput)(nil)).Elem(), ElasticPoolMap{})
	pulumi.RegisterOutputType(ElasticPoolOutput{})
	pulumi.RegisterOutputType(ElasticPoolArrayOutput{})
	pulumi.RegisterOutputType(ElasticPoolMapOutput{})
}
