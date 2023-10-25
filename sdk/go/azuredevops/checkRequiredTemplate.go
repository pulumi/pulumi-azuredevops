// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azuredevops/sdk/v2/go/azuredevops/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Manages a Required Template Check.
//
// ## Example Usage
// ### Protect a service connection
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
//			exampleProject, err := azuredevops.NewProject(ctx, "exampleProject", nil)
//			if err != nil {
//				return err
//			}
//			exampleServiceEndpointGeneric, err := azuredevops.NewServiceEndpointGeneric(ctx, "exampleServiceEndpointGeneric", &azuredevops.ServiceEndpointGenericArgs{
//				ProjectId:           exampleProject.ID(),
//				ServerUrl:           pulumi.String("https://some-server.example.com"),
//				Username:            pulumi.String("username"),
//				Password:            pulumi.String("password"),
//				ServiceEndpointName: pulumi.String("Example Generic"),
//				Description:         pulumi.String("Managed by Terraform"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuredevops.NewCheckRequiredTemplate(ctx, "exampleCheckRequiredTemplate", &azuredevops.CheckRequiredTemplateArgs{
//				ProjectId:          exampleProject.ID(),
//				TargetResourceId:   exampleServiceEndpointGeneric.ID(),
//				TargetResourceType: pulumi.String("endpoint"),
//				RequiredTemplates: azuredevops.CheckRequiredTemplateRequiredTemplateArray{
//					&azuredevops.CheckRequiredTemplateRequiredTemplateArgs{
//						RepositoryType: pulumi.String("azuregit"),
//						RepositoryName: pulumi.String("project/repository"),
//						RepositoryRef:  pulumi.String("refs/heads/main"),
//						TemplatePath:   pulumi.String("template/path.yml"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Protect an environment
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
//			exampleProject, err := azuredevops.NewProject(ctx, "exampleProject", nil)
//			if err != nil {
//				return err
//			}
//			exampleEnvironment, err := azuredevops.NewEnvironment(ctx, "exampleEnvironment", &azuredevops.EnvironmentArgs{
//				ProjectId: exampleProject.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuredevops.NewCheckRequiredTemplate(ctx, "exampleCheckRequiredTemplate", &azuredevops.CheckRequiredTemplateArgs{
//				ProjectId:          exampleProject.ID(),
//				TargetResourceId:   exampleEnvironment.ID(),
//				TargetResourceType: pulumi.String("environment"),
//				RequiredTemplates: azuredevops.CheckRequiredTemplateRequiredTemplateArray{
//					&azuredevops.CheckRequiredTemplateRequiredTemplateArgs{
//						RepositoryName: pulumi.String("project/repository"),
//						RepositoryRef:  pulumi.String("refs/heads/main"),
//						TemplatePath:   pulumi.String("template/path.yml"),
//					},
//					&azuredevops.CheckRequiredTemplateRequiredTemplateArgs{
//						RepositoryName: pulumi.String("project/repository"),
//						RepositoryRef:  pulumi.String("refs/heads/main"),
//						TemplatePath:   pulumi.String("template/alternate-path.yml"),
//					},
//				},
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
// ## Import
//
// Importing this resource is not supported.
type CheckRequiredTemplate struct {
	pulumi.CustomResourceState

	// The project ID. Changing this forces a new Required Template Check to be created.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// One or more `requiredTemplate` blocks documented below.
	RequiredTemplates CheckRequiredTemplateRequiredTemplateArrayOutput `pulumi:"requiredTemplates"`
	// The ID of the resource being protected by the check. Changing this forces a new Required Template Check to be created.
	TargetResourceId pulumi.StringOutput `pulumi:"targetResourceId"`
	// The type of resource being protected by the check. Valid values: `endpoint`, `environment`, `queue`, `repository`, `securefile`, `variablegroup`. Changing this forces a new Required Template Check to be created.
	TargetResourceType pulumi.StringOutput `pulumi:"targetResourceType"`
}

// NewCheckRequiredTemplate registers a new resource with the given unique name, arguments, and options.
func NewCheckRequiredTemplate(ctx *pulumi.Context,
	name string, args *CheckRequiredTemplateArgs, opts ...pulumi.ResourceOption) (*CheckRequiredTemplate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.RequiredTemplates == nil {
		return nil, errors.New("invalid value for required argument 'RequiredTemplates'")
	}
	if args.TargetResourceId == nil {
		return nil, errors.New("invalid value for required argument 'TargetResourceId'")
	}
	if args.TargetResourceType == nil {
		return nil, errors.New("invalid value for required argument 'TargetResourceType'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CheckRequiredTemplate
	err := ctx.RegisterResource("azuredevops:index/checkRequiredTemplate:CheckRequiredTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCheckRequiredTemplate gets an existing CheckRequiredTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCheckRequiredTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CheckRequiredTemplateState, opts ...pulumi.ResourceOption) (*CheckRequiredTemplate, error) {
	var resource CheckRequiredTemplate
	err := ctx.ReadResource("azuredevops:index/checkRequiredTemplate:CheckRequiredTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CheckRequiredTemplate resources.
type checkRequiredTemplateState struct {
	// The project ID. Changing this forces a new Required Template Check to be created.
	ProjectId *string `pulumi:"projectId"`
	// One or more `requiredTemplate` blocks documented below.
	RequiredTemplates []CheckRequiredTemplateRequiredTemplate `pulumi:"requiredTemplates"`
	// The ID of the resource being protected by the check. Changing this forces a new Required Template Check to be created.
	TargetResourceId *string `pulumi:"targetResourceId"`
	// The type of resource being protected by the check. Valid values: `endpoint`, `environment`, `queue`, `repository`, `securefile`, `variablegroup`. Changing this forces a new Required Template Check to be created.
	TargetResourceType *string `pulumi:"targetResourceType"`
}

type CheckRequiredTemplateState struct {
	// The project ID. Changing this forces a new Required Template Check to be created.
	ProjectId pulumi.StringPtrInput
	// One or more `requiredTemplate` blocks documented below.
	RequiredTemplates CheckRequiredTemplateRequiredTemplateArrayInput
	// The ID of the resource being protected by the check. Changing this forces a new Required Template Check to be created.
	TargetResourceId pulumi.StringPtrInput
	// The type of resource being protected by the check. Valid values: `endpoint`, `environment`, `queue`, `repository`, `securefile`, `variablegroup`. Changing this forces a new Required Template Check to be created.
	TargetResourceType pulumi.StringPtrInput
}

func (CheckRequiredTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*checkRequiredTemplateState)(nil)).Elem()
}

type checkRequiredTemplateArgs struct {
	// The project ID. Changing this forces a new Required Template Check to be created.
	ProjectId string `pulumi:"projectId"`
	// One or more `requiredTemplate` blocks documented below.
	RequiredTemplates []CheckRequiredTemplateRequiredTemplate `pulumi:"requiredTemplates"`
	// The ID of the resource being protected by the check. Changing this forces a new Required Template Check to be created.
	TargetResourceId string `pulumi:"targetResourceId"`
	// The type of resource being protected by the check. Valid values: `endpoint`, `environment`, `queue`, `repository`, `securefile`, `variablegroup`. Changing this forces a new Required Template Check to be created.
	TargetResourceType string `pulumi:"targetResourceType"`
}

// The set of arguments for constructing a CheckRequiredTemplate resource.
type CheckRequiredTemplateArgs struct {
	// The project ID. Changing this forces a new Required Template Check to be created.
	ProjectId pulumi.StringInput
	// One or more `requiredTemplate` blocks documented below.
	RequiredTemplates CheckRequiredTemplateRequiredTemplateArrayInput
	// The ID of the resource being protected by the check. Changing this forces a new Required Template Check to be created.
	TargetResourceId pulumi.StringInput
	// The type of resource being protected by the check. Valid values: `endpoint`, `environment`, `queue`, `repository`, `securefile`, `variablegroup`. Changing this forces a new Required Template Check to be created.
	TargetResourceType pulumi.StringInput
}

func (CheckRequiredTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*checkRequiredTemplateArgs)(nil)).Elem()
}

type CheckRequiredTemplateInput interface {
	pulumi.Input

	ToCheckRequiredTemplateOutput() CheckRequiredTemplateOutput
	ToCheckRequiredTemplateOutputWithContext(ctx context.Context) CheckRequiredTemplateOutput
}

func (*CheckRequiredTemplate) ElementType() reflect.Type {
	return reflect.TypeOf((**CheckRequiredTemplate)(nil)).Elem()
}

func (i *CheckRequiredTemplate) ToCheckRequiredTemplateOutput() CheckRequiredTemplateOutput {
	return i.ToCheckRequiredTemplateOutputWithContext(context.Background())
}

func (i *CheckRequiredTemplate) ToCheckRequiredTemplateOutputWithContext(ctx context.Context) CheckRequiredTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CheckRequiredTemplateOutput)
}

func (i *CheckRequiredTemplate) ToOutput(ctx context.Context) pulumix.Output[*CheckRequiredTemplate] {
	return pulumix.Output[*CheckRequiredTemplate]{
		OutputState: i.ToCheckRequiredTemplateOutputWithContext(ctx).OutputState,
	}
}

// CheckRequiredTemplateArrayInput is an input type that accepts CheckRequiredTemplateArray and CheckRequiredTemplateArrayOutput values.
// You can construct a concrete instance of `CheckRequiredTemplateArrayInput` via:
//
//	CheckRequiredTemplateArray{ CheckRequiredTemplateArgs{...} }
type CheckRequiredTemplateArrayInput interface {
	pulumi.Input

	ToCheckRequiredTemplateArrayOutput() CheckRequiredTemplateArrayOutput
	ToCheckRequiredTemplateArrayOutputWithContext(context.Context) CheckRequiredTemplateArrayOutput
}

type CheckRequiredTemplateArray []CheckRequiredTemplateInput

func (CheckRequiredTemplateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CheckRequiredTemplate)(nil)).Elem()
}

func (i CheckRequiredTemplateArray) ToCheckRequiredTemplateArrayOutput() CheckRequiredTemplateArrayOutput {
	return i.ToCheckRequiredTemplateArrayOutputWithContext(context.Background())
}

func (i CheckRequiredTemplateArray) ToCheckRequiredTemplateArrayOutputWithContext(ctx context.Context) CheckRequiredTemplateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CheckRequiredTemplateArrayOutput)
}

func (i CheckRequiredTemplateArray) ToOutput(ctx context.Context) pulumix.Output[[]*CheckRequiredTemplate] {
	return pulumix.Output[[]*CheckRequiredTemplate]{
		OutputState: i.ToCheckRequiredTemplateArrayOutputWithContext(ctx).OutputState,
	}
}

// CheckRequiredTemplateMapInput is an input type that accepts CheckRequiredTemplateMap and CheckRequiredTemplateMapOutput values.
// You can construct a concrete instance of `CheckRequiredTemplateMapInput` via:
//
//	CheckRequiredTemplateMap{ "key": CheckRequiredTemplateArgs{...} }
type CheckRequiredTemplateMapInput interface {
	pulumi.Input

	ToCheckRequiredTemplateMapOutput() CheckRequiredTemplateMapOutput
	ToCheckRequiredTemplateMapOutputWithContext(context.Context) CheckRequiredTemplateMapOutput
}

type CheckRequiredTemplateMap map[string]CheckRequiredTemplateInput

func (CheckRequiredTemplateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CheckRequiredTemplate)(nil)).Elem()
}

func (i CheckRequiredTemplateMap) ToCheckRequiredTemplateMapOutput() CheckRequiredTemplateMapOutput {
	return i.ToCheckRequiredTemplateMapOutputWithContext(context.Background())
}

func (i CheckRequiredTemplateMap) ToCheckRequiredTemplateMapOutputWithContext(ctx context.Context) CheckRequiredTemplateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CheckRequiredTemplateMapOutput)
}

func (i CheckRequiredTemplateMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*CheckRequiredTemplate] {
	return pulumix.Output[map[string]*CheckRequiredTemplate]{
		OutputState: i.ToCheckRequiredTemplateMapOutputWithContext(ctx).OutputState,
	}
}

type CheckRequiredTemplateOutput struct{ *pulumi.OutputState }

func (CheckRequiredTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CheckRequiredTemplate)(nil)).Elem()
}

func (o CheckRequiredTemplateOutput) ToCheckRequiredTemplateOutput() CheckRequiredTemplateOutput {
	return o
}

func (o CheckRequiredTemplateOutput) ToCheckRequiredTemplateOutputWithContext(ctx context.Context) CheckRequiredTemplateOutput {
	return o
}

func (o CheckRequiredTemplateOutput) ToOutput(ctx context.Context) pulumix.Output[*CheckRequiredTemplate] {
	return pulumix.Output[*CheckRequiredTemplate]{
		OutputState: o.OutputState,
	}
}

// The project ID. Changing this forces a new Required Template Check to be created.
func (o CheckRequiredTemplateOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *CheckRequiredTemplate) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// One or more `requiredTemplate` blocks documented below.
func (o CheckRequiredTemplateOutput) RequiredTemplates() CheckRequiredTemplateRequiredTemplateArrayOutput {
	return o.ApplyT(func(v *CheckRequiredTemplate) CheckRequiredTemplateRequiredTemplateArrayOutput {
		return v.RequiredTemplates
	}).(CheckRequiredTemplateRequiredTemplateArrayOutput)
}

// The ID of the resource being protected by the check. Changing this forces a new Required Template Check to be created.
func (o CheckRequiredTemplateOutput) TargetResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *CheckRequiredTemplate) pulumi.StringOutput { return v.TargetResourceId }).(pulumi.StringOutput)
}

// The type of resource being protected by the check. Valid values: `endpoint`, `environment`, `queue`, `repository`, `securefile`, `variablegroup`. Changing this forces a new Required Template Check to be created.
func (o CheckRequiredTemplateOutput) TargetResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *CheckRequiredTemplate) pulumi.StringOutput { return v.TargetResourceType }).(pulumi.StringOutput)
}

type CheckRequiredTemplateArrayOutput struct{ *pulumi.OutputState }

func (CheckRequiredTemplateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CheckRequiredTemplate)(nil)).Elem()
}

func (o CheckRequiredTemplateArrayOutput) ToCheckRequiredTemplateArrayOutput() CheckRequiredTemplateArrayOutput {
	return o
}

func (o CheckRequiredTemplateArrayOutput) ToCheckRequiredTemplateArrayOutputWithContext(ctx context.Context) CheckRequiredTemplateArrayOutput {
	return o
}

func (o CheckRequiredTemplateArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*CheckRequiredTemplate] {
	return pulumix.Output[[]*CheckRequiredTemplate]{
		OutputState: o.OutputState,
	}
}

func (o CheckRequiredTemplateArrayOutput) Index(i pulumi.IntInput) CheckRequiredTemplateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CheckRequiredTemplate {
		return vs[0].([]*CheckRequiredTemplate)[vs[1].(int)]
	}).(CheckRequiredTemplateOutput)
}

type CheckRequiredTemplateMapOutput struct{ *pulumi.OutputState }

func (CheckRequiredTemplateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CheckRequiredTemplate)(nil)).Elem()
}

func (o CheckRequiredTemplateMapOutput) ToCheckRequiredTemplateMapOutput() CheckRequiredTemplateMapOutput {
	return o
}

func (o CheckRequiredTemplateMapOutput) ToCheckRequiredTemplateMapOutputWithContext(ctx context.Context) CheckRequiredTemplateMapOutput {
	return o
}

func (o CheckRequiredTemplateMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*CheckRequiredTemplate] {
	return pulumix.Output[map[string]*CheckRequiredTemplate]{
		OutputState: o.OutputState,
	}
}

func (o CheckRequiredTemplateMapOutput) MapIndex(k pulumi.StringInput) CheckRequiredTemplateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CheckRequiredTemplate {
		return vs[0].(map[string]*CheckRequiredTemplate)[vs[1].(string)]
	}).(CheckRequiredTemplateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CheckRequiredTemplateInput)(nil)).Elem(), &CheckRequiredTemplate{})
	pulumi.RegisterInputType(reflect.TypeOf((*CheckRequiredTemplateArrayInput)(nil)).Elem(), CheckRequiredTemplateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CheckRequiredTemplateMapInput)(nil)).Elem(), CheckRequiredTemplateMap{})
	pulumi.RegisterOutputType(CheckRequiredTemplateOutput{})
	pulumi.RegisterOutputType(CheckRequiredTemplateArrayOutput{})
	pulumi.RegisterOutputType(CheckRequiredTemplateMapOutput{})
}
