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

// Manages a SonarQube Cloud service endpoint within Azure DevOps.
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
//			_, err = azuredevops.NewServiceEndpointSonarCloud(ctx, "example", &azuredevops.ServiceEndpointSonarCloudArgs{
//				ProjectId:           example.ID(),
//				ServiceEndpointName: pulumi.String("Example SonarCloud"),
//				Token:               pulumi.String("0000000000000000000000000000000000000000"),
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
// - [Azure DevOps Service REST API 7.0 - Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-7.0)
// - [Azure DevOps Service Connections](https://docs.microsoft.com/en-us/azure/devops/pipelines/library/service-endpoints?view=azure-devops&tabs=yaml)
// - [SonarCloud User Token](https://docs.sonarcloud.io/advanced-setup/user-accounts/)
//
// ## Import
//
// Azure DevOps SonarQube Cloud Service Endpoint can be imported using the **projectID/serviceEndpointID**, e.g.
//
// ```sh
// $ pulumi import azuredevops:index/serviceEndpointSonarCloud:ServiceEndpointSonarCloud example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
// ```
type ServiceEndpointSonarCloud struct {
	pulumi.CustomResourceState

	Authorization pulumi.StringMapOutput `pulumi:"authorization"`
	// The Service Endpoint description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The ID of the project.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The Service Endpoint name.
	ServiceEndpointName pulumi.StringOutput `pulumi:"serviceEndpointName"`
	// The Authentication Token generated through SonarCloud (go to `My Account > Security > Generate Tokens`).
	Token pulumi.StringOutput `pulumi:"token"`
}

// NewServiceEndpointSonarCloud registers a new resource with the given unique name, arguments, and options.
func NewServiceEndpointSonarCloud(ctx *pulumi.Context,
	name string, args *ServiceEndpointSonarCloudArgs, opts ...pulumi.ResourceOption) (*ServiceEndpointSonarCloud, error) {
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
	if args.Token != nil {
		args.Token = pulumi.ToSecret(args.Token).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"token",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ServiceEndpointSonarCloud
	err := ctx.RegisterResource("azuredevops:index/serviceEndpointSonarCloud:ServiceEndpointSonarCloud", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceEndpointSonarCloud gets an existing ServiceEndpointSonarCloud resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceEndpointSonarCloud(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceEndpointSonarCloudState, opts ...pulumi.ResourceOption) (*ServiceEndpointSonarCloud, error) {
	var resource ServiceEndpointSonarCloud
	err := ctx.ReadResource("azuredevops:index/serviceEndpointSonarCloud:ServiceEndpointSonarCloud", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceEndpointSonarCloud resources.
type serviceEndpointSonarCloudState struct {
	Authorization map[string]string `pulumi:"authorization"`
	// The Service Endpoint description.
	Description *string `pulumi:"description"`
	// The ID of the project.
	ProjectId *string `pulumi:"projectId"`
	// The Service Endpoint name.
	ServiceEndpointName *string `pulumi:"serviceEndpointName"`
	// The Authentication Token generated through SonarCloud (go to `My Account > Security > Generate Tokens`).
	Token *string `pulumi:"token"`
}

type ServiceEndpointSonarCloudState struct {
	Authorization pulumi.StringMapInput
	// The Service Endpoint description.
	Description pulumi.StringPtrInput
	// The ID of the project.
	ProjectId pulumi.StringPtrInput
	// The Service Endpoint name.
	ServiceEndpointName pulumi.StringPtrInput
	// The Authentication Token generated through SonarCloud (go to `My Account > Security > Generate Tokens`).
	Token pulumi.StringPtrInput
}

func (ServiceEndpointSonarCloudState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceEndpointSonarCloudState)(nil)).Elem()
}

type serviceEndpointSonarCloudArgs struct {
	// The Service Endpoint description.
	Description *string `pulumi:"description"`
	// The ID of the project.
	ProjectId string `pulumi:"projectId"`
	// The Service Endpoint name.
	ServiceEndpointName string `pulumi:"serviceEndpointName"`
	// The Authentication Token generated through SonarCloud (go to `My Account > Security > Generate Tokens`).
	Token string `pulumi:"token"`
}

// The set of arguments for constructing a ServiceEndpointSonarCloud resource.
type ServiceEndpointSonarCloudArgs struct {
	// The Service Endpoint description.
	Description pulumi.StringPtrInput
	// The ID of the project.
	ProjectId pulumi.StringInput
	// The Service Endpoint name.
	ServiceEndpointName pulumi.StringInput
	// The Authentication Token generated through SonarCloud (go to `My Account > Security > Generate Tokens`).
	Token pulumi.StringInput
}

func (ServiceEndpointSonarCloudArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceEndpointSonarCloudArgs)(nil)).Elem()
}

type ServiceEndpointSonarCloudInput interface {
	pulumi.Input

	ToServiceEndpointSonarCloudOutput() ServiceEndpointSonarCloudOutput
	ToServiceEndpointSonarCloudOutputWithContext(ctx context.Context) ServiceEndpointSonarCloudOutput
}

func (*ServiceEndpointSonarCloud) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceEndpointSonarCloud)(nil)).Elem()
}

func (i *ServiceEndpointSonarCloud) ToServiceEndpointSonarCloudOutput() ServiceEndpointSonarCloudOutput {
	return i.ToServiceEndpointSonarCloudOutputWithContext(context.Background())
}

func (i *ServiceEndpointSonarCloud) ToServiceEndpointSonarCloudOutputWithContext(ctx context.Context) ServiceEndpointSonarCloudOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointSonarCloudOutput)
}

// ServiceEndpointSonarCloudArrayInput is an input type that accepts ServiceEndpointSonarCloudArray and ServiceEndpointSonarCloudArrayOutput values.
// You can construct a concrete instance of `ServiceEndpointSonarCloudArrayInput` via:
//
//	ServiceEndpointSonarCloudArray{ ServiceEndpointSonarCloudArgs{...} }
type ServiceEndpointSonarCloudArrayInput interface {
	pulumi.Input

	ToServiceEndpointSonarCloudArrayOutput() ServiceEndpointSonarCloudArrayOutput
	ToServiceEndpointSonarCloudArrayOutputWithContext(context.Context) ServiceEndpointSonarCloudArrayOutput
}

type ServiceEndpointSonarCloudArray []ServiceEndpointSonarCloudInput

func (ServiceEndpointSonarCloudArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceEndpointSonarCloud)(nil)).Elem()
}

func (i ServiceEndpointSonarCloudArray) ToServiceEndpointSonarCloudArrayOutput() ServiceEndpointSonarCloudArrayOutput {
	return i.ToServiceEndpointSonarCloudArrayOutputWithContext(context.Background())
}

func (i ServiceEndpointSonarCloudArray) ToServiceEndpointSonarCloudArrayOutputWithContext(ctx context.Context) ServiceEndpointSonarCloudArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointSonarCloudArrayOutput)
}

// ServiceEndpointSonarCloudMapInput is an input type that accepts ServiceEndpointSonarCloudMap and ServiceEndpointSonarCloudMapOutput values.
// You can construct a concrete instance of `ServiceEndpointSonarCloudMapInput` via:
//
//	ServiceEndpointSonarCloudMap{ "key": ServiceEndpointSonarCloudArgs{...} }
type ServiceEndpointSonarCloudMapInput interface {
	pulumi.Input

	ToServiceEndpointSonarCloudMapOutput() ServiceEndpointSonarCloudMapOutput
	ToServiceEndpointSonarCloudMapOutputWithContext(context.Context) ServiceEndpointSonarCloudMapOutput
}

type ServiceEndpointSonarCloudMap map[string]ServiceEndpointSonarCloudInput

func (ServiceEndpointSonarCloudMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceEndpointSonarCloud)(nil)).Elem()
}

func (i ServiceEndpointSonarCloudMap) ToServiceEndpointSonarCloudMapOutput() ServiceEndpointSonarCloudMapOutput {
	return i.ToServiceEndpointSonarCloudMapOutputWithContext(context.Background())
}

func (i ServiceEndpointSonarCloudMap) ToServiceEndpointSonarCloudMapOutputWithContext(ctx context.Context) ServiceEndpointSonarCloudMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointSonarCloudMapOutput)
}

type ServiceEndpointSonarCloudOutput struct{ *pulumi.OutputState }

func (ServiceEndpointSonarCloudOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceEndpointSonarCloud)(nil)).Elem()
}

func (o ServiceEndpointSonarCloudOutput) ToServiceEndpointSonarCloudOutput() ServiceEndpointSonarCloudOutput {
	return o
}

func (o ServiceEndpointSonarCloudOutput) ToServiceEndpointSonarCloudOutputWithContext(ctx context.Context) ServiceEndpointSonarCloudOutput {
	return o
}

func (o ServiceEndpointSonarCloudOutput) Authorization() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ServiceEndpointSonarCloud) pulumi.StringMapOutput { return v.Authorization }).(pulumi.StringMapOutput)
}

// The Service Endpoint description.
func (o ServiceEndpointSonarCloudOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceEndpointSonarCloud) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The ID of the project.
func (o ServiceEndpointSonarCloudOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEndpointSonarCloud) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The Service Endpoint name.
func (o ServiceEndpointSonarCloudOutput) ServiceEndpointName() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEndpointSonarCloud) pulumi.StringOutput { return v.ServiceEndpointName }).(pulumi.StringOutput)
}

// The Authentication Token generated through SonarCloud (go to `My Account > Security > Generate Tokens`).
func (o ServiceEndpointSonarCloudOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEndpointSonarCloud) pulumi.StringOutput { return v.Token }).(pulumi.StringOutput)
}

type ServiceEndpointSonarCloudArrayOutput struct{ *pulumi.OutputState }

func (ServiceEndpointSonarCloudArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceEndpointSonarCloud)(nil)).Elem()
}

func (o ServiceEndpointSonarCloudArrayOutput) ToServiceEndpointSonarCloudArrayOutput() ServiceEndpointSonarCloudArrayOutput {
	return o
}

func (o ServiceEndpointSonarCloudArrayOutput) ToServiceEndpointSonarCloudArrayOutputWithContext(ctx context.Context) ServiceEndpointSonarCloudArrayOutput {
	return o
}

func (o ServiceEndpointSonarCloudArrayOutput) Index(i pulumi.IntInput) ServiceEndpointSonarCloudOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServiceEndpointSonarCloud {
		return vs[0].([]*ServiceEndpointSonarCloud)[vs[1].(int)]
	}).(ServiceEndpointSonarCloudOutput)
}

type ServiceEndpointSonarCloudMapOutput struct{ *pulumi.OutputState }

func (ServiceEndpointSonarCloudMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceEndpointSonarCloud)(nil)).Elem()
}

func (o ServiceEndpointSonarCloudMapOutput) ToServiceEndpointSonarCloudMapOutput() ServiceEndpointSonarCloudMapOutput {
	return o
}

func (o ServiceEndpointSonarCloudMapOutput) ToServiceEndpointSonarCloudMapOutputWithContext(ctx context.Context) ServiceEndpointSonarCloudMapOutput {
	return o
}

func (o ServiceEndpointSonarCloudMapOutput) MapIndex(k pulumi.StringInput) ServiceEndpointSonarCloudOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServiceEndpointSonarCloud {
		return vs[0].(map[string]*ServiceEndpointSonarCloud)[vs[1].(string)]
	}).(ServiceEndpointSonarCloudOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceEndpointSonarCloudInput)(nil)).Elem(), &ServiceEndpointSonarCloud{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceEndpointSonarCloudArrayInput)(nil)).Elem(), ServiceEndpointSonarCloudArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceEndpointSonarCloudMapInput)(nil)).Elem(), ServiceEndpointSonarCloudMap{})
	pulumi.RegisterOutputType(ServiceEndpointSonarCloudOutput{})
	pulumi.RegisterOutputType(ServiceEndpointSonarCloudArrayOutput{})
	pulumi.RegisterOutputType(ServiceEndpointSonarCloudMapOutput{})
}
