// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a generic service endpoint within Azure DevOps, which can be used to authenticate to any external server using
// basic authentication via a username and password.
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
// 		_, err = azuredevops.NewServiceEndpointGeneric(ctx, "serviceendpoint", &azuredevops.ServiceEndpointGenericArgs{
// 			ProjectId:           project.ID(),
// 			ServerUrl:           pulumi.String("https://some-server.example.com"),
// 			Username:            pulumi.String("username"),
// 			Password:            pulumi.String("password"),
// 			ServiceEndpointName: pulumi.String("Sample Generic"),
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
// - [Azure DevOps Service REST API 5.1 - Agent Pools](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-5.1)
//
// ## Import
//
// Azure DevOps Service Endpoint Generic can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
//
// ```sh
//  $ pulumi import azuredevops:index/serviceEndpointGeneric:ServiceEndpointGeneric serviceendpoint 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
// ```
type ServiceEndpointGeneric struct {
	pulumi.CustomResourceState

	Authorization pulumi.StringMapOutput `pulumi:"authorization"`
	Description   pulumi.StringPtrOutput `pulumi:"description"`
	// The password or token key used to authenticate to the server url using basic authentication.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// A bcrypted hash of the attribute 'password'
	PasswordHash pulumi.StringOutput `pulumi:"passwordHash"`
	// The project ID or project name.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The URL of the server associated with the service endpoint.
	ServerUrl pulumi.StringOutput `pulumi:"serverUrl"`
	// The service endpoint name.
	ServiceEndpointName pulumi.StringOutput `pulumi:"serviceEndpointName"`
	// The username used to authenticate to the server url using basic authentication.
	Username pulumi.StringPtrOutput `pulumi:"username"`
}

// NewServiceEndpointGeneric registers a new resource with the given unique name, arguments, and options.
func NewServiceEndpointGeneric(ctx *pulumi.Context,
	name string, args *ServiceEndpointGenericArgs, opts ...pulumi.ResourceOption) (*ServiceEndpointGeneric, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.ServerUrl == nil {
		return nil, errors.New("invalid value for required argument 'ServerUrl'")
	}
	if args.ServiceEndpointName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceEndpointName'")
	}
	var resource ServiceEndpointGeneric
	err := ctx.RegisterResource("azuredevops:index/serviceEndpointGeneric:ServiceEndpointGeneric", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceEndpointGeneric gets an existing ServiceEndpointGeneric resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceEndpointGeneric(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceEndpointGenericState, opts ...pulumi.ResourceOption) (*ServiceEndpointGeneric, error) {
	var resource ServiceEndpointGeneric
	err := ctx.ReadResource("azuredevops:index/serviceEndpointGeneric:ServiceEndpointGeneric", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceEndpointGeneric resources.
type serviceEndpointGenericState struct {
	Authorization map[string]string `pulumi:"authorization"`
	Description   *string           `pulumi:"description"`
	// The password or token key used to authenticate to the server url using basic authentication.
	Password *string `pulumi:"password"`
	// A bcrypted hash of the attribute 'password'
	PasswordHash *string `pulumi:"passwordHash"`
	// The project ID or project name.
	ProjectId *string `pulumi:"projectId"`
	// The URL of the server associated with the service endpoint.
	ServerUrl *string `pulumi:"serverUrl"`
	// The service endpoint name.
	ServiceEndpointName *string `pulumi:"serviceEndpointName"`
	// The username used to authenticate to the server url using basic authentication.
	Username *string `pulumi:"username"`
}

type ServiceEndpointGenericState struct {
	Authorization pulumi.StringMapInput
	Description   pulumi.StringPtrInput
	// The password or token key used to authenticate to the server url using basic authentication.
	Password pulumi.StringPtrInput
	// A bcrypted hash of the attribute 'password'
	PasswordHash pulumi.StringPtrInput
	// The project ID or project name.
	ProjectId pulumi.StringPtrInput
	// The URL of the server associated with the service endpoint.
	ServerUrl pulumi.StringPtrInput
	// The service endpoint name.
	ServiceEndpointName pulumi.StringPtrInput
	// The username used to authenticate to the server url using basic authentication.
	Username pulumi.StringPtrInput
}

func (ServiceEndpointGenericState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceEndpointGenericState)(nil)).Elem()
}

type serviceEndpointGenericArgs struct {
	Authorization map[string]string `pulumi:"authorization"`
	Description   *string           `pulumi:"description"`
	// The password or token key used to authenticate to the server url using basic authentication.
	Password *string `pulumi:"password"`
	// The project ID or project name.
	ProjectId string `pulumi:"projectId"`
	// The URL of the server associated with the service endpoint.
	ServerUrl string `pulumi:"serverUrl"`
	// The service endpoint name.
	ServiceEndpointName string `pulumi:"serviceEndpointName"`
	// The username used to authenticate to the server url using basic authentication.
	Username *string `pulumi:"username"`
}

// The set of arguments for constructing a ServiceEndpointGeneric resource.
type ServiceEndpointGenericArgs struct {
	Authorization pulumi.StringMapInput
	Description   pulumi.StringPtrInput
	// The password or token key used to authenticate to the server url using basic authentication.
	Password pulumi.StringPtrInput
	// The project ID or project name.
	ProjectId pulumi.StringInput
	// The URL of the server associated with the service endpoint.
	ServerUrl pulumi.StringInput
	// The service endpoint name.
	ServiceEndpointName pulumi.StringInput
	// The username used to authenticate to the server url using basic authentication.
	Username pulumi.StringPtrInput
}

func (ServiceEndpointGenericArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceEndpointGenericArgs)(nil)).Elem()
}

type ServiceEndpointGenericInput interface {
	pulumi.Input

	ToServiceEndpointGenericOutput() ServiceEndpointGenericOutput
	ToServiceEndpointGenericOutputWithContext(ctx context.Context) ServiceEndpointGenericOutput
}

func (*ServiceEndpointGeneric) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceEndpointGeneric)(nil))
}

func (i *ServiceEndpointGeneric) ToServiceEndpointGenericOutput() ServiceEndpointGenericOutput {
	return i.ToServiceEndpointGenericOutputWithContext(context.Background())
}

func (i *ServiceEndpointGeneric) ToServiceEndpointGenericOutputWithContext(ctx context.Context) ServiceEndpointGenericOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointGenericOutput)
}

func (i *ServiceEndpointGeneric) ToServiceEndpointGenericPtrOutput() ServiceEndpointGenericPtrOutput {
	return i.ToServiceEndpointGenericPtrOutputWithContext(context.Background())
}

func (i *ServiceEndpointGeneric) ToServiceEndpointGenericPtrOutputWithContext(ctx context.Context) ServiceEndpointGenericPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointGenericPtrOutput)
}

type ServiceEndpointGenericPtrInput interface {
	pulumi.Input

	ToServiceEndpointGenericPtrOutput() ServiceEndpointGenericPtrOutput
	ToServiceEndpointGenericPtrOutputWithContext(ctx context.Context) ServiceEndpointGenericPtrOutput
}

type serviceEndpointGenericPtrType ServiceEndpointGenericArgs

func (*serviceEndpointGenericPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceEndpointGeneric)(nil))
}

func (i *serviceEndpointGenericPtrType) ToServiceEndpointGenericPtrOutput() ServiceEndpointGenericPtrOutput {
	return i.ToServiceEndpointGenericPtrOutputWithContext(context.Background())
}

func (i *serviceEndpointGenericPtrType) ToServiceEndpointGenericPtrOutputWithContext(ctx context.Context) ServiceEndpointGenericPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointGenericPtrOutput)
}

// ServiceEndpointGenericArrayInput is an input type that accepts ServiceEndpointGenericArray and ServiceEndpointGenericArrayOutput values.
// You can construct a concrete instance of `ServiceEndpointGenericArrayInput` via:
//
//          ServiceEndpointGenericArray{ ServiceEndpointGenericArgs{...} }
type ServiceEndpointGenericArrayInput interface {
	pulumi.Input

	ToServiceEndpointGenericArrayOutput() ServiceEndpointGenericArrayOutput
	ToServiceEndpointGenericArrayOutputWithContext(context.Context) ServiceEndpointGenericArrayOutput
}

type ServiceEndpointGenericArray []ServiceEndpointGenericInput

func (ServiceEndpointGenericArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*ServiceEndpointGeneric)(nil))
}

func (i ServiceEndpointGenericArray) ToServiceEndpointGenericArrayOutput() ServiceEndpointGenericArrayOutput {
	return i.ToServiceEndpointGenericArrayOutputWithContext(context.Background())
}

func (i ServiceEndpointGenericArray) ToServiceEndpointGenericArrayOutputWithContext(ctx context.Context) ServiceEndpointGenericArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointGenericArrayOutput)
}

// ServiceEndpointGenericMapInput is an input type that accepts ServiceEndpointGenericMap and ServiceEndpointGenericMapOutput values.
// You can construct a concrete instance of `ServiceEndpointGenericMapInput` via:
//
//          ServiceEndpointGenericMap{ "key": ServiceEndpointGenericArgs{...} }
type ServiceEndpointGenericMapInput interface {
	pulumi.Input

	ToServiceEndpointGenericMapOutput() ServiceEndpointGenericMapOutput
	ToServiceEndpointGenericMapOutputWithContext(context.Context) ServiceEndpointGenericMapOutput
}

type ServiceEndpointGenericMap map[string]ServiceEndpointGenericInput

func (ServiceEndpointGenericMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*ServiceEndpointGeneric)(nil))
}

func (i ServiceEndpointGenericMap) ToServiceEndpointGenericMapOutput() ServiceEndpointGenericMapOutput {
	return i.ToServiceEndpointGenericMapOutputWithContext(context.Background())
}

func (i ServiceEndpointGenericMap) ToServiceEndpointGenericMapOutputWithContext(ctx context.Context) ServiceEndpointGenericMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointGenericMapOutput)
}

type ServiceEndpointGenericOutput struct {
	*pulumi.OutputState
}

func (ServiceEndpointGenericOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceEndpointGeneric)(nil))
}

func (o ServiceEndpointGenericOutput) ToServiceEndpointGenericOutput() ServiceEndpointGenericOutput {
	return o
}

func (o ServiceEndpointGenericOutput) ToServiceEndpointGenericOutputWithContext(ctx context.Context) ServiceEndpointGenericOutput {
	return o
}

func (o ServiceEndpointGenericOutput) ToServiceEndpointGenericPtrOutput() ServiceEndpointGenericPtrOutput {
	return o.ToServiceEndpointGenericPtrOutputWithContext(context.Background())
}

func (o ServiceEndpointGenericOutput) ToServiceEndpointGenericPtrOutputWithContext(ctx context.Context) ServiceEndpointGenericPtrOutput {
	return o.ApplyT(func(v ServiceEndpointGeneric) *ServiceEndpointGeneric {
		return &v
	}).(ServiceEndpointGenericPtrOutput)
}

type ServiceEndpointGenericPtrOutput struct {
	*pulumi.OutputState
}

func (ServiceEndpointGenericPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceEndpointGeneric)(nil))
}

func (o ServiceEndpointGenericPtrOutput) ToServiceEndpointGenericPtrOutput() ServiceEndpointGenericPtrOutput {
	return o
}

func (o ServiceEndpointGenericPtrOutput) ToServiceEndpointGenericPtrOutputWithContext(ctx context.Context) ServiceEndpointGenericPtrOutput {
	return o
}

type ServiceEndpointGenericArrayOutput struct{ *pulumi.OutputState }

func (ServiceEndpointGenericArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceEndpointGeneric)(nil))
}

func (o ServiceEndpointGenericArrayOutput) ToServiceEndpointGenericArrayOutput() ServiceEndpointGenericArrayOutput {
	return o
}

func (o ServiceEndpointGenericArrayOutput) ToServiceEndpointGenericArrayOutputWithContext(ctx context.Context) ServiceEndpointGenericArrayOutput {
	return o
}

func (o ServiceEndpointGenericArrayOutput) Index(i pulumi.IntInput) ServiceEndpointGenericOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServiceEndpointGeneric {
		return vs[0].([]ServiceEndpointGeneric)[vs[1].(int)]
	}).(ServiceEndpointGenericOutput)
}

type ServiceEndpointGenericMapOutput struct{ *pulumi.OutputState }

func (ServiceEndpointGenericMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ServiceEndpointGeneric)(nil))
}

func (o ServiceEndpointGenericMapOutput) ToServiceEndpointGenericMapOutput() ServiceEndpointGenericMapOutput {
	return o
}

func (o ServiceEndpointGenericMapOutput) ToServiceEndpointGenericMapOutputWithContext(ctx context.Context) ServiceEndpointGenericMapOutput {
	return o
}

func (o ServiceEndpointGenericMapOutput) MapIndex(k pulumi.StringInput) ServiceEndpointGenericOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ServiceEndpointGeneric {
		return vs[0].(map[string]ServiceEndpointGeneric)[vs[1].(string)]
	}).(ServiceEndpointGenericOutput)
}

func init() {
	pulumi.RegisterOutputType(ServiceEndpointGenericOutput{})
	pulumi.RegisterOutputType(ServiceEndpointGenericPtrOutput{})
	pulumi.RegisterOutputType(ServiceEndpointGenericArrayOutput{})
	pulumi.RegisterOutputType(ServiceEndpointGenericMapOutput{})
}
