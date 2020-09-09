// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Docker Registry service endpoint within Azure DevOps.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azuredevops/sdk/go/azuredevops"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		project, err := azuredevops.NewProject(ctx, "project", &azuredevops.ProjectArgs{
// 			ProjectName:      pulumi.String("Sample Project"),
// 			Visibility:       pulumi.String("private"),
// 			VersionControl:   pulumi.String("Git"),
// 			WorkItemTemplate: pulumi.String("Agile"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = azuredevops.NewDockerRegistry(ctx, "dockerhubregistry", &azuredevops.DockerRegistryArgs{
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
// 		_, err = azuredevops.NewDockerRegistry(ctx, "otherregistry", &azuredevops.DockerRegistryArgs{
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
// * [Azure DevOps Service REST API 5.1 - Service Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-5.1)
// * [Docker Registry Service Connection](https://docs.microsoft.com/en-us/azure/devops/pipelines/library/service-endpoints?view=azure-devops&tabs=yaml#sep-docreg)
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
	if args == nil || args.DockerRegistry == nil {
		return nil, errors.New("missing required argument 'DockerRegistry'")
	}
	if args == nil || args.ProjectId == nil {
		return nil, errors.New("missing required argument 'ProjectId'")
	}
	if args == nil || args.RegistryType == nil {
		return nil, errors.New("missing required argument 'RegistryType'")
	}
	if args == nil || args.ServiceEndpointName == nil {
		return nil, errors.New("missing required argument 'ServiceEndpointName'")
	}
	if args == nil {
		args = &DockerRegistryArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azuredevops:ServiceEndpoint/dockerRegistry:DockerRegistry"),
		},
	})
	opts = append(opts, aliases)
	var resource DockerRegistry
	err := ctx.RegisterResource("azuredevops:index/dockerRegistry:DockerRegistry", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azuredevops:index/dockerRegistry:DockerRegistry", name, id, state, &resource, opts...)
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