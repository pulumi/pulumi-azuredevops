// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azuredevops/sdk/v2/go/azuredevops/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Manages a Jenkins service endpoint within Azure DevOps, which can be used as a resource in YAML pipelines to connect to a Jenkins instance.
//
// ## Import
//
// Service Connection Jenkins can be imported using the `projectId/id` or or `projectName/id`, e.g.
//
// ```sh
//
//	$ pulumi import azuredevops:index/serviceendpointJenkins:ServiceendpointJenkins example projectName/00000000-0000-0000-0000-000000000000
//
// ```
type ServiceendpointJenkins struct {
	pulumi.CustomResourceState

	// Allows the Jenkins clients to accept self-signed SSL server certificates. Defaults to `false.`
	AcceptUntrustedCerts pulumi.BoolPtrOutput   `pulumi:"acceptUntrustedCerts"`
	Authorization        pulumi.StringMapOutput `pulumi:"authorization"`
	Description          pulumi.StringPtrOutput `pulumi:"description"`
	// The Service Endpoint password to authenticate at the Jenkins Instance.
	Password pulumi.StringOutput `pulumi:"password"`
	// The ID of the project. Changing this forces a new Service Connection Jenkins to be created.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The name of the service endpoint. Changing this forces a new Service Connection Jenkins to be created.
	ServiceEndpointName pulumi.StringOutput `pulumi:"serviceEndpointName"`
	// The Service Endpoint url.
	Url pulumi.StringOutput `pulumi:"url"`
	// The Service Endpoint username to authenticate at the Jenkins Instance.
	Username pulumi.StringOutput `pulumi:"username"`
}

// NewServiceendpointJenkins registers a new resource with the given unique name, arguments, and options.
func NewServiceendpointJenkins(ctx *pulumi.Context,
	name string, args *ServiceendpointJenkinsArgs, opts ...pulumi.ResourceOption) (*ServiceendpointJenkins, error) {
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
	var resource ServiceendpointJenkins
	err := ctx.RegisterResource("azuredevops:index/serviceendpointJenkins:ServiceendpointJenkins", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceendpointJenkins gets an existing ServiceendpointJenkins resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceendpointJenkins(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceendpointJenkinsState, opts ...pulumi.ResourceOption) (*ServiceendpointJenkins, error) {
	var resource ServiceendpointJenkins
	err := ctx.ReadResource("azuredevops:index/serviceendpointJenkins:ServiceendpointJenkins", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceendpointJenkins resources.
type serviceendpointJenkinsState struct {
	// Allows the Jenkins clients to accept self-signed SSL server certificates. Defaults to `false.`
	AcceptUntrustedCerts *bool             `pulumi:"acceptUntrustedCerts"`
	Authorization        map[string]string `pulumi:"authorization"`
	Description          *string           `pulumi:"description"`
	// The Service Endpoint password to authenticate at the Jenkins Instance.
	Password *string `pulumi:"password"`
	// The ID of the project. Changing this forces a new Service Connection Jenkins to be created.
	ProjectId *string `pulumi:"projectId"`
	// The name of the service endpoint. Changing this forces a new Service Connection Jenkins to be created.
	ServiceEndpointName *string `pulumi:"serviceEndpointName"`
	// The Service Endpoint url.
	Url *string `pulumi:"url"`
	// The Service Endpoint username to authenticate at the Jenkins Instance.
	Username *string `pulumi:"username"`
}

type ServiceendpointJenkinsState struct {
	// Allows the Jenkins clients to accept self-signed SSL server certificates. Defaults to `false.`
	AcceptUntrustedCerts pulumi.BoolPtrInput
	Authorization        pulumi.StringMapInput
	Description          pulumi.StringPtrInput
	// The Service Endpoint password to authenticate at the Jenkins Instance.
	Password pulumi.StringPtrInput
	// The ID of the project. Changing this forces a new Service Connection Jenkins to be created.
	ProjectId pulumi.StringPtrInput
	// The name of the service endpoint. Changing this forces a new Service Connection Jenkins to be created.
	ServiceEndpointName pulumi.StringPtrInput
	// The Service Endpoint url.
	Url pulumi.StringPtrInput
	// The Service Endpoint username to authenticate at the Jenkins Instance.
	Username pulumi.StringPtrInput
}

func (ServiceendpointJenkinsState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceendpointJenkinsState)(nil)).Elem()
}

type serviceendpointJenkinsArgs struct {
	// Allows the Jenkins clients to accept self-signed SSL server certificates. Defaults to `false.`
	AcceptUntrustedCerts *bool             `pulumi:"acceptUntrustedCerts"`
	Authorization        map[string]string `pulumi:"authorization"`
	Description          *string           `pulumi:"description"`
	// The Service Endpoint password to authenticate at the Jenkins Instance.
	Password string `pulumi:"password"`
	// The ID of the project. Changing this forces a new Service Connection Jenkins to be created.
	ProjectId string `pulumi:"projectId"`
	// The name of the service endpoint. Changing this forces a new Service Connection Jenkins to be created.
	ServiceEndpointName string `pulumi:"serviceEndpointName"`
	// The Service Endpoint url.
	Url string `pulumi:"url"`
	// The Service Endpoint username to authenticate at the Jenkins Instance.
	Username string `pulumi:"username"`
}

// The set of arguments for constructing a ServiceendpointJenkins resource.
type ServiceendpointJenkinsArgs struct {
	// Allows the Jenkins clients to accept self-signed SSL server certificates. Defaults to `false.`
	AcceptUntrustedCerts pulumi.BoolPtrInput
	Authorization        pulumi.StringMapInput
	Description          pulumi.StringPtrInput
	// The Service Endpoint password to authenticate at the Jenkins Instance.
	Password pulumi.StringInput
	// The ID of the project. Changing this forces a new Service Connection Jenkins to be created.
	ProjectId pulumi.StringInput
	// The name of the service endpoint. Changing this forces a new Service Connection Jenkins to be created.
	ServiceEndpointName pulumi.StringInput
	// The Service Endpoint url.
	Url pulumi.StringInput
	// The Service Endpoint username to authenticate at the Jenkins Instance.
	Username pulumi.StringInput
}

func (ServiceendpointJenkinsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceendpointJenkinsArgs)(nil)).Elem()
}

type ServiceendpointJenkinsInput interface {
	pulumi.Input

	ToServiceendpointJenkinsOutput() ServiceendpointJenkinsOutput
	ToServiceendpointJenkinsOutputWithContext(ctx context.Context) ServiceendpointJenkinsOutput
}

func (*ServiceendpointJenkins) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceendpointJenkins)(nil)).Elem()
}

func (i *ServiceendpointJenkins) ToServiceendpointJenkinsOutput() ServiceendpointJenkinsOutput {
	return i.ToServiceendpointJenkinsOutputWithContext(context.Background())
}

func (i *ServiceendpointJenkins) ToServiceendpointJenkinsOutputWithContext(ctx context.Context) ServiceendpointJenkinsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceendpointJenkinsOutput)
}

func (i *ServiceendpointJenkins) ToOutput(ctx context.Context) pulumix.Output[*ServiceendpointJenkins] {
	return pulumix.Output[*ServiceendpointJenkins]{
		OutputState: i.ToServiceendpointJenkinsOutputWithContext(ctx).OutputState,
	}
}

// ServiceendpointJenkinsArrayInput is an input type that accepts ServiceendpointJenkinsArray and ServiceendpointJenkinsArrayOutput values.
// You can construct a concrete instance of `ServiceendpointJenkinsArrayInput` via:
//
//	ServiceendpointJenkinsArray{ ServiceendpointJenkinsArgs{...} }
type ServiceendpointJenkinsArrayInput interface {
	pulumi.Input

	ToServiceendpointJenkinsArrayOutput() ServiceendpointJenkinsArrayOutput
	ToServiceendpointJenkinsArrayOutputWithContext(context.Context) ServiceendpointJenkinsArrayOutput
}

type ServiceendpointJenkinsArray []ServiceendpointJenkinsInput

func (ServiceendpointJenkinsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceendpointJenkins)(nil)).Elem()
}

func (i ServiceendpointJenkinsArray) ToServiceendpointJenkinsArrayOutput() ServiceendpointJenkinsArrayOutput {
	return i.ToServiceendpointJenkinsArrayOutputWithContext(context.Background())
}

func (i ServiceendpointJenkinsArray) ToServiceendpointJenkinsArrayOutputWithContext(ctx context.Context) ServiceendpointJenkinsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceendpointJenkinsArrayOutput)
}

func (i ServiceendpointJenkinsArray) ToOutput(ctx context.Context) pulumix.Output[[]*ServiceendpointJenkins] {
	return pulumix.Output[[]*ServiceendpointJenkins]{
		OutputState: i.ToServiceendpointJenkinsArrayOutputWithContext(ctx).OutputState,
	}
}

// ServiceendpointJenkinsMapInput is an input type that accepts ServiceendpointJenkinsMap and ServiceendpointJenkinsMapOutput values.
// You can construct a concrete instance of `ServiceendpointJenkinsMapInput` via:
//
//	ServiceendpointJenkinsMap{ "key": ServiceendpointJenkinsArgs{...} }
type ServiceendpointJenkinsMapInput interface {
	pulumi.Input

	ToServiceendpointJenkinsMapOutput() ServiceendpointJenkinsMapOutput
	ToServiceendpointJenkinsMapOutputWithContext(context.Context) ServiceendpointJenkinsMapOutput
}

type ServiceendpointJenkinsMap map[string]ServiceendpointJenkinsInput

func (ServiceendpointJenkinsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceendpointJenkins)(nil)).Elem()
}

func (i ServiceendpointJenkinsMap) ToServiceendpointJenkinsMapOutput() ServiceendpointJenkinsMapOutput {
	return i.ToServiceendpointJenkinsMapOutputWithContext(context.Background())
}

func (i ServiceendpointJenkinsMap) ToServiceendpointJenkinsMapOutputWithContext(ctx context.Context) ServiceendpointJenkinsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceendpointJenkinsMapOutput)
}

func (i ServiceendpointJenkinsMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*ServiceendpointJenkins] {
	return pulumix.Output[map[string]*ServiceendpointJenkins]{
		OutputState: i.ToServiceendpointJenkinsMapOutputWithContext(ctx).OutputState,
	}
}

type ServiceendpointJenkinsOutput struct{ *pulumi.OutputState }

func (ServiceendpointJenkinsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceendpointJenkins)(nil)).Elem()
}

func (o ServiceendpointJenkinsOutput) ToServiceendpointJenkinsOutput() ServiceendpointJenkinsOutput {
	return o
}

func (o ServiceendpointJenkinsOutput) ToServiceendpointJenkinsOutputWithContext(ctx context.Context) ServiceendpointJenkinsOutput {
	return o
}

func (o ServiceendpointJenkinsOutput) ToOutput(ctx context.Context) pulumix.Output[*ServiceendpointJenkins] {
	return pulumix.Output[*ServiceendpointJenkins]{
		OutputState: o.OutputState,
	}
}

// Allows the Jenkins clients to accept self-signed SSL server certificates. Defaults to `false.`
func (o ServiceendpointJenkinsOutput) AcceptUntrustedCerts() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ServiceendpointJenkins) pulumi.BoolPtrOutput { return v.AcceptUntrustedCerts }).(pulumi.BoolPtrOutput)
}

func (o ServiceendpointJenkinsOutput) Authorization() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ServiceendpointJenkins) pulumi.StringMapOutput { return v.Authorization }).(pulumi.StringMapOutput)
}

func (o ServiceendpointJenkinsOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceendpointJenkins) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The Service Endpoint password to authenticate at the Jenkins Instance.
func (o ServiceendpointJenkinsOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceendpointJenkins) pulumi.StringOutput { return v.Password }).(pulumi.StringOutput)
}

// The ID of the project. Changing this forces a new Service Connection Jenkins to be created.
func (o ServiceendpointJenkinsOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceendpointJenkins) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The name of the service endpoint. Changing this forces a new Service Connection Jenkins to be created.
func (o ServiceendpointJenkinsOutput) ServiceEndpointName() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceendpointJenkins) pulumi.StringOutput { return v.ServiceEndpointName }).(pulumi.StringOutput)
}

// The Service Endpoint url.
func (o ServiceendpointJenkinsOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceendpointJenkins) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

// The Service Endpoint username to authenticate at the Jenkins Instance.
func (o ServiceendpointJenkinsOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceendpointJenkins) pulumi.StringOutput { return v.Username }).(pulumi.StringOutput)
}

type ServiceendpointJenkinsArrayOutput struct{ *pulumi.OutputState }

func (ServiceendpointJenkinsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceendpointJenkins)(nil)).Elem()
}

func (o ServiceendpointJenkinsArrayOutput) ToServiceendpointJenkinsArrayOutput() ServiceendpointJenkinsArrayOutput {
	return o
}

func (o ServiceendpointJenkinsArrayOutput) ToServiceendpointJenkinsArrayOutputWithContext(ctx context.Context) ServiceendpointJenkinsArrayOutput {
	return o
}

func (o ServiceendpointJenkinsArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*ServiceendpointJenkins] {
	return pulumix.Output[[]*ServiceendpointJenkins]{
		OutputState: o.OutputState,
	}
}

func (o ServiceendpointJenkinsArrayOutput) Index(i pulumi.IntInput) ServiceendpointJenkinsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServiceendpointJenkins {
		return vs[0].([]*ServiceendpointJenkins)[vs[1].(int)]
	}).(ServiceendpointJenkinsOutput)
}

type ServiceendpointJenkinsMapOutput struct{ *pulumi.OutputState }

func (ServiceendpointJenkinsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceendpointJenkins)(nil)).Elem()
}

func (o ServiceendpointJenkinsMapOutput) ToServiceendpointJenkinsMapOutput() ServiceendpointJenkinsMapOutput {
	return o
}

func (o ServiceendpointJenkinsMapOutput) ToServiceendpointJenkinsMapOutputWithContext(ctx context.Context) ServiceendpointJenkinsMapOutput {
	return o
}

func (o ServiceendpointJenkinsMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*ServiceendpointJenkins] {
	return pulumix.Output[map[string]*ServiceendpointJenkins]{
		OutputState: o.OutputState,
	}
}

func (o ServiceendpointJenkinsMapOutput) MapIndex(k pulumi.StringInput) ServiceendpointJenkinsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServiceendpointJenkins {
		return vs[0].(map[string]*ServiceendpointJenkins)[vs[1].(string)]
	}).(ServiceendpointJenkinsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceendpointJenkinsInput)(nil)).Elem(), &ServiceendpointJenkins{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceendpointJenkinsArrayInput)(nil)).Elem(), ServiceendpointJenkinsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceendpointJenkinsMapInput)(nil)).Elem(), ServiceendpointJenkinsMap{})
	pulumi.RegisterOutputType(ServiceendpointJenkinsOutput{})
	pulumi.RegisterOutputType(ServiceendpointJenkinsArrayOutput{})
	pulumi.RegisterOutputType(ServiceendpointJenkinsMapOutput{})
}
