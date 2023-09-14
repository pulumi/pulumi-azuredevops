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

// Manages a Docker Registry service endpoint within Azure DevOps.
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
//			_, err = azuredevops.NewServiceEndpointDockerRegistry(ctx, "exampleServiceEndpointDockerRegistry", &azuredevops.ServiceEndpointDockerRegistryArgs{
//				ProjectId:           exampleProject.ID(),
//				ServiceEndpointName: pulumi.String("Example Docker Hub"),
//				DockerUsername:      pulumi.String("example"),
//				DockerEmail:         pulumi.String("email@example.com"),
//				DockerPassword:      pulumi.String("12345"),
//				RegistryType:        pulumi.String("DockerHub"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuredevops.NewServiceEndpointDockerRegistry(ctx, "example-other", &azuredevops.ServiceEndpointDockerRegistryArgs{
//				ProjectId:           exampleProject.ID(),
//				ServiceEndpointName: pulumi.String("Example Docker Registry"),
//				DockerRegistry:      pulumi.String("https://sample.azurecr.io/v1"),
//				DockerUsername:      pulumi.String("sample"),
//				DockerPassword:      pulumi.String("12345"),
//				RegistryType:        pulumi.String("Others"),
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
// - [Docker Registry Service Connection](https://docs.microsoft.com/en-us/azure/devops/pipelines/library/service-endpoints?view=azure-devops&tabs=yaml#sep-docreg)
//
// ## Import
//
// Azure DevOps Service Endpoint Docker Registry can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
//
// ```sh
//
//	$ pulumi import azuredevops:index/serviceEndpointDockerRegistry:ServiceEndpointDockerRegistry example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
//
// ```
type ServiceEndpointDockerRegistry struct {
	pulumi.CustomResourceState

	Authorization pulumi.StringMapOutput `pulumi:"authorization"`
	Description   pulumi.StringPtrOutput `pulumi:"description"`
	// The email for Docker account user.
	DockerEmail pulumi.StringPtrOutput `pulumi:"dockerEmail"`
	// The password for the account user identified above.
	DockerPassword pulumi.StringPtrOutput `pulumi:"dockerPassword"`
	// The URL of the Docker registry. (Default: "https://index.docker.io/v1/")
	DockerRegistry pulumi.StringOutput `pulumi:"dockerRegistry"`
	// The identifier of the Docker account user.
	DockerUsername pulumi.StringPtrOutput `pulumi:"dockerUsername"`
	// The ID of the project.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Can be "DockerHub" or "Others" (Default "DockerHub")
	RegistryType pulumi.StringOutput `pulumi:"registryType"`
	// The name you will use to refer to this service connection in task inputs.
	ServiceEndpointName pulumi.StringOutput `pulumi:"serviceEndpointName"`
}

// NewServiceEndpointDockerRegistry registers a new resource with the given unique name, arguments, and options.
func NewServiceEndpointDockerRegistry(ctx *pulumi.Context,
	name string, args *ServiceEndpointDockerRegistryArgs, opts ...pulumi.ResourceOption) (*ServiceEndpointDockerRegistry, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DockerRegistry == nil {
		return nil, errors.New("invalid value for required argument 'DockerRegistry'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.RegistryType == nil {
		return nil, errors.New("invalid value for required argument 'RegistryType'")
	}
	if args.ServiceEndpointName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceEndpointName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azuredevops:ServiceEndpoint/dockerRegistry:DockerRegistry"),
		},
	})
	opts = append(opts, aliases)
	if args.DockerPassword != nil {
		args.DockerPassword = pulumi.ToSecret(args.DockerPassword).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"dockerPassword",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ServiceEndpointDockerRegistry
	err := ctx.RegisterResource("azuredevops:index/serviceEndpointDockerRegistry:ServiceEndpointDockerRegistry", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceEndpointDockerRegistry gets an existing ServiceEndpointDockerRegistry resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceEndpointDockerRegistry(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceEndpointDockerRegistryState, opts ...pulumi.ResourceOption) (*ServiceEndpointDockerRegistry, error) {
	var resource ServiceEndpointDockerRegistry
	err := ctx.ReadResource("azuredevops:index/serviceEndpointDockerRegistry:ServiceEndpointDockerRegistry", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceEndpointDockerRegistry resources.
type serviceEndpointDockerRegistryState struct {
	Authorization map[string]string `pulumi:"authorization"`
	Description   *string           `pulumi:"description"`
	// The email for Docker account user.
	DockerEmail *string `pulumi:"dockerEmail"`
	// The password for the account user identified above.
	DockerPassword *string `pulumi:"dockerPassword"`
	// The URL of the Docker registry. (Default: "https://index.docker.io/v1/")
	DockerRegistry *string `pulumi:"dockerRegistry"`
	// The identifier of the Docker account user.
	DockerUsername *string `pulumi:"dockerUsername"`
	// The ID of the project.
	ProjectId *string `pulumi:"projectId"`
	// Can be "DockerHub" or "Others" (Default "DockerHub")
	RegistryType *string `pulumi:"registryType"`
	// The name you will use to refer to this service connection in task inputs.
	ServiceEndpointName *string `pulumi:"serviceEndpointName"`
}

type ServiceEndpointDockerRegistryState struct {
	Authorization pulumi.StringMapInput
	Description   pulumi.StringPtrInput
	// The email for Docker account user.
	DockerEmail pulumi.StringPtrInput
	// The password for the account user identified above.
	DockerPassword pulumi.StringPtrInput
	// The URL of the Docker registry. (Default: "https://index.docker.io/v1/")
	DockerRegistry pulumi.StringPtrInput
	// The identifier of the Docker account user.
	DockerUsername pulumi.StringPtrInput
	// The ID of the project.
	ProjectId pulumi.StringPtrInput
	// Can be "DockerHub" or "Others" (Default "DockerHub")
	RegistryType pulumi.StringPtrInput
	// The name you will use to refer to this service connection in task inputs.
	ServiceEndpointName pulumi.StringPtrInput
}

func (ServiceEndpointDockerRegistryState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceEndpointDockerRegistryState)(nil)).Elem()
}

type serviceEndpointDockerRegistryArgs struct {
	Authorization map[string]string `pulumi:"authorization"`
	Description   *string           `pulumi:"description"`
	// The email for Docker account user.
	DockerEmail *string `pulumi:"dockerEmail"`
	// The password for the account user identified above.
	DockerPassword *string `pulumi:"dockerPassword"`
	// The URL of the Docker registry. (Default: "https://index.docker.io/v1/")
	DockerRegistry string `pulumi:"dockerRegistry"`
	// The identifier of the Docker account user.
	DockerUsername *string `pulumi:"dockerUsername"`
	// The ID of the project.
	ProjectId string `pulumi:"projectId"`
	// Can be "DockerHub" or "Others" (Default "DockerHub")
	RegistryType string `pulumi:"registryType"`
	// The name you will use to refer to this service connection in task inputs.
	ServiceEndpointName string `pulumi:"serviceEndpointName"`
}

// The set of arguments for constructing a ServiceEndpointDockerRegistry resource.
type ServiceEndpointDockerRegistryArgs struct {
	Authorization pulumi.StringMapInput
	Description   pulumi.StringPtrInput
	// The email for Docker account user.
	DockerEmail pulumi.StringPtrInput
	// The password for the account user identified above.
	DockerPassword pulumi.StringPtrInput
	// The URL of the Docker registry. (Default: "https://index.docker.io/v1/")
	DockerRegistry pulumi.StringInput
	// The identifier of the Docker account user.
	DockerUsername pulumi.StringPtrInput
	// The ID of the project.
	ProjectId pulumi.StringInput
	// Can be "DockerHub" or "Others" (Default "DockerHub")
	RegistryType pulumi.StringInput
	// The name you will use to refer to this service connection in task inputs.
	ServiceEndpointName pulumi.StringInput
}

func (ServiceEndpointDockerRegistryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceEndpointDockerRegistryArgs)(nil)).Elem()
}

type ServiceEndpointDockerRegistryInput interface {
	pulumi.Input

	ToServiceEndpointDockerRegistryOutput() ServiceEndpointDockerRegistryOutput
	ToServiceEndpointDockerRegistryOutputWithContext(ctx context.Context) ServiceEndpointDockerRegistryOutput
}

func (*ServiceEndpointDockerRegistry) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceEndpointDockerRegistry)(nil)).Elem()
}

func (i *ServiceEndpointDockerRegistry) ToServiceEndpointDockerRegistryOutput() ServiceEndpointDockerRegistryOutput {
	return i.ToServiceEndpointDockerRegistryOutputWithContext(context.Background())
}

func (i *ServiceEndpointDockerRegistry) ToServiceEndpointDockerRegistryOutputWithContext(ctx context.Context) ServiceEndpointDockerRegistryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointDockerRegistryOutput)
}

func (i *ServiceEndpointDockerRegistry) ToOutput(ctx context.Context) pulumix.Output[*ServiceEndpointDockerRegistry] {
	return pulumix.Output[*ServiceEndpointDockerRegistry]{
		OutputState: i.ToServiceEndpointDockerRegistryOutputWithContext(ctx).OutputState,
	}
}

// ServiceEndpointDockerRegistryArrayInput is an input type that accepts ServiceEndpointDockerRegistryArray and ServiceEndpointDockerRegistryArrayOutput values.
// You can construct a concrete instance of `ServiceEndpointDockerRegistryArrayInput` via:
//
//	ServiceEndpointDockerRegistryArray{ ServiceEndpointDockerRegistryArgs{...} }
type ServiceEndpointDockerRegistryArrayInput interface {
	pulumi.Input

	ToServiceEndpointDockerRegistryArrayOutput() ServiceEndpointDockerRegistryArrayOutput
	ToServiceEndpointDockerRegistryArrayOutputWithContext(context.Context) ServiceEndpointDockerRegistryArrayOutput
}

type ServiceEndpointDockerRegistryArray []ServiceEndpointDockerRegistryInput

func (ServiceEndpointDockerRegistryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceEndpointDockerRegistry)(nil)).Elem()
}

func (i ServiceEndpointDockerRegistryArray) ToServiceEndpointDockerRegistryArrayOutput() ServiceEndpointDockerRegistryArrayOutput {
	return i.ToServiceEndpointDockerRegistryArrayOutputWithContext(context.Background())
}

func (i ServiceEndpointDockerRegistryArray) ToServiceEndpointDockerRegistryArrayOutputWithContext(ctx context.Context) ServiceEndpointDockerRegistryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointDockerRegistryArrayOutput)
}

func (i ServiceEndpointDockerRegistryArray) ToOutput(ctx context.Context) pulumix.Output[[]*ServiceEndpointDockerRegistry] {
	return pulumix.Output[[]*ServiceEndpointDockerRegistry]{
		OutputState: i.ToServiceEndpointDockerRegistryArrayOutputWithContext(ctx).OutputState,
	}
}

// ServiceEndpointDockerRegistryMapInput is an input type that accepts ServiceEndpointDockerRegistryMap and ServiceEndpointDockerRegistryMapOutput values.
// You can construct a concrete instance of `ServiceEndpointDockerRegistryMapInput` via:
//
//	ServiceEndpointDockerRegistryMap{ "key": ServiceEndpointDockerRegistryArgs{...} }
type ServiceEndpointDockerRegistryMapInput interface {
	pulumi.Input

	ToServiceEndpointDockerRegistryMapOutput() ServiceEndpointDockerRegistryMapOutput
	ToServiceEndpointDockerRegistryMapOutputWithContext(context.Context) ServiceEndpointDockerRegistryMapOutput
}

type ServiceEndpointDockerRegistryMap map[string]ServiceEndpointDockerRegistryInput

func (ServiceEndpointDockerRegistryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceEndpointDockerRegistry)(nil)).Elem()
}

func (i ServiceEndpointDockerRegistryMap) ToServiceEndpointDockerRegistryMapOutput() ServiceEndpointDockerRegistryMapOutput {
	return i.ToServiceEndpointDockerRegistryMapOutputWithContext(context.Background())
}

func (i ServiceEndpointDockerRegistryMap) ToServiceEndpointDockerRegistryMapOutputWithContext(ctx context.Context) ServiceEndpointDockerRegistryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointDockerRegistryMapOutput)
}

func (i ServiceEndpointDockerRegistryMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*ServiceEndpointDockerRegistry] {
	return pulumix.Output[map[string]*ServiceEndpointDockerRegistry]{
		OutputState: i.ToServiceEndpointDockerRegistryMapOutputWithContext(ctx).OutputState,
	}
}

type ServiceEndpointDockerRegistryOutput struct{ *pulumi.OutputState }

func (ServiceEndpointDockerRegistryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceEndpointDockerRegistry)(nil)).Elem()
}

func (o ServiceEndpointDockerRegistryOutput) ToServiceEndpointDockerRegistryOutput() ServiceEndpointDockerRegistryOutput {
	return o
}

func (o ServiceEndpointDockerRegistryOutput) ToServiceEndpointDockerRegistryOutputWithContext(ctx context.Context) ServiceEndpointDockerRegistryOutput {
	return o
}

func (o ServiceEndpointDockerRegistryOutput) ToOutput(ctx context.Context) pulumix.Output[*ServiceEndpointDockerRegistry] {
	return pulumix.Output[*ServiceEndpointDockerRegistry]{
		OutputState: o.OutputState,
	}
}

func (o ServiceEndpointDockerRegistryOutput) Authorization() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ServiceEndpointDockerRegistry) pulumi.StringMapOutput { return v.Authorization }).(pulumi.StringMapOutput)
}

func (o ServiceEndpointDockerRegistryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceEndpointDockerRegistry) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The email for Docker account user.
func (o ServiceEndpointDockerRegistryOutput) DockerEmail() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceEndpointDockerRegistry) pulumi.StringPtrOutput { return v.DockerEmail }).(pulumi.StringPtrOutput)
}

// The password for the account user identified above.
func (o ServiceEndpointDockerRegistryOutput) DockerPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceEndpointDockerRegistry) pulumi.StringPtrOutput { return v.DockerPassword }).(pulumi.StringPtrOutput)
}

// The URL of the Docker registry. (Default: "https://index.docker.io/v1/")
func (o ServiceEndpointDockerRegistryOutput) DockerRegistry() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEndpointDockerRegistry) pulumi.StringOutput { return v.DockerRegistry }).(pulumi.StringOutput)
}

// The identifier of the Docker account user.
func (o ServiceEndpointDockerRegistryOutput) DockerUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceEndpointDockerRegistry) pulumi.StringPtrOutput { return v.DockerUsername }).(pulumi.StringPtrOutput)
}

// The ID of the project.
func (o ServiceEndpointDockerRegistryOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEndpointDockerRegistry) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// Can be "DockerHub" or "Others" (Default "DockerHub")
func (o ServiceEndpointDockerRegistryOutput) RegistryType() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEndpointDockerRegistry) pulumi.StringOutput { return v.RegistryType }).(pulumi.StringOutput)
}

// The name you will use to refer to this service connection in task inputs.
func (o ServiceEndpointDockerRegistryOutput) ServiceEndpointName() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEndpointDockerRegistry) pulumi.StringOutput { return v.ServiceEndpointName }).(pulumi.StringOutput)
}

type ServiceEndpointDockerRegistryArrayOutput struct{ *pulumi.OutputState }

func (ServiceEndpointDockerRegistryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceEndpointDockerRegistry)(nil)).Elem()
}

func (o ServiceEndpointDockerRegistryArrayOutput) ToServiceEndpointDockerRegistryArrayOutput() ServiceEndpointDockerRegistryArrayOutput {
	return o
}

func (o ServiceEndpointDockerRegistryArrayOutput) ToServiceEndpointDockerRegistryArrayOutputWithContext(ctx context.Context) ServiceEndpointDockerRegistryArrayOutput {
	return o
}

func (o ServiceEndpointDockerRegistryArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*ServiceEndpointDockerRegistry] {
	return pulumix.Output[[]*ServiceEndpointDockerRegistry]{
		OutputState: o.OutputState,
	}
}

func (o ServiceEndpointDockerRegistryArrayOutput) Index(i pulumi.IntInput) ServiceEndpointDockerRegistryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServiceEndpointDockerRegistry {
		return vs[0].([]*ServiceEndpointDockerRegistry)[vs[1].(int)]
	}).(ServiceEndpointDockerRegistryOutput)
}

type ServiceEndpointDockerRegistryMapOutput struct{ *pulumi.OutputState }

func (ServiceEndpointDockerRegistryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceEndpointDockerRegistry)(nil)).Elem()
}

func (o ServiceEndpointDockerRegistryMapOutput) ToServiceEndpointDockerRegistryMapOutput() ServiceEndpointDockerRegistryMapOutput {
	return o
}

func (o ServiceEndpointDockerRegistryMapOutput) ToServiceEndpointDockerRegistryMapOutputWithContext(ctx context.Context) ServiceEndpointDockerRegistryMapOutput {
	return o
}

func (o ServiceEndpointDockerRegistryMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*ServiceEndpointDockerRegistry] {
	return pulumix.Output[map[string]*ServiceEndpointDockerRegistry]{
		OutputState: o.OutputState,
	}
}

func (o ServiceEndpointDockerRegistryMapOutput) MapIndex(k pulumi.StringInput) ServiceEndpointDockerRegistryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServiceEndpointDockerRegistry {
		return vs[0].(map[string]*ServiceEndpointDockerRegistry)[vs[1].(string)]
	}).(ServiceEndpointDockerRegistryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceEndpointDockerRegistryInput)(nil)).Elem(), &ServiceEndpointDockerRegistry{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceEndpointDockerRegistryArrayInput)(nil)).Elem(), ServiceEndpointDockerRegistryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceEndpointDockerRegistryMapInput)(nil)).Elem(), ServiceEndpointDockerRegistryMap{})
	pulumi.RegisterOutputType(ServiceEndpointDockerRegistryOutput{})
	pulumi.RegisterOutputType(ServiceEndpointDockerRegistryArrayOutput{})
	pulumi.RegisterOutputType(ServiceEndpointDockerRegistryMapOutput{})
}
