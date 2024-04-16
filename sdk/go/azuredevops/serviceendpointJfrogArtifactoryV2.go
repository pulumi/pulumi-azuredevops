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

// Manages a JFrog Artifactory V2 server endpoint within an Azure DevOps organization.
//
// > **Note:** Using this service endpoint requires you to first install [JFrog Extension](https://marketplace.visualstudio.com/items?itemName=JFrog.jfrog-azure-devops-extension).
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
//			_, err = azuredevops.NewServiceendpointJfrogArtifactoryV2(ctx, "example", &azuredevops.ServiceendpointJfrogArtifactoryV2Args{
//				ProjectId:           example.ID(),
//				ServiceEndpointName: pulumi.String("Example JFrog Artifactory V2"),
//				Description:         pulumi.String("Managed by Terraform"),
//				Url:                 pulumi.String("https://artifactory.my.com"),
//				AuthenticationToken: &azuredevops.ServiceendpointJfrogArtifactoryV2AuthenticationTokenArgs{
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
// <!--End PulumiCodeChooser -->
// Alternatively a username and password may be used.
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
//			_, err = azuredevops.NewServiceendpointJfrogArtifactoryV2(ctx, "example", &azuredevops.ServiceendpointJfrogArtifactoryV2Args{
//				ProjectId:           example.ID(),
//				ServiceEndpointName: pulumi.String("Example JFrog Artifactory V2"),
//				Description:         pulumi.String("Managed by Terraform"),
//				Url:                 pulumi.String("https://artifactory.my.com"),
//				AuthenticationBasic: &azuredevops.ServiceendpointJfrogArtifactoryV2AuthenticationBasicArgs{
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
// <!--End PulumiCodeChooser -->
//
// ## Relevant Links
//
// * [Azure DevOps Service Connections](https://docs.microsoft.com/en-us/azure/devops/pipelines/library/service-endpoints?view=azure-devops&tabs=yaml)
// * [Artifactory User Token](https://docs.artifactory.org/latest/user-guide/user-token/)
//
// ## Import
//
// Azure DevOps Service Endpoint JFrog Artifactory V2 can be imported using the **projectID/serviceEndpointID**, e.g.
//
// ```sh
// $ pulumi import azuredevops:index/serviceendpointJfrogArtifactoryV2:ServiceendpointJfrogArtifactoryV2 example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
// ```
type ServiceendpointJfrogArtifactoryV2 struct {
	pulumi.CustomResourceState

	// A `authenticationBasic` block as documented below.
	AuthenticationBasic ServiceendpointJfrogArtifactoryV2AuthenticationBasicPtrOutput `pulumi:"authenticationBasic"`
	// A `authenticationToken` block as documented below.
	AuthenticationToken ServiceendpointJfrogArtifactoryV2AuthenticationTokenPtrOutput `pulumi:"authenticationToken"`
	Authorization       pulumi.StringMapOutput                                        `pulumi:"authorization"`
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

// NewServiceendpointJfrogArtifactoryV2 registers a new resource with the given unique name, arguments, and options.
func NewServiceendpointJfrogArtifactoryV2(ctx *pulumi.Context,
	name string, args *ServiceendpointJfrogArtifactoryV2Args, opts ...pulumi.ResourceOption) (*ServiceendpointJfrogArtifactoryV2, error) {
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
	var resource ServiceendpointJfrogArtifactoryV2
	err := ctx.RegisterResource("azuredevops:index/serviceendpointJfrogArtifactoryV2:ServiceendpointJfrogArtifactoryV2", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceendpointJfrogArtifactoryV2 gets an existing ServiceendpointJfrogArtifactoryV2 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceendpointJfrogArtifactoryV2(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceendpointJfrogArtifactoryV2State, opts ...pulumi.ResourceOption) (*ServiceendpointJfrogArtifactoryV2, error) {
	var resource ServiceendpointJfrogArtifactoryV2
	err := ctx.ReadResource("azuredevops:index/serviceendpointJfrogArtifactoryV2:ServiceendpointJfrogArtifactoryV2", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceendpointJfrogArtifactoryV2 resources.
type serviceendpointJfrogArtifactoryV2State struct {
	// A `authenticationBasic` block as documented below.
	AuthenticationBasic *ServiceendpointJfrogArtifactoryV2AuthenticationBasic `pulumi:"authenticationBasic"`
	// A `authenticationToken` block as documented below.
	AuthenticationToken *ServiceendpointJfrogArtifactoryV2AuthenticationToken `pulumi:"authenticationToken"`
	Authorization       map[string]string                                     `pulumi:"authorization"`
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

type ServiceendpointJfrogArtifactoryV2State struct {
	// A `authenticationBasic` block as documented below.
	AuthenticationBasic ServiceendpointJfrogArtifactoryV2AuthenticationBasicPtrInput
	// A `authenticationToken` block as documented below.
	AuthenticationToken ServiceendpointJfrogArtifactoryV2AuthenticationTokenPtrInput
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

func (ServiceendpointJfrogArtifactoryV2State) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceendpointJfrogArtifactoryV2State)(nil)).Elem()
}

type serviceendpointJfrogArtifactoryV2Args struct {
	// A `authenticationBasic` block as documented below.
	AuthenticationBasic *ServiceendpointJfrogArtifactoryV2AuthenticationBasic `pulumi:"authenticationBasic"`
	// A `authenticationToken` block as documented below.
	AuthenticationToken *ServiceendpointJfrogArtifactoryV2AuthenticationToken `pulumi:"authenticationToken"`
	Authorization       map[string]string                                     `pulumi:"authorization"`
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

// The set of arguments for constructing a ServiceendpointJfrogArtifactoryV2 resource.
type ServiceendpointJfrogArtifactoryV2Args struct {
	// A `authenticationBasic` block as documented below.
	AuthenticationBasic ServiceendpointJfrogArtifactoryV2AuthenticationBasicPtrInput
	// A `authenticationToken` block as documented below.
	AuthenticationToken ServiceendpointJfrogArtifactoryV2AuthenticationTokenPtrInput
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

func (ServiceendpointJfrogArtifactoryV2Args) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceendpointJfrogArtifactoryV2Args)(nil)).Elem()
}

type ServiceendpointJfrogArtifactoryV2Input interface {
	pulumi.Input

	ToServiceendpointJfrogArtifactoryV2Output() ServiceendpointJfrogArtifactoryV2Output
	ToServiceendpointJfrogArtifactoryV2OutputWithContext(ctx context.Context) ServiceendpointJfrogArtifactoryV2Output
}

func (*ServiceendpointJfrogArtifactoryV2) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceendpointJfrogArtifactoryV2)(nil)).Elem()
}

func (i *ServiceendpointJfrogArtifactoryV2) ToServiceendpointJfrogArtifactoryV2Output() ServiceendpointJfrogArtifactoryV2Output {
	return i.ToServiceendpointJfrogArtifactoryV2OutputWithContext(context.Background())
}

func (i *ServiceendpointJfrogArtifactoryV2) ToServiceendpointJfrogArtifactoryV2OutputWithContext(ctx context.Context) ServiceendpointJfrogArtifactoryV2Output {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceendpointJfrogArtifactoryV2Output)
}

// ServiceendpointJfrogArtifactoryV2ArrayInput is an input type that accepts ServiceendpointJfrogArtifactoryV2Array and ServiceendpointJfrogArtifactoryV2ArrayOutput values.
// You can construct a concrete instance of `ServiceendpointJfrogArtifactoryV2ArrayInput` via:
//
//	ServiceendpointJfrogArtifactoryV2Array{ ServiceendpointJfrogArtifactoryV2Args{...} }
type ServiceendpointJfrogArtifactoryV2ArrayInput interface {
	pulumi.Input

	ToServiceendpointJfrogArtifactoryV2ArrayOutput() ServiceendpointJfrogArtifactoryV2ArrayOutput
	ToServiceendpointJfrogArtifactoryV2ArrayOutputWithContext(context.Context) ServiceendpointJfrogArtifactoryV2ArrayOutput
}

type ServiceendpointJfrogArtifactoryV2Array []ServiceendpointJfrogArtifactoryV2Input

func (ServiceendpointJfrogArtifactoryV2Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceendpointJfrogArtifactoryV2)(nil)).Elem()
}

func (i ServiceendpointJfrogArtifactoryV2Array) ToServiceendpointJfrogArtifactoryV2ArrayOutput() ServiceendpointJfrogArtifactoryV2ArrayOutput {
	return i.ToServiceendpointJfrogArtifactoryV2ArrayOutputWithContext(context.Background())
}

func (i ServiceendpointJfrogArtifactoryV2Array) ToServiceendpointJfrogArtifactoryV2ArrayOutputWithContext(ctx context.Context) ServiceendpointJfrogArtifactoryV2ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceendpointJfrogArtifactoryV2ArrayOutput)
}

// ServiceendpointJfrogArtifactoryV2MapInput is an input type that accepts ServiceendpointJfrogArtifactoryV2Map and ServiceendpointJfrogArtifactoryV2MapOutput values.
// You can construct a concrete instance of `ServiceendpointJfrogArtifactoryV2MapInput` via:
//
//	ServiceendpointJfrogArtifactoryV2Map{ "key": ServiceendpointJfrogArtifactoryV2Args{...} }
type ServiceendpointJfrogArtifactoryV2MapInput interface {
	pulumi.Input

	ToServiceendpointJfrogArtifactoryV2MapOutput() ServiceendpointJfrogArtifactoryV2MapOutput
	ToServiceendpointJfrogArtifactoryV2MapOutputWithContext(context.Context) ServiceendpointJfrogArtifactoryV2MapOutput
}

type ServiceendpointJfrogArtifactoryV2Map map[string]ServiceendpointJfrogArtifactoryV2Input

func (ServiceendpointJfrogArtifactoryV2Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceendpointJfrogArtifactoryV2)(nil)).Elem()
}

func (i ServiceendpointJfrogArtifactoryV2Map) ToServiceendpointJfrogArtifactoryV2MapOutput() ServiceendpointJfrogArtifactoryV2MapOutput {
	return i.ToServiceendpointJfrogArtifactoryV2MapOutputWithContext(context.Background())
}

func (i ServiceendpointJfrogArtifactoryV2Map) ToServiceendpointJfrogArtifactoryV2MapOutputWithContext(ctx context.Context) ServiceendpointJfrogArtifactoryV2MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceendpointJfrogArtifactoryV2MapOutput)
}

type ServiceendpointJfrogArtifactoryV2Output struct{ *pulumi.OutputState }

func (ServiceendpointJfrogArtifactoryV2Output) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceendpointJfrogArtifactoryV2)(nil)).Elem()
}

func (o ServiceendpointJfrogArtifactoryV2Output) ToServiceendpointJfrogArtifactoryV2Output() ServiceendpointJfrogArtifactoryV2Output {
	return o
}

func (o ServiceendpointJfrogArtifactoryV2Output) ToServiceendpointJfrogArtifactoryV2OutputWithContext(ctx context.Context) ServiceendpointJfrogArtifactoryV2Output {
	return o
}

// A `authenticationBasic` block as documented below.
func (o ServiceendpointJfrogArtifactoryV2Output) AuthenticationBasic() ServiceendpointJfrogArtifactoryV2AuthenticationBasicPtrOutput {
	return o.ApplyT(func(v *ServiceendpointJfrogArtifactoryV2) ServiceendpointJfrogArtifactoryV2AuthenticationBasicPtrOutput {
		return v.AuthenticationBasic
	}).(ServiceendpointJfrogArtifactoryV2AuthenticationBasicPtrOutput)
}

// A `authenticationToken` block as documented below.
func (o ServiceendpointJfrogArtifactoryV2Output) AuthenticationToken() ServiceendpointJfrogArtifactoryV2AuthenticationTokenPtrOutput {
	return o.ApplyT(func(v *ServiceendpointJfrogArtifactoryV2) ServiceendpointJfrogArtifactoryV2AuthenticationTokenPtrOutput {
		return v.AuthenticationToken
	}).(ServiceendpointJfrogArtifactoryV2AuthenticationTokenPtrOutput)
}

func (o ServiceendpointJfrogArtifactoryV2Output) Authorization() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ServiceendpointJfrogArtifactoryV2) pulumi.StringMapOutput { return v.Authorization }).(pulumi.StringMapOutput)
}

// The Service Endpoint description.
func (o ServiceendpointJfrogArtifactoryV2Output) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceendpointJfrogArtifactoryV2) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The ID of the project.
func (o ServiceendpointJfrogArtifactoryV2Output) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceendpointJfrogArtifactoryV2) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The Service Endpoint name.
func (o ServiceendpointJfrogArtifactoryV2Output) ServiceEndpointName() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceendpointJfrogArtifactoryV2) pulumi.StringOutput { return v.ServiceEndpointName }).(pulumi.StringOutput)
}

// URL of the Artifactory server to connect with.
//
// > **NOTE:** URL should not end in a slash character.
func (o ServiceendpointJfrogArtifactoryV2Output) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceendpointJfrogArtifactoryV2) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

type ServiceendpointJfrogArtifactoryV2ArrayOutput struct{ *pulumi.OutputState }

func (ServiceendpointJfrogArtifactoryV2ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceendpointJfrogArtifactoryV2)(nil)).Elem()
}

func (o ServiceendpointJfrogArtifactoryV2ArrayOutput) ToServiceendpointJfrogArtifactoryV2ArrayOutput() ServiceendpointJfrogArtifactoryV2ArrayOutput {
	return o
}

func (o ServiceendpointJfrogArtifactoryV2ArrayOutput) ToServiceendpointJfrogArtifactoryV2ArrayOutputWithContext(ctx context.Context) ServiceendpointJfrogArtifactoryV2ArrayOutput {
	return o
}

func (o ServiceendpointJfrogArtifactoryV2ArrayOutput) Index(i pulumi.IntInput) ServiceendpointJfrogArtifactoryV2Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServiceendpointJfrogArtifactoryV2 {
		return vs[0].([]*ServiceendpointJfrogArtifactoryV2)[vs[1].(int)]
	}).(ServiceendpointJfrogArtifactoryV2Output)
}

type ServiceendpointJfrogArtifactoryV2MapOutput struct{ *pulumi.OutputState }

func (ServiceendpointJfrogArtifactoryV2MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceendpointJfrogArtifactoryV2)(nil)).Elem()
}

func (o ServiceendpointJfrogArtifactoryV2MapOutput) ToServiceendpointJfrogArtifactoryV2MapOutput() ServiceendpointJfrogArtifactoryV2MapOutput {
	return o
}

func (o ServiceendpointJfrogArtifactoryV2MapOutput) ToServiceendpointJfrogArtifactoryV2MapOutputWithContext(ctx context.Context) ServiceendpointJfrogArtifactoryV2MapOutput {
	return o
}

func (o ServiceendpointJfrogArtifactoryV2MapOutput) MapIndex(k pulumi.StringInput) ServiceendpointJfrogArtifactoryV2Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServiceendpointJfrogArtifactoryV2 {
		return vs[0].(map[string]*ServiceendpointJfrogArtifactoryV2)[vs[1].(string)]
	}).(ServiceendpointJfrogArtifactoryV2Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceendpointJfrogArtifactoryV2Input)(nil)).Elem(), &ServiceendpointJfrogArtifactoryV2{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceendpointJfrogArtifactoryV2ArrayInput)(nil)).Elem(), ServiceendpointJfrogArtifactoryV2Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceendpointJfrogArtifactoryV2MapInput)(nil)).Elem(), ServiceendpointJfrogArtifactoryV2Map{})
	pulumi.RegisterOutputType(ServiceendpointJfrogArtifactoryV2Output{})
	pulumi.RegisterOutputType(ServiceendpointJfrogArtifactoryV2ArrayOutput{})
	pulumi.RegisterOutputType(ServiceendpointJfrogArtifactoryV2MapOutput{})
}
