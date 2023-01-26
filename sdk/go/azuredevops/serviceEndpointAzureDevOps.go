// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an Azure DevOps service endpoint within Azure DevOps.
//
// > **Note** This resource is duplicate with `ServiceEndpointPipeline`,  will be removed in the future, use `ServiceEndpointPipeline` instead.
//
// > **Note** Prerequisite: Extension [Configurable Pipeline Runner](https://marketplace.visualstudio.com/items?itemName=CSE-DevOps.RunPipelines) has been installed for the organization.
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
//			_, err = azuredevops.NewServiceEndpointAzureDevOps(ctx, "exampleServiceEndpointAzureDevOps", &azuredevops.ServiceEndpointAzureDevOpsArgs{
//				ProjectId:           exampleProject.ID(),
//				ServiceEndpointName: pulumi.String("Example Azure DevOps"),
//				OrgUrl:              pulumi.String("https://dev.azure.com/testorganization"),
//				ReleaseApiUrl:       pulumi.String("https://vsrm.dev.azure.com/testorganization"),
//				PersonalAccessToken: pulumi.String("0000000000000000000000000000000000000000000000000000"),
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
// - [Azure DevOps Service REST API 6.0 - Service Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-6.0)
//
// ## Import
//
// Azure DevOps Service Endpoint Azure DevOps can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
//
// ```sh
//
//	$ pulumi import azuredevops:index/serviceEndpointAzureDevOps:ServiceEndpointAzureDevOps example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
//
// ```
type ServiceEndpointAzureDevOps struct {
	pulumi.CustomResourceState

	Authorization pulumi.StringMapOutput `pulumi:"authorization"`
	Description   pulumi.StringPtrOutput `pulumi:"description"`
	// The organization URL.
	OrgUrl pulumi.StringOutput `pulumi:"orgUrl"`
	// The Azure DevOps personal access token.
	PersonalAccessToken pulumi.StringOutput `pulumi:"personalAccessToken"`
	// A bcrypted hash of the attribute 'personal_access_token'
	PersonalAccessTokenHash pulumi.StringOutput `pulumi:"personalAccessTokenHash"`
	// The ID of the project.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The URL of the release API.
	ReleaseApiUrl pulumi.StringOutput `pulumi:"releaseApiUrl"`
	// The Service Endpoint name.
	ServiceEndpointName pulumi.StringOutput `pulumi:"serviceEndpointName"`
}

// NewServiceEndpointAzureDevOps registers a new resource with the given unique name, arguments, and options.
func NewServiceEndpointAzureDevOps(ctx *pulumi.Context,
	name string, args *ServiceEndpointAzureDevOpsArgs, opts ...pulumi.ResourceOption) (*ServiceEndpointAzureDevOps, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.OrgUrl == nil {
		return nil, errors.New("invalid value for required argument 'OrgUrl'")
	}
	if args.PersonalAccessToken == nil {
		return nil, errors.New("invalid value for required argument 'PersonalAccessToken'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.ReleaseApiUrl == nil {
		return nil, errors.New("invalid value for required argument 'ReleaseApiUrl'")
	}
	if args.ServiceEndpointName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceEndpointName'")
	}
	var resource ServiceEndpointAzureDevOps
	err := ctx.RegisterResource("azuredevops:index/serviceEndpointAzureDevOps:ServiceEndpointAzureDevOps", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceEndpointAzureDevOps gets an existing ServiceEndpointAzureDevOps resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceEndpointAzureDevOps(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceEndpointAzureDevOpsState, opts ...pulumi.ResourceOption) (*ServiceEndpointAzureDevOps, error) {
	var resource ServiceEndpointAzureDevOps
	err := ctx.ReadResource("azuredevops:index/serviceEndpointAzureDevOps:ServiceEndpointAzureDevOps", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceEndpointAzureDevOps resources.
type serviceEndpointAzureDevOpsState struct {
	Authorization map[string]string `pulumi:"authorization"`
	Description   *string           `pulumi:"description"`
	// The organization URL.
	OrgUrl *string `pulumi:"orgUrl"`
	// The Azure DevOps personal access token.
	PersonalAccessToken *string `pulumi:"personalAccessToken"`
	// A bcrypted hash of the attribute 'personal_access_token'
	PersonalAccessTokenHash *string `pulumi:"personalAccessTokenHash"`
	// The ID of the project.
	ProjectId *string `pulumi:"projectId"`
	// The URL of the release API.
	ReleaseApiUrl *string `pulumi:"releaseApiUrl"`
	// The Service Endpoint name.
	ServiceEndpointName *string `pulumi:"serviceEndpointName"`
}

type ServiceEndpointAzureDevOpsState struct {
	Authorization pulumi.StringMapInput
	Description   pulumi.StringPtrInput
	// The organization URL.
	OrgUrl pulumi.StringPtrInput
	// The Azure DevOps personal access token.
	PersonalAccessToken pulumi.StringPtrInput
	// A bcrypted hash of the attribute 'personal_access_token'
	PersonalAccessTokenHash pulumi.StringPtrInput
	// The ID of the project.
	ProjectId pulumi.StringPtrInput
	// The URL of the release API.
	ReleaseApiUrl pulumi.StringPtrInput
	// The Service Endpoint name.
	ServiceEndpointName pulumi.StringPtrInput
}

func (ServiceEndpointAzureDevOpsState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceEndpointAzureDevOpsState)(nil)).Elem()
}

type serviceEndpointAzureDevOpsArgs struct {
	Authorization map[string]string `pulumi:"authorization"`
	Description   *string           `pulumi:"description"`
	// The organization URL.
	OrgUrl string `pulumi:"orgUrl"`
	// The Azure DevOps personal access token.
	PersonalAccessToken string `pulumi:"personalAccessToken"`
	// The ID of the project.
	ProjectId string `pulumi:"projectId"`
	// The URL of the release API.
	ReleaseApiUrl string `pulumi:"releaseApiUrl"`
	// The Service Endpoint name.
	ServiceEndpointName string `pulumi:"serviceEndpointName"`
}

// The set of arguments for constructing a ServiceEndpointAzureDevOps resource.
type ServiceEndpointAzureDevOpsArgs struct {
	Authorization pulumi.StringMapInput
	Description   pulumi.StringPtrInput
	// The organization URL.
	OrgUrl pulumi.StringInput
	// The Azure DevOps personal access token.
	PersonalAccessToken pulumi.StringInput
	// The ID of the project.
	ProjectId pulumi.StringInput
	// The URL of the release API.
	ReleaseApiUrl pulumi.StringInput
	// The Service Endpoint name.
	ServiceEndpointName pulumi.StringInput
}

func (ServiceEndpointAzureDevOpsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceEndpointAzureDevOpsArgs)(nil)).Elem()
}

type ServiceEndpointAzureDevOpsInput interface {
	pulumi.Input

	ToServiceEndpointAzureDevOpsOutput() ServiceEndpointAzureDevOpsOutput
	ToServiceEndpointAzureDevOpsOutputWithContext(ctx context.Context) ServiceEndpointAzureDevOpsOutput
}

func (*ServiceEndpointAzureDevOps) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceEndpointAzureDevOps)(nil)).Elem()
}

func (i *ServiceEndpointAzureDevOps) ToServiceEndpointAzureDevOpsOutput() ServiceEndpointAzureDevOpsOutput {
	return i.ToServiceEndpointAzureDevOpsOutputWithContext(context.Background())
}

func (i *ServiceEndpointAzureDevOps) ToServiceEndpointAzureDevOpsOutputWithContext(ctx context.Context) ServiceEndpointAzureDevOpsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointAzureDevOpsOutput)
}

// ServiceEndpointAzureDevOpsArrayInput is an input type that accepts ServiceEndpointAzureDevOpsArray and ServiceEndpointAzureDevOpsArrayOutput values.
// You can construct a concrete instance of `ServiceEndpointAzureDevOpsArrayInput` via:
//
//	ServiceEndpointAzureDevOpsArray{ ServiceEndpointAzureDevOpsArgs{...} }
type ServiceEndpointAzureDevOpsArrayInput interface {
	pulumi.Input

	ToServiceEndpointAzureDevOpsArrayOutput() ServiceEndpointAzureDevOpsArrayOutput
	ToServiceEndpointAzureDevOpsArrayOutputWithContext(context.Context) ServiceEndpointAzureDevOpsArrayOutput
}

type ServiceEndpointAzureDevOpsArray []ServiceEndpointAzureDevOpsInput

func (ServiceEndpointAzureDevOpsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceEndpointAzureDevOps)(nil)).Elem()
}

func (i ServiceEndpointAzureDevOpsArray) ToServiceEndpointAzureDevOpsArrayOutput() ServiceEndpointAzureDevOpsArrayOutput {
	return i.ToServiceEndpointAzureDevOpsArrayOutputWithContext(context.Background())
}

func (i ServiceEndpointAzureDevOpsArray) ToServiceEndpointAzureDevOpsArrayOutputWithContext(ctx context.Context) ServiceEndpointAzureDevOpsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointAzureDevOpsArrayOutput)
}

// ServiceEndpointAzureDevOpsMapInput is an input type that accepts ServiceEndpointAzureDevOpsMap and ServiceEndpointAzureDevOpsMapOutput values.
// You can construct a concrete instance of `ServiceEndpointAzureDevOpsMapInput` via:
//
//	ServiceEndpointAzureDevOpsMap{ "key": ServiceEndpointAzureDevOpsArgs{...} }
type ServiceEndpointAzureDevOpsMapInput interface {
	pulumi.Input

	ToServiceEndpointAzureDevOpsMapOutput() ServiceEndpointAzureDevOpsMapOutput
	ToServiceEndpointAzureDevOpsMapOutputWithContext(context.Context) ServiceEndpointAzureDevOpsMapOutput
}

type ServiceEndpointAzureDevOpsMap map[string]ServiceEndpointAzureDevOpsInput

func (ServiceEndpointAzureDevOpsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceEndpointAzureDevOps)(nil)).Elem()
}

func (i ServiceEndpointAzureDevOpsMap) ToServiceEndpointAzureDevOpsMapOutput() ServiceEndpointAzureDevOpsMapOutput {
	return i.ToServiceEndpointAzureDevOpsMapOutputWithContext(context.Background())
}

func (i ServiceEndpointAzureDevOpsMap) ToServiceEndpointAzureDevOpsMapOutputWithContext(ctx context.Context) ServiceEndpointAzureDevOpsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointAzureDevOpsMapOutput)
}

type ServiceEndpointAzureDevOpsOutput struct{ *pulumi.OutputState }

func (ServiceEndpointAzureDevOpsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceEndpointAzureDevOps)(nil)).Elem()
}

func (o ServiceEndpointAzureDevOpsOutput) ToServiceEndpointAzureDevOpsOutput() ServiceEndpointAzureDevOpsOutput {
	return o
}

func (o ServiceEndpointAzureDevOpsOutput) ToServiceEndpointAzureDevOpsOutputWithContext(ctx context.Context) ServiceEndpointAzureDevOpsOutput {
	return o
}

func (o ServiceEndpointAzureDevOpsOutput) Authorization() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ServiceEndpointAzureDevOps) pulumi.StringMapOutput { return v.Authorization }).(pulumi.StringMapOutput)
}

func (o ServiceEndpointAzureDevOpsOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceEndpointAzureDevOps) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The organization URL.
func (o ServiceEndpointAzureDevOpsOutput) OrgUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEndpointAzureDevOps) pulumi.StringOutput { return v.OrgUrl }).(pulumi.StringOutput)
}

// The Azure DevOps personal access token.
func (o ServiceEndpointAzureDevOpsOutput) PersonalAccessToken() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEndpointAzureDevOps) pulumi.StringOutput { return v.PersonalAccessToken }).(pulumi.StringOutput)
}

// A bcrypted hash of the attribute 'personal_access_token'
func (o ServiceEndpointAzureDevOpsOutput) PersonalAccessTokenHash() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEndpointAzureDevOps) pulumi.StringOutput { return v.PersonalAccessTokenHash }).(pulumi.StringOutput)
}

// The ID of the project.
func (o ServiceEndpointAzureDevOpsOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEndpointAzureDevOps) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The URL of the release API.
func (o ServiceEndpointAzureDevOpsOutput) ReleaseApiUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEndpointAzureDevOps) pulumi.StringOutput { return v.ReleaseApiUrl }).(pulumi.StringOutput)
}

// The Service Endpoint name.
func (o ServiceEndpointAzureDevOpsOutput) ServiceEndpointName() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEndpointAzureDevOps) pulumi.StringOutput { return v.ServiceEndpointName }).(pulumi.StringOutput)
}

type ServiceEndpointAzureDevOpsArrayOutput struct{ *pulumi.OutputState }

func (ServiceEndpointAzureDevOpsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceEndpointAzureDevOps)(nil)).Elem()
}

func (o ServiceEndpointAzureDevOpsArrayOutput) ToServiceEndpointAzureDevOpsArrayOutput() ServiceEndpointAzureDevOpsArrayOutput {
	return o
}

func (o ServiceEndpointAzureDevOpsArrayOutput) ToServiceEndpointAzureDevOpsArrayOutputWithContext(ctx context.Context) ServiceEndpointAzureDevOpsArrayOutput {
	return o
}

func (o ServiceEndpointAzureDevOpsArrayOutput) Index(i pulumi.IntInput) ServiceEndpointAzureDevOpsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServiceEndpointAzureDevOps {
		return vs[0].([]*ServiceEndpointAzureDevOps)[vs[1].(int)]
	}).(ServiceEndpointAzureDevOpsOutput)
}

type ServiceEndpointAzureDevOpsMapOutput struct{ *pulumi.OutputState }

func (ServiceEndpointAzureDevOpsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceEndpointAzureDevOps)(nil)).Elem()
}

func (o ServiceEndpointAzureDevOpsMapOutput) ToServiceEndpointAzureDevOpsMapOutput() ServiceEndpointAzureDevOpsMapOutput {
	return o
}

func (o ServiceEndpointAzureDevOpsMapOutput) ToServiceEndpointAzureDevOpsMapOutputWithContext(ctx context.Context) ServiceEndpointAzureDevOpsMapOutput {
	return o
}

func (o ServiceEndpointAzureDevOpsMapOutput) MapIndex(k pulumi.StringInput) ServiceEndpointAzureDevOpsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServiceEndpointAzureDevOps {
		return vs[0].(map[string]*ServiceEndpointAzureDevOps)[vs[1].(string)]
	}).(ServiceEndpointAzureDevOpsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceEndpointAzureDevOpsInput)(nil)).Elem(), &ServiceEndpointAzureDevOps{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceEndpointAzureDevOpsArrayInput)(nil)).Elem(), ServiceEndpointAzureDevOpsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceEndpointAzureDevOpsMapInput)(nil)).Elem(), ServiceEndpointAzureDevOpsMap{})
	pulumi.RegisterOutputType(ServiceEndpointAzureDevOpsOutput{})
	pulumi.RegisterOutputType(ServiceEndpointAzureDevOpsArrayOutput{})
	pulumi.RegisterOutputType(ServiceEndpointAzureDevOpsMapOutput{})
}
