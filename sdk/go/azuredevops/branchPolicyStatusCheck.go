// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azuredevops/sdk/v3/go/azuredevops/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a status check branch policy within Azure DevOps.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-azuredevops/sdk/v3/go/azuredevops"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			example, err := azuredevops.NewProject(ctx, "example", &azuredevops.ProjectArgs{
//				Name:             pulumi.String("Example Project"),
//				Visibility:       pulumi.String("private"),
//				VersionControl:   pulumi.String("Git"),
//				WorkItemTemplate: pulumi.String("Agile"),
//				Features: pulumi.StringMap{
//					"testplans": pulumi.String("disabled"),
//					"artifacts": pulumi.String("disabled"),
//				},
//				Description: pulumi.String("Managed by Pulumi"),
//			})
//			if err != nil {
//				return err
//			}
//			exampleGit, err := azuredevops.NewGit(ctx, "example", &azuredevops.GitArgs{
//				ProjectId: example.ID(),
//				Name:      pulumi.String("Example Repository"),
//				Initialization: &azuredevops.GitInitializationArgs{
//					InitType: pulumi.String("Clean"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			exampleUser, err := azuredevops.NewUser(ctx, "example", &azuredevops.UserArgs{
//				PrincipalName:      pulumi.String("mail@email.com"),
//				AccountLicenseType: pulumi.String("basic"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuredevops.NewBranchPolicyStatusCheck(ctx, "example", &azuredevops.BranchPolicyStatusCheckArgs{
//				ProjectId: example.ID(),
//				Enabled:   pulumi.Bool(true),
//				Blocking:  pulumi.Bool(true),
//				Settings: &azuredevops.BranchPolicyStatusCheckSettingsArgs{
//					Name:               pulumi.String("Release"),
//					AuthorId:           exampleUser.ID(),
//					InvalidateOnUpdate: pulumi.Bool(true),
//					Applicability:      pulumi.String("conditional"),
//					DisplayName:        pulumi.String("PreCheck"),
//					Scopes: azuredevops.BranchPolicyStatusCheckSettingsScopeArray{
//						&azuredevops.BranchPolicyStatusCheckSettingsScopeArgs{
//							RepositoryId:  exampleGit.ID(),
//							RepositoryRef: exampleGit.DefaultBranch,
//							MatchType:     pulumi.String("Exact"),
//						},
//						&azuredevops.BranchPolicyStatusCheckSettingsScopeArgs{
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
// $ pulumi import azuredevops:index/branchPolicyStatusCheck:BranchPolicyStatusCheck example 00000000-0000-0000-0000-000000000000/0
// ```
type BranchPolicyStatusCheck struct {
	pulumi.CustomResourceState

	// A flag indicating if the policy should be blocking. Defaults to `true`.
	Blocking pulumi.BoolPtrOutput `pulumi:"blocking"`
	// A flag indicating if the policy should be enabled. Defaults to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// The ID of the project in which the policy will be created.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// A `settings` block as defined below. Configuration for the policy. This block must be defined exactly once.
	Settings BranchPolicyStatusCheckSettingsOutput `pulumi:"settings"`
}

// NewBranchPolicyStatusCheck registers a new resource with the given unique name, arguments, and options.
func NewBranchPolicyStatusCheck(ctx *pulumi.Context,
	name string, args *BranchPolicyStatusCheckArgs, opts ...pulumi.ResourceOption) (*BranchPolicyStatusCheck, error) {
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
	var resource BranchPolicyStatusCheck
	err := ctx.RegisterResource("azuredevops:index/branchPolicyStatusCheck:BranchPolicyStatusCheck", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBranchPolicyStatusCheck gets an existing BranchPolicyStatusCheck resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBranchPolicyStatusCheck(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BranchPolicyStatusCheckState, opts ...pulumi.ResourceOption) (*BranchPolicyStatusCheck, error) {
	var resource BranchPolicyStatusCheck
	err := ctx.ReadResource("azuredevops:index/branchPolicyStatusCheck:BranchPolicyStatusCheck", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BranchPolicyStatusCheck resources.
type branchPolicyStatusCheckState struct {
	// A flag indicating if the policy should be blocking. Defaults to `true`.
	Blocking *bool `pulumi:"blocking"`
	// A flag indicating if the policy should be enabled. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// The ID of the project in which the policy will be created.
	ProjectId *string `pulumi:"projectId"`
	// A `settings` block as defined below. Configuration for the policy. This block must be defined exactly once.
	Settings *BranchPolicyStatusCheckSettings `pulumi:"settings"`
}

type BranchPolicyStatusCheckState struct {
	// A flag indicating if the policy should be blocking. Defaults to `true`.
	Blocking pulumi.BoolPtrInput
	// A flag indicating if the policy should be enabled. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// The ID of the project in which the policy will be created.
	ProjectId pulumi.StringPtrInput
	// A `settings` block as defined below. Configuration for the policy. This block must be defined exactly once.
	Settings BranchPolicyStatusCheckSettingsPtrInput
}

func (BranchPolicyStatusCheckState) ElementType() reflect.Type {
	return reflect.TypeOf((*branchPolicyStatusCheckState)(nil)).Elem()
}

type branchPolicyStatusCheckArgs struct {
	// A flag indicating if the policy should be blocking. Defaults to `true`.
	Blocking *bool `pulumi:"blocking"`
	// A flag indicating if the policy should be enabled. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// The ID of the project in which the policy will be created.
	ProjectId string `pulumi:"projectId"`
	// A `settings` block as defined below. Configuration for the policy. This block must be defined exactly once.
	Settings BranchPolicyStatusCheckSettings `pulumi:"settings"`
}

// The set of arguments for constructing a BranchPolicyStatusCheck resource.
type BranchPolicyStatusCheckArgs struct {
	// A flag indicating if the policy should be blocking. Defaults to `true`.
	Blocking pulumi.BoolPtrInput
	// A flag indicating if the policy should be enabled. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// The ID of the project in which the policy will be created.
	ProjectId pulumi.StringInput
	// A `settings` block as defined below. Configuration for the policy. This block must be defined exactly once.
	Settings BranchPolicyStatusCheckSettingsInput
}

func (BranchPolicyStatusCheckArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*branchPolicyStatusCheckArgs)(nil)).Elem()
}

type BranchPolicyStatusCheckInput interface {
	pulumi.Input

	ToBranchPolicyStatusCheckOutput() BranchPolicyStatusCheckOutput
	ToBranchPolicyStatusCheckOutputWithContext(ctx context.Context) BranchPolicyStatusCheckOutput
}

func (*BranchPolicyStatusCheck) ElementType() reflect.Type {
	return reflect.TypeOf((**BranchPolicyStatusCheck)(nil)).Elem()
}

func (i *BranchPolicyStatusCheck) ToBranchPolicyStatusCheckOutput() BranchPolicyStatusCheckOutput {
	return i.ToBranchPolicyStatusCheckOutputWithContext(context.Background())
}

func (i *BranchPolicyStatusCheck) ToBranchPolicyStatusCheckOutputWithContext(ctx context.Context) BranchPolicyStatusCheckOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BranchPolicyStatusCheckOutput)
}

// BranchPolicyStatusCheckArrayInput is an input type that accepts BranchPolicyStatusCheckArray and BranchPolicyStatusCheckArrayOutput values.
// You can construct a concrete instance of `BranchPolicyStatusCheckArrayInput` via:
//
//	BranchPolicyStatusCheckArray{ BranchPolicyStatusCheckArgs{...} }
type BranchPolicyStatusCheckArrayInput interface {
	pulumi.Input

	ToBranchPolicyStatusCheckArrayOutput() BranchPolicyStatusCheckArrayOutput
	ToBranchPolicyStatusCheckArrayOutputWithContext(context.Context) BranchPolicyStatusCheckArrayOutput
}

type BranchPolicyStatusCheckArray []BranchPolicyStatusCheckInput

func (BranchPolicyStatusCheckArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BranchPolicyStatusCheck)(nil)).Elem()
}

func (i BranchPolicyStatusCheckArray) ToBranchPolicyStatusCheckArrayOutput() BranchPolicyStatusCheckArrayOutput {
	return i.ToBranchPolicyStatusCheckArrayOutputWithContext(context.Background())
}

func (i BranchPolicyStatusCheckArray) ToBranchPolicyStatusCheckArrayOutputWithContext(ctx context.Context) BranchPolicyStatusCheckArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BranchPolicyStatusCheckArrayOutput)
}

// BranchPolicyStatusCheckMapInput is an input type that accepts BranchPolicyStatusCheckMap and BranchPolicyStatusCheckMapOutput values.
// You can construct a concrete instance of `BranchPolicyStatusCheckMapInput` via:
//
//	BranchPolicyStatusCheckMap{ "key": BranchPolicyStatusCheckArgs{...} }
type BranchPolicyStatusCheckMapInput interface {
	pulumi.Input

	ToBranchPolicyStatusCheckMapOutput() BranchPolicyStatusCheckMapOutput
	ToBranchPolicyStatusCheckMapOutputWithContext(context.Context) BranchPolicyStatusCheckMapOutput
}

type BranchPolicyStatusCheckMap map[string]BranchPolicyStatusCheckInput

func (BranchPolicyStatusCheckMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BranchPolicyStatusCheck)(nil)).Elem()
}

func (i BranchPolicyStatusCheckMap) ToBranchPolicyStatusCheckMapOutput() BranchPolicyStatusCheckMapOutput {
	return i.ToBranchPolicyStatusCheckMapOutputWithContext(context.Background())
}

func (i BranchPolicyStatusCheckMap) ToBranchPolicyStatusCheckMapOutputWithContext(ctx context.Context) BranchPolicyStatusCheckMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BranchPolicyStatusCheckMapOutput)
}

type BranchPolicyStatusCheckOutput struct{ *pulumi.OutputState }

func (BranchPolicyStatusCheckOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BranchPolicyStatusCheck)(nil)).Elem()
}

func (o BranchPolicyStatusCheckOutput) ToBranchPolicyStatusCheckOutput() BranchPolicyStatusCheckOutput {
	return o
}

func (o BranchPolicyStatusCheckOutput) ToBranchPolicyStatusCheckOutputWithContext(ctx context.Context) BranchPolicyStatusCheckOutput {
	return o
}

// A flag indicating if the policy should be blocking. Defaults to `true`.
func (o BranchPolicyStatusCheckOutput) Blocking() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BranchPolicyStatusCheck) pulumi.BoolPtrOutput { return v.Blocking }).(pulumi.BoolPtrOutput)
}

// A flag indicating if the policy should be enabled. Defaults to `true`.
func (o BranchPolicyStatusCheckOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BranchPolicyStatusCheck) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// The ID of the project in which the policy will be created.
func (o BranchPolicyStatusCheckOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *BranchPolicyStatusCheck) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// A `settings` block as defined below. Configuration for the policy. This block must be defined exactly once.
func (o BranchPolicyStatusCheckOutput) Settings() BranchPolicyStatusCheckSettingsOutput {
	return o.ApplyT(func(v *BranchPolicyStatusCheck) BranchPolicyStatusCheckSettingsOutput { return v.Settings }).(BranchPolicyStatusCheckSettingsOutput)
}

type BranchPolicyStatusCheckArrayOutput struct{ *pulumi.OutputState }

func (BranchPolicyStatusCheckArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BranchPolicyStatusCheck)(nil)).Elem()
}

func (o BranchPolicyStatusCheckArrayOutput) ToBranchPolicyStatusCheckArrayOutput() BranchPolicyStatusCheckArrayOutput {
	return o
}

func (o BranchPolicyStatusCheckArrayOutput) ToBranchPolicyStatusCheckArrayOutputWithContext(ctx context.Context) BranchPolicyStatusCheckArrayOutput {
	return o
}

func (o BranchPolicyStatusCheckArrayOutput) Index(i pulumi.IntInput) BranchPolicyStatusCheckOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BranchPolicyStatusCheck {
		return vs[0].([]*BranchPolicyStatusCheck)[vs[1].(int)]
	}).(BranchPolicyStatusCheckOutput)
}

type BranchPolicyStatusCheckMapOutput struct{ *pulumi.OutputState }

func (BranchPolicyStatusCheckMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BranchPolicyStatusCheck)(nil)).Elem()
}

func (o BranchPolicyStatusCheckMapOutput) ToBranchPolicyStatusCheckMapOutput() BranchPolicyStatusCheckMapOutput {
	return o
}

func (o BranchPolicyStatusCheckMapOutput) ToBranchPolicyStatusCheckMapOutputWithContext(ctx context.Context) BranchPolicyStatusCheckMapOutput {
	return o
}

func (o BranchPolicyStatusCheckMapOutput) MapIndex(k pulumi.StringInput) BranchPolicyStatusCheckOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BranchPolicyStatusCheck {
		return vs[0].(map[string]*BranchPolicyStatusCheck)[vs[1].(string)]
	}).(BranchPolicyStatusCheckOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BranchPolicyStatusCheckInput)(nil)).Elem(), &BranchPolicyStatusCheck{})
	pulumi.RegisterInputType(reflect.TypeOf((*BranchPolicyStatusCheckArrayInput)(nil)).Elem(), BranchPolicyStatusCheckArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BranchPolicyStatusCheckMapInput)(nil)).Elem(), BranchPolicyStatusCheckMap{})
	pulumi.RegisterOutputType(BranchPolicyStatusCheckOutput{})
	pulumi.RegisterOutputType(BranchPolicyStatusCheckArrayOutput{})
	pulumi.RegisterOutputType(BranchPolicyStatusCheckMapOutput{})
}
