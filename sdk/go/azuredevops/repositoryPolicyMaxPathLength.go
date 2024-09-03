// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azuredevops/sdk/v3/go/azuredevops/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// layout: "azuredevops"
// page_title: "AzureDevops: RepositoryPolicyMaxPathLength"
// description: |- Manages a max path length repository policy within Azure DevOps project.
// <!-- yaml: line 3: did not find expected comment or line break -->
//
// # RepositoryPolicyMaxPathLength
//
// Manage a max path length repository policy within Azure DevOps project.
//
// > If both project and project policy are enabled, the repository policy has high priority.
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
//				Description:      pulumi.String("Managed by Terraform"),
//			})
//			if err != nil {
//				return err
//			}
//			exampleGit, err := azuredevops.NewGit(ctx, "example", &azuredevops.GitArgs{
//				ProjectId: example.ID(),
//				Name:      pulumi.String("Sample Repo"),
//				Initialization: &azuredevops.GitInitializationArgs{
//					InitType: pulumi.String("Clean"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuredevops.NewRepositoryPolicyMaxPathLength(ctx, "example", &azuredevops.RepositoryPolicyMaxPathLengthArgs{
//				ProjectId:     example.ID(),
//				Enabled:       pulumi.Bool(true),
//				Blocking:      pulumi.Bool(true),
//				MaxPathLength: pulumi.Int(500),
//				RepositoryIds: pulumi.StringArray{
//					exampleGit.ID(),
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
// # Set project level repository policy
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
//				Description:      pulumi.String("Managed by Terraform"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuredevops.NewRepositoryPolicyMaxPathLength(ctx, "example", &azuredevops.RepositoryPolicyMaxPathLengthArgs{
//				ProjectId:     example.ID(),
//				Enabled:       pulumi.Bool(true),
//				Blocking:      pulumi.Bool(true),
//				MaxPathLength: pulumi.Int(1000),
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
// - [Azure DevOps Service REST API 7.0 - Policy Configurations](https://docs.microsoft.com/en-us/rest/api/azure/devops/policy/configurations?view=azure-devops-rest-7.0)
//
// ## Import
//
// Azure DevOps repository policies can be imported using the projectID/policyID or projectName/policyID:
//
// ```sh
// $ pulumi import azuredevops:index/repositoryPolicyMaxPathLength:RepositoryPolicyMaxPathLength example 00000000-0000-0000-0000-000000000000/0
// ```
type RepositoryPolicyMaxPathLength struct {
	pulumi.CustomResourceState

	// A flag indicating if the policy should be blocking. Defaults to `true`.
	Blocking pulumi.BoolPtrOutput `pulumi:"blocking"`
	// A flag indicating if the policy should be enabled. Defaults to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// Block pushes that introduce paths that exceed the specified length.
	MaxPathLength pulumi.IntOutput `pulumi:"maxPathLength"`
	// The ID of the project in which the policy will be created.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Control whether the policy is enabled for the repository or the project. If `repositoryIds` not configured, the policy will be set to the project.
	RepositoryIds pulumi.StringArrayOutput `pulumi:"repositoryIds"`
}

// NewRepositoryPolicyMaxPathLength registers a new resource with the given unique name, arguments, and options.
func NewRepositoryPolicyMaxPathLength(ctx *pulumi.Context,
	name string, args *RepositoryPolicyMaxPathLengthArgs, opts ...pulumi.ResourceOption) (*RepositoryPolicyMaxPathLength, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MaxPathLength == nil {
		return nil, errors.New("invalid value for required argument 'MaxPathLength'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RepositoryPolicyMaxPathLength
	err := ctx.RegisterResource("azuredevops:index/repositoryPolicyMaxPathLength:RepositoryPolicyMaxPathLength", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRepositoryPolicyMaxPathLength gets an existing RepositoryPolicyMaxPathLength resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRepositoryPolicyMaxPathLength(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RepositoryPolicyMaxPathLengthState, opts ...pulumi.ResourceOption) (*RepositoryPolicyMaxPathLength, error) {
	var resource RepositoryPolicyMaxPathLength
	err := ctx.ReadResource("azuredevops:index/repositoryPolicyMaxPathLength:RepositoryPolicyMaxPathLength", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RepositoryPolicyMaxPathLength resources.
type repositoryPolicyMaxPathLengthState struct {
	// A flag indicating if the policy should be blocking. Defaults to `true`.
	Blocking *bool `pulumi:"blocking"`
	// A flag indicating if the policy should be enabled. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// Block pushes that introduce paths that exceed the specified length.
	MaxPathLength *int `pulumi:"maxPathLength"`
	// The ID of the project in which the policy will be created.
	ProjectId *string `pulumi:"projectId"`
	// Control whether the policy is enabled for the repository or the project. If `repositoryIds` not configured, the policy will be set to the project.
	RepositoryIds []string `pulumi:"repositoryIds"`
}

type RepositoryPolicyMaxPathLengthState struct {
	// A flag indicating if the policy should be blocking. Defaults to `true`.
	Blocking pulumi.BoolPtrInput
	// A flag indicating if the policy should be enabled. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// Block pushes that introduce paths that exceed the specified length.
	MaxPathLength pulumi.IntPtrInput
	// The ID of the project in which the policy will be created.
	ProjectId pulumi.StringPtrInput
	// Control whether the policy is enabled for the repository or the project. If `repositoryIds` not configured, the policy will be set to the project.
	RepositoryIds pulumi.StringArrayInput
}

func (RepositoryPolicyMaxPathLengthState) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryPolicyMaxPathLengthState)(nil)).Elem()
}

type repositoryPolicyMaxPathLengthArgs struct {
	// A flag indicating if the policy should be blocking. Defaults to `true`.
	Blocking *bool `pulumi:"blocking"`
	// A flag indicating if the policy should be enabled. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// Block pushes that introduce paths that exceed the specified length.
	MaxPathLength int `pulumi:"maxPathLength"`
	// The ID of the project in which the policy will be created.
	ProjectId string `pulumi:"projectId"`
	// Control whether the policy is enabled for the repository or the project. If `repositoryIds` not configured, the policy will be set to the project.
	RepositoryIds []string `pulumi:"repositoryIds"`
}

// The set of arguments for constructing a RepositoryPolicyMaxPathLength resource.
type RepositoryPolicyMaxPathLengthArgs struct {
	// A flag indicating if the policy should be blocking. Defaults to `true`.
	Blocking pulumi.BoolPtrInput
	// A flag indicating if the policy should be enabled. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// Block pushes that introduce paths that exceed the specified length.
	MaxPathLength pulumi.IntInput
	// The ID of the project in which the policy will be created.
	ProjectId pulumi.StringInput
	// Control whether the policy is enabled for the repository or the project. If `repositoryIds` not configured, the policy will be set to the project.
	RepositoryIds pulumi.StringArrayInput
}

func (RepositoryPolicyMaxPathLengthArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryPolicyMaxPathLengthArgs)(nil)).Elem()
}

type RepositoryPolicyMaxPathLengthInput interface {
	pulumi.Input

	ToRepositoryPolicyMaxPathLengthOutput() RepositoryPolicyMaxPathLengthOutput
	ToRepositoryPolicyMaxPathLengthOutputWithContext(ctx context.Context) RepositoryPolicyMaxPathLengthOutput
}

func (*RepositoryPolicyMaxPathLength) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryPolicyMaxPathLength)(nil)).Elem()
}

func (i *RepositoryPolicyMaxPathLength) ToRepositoryPolicyMaxPathLengthOutput() RepositoryPolicyMaxPathLengthOutput {
	return i.ToRepositoryPolicyMaxPathLengthOutputWithContext(context.Background())
}

func (i *RepositoryPolicyMaxPathLength) ToRepositoryPolicyMaxPathLengthOutputWithContext(ctx context.Context) RepositoryPolicyMaxPathLengthOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryPolicyMaxPathLengthOutput)
}

// RepositoryPolicyMaxPathLengthArrayInput is an input type that accepts RepositoryPolicyMaxPathLengthArray and RepositoryPolicyMaxPathLengthArrayOutput values.
// You can construct a concrete instance of `RepositoryPolicyMaxPathLengthArrayInput` via:
//
//	RepositoryPolicyMaxPathLengthArray{ RepositoryPolicyMaxPathLengthArgs{...} }
type RepositoryPolicyMaxPathLengthArrayInput interface {
	pulumi.Input

	ToRepositoryPolicyMaxPathLengthArrayOutput() RepositoryPolicyMaxPathLengthArrayOutput
	ToRepositoryPolicyMaxPathLengthArrayOutputWithContext(context.Context) RepositoryPolicyMaxPathLengthArrayOutput
}

type RepositoryPolicyMaxPathLengthArray []RepositoryPolicyMaxPathLengthInput

func (RepositoryPolicyMaxPathLengthArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RepositoryPolicyMaxPathLength)(nil)).Elem()
}

func (i RepositoryPolicyMaxPathLengthArray) ToRepositoryPolicyMaxPathLengthArrayOutput() RepositoryPolicyMaxPathLengthArrayOutput {
	return i.ToRepositoryPolicyMaxPathLengthArrayOutputWithContext(context.Background())
}

func (i RepositoryPolicyMaxPathLengthArray) ToRepositoryPolicyMaxPathLengthArrayOutputWithContext(ctx context.Context) RepositoryPolicyMaxPathLengthArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryPolicyMaxPathLengthArrayOutput)
}

// RepositoryPolicyMaxPathLengthMapInput is an input type that accepts RepositoryPolicyMaxPathLengthMap and RepositoryPolicyMaxPathLengthMapOutput values.
// You can construct a concrete instance of `RepositoryPolicyMaxPathLengthMapInput` via:
//
//	RepositoryPolicyMaxPathLengthMap{ "key": RepositoryPolicyMaxPathLengthArgs{...} }
type RepositoryPolicyMaxPathLengthMapInput interface {
	pulumi.Input

	ToRepositoryPolicyMaxPathLengthMapOutput() RepositoryPolicyMaxPathLengthMapOutput
	ToRepositoryPolicyMaxPathLengthMapOutputWithContext(context.Context) RepositoryPolicyMaxPathLengthMapOutput
}

type RepositoryPolicyMaxPathLengthMap map[string]RepositoryPolicyMaxPathLengthInput

func (RepositoryPolicyMaxPathLengthMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RepositoryPolicyMaxPathLength)(nil)).Elem()
}

func (i RepositoryPolicyMaxPathLengthMap) ToRepositoryPolicyMaxPathLengthMapOutput() RepositoryPolicyMaxPathLengthMapOutput {
	return i.ToRepositoryPolicyMaxPathLengthMapOutputWithContext(context.Background())
}

func (i RepositoryPolicyMaxPathLengthMap) ToRepositoryPolicyMaxPathLengthMapOutputWithContext(ctx context.Context) RepositoryPolicyMaxPathLengthMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryPolicyMaxPathLengthMapOutput)
}

type RepositoryPolicyMaxPathLengthOutput struct{ *pulumi.OutputState }

func (RepositoryPolicyMaxPathLengthOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryPolicyMaxPathLength)(nil)).Elem()
}

func (o RepositoryPolicyMaxPathLengthOutput) ToRepositoryPolicyMaxPathLengthOutput() RepositoryPolicyMaxPathLengthOutput {
	return o
}

func (o RepositoryPolicyMaxPathLengthOutput) ToRepositoryPolicyMaxPathLengthOutputWithContext(ctx context.Context) RepositoryPolicyMaxPathLengthOutput {
	return o
}

// A flag indicating if the policy should be blocking. Defaults to `true`.
func (o RepositoryPolicyMaxPathLengthOutput) Blocking() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RepositoryPolicyMaxPathLength) pulumi.BoolPtrOutput { return v.Blocking }).(pulumi.BoolPtrOutput)
}

// A flag indicating if the policy should be enabled. Defaults to `true`.
func (o RepositoryPolicyMaxPathLengthOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RepositoryPolicyMaxPathLength) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// Block pushes that introduce paths that exceed the specified length.
func (o RepositoryPolicyMaxPathLengthOutput) MaxPathLength() pulumi.IntOutput {
	return o.ApplyT(func(v *RepositoryPolicyMaxPathLength) pulumi.IntOutput { return v.MaxPathLength }).(pulumi.IntOutput)
}

// The ID of the project in which the policy will be created.
func (o RepositoryPolicyMaxPathLengthOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryPolicyMaxPathLength) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// Control whether the policy is enabled for the repository or the project. If `repositoryIds` not configured, the policy will be set to the project.
func (o RepositoryPolicyMaxPathLengthOutput) RepositoryIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RepositoryPolicyMaxPathLength) pulumi.StringArrayOutput { return v.RepositoryIds }).(pulumi.StringArrayOutput)
}

type RepositoryPolicyMaxPathLengthArrayOutput struct{ *pulumi.OutputState }

func (RepositoryPolicyMaxPathLengthArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RepositoryPolicyMaxPathLength)(nil)).Elem()
}

func (o RepositoryPolicyMaxPathLengthArrayOutput) ToRepositoryPolicyMaxPathLengthArrayOutput() RepositoryPolicyMaxPathLengthArrayOutput {
	return o
}

func (o RepositoryPolicyMaxPathLengthArrayOutput) ToRepositoryPolicyMaxPathLengthArrayOutputWithContext(ctx context.Context) RepositoryPolicyMaxPathLengthArrayOutput {
	return o
}

func (o RepositoryPolicyMaxPathLengthArrayOutput) Index(i pulumi.IntInput) RepositoryPolicyMaxPathLengthOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RepositoryPolicyMaxPathLength {
		return vs[0].([]*RepositoryPolicyMaxPathLength)[vs[1].(int)]
	}).(RepositoryPolicyMaxPathLengthOutput)
}

type RepositoryPolicyMaxPathLengthMapOutput struct{ *pulumi.OutputState }

func (RepositoryPolicyMaxPathLengthMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RepositoryPolicyMaxPathLength)(nil)).Elem()
}

func (o RepositoryPolicyMaxPathLengthMapOutput) ToRepositoryPolicyMaxPathLengthMapOutput() RepositoryPolicyMaxPathLengthMapOutput {
	return o
}

func (o RepositoryPolicyMaxPathLengthMapOutput) ToRepositoryPolicyMaxPathLengthMapOutputWithContext(ctx context.Context) RepositoryPolicyMaxPathLengthMapOutput {
	return o
}

func (o RepositoryPolicyMaxPathLengthMapOutput) MapIndex(k pulumi.StringInput) RepositoryPolicyMaxPathLengthOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RepositoryPolicyMaxPathLength {
		return vs[0].(map[string]*RepositoryPolicyMaxPathLength)[vs[1].(string)]
	}).(RepositoryPolicyMaxPathLengthOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryPolicyMaxPathLengthInput)(nil)).Elem(), &RepositoryPolicyMaxPathLength{})
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryPolicyMaxPathLengthArrayInput)(nil)).Elem(), RepositoryPolicyMaxPathLengthArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryPolicyMaxPathLengthMapInput)(nil)).Elem(), RepositoryPolicyMaxPathLengthMap{})
	pulumi.RegisterOutputType(RepositoryPolicyMaxPathLengthOutput{})
	pulumi.RegisterOutputType(RepositoryPolicyMaxPathLengthArrayOutput{})
	pulumi.RegisterOutputType(RepositoryPolicyMaxPathLengthMapOutput{})
}
