// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azuredevops/sdk/v2/go/azuredevops/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages Pipeline Settings for Azure DevOps projects
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
//			_, err = azuredevops.NewProjectPipelineSettings(ctx, "exampleProjectPipelineSettings", &azuredevops.ProjectPipelineSettingsArgs{
//				ProjectId:                        exampleProject.ID(),
//				EnforceJobScope:                  pulumi.Bool(true),
//				EnforceReferencedRepoScopedToken: pulumi.Bool(false),
//				EnforceSettableVar:               pulumi.Bool(true),
//				PublishPipelineMetadata:          pulumi.Bool(false),
//				StatusBadgesArePrivate:           pulumi.Bool(true),
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
// # No official documentation available
//
// ## PAT Permissions Required
//
// - Full Access
//
// ## Import
//
// Azure DevOps feature settings can be imported using the project id, e.g.
//
// ```sh
//
//	$ pulumi import azuredevops:index/projectPipelineSettings:ProjectPipelineSettings example 00000000-0000-0000-0000-000000000000
//
// ```
type ProjectPipelineSettings struct {
	pulumi.CustomResourceState

	// Limit job authorization scope to current project for non-release pipelines.
	EnforceJobScope pulumi.BoolOutput `pulumi:"enforceJobScope"`
	// Limit job authorization scope to current project for release pipelines.**NOTE:**\
	// The settings at the organization will override settings specified on the project.
	// For example, if `enforceJobScope` is true at the organization, the `ProjectPipelineSettings` resource cannot set it to false.
	// In this scenario, the plan will always show that the resource is trying to change `enforceJobScope` from `true` to `false`.
	EnforceJobScopeForRelease pulumi.BoolOutput `pulumi:"enforceJobScopeForRelease"`
	// Protect access to repositories in YAML pipelines.
	EnforceReferencedRepoScopedToken pulumi.BoolOutput `pulumi:"enforceReferencedRepoScopedToken"`
	// Limit variables that can be set at queue time.
	EnforceSettableVar pulumi.BoolOutput `pulumi:"enforceSettableVar"`
	// The `id` of the project for which the project pipeline settings will be managed.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Publish metadata from pipelines.
	PublishPipelineMetadata pulumi.BoolOutput `pulumi:"publishPipelineMetadata"`
	// Disable anonymous access to badges.
	StatusBadgesArePrivate pulumi.BoolOutput `pulumi:"statusBadgesArePrivate"`
}

// NewProjectPipelineSettings registers a new resource with the given unique name, arguments, and options.
func NewProjectPipelineSettings(ctx *pulumi.Context,
	name string, args *ProjectPipelineSettingsArgs, opts ...pulumi.ResourceOption) (*ProjectPipelineSettings, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ProjectPipelineSettings
	err := ctx.RegisterResource("azuredevops:index/projectPipelineSettings:ProjectPipelineSettings", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProjectPipelineSettings gets an existing ProjectPipelineSettings resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProjectPipelineSettings(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectPipelineSettingsState, opts ...pulumi.ResourceOption) (*ProjectPipelineSettings, error) {
	var resource ProjectPipelineSettings
	err := ctx.ReadResource("azuredevops:index/projectPipelineSettings:ProjectPipelineSettings", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProjectPipelineSettings resources.
type projectPipelineSettingsState struct {
	// Limit job authorization scope to current project for non-release pipelines.
	EnforceJobScope *bool `pulumi:"enforceJobScope"`
	// Limit job authorization scope to current project for release pipelines.**NOTE:**\
	// The settings at the organization will override settings specified on the project.
	// For example, if `enforceJobScope` is true at the organization, the `ProjectPipelineSettings` resource cannot set it to false.
	// In this scenario, the plan will always show that the resource is trying to change `enforceJobScope` from `true` to `false`.
	EnforceJobScopeForRelease *bool `pulumi:"enforceJobScopeForRelease"`
	// Protect access to repositories in YAML pipelines.
	EnforceReferencedRepoScopedToken *bool `pulumi:"enforceReferencedRepoScopedToken"`
	// Limit variables that can be set at queue time.
	EnforceSettableVar *bool `pulumi:"enforceSettableVar"`
	// The `id` of the project for which the project pipeline settings will be managed.
	ProjectId *string `pulumi:"projectId"`
	// Publish metadata from pipelines.
	PublishPipelineMetadata *bool `pulumi:"publishPipelineMetadata"`
	// Disable anonymous access to badges.
	StatusBadgesArePrivate *bool `pulumi:"statusBadgesArePrivate"`
}

type ProjectPipelineSettingsState struct {
	// Limit job authorization scope to current project for non-release pipelines.
	EnforceJobScope pulumi.BoolPtrInput
	// Limit job authorization scope to current project for release pipelines.**NOTE:**\
	// The settings at the organization will override settings specified on the project.
	// For example, if `enforceJobScope` is true at the organization, the `ProjectPipelineSettings` resource cannot set it to false.
	// In this scenario, the plan will always show that the resource is trying to change `enforceJobScope` from `true` to `false`.
	EnforceJobScopeForRelease pulumi.BoolPtrInput
	// Protect access to repositories in YAML pipelines.
	EnforceReferencedRepoScopedToken pulumi.BoolPtrInput
	// Limit variables that can be set at queue time.
	EnforceSettableVar pulumi.BoolPtrInput
	// The `id` of the project for which the project pipeline settings will be managed.
	ProjectId pulumi.StringPtrInput
	// Publish metadata from pipelines.
	PublishPipelineMetadata pulumi.BoolPtrInput
	// Disable anonymous access to badges.
	StatusBadgesArePrivate pulumi.BoolPtrInput
}

func (ProjectPipelineSettingsState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectPipelineSettingsState)(nil)).Elem()
}

type projectPipelineSettingsArgs struct {
	// Limit job authorization scope to current project for non-release pipelines.
	EnforceJobScope *bool `pulumi:"enforceJobScope"`
	// Limit job authorization scope to current project for release pipelines.**NOTE:**\
	// The settings at the organization will override settings specified on the project.
	// For example, if `enforceJobScope` is true at the organization, the `ProjectPipelineSettings` resource cannot set it to false.
	// In this scenario, the plan will always show that the resource is trying to change `enforceJobScope` from `true` to `false`.
	EnforceJobScopeForRelease *bool `pulumi:"enforceJobScopeForRelease"`
	// Protect access to repositories in YAML pipelines.
	EnforceReferencedRepoScopedToken *bool `pulumi:"enforceReferencedRepoScopedToken"`
	// Limit variables that can be set at queue time.
	EnforceSettableVar *bool `pulumi:"enforceSettableVar"`
	// The `id` of the project for which the project pipeline settings will be managed.
	ProjectId string `pulumi:"projectId"`
	// Publish metadata from pipelines.
	PublishPipelineMetadata *bool `pulumi:"publishPipelineMetadata"`
	// Disable anonymous access to badges.
	StatusBadgesArePrivate *bool `pulumi:"statusBadgesArePrivate"`
}

// The set of arguments for constructing a ProjectPipelineSettings resource.
type ProjectPipelineSettingsArgs struct {
	// Limit job authorization scope to current project for non-release pipelines.
	EnforceJobScope pulumi.BoolPtrInput
	// Limit job authorization scope to current project for release pipelines.**NOTE:**\
	// The settings at the organization will override settings specified on the project.
	// For example, if `enforceJobScope` is true at the organization, the `ProjectPipelineSettings` resource cannot set it to false.
	// In this scenario, the plan will always show that the resource is trying to change `enforceJobScope` from `true` to `false`.
	EnforceJobScopeForRelease pulumi.BoolPtrInput
	// Protect access to repositories in YAML pipelines.
	EnforceReferencedRepoScopedToken pulumi.BoolPtrInput
	// Limit variables that can be set at queue time.
	EnforceSettableVar pulumi.BoolPtrInput
	// The `id` of the project for which the project pipeline settings will be managed.
	ProjectId pulumi.StringInput
	// Publish metadata from pipelines.
	PublishPipelineMetadata pulumi.BoolPtrInput
	// Disable anonymous access to badges.
	StatusBadgesArePrivate pulumi.BoolPtrInput
}

func (ProjectPipelineSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectPipelineSettingsArgs)(nil)).Elem()
}

type ProjectPipelineSettingsInput interface {
	pulumi.Input

	ToProjectPipelineSettingsOutput() ProjectPipelineSettingsOutput
	ToProjectPipelineSettingsOutputWithContext(ctx context.Context) ProjectPipelineSettingsOutput
}

func (*ProjectPipelineSettings) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectPipelineSettings)(nil)).Elem()
}

func (i *ProjectPipelineSettings) ToProjectPipelineSettingsOutput() ProjectPipelineSettingsOutput {
	return i.ToProjectPipelineSettingsOutputWithContext(context.Background())
}

func (i *ProjectPipelineSettings) ToProjectPipelineSettingsOutputWithContext(ctx context.Context) ProjectPipelineSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectPipelineSettingsOutput)
}

// ProjectPipelineSettingsArrayInput is an input type that accepts ProjectPipelineSettingsArray and ProjectPipelineSettingsArrayOutput values.
// You can construct a concrete instance of `ProjectPipelineSettingsArrayInput` via:
//
//	ProjectPipelineSettingsArray{ ProjectPipelineSettingsArgs{...} }
type ProjectPipelineSettingsArrayInput interface {
	pulumi.Input

	ToProjectPipelineSettingsArrayOutput() ProjectPipelineSettingsArrayOutput
	ToProjectPipelineSettingsArrayOutputWithContext(context.Context) ProjectPipelineSettingsArrayOutput
}

type ProjectPipelineSettingsArray []ProjectPipelineSettingsInput

func (ProjectPipelineSettingsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectPipelineSettings)(nil)).Elem()
}

func (i ProjectPipelineSettingsArray) ToProjectPipelineSettingsArrayOutput() ProjectPipelineSettingsArrayOutput {
	return i.ToProjectPipelineSettingsArrayOutputWithContext(context.Background())
}

func (i ProjectPipelineSettingsArray) ToProjectPipelineSettingsArrayOutputWithContext(ctx context.Context) ProjectPipelineSettingsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectPipelineSettingsArrayOutput)
}

// ProjectPipelineSettingsMapInput is an input type that accepts ProjectPipelineSettingsMap and ProjectPipelineSettingsMapOutput values.
// You can construct a concrete instance of `ProjectPipelineSettingsMapInput` via:
//
//	ProjectPipelineSettingsMap{ "key": ProjectPipelineSettingsArgs{...} }
type ProjectPipelineSettingsMapInput interface {
	pulumi.Input

	ToProjectPipelineSettingsMapOutput() ProjectPipelineSettingsMapOutput
	ToProjectPipelineSettingsMapOutputWithContext(context.Context) ProjectPipelineSettingsMapOutput
}

type ProjectPipelineSettingsMap map[string]ProjectPipelineSettingsInput

func (ProjectPipelineSettingsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectPipelineSettings)(nil)).Elem()
}

func (i ProjectPipelineSettingsMap) ToProjectPipelineSettingsMapOutput() ProjectPipelineSettingsMapOutput {
	return i.ToProjectPipelineSettingsMapOutputWithContext(context.Background())
}

func (i ProjectPipelineSettingsMap) ToProjectPipelineSettingsMapOutputWithContext(ctx context.Context) ProjectPipelineSettingsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectPipelineSettingsMapOutput)
}

type ProjectPipelineSettingsOutput struct{ *pulumi.OutputState }

func (ProjectPipelineSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectPipelineSettings)(nil)).Elem()
}

func (o ProjectPipelineSettingsOutput) ToProjectPipelineSettingsOutput() ProjectPipelineSettingsOutput {
	return o
}

func (o ProjectPipelineSettingsOutput) ToProjectPipelineSettingsOutputWithContext(ctx context.Context) ProjectPipelineSettingsOutput {
	return o
}

// Limit job authorization scope to current project for non-release pipelines.
func (o ProjectPipelineSettingsOutput) EnforceJobScope() pulumi.BoolOutput {
	return o.ApplyT(func(v *ProjectPipelineSettings) pulumi.BoolOutput { return v.EnforceJobScope }).(pulumi.BoolOutput)
}

// Limit job authorization scope to current project for release pipelines.**NOTE:**\
// The settings at the organization will override settings specified on the project.
// For example, if `enforceJobScope` is true at the organization, the `ProjectPipelineSettings` resource cannot set it to false.
// In this scenario, the plan will always show that the resource is trying to change `enforceJobScope` from `true` to `false`.
func (o ProjectPipelineSettingsOutput) EnforceJobScopeForRelease() pulumi.BoolOutput {
	return o.ApplyT(func(v *ProjectPipelineSettings) pulumi.BoolOutput { return v.EnforceJobScopeForRelease }).(pulumi.BoolOutput)
}

// Protect access to repositories in YAML pipelines.
func (o ProjectPipelineSettingsOutput) EnforceReferencedRepoScopedToken() pulumi.BoolOutput {
	return o.ApplyT(func(v *ProjectPipelineSettings) pulumi.BoolOutput { return v.EnforceReferencedRepoScopedToken }).(pulumi.BoolOutput)
}

// Limit variables that can be set at queue time.
func (o ProjectPipelineSettingsOutput) EnforceSettableVar() pulumi.BoolOutput {
	return o.ApplyT(func(v *ProjectPipelineSettings) pulumi.BoolOutput { return v.EnforceSettableVar }).(pulumi.BoolOutput)
}

// The `id` of the project for which the project pipeline settings will be managed.
func (o ProjectPipelineSettingsOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectPipelineSettings) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// Publish metadata from pipelines.
func (o ProjectPipelineSettingsOutput) PublishPipelineMetadata() pulumi.BoolOutput {
	return o.ApplyT(func(v *ProjectPipelineSettings) pulumi.BoolOutput { return v.PublishPipelineMetadata }).(pulumi.BoolOutput)
}

// Disable anonymous access to badges.
func (o ProjectPipelineSettingsOutput) StatusBadgesArePrivate() pulumi.BoolOutput {
	return o.ApplyT(func(v *ProjectPipelineSettings) pulumi.BoolOutput { return v.StatusBadgesArePrivate }).(pulumi.BoolOutput)
}

type ProjectPipelineSettingsArrayOutput struct{ *pulumi.OutputState }

func (ProjectPipelineSettingsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectPipelineSettings)(nil)).Elem()
}

func (o ProjectPipelineSettingsArrayOutput) ToProjectPipelineSettingsArrayOutput() ProjectPipelineSettingsArrayOutput {
	return o
}

func (o ProjectPipelineSettingsArrayOutput) ToProjectPipelineSettingsArrayOutputWithContext(ctx context.Context) ProjectPipelineSettingsArrayOutput {
	return o
}

func (o ProjectPipelineSettingsArrayOutput) Index(i pulumi.IntInput) ProjectPipelineSettingsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ProjectPipelineSettings {
		return vs[0].([]*ProjectPipelineSettings)[vs[1].(int)]
	}).(ProjectPipelineSettingsOutput)
}

type ProjectPipelineSettingsMapOutput struct{ *pulumi.OutputState }

func (ProjectPipelineSettingsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectPipelineSettings)(nil)).Elem()
}

func (o ProjectPipelineSettingsMapOutput) ToProjectPipelineSettingsMapOutput() ProjectPipelineSettingsMapOutput {
	return o
}

func (o ProjectPipelineSettingsMapOutput) ToProjectPipelineSettingsMapOutputWithContext(ctx context.Context) ProjectPipelineSettingsMapOutput {
	return o
}

func (o ProjectPipelineSettingsMapOutput) MapIndex(k pulumi.StringInput) ProjectPipelineSettingsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ProjectPipelineSettings {
		return vs[0].(map[string]*ProjectPipelineSettings)[vs[1].(string)]
	}).(ProjectPipelineSettingsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectPipelineSettingsInput)(nil)).Elem(), &ProjectPipelineSettings{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectPipelineSettingsArrayInput)(nil)).Elem(), ProjectPipelineSettingsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectPipelineSettingsMapInput)(nil)).Elem(), ProjectPipelineSettingsMap{})
	pulumi.RegisterOutputType(ProjectPipelineSettingsOutput{})
	pulumi.RegisterOutputType(ProjectPipelineSettingsArrayOutput{})
	pulumi.RegisterOutputType(ProjectPipelineSettingsMapOutput{})
}
