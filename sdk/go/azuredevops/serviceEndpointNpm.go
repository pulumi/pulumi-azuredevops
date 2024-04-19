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

// Manages a npm service endpoint within Azure DevOps.
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
//			example, err := azuredevops.NewProject(ctx, "example", &azuredevops.ProjectArgs{
//				Name:             pulumi.String("Example Project"),
//				Visibility:       pulumi.String("private"),
//				VersionControl:   pulumi.String("Git"),
//				WorkItemTemplate: pulumi.String("Agile"),
//				Description:      pulumi.String("Managed by Terraform"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuredevops.NewServiceEndpointNpm(ctx, "example", &azuredevops.ServiceEndpointNpmArgs{
//				ProjectId:           example.ID(),
//				ServiceEndpointName: pulumi.String("Example npm"),
//				Url:                 pulumi.String("https://registry.npmjs.org"),
//				AccessToken:         pulumi.String("00000000-0000-0000-0000-000000000000"),
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
// - [Azure DevOps Service REST API 7.0 - Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-7.0)
// - [Azure DevOps Service Connections](https://docs.microsoft.com/en-us/azure/devops/pipelines/library/service-endpoints?view=azure-devops&tabs=yaml)
// - [npm User Token](https://docs.npmjs.com/about-access-tokens)
//
// ## Import
//
// Azure DevOps Service Endpoint npm can be imported using the **projectID/serviceEndpointID**, e.g.
//
// ```sh
// $ pulumi import azuredevops:index/serviceEndpointNpm:ServiceEndpointNpm example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
// ```
type ServiceEndpointNpm struct {
	pulumi.CustomResourceState

	// The access token for npm registry.
	AccessToken   pulumi.StringOutput    `pulumi:"accessToken"`
	Authorization pulumi.StringMapOutput `pulumi:"authorization"`
	// The Service Endpoint description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The ID of the project.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The Service Endpoint name.
	ServiceEndpointName pulumi.StringOutput `pulumi:"serviceEndpointName"`
	// URL of the npm registry to connect with.
	Url pulumi.StringOutput `pulumi:"url"`
}

// NewServiceEndpointNpm registers a new resource with the given unique name, arguments, and options.
func NewServiceEndpointNpm(ctx *pulumi.Context,
	name string, args *ServiceEndpointNpmArgs, opts ...pulumi.ResourceOption) (*ServiceEndpointNpm, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccessToken == nil {
		return nil, errors.New("invalid value for required argument 'AccessToken'")
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
	if args.AccessToken != nil {
		args.AccessToken = pulumi.ToSecret(args.AccessToken).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"accessToken",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ServiceEndpointNpm
	err := ctx.RegisterResource("azuredevops:index/serviceEndpointNpm:ServiceEndpointNpm", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceEndpointNpm gets an existing ServiceEndpointNpm resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceEndpointNpm(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceEndpointNpmState, opts ...pulumi.ResourceOption) (*ServiceEndpointNpm, error) {
	var resource ServiceEndpointNpm
	err := ctx.ReadResource("azuredevops:index/serviceEndpointNpm:ServiceEndpointNpm", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceEndpointNpm resources.
type serviceEndpointNpmState struct {
	// The access token for npm registry.
	AccessToken   *string           `pulumi:"accessToken"`
	Authorization map[string]string `pulumi:"authorization"`
	// The Service Endpoint description.
	Description *string `pulumi:"description"`
	// The ID of the project.
	ProjectId *string `pulumi:"projectId"`
	// The Service Endpoint name.
	ServiceEndpointName *string `pulumi:"serviceEndpointName"`
	// URL of the npm registry to connect with.
	Url *string `pulumi:"url"`
}

type ServiceEndpointNpmState struct {
	// The access token for npm registry.
	AccessToken   pulumi.StringPtrInput
	Authorization pulumi.StringMapInput
	// The Service Endpoint description.
	Description pulumi.StringPtrInput
	// The ID of the project.
	ProjectId pulumi.StringPtrInput
	// The Service Endpoint name.
	ServiceEndpointName pulumi.StringPtrInput
	// URL of the npm registry to connect with.
	Url pulumi.StringPtrInput
}

func (ServiceEndpointNpmState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceEndpointNpmState)(nil)).Elem()
}

type serviceEndpointNpmArgs struct {
	// The access token for npm registry.
	AccessToken   string            `pulumi:"accessToken"`
	Authorization map[string]string `pulumi:"authorization"`
	// The Service Endpoint description.
	Description *string `pulumi:"description"`
	// The ID of the project.
	ProjectId string `pulumi:"projectId"`
	// The Service Endpoint name.
	ServiceEndpointName string `pulumi:"serviceEndpointName"`
	// URL of the npm registry to connect with.
	Url string `pulumi:"url"`
}

// The set of arguments for constructing a ServiceEndpointNpm resource.
type ServiceEndpointNpmArgs struct {
	// The access token for npm registry.
	AccessToken   pulumi.StringInput
	Authorization pulumi.StringMapInput
	// The Service Endpoint description.
	Description pulumi.StringPtrInput
	// The ID of the project.
	ProjectId pulumi.StringInput
	// The Service Endpoint name.
	ServiceEndpointName pulumi.StringInput
	// URL of the npm registry to connect with.
	Url pulumi.StringInput
}

func (ServiceEndpointNpmArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceEndpointNpmArgs)(nil)).Elem()
}

type ServiceEndpointNpmInput interface {
	pulumi.Input

	ToServiceEndpointNpmOutput() ServiceEndpointNpmOutput
	ToServiceEndpointNpmOutputWithContext(ctx context.Context) ServiceEndpointNpmOutput
}

func (*ServiceEndpointNpm) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceEndpointNpm)(nil)).Elem()
}

func (i *ServiceEndpointNpm) ToServiceEndpointNpmOutput() ServiceEndpointNpmOutput {
	return i.ToServiceEndpointNpmOutputWithContext(context.Background())
}

func (i *ServiceEndpointNpm) ToServiceEndpointNpmOutputWithContext(ctx context.Context) ServiceEndpointNpmOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointNpmOutput)
}

// ServiceEndpointNpmArrayInput is an input type that accepts ServiceEndpointNpmArray and ServiceEndpointNpmArrayOutput values.
// You can construct a concrete instance of `ServiceEndpointNpmArrayInput` via:
//
//	ServiceEndpointNpmArray{ ServiceEndpointNpmArgs{...} }
type ServiceEndpointNpmArrayInput interface {
	pulumi.Input

	ToServiceEndpointNpmArrayOutput() ServiceEndpointNpmArrayOutput
	ToServiceEndpointNpmArrayOutputWithContext(context.Context) ServiceEndpointNpmArrayOutput
}

type ServiceEndpointNpmArray []ServiceEndpointNpmInput

func (ServiceEndpointNpmArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceEndpointNpm)(nil)).Elem()
}

func (i ServiceEndpointNpmArray) ToServiceEndpointNpmArrayOutput() ServiceEndpointNpmArrayOutput {
	return i.ToServiceEndpointNpmArrayOutputWithContext(context.Background())
}

func (i ServiceEndpointNpmArray) ToServiceEndpointNpmArrayOutputWithContext(ctx context.Context) ServiceEndpointNpmArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointNpmArrayOutput)
}

// ServiceEndpointNpmMapInput is an input type that accepts ServiceEndpointNpmMap and ServiceEndpointNpmMapOutput values.
// You can construct a concrete instance of `ServiceEndpointNpmMapInput` via:
//
//	ServiceEndpointNpmMap{ "key": ServiceEndpointNpmArgs{...} }
type ServiceEndpointNpmMapInput interface {
	pulumi.Input

	ToServiceEndpointNpmMapOutput() ServiceEndpointNpmMapOutput
	ToServiceEndpointNpmMapOutputWithContext(context.Context) ServiceEndpointNpmMapOutput
}

type ServiceEndpointNpmMap map[string]ServiceEndpointNpmInput

func (ServiceEndpointNpmMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceEndpointNpm)(nil)).Elem()
}

func (i ServiceEndpointNpmMap) ToServiceEndpointNpmMapOutput() ServiceEndpointNpmMapOutput {
	return i.ToServiceEndpointNpmMapOutputWithContext(context.Background())
}

func (i ServiceEndpointNpmMap) ToServiceEndpointNpmMapOutputWithContext(ctx context.Context) ServiceEndpointNpmMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointNpmMapOutput)
}

type ServiceEndpointNpmOutput struct{ *pulumi.OutputState }

func (ServiceEndpointNpmOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceEndpointNpm)(nil)).Elem()
}

func (o ServiceEndpointNpmOutput) ToServiceEndpointNpmOutput() ServiceEndpointNpmOutput {
	return o
}

func (o ServiceEndpointNpmOutput) ToServiceEndpointNpmOutputWithContext(ctx context.Context) ServiceEndpointNpmOutput {
	return o
}

// The access token for npm registry.
func (o ServiceEndpointNpmOutput) AccessToken() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEndpointNpm) pulumi.StringOutput { return v.AccessToken }).(pulumi.StringOutput)
}

func (o ServiceEndpointNpmOutput) Authorization() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ServiceEndpointNpm) pulumi.StringMapOutput { return v.Authorization }).(pulumi.StringMapOutput)
}

// The Service Endpoint description.
func (o ServiceEndpointNpmOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceEndpointNpm) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The ID of the project.
func (o ServiceEndpointNpmOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEndpointNpm) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The Service Endpoint name.
func (o ServiceEndpointNpmOutput) ServiceEndpointName() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEndpointNpm) pulumi.StringOutput { return v.ServiceEndpointName }).(pulumi.StringOutput)
}

// URL of the npm registry to connect with.
func (o ServiceEndpointNpmOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEndpointNpm) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

type ServiceEndpointNpmArrayOutput struct{ *pulumi.OutputState }

func (ServiceEndpointNpmArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceEndpointNpm)(nil)).Elem()
}

func (o ServiceEndpointNpmArrayOutput) ToServiceEndpointNpmArrayOutput() ServiceEndpointNpmArrayOutput {
	return o
}

func (o ServiceEndpointNpmArrayOutput) ToServiceEndpointNpmArrayOutputWithContext(ctx context.Context) ServiceEndpointNpmArrayOutput {
	return o
}

func (o ServiceEndpointNpmArrayOutput) Index(i pulumi.IntInput) ServiceEndpointNpmOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServiceEndpointNpm {
		return vs[0].([]*ServiceEndpointNpm)[vs[1].(int)]
	}).(ServiceEndpointNpmOutput)
}

type ServiceEndpointNpmMapOutput struct{ *pulumi.OutputState }

func (ServiceEndpointNpmMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceEndpointNpm)(nil)).Elem()
}

func (o ServiceEndpointNpmMapOutput) ToServiceEndpointNpmMapOutput() ServiceEndpointNpmMapOutput {
	return o
}

func (o ServiceEndpointNpmMapOutput) ToServiceEndpointNpmMapOutputWithContext(ctx context.Context) ServiceEndpointNpmMapOutput {
	return o
}

func (o ServiceEndpointNpmMapOutput) MapIndex(k pulumi.StringInput) ServiceEndpointNpmOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServiceEndpointNpm {
		return vs[0].(map[string]*ServiceEndpointNpm)[vs[1].(string)]
	}).(ServiceEndpointNpmOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceEndpointNpmInput)(nil)).Elem(), &ServiceEndpointNpm{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceEndpointNpmArrayInput)(nil)).Elem(), ServiceEndpointNpmArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceEndpointNpmMapInput)(nil)).Elem(), ServiceEndpointNpmMap{})
	pulumi.RegisterOutputType(ServiceEndpointNpmOutput{})
	pulumi.RegisterOutputType(ServiceEndpointNpmArrayOutput{})
	pulumi.RegisterOutputType(ServiceEndpointNpmMapOutput{})
}
