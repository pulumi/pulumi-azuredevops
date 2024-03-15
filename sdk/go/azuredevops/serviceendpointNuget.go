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

// Manages a NuGet service endpoint within Azure DevOps.
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
//			_, err = azuredevops.NewServiceendpointNuget(ctx, "exampleServiceendpointNuget", &azuredevops.ServiceendpointNugetArgs{
//				ProjectId:           exampleProject.ID(),
//				ApiKey:              pulumi.String("apikey"),
//				ServiceEndpointName: pulumi.String("Example NuGet"),
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
// <!--End PulumiCodeChooser -->
//
// ## Relevant Links
//
// - [Azure DevOps Service REST API 7.0 - Agent Pools](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-7.0)
//
// ## Import
//
// Azure DevOps Service Endpoint NuGet can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
//
// ```sh
// $ pulumi import azuredevops:index/serviceendpointNuget:ServiceendpointNuget example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
// ```
type ServiceendpointNuget struct {
	pulumi.CustomResourceState

	// The API Key used to connect to the endpoint.
	ApiKey        pulumi.StringPtrOutput `pulumi:"apiKey"`
	Authorization pulumi.StringMapOutput `pulumi:"authorization"`
	Description   pulumi.StringPtrOutput `pulumi:"description"`
	// The URL for the feed. This will generally end with `index.json`.
	FeedUrl pulumi.StringOutput `pulumi:"feedUrl"`
	// The account password used to connect to the endpoint
	//
	// > **Note** Only one of `apiKey` or `personalAccessToken` or  `username`, `password` can be set at the same time.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// The Personal access token used to  connect to the endpoint. Personal access tokens are applicable only for NuGet feeds hosted on other Azure DevOps Services organizations or Azure DevOps Server 2019 (or later).
	PersonalAccessToken pulumi.StringPtrOutput `pulumi:"personalAccessToken"`
	// The ID of the project.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The Service Endpoint name.
	ServiceEndpointName pulumi.StringOutput `pulumi:"serviceEndpointName"`
	// The account username used to connect to the endpoint.
	Username pulumi.StringPtrOutput `pulumi:"username"`
}

// NewServiceendpointNuget registers a new resource with the given unique name, arguments, and options.
func NewServiceendpointNuget(ctx *pulumi.Context,
	name string, args *ServiceendpointNugetArgs, opts ...pulumi.ResourceOption) (*ServiceendpointNuget, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FeedUrl == nil {
		return nil, errors.New("invalid value for required argument 'FeedUrl'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.ServiceEndpointName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceEndpointName'")
	}
	if args.ApiKey != nil {
		args.ApiKey = pulumi.ToSecret(args.ApiKey).(pulumi.StringPtrInput)
	}
	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringPtrInput)
	}
	if args.PersonalAccessToken != nil {
		args.PersonalAccessToken = pulumi.ToSecret(args.PersonalAccessToken).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"apiKey",
		"password",
		"personalAccessToken",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ServiceendpointNuget
	err := ctx.RegisterResource("azuredevops:index/serviceendpointNuget:ServiceendpointNuget", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceendpointNuget gets an existing ServiceendpointNuget resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceendpointNuget(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceendpointNugetState, opts ...pulumi.ResourceOption) (*ServiceendpointNuget, error) {
	var resource ServiceendpointNuget
	err := ctx.ReadResource("azuredevops:index/serviceendpointNuget:ServiceendpointNuget", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceendpointNuget resources.
type serviceendpointNugetState struct {
	// The API Key used to connect to the endpoint.
	ApiKey        *string           `pulumi:"apiKey"`
	Authorization map[string]string `pulumi:"authorization"`
	Description   *string           `pulumi:"description"`
	// The URL for the feed. This will generally end with `index.json`.
	FeedUrl *string `pulumi:"feedUrl"`
	// The account password used to connect to the endpoint
	//
	// > **Note** Only one of `apiKey` or `personalAccessToken` or  `username`, `password` can be set at the same time.
	Password *string `pulumi:"password"`
	// The Personal access token used to  connect to the endpoint. Personal access tokens are applicable only for NuGet feeds hosted on other Azure DevOps Services organizations or Azure DevOps Server 2019 (or later).
	PersonalAccessToken *string `pulumi:"personalAccessToken"`
	// The ID of the project.
	ProjectId *string `pulumi:"projectId"`
	// The Service Endpoint name.
	ServiceEndpointName *string `pulumi:"serviceEndpointName"`
	// The account username used to connect to the endpoint.
	Username *string `pulumi:"username"`
}

type ServiceendpointNugetState struct {
	// The API Key used to connect to the endpoint.
	ApiKey        pulumi.StringPtrInput
	Authorization pulumi.StringMapInput
	Description   pulumi.StringPtrInput
	// The URL for the feed. This will generally end with `index.json`.
	FeedUrl pulumi.StringPtrInput
	// The account password used to connect to the endpoint
	//
	// > **Note** Only one of `apiKey` or `personalAccessToken` or  `username`, `password` can be set at the same time.
	Password pulumi.StringPtrInput
	// The Personal access token used to  connect to the endpoint. Personal access tokens are applicable only for NuGet feeds hosted on other Azure DevOps Services organizations or Azure DevOps Server 2019 (or later).
	PersonalAccessToken pulumi.StringPtrInput
	// The ID of the project.
	ProjectId pulumi.StringPtrInput
	// The Service Endpoint name.
	ServiceEndpointName pulumi.StringPtrInput
	// The account username used to connect to the endpoint.
	Username pulumi.StringPtrInput
}

func (ServiceendpointNugetState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceendpointNugetState)(nil)).Elem()
}

type serviceendpointNugetArgs struct {
	// The API Key used to connect to the endpoint.
	ApiKey        *string           `pulumi:"apiKey"`
	Authorization map[string]string `pulumi:"authorization"`
	Description   *string           `pulumi:"description"`
	// The URL for the feed. This will generally end with `index.json`.
	FeedUrl string `pulumi:"feedUrl"`
	// The account password used to connect to the endpoint
	//
	// > **Note** Only one of `apiKey` or `personalAccessToken` or  `username`, `password` can be set at the same time.
	Password *string `pulumi:"password"`
	// The Personal access token used to  connect to the endpoint. Personal access tokens are applicable only for NuGet feeds hosted on other Azure DevOps Services organizations or Azure DevOps Server 2019 (or later).
	PersonalAccessToken *string `pulumi:"personalAccessToken"`
	// The ID of the project.
	ProjectId string `pulumi:"projectId"`
	// The Service Endpoint name.
	ServiceEndpointName string `pulumi:"serviceEndpointName"`
	// The account username used to connect to the endpoint.
	Username *string `pulumi:"username"`
}

// The set of arguments for constructing a ServiceendpointNuget resource.
type ServiceendpointNugetArgs struct {
	// The API Key used to connect to the endpoint.
	ApiKey        pulumi.StringPtrInput
	Authorization pulumi.StringMapInput
	Description   pulumi.StringPtrInput
	// The URL for the feed. This will generally end with `index.json`.
	FeedUrl pulumi.StringInput
	// The account password used to connect to the endpoint
	//
	// > **Note** Only one of `apiKey` or `personalAccessToken` or  `username`, `password` can be set at the same time.
	Password pulumi.StringPtrInput
	// The Personal access token used to  connect to the endpoint. Personal access tokens are applicable only for NuGet feeds hosted on other Azure DevOps Services organizations or Azure DevOps Server 2019 (or later).
	PersonalAccessToken pulumi.StringPtrInput
	// The ID of the project.
	ProjectId pulumi.StringInput
	// The Service Endpoint name.
	ServiceEndpointName pulumi.StringInput
	// The account username used to connect to the endpoint.
	Username pulumi.StringPtrInput
}

func (ServiceendpointNugetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceendpointNugetArgs)(nil)).Elem()
}

type ServiceendpointNugetInput interface {
	pulumi.Input

	ToServiceendpointNugetOutput() ServiceendpointNugetOutput
	ToServiceendpointNugetOutputWithContext(ctx context.Context) ServiceendpointNugetOutput
}

func (*ServiceendpointNuget) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceendpointNuget)(nil)).Elem()
}

func (i *ServiceendpointNuget) ToServiceendpointNugetOutput() ServiceendpointNugetOutput {
	return i.ToServiceendpointNugetOutputWithContext(context.Background())
}

func (i *ServiceendpointNuget) ToServiceendpointNugetOutputWithContext(ctx context.Context) ServiceendpointNugetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceendpointNugetOutput)
}

// ServiceendpointNugetArrayInput is an input type that accepts ServiceendpointNugetArray and ServiceendpointNugetArrayOutput values.
// You can construct a concrete instance of `ServiceendpointNugetArrayInput` via:
//
//	ServiceendpointNugetArray{ ServiceendpointNugetArgs{...} }
type ServiceendpointNugetArrayInput interface {
	pulumi.Input

	ToServiceendpointNugetArrayOutput() ServiceendpointNugetArrayOutput
	ToServiceendpointNugetArrayOutputWithContext(context.Context) ServiceendpointNugetArrayOutput
}

type ServiceendpointNugetArray []ServiceendpointNugetInput

func (ServiceendpointNugetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceendpointNuget)(nil)).Elem()
}

func (i ServiceendpointNugetArray) ToServiceendpointNugetArrayOutput() ServiceendpointNugetArrayOutput {
	return i.ToServiceendpointNugetArrayOutputWithContext(context.Background())
}

func (i ServiceendpointNugetArray) ToServiceendpointNugetArrayOutputWithContext(ctx context.Context) ServiceendpointNugetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceendpointNugetArrayOutput)
}

// ServiceendpointNugetMapInput is an input type that accepts ServiceendpointNugetMap and ServiceendpointNugetMapOutput values.
// You can construct a concrete instance of `ServiceendpointNugetMapInput` via:
//
//	ServiceendpointNugetMap{ "key": ServiceendpointNugetArgs{...} }
type ServiceendpointNugetMapInput interface {
	pulumi.Input

	ToServiceendpointNugetMapOutput() ServiceendpointNugetMapOutput
	ToServiceendpointNugetMapOutputWithContext(context.Context) ServiceendpointNugetMapOutput
}

type ServiceendpointNugetMap map[string]ServiceendpointNugetInput

func (ServiceendpointNugetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceendpointNuget)(nil)).Elem()
}

func (i ServiceendpointNugetMap) ToServiceendpointNugetMapOutput() ServiceendpointNugetMapOutput {
	return i.ToServiceendpointNugetMapOutputWithContext(context.Background())
}

func (i ServiceendpointNugetMap) ToServiceendpointNugetMapOutputWithContext(ctx context.Context) ServiceendpointNugetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceendpointNugetMapOutput)
}

type ServiceendpointNugetOutput struct{ *pulumi.OutputState }

func (ServiceendpointNugetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceendpointNuget)(nil)).Elem()
}

func (o ServiceendpointNugetOutput) ToServiceendpointNugetOutput() ServiceendpointNugetOutput {
	return o
}

func (o ServiceendpointNugetOutput) ToServiceendpointNugetOutputWithContext(ctx context.Context) ServiceendpointNugetOutput {
	return o
}

// The API Key used to connect to the endpoint.
func (o ServiceendpointNugetOutput) ApiKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceendpointNuget) pulumi.StringPtrOutput { return v.ApiKey }).(pulumi.StringPtrOutput)
}

func (o ServiceendpointNugetOutput) Authorization() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ServiceendpointNuget) pulumi.StringMapOutput { return v.Authorization }).(pulumi.StringMapOutput)
}

func (o ServiceendpointNugetOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceendpointNuget) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The URL for the feed. This will generally end with `index.json`.
func (o ServiceendpointNugetOutput) FeedUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceendpointNuget) pulumi.StringOutput { return v.FeedUrl }).(pulumi.StringOutput)
}

// The account password used to connect to the endpoint
//
// > **Note** Only one of `apiKey` or `personalAccessToken` or  `username`, `password` can be set at the same time.
func (o ServiceendpointNugetOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceendpointNuget) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

// The Personal access token used to  connect to the endpoint. Personal access tokens are applicable only for NuGet feeds hosted on other Azure DevOps Services organizations or Azure DevOps Server 2019 (or later).
func (o ServiceendpointNugetOutput) PersonalAccessToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceendpointNuget) pulumi.StringPtrOutput { return v.PersonalAccessToken }).(pulumi.StringPtrOutput)
}

// The ID of the project.
func (o ServiceendpointNugetOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceendpointNuget) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The Service Endpoint name.
func (o ServiceendpointNugetOutput) ServiceEndpointName() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceendpointNuget) pulumi.StringOutput { return v.ServiceEndpointName }).(pulumi.StringOutput)
}

// The account username used to connect to the endpoint.
func (o ServiceendpointNugetOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceendpointNuget) pulumi.StringPtrOutput { return v.Username }).(pulumi.StringPtrOutput)
}

type ServiceendpointNugetArrayOutput struct{ *pulumi.OutputState }

func (ServiceendpointNugetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceendpointNuget)(nil)).Elem()
}

func (o ServiceendpointNugetArrayOutput) ToServiceendpointNugetArrayOutput() ServiceendpointNugetArrayOutput {
	return o
}

func (o ServiceendpointNugetArrayOutput) ToServiceendpointNugetArrayOutputWithContext(ctx context.Context) ServiceendpointNugetArrayOutput {
	return o
}

func (o ServiceendpointNugetArrayOutput) Index(i pulumi.IntInput) ServiceendpointNugetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServiceendpointNuget {
		return vs[0].([]*ServiceendpointNuget)[vs[1].(int)]
	}).(ServiceendpointNugetOutput)
}

type ServiceendpointNugetMapOutput struct{ *pulumi.OutputState }

func (ServiceendpointNugetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceendpointNuget)(nil)).Elem()
}

func (o ServiceendpointNugetMapOutput) ToServiceendpointNugetMapOutput() ServiceendpointNugetMapOutput {
	return o
}

func (o ServiceendpointNugetMapOutput) ToServiceendpointNugetMapOutputWithContext(ctx context.Context) ServiceendpointNugetMapOutput {
	return o
}

func (o ServiceendpointNugetMapOutput) MapIndex(k pulumi.StringInput) ServiceendpointNugetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServiceendpointNuget {
		return vs[0].(map[string]*ServiceendpointNuget)[vs[1].(string)]
	}).(ServiceendpointNugetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceendpointNugetInput)(nil)).Elem(), &ServiceendpointNuget{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceendpointNugetArrayInput)(nil)).Elem(), ServiceendpointNugetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceendpointNugetMapInput)(nil)).Elem(), ServiceendpointNugetMap{})
	pulumi.RegisterOutputType(ServiceendpointNugetOutput{})
	pulumi.RegisterOutputType(ServiceendpointNugetArrayOutput{})
	pulumi.RegisterOutputType(ServiceendpointNugetMapOutput{})
}
