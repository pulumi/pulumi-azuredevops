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

// Manages a Nexus IQ service endpoint within Azure DevOps, which can be used as a resource in YAML pipelines to connect to a Nexus IQ instance.
// Nexus IQ is not supported by default, to manage a nexus service connection resource, it is necessary to install the [Nexus Extension](https://marketplace.visualstudio.com/items?itemName=SonatypeIntegrations.nexus-iq-azure-extension) in Azure DevOps.
//
// ## Example Usage
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
//			exampleProject, err := azuredevops.NewProject(ctx, "exampleProject", &azuredevops.ProjectArgs{
//				Visibility:       pulumi.String("private"),
//				VersionControl:   pulumi.String("Git"),
//				WorkItemTemplate: pulumi.String("Agile"),
//				Description:      pulumi.String("Managed by Terraform"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuredevops.NewServiceendpointNexus(ctx, "exampleServiceendpointNexus", &azuredevops.ServiceendpointNexusArgs{
//				ProjectId:           exampleProject.ID(),
//				ServiceEndpointName: pulumi.String("nexus-example"),
//				Description:         pulumi.String("Service Endpoint for 'Nexus IQ' (Managed by Terraform)"),
//				Url:                 pulumi.String("https://example.com"),
//				Username:            pulumi.String("username"),
//				Password:            pulumi.String("password"),
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
// Service Connection Nexus can be imported using the `projectId/id` or or `projectName/id`, e.g.
//
// ```sh
// $ pulumi import azuredevops:index/serviceendpointNexus:ServiceendpointNexus example projectName/00000000-0000-0000-0000-000000000000
// ```
type ServiceendpointNexus struct {
	pulumi.CustomResourceState

	Authorization pulumi.StringMapOutput `pulumi:"authorization"`
	Description   pulumi.StringPtrOutput `pulumi:"description"`
	// The Service Endpoint password to authenticate at the Nexus IQ Instance.
	Password pulumi.StringOutput `pulumi:"password"`
	// The ID of the project. Changing this forces a new Service Connection Nexus to be created.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The name of the service endpoint. Changing this forces a new Service Connection Nexus to be created.
	ServiceEndpointName pulumi.StringOutput `pulumi:"serviceEndpointName"`
	// The Service Endpoint url.
	Url pulumi.StringOutput `pulumi:"url"`
	// The Service Endpoint username to authenticate at the Nexus IQ Instance.
	Username pulumi.StringOutput `pulumi:"username"`
}

// NewServiceendpointNexus registers a new resource with the given unique name, arguments, and options.
func NewServiceendpointNexus(ctx *pulumi.Context,
	name string, args *ServiceendpointNexusArgs, opts ...pulumi.ResourceOption) (*ServiceendpointNexus, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Password == nil {
		return nil, errors.New("invalid value for required argument 'Password'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.ServiceEndpointName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceEndpointName'")
	}
	if args.Url == nil {
		return nil, errors.New("invalid value for required argument 'Url'")
	}
	if args.Username == nil {
		return nil, errors.New("invalid value for required argument 'Username'")
	}
	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ServiceendpointNexus
	err := ctx.RegisterResource("azuredevops:index/serviceendpointNexus:ServiceendpointNexus", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceendpointNexus gets an existing ServiceendpointNexus resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceendpointNexus(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceendpointNexusState, opts ...pulumi.ResourceOption) (*ServiceendpointNexus, error) {
	var resource ServiceendpointNexus
	err := ctx.ReadResource("azuredevops:index/serviceendpointNexus:ServiceendpointNexus", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceendpointNexus resources.
type serviceendpointNexusState struct {
	Authorization map[string]string `pulumi:"authorization"`
	Description   *string           `pulumi:"description"`
	// The Service Endpoint password to authenticate at the Nexus IQ Instance.
	Password *string `pulumi:"password"`
	// The ID of the project. Changing this forces a new Service Connection Nexus to be created.
	ProjectId *string `pulumi:"projectId"`
	// The name of the service endpoint. Changing this forces a new Service Connection Nexus to be created.
	ServiceEndpointName *string `pulumi:"serviceEndpointName"`
	// The Service Endpoint url.
	Url *string `pulumi:"url"`
	// The Service Endpoint username to authenticate at the Nexus IQ Instance.
	Username *string `pulumi:"username"`
}

type ServiceendpointNexusState struct {
	Authorization pulumi.StringMapInput
	Description   pulumi.StringPtrInput
	// The Service Endpoint password to authenticate at the Nexus IQ Instance.
	Password pulumi.StringPtrInput
	// The ID of the project. Changing this forces a new Service Connection Nexus to be created.
	ProjectId pulumi.StringPtrInput
	// The name of the service endpoint. Changing this forces a new Service Connection Nexus to be created.
	ServiceEndpointName pulumi.StringPtrInput
	// The Service Endpoint url.
	Url pulumi.StringPtrInput
	// The Service Endpoint username to authenticate at the Nexus IQ Instance.
	Username pulumi.StringPtrInput
}

func (ServiceendpointNexusState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceendpointNexusState)(nil)).Elem()
}

type serviceendpointNexusArgs struct {
	Authorization map[string]string `pulumi:"authorization"`
	Description   *string           `pulumi:"description"`
	// The Service Endpoint password to authenticate at the Nexus IQ Instance.
	Password string `pulumi:"password"`
	// The ID of the project. Changing this forces a new Service Connection Nexus to be created.
	ProjectId string `pulumi:"projectId"`
	// The name of the service endpoint. Changing this forces a new Service Connection Nexus to be created.
	ServiceEndpointName string `pulumi:"serviceEndpointName"`
	// The Service Endpoint url.
	Url string `pulumi:"url"`
	// The Service Endpoint username to authenticate at the Nexus IQ Instance.
	Username string `pulumi:"username"`
}

// The set of arguments for constructing a ServiceendpointNexus resource.
type ServiceendpointNexusArgs struct {
	Authorization pulumi.StringMapInput
	Description   pulumi.StringPtrInput
	// The Service Endpoint password to authenticate at the Nexus IQ Instance.
	Password pulumi.StringInput
	// The ID of the project. Changing this forces a new Service Connection Nexus to be created.
	ProjectId pulumi.StringInput
	// The name of the service endpoint. Changing this forces a new Service Connection Nexus to be created.
	ServiceEndpointName pulumi.StringInput
	// The Service Endpoint url.
	Url pulumi.StringInput
	// The Service Endpoint username to authenticate at the Nexus IQ Instance.
	Username pulumi.StringInput
}

func (ServiceendpointNexusArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceendpointNexusArgs)(nil)).Elem()
}

type ServiceendpointNexusInput interface {
	pulumi.Input

	ToServiceendpointNexusOutput() ServiceendpointNexusOutput
	ToServiceendpointNexusOutputWithContext(ctx context.Context) ServiceendpointNexusOutput
}

func (*ServiceendpointNexus) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceendpointNexus)(nil)).Elem()
}

func (i *ServiceendpointNexus) ToServiceendpointNexusOutput() ServiceendpointNexusOutput {
	return i.ToServiceendpointNexusOutputWithContext(context.Background())
}

func (i *ServiceendpointNexus) ToServiceendpointNexusOutputWithContext(ctx context.Context) ServiceendpointNexusOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceendpointNexusOutput)
}

// ServiceendpointNexusArrayInput is an input type that accepts ServiceendpointNexusArray and ServiceendpointNexusArrayOutput values.
// You can construct a concrete instance of `ServiceendpointNexusArrayInput` via:
//
//	ServiceendpointNexusArray{ ServiceendpointNexusArgs{...} }
type ServiceendpointNexusArrayInput interface {
	pulumi.Input

	ToServiceendpointNexusArrayOutput() ServiceendpointNexusArrayOutput
	ToServiceendpointNexusArrayOutputWithContext(context.Context) ServiceendpointNexusArrayOutput
}

type ServiceendpointNexusArray []ServiceendpointNexusInput

func (ServiceendpointNexusArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceendpointNexus)(nil)).Elem()
}

func (i ServiceendpointNexusArray) ToServiceendpointNexusArrayOutput() ServiceendpointNexusArrayOutput {
	return i.ToServiceendpointNexusArrayOutputWithContext(context.Background())
}

func (i ServiceendpointNexusArray) ToServiceendpointNexusArrayOutputWithContext(ctx context.Context) ServiceendpointNexusArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceendpointNexusArrayOutput)
}

// ServiceendpointNexusMapInput is an input type that accepts ServiceendpointNexusMap and ServiceendpointNexusMapOutput values.
// You can construct a concrete instance of `ServiceendpointNexusMapInput` via:
//
//	ServiceendpointNexusMap{ "key": ServiceendpointNexusArgs{...} }
type ServiceendpointNexusMapInput interface {
	pulumi.Input

	ToServiceendpointNexusMapOutput() ServiceendpointNexusMapOutput
	ToServiceendpointNexusMapOutputWithContext(context.Context) ServiceendpointNexusMapOutput
}

type ServiceendpointNexusMap map[string]ServiceendpointNexusInput

func (ServiceendpointNexusMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceendpointNexus)(nil)).Elem()
}

func (i ServiceendpointNexusMap) ToServiceendpointNexusMapOutput() ServiceendpointNexusMapOutput {
	return i.ToServiceendpointNexusMapOutputWithContext(context.Background())
}

func (i ServiceendpointNexusMap) ToServiceendpointNexusMapOutputWithContext(ctx context.Context) ServiceendpointNexusMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceendpointNexusMapOutput)
}

type ServiceendpointNexusOutput struct{ *pulumi.OutputState }

func (ServiceendpointNexusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceendpointNexus)(nil)).Elem()
}

func (o ServiceendpointNexusOutput) ToServiceendpointNexusOutput() ServiceendpointNexusOutput {
	return o
}

func (o ServiceendpointNexusOutput) ToServiceendpointNexusOutputWithContext(ctx context.Context) ServiceendpointNexusOutput {
	return o
}

func (o ServiceendpointNexusOutput) Authorization() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ServiceendpointNexus) pulumi.StringMapOutput { return v.Authorization }).(pulumi.StringMapOutput)
}

func (o ServiceendpointNexusOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceendpointNexus) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The Service Endpoint password to authenticate at the Nexus IQ Instance.
func (o ServiceendpointNexusOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceendpointNexus) pulumi.StringOutput { return v.Password }).(pulumi.StringOutput)
}

// The ID of the project. Changing this forces a new Service Connection Nexus to be created.
func (o ServiceendpointNexusOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceendpointNexus) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The name of the service endpoint. Changing this forces a new Service Connection Nexus to be created.
func (o ServiceendpointNexusOutput) ServiceEndpointName() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceendpointNexus) pulumi.StringOutput { return v.ServiceEndpointName }).(pulumi.StringOutput)
}

// The Service Endpoint url.
func (o ServiceendpointNexusOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceendpointNexus) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

// The Service Endpoint username to authenticate at the Nexus IQ Instance.
func (o ServiceendpointNexusOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceendpointNexus) pulumi.StringOutput { return v.Username }).(pulumi.StringOutput)
}

type ServiceendpointNexusArrayOutput struct{ *pulumi.OutputState }

func (ServiceendpointNexusArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceendpointNexus)(nil)).Elem()
}

func (o ServiceendpointNexusArrayOutput) ToServiceendpointNexusArrayOutput() ServiceendpointNexusArrayOutput {
	return o
}

func (o ServiceendpointNexusArrayOutput) ToServiceendpointNexusArrayOutputWithContext(ctx context.Context) ServiceendpointNexusArrayOutput {
	return o
}

func (o ServiceendpointNexusArrayOutput) Index(i pulumi.IntInput) ServiceendpointNexusOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServiceendpointNexus {
		return vs[0].([]*ServiceendpointNexus)[vs[1].(int)]
	}).(ServiceendpointNexusOutput)
}

type ServiceendpointNexusMapOutput struct{ *pulumi.OutputState }

func (ServiceendpointNexusMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceendpointNexus)(nil)).Elem()
}

func (o ServiceendpointNexusMapOutput) ToServiceendpointNexusMapOutput() ServiceendpointNexusMapOutput {
	return o
}

func (o ServiceendpointNexusMapOutput) ToServiceendpointNexusMapOutputWithContext(ctx context.Context) ServiceendpointNexusMapOutput {
	return o
}

func (o ServiceendpointNexusMapOutput) MapIndex(k pulumi.StringInput) ServiceendpointNexusOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServiceendpointNexus {
		return vs[0].(map[string]*ServiceendpointNexus)[vs[1].(string)]
	}).(ServiceendpointNexusOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceendpointNexusInput)(nil)).Elem(), &ServiceendpointNexus{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceendpointNexusArrayInput)(nil)).Elem(), ServiceendpointNexusArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceendpointNexusMapInput)(nil)).Elem(), ServiceendpointNexusMap{})
	pulumi.RegisterOutputType(ServiceendpointNexusOutput{})
	pulumi.RegisterOutputType(ServiceendpointNexusArrayOutput{})
	pulumi.RegisterOutputType(ServiceendpointNexusMapOutput{})
}
