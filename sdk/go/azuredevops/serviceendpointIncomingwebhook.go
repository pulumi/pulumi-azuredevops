// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an Incoming WebHook service endpoint within Azure DevOps, which can be used as a resource in YAML pipelines to subscribe to a webhook event.
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
//			exampleProject, err := azuredevops.NewProject(ctx, "exampleProject", &azuredevops.ProjectArgs{
//				Visibility:       pulumi.String("private"),
//				VersionControl:   pulumi.String("Git"),
//				WorkItemTemplate: pulumi.String("Agile"),
//				Description:      pulumi.String("Managed by Terraform"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuredevops.NewServiceendpointIncomingwebhook(ctx, "exampleServiceendpointIncomingwebhook", &azuredevops.ServiceendpointIncomingwebhookArgs{
//				ProjectId:           exampleProject.ID(),
//				WebhookName:         pulumi.String("example_webhook"),
//				Secret:              pulumi.String("secret"),
//				HttpHeader:          pulumi.String("X-Hub-Signature"),
//				ServiceEndpointName: pulumi.String("Example IncomingWebhook"),
//				Description:         pulumi.String("Managed by Terraform"),
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
// ## Import
//
// Azure DevOps Service Endpoint Incoming WebHook can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
//
// ```sh
//
//	$ pulumi import azuredevops:index/serviceendpointIncomingwebhook:ServiceendpointIncomingwebhook example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
//
// ```
type ServiceendpointIncomingwebhook struct {
	pulumi.CustomResourceState

	Authorization pulumi.StringMapOutput `pulumi:"authorization"`
	Description   pulumi.StringPtrOutput `pulumi:"description"`
	// Http header name on which checksum will be sent.
	HttpHeader pulumi.StringPtrOutput `pulumi:"httpHeader"`
	// The ID of the project. Changing this forces a new Service Connection Incoming WebHook to be created.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Secret for the WebHook. WebHook service will use this secret to calculate the payload checksum.
	Secret pulumi.StringPtrOutput `pulumi:"secret"`
	// The name of the service endpoint. Changing this forces a new Service Connection Incoming WebHook to be created.
	ServiceEndpointName pulumi.StringOutput `pulumi:"serviceEndpointName"`
	// The name of the WebHook.
	WebhookName pulumi.StringOutput `pulumi:"webhookName"`
}

// NewServiceendpointIncomingwebhook registers a new resource with the given unique name, arguments, and options.
func NewServiceendpointIncomingwebhook(ctx *pulumi.Context,
	name string, args *ServiceendpointIncomingwebhookArgs, opts ...pulumi.ResourceOption) (*ServiceendpointIncomingwebhook, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.ServiceEndpointName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceEndpointName'")
	}
	if args.WebhookName == nil {
		return nil, errors.New("invalid value for required argument 'WebhookName'")
	}
	if args.Secret != nil {
		args.Secret = pulumi.ToSecret(args.Secret).(pulumi.StringPtrOutput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"secret",
	})
	opts = append(opts, secrets)
	var resource ServiceendpointIncomingwebhook
	err := ctx.RegisterResource("azuredevops:index/serviceendpointIncomingwebhook:ServiceendpointIncomingwebhook", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceendpointIncomingwebhook gets an existing ServiceendpointIncomingwebhook resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceendpointIncomingwebhook(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceendpointIncomingwebhookState, opts ...pulumi.ResourceOption) (*ServiceendpointIncomingwebhook, error) {
	var resource ServiceendpointIncomingwebhook
	err := ctx.ReadResource("azuredevops:index/serviceendpointIncomingwebhook:ServiceendpointIncomingwebhook", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceendpointIncomingwebhook resources.
type serviceendpointIncomingwebhookState struct {
	Authorization map[string]string `pulumi:"authorization"`
	Description   *string           `pulumi:"description"`
	// Http header name on which checksum will be sent.
	HttpHeader *string `pulumi:"httpHeader"`
	// The ID of the project. Changing this forces a new Service Connection Incoming WebHook to be created.
	ProjectId *string `pulumi:"projectId"`
	// Secret for the WebHook. WebHook service will use this secret to calculate the payload checksum.
	Secret *string `pulumi:"secret"`
	// The name of the service endpoint. Changing this forces a new Service Connection Incoming WebHook to be created.
	ServiceEndpointName *string `pulumi:"serviceEndpointName"`
	// The name of the WebHook.
	WebhookName *string `pulumi:"webhookName"`
}

type ServiceendpointIncomingwebhookState struct {
	Authorization pulumi.StringMapInput
	Description   pulumi.StringPtrInput
	// Http header name on which checksum will be sent.
	HttpHeader pulumi.StringPtrInput
	// The ID of the project. Changing this forces a new Service Connection Incoming WebHook to be created.
	ProjectId pulumi.StringPtrInput
	// Secret for the WebHook. WebHook service will use this secret to calculate the payload checksum.
	Secret pulumi.StringPtrInput
	// The name of the service endpoint. Changing this forces a new Service Connection Incoming WebHook to be created.
	ServiceEndpointName pulumi.StringPtrInput
	// The name of the WebHook.
	WebhookName pulumi.StringPtrInput
}

func (ServiceendpointIncomingwebhookState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceendpointIncomingwebhookState)(nil)).Elem()
}

type serviceendpointIncomingwebhookArgs struct {
	Authorization map[string]string `pulumi:"authorization"`
	Description   *string           `pulumi:"description"`
	// Http header name on which checksum will be sent.
	HttpHeader *string `pulumi:"httpHeader"`
	// The ID of the project. Changing this forces a new Service Connection Incoming WebHook to be created.
	ProjectId string `pulumi:"projectId"`
	// Secret for the WebHook. WebHook service will use this secret to calculate the payload checksum.
	Secret *string `pulumi:"secret"`
	// The name of the service endpoint. Changing this forces a new Service Connection Incoming WebHook to be created.
	ServiceEndpointName string `pulumi:"serviceEndpointName"`
	// The name of the WebHook.
	WebhookName string `pulumi:"webhookName"`
}

// The set of arguments for constructing a ServiceendpointIncomingwebhook resource.
type ServiceendpointIncomingwebhookArgs struct {
	Authorization pulumi.StringMapInput
	Description   pulumi.StringPtrInput
	// Http header name on which checksum will be sent.
	HttpHeader pulumi.StringPtrInput
	// The ID of the project. Changing this forces a new Service Connection Incoming WebHook to be created.
	ProjectId pulumi.StringInput
	// Secret for the WebHook. WebHook service will use this secret to calculate the payload checksum.
	Secret pulumi.StringPtrInput
	// The name of the service endpoint. Changing this forces a new Service Connection Incoming WebHook to be created.
	ServiceEndpointName pulumi.StringInput
	// The name of the WebHook.
	WebhookName pulumi.StringInput
}

func (ServiceendpointIncomingwebhookArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceendpointIncomingwebhookArgs)(nil)).Elem()
}

type ServiceendpointIncomingwebhookInput interface {
	pulumi.Input

	ToServiceendpointIncomingwebhookOutput() ServiceendpointIncomingwebhookOutput
	ToServiceendpointIncomingwebhookOutputWithContext(ctx context.Context) ServiceendpointIncomingwebhookOutput
}

func (*ServiceendpointIncomingwebhook) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceendpointIncomingwebhook)(nil)).Elem()
}

func (i *ServiceendpointIncomingwebhook) ToServiceendpointIncomingwebhookOutput() ServiceendpointIncomingwebhookOutput {
	return i.ToServiceendpointIncomingwebhookOutputWithContext(context.Background())
}

func (i *ServiceendpointIncomingwebhook) ToServiceendpointIncomingwebhookOutputWithContext(ctx context.Context) ServiceendpointIncomingwebhookOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceendpointIncomingwebhookOutput)
}

// ServiceendpointIncomingwebhookArrayInput is an input type that accepts ServiceendpointIncomingwebhookArray and ServiceendpointIncomingwebhookArrayOutput values.
// You can construct a concrete instance of `ServiceendpointIncomingwebhookArrayInput` via:
//
//	ServiceendpointIncomingwebhookArray{ ServiceendpointIncomingwebhookArgs{...} }
type ServiceendpointIncomingwebhookArrayInput interface {
	pulumi.Input

	ToServiceendpointIncomingwebhookArrayOutput() ServiceendpointIncomingwebhookArrayOutput
	ToServiceendpointIncomingwebhookArrayOutputWithContext(context.Context) ServiceendpointIncomingwebhookArrayOutput
}

type ServiceendpointIncomingwebhookArray []ServiceendpointIncomingwebhookInput

func (ServiceendpointIncomingwebhookArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceendpointIncomingwebhook)(nil)).Elem()
}

func (i ServiceendpointIncomingwebhookArray) ToServiceendpointIncomingwebhookArrayOutput() ServiceendpointIncomingwebhookArrayOutput {
	return i.ToServiceendpointIncomingwebhookArrayOutputWithContext(context.Background())
}

func (i ServiceendpointIncomingwebhookArray) ToServiceendpointIncomingwebhookArrayOutputWithContext(ctx context.Context) ServiceendpointIncomingwebhookArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceendpointIncomingwebhookArrayOutput)
}

// ServiceendpointIncomingwebhookMapInput is an input type that accepts ServiceendpointIncomingwebhookMap and ServiceendpointIncomingwebhookMapOutput values.
// You can construct a concrete instance of `ServiceendpointIncomingwebhookMapInput` via:
//
//	ServiceendpointIncomingwebhookMap{ "key": ServiceendpointIncomingwebhookArgs{...} }
type ServiceendpointIncomingwebhookMapInput interface {
	pulumi.Input

	ToServiceendpointIncomingwebhookMapOutput() ServiceendpointIncomingwebhookMapOutput
	ToServiceendpointIncomingwebhookMapOutputWithContext(context.Context) ServiceendpointIncomingwebhookMapOutput
}

type ServiceendpointIncomingwebhookMap map[string]ServiceendpointIncomingwebhookInput

func (ServiceendpointIncomingwebhookMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceendpointIncomingwebhook)(nil)).Elem()
}

func (i ServiceendpointIncomingwebhookMap) ToServiceendpointIncomingwebhookMapOutput() ServiceendpointIncomingwebhookMapOutput {
	return i.ToServiceendpointIncomingwebhookMapOutputWithContext(context.Background())
}

func (i ServiceendpointIncomingwebhookMap) ToServiceendpointIncomingwebhookMapOutputWithContext(ctx context.Context) ServiceendpointIncomingwebhookMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceendpointIncomingwebhookMapOutput)
}

type ServiceendpointIncomingwebhookOutput struct{ *pulumi.OutputState }

func (ServiceendpointIncomingwebhookOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceendpointIncomingwebhook)(nil)).Elem()
}

func (o ServiceendpointIncomingwebhookOutput) ToServiceendpointIncomingwebhookOutput() ServiceendpointIncomingwebhookOutput {
	return o
}

func (o ServiceendpointIncomingwebhookOutput) ToServiceendpointIncomingwebhookOutputWithContext(ctx context.Context) ServiceendpointIncomingwebhookOutput {
	return o
}

func (o ServiceendpointIncomingwebhookOutput) Authorization() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ServiceendpointIncomingwebhook) pulumi.StringMapOutput { return v.Authorization }).(pulumi.StringMapOutput)
}

func (o ServiceendpointIncomingwebhookOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceendpointIncomingwebhook) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Http header name on which checksum will be sent.
func (o ServiceendpointIncomingwebhookOutput) HttpHeader() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceendpointIncomingwebhook) pulumi.StringPtrOutput { return v.HttpHeader }).(pulumi.StringPtrOutput)
}

// The ID of the project. Changing this forces a new Service Connection Incoming WebHook to be created.
func (o ServiceendpointIncomingwebhookOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceendpointIncomingwebhook) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// Secret for the WebHook. WebHook service will use this secret to calculate the payload checksum.
func (o ServiceendpointIncomingwebhookOutput) Secret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceendpointIncomingwebhook) pulumi.StringPtrOutput { return v.Secret }).(pulumi.StringPtrOutput)
}

// The name of the service endpoint. Changing this forces a new Service Connection Incoming WebHook to be created.
func (o ServiceendpointIncomingwebhookOutput) ServiceEndpointName() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceendpointIncomingwebhook) pulumi.StringOutput { return v.ServiceEndpointName }).(pulumi.StringOutput)
}

// The name of the WebHook.
func (o ServiceendpointIncomingwebhookOutput) WebhookName() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceendpointIncomingwebhook) pulumi.StringOutput { return v.WebhookName }).(pulumi.StringOutput)
}

type ServiceendpointIncomingwebhookArrayOutput struct{ *pulumi.OutputState }

func (ServiceendpointIncomingwebhookArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceendpointIncomingwebhook)(nil)).Elem()
}

func (o ServiceendpointIncomingwebhookArrayOutput) ToServiceendpointIncomingwebhookArrayOutput() ServiceendpointIncomingwebhookArrayOutput {
	return o
}

func (o ServiceendpointIncomingwebhookArrayOutput) ToServiceendpointIncomingwebhookArrayOutputWithContext(ctx context.Context) ServiceendpointIncomingwebhookArrayOutput {
	return o
}

func (o ServiceendpointIncomingwebhookArrayOutput) Index(i pulumi.IntInput) ServiceendpointIncomingwebhookOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServiceendpointIncomingwebhook {
		return vs[0].([]*ServiceendpointIncomingwebhook)[vs[1].(int)]
	}).(ServiceendpointIncomingwebhookOutput)
}

type ServiceendpointIncomingwebhookMapOutput struct{ *pulumi.OutputState }

func (ServiceendpointIncomingwebhookMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceendpointIncomingwebhook)(nil)).Elem()
}

func (o ServiceendpointIncomingwebhookMapOutput) ToServiceendpointIncomingwebhookMapOutput() ServiceendpointIncomingwebhookMapOutput {
	return o
}

func (o ServiceendpointIncomingwebhookMapOutput) ToServiceendpointIncomingwebhookMapOutputWithContext(ctx context.Context) ServiceendpointIncomingwebhookMapOutput {
	return o
}

func (o ServiceendpointIncomingwebhookMapOutput) MapIndex(k pulumi.StringInput) ServiceendpointIncomingwebhookOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServiceendpointIncomingwebhook {
		return vs[0].(map[string]*ServiceendpointIncomingwebhook)[vs[1].(string)]
	}).(ServiceendpointIncomingwebhookOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceendpointIncomingwebhookInput)(nil)).Elem(), &ServiceendpointIncomingwebhook{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceendpointIncomingwebhookArrayInput)(nil)).Elem(), ServiceendpointIncomingwebhookArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceendpointIncomingwebhookMapInput)(nil)).Elem(), ServiceendpointIncomingwebhookMap{})
	pulumi.RegisterOutputType(ServiceendpointIncomingwebhookOutput{})
	pulumi.RegisterOutputType(ServiceendpointIncomingwebhookArrayOutput{})
	pulumi.RegisterOutputType(ServiceendpointIncomingwebhookMapOutput{})
}
