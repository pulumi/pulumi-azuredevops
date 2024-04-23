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

// Manages a JFrog Platform V2 server endpoint within an Azure DevOps organization.
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
//				Description:      pulumi.String("Managed by Terraform"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuredevops.NewServiceendpointJfrogPlatformV2(ctx, "example", &azuredevops.ServiceendpointJfrogPlatformV2Args{
//				ProjectId:           example.ID(),
//				ServiceEndpointName: pulumi.String("Example Artifactory"),
//				Description:         pulumi.String("Managed by Terraform"),
//				Url:                 pulumi.String("https://artifactory.my.com"),
//				AuthenticationToken: &azuredevops.ServiceendpointJfrogPlatformV2AuthenticationTokenArgs{
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
//				Description:      pulumi.String("Managed by Terraform"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuredevops.NewServiceendpointJfrogPlatformV2(ctx, "example", &azuredevops.ServiceendpointJfrogPlatformV2Args{
//				ProjectId:           example.ID(),
//				ServiceEndpointName: pulumi.String("Example Artifactory"),
//				Description:         pulumi.String("Managed by Terraform"),
//				Url:                 pulumi.String("https://artifactory.my.com"),
//				AuthenticationBasic: &azuredevops.ServiceendpointJfrogPlatformV2AuthenticationBasicArgs{
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
//
// ## Relevant Links
//
// * [Azure DevOps Service Connections](https://docs.microsoft.com/en-us/azure/devops/pipelines/library/service-endpoints?view=azure-devops&tabs=yaml)
// * [Artifactory User Token](https://docs.artifactory.org/latest/user-guide/user-token/)
//
// ## Import
//
// Azure DevOps Service Endpoint JFrog Platform V2 can be imported using the **projectID/serviceEndpointID**, e.g.
//
// ```sh
// $ pulumi import azuredevops:index/serviceendpointJfrogPlatformV2:ServiceendpointJfrogPlatformV2 example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
// ```
type ServiceendpointJfrogPlatformV2 struct {
	pulumi.CustomResourceState

	// A `authenticationBasic` block as documented below.
	AuthenticationBasic ServiceendpointJfrogPlatformV2AuthenticationBasicPtrOutput `pulumi:"authenticationBasic"`
	// A `authenticationToken` block as documented below.
	AuthenticationToken ServiceendpointJfrogPlatformV2AuthenticationTokenPtrOutput `pulumi:"authenticationToken"`
	Authorization       pulumi.StringMapOutput                                     `pulumi:"authorization"`
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

// NewServiceendpointJfrogPlatformV2 registers a new resource with the given unique name, arguments, and options.
func NewServiceendpointJfrogPlatformV2(ctx *pulumi.Context,
	name string, args *ServiceendpointJfrogPlatformV2Args, opts ...pulumi.ResourceOption) (*ServiceendpointJfrogPlatformV2, error) {
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
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ServiceendpointJfrogPlatformV2
	err := ctx.RegisterResource("azuredevops:index/serviceendpointJfrogPlatformV2:ServiceendpointJfrogPlatformV2", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceendpointJfrogPlatformV2 gets an existing ServiceendpointJfrogPlatformV2 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceendpointJfrogPlatformV2(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceendpointJfrogPlatformV2State, opts ...pulumi.ResourceOption) (*ServiceendpointJfrogPlatformV2, error) {
	var resource ServiceendpointJfrogPlatformV2
	err := ctx.ReadResource("azuredevops:index/serviceendpointJfrogPlatformV2:ServiceendpointJfrogPlatformV2", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceendpointJfrogPlatformV2 resources.
type serviceendpointJfrogPlatformV2State struct {
	// A `authenticationBasic` block as documented below.
	AuthenticationBasic *ServiceendpointJfrogPlatformV2AuthenticationBasic `pulumi:"authenticationBasic"`
	// A `authenticationToken` block as documented below.
	AuthenticationToken *ServiceendpointJfrogPlatformV2AuthenticationToken `pulumi:"authenticationToken"`
	Authorization       map[string]string                                  `pulumi:"authorization"`
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

type ServiceendpointJfrogPlatformV2State struct {
	// A `authenticationBasic` block as documented below.
	AuthenticationBasic ServiceendpointJfrogPlatformV2AuthenticationBasicPtrInput
	// A `authenticationToken` block as documented below.
	AuthenticationToken ServiceendpointJfrogPlatformV2AuthenticationTokenPtrInput
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

func (ServiceendpointJfrogPlatformV2State) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceendpointJfrogPlatformV2State)(nil)).Elem()
}

type serviceendpointJfrogPlatformV2Args struct {
	// A `authenticationBasic` block as documented below.
	AuthenticationBasic *ServiceendpointJfrogPlatformV2AuthenticationBasic `pulumi:"authenticationBasic"`
	// A `authenticationToken` block as documented below.
	AuthenticationToken *ServiceendpointJfrogPlatformV2AuthenticationToken `pulumi:"authenticationToken"`
	Authorization       map[string]string                                  `pulumi:"authorization"`
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

// The set of arguments for constructing a ServiceendpointJfrogPlatformV2 resource.
type ServiceendpointJfrogPlatformV2Args struct {
	// A `authenticationBasic` block as documented below.
	AuthenticationBasic ServiceendpointJfrogPlatformV2AuthenticationBasicPtrInput
	// A `authenticationToken` block as documented below.
	AuthenticationToken ServiceendpointJfrogPlatformV2AuthenticationTokenPtrInput
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

func (ServiceendpointJfrogPlatformV2Args) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceendpointJfrogPlatformV2Args)(nil)).Elem()
}

type ServiceendpointJfrogPlatformV2Input interface {
	pulumi.Input

	ToServiceendpointJfrogPlatformV2Output() ServiceendpointJfrogPlatformV2Output
	ToServiceendpointJfrogPlatformV2OutputWithContext(ctx context.Context) ServiceendpointJfrogPlatformV2Output
}

func (*ServiceendpointJfrogPlatformV2) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceendpointJfrogPlatformV2)(nil)).Elem()
}

func (i *ServiceendpointJfrogPlatformV2) ToServiceendpointJfrogPlatformV2Output() ServiceendpointJfrogPlatformV2Output {
	return i.ToServiceendpointJfrogPlatformV2OutputWithContext(context.Background())
}

func (i *ServiceendpointJfrogPlatformV2) ToServiceendpointJfrogPlatformV2OutputWithContext(ctx context.Context) ServiceendpointJfrogPlatformV2Output {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceendpointJfrogPlatformV2Output)
}

// ServiceendpointJfrogPlatformV2ArrayInput is an input type that accepts ServiceendpointJfrogPlatformV2Array and ServiceendpointJfrogPlatformV2ArrayOutput values.
// You can construct a concrete instance of `ServiceendpointJfrogPlatformV2ArrayInput` via:
//
//	ServiceendpointJfrogPlatformV2Array{ ServiceendpointJfrogPlatformV2Args{...} }
type ServiceendpointJfrogPlatformV2ArrayInput interface {
	pulumi.Input

	ToServiceendpointJfrogPlatformV2ArrayOutput() ServiceendpointJfrogPlatformV2ArrayOutput
	ToServiceendpointJfrogPlatformV2ArrayOutputWithContext(context.Context) ServiceendpointJfrogPlatformV2ArrayOutput
}

type ServiceendpointJfrogPlatformV2Array []ServiceendpointJfrogPlatformV2Input

func (ServiceendpointJfrogPlatformV2Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceendpointJfrogPlatformV2)(nil)).Elem()
}

func (i ServiceendpointJfrogPlatformV2Array) ToServiceendpointJfrogPlatformV2ArrayOutput() ServiceendpointJfrogPlatformV2ArrayOutput {
	return i.ToServiceendpointJfrogPlatformV2ArrayOutputWithContext(context.Background())
}

func (i ServiceendpointJfrogPlatformV2Array) ToServiceendpointJfrogPlatformV2ArrayOutputWithContext(ctx context.Context) ServiceendpointJfrogPlatformV2ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceendpointJfrogPlatformV2ArrayOutput)
}

// ServiceendpointJfrogPlatformV2MapInput is an input type that accepts ServiceendpointJfrogPlatformV2Map and ServiceendpointJfrogPlatformV2MapOutput values.
// You can construct a concrete instance of `ServiceendpointJfrogPlatformV2MapInput` via:
//
//	ServiceendpointJfrogPlatformV2Map{ "key": ServiceendpointJfrogPlatformV2Args{...} }
type ServiceendpointJfrogPlatformV2MapInput interface {
	pulumi.Input

	ToServiceendpointJfrogPlatformV2MapOutput() ServiceendpointJfrogPlatformV2MapOutput
	ToServiceendpointJfrogPlatformV2MapOutputWithContext(context.Context) ServiceendpointJfrogPlatformV2MapOutput
}

type ServiceendpointJfrogPlatformV2Map map[string]ServiceendpointJfrogPlatformV2Input

func (ServiceendpointJfrogPlatformV2Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceendpointJfrogPlatformV2)(nil)).Elem()
}

func (i ServiceendpointJfrogPlatformV2Map) ToServiceendpointJfrogPlatformV2MapOutput() ServiceendpointJfrogPlatformV2MapOutput {
	return i.ToServiceendpointJfrogPlatformV2MapOutputWithContext(context.Background())
}

func (i ServiceendpointJfrogPlatformV2Map) ToServiceendpointJfrogPlatformV2MapOutputWithContext(ctx context.Context) ServiceendpointJfrogPlatformV2MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceendpointJfrogPlatformV2MapOutput)
}

type ServiceendpointJfrogPlatformV2Output struct{ *pulumi.OutputState }

func (ServiceendpointJfrogPlatformV2Output) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceendpointJfrogPlatformV2)(nil)).Elem()
}

func (o ServiceendpointJfrogPlatformV2Output) ToServiceendpointJfrogPlatformV2Output() ServiceendpointJfrogPlatformV2Output {
	return o
}

func (o ServiceendpointJfrogPlatformV2Output) ToServiceendpointJfrogPlatformV2OutputWithContext(ctx context.Context) ServiceendpointJfrogPlatformV2Output {
	return o
}

// A `authenticationBasic` block as documented below.
func (o ServiceendpointJfrogPlatformV2Output) AuthenticationBasic() ServiceendpointJfrogPlatformV2AuthenticationBasicPtrOutput {
	return o.ApplyT(func(v *ServiceendpointJfrogPlatformV2) ServiceendpointJfrogPlatformV2AuthenticationBasicPtrOutput {
		return v.AuthenticationBasic
	}).(ServiceendpointJfrogPlatformV2AuthenticationBasicPtrOutput)
}

// A `authenticationToken` block as documented below.
func (o ServiceendpointJfrogPlatformV2Output) AuthenticationToken() ServiceendpointJfrogPlatformV2AuthenticationTokenPtrOutput {
	return o.ApplyT(func(v *ServiceendpointJfrogPlatformV2) ServiceendpointJfrogPlatformV2AuthenticationTokenPtrOutput {
		return v.AuthenticationToken
	}).(ServiceendpointJfrogPlatformV2AuthenticationTokenPtrOutput)
}

func (o ServiceendpointJfrogPlatformV2Output) Authorization() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ServiceendpointJfrogPlatformV2) pulumi.StringMapOutput { return v.Authorization }).(pulumi.StringMapOutput)
}

// The Service Endpoint description.
func (o ServiceendpointJfrogPlatformV2Output) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceendpointJfrogPlatformV2) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The ID of the project.
func (o ServiceendpointJfrogPlatformV2Output) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceendpointJfrogPlatformV2) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The Service Endpoint name.
func (o ServiceendpointJfrogPlatformV2Output) ServiceEndpointName() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceendpointJfrogPlatformV2) pulumi.StringOutput { return v.ServiceEndpointName }).(pulumi.StringOutput)
}

// URL of the Artifactory server to connect with.
//
// > **NOTE:** URL should not end in a slash character.
func (o ServiceendpointJfrogPlatformV2Output) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceendpointJfrogPlatformV2) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

type ServiceendpointJfrogPlatformV2ArrayOutput struct{ *pulumi.OutputState }

func (ServiceendpointJfrogPlatformV2ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceendpointJfrogPlatformV2)(nil)).Elem()
}

func (o ServiceendpointJfrogPlatformV2ArrayOutput) ToServiceendpointJfrogPlatformV2ArrayOutput() ServiceendpointJfrogPlatformV2ArrayOutput {
	return o
}

func (o ServiceendpointJfrogPlatformV2ArrayOutput) ToServiceendpointJfrogPlatformV2ArrayOutputWithContext(ctx context.Context) ServiceendpointJfrogPlatformV2ArrayOutput {
	return o
}

func (o ServiceendpointJfrogPlatformV2ArrayOutput) Index(i pulumi.IntInput) ServiceendpointJfrogPlatformV2Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServiceendpointJfrogPlatformV2 {
		return vs[0].([]*ServiceendpointJfrogPlatformV2)[vs[1].(int)]
	}).(ServiceendpointJfrogPlatformV2Output)
}

type ServiceendpointJfrogPlatformV2MapOutput struct{ *pulumi.OutputState }

func (ServiceendpointJfrogPlatformV2MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceendpointJfrogPlatformV2)(nil)).Elem()
}

func (o ServiceendpointJfrogPlatformV2MapOutput) ToServiceendpointJfrogPlatformV2MapOutput() ServiceendpointJfrogPlatformV2MapOutput {
	return o
}

func (o ServiceendpointJfrogPlatformV2MapOutput) ToServiceendpointJfrogPlatformV2MapOutputWithContext(ctx context.Context) ServiceendpointJfrogPlatformV2MapOutput {
	return o
}

func (o ServiceendpointJfrogPlatformV2MapOutput) MapIndex(k pulumi.StringInput) ServiceendpointJfrogPlatformV2Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServiceendpointJfrogPlatformV2 {
		return vs[0].(map[string]*ServiceendpointJfrogPlatformV2)[vs[1].(string)]
	}).(ServiceendpointJfrogPlatformV2Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceendpointJfrogPlatformV2Input)(nil)).Elem(), &ServiceendpointJfrogPlatformV2{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceendpointJfrogPlatformV2ArrayInput)(nil)).Elem(), ServiceendpointJfrogPlatformV2Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceendpointJfrogPlatformV2MapInput)(nil)).Elem(), ServiceendpointJfrogPlatformV2Map{})
	pulumi.RegisterOutputType(ServiceendpointJfrogPlatformV2Output{})
	pulumi.RegisterOutputType(ServiceendpointJfrogPlatformV2ArrayOutput{})
	pulumi.RegisterOutputType(ServiceendpointJfrogPlatformV2MapOutput{})
}
