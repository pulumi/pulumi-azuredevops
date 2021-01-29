// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Require associations between branches and a work item within Azure DevOps.
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
// 		_, err = azuredevops.NewBranchPolicyWorkItemLinking(ctx, "branchPolicyWorkItemLinking", &azuredevops.BranchPolicyWorkItemLinkingArgs{
// 			ProjectId: project.ID(),
// 			Enabled:   pulumi.Bool(true),
// 			Blocking:  pulumi.Bool(true),
// 			Settings: &azuredevops.BranchPolicyWorkItemLinkingSettingsArgs{
// 				Scopes: azuredevops.BranchPolicyWorkItemLinkingSettingsScopeArray{
// 					&azuredevops.BranchPolicyWorkItemLinkingSettingsScopeArgs{
// 						RepositoryId:  git.ID(),
// 						RepositoryRef: git.DefaultBranch,
// 						MatchType:     pulumi.String("Exact"),
// 					},
// 					&azuredevops.BranchPolicyWorkItemLinkingSettingsScopeArgs{
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
//  $ pulumi import azuredevops:index/branchPolicyWorkItemLinking:BranchPolicyWorkItemLinking p 00000000-0000-0000-0000-000000000000/0
// ```
type BranchPolicyWorkItemLinking struct {
	pulumi.CustomResourceState

	// A flag indicating if the policy should be blocking. Defaults to `true`.
	Blocking pulumi.BoolPtrOutput `pulumi:"blocking"`
	// A flag indicating if the policy should be enabled. Defaults to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// The ID of the project in which the policy will be created.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Configuration for the policy. This block must be defined exactly once.
	Settings BranchPolicyWorkItemLinkingSettingsOutput `pulumi:"settings"`
}

// NewBranchPolicyWorkItemLinking registers a new resource with the given unique name, arguments, and options.
func NewBranchPolicyWorkItemLinking(ctx *pulumi.Context,
	name string, args *BranchPolicyWorkItemLinkingArgs, opts ...pulumi.ResourceOption) (*BranchPolicyWorkItemLinking, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.Settings == nil {
		return nil, errors.New("invalid value for required argument 'Settings'")
	}
	var resource BranchPolicyWorkItemLinking
	err := ctx.RegisterResource("azuredevops:index/branchPolicyWorkItemLinking:BranchPolicyWorkItemLinking", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBranchPolicyWorkItemLinking gets an existing BranchPolicyWorkItemLinking resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBranchPolicyWorkItemLinking(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BranchPolicyWorkItemLinkingState, opts ...pulumi.ResourceOption) (*BranchPolicyWorkItemLinking, error) {
	var resource BranchPolicyWorkItemLinking
	err := ctx.ReadResource("azuredevops:index/branchPolicyWorkItemLinking:BranchPolicyWorkItemLinking", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BranchPolicyWorkItemLinking resources.
type branchPolicyWorkItemLinkingState struct {
	// A flag indicating if the policy should be blocking. Defaults to `true`.
	Blocking *bool `pulumi:"blocking"`
	// A flag indicating if the policy should be enabled. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// The ID of the project in which the policy will be created.
	ProjectId *string `pulumi:"projectId"`
	// Configuration for the policy. This block must be defined exactly once.
	Settings *BranchPolicyWorkItemLinkingSettings `pulumi:"settings"`
}

type BranchPolicyWorkItemLinkingState struct {
	// A flag indicating if the policy should be blocking. Defaults to `true`.
	Blocking pulumi.BoolPtrInput
	// A flag indicating if the policy should be enabled. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// The ID of the project in which the policy will be created.
	ProjectId pulumi.StringPtrInput
	// Configuration for the policy. This block must be defined exactly once.
	Settings BranchPolicyWorkItemLinkingSettingsPtrInput
}

func (BranchPolicyWorkItemLinkingState) ElementType() reflect.Type {
	return reflect.TypeOf((*branchPolicyWorkItemLinkingState)(nil)).Elem()
}

type branchPolicyWorkItemLinkingArgs struct {
	// A flag indicating if the policy should be blocking. Defaults to `true`.
	Blocking *bool `pulumi:"blocking"`
	// A flag indicating if the policy should be enabled. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// The ID of the project in which the policy will be created.
	ProjectId string `pulumi:"projectId"`
	// Configuration for the policy. This block must be defined exactly once.
	Settings BranchPolicyWorkItemLinkingSettings `pulumi:"settings"`
}

// The set of arguments for constructing a BranchPolicyWorkItemLinking resource.
type BranchPolicyWorkItemLinkingArgs struct {
	// A flag indicating if the policy should be blocking. Defaults to `true`.
	Blocking pulumi.BoolPtrInput
	// A flag indicating if the policy should be enabled. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// The ID of the project in which the policy will be created.
	ProjectId pulumi.StringInput
	// Configuration for the policy. This block must be defined exactly once.
	Settings BranchPolicyWorkItemLinkingSettingsInput
}

func (BranchPolicyWorkItemLinkingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*branchPolicyWorkItemLinkingArgs)(nil)).Elem()
}

type BranchPolicyWorkItemLinkingInput interface {
	pulumi.Input

	ToBranchPolicyWorkItemLinkingOutput() BranchPolicyWorkItemLinkingOutput
	ToBranchPolicyWorkItemLinkingOutputWithContext(ctx context.Context) BranchPolicyWorkItemLinkingOutput
}

func (*BranchPolicyWorkItemLinking) ElementType() reflect.Type {
	return reflect.TypeOf((*BranchPolicyWorkItemLinking)(nil))
}

func (i *BranchPolicyWorkItemLinking) ToBranchPolicyWorkItemLinkingOutput() BranchPolicyWorkItemLinkingOutput {
	return i.ToBranchPolicyWorkItemLinkingOutputWithContext(context.Background())
}

func (i *BranchPolicyWorkItemLinking) ToBranchPolicyWorkItemLinkingOutputWithContext(ctx context.Context) BranchPolicyWorkItemLinkingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BranchPolicyWorkItemLinkingOutput)
}

type BranchPolicyWorkItemLinkingOutput struct {
	*pulumi.OutputState
}

func (BranchPolicyWorkItemLinkingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BranchPolicyWorkItemLinking)(nil))
}

func (o BranchPolicyWorkItemLinkingOutput) ToBranchPolicyWorkItemLinkingOutput() BranchPolicyWorkItemLinkingOutput {
	return o
}

func (o BranchPolicyWorkItemLinkingOutput) ToBranchPolicyWorkItemLinkingOutputWithContext(ctx context.Context) BranchPolicyWorkItemLinkingOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(BranchPolicyWorkItemLinkingOutput{})
}
