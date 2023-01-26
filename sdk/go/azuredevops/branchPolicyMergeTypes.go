// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Branch policy for merge types allowed on a specified branch.
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
//				Initialization: &GitInitializationArgs{
//					InitType: pulumi.String("Clean"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuredevops.NewBranchPolicyMergeTypes(ctx, "exampleBranchPolicyMergeTypes", &azuredevops.BranchPolicyMergeTypesArgs{
//				ProjectId: exampleProject.ID(),
//				Enabled:   pulumi.Bool(true),
//				Blocking:  pulumi.Bool(true),
//				Settings: &BranchPolicyMergeTypesSettingsArgs{
//					AllowSquash:               pulumi.Bool(true),
//					AllowRebaseAndFastForward: pulumi.Bool(true),
//					AllowBasicNoFastForward:   pulumi.Bool(true),
//					AllowRebaseWithMerge:      pulumi.Bool(true),
//					Scopes: BranchPolicyMergeTypesSettingsScopeArray{
//						&BranchPolicyMergeTypesSettingsScopeArgs{
//							RepositoryId:  exampleGit.ID(),
//							RepositoryRef: exampleGit.DefaultBranch,
//							MatchType:     pulumi.String("Exact"),
//						},
//						&BranchPolicyMergeTypesSettingsScopeArgs{
//							RepositoryId:  nil,
//							RepositoryRef: pulumi.String("refs/heads/releases"),
//							MatchType:     pulumi.String("Prefix"),
//						},
//						&BranchPolicyMergeTypesSettingsScopeArgs{
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
// - [Azure DevOps Service REST API 6.0 - Policy Configurations](https://docs.microsoft.com/en-us/rest/api/azure/devops/policy/configurations/create?view=azure-devops-rest-6.0)
//
// ## Import
//
// # Azure DevOps Branch Policies can be imported using the project ID and policy configuration ID
//
// ```sh
//
//	$ pulumi import azuredevops:index/branchPolicyMergeTypes:BranchPolicyMergeTypes example 00000000-0000-0000-0000-000000000000/0
//
// ```
type BranchPolicyMergeTypes struct {
	pulumi.CustomResourceState

	// A flag indicating if the policy should be blocking. Defaults to `true`.
	Blocking pulumi.BoolPtrOutput `pulumi:"blocking"`
	// A flag indicating if the policy should be enabled. Defaults to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// The ID of the project in which the policy will be created.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Configuration for the policy. This block must be defined exactly once.
	Settings BranchPolicyMergeTypesSettingsOutput `pulumi:"settings"`
}

// NewBranchPolicyMergeTypes registers a new resource with the given unique name, arguments, and options.
func NewBranchPolicyMergeTypes(ctx *pulumi.Context,
	name string, args *BranchPolicyMergeTypesArgs, opts ...pulumi.ResourceOption) (*BranchPolicyMergeTypes, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.Settings == nil {
		return nil, errors.New("invalid value for required argument 'Settings'")
	}
	var resource BranchPolicyMergeTypes
	err := ctx.RegisterResource("azuredevops:index/branchPolicyMergeTypes:BranchPolicyMergeTypes", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBranchPolicyMergeTypes gets an existing BranchPolicyMergeTypes resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBranchPolicyMergeTypes(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BranchPolicyMergeTypesState, opts ...pulumi.ResourceOption) (*BranchPolicyMergeTypes, error) {
	var resource BranchPolicyMergeTypes
	err := ctx.ReadResource("azuredevops:index/branchPolicyMergeTypes:BranchPolicyMergeTypes", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BranchPolicyMergeTypes resources.
type branchPolicyMergeTypesState struct {
	// A flag indicating if the policy should be blocking. Defaults to `true`.
	Blocking *bool `pulumi:"blocking"`
	// A flag indicating if the policy should be enabled. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// The ID of the project in which the policy will be created.
	ProjectId *string `pulumi:"projectId"`
	// Configuration for the policy. This block must be defined exactly once.
	Settings *BranchPolicyMergeTypesSettings `pulumi:"settings"`
}

type BranchPolicyMergeTypesState struct {
	// A flag indicating if the policy should be blocking. Defaults to `true`.
	Blocking pulumi.BoolPtrInput
	// A flag indicating if the policy should be enabled. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// The ID of the project in which the policy will be created.
	ProjectId pulumi.StringPtrInput
	// Configuration for the policy. This block must be defined exactly once.
	Settings BranchPolicyMergeTypesSettingsPtrInput
}

func (BranchPolicyMergeTypesState) ElementType() reflect.Type {
	return reflect.TypeOf((*branchPolicyMergeTypesState)(nil)).Elem()
}

type branchPolicyMergeTypesArgs struct {
	// A flag indicating if the policy should be blocking. Defaults to `true`.
	Blocking *bool `pulumi:"blocking"`
	// A flag indicating if the policy should be enabled. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// The ID of the project in which the policy will be created.
	ProjectId string `pulumi:"projectId"`
	// Configuration for the policy. This block must be defined exactly once.
	Settings BranchPolicyMergeTypesSettings `pulumi:"settings"`
}

// The set of arguments for constructing a BranchPolicyMergeTypes resource.
type BranchPolicyMergeTypesArgs struct {
	// A flag indicating if the policy should be blocking. Defaults to `true`.
	Blocking pulumi.BoolPtrInput
	// A flag indicating if the policy should be enabled. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// The ID of the project in which the policy will be created.
	ProjectId pulumi.StringInput
	// Configuration for the policy. This block must be defined exactly once.
	Settings BranchPolicyMergeTypesSettingsInput
}

func (BranchPolicyMergeTypesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*branchPolicyMergeTypesArgs)(nil)).Elem()
}

type BranchPolicyMergeTypesInput interface {
	pulumi.Input

	ToBranchPolicyMergeTypesOutput() BranchPolicyMergeTypesOutput
	ToBranchPolicyMergeTypesOutputWithContext(ctx context.Context) BranchPolicyMergeTypesOutput
}

func (*BranchPolicyMergeTypes) ElementType() reflect.Type {
	return reflect.TypeOf((**BranchPolicyMergeTypes)(nil)).Elem()
}

func (i *BranchPolicyMergeTypes) ToBranchPolicyMergeTypesOutput() BranchPolicyMergeTypesOutput {
	return i.ToBranchPolicyMergeTypesOutputWithContext(context.Background())
}

func (i *BranchPolicyMergeTypes) ToBranchPolicyMergeTypesOutputWithContext(ctx context.Context) BranchPolicyMergeTypesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BranchPolicyMergeTypesOutput)
}

// BranchPolicyMergeTypesArrayInput is an input type that accepts BranchPolicyMergeTypesArray and BranchPolicyMergeTypesArrayOutput values.
// You can construct a concrete instance of `BranchPolicyMergeTypesArrayInput` via:
//
//	BranchPolicyMergeTypesArray{ BranchPolicyMergeTypesArgs{...} }
type BranchPolicyMergeTypesArrayInput interface {
	pulumi.Input

	ToBranchPolicyMergeTypesArrayOutput() BranchPolicyMergeTypesArrayOutput
	ToBranchPolicyMergeTypesArrayOutputWithContext(context.Context) BranchPolicyMergeTypesArrayOutput
}

type BranchPolicyMergeTypesArray []BranchPolicyMergeTypesInput

func (BranchPolicyMergeTypesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BranchPolicyMergeTypes)(nil)).Elem()
}

func (i BranchPolicyMergeTypesArray) ToBranchPolicyMergeTypesArrayOutput() BranchPolicyMergeTypesArrayOutput {
	return i.ToBranchPolicyMergeTypesArrayOutputWithContext(context.Background())
}

func (i BranchPolicyMergeTypesArray) ToBranchPolicyMergeTypesArrayOutputWithContext(ctx context.Context) BranchPolicyMergeTypesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BranchPolicyMergeTypesArrayOutput)
}

// BranchPolicyMergeTypesMapInput is an input type that accepts BranchPolicyMergeTypesMap and BranchPolicyMergeTypesMapOutput values.
// You can construct a concrete instance of `BranchPolicyMergeTypesMapInput` via:
//
//	BranchPolicyMergeTypesMap{ "key": BranchPolicyMergeTypesArgs{...} }
type BranchPolicyMergeTypesMapInput interface {
	pulumi.Input

	ToBranchPolicyMergeTypesMapOutput() BranchPolicyMergeTypesMapOutput
	ToBranchPolicyMergeTypesMapOutputWithContext(context.Context) BranchPolicyMergeTypesMapOutput
}

type BranchPolicyMergeTypesMap map[string]BranchPolicyMergeTypesInput

func (BranchPolicyMergeTypesMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BranchPolicyMergeTypes)(nil)).Elem()
}

func (i BranchPolicyMergeTypesMap) ToBranchPolicyMergeTypesMapOutput() BranchPolicyMergeTypesMapOutput {
	return i.ToBranchPolicyMergeTypesMapOutputWithContext(context.Background())
}

func (i BranchPolicyMergeTypesMap) ToBranchPolicyMergeTypesMapOutputWithContext(ctx context.Context) BranchPolicyMergeTypesMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BranchPolicyMergeTypesMapOutput)
}

type BranchPolicyMergeTypesOutput struct{ *pulumi.OutputState }

func (BranchPolicyMergeTypesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BranchPolicyMergeTypes)(nil)).Elem()
}

func (o BranchPolicyMergeTypesOutput) ToBranchPolicyMergeTypesOutput() BranchPolicyMergeTypesOutput {
	return o
}

func (o BranchPolicyMergeTypesOutput) ToBranchPolicyMergeTypesOutputWithContext(ctx context.Context) BranchPolicyMergeTypesOutput {
	return o
}

// A flag indicating if the policy should be blocking. Defaults to `true`.
func (o BranchPolicyMergeTypesOutput) Blocking() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BranchPolicyMergeTypes) pulumi.BoolPtrOutput { return v.Blocking }).(pulumi.BoolPtrOutput)
}

// A flag indicating if the policy should be enabled. Defaults to `true`.
func (o BranchPolicyMergeTypesOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BranchPolicyMergeTypes) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// The ID of the project in which the policy will be created.
func (o BranchPolicyMergeTypesOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *BranchPolicyMergeTypes) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// Configuration for the policy. This block must be defined exactly once.
func (o BranchPolicyMergeTypesOutput) Settings() BranchPolicyMergeTypesSettingsOutput {
	return o.ApplyT(func(v *BranchPolicyMergeTypes) BranchPolicyMergeTypesSettingsOutput { return v.Settings }).(BranchPolicyMergeTypesSettingsOutput)
}

type BranchPolicyMergeTypesArrayOutput struct{ *pulumi.OutputState }

func (BranchPolicyMergeTypesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BranchPolicyMergeTypes)(nil)).Elem()
}

func (o BranchPolicyMergeTypesArrayOutput) ToBranchPolicyMergeTypesArrayOutput() BranchPolicyMergeTypesArrayOutput {
	return o
}

func (o BranchPolicyMergeTypesArrayOutput) ToBranchPolicyMergeTypesArrayOutputWithContext(ctx context.Context) BranchPolicyMergeTypesArrayOutput {
	return o
}

func (o BranchPolicyMergeTypesArrayOutput) Index(i pulumi.IntInput) BranchPolicyMergeTypesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BranchPolicyMergeTypes {
		return vs[0].([]*BranchPolicyMergeTypes)[vs[1].(int)]
	}).(BranchPolicyMergeTypesOutput)
}

type BranchPolicyMergeTypesMapOutput struct{ *pulumi.OutputState }

func (BranchPolicyMergeTypesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BranchPolicyMergeTypes)(nil)).Elem()
}

func (o BranchPolicyMergeTypesMapOutput) ToBranchPolicyMergeTypesMapOutput() BranchPolicyMergeTypesMapOutput {
	return o
}

func (o BranchPolicyMergeTypesMapOutput) ToBranchPolicyMergeTypesMapOutputWithContext(ctx context.Context) BranchPolicyMergeTypesMapOutput {
	return o
}

func (o BranchPolicyMergeTypesMapOutput) MapIndex(k pulumi.StringInput) BranchPolicyMergeTypesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BranchPolicyMergeTypes {
		return vs[0].(map[string]*BranchPolicyMergeTypes)[vs[1].(string)]
	}).(BranchPolicyMergeTypesOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BranchPolicyMergeTypesInput)(nil)).Elem(), &BranchPolicyMergeTypes{})
	pulumi.RegisterInputType(reflect.TypeOf((*BranchPolicyMergeTypesArrayInput)(nil)).Elem(), BranchPolicyMergeTypesArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BranchPolicyMergeTypesMapInput)(nil)).Elem(), BranchPolicyMergeTypesMap{})
	pulumi.RegisterOutputType(BranchPolicyMergeTypesOutput{})
	pulumi.RegisterOutputType(BranchPolicyMergeTypesArrayOutput{})
	pulumi.RegisterOutputType(BranchPolicyMergeTypesMapOutput{})
}
