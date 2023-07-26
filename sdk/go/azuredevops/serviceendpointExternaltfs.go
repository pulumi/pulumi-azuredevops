// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azuredevops/sdk/v2/go/azuredevops/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an Azure Repos/Team Foundation Server service endpoint within Azure DevOps.
//
// ## Relevant Links
//
// - [Azure DevOps Service REST API 6.0 - Service Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-6.0)
//
// ## Import
//
// Azure DevOps Service Endpoint External TFS can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
//
// ```sh
//
//	$ pulumi import azuredevops:index/serviceendpointExternaltfs:ServiceendpointExternaltfs example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
//
// ```
type ServiceendpointExternaltfs struct {
	pulumi.CustomResourceState

	// An `authPersonal` block as documented below. Allows connecting using a personal access token.
	AuthPersonal  ServiceendpointExternaltfsAuthPersonalOutput `pulumi:"authPersonal"`
	Authorization pulumi.StringMapOutput                       `pulumi:"authorization"`
	// URL of the Azure DevOps organization or the TFS Project Collection to connect to.
	ConnectionUrl pulumi.StringOutput    `pulumi:"connectionUrl"`
	Description   pulumi.StringPtrOutput `pulumi:"description"`
	// The ID of the project.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The Service Endpoint name.
	ServiceEndpointName pulumi.StringOutput `pulumi:"serviceEndpointName"`
}

// NewServiceendpointExternaltfs registers a new resource with the given unique name, arguments, and options.
func NewServiceendpointExternaltfs(ctx *pulumi.Context,
	name string, args *ServiceendpointExternaltfsArgs, opts ...pulumi.ResourceOption) (*ServiceendpointExternaltfs, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AuthPersonal == nil {
		return nil, errors.New("invalid value for required argument 'AuthPersonal'")
	}
	if args.ConnectionUrl == nil {
		return nil, errors.New("invalid value for required argument 'ConnectionUrl'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.ServiceEndpointName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceEndpointName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ServiceendpointExternaltfs
	err := ctx.RegisterResource("azuredevops:index/serviceendpointExternaltfs:ServiceendpointExternaltfs", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceendpointExternaltfs gets an existing ServiceendpointExternaltfs resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceendpointExternaltfs(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceendpointExternaltfsState, opts ...pulumi.ResourceOption) (*ServiceendpointExternaltfs, error) {
	var resource ServiceendpointExternaltfs
	err := ctx.ReadResource("azuredevops:index/serviceendpointExternaltfs:ServiceendpointExternaltfs", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceendpointExternaltfs resources.
type serviceendpointExternaltfsState struct {
	// An `authPersonal` block as documented below. Allows connecting using a personal access token.
	AuthPersonal  *ServiceendpointExternaltfsAuthPersonal `pulumi:"authPersonal"`
	Authorization map[string]string                       `pulumi:"authorization"`
	// URL of the Azure DevOps organization or the TFS Project Collection to connect to.
	ConnectionUrl *string `pulumi:"connectionUrl"`
	Description   *string `pulumi:"description"`
	// The ID of the project.
	ProjectId *string `pulumi:"projectId"`
	// The Service Endpoint name.
	ServiceEndpointName *string `pulumi:"serviceEndpointName"`
}

type ServiceendpointExternaltfsState struct {
	// An `authPersonal` block as documented below. Allows connecting using a personal access token.
	AuthPersonal  ServiceendpointExternaltfsAuthPersonalPtrInput
	Authorization pulumi.StringMapInput
	// URL of the Azure DevOps organization or the TFS Project Collection to connect to.
	ConnectionUrl pulumi.StringPtrInput
	Description   pulumi.StringPtrInput
	// The ID of the project.
	ProjectId pulumi.StringPtrInput
	// The Service Endpoint name.
	ServiceEndpointName pulumi.StringPtrInput
}

func (ServiceendpointExternaltfsState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceendpointExternaltfsState)(nil)).Elem()
}

type serviceendpointExternaltfsArgs struct {
	// An `authPersonal` block as documented below. Allows connecting using a personal access token.
	AuthPersonal  ServiceendpointExternaltfsAuthPersonal `pulumi:"authPersonal"`
	Authorization map[string]string                      `pulumi:"authorization"`
	// URL of the Azure DevOps organization or the TFS Project Collection to connect to.
	ConnectionUrl string  `pulumi:"connectionUrl"`
	Description   *string `pulumi:"description"`
	// The ID of the project.
	ProjectId string `pulumi:"projectId"`
	// The Service Endpoint name.
	ServiceEndpointName string `pulumi:"serviceEndpointName"`
}

// The set of arguments for constructing a ServiceendpointExternaltfs resource.
type ServiceendpointExternaltfsArgs struct {
	// An `authPersonal` block as documented below. Allows connecting using a personal access token.
	AuthPersonal  ServiceendpointExternaltfsAuthPersonalInput
	Authorization pulumi.StringMapInput
	// URL of the Azure DevOps organization or the TFS Project Collection to connect to.
	ConnectionUrl pulumi.StringInput
	Description   pulumi.StringPtrInput
	// The ID of the project.
	ProjectId pulumi.StringInput
	// The Service Endpoint name.
	ServiceEndpointName pulumi.StringInput
}

func (ServiceendpointExternaltfsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceendpointExternaltfsArgs)(nil)).Elem()
}

type ServiceendpointExternaltfsInput interface {
	pulumi.Input

	ToServiceendpointExternaltfsOutput() ServiceendpointExternaltfsOutput
	ToServiceendpointExternaltfsOutputWithContext(ctx context.Context) ServiceendpointExternaltfsOutput
}

func (*ServiceendpointExternaltfs) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceendpointExternaltfs)(nil)).Elem()
}

func (i *ServiceendpointExternaltfs) ToServiceendpointExternaltfsOutput() ServiceendpointExternaltfsOutput {
	return i.ToServiceendpointExternaltfsOutputWithContext(context.Background())
}

func (i *ServiceendpointExternaltfs) ToServiceendpointExternaltfsOutputWithContext(ctx context.Context) ServiceendpointExternaltfsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceendpointExternaltfsOutput)
}

// ServiceendpointExternaltfsArrayInput is an input type that accepts ServiceendpointExternaltfsArray and ServiceendpointExternaltfsArrayOutput values.
// You can construct a concrete instance of `ServiceendpointExternaltfsArrayInput` via:
//
//	ServiceendpointExternaltfsArray{ ServiceendpointExternaltfsArgs{...} }
type ServiceendpointExternaltfsArrayInput interface {
	pulumi.Input

	ToServiceendpointExternaltfsArrayOutput() ServiceendpointExternaltfsArrayOutput
	ToServiceendpointExternaltfsArrayOutputWithContext(context.Context) ServiceendpointExternaltfsArrayOutput
}

type ServiceendpointExternaltfsArray []ServiceendpointExternaltfsInput

func (ServiceendpointExternaltfsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceendpointExternaltfs)(nil)).Elem()
}

func (i ServiceendpointExternaltfsArray) ToServiceendpointExternaltfsArrayOutput() ServiceendpointExternaltfsArrayOutput {
	return i.ToServiceendpointExternaltfsArrayOutputWithContext(context.Background())
}

func (i ServiceendpointExternaltfsArray) ToServiceendpointExternaltfsArrayOutputWithContext(ctx context.Context) ServiceendpointExternaltfsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceendpointExternaltfsArrayOutput)
}

// ServiceendpointExternaltfsMapInput is an input type that accepts ServiceendpointExternaltfsMap and ServiceendpointExternaltfsMapOutput values.
// You can construct a concrete instance of `ServiceendpointExternaltfsMapInput` via:
//
//	ServiceendpointExternaltfsMap{ "key": ServiceendpointExternaltfsArgs{...} }
type ServiceendpointExternaltfsMapInput interface {
	pulumi.Input

	ToServiceendpointExternaltfsMapOutput() ServiceendpointExternaltfsMapOutput
	ToServiceendpointExternaltfsMapOutputWithContext(context.Context) ServiceendpointExternaltfsMapOutput
}

type ServiceendpointExternaltfsMap map[string]ServiceendpointExternaltfsInput

func (ServiceendpointExternaltfsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceendpointExternaltfs)(nil)).Elem()
}

func (i ServiceendpointExternaltfsMap) ToServiceendpointExternaltfsMapOutput() ServiceendpointExternaltfsMapOutput {
	return i.ToServiceendpointExternaltfsMapOutputWithContext(context.Background())
}

func (i ServiceendpointExternaltfsMap) ToServiceendpointExternaltfsMapOutputWithContext(ctx context.Context) ServiceendpointExternaltfsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceendpointExternaltfsMapOutput)
}

type ServiceendpointExternaltfsOutput struct{ *pulumi.OutputState }

func (ServiceendpointExternaltfsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceendpointExternaltfs)(nil)).Elem()
}

func (o ServiceendpointExternaltfsOutput) ToServiceendpointExternaltfsOutput() ServiceendpointExternaltfsOutput {
	return o
}

func (o ServiceendpointExternaltfsOutput) ToServiceendpointExternaltfsOutputWithContext(ctx context.Context) ServiceendpointExternaltfsOutput {
	return o
}

// An `authPersonal` block as documented below. Allows connecting using a personal access token.
func (o ServiceendpointExternaltfsOutput) AuthPersonal() ServiceendpointExternaltfsAuthPersonalOutput {
	return o.ApplyT(func(v *ServiceendpointExternaltfs) ServiceendpointExternaltfsAuthPersonalOutput {
		return v.AuthPersonal
	}).(ServiceendpointExternaltfsAuthPersonalOutput)
}

func (o ServiceendpointExternaltfsOutput) Authorization() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ServiceendpointExternaltfs) pulumi.StringMapOutput { return v.Authorization }).(pulumi.StringMapOutput)
}

// URL of the Azure DevOps organization or the TFS Project Collection to connect to.
func (o ServiceendpointExternaltfsOutput) ConnectionUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceendpointExternaltfs) pulumi.StringOutput { return v.ConnectionUrl }).(pulumi.StringOutput)
}

func (o ServiceendpointExternaltfsOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceendpointExternaltfs) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The ID of the project.
func (o ServiceendpointExternaltfsOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceendpointExternaltfs) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The Service Endpoint name.
func (o ServiceendpointExternaltfsOutput) ServiceEndpointName() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceendpointExternaltfs) pulumi.StringOutput { return v.ServiceEndpointName }).(pulumi.StringOutput)
}

type ServiceendpointExternaltfsArrayOutput struct{ *pulumi.OutputState }

func (ServiceendpointExternaltfsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceendpointExternaltfs)(nil)).Elem()
}

func (o ServiceendpointExternaltfsArrayOutput) ToServiceendpointExternaltfsArrayOutput() ServiceendpointExternaltfsArrayOutput {
	return o
}

func (o ServiceendpointExternaltfsArrayOutput) ToServiceendpointExternaltfsArrayOutputWithContext(ctx context.Context) ServiceendpointExternaltfsArrayOutput {
	return o
}

func (o ServiceendpointExternaltfsArrayOutput) Index(i pulumi.IntInput) ServiceendpointExternaltfsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServiceendpointExternaltfs {
		return vs[0].([]*ServiceendpointExternaltfs)[vs[1].(int)]
	}).(ServiceendpointExternaltfsOutput)
}

type ServiceendpointExternaltfsMapOutput struct{ *pulumi.OutputState }

func (ServiceendpointExternaltfsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceendpointExternaltfs)(nil)).Elem()
}

func (o ServiceendpointExternaltfsMapOutput) ToServiceendpointExternaltfsMapOutput() ServiceendpointExternaltfsMapOutput {
	return o
}

func (o ServiceendpointExternaltfsMapOutput) ToServiceendpointExternaltfsMapOutputWithContext(ctx context.Context) ServiceendpointExternaltfsMapOutput {
	return o
}

func (o ServiceendpointExternaltfsMapOutput) MapIndex(k pulumi.StringInput) ServiceendpointExternaltfsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServiceendpointExternaltfs {
		return vs[0].(map[string]*ServiceendpointExternaltfs)[vs[1].(string)]
	}).(ServiceendpointExternaltfsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceendpointExternaltfsInput)(nil)).Elem(), &ServiceendpointExternaltfs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceendpointExternaltfsArrayInput)(nil)).Elem(), ServiceendpointExternaltfsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceendpointExternaltfsMapInput)(nil)).Elem(), ServiceendpointExternaltfsMap{})
	pulumi.RegisterOutputType(ServiceendpointExternaltfsOutput{})
	pulumi.RegisterOutputType(ServiceendpointExternaltfsArrayOutput{})
	pulumi.RegisterOutputType(ServiceendpointExternaltfsMapOutput{})
}
