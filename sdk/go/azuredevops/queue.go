// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an agent queue within Azure DevOps. In the UI, this is equivalent to adding an
// Organization defined pool to a project.
//
// The created queue is not authorized for use by all pipelines in the project. However,
// the `ResourceAuthorization` resource can be used to grant authorization.
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
// 		project, err := azuredevops.NewProject(ctx, "project", nil)
// 		if err != nil {
// 			return err
// 		}
// 		pool, err := azuredevops.LookupPool(ctx, &azuredevops.LookupPoolArgs{
// 			Name: "contoso-pool",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		queue, err := azuredevops.NewQueue(ctx, "queue", &azuredevops.QueueArgs{
// 			ProjectId:   project.ID(),
// 			AgentPoolId: pulumi.String(pool.Id),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = azuredevops.NewResourceAuthorization(ctx, "auth", &azuredevops.ResourceAuthorizationArgs{
// 			ProjectId:  project.ID(),
// 			ResourceId: queue.ID(),
// 			Type:       pulumi.String("queue"),
// 			Authorized: pulumi.Bool(true),
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
// - [Azure DevOps Service REST API 5.1 - Agent Queues](https://docs.microsoft.com/en-us/rest/api/azure/devops/distributedtask/queues?view=azure-devops-rest-5.1)
//
// ## Import
//
// Azure DevOps Agent Pools can be imported using the project ID and agent queue ID, e.g.
//
// ```sh
//  $ pulumi import azuredevops:index/queue:Queue q 00000000-0000-0000-0000-000000000000/0
// ```
type Queue struct {
	pulumi.CustomResourceState

	// The ID of the organization agent pool.
	AgentPoolId pulumi.IntOutput `pulumi:"agentPoolId"`
	// The ID of the project in which to create the resource.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
}

// NewQueue registers a new resource with the given unique name, arguments, and options.
func NewQueue(ctx *pulumi.Context,
	name string, args *QueueArgs, opts ...pulumi.ResourceOption) (*Queue, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AgentPoolId == nil {
		return nil, errors.New("invalid value for required argument 'AgentPoolId'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azuredevops:Agent/queue:Queue"),
		},
	})
	opts = append(opts, aliases)
	var resource Queue
	err := ctx.RegisterResource("azuredevops:index/queue:Queue", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azuredevops:index/queue:Queue", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Queue resources.
type queueState struct {
	// The ID of the organization agent pool.
	AgentPoolId *int `pulumi:"agentPoolId"`
	// The ID of the project in which to create the resource.
	ProjectId *string `pulumi:"projectId"`
}

type QueueState struct {
	// The ID of the organization agent pool.
	AgentPoolId pulumi.IntPtrInput
	// The ID of the project in which to create the resource.
	ProjectId pulumi.StringPtrInput
}

func (QueueState) ElementType() reflect.Type {
	return reflect.TypeOf((*queueState)(nil)).Elem()
}

type queueArgs struct {
	// The ID of the organization agent pool.
	AgentPoolId int `pulumi:"agentPoolId"`
	// The ID of the project in which to create the resource.
	ProjectId string `pulumi:"projectId"`
}

// The set of arguments for constructing a Queue resource.
type QueueArgs struct {
	// The ID of the organization agent pool.
	AgentPoolId pulumi.IntInput
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
	return reflect.TypeOf((*Queue)(nil))
}

func (i *Queue) ToQueueOutput() QueueOutput {
	return i.ToQueueOutputWithContext(context.Background())
}

func (i *Queue) ToQueueOutputWithContext(ctx context.Context) QueueOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueueOutput)
}

func (i *Queue) ToQueuePtrOutput() QueuePtrOutput {
	return i.ToQueuePtrOutputWithContext(context.Background())
}

func (i *Queue) ToQueuePtrOutputWithContext(ctx context.Context) QueuePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueuePtrOutput)
}

type QueuePtrInput interface {
	pulumi.Input

	ToQueuePtrOutput() QueuePtrOutput
	ToQueuePtrOutputWithContext(ctx context.Context) QueuePtrOutput
}

type queuePtrType QueueArgs

func (*queuePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Queue)(nil))
}

func (i *queuePtrType) ToQueuePtrOutput() QueuePtrOutput {
	return i.ToQueuePtrOutputWithContext(context.Background())
}

func (i *queuePtrType) ToQueuePtrOutputWithContext(ctx context.Context) QueuePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueuePtrOutput)
}

// QueueArrayInput is an input type that accepts QueueArray and QueueArrayOutput values.
// You can construct a concrete instance of `QueueArrayInput` via:
//
//          QueueArray{ QueueArgs{...} }
type QueueArrayInput interface {
	pulumi.Input

	ToQueueArrayOutput() QueueArrayOutput
	ToQueueArrayOutputWithContext(context.Context) QueueArrayOutput
}

type QueueArray []QueueInput

func (QueueArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Queue)(nil))
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
//          QueueMap{ "key": QueueArgs{...} }
type QueueMapInput interface {
	pulumi.Input

	ToQueueMapOutput() QueueMapOutput
	ToQueueMapOutputWithContext(context.Context) QueueMapOutput
}

type QueueMap map[string]QueueInput

func (QueueMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Queue)(nil))
}

func (i QueueMap) ToQueueMapOutput() QueueMapOutput {
	return i.ToQueueMapOutputWithContext(context.Background())
}

func (i QueueMap) ToQueueMapOutputWithContext(ctx context.Context) QueueMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueueMapOutput)
}

type QueueOutput struct {
	*pulumi.OutputState
}

func (QueueOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Queue)(nil))
}

func (o QueueOutput) ToQueueOutput() QueueOutput {
	return o
}

func (o QueueOutput) ToQueueOutputWithContext(ctx context.Context) QueueOutput {
	return o
}

func (o QueueOutput) ToQueuePtrOutput() QueuePtrOutput {
	return o.ToQueuePtrOutputWithContext(context.Background())
}

func (o QueueOutput) ToQueuePtrOutputWithContext(ctx context.Context) QueuePtrOutput {
	return o.ApplyT(func(v Queue) *Queue {
		return &v
	}).(QueuePtrOutput)
}

type QueuePtrOutput struct {
	*pulumi.OutputState
}

func (QueuePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Queue)(nil))
}

func (o QueuePtrOutput) ToQueuePtrOutput() QueuePtrOutput {
	return o
}

func (o QueuePtrOutput) ToQueuePtrOutputWithContext(ctx context.Context) QueuePtrOutput {
	return o
}

type QueueArrayOutput struct{ *pulumi.OutputState }

func (QueueArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Queue)(nil))
}

func (o QueueArrayOutput) ToQueueArrayOutput() QueueArrayOutput {
	return o
}

func (o QueueArrayOutput) ToQueueArrayOutputWithContext(ctx context.Context) QueueArrayOutput {
	return o
}

func (o QueueArrayOutput) Index(i pulumi.IntInput) QueueOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Queue {
		return vs[0].([]Queue)[vs[1].(int)]
	}).(QueueOutput)
}

type QueueMapOutput struct{ *pulumi.OutputState }

func (QueueMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Queue)(nil))
}

func (o QueueMapOutput) ToQueueMapOutput() QueueMapOutput {
	return o
}

func (o QueueMapOutput) ToQueueMapOutputWithContext(ctx context.Context) QueueMapOutput {
	return o
}

func (o QueueMapOutput) MapIndex(k pulumi.StringInput) QueueOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Queue {
		return vs[0].(map[string]Queue)[vs[1].(string)]
	}).(QueueOutput)
}

func init() {
	pulumi.RegisterOutputType(QueueOutput{})
	pulumi.RegisterOutputType(QueuePtrOutput{})
	pulumi.RegisterOutputType(QueueArrayOutput{})
	pulumi.RegisterOutputType(QueueMapOutput{})
}
