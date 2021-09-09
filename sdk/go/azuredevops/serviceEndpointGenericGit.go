// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a generic service endpoint within Azure DevOps, which can be used to authenticate to any external git service
// using basic authentication via a username and password. This is mostly useful for importing private git repositories.
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
// 		_, err = azuredevops.NewServiceEndpointGenericGit(ctx, "serviceendpoint", &azuredevops.ServiceEndpointGenericGitArgs{
// 			ProjectId:           project.ID(),
// 			RepositoryUrl:       pulumi.String("https://dev.azure.com/org/project/_git/repository"),
// 			Username:            pulumi.String("username"),
// 			Password:            pulumi.String("password"),
// 			ServiceEndpointName: pulumi.String("Sample Generic Git"),
// 			Description:         pulumi.String("Managed by Terraform"),
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
// Azure DevOps Service Endpoint Generic Git can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
//
// ```sh
//  $ pulumi import azuredevops:index/serviceEndpointGenericGit:ServiceEndpointGenericGit serviceendpoint 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
// ```
type ServiceEndpointGenericGit struct {
	pulumi.CustomResourceState

	Authorization pulumi.StringMapOutput `pulumi:"authorization"`
	Description   pulumi.StringPtrOutput `pulumi:"description"`
	// A value indicating whether or not to attempt accessing this git server from Azure Pipelines.
	EnablePipelinesAccess pulumi.BoolPtrOutput `pulumi:"enablePipelinesAccess"`
	// The password or token key used to authenticate to the git repository using basic authentication.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// A bcrypted hash of the attribute 'password'
	PasswordHash pulumi.StringOutput `pulumi:"passwordHash"`
	// The project ID or project name to associate with the service endpoint.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The URL of the repository associated with the service endpoint.
	RepositoryUrl pulumi.StringOutput `pulumi:"repositoryUrl"`
	// The name of the service endpoint.
	ServiceEndpointName pulumi.StringOutput `pulumi:"serviceEndpointName"`
	// The username used to authenticate to the git repository using basic authentication.
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
	// The password or token key used to authenticate to the git repository using basic authentication.
	Password *string `pulumi:"password"`
	// A bcrypted hash of the attribute 'password'
	PasswordHash *string `pulumi:"passwordHash"`
	// The project ID or project name to associate with the service endpoint.
	ProjectId *string `pulumi:"projectId"`
	// The URL of the repository associated with the service endpoint.
	RepositoryUrl *string `pulumi:"repositoryUrl"`
	// The name of the service endpoint.
	ServiceEndpointName *string `pulumi:"serviceEndpointName"`
	// The username used to authenticate to the git repository using basic authentication.
	Username *string `pulumi:"username"`
}

type ServiceEndpointGenericGitState struct {
	Authorization pulumi.StringMapInput
	Description   pulumi.StringPtrInput
	// A value indicating whether or not to attempt accessing this git server from Azure Pipelines.
	EnablePipelinesAccess pulumi.BoolPtrInput
	// The password or token key used to authenticate to the git repository using basic authentication.
	Password pulumi.StringPtrInput
	// A bcrypted hash of the attribute 'password'
	PasswordHash pulumi.StringPtrInput
	// The project ID or project name to associate with the service endpoint.
	ProjectId pulumi.StringPtrInput
	// The URL of the repository associated with the service endpoint.
	RepositoryUrl pulumi.StringPtrInput
	// The name of the service endpoint.
	ServiceEndpointName pulumi.StringPtrInput
	// The username used to authenticate to the git repository using basic authentication.
	Username pulumi.StringPtrInput
}

func (ServiceEndpointGenericGitState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceEndpointGenericGitState)(nil)).Elem()
}

type serviceEndpointGenericGitArgs struct {
	Authorization map[string]string `pulumi:"authorization"`
	Description   *string           `pulumi:"description"`
	// A value indicating whether or not to attempt accessing this git server from Azure Pipelines.
	EnablePipelinesAccess *bool `pulumi:"enablePipelinesAccess"`
	// The password or token key used to authenticate to the git repository using basic authentication.
	Password *string `pulumi:"password"`
	// The project ID or project name to associate with the service endpoint.
	ProjectId string `pulumi:"projectId"`
	// The URL of the repository associated with the service endpoint.
	RepositoryUrl string `pulumi:"repositoryUrl"`
	// The name of the service endpoint.
	ServiceEndpointName string `pulumi:"serviceEndpointName"`
	// The username used to authenticate to the git repository using basic authentication.
	Username *string `pulumi:"username"`
}

// The set of arguments for constructing a ServiceEndpointGenericGit resource.
type ServiceEndpointGenericGitArgs struct {
	Authorization pulumi.StringMapInput
	Description   pulumi.StringPtrInput
	// A value indicating whether or not to attempt accessing this git server from Azure Pipelines.
	EnablePipelinesAccess pulumi.BoolPtrInput
	// The password or token key used to authenticate to the git repository using basic authentication.
	Password pulumi.StringPtrInput
	// The project ID or project name to associate with the service endpoint.
	ProjectId pulumi.StringInput
	// The URL of the repository associated with the service endpoint.
	RepositoryUrl pulumi.StringInput
	// The name of the service endpoint.
	ServiceEndpointName pulumi.StringInput
	// The username used to authenticate to the git repository using basic authentication.
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
	return reflect.TypeOf((*ServiceEndpointGenericGit)(nil))
}

func (i *ServiceEndpointGenericGit) ToServiceEndpointGenericGitOutput() ServiceEndpointGenericGitOutput {
	return i.ToServiceEndpointGenericGitOutputWithContext(context.Background())
}

func (i *ServiceEndpointGenericGit) ToServiceEndpointGenericGitOutputWithContext(ctx context.Context) ServiceEndpointGenericGitOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointGenericGitOutput)
}

func (i *ServiceEndpointGenericGit) ToServiceEndpointGenericGitPtrOutput() ServiceEndpointGenericGitPtrOutput {
	return i.ToServiceEndpointGenericGitPtrOutputWithContext(context.Background())
}

func (i *ServiceEndpointGenericGit) ToServiceEndpointGenericGitPtrOutputWithContext(ctx context.Context) ServiceEndpointGenericGitPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointGenericGitPtrOutput)
}

type ServiceEndpointGenericGitPtrInput interface {
	pulumi.Input

	ToServiceEndpointGenericGitPtrOutput() ServiceEndpointGenericGitPtrOutput
	ToServiceEndpointGenericGitPtrOutputWithContext(ctx context.Context) ServiceEndpointGenericGitPtrOutput
}

type serviceEndpointGenericGitPtrType ServiceEndpointGenericGitArgs

func (*serviceEndpointGenericGitPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceEndpointGenericGit)(nil))
}

func (i *serviceEndpointGenericGitPtrType) ToServiceEndpointGenericGitPtrOutput() ServiceEndpointGenericGitPtrOutput {
	return i.ToServiceEndpointGenericGitPtrOutputWithContext(context.Background())
}

func (i *serviceEndpointGenericGitPtrType) ToServiceEndpointGenericGitPtrOutputWithContext(ctx context.Context) ServiceEndpointGenericGitPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointGenericGitPtrOutput)
}

// ServiceEndpointGenericGitArrayInput is an input type that accepts ServiceEndpointGenericGitArray and ServiceEndpointGenericGitArrayOutput values.
// You can construct a concrete instance of `ServiceEndpointGenericGitArrayInput` via:
//
//          ServiceEndpointGenericGitArray{ ServiceEndpointGenericGitArgs{...} }
type ServiceEndpointGenericGitArrayInput interface {
	pulumi.Input

	ToServiceEndpointGenericGitArrayOutput() ServiceEndpointGenericGitArrayOutput
	ToServiceEndpointGenericGitArrayOutputWithContext(context.Context) ServiceEndpointGenericGitArrayOutput
}

type ServiceEndpointGenericGitArray []ServiceEndpointGenericGitInput

func (ServiceEndpointGenericGitArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*ServiceEndpointGenericGit)(nil))
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
//          ServiceEndpointGenericGitMap{ "key": ServiceEndpointGenericGitArgs{...} }
type ServiceEndpointGenericGitMapInput interface {
	pulumi.Input

	ToServiceEndpointGenericGitMapOutput() ServiceEndpointGenericGitMapOutput
	ToServiceEndpointGenericGitMapOutputWithContext(context.Context) ServiceEndpointGenericGitMapOutput
}

type ServiceEndpointGenericGitMap map[string]ServiceEndpointGenericGitInput

func (ServiceEndpointGenericGitMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*ServiceEndpointGenericGit)(nil))
}

func (i ServiceEndpointGenericGitMap) ToServiceEndpointGenericGitMapOutput() ServiceEndpointGenericGitMapOutput {
	return i.ToServiceEndpointGenericGitMapOutputWithContext(context.Background())
}

func (i ServiceEndpointGenericGitMap) ToServiceEndpointGenericGitMapOutputWithContext(ctx context.Context) ServiceEndpointGenericGitMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointGenericGitMapOutput)
}

type ServiceEndpointGenericGitOutput struct {
	*pulumi.OutputState
}

func (ServiceEndpointGenericGitOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceEndpointGenericGit)(nil))
}

func (o ServiceEndpointGenericGitOutput) ToServiceEndpointGenericGitOutput() ServiceEndpointGenericGitOutput {
	return o
}

func (o ServiceEndpointGenericGitOutput) ToServiceEndpointGenericGitOutputWithContext(ctx context.Context) ServiceEndpointGenericGitOutput {
	return o
}

func (o ServiceEndpointGenericGitOutput) ToServiceEndpointGenericGitPtrOutput() ServiceEndpointGenericGitPtrOutput {
	return o.ToServiceEndpointGenericGitPtrOutputWithContext(context.Background())
}

func (o ServiceEndpointGenericGitOutput) ToServiceEndpointGenericGitPtrOutputWithContext(ctx context.Context) ServiceEndpointGenericGitPtrOutput {
	return o.ApplyT(func(v ServiceEndpointGenericGit) *ServiceEndpointGenericGit {
		return &v
	}).(ServiceEndpointGenericGitPtrOutput)
}

type ServiceEndpointGenericGitPtrOutput struct {
	*pulumi.OutputState
}

func (ServiceEndpointGenericGitPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceEndpointGenericGit)(nil))
}

func (o ServiceEndpointGenericGitPtrOutput) ToServiceEndpointGenericGitPtrOutput() ServiceEndpointGenericGitPtrOutput {
	return o
}

func (o ServiceEndpointGenericGitPtrOutput) ToServiceEndpointGenericGitPtrOutputWithContext(ctx context.Context) ServiceEndpointGenericGitPtrOutput {
	return o
}

type ServiceEndpointGenericGitArrayOutput struct{ *pulumi.OutputState }

func (ServiceEndpointGenericGitArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceEndpointGenericGit)(nil))
}

func (o ServiceEndpointGenericGitArrayOutput) ToServiceEndpointGenericGitArrayOutput() ServiceEndpointGenericGitArrayOutput {
	return o
}

func (o ServiceEndpointGenericGitArrayOutput) ToServiceEndpointGenericGitArrayOutputWithContext(ctx context.Context) ServiceEndpointGenericGitArrayOutput {
	return o
}

func (o ServiceEndpointGenericGitArrayOutput) Index(i pulumi.IntInput) ServiceEndpointGenericGitOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServiceEndpointGenericGit {
		return vs[0].([]ServiceEndpointGenericGit)[vs[1].(int)]
	}).(ServiceEndpointGenericGitOutput)
}

type ServiceEndpointGenericGitMapOutput struct{ *pulumi.OutputState }

func (ServiceEndpointGenericGitMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ServiceEndpointGenericGit)(nil))
}

func (o ServiceEndpointGenericGitMapOutput) ToServiceEndpointGenericGitMapOutput() ServiceEndpointGenericGitMapOutput {
	return o
}

func (o ServiceEndpointGenericGitMapOutput) ToServiceEndpointGenericGitMapOutputWithContext(ctx context.Context) ServiceEndpointGenericGitMapOutput {
	return o
}

func (o ServiceEndpointGenericGitMapOutput) MapIndex(k pulumi.StringInput) ServiceEndpointGenericGitOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ServiceEndpointGenericGit {
		return vs[0].(map[string]ServiceEndpointGenericGit)[vs[1].(string)]
	}).(ServiceEndpointGenericGitOutput)
}

func init() {
	pulumi.RegisterOutputType(ServiceEndpointGenericGitOutput{})
	pulumi.RegisterOutputType(ServiceEndpointGenericGitPtrOutput{})
	pulumi.RegisterOutputType(ServiceEndpointGenericGitArrayOutput{})
	pulumi.RegisterOutputType(ServiceEndpointGenericGitMapOutput{})
}
