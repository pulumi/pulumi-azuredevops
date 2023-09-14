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

// Manages a build validation branch policy within Azure DevOps.
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
//			exampleProject, err := azuredevops.NewProject(ctx, "exampleProject", nil)
//			if err != nil {
//				return err
//			}
//			exampleGit, err := azuredevops.NewGit(ctx, "exampleGit", &azuredevops.GitArgs{
//				ProjectId: exampleProject.ID(),
//				Initialization: &azuredevops.GitInitializationArgs{
//					InitType: pulumi.String("Clean"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			exampleBuildDefinition, err := azuredevops.NewBuildDefinition(ctx, "exampleBuildDefinition", &azuredevops.BuildDefinitionArgs{
//				ProjectId: exampleProject.ID(),
//				Repository: &azuredevops.BuildDefinitionRepositoryArgs{
//					RepoType: pulumi.String("TfsGit"),
//					RepoId:   exampleGit.ID(),
//					YmlPath:  pulumi.String("azure-pipelines.yml"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuredevops.NewBranchPolicyBuildValidation(ctx, "exampleBranchPolicyBuildValidation", &azuredevops.BranchPolicyBuildValidationArgs{
//				ProjectId: exampleProject.ID(),
//				Enabled:   pulumi.Bool(true),
//				Blocking:  pulumi.Bool(true),
//				Settings: &azuredevops.BranchPolicyBuildValidationSettingsArgs{
//					DisplayName:       pulumi.String("Example build validation policy"),
//					BuildDefinitionId: exampleBuildDefinition.ID(),
//					ValidDuration:     pulumi.Int(720),
//					FilenamePatterns: pulumi.StringArray{
//						pulumi.String("/WebApp/*"),
//						pulumi.String("!/WebApp/Tests/*"),
//						pulumi.String("*.cs"),
//					},
//					Scopes: azuredevops.BranchPolicyBuildValidationSettingsScopeArray{
//						&azuredevops.BranchPolicyBuildValidationSettingsScopeArgs{
//							RepositoryId:  exampleGit.ID(),
//							RepositoryRef: exampleGit.DefaultBranch,
//							MatchType:     pulumi.String("Exact"),
//						},
//						&azuredevops.BranchPolicyBuildValidationSettingsScopeArgs{
//							RepositoryId:  exampleGit.ID(),
//							RepositoryRef: pulumi.String("refs/heads/releases"),
//							MatchType:     pulumi.String("Prefix"),
//						},
//						&azuredevops.BranchPolicyBuildValidationSettingsScopeArgs{
//							MatchType: pulumi.String("DefaultBranch"),
//						},
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
// ## Relevant Links
//
// - [Azure DevOps Service REST API 7.0 - Policy Configurations](https://docs.microsoft.com/en-us/rest/api/azure/devops/policy/configurations/create?view=azure-devops-rest-7.0)
//
// ## Import
//
// # Azure DevOps Branch Policies can be imported using the project ID and policy configuration ID
//
// ```sh
//
//	$ pulumi import azuredevops:index/branchPolicyBuildValidation:BranchPolicyBuildValidation example 00000000-0000-0000-0000-000000000000/0
//
// ```
type BranchPolicyBuildValidation struct {
	pulumi.CustomResourceState

	// A flag indicating if the policy should be blocking. Defaults to `true`.
	Blocking pulumi.BoolPtrOutput `pulumi:"blocking"`
	// A flag indicating if the policy should be enabled. Defaults to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// The ID of the project in which the policy will be created.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Configuration for the policy. This block must be defined exactly once.
	Settings BranchPolicyBuildValidationSettingsOutput `pulumi:"settings"`
}

// NewBranchPolicyBuildValidation registers a new resource with the given unique name, arguments, and options.
func NewBranchPolicyBuildValidation(ctx *pulumi.Context,
	name string, args *BranchPolicyBuildValidationArgs, opts ...pulumi.ResourceOption) (*BranchPolicyBuildValidation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.Settings == nil {
		return nil, errors.New("invalid value for required argument 'Settings'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azuredevops:Policy/branchPolicyBuildValidation:BranchPolicyBuildValidation"),
		},
	})
	opts = append(opts, aliases)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource BranchPolicyBuildValidation
	err := ctx.RegisterResource("azuredevops:index/branchPolicyBuildValidation:BranchPolicyBuildValidation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBranchPolicyBuildValidation gets an existing BranchPolicyBuildValidation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBranchPolicyBuildValidation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BranchPolicyBuildValidationState, opts ...pulumi.ResourceOption) (*BranchPolicyBuildValidation, error) {
	var resource BranchPolicyBuildValidation
	err := ctx.ReadResource("azuredevops:index/branchPolicyBuildValidation:BranchPolicyBuildValidation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BranchPolicyBuildValidation resources.
type branchPolicyBuildValidationState struct {
	// A flag indicating if the policy should be blocking. Defaults to `true`.
	Blocking *bool `pulumi:"blocking"`
	// A flag indicating if the policy should be enabled. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// The ID of the project in which the policy will be created.
	ProjectId *string `pulumi:"projectId"`
	// Configuration for the policy. This block must be defined exactly once.
	Settings *BranchPolicyBuildValidationSettings `pulumi:"settings"`
}

type BranchPolicyBuildValidationState struct {
	// A flag indicating if the policy should be blocking. Defaults to `true`.
	Blocking pulumi.BoolPtrInput
	// A flag indicating if the policy should be enabled. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// The ID of the project in which the policy will be created.
	ProjectId pulumi.StringPtrInput
	// Configuration for the policy. This block must be defined exactly once.
	Settings BranchPolicyBuildValidationSettingsPtrInput
}

func (BranchPolicyBuildValidationState) ElementType() reflect.Type {
	return reflect.TypeOf((*branchPolicyBuildValidationState)(nil)).Elem()
}

type branchPolicyBuildValidationArgs struct {
	// A flag indicating if the policy should be blocking. Defaults to `true`.
	Blocking *bool `pulumi:"blocking"`
	// A flag indicating if the policy should be enabled. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// The ID of the project in which the policy will be created.
	ProjectId string `pulumi:"projectId"`
	// Configuration for the policy. This block must be defined exactly once.
	Settings BranchPolicyBuildValidationSettings `pulumi:"settings"`
}

// The set of arguments for constructing a BranchPolicyBuildValidation resource.
type BranchPolicyBuildValidationArgs struct {
	// A flag indicating if the policy should be blocking. Defaults to `true`.
	Blocking pulumi.BoolPtrInput
	// A flag indicating if the policy should be enabled. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// The ID of the project in which the policy will be created.
	ProjectId pulumi.StringInput
	// Configuration for the policy. This block must be defined exactly once.
	Settings BranchPolicyBuildValidationSettingsInput
}

func (BranchPolicyBuildValidationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*branchPolicyBuildValidationArgs)(nil)).Elem()
}

type BranchPolicyBuildValidationInput interface {
	pulumi.Input

	ToBranchPolicyBuildValidationOutput() BranchPolicyBuildValidationOutput
	ToBranchPolicyBuildValidationOutputWithContext(ctx context.Context) BranchPolicyBuildValidationOutput
}

func (*BranchPolicyBuildValidation) ElementType() reflect.Type {
	return reflect.TypeOf((**BranchPolicyBuildValidation)(nil)).Elem()
}

func (i *BranchPolicyBuildValidation) ToBranchPolicyBuildValidationOutput() BranchPolicyBuildValidationOutput {
	return i.ToBranchPolicyBuildValidationOutputWithContext(context.Background())
}

func (i *BranchPolicyBuildValidation) ToBranchPolicyBuildValidationOutputWithContext(ctx context.Context) BranchPolicyBuildValidationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BranchPolicyBuildValidationOutput)
}

func (i *BranchPolicyBuildValidation) ToOutput(ctx context.Context) pulumix.Output[*BranchPolicyBuildValidation] {
	return pulumix.Output[*BranchPolicyBuildValidation]{
		OutputState: i.ToBranchPolicyBuildValidationOutputWithContext(ctx).OutputState,
	}
}

// BranchPolicyBuildValidationArrayInput is an input type that accepts BranchPolicyBuildValidationArray and BranchPolicyBuildValidationArrayOutput values.
// You can construct a concrete instance of `BranchPolicyBuildValidationArrayInput` via:
//
//	BranchPolicyBuildValidationArray{ BranchPolicyBuildValidationArgs{...} }
type BranchPolicyBuildValidationArrayInput interface {
	pulumi.Input

	ToBranchPolicyBuildValidationArrayOutput() BranchPolicyBuildValidationArrayOutput
	ToBranchPolicyBuildValidationArrayOutputWithContext(context.Context) BranchPolicyBuildValidationArrayOutput
}

type BranchPolicyBuildValidationArray []BranchPolicyBuildValidationInput

func (BranchPolicyBuildValidationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BranchPolicyBuildValidation)(nil)).Elem()
}

func (i BranchPolicyBuildValidationArray) ToBranchPolicyBuildValidationArrayOutput() BranchPolicyBuildValidationArrayOutput {
	return i.ToBranchPolicyBuildValidationArrayOutputWithContext(context.Background())
}

func (i BranchPolicyBuildValidationArray) ToBranchPolicyBuildValidationArrayOutputWithContext(ctx context.Context) BranchPolicyBuildValidationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BranchPolicyBuildValidationArrayOutput)
}

func (i BranchPolicyBuildValidationArray) ToOutput(ctx context.Context) pulumix.Output[[]*BranchPolicyBuildValidation] {
	return pulumix.Output[[]*BranchPolicyBuildValidation]{
		OutputState: i.ToBranchPolicyBuildValidationArrayOutputWithContext(ctx).OutputState,
	}
}

// BranchPolicyBuildValidationMapInput is an input type that accepts BranchPolicyBuildValidationMap and BranchPolicyBuildValidationMapOutput values.
// You can construct a concrete instance of `BranchPolicyBuildValidationMapInput` via:
//
//	BranchPolicyBuildValidationMap{ "key": BranchPolicyBuildValidationArgs{...} }
type BranchPolicyBuildValidationMapInput interface {
	pulumi.Input

	ToBranchPolicyBuildValidationMapOutput() BranchPolicyBuildValidationMapOutput
	ToBranchPolicyBuildValidationMapOutputWithContext(context.Context) BranchPolicyBuildValidationMapOutput
}

type BranchPolicyBuildValidationMap map[string]BranchPolicyBuildValidationInput

func (BranchPolicyBuildValidationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BranchPolicyBuildValidation)(nil)).Elem()
}

func (i BranchPolicyBuildValidationMap) ToBranchPolicyBuildValidationMapOutput() BranchPolicyBuildValidationMapOutput {
	return i.ToBranchPolicyBuildValidationMapOutputWithContext(context.Background())
}

func (i BranchPolicyBuildValidationMap) ToBranchPolicyBuildValidationMapOutputWithContext(ctx context.Context) BranchPolicyBuildValidationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BranchPolicyBuildValidationMapOutput)
}

func (i BranchPolicyBuildValidationMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*BranchPolicyBuildValidation] {
	return pulumix.Output[map[string]*BranchPolicyBuildValidation]{
		OutputState: i.ToBranchPolicyBuildValidationMapOutputWithContext(ctx).OutputState,
	}
}

type BranchPolicyBuildValidationOutput struct{ *pulumi.OutputState }

func (BranchPolicyBuildValidationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BranchPolicyBuildValidation)(nil)).Elem()
}

func (o BranchPolicyBuildValidationOutput) ToBranchPolicyBuildValidationOutput() BranchPolicyBuildValidationOutput {
	return o
}

func (o BranchPolicyBuildValidationOutput) ToBranchPolicyBuildValidationOutputWithContext(ctx context.Context) BranchPolicyBuildValidationOutput {
	return o
}

func (o BranchPolicyBuildValidationOutput) ToOutput(ctx context.Context) pulumix.Output[*BranchPolicyBuildValidation] {
	return pulumix.Output[*BranchPolicyBuildValidation]{
		OutputState: o.OutputState,
	}
}

// A flag indicating if the policy should be blocking. Defaults to `true`.
func (o BranchPolicyBuildValidationOutput) Blocking() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BranchPolicyBuildValidation) pulumi.BoolPtrOutput { return v.Blocking }).(pulumi.BoolPtrOutput)
}

// A flag indicating if the policy should be enabled. Defaults to `true`.
func (o BranchPolicyBuildValidationOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BranchPolicyBuildValidation) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// The ID of the project in which the policy will be created.
func (o BranchPolicyBuildValidationOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *BranchPolicyBuildValidation) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// Configuration for the policy. This block must be defined exactly once.
func (o BranchPolicyBuildValidationOutput) Settings() BranchPolicyBuildValidationSettingsOutput {
	return o.ApplyT(func(v *BranchPolicyBuildValidation) BranchPolicyBuildValidationSettingsOutput { return v.Settings }).(BranchPolicyBuildValidationSettingsOutput)
}

type BranchPolicyBuildValidationArrayOutput struct{ *pulumi.OutputState }

func (BranchPolicyBuildValidationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BranchPolicyBuildValidation)(nil)).Elem()
}

func (o BranchPolicyBuildValidationArrayOutput) ToBranchPolicyBuildValidationArrayOutput() BranchPolicyBuildValidationArrayOutput {
	return o
}

func (o BranchPolicyBuildValidationArrayOutput) ToBranchPolicyBuildValidationArrayOutputWithContext(ctx context.Context) BranchPolicyBuildValidationArrayOutput {
	return o
}

func (o BranchPolicyBuildValidationArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*BranchPolicyBuildValidation] {
	return pulumix.Output[[]*BranchPolicyBuildValidation]{
		OutputState: o.OutputState,
	}
}

func (o BranchPolicyBuildValidationArrayOutput) Index(i pulumi.IntInput) BranchPolicyBuildValidationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BranchPolicyBuildValidation {
		return vs[0].([]*BranchPolicyBuildValidation)[vs[1].(int)]
	}).(BranchPolicyBuildValidationOutput)
}

type BranchPolicyBuildValidationMapOutput struct{ *pulumi.OutputState }

func (BranchPolicyBuildValidationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BranchPolicyBuildValidation)(nil)).Elem()
}

func (o BranchPolicyBuildValidationMapOutput) ToBranchPolicyBuildValidationMapOutput() BranchPolicyBuildValidationMapOutput {
	return o
}

func (o BranchPolicyBuildValidationMapOutput) ToBranchPolicyBuildValidationMapOutputWithContext(ctx context.Context) BranchPolicyBuildValidationMapOutput {
	return o
}

func (o BranchPolicyBuildValidationMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*BranchPolicyBuildValidation] {
	return pulumix.Output[map[string]*BranchPolicyBuildValidation]{
		OutputState: o.OutputState,
	}
}

func (o BranchPolicyBuildValidationMapOutput) MapIndex(k pulumi.StringInput) BranchPolicyBuildValidationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BranchPolicyBuildValidation {
		return vs[0].(map[string]*BranchPolicyBuildValidation)[vs[1].(string)]
	}).(BranchPolicyBuildValidationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BranchPolicyBuildValidationInput)(nil)).Elem(), &BranchPolicyBuildValidation{})
	pulumi.RegisterInputType(reflect.TypeOf((*BranchPolicyBuildValidationArrayInput)(nil)).Elem(), BranchPolicyBuildValidationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BranchPolicyBuildValidationMapInput)(nil)).Elem(), BranchPolicyBuildValidationMap{})
	pulumi.RegisterOutputType(BranchPolicyBuildValidationOutput{})
	pulumi.RegisterOutputType(BranchPolicyBuildValidationArrayOutput{})
	pulumi.RegisterOutputType(BranchPolicyBuildValidationMapOutput{})
}
