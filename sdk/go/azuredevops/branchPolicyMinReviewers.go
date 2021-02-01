// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a minimum reviewer branch policy within Azure DevOps.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azuredevops/sdk/go/azuredevops"
// 	"github.com/pulumi/pulumi-azuredevops/sdk/go/azuredevops/"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
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
// 		_, err = azuredevops.NewBranchPolicyMinReviewers(ctx, "branchPolicyMinReviewers", &azuredevops.BranchPolicyMinReviewersArgs{
// 			ProjectId: project.ID(),
// 			Enabled:   pulumi.Bool(true),
// 			Blocking:  pulumi.Bool(true),
// 			Settings: &azuredevops.BranchPolicyMinReviewersSettingsArgs{
// 				ReviewerCount:    pulumi.Int(2),
// 				SubmitterCanVote: pulumi.Bool(false),
// 				Scopes: azuredevops.BranchPolicyMinReviewersSettingsScopeArray{
// 					&azuredevops.BranchPolicyMinReviewersSettingsScopeArgs{
// 						RepositoryId:  git.ID(),
// 						RepositoryRef: git.DefaultBranch,
// 						MatchType:     pulumi.String("Exact"),
// 					},
// 					&azuredevops.BranchPolicyMinReviewersSettingsScopeArgs{
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
//  $ pulumi import azuredevops:index/branchPolicyMinReviewers:BranchPolicyMinReviewers p 00000000-0000-0000-0000-000000000000/0
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
	return reflect.TypeOf((*BranchPolicyMinReviewers)(nil))
}

func (i *BranchPolicyMinReviewers) ToBranchPolicyMinReviewersOutput() BranchPolicyMinReviewersOutput {
	return i.ToBranchPolicyMinReviewersOutputWithContext(context.Background())
}

func (i *BranchPolicyMinReviewers) ToBranchPolicyMinReviewersOutputWithContext(ctx context.Context) BranchPolicyMinReviewersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BranchPolicyMinReviewersOutput)
}

type BranchPolicyMinReviewersOutput struct {
	*pulumi.OutputState
}

func (BranchPolicyMinReviewersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BranchPolicyMinReviewers)(nil))
}

func (o BranchPolicyMinReviewersOutput) ToBranchPolicyMinReviewersOutput() BranchPolicyMinReviewersOutput {
	return o
}

func (o BranchPolicyMinReviewersOutput) ToBranchPolicyMinReviewersOutputWithContext(ctx context.Context) BranchPolicyMinReviewersOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(BranchPolicyMinReviewersOutput{})
}
