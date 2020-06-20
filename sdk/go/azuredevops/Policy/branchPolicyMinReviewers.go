// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package Policy

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a minimum reviewer branch policy within Azure DevOps.
//
// ## Example Usage
//
//
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azuredevops/sdk/go/azuredevops/Core"
// 	"github.com/pulumi/pulumi-azuredevops/sdk/go/azuredevops/Policy"
// 	"github.com/pulumi/pulumi-azuredevops/sdk/go/azuredevops/Repository"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		project, err := Core.NewProject(ctx, "project", &Core.ProjectArgs{
// 			ProjectName: pulumi.String("Sample Project"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		git, err := Repository.NewGit(ctx, "git", &Repository.GitArgs{
// 			ProjectId: project.ID(),
// 			Initialization: &Repository.GitInitializationArgs{
// 				InitType: pulumi.String("Clean"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		branchPolicyMinReviewers, err := Policy.NewBranchPolicyMinReviewers(ctx, "branchPolicyMinReviewers", &Policy.BranchPolicyMinReviewersArgs{
// 			ProjectId: project.ID(),
// 			Enabled:   pulumi.Bool(true),
// 			Blocking:  pulumi.Bool(true),
// 			Settings: &Policy.BranchPolicyMinReviewersSettingsArgs{
// 				ReviewerCount:    pulumi.Int(2),
// 				SubmitterCanVote: pulumi.Bool(false),
// 				Scope: []interface{}{
// 					map[string]interface{}{
// 						"repositoryId":  git.ID(),
// 						"repositoryRef": git.DefaultBranch,
// 						"matchType":     "Exact",
// 					},
// 					map[string]interface{}{
// 						"repositoryId":  git.ID(),
// 						"repositoryRef": "refs/heads/releases",
// 						"matchType":     "Prefix",
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
//
// ## Relevant Links
//
// * [Azure DevOps Service REST API 5.1 - Policy Configurations](https://docs.microsoft.com/en-us/rest/api/azure/devops/policy/configurations/create?view=azure-devops-rest-5.1)
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
	if args == nil || args.ProjectId == nil {
		return nil, errors.New("missing required argument 'ProjectId'")
	}
	if args == nil || args.Settings == nil {
		return nil, errors.New("missing required argument 'Settings'")
	}
	if args == nil {
		args = &BranchPolicyMinReviewersArgs{}
	}
	var resource BranchPolicyMinReviewers
	err := ctx.RegisterResource("azuredevops:Policy/branchPolicyMinReviewers:BranchPolicyMinReviewers", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azuredevops:Policy/branchPolicyMinReviewers:BranchPolicyMinReviewers", name, id, state, &resource, opts...)
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
