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

// Manages a case enforcement repository policy within Azure DevOps project.
//
// > If both project and project policy are enabled, the project policy has high priority.
//
// ## Relevant Links
//
// - [Azure DevOps Service REST API 7.0 - Policy Configurations](https://docs.microsoft.com/en-us/rest/api/azure/devops/policy/configurations?view=azure-devops-rest-7.0)
//
// ## Import
//
// Azure DevOps repository policies can be imported using the projectID/policyID or projectName/policyID
//
// ```sh
//
//	$ pulumi import azuredevops:index/repositoryPolicyCaseEnforcement:RepositoryPolicyCaseEnforcement example 00000000-0000-0000-0000-000000000000/0
//
// ```
type RepositoryPolicyCaseEnforcement struct {
	pulumi.CustomResourceState

	// A flag indicating if the policy should be blocking. Defaults to `true`.
	Blocking pulumi.BoolPtrOutput `pulumi:"blocking"`
	// A flag indicating if the policy should be enabled. Defaults to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// Avoid case-sensitivity conflicts by blocking pushes that change name casing on files, folders, branches, and tags.
	EnforceConsistentCase pulumi.BoolOutput `pulumi:"enforceConsistentCase"`
	// The ID of the project in which the policy will be created.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Control whether the policy is enabled for the repository or the project. If `repositoryIds` not configured, the policy will be set to the project.
	RepositoryIds pulumi.StringArrayOutput `pulumi:"repositoryIds"`
}

// NewRepositoryPolicyCaseEnforcement registers a new resource with the given unique name, arguments, and options.
func NewRepositoryPolicyCaseEnforcement(ctx *pulumi.Context,
	name string, args *RepositoryPolicyCaseEnforcementArgs, opts ...pulumi.ResourceOption) (*RepositoryPolicyCaseEnforcement, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EnforceConsistentCase == nil {
		return nil, errors.New("invalid value for required argument 'EnforceConsistentCase'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RepositoryPolicyCaseEnforcement
	err := ctx.RegisterResource("azuredevops:index/repositoryPolicyCaseEnforcement:RepositoryPolicyCaseEnforcement", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRepositoryPolicyCaseEnforcement gets an existing RepositoryPolicyCaseEnforcement resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRepositoryPolicyCaseEnforcement(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RepositoryPolicyCaseEnforcementState, opts ...pulumi.ResourceOption) (*RepositoryPolicyCaseEnforcement, error) {
	var resource RepositoryPolicyCaseEnforcement
	err := ctx.ReadResource("azuredevops:index/repositoryPolicyCaseEnforcement:RepositoryPolicyCaseEnforcement", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RepositoryPolicyCaseEnforcement resources.
type repositoryPolicyCaseEnforcementState struct {
	// A flag indicating if the policy should be blocking. Defaults to `true`.
	Blocking *bool `pulumi:"blocking"`
	// A flag indicating if the policy should be enabled. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// Avoid case-sensitivity conflicts by blocking pushes that change name casing on files, folders, branches, and tags.
	EnforceConsistentCase *bool `pulumi:"enforceConsistentCase"`
	// The ID of the project in which the policy will be created.
	ProjectId *string `pulumi:"projectId"`
	// Control whether the policy is enabled for the repository or the project. If `repositoryIds` not configured, the policy will be set to the project.
	RepositoryIds []string `pulumi:"repositoryIds"`
}

type RepositoryPolicyCaseEnforcementState struct {
	// A flag indicating if the policy should be blocking. Defaults to `true`.
	Blocking pulumi.BoolPtrInput
	// A flag indicating if the policy should be enabled. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// Avoid case-sensitivity conflicts by blocking pushes that change name casing on files, folders, branches, and tags.
	EnforceConsistentCase pulumi.BoolPtrInput
	// The ID of the project in which the policy will be created.
	ProjectId pulumi.StringPtrInput
	// Control whether the policy is enabled for the repository or the project. If `repositoryIds` not configured, the policy will be set to the project.
	RepositoryIds pulumi.StringArrayInput
}

func (RepositoryPolicyCaseEnforcementState) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryPolicyCaseEnforcementState)(nil)).Elem()
}

type repositoryPolicyCaseEnforcementArgs struct {
	// A flag indicating if the policy should be blocking. Defaults to `true`.
	Blocking *bool `pulumi:"blocking"`
	// A flag indicating if the policy should be enabled. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// Avoid case-sensitivity conflicts by blocking pushes that change name casing on files, folders, branches, and tags.
	EnforceConsistentCase bool `pulumi:"enforceConsistentCase"`
	// The ID of the project in which the policy will be created.
	ProjectId string `pulumi:"projectId"`
	// Control whether the policy is enabled for the repository or the project. If `repositoryIds` not configured, the policy will be set to the project.
	RepositoryIds []string `pulumi:"repositoryIds"`
}

// The set of arguments for constructing a RepositoryPolicyCaseEnforcement resource.
type RepositoryPolicyCaseEnforcementArgs struct {
	// A flag indicating if the policy should be blocking. Defaults to `true`.
	Blocking pulumi.BoolPtrInput
	// A flag indicating if the policy should be enabled. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// Avoid case-sensitivity conflicts by blocking pushes that change name casing on files, folders, branches, and tags.
	EnforceConsistentCase pulumi.BoolInput
	// The ID of the project in which the policy will be created.
	ProjectId pulumi.StringInput
	// Control whether the policy is enabled for the repository or the project. If `repositoryIds` not configured, the policy will be set to the project.
	RepositoryIds pulumi.StringArrayInput
}

func (RepositoryPolicyCaseEnforcementArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryPolicyCaseEnforcementArgs)(nil)).Elem()
}

type RepositoryPolicyCaseEnforcementInput interface {
	pulumi.Input

	ToRepositoryPolicyCaseEnforcementOutput() RepositoryPolicyCaseEnforcementOutput
	ToRepositoryPolicyCaseEnforcementOutputWithContext(ctx context.Context) RepositoryPolicyCaseEnforcementOutput
}

func (*RepositoryPolicyCaseEnforcement) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryPolicyCaseEnforcement)(nil)).Elem()
}

func (i *RepositoryPolicyCaseEnforcement) ToRepositoryPolicyCaseEnforcementOutput() RepositoryPolicyCaseEnforcementOutput {
	return i.ToRepositoryPolicyCaseEnforcementOutputWithContext(context.Background())
}

func (i *RepositoryPolicyCaseEnforcement) ToRepositoryPolicyCaseEnforcementOutputWithContext(ctx context.Context) RepositoryPolicyCaseEnforcementOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryPolicyCaseEnforcementOutput)
}

func (i *RepositoryPolicyCaseEnforcement) ToOutput(ctx context.Context) pulumix.Output[*RepositoryPolicyCaseEnforcement] {
	return pulumix.Output[*RepositoryPolicyCaseEnforcement]{
		OutputState: i.ToRepositoryPolicyCaseEnforcementOutputWithContext(ctx).OutputState,
	}
}

// RepositoryPolicyCaseEnforcementArrayInput is an input type that accepts RepositoryPolicyCaseEnforcementArray and RepositoryPolicyCaseEnforcementArrayOutput values.
// You can construct a concrete instance of `RepositoryPolicyCaseEnforcementArrayInput` via:
//
//	RepositoryPolicyCaseEnforcementArray{ RepositoryPolicyCaseEnforcementArgs{...} }
type RepositoryPolicyCaseEnforcementArrayInput interface {
	pulumi.Input

	ToRepositoryPolicyCaseEnforcementArrayOutput() RepositoryPolicyCaseEnforcementArrayOutput
	ToRepositoryPolicyCaseEnforcementArrayOutputWithContext(context.Context) RepositoryPolicyCaseEnforcementArrayOutput
}

type RepositoryPolicyCaseEnforcementArray []RepositoryPolicyCaseEnforcementInput

func (RepositoryPolicyCaseEnforcementArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RepositoryPolicyCaseEnforcement)(nil)).Elem()
}

func (i RepositoryPolicyCaseEnforcementArray) ToRepositoryPolicyCaseEnforcementArrayOutput() RepositoryPolicyCaseEnforcementArrayOutput {
	return i.ToRepositoryPolicyCaseEnforcementArrayOutputWithContext(context.Background())
}

func (i RepositoryPolicyCaseEnforcementArray) ToRepositoryPolicyCaseEnforcementArrayOutputWithContext(ctx context.Context) RepositoryPolicyCaseEnforcementArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryPolicyCaseEnforcementArrayOutput)
}

func (i RepositoryPolicyCaseEnforcementArray) ToOutput(ctx context.Context) pulumix.Output[[]*RepositoryPolicyCaseEnforcement] {
	return pulumix.Output[[]*RepositoryPolicyCaseEnforcement]{
		OutputState: i.ToRepositoryPolicyCaseEnforcementArrayOutputWithContext(ctx).OutputState,
	}
}

// RepositoryPolicyCaseEnforcementMapInput is an input type that accepts RepositoryPolicyCaseEnforcementMap and RepositoryPolicyCaseEnforcementMapOutput values.
// You can construct a concrete instance of `RepositoryPolicyCaseEnforcementMapInput` via:
//
//	RepositoryPolicyCaseEnforcementMap{ "key": RepositoryPolicyCaseEnforcementArgs{...} }
type RepositoryPolicyCaseEnforcementMapInput interface {
	pulumi.Input

	ToRepositoryPolicyCaseEnforcementMapOutput() RepositoryPolicyCaseEnforcementMapOutput
	ToRepositoryPolicyCaseEnforcementMapOutputWithContext(context.Context) RepositoryPolicyCaseEnforcementMapOutput
}

type RepositoryPolicyCaseEnforcementMap map[string]RepositoryPolicyCaseEnforcementInput

func (RepositoryPolicyCaseEnforcementMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RepositoryPolicyCaseEnforcement)(nil)).Elem()
}

func (i RepositoryPolicyCaseEnforcementMap) ToRepositoryPolicyCaseEnforcementMapOutput() RepositoryPolicyCaseEnforcementMapOutput {
	return i.ToRepositoryPolicyCaseEnforcementMapOutputWithContext(context.Background())
}

func (i RepositoryPolicyCaseEnforcementMap) ToRepositoryPolicyCaseEnforcementMapOutputWithContext(ctx context.Context) RepositoryPolicyCaseEnforcementMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryPolicyCaseEnforcementMapOutput)
}

func (i RepositoryPolicyCaseEnforcementMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*RepositoryPolicyCaseEnforcement] {
	return pulumix.Output[map[string]*RepositoryPolicyCaseEnforcement]{
		OutputState: i.ToRepositoryPolicyCaseEnforcementMapOutputWithContext(ctx).OutputState,
	}
}

type RepositoryPolicyCaseEnforcementOutput struct{ *pulumi.OutputState }

func (RepositoryPolicyCaseEnforcementOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryPolicyCaseEnforcement)(nil)).Elem()
}

func (o RepositoryPolicyCaseEnforcementOutput) ToRepositoryPolicyCaseEnforcementOutput() RepositoryPolicyCaseEnforcementOutput {
	return o
}

func (o RepositoryPolicyCaseEnforcementOutput) ToRepositoryPolicyCaseEnforcementOutputWithContext(ctx context.Context) RepositoryPolicyCaseEnforcementOutput {
	return o
}

func (o RepositoryPolicyCaseEnforcementOutput) ToOutput(ctx context.Context) pulumix.Output[*RepositoryPolicyCaseEnforcement] {
	return pulumix.Output[*RepositoryPolicyCaseEnforcement]{
		OutputState: o.OutputState,
	}
}

// A flag indicating if the policy should be blocking. Defaults to `true`.
func (o RepositoryPolicyCaseEnforcementOutput) Blocking() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RepositoryPolicyCaseEnforcement) pulumi.BoolPtrOutput { return v.Blocking }).(pulumi.BoolPtrOutput)
}

// A flag indicating if the policy should be enabled. Defaults to `true`.
func (o RepositoryPolicyCaseEnforcementOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RepositoryPolicyCaseEnforcement) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// Avoid case-sensitivity conflicts by blocking pushes that change name casing on files, folders, branches, and tags.
func (o RepositoryPolicyCaseEnforcementOutput) EnforceConsistentCase() pulumi.BoolOutput {
	return o.ApplyT(func(v *RepositoryPolicyCaseEnforcement) pulumi.BoolOutput { return v.EnforceConsistentCase }).(pulumi.BoolOutput)
}

// The ID of the project in which the policy will be created.
func (o RepositoryPolicyCaseEnforcementOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryPolicyCaseEnforcement) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// Control whether the policy is enabled for the repository or the project. If `repositoryIds` not configured, the policy will be set to the project.
func (o RepositoryPolicyCaseEnforcementOutput) RepositoryIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RepositoryPolicyCaseEnforcement) pulumi.StringArrayOutput { return v.RepositoryIds }).(pulumi.StringArrayOutput)
}

type RepositoryPolicyCaseEnforcementArrayOutput struct{ *pulumi.OutputState }

func (RepositoryPolicyCaseEnforcementArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RepositoryPolicyCaseEnforcement)(nil)).Elem()
}

func (o RepositoryPolicyCaseEnforcementArrayOutput) ToRepositoryPolicyCaseEnforcementArrayOutput() RepositoryPolicyCaseEnforcementArrayOutput {
	return o
}

func (o RepositoryPolicyCaseEnforcementArrayOutput) ToRepositoryPolicyCaseEnforcementArrayOutputWithContext(ctx context.Context) RepositoryPolicyCaseEnforcementArrayOutput {
	return o
}

func (o RepositoryPolicyCaseEnforcementArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*RepositoryPolicyCaseEnforcement] {
	return pulumix.Output[[]*RepositoryPolicyCaseEnforcement]{
		OutputState: o.OutputState,
	}
}

func (o RepositoryPolicyCaseEnforcementArrayOutput) Index(i pulumi.IntInput) RepositoryPolicyCaseEnforcementOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RepositoryPolicyCaseEnforcement {
		return vs[0].([]*RepositoryPolicyCaseEnforcement)[vs[1].(int)]
	}).(RepositoryPolicyCaseEnforcementOutput)
}

type RepositoryPolicyCaseEnforcementMapOutput struct{ *pulumi.OutputState }

func (RepositoryPolicyCaseEnforcementMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RepositoryPolicyCaseEnforcement)(nil)).Elem()
}

func (o RepositoryPolicyCaseEnforcementMapOutput) ToRepositoryPolicyCaseEnforcementMapOutput() RepositoryPolicyCaseEnforcementMapOutput {
	return o
}

func (o RepositoryPolicyCaseEnforcementMapOutput) ToRepositoryPolicyCaseEnforcementMapOutputWithContext(ctx context.Context) RepositoryPolicyCaseEnforcementMapOutput {
	return o
}

func (o RepositoryPolicyCaseEnforcementMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*RepositoryPolicyCaseEnforcement] {
	return pulumix.Output[map[string]*RepositoryPolicyCaseEnforcement]{
		OutputState: o.OutputState,
	}
}

func (o RepositoryPolicyCaseEnforcementMapOutput) MapIndex(k pulumi.StringInput) RepositoryPolicyCaseEnforcementOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RepositoryPolicyCaseEnforcement {
		return vs[0].(map[string]*RepositoryPolicyCaseEnforcement)[vs[1].(string)]
	}).(RepositoryPolicyCaseEnforcementOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryPolicyCaseEnforcementInput)(nil)).Elem(), &RepositoryPolicyCaseEnforcement{})
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryPolicyCaseEnforcementArrayInput)(nil)).Elem(), RepositoryPolicyCaseEnforcementArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryPolicyCaseEnforcementMapInput)(nil)).Elem(), RepositoryPolicyCaseEnforcementMap{})
	pulumi.RegisterOutputType(RepositoryPolicyCaseEnforcementOutput{})
	pulumi.RegisterOutputType(RepositoryPolicyCaseEnforcementArrayOutput{})
	pulumi.RegisterOutputType(RepositoryPolicyCaseEnforcementMapOutput{})
}
