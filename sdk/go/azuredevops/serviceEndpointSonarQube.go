// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a SonarQube service endpoint within Azure DevOps.
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
// 		_, err = azuredevops.NewServiceEndpointSonarQube(ctx, "serviceendpoint", &azuredevops.ServiceEndpointSonarQubeArgs{
// 			ProjectId:           project.ID(),
// 			ServiceEndpointName: pulumi.String("Sample SonarQube"),
// 			Url:                 pulumi.String("https://sonarqube.my.com"),
// 			Token:               pulumi.String("0000000000000000000000000000000000000000"),
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
// * [Azure DevOps Service Connections](https://docs.microsoft.com/en-us/azure/devops/pipelines/library/service-endpoints?view=azure-devops&tabs=yaml)
// * [SonarQube User Token](https://docs.sonarqube.org/latest/user-guide/user-token/)
//
// ## Import
//
// Azure DevOps Service Endpoint SonarQube can be imported using the **projectID/serviceEndpointID**, e.g.
//
// ```sh
//  $ pulumi import azuredevops:index/serviceEndpointSonarQube:ServiceEndpointSonarQube serviceendpoint 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
// ```
type ServiceEndpointSonarQube struct {
	pulumi.CustomResourceState

	Authorization pulumi.StringMapOutput `pulumi:"authorization"`
	// The Service Endpoint description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The project ID or project name.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The Service Endpoint name.
	ServiceEndpointName pulumi.StringOutput `pulumi:"serviceEndpointName"`
	// Authentication Token generated through SonarQube (go to My Account > Security > Generate Tokens).
	Token pulumi.StringOutput `pulumi:"token"`
	// A bcrypted hash of the attribute 'token'
	TokenHash pulumi.StringOutput `pulumi:"tokenHash"`
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
	// The project ID or project name.
	ProjectId *string `pulumi:"projectId"`
	// The Service Endpoint name.
	ServiceEndpointName *string `pulumi:"serviceEndpointName"`
	// Authentication Token generated through SonarQube (go to My Account > Security > Generate Tokens).
	Token *string `pulumi:"token"`
	// A bcrypted hash of the attribute 'token'
	TokenHash *string `pulumi:"tokenHash"`
	// URL of the SonarQube server to connect with.
	Url *string `pulumi:"url"`
}

type ServiceEndpointSonarQubeState struct {
	Authorization pulumi.StringMapInput
	// The Service Endpoint description.
	Description pulumi.StringPtrInput
	// The project ID or project name.
	ProjectId pulumi.StringPtrInput
	// The Service Endpoint name.
	ServiceEndpointName pulumi.StringPtrInput
	// Authentication Token generated through SonarQube (go to My Account > Security > Generate Tokens).
	Token pulumi.StringPtrInput
	// A bcrypted hash of the attribute 'token'
	TokenHash pulumi.StringPtrInput
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
	// The project ID or project name.
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
	// The project ID or project name.
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
	return reflect.TypeOf((*ServiceEndpointSonarQube)(nil))
}

func (i *ServiceEndpointSonarQube) ToServiceEndpointSonarQubeOutput() ServiceEndpointSonarQubeOutput {
	return i.ToServiceEndpointSonarQubeOutputWithContext(context.Background())
}

func (i *ServiceEndpointSonarQube) ToServiceEndpointSonarQubeOutputWithContext(ctx context.Context) ServiceEndpointSonarQubeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointSonarQubeOutput)
}

func (i *ServiceEndpointSonarQube) ToServiceEndpointSonarQubePtrOutput() ServiceEndpointSonarQubePtrOutput {
	return i.ToServiceEndpointSonarQubePtrOutputWithContext(context.Background())
}

func (i *ServiceEndpointSonarQube) ToServiceEndpointSonarQubePtrOutputWithContext(ctx context.Context) ServiceEndpointSonarQubePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointSonarQubePtrOutput)
}

type ServiceEndpointSonarQubePtrInput interface {
	pulumi.Input

	ToServiceEndpointSonarQubePtrOutput() ServiceEndpointSonarQubePtrOutput
	ToServiceEndpointSonarQubePtrOutputWithContext(ctx context.Context) ServiceEndpointSonarQubePtrOutput
}

type serviceEndpointSonarQubePtrType ServiceEndpointSonarQubeArgs

func (*serviceEndpointSonarQubePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceEndpointSonarQube)(nil))
}

func (i *serviceEndpointSonarQubePtrType) ToServiceEndpointSonarQubePtrOutput() ServiceEndpointSonarQubePtrOutput {
	return i.ToServiceEndpointSonarQubePtrOutputWithContext(context.Background())
}

func (i *serviceEndpointSonarQubePtrType) ToServiceEndpointSonarQubePtrOutputWithContext(ctx context.Context) ServiceEndpointSonarQubePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointSonarQubePtrOutput)
}

// ServiceEndpointSonarQubeArrayInput is an input type that accepts ServiceEndpointSonarQubeArray and ServiceEndpointSonarQubeArrayOutput values.
// You can construct a concrete instance of `ServiceEndpointSonarQubeArrayInput` via:
//
//          ServiceEndpointSonarQubeArray{ ServiceEndpointSonarQubeArgs{...} }
type ServiceEndpointSonarQubeArrayInput interface {
	pulumi.Input

	ToServiceEndpointSonarQubeArrayOutput() ServiceEndpointSonarQubeArrayOutput
	ToServiceEndpointSonarQubeArrayOutputWithContext(context.Context) ServiceEndpointSonarQubeArrayOutput
}

type ServiceEndpointSonarQubeArray []ServiceEndpointSonarQubeInput

func (ServiceEndpointSonarQubeArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*ServiceEndpointSonarQube)(nil))
}

func (i ServiceEndpointSonarQubeArray) ToServiceEndpointSonarQubeArrayOutput() ServiceEndpointSonarQubeArrayOutput {
	return i.ToServiceEndpointSonarQubeArrayOutputWithContext(context.Background())
}

func (i ServiceEndpointSonarQubeArray) ToServiceEndpointSonarQubeArrayOutputWithContext(ctx context.Context) ServiceEndpointSonarQubeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointSonarQubeArrayOutput)
}

// ServiceEndpointSonarQubeMapInput is an input type that accepts ServiceEndpointSonarQubeMap and ServiceEndpointSonarQubeMapOutput values.
// You can construct a concrete instance of `ServiceEndpointSonarQubeMapInput` via:
//
//          ServiceEndpointSonarQubeMap{ "key": ServiceEndpointSonarQubeArgs{...} }
type ServiceEndpointSonarQubeMapInput interface {
	pulumi.Input

	ToServiceEndpointSonarQubeMapOutput() ServiceEndpointSonarQubeMapOutput
	ToServiceEndpointSonarQubeMapOutputWithContext(context.Context) ServiceEndpointSonarQubeMapOutput
}

type ServiceEndpointSonarQubeMap map[string]ServiceEndpointSonarQubeInput

func (ServiceEndpointSonarQubeMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*ServiceEndpointSonarQube)(nil))
}

func (i ServiceEndpointSonarQubeMap) ToServiceEndpointSonarQubeMapOutput() ServiceEndpointSonarQubeMapOutput {
	return i.ToServiceEndpointSonarQubeMapOutputWithContext(context.Background())
}

func (i ServiceEndpointSonarQubeMap) ToServiceEndpointSonarQubeMapOutputWithContext(ctx context.Context) ServiceEndpointSonarQubeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointSonarQubeMapOutput)
}

type ServiceEndpointSonarQubeOutput struct {
	*pulumi.OutputState
}

func (ServiceEndpointSonarQubeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceEndpointSonarQube)(nil))
}

func (o ServiceEndpointSonarQubeOutput) ToServiceEndpointSonarQubeOutput() ServiceEndpointSonarQubeOutput {
	return o
}

func (o ServiceEndpointSonarQubeOutput) ToServiceEndpointSonarQubeOutputWithContext(ctx context.Context) ServiceEndpointSonarQubeOutput {
	return o
}

func (o ServiceEndpointSonarQubeOutput) ToServiceEndpointSonarQubePtrOutput() ServiceEndpointSonarQubePtrOutput {
	return o.ToServiceEndpointSonarQubePtrOutputWithContext(context.Background())
}

func (o ServiceEndpointSonarQubeOutput) ToServiceEndpointSonarQubePtrOutputWithContext(ctx context.Context) ServiceEndpointSonarQubePtrOutput {
	return o.ApplyT(func(v ServiceEndpointSonarQube) *ServiceEndpointSonarQube {
		return &v
	}).(ServiceEndpointSonarQubePtrOutput)
}

type ServiceEndpointSonarQubePtrOutput struct {
	*pulumi.OutputState
}

func (ServiceEndpointSonarQubePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceEndpointSonarQube)(nil))
}

func (o ServiceEndpointSonarQubePtrOutput) ToServiceEndpointSonarQubePtrOutput() ServiceEndpointSonarQubePtrOutput {
	return o
}

func (o ServiceEndpointSonarQubePtrOutput) ToServiceEndpointSonarQubePtrOutputWithContext(ctx context.Context) ServiceEndpointSonarQubePtrOutput {
	return o
}

type ServiceEndpointSonarQubeArrayOutput struct{ *pulumi.OutputState }

func (ServiceEndpointSonarQubeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceEndpointSonarQube)(nil))
}

func (o ServiceEndpointSonarQubeArrayOutput) ToServiceEndpointSonarQubeArrayOutput() ServiceEndpointSonarQubeArrayOutput {
	return o
}

func (o ServiceEndpointSonarQubeArrayOutput) ToServiceEndpointSonarQubeArrayOutputWithContext(ctx context.Context) ServiceEndpointSonarQubeArrayOutput {
	return o
}

func (o ServiceEndpointSonarQubeArrayOutput) Index(i pulumi.IntInput) ServiceEndpointSonarQubeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServiceEndpointSonarQube {
		return vs[0].([]ServiceEndpointSonarQube)[vs[1].(int)]
	}).(ServiceEndpointSonarQubeOutput)
}

type ServiceEndpointSonarQubeMapOutput struct{ *pulumi.OutputState }

func (ServiceEndpointSonarQubeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ServiceEndpointSonarQube)(nil))
}

func (o ServiceEndpointSonarQubeMapOutput) ToServiceEndpointSonarQubeMapOutput() ServiceEndpointSonarQubeMapOutput {
	return o
}

func (o ServiceEndpointSonarQubeMapOutput) ToServiceEndpointSonarQubeMapOutputWithContext(ctx context.Context) ServiceEndpointSonarQubeMapOutput {
	return o
}

func (o ServiceEndpointSonarQubeMapOutput) MapIndex(k pulumi.StringInput) ServiceEndpointSonarQubeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ServiceEndpointSonarQube {
		return vs[0].(map[string]ServiceEndpointSonarQube)[vs[1].(string)]
	}).(ServiceEndpointSonarQubeOutput)
}

func init() {
	pulumi.RegisterOutputType(ServiceEndpointSonarQubeOutput{})
	pulumi.RegisterOutputType(ServiceEndpointSonarQubePtrOutput{})
	pulumi.RegisterOutputType(ServiceEndpointSonarQubeArrayOutput{})
	pulumi.RegisterOutputType(ServiceEndpointSonarQubeMapOutput{})
}
