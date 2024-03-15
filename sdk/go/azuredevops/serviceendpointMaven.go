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

// Manages a Maven service endpoint within Azure DevOps, which can be used as a resource in YAML pipelines to connect to a Maven instance.
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
//			_, err = azuredevops.NewServiceendpointMaven(ctx, "exampleServiceendpointMaven", &azuredevops.ServiceendpointMavenArgs{
//				ProjectId:           exampleProject.ID(),
//				ServiceEndpointName: pulumi.String("maven-example"),
//				Description:         pulumi.String("Service Endpoint for 'Maven' (Managed by Terraform)"),
//				Url:                 pulumi.String("https://example.com"),
//				RepositoryId:        pulumi.String("example"),
//				AuthenticationToken: &azuredevops.ServiceendpointMavenAuthenticationTokenArgs{
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
//
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
//			exampleProject, err := azuredevops.NewProject(ctx, "exampleProject", &azuredevops.ProjectArgs{
//				Visibility:       pulumi.String("private"),
//				VersionControl:   pulumi.String("Git"),
//				WorkItemTemplate: pulumi.String("Agile"),
//				Description:      pulumi.String("Managed by Terraform"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuredevops.NewServiceendpointMaven(ctx, "exampleServiceendpointMaven", &azuredevops.ServiceendpointMavenArgs{
//				ProjectId:           exampleProject.ID(),
//				ServiceEndpointName: pulumi.String("maven-example"),
//				Description:         pulumi.String("Service Endpoint for 'Maven' (Managed by Terraform)"),
//				Url:                 pulumi.String("https://example.com"),
//				RepositoryId:        pulumi.String("example"),
//				AuthenticationBasic: &azuredevops.ServiceendpointMavenAuthenticationBasicArgs{
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
// ## Import
//
// Service Connection Maven can be imported using the `projectId/id` or or `projectName/id`, e.g.
//
// ```sh
// $ pulumi import azuredevops:index/serviceendpointMaven:ServiceendpointMaven example projectName/00000000-0000-0000-0000-000000000000
// ```
type ServiceendpointMaven struct {
	pulumi.CustomResourceState

	// A `authenticationBasic` block as documented below.
	AuthenticationBasic ServiceendpointMavenAuthenticationBasicPtrOutput `pulumi:"authenticationBasic"`
	// A `authenticationToken` block as documented below.
	AuthenticationToken ServiceendpointMavenAuthenticationTokenPtrOutput `pulumi:"authenticationToken"`
	Authorization       pulumi.StringMapOutput                           `pulumi:"authorization"`
	Description         pulumi.StringPtrOutput                           `pulumi:"description"`
	// The ID of the project. Changing this forces a new Service Connection Maven to be created.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The ID of the server that matches the id element of the `repository/mirror` that Maven tries to connect to.
	RepositoryId pulumi.StringOutput `pulumi:"repositoryId"`
	// The name of the service endpoint. Changing this forces a new Service Connection Maven to be created.
	ServiceEndpointName pulumi.StringOutput `pulumi:"serviceEndpointName"`
	// The URL of the Maven Repository.
	Url pulumi.StringOutput `pulumi:"url"`
}

// NewServiceendpointMaven registers a new resource with the given unique name, arguments, and options.
func NewServiceendpointMaven(ctx *pulumi.Context,
	name string, args *ServiceendpointMavenArgs, opts ...pulumi.ResourceOption) (*ServiceendpointMaven, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.RepositoryId == nil {
		return nil, errors.New("invalid value for required argument 'RepositoryId'")
	}
	if args.ServiceEndpointName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceEndpointName'")
	}
	if args.Url == nil {
		return nil, errors.New("invalid value for required argument 'Url'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ServiceendpointMaven
	err := ctx.RegisterResource("azuredevops:index/serviceendpointMaven:ServiceendpointMaven", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceendpointMaven gets an existing ServiceendpointMaven resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceendpointMaven(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceendpointMavenState, opts ...pulumi.ResourceOption) (*ServiceendpointMaven, error) {
	var resource ServiceendpointMaven
	err := ctx.ReadResource("azuredevops:index/serviceendpointMaven:ServiceendpointMaven", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceendpointMaven resources.
type serviceendpointMavenState struct {
	// A `authenticationBasic` block as documented below.
	AuthenticationBasic *ServiceendpointMavenAuthenticationBasic `pulumi:"authenticationBasic"`
	// A `authenticationToken` block as documented below.
	AuthenticationToken *ServiceendpointMavenAuthenticationToken `pulumi:"authenticationToken"`
	Authorization       map[string]string                        `pulumi:"authorization"`
	Description         *string                                  `pulumi:"description"`
	// The ID of the project. Changing this forces a new Service Connection Maven to be created.
	ProjectId *string `pulumi:"projectId"`
	// The ID of the server that matches the id element of the `repository/mirror` that Maven tries to connect to.
	RepositoryId *string `pulumi:"repositoryId"`
	// The name of the service endpoint. Changing this forces a new Service Connection Maven to be created.
	ServiceEndpointName *string `pulumi:"serviceEndpointName"`
	// The URL of the Maven Repository.
	Url *string `pulumi:"url"`
}

type ServiceendpointMavenState struct {
	// A `authenticationBasic` block as documented below.
	AuthenticationBasic ServiceendpointMavenAuthenticationBasicPtrInput
	// A `authenticationToken` block as documented below.
	AuthenticationToken ServiceendpointMavenAuthenticationTokenPtrInput
	Authorization       pulumi.StringMapInput
	Description         pulumi.StringPtrInput
	// The ID of the project. Changing this forces a new Service Connection Maven to be created.
	ProjectId pulumi.StringPtrInput
	// The ID of the server that matches the id element of the `repository/mirror` that Maven tries to connect to.
	RepositoryId pulumi.StringPtrInput
	// The name of the service endpoint. Changing this forces a new Service Connection Maven to be created.
	ServiceEndpointName pulumi.StringPtrInput
	// The URL of the Maven Repository.
	Url pulumi.StringPtrInput
}

func (ServiceendpointMavenState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceendpointMavenState)(nil)).Elem()
}

type serviceendpointMavenArgs struct {
	// A `authenticationBasic` block as documented below.
	AuthenticationBasic *ServiceendpointMavenAuthenticationBasic `pulumi:"authenticationBasic"`
	// A `authenticationToken` block as documented below.
	AuthenticationToken *ServiceendpointMavenAuthenticationToken `pulumi:"authenticationToken"`
	Authorization       map[string]string                        `pulumi:"authorization"`
	Description         *string                                  `pulumi:"description"`
	// The ID of the project. Changing this forces a new Service Connection Maven to be created.
	ProjectId string `pulumi:"projectId"`
	// The ID of the server that matches the id element of the `repository/mirror` that Maven tries to connect to.
	RepositoryId string `pulumi:"repositoryId"`
	// The name of the service endpoint. Changing this forces a new Service Connection Maven to be created.
	ServiceEndpointName string `pulumi:"serviceEndpointName"`
	// The URL of the Maven Repository.
	Url string `pulumi:"url"`
}

// The set of arguments for constructing a ServiceendpointMaven resource.
type ServiceendpointMavenArgs struct {
	// A `authenticationBasic` block as documented below.
	AuthenticationBasic ServiceendpointMavenAuthenticationBasicPtrInput
	// A `authenticationToken` block as documented below.
	AuthenticationToken ServiceendpointMavenAuthenticationTokenPtrInput
	Authorization       pulumi.StringMapInput
	Description         pulumi.StringPtrInput
	// The ID of the project. Changing this forces a new Service Connection Maven to be created.
	ProjectId pulumi.StringInput
	// The ID of the server that matches the id element of the `repository/mirror` that Maven tries to connect to.
	RepositoryId pulumi.StringInput
	// The name of the service endpoint. Changing this forces a new Service Connection Maven to be created.
	ServiceEndpointName pulumi.StringInput
	// The URL of the Maven Repository.
	Url pulumi.StringInput
}

func (ServiceendpointMavenArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceendpointMavenArgs)(nil)).Elem()
}

type ServiceendpointMavenInput interface {
	pulumi.Input

	ToServiceendpointMavenOutput() ServiceendpointMavenOutput
	ToServiceendpointMavenOutputWithContext(ctx context.Context) ServiceendpointMavenOutput
}

func (*ServiceendpointMaven) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceendpointMaven)(nil)).Elem()
}

func (i *ServiceendpointMaven) ToServiceendpointMavenOutput() ServiceendpointMavenOutput {
	return i.ToServiceendpointMavenOutputWithContext(context.Background())
}

func (i *ServiceendpointMaven) ToServiceendpointMavenOutputWithContext(ctx context.Context) ServiceendpointMavenOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceendpointMavenOutput)
}

// ServiceendpointMavenArrayInput is an input type that accepts ServiceendpointMavenArray and ServiceendpointMavenArrayOutput values.
// You can construct a concrete instance of `ServiceendpointMavenArrayInput` via:
//
//	ServiceendpointMavenArray{ ServiceendpointMavenArgs{...} }
type ServiceendpointMavenArrayInput interface {
	pulumi.Input

	ToServiceendpointMavenArrayOutput() ServiceendpointMavenArrayOutput
	ToServiceendpointMavenArrayOutputWithContext(context.Context) ServiceendpointMavenArrayOutput
}

type ServiceendpointMavenArray []ServiceendpointMavenInput

func (ServiceendpointMavenArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceendpointMaven)(nil)).Elem()
}

func (i ServiceendpointMavenArray) ToServiceendpointMavenArrayOutput() ServiceendpointMavenArrayOutput {
	return i.ToServiceendpointMavenArrayOutputWithContext(context.Background())
}

func (i ServiceendpointMavenArray) ToServiceendpointMavenArrayOutputWithContext(ctx context.Context) ServiceendpointMavenArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceendpointMavenArrayOutput)
}

// ServiceendpointMavenMapInput is an input type that accepts ServiceendpointMavenMap and ServiceendpointMavenMapOutput values.
// You can construct a concrete instance of `ServiceendpointMavenMapInput` via:
//
//	ServiceendpointMavenMap{ "key": ServiceendpointMavenArgs{...} }
type ServiceendpointMavenMapInput interface {
	pulumi.Input

	ToServiceendpointMavenMapOutput() ServiceendpointMavenMapOutput
	ToServiceendpointMavenMapOutputWithContext(context.Context) ServiceendpointMavenMapOutput
}

type ServiceendpointMavenMap map[string]ServiceendpointMavenInput

func (ServiceendpointMavenMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceendpointMaven)(nil)).Elem()
}

func (i ServiceendpointMavenMap) ToServiceendpointMavenMapOutput() ServiceendpointMavenMapOutput {
	return i.ToServiceendpointMavenMapOutputWithContext(context.Background())
}

func (i ServiceendpointMavenMap) ToServiceendpointMavenMapOutputWithContext(ctx context.Context) ServiceendpointMavenMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceendpointMavenMapOutput)
}

type ServiceendpointMavenOutput struct{ *pulumi.OutputState }

func (ServiceendpointMavenOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceendpointMaven)(nil)).Elem()
}

func (o ServiceendpointMavenOutput) ToServiceendpointMavenOutput() ServiceendpointMavenOutput {
	return o
}

func (o ServiceendpointMavenOutput) ToServiceendpointMavenOutputWithContext(ctx context.Context) ServiceendpointMavenOutput {
	return o
}

// A `authenticationBasic` block as documented below.
func (o ServiceendpointMavenOutput) AuthenticationBasic() ServiceendpointMavenAuthenticationBasicPtrOutput {
	return o.ApplyT(func(v *ServiceendpointMaven) ServiceendpointMavenAuthenticationBasicPtrOutput {
		return v.AuthenticationBasic
	}).(ServiceendpointMavenAuthenticationBasicPtrOutput)
}

// A `authenticationToken` block as documented below.
func (o ServiceendpointMavenOutput) AuthenticationToken() ServiceendpointMavenAuthenticationTokenPtrOutput {
	return o.ApplyT(func(v *ServiceendpointMaven) ServiceendpointMavenAuthenticationTokenPtrOutput {
		return v.AuthenticationToken
	}).(ServiceendpointMavenAuthenticationTokenPtrOutput)
}

func (o ServiceendpointMavenOutput) Authorization() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ServiceendpointMaven) pulumi.StringMapOutput { return v.Authorization }).(pulumi.StringMapOutput)
}

func (o ServiceendpointMavenOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceendpointMaven) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The ID of the project. Changing this forces a new Service Connection Maven to be created.
func (o ServiceendpointMavenOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceendpointMaven) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The ID of the server that matches the id element of the `repository/mirror` that Maven tries to connect to.
func (o ServiceendpointMavenOutput) RepositoryId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceendpointMaven) pulumi.StringOutput { return v.RepositoryId }).(pulumi.StringOutput)
}

// The name of the service endpoint. Changing this forces a new Service Connection Maven to be created.
func (o ServiceendpointMavenOutput) ServiceEndpointName() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceendpointMaven) pulumi.StringOutput { return v.ServiceEndpointName }).(pulumi.StringOutput)
}

// The URL of the Maven Repository.
func (o ServiceendpointMavenOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceendpointMaven) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

type ServiceendpointMavenArrayOutput struct{ *pulumi.OutputState }

func (ServiceendpointMavenArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceendpointMaven)(nil)).Elem()
}

func (o ServiceendpointMavenArrayOutput) ToServiceendpointMavenArrayOutput() ServiceendpointMavenArrayOutput {
	return o
}

func (o ServiceendpointMavenArrayOutput) ToServiceendpointMavenArrayOutputWithContext(ctx context.Context) ServiceendpointMavenArrayOutput {
	return o
}

func (o ServiceendpointMavenArrayOutput) Index(i pulumi.IntInput) ServiceendpointMavenOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServiceendpointMaven {
		return vs[0].([]*ServiceendpointMaven)[vs[1].(int)]
	}).(ServiceendpointMavenOutput)
}

type ServiceendpointMavenMapOutput struct{ *pulumi.OutputState }

func (ServiceendpointMavenMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceendpointMaven)(nil)).Elem()
}

func (o ServiceendpointMavenMapOutput) ToServiceendpointMavenMapOutput() ServiceendpointMavenMapOutput {
	return o
}

func (o ServiceendpointMavenMapOutput) ToServiceendpointMavenMapOutputWithContext(ctx context.Context) ServiceendpointMavenMapOutput {
	return o
}

func (o ServiceendpointMavenMapOutput) MapIndex(k pulumi.StringInput) ServiceendpointMavenOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServiceendpointMaven {
		return vs[0].(map[string]*ServiceendpointMaven)[vs[1].(string)]
	}).(ServiceendpointMavenOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceendpointMavenInput)(nil)).Elem(), &ServiceendpointMaven{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceendpointMavenArrayInput)(nil)).Elem(), ServiceendpointMavenArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceendpointMavenMapInput)(nil)).Elem(), ServiceendpointMavenMap{})
	pulumi.RegisterOutputType(ServiceendpointMavenOutput{})
	pulumi.RegisterOutputType(ServiceendpointMavenArrayOutput{})
	pulumi.RegisterOutputType(ServiceendpointMavenMapOutput{})
}
