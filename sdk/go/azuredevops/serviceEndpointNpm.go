// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a npm service endpoint within Azure DevOps.
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
// 		_, err = azuredevops.NewServiceEndpointNpm(ctx, "serviceendpoint", &azuredevops.ServiceEndpointNpmArgs{
// 			ProjectId:           project.ID(),
// 			ServiceEndpointName: pulumi.String("Sample npm"),
// 			Url:                 pulumi.String("https://registry.npmjs.org"),
// 			AccessToken:         pulumi.String("00000000-0000-0000-0000-000000000000"),
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
// - [Azure DevOps Service Connections](https://docs.microsoft.com/en-us/azure/devops/pipelines/library/service-endpoints?view=azure-devops&tabs=yaml)
// - [npm User Token](https://docs.npmjs.com/about-access-tokens)
//
// ## Import
//
// Azure DevOps Service Endpoint npm can be imported using the **projectID/serviceEndpointID**, e.g.
//
// ```sh
//  $ pulumi import azuredevops:index/serviceEndpointNpm:ServiceEndpointNpm serviceendpoint 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
// ```
type ServiceEndpointNpm struct {
	pulumi.CustomResourceState

	// The access token for npm registry.
	AccessToken pulumi.StringOutput `pulumi:"accessToken"`
	// A bcrypted hash of the attribute 'access_token'
	AccessTokenHash pulumi.StringOutput    `pulumi:"accessTokenHash"`
	Authorization   pulumi.StringMapOutput `pulumi:"authorization"`
	// The Service Endpoint description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The project ID or project name.
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
	AccessToken *string `pulumi:"accessToken"`
	// A bcrypted hash of the attribute 'access_token'
	AccessTokenHash *string           `pulumi:"accessTokenHash"`
	Authorization   map[string]string `pulumi:"authorization"`
	// The Service Endpoint description.
	Description *string `pulumi:"description"`
	// The project ID or project name.
	ProjectId *string `pulumi:"projectId"`
	// The Service Endpoint name.
	ServiceEndpointName *string `pulumi:"serviceEndpointName"`
	// URL of the npm registry to connect with.
	Url *string `pulumi:"url"`
}

type ServiceEndpointNpmState struct {
	// The access token for npm registry.
	AccessToken pulumi.StringPtrInput
	// A bcrypted hash of the attribute 'access_token'
	AccessTokenHash pulumi.StringPtrInput
	Authorization   pulumi.StringMapInput
	// The Service Endpoint description.
	Description pulumi.StringPtrInput
	// The project ID or project name.
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
	// The project ID or project name.
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
	// The project ID or project name.
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
	return reflect.TypeOf((*ServiceEndpointNpm)(nil))
}

func (i *ServiceEndpointNpm) ToServiceEndpointNpmOutput() ServiceEndpointNpmOutput {
	return i.ToServiceEndpointNpmOutputWithContext(context.Background())
}

func (i *ServiceEndpointNpm) ToServiceEndpointNpmOutputWithContext(ctx context.Context) ServiceEndpointNpmOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointNpmOutput)
}

func (i *ServiceEndpointNpm) ToServiceEndpointNpmPtrOutput() ServiceEndpointNpmPtrOutput {
	return i.ToServiceEndpointNpmPtrOutputWithContext(context.Background())
}

func (i *ServiceEndpointNpm) ToServiceEndpointNpmPtrOutputWithContext(ctx context.Context) ServiceEndpointNpmPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointNpmPtrOutput)
}

type ServiceEndpointNpmPtrInput interface {
	pulumi.Input

	ToServiceEndpointNpmPtrOutput() ServiceEndpointNpmPtrOutput
	ToServiceEndpointNpmPtrOutputWithContext(ctx context.Context) ServiceEndpointNpmPtrOutput
}

type serviceEndpointNpmPtrType ServiceEndpointNpmArgs

func (*serviceEndpointNpmPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceEndpointNpm)(nil))
}

func (i *serviceEndpointNpmPtrType) ToServiceEndpointNpmPtrOutput() ServiceEndpointNpmPtrOutput {
	return i.ToServiceEndpointNpmPtrOutputWithContext(context.Background())
}

func (i *serviceEndpointNpmPtrType) ToServiceEndpointNpmPtrOutputWithContext(ctx context.Context) ServiceEndpointNpmPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointNpmPtrOutput)
}

// ServiceEndpointNpmArrayInput is an input type that accepts ServiceEndpointNpmArray and ServiceEndpointNpmArrayOutput values.
// You can construct a concrete instance of `ServiceEndpointNpmArrayInput` via:
//
//          ServiceEndpointNpmArray{ ServiceEndpointNpmArgs{...} }
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
//          ServiceEndpointNpmMap{ "key": ServiceEndpointNpmArgs{...} }
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
	return reflect.TypeOf((*ServiceEndpointNpm)(nil))
}

func (o ServiceEndpointNpmOutput) ToServiceEndpointNpmOutput() ServiceEndpointNpmOutput {
	return o
}

func (o ServiceEndpointNpmOutput) ToServiceEndpointNpmOutputWithContext(ctx context.Context) ServiceEndpointNpmOutput {
	return o
}

func (o ServiceEndpointNpmOutput) ToServiceEndpointNpmPtrOutput() ServiceEndpointNpmPtrOutput {
	return o.ToServiceEndpointNpmPtrOutputWithContext(context.Background())
}

func (o ServiceEndpointNpmOutput) ToServiceEndpointNpmPtrOutputWithContext(ctx context.Context) ServiceEndpointNpmPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServiceEndpointNpm) *ServiceEndpointNpm {
		return &v
	}).(ServiceEndpointNpmPtrOutput)
}

type ServiceEndpointNpmPtrOutput struct{ *pulumi.OutputState }

func (ServiceEndpointNpmPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceEndpointNpm)(nil))
}

func (o ServiceEndpointNpmPtrOutput) ToServiceEndpointNpmPtrOutput() ServiceEndpointNpmPtrOutput {
	return o
}

func (o ServiceEndpointNpmPtrOutput) ToServiceEndpointNpmPtrOutputWithContext(ctx context.Context) ServiceEndpointNpmPtrOutput {
	return o
}

func (o ServiceEndpointNpmPtrOutput) Elem() ServiceEndpointNpmOutput {
	return o.ApplyT(func(v *ServiceEndpointNpm) ServiceEndpointNpm {
		if v != nil {
			return *v
		}
		var ret ServiceEndpointNpm
		return ret
	}).(ServiceEndpointNpmOutput)
}

type ServiceEndpointNpmArrayOutput struct{ *pulumi.OutputState }

func (ServiceEndpointNpmArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceEndpointNpm)(nil))
}

func (o ServiceEndpointNpmArrayOutput) ToServiceEndpointNpmArrayOutput() ServiceEndpointNpmArrayOutput {
	return o
}

func (o ServiceEndpointNpmArrayOutput) ToServiceEndpointNpmArrayOutputWithContext(ctx context.Context) ServiceEndpointNpmArrayOutput {
	return o
}

func (o ServiceEndpointNpmArrayOutput) Index(i pulumi.IntInput) ServiceEndpointNpmOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServiceEndpointNpm {
		return vs[0].([]ServiceEndpointNpm)[vs[1].(int)]
	}).(ServiceEndpointNpmOutput)
}

type ServiceEndpointNpmMapOutput struct{ *pulumi.OutputState }

func (ServiceEndpointNpmMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ServiceEndpointNpm)(nil))
}

func (o ServiceEndpointNpmMapOutput) ToServiceEndpointNpmMapOutput() ServiceEndpointNpmMapOutput {
	return o
}

func (o ServiceEndpointNpmMapOutput) ToServiceEndpointNpmMapOutputWithContext(ctx context.Context) ServiceEndpointNpmMapOutput {
	return o
}

func (o ServiceEndpointNpmMapOutput) MapIndex(k pulumi.StringInput) ServiceEndpointNpmOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ServiceEndpointNpm {
		return vs[0].(map[string]ServiceEndpointNpm)[vs[1].(string)]
	}).(ServiceEndpointNpmOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceEndpointNpmInput)(nil)).Elem(), &ServiceEndpointNpm{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceEndpointNpmPtrInput)(nil)).Elem(), &ServiceEndpointNpm{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceEndpointNpmArrayInput)(nil)).Elem(), ServiceEndpointNpmArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceEndpointNpmMapInput)(nil)).Elem(), ServiceEndpointNpmMap{})
	pulumi.RegisterOutputType(ServiceEndpointNpmOutput{})
	pulumi.RegisterOutputType(ServiceEndpointNpmPtrOutput{})
	pulumi.RegisterOutputType(ServiceEndpointNpmArrayOutput{})
	pulumi.RegisterOutputType(ServiceEndpointNpmMapOutput{})
}
