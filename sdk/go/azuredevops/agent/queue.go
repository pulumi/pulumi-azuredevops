// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package agent

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azuredevops/sdk/v2/go/azuredevops/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an agent queue within Azure DevOps. In the UI, this is equivalent to adding an
// Organization defined pool to a project.
//
// The created queue is not authorized for use by all pipelines in the project. However,
// the `ResourceAuthorization` resource can be used to grant authorization.
//
// ## Example Usage
// ### Creating a Queue from an organization-level pool
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
//			exampleProject, err := azuredevops.NewProject(ctx, "exampleProject", nil)
//			if err != nil {
//				return err
//			}
//			examplePool, err := azuredevops.LookupPool(ctx, &azuredevops.LookupPoolArgs{
//				Name: "example-pool",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			exampleQueue, err := azuredevops.NewQueue(ctx, "exampleQueue", &azuredevops.QueueArgs{
//				ProjectId:   exampleProject.ID(),
//				AgentPoolId: *pulumi.String(examplePool.Id),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuredevops.NewResourceAuthorization(ctx, "exampleResourceAuthorization", &azuredevops.ResourceAuthorizationArgs{
//				ProjectId:  exampleProject.ID(),
//				ResourceId: exampleQueue.ID(),
//				Type:       pulumi.String("queue"),
//				Authorized: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Creating a Queue at the project level (Organization-level permissions not required)
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
//			_, err = azuredevops.NewQueue(ctx, "exampleQueue", &azuredevops.QueueArgs{
//				ProjectId: *pulumi.String(exampleProject.Id),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ## Relevant Links
//
// - [Azure DevOps Service REST API 7.0 - Agent Queues](https://docs.microsoft.com/en-us/rest/api/azure/devops/distributedtask/queues?view=azure-devops-rest-7.0)
//
// ## Import
//
// Azure DevOps Agent Pools can be imported using the project ID and agent queue ID, e.g.
//
// ```sh
//
//	$ pulumi import azuredevops:Agent/queue:Queue example 00000000-0000-0000-0000-000000000000/0
//
// ```
//
// Deprecated: azuredevops.agent.Queue has been deprecated in favor of azuredevops.Queue
type Queue struct {
	pulumi.CustomResourceState

	// The ID of the organization agent pool. Conflicts with `name`.
	//
	// > **NOTE:**
	// One of `name` or `agentPoolId` must be specified, but not both.
	// When `agentPoolId` is specified, the agent queue name will be derived from the agent pool name.
	AgentPoolId pulumi.IntOutput `pulumi:"agentPoolId"`
	// The name of the agent queue. Defaults to the ID of the agent pool. Conflicts with `agentPoolId`.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which to create the resource.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
}

// NewQueue registers a new resource with the given unique name, arguments, and options.
func NewQueue(ctx *pulumi.Context,
	name string, args *QueueArgs, opts ...pulumi.ResourceOption) (*Queue, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Queue
	err := ctx.RegisterResource("azuredevops:Agent/queue:Queue", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetQueue gets an existing Queue resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetQueue(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *QueueState, opts ...pulumi.ResourceOption) (*Queue, error) {
	var resource Queue
	err := ctx.ReadResource("azuredevops:Agent/queue:Queue", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Queue resources.
type queueState struct {
	// The ID of the organization agent pool. Conflicts with `name`.
	//
	// > **NOTE:**
	// One of `name` or `agentPoolId` must be specified, but not both.
	// When `agentPoolId` is specified, the agent queue name will be derived from the agent pool name.
	AgentPoolId *int `pulumi:"agentPoolId"`
	// The name of the agent queue. Defaults to the ID of the agent pool. Conflicts with `agentPoolId`.
	Name *string `pulumi:"name"`
	// The ID of the project in which to create the resource.
	ProjectId *string `pulumi:"projectId"`
}

type QueueState struct {
	// The ID of the organization agent pool. Conflicts with `name`.
	//
	// > **NOTE:**
	// One of `name` or `agentPoolId` must be specified, but not both.
	// When `agentPoolId` is specified, the agent queue name will be derived from the agent pool name.
	AgentPoolId pulumi.IntPtrInput
	// The name of the agent queue. Defaults to the ID of the agent pool. Conflicts with `agentPoolId`.
	Name pulumi.StringPtrInput
	// The ID of the project in which to create the resource.
	ProjectId pulumi.StringPtrInput
}

func (QueueState) ElementType() reflect.Type {
	return reflect.TypeOf((*queueState)(nil)).Elem()
}

type queueArgs struct {
	// The ID of the organization agent pool. Conflicts with `name`.
	//
	// > **NOTE:**
	// One of `name` or `agentPoolId` must be specified, but not both.
	// When `agentPoolId` is specified, the agent queue name will be derived from the agent pool name.
	AgentPoolId *int `pulumi:"agentPoolId"`
	// The name of the agent queue. Defaults to the ID of the agent pool. Conflicts with `agentPoolId`.
	Name *string `pulumi:"name"`
	// The ID of the project in which to create the resource.
	ProjectId string `pulumi:"projectId"`
}

// The set of arguments for constructing a Queue resource.
type QueueArgs struct {
	// The ID of the organization agent pool. Conflicts with `name`.
	//
	// > **NOTE:**
	// One of `name` or `agentPoolId` must be specified, but not both.
	// When `agentPoolId` is specified, the agent queue name will be derived from the agent pool name.
	AgentPoolId pulumi.IntPtrInput
	// The name of the agent queue. Defaults to the ID of the agent pool. Conflicts with `agentPoolId`.
	Name pulumi.StringPtrInput
	// The ID of the project in which to create the resource.
	ProjectId pulumi.StringInput
}

func (QueueArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*queueArgs)(nil)).Elem()
}

type QueueInput interface {
	pulumi.Input

	ToQueueOutput() QueueOutput
	ToQueueOutputWithContext(ctx context.Context) QueueOutput
}

func (*Queue) ElementType() reflect.Type {
	return reflect.TypeOf((**Queue)(nil)).Elem()
}

func (i *Queue) ToQueueOutput() QueueOutput {
	return i.ToQueueOutputWithContext(context.Background())
}

func (i *Queue) ToQueueOutputWithContext(ctx context.Context) QueueOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueueOutput)
}

// QueueArrayInput is an input type that accepts QueueArray and QueueArrayOutput values.
// You can construct a concrete instance of `QueueArrayInput` via:
//
//	QueueArray{ QueueArgs{...} }
type QueueArrayInput interface {
	pulumi.Input

	ToQueueArrayOutput() QueueArrayOutput
	ToQueueArrayOutputWithContext(context.Context) QueueArrayOutput
}

type QueueArray []QueueInput

func (QueueArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Queue)(nil)).Elem()
}

func (i QueueArray) ToQueueArrayOutput() QueueArrayOutput {
	return i.ToQueueArrayOutputWithContext(context.Background())
}

func (i QueueArray) ToQueueArrayOutputWithContext(ctx context.Context) QueueArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueueArrayOutput)
}

// QueueMapInput is an input type that accepts QueueMap and QueueMapOutput values.
// You can construct a concrete instance of `QueueMapInput` via:
//
//	QueueMap{ "key": QueueArgs{...} }
type QueueMapInput interface {
	pulumi.Input

	ToQueueMapOutput() QueueMapOutput
	ToQueueMapOutputWithContext(context.Context) QueueMapOutput
}

type QueueMap map[string]QueueInput

func (QueueMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Queue)(nil)).Elem()
}

func (i QueueMap) ToQueueMapOutput() QueueMapOutput {
	return i.ToQueueMapOutputWithContext(context.Background())
}

func (i QueueMap) ToQueueMapOutputWithContext(ctx context.Context) QueueMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueueMapOutput)
}

type QueueOutput struct{ *pulumi.OutputState }

func (QueueOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Queue)(nil)).Elem()
}

func (o QueueOutput) ToQueueOutput() QueueOutput {
	return o
}

func (o QueueOutput) ToQueueOutputWithContext(ctx context.Context) QueueOutput {
	return o
}

// The ID of the organization agent pool. Conflicts with `name`.
//
// > **NOTE:**
// One of `name` or `agentPoolId` must be specified, but not both.
// When `agentPoolId` is specified, the agent queue name will be derived from the agent pool name.
func (o QueueOutput) AgentPoolId() pulumi.IntOutput {
	return o.ApplyT(func(v *Queue) pulumi.IntOutput { return v.AgentPoolId }).(pulumi.IntOutput)
}

// The name of the agent queue. Defaults to the ID of the agent pool. Conflicts with `agentPoolId`.
func (o QueueOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Queue) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the project in which to create the resource.
func (o QueueOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *Queue) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

type QueueArrayOutput struct{ *pulumi.OutputState }

func (QueueArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Queue)(nil)).Elem()
}

func (o QueueArrayOutput) ToQueueArrayOutput() QueueArrayOutput {
	return o
}

func (o QueueArrayOutput) ToQueueArrayOutputWithContext(ctx context.Context) QueueArrayOutput {
	return o
}

func (o QueueArrayOutput) Index(i pulumi.IntInput) QueueOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Queue {
		return vs[0].([]*Queue)[vs[1].(int)]
	}).(QueueOutput)
}

type QueueMapOutput struct{ *pulumi.OutputState }

func (QueueMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Queue)(nil)).Elem()
}

func (o QueueMapOutput) ToQueueMapOutput() QueueMapOutput {
	return o
}

func (o QueueMapOutput) ToQueueMapOutputWithContext(ctx context.Context) QueueMapOutput {
	return o
}

func (o QueueMapOutput) MapIndex(k pulumi.StringInput) QueueOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Queue {
		return vs[0].(map[string]*Queue)[vs[1].(string)]
	}).(QueueOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*QueueInput)(nil)).Elem(), &Queue{})
	pulumi.RegisterInputType(reflect.TypeOf((*QueueArrayInput)(nil)).Elem(), QueueArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*QueueMapInput)(nil)).Elem(), QueueMap{})
	pulumi.RegisterOutputType(QueueOutput{})
	pulumi.RegisterOutputType(QueueArrayOutput{})
	pulumi.RegisterOutputType(QueueMapOutput{})
}
