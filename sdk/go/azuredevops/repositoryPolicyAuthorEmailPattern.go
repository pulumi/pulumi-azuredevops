// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manage author email pattern repository policy within Azure DevOps project.
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
// 		project, err := azuredevops.NewProject(ctx, "project", &azuredevops.ProjectArgs{
// 			Description:      pulumi.String("Managed by Terraform"),
// 			Visibility:       pulumi.String("private"),
// 			VersionControl:   pulumi.String("Git"),
// 			WorkItemTemplate: pulumi.String("Agile"),
// 		})
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
// 		_, err = azuredevops.NewRepositoryPolicyAuthorEmailPattern(ctx, "repositoryPolicyAuthorEmailPattern", &azuredevops.RepositoryPolicyAuthorEmailPatternArgs{
// 			ProjectId: project.ID(),
// 			Enabled:   pulumi.Bool(true),
// 			Blocking:  pulumi.Bool(true),
// 			AuthorEmailPatterns: pulumi.StringArray{
// 				pulumi.String("user1@test.com"),
// 				pulumi.String("user2@test.com"),
// 			},
// 			RepositoryIds: pulumi.StringArray{
// 				git.ID(),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## Set project level repository policy
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
// 		_, err := azuredevops.NewRepositoryPolicyAuthorEmailPattern(ctx, "repositoryPolicyAuthorEmailPattern", &azuredevops.RepositoryPolicyAuthorEmailPatternArgs{
// 			ProjectId: pulumi.Any(azuredevops_project.P.Id),
// 			Enabled:   pulumi.Bool(true),
// 			Blocking:  pulumi.Bool(true),
// 			AuthorEmailPatterns: pulumi.StringArray{
// 				pulumi.String("user1@test.com"),
// 				pulumi.String("user2@test.com"),
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
// - [Azure DevOps Service REST API 5.1 - Policy Configurations](https://docs.microsoft.com/en-us/rest/api/azure/devops/policy/configurations/create?view=azure-devops-rest-5.1)
//
// ## Import
//
// Azure DevOps Branch Policies can be imported using the project ID and policy configuration ID
//
// ```sh
//  $ pulumi import azuredevops:index/repositoryPolicyAuthorEmailPattern:RepositoryPolicyAuthorEmailPattern p 00000000-0000-0000-0000-000000000000/0
// ```
type RepositoryPolicyAuthorEmailPattern struct {
	pulumi.CustomResourceState

	// Block pushes with a commit author email that does not match the patterns. You can specify exact emails or use wildcards.
	// Email patterns prefixed with "!" are excluded. Order is important.
	AuthorEmailPatterns pulumi.StringArrayOutput `pulumi:"authorEmailPatterns"`
	// A flag indicating if the policy should be blocking. Defaults to `true`.
	Blocking pulumi.BoolPtrOutput `pulumi:"blocking"`
	// A flag indicating if the policy should be enabled. Defaults to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// The ID of the project in which the policy will be created.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Control whether the policy is enabled for the repository or the project. If `repositoryIds` not configured, the policy will be set to the project.
	RepositoryIds pulumi.StringArrayOutput `pulumi:"repositoryIds"`
}

// NewRepositoryPolicyAuthorEmailPattern registers a new resource with the given unique name, arguments, and options.
func NewRepositoryPolicyAuthorEmailPattern(ctx *pulumi.Context,
	name string, args *RepositoryPolicyAuthorEmailPatternArgs, opts ...pulumi.ResourceOption) (*RepositoryPolicyAuthorEmailPattern, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AuthorEmailPatterns == nil {
		return nil, errors.New("invalid value for required argument 'AuthorEmailPatterns'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	var resource RepositoryPolicyAuthorEmailPattern
	err := ctx.RegisterResource("azuredevops:index/repositoryPolicyAuthorEmailPattern:RepositoryPolicyAuthorEmailPattern", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRepositoryPolicyAuthorEmailPattern gets an existing RepositoryPolicyAuthorEmailPattern resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRepositoryPolicyAuthorEmailPattern(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RepositoryPolicyAuthorEmailPatternState, opts ...pulumi.ResourceOption) (*RepositoryPolicyAuthorEmailPattern, error) {
	var resource RepositoryPolicyAuthorEmailPattern
	err := ctx.ReadResource("azuredevops:index/repositoryPolicyAuthorEmailPattern:RepositoryPolicyAuthorEmailPattern", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RepositoryPolicyAuthorEmailPattern resources.
type repositoryPolicyAuthorEmailPatternState struct {
	// Block pushes with a commit author email that does not match the patterns. You can specify exact emails or use wildcards.
	// Email patterns prefixed with "!" are excluded. Order is important.
	AuthorEmailPatterns []string `pulumi:"authorEmailPatterns"`
	// A flag indicating if the policy should be blocking. Defaults to `true`.
	Blocking *bool `pulumi:"blocking"`
	// A flag indicating if the policy should be enabled. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// The ID of the project in which the policy will be created.
	ProjectId *string `pulumi:"projectId"`
	// Control whether the policy is enabled for the repository or the project. If `repositoryIds` not configured, the policy will be set to the project.
	RepositoryIds []string `pulumi:"repositoryIds"`
}

type RepositoryPolicyAuthorEmailPatternState struct {
	// Block pushes with a commit author email that does not match the patterns. You can specify exact emails or use wildcards.
	// Email patterns prefixed with "!" are excluded. Order is important.
	AuthorEmailPatterns pulumi.StringArrayInput
	// A flag indicating if the policy should be blocking. Defaults to `true`.
	Blocking pulumi.BoolPtrInput
	// A flag indicating if the policy should be enabled. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// The ID of the project in which the policy will be created.
	ProjectId pulumi.StringPtrInput
	// Control whether the policy is enabled for the repository or the project. If `repositoryIds` not configured, the policy will be set to the project.
	RepositoryIds pulumi.StringArrayInput
}

func (RepositoryPolicyAuthorEmailPatternState) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryPolicyAuthorEmailPatternState)(nil)).Elem()
}

type repositoryPolicyAuthorEmailPatternArgs struct {
	// Block pushes with a commit author email that does not match the patterns. You can specify exact emails or use wildcards.
	// Email patterns prefixed with "!" are excluded. Order is important.
	AuthorEmailPatterns []string `pulumi:"authorEmailPatterns"`
	// A flag indicating if the policy should be blocking. Defaults to `true`.
	Blocking *bool `pulumi:"blocking"`
	// A flag indicating if the policy should be enabled. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// The ID of the project in which the policy will be created.
	ProjectId string `pulumi:"projectId"`
	// Control whether the policy is enabled for the repository or the project. If `repositoryIds` not configured, the policy will be set to the project.
	RepositoryIds []string `pulumi:"repositoryIds"`
}

// The set of arguments for constructing a RepositoryPolicyAuthorEmailPattern resource.
type RepositoryPolicyAuthorEmailPatternArgs struct {
	// Block pushes with a commit author email that does not match the patterns. You can specify exact emails or use wildcards.
	// Email patterns prefixed with "!" are excluded. Order is important.
	AuthorEmailPatterns pulumi.StringArrayInput
	// A flag indicating if the policy should be blocking. Defaults to `true`.
	Blocking pulumi.BoolPtrInput
	// A flag indicating if the policy should be enabled. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// The ID of the project in which the policy will be created.
	ProjectId pulumi.StringInput
	// Control whether the policy is enabled for the repository or the project. If `repositoryIds` not configured, the policy will be set to the project.
	RepositoryIds pulumi.StringArrayInput
}

func (RepositoryPolicyAuthorEmailPatternArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryPolicyAuthorEmailPatternArgs)(nil)).Elem()
}

type RepositoryPolicyAuthorEmailPatternInput interface {
	pulumi.Input

	ToRepositoryPolicyAuthorEmailPatternOutput() RepositoryPolicyAuthorEmailPatternOutput
	ToRepositoryPolicyAuthorEmailPatternOutputWithContext(ctx context.Context) RepositoryPolicyAuthorEmailPatternOutput
}

func (*RepositoryPolicyAuthorEmailPattern) ElementType() reflect.Type {
	return reflect.TypeOf((*RepositoryPolicyAuthorEmailPattern)(nil))
}

func (i *RepositoryPolicyAuthorEmailPattern) ToRepositoryPolicyAuthorEmailPatternOutput() RepositoryPolicyAuthorEmailPatternOutput {
	return i.ToRepositoryPolicyAuthorEmailPatternOutputWithContext(context.Background())
}

func (i *RepositoryPolicyAuthorEmailPattern) ToRepositoryPolicyAuthorEmailPatternOutputWithContext(ctx context.Context) RepositoryPolicyAuthorEmailPatternOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryPolicyAuthorEmailPatternOutput)
}

func (i *RepositoryPolicyAuthorEmailPattern) ToRepositoryPolicyAuthorEmailPatternPtrOutput() RepositoryPolicyAuthorEmailPatternPtrOutput {
	return i.ToRepositoryPolicyAuthorEmailPatternPtrOutputWithContext(context.Background())
}

func (i *RepositoryPolicyAuthorEmailPattern) ToRepositoryPolicyAuthorEmailPatternPtrOutputWithContext(ctx context.Context) RepositoryPolicyAuthorEmailPatternPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryPolicyAuthorEmailPatternPtrOutput)
}

type RepositoryPolicyAuthorEmailPatternPtrInput interface {
	pulumi.Input

	ToRepositoryPolicyAuthorEmailPatternPtrOutput() RepositoryPolicyAuthorEmailPatternPtrOutput
	ToRepositoryPolicyAuthorEmailPatternPtrOutputWithContext(ctx context.Context) RepositoryPolicyAuthorEmailPatternPtrOutput
}

type repositoryPolicyAuthorEmailPatternPtrType RepositoryPolicyAuthorEmailPatternArgs

func (*repositoryPolicyAuthorEmailPatternPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryPolicyAuthorEmailPattern)(nil))
}

func (i *repositoryPolicyAuthorEmailPatternPtrType) ToRepositoryPolicyAuthorEmailPatternPtrOutput() RepositoryPolicyAuthorEmailPatternPtrOutput {
	return i.ToRepositoryPolicyAuthorEmailPatternPtrOutputWithContext(context.Background())
}

func (i *repositoryPolicyAuthorEmailPatternPtrType) ToRepositoryPolicyAuthorEmailPatternPtrOutputWithContext(ctx context.Context) RepositoryPolicyAuthorEmailPatternPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryPolicyAuthorEmailPatternPtrOutput)
}

// RepositoryPolicyAuthorEmailPatternArrayInput is an input type that accepts RepositoryPolicyAuthorEmailPatternArray and RepositoryPolicyAuthorEmailPatternArrayOutput values.
// You can construct a concrete instance of `RepositoryPolicyAuthorEmailPatternArrayInput` via:
//
//          RepositoryPolicyAuthorEmailPatternArray{ RepositoryPolicyAuthorEmailPatternArgs{...} }
type RepositoryPolicyAuthorEmailPatternArrayInput interface {
	pulumi.Input

	ToRepositoryPolicyAuthorEmailPatternArrayOutput() RepositoryPolicyAuthorEmailPatternArrayOutput
	ToRepositoryPolicyAuthorEmailPatternArrayOutputWithContext(context.Context) RepositoryPolicyAuthorEmailPatternArrayOutput
}

type RepositoryPolicyAuthorEmailPatternArray []RepositoryPolicyAuthorEmailPatternInput

func (RepositoryPolicyAuthorEmailPatternArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*RepositoryPolicyAuthorEmailPattern)(nil))
}

func (i RepositoryPolicyAuthorEmailPatternArray) ToRepositoryPolicyAuthorEmailPatternArrayOutput() RepositoryPolicyAuthorEmailPatternArrayOutput {
	return i.ToRepositoryPolicyAuthorEmailPatternArrayOutputWithContext(context.Background())
}

func (i RepositoryPolicyAuthorEmailPatternArray) ToRepositoryPolicyAuthorEmailPatternArrayOutputWithContext(ctx context.Context) RepositoryPolicyAuthorEmailPatternArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryPolicyAuthorEmailPatternArrayOutput)
}

// RepositoryPolicyAuthorEmailPatternMapInput is an input type that accepts RepositoryPolicyAuthorEmailPatternMap and RepositoryPolicyAuthorEmailPatternMapOutput values.
// You can construct a concrete instance of `RepositoryPolicyAuthorEmailPatternMapInput` via:
//
//          RepositoryPolicyAuthorEmailPatternMap{ "key": RepositoryPolicyAuthorEmailPatternArgs{...} }
type RepositoryPolicyAuthorEmailPatternMapInput interface {
	pulumi.Input

	ToRepositoryPolicyAuthorEmailPatternMapOutput() RepositoryPolicyAuthorEmailPatternMapOutput
	ToRepositoryPolicyAuthorEmailPatternMapOutputWithContext(context.Context) RepositoryPolicyAuthorEmailPatternMapOutput
}

type RepositoryPolicyAuthorEmailPatternMap map[string]RepositoryPolicyAuthorEmailPatternInput

func (RepositoryPolicyAuthorEmailPatternMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*RepositoryPolicyAuthorEmailPattern)(nil))
}

func (i RepositoryPolicyAuthorEmailPatternMap) ToRepositoryPolicyAuthorEmailPatternMapOutput() RepositoryPolicyAuthorEmailPatternMapOutput {
	return i.ToRepositoryPolicyAuthorEmailPatternMapOutputWithContext(context.Background())
}

func (i RepositoryPolicyAuthorEmailPatternMap) ToRepositoryPolicyAuthorEmailPatternMapOutputWithContext(ctx context.Context) RepositoryPolicyAuthorEmailPatternMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryPolicyAuthorEmailPatternMapOutput)
}

type RepositoryPolicyAuthorEmailPatternOutput struct {
	*pulumi.OutputState
}

func (RepositoryPolicyAuthorEmailPatternOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RepositoryPolicyAuthorEmailPattern)(nil))
}

func (o RepositoryPolicyAuthorEmailPatternOutput) ToRepositoryPolicyAuthorEmailPatternOutput() RepositoryPolicyAuthorEmailPatternOutput {
	return o
}

func (o RepositoryPolicyAuthorEmailPatternOutput) ToRepositoryPolicyAuthorEmailPatternOutputWithContext(ctx context.Context) RepositoryPolicyAuthorEmailPatternOutput {
	return o
}

func (o RepositoryPolicyAuthorEmailPatternOutput) ToRepositoryPolicyAuthorEmailPatternPtrOutput() RepositoryPolicyAuthorEmailPatternPtrOutput {
	return o.ToRepositoryPolicyAuthorEmailPatternPtrOutputWithContext(context.Background())
}

func (o RepositoryPolicyAuthorEmailPatternOutput) ToRepositoryPolicyAuthorEmailPatternPtrOutputWithContext(ctx context.Context) RepositoryPolicyAuthorEmailPatternPtrOutput {
	return o.ApplyT(func(v RepositoryPolicyAuthorEmailPattern) *RepositoryPolicyAuthorEmailPattern {
		return &v
	}).(RepositoryPolicyAuthorEmailPatternPtrOutput)
}

type RepositoryPolicyAuthorEmailPatternPtrOutput struct {
	*pulumi.OutputState
}

func (RepositoryPolicyAuthorEmailPatternPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryPolicyAuthorEmailPattern)(nil))
}

func (o RepositoryPolicyAuthorEmailPatternPtrOutput) ToRepositoryPolicyAuthorEmailPatternPtrOutput() RepositoryPolicyAuthorEmailPatternPtrOutput {
	return o
}

func (o RepositoryPolicyAuthorEmailPatternPtrOutput) ToRepositoryPolicyAuthorEmailPatternPtrOutputWithContext(ctx context.Context) RepositoryPolicyAuthorEmailPatternPtrOutput {
	return o
}

type RepositoryPolicyAuthorEmailPatternArrayOutput struct{ *pulumi.OutputState }

func (RepositoryPolicyAuthorEmailPatternArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RepositoryPolicyAuthorEmailPattern)(nil))
}

func (o RepositoryPolicyAuthorEmailPatternArrayOutput) ToRepositoryPolicyAuthorEmailPatternArrayOutput() RepositoryPolicyAuthorEmailPatternArrayOutput {
	return o
}

func (o RepositoryPolicyAuthorEmailPatternArrayOutput) ToRepositoryPolicyAuthorEmailPatternArrayOutputWithContext(ctx context.Context) RepositoryPolicyAuthorEmailPatternArrayOutput {
	return o
}

func (o RepositoryPolicyAuthorEmailPatternArrayOutput) Index(i pulumi.IntInput) RepositoryPolicyAuthorEmailPatternOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RepositoryPolicyAuthorEmailPattern {
		return vs[0].([]RepositoryPolicyAuthorEmailPattern)[vs[1].(int)]
	}).(RepositoryPolicyAuthorEmailPatternOutput)
}

type RepositoryPolicyAuthorEmailPatternMapOutput struct{ *pulumi.OutputState }

func (RepositoryPolicyAuthorEmailPatternMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]RepositoryPolicyAuthorEmailPattern)(nil))
}

func (o RepositoryPolicyAuthorEmailPatternMapOutput) ToRepositoryPolicyAuthorEmailPatternMapOutput() RepositoryPolicyAuthorEmailPatternMapOutput {
	return o
}

func (o RepositoryPolicyAuthorEmailPatternMapOutput) ToRepositoryPolicyAuthorEmailPatternMapOutputWithContext(ctx context.Context) RepositoryPolicyAuthorEmailPatternMapOutput {
	return o
}

func (o RepositoryPolicyAuthorEmailPatternMapOutput) MapIndex(k pulumi.StringInput) RepositoryPolicyAuthorEmailPatternOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) RepositoryPolicyAuthorEmailPattern {
		return vs[0].(map[string]RepositoryPolicyAuthorEmailPattern)[vs[1].(string)]
	}).(RepositoryPolicyAuthorEmailPatternOutput)
}

func init() {
	pulumi.RegisterOutputType(RepositoryPolicyAuthorEmailPatternOutput{})
	pulumi.RegisterOutputType(RepositoryPolicyAuthorEmailPatternPtrOutput{})
	pulumi.RegisterOutputType(RepositoryPolicyAuthorEmailPatternArrayOutput{})
	pulumi.RegisterOutputType(RepositoryPolicyAuthorEmailPatternMapOutput{})
}
