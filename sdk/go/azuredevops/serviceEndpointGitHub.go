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

// Manages a GitHub service endpoint within Azure DevOps.
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
//			_, err = azuredevops.NewServiceEndpointGitHub(ctx, "exampleServiceEndpointGitHub", &azuredevops.ServiceEndpointGitHubArgs{
//				ProjectId:           exampleProject.ID(),
//				ServiceEndpointName: pulumi.String("Example GitHub Personal Access Token"),
//				AuthPersonal: &azuredevops.ServiceEndpointGitHubAuthPersonalArgs{
//					PersonalAccessToken: pulumi.String("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
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
//			_, err = azuredevops.NewServiceEndpointGitHub(ctx, "exampleServiceEndpointGitHub", &azuredevops.ServiceEndpointGitHubArgs{
//				ProjectId:           exampleProject.ID(),
//				ServiceEndpointName: pulumi.String("Example GitHub"),
//				AuthOauth: &azuredevops.ServiceEndpointGitHubAuthOauthArgs{
//					OauthConfigurationId: pulumi.String("00000000-0000-0000-0000-000000000000"),
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
//			_, err = azuredevops.NewServiceEndpointGitHub(ctx, "exampleServiceEndpointGitHub", &azuredevops.ServiceEndpointGitHubArgs{
//				ProjectId:           exampleProject.ID(),
//				ServiceEndpointName: pulumi.String("Example GitHub Apps: Azure Pipelines"),
//				Description:         pulumi.String("Managed by Terraform"),
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
// - [Azure DevOps Service REST API 7.0 - Service Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-7.0)
//
// ## Import
//
// Azure DevOps Service Endpoint GitHub can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
//
// ```sh
//
//	$ pulumi import azuredevops:index/serviceEndpointGitHub:ServiceEndpointGitHub example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
//
// ```
type ServiceEndpointGitHub struct {
	pulumi.CustomResourceState

	AuthOauth ServiceEndpointGitHubAuthOauthPtrOutput `pulumi:"authOauth"`
	// An `authPersonal` block as documented below. Allows connecting using a personal access token.
	AuthPersonal  ServiceEndpointGitHubAuthPersonalPtrOutput `pulumi:"authPersonal"`
	Authorization pulumi.StringMapOutput                     `pulumi:"authorization"`
	Description   pulumi.StringPtrOutput                     `pulumi:"description"`
	// The ID of the project.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The Service Endpoint name.
	ServiceEndpointName pulumi.StringOutput `pulumi:"serviceEndpointName"`
}

// NewServiceEndpointGitHub registers a new resource with the given unique name, arguments, and options.
func NewServiceEndpointGitHub(ctx *pulumi.Context,
	name string, args *ServiceEndpointGitHubArgs, opts ...pulumi.ResourceOption) (*ServiceEndpointGitHub, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.ServiceEndpointName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceEndpointName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azuredevops:ServiceEndpoint/gitHub:GitHub"),
		},
	})
	opts = append(opts, aliases)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ServiceEndpointGitHub
	err := ctx.RegisterResource("azuredevops:index/serviceEndpointGitHub:ServiceEndpointGitHub", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceEndpointGitHub gets an existing ServiceEndpointGitHub resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceEndpointGitHub(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceEndpointGitHubState, opts ...pulumi.ResourceOption) (*ServiceEndpointGitHub, error) {
	var resource ServiceEndpointGitHub
	err := ctx.ReadResource("azuredevops:index/serviceEndpointGitHub:ServiceEndpointGitHub", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceEndpointGitHub resources.
type serviceEndpointGitHubState struct {
	AuthOauth *ServiceEndpointGitHubAuthOauth `pulumi:"authOauth"`
	// An `authPersonal` block as documented below. Allows connecting using a personal access token.
	AuthPersonal  *ServiceEndpointGitHubAuthPersonal `pulumi:"authPersonal"`
	Authorization map[string]string                  `pulumi:"authorization"`
	Description   *string                            `pulumi:"description"`
	// The ID of the project.
	ProjectId *string `pulumi:"projectId"`
	// The Service Endpoint name.
	ServiceEndpointName *string `pulumi:"serviceEndpointName"`
}

type ServiceEndpointGitHubState struct {
	AuthOauth ServiceEndpointGitHubAuthOauthPtrInput
	// An `authPersonal` block as documented below. Allows connecting using a personal access token.
	AuthPersonal  ServiceEndpointGitHubAuthPersonalPtrInput
	Authorization pulumi.StringMapInput
	Description   pulumi.StringPtrInput
	// The ID of the project.
	ProjectId pulumi.StringPtrInput
	// The Service Endpoint name.
	ServiceEndpointName pulumi.StringPtrInput
}

func (ServiceEndpointGitHubState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceEndpointGitHubState)(nil)).Elem()
}

type serviceEndpointGitHubArgs struct {
	AuthOauth *ServiceEndpointGitHubAuthOauth `pulumi:"authOauth"`
	// An `authPersonal` block as documented below. Allows connecting using a personal access token.
	AuthPersonal  *ServiceEndpointGitHubAuthPersonal `pulumi:"authPersonal"`
	Authorization map[string]string                  `pulumi:"authorization"`
	Description   *string                            `pulumi:"description"`
	// The ID of the project.
	ProjectId string `pulumi:"projectId"`
	// The Service Endpoint name.
	ServiceEndpointName string `pulumi:"serviceEndpointName"`
}

// The set of arguments for constructing a ServiceEndpointGitHub resource.
type ServiceEndpointGitHubArgs struct {
	AuthOauth ServiceEndpointGitHubAuthOauthPtrInput
	// An `authPersonal` block as documented below. Allows connecting using a personal access token.
	AuthPersonal  ServiceEndpointGitHubAuthPersonalPtrInput
	Authorization pulumi.StringMapInput
	Description   pulumi.StringPtrInput
	// The ID of the project.
	ProjectId pulumi.StringInput
	// The Service Endpoint name.
	ServiceEndpointName pulumi.StringInput
}

func (ServiceEndpointGitHubArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceEndpointGitHubArgs)(nil)).Elem()
}

type ServiceEndpointGitHubInput interface {
	pulumi.Input

	ToServiceEndpointGitHubOutput() ServiceEndpointGitHubOutput
	ToServiceEndpointGitHubOutputWithContext(ctx context.Context) ServiceEndpointGitHubOutput
}

func (*ServiceEndpointGitHub) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceEndpointGitHub)(nil)).Elem()
}

func (i *ServiceEndpointGitHub) ToServiceEndpointGitHubOutput() ServiceEndpointGitHubOutput {
	return i.ToServiceEndpointGitHubOutputWithContext(context.Background())
}

func (i *ServiceEndpointGitHub) ToServiceEndpointGitHubOutputWithContext(ctx context.Context) ServiceEndpointGitHubOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointGitHubOutput)
}

// ServiceEndpointGitHubArrayInput is an input type that accepts ServiceEndpointGitHubArray and ServiceEndpointGitHubArrayOutput values.
// You can construct a concrete instance of `ServiceEndpointGitHubArrayInput` via:
//
//	ServiceEndpointGitHubArray{ ServiceEndpointGitHubArgs{...} }
type ServiceEndpointGitHubArrayInput interface {
	pulumi.Input

	ToServiceEndpointGitHubArrayOutput() ServiceEndpointGitHubArrayOutput
	ToServiceEndpointGitHubArrayOutputWithContext(context.Context) ServiceEndpointGitHubArrayOutput
}

type ServiceEndpointGitHubArray []ServiceEndpointGitHubInput

func (ServiceEndpointGitHubArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceEndpointGitHub)(nil)).Elem()
}

func (i ServiceEndpointGitHubArray) ToServiceEndpointGitHubArrayOutput() ServiceEndpointGitHubArrayOutput {
	return i.ToServiceEndpointGitHubArrayOutputWithContext(context.Background())
}

func (i ServiceEndpointGitHubArray) ToServiceEndpointGitHubArrayOutputWithContext(ctx context.Context) ServiceEndpointGitHubArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointGitHubArrayOutput)
}

// ServiceEndpointGitHubMapInput is an input type that accepts ServiceEndpointGitHubMap and ServiceEndpointGitHubMapOutput values.
// You can construct a concrete instance of `ServiceEndpointGitHubMapInput` via:
//
//	ServiceEndpointGitHubMap{ "key": ServiceEndpointGitHubArgs{...} }
type ServiceEndpointGitHubMapInput interface {
	pulumi.Input

	ToServiceEndpointGitHubMapOutput() ServiceEndpointGitHubMapOutput
	ToServiceEndpointGitHubMapOutputWithContext(context.Context) ServiceEndpointGitHubMapOutput
}

type ServiceEndpointGitHubMap map[string]ServiceEndpointGitHubInput

func (ServiceEndpointGitHubMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceEndpointGitHub)(nil)).Elem()
}

func (i ServiceEndpointGitHubMap) ToServiceEndpointGitHubMapOutput() ServiceEndpointGitHubMapOutput {
	return i.ToServiceEndpointGitHubMapOutputWithContext(context.Background())
}

func (i ServiceEndpointGitHubMap) ToServiceEndpointGitHubMapOutputWithContext(ctx context.Context) ServiceEndpointGitHubMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointGitHubMapOutput)
}

type ServiceEndpointGitHubOutput struct{ *pulumi.OutputState }

func (ServiceEndpointGitHubOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceEndpointGitHub)(nil)).Elem()
}

func (o ServiceEndpointGitHubOutput) ToServiceEndpointGitHubOutput() ServiceEndpointGitHubOutput {
	return o
}

func (o ServiceEndpointGitHubOutput) ToServiceEndpointGitHubOutputWithContext(ctx context.Context) ServiceEndpointGitHubOutput {
	return o
}

func (o ServiceEndpointGitHubOutput) AuthOauth() ServiceEndpointGitHubAuthOauthPtrOutput {
	return o.ApplyT(func(v *ServiceEndpointGitHub) ServiceEndpointGitHubAuthOauthPtrOutput { return v.AuthOauth }).(ServiceEndpointGitHubAuthOauthPtrOutput)
}

// An `authPersonal` block as documented below. Allows connecting using a personal access token.
func (o ServiceEndpointGitHubOutput) AuthPersonal() ServiceEndpointGitHubAuthPersonalPtrOutput {
	return o.ApplyT(func(v *ServiceEndpointGitHub) ServiceEndpointGitHubAuthPersonalPtrOutput { return v.AuthPersonal }).(ServiceEndpointGitHubAuthPersonalPtrOutput)
}

func (o ServiceEndpointGitHubOutput) Authorization() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ServiceEndpointGitHub) pulumi.StringMapOutput { return v.Authorization }).(pulumi.StringMapOutput)
}

func (o ServiceEndpointGitHubOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceEndpointGitHub) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The ID of the project.
func (o ServiceEndpointGitHubOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEndpointGitHub) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The Service Endpoint name.
func (o ServiceEndpointGitHubOutput) ServiceEndpointName() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEndpointGitHub) pulumi.StringOutput { return v.ServiceEndpointName }).(pulumi.StringOutput)
}

type ServiceEndpointGitHubArrayOutput struct{ *pulumi.OutputState }

func (ServiceEndpointGitHubArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceEndpointGitHub)(nil)).Elem()
}

func (o ServiceEndpointGitHubArrayOutput) ToServiceEndpointGitHubArrayOutput() ServiceEndpointGitHubArrayOutput {
	return o
}

func (o ServiceEndpointGitHubArrayOutput) ToServiceEndpointGitHubArrayOutputWithContext(ctx context.Context) ServiceEndpointGitHubArrayOutput {
	return o
}

func (o ServiceEndpointGitHubArrayOutput) Index(i pulumi.IntInput) ServiceEndpointGitHubOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServiceEndpointGitHub {
		return vs[0].([]*ServiceEndpointGitHub)[vs[1].(int)]
	}).(ServiceEndpointGitHubOutput)
}

type ServiceEndpointGitHubMapOutput struct{ *pulumi.OutputState }

func (ServiceEndpointGitHubMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceEndpointGitHub)(nil)).Elem()
}

func (o ServiceEndpointGitHubMapOutput) ToServiceEndpointGitHubMapOutput() ServiceEndpointGitHubMapOutput {
	return o
}

func (o ServiceEndpointGitHubMapOutput) ToServiceEndpointGitHubMapOutputWithContext(ctx context.Context) ServiceEndpointGitHubMapOutput {
	return o
}

func (o ServiceEndpointGitHubMapOutput) MapIndex(k pulumi.StringInput) ServiceEndpointGitHubOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServiceEndpointGitHub {
		return vs[0].(map[string]*ServiceEndpointGitHub)[vs[1].(string)]
	}).(ServiceEndpointGitHubOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceEndpointGitHubInput)(nil)).Elem(), &ServiceEndpointGitHub{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceEndpointGitHubArrayInput)(nil)).Elem(), ServiceEndpointGitHubArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceEndpointGitHubMapInput)(nil)).Elem(), ServiceEndpointGitHubMap{})
	pulumi.RegisterOutputType(ServiceEndpointGitHubOutput{})
	pulumi.RegisterOutputType(ServiceEndpointGitHubArrayOutput{})
	pulumi.RegisterOutputType(ServiceEndpointGitHubMapOutput{})
}
