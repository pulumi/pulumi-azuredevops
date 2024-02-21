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

// Configure a comment resolution policy for your branch within Azure DevOps project.
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
//			_, err = azuredevops.NewBranchPolicyCommentResolution(ctx, "exampleBranchPolicyCommentResolution", &azuredevops.BranchPolicyCommentResolutionArgs{
//				ProjectId: exampleProject.ID(),
//				Enabled:   pulumi.Bool(true),
//				Blocking:  pulumi.Bool(true),
//				Settings: &azuredevops.BranchPolicyCommentResolutionSettingsArgs{
//					Scopes: azuredevops.BranchPolicyCommentResolutionSettingsScopeArray{
//						&azuredevops.BranchPolicyCommentResolutionSettingsScopeArgs{
//							RepositoryId:  exampleGit.ID(),
//							RepositoryRef: exampleGit.DefaultBranch,
//							MatchType:     pulumi.String("Exact"),
//						},
//						&azuredevops.BranchPolicyCommentResolutionSettingsScopeArgs{
//							RepositoryId:  exampleGit.ID(),
//							RepositoryRef: pulumi.String("refs/heads/releases"),
//							MatchType:     pulumi.String("Prefix"),
//						},
//						&azuredevops.BranchPolicyCommentResolutionSettingsScopeArgs{
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
// Azure DevOps Branch Policies can be imported using the project ID and policy configuration ID:
//
// ```sh
// $ pulumi import azuredevops:index/branchPolicyCommentResolution:BranchPolicyCommentResolution example 00000000-0000-0000-0000-000000000000/0
// ```
type BranchPolicyCommentResolution struct {
	pulumi.CustomResourceState

	// A flag indicating if the policy should be blocking. Defaults to `true`.
	Blocking pulumi.BoolPtrOutput `pulumi:"blocking"`
	// A flag indicating if the policy should be enabled. Defaults to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// The ID of the project in which the policy will be created.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Configuration for the policy. This block must be defined exactly once.
	Settings BranchPolicyCommentResolutionSettingsOutput `pulumi:"settings"`
}

// NewBranchPolicyCommentResolution registers a new resource with the given unique name, arguments, and options.
func NewBranchPolicyCommentResolution(ctx *pulumi.Context,
	name string, args *BranchPolicyCommentResolutionArgs, opts ...pulumi.ResourceOption) (*BranchPolicyCommentResolution, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.Settings == nil {
		return nil, errors.New("invalid value for required argument 'Settings'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource BranchPolicyCommentResolution
	err := ctx.RegisterResource("azuredevops:index/branchPolicyCommentResolution:BranchPolicyCommentResolution", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBranchPolicyCommentResolution gets an existing BranchPolicyCommentResolution resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBranchPolicyCommentResolution(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BranchPolicyCommentResolutionState, opts ...pulumi.ResourceOption) (*BranchPolicyCommentResolution, error) {
	var resource BranchPolicyCommentResolution
	err := ctx.ReadResource("azuredevops:index/branchPolicyCommentResolution:BranchPolicyCommentResolution", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BranchPolicyCommentResolution resources.
type branchPolicyCommentResolutionState struct {
	// A flag indicating if the policy should be blocking. Defaults to `true`.
	Blocking *bool `pulumi:"blocking"`
	// A flag indicating if the policy should be enabled. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// The ID of the project in which the policy will be created.
	ProjectId *string `pulumi:"projectId"`
	// Configuration for the policy. This block must be defined exactly once.
	Settings *BranchPolicyCommentResolutionSettings `pulumi:"settings"`
}

type BranchPolicyCommentResolutionState struct {
	// A flag indicating if the policy should be blocking. Defaults to `true`.
	Blocking pulumi.BoolPtrInput
	// A flag indicating if the policy should be enabled. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// The ID of the project in which the policy will be created.
	ProjectId pulumi.StringPtrInput
	// Configuration for the policy. This block must be defined exactly once.
	Settings BranchPolicyCommentResolutionSettingsPtrInput
}

func (BranchPolicyCommentResolutionState) ElementType() reflect.Type {
	return reflect.TypeOf((*branchPolicyCommentResolutionState)(nil)).Elem()
}

type branchPolicyCommentResolutionArgs struct {
	// A flag indicating if the policy should be blocking. Defaults to `true`.
	Blocking *bool `pulumi:"blocking"`
	// A flag indicating if the policy should be enabled. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// The ID of the project in which the policy will be created.
	ProjectId string `pulumi:"projectId"`
	// Configuration for the policy. This block must be defined exactly once.
	Settings BranchPolicyCommentResolutionSettings `pulumi:"settings"`
}

// The set of arguments for constructing a BranchPolicyCommentResolution resource.
type BranchPolicyCommentResolutionArgs struct {
	// A flag indicating if the policy should be blocking. Defaults to `true`.
	Blocking pulumi.BoolPtrInput
	// A flag indicating if the policy should be enabled. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// The ID of the project in which the policy will be created.
	ProjectId pulumi.StringInput
	// Configuration for the policy. This block must be defined exactly once.
	Settings BranchPolicyCommentResolutionSettingsInput
}

func (BranchPolicyCommentResolutionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*branchPolicyCommentResolutionArgs)(nil)).Elem()
}

type BranchPolicyCommentResolutionInput interface {
	pulumi.Input

	ToBranchPolicyCommentResolutionOutput() BranchPolicyCommentResolutionOutput
	ToBranchPolicyCommentResolutionOutputWithContext(ctx context.Context) BranchPolicyCommentResolutionOutput
}

func (*BranchPolicyCommentResolution) ElementType() reflect.Type {
	return reflect.TypeOf((**BranchPolicyCommentResolution)(nil)).Elem()
}

func (i *BranchPolicyCommentResolution) ToBranchPolicyCommentResolutionOutput() BranchPolicyCommentResolutionOutput {
	return i.ToBranchPolicyCommentResolutionOutputWithContext(context.Background())
}

func (i *BranchPolicyCommentResolution) ToBranchPolicyCommentResolutionOutputWithContext(ctx context.Context) BranchPolicyCommentResolutionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BranchPolicyCommentResolutionOutput)
}

// BranchPolicyCommentResolutionArrayInput is an input type that accepts BranchPolicyCommentResolutionArray and BranchPolicyCommentResolutionArrayOutput values.
// You can construct a concrete instance of `BranchPolicyCommentResolutionArrayInput` via:
//
//	BranchPolicyCommentResolutionArray{ BranchPolicyCommentResolutionArgs{...} }
type BranchPolicyCommentResolutionArrayInput interface {
	pulumi.Input

	ToBranchPolicyCommentResolutionArrayOutput() BranchPolicyCommentResolutionArrayOutput
	ToBranchPolicyCommentResolutionArrayOutputWithContext(context.Context) BranchPolicyCommentResolutionArrayOutput
}

type BranchPolicyCommentResolutionArray []BranchPolicyCommentResolutionInput

func (BranchPolicyCommentResolutionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BranchPolicyCommentResolution)(nil)).Elem()
}

func (i BranchPolicyCommentResolutionArray) ToBranchPolicyCommentResolutionArrayOutput() BranchPolicyCommentResolutionArrayOutput {
	return i.ToBranchPolicyCommentResolutionArrayOutputWithContext(context.Background())
}

func (i BranchPolicyCommentResolutionArray) ToBranchPolicyCommentResolutionArrayOutputWithContext(ctx context.Context) BranchPolicyCommentResolutionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BranchPolicyCommentResolutionArrayOutput)
}

// BranchPolicyCommentResolutionMapInput is an input type that accepts BranchPolicyCommentResolutionMap and BranchPolicyCommentResolutionMapOutput values.
// You can construct a concrete instance of `BranchPolicyCommentResolutionMapInput` via:
//
//	BranchPolicyCommentResolutionMap{ "key": BranchPolicyCommentResolutionArgs{...} }
type BranchPolicyCommentResolutionMapInput interface {
	pulumi.Input

	ToBranchPolicyCommentResolutionMapOutput() BranchPolicyCommentResolutionMapOutput
	ToBranchPolicyCommentResolutionMapOutputWithContext(context.Context) BranchPolicyCommentResolutionMapOutput
}

type BranchPolicyCommentResolutionMap map[string]BranchPolicyCommentResolutionInput

func (BranchPolicyCommentResolutionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BranchPolicyCommentResolution)(nil)).Elem()
}

func (i BranchPolicyCommentResolutionMap) ToBranchPolicyCommentResolutionMapOutput() BranchPolicyCommentResolutionMapOutput {
	return i.ToBranchPolicyCommentResolutionMapOutputWithContext(context.Background())
}

func (i BranchPolicyCommentResolutionMap) ToBranchPolicyCommentResolutionMapOutputWithContext(ctx context.Context) BranchPolicyCommentResolutionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BranchPolicyCommentResolutionMapOutput)
}

type BranchPolicyCommentResolutionOutput struct{ *pulumi.OutputState }

func (BranchPolicyCommentResolutionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BranchPolicyCommentResolution)(nil)).Elem()
}

func (o BranchPolicyCommentResolutionOutput) ToBranchPolicyCommentResolutionOutput() BranchPolicyCommentResolutionOutput {
	return o
}

func (o BranchPolicyCommentResolutionOutput) ToBranchPolicyCommentResolutionOutputWithContext(ctx context.Context) BranchPolicyCommentResolutionOutput {
	return o
}

// A flag indicating if the policy should be blocking. Defaults to `true`.
func (o BranchPolicyCommentResolutionOutput) Blocking() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BranchPolicyCommentResolution) pulumi.BoolPtrOutput { return v.Blocking }).(pulumi.BoolPtrOutput)
}

// A flag indicating if the policy should be enabled. Defaults to `true`.
func (o BranchPolicyCommentResolutionOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BranchPolicyCommentResolution) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// The ID of the project in which the policy will be created.
func (o BranchPolicyCommentResolutionOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *BranchPolicyCommentResolution) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// Configuration for the policy. This block must be defined exactly once.
func (o BranchPolicyCommentResolutionOutput) Settings() BranchPolicyCommentResolutionSettingsOutput {
	return o.ApplyT(func(v *BranchPolicyCommentResolution) BranchPolicyCommentResolutionSettingsOutput { return v.Settings }).(BranchPolicyCommentResolutionSettingsOutput)
}

type BranchPolicyCommentResolutionArrayOutput struct{ *pulumi.OutputState }

func (BranchPolicyCommentResolutionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BranchPolicyCommentResolution)(nil)).Elem()
}

func (o BranchPolicyCommentResolutionArrayOutput) ToBranchPolicyCommentResolutionArrayOutput() BranchPolicyCommentResolutionArrayOutput {
	return o
}

func (o BranchPolicyCommentResolutionArrayOutput) ToBranchPolicyCommentResolutionArrayOutputWithContext(ctx context.Context) BranchPolicyCommentResolutionArrayOutput {
	return o
}

func (o BranchPolicyCommentResolutionArrayOutput) Index(i pulumi.IntInput) BranchPolicyCommentResolutionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BranchPolicyCommentResolution {
		return vs[0].([]*BranchPolicyCommentResolution)[vs[1].(int)]
	}).(BranchPolicyCommentResolutionOutput)
}

type BranchPolicyCommentResolutionMapOutput struct{ *pulumi.OutputState }

func (BranchPolicyCommentResolutionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BranchPolicyCommentResolution)(nil)).Elem()
}

func (o BranchPolicyCommentResolutionMapOutput) ToBranchPolicyCommentResolutionMapOutput() BranchPolicyCommentResolutionMapOutput {
	return o
}

func (o BranchPolicyCommentResolutionMapOutput) ToBranchPolicyCommentResolutionMapOutputWithContext(ctx context.Context) BranchPolicyCommentResolutionMapOutput {
	return o
}

func (o BranchPolicyCommentResolutionMapOutput) MapIndex(k pulumi.StringInput) BranchPolicyCommentResolutionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BranchPolicyCommentResolution {
		return vs[0].(map[string]*BranchPolicyCommentResolution)[vs[1].(string)]
	}).(BranchPolicyCommentResolutionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BranchPolicyCommentResolutionInput)(nil)).Elem(), &BranchPolicyCommentResolution{})
	pulumi.RegisterInputType(reflect.TypeOf((*BranchPolicyCommentResolutionArrayInput)(nil)).Elem(), BranchPolicyCommentResolutionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BranchPolicyCommentResolutionMapInput)(nil)).Elem(), BranchPolicyCommentResolutionMap{})
	pulumi.RegisterOutputType(BranchPolicyCommentResolutionOutput{})
	pulumi.RegisterOutputType(BranchPolicyCommentResolutionArrayOutput{})
	pulumi.RegisterOutputType(BranchPolicyCommentResolutionMapOutput{})
}
