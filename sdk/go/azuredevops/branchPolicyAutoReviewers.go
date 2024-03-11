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

// Manages required reviewer policy branch policy within Azure DevOps.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
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
//			exampleUser, err := azuredevops.NewUser(ctx, "exampleUser", &azuredevops.UserArgs{
//				PrincipalName:      pulumi.String("mail@email.com"),
//				AccountLicenseType: pulumi.String("basic"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuredevops.NewBranchPolicyAutoReviewers(ctx, "exampleBranchPolicyAutoReviewers", &azuredevops.BranchPolicyAutoReviewersArgs{
//				ProjectId: exampleProject.ID(),
//				Enabled:   pulumi.Bool(true),
//				Blocking:  pulumi.Bool(true),
//				Settings: &azuredevops.BranchPolicyAutoReviewersSettingsArgs{
//					AutoReviewerIds: pulumi.StringArray{
//						exampleUser.ID(),
//					},
//					SubmitterCanVote: pulumi.Bool(false),
//					Message:          pulumi.String("Auto reviewer"),
//					PathFilters: pulumi.StringArray{
//						pulumi.String("*/src/*.ts"),
//					},
//					Scopes: azuredevops.BranchPolicyAutoReviewersSettingsScopeArray{
//						&azuredevops.BranchPolicyAutoReviewersSettingsScopeArgs{
//							RepositoryId:  exampleGit.ID(),
//							RepositoryRef: exampleGit.DefaultBranch,
//							MatchType:     pulumi.String("Exact"),
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
// <!--End PulumiCodeChooser -->
//
// ## Relevant Links
//
// - [Azure DevOps Service REST API 7.0 - Policy Configurations](https://docs.microsoft.com/en-us/rest/api/azure/devops/policy/configurations/create?view=azure-devops-rest-7.0)
//
// ## Import
//
// Azure DevOps Branch Policies can be imported using the project ID and policy configuration ID:
//
// ```sh
// $ pulumi import azuredevops:index/branchPolicyAutoReviewers:BranchPolicyAutoReviewers example 00000000-0000-0000-0000-000000000000/0
// ```
type BranchPolicyAutoReviewers struct {
	pulumi.CustomResourceState

	// A flag indicating if the policy should be blocking. This relates to the Azure DevOps terms "optional" and "required" reviewers. Defaults to `true`.
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
	opts = internal.PkgResourceDefaultOpts(opts)
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
	// A flag indicating if the policy should be blocking. This relates to the Azure DevOps terms "optional" and "required" reviewers. Defaults to `true`.
	Blocking *bool `pulumi:"blocking"`
	// A flag indicating if the policy should be enabled. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// The ID of the project in which the policy will be created.
	ProjectId *string `pulumi:"projectId"`
	// Configuration for the policy. This block must be defined exactly once.
	Settings *BranchPolicyAutoReviewersSettings `pulumi:"settings"`
}

type BranchPolicyAutoReviewersState struct {
	// A flag indicating if the policy should be blocking. This relates to the Azure DevOps terms "optional" and "required" reviewers. Defaults to `true`.
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
	// A flag indicating if the policy should be blocking. This relates to the Azure DevOps terms "optional" and "required" reviewers. Defaults to `true`.
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
	// A flag indicating if the policy should be blocking. This relates to the Azure DevOps terms "optional" and "required" reviewers. Defaults to `true`.
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

func (*BranchPolicyAutoReviewers) ElementType() reflect.Type {
	return reflect.TypeOf((**BranchPolicyAutoReviewers)(nil)).Elem()
}

func (i *BranchPolicyAutoReviewers) ToBranchPolicyAutoReviewersOutput() BranchPolicyAutoReviewersOutput {
	return i.ToBranchPolicyAutoReviewersOutputWithContext(context.Background())
}

func (i *BranchPolicyAutoReviewers) ToBranchPolicyAutoReviewersOutputWithContext(ctx context.Context) BranchPolicyAutoReviewersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BranchPolicyAutoReviewersOutput)
}

// BranchPolicyAutoReviewersArrayInput is an input type that accepts BranchPolicyAutoReviewersArray and BranchPolicyAutoReviewersArrayOutput values.
// You can construct a concrete instance of `BranchPolicyAutoReviewersArrayInput` via:
//
//	BranchPolicyAutoReviewersArray{ BranchPolicyAutoReviewersArgs{...} }
type BranchPolicyAutoReviewersArrayInput interface {
	pulumi.Input

	ToBranchPolicyAutoReviewersArrayOutput() BranchPolicyAutoReviewersArrayOutput
	ToBranchPolicyAutoReviewersArrayOutputWithContext(context.Context) BranchPolicyAutoReviewersArrayOutput
}

type BranchPolicyAutoReviewersArray []BranchPolicyAutoReviewersInput

func (BranchPolicyAutoReviewersArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BranchPolicyAutoReviewers)(nil)).Elem()
}

func (i BranchPolicyAutoReviewersArray) ToBranchPolicyAutoReviewersArrayOutput() BranchPolicyAutoReviewersArrayOutput {
	return i.ToBranchPolicyAutoReviewersArrayOutputWithContext(context.Background())
}

func (i BranchPolicyAutoReviewersArray) ToBranchPolicyAutoReviewersArrayOutputWithContext(ctx context.Context) BranchPolicyAutoReviewersArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BranchPolicyAutoReviewersArrayOutput)
}

// BranchPolicyAutoReviewersMapInput is an input type that accepts BranchPolicyAutoReviewersMap and BranchPolicyAutoReviewersMapOutput values.
// You can construct a concrete instance of `BranchPolicyAutoReviewersMapInput` via:
//
//	BranchPolicyAutoReviewersMap{ "key": BranchPolicyAutoReviewersArgs{...} }
type BranchPolicyAutoReviewersMapInput interface {
	pulumi.Input

	ToBranchPolicyAutoReviewersMapOutput() BranchPolicyAutoReviewersMapOutput
	ToBranchPolicyAutoReviewersMapOutputWithContext(context.Context) BranchPolicyAutoReviewersMapOutput
}

type BranchPolicyAutoReviewersMap map[string]BranchPolicyAutoReviewersInput

func (BranchPolicyAutoReviewersMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BranchPolicyAutoReviewers)(nil)).Elem()
}

func (i BranchPolicyAutoReviewersMap) ToBranchPolicyAutoReviewersMapOutput() BranchPolicyAutoReviewersMapOutput {
	return i.ToBranchPolicyAutoReviewersMapOutputWithContext(context.Background())
}

func (i BranchPolicyAutoReviewersMap) ToBranchPolicyAutoReviewersMapOutputWithContext(ctx context.Context) BranchPolicyAutoReviewersMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BranchPolicyAutoReviewersMapOutput)
}

type BranchPolicyAutoReviewersOutput struct{ *pulumi.OutputState }

func (BranchPolicyAutoReviewersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BranchPolicyAutoReviewers)(nil)).Elem()
}

func (o BranchPolicyAutoReviewersOutput) ToBranchPolicyAutoReviewersOutput() BranchPolicyAutoReviewersOutput {
	return o
}

func (o BranchPolicyAutoReviewersOutput) ToBranchPolicyAutoReviewersOutputWithContext(ctx context.Context) BranchPolicyAutoReviewersOutput {
	return o
}

// A flag indicating if the policy should be blocking. This relates to the Azure DevOps terms "optional" and "required" reviewers. Defaults to `true`.
func (o BranchPolicyAutoReviewersOutput) Blocking() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BranchPolicyAutoReviewers) pulumi.BoolPtrOutput { return v.Blocking }).(pulumi.BoolPtrOutput)
}

// A flag indicating if the policy should be enabled. Defaults to `true`.
func (o BranchPolicyAutoReviewersOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BranchPolicyAutoReviewers) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// The ID of the project in which the policy will be created.
func (o BranchPolicyAutoReviewersOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *BranchPolicyAutoReviewers) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// Configuration for the policy. This block must be defined exactly once.
func (o BranchPolicyAutoReviewersOutput) Settings() BranchPolicyAutoReviewersSettingsOutput {
	return o.ApplyT(func(v *BranchPolicyAutoReviewers) BranchPolicyAutoReviewersSettingsOutput { return v.Settings }).(BranchPolicyAutoReviewersSettingsOutput)
}

type BranchPolicyAutoReviewersArrayOutput struct{ *pulumi.OutputState }

func (BranchPolicyAutoReviewersArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BranchPolicyAutoReviewers)(nil)).Elem()
}

func (o BranchPolicyAutoReviewersArrayOutput) ToBranchPolicyAutoReviewersArrayOutput() BranchPolicyAutoReviewersArrayOutput {
	return o
}

func (o BranchPolicyAutoReviewersArrayOutput) ToBranchPolicyAutoReviewersArrayOutputWithContext(ctx context.Context) BranchPolicyAutoReviewersArrayOutput {
	return o
}

func (o BranchPolicyAutoReviewersArrayOutput) Index(i pulumi.IntInput) BranchPolicyAutoReviewersOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BranchPolicyAutoReviewers {
		return vs[0].([]*BranchPolicyAutoReviewers)[vs[1].(int)]
	}).(BranchPolicyAutoReviewersOutput)
}

type BranchPolicyAutoReviewersMapOutput struct{ *pulumi.OutputState }

func (BranchPolicyAutoReviewersMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BranchPolicyAutoReviewers)(nil)).Elem()
}

func (o BranchPolicyAutoReviewersMapOutput) ToBranchPolicyAutoReviewersMapOutput() BranchPolicyAutoReviewersMapOutput {
	return o
}

func (o BranchPolicyAutoReviewersMapOutput) ToBranchPolicyAutoReviewersMapOutputWithContext(ctx context.Context) BranchPolicyAutoReviewersMapOutput {
	return o
}

func (o BranchPolicyAutoReviewersMapOutput) MapIndex(k pulumi.StringInput) BranchPolicyAutoReviewersOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BranchPolicyAutoReviewers {
		return vs[0].(map[string]*BranchPolicyAutoReviewers)[vs[1].(string)]
	}).(BranchPolicyAutoReviewersOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BranchPolicyAutoReviewersInput)(nil)).Elem(), &BranchPolicyAutoReviewers{})
	pulumi.RegisterInputType(reflect.TypeOf((*BranchPolicyAutoReviewersArrayInput)(nil)).Elem(), BranchPolicyAutoReviewersArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BranchPolicyAutoReviewersMapInput)(nil)).Elem(), BranchPolicyAutoReviewersMap{})
	pulumi.RegisterOutputType(BranchPolicyAutoReviewersOutput{})
	pulumi.RegisterOutputType(BranchPolicyAutoReviewersArrayOutput{})
	pulumi.RegisterOutputType(BranchPolicyAutoReviewersMapOutput{})
}
