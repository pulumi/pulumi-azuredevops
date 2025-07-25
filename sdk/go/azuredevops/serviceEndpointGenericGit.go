// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azuredevops/sdk/v3/go/azuredevops/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an Other Git service endpoint within Azure DevOps, which can be used to authenticate to any external git service
// using basic authentication via a username and password. This is mostly useful for importing private git repositories.
//
// ## Example Usage
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
//			_, err = azuredevops.NewServiceEndpointGenericGit(ctx, "example", &azuredevops.ServiceEndpointGenericGitArgs{
//				ProjectId:           example.ID(),
//				RepositoryUrl:       pulumi.String("https://dev.azure.com/org/project/_git/repository"),
//				Username:            pulumi.String("username"),
//				Password:            pulumi.String("password"),
//				ServiceEndpointName: pulumi.String("Example Generic Git"),
//				Description:         pulumi.String("Managed by Pulumi"),
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
// - [Azure DevOps Service REST API 7.0 - Agent Pools](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-7.0)
//
// ## Import
//
// Azure DevOps Other Git Service Endpoint can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
//
// ```sh
// $ pulumi import azuredevops:index/serviceEndpointGenericGit:ServiceEndpointGenericGit example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
// ```
type ServiceEndpointGenericGit struct {
	pulumi.CustomResourceState

	Authorization pulumi.StringMapOutput `pulumi:"authorization"`
	Description   pulumi.StringPtrOutput `pulumi:"description"`
	// A value indicating whether or not to attempt accessing this git server from Azure Pipelines.
	EnablePipelinesAccess pulumi.BoolPtrOutput `pulumi:"enablePipelinesAccess"`
	// The PAT or password used to authenticate to the git repository.
	//
	// > **Note** For AzureDevOps Git, PAT should be used as the password.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// The ID of the project.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The URL of the repository associated with the service endpoint.
	RepositoryUrl pulumi.StringOutput `pulumi:"repositoryUrl"`
	// The name of the service endpoint.
	ServiceEndpointName pulumi.StringOutput `pulumi:"serviceEndpointName"`
	// The username used to authenticate to the git repository.
	Username pulumi.StringPtrOutput `pulumi:"username"`
}

// NewServiceEndpointGenericGit registers a new resource with the given unique name, arguments, and options.
func NewServiceEndpointGenericGit(ctx *pulumi.Context,
	name string, args *ServiceEndpointGenericGitArgs, opts ...pulumi.ResourceOption) (*ServiceEndpointGenericGit, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.RepositoryUrl == nil {
		return nil, errors.New("invalid value for required argument 'RepositoryUrl'")
	}
	if args.ServiceEndpointName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceEndpointName'")
	}
	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ServiceEndpointGenericGit
	err := ctx.RegisterResource("azuredevops:index/serviceEndpointGenericGit:ServiceEndpointGenericGit", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceEndpointGenericGit gets an existing ServiceEndpointGenericGit resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceEndpointGenericGit(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceEndpointGenericGitState, opts ...pulumi.ResourceOption) (*ServiceEndpointGenericGit, error) {
	var resource ServiceEndpointGenericGit
	err := ctx.ReadResource("azuredevops:index/serviceEndpointGenericGit:ServiceEndpointGenericGit", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceEndpointGenericGit resources.
type serviceEndpointGenericGitState struct {
	Authorization map[string]string `pulumi:"authorization"`
	Description   *string           `pulumi:"description"`
	// A value indicating whether or not to attempt accessing this git server from Azure Pipelines.
	EnablePipelinesAccess *bool `pulumi:"enablePipelinesAccess"`
	// The PAT or password used to authenticate to the git repository.
	//
	// > **Note** For AzureDevOps Git, PAT should be used as the password.
	Password *string `pulumi:"password"`
	// The ID of the project.
	ProjectId *string `pulumi:"projectId"`
	// The URL of the repository associated with the service endpoint.
	RepositoryUrl *string `pulumi:"repositoryUrl"`
	// The name of the service endpoint.
	ServiceEndpointName *string `pulumi:"serviceEndpointName"`
	// The username used to authenticate to the git repository.
	Username *string `pulumi:"username"`
}

type ServiceEndpointGenericGitState struct {
	Authorization pulumi.StringMapInput
	Description   pulumi.StringPtrInput
	// A value indicating whether or not to attempt accessing this git server from Azure Pipelines.
	EnablePipelinesAccess pulumi.BoolPtrInput
	// The PAT or password used to authenticate to the git repository.
	//
	// > **Note** For AzureDevOps Git, PAT should be used as the password.
	Password pulumi.StringPtrInput
	// The ID of the project.
	ProjectId pulumi.StringPtrInput
	// The URL of the repository associated with the service endpoint.
	RepositoryUrl pulumi.StringPtrInput
	// The name of the service endpoint.
	ServiceEndpointName pulumi.StringPtrInput
	// The username used to authenticate to the git repository.
	Username pulumi.StringPtrInput
}

func (ServiceEndpointGenericGitState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceEndpointGenericGitState)(nil)).Elem()
}

type serviceEndpointGenericGitArgs struct {
	Description *string `pulumi:"description"`
	// A value indicating whether or not to attempt accessing this git server from Azure Pipelines.
	EnablePipelinesAccess *bool `pulumi:"enablePipelinesAccess"`
	// The PAT or password used to authenticate to the git repository.
	//
	// > **Note** For AzureDevOps Git, PAT should be used as the password.
	Password *string `pulumi:"password"`
	// The ID of the project.
	ProjectId string `pulumi:"projectId"`
	// The URL of the repository associated with the service endpoint.
	RepositoryUrl string `pulumi:"repositoryUrl"`
	// The name of the service endpoint.
	ServiceEndpointName string `pulumi:"serviceEndpointName"`
	// The username used to authenticate to the git repository.
	Username *string `pulumi:"username"`
}

// The set of arguments for constructing a ServiceEndpointGenericGit resource.
type ServiceEndpointGenericGitArgs struct {
	Description pulumi.StringPtrInput
	// A value indicating whether or not to attempt accessing this git server from Azure Pipelines.
	EnablePipelinesAccess pulumi.BoolPtrInput
	// The PAT or password used to authenticate to the git repository.
	//
	// > **Note** For AzureDevOps Git, PAT should be used as the password.
	Password pulumi.StringPtrInput
	// The ID of the project.
	ProjectId pulumi.StringInput
	// The URL of the repository associated with the service endpoint.
	RepositoryUrl pulumi.StringInput
	// The name of the service endpoint.
	ServiceEndpointName pulumi.StringInput
	// The username used to authenticate to the git repository.
	Username pulumi.StringPtrInput
}

func (ServiceEndpointGenericGitArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceEndpointGenericGitArgs)(nil)).Elem()
}

type ServiceEndpointGenericGitInput interface {
	pulumi.Input

	ToServiceEndpointGenericGitOutput() ServiceEndpointGenericGitOutput
	ToServiceEndpointGenericGitOutputWithContext(ctx context.Context) ServiceEndpointGenericGitOutput
}

func (*ServiceEndpointGenericGit) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceEndpointGenericGit)(nil)).Elem()
}

func (i *ServiceEndpointGenericGit) ToServiceEndpointGenericGitOutput() ServiceEndpointGenericGitOutput {
	return i.ToServiceEndpointGenericGitOutputWithContext(context.Background())
}

func (i *ServiceEndpointGenericGit) ToServiceEndpointGenericGitOutputWithContext(ctx context.Context) ServiceEndpointGenericGitOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointGenericGitOutput)
}

// ServiceEndpointGenericGitArrayInput is an input type that accepts ServiceEndpointGenericGitArray and ServiceEndpointGenericGitArrayOutput values.
// You can construct a concrete instance of `ServiceEndpointGenericGitArrayInput` via:
//
//	ServiceEndpointGenericGitArray{ ServiceEndpointGenericGitArgs{...} }
type ServiceEndpointGenericGitArrayInput interface {
	pulumi.Input

	ToServiceEndpointGenericGitArrayOutput() ServiceEndpointGenericGitArrayOutput
	ToServiceEndpointGenericGitArrayOutputWithContext(context.Context) ServiceEndpointGenericGitArrayOutput
}

type ServiceEndpointGenericGitArray []ServiceEndpointGenericGitInput

func (ServiceEndpointGenericGitArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceEndpointGenericGit)(nil)).Elem()
}

func (i ServiceEndpointGenericGitArray) ToServiceEndpointGenericGitArrayOutput() ServiceEndpointGenericGitArrayOutput {
	return i.ToServiceEndpointGenericGitArrayOutputWithContext(context.Background())
}

func (i ServiceEndpointGenericGitArray) ToServiceEndpointGenericGitArrayOutputWithContext(ctx context.Context) ServiceEndpointGenericGitArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointGenericGitArrayOutput)
}

// ServiceEndpointGenericGitMapInput is an input type that accepts ServiceEndpointGenericGitMap and ServiceEndpointGenericGitMapOutput values.
// You can construct a concrete instance of `ServiceEndpointGenericGitMapInput` via:
//
//	ServiceEndpointGenericGitMap{ "key": ServiceEndpointGenericGitArgs{...} }
type ServiceEndpointGenericGitMapInput interface {
	pulumi.Input

	ToServiceEndpointGenericGitMapOutput() ServiceEndpointGenericGitMapOutput
	ToServiceEndpointGenericGitMapOutputWithContext(context.Context) ServiceEndpointGenericGitMapOutput
}

type ServiceEndpointGenericGitMap map[string]ServiceEndpointGenericGitInput

func (ServiceEndpointGenericGitMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceEndpointGenericGit)(nil)).Elem()
}

func (i ServiceEndpointGenericGitMap) ToServiceEndpointGenericGitMapOutput() ServiceEndpointGenericGitMapOutput {
	return i.ToServiceEndpointGenericGitMapOutputWithContext(context.Background())
}

func (i ServiceEndpointGenericGitMap) ToServiceEndpointGenericGitMapOutputWithContext(ctx context.Context) ServiceEndpointGenericGitMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointGenericGitMapOutput)
}

type ServiceEndpointGenericGitOutput struct{ *pulumi.OutputState }

func (ServiceEndpointGenericGitOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceEndpointGenericGit)(nil)).Elem()
}

func (o ServiceEndpointGenericGitOutput) ToServiceEndpointGenericGitOutput() ServiceEndpointGenericGitOutput {
	return o
}

func (o ServiceEndpointGenericGitOutput) ToServiceEndpointGenericGitOutputWithContext(ctx context.Context) ServiceEndpointGenericGitOutput {
	return o
}

func (o ServiceEndpointGenericGitOutput) Authorization() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ServiceEndpointGenericGit) pulumi.StringMapOutput { return v.Authorization }).(pulumi.StringMapOutput)
}

func (o ServiceEndpointGenericGitOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceEndpointGenericGit) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// A value indicating whether or not to attempt accessing this git server from Azure Pipelines.
func (o ServiceEndpointGenericGitOutput) EnablePipelinesAccess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ServiceEndpointGenericGit) pulumi.BoolPtrOutput { return v.EnablePipelinesAccess }).(pulumi.BoolPtrOutput)
}

// The PAT or password used to authenticate to the git repository.
//
// > **Note** For AzureDevOps Git, PAT should be used as the password.
func (o ServiceEndpointGenericGitOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceEndpointGenericGit) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

// The ID of the project.
func (o ServiceEndpointGenericGitOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEndpointGenericGit) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The URL of the repository associated with the service endpoint.
func (o ServiceEndpointGenericGitOutput) RepositoryUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEndpointGenericGit) pulumi.StringOutput { return v.RepositoryUrl }).(pulumi.StringOutput)
}

// The name of the service endpoint.
func (o ServiceEndpointGenericGitOutput) ServiceEndpointName() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEndpointGenericGit) pulumi.StringOutput { return v.ServiceEndpointName }).(pulumi.StringOutput)
}

// The username used to authenticate to the git repository.
func (o ServiceEndpointGenericGitOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceEndpointGenericGit) pulumi.StringPtrOutput { return v.Username }).(pulumi.StringPtrOutput)
}

type ServiceEndpointGenericGitArrayOutput struct{ *pulumi.OutputState }

func (ServiceEndpointGenericGitArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceEndpointGenericGit)(nil)).Elem()
}

func (o ServiceEndpointGenericGitArrayOutput) ToServiceEndpointGenericGitArrayOutput() ServiceEndpointGenericGitArrayOutput {
	return o
}

func (o ServiceEndpointGenericGitArrayOutput) ToServiceEndpointGenericGitArrayOutputWithContext(ctx context.Context) ServiceEndpointGenericGitArrayOutput {
	return o
}

func (o ServiceEndpointGenericGitArrayOutput) Index(i pulumi.IntInput) ServiceEndpointGenericGitOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServiceEndpointGenericGit {
		return vs[0].([]*ServiceEndpointGenericGit)[vs[1].(int)]
	}).(ServiceEndpointGenericGitOutput)
}

type ServiceEndpointGenericGitMapOutput struct{ *pulumi.OutputState }

func (ServiceEndpointGenericGitMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceEndpointGenericGit)(nil)).Elem()
}

func (o ServiceEndpointGenericGitMapOutput) ToServiceEndpointGenericGitMapOutput() ServiceEndpointGenericGitMapOutput {
	return o
}

func (o ServiceEndpointGenericGitMapOutput) ToServiceEndpointGenericGitMapOutputWithContext(ctx context.Context) ServiceEndpointGenericGitMapOutput {
	return o
}

func (o ServiceEndpointGenericGitMapOutput) MapIndex(k pulumi.StringInput) ServiceEndpointGenericGitOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServiceEndpointGenericGit {
		return vs[0].(map[string]*ServiceEndpointGenericGit)[vs[1].(string)]
	}).(ServiceEndpointGenericGitOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceEndpointGenericGitInput)(nil)).Elem(), &ServiceEndpointGenericGit{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceEndpointGenericGitArrayInput)(nil)).Elem(), ServiceEndpointGenericGitArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceEndpointGenericGitMapInput)(nil)).Elem(), ServiceEndpointGenericGitMap{})
	pulumi.RegisterOutputType(ServiceEndpointGenericGitOutput{})
	pulumi.RegisterOutputType(ServiceEndpointGenericGitArrayOutput{})
	pulumi.RegisterOutputType(ServiceEndpointGenericGitMapOutput{})
}
