// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manage a max file size repository policy within Azure DevOps project.
//
// > If both project and project policy are enabled, the repository policy has high priority.
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
// 		_, err = azuredevops.NewRepositoryPolicyMaxFileSize(ctx, "repositoryPolicyMaxFileSize", &azuredevops.RepositoryPolicyMaxFileSizeArgs{
// 			ProjectId:   project.ID(),
// 			Enabled:     pulumi.Bool(true),
// 			Blocking:    pulumi.Bool(true),
// 			MaxFileSize: pulumi.Int(1),
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
//
// # Set project level repository policy
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
// 		_, err := azuredevops.NewRepositoryPolicyMaxFileSize(ctx, "repositoryPolicyMaxFileSize", &azuredevops.RepositoryPolicyMaxFileSizeArgs{
// 			ProjectId:   pulumi.Any(azuredevops_project.P.Id),
// 			Enabled:     pulumi.Bool(true),
// 			Blocking:    pulumi.Bool(true),
// 			MaxFileSize: pulumi.Int(1),
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
// Azure DevOps repository policies can be imported using the projectID/policyID or projectName/policyID
//
// ```sh
//  $ pulumi import azuredevops:index/repositoryPolicyMaxFileSize:RepositoryPolicyMaxFileSize p 00000000-0000-0000-0000-000000000000/0
// ```
type RepositoryPolicyMaxFileSize struct {
	pulumi.CustomResourceState

	// A flag indicating if the policy should be blocking. Defaults to `true`.
	Blocking pulumi.BoolPtrOutput `pulumi:"blocking"`
	// A flag indicating if the policy should be enabled. Defaults to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// Block pushes that contain new or updated files larger than this limit. Available values is: `1, 2, 5, 10, 100, 200` (MB).
	MaxFileSize pulumi.IntOutput `pulumi:"maxFileSize"`
	// The ID of the project in which the policy will be created.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Control whether the policy is enabled for the repository or the project. If `repositoryIds` not configured, the policy will be set to the project.
	RepositoryIds pulumi.StringArrayOutput `pulumi:"repositoryIds"`
}

// NewRepositoryPolicyMaxFileSize registers a new resource with the given unique name, arguments, and options.
func NewRepositoryPolicyMaxFileSize(ctx *pulumi.Context,
	name string, args *RepositoryPolicyMaxFileSizeArgs, opts ...pulumi.ResourceOption) (*RepositoryPolicyMaxFileSize, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MaxFileSize == nil {
		return nil, errors.New("invalid value for required argument 'MaxFileSize'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	var resource RepositoryPolicyMaxFileSize
	err := ctx.RegisterResource("azuredevops:index/repositoryPolicyMaxFileSize:RepositoryPolicyMaxFileSize", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRepositoryPolicyMaxFileSize gets an existing RepositoryPolicyMaxFileSize resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRepositoryPolicyMaxFileSize(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RepositoryPolicyMaxFileSizeState, opts ...pulumi.ResourceOption) (*RepositoryPolicyMaxFileSize, error) {
	var resource RepositoryPolicyMaxFileSize
	err := ctx.ReadResource("azuredevops:index/repositoryPolicyMaxFileSize:RepositoryPolicyMaxFileSize", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RepositoryPolicyMaxFileSize resources.
type repositoryPolicyMaxFileSizeState struct {
	// A flag indicating if the policy should be blocking. Defaults to `true`.
	Blocking *bool `pulumi:"blocking"`
	// A flag indicating if the policy should be enabled. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// Block pushes that contain new or updated files larger than this limit. Available values is: `1, 2, 5, 10, 100, 200` (MB).
	MaxFileSize *int `pulumi:"maxFileSize"`
	// The ID of the project in which the policy will be created.
	ProjectId *string `pulumi:"projectId"`
	// Control whether the policy is enabled for the repository or the project. If `repositoryIds` not configured, the policy will be set to the project.
	RepositoryIds []string `pulumi:"repositoryIds"`
}

type RepositoryPolicyMaxFileSizeState struct {
	// A flag indicating if the policy should be blocking. Defaults to `true`.
	Blocking pulumi.BoolPtrInput
	// A flag indicating if the policy should be enabled. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// Block pushes that contain new or updated files larger than this limit. Available values is: `1, 2, 5, 10, 100, 200` (MB).
	MaxFileSize pulumi.IntPtrInput
	// The ID of the project in which the policy will be created.
	ProjectId pulumi.StringPtrInput
	// Control whether the policy is enabled for the repository or the project. If `repositoryIds` not configured, the policy will be set to the project.
	RepositoryIds pulumi.StringArrayInput
}

func (RepositoryPolicyMaxFileSizeState) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryPolicyMaxFileSizeState)(nil)).Elem()
}

type repositoryPolicyMaxFileSizeArgs struct {
	// A flag indicating if the policy should be blocking. Defaults to `true`.
	Blocking *bool `pulumi:"blocking"`
	// A flag indicating if the policy should be enabled. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// Block pushes that contain new or updated files larger than this limit. Available values is: `1, 2, 5, 10, 100, 200` (MB).
	MaxFileSize int `pulumi:"maxFileSize"`
	// The ID of the project in which the policy will be created.
	ProjectId string `pulumi:"projectId"`
	// Control whether the policy is enabled for the repository or the project. If `repositoryIds` not configured, the policy will be set to the project.
	RepositoryIds []string `pulumi:"repositoryIds"`
}

// The set of arguments for constructing a RepositoryPolicyMaxFileSize resource.
type RepositoryPolicyMaxFileSizeArgs struct {
	// A flag indicating if the policy should be blocking. Defaults to `true`.
	Blocking pulumi.BoolPtrInput
	// A flag indicating if the policy should be enabled. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// Block pushes that contain new or updated files larger than this limit. Available values is: `1, 2, 5, 10, 100, 200` (MB).
	MaxFileSize pulumi.IntInput
	// The ID of the project in which the policy will be created.
	ProjectId pulumi.StringInput
	// Control whether the policy is enabled for the repository or the project. If `repositoryIds` not configured, the policy will be set to the project.
	RepositoryIds pulumi.StringArrayInput
}

func (RepositoryPolicyMaxFileSizeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryPolicyMaxFileSizeArgs)(nil)).Elem()
}

type RepositoryPolicyMaxFileSizeInput interface {
	pulumi.Input

	ToRepositoryPolicyMaxFileSizeOutput() RepositoryPolicyMaxFileSizeOutput
	ToRepositoryPolicyMaxFileSizeOutputWithContext(ctx context.Context) RepositoryPolicyMaxFileSizeOutput
}

func (*RepositoryPolicyMaxFileSize) ElementType() reflect.Type {
	return reflect.TypeOf((*RepositoryPolicyMaxFileSize)(nil))
}

func (i *RepositoryPolicyMaxFileSize) ToRepositoryPolicyMaxFileSizeOutput() RepositoryPolicyMaxFileSizeOutput {
	return i.ToRepositoryPolicyMaxFileSizeOutputWithContext(context.Background())
}

func (i *RepositoryPolicyMaxFileSize) ToRepositoryPolicyMaxFileSizeOutputWithContext(ctx context.Context) RepositoryPolicyMaxFileSizeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryPolicyMaxFileSizeOutput)
}

func (i *RepositoryPolicyMaxFileSize) ToRepositoryPolicyMaxFileSizePtrOutput() RepositoryPolicyMaxFileSizePtrOutput {
	return i.ToRepositoryPolicyMaxFileSizePtrOutputWithContext(context.Background())
}

func (i *RepositoryPolicyMaxFileSize) ToRepositoryPolicyMaxFileSizePtrOutputWithContext(ctx context.Context) RepositoryPolicyMaxFileSizePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryPolicyMaxFileSizePtrOutput)
}

type RepositoryPolicyMaxFileSizePtrInput interface {
	pulumi.Input

	ToRepositoryPolicyMaxFileSizePtrOutput() RepositoryPolicyMaxFileSizePtrOutput
	ToRepositoryPolicyMaxFileSizePtrOutputWithContext(ctx context.Context) RepositoryPolicyMaxFileSizePtrOutput
}

type repositoryPolicyMaxFileSizePtrType RepositoryPolicyMaxFileSizeArgs

func (*repositoryPolicyMaxFileSizePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryPolicyMaxFileSize)(nil))
}

func (i *repositoryPolicyMaxFileSizePtrType) ToRepositoryPolicyMaxFileSizePtrOutput() RepositoryPolicyMaxFileSizePtrOutput {
	return i.ToRepositoryPolicyMaxFileSizePtrOutputWithContext(context.Background())
}

func (i *repositoryPolicyMaxFileSizePtrType) ToRepositoryPolicyMaxFileSizePtrOutputWithContext(ctx context.Context) RepositoryPolicyMaxFileSizePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryPolicyMaxFileSizePtrOutput)
}

// RepositoryPolicyMaxFileSizeArrayInput is an input type that accepts RepositoryPolicyMaxFileSizeArray and RepositoryPolicyMaxFileSizeArrayOutput values.
// You can construct a concrete instance of `RepositoryPolicyMaxFileSizeArrayInput` via:
//
//          RepositoryPolicyMaxFileSizeArray{ RepositoryPolicyMaxFileSizeArgs{...} }
type RepositoryPolicyMaxFileSizeArrayInput interface {
	pulumi.Input

	ToRepositoryPolicyMaxFileSizeArrayOutput() RepositoryPolicyMaxFileSizeArrayOutput
	ToRepositoryPolicyMaxFileSizeArrayOutputWithContext(context.Context) RepositoryPolicyMaxFileSizeArrayOutput
}

type RepositoryPolicyMaxFileSizeArray []RepositoryPolicyMaxFileSizeInput

func (RepositoryPolicyMaxFileSizeArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*RepositoryPolicyMaxFileSize)(nil))
}

func (i RepositoryPolicyMaxFileSizeArray) ToRepositoryPolicyMaxFileSizeArrayOutput() RepositoryPolicyMaxFileSizeArrayOutput {
	return i.ToRepositoryPolicyMaxFileSizeArrayOutputWithContext(context.Background())
}

func (i RepositoryPolicyMaxFileSizeArray) ToRepositoryPolicyMaxFileSizeArrayOutputWithContext(ctx context.Context) RepositoryPolicyMaxFileSizeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryPolicyMaxFileSizeArrayOutput)
}

// RepositoryPolicyMaxFileSizeMapInput is an input type that accepts RepositoryPolicyMaxFileSizeMap and RepositoryPolicyMaxFileSizeMapOutput values.
// You can construct a concrete instance of `RepositoryPolicyMaxFileSizeMapInput` via:
//
//          RepositoryPolicyMaxFileSizeMap{ "key": RepositoryPolicyMaxFileSizeArgs{...} }
type RepositoryPolicyMaxFileSizeMapInput interface {
	pulumi.Input

	ToRepositoryPolicyMaxFileSizeMapOutput() RepositoryPolicyMaxFileSizeMapOutput
	ToRepositoryPolicyMaxFileSizeMapOutputWithContext(context.Context) RepositoryPolicyMaxFileSizeMapOutput
}

type RepositoryPolicyMaxFileSizeMap map[string]RepositoryPolicyMaxFileSizeInput

func (RepositoryPolicyMaxFileSizeMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*RepositoryPolicyMaxFileSize)(nil))
}

func (i RepositoryPolicyMaxFileSizeMap) ToRepositoryPolicyMaxFileSizeMapOutput() RepositoryPolicyMaxFileSizeMapOutput {
	return i.ToRepositoryPolicyMaxFileSizeMapOutputWithContext(context.Background())
}

func (i RepositoryPolicyMaxFileSizeMap) ToRepositoryPolicyMaxFileSizeMapOutputWithContext(ctx context.Context) RepositoryPolicyMaxFileSizeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryPolicyMaxFileSizeMapOutput)
}

type RepositoryPolicyMaxFileSizeOutput struct {
	*pulumi.OutputState
}

func (RepositoryPolicyMaxFileSizeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RepositoryPolicyMaxFileSize)(nil))
}

func (o RepositoryPolicyMaxFileSizeOutput) ToRepositoryPolicyMaxFileSizeOutput() RepositoryPolicyMaxFileSizeOutput {
	return o
}

func (o RepositoryPolicyMaxFileSizeOutput) ToRepositoryPolicyMaxFileSizeOutputWithContext(ctx context.Context) RepositoryPolicyMaxFileSizeOutput {
	return o
}

func (o RepositoryPolicyMaxFileSizeOutput) ToRepositoryPolicyMaxFileSizePtrOutput() RepositoryPolicyMaxFileSizePtrOutput {
	return o.ToRepositoryPolicyMaxFileSizePtrOutputWithContext(context.Background())
}

func (o RepositoryPolicyMaxFileSizeOutput) ToRepositoryPolicyMaxFileSizePtrOutputWithContext(ctx context.Context) RepositoryPolicyMaxFileSizePtrOutput {
	return o.ApplyT(func(v RepositoryPolicyMaxFileSize) *RepositoryPolicyMaxFileSize {
		return &v
	}).(RepositoryPolicyMaxFileSizePtrOutput)
}

type RepositoryPolicyMaxFileSizePtrOutput struct {
	*pulumi.OutputState
}

func (RepositoryPolicyMaxFileSizePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryPolicyMaxFileSize)(nil))
}

func (o RepositoryPolicyMaxFileSizePtrOutput) ToRepositoryPolicyMaxFileSizePtrOutput() RepositoryPolicyMaxFileSizePtrOutput {
	return o
}

func (o RepositoryPolicyMaxFileSizePtrOutput) ToRepositoryPolicyMaxFileSizePtrOutputWithContext(ctx context.Context) RepositoryPolicyMaxFileSizePtrOutput {
	return o
}

type RepositoryPolicyMaxFileSizeArrayOutput struct{ *pulumi.OutputState }

func (RepositoryPolicyMaxFileSizeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RepositoryPolicyMaxFileSize)(nil))
}

func (o RepositoryPolicyMaxFileSizeArrayOutput) ToRepositoryPolicyMaxFileSizeArrayOutput() RepositoryPolicyMaxFileSizeArrayOutput {
	return o
}

func (o RepositoryPolicyMaxFileSizeArrayOutput) ToRepositoryPolicyMaxFileSizeArrayOutputWithContext(ctx context.Context) RepositoryPolicyMaxFileSizeArrayOutput {
	return o
}

func (o RepositoryPolicyMaxFileSizeArrayOutput) Index(i pulumi.IntInput) RepositoryPolicyMaxFileSizeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RepositoryPolicyMaxFileSize {
		return vs[0].([]RepositoryPolicyMaxFileSize)[vs[1].(int)]
	}).(RepositoryPolicyMaxFileSizeOutput)
}

type RepositoryPolicyMaxFileSizeMapOutput struct{ *pulumi.OutputState }

func (RepositoryPolicyMaxFileSizeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]RepositoryPolicyMaxFileSize)(nil))
}

func (o RepositoryPolicyMaxFileSizeMapOutput) ToRepositoryPolicyMaxFileSizeMapOutput() RepositoryPolicyMaxFileSizeMapOutput {
	return o
}

func (o RepositoryPolicyMaxFileSizeMapOutput) ToRepositoryPolicyMaxFileSizeMapOutputWithContext(ctx context.Context) RepositoryPolicyMaxFileSizeMapOutput {
	return o
}

func (o RepositoryPolicyMaxFileSizeMapOutput) MapIndex(k pulumi.StringInput) RepositoryPolicyMaxFileSizeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) RepositoryPolicyMaxFileSize {
		return vs[0].(map[string]RepositoryPolicyMaxFileSize)[vs[1].(string)]
	}).(RepositoryPolicyMaxFileSizeOutput)
}

func init() {
	pulumi.RegisterOutputType(RepositoryPolicyMaxFileSizeOutput{})
	pulumi.RegisterOutputType(RepositoryPolicyMaxFileSizePtrOutput{})
	pulumi.RegisterOutputType(RepositoryPolicyMaxFileSizeArrayOutput{})
	pulumi.RegisterOutputType(RepositoryPolicyMaxFileSizeMapOutput{})
}
