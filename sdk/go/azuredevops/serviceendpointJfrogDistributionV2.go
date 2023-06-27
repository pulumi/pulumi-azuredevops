// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a JFrog Distribution V2 server endpoint within an Azure DevOps organization.
//
// > **Note:** Using this service endpoint requires you to first install [JFrog Extension](https://marketplace.visualstudio.com/items?itemName=JFrog.jfrog-azure-devops-extension).
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
//			_, err = azuredevops.NewServiceendpointJfrogDistributionV2(ctx, "exampleServiceendpointJfrogDistributionV2", &azuredevops.ServiceendpointJfrogDistributionV2Args{
//				ProjectId:           exampleProject.ID(),
//				ServiceEndpointName: pulumi.String("Example JFrog Distribution V2"),
//				Description:         pulumi.String("Managed by Terraform"),
//				Url:                 pulumi.String("https://artifactory.my.com"),
//				AuthenticationToken: &azuredevops.ServiceendpointJfrogDistributionV2AuthenticationTokenArgs{
//					Token: pulumi.String("0000000000000000000000000000000000000000"),
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
// Alternatively a username and password may be used.
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
//			_, err = azuredevops.NewServiceendpointJfrogDistributionV2(ctx, "exampleServiceendpointJfrogDistributionV2", &azuredevops.ServiceendpointJfrogDistributionV2Args{
//				ProjectId:           exampleProject.ID(),
//				ServiceEndpointName: pulumi.String("Example JFrog Distribution V2"),
//				Description:         pulumi.String("Managed by Terraform"),
//				Url:                 pulumi.String("https://artifactory.my.com"),
//				AuthenticationBasic: &azuredevops.ServiceendpointJfrogDistributionV2AuthenticationBasicArgs{
//					Username: pulumi.String("username"),
//					Password: pulumi.String("password"),
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
// ## Relevant Links
//
// * [Azure DevOps Service Connections](https://docs.microsoft.com/en-us/azure/devops/pipelines/library/service-endpoints?view=azure-devops&tabs=yaml)
// * [Artifactory User Token](https://docs.artifactory.org/latest/user-guide/user-token/)
//
// ## Import
//
// Azure DevOps Service Endpoint JFrog Distribution V2 can be imported using the **projectID/serviceEndpointID**, e.g.
//
// ```sh
//
//	$ pulumi import azuredevops:index/serviceendpointJfrogDistributionV2:ServiceendpointJfrogDistributionV2 example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
//
// ```
type ServiceendpointJfrogDistributionV2 struct {
	pulumi.CustomResourceState

	// A `authenticationBasic` block as documented below.
	AuthenticationBasic ServiceendpointJfrogDistributionV2AuthenticationBasicPtrOutput `pulumi:"authenticationBasic"`
	// A `authenticationToken` block as documented below.
	AuthenticationToken ServiceendpointJfrogDistributionV2AuthenticationTokenPtrOutput `pulumi:"authenticationToken"`
	Authorization       pulumi.StringMapOutput                                         `pulumi:"authorization"`
	// The Service Endpoint description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The ID of the project.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The Service Endpoint name.
	ServiceEndpointName pulumi.StringOutput `pulumi:"serviceEndpointName"`
	// URL of the Artifactory server to connect with.
	//
	// > **NOTE:** URL should not end in a slash character.
	Url pulumi.StringOutput `pulumi:"url"`
}

// NewServiceendpointJfrogDistributionV2 registers a new resource with the given unique name, arguments, and options.
func NewServiceendpointJfrogDistributionV2(ctx *pulumi.Context,
	name string, args *ServiceendpointJfrogDistributionV2Args, opts ...pulumi.ResourceOption) (*ServiceendpointJfrogDistributionV2, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
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
	var resource ServiceendpointJfrogDistributionV2
	err := ctx.RegisterResource("azuredevops:index/serviceendpointJfrogDistributionV2:ServiceendpointJfrogDistributionV2", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceendpointJfrogDistributionV2 gets an existing ServiceendpointJfrogDistributionV2 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceendpointJfrogDistributionV2(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceendpointJfrogDistributionV2State, opts ...pulumi.ResourceOption) (*ServiceendpointJfrogDistributionV2, error) {
	var resource ServiceendpointJfrogDistributionV2
	err := ctx.ReadResource("azuredevops:index/serviceendpointJfrogDistributionV2:ServiceendpointJfrogDistributionV2", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceendpointJfrogDistributionV2 resources.
type serviceendpointJfrogDistributionV2State struct {
	// A `authenticationBasic` block as documented below.
	AuthenticationBasic *ServiceendpointJfrogDistributionV2AuthenticationBasic `pulumi:"authenticationBasic"`
	// A `authenticationToken` block as documented below.
	AuthenticationToken *ServiceendpointJfrogDistributionV2AuthenticationToken `pulumi:"authenticationToken"`
	Authorization       map[string]string                                      `pulumi:"authorization"`
	// The Service Endpoint description.
	Description *string `pulumi:"description"`
	// The ID of the project.
	ProjectId *string `pulumi:"projectId"`
	// The Service Endpoint name.
	ServiceEndpointName *string `pulumi:"serviceEndpointName"`
	// URL of the Artifactory server to connect with.
	//
	// > **NOTE:** URL should not end in a slash character.
	Url *string `pulumi:"url"`
}

type ServiceendpointJfrogDistributionV2State struct {
	// A `authenticationBasic` block as documented below.
	AuthenticationBasic ServiceendpointJfrogDistributionV2AuthenticationBasicPtrInput
	// A `authenticationToken` block as documented below.
	AuthenticationToken ServiceendpointJfrogDistributionV2AuthenticationTokenPtrInput
	Authorization       pulumi.StringMapInput
	// The Service Endpoint description.
	Description pulumi.StringPtrInput
	// The ID of the project.
	ProjectId pulumi.StringPtrInput
	// The Service Endpoint name.
	ServiceEndpointName pulumi.StringPtrInput
	// URL of the Artifactory server to connect with.
	//
	// > **NOTE:** URL should not end in a slash character.
	Url pulumi.StringPtrInput
}

func (ServiceendpointJfrogDistributionV2State) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceendpointJfrogDistributionV2State)(nil)).Elem()
}

type serviceendpointJfrogDistributionV2Args struct {
	// A `authenticationBasic` block as documented below.
	AuthenticationBasic *ServiceendpointJfrogDistributionV2AuthenticationBasic `pulumi:"authenticationBasic"`
	// A `authenticationToken` block as documented below.
	AuthenticationToken *ServiceendpointJfrogDistributionV2AuthenticationToken `pulumi:"authenticationToken"`
	Authorization       map[string]string                                      `pulumi:"authorization"`
	// The Service Endpoint description.
	Description *string `pulumi:"description"`
	// The ID of the project.
	ProjectId string `pulumi:"projectId"`
	// The Service Endpoint name.
	ServiceEndpointName string `pulumi:"serviceEndpointName"`
	// URL of the Artifactory server to connect with.
	//
	// > **NOTE:** URL should not end in a slash character.
	Url string `pulumi:"url"`
}

// The set of arguments for constructing a ServiceendpointJfrogDistributionV2 resource.
type ServiceendpointJfrogDistributionV2Args struct {
	// A `authenticationBasic` block as documented below.
	AuthenticationBasic ServiceendpointJfrogDistributionV2AuthenticationBasicPtrInput
	// A `authenticationToken` block as documented below.
	AuthenticationToken ServiceendpointJfrogDistributionV2AuthenticationTokenPtrInput
	Authorization       pulumi.StringMapInput
	// The Service Endpoint description.
	Description pulumi.StringPtrInput
	// The ID of the project.
	ProjectId pulumi.StringInput
	// The Service Endpoint name.
	ServiceEndpointName pulumi.StringInput
	// URL of the Artifactory server to connect with.
	//
	// > **NOTE:** URL should not end in a slash character.
	Url pulumi.StringInput
}

func (ServiceendpointJfrogDistributionV2Args) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceendpointJfrogDistributionV2Args)(nil)).Elem()
}

type ServiceendpointJfrogDistributionV2Input interface {
	pulumi.Input

	ToServiceendpointJfrogDistributionV2Output() ServiceendpointJfrogDistributionV2Output
	ToServiceendpointJfrogDistributionV2OutputWithContext(ctx context.Context) ServiceendpointJfrogDistributionV2Output
}

func (*ServiceendpointJfrogDistributionV2) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceendpointJfrogDistributionV2)(nil)).Elem()
}

func (i *ServiceendpointJfrogDistributionV2) ToServiceendpointJfrogDistributionV2Output() ServiceendpointJfrogDistributionV2Output {
	return i.ToServiceendpointJfrogDistributionV2OutputWithContext(context.Background())
}

func (i *ServiceendpointJfrogDistributionV2) ToServiceendpointJfrogDistributionV2OutputWithContext(ctx context.Context) ServiceendpointJfrogDistributionV2Output {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceendpointJfrogDistributionV2Output)
}

// ServiceendpointJfrogDistributionV2ArrayInput is an input type that accepts ServiceendpointJfrogDistributionV2Array and ServiceendpointJfrogDistributionV2ArrayOutput values.
// You can construct a concrete instance of `ServiceendpointJfrogDistributionV2ArrayInput` via:
//
//	ServiceendpointJfrogDistributionV2Array{ ServiceendpointJfrogDistributionV2Args{...} }
type ServiceendpointJfrogDistributionV2ArrayInput interface {
	pulumi.Input

	ToServiceendpointJfrogDistributionV2ArrayOutput() ServiceendpointJfrogDistributionV2ArrayOutput
	ToServiceendpointJfrogDistributionV2ArrayOutputWithContext(context.Context) ServiceendpointJfrogDistributionV2ArrayOutput
}

type ServiceendpointJfrogDistributionV2Array []ServiceendpointJfrogDistributionV2Input

func (ServiceendpointJfrogDistributionV2Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceendpointJfrogDistributionV2)(nil)).Elem()
}

func (i ServiceendpointJfrogDistributionV2Array) ToServiceendpointJfrogDistributionV2ArrayOutput() ServiceendpointJfrogDistributionV2ArrayOutput {
	return i.ToServiceendpointJfrogDistributionV2ArrayOutputWithContext(context.Background())
}

func (i ServiceendpointJfrogDistributionV2Array) ToServiceendpointJfrogDistributionV2ArrayOutputWithContext(ctx context.Context) ServiceendpointJfrogDistributionV2ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceendpointJfrogDistributionV2ArrayOutput)
}

// ServiceendpointJfrogDistributionV2MapInput is an input type that accepts ServiceendpointJfrogDistributionV2Map and ServiceendpointJfrogDistributionV2MapOutput values.
// You can construct a concrete instance of `ServiceendpointJfrogDistributionV2MapInput` via:
//
//	ServiceendpointJfrogDistributionV2Map{ "key": ServiceendpointJfrogDistributionV2Args{...} }
type ServiceendpointJfrogDistributionV2MapInput interface {
	pulumi.Input

	ToServiceendpointJfrogDistributionV2MapOutput() ServiceendpointJfrogDistributionV2MapOutput
	ToServiceendpointJfrogDistributionV2MapOutputWithContext(context.Context) ServiceendpointJfrogDistributionV2MapOutput
}

type ServiceendpointJfrogDistributionV2Map map[string]ServiceendpointJfrogDistributionV2Input

func (ServiceendpointJfrogDistributionV2Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceendpointJfrogDistributionV2)(nil)).Elem()
}

func (i ServiceendpointJfrogDistributionV2Map) ToServiceendpointJfrogDistributionV2MapOutput() ServiceendpointJfrogDistributionV2MapOutput {
	return i.ToServiceendpointJfrogDistributionV2MapOutputWithContext(context.Background())
}

func (i ServiceendpointJfrogDistributionV2Map) ToServiceendpointJfrogDistributionV2MapOutputWithContext(ctx context.Context) ServiceendpointJfrogDistributionV2MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceendpointJfrogDistributionV2MapOutput)
}

type ServiceendpointJfrogDistributionV2Output struct{ *pulumi.OutputState }

func (ServiceendpointJfrogDistributionV2Output) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceendpointJfrogDistributionV2)(nil)).Elem()
}

func (o ServiceendpointJfrogDistributionV2Output) ToServiceendpointJfrogDistributionV2Output() ServiceendpointJfrogDistributionV2Output {
	return o
}

func (o ServiceendpointJfrogDistributionV2Output) ToServiceendpointJfrogDistributionV2OutputWithContext(ctx context.Context) ServiceendpointJfrogDistributionV2Output {
	return o
}

// A `authenticationBasic` block as documented below.
func (o ServiceendpointJfrogDistributionV2Output) AuthenticationBasic() ServiceendpointJfrogDistributionV2AuthenticationBasicPtrOutput {
	return o.ApplyT(func(v *ServiceendpointJfrogDistributionV2) ServiceendpointJfrogDistributionV2AuthenticationBasicPtrOutput {
		return v.AuthenticationBasic
	}).(ServiceendpointJfrogDistributionV2AuthenticationBasicPtrOutput)
}

// A `authenticationToken` block as documented below.
func (o ServiceendpointJfrogDistributionV2Output) AuthenticationToken() ServiceendpointJfrogDistributionV2AuthenticationTokenPtrOutput {
	return o.ApplyT(func(v *ServiceendpointJfrogDistributionV2) ServiceendpointJfrogDistributionV2AuthenticationTokenPtrOutput {
		return v.AuthenticationToken
	}).(ServiceendpointJfrogDistributionV2AuthenticationTokenPtrOutput)
}

func (o ServiceendpointJfrogDistributionV2Output) Authorization() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ServiceendpointJfrogDistributionV2) pulumi.StringMapOutput { return v.Authorization }).(pulumi.StringMapOutput)
}

// The Service Endpoint description.
func (o ServiceendpointJfrogDistributionV2Output) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceendpointJfrogDistributionV2) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The ID of the project.
func (o ServiceendpointJfrogDistributionV2Output) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceendpointJfrogDistributionV2) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The Service Endpoint name.
func (o ServiceendpointJfrogDistributionV2Output) ServiceEndpointName() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceendpointJfrogDistributionV2) pulumi.StringOutput { return v.ServiceEndpointName }).(pulumi.StringOutput)
}

// URL of the Artifactory server to connect with.
//
// > **NOTE:** URL should not end in a slash character.
func (o ServiceendpointJfrogDistributionV2Output) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceendpointJfrogDistributionV2) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

type ServiceendpointJfrogDistributionV2ArrayOutput struct{ *pulumi.OutputState }

func (ServiceendpointJfrogDistributionV2ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceendpointJfrogDistributionV2)(nil)).Elem()
}

func (o ServiceendpointJfrogDistributionV2ArrayOutput) ToServiceendpointJfrogDistributionV2ArrayOutput() ServiceendpointJfrogDistributionV2ArrayOutput {
	return o
}

func (o ServiceendpointJfrogDistributionV2ArrayOutput) ToServiceendpointJfrogDistributionV2ArrayOutputWithContext(ctx context.Context) ServiceendpointJfrogDistributionV2ArrayOutput {
	return o
}

func (o ServiceendpointJfrogDistributionV2ArrayOutput) Index(i pulumi.IntInput) ServiceendpointJfrogDistributionV2Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServiceendpointJfrogDistributionV2 {
		return vs[0].([]*ServiceendpointJfrogDistributionV2)[vs[1].(int)]
	}).(ServiceendpointJfrogDistributionV2Output)
}

type ServiceendpointJfrogDistributionV2MapOutput struct{ *pulumi.OutputState }

func (ServiceendpointJfrogDistributionV2MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceendpointJfrogDistributionV2)(nil)).Elem()
}

func (o ServiceendpointJfrogDistributionV2MapOutput) ToServiceendpointJfrogDistributionV2MapOutput() ServiceendpointJfrogDistributionV2MapOutput {
	return o
}

func (o ServiceendpointJfrogDistributionV2MapOutput) ToServiceendpointJfrogDistributionV2MapOutputWithContext(ctx context.Context) ServiceendpointJfrogDistributionV2MapOutput {
	return o
}

func (o ServiceendpointJfrogDistributionV2MapOutput) MapIndex(k pulumi.StringInput) ServiceendpointJfrogDistributionV2Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServiceendpointJfrogDistributionV2 {
		return vs[0].(map[string]*ServiceendpointJfrogDistributionV2)[vs[1].(string)]
	}).(ServiceendpointJfrogDistributionV2Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceendpointJfrogDistributionV2Input)(nil)).Elem(), &ServiceendpointJfrogDistributionV2{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceendpointJfrogDistributionV2ArrayInput)(nil)).Elem(), ServiceendpointJfrogDistributionV2Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceendpointJfrogDistributionV2MapInput)(nil)).Elem(), ServiceendpointJfrogDistributionV2Map{})
	pulumi.RegisterOutputType(ServiceendpointJfrogDistributionV2Output{})
	pulumi.RegisterOutputType(ServiceendpointJfrogDistributionV2ArrayOutput{})
	pulumi.RegisterOutputType(ServiceendpointJfrogDistributionV2MapOutput{})
}
