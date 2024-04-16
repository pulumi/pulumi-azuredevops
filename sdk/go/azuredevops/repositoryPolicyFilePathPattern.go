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

// Manage a file path pattern repository policy within Azure DevOps project.
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
//			exampleGit, err := azuredevops.NewGit(ctx, "example", &azuredevops.GitArgs{
//				ProjectId: example.ID(),
//				Name:      pulumi.String("Example Repository"),
//				Initialization: &azuredevops.GitInitializationArgs{
//					InitType: pulumi.String("Clean"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuredevops.NewRepositoryPolicyFilePathPattern(ctx, "example", &azuredevops.RepositoryPolicyFilePathPatternArgs{
//				ProjectId: example.ID(),
//				Enabled:   pulumi.Bool(true),
//				Blocking:  pulumi.Bool(true),
//				FilepathPatterns: pulumi.StringArray{
//					pulumi.String("*.go"),
//					pulumi.String("/home/test/*.ts"),
//				},
//				RepositoryIds: pulumi.StringArray{
//					exampleGit.ID(),
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
// <!--End PulumiCodeChooser -->
//
// # Set project level repository policy
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
//			_, err = azuredevops.NewRepositoryPolicyFilePathPattern(ctx, "examplep", &azuredevops.RepositoryPolicyFilePathPatternArgs{
//				ProjectId: example.ID(),
//				Enabled:   pulumi.Bool(true),
//				Blocking:  pulumi.Bool(true),
//				FilepathPatterns: pulumi.StringArray{
//					pulumi.String("*.go"),
//					pulumi.String("/home/test/*.ts"),
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
// <!--End PulumiCodeChooser -->
//
// ## Relevant Links
//
// - [Azure DevOps Service REST API 7.0 - Policy Configurations](https://docs.microsoft.com/en-us/rest/api/azure/devops/policy/configurations?view=azure-devops-rest-7.0)
//
// ## Import
//
// Azure DevOps repository policies can be imported using the projectID/policyID or projectName/policyID:
//
// ```sh
// $ pulumi import azuredevops:index/repositoryPolicyFilePathPattern:RepositoryPolicyFilePathPattern example 00000000-0000-0000-0000-000000000000/0
// ```
type RepositoryPolicyFilePathPattern struct {
	pulumi.CustomResourceState

	// A flag indicating if the policy should be blocking. Defaults to `true`.
	Blocking pulumi.BoolPtrOutput `pulumi:"blocking"`
	// A flag indicating if the policy should be enabled. Defaults to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// Block pushes from introducing file paths that match the following patterns. Exact paths begin with "/". You can specify exact paths and wildcards. You can also specify multiple paths using ";" as a separator. Paths prefixed with "!" are excluded. Order is important.
	FilepathPatterns pulumi.StringArrayOutput `pulumi:"filepathPatterns"`
	// The ID of the project in which the policy will be created.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Control whether the policy is enabled for the repository or the project. If `repositoryIds` not configured, the policy will be set to the project.
	RepositoryIds pulumi.StringArrayOutput `pulumi:"repositoryIds"`
}

// NewRepositoryPolicyFilePathPattern registers a new resource with the given unique name, arguments, and options.
func NewRepositoryPolicyFilePathPattern(ctx *pulumi.Context,
	name string, args *RepositoryPolicyFilePathPatternArgs, opts ...pulumi.ResourceOption) (*RepositoryPolicyFilePathPattern, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FilepathPatterns == nil {
		return nil, errors.New("invalid value for required argument 'FilepathPatterns'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RepositoryPolicyFilePathPattern
	err := ctx.RegisterResource("azuredevops:index/repositoryPolicyFilePathPattern:RepositoryPolicyFilePathPattern", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRepositoryPolicyFilePathPattern gets an existing RepositoryPolicyFilePathPattern resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRepositoryPolicyFilePathPattern(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RepositoryPolicyFilePathPatternState, opts ...pulumi.ResourceOption) (*RepositoryPolicyFilePathPattern, error) {
	var resource RepositoryPolicyFilePathPattern
	err := ctx.ReadResource("azuredevops:index/repositoryPolicyFilePathPattern:RepositoryPolicyFilePathPattern", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RepositoryPolicyFilePathPattern resources.
type repositoryPolicyFilePathPatternState struct {
	// A flag indicating if the policy should be blocking. Defaults to `true`.
	Blocking *bool `pulumi:"blocking"`
	// A flag indicating if the policy should be enabled. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// Block pushes from introducing file paths that match the following patterns. Exact paths begin with "/". You can specify exact paths and wildcards. You can also specify multiple paths using ";" as a separator. Paths prefixed with "!" are excluded. Order is important.
	FilepathPatterns []string `pulumi:"filepathPatterns"`
	// The ID of the project in which the policy will be created.
	ProjectId *string `pulumi:"projectId"`
	// Control whether the policy is enabled for the repository or the project. If `repositoryIds` not configured, the policy will be set to the project.
	RepositoryIds []string `pulumi:"repositoryIds"`
}

type RepositoryPolicyFilePathPatternState struct {
	// A flag indicating if the policy should be blocking. Defaults to `true`.
	Blocking pulumi.BoolPtrInput
	// A flag indicating if the policy should be enabled. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// Block pushes from introducing file paths that match the following patterns. Exact paths begin with "/". You can specify exact paths and wildcards. You can also specify multiple paths using ";" as a separator. Paths prefixed with "!" are excluded. Order is important.
	FilepathPatterns pulumi.StringArrayInput
	// The ID of the project in which the policy will be created.
	ProjectId pulumi.StringPtrInput
	// Control whether the policy is enabled for the repository or the project. If `repositoryIds` not configured, the policy will be set to the project.
	RepositoryIds pulumi.StringArrayInput
}

func (RepositoryPolicyFilePathPatternState) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryPolicyFilePathPatternState)(nil)).Elem()
}

type repositoryPolicyFilePathPatternArgs struct {
	// A flag indicating if the policy should be blocking. Defaults to `true`.
	Blocking *bool `pulumi:"blocking"`
	// A flag indicating if the policy should be enabled. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// Block pushes from introducing file paths that match the following patterns. Exact paths begin with "/". You can specify exact paths and wildcards. You can also specify multiple paths using ";" as a separator. Paths prefixed with "!" are excluded. Order is important.
	FilepathPatterns []string `pulumi:"filepathPatterns"`
	// The ID of the project in which the policy will be created.
	ProjectId string `pulumi:"projectId"`
	// Control whether the policy is enabled for the repository or the project. If `repositoryIds` not configured, the policy will be set to the project.
	RepositoryIds []string `pulumi:"repositoryIds"`
}

// The set of arguments for constructing a RepositoryPolicyFilePathPattern resource.
type RepositoryPolicyFilePathPatternArgs struct {
	// A flag indicating if the policy should be blocking. Defaults to `true`.
	Blocking pulumi.BoolPtrInput
	// A flag indicating if the policy should be enabled. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// Block pushes from introducing file paths that match the following patterns. Exact paths begin with "/". You can specify exact paths and wildcards. You can also specify multiple paths using ";" as a separator. Paths prefixed with "!" are excluded. Order is important.
	FilepathPatterns pulumi.StringArrayInput
	// The ID of the project in which the policy will be created.
	ProjectId pulumi.StringInput
	// Control whether the policy is enabled for the repository or the project. If `repositoryIds` not configured, the policy will be set to the project.
	RepositoryIds pulumi.StringArrayInput
}

func (RepositoryPolicyFilePathPatternArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryPolicyFilePathPatternArgs)(nil)).Elem()
}

type RepositoryPolicyFilePathPatternInput interface {
	pulumi.Input

	ToRepositoryPolicyFilePathPatternOutput() RepositoryPolicyFilePathPatternOutput
	ToRepositoryPolicyFilePathPatternOutputWithContext(ctx context.Context) RepositoryPolicyFilePathPatternOutput
}

func (*RepositoryPolicyFilePathPattern) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryPolicyFilePathPattern)(nil)).Elem()
}

func (i *RepositoryPolicyFilePathPattern) ToRepositoryPolicyFilePathPatternOutput() RepositoryPolicyFilePathPatternOutput {
	return i.ToRepositoryPolicyFilePathPatternOutputWithContext(context.Background())
}

func (i *RepositoryPolicyFilePathPattern) ToRepositoryPolicyFilePathPatternOutputWithContext(ctx context.Context) RepositoryPolicyFilePathPatternOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryPolicyFilePathPatternOutput)
}

// RepositoryPolicyFilePathPatternArrayInput is an input type that accepts RepositoryPolicyFilePathPatternArray and RepositoryPolicyFilePathPatternArrayOutput values.
// You can construct a concrete instance of `RepositoryPolicyFilePathPatternArrayInput` via:
//
//	RepositoryPolicyFilePathPatternArray{ RepositoryPolicyFilePathPatternArgs{...} }
type RepositoryPolicyFilePathPatternArrayInput interface {
	pulumi.Input

	ToRepositoryPolicyFilePathPatternArrayOutput() RepositoryPolicyFilePathPatternArrayOutput
	ToRepositoryPolicyFilePathPatternArrayOutputWithContext(context.Context) RepositoryPolicyFilePathPatternArrayOutput
}

type RepositoryPolicyFilePathPatternArray []RepositoryPolicyFilePathPatternInput

func (RepositoryPolicyFilePathPatternArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RepositoryPolicyFilePathPattern)(nil)).Elem()
}

func (i RepositoryPolicyFilePathPatternArray) ToRepositoryPolicyFilePathPatternArrayOutput() RepositoryPolicyFilePathPatternArrayOutput {
	return i.ToRepositoryPolicyFilePathPatternArrayOutputWithContext(context.Background())
}

func (i RepositoryPolicyFilePathPatternArray) ToRepositoryPolicyFilePathPatternArrayOutputWithContext(ctx context.Context) RepositoryPolicyFilePathPatternArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryPolicyFilePathPatternArrayOutput)
}

// RepositoryPolicyFilePathPatternMapInput is an input type that accepts RepositoryPolicyFilePathPatternMap and RepositoryPolicyFilePathPatternMapOutput values.
// You can construct a concrete instance of `RepositoryPolicyFilePathPatternMapInput` via:
//
//	RepositoryPolicyFilePathPatternMap{ "key": RepositoryPolicyFilePathPatternArgs{...} }
type RepositoryPolicyFilePathPatternMapInput interface {
	pulumi.Input

	ToRepositoryPolicyFilePathPatternMapOutput() RepositoryPolicyFilePathPatternMapOutput
	ToRepositoryPolicyFilePathPatternMapOutputWithContext(context.Context) RepositoryPolicyFilePathPatternMapOutput
}

type RepositoryPolicyFilePathPatternMap map[string]RepositoryPolicyFilePathPatternInput

func (RepositoryPolicyFilePathPatternMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RepositoryPolicyFilePathPattern)(nil)).Elem()
}

func (i RepositoryPolicyFilePathPatternMap) ToRepositoryPolicyFilePathPatternMapOutput() RepositoryPolicyFilePathPatternMapOutput {
	return i.ToRepositoryPolicyFilePathPatternMapOutputWithContext(context.Background())
}

func (i RepositoryPolicyFilePathPatternMap) ToRepositoryPolicyFilePathPatternMapOutputWithContext(ctx context.Context) RepositoryPolicyFilePathPatternMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryPolicyFilePathPatternMapOutput)
}

type RepositoryPolicyFilePathPatternOutput struct{ *pulumi.OutputState }

func (RepositoryPolicyFilePathPatternOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryPolicyFilePathPattern)(nil)).Elem()
}

func (o RepositoryPolicyFilePathPatternOutput) ToRepositoryPolicyFilePathPatternOutput() RepositoryPolicyFilePathPatternOutput {
	return o
}

func (o RepositoryPolicyFilePathPatternOutput) ToRepositoryPolicyFilePathPatternOutputWithContext(ctx context.Context) RepositoryPolicyFilePathPatternOutput {
	return o
}

// A flag indicating if the policy should be blocking. Defaults to `true`.
func (o RepositoryPolicyFilePathPatternOutput) Blocking() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RepositoryPolicyFilePathPattern) pulumi.BoolPtrOutput { return v.Blocking }).(pulumi.BoolPtrOutput)
}

// A flag indicating if the policy should be enabled. Defaults to `true`.
func (o RepositoryPolicyFilePathPatternOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RepositoryPolicyFilePathPattern) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// Block pushes from introducing file paths that match the following patterns. Exact paths begin with "/". You can specify exact paths and wildcards. You can also specify multiple paths using ";" as a separator. Paths prefixed with "!" are excluded. Order is important.
func (o RepositoryPolicyFilePathPatternOutput) FilepathPatterns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RepositoryPolicyFilePathPattern) pulumi.StringArrayOutput { return v.FilepathPatterns }).(pulumi.StringArrayOutput)
}

// The ID of the project in which the policy will be created.
func (o RepositoryPolicyFilePathPatternOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryPolicyFilePathPattern) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// Control whether the policy is enabled for the repository or the project. If `repositoryIds` not configured, the policy will be set to the project.
func (o RepositoryPolicyFilePathPatternOutput) RepositoryIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RepositoryPolicyFilePathPattern) pulumi.StringArrayOutput { return v.RepositoryIds }).(pulumi.StringArrayOutput)
}

type RepositoryPolicyFilePathPatternArrayOutput struct{ *pulumi.OutputState }

func (RepositoryPolicyFilePathPatternArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RepositoryPolicyFilePathPattern)(nil)).Elem()
}

func (o RepositoryPolicyFilePathPatternArrayOutput) ToRepositoryPolicyFilePathPatternArrayOutput() RepositoryPolicyFilePathPatternArrayOutput {
	return o
}

func (o RepositoryPolicyFilePathPatternArrayOutput) ToRepositoryPolicyFilePathPatternArrayOutputWithContext(ctx context.Context) RepositoryPolicyFilePathPatternArrayOutput {
	return o
}

func (o RepositoryPolicyFilePathPatternArrayOutput) Index(i pulumi.IntInput) RepositoryPolicyFilePathPatternOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RepositoryPolicyFilePathPattern {
		return vs[0].([]*RepositoryPolicyFilePathPattern)[vs[1].(int)]
	}).(RepositoryPolicyFilePathPatternOutput)
}

type RepositoryPolicyFilePathPatternMapOutput struct{ *pulumi.OutputState }

func (RepositoryPolicyFilePathPatternMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RepositoryPolicyFilePathPattern)(nil)).Elem()
}

func (o RepositoryPolicyFilePathPatternMapOutput) ToRepositoryPolicyFilePathPatternMapOutput() RepositoryPolicyFilePathPatternMapOutput {
	return o
}

func (o RepositoryPolicyFilePathPatternMapOutput) ToRepositoryPolicyFilePathPatternMapOutputWithContext(ctx context.Context) RepositoryPolicyFilePathPatternMapOutput {
	return o
}

func (o RepositoryPolicyFilePathPatternMapOutput) MapIndex(k pulumi.StringInput) RepositoryPolicyFilePathPatternOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RepositoryPolicyFilePathPattern {
		return vs[0].(map[string]*RepositoryPolicyFilePathPattern)[vs[1].(string)]
	}).(RepositoryPolicyFilePathPatternOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryPolicyFilePathPatternInput)(nil)).Elem(), &RepositoryPolicyFilePathPattern{})
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryPolicyFilePathPatternArrayInput)(nil)).Elem(), RepositoryPolicyFilePathPatternArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryPolicyFilePathPatternMapInput)(nil)).Elem(), RepositoryPolicyFilePathPatternMap{})
	pulumi.RegisterOutputType(RepositoryPolicyFilePathPatternOutput{})
	pulumi.RegisterOutputType(RepositoryPolicyFilePathPatternArrayOutput{})
	pulumi.RegisterOutputType(RepositoryPolicyFilePathPatternMapOutput{})
}
