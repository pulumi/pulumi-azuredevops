// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages authorization of resources, e.g. for access in build pipelines.
//
// Currently supported resources: service endpoint (aka service connection, endpoint).
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
//			exampleServiceEndpointBitBucket, err := azuredevops.NewServiceEndpointBitBucket(ctx, "exampleServiceEndpointBitBucket", &azuredevops.ServiceEndpointBitBucketArgs{
//				ProjectId:           exampleProject.ID(),
//				Username:            pulumi.String("username"),
//				Password:            pulumi.String("password"),
//				ServiceEndpointName: pulumi.String("example-bitbucket"),
//				Description:         pulumi.String("Managed by Terraform"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuredevops.NewResourceAuthorization(ctx, "exampleResourceAuthorization", &azuredevops.ResourceAuthorizationArgs{
//				ProjectId:  exampleProject.ID(),
//				ResourceId: exampleServiceEndpointBitBucket.ID(),
//				Authorized: pulumi.Bool(true),
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
// - [Azure DevOps Service REST API 6.0 - Authorize Definition Resource](https://docs.microsoft.com/en-us/rest/api/azure/devops/build/resources/authorize%20definition%20resources?view=azure-devops-rest-6.0)
type ResourceAuthorization struct {
	pulumi.CustomResourceState

	// Set to true to allow public access in the project. Type: boolean.
	Authorized pulumi.BoolOutput `pulumi:"authorized"`
	// The ID of the build definition to authorize. Type: string.
	DefinitionId pulumi.IntPtrOutput `pulumi:"definitionId"`
	// The project ID or project name. Type: string.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The ID of the resource to authorize. Type: string.
	ResourceId pulumi.StringOutput `pulumi:"resourceId"`
	// The type of the resource to authorize. Type: string. Valid values: `endpoint`, `queue`, `variablegroup`. Default value: `endpoint`.
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewResourceAuthorization registers a new resource with the given unique name, arguments, and options.
func NewResourceAuthorization(ctx *pulumi.Context,
	name string, args *ResourceAuthorizationArgs, opts ...pulumi.ResourceOption) (*ResourceAuthorization, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Authorized == nil {
		return nil, errors.New("invalid value for required argument 'Authorized'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.ResourceId == nil {
		return nil, errors.New("invalid value for required argument 'ResourceId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azuredevops:Security/resourceAuthorization:ResourceAuthorization"),
		},
	})
	opts = append(opts, aliases)
	var resource ResourceAuthorization
	err := ctx.RegisterResource("azuredevops:index/resourceAuthorization:ResourceAuthorization", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResourceAuthorization gets an existing ResourceAuthorization resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResourceAuthorization(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourceAuthorizationState, opts ...pulumi.ResourceOption) (*ResourceAuthorization, error) {
	var resource ResourceAuthorization
	err := ctx.ReadResource("azuredevops:index/resourceAuthorization:ResourceAuthorization", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResourceAuthorization resources.
type resourceAuthorizationState struct {
	// Set to true to allow public access in the project. Type: boolean.
	Authorized *bool `pulumi:"authorized"`
	// The ID of the build definition to authorize. Type: string.
	DefinitionId *int `pulumi:"definitionId"`
	// The project ID or project name. Type: string.
	ProjectId *string `pulumi:"projectId"`
	// The ID of the resource to authorize. Type: string.
	ResourceId *string `pulumi:"resourceId"`
	// The type of the resource to authorize. Type: string. Valid values: `endpoint`, `queue`, `variablegroup`. Default value: `endpoint`.
	Type *string `pulumi:"type"`
}

type ResourceAuthorizationState struct {
	// Set to true to allow public access in the project. Type: boolean.
	Authorized pulumi.BoolPtrInput
	// The ID of the build definition to authorize. Type: string.
	DefinitionId pulumi.IntPtrInput
	// The project ID or project name. Type: string.
	ProjectId pulumi.StringPtrInput
	// The ID of the resource to authorize. Type: string.
	ResourceId pulumi.StringPtrInput
	// The type of the resource to authorize. Type: string. Valid values: `endpoint`, `queue`, `variablegroup`. Default value: `endpoint`.
	Type pulumi.StringPtrInput
}

func (ResourceAuthorizationState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceAuthorizationState)(nil)).Elem()
}

type resourceAuthorizationArgs struct {
	// Set to true to allow public access in the project. Type: boolean.
	Authorized bool `pulumi:"authorized"`
	// The ID of the build definition to authorize. Type: string.
	DefinitionId *int `pulumi:"definitionId"`
	// The project ID or project name. Type: string.
	ProjectId string `pulumi:"projectId"`
	// The ID of the resource to authorize. Type: string.
	ResourceId string `pulumi:"resourceId"`
	// The type of the resource to authorize. Type: string. Valid values: `endpoint`, `queue`, `variablegroup`. Default value: `endpoint`.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a ResourceAuthorization resource.
type ResourceAuthorizationArgs struct {
	// Set to true to allow public access in the project. Type: boolean.
	Authorized pulumi.BoolInput
	// The ID of the build definition to authorize. Type: string.
	DefinitionId pulumi.IntPtrInput
	// The project ID or project name. Type: string.
	ProjectId pulumi.StringInput
	// The ID of the resource to authorize. Type: string.
	ResourceId pulumi.StringInput
	// The type of the resource to authorize. Type: string. Valid values: `endpoint`, `queue`, `variablegroup`. Default value: `endpoint`.
	Type pulumi.StringPtrInput
}

func (ResourceAuthorizationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceAuthorizationArgs)(nil)).Elem()
}

type ResourceAuthorizationInput interface {
	pulumi.Input

	ToResourceAuthorizationOutput() ResourceAuthorizationOutput
	ToResourceAuthorizationOutputWithContext(ctx context.Context) ResourceAuthorizationOutput
}

func (*ResourceAuthorization) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceAuthorization)(nil)).Elem()
}

func (i *ResourceAuthorization) ToResourceAuthorizationOutput() ResourceAuthorizationOutput {
	return i.ToResourceAuthorizationOutputWithContext(context.Background())
}

func (i *ResourceAuthorization) ToResourceAuthorizationOutputWithContext(ctx context.Context) ResourceAuthorizationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceAuthorizationOutput)
}

// ResourceAuthorizationArrayInput is an input type that accepts ResourceAuthorizationArray and ResourceAuthorizationArrayOutput values.
// You can construct a concrete instance of `ResourceAuthorizationArrayInput` via:
//
//	ResourceAuthorizationArray{ ResourceAuthorizationArgs{...} }
type ResourceAuthorizationArrayInput interface {
	pulumi.Input

	ToResourceAuthorizationArrayOutput() ResourceAuthorizationArrayOutput
	ToResourceAuthorizationArrayOutputWithContext(context.Context) ResourceAuthorizationArrayOutput
}

type ResourceAuthorizationArray []ResourceAuthorizationInput

func (ResourceAuthorizationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ResourceAuthorization)(nil)).Elem()
}

func (i ResourceAuthorizationArray) ToResourceAuthorizationArrayOutput() ResourceAuthorizationArrayOutput {
	return i.ToResourceAuthorizationArrayOutputWithContext(context.Background())
}

func (i ResourceAuthorizationArray) ToResourceAuthorizationArrayOutputWithContext(ctx context.Context) ResourceAuthorizationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceAuthorizationArrayOutput)
}

// ResourceAuthorizationMapInput is an input type that accepts ResourceAuthorizationMap and ResourceAuthorizationMapOutput values.
// You can construct a concrete instance of `ResourceAuthorizationMapInput` via:
//
//	ResourceAuthorizationMap{ "key": ResourceAuthorizationArgs{...} }
type ResourceAuthorizationMapInput interface {
	pulumi.Input

	ToResourceAuthorizationMapOutput() ResourceAuthorizationMapOutput
	ToResourceAuthorizationMapOutputWithContext(context.Context) ResourceAuthorizationMapOutput
}

type ResourceAuthorizationMap map[string]ResourceAuthorizationInput

func (ResourceAuthorizationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ResourceAuthorization)(nil)).Elem()
}

func (i ResourceAuthorizationMap) ToResourceAuthorizationMapOutput() ResourceAuthorizationMapOutput {
	return i.ToResourceAuthorizationMapOutputWithContext(context.Background())
}

func (i ResourceAuthorizationMap) ToResourceAuthorizationMapOutputWithContext(ctx context.Context) ResourceAuthorizationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceAuthorizationMapOutput)
}

type ResourceAuthorizationOutput struct{ *pulumi.OutputState }

func (ResourceAuthorizationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceAuthorization)(nil)).Elem()
}

func (o ResourceAuthorizationOutput) ToResourceAuthorizationOutput() ResourceAuthorizationOutput {
	return o
}

func (o ResourceAuthorizationOutput) ToResourceAuthorizationOutputWithContext(ctx context.Context) ResourceAuthorizationOutput {
	return o
}

// Set to true to allow public access in the project. Type: boolean.
func (o ResourceAuthorizationOutput) Authorized() pulumi.BoolOutput {
	return o.ApplyT(func(v *ResourceAuthorization) pulumi.BoolOutput { return v.Authorized }).(pulumi.BoolOutput)
}

// The ID of the build definition to authorize. Type: string.
func (o ResourceAuthorizationOutput) DefinitionId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ResourceAuthorization) pulumi.IntPtrOutput { return v.DefinitionId }).(pulumi.IntPtrOutput)
}

// The project ID or project name. Type: string.
func (o ResourceAuthorizationOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceAuthorization) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The ID of the resource to authorize. Type: string.
func (o ResourceAuthorizationOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceAuthorization) pulumi.StringOutput { return v.ResourceId }).(pulumi.StringOutput)
}

// The type of the resource to authorize. Type: string. Valid values: `endpoint`, `queue`, `variablegroup`. Default value: `endpoint`.
func (o ResourceAuthorizationOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceAuthorization) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

type ResourceAuthorizationArrayOutput struct{ *pulumi.OutputState }

func (ResourceAuthorizationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ResourceAuthorization)(nil)).Elem()
}

func (o ResourceAuthorizationArrayOutput) ToResourceAuthorizationArrayOutput() ResourceAuthorizationArrayOutput {
	return o
}

func (o ResourceAuthorizationArrayOutput) ToResourceAuthorizationArrayOutputWithContext(ctx context.Context) ResourceAuthorizationArrayOutput {
	return o
}

func (o ResourceAuthorizationArrayOutput) Index(i pulumi.IntInput) ResourceAuthorizationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ResourceAuthorization {
		return vs[0].([]*ResourceAuthorization)[vs[1].(int)]
	}).(ResourceAuthorizationOutput)
}

type ResourceAuthorizationMapOutput struct{ *pulumi.OutputState }

func (ResourceAuthorizationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ResourceAuthorization)(nil)).Elem()
}

func (o ResourceAuthorizationMapOutput) ToResourceAuthorizationMapOutput() ResourceAuthorizationMapOutput {
	return o
}

func (o ResourceAuthorizationMapOutput) ToResourceAuthorizationMapOutputWithContext(ctx context.Context) ResourceAuthorizationMapOutput {
	return o
}

func (o ResourceAuthorizationMapOutput) MapIndex(k pulumi.StringInput) ResourceAuthorizationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ResourceAuthorization {
		return vs[0].(map[string]*ResourceAuthorization)[vs[1].(string)]
	}).(ResourceAuthorizationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceAuthorizationInput)(nil)).Elem(), &ResourceAuthorization{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceAuthorizationArrayInput)(nil)).Elem(), ResourceAuthorizationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceAuthorizationMapInput)(nil)).Elem(), ResourceAuthorizationMap{})
	pulumi.RegisterOutputType(ResourceAuthorizationOutput{})
	pulumi.RegisterOutputType(ResourceAuthorizationArrayOutput{})
	pulumi.RegisterOutputType(ResourceAuthorizationMapOutput{})
}
