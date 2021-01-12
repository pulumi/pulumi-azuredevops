// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Azure DevOps Service Connection service endpoint within Azure DevOps. Allows to run downstream pipelines, monitoring their execution, collecting and consolidating artefacts produced in the delegate pipelines (yaml block `task: RunPipelines@1`). More details on Marketplace page: [RunPipelines](https://marketplace.visualstudio.com/items?itemName=CSE-DevOps.RunPipelines)
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
// 			Visibility:       pulumi.String("private"),
// 			VersionControl:   pulumi.String("Git"),
// 			WorkItemTemplate: pulumi.String("Agile"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = azuredevops.NewServiceEndpointPipeline(ctx, "serviceendpoint", &azuredevops.ServiceEndpointPipelineArgs{
// 			ProjectId:           project.ID(),
// 			ServiceEndpointName: pulumi.String("Sample Pipeline Runner"),
// 			OrganizationName:    pulumi.String("MyOrganization"),
// 			AuthPersonal: &azuredevops.ServiceEndpointPipelineAuthPersonalArgs{
// 				PersonalAccessToken: pulumi.String("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
// 			},
// 			Description: pulumi.String("Managed by Terraform"),
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
// - [Azure DevOps Service REST API 5.1 - Agent Pools](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-5.1)
//
// ## Import
//
// Azure DevOps Service Endpoint can be imported using the `project id`, `service connection id`, e.g.
//
// ```sh
//  $ pulumi import azuredevops:index/serviceEndpointPipeline:ServiceEndpointPipeline serviceendpoint projectID/00000000-0000-0000-0000-000000000000
// ```
type ServiceEndpointPipeline struct {
	pulumi.CustomResourceState

	// An `authPersonal` block as documented below. Allows connecting using a personal access token.
	AuthPersonal  ServiceEndpointPipelineAuthPersonalOutput `pulumi:"authPersonal"`
	Authorization pulumi.StringMapOutput                    `pulumi:"authorization"`
	Description   pulumi.StringPtrOutput                    `pulumi:"description"`
	// The organization name used for `Organization Url` and `Release API Url` fields.
	OrganizationName pulumi.StringOutput `pulumi:"organizationName"`
	// The project ID or project name.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The Service Endpoint name.
	ServiceEndpointName pulumi.StringOutput `pulumi:"serviceEndpointName"`
}

// NewServiceEndpointPipeline registers a new resource with the given unique name, arguments, and options.
func NewServiceEndpointPipeline(ctx *pulumi.Context,
	name string, args *ServiceEndpointPipelineArgs, opts ...pulumi.ResourceOption) (*ServiceEndpointPipeline, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AuthPersonal == nil {
		return nil, errors.New("invalid value for required argument 'AuthPersonal'")
	}
	if args.OrganizationName == nil {
		return nil, errors.New("invalid value for required argument 'OrganizationName'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.ServiceEndpointName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceEndpointName'")
	}
	var resource ServiceEndpointPipeline
	err := ctx.RegisterResource("azuredevops:index/serviceEndpointPipeline:ServiceEndpointPipeline", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceEndpointPipeline gets an existing ServiceEndpointPipeline resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceEndpointPipeline(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceEndpointPipelineState, opts ...pulumi.ResourceOption) (*ServiceEndpointPipeline, error) {
	var resource ServiceEndpointPipeline
	err := ctx.ReadResource("azuredevops:index/serviceEndpointPipeline:ServiceEndpointPipeline", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceEndpointPipeline resources.
type serviceEndpointPipelineState struct {
	// An `authPersonal` block as documented below. Allows connecting using a personal access token.
	AuthPersonal  *ServiceEndpointPipelineAuthPersonal `pulumi:"authPersonal"`
	Authorization map[string]string                    `pulumi:"authorization"`
	Description   *string                              `pulumi:"description"`
	// The organization name used for `Organization Url` and `Release API Url` fields.
	OrganizationName *string `pulumi:"organizationName"`
	// The project ID or project name.
	ProjectId *string `pulumi:"projectId"`
	// The Service Endpoint name.
	ServiceEndpointName *string `pulumi:"serviceEndpointName"`
}

type ServiceEndpointPipelineState struct {
	// An `authPersonal` block as documented below. Allows connecting using a personal access token.
	AuthPersonal  ServiceEndpointPipelineAuthPersonalPtrInput
	Authorization pulumi.StringMapInput
	Description   pulumi.StringPtrInput
	// The organization name used for `Organization Url` and `Release API Url` fields.
	OrganizationName pulumi.StringPtrInput
	// The project ID or project name.
	ProjectId pulumi.StringPtrInput
	// The Service Endpoint name.
	ServiceEndpointName pulumi.StringPtrInput
}

func (ServiceEndpointPipelineState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceEndpointPipelineState)(nil)).Elem()
}

type serviceEndpointPipelineArgs struct {
	// An `authPersonal` block as documented below. Allows connecting using a personal access token.
	AuthPersonal  ServiceEndpointPipelineAuthPersonal `pulumi:"authPersonal"`
	Authorization map[string]string                   `pulumi:"authorization"`
	Description   *string                             `pulumi:"description"`
	// The organization name used for `Organization Url` and `Release API Url` fields.
	OrganizationName string `pulumi:"organizationName"`
	// The project ID or project name.
	ProjectId string `pulumi:"projectId"`
	// The Service Endpoint name.
	ServiceEndpointName string `pulumi:"serviceEndpointName"`
}

// The set of arguments for constructing a ServiceEndpointPipeline resource.
type ServiceEndpointPipelineArgs struct {
	// An `authPersonal` block as documented below. Allows connecting using a personal access token.
	AuthPersonal  ServiceEndpointPipelineAuthPersonalInput
	Authorization pulumi.StringMapInput
	Description   pulumi.StringPtrInput
	// The organization name used for `Organization Url` and `Release API Url` fields.
	OrganizationName pulumi.StringInput
	// The project ID or project name.
	ProjectId pulumi.StringInput
	// The Service Endpoint name.
	ServiceEndpointName pulumi.StringInput
}

func (ServiceEndpointPipelineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceEndpointPipelineArgs)(nil)).Elem()
}

type ServiceEndpointPipelineInput interface {
	pulumi.Input

	ToServiceEndpointPipelineOutput() ServiceEndpointPipelineOutput
	ToServiceEndpointPipelineOutputWithContext(ctx context.Context) ServiceEndpointPipelineOutput
}

func (ServiceEndpointPipeline) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceEndpointPipeline)(nil)).Elem()
}

func (i ServiceEndpointPipeline) ToServiceEndpointPipelineOutput() ServiceEndpointPipelineOutput {
	return i.ToServiceEndpointPipelineOutputWithContext(context.Background())
}

func (i ServiceEndpointPipeline) ToServiceEndpointPipelineOutputWithContext(ctx context.Context) ServiceEndpointPipelineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointPipelineOutput)
}

type ServiceEndpointPipelineOutput struct {
	*pulumi.OutputState
}

func (ServiceEndpointPipelineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceEndpointPipelineOutput)(nil)).Elem()
}

func (o ServiceEndpointPipelineOutput) ToServiceEndpointPipelineOutput() ServiceEndpointPipelineOutput {
	return o
}

func (o ServiceEndpointPipelineOutput) ToServiceEndpointPipelineOutputWithContext(ctx context.Context) ServiceEndpointPipelineOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ServiceEndpointPipelineOutput{})
}
