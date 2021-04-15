// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package policy

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a build validation branch policy within Azure DevOps.
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
// 		project, err := azuredevops.NewProject(ctx, "project", nil)
// 		if err != nil {
// 			return err
// 		}
// 		git, err := azuredevops.NewGit(ctx, "git", &azuredevops.GitArgs{
// 			ProjectId: project.ID(),
// 			Initialization: &azuredevops.GitInitializationArgs{
// 				InitType: pulumi.String("Clean"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		buildDefinition, err := azuredevops.NewBuildDefinition(ctx, "buildDefinition", &azuredevops.BuildDefinitionArgs{
// 			ProjectId: project.ID(),
// 			Repository: &azuredevops.BuildDefinitionRepositoryArgs{
// 				RepoType: pulumi.String("TfsGit"),
// 				RepoId:   git.ID(),
// 				YmlPath:  pulumi.String("azure-pipelines.yml"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = azuredevops.NewBranchPolicyBuildValidation(ctx, "branchPolicyBuildValidation", &azuredevops.BranchPolicyBuildValidationArgs{
// 			ProjectId: project.ID(),
// 			Enabled:   pulumi.Bool(true),
// 			Blocking:  pulumi.Bool(true),
// 			Settings: &azuredevops.BranchPolicyBuildValidationSettingsArgs{
// 				DisplayName:       pulumi.String("Don't break the build!"),
// 				BuildDefinitionId: buildDefinition.ID(),
// 				ValidDuration:     pulumi.Int(720),
// 				FilenamePatterns: pulumi.StringArray{
// 					pulumi.String("/WebApp/*"),
// 					pulumi.String("!/WebApp/Tests/*"),
// 					pulumi.String("*.cs"),
// 				},
// 				Scopes: azuredevops.BranchPolicyBuildValidationSettingsScopeArray{
// 					&azuredevops.BranchPolicyBuildValidationSettingsScopeArgs{
// 						RepositoryId:  git.ID(),
// 						RepositoryRef: git.DefaultBranch,
// 						MatchType:     pulumi.String("Exact"),
// 					},
// 					&azuredevops.BranchPolicyBuildValidationSettingsScopeArgs{
// 						RepositoryId:  git.ID(),
// 						RepositoryRef: pulumi.String("refs/heads/releases"),
// 						MatchType:     pulumi.String("Prefix"),
// 					},
// 				},
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
// - [Azure DevOps Service REST API 5.1 - Policy Configurations](https://docs.microsoft.com/en-us/rest/api/azure/devops/policy/configurations/create?view=azure-devops-rest-5.1)
//
// ## Import
//
// Azure DevOps Branch Policies can be imported using the project ID and policy configuration ID
//
// ```sh
//  $ pulumi import azuredevops:Policy/branchPolicyBuildValidation:BranchPolicyBuildValidation p 00000000-0000-0000-0000-000000000000/0
// ```
//
// Deprecated: azuredevops.policy.BranchPolicyBuildValidation has been deprecated in favor of azuredevops.BranchPolicyBuildValidation
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
	var resource BranchPolicyBuildValidation
	err := ctx.RegisterResource("azuredevops:Policy/branchPolicyBuildValidation:BranchPolicyBuildValidation", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azuredevops:Policy/branchPolicyBuildValidation:BranchPolicyBuildValidation", name, id, state, &resource, opts...)
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
	return reflect.TypeOf((*BranchPolicyBuildValidation)(nil))
}

func (i *BranchPolicyBuildValidation) ToBranchPolicyBuildValidationOutput() BranchPolicyBuildValidationOutput {
	return i.ToBranchPolicyBuildValidationOutputWithContext(context.Background())
}

func (i *BranchPolicyBuildValidation) ToBranchPolicyBuildValidationOutputWithContext(ctx context.Context) BranchPolicyBuildValidationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BranchPolicyBuildValidationOutput)
}

func (i *BranchPolicyBuildValidation) ToBranchPolicyBuildValidationPtrOutput() BranchPolicyBuildValidationPtrOutput {
	return i.ToBranchPolicyBuildValidationPtrOutputWithContext(context.Background())
}

func (i *BranchPolicyBuildValidation) ToBranchPolicyBuildValidationPtrOutputWithContext(ctx context.Context) BranchPolicyBuildValidationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BranchPolicyBuildValidationPtrOutput)
}

type BranchPolicyBuildValidationPtrInput interface {
	pulumi.Input

	ToBranchPolicyBuildValidationPtrOutput() BranchPolicyBuildValidationPtrOutput
	ToBranchPolicyBuildValidationPtrOutputWithContext(ctx context.Context) BranchPolicyBuildValidationPtrOutput
}

type branchPolicyBuildValidationPtrType BranchPolicyBuildValidationArgs

func (*branchPolicyBuildValidationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BranchPolicyBuildValidation)(nil))
}

func (i *branchPolicyBuildValidationPtrType) ToBranchPolicyBuildValidationPtrOutput() BranchPolicyBuildValidationPtrOutput {
	return i.ToBranchPolicyBuildValidationPtrOutputWithContext(context.Background())
}

func (i *branchPolicyBuildValidationPtrType) ToBranchPolicyBuildValidationPtrOutputWithContext(ctx context.Context) BranchPolicyBuildValidationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BranchPolicyBuildValidationPtrOutput)
}

// BranchPolicyBuildValidationArrayInput is an input type that accepts BranchPolicyBuildValidationArray and BranchPolicyBuildValidationArrayOutput values.
// You can construct a concrete instance of `BranchPolicyBuildValidationArrayInput` via:
//
//          BranchPolicyBuildValidationArray{ BranchPolicyBuildValidationArgs{...} }
type BranchPolicyBuildValidationArrayInput interface {
	pulumi.Input

	ToBranchPolicyBuildValidationArrayOutput() BranchPolicyBuildValidationArrayOutput
	ToBranchPolicyBuildValidationArrayOutputWithContext(context.Context) BranchPolicyBuildValidationArrayOutput
}

type BranchPolicyBuildValidationArray []BranchPolicyBuildValidationInput

func (BranchPolicyBuildValidationArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*BranchPolicyBuildValidation)(nil))
}

func (i BranchPolicyBuildValidationArray) ToBranchPolicyBuildValidationArrayOutput() BranchPolicyBuildValidationArrayOutput {
	return i.ToBranchPolicyBuildValidationArrayOutputWithContext(context.Background())
}

func (i BranchPolicyBuildValidationArray) ToBranchPolicyBuildValidationArrayOutputWithContext(ctx context.Context) BranchPolicyBuildValidationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BranchPolicyBuildValidationArrayOutput)
}

// BranchPolicyBuildValidationMapInput is an input type that accepts BranchPolicyBuildValidationMap and BranchPolicyBuildValidationMapOutput values.
// You can construct a concrete instance of `BranchPolicyBuildValidationMapInput` via:
//
//          BranchPolicyBuildValidationMap{ "key": BranchPolicyBuildValidationArgs{...} }
type BranchPolicyBuildValidationMapInput interface {
	pulumi.Input

	ToBranchPolicyBuildValidationMapOutput() BranchPolicyBuildValidationMapOutput
	ToBranchPolicyBuildValidationMapOutputWithContext(context.Context) BranchPolicyBuildValidationMapOutput
}

type BranchPolicyBuildValidationMap map[string]BranchPolicyBuildValidationInput

func (BranchPolicyBuildValidationMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*BranchPolicyBuildValidation)(nil))
}

func (i BranchPolicyBuildValidationMap) ToBranchPolicyBuildValidationMapOutput() BranchPolicyBuildValidationMapOutput {
	return i.ToBranchPolicyBuildValidationMapOutputWithContext(context.Background())
}

func (i BranchPolicyBuildValidationMap) ToBranchPolicyBuildValidationMapOutputWithContext(ctx context.Context) BranchPolicyBuildValidationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BranchPolicyBuildValidationMapOutput)
}

type BranchPolicyBuildValidationOutput struct {
	*pulumi.OutputState
}

func (BranchPolicyBuildValidationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BranchPolicyBuildValidation)(nil))
}

func (o BranchPolicyBuildValidationOutput) ToBranchPolicyBuildValidationOutput() BranchPolicyBuildValidationOutput {
	return o
}

func (o BranchPolicyBuildValidationOutput) ToBranchPolicyBuildValidationOutputWithContext(ctx context.Context) BranchPolicyBuildValidationOutput {
	return o
}

func (o BranchPolicyBuildValidationOutput) ToBranchPolicyBuildValidationPtrOutput() BranchPolicyBuildValidationPtrOutput {
	return o.ToBranchPolicyBuildValidationPtrOutputWithContext(context.Background())
}

func (o BranchPolicyBuildValidationOutput) ToBranchPolicyBuildValidationPtrOutputWithContext(ctx context.Context) BranchPolicyBuildValidationPtrOutput {
	return o.ApplyT(func(v BranchPolicyBuildValidation) *BranchPolicyBuildValidation {
		return &v
	}).(BranchPolicyBuildValidationPtrOutput)
}

type BranchPolicyBuildValidationPtrOutput struct {
	*pulumi.OutputState
}

func (BranchPolicyBuildValidationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BranchPolicyBuildValidation)(nil))
}

func (o BranchPolicyBuildValidationPtrOutput) ToBranchPolicyBuildValidationPtrOutput() BranchPolicyBuildValidationPtrOutput {
	return o
}

func (o BranchPolicyBuildValidationPtrOutput) ToBranchPolicyBuildValidationPtrOutputWithContext(ctx context.Context) BranchPolicyBuildValidationPtrOutput {
	return o
}

type BranchPolicyBuildValidationArrayOutput struct{ *pulumi.OutputState }

func (BranchPolicyBuildValidationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BranchPolicyBuildValidation)(nil))
}

func (o BranchPolicyBuildValidationArrayOutput) ToBranchPolicyBuildValidationArrayOutput() BranchPolicyBuildValidationArrayOutput {
	return o
}

func (o BranchPolicyBuildValidationArrayOutput) ToBranchPolicyBuildValidationArrayOutputWithContext(ctx context.Context) BranchPolicyBuildValidationArrayOutput {
	return o
}

func (o BranchPolicyBuildValidationArrayOutput) Index(i pulumi.IntInput) BranchPolicyBuildValidationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) BranchPolicyBuildValidation {
		return vs[0].([]BranchPolicyBuildValidation)[vs[1].(int)]
	}).(BranchPolicyBuildValidationOutput)
}

type BranchPolicyBuildValidationMapOutput struct{ *pulumi.OutputState }

func (BranchPolicyBuildValidationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]BranchPolicyBuildValidation)(nil))
}

func (o BranchPolicyBuildValidationMapOutput) ToBranchPolicyBuildValidationMapOutput() BranchPolicyBuildValidationMapOutput {
	return o
}

func (o BranchPolicyBuildValidationMapOutput) ToBranchPolicyBuildValidationMapOutputWithContext(ctx context.Context) BranchPolicyBuildValidationMapOutput {
	return o
}

func (o BranchPolicyBuildValidationMapOutput) MapIndex(k pulumi.StringInput) BranchPolicyBuildValidationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) BranchPolicyBuildValidation {
		return vs[0].(map[string]BranchPolicyBuildValidation)[vs[1].(string)]
	}).(BranchPolicyBuildValidationOutput)
}

func init() {
	pulumi.RegisterOutputType(BranchPolicyBuildValidationOutput{})
	pulumi.RegisterOutputType(BranchPolicyBuildValidationPtrOutput{})
	pulumi.RegisterOutputType(BranchPolicyBuildValidationArrayOutput{})
	pulumi.RegisterOutputType(BranchPolicyBuildValidationMapOutput{})
}
