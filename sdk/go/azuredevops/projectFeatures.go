// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages features for Azure DevOps projects
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
// 		tf_project_test_001, err := azuredevops.LookupProject(ctx, &GetProjectArgs{
// 			Name: pulumi.StringRef("Test Project"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = azuredevops.NewProjectFeatures(ctx, "my-project-features", &azuredevops.ProjectFeaturesArgs{
// 			ProjectId: pulumi.String(tf_project_test_001.Id),
// 			Features: pulumi.StringMap{
// 				"testplans": pulumi.String("disabled"),
// 				"artifacts": pulumi.String("enabled"),
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
// No official documentation available
//
// ## PAT Permissions Required
//
// - **Project & Team**: Read, Write, & Manage
//
// ## Import
//
// Azure DevOps feature settings can be imported using the project id, e.g.
//
// ```sh
//  $ pulumi import azuredevops:index/projectFeatures:ProjectFeatures project_id 00000000-0000-0000-0000-000000000000
// ```
type ProjectFeatures struct {
	pulumi.CustomResourceState

	// Defines the status (`enabled`, `disabled`) of the project features.\
	// Valid features `boards`, `repositories`, `pipelines`, `testplans`, `artifacts`
	Features  pulumi.StringMapOutput `pulumi:"features"`
	ProjectId pulumi.StringOutput    `pulumi:"projectId"`
}

// NewProjectFeatures registers a new resource with the given unique name, arguments, and options.
func NewProjectFeatures(ctx *pulumi.Context,
	name string, args *ProjectFeaturesArgs, opts ...pulumi.ResourceOption) (*ProjectFeatures, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Features == nil {
		return nil, errors.New("invalid value for required argument 'Features'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azuredevops:Core/projectFeatures:ProjectFeatures"),
		},
	})
	opts = append(opts, aliases)
	var resource ProjectFeatures
	err := ctx.RegisterResource("azuredevops:index/projectFeatures:ProjectFeatures", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProjectFeatures gets an existing ProjectFeatures resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProjectFeatures(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectFeaturesState, opts ...pulumi.ResourceOption) (*ProjectFeatures, error) {
	var resource ProjectFeatures
	err := ctx.ReadResource("azuredevops:index/projectFeatures:ProjectFeatures", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProjectFeatures resources.
type projectFeaturesState struct {
	// Defines the status (`enabled`, `disabled`) of the project features.\
	// Valid features `boards`, `repositories`, `pipelines`, `testplans`, `artifacts`
	Features  map[string]string `pulumi:"features"`
	ProjectId *string           `pulumi:"projectId"`
}

type ProjectFeaturesState struct {
	// Defines the status (`enabled`, `disabled`) of the project features.\
	// Valid features `boards`, `repositories`, `pipelines`, `testplans`, `artifacts`
	Features  pulumi.StringMapInput
	ProjectId pulumi.StringPtrInput
}

func (ProjectFeaturesState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectFeaturesState)(nil)).Elem()
}

type projectFeaturesArgs struct {
	// Defines the status (`enabled`, `disabled`) of the project features.\
	// Valid features `boards`, `repositories`, `pipelines`, `testplans`, `artifacts`
	Features  map[string]string `pulumi:"features"`
	ProjectId string            `pulumi:"projectId"`
}

// The set of arguments for constructing a ProjectFeatures resource.
type ProjectFeaturesArgs struct {
	// Defines the status (`enabled`, `disabled`) of the project features.\
	// Valid features `boards`, `repositories`, `pipelines`, `testplans`, `artifacts`
	Features  pulumi.StringMapInput
	ProjectId pulumi.StringInput
}

func (ProjectFeaturesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectFeaturesArgs)(nil)).Elem()
}

type ProjectFeaturesInput interface {
	pulumi.Input

	ToProjectFeaturesOutput() ProjectFeaturesOutput
	ToProjectFeaturesOutputWithContext(ctx context.Context) ProjectFeaturesOutput
}

func (*ProjectFeatures) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectFeatures)(nil)).Elem()
}

func (i *ProjectFeatures) ToProjectFeaturesOutput() ProjectFeaturesOutput {
	return i.ToProjectFeaturesOutputWithContext(context.Background())
}

func (i *ProjectFeatures) ToProjectFeaturesOutputWithContext(ctx context.Context) ProjectFeaturesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectFeaturesOutput)
}

// ProjectFeaturesArrayInput is an input type that accepts ProjectFeaturesArray and ProjectFeaturesArrayOutput values.
// You can construct a concrete instance of `ProjectFeaturesArrayInput` via:
//
//          ProjectFeaturesArray{ ProjectFeaturesArgs{...} }
type ProjectFeaturesArrayInput interface {
	pulumi.Input

	ToProjectFeaturesArrayOutput() ProjectFeaturesArrayOutput
	ToProjectFeaturesArrayOutputWithContext(context.Context) ProjectFeaturesArrayOutput
}

type ProjectFeaturesArray []ProjectFeaturesInput

func (ProjectFeaturesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectFeatures)(nil)).Elem()
}

func (i ProjectFeaturesArray) ToProjectFeaturesArrayOutput() ProjectFeaturesArrayOutput {
	return i.ToProjectFeaturesArrayOutputWithContext(context.Background())
}

func (i ProjectFeaturesArray) ToProjectFeaturesArrayOutputWithContext(ctx context.Context) ProjectFeaturesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectFeaturesArrayOutput)
}

// ProjectFeaturesMapInput is an input type that accepts ProjectFeaturesMap and ProjectFeaturesMapOutput values.
// You can construct a concrete instance of `ProjectFeaturesMapInput` via:
//
//          ProjectFeaturesMap{ "key": ProjectFeaturesArgs{...} }
type ProjectFeaturesMapInput interface {
	pulumi.Input

	ToProjectFeaturesMapOutput() ProjectFeaturesMapOutput
	ToProjectFeaturesMapOutputWithContext(context.Context) ProjectFeaturesMapOutput
}

type ProjectFeaturesMap map[string]ProjectFeaturesInput

func (ProjectFeaturesMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectFeatures)(nil)).Elem()
}

func (i ProjectFeaturesMap) ToProjectFeaturesMapOutput() ProjectFeaturesMapOutput {
	return i.ToProjectFeaturesMapOutputWithContext(context.Background())
}

func (i ProjectFeaturesMap) ToProjectFeaturesMapOutputWithContext(ctx context.Context) ProjectFeaturesMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectFeaturesMapOutput)
}

type ProjectFeaturesOutput struct{ *pulumi.OutputState }

func (ProjectFeaturesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectFeatures)(nil)).Elem()
}

func (o ProjectFeaturesOutput) ToProjectFeaturesOutput() ProjectFeaturesOutput {
	return o
}

func (o ProjectFeaturesOutput) ToProjectFeaturesOutputWithContext(ctx context.Context) ProjectFeaturesOutput {
	return o
}

type ProjectFeaturesArrayOutput struct{ *pulumi.OutputState }

func (ProjectFeaturesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectFeatures)(nil)).Elem()
}

func (o ProjectFeaturesArrayOutput) ToProjectFeaturesArrayOutput() ProjectFeaturesArrayOutput {
	return o
}

func (o ProjectFeaturesArrayOutput) ToProjectFeaturesArrayOutputWithContext(ctx context.Context) ProjectFeaturesArrayOutput {
	return o
}

func (o ProjectFeaturesArrayOutput) Index(i pulumi.IntInput) ProjectFeaturesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ProjectFeatures {
		return vs[0].([]*ProjectFeatures)[vs[1].(int)]
	}).(ProjectFeaturesOutput)
}

type ProjectFeaturesMapOutput struct{ *pulumi.OutputState }

func (ProjectFeaturesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectFeatures)(nil)).Elem()
}

func (o ProjectFeaturesMapOutput) ToProjectFeaturesMapOutput() ProjectFeaturesMapOutput {
	return o
}

func (o ProjectFeaturesMapOutput) ToProjectFeaturesMapOutputWithContext(ctx context.Context) ProjectFeaturesMapOutput {
	return o
}

func (o ProjectFeaturesMapOutput) MapIndex(k pulumi.StringInput) ProjectFeaturesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ProjectFeatures {
		return vs[0].(map[string]*ProjectFeatures)[vs[1].(string)]
	}).(ProjectFeaturesOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectFeaturesInput)(nil)).Elem(), &ProjectFeatures{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectFeaturesArrayInput)(nil)).Elem(), ProjectFeaturesArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectFeaturesMapInput)(nil)).Elem(), ProjectFeaturesMap{})
	pulumi.RegisterOutputType(ProjectFeaturesOutput{})
	pulumi.RegisterOutputType(ProjectFeaturesArrayOutput{})
	pulumi.RegisterOutputType(ProjectFeaturesMapOutput{})
}
