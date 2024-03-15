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

// Manages a Service Hook Storage Queue Pipelines.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-azure/sdk/v5/go/azure/core"
//	"github.com/pulumi/pulumi-azure/sdk/v5/go/azure/storage"
//	"github.com/pulumi/pulumi-azuredevops/sdk/v3/go/azuredevops"
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
//			exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
//				Location: pulumi.String("West Europe"),
//			})
//			if err != nil {
//				return err
//			}
//			exampleAccount, err := storage.NewAccount(ctx, "exampleAccount", &storage.AccountArgs{
//				ResourceGroupName:      exampleResourceGroup.Name,
//				Location:               exampleResourceGroup.Location,
//				AccountTier:            pulumi.String("Standard"),
//				AccountReplicationType: pulumi.String("LRS"),
//			})
//			if err != nil {
//				return err
//			}
//			exampleQueue, err := storage.NewQueue(ctx, "exampleQueue", &storage.QueueArgs{
//				StorageAccountName: exampleAccount.Name,
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuredevops.NewServicehookStorageQueuePipelines(ctx, "exampleServicehookStorageQueuePipelines", &azuredevops.ServicehookStorageQueuePipelinesArgs{
//				ProjectId:   exampleProject.ID(),
//				AccountName: exampleAccount.Name,
//				AccountKey:  exampleAccount.PrimaryAccessKey,
//				QueueName:   exampleQueue.Name,
//				VisiTimeout: pulumi.Int(30),
//				RunStateChangedEvent: &azuredevops.ServicehookStorageQueuePipelinesRunStateChangedEventArgs{
//					RunStateFilter:  pulumi.String("Completed"),
//					RunResultFilter: pulumi.String("Succeeded"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// An empty configuration block will occur in all events triggering the associated action.
//
// <!--Start PulumiCodeChooser -->
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
//			_, err := azuredevops.NewServicehookStorageQueuePipelines(ctx, "example", &azuredevops.ServicehookStorageQueuePipelinesArgs{
//				ProjectId:            pulumi.Any(azuredevops_project.Example.Id),
//				AccountName:          pulumi.Any(azurerm_storage_account.Example.Name),
//				AccountKey:           pulumi.Any(azurerm_storage_account.Example.Primary_access_key),
//				QueueName:            pulumi.Any(azurerm_storage_queue.Example.Name),
//				VisiTimeout:          pulumi.Int(30),
//				RunStateChangedEvent: nil,
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// Service Hook Storage Queue Pipeliness can be imported using the `resource id`, e.g.
//
// ```sh
// $ pulumi import azuredevops:index/servicehookStorageQueuePipelines:ServicehookStorageQueuePipelines example 00000000-0000-0000-0000-000000000000
// ```
type ServicehookStorageQueuePipelines struct {
	pulumi.CustomResourceState

	// A valid account key from the queue's storage account.
	AccountKey pulumi.StringOutput `pulumi:"accountKey"`
	// The queue's storage account name.
	AccountName pulumi.StringOutput `pulumi:"accountName"`
	// The ID of the associated project. Changing this forces a new Service Hook Storage Queue Pipelines to be created.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The name of the queue that will store the events.
	QueueName pulumi.StringOutput `pulumi:"queueName"`
	// A `runStateChangedEvent` block as defined below. Conflicts with `stageStateChangedEvent`
	RunStateChangedEvent ServicehookStorageQueuePipelinesRunStateChangedEventPtrOutput `pulumi:"runStateChangedEvent"`
	// A `stageStateChangedEvent` block as defined below. Conflicts with `runStateChangedEvent`
	//
	// > **Note** At least one of `runStateChangedEvent` and `stageStateChangedEvent` has to be set.
	StageStateChangedEvent ServicehookStorageQueuePipelinesStageStateChangedEventPtrOutput `pulumi:"stageStateChangedEvent"`
	// event time-to-live - the duration a message can remain in the queue before it's automatically removed. Defaults to `604800`.
	Ttl pulumi.IntPtrOutput `pulumi:"ttl"`
	// event visibility timout - how long a message is invisible to other consumers after it's been dequeued. Defaults to `0`.
	VisiTimeout pulumi.IntPtrOutput `pulumi:"visiTimeout"`
}

// NewServicehookStorageQueuePipelines registers a new resource with the given unique name, arguments, and options.
func NewServicehookStorageQueuePipelines(ctx *pulumi.Context,
	name string, args *ServicehookStorageQueuePipelinesArgs, opts ...pulumi.ResourceOption) (*ServicehookStorageQueuePipelines, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountKey == nil {
		return nil, errors.New("invalid value for required argument 'AccountKey'")
	}
	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.QueueName == nil {
		return nil, errors.New("invalid value for required argument 'QueueName'")
	}
	if args.AccountKey != nil {
		args.AccountKey = pulumi.ToSecret(args.AccountKey).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"accountKey",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ServicehookStorageQueuePipelines
	err := ctx.RegisterResource("azuredevops:index/servicehookStorageQueuePipelines:ServicehookStorageQueuePipelines", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServicehookStorageQueuePipelines gets an existing ServicehookStorageQueuePipelines resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServicehookStorageQueuePipelines(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServicehookStorageQueuePipelinesState, opts ...pulumi.ResourceOption) (*ServicehookStorageQueuePipelines, error) {
	var resource ServicehookStorageQueuePipelines
	err := ctx.ReadResource("azuredevops:index/servicehookStorageQueuePipelines:ServicehookStorageQueuePipelines", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServicehookStorageQueuePipelines resources.
type servicehookStorageQueuePipelinesState struct {
	// A valid account key from the queue's storage account.
	AccountKey *string `pulumi:"accountKey"`
	// The queue's storage account name.
	AccountName *string `pulumi:"accountName"`
	// The ID of the associated project. Changing this forces a new Service Hook Storage Queue Pipelines to be created.
	ProjectId *string `pulumi:"projectId"`
	// The name of the queue that will store the events.
	QueueName *string `pulumi:"queueName"`
	// A `runStateChangedEvent` block as defined below. Conflicts with `stageStateChangedEvent`
	RunStateChangedEvent *ServicehookStorageQueuePipelinesRunStateChangedEvent `pulumi:"runStateChangedEvent"`
	// A `stageStateChangedEvent` block as defined below. Conflicts with `runStateChangedEvent`
	//
	// > **Note** At least one of `runStateChangedEvent` and `stageStateChangedEvent` has to be set.
	StageStateChangedEvent *ServicehookStorageQueuePipelinesStageStateChangedEvent `pulumi:"stageStateChangedEvent"`
	// event time-to-live - the duration a message can remain in the queue before it's automatically removed. Defaults to `604800`.
	Ttl *int `pulumi:"ttl"`
	// event visibility timout - how long a message is invisible to other consumers after it's been dequeued. Defaults to `0`.
	VisiTimeout *int `pulumi:"visiTimeout"`
}

type ServicehookStorageQueuePipelinesState struct {
	// A valid account key from the queue's storage account.
	AccountKey pulumi.StringPtrInput
	// The queue's storage account name.
	AccountName pulumi.StringPtrInput
	// The ID of the associated project. Changing this forces a new Service Hook Storage Queue Pipelines to be created.
	ProjectId pulumi.StringPtrInput
	// The name of the queue that will store the events.
	QueueName pulumi.StringPtrInput
	// A `runStateChangedEvent` block as defined below. Conflicts with `stageStateChangedEvent`
	RunStateChangedEvent ServicehookStorageQueuePipelinesRunStateChangedEventPtrInput
	// A `stageStateChangedEvent` block as defined below. Conflicts with `runStateChangedEvent`
	//
	// > **Note** At least one of `runStateChangedEvent` and `stageStateChangedEvent` has to be set.
	StageStateChangedEvent ServicehookStorageQueuePipelinesStageStateChangedEventPtrInput
	// event time-to-live - the duration a message can remain in the queue before it's automatically removed. Defaults to `604800`.
	Ttl pulumi.IntPtrInput
	// event visibility timout - how long a message is invisible to other consumers after it's been dequeued. Defaults to `0`.
	VisiTimeout pulumi.IntPtrInput
}

func (ServicehookStorageQueuePipelinesState) ElementType() reflect.Type {
	return reflect.TypeOf((*servicehookStorageQueuePipelinesState)(nil)).Elem()
}

type servicehookStorageQueuePipelinesArgs struct {
	// A valid account key from the queue's storage account.
	AccountKey string `pulumi:"accountKey"`
	// The queue's storage account name.
	AccountName string `pulumi:"accountName"`
	// The ID of the associated project. Changing this forces a new Service Hook Storage Queue Pipelines to be created.
	ProjectId string `pulumi:"projectId"`
	// The name of the queue that will store the events.
	QueueName string `pulumi:"queueName"`
	// A `runStateChangedEvent` block as defined below. Conflicts with `stageStateChangedEvent`
	RunStateChangedEvent *ServicehookStorageQueuePipelinesRunStateChangedEvent `pulumi:"runStateChangedEvent"`
	// A `stageStateChangedEvent` block as defined below. Conflicts with `runStateChangedEvent`
	//
	// > **Note** At least one of `runStateChangedEvent` and `stageStateChangedEvent` has to be set.
	StageStateChangedEvent *ServicehookStorageQueuePipelinesStageStateChangedEvent `pulumi:"stageStateChangedEvent"`
	// event time-to-live - the duration a message can remain in the queue before it's automatically removed. Defaults to `604800`.
	Ttl *int `pulumi:"ttl"`
	// event visibility timout - how long a message is invisible to other consumers after it's been dequeued. Defaults to `0`.
	VisiTimeout *int `pulumi:"visiTimeout"`
}

// The set of arguments for constructing a ServicehookStorageQueuePipelines resource.
type ServicehookStorageQueuePipelinesArgs struct {
	// A valid account key from the queue's storage account.
	AccountKey pulumi.StringInput
	// The queue's storage account name.
	AccountName pulumi.StringInput
	// The ID of the associated project. Changing this forces a new Service Hook Storage Queue Pipelines to be created.
	ProjectId pulumi.StringInput
	// The name of the queue that will store the events.
	QueueName pulumi.StringInput
	// A `runStateChangedEvent` block as defined below. Conflicts with `stageStateChangedEvent`
	RunStateChangedEvent ServicehookStorageQueuePipelinesRunStateChangedEventPtrInput
	// A `stageStateChangedEvent` block as defined below. Conflicts with `runStateChangedEvent`
	//
	// > **Note** At least one of `runStateChangedEvent` and `stageStateChangedEvent` has to be set.
	StageStateChangedEvent ServicehookStorageQueuePipelinesStageStateChangedEventPtrInput
	// event time-to-live - the duration a message can remain in the queue before it's automatically removed. Defaults to `604800`.
	Ttl pulumi.IntPtrInput
	// event visibility timout - how long a message is invisible to other consumers after it's been dequeued. Defaults to `0`.
	VisiTimeout pulumi.IntPtrInput
}

func (ServicehookStorageQueuePipelinesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*servicehookStorageQueuePipelinesArgs)(nil)).Elem()
}

type ServicehookStorageQueuePipelinesInput interface {
	pulumi.Input

	ToServicehookStorageQueuePipelinesOutput() ServicehookStorageQueuePipelinesOutput
	ToServicehookStorageQueuePipelinesOutputWithContext(ctx context.Context) ServicehookStorageQueuePipelinesOutput
}

func (*ServicehookStorageQueuePipelines) ElementType() reflect.Type {
	return reflect.TypeOf((**ServicehookStorageQueuePipelines)(nil)).Elem()
}

func (i *ServicehookStorageQueuePipelines) ToServicehookStorageQueuePipelinesOutput() ServicehookStorageQueuePipelinesOutput {
	return i.ToServicehookStorageQueuePipelinesOutputWithContext(context.Background())
}

func (i *ServicehookStorageQueuePipelines) ToServicehookStorageQueuePipelinesOutputWithContext(ctx context.Context) ServicehookStorageQueuePipelinesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicehookStorageQueuePipelinesOutput)
}

// ServicehookStorageQueuePipelinesArrayInput is an input type that accepts ServicehookStorageQueuePipelinesArray and ServicehookStorageQueuePipelinesArrayOutput values.
// You can construct a concrete instance of `ServicehookStorageQueuePipelinesArrayInput` via:
//
//	ServicehookStorageQueuePipelinesArray{ ServicehookStorageQueuePipelinesArgs{...} }
type ServicehookStorageQueuePipelinesArrayInput interface {
	pulumi.Input

	ToServicehookStorageQueuePipelinesArrayOutput() ServicehookStorageQueuePipelinesArrayOutput
	ToServicehookStorageQueuePipelinesArrayOutputWithContext(context.Context) ServicehookStorageQueuePipelinesArrayOutput
}

type ServicehookStorageQueuePipelinesArray []ServicehookStorageQueuePipelinesInput

func (ServicehookStorageQueuePipelinesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServicehookStorageQueuePipelines)(nil)).Elem()
}

func (i ServicehookStorageQueuePipelinesArray) ToServicehookStorageQueuePipelinesArrayOutput() ServicehookStorageQueuePipelinesArrayOutput {
	return i.ToServicehookStorageQueuePipelinesArrayOutputWithContext(context.Background())
}

func (i ServicehookStorageQueuePipelinesArray) ToServicehookStorageQueuePipelinesArrayOutputWithContext(ctx context.Context) ServicehookStorageQueuePipelinesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicehookStorageQueuePipelinesArrayOutput)
}

// ServicehookStorageQueuePipelinesMapInput is an input type that accepts ServicehookStorageQueuePipelinesMap and ServicehookStorageQueuePipelinesMapOutput values.
// You can construct a concrete instance of `ServicehookStorageQueuePipelinesMapInput` via:
//
//	ServicehookStorageQueuePipelinesMap{ "key": ServicehookStorageQueuePipelinesArgs{...} }
type ServicehookStorageQueuePipelinesMapInput interface {
	pulumi.Input

	ToServicehookStorageQueuePipelinesMapOutput() ServicehookStorageQueuePipelinesMapOutput
	ToServicehookStorageQueuePipelinesMapOutputWithContext(context.Context) ServicehookStorageQueuePipelinesMapOutput
}

type ServicehookStorageQueuePipelinesMap map[string]ServicehookStorageQueuePipelinesInput

func (ServicehookStorageQueuePipelinesMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServicehookStorageQueuePipelines)(nil)).Elem()
}

func (i ServicehookStorageQueuePipelinesMap) ToServicehookStorageQueuePipelinesMapOutput() ServicehookStorageQueuePipelinesMapOutput {
	return i.ToServicehookStorageQueuePipelinesMapOutputWithContext(context.Background())
}

func (i ServicehookStorageQueuePipelinesMap) ToServicehookStorageQueuePipelinesMapOutputWithContext(ctx context.Context) ServicehookStorageQueuePipelinesMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicehookStorageQueuePipelinesMapOutput)
}

type ServicehookStorageQueuePipelinesOutput struct{ *pulumi.OutputState }

func (ServicehookStorageQueuePipelinesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServicehookStorageQueuePipelines)(nil)).Elem()
}

func (o ServicehookStorageQueuePipelinesOutput) ToServicehookStorageQueuePipelinesOutput() ServicehookStorageQueuePipelinesOutput {
	return o
}

func (o ServicehookStorageQueuePipelinesOutput) ToServicehookStorageQueuePipelinesOutputWithContext(ctx context.Context) ServicehookStorageQueuePipelinesOutput {
	return o
}

// A valid account key from the queue's storage account.
func (o ServicehookStorageQueuePipelinesOutput) AccountKey() pulumi.StringOutput {
	return o.ApplyT(func(v *ServicehookStorageQueuePipelines) pulumi.StringOutput { return v.AccountKey }).(pulumi.StringOutput)
}

// The queue's storage account name.
func (o ServicehookStorageQueuePipelinesOutput) AccountName() pulumi.StringOutput {
	return o.ApplyT(func(v *ServicehookStorageQueuePipelines) pulumi.StringOutput { return v.AccountName }).(pulumi.StringOutput)
}

// The ID of the associated project. Changing this forces a new Service Hook Storage Queue Pipelines to be created.
func (o ServicehookStorageQueuePipelinesOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServicehookStorageQueuePipelines) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The name of the queue that will store the events.
func (o ServicehookStorageQueuePipelinesOutput) QueueName() pulumi.StringOutput {
	return o.ApplyT(func(v *ServicehookStorageQueuePipelines) pulumi.StringOutput { return v.QueueName }).(pulumi.StringOutput)
}

// A `runStateChangedEvent` block as defined below. Conflicts with `stageStateChangedEvent`
func (o ServicehookStorageQueuePipelinesOutput) RunStateChangedEvent() ServicehookStorageQueuePipelinesRunStateChangedEventPtrOutput {
	return o.ApplyT(func(v *ServicehookStorageQueuePipelines) ServicehookStorageQueuePipelinesRunStateChangedEventPtrOutput {
		return v.RunStateChangedEvent
	}).(ServicehookStorageQueuePipelinesRunStateChangedEventPtrOutput)
}

// A `stageStateChangedEvent` block as defined below. Conflicts with `runStateChangedEvent`
//
// > **Note** At least one of `runStateChangedEvent` and `stageStateChangedEvent` has to be set.
func (o ServicehookStorageQueuePipelinesOutput) StageStateChangedEvent() ServicehookStorageQueuePipelinesStageStateChangedEventPtrOutput {
	return o.ApplyT(func(v *ServicehookStorageQueuePipelines) ServicehookStorageQueuePipelinesStageStateChangedEventPtrOutput {
		return v.StageStateChangedEvent
	}).(ServicehookStorageQueuePipelinesStageStateChangedEventPtrOutput)
}

// event time-to-live - the duration a message can remain in the queue before it's automatically removed. Defaults to `604800`.
func (o ServicehookStorageQueuePipelinesOutput) Ttl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ServicehookStorageQueuePipelines) pulumi.IntPtrOutput { return v.Ttl }).(pulumi.IntPtrOutput)
}

// event visibility timout - how long a message is invisible to other consumers after it's been dequeued. Defaults to `0`.
func (o ServicehookStorageQueuePipelinesOutput) VisiTimeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ServicehookStorageQueuePipelines) pulumi.IntPtrOutput { return v.VisiTimeout }).(pulumi.IntPtrOutput)
}

type ServicehookStorageQueuePipelinesArrayOutput struct{ *pulumi.OutputState }

func (ServicehookStorageQueuePipelinesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServicehookStorageQueuePipelines)(nil)).Elem()
}

func (o ServicehookStorageQueuePipelinesArrayOutput) ToServicehookStorageQueuePipelinesArrayOutput() ServicehookStorageQueuePipelinesArrayOutput {
	return o
}

func (o ServicehookStorageQueuePipelinesArrayOutput) ToServicehookStorageQueuePipelinesArrayOutputWithContext(ctx context.Context) ServicehookStorageQueuePipelinesArrayOutput {
	return o
}

func (o ServicehookStorageQueuePipelinesArrayOutput) Index(i pulumi.IntInput) ServicehookStorageQueuePipelinesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServicehookStorageQueuePipelines {
		return vs[0].([]*ServicehookStorageQueuePipelines)[vs[1].(int)]
	}).(ServicehookStorageQueuePipelinesOutput)
}

type ServicehookStorageQueuePipelinesMapOutput struct{ *pulumi.OutputState }

func (ServicehookStorageQueuePipelinesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServicehookStorageQueuePipelines)(nil)).Elem()
}

func (o ServicehookStorageQueuePipelinesMapOutput) ToServicehookStorageQueuePipelinesMapOutput() ServicehookStorageQueuePipelinesMapOutput {
	return o
}

func (o ServicehookStorageQueuePipelinesMapOutput) ToServicehookStorageQueuePipelinesMapOutputWithContext(ctx context.Context) ServicehookStorageQueuePipelinesMapOutput {
	return o
}

func (o ServicehookStorageQueuePipelinesMapOutput) MapIndex(k pulumi.StringInput) ServicehookStorageQueuePipelinesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServicehookStorageQueuePipelines {
		return vs[0].(map[string]*ServicehookStorageQueuePipelines)[vs[1].(string)]
	}).(ServicehookStorageQueuePipelinesOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServicehookStorageQueuePipelinesInput)(nil)).Elem(), &ServicehookStorageQueuePipelines{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServicehookStorageQueuePipelinesArrayInput)(nil)).Elem(), ServicehookStorageQueuePipelinesArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServicehookStorageQueuePipelinesMapInput)(nil)).Elem(), ServicehookStorageQueuePipelinesMap{})
	pulumi.RegisterOutputType(ServicehookStorageQueuePipelinesOutput{})
	pulumi.RegisterOutputType(ServicehookStorageQueuePipelinesArrayOutput{})
	pulumi.RegisterOutputType(ServicehookStorageQueuePipelinesMapOutput{})
}
