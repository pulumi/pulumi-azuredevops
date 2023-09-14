// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azuredevops/sdk/v2/go/azuredevops/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Manages a SonarQube service endpoint within Azure DevOps.
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
//			_, err = azuredevops.NewServiceEndpointSonarQube(ctx, "exampleServiceEndpointSonarQube", &azuredevops.ServiceEndpointSonarQubeArgs{
//				ProjectId:           exampleProject.ID(),
//				ServiceEndpointName: pulumi.String("Example SonarQube"),
//				Url:                 pulumi.String("https://sonarqube.my.com"),
//				Token:               pulumi.String("0000000000000000000000000000000000000000"),
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
// - [Azure DevOps Service REST API 7.0 - Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-7.0)
// - [Azure DevOps Service Connections](https://docs.microsoft.com/en-us/azure/devops/pipelines/library/service-endpoints?view=azure-devops&tabs=yaml)
// - [SonarQube User Token](https://docs.sonarqube.org/latest/user-guide/user-token/)
//
// ## Import
//
// Azure DevOps Service Endpoint SonarQube can be imported using the **projectID/serviceEndpointID**, e.g.
//
// ```sh
//
//	$ pulumi import azuredevops:index/serviceEndpointSonarQube:ServiceEndpointSonarQube example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
//
// ```
type ServiceEndpointSonarQube struct {
	pulumi.CustomResourceState

	Authorization pulumi.StringMapOutput `pulumi:"authorization"`
	// The Service Endpoint description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The ID of the project.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The Service Endpoint name.
	ServiceEndpointName pulumi.StringOutput `pulumi:"serviceEndpointName"`
	// Authentication Token generated through SonarQube (go to My Account > Security > Generate Tokens).
	Token pulumi.StringOutput `pulumi:"token"`
	// URL of the SonarQube server to connect with.
	Url pulumi.StringOutput `pulumi:"url"`
}

// NewServiceEndpointSonarQube registers a new resource with the given unique name, arguments, and options.
func NewServiceEndpointSonarQube(ctx *pulumi.Context,
	name string, args *ServiceEndpointSonarQubeArgs, opts ...pulumi.ResourceOption) (*ServiceEndpointSonarQube, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.ServiceEndpointName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceEndpointName'")
	}
	if args.Token == nil {
		return nil, errors.New("invalid value for required argument 'Token'")
	}
	if args.Url == nil {
		return nil, errors.New("invalid value for required argument 'Url'")
	}
	if args.Token != nil {
		args.Token = pulumi.ToSecret(args.Token).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"token",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ServiceEndpointSonarQube
	err := ctx.RegisterResource("azuredevops:index/serviceEndpointSonarQube:ServiceEndpointSonarQube", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceEndpointSonarQube gets an existing ServiceEndpointSonarQube resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceEndpointSonarQube(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceEndpointSonarQubeState, opts ...pulumi.ResourceOption) (*ServiceEndpointSonarQube, error) {
	var resource ServiceEndpointSonarQube
	err := ctx.ReadResource("azuredevops:index/serviceEndpointSonarQube:ServiceEndpointSonarQube", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceEndpointSonarQube resources.
type serviceEndpointSonarQubeState struct {
	Authorization map[string]string `pulumi:"authorization"`
	// The Service Endpoint description.
	Description *string `pulumi:"description"`
	// The ID of the project.
	ProjectId *string `pulumi:"projectId"`
	// The Service Endpoint name.
	ServiceEndpointName *string `pulumi:"serviceEndpointName"`
	// Authentication Token generated through SonarQube (go to My Account > Security > Generate Tokens).
	Token *string `pulumi:"token"`
	// URL of the SonarQube server to connect with.
	Url *string `pulumi:"url"`
}

type ServiceEndpointSonarQubeState struct {
	Authorization pulumi.StringMapInput
	// The Service Endpoint description.
	Description pulumi.StringPtrInput
	// The ID of the project.
	ProjectId pulumi.StringPtrInput
	// The Service Endpoint name.
	ServiceEndpointName pulumi.StringPtrInput
	// Authentication Token generated through SonarQube (go to My Account > Security > Generate Tokens).
	Token pulumi.StringPtrInput
	// URL of the SonarQube server to connect with.
	Url pulumi.StringPtrInput
}

func (ServiceEndpointSonarQubeState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceEndpointSonarQubeState)(nil)).Elem()
}

type serviceEndpointSonarQubeArgs struct {
	Authorization map[string]string `pulumi:"authorization"`
	// The Service Endpoint description.
	Description *string `pulumi:"description"`
	// The ID of the project.
	ProjectId string `pulumi:"projectId"`
	// The Service Endpoint name.
	ServiceEndpointName string `pulumi:"serviceEndpointName"`
	// Authentication Token generated through SonarQube (go to My Account > Security > Generate Tokens).
	Token string `pulumi:"token"`
	// URL of the SonarQube server to connect with.
	Url string `pulumi:"url"`
}

// The set of arguments for constructing a ServiceEndpointSonarQube resource.
type ServiceEndpointSonarQubeArgs struct {
	Authorization pulumi.StringMapInput
	// The Service Endpoint description.
	Description pulumi.StringPtrInput
	// The ID of the project.
	ProjectId pulumi.StringInput
	// The Service Endpoint name.
	ServiceEndpointName pulumi.StringInput
	// Authentication Token generated through SonarQube (go to My Account > Security > Generate Tokens).
	Token pulumi.StringInput
	// URL of the SonarQube server to connect with.
	Url pulumi.StringInput
}

func (ServiceEndpointSonarQubeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceEndpointSonarQubeArgs)(nil)).Elem()
}

type ServiceEndpointSonarQubeInput interface {
	pulumi.Input

	ToServiceEndpointSonarQubeOutput() ServiceEndpointSonarQubeOutput
	ToServiceEndpointSonarQubeOutputWithContext(ctx context.Context) ServiceEndpointSonarQubeOutput
}

func (*ServiceEndpointSonarQube) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceEndpointSonarQube)(nil)).Elem()
}

func (i *ServiceEndpointSonarQube) ToServiceEndpointSonarQubeOutput() ServiceEndpointSonarQubeOutput {
	return i.ToServiceEndpointSonarQubeOutputWithContext(context.Background())
}

func (i *ServiceEndpointSonarQube) ToServiceEndpointSonarQubeOutputWithContext(ctx context.Context) ServiceEndpointSonarQubeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointSonarQubeOutput)
}

func (i *ServiceEndpointSonarQube) ToOutput(ctx context.Context) pulumix.Output[*ServiceEndpointSonarQube] {
	return pulumix.Output[*ServiceEndpointSonarQube]{
		OutputState: i.ToServiceEndpointSonarQubeOutputWithContext(ctx).OutputState,
	}
}

// ServiceEndpointSonarQubeArrayInput is an input type that accepts ServiceEndpointSonarQubeArray and ServiceEndpointSonarQubeArrayOutput values.
// You can construct a concrete instance of `ServiceEndpointSonarQubeArrayInput` via:
//
//	ServiceEndpointSonarQubeArray{ ServiceEndpointSonarQubeArgs{...} }
type ServiceEndpointSonarQubeArrayInput interface {
	pulumi.Input

	ToServiceEndpointSonarQubeArrayOutput() ServiceEndpointSonarQubeArrayOutput
	ToServiceEndpointSonarQubeArrayOutputWithContext(context.Context) ServiceEndpointSonarQubeArrayOutput
}

type ServiceEndpointSonarQubeArray []ServiceEndpointSonarQubeInput

func (ServiceEndpointSonarQubeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceEndpointSonarQube)(nil)).Elem()
}

func (i ServiceEndpointSonarQubeArray) ToServiceEndpointSonarQubeArrayOutput() ServiceEndpointSonarQubeArrayOutput {
	return i.ToServiceEndpointSonarQubeArrayOutputWithContext(context.Background())
}

func (i ServiceEndpointSonarQubeArray) ToServiceEndpointSonarQubeArrayOutputWithContext(ctx context.Context) ServiceEndpointSonarQubeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointSonarQubeArrayOutput)
}

func (i ServiceEndpointSonarQubeArray) ToOutput(ctx context.Context) pulumix.Output[[]*ServiceEndpointSonarQube] {
	return pulumix.Output[[]*ServiceEndpointSonarQube]{
		OutputState: i.ToServiceEndpointSonarQubeArrayOutputWithContext(ctx).OutputState,
	}
}

// ServiceEndpointSonarQubeMapInput is an input type that accepts ServiceEndpointSonarQubeMap and ServiceEndpointSonarQubeMapOutput values.
// You can construct a concrete instance of `ServiceEndpointSonarQubeMapInput` via:
//
//	ServiceEndpointSonarQubeMap{ "key": ServiceEndpointSonarQubeArgs{...} }
type ServiceEndpointSonarQubeMapInput interface {
	pulumi.Input

	ToServiceEndpointSonarQubeMapOutput() ServiceEndpointSonarQubeMapOutput
	ToServiceEndpointSonarQubeMapOutputWithContext(context.Context) ServiceEndpointSonarQubeMapOutput
}

type ServiceEndpointSonarQubeMap map[string]ServiceEndpointSonarQubeInput

func (ServiceEndpointSonarQubeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceEndpointSonarQube)(nil)).Elem()
}

func (i ServiceEndpointSonarQubeMap) ToServiceEndpointSonarQubeMapOutput() ServiceEndpointSonarQubeMapOutput {
	return i.ToServiceEndpointSonarQubeMapOutputWithContext(context.Background())
}

func (i ServiceEndpointSonarQubeMap) ToServiceEndpointSonarQubeMapOutputWithContext(ctx context.Context) ServiceEndpointSonarQubeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointSonarQubeMapOutput)
}

func (i ServiceEndpointSonarQubeMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*ServiceEndpointSonarQube] {
	return pulumix.Output[map[string]*ServiceEndpointSonarQube]{
		OutputState: i.ToServiceEndpointSonarQubeMapOutputWithContext(ctx).OutputState,
	}
}

type ServiceEndpointSonarQubeOutput struct{ *pulumi.OutputState }

func (ServiceEndpointSonarQubeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceEndpointSonarQube)(nil)).Elem()
}

func (o ServiceEndpointSonarQubeOutput) ToServiceEndpointSonarQubeOutput() ServiceEndpointSonarQubeOutput {
	return o
}

func (o ServiceEndpointSonarQubeOutput) ToServiceEndpointSonarQubeOutputWithContext(ctx context.Context) ServiceEndpointSonarQubeOutput {
	return o
}

func (o ServiceEndpointSonarQubeOutput) ToOutput(ctx context.Context) pulumix.Output[*ServiceEndpointSonarQube] {
	return pulumix.Output[*ServiceEndpointSonarQube]{
		OutputState: o.OutputState,
	}
}

func (o ServiceEndpointSonarQubeOutput) Authorization() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ServiceEndpointSonarQube) pulumi.StringMapOutput { return v.Authorization }).(pulumi.StringMapOutput)
}

// The Service Endpoint description.
func (o ServiceEndpointSonarQubeOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceEndpointSonarQube) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The ID of the project.
func (o ServiceEndpointSonarQubeOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEndpointSonarQube) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The Service Endpoint name.
func (o ServiceEndpointSonarQubeOutput) ServiceEndpointName() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEndpointSonarQube) pulumi.StringOutput { return v.ServiceEndpointName }).(pulumi.StringOutput)
}

// Authentication Token generated through SonarQube (go to My Account > Security > Generate Tokens).
func (o ServiceEndpointSonarQubeOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEndpointSonarQube) pulumi.StringOutput { return v.Token }).(pulumi.StringOutput)
}

// URL of the SonarQube server to connect with.
func (o ServiceEndpointSonarQubeOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEndpointSonarQube) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

type ServiceEndpointSonarQubeArrayOutput struct{ *pulumi.OutputState }

func (ServiceEndpointSonarQubeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceEndpointSonarQube)(nil)).Elem()
}

func (o ServiceEndpointSonarQubeArrayOutput) ToServiceEndpointSonarQubeArrayOutput() ServiceEndpointSonarQubeArrayOutput {
	return o
}

func (o ServiceEndpointSonarQubeArrayOutput) ToServiceEndpointSonarQubeArrayOutputWithContext(ctx context.Context) ServiceEndpointSonarQubeArrayOutput {
	return o
}

func (o ServiceEndpointSonarQubeArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*ServiceEndpointSonarQube] {
	return pulumix.Output[[]*ServiceEndpointSonarQube]{
		OutputState: o.OutputState,
	}
}

func (o ServiceEndpointSonarQubeArrayOutput) Index(i pulumi.IntInput) ServiceEndpointSonarQubeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServiceEndpointSonarQube {
		return vs[0].([]*ServiceEndpointSonarQube)[vs[1].(int)]
	}).(ServiceEndpointSonarQubeOutput)
}

type ServiceEndpointSonarQubeMapOutput struct{ *pulumi.OutputState }

func (ServiceEndpointSonarQubeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceEndpointSonarQube)(nil)).Elem()
}

func (o ServiceEndpointSonarQubeMapOutput) ToServiceEndpointSonarQubeMapOutput() ServiceEndpointSonarQubeMapOutput {
	return o
}

func (o ServiceEndpointSonarQubeMapOutput) ToServiceEndpointSonarQubeMapOutputWithContext(ctx context.Context) ServiceEndpointSonarQubeMapOutput {
	return o
}

func (o ServiceEndpointSonarQubeMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*ServiceEndpointSonarQube] {
	return pulumix.Output[map[string]*ServiceEndpointSonarQube]{
		OutputState: o.OutputState,
	}
}

func (o ServiceEndpointSonarQubeMapOutput) MapIndex(k pulumi.StringInput) ServiceEndpointSonarQubeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServiceEndpointSonarQube {
		return vs[0].(map[string]*ServiceEndpointSonarQube)[vs[1].(string)]
	}).(ServiceEndpointSonarQubeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceEndpointSonarQubeInput)(nil)).Elem(), &ServiceEndpointSonarQube{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceEndpointSonarQubeArrayInput)(nil)).Elem(), ServiceEndpointSonarQubeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceEndpointSonarQubeMapInput)(nil)).Elem(), ServiceEndpointSonarQubeMap{})
	pulumi.RegisterOutputType(ServiceEndpointSonarQubeOutput{})
	pulumi.RegisterOutputType(ServiceEndpointSonarQubeArrayOutput{})
	pulumi.RegisterOutputType(ServiceEndpointSonarQubeMapOutput{})
}
