// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an Artifactory server endpoint within an Azure DevOps organization.
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
// 		_, err = azuredevops.NewServiceEndpointArtifactory(ctx, "serviceendpoint", &azuredevops.ServiceEndpointArtifactoryArgs{
// 			ProjectId:           project.ID(),
// 			ServiceEndpointName: pulumi.String("Sample Artifactory"),
// 			Description:         pulumi.String("Managed by Terraform"),
// 			Url:                 pulumi.String("https://artifactory.my.com"),
// 			AuthenticationToken: &azuredevops.ServiceEndpointArtifactoryAuthenticationTokenArgs{
// 				Token: pulumi.String("0000000000000000000000000000000000000000"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// Alternatively a username and password may be used.
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
// 		_, err := azuredevops.NewServiceEndpointArtifactory(ctx, "serviceendpoint", &azuredevops.ServiceEndpointArtifactoryArgs{
// 			ProjectId:           pulumi.Any(azuredevops_project.Project.Id),
// 			ServiceEndpointName: pulumi.String("Sample Artifactory"),
// 			Description:         pulumi.String("Managed by Terraform"),
// 			Url:                 pulumi.String("https://artifactory.my.com"),
// 			AuthenticationBasic: &azuredevops.ServiceEndpointArtifactoryAuthenticationBasicArgs{
// 				Username: pulumi.String("sampleuser"),
// 				Password: pulumi.String("0000000000000000000000000000000000000000"),
// 			},
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
// * [Azure DevOps Service Connections](https://docs.microsoft.com/en-us/azure/devops/pipelines/library/service-endpoints?view=azure-devops&tabs=yaml)
// * [Artifactory User Token](https://docs.artifactory.org/latest/user-guide/user-token/)
//
// ## Import
//
// Azure DevOps Service Endpoint Artifactory can be imported using the **projectID/serviceEndpointID**, e.g.
//
// ```sh
//  $ pulumi import azuredevops:index/serviceEndpointArtifactory:ServiceEndpointArtifactory serviceendpoint 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
// ```
type ServiceEndpointArtifactory struct {
	pulumi.CustomResourceState

	AuthenticationBasic ServiceEndpointArtifactoryAuthenticationBasicPtrOutput `pulumi:"authenticationBasic"`
	AuthenticationToken ServiceEndpointArtifactoryAuthenticationTokenPtrOutput `pulumi:"authenticationToken"`
	Authorization       pulumi.StringMapOutput                                 `pulumi:"authorization"`
	// The Service Endpoint description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The project ID or project name.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The Service Endpoint name.
	ServiceEndpointName pulumi.StringOutput `pulumi:"serviceEndpointName"`
	// URL of the Artifactory server to connect with.
	Url pulumi.StringOutput `pulumi:"url"`
}

// NewServiceEndpointArtifactory registers a new resource with the given unique name, arguments, and options.
func NewServiceEndpointArtifactory(ctx *pulumi.Context,
	name string, args *ServiceEndpointArtifactoryArgs, opts ...pulumi.ResourceOption) (*ServiceEndpointArtifactory, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.ServiceEndpointName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceEndpointName'")
	}
	if args.Url == nil {
		return nil, errors.New("invalid value for required argument 'Url'")
	}
	var resource ServiceEndpointArtifactory
	err := ctx.RegisterResource("azuredevops:index/serviceEndpointArtifactory:ServiceEndpointArtifactory", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceEndpointArtifactory gets an existing ServiceEndpointArtifactory resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceEndpointArtifactory(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceEndpointArtifactoryState, opts ...pulumi.ResourceOption) (*ServiceEndpointArtifactory, error) {
	var resource ServiceEndpointArtifactory
	err := ctx.ReadResource("azuredevops:index/serviceEndpointArtifactory:ServiceEndpointArtifactory", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceEndpointArtifactory resources.
type serviceEndpointArtifactoryState struct {
	AuthenticationBasic *ServiceEndpointArtifactoryAuthenticationBasic `pulumi:"authenticationBasic"`
	AuthenticationToken *ServiceEndpointArtifactoryAuthenticationToken `pulumi:"authenticationToken"`
	Authorization       map[string]string                              `pulumi:"authorization"`
	// The Service Endpoint description.
	Description *string `pulumi:"description"`
	// The project ID or project name.
	ProjectId *string `pulumi:"projectId"`
	// The Service Endpoint name.
	ServiceEndpointName *string `pulumi:"serviceEndpointName"`
	// URL of the Artifactory server to connect with.
	Url *string `pulumi:"url"`
}

type ServiceEndpointArtifactoryState struct {
	AuthenticationBasic ServiceEndpointArtifactoryAuthenticationBasicPtrInput
	AuthenticationToken ServiceEndpointArtifactoryAuthenticationTokenPtrInput
	Authorization       pulumi.StringMapInput
	// The Service Endpoint description.
	Description pulumi.StringPtrInput
	// The project ID or project name.
	ProjectId pulumi.StringPtrInput
	// The Service Endpoint name.
	ServiceEndpointName pulumi.StringPtrInput
	// URL of the Artifactory server to connect with.
	Url pulumi.StringPtrInput
}

func (ServiceEndpointArtifactoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceEndpointArtifactoryState)(nil)).Elem()
}

type serviceEndpointArtifactoryArgs struct {
	AuthenticationBasic *ServiceEndpointArtifactoryAuthenticationBasic `pulumi:"authenticationBasic"`
	AuthenticationToken *ServiceEndpointArtifactoryAuthenticationToken `pulumi:"authenticationToken"`
	Authorization       map[string]string                              `pulumi:"authorization"`
	// The Service Endpoint description.
	Description *string `pulumi:"description"`
	// The project ID or project name.
	ProjectId string `pulumi:"projectId"`
	// The Service Endpoint name.
	ServiceEndpointName string `pulumi:"serviceEndpointName"`
	// URL of the Artifactory server to connect with.
	Url string `pulumi:"url"`
}

// The set of arguments for constructing a ServiceEndpointArtifactory resource.
type ServiceEndpointArtifactoryArgs struct {
	AuthenticationBasic ServiceEndpointArtifactoryAuthenticationBasicPtrInput
	AuthenticationToken ServiceEndpointArtifactoryAuthenticationTokenPtrInput
	Authorization       pulumi.StringMapInput
	// The Service Endpoint description.
	Description pulumi.StringPtrInput
	// The project ID or project name.
	ProjectId pulumi.StringInput
	// The Service Endpoint name.
	ServiceEndpointName pulumi.StringInput
	// URL of the Artifactory server to connect with.
	Url pulumi.StringInput
}

func (ServiceEndpointArtifactoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceEndpointArtifactoryArgs)(nil)).Elem()
}

type ServiceEndpointArtifactoryInput interface {
	pulumi.Input

	ToServiceEndpointArtifactoryOutput() ServiceEndpointArtifactoryOutput
	ToServiceEndpointArtifactoryOutputWithContext(ctx context.Context) ServiceEndpointArtifactoryOutput
}

func (*ServiceEndpointArtifactory) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceEndpointArtifactory)(nil))
}

func (i *ServiceEndpointArtifactory) ToServiceEndpointArtifactoryOutput() ServiceEndpointArtifactoryOutput {
	return i.ToServiceEndpointArtifactoryOutputWithContext(context.Background())
}

func (i *ServiceEndpointArtifactory) ToServiceEndpointArtifactoryOutputWithContext(ctx context.Context) ServiceEndpointArtifactoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointArtifactoryOutput)
}

func (i *ServiceEndpointArtifactory) ToServiceEndpointArtifactoryPtrOutput() ServiceEndpointArtifactoryPtrOutput {
	return i.ToServiceEndpointArtifactoryPtrOutputWithContext(context.Background())
}

func (i *ServiceEndpointArtifactory) ToServiceEndpointArtifactoryPtrOutputWithContext(ctx context.Context) ServiceEndpointArtifactoryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointArtifactoryPtrOutput)
}

type ServiceEndpointArtifactoryPtrInput interface {
	pulumi.Input

	ToServiceEndpointArtifactoryPtrOutput() ServiceEndpointArtifactoryPtrOutput
	ToServiceEndpointArtifactoryPtrOutputWithContext(ctx context.Context) ServiceEndpointArtifactoryPtrOutput
}

type serviceEndpointArtifactoryPtrType ServiceEndpointArtifactoryArgs

func (*serviceEndpointArtifactoryPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceEndpointArtifactory)(nil))
}

func (i *serviceEndpointArtifactoryPtrType) ToServiceEndpointArtifactoryPtrOutput() ServiceEndpointArtifactoryPtrOutput {
	return i.ToServiceEndpointArtifactoryPtrOutputWithContext(context.Background())
}

func (i *serviceEndpointArtifactoryPtrType) ToServiceEndpointArtifactoryPtrOutputWithContext(ctx context.Context) ServiceEndpointArtifactoryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointArtifactoryPtrOutput)
}

// ServiceEndpointArtifactoryArrayInput is an input type that accepts ServiceEndpointArtifactoryArray and ServiceEndpointArtifactoryArrayOutput values.
// You can construct a concrete instance of `ServiceEndpointArtifactoryArrayInput` via:
//
//          ServiceEndpointArtifactoryArray{ ServiceEndpointArtifactoryArgs{...} }
type ServiceEndpointArtifactoryArrayInput interface {
	pulumi.Input

	ToServiceEndpointArtifactoryArrayOutput() ServiceEndpointArtifactoryArrayOutput
	ToServiceEndpointArtifactoryArrayOutputWithContext(context.Context) ServiceEndpointArtifactoryArrayOutput
}

type ServiceEndpointArtifactoryArray []ServiceEndpointArtifactoryInput

func (ServiceEndpointArtifactoryArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*ServiceEndpointArtifactory)(nil))
}

func (i ServiceEndpointArtifactoryArray) ToServiceEndpointArtifactoryArrayOutput() ServiceEndpointArtifactoryArrayOutput {
	return i.ToServiceEndpointArtifactoryArrayOutputWithContext(context.Background())
}

func (i ServiceEndpointArtifactoryArray) ToServiceEndpointArtifactoryArrayOutputWithContext(ctx context.Context) ServiceEndpointArtifactoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointArtifactoryArrayOutput)
}

// ServiceEndpointArtifactoryMapInput is an input type that accepts ServiceEndpointArtifactoryMap and ServiceEndpointArtifactoryMapOutput values.
// You can construct a concrete instance of `ServiceEndpointArtifactoryMapInput` via:
//
//          ServiceEndpointArtifactoryMap{ "key": ServiceEndpointArtifactoryArgs{...} }
type ServiceEndpointArtifactoryMapInput interface {
	pulumi.Input

	ToServiceEndpointArtifactoryMapOutput() ServiceEndpointArtifactoryMapOutput
	ToServiceEndpointArtifactoryMapOutputWithContext(context.Context) ServiceEndpointArtifactoryMapOutput
}

type ServiceEndpointArtifactoryMap map[string]ServiceEndpointArtifactoryInput

func (ServiceEndpointArtifactoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*ServiceEndpointArtifactory)(nil))
}

func (i ServiceEndpointArtifactoryMap) ToServiceEndpointArtifactoryMapOutput() ServiceEndpointArtifactoryMapOutput {
	return i.ToServiceEndpointArtifactoryMapOutputWithContext(context.Background())
}

func (i ServiceEndpointArtifactoryMap) ToServiceEndpointArtifactoryMapOutputWithContext(ctx context.Context) ServiceEndpointArtifactoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointArtifactoryMapOutput)
}

type ServiceEndpointArtifactoryOutput struct {
	*pulumi.OutputState
}

func (ServiceEndpointArtifactoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceEndpointArtifactory)(nil))
}

func (o ServiceEndpointArtifactoryOutput) ToServiceEndpointArtifactoryOutput() ServiceEndpointArtifactoryOutput {
	return o
}

func (o ServiceEndpointArtifactoryOutput) ToServiceEndpointArtifactoryOutputWithContext(ctx context.Context) ServiceEndpointArtifactoryOutput {
	return o
}

func (o ServiceEndpointArtifactoryOutput) ToServiceEndpointArtifactoryPtrOutput() ServiceEndpointArtifactoryPtrOutput {
	return o.ToServiceEndpointArtifactoryPtrOutputWithContext(context.Background())
}

func (o ServiceEndpointArtifactoryOutput) ToServiceEndpointArtifactoryPtrOutputWithContext(ctx context.Context) ServiceEndpointArtifactoryPtrOutput {
	return o.ApplyT(func(v ServiceEndpointArtifactory) *ServiceEndpointArtifactory {
		return &v
	}).(ServiceEndpointArtifactoryPtrOutput)
}

type ServiceEndpointArtifactoryPtrOutput struct {
	*pulumi.OutputState
}

func (ServiceEndpointArtifactoryPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceEndpointArtifactory)(nil))
}

func (o ServiceEndpointArtifactoryPtrOutput) ToServiceEndpointArtifactoryPtrOutput() ServiceEndpointArtifactoryPtrOutput {
	return o
}

func (o ServiceEndpointArtifactoryPtrOutput) ToServiceEndpointArtifactoryPtrOutputWithContext(ctx context.Context) ServiceEndpointArtifactoryPtrOutput {
	return o
}

type ServiceEndpointArtifactoryArrayOutput struct{ *pulumi.OutputState }

func (ServiceEndpointArtifactoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceEndpointArtifactory)(nil))
}

func (o ServiceEndpointArtifactoryArrayOutput) ToServiceEndpointArtifactoryArrayOutput() ServiceEndpointArtifactoryArrayOutput {
	return o
}

func (o ServiceEndpointArtifactoryArrayOutput) ToServiceEndpointArtifactoryArrayOutputWithContext(ctx context.Context) ServiceEndpointArtifactoryArrayOutput {
	return o
}

func (o ServiceEndpointArtifactoryArrayOutput) Index(i pulumi.IntInput) ServiceEndpointArtifactoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServiceEndpointArtifactory {
		return vs[0].([]ServiceEndpointArtifactory)[vs[1].(int)]
	}).(ServiceEndpointArtifactoryOutput)
}

type ServiceEndpointArtifactoryMapOutput struct{ *pulumi.OutputState }

func (ServiceEndpointArtifactoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ServiceEndpointArtifactory)(nil))
}

func (o ServiceEndpointArtifactoryMapOutput) ToServiceEndpointArtifactoryMapOutput() ServiceEndpointArtifactoryMapOutput {
	return o
}

func (o ServiceEndpointArtifactoryMapOutput) ToServiceEndpointArtifactoryMapOutputWithContext(ctx context.Context) ServiceEndpointArtifactoryMapOutput {
	return o
}

func (o ServiceEndpointArtifactoryMapOutput) MapIndex(k pulumi.StringInput) ServiceEndpointArtifactoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ServiceEndpointArtifactory {
		return vs[0].(map[string]ServiceEndpointArtifactory)[vs[1].(string)]
	}).(ServiceEndpointArtifactoryOutput)
}

func init() {
	pulumi.RegisterOutputType(ServiceEndpointArtifactoryOutput{})
	pulumi.RegisterOutputType(ServiceEndpointArtifactoryPtrOutput{})
	pulumi.RegisterOutputType(ServiceEndpointArtifactoryArrayOutput{})
	pulumi.RegisterOutputType(ServiceEndpointArtifactoryMapOutput{})
}
