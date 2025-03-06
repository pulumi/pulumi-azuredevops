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

// Manages a Visual Studio Marketplace service endpoint within Azure DevOps. Using this service endpoint requires you to install: [Azure DevOps Extension Tasks](https://marketplace.visualstudio.com/items?itemName=ms-devlabs.vsts-developer-tools-build-tasks)
//
// ## Example Usage
//
// ### Authorize with token
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
//			_, err = azuredevops.NewServiceendpointVisualstudiomarketplace(ctx, "example", &azuredevops.ServiceendpointVisualstudiomarketplaceArgs{
//				ProjectId:           example.ID(),
//				ServiceEndpointName: pulumi.String("Example Marketplace"),
//				Url:                 pulumi.String("https://markpetplace.com"),
//				AuthenticationToken: &azuredevops.ServiceendpointVisualstudiomarketplaceAuthenticationTokenArgs{
//					Token: pulumi.String("token"),
//				},
//				Description: pulumi.String("Managed by Pulumi"),
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
// ### Authorize with username and password
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
//			_, err = azuredevops.NewServiceendpointVisualstudiomarketplace(ctx, "example", &azuredevops.ServiceendpointVisualstudiomarketplaceArgs{
//				ProjectId:           example.ID(),
//				ServiceEndpointName: pulumi.String("Example Marketplace"),
//				Url:                 pulumi.String("https://markpetplace.com"),
//				AuthenticationBasic: &azuredevops.ServiceendpointVisualstudiomarketplaceAuthenticationBasicArgs{
//					Username: pulumi.String("username"),
//					Password: pulumi.String("password"),
//				},
//				Description: pulumi.String("Managed by Pulumi"),
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
// - [Azure DevOps Service REST API 7.0 - Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-7.0)
//
// ## Import
//
// Azure DevOps Visual Studio Marketplace Service Endpoint can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
//
// ```sh
// $ pulumi import azuredevops:index/serviceendpointVisualstudiomarketplace:ServiceendpointVisualstudiomarketplace example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
// ```
type ServiceendpointVisualstudiomarketplace struct {
	pulumi.CustomResourceState

	// An `authenticationBasic` block as documented below.
	//
	// > **NOTE:** `authenticationBasic` and `authenticationToken` conflict with each other, only one is required.
	AuthenticationBasic ServiceendpointVisualstudiomarketplaceAuthenticationBasicPtrOutput `pulumi:"authenticationBasic"`
	// An `authenticationToken` block as documented below.
	AuthenticationToken ServiceendpointVisualstudiomarketplaceAuthenticationTokenPtrOutput `pulumi:"authenticationToken"`
	Authorization       pulumi.StringMapOutput                                             `pulumi:"authorization"`
	Description         pulumi.StringPtrOutput                                             `pulumi:"description"`
	// The ID of the project.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The Service Endpoint name.
	ServiceEndpointName pulumi.StringOutput `pulumi:"serviceEndpointName"`
	// The server URL for Visual Studio Marketplace.
	Url pulumi.StringOutput `pulumi:"url"`
}

// NewServiceendpointVisualstudiomarketplace registers a new resource with the given unique name, arguments, and options.
func NewServiceendpointVisualstudiomarketplace(ctx *pulumi.Context,
	name string, args *ServiceendpointVisualstudiomarketplaceArgs, opts ...pulumi.ResourceOption) (*ServiceendpointVisualstudiomarketplace, error) {
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
	var resource ServiceendpointVisualstudiomarketplace
	err := ctx.RegisterResource("azuredevops:index/serviceendpointVisualstudiomarketplace:ServiceendpointVisualstudiomarketplace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceendpointVisualstudiomarketplace gets an existing ServiceendpointVisualstudiomarketplace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceendpointVisualstudiomarketplace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceendpointVisualstudiomarketplaceState, opts ...pulumi.ResourceOption) (*ServiceendpointVisualstudiomarketplace, error) {
	var resource ServiceendpointVisualstudiomarketplace
	err := ctx.ReadResource("azuredevops:index/serviceendpointVisualstudiomarketplace:ServiceendpointVisualstudiomarketplace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceendpointVisualstudiomarketplace resources.
type serviceendpointVisualstudiomarketplaceState struct {
	// An `authenticationBasic` block as documented below.
	//
	// > **NOTE:** `authenticationBasic` and `authenticationToken` conflict with each other, only one is required.
	AuthenticationBasic *ServiceendpointVisualstudiomarketplaceAuthenticationBasic `pulumi:"authenticationBasic"`
	// An `authenticationToken` block as documented below.
	AuthenticationToken *ServiceendpointVisualstudiomarketplaceAuthenticationToken `pulumi:"authenticationToken"`
	Authorization       map[string]string                                          `pulumi:"authorization"`
	Description         *string                                                    `pulumi:"description"`
	// The ID of the project.
	ProjectId *string `pulumi:"projectId"`
	// The Service Endpoint name.
	ServiceEndpointName *string `pulumi:"serviceEndpointName"`
	// The server URL for Visual Studio Marketplace.
	Url *string `pulumi:"url"`
}

type ServiceendpointVisualstudiomarketplaceState struct {
	// An `authenticationBasic` block as documented below.
	//
	// > **NOTE:** `authenticationBasic` and `authenticationToken` conflict with each other, only one is required.
	AuthenticationBasic ServiceendpointVisualstudiomarketplaceAuthenticationBasicPtrInput
	// An `authenticationToken` block as documented below.
	AuthenticationToken ServiceendpointVisualstudiomarketplaceAuthenticationTokenPtrInput
	Authorization       pulumi.StringMapInput
	Description         pulumi.StringPtrInput
	// The ID of the project.
	ProjectId pulumi.StringPtrInput
	// The Service Endpoint name.
	ServiceEndpointName pulumi.StringPtrInput
	// The server URL for Visual Studio Marketplace.
	Url pulumi.StringPtrInput
}

func (ServiceendpointVisualstudiomarketplaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceendpointVisualstudiomarketplaceState)(nil)).Elem()
}

type serviceendpointVisualstudiomarketplaceArgs struct {
	// An `authenticationBasic` block as documented below.
	//
	// > **NOTE:** `authenticationBasic` and `authenticationToken` conflict with each other, only one is required.
	AuthenticationBasic *ServiceendpointVisualstudiomarketplaceAuthenticationBasic `pulumi:"authenticationBasic"`
	// An `authenticationToken` block as documented below.
	AuthenticationToken *ServiceendpointVisualstudiomarketplaceAuthenticationToken `pulumi:"authenticationToken"`
	Description         *string                                                    `pulumi:"description"`
	// The ID of the project.
	ProjectId string `pulumi:"projectId"`
	// The Service Endpoint name.
	ServiceEndpointName string `pulumi:"serviceEndpointName"`
	// The server URL for Visual Studio Marketplace.
	Url string `pulumi:"url"`
}

// The set of arguments for constructing a ServiceendpointVisualstudiomarketplace resource.
type ServiceendpointVisualstudiomarketplaceArgs struct {
	// An `authenticationBasic` block as documented below.
	//
	// > **NOTE:** `authenticationBasic` and `authenticationToken` conflict with each other, only one is required.
	AuthenticationBasic ServiceendpointVisualstudiomarketplaceAuthenticationBasicPtrInput
	// An `authenticationToken` block as documented below.
	AuthenticationToken ServiceendpointVisualstudiomarketplaceAuthenticationTokenPtrInput
	Description         pulumi.StringPtrInput
	// The ID of the project.
	ProjectId pulumi.StringInput
	// The Service Endpoint name.
	ServiceEndpointName pulumi.StringInput
	// The server URL for Visual Studio Marketplace.
	Url pulumi.StringInput
}

func (ServiceendpointVisualstudiomarketplaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceendpointVisualstudiomarketplaceArgs)(nil)).Elem()
}

type ServiceendpointVisualstudiomarketplaceInput interface {
	pulumi.Input

	ToServiceendpointVisualstudiomarketplaceOutput() ServiceendpointVisualstudiomarketplaceOutput
	ToServiceendpointVisualstudiomarketplaceOutputWithContext(ctx context.Context) ServiceendpointVisualstudiomarketplaceOutput
}

func (*ServiceendpointVisualstudiomarketplace) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceendpointVisualstudiomarketplace)(nil)).Elem()
}

func (i *ServiceendpointVisualstudiomarketplace) ToServiceendpointVisualstudiomarketplaceOutput() ServiceendpointVisualstudiomarketplaceOutput {
	return i.ToServiceendpointVisualstudiomarketplaceOutputWithContext(context.Background())
}

func (i *ServiceendpointVisualstudiomarketplace) ToServiceendpointVisualstudiomarketplaceOutputWithContext(ctx context.Context) ServiceendpointVisualstudiomarketplaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceendpointVisualstudiomarketplaceOutput)
}

// ServiceendpointVisualstudiomarketplaceArrayInput is an input type that accepts ServiceendpointVisualstudiomarketplaceArray and ServiceendpointVisualstudiomarketplaceArrayOutput values.
// You can construct a concrete instance of `ServiceendpointVisualstudiomarketplaceArrayInput` via:
//
//	ServiceendpointVisualstudiomarketplaceArray{ ServiceendpointVisualstudiomarketplaceArgs{...} }
type ServiceendpointVisualstudiomarketplaceArrayInput interface {
	pulumi.Input

	ToServiceendpointVisualstudiomarketplaceArrayOutput() ServiceendpointVisualstudiomarketplaceArrayOutput
	ToServiceendpointVisualstudiomarketplaceArrayOutputWithContext(context.Context) ServiceendpointVisualstudiomarketplaceArrayOutput
}

type ServiceendpointVisualstudiomarketplaceArray []ServiceendpointVisualstudiomarketplaceInput

func (ServiceendpointVisualstudiomarketplaceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceendpointVisualstudiomarketplace)(nil)).Elem()
}

func (i ServiceendpointVisualstudiomarketplaceArray) ToServiceendpointVisualstudiomarketplaceArrayOutput() ServiceendpointVisualstudiomarketplaceArrayOutput {
	return i.ToServiceendpointVisualstudiomarketplaceArrayOutputWithContext(context.Background())
}

func (i ServiceendpointVisualstudiomarketplaceArray) ToServiceendpointVisualstudiomarketplaceArrayOutputWithContext(ctx context.Context) ServiceendpointVisualstudiomarketplaceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceendpointVisualstudiomarketplaceArrayOutput)
}

// ServiceendpointVisualstudiomarketplaceMapInput is an input type that accepts ServiceendpointVisualstudiomarketplaceMap and ServiceendpointVisualstudiomarketplaceMapOutput values.
// You can construct a concrete instance of `ServiceendpointVisualstudiomarketplaceMapInput` via:
//
//	ServiceendpointVisualstudiomarketplaceMap{ "key": ServiceendpointVisualstudiomarketplaceArgs{...} }
type ServiceendpointVisualstudiomarketplaceMapInput interface {
	pulumi.Input

	ToServiceendpointVisualstudiomarketplaceMapOutput() ServiceendpointVisualstudiomarketplaceMapOutput
	ToServiceendpointVisualstudiomarketplaceMapOutputWithContext(context.Context) ServiceendpointVisualstudiomarketplaceMapOutput
}

type ServiceendpointVisualstudiomarketplaceMap map[string]ServiceendpointVisualstudiomarketplaceInput

func (ServiceendpointVisualstudiomarketplaceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceendpointVisualstudiomarketplace)(nil)).Elem()
}

func (i ServiceendpointVisualstudiomarketplaceMap) ToServiceendpointVisualstudiomarketplaceMapOutput() ServiceendpointVisualstudiomarketplaceMapOutput {
	return i.ToServiceendpointVisualstudiomarketplaceMapOutputWithContext(context.Background())
}

func (i ServiceendpointVisualstudiomarketplaceMap) ToServiceendpointVisualstudiomarketplaceMapOutputWithContext(ctx context.Context) ServiceendpointVisualstudiomarketplaceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceendpointVisualstudiomarketplaceMapOutput)
}

type ServiceendpointVisualstudiomarketplaceOutput struct{ *pulumi.OutputState }

func (ServiceendpointVisualstudiomarketplaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceendpointVisualstudiomarketplace)(nil)).Elem()
}

func (o ServiceendpointVisualstudiomarketplaceOutput) ToServiceendpointVisualstudiomarketplaceOutput() ServiceendpointVisualstudiomarketplaceOutput {
	return o
}

func (o ServiceendpointVisualstudiomarketplaceOutput) ToServiceendpointVisualstudiomarketplaceOutputWithContext(ctx context.Context) ServiceendpointVisualstudiomarketplaceOutput {
	return o
}

// An `authenticationBasic` block as documented below.
//
// > **NOTE:** `authenticationBasic` and `authenticationToken` conflict with each other, only one is required.
func (o ServiceendpointVisualstudiomarketplaceOutput) AuthenticationBasic() ServiceendpointVisualstudiomarketplaceAuthenticationBasicPtrOutput {
	return o.ApplyT(func(v *ServiceendpointVisualstudiomarketplace) ServiceendpointVisualstudiomarketplaceAuthenticationBasicPtrOutput {
		return v.AuthenticationBasic
	}).(ServiceendpointVisualstudiomarketplaceAuthenticationBasicPtrOutput)
}

// An `authenticationToken` block as documented below.
func (o ServiceendpointVisualstudiomarketplaceOutput) AuthenticationToken() ServiceendpointVisualstudiomarketplaceAuthenticationTokenPtrOutput {
	return o.ApplyT(func(v *ServiceendpointVisualstudiomarketplace) ServiceendpointVisualstudiomarketplaceAuthenticationTokenPtrOutput {
		return v.AuthenticationToken
	}).(ServiceendpointVisualstudiomarketplaceAuthenticationTokenPtrOutput)
}

func (o ServiceendpointVisualstudiomarketplaceOutput) Authorization() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ServiceendpointVisualstudiomarketplace) pulumi.StringMapOutput { return v.Authorization }).(pulumi.StringMapOutput)
}

func (o ServiceendpointVisualstudiomarketplaceOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceendpointVisualstudiomarketplace) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The ID of the project.
func (o ServiceendpointVisualstudiomarketplaceOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceendpointVisualstudiomarketplace) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The Service Endpoint name.
func (o ServiceendpointVisualstudiomarketplaceOutput) ServiceEndpointName() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceendpointVisualstudiomarketplace) pulumi.StringOutput { return v.ServiceEndpointName }).(pulumi.StringOutput)
}

// The server URL for Visual Studio Marketplace.
func (o ServiceendpointVisualstudiomarketplaceOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceendpointVisualstudiomarketplace) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

type ServiceendpointVisualstudiomarketplaceArrayOutput struct{ *pulumi.OutputState }

func (ServiceendpointVisualstudiomarketplaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceendpointVisualstudiomarketplace)(nil)).Elem()
}

func (o ServiceendpointVisualstudiomarketplaceArrayOutput) ToServiceendpointVisualstudiomarketplaceArrayOutput() ServiceendpointVisualstudiomarketplaceArrayOutput {
	return o
}

func (o ServiceendpointVisualstudiomarketplaceArrayOutput) ToServiceendpointVisualstudiomarketplaceArrayOutputWithContext(ctx context.Context) ServiceendpointVisualstudiomarketplaceArrayOutput {
	return o
}

func (o ServiceendpointVisualstudiomarketplaceArrayOutput) Index(i pulumi.IntInput) ServiceendpointVisualstudiomarketplaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServiceendpointVisualstudiomarketplace {
		return vs[0].([]*ServiceendpointVisualstudiomarketplace)[vs[1].(int)]
	}).(ServiceendpointVisualstudiomarketplaceOutput)
}

type ServiceendpointVisualstudiomarketplaceMapOutput struct{ *pulumi.OutputState }

func (ServiceendpointVisualstudiomarketplaceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceendpointVisualstudiomarketplace)(nil)).Elem()
}

func (o ServiceendpointVisualstudiomarketplaceMapOutput) ToServiceendpointVisualstudiomarketplaceMapOutput() ServiceendpointVisualstudiomarketplaceMapOutput {
	return o
}

func (o ServiceendpointVisualstudiomarketplaceMapOutput) ToServiceendpointVisualstudiomarketplaceMapOutputWithContext(ctx context.Context) ServiceendpointVisualstudiomarketplaceMapOutput {
	return o
}

func (o ServiceendpointVisualstudiomarketplaceMapOutput) MapIndex(k pulumi.StringInput) ServiceendpointVisualstudiomarketplaceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServiceendpointVisualstudiomarketplace {
		return vs[0].(map[string]*ServiceendpointVisualstudiomarketplace)[vs[1].(string)]
	}).(ServiceendpointVisualstudiomarketplaceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceendpointVisualstudiomarketplaceInput)(nil)).Elem(), &ServiceendpointVisualstudiomarketplace{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceendpointVisualstudiomarketplaceArrayInput)(nil)).Elem(), ServiceendpointVisualstudiomarketplaceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceendpointVisualstudiomarketplaceMapInput)(nil)).Elem(), ServiceendpointVisualstudiomarketplaceMap{})
	pulumi.RegisterOutputType(ServiceendpointVisualstudiomarketplaceOutput{})
	pulumi.RegisterOutputType(ServiceendpointVisualstudiomarketplaceArrayOutput{})
	pulumi.RegisterOutputType(ServiceendpointVisualstudiomarketplaceMapOutput{})
}
