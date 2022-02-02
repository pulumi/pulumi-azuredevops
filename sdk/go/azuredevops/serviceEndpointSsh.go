// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a SSH service endpoint within Azure DevOps.
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
// 		project, err := azuredevops.NewProject(ctx, "project", &azuredevops.ProjectArgs{
// 			Visibility:       pulumi.String("private"),
// 			VersionControl:   pulumi.String("Git"),
// 			WorkItemTemplate: pulumi.String("Agile"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = azuredevops.NewServiceEndpointSsh(ctx, "test", &azuredevops.ServiceEndpointSshArgs{
// 			ProjectId:           project.ID(),
// 			ServiceEndpointName: pulumi.String("Sample SSH"),
// 			Host:                pulumi.String("1.2.3.4"),
// 			Username:            pulumi.String("username"),
// 			Description:         pulumi.String("Managed by Terraform"),
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
// - [Azure DevOps Service REST API 5.1 - Service Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-5.1)
//
// ## Import
//
// Azure DevOps Service Endpoint SSH can be imported using **projectID/serviceEndpointID** or ** projectName/serviceEndpointID**
//
// ```sh
//  $ pulumi import azuredevops:index/serviceEndpointSsh:ServiceEndpointSsh serviceendpoint 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
// ```
type ServiceEndpointSsh struct {
	pulumi.CustomResourceState

	Authorization pulumi.StringMapOutput `pulumi:"authorization"`
	Description   pulumi.StringPtrOutput `pulumi:"description"`
	// The Host name or IP address of the remote machine.
	Host pulumi.StringOutput `pulumi:"host"`
	// Password for connecting to the endpoint.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// A bcrypted hash of the attribute 'password'
	PasswordHash pulumi.StringOutput `pulumi:"passwordHash"`
	// Port number on the remote machine to use for connecting. Defaults to `22`.
	Port pulumi.IntPtrOutput `pulumi:"port"`
	// Private Key for connecting to the endpoint.
	PrivateKey pulumi.StringPtrOutput `pulumi:"privateKey"`
	// A bcrypted hash of the attribute 'private_key'
	PrivateKeyHash pulumi.StringOutput `pulumi:"privateKeyHash"`
	// The project ID or project name.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The Service Endpoint name.
	ServiceEndpointName pulumi.StringOutput `pulumi:"serviceEndpointName"`
	// Username for connecting to the endpoint.
	Username pulumi.StringOutput `pulumi:"username"`
}

// NewServiceEndpointSsh registers a new resource with the given unique name, arguments, and options.
func NewServiceEndpointSsh(ctx *pulumi.Context,
	name string, args *ServiceEndpointSshArgs, opts ...pulumi.ResourceOption) (*ServiceEndpointSsh, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Host == nil {
		return nil, errors.New("invalid value for required argument 'Host'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.ServiceEndpointName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceEndpointName'")
	}
	if args.Username == nil {
		return nil, errors.New("invalid value for required argument 'Username'")
	}
	var resource ServiceEndpointSsh
	err := ctx.RegisterResource("azuredevops:index/serviceEndpointSsh:ServiceEndpointSsh", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceEndpointSsh gets an existing ServiceEndpointSsh resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceEndpointSsh(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceEndpointSshState, opts ...pulumi.ResourceOption) (*ServiceEndpointSsh, error) {
	var resource ServiceEndpointSsh
	err := ctx.ReadResource("azuredevops:index/serviceEndpointSsh:ServiceEndpointSsh", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceEndpointSsh resources.
type serviceEndpointSshState struct {
	Authorization map[string]string `pulumi:"authorization"`
	Description   *string           `pulumi:"description"`
	// The Host name or IP address of the remote machine.
	Host *string `pulumi:"host"`
	// Password for connecting to the endpoint.
	Password *string `pulumi:"password"`
	// A bcrypted hash of the attribute 'password'
	PasswordHash *string `pulumi:"passwordHash"`
	// Port number on the remote machine to use for connecting. Defaults to `22`.
	Port *int `pulumi:"port"`
	// Private Key for connecting to the endpoint.
	PrivateKey *string `pulumi:"privateKey"`
	// A bcrypted hash of the attribute 'private_key'
	PrivateKeyHash *string `pulumi:"privateKeyHash"`
	// The project ID or project name.
	ProjectId *string `pulumi:"projectId"`
	// The Service Endpoint name.
	ServiceEndpointName *string `pulumi:"serviceEndpointName"`
	// Username for connecting to the endpoint.
	Username *string `pulumi:"username"`
}

type ServiceEndpointSshState struct {
	Authorization pulumi.StringMapInput
	Description   pulumi.StringPtrInput
	// The Host name or IP address of the remote machine.
	Host pulumi.StringPtrInput
	// Password for connecting to the endpoint.
	Password pulumi.StringPtrInput
	// A bcrypted hash of the attribute 'password'
	PasswordHash pulumi.StringPtrInput
	// Port number on the remote machine to use for connecting. Defaults to `22`.
	Port pulumi.IntPtrInput
	// Private Key for connecting to the endpoint.
	PrivateKey pulumi.StringPtrInput
	// A bcrypted hash of the attribute 'private_key'
	PrivateKeyHash pulumi.StringPtrInput
	// The project ID or project name.
	ProjectId pulumi.StringPtrInput
	// The Service Endpoint name.
	ServiceEndpointName pulumi.StringPtrInput
	// Username for connecting to the endpoint.
	Username pulumi.StringPtrInput
}

func (ServiceEndpointSshState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceEndpointSshState)(nil)).Elem()
}

type serviceEndpointSshArgs struct {
	Authorization map[string]string `pulumi:"authorization"`
	Description   *string           `pulumi:"description"`
	// The Host name or IP address of the remote machine.
	Host string `pulumi:"host"`
	// Password for connecting to the endpoint.
	Password *string `pulumi:"password"`
	// Port number on the remote machine to use for connecting. Defaults to `22`.
	Port *int `pulumi:"port"`
	// Private Key for connecting to the endpoint.
	PrivateKey *string `pulumi:"privateKey"`
	// The project ID or project name.
	ProjectId string `pulumi:"projectId"`
	// The Service Endpoint name.
	ServiceEndpointName string `pulumi:"serviceEndpointName"`
	// Username for connecting to the endpoint.
	Username string `pulumi:"username"`
}

// The set of arguments for constructing a ServiceEndpointSsh resource.
type ServiceEndpointSshArgs struct {
	Authorization pulumi.StringMapInput
	Description   pulumi.StringPtrInput
	// The Host name or IP address of the remote machine.
	Host pulumi.StringInput
	// Password for connecting to the endpoint.
	Password pulumi.StringPtrInput
	// Port number on the remote machine to use for connecting. Defaults to `22`.
	Port pulumi.IntPtrInput
	// Private Key for connecting to the endpoint.
	PrivateKey pulumi.StringPtrInput
	// The project ID or project name.
	ProjectId pulumi.StringInput
	// The Service Endpoint name.
	ServiceEndpointName pulumi.StringInput
	// Username for connecting to the endpoint.
	Username pulumi.StringInput
}

func (ServiceEndpointSshArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceEndpointSshArgs)(nil)).Elem()
}

type ServiceEndpointSshInput interface {
	pulumi.Input

	ToServiceEndpointSshOutput() ServiceEndpointSshOutput
	ToServiceEndpointSshOutputWithContext(ctx context.Context) ServiceEndpointSshOutput
}

func (*ServiceEndpointSsh) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceEndpointSsh)(nil)).Elem()
}

func (i *ServiceEndpointSsh) ToServiceEndpointSshOutput() ServiceEndpointSshOutput {
	return i.ToServiceEndpointSshOutputWithContext(context.Background())
}

func (i *ServiceEndpointSsh) ToServiceEndpointSshOutputWithContext(ctx context.Context) ServiceEndpointSshOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointSshOutput)
}

// ServiceEndpointSshArrayInput is an input type that accepts ServiceEndpointSshArray and ServiceEndpointSshArrayOutput values.
// You can construct a concrete instance of `ServiceEndpointSshArrayInput` via:
//
//          ServiceEndpointSshArray{ ServiceEndpointSshArgs{...} }
type ServiceEndpointSshArrayInput interface {
	pulumi.Input

	ToServiceEndpointSshArrayOutput() ServiceEndpointSshArrayOutput
	ToServiceEndpointSshArrayOutputWithContext(context.Context) ServiceEndpointSshArrayOutput
}

type ServiceEndpointSshArray []ServiceEndpointSshInput

func (ServiceEndpointSshArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceEndpointSsh)(nil)).Elem()
}

func (i ServiceEndpointSshArray) ToServiceEndpointSshArrayOutput() ServiceEndpointSshArrayOutput {
	return i.ToServiceEndpointSshArrayOutputWithContext(context.Background())
}

func (i ServiceEndpointSshArray) ToServiceEndpointSshArrayOutputWithContext(ctx context.Context) ServiceEndpointSshArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointSshArrayOutput)
}

// ServiceEndpointSshMapInput is an input type that accepts ServiceEndpointSshMap and ServiceEndpointSshMapOutput values.
// You can construct a concrete instance of `ServiceEndpointSshMapInput` via:
//
//          ServiceEndpointSshMap{ "key": ServiceEndpointSshArgs{...} }
type ServiceEndpointSshMapInput interface {
	pulumi.Input

	ToServiceEndpointSshMapOutput() ServiceEndpointSshMapOutput
	ToServiceEndpointSshMapOutputWithContext(context.Context) ServiceEndpointSshMapOutput
}

type ServiceEndpointSshMap map[string]ServiceEndpointSshInput

func (ServiceEndpointSshMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceEndpointSsh)(nil)).Elem()
}

func (i ServiceEndpointSshMap) ToServiceEndpointSshMapOutput() ServiceEndpointSshMapOutput {
	return i.ToServiceEndpointSshMapOutputWithContext(context.Background())
}

func (i ServiceEndpointSshMap) ToServiceEndpointSshMapOutputWithContext(ctx context.Context) ServiceEndpointSshMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointSshMapOutput)
}

type ServiceEndpointSshOutput struct{ *pulumi.OutputState }

func (ServiceEndpointSshOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceEndpointSsh)(nil)).Elem()
}

func (o ServiceEndpointSshOutput) ToServiceEndpointSshOutput() ServiceEndpointSshOutput {
	return o
}

func (o ServiceEndpointSshOutput) ToServiceEndpointSshOutputWithContext(ctx context.Context) ServiceEndpointSshOutput {
	return o
}

type ServiceEndpointSshArrayOutput struct{ *pulumi.OutputState }

func (ServiceEndpointSshArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceEndpointSsh)(nil)).Elem()
}

func (o ServiceEndpointSshArrayOutput) ToServiceEndpointSshArrayOutput() ServiceEndpointSshArrayOutput {
	return o
}

func (o ServiceEndpointSshArrayOutput) ToServiceEndpointSshArrayOutputWithContext(ctx context.Context) ServiceEndpointSshArrayOutput {
	return o
}

func (o ServiceEndpointSshArrayOutput) Index(i pulumi.IntInput) ServiceEndpointSshOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServiceEndpointSsh {
		return vs[0].([]*ServiceEndpointSsh)[vs[1].(int)]
	}).(ServiceEndpointSshOutput)
}

type ServiceEndpointSshMapOutput struct{ *pulumi.OutputState }

func (ServiceEndpointSshMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceEndpointSsh)(nil)).Elem()
}

func (o ServiceEndpointSshMapOutput) ToServiceEndpointSshMapOutput() ServiceEndpointSshMapOutput {
	return o
}

func (o ServiceEndpointSshMapOutput) ToServiceEndpointSshMapOutputWithContext(ctx context.Context) ServiceEndpointSshMapOutput {
	return o
}

func (o ServiceEndpointSshMapOutput) MapIndex(k pulumi.StringInput) ServiceEndpointSshOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServiceEndpointSsh {
		return vs[0].(map[string]*ServiceEndpointSsh)[vs[1].(string)]
	}).(ServiceEndpointSshOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceEndpointSshInput)(nil)).Elem(), &ServiceEndpointSsh{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceEndpointSshArrayInput)(nil)).Elem(), ServiceEndpointSshArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceEndpointSshMapInput)(nil)).Elem(), ServiceEndpointSshMap{})
	pulumi.RegisterOutputType(ServiceEndpointSshOutput{})
	pulumi.RegisterOutputType(ServiceEndpointSshArrayOutput{})
	pulumi.RegisterOutputType(ServiceEndpointSshMapOutput{})
}
