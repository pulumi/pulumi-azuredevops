// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Branch policy for reviewers on pull requests. Includes the minimum number of reviewers and other conditions.
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
// 		exampleProject, err := azuredevops.NewProject(ctx, "exampleProject", nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleGit, err := azuredevops.NewGit(ctx, "exampleGit", &azuredevops.GitArgs{
// 			ProjectId: exampleProject.ID(),
// 			Initialization: &GitInitializationArgs{
// 				InitType: pulumi.String("Clean"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = azuredevops.NewBranchPolicyMinReviewers(ctx, "exampleBranchPolicyMinReviewers", &azuredevops.BranchPolicyMinReviewersArgs{
// 			ProjectId: exampleProject.ID(),
// 			Enabled:   pulumi.Bool(true),
// 			Blocking:  pulumi.Bool(true),
// 			Settings: &BranchPolicyMinReviewersSettingsArgs{
// 				ReviewerCount:                     pulumi.Int(7),
// 				SubmitterCanVote:                  pulumi.Bool(false),
// 				LastPusherCannotApprove:           pulumi.Bool(true),
// 				AllowCompletionWithRejectsOrWaits: pulumi.Bool(false),
// 				OnPushResetApprovedVotes:          pulumi.Bool(true),
// 				OnLastIterationRequireVote:        pulumi.Bool(false),
// 				Scopes: BranchPolicyMinReviewersSettingsScopeArray{
// 					&BranchPolicyMinReviewersSettingsScopeArgs{
// 						RepositoryId:  exampleGit.ID(),
// 						RepositoryRef: exampleGit.DefaultBranch,
// 						MatchType:     pulumi.String("Exact"),
// 					},
// 					&BranchPolicyMinReviewersSettingsScopeArgs{
// 						RepositoryId:  nil,
// 						RepositoryRef: pulumi.String("refs/heads/releases"),
// 						MatchType:     pulumi.String("Prefix"),
// 					},
// 					&BranchPolicyMinReviewersSettingsScopeArgs{
// 						MatchType: pulumi.String("DefaultBranch"),
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
// - [Azure DevOps Service REST API 6.0 - Policy Configurations](https://docs.microsoft.com/en-us/rest/api/azure/devops/policy/configurations/create?view=azure-devops-rest-6.0)
//
// ## Import
//
// Azure DevOps Branch Policies can be imported using the project ID and policy configuration ID
//
// ```sh
//  $ pulumi import azuredevops:index/branchPolicyMinReviewers:BranchPolicyMinReviewers example 00000000-0000-0000-0000-000000000000/0
// ```
type BranchPolicyMinReviewers struct {
	pulumi.CustomResourceState

	// A flag indicating if the policy should be blocking. Defaults to `true`.
	Blocking pulumi.BoolPtrOutput `pulumi:"blocking"`
	// A flag indicating if the policy should be enabled. Defaults to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// The ID of the project in which the policy will be created.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Configuration for the policy. This block must be defined exactly once.
	Settings BranchPolicyMinReviewersSettingsOutput `pulumi:"settings"`
}

// NewBranchPolicyMinReviewers registers a new resource with the given unique name, arguments, and options.
func NewBranchPolicyMinReviewers(ctx *pulumi.Context,
	name string, args *BranchPolicyMinReviewersArgs, opts ...pulumi.ResourceOption) (*BranchPolicyMinReviewers, error) {
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
			Type: pulumi.String("azuredevops:Policy/branchPolicyMinReviewers:BranchPolicyMinReviewers"),
		},
	})
	opts = append(opts, aliases)
	var resource BranchPolicyMinReviewers
	err := ctx.RegisterResource("azuredevops:index/branchPolicyMinReviewers:BranchPolicyMinReviewers", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBranchPolicyMinReviewers gets an existing BranchPolicyMinReviewers resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBranchPolicyMinReviewers(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BranchPolicyMinReviewersState, opts ...pulumi.ResourceOption) (*BranchPolicyMinReviewers, error) {
	var resource BranchPolicyMinReviewers
	err := ctx.ReadResource("azuredevops:index/branchPolicyMinReviewers:BranchPolicyMinReviewers", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BranchPolicyMinReviewers resources.
type branchPolicyMinReviewersState struct {
	// A flag indicating if the policy should be blocking. Defaults to `true`.
	Blocking *bool `pulumi:"blocking"`
	// A flag indicating if the policy should be enabled. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// The ID of the project in which the policy will be created.
	ProjectId *string `pulumi:"projectId"`
	// Configuration for the policy. This block must be defined exactly once.
	Settings *BranchPolicyMinReviewersSettings `pulumi:"settings"`
}

type BranchPolicyMinReviewersState struct {
	// A flag indicating if the policy should be blocking. Defaults to `true`.
	Blocking pulumi.BoolPtrInput
	// A flag indicating if the policy should be enabled. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// The ID of the project in which the policy will be created.
	ProjectId pulumi.StringPtrInput
	// Configuration for the policy. This block must be defined exactly once.
	Settings BranchPolicyMinReviewersSettingsPtrInput
}

func (BranchPolicyMinReviewersState) ElementType() reflect.Type {
	return reflect.TypeOf((*branchPolicyMinReviewersState)(nil)).Elem()
}

type branchPolicyMinReviewersArgs struct {
	// A flag indicating if the policy should be blocking. Defaults to `true`.
	Blocking *bool `pulumi:"blocking"`
	// A flag indicating if the policy should be enabled. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// The ID of the project in which the policy will be created.
	ProjectId string `pulumi:"projectId"`
	// Configuration for the policy. This block must be defined exactly once.
	Settings BranchPolicyMinReviewersSettings `pulumi:"settings"`
}

// The set of arguments for constructing a BranchPolicyMinReviewers resource.
type BranchPolicyMinReviewersArgs struct {
	// A flag indicating if the policy should be blocking. Defaults to `true`.
	Blocking pulumi.BoolPtrInput
	// A flag indicating if the policy should be enabled. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// The ID of the project in which the policy will be created.
	ProjectId pulumi.StringInput
	// Configuration for the policy. This block must be defined exactly once.
	Settings BranchPolicyMinReviewersSettingsInput
}

func (BranchPolicyMinReviewersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*branchPolicyMinReviewersArgs)(nil)).Elem()
}

type BranchPolicyMinReviewersInput interface {
	pulumi.Input

	ToBranchPolicyMinReviewersOutput() BranchPolicyMinReviewersOutput
	ToBranchPolicyMinReviewersOutputWithContext(ctx context.Context) BranchPolicyMinReviewersOutput
}

func (*BranchPolicyMinReviewers) ElementType() reflect.Type {
	return reflect.TypeOf((**BranchPolicyMinReviewers)(nil)).Elem()
}

func (i *BranchPolicyMinReviewers) ToBranchPolicyMinReviewersOutput() BranchPolicyMinReviewersOutput {
	return i.ToBranchPolicyMinReviewersOutputWithContext(context.Background())
}

func (i *BranchPolicyMinReviewers) ToBranchPolicyMinReviewersOutputWithContext(ctx context.Context) BranchPolicyMinReviewersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BranchPolicyMinReviewersOutput)
}

// BranchPolicyMinReviewersArrayInput is an input type that accepts BranchPolicyMinReviewersArray and BranchPolicyMinReviewersArrayOutput values.
// You can construct a concrete instance of `BranchPolicyMinReviewersArrayInput` via:
//
//          BranchPolicyMinReviewersArray{ BranchPolicyMinReviewersArgs{...} }
type BranchPolicyMinReviewersArrayInput interface {
	pulumi.Input

	ToBranchPolicyMinReviewersArrayOutput() BranchPolicyMinReviewersArrayOutput
	ToBranchPolicyMinReviewersArrayOutputWithContext(context.Context) BranchPolicyMinReviewersArrayOutput
}

type BranchPolicyMinReviewersArray []BranchPolicyMinReviewersInput

func (BranchPolicyMinReviewersArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BranchPolicyMinReviewers)(nil)).Elem()
}

func (i BranchPolicyMinReviewersArray) ToBranchPolicyMinReviewersArrayOutput() BranchPolicyMinReviewersArrayOutput {
	return i.ToBranchPolicyMinReviewersArrayOutputWithContext(context.Background())
}

func (i BranchPolicyMinReviewersArray) ToBranchPolicyMinReviewersArrayOutputWithContext(ctx context.Context) BranchPolicyMinReviewersArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BranchPolicyMinReviewersArrayOutput)
}

// BranchPolicyMinReviewersMapInput is an input type that accepts BranchPolicyMinReviewersMap and BranchPolicyMinReviewersMapOutput values.
// You can construct a concrete instance of `BranchPolicyMinReviewersMapInput` via:
//
//          BranchPolicyMinReviewersMap{ "key": BranchPolicyMinReviewersArgs{...} }
type BranchPolicyMinReviewersMapInput interface {
	pulumi.Input

	ToBranchPolicyMinReviewersMapOutput() BranchPolicyMinReviewersMapOutput
	ToBranchPolicyMinReviewersMapOutputWithContext(context.Context) BranchPolicyMinReviewersMapOutput
}

type BranchPolicyMinReviewersMap map[string]BranchPolicyMinReviewersInput

func (BranchPolicyMinReviewersMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BranchPolicyMinReviewers)(nil)).Elem()
}

func (i BranchPolicyMinReviewersMap) ToBranchPolicyMinReviewersMapOutput() BranchPolicyMinReviewersMapOutput {
	return i.ToBranchPolicyMinReviewersMapOutputWithContext(context.Background())
}

func (i BranchPolicyMinReviewersMap) ToBranchPolicyMinReviewersMapOutputWithContext(ctx context.Context) BranchPolicyMinReviewersMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BranchPolicyMinReviewersMapOutput)
}

type BranchPolicyMinReviewersOutput struct{ *pulumi.OutputState }

func (BranchPolicyMinReviewersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BranchPolicyMinReviewers)(nil)).Elem()
}

func (o BranchPolicyMinReviewersOutput) ToBranchPolicyMinReviewersOutput() BranchPolicyMinReviewersOutput {
	return o
}

func (o BranchPolicyMinReviewersOutput) ToBranchPolicyMinReviewersOutputWithContext(ctx context.Context) BranchPolicyMinReviewersOutput {
	return o
}

// A flag indicating if the policy should be blocking. Defaults to `true`.
func (o BranchPolicyMinReviewersOutput) Blocking() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BranchPolicyMinReviewers) pulumi.BoolPtrOutput { return v.Blocking }).(pulumi.BoolPtrOutput)
}

// A flag indicating if the policy should be enabled. Defaults to `true`.
func (o BranchPolicyMinReviewersOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BranchPolicyMinReviewers) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// The ID of the project in which the policy will be created.
func (o BranchPolicyMinReviewersOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *BranchPolicyMinReviewers) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// Configuration for the policy. This block must be defined exactly once.
func (o BranchPolicyMinReviewersOutput) Settings() BranchPolicyMinReviewersSettingsOutput {
	return o.ApplyT(func(v *BranchPolicyMinReviewers) BranchPolicyMinReviewersSettingsOutput { return v.Settings }).(BranchPolicyMinReviewersSettingsOutput)
}

type BranchPolicyMinReviewersArrayOutput struct{ *pulumi.OutputState }

func (BranchPolicyMinReviewersArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BranchPolicyMinReviewers)(nil)).Elem()
}

func (o BranchPolicyMinReviewersArrayOutput) ToBranchPolicyMinReviewersArrayOutput() BranchPolicyMinReviewersArrayOutput {
	return o
}

func (o BranchPolicyMinReviewersArrayOutput) ToBranchPolicyMinReviewersArrayOutputWithContext(ctx context.Context) BranchPolicyMinReviewersArrayOutput {
	return o
}

func (o BranchPolicyMinReviewersArrayOutput) Index(i pulumi.IntInput) BranchPolicyMinReviewersOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BranchPolicyMinReviewers {
		return vs[0].([]*BranchPolicyMinReviewers)[vs[1].(int)]
	}).(BranchPolicyMinReviewersOutput)
}

type BranchPolicyMinReviewersMapOutput struct{ *pulumi.OutputState }

func (BranchPolicyMinReviewersMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BranchPolicyMinReviewers)(nil)).Elem()
}

func (o BranchPolicyMinReviewersMapOutput) ToBranchPolicyMinReviewersMapOutput() BranchPolicyMinReviewersMapOutput {
	return o
}

func (o BranchPolicyMinReviewersMapOutput) ToBranchPolicyMinReviewersMapOutputWithContext(ctx context.Context) BranchPolicyMinReviewersMapOutput {
	return o
}

func (o BranchPolicyMinReviewersMapOutput) MapIndex(k pulumi.StringInput) BranchPolicyMinReviewersOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BranchPolicyMinReviewers {
		return vs[0].(map[string]*BranchPolicyMinReviewers)[vs[1].(string)]
	}).(BranchPolicyMinReviewersOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BranchPolicyMinReviewersInput)(nil)).Elem(), &BranchPolicyMinReviewers{})
	pulumi.RegisterInputType(reflect.TypeOf((*BranchPolicyMinReviewersArrayInput)(nil)).Elem(), BranchPolicyMinReviewersArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BranchPolicyMinReviewersMapInput)(nil)).Elem(), BranchPolicyMinReviewersMap{})
	pulumi.RegisterOutputType(BranchPolicyMinReviewersOutput{})
	pulumi.RegisterOutputType(BranchPolicyMinReviewersArrayOutput{})
	pulumi.RegisterOutputType(BranchPolicyMinReviewersMapOutput{})
}
