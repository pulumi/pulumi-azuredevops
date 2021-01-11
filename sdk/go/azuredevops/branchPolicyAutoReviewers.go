// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages required reviewer policy branch policy within Azure DevOps.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azuredevops/sdk/go/azuredevops"
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
// 		user, err := azuredevops.NewUser(ctx, "user", &azuredevops.UserArgs{
// 			PrincipalName:      pulumi.String("mail@email.com"),
// 			AccountLicenseType: pulumi.String("basic"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = azuredevops.NewBranchPolicyAutoReviewers(ctx, "branchPolicyAutoReviewers", &azuredevops.BranchPolicyAutoReviewersArgs{
// 			ProjectId: project.ID(),
// 			Enabled:   pulumi.Bool(true),
// 			Blocking:  pulumi.Bool(true),
// 			Settings: &azuredevops.BranchPolicyAutoReviewersSettingsArgs{
// 				AutoReviewerIds: pulumi.StringArray{
// 					user.ID(),
// 				},
// 				SubmitterCanVote: pulumi.Bool(false),
// 				Message:          pulumi.String("Auto reviewer"),
// 				PathFilters: pulumi.StringArray{
// 					pulumi.String("*/src/*.ts"),
// 				},
// 				Scopes: azuredevops.BranchPolicyAutoReviewersSettingsScopeArray{
// 					&azuredevops.BranchPolicyAutoReviewersSettingsScopeArgs{
// 						RepositoryId:  git.ID(),
// 						RepositoryRef: git.DefaultBranch,
// 						MatchType:     pulumi.String("Exact"),
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
//  $ pulumi import azuredevops:index/branchPolicyAutoReviewers:BranchPolicyAutoReviewers p 00000000-0000-0000-0000-000000000000/0
// ```
type BranchPolicyAutoReviewers struct {
	pulumi.CustomResourceState

	// A flag indicating if the policy should be blocking. Defaults to `true`.
	Blocking pulumi.BoolPtrOutput `pulumi:"blocking"`
	// A flag indicating if the policy should be enabled. Defaults to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// The ID of the project in which the policy will be created.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Configuration for the policy. This block must be defined exactly once.
	Settings BranchPolicyAutoReviewersSettingsOutput `pulumi:"settings"`
}

// NewBranchPolicyAutoReviewers registers a new resource with the given unique name, arguments, and options.
func NewBranchPolicyAutoReviewers(ctx *pulumi.Context,
	name string, args *BranchPolicyAutoReviewersArgs, opts ...pulumi.ResourceOption) (*BranchPolicyAutoReviewers, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.Settings == nil {
		return nil, errors.New("invalid value for required argument 'Settings'")
	}
	var resource BranchPolicyAutoReviewers
	err := ctx.RegisterResource("azuredevops:index/branchPolicyAutoReviewers:BranchPolicyAutoReviewers", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBranchPolicyAutoReviewers gets an existing BranchPolicyAutoReviewers resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBranchPolicyAutoReviewers(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BranchPolicyAutoReviewersState, opts ...pulumi.ResourceOption) (*BranchPolicyAutoReviewers, error) {
	var resource BranchPolicyAutoReviewers
	err := ctx.ReadResource("azuredevops:index/branchPolicyAutoReviewers:BranchPolicyAutoReviewers", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BranchPolicyAutoReviewers resources.
type branchPolicyAutoReviewersState struct {
	// A flag indicating if the policy should be blocking. Defaults to `true`.
	Blocking *bool `pulumi:"blocking"`
	// A flag indicating if the policy should be enabled. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// The ID of the project in which the policy will be created.
	ProjectId *string `pulumi:"projectId"`
	// Configuration for the policy. This block must be defined exactly once.
	Settings *BranchPolicyAutoReviewersSettings `pulumi:"settings"`
}

type BranchPolicyAutoReviewersState struct {
	// A flag indicating if the policy should be blocking. Defaults to `true`.
	Blocking pulumi.BoolPtrInput
	// A flag indicating if the policy should be enabled. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// The ID of the project in which the policy will be created.
	ProjectId pulumi.StringPtrInput
	// Configuration for the policy. This block must be defined exactly once.
	Settings BranchPolicyAutoReviewersSettingsPtrInput
}

func (BranchPolicyAutoReviewersState) ElementType() reflect.Type {
	return reflect.TypeOf((*branchPolicyAutoReviewersState)(nil)).Elem()
}

type branchPolicyAutoReviewersArgs struct {
	// A flag indicating if the policy should be blocking. Defaults to `true`.
	Blocking *bool `pulumi:"blocking"`
	// A flag indicating if the policy should be enabled. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// The ID of the project in which the policy will be created.
	ProjectId string `pulumi:"projectId"`
	// Configuration for the policy. This block must be defined exactly once.
	Settings BranchPolicyAutoReviewersSettings `pulumi:"settings"`
}

// The set of arguments for constructing a BranchPolicyAutoReviewers resource.
type BranchPolicyAutoReviewersArgs struct {
	// A flag indicating if the policy should be blocking. Defaults to `true`.
	Blocking pulumi.BoolPtrInput
	// A flag indicating if the policy should be enabled. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// The ID of the project in which the policy will be created.
	ProjectId pulumi.StringInput
	// Configuration for the policy. This block must be defined exactly once.
	Settings BranchPolicyAutoReviewersSettingsInput
}

func (BranchPolicyAutoReviewersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*branchPolicyAutoReviewersArgs)(nil)).Elem()
}

type BranchPolicyAutoReviewersInput interface {
	pulumi.Input

	ToBranchPolicyAutoReviewersOutput() BranchPolicyAutoReviewersOutput
	ToBranchPolicyAutoReviewersOutputWithContext(ctx context.Context) BranchPolicyAutoReviewersOutput
}

func (BranchPolicyAutoReviewers) ElementType() reflect.Type {
	return reflect.TypeOf((*BranchPolicyAutoReviewers)(nil)).Elem()
}

func (i BranchPolicyAutoReviewers) ToBranchPolicyAutoReviewersOutput() BranchPolicyAutoReviewersOutput {
	return i.ToBranchPolicyAutoReviewersOutputWithContext(context.Background())
}

func (i BranchPolicyAutoReviewers) ToBranchPolicyAutoReviewersOutputWithContext(ctx context.Context) BranchPolicyAutoReviewersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BranchPolicyAutoReviewersOutput)
}

type BranchPolicyAutoReviewersOutput struct {
	*pulumi.OutputState
}

func (BranchPolicyAutoReviewersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BranchPolicyAutoReviewersOutput)(nil)).Elem()
}

func (o BranchPolicyAutoReviewersOutput) ToBranchPolicyAutoReviewersOutput() BranchPolicyAutoReviewersOutput {
	return o
}

func (o BranchPolicyAutoReviewersOutput) ToBranchPolicyAutoReviewersOutputWithContext(ctx context.Context) BranchPolicyAutoReviewersOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(BranchPolicyAutoReviewersOutput{})
}
