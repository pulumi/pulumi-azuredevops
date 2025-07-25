// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azuredevops/sdk/v3/go/azuredevops/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an Octopus Deploy service endpoint within Azure DevOps. Using this service endpoint requires you to install [Octopus Deploy](https://marketplace.visualstudio.com/items?itemName=octopusdeploy.octopus-deploy-build-release-tasks).
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
//			_, err = azuredevops.NewServiceendpointOctopusdeploy(ctx, "example", &azuredevops.ServiceendpointOctopusdeployArgs{
//				ProjectId:           example.ID(),
//				Url:                 pulumi.String("https://octopus.com"),
//				ApiKey:              pulumi.String("000000000000000000000000000000000000"),
//				ServiceEndpointName: pulumi.String("Example Octopus Deploy"),
//				Description:         pulumi.String("Managed by Pulumi"),
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
// - [Azure DevOps Service REST API 7.0 - Agent Pools](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-7.0)
//
// ## Import
//
// Azure DevOps Octopus Deploy Service Endpoint can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
//
// ```sh
// $ pulumi import azuredevops:index/serviceendpointOctopusdeploy:ServiceendpointOctopusdeploy example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
// ```
type ServiceendpointOctopusdeploy struct {
	pulumi.CustomResourceState

	// API key to connect to Octopus Deploy.
	ApiKey        pulumi.StringOutput    `pulumi:"apiKey"`
	Authorization pulumi.StringMapOutput `pulumi:"authorization"`
	Description   pulumi.StringPtrOutput `pulumi:"description"`
	// Whether to ignore SSL errors when connecting to the Octopus server from the agent. Default to `false`.
	IgnoreSslError pulumi.BoolPtrOutput `pulumi:"ignoreSslError"`
	// The ID of the project.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The Service Endpoint name.
	ServiceEndpointName pulumi.StringOutput `pulumi:"serviceEndpointName"`
	// Octopus Server url.
	Url pulumi.StringOutput `pulumi:"url"`
}

// NewServiceendpointOctopusdeploy registers a new resource with the given unique name, arguments, and options.
func NewServiceendpointOctopusdeploy(ctx *pulumi.Context,
	name string, args *ServiceendpointOctopusdeployArgs, opts ...pulumi.ResourceOption) (*ServiceendpointOctopusdeploy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiKey == nil {
		return nil, errors.New("invalid value for required argument 'ApiKey'")
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
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ServiceendpointOctopusdeploy
	err := ctx.RegisterResource("azuredevops:index/serviceendpointOctopusdeploy:ServiceendpointOctopusdeploy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceendpointOctopusdeploy gets an existing ServiceendpointOctopusdeploy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceendpointOctopusdeploy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceendpointOctopusdeployState, opts ...pulumi.ResourceOption) (*ServiceendpointOctopusdeploy, error) {
	var resource ServiceendpointOctopusdeploy
	err := ctx.ReadResource("azuredevops:index/serviceendpointOctopusdeploy:ServiceendpointOctopusdeploy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceendpointOctopusdeploy resources.
type serviceendpointOctopusdeployState struct {
	// API key to connect to Octopus Deploy.
	ApiKey        *string           `pulumi:"apiKey"`
	Authorization map[string]string `pulumi:"authorization"`
	Description   *string           `pulumi:"description"`
	// Whether to ignore SSL errors when connecting to the Octopus server from the agent. Default to `false`.
	IgnoreSslError *bool `pulumi:"ignoreSslError"`
	// The ID of the project.
	ProjectId *string `pulumi:"projectId"`
	// The Service Endpoint name.
	ServiceEndpointName *string `pulumi:"serviceEndpointName"`
	// Octopus Server url.
	Url *string `pulumi:"url"`
}

type ServiceendpointOctopusdeployState struct {
	// API key to connect to Octopus Deploy.
	ApiKey        pulumi.StringPtrInput
	Authorization pulumi.StringMapInput
	Description   pulumi.StringPtrInput
	// Whether to ignore SSL errors when connecting to the Octopus server from the agent. Default to `false`.
	IgnoreSslError pulumi.BoolPtrInput
	// The ID of the project.
	ProjectId pulumi.StringPtrInput
	// The Service Endpoint name.
	ServiceEndpointName pulumi.StringPtrInput
	// Octopus Server url.
	Url pulumi.StringPtrInput
}

func (ServiceendpointOctopusdeployState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceendpointOctopusdeployState)(nil)).Elem()
}

type serviceendpointOctopusdeployArgs struct {
	// API key to connect to Octopus Deploy.
	ApiKey      string  `pulumi:"apiKey"`
	Description *string `pulumi:"description"`
	// Whether to ignore SSL errors when connecting to the Octopus server from the agent. Default to `false`.
	IgnoreSslError *bool `pulumi:"ignoreSslError"`
	// The ID of the project.
	ProjectId string `pulumi:"projectId"`
	// The Service Endpoint name.
	ServiceEndpointName string `pulumi:"serviceEndpointName"`
	// Octopus Server url.
	Url string `pulumi:"url"`
}

// The set of arguments for constructing a ServiceendpointOctopusdeploy resource.
type ServiceendpointOctopusdeployArgs struct {
	// API key to connect to Octopus Deploy.
	ApiKey      pulumi.StringInput
	Description pulumi.StringPtrInput
	// Whether to ignore SSL errors when connecting to the Octopus server from the agent. Default to `false`.
	IgnoreSslError pulumi.BoolPtrInput
	// The ID of the project.
	ProjectId pulumi.StringInput
	// The Service Endpoint name.
	ServiceEndpointName pulumi.StringInput
	// Octopus Server url.
	Url pulumi.StringInput
}

func (ServiceendpointOctopusdeployArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceendpointOctopusdeployArgs)(nil)).Elem()
}

type ServiceendpointOctopusdeployInput interface {
	pulumi.Input

	ToServiceendpointOctopusdeployOutput() ServiceendpointOctopusdeployOutput
	ToServiceendpointOctopusdeployOutputWithContext(ctx context.Context) ServiceendpointOctopusdeployOutput
}

func (*ServiceendpointOctopusdeploy) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceendpointOctopusdeploy)(nil)).Elem()
}

func (i *ServiceendpointOctopusdeploy) ToServiceendpointOctopusdeployOutput() ServiceendpointOctopusdeployOutput {
	return i.ToServiceendpointOctopusdeployOutputWithContext(context.Background())
}

func (i *ServiceendpointOctopusdeploy) ToServiceendpointOctopusdeployOutputWithContext(ctx context.Context) ServiceendpointOctopusdeployOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceendpointOctopusdeployOutput)
}

// ServiceendpointOctopusdeployArrayInput is an input type that accepts ServiceendpointOctopusdeployArray and ServiceendpointOctopusdeployArrayOutput values.
// You can construct a concrete instance of `ServiceendpointOctopusdeployArrayInput` via:
//
//	ServiceendpointOctopusdeployArray{ ServiceendpointOctopusdeployArgs{...} }
type ServiceendpointOctopusdeployArrayInput interface {
	pulumi.Input

	ToServiceendpointOctopusdeployArrayOutput() ServiceendpointOctopusdeployArrayOutput
	ToServiceendpointOctopusdeployArrayOutputWithContext(context.Context) ServiceendpointOctopusdeployArrayOutput
}

type ServiceendpointOctopusdeployArray []ServiceendpointOctopusdeployInput

func (ServiceendpointOctopusdeployArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceendpointOctopusdeploy)(nil)).Elem()
}

func (i ServiceendpointOctopusdeployArray) ToServiceendpointOctopusdeployArrayOutput() ServiceendpointOctopusdeployArrayOutput {
	return i.ToServiceendpointOctopusdeployArrayOutputWithContext(context.Background())
}

func (i ServiceendpointOctopusdeployArray) ToServiceendpointOctopusdeployArrayOutputWithContext(ctx context.Context) ServiceendpointOctopusdeployArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceendpointOctopusdeployArrayOutput)
}

// ServiceendpointOctopusdeployMapInput is an input type that accepts ServiceendpointOctopusdeployMap and ServiceendpointOctopusdeployMapOutput values.
// You can construct a concrete instance of `ServiceendpointOctopusdeployMapInput` via:
//
//	ServiceendpointOctopusdeployMap{ "key": ServiceendpointOctopusdeployArgs{...} }
type ServiceendpointOctopusdeployMapInput interface {
	pulumi.Input

	ToServiceendpointOctopusdeployMapOutput() ServiceendpointOctopusdeployMapOutput
	ToServiceendpointOctopusdeployMapOutputWithContext(context.Context) ServiceendpointOctopusdeployMapOutput
}

type ServiceendpointOctopusdeployMap map[string]ServiceendpointOctopusdeployInput

func (ServiceendpointOctopusdeployMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceendpointOctopusdeploy)(nil)).Elem()
}

func (i ServiceendpointOctopusdeployMap) ToServiceendpointOctopusdeployMapOutput() ServiceendpointOctopusdeployMapOutput {
	return i.ToServiceendpointOctopusdeployMapOutputWithContext(context.Background())
}

func (i ServiceendpointOctopusdeployMap) ToServiceendpointOctopusdeployMapOutputWithContext(ctx context.Context) ServiceendpointOctopusdeployMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceendpointOctopusdeployMapOutput)
}

type ServiceendpointOctopusdeployOutput struct{ *pulumi.OutputState }

func (ServiceendpointOctopusdeployOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceendpointOctopusdeploy)(nil)).Elem()
}

func (o ServiceendpointOctopusdeployOutput) ToServiceendpointOctopusdeployOutput() ServiceendpointOctopusdeployOutput {
	return o
}

func (o ServiceendpointOctopusdeployOutput) ToServiceendpointOctopusdeployOutputWithContext(ctx context.Context) ServiceendpointOctopusdeployOutput {
	return o
}

// API key to connect to Octopus Deploy.
func (o ServiceendpointOctopusdeployOutput) ApiKey() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceendpointOctopusdeploy) pulumi.StringOutput { return v.ApiKey }).(pulumi.StringOutput)
}

func (o ServiceendpointOctopusdeployOutput) Authorization() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ServiceendpointOctopusdeploy) pulumi.StringMapOutput { return v.Authorization }).(pulumi.StringMapOutput)
}

func (o ServiceendpointOctopusdeployOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceendpointOctopusdeploy) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Whether to ignore SSL errors when connecting to the Octopus server from the agent. Default to `false`.
func (o ServiceendpointOctopusdeployOutput) IgnoreSslError() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ServiceendpointOctopusdeploy) pulumi.BoolPtrOutput { return v.IgnoreSslError }).(pulumi.BoolPtrOutput)
}

// The ID of the project.
func (o ServiceendpointOctopusdeployOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceendpointOctopusdeploy) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The Service Endpoint name.
func (o ServiceendpointOctopusdeployOutput) ServiceEndpointName() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceendpointOctopusdeploy) pulumi.StringOutput { return v.ServiceEndpointName }).(pulumi.StringOutput)
}

// Octopus Server url.
func (o ServiceendpointOctopusdeployOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceendpointOctopusdeploy) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

type ServiceendpointOctopusdeployArrayOutput struct{ *pulumi.OutputState }

func (ServiceendpointOctopusdeployArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceendpointOctopusdeploy)(nil)).Elem()
}

func (o ServiceendpointOctopusdeployArrayOutput) ToServiceendpointOctopusdeployArrayOutput() ServiceendpointOctopusdeployArrayOutput {
	return o
}

func (o ServiceendpointOctopusdeployArrayOutput) ToServiceendpointOctopusdeployArrayOutputWithContext(ctx context.Context) ServiceendpointOctopusdeployArrayOutput {
	return o
}

func (o ServiceendpointOctopusdeployArrayOutput) Index(i pulumi.IntInput) ServiceendpointOctopusdeployOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServiceendpointOctopusdeploy {
		return vs[0].([]*ServiceendpointOctopusdeploy)[vs[1].(int)]
	}).(ServiceendpointOctopusdeployOutput)
}

type ServiceendpointOctopusdeployMapOutput struct{ *pulumi.OutputState }

func (ServiceendpointOctopusdeployMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceendpointOctopusdeploy)(nil)).Elem()
}

func (o ServiceendpointOctopusdeployMapOutput) ToServiceendpointOctopusdeployMapOutput() ServiceendpointOctopusdeployMapOutput {
	return o
}

func (o ServiceendpointOctopusdeployMapOutput) ToServiceendpointOctopusdeployMapOutputWithContext(ctx context.Context) ServiceendpointOctopusdeployMapOutput {
	return o
}

func (o ServiceendpointOctopusdeployMapOutput) MapIndex(k pulumi.StringInput) ServiceendpointOctopusdeployOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServiceendpointOctopusdeploy {
		return vs[0].(map[string]*ServiceendpointOctopusdeploy)[vs[1].(string)]
	}).(ServiceendpointOctopusdeployOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceendpointOctopusdeployInput)(nil)).Elem(), &ServiceendpointOctopusdeploy{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceendpointOctopusdeployArrayInput)(nil)).Elem(), ServiceendpointOctopusdeployArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceendpointOctopusdeployMapInput)(nil)).Elem(), ServiceendpointOctopusdeployMap{})
	pulumi.RegisterOutputType(ServiceendpointOctopusdeployOutput{})
	pulumi.RegisterOutputType(ServiceendpointOctopusdeployArrayOutput{})
	pulumi.RegisterOutputType(ServiceendpointOctopusdeployMapOutput{})
}
