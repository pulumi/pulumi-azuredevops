// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package serviceendpoint

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Docker Registry service endpoint within Azure DevOps.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azuredevops/sdk/v2/go/azuredevops"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		project, err := azuredevops.NewProject(ctx, "project", &azuredevops.ProjectArgs{
// 			Visibility:       pulumi.String("private"),
// 			VersionControl:   pulumi.String("Git"),
// 			WorkItemTemplate: pulumi.String("Agile"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = azuredevops.NewServiceEndpointDockerRegistry(ctx, "dockerhubregistry", &azuredevops.ServiceEndpointDockerRegistryArgs{
// 			ProjectId:           project.ID(),
// 			ServiceEndpointName: pulumi.String("Sample Docker Hub"),
// 			DockerUsername:      pulumi.String("sample"),
// 			DockerEmail:         pulumi.String("email@example.com"),
// 			DockerPassword:      pulumi.String("12345"),
// 			RegistryType:        pulumi.String("DockerHub"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = azuredevops.NewServiceEndpointDockerRegistry(ctx, "otherregistry", &azuredevops.ServiceEndpointDockerRegistryArgs{
// 			ProjectId:           project.ID(),
// 			ServiceEndpointName: pulumi.String("Sample Docker Registry"),
// 			DockerRegistry:      pulumi.String("https://sample.azurecr.io/v1"),
// 			DockerUsername:      pulumi.String("sample"),
// 			DockerPassword:      pulumi.String("12345"),
// 			RegistryType:        pulumi.String("Others"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## Relevant Links
//
// - [Azure DevOps Service REST API 5.1 - Service Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-5.1)
// - [Docker Registry Service Connection](https://docs.microsoft.com/en-us/azure/devops/pipelines/library/service-endpoints?view=azure-devops&tabs=yaml#sep-docreg)
//
// ## Import
//
// Azure DevOps Service Endpoint Docker Registry can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
//
// ```sh
//  $ pulumi import azuredevops:ServiceEndpoint/dockerRegistry:DockerRegistry serviceendpoint 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
// ```
//
// Deprecated: azuredevops.serviceendpoint.DockerRegistry has been deprecated in favor of azuredevops.ServiceEndpointDockerRegistry
type DockerRegistry struct {
	pulumi.CustomResourceState

	Authorization pulumi.StringMapOutput `pulumi:"authorization"`
	// The name you will use to refer to this service connection in task inputs.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The email for Docker account user.
	DockerEmail pulumi.StringPtrOutput `pulumi:"dockerEmail"`
	// The password for the account user identified above.
	DockerPassword pulumi.StringPtrOutput `pulumi:"dockerPassword"`
	// A bcrypted hash of the attribute 'docker_password'
	DockerPasswordHash pulumi.StringOutput `pulumi:"dockerPasswordHash"`
	// The URL of the Docker registry. (Default: "https://index.docker.io/v1/")
	DockerRegistry pulumi.StringOutput `pulumi:"dockerRegistry"`
	// The identifier of the Docker account user.
	DockerUsername pulumi.StringPtrOutput `pulumi:"dockerUsername"`
	// The project ID or project name.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Can be "DockerHub" or "Others" (Default "DockerHub")
	RegistryType pulumi.StringOutput `pulumi:"registryType"`
	// The name you will use to refer to this service connection in task inputs.
	ServiceEndpointName pulumi.StringOutput `pulumi:"serviceEndpointName"`
}

// NewDockerRegistry registers a new resource with the given unique name, arguments, and options.
func NewDockerRegistry(ctx *pulumi.Context,
	name string, args *DockerRegistryArgs, opts ...pulumi.ResourceOption) (*DockerRegistry, error) {
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
	var resource DockerRegistry
	err := ctx.RegisterResource("azuredevops:ServiceEndpoint/dockerRegistry:DockerRegistry", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDockerRegistry gets an existing DockerRegistry resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDockerRegistry(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DockerRegistryState, opts ...pulumi.ResourceOption) (*DockerRegistry, error) {
	var resource DockerRegistry
	err := ctx.ReadResource("azuredevops:ServiceEndpoint/dockerRegistry:DockerRegistry", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DockerRegistry resources.
type dockerRegistryState struct {
	Authorization map[string]string `pulumi:"authorization"`
	// The name you will use to refer to this service connection in task inputs.
	Description *string `pulumi:"description"`
	// The email for Docker account user.
	DockerEmail *string `pulumi:"dockerEmail"`
	// The password for the account user identified above.
	DockerPassword *string `pulumi:"dockerPassword"`
	// A bcrypted hash of the attribute 'docker_password'
	DockerPasswordHash *string `pulumi:"dockerPasswordHash"`
	// The URL of the Docker registry. (Default: "https://index.docker.io/v1/")
	DockerRegistry *string `pulumi:"dockerRegistry"`
	// The identifier of the Docker account user.
	DockerUsername *string `pulumi:"dockerUsername"`
	// The project ID or project name.
	ProjectId *string `pulumi:"projectId"`
	// Can be "DockerHub" or "Others" (Default "DockerHub")
	RegistryType *string `pulumi:"registryType"`
	// The name you will use to refer to this service connection in task inputs.
	ServiceEndpointName *string `pulumi:"serviceEndpointName"`
}

type DockerRegistryState struct {
	Authorization pulumi.StringMapInput
	// The name you will use to refer to this service connection in task inputs.
	Description pulumi.StringPtrInput
	// The email for Docker account user.
	DockerEmail pulumi.StringPtrInput
	// The password for the account user identified above.
	DockerPassword pulumi.StringPtrInput
	// A bcrypted hash of the attribute 'docker_password'
	DockerPasswordHash pulumi.StringPtrInput
	// The URL of the Docker registry. (Default: "https://index.docker.io/v1/")
	DockerRegistry pulumi.StringPtrInput
	// The identifier of the Docker account user.
	DockerUsername pulumi.StringPtrInput
	// The project ID or project name.
	ProjectId pulumi.StringPtrInput
	// Can be "DockerHub" or "Others" (Default "DockerHub")
	RegistryType pulumi.StringPtrInput
	// The name you will use to refer to this service connection in task inputs.
	ServiceEndpointName pulumi.StringPtrInput
}

func (DockerRegistryState) ElementType() reflect.Type {
	return reflect.TypeOf((*dockerRegistryState)(nil)).Elem()
}

type dockerRegistryArgs struct {
	Authorization map[string]string `pulumi:"authorization"`
	// The name you will use to refer to this service connection in task inputs.
	Description *string `pulumi:"description"`
	// The email for Docker account user.
	DockerEmail *string `pulumi:"dockerEmail"`
	// The password for the account user identified above.
	DockerPassword *string `pulumi:"dockerPassword"`
	// The URL of the Docker registry. (Default: "https://index.docker.io/v1/")
	DockerRegistry string `pulumi:"dockerRegistry"`
	// The identifier of the Docker account user.
	DockerUsername *string `pulumi:"dockerUsername"`
	// The project ID or project name.
	ProjectId string `pulumi:"projectId"`
	// Can be "DockerHub" or "Others" (Default "DockerHub")
	RegistryType string `pulumi:"registryType"`
	// The name you will use to refer to this service connection in task inputs.
	ServiceEndpointName string `pulumi:"serviceEndpointName"`
}

// The set of arguments for constructing a DockerRegistry resource.
type DockerRegistryArgs struct {
	Authorization pulumi.StringMapInput
	// The name you will use to refer to this service connection in task inputs.
	Description pulumi.StringPtrInput
	// The email for Docker account user.
	DockerEmail pulumi.StringPtrInput
	// The password for the account user identified above.
	DockerPassword pulumi.StringPtrInput
	// The URL of the Docker registry. (Default: "https://index.docker.io/v1/")
	DockerRegistry pulumi.StringInput
	// The identifier of the Docker account user.
	DockerUsername pulumi.StringPtrInput
	// The project ID or project name.
	ProjectId pulumi.StringInput
	// Can be "DockerHub" or "Others" (Default "DockerHub")
	RegistryType pulumi.StringInput
	// The name you will use to refer to this service connection in task inputs.
	ServiceEndpointName pulumi.StringInput
}

func (DockerRegistryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dockerRegistryArgs)(nil)).Elem()
}

type DockerRegistryInput interface {
	pulumi.Input

	ToDockerRegistryOutput() DockerRegistryOutput
	ToDockerRegistryOutputWithContext(ctx context.Context) DockerRegistryOutput
}

func (*DockerRegistry) ElementType() reflect.Type {
	return reflect.TypeOf((*DockerRegistry)(nil))
}

func (i *DockerRegistry) ToDockerRegistryOutput() DockerRegistryOutput {
	return i.ToDockerRegistryOutputWithContext(context.Background())
}

func (i *DockerRegistry) ToDockerRegistryOutputWithContext(ctx context.Context) DockerRegistryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DockerRegistryOutput)
}

func (i *DockerRegistry) ToDockerRegistryPtrOutput() DockerRegistryPtrOutput {
	return i.ToDockerRegistryPtrOutputWithContext(context.Background())
}

func (i *DockerRegistry) ToDockerRegistryPtrOutputWithContext(ctx context.Context) DockerRegistryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DockerRegistryPtrOutput)
}

type DockerRegistryPtrInput interface {
	pulumi.Input

	ToDockerRegistryPtrOutput() DockerRegistryPtrOutput
	ToDockerRegistryPtrOutputWithContext(ctx context.Context) DockerRegistryPtrOutput
}

type dockerRegistryPtrType DockerRegistryArgs

func (*dockerRegistryPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DockerRegistry)(nil))
}

func (i *dockerRegistryPtrType) ToDockerRegistryPtrOutput() DockerRegistryPtrOutput {
	return i.ToDockerRegistryPtrOutputWithContext(context.Background())
}

func (i *dockerRegistryPtrType) ToDockerRegistryPtrOutputWithContext(ctx context.Context) DockerRegistryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DockerRegistryPtrOutput)
}

// DockerRegistryArrayInput is an input type that accepts DockerRegistryArray and DockerRegistryArrayOutput values.
// You can construct a concrete instance of `DockerRegistryArrayInput` via:
//
//          DockerRegistryArray{ DockerRegistryArgs{...} }
type DockerRegistryArrayInput interface {
	pulumi.Input

	ToDockerRegistryArrayOutput() DockerRegistryArrayOutput
	ToDockerRegistryArrayOutputWithContext(context.Context) DockerRegistryArrayOutput
}

type DockerRegistryArray []DockerRegistryInput

func (DockerRegistryArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*DockerRegistry)(nil))
}

func (i DockerRegistryArray) ToDockerRegistryArrayOutput() DockerRegistryArrayOutput {
	return i.ToDockerRegistryArrayOutputWithContext(context.Background())
}

func (i DockerRegistryArray) ToDockerRegistryArrayOutputWithContext(ctx context.Context) DockerRegistryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DockerRegistryArrayOutput)
}

// DockerRegistryMapInput is an input type that accepts DockerRegistryMap and DockerRegistryMapOutput values.
// You can construct a concrete instance of `DockerRegistryMapInput` via:
//
//          DockerRegistryMap{ "key": DockerRegistryArgs{...} }
type DockerRegistryMapInput interface {
	pulumi.Input

	ToDockerRegistryMapOutput() DockerRegistryMapOutput
	ToDockerRegistryMapOutputWithContext(context.Context) DockerRegistryMapOutput
}

type DockerRegistryMap map[string]DockerRegistryInput

func (DockerRegistryMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*DockerRegistry)(nil))
}

func (i DockerRegistryMap) ToDockerRegistryMapOutput() DockerRegistryMapOutput {
	return i.ToDockerRegistryMapOutputWithContext(context.Background())
}

func (i DockerRegistryMap) ToDockerRegistryMapOutputWithContext(ctx context.Context) DockerRegistryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DockerRegistryMapOutput)
}

type DockerRegistryOutput struct {
	*pulumi.OutputState
}

func (DockerRegistryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DockerRegistry)(nil))
}

func (o DockerRegistryOutput) ToDockerRegistryOutput() DockerRegistryOutput {
	return o
}

func (o DockerRegistryOutput) ToDockerRegistryOutputWithContext(ctx context.Context) DockerRegistryOutput {
	return o
}

func (o DockerRegistryOutput) ToDockerRegistryPtrOutput() DockerRegistryPtrOutput {
	return o.ToDockerRegistryPtrOutputWithContext(context.Background())
}

func (o DockerRegistryOutput) ToDockerRegistryPtrOutputWithContext(ctx context.Context) DockerRegistryPtrOutput {
	return o.ApplyT(func(v DockerRegistry) *DockerRegistry {
		return &v
	}).(DockerRegistryPtrOutput)
}

type DockerRegistryPtrOutput struct {
	*pulumi.OutputState
}

func (DockerRegistryPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DockerRegistry)(nil))
}

func (o DockerRegistryPtrOutput) ToDockerRegistryPtrOutput() DockerRegistryPtrOutput {
	return o
}

func (o DockerRegistryPtrOutput) ToDockerRegistryPtrOutputWithContext(ctx context.Context) DockerRegistryPtrOutput {
	return o
}

type DockerRegistryArrayOutput struct{ *pulumi.OutputState }

func (DockerRegistryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DockerRegistry)(nil))
}

func (o DockerRegistryArrayOutput) ToDockerRegistryArrayOutput() DockerRegistryArrayOutput {
	return o
}

func (o DockerRegistryArrayOutput) ToDockerRegistryArrayOutputWithContext(ctx context.Context) DockerRegistryArrayOutput {
	return o
}

func (o DockerRegistryArrayOutput) Index(i pulumi.IntInput) DockerRegistryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DockerRegistry {
		return vs[0].([]DockerRegistry)[vs[1].(int)]
	}).(DockerRegistryOutput)
}

type DockerRegistryMapOutput struct{ *pulumi.OutputState }

func (DockerRegistryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DockerRegistry)(nil))
}

func (o DockerRegistryMapOutput) ToDockerRegistryMapOutput() DockerRegistryMapOutput {
	return o
}

func (o DockerRegistryMapOutput) ToDockerRegistryMapOutputWithContext(ctx context.Context) DockerRegistryMapOutput {
	return o
}

func (o DockerRegistryMapOutput) MapIndex(k pulumi.StringInput) DockerRegistryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) DockerRegistry {
		return vs[0].(map[string]DockerRegistry)[vs[1].(string)]
	}).(DockerRegistryOutput)
}

func init() {
	pulumi.RegisterOutputType(DockerRegistryOutput{})
	pulumi.RegisterOutputType(DockerRegistryPtrOutput{})
	pulumi.RegisterOutputType(DockerRegistryArrayOutput{})
	pulumi.RegisterOutputType(DockerRegistryMapOutput{})
}
