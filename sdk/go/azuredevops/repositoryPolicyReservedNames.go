// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manage a reserved names repository policy within Azure DevOps project. Block pushes that introduce files, folders, or branch names that include platform reserved names or incompatible characters.
//
// > If both project and project policy are enabled, the project policy has high priority.
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
// 		_, err = azuredevops.NewRepositoryPolicyReservedNames(ctx, "repositoryPolicyReservedNames", &azuredevops.RepositoryPolicyReservedNamesArgs{
// 			ProjectId: project.ID(),
// 			Enabled:   pulumi.Bool(true),
// 			Blocking:  pulumi.Bool(true),
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
// 		_, err := azuredevops.NewRepositoryPolicyReservedNames(ctx, "repositoryPolicyReservedNames", &azuredevops.RepositoryPolicyReservedNamesArgs{
// 			ProjectId: pulumi.Any(azuredevops_project.P.Id),
// 			Enabled:   pulumi.Bool(true),
// 			Blocking:  pulumi.Bool(true),
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
//  $ pulumi import azuredevops:index/repositoryPolicyReservedNames:RepositoryPolicyReservedNames p 00000000-0000-0000-0000-000000000000/0
// ```
type RepositoryPolicyReservedNames struct {
	pulumi.CustomResourceState

	// A flag indicating if the policy should be blocking. Defaults to `true`.
	Blocking pulumi.BoolPtrOutput `pulumi:"blocking"`
	// A flag indicating if the policy should be enabled. Defaults to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// The ID of the project in which the policy will be created.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Control whether the policy is enabled for the repository or the project. If `repositoryIds` not configured, the policy will be set to the project.
	RepositoryIds pulumi.StringArrayOutput `pulumi:"repositoryIds"`
}

// NewRepositoryPolicyReservedNames registers a new resource with the given unique name, arguments, and options.
func NewRepositoryPolicyReservedNames(ctx *pulumi.Context,
	name string, args *RepositoryPolicyReservedNamesArgs, opts ...pulumi.ResourceOption) (*RepositoryPolicyReservedNames, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	var resource RepositoryPolicyReservedNames
	err := ctx.RegisterResource("azuredevops:index/repositoryPolicyReservedNames:RepositoryPolicyReservedNames", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRepositoryPolicyReservedNames gets an existing RepositoryPolicyReservedNames resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRepositoryPolicyReservedNames(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RepositoryPolicyReservedNamesState, opts ...pulumi.ResourceOption) (*RepositoryPolicyReservedNames, error) {
	var resource RepositoryPolicyReservedNames
	err := ctx.ReadResource("azuredevops:index/repositoryPolicyReservedNames:RepositoryPolicyReservedNames", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RepositoryPolicyReservedNames resources.
type repositoryPolicyReservedNamesState struct {
	// A flag indicating if the policy should be blocking. Defaults to `true`.
	Blocking *bool `pulumi:"blocking"`
	// A flag indicating if the policy should be enabled. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// The ID of the project in which the policy will be created.
	ProjectId *string `pulumi:"projectId"`
	// Control whether the policy is enabled for the repository or the project. If `repositoryIds` not configured, the policy will be set to the project.
	RepositoryIds []string `pulumi:"repositoryIds"`
}

type RepositoryPolicyReservedNamesState struct {
	// A flag indicating if the policy should be blocking. Defaults to `true`.
	Blocking pulumi.BoolPtrInput
	// A flag indicating if the policy should be enabled. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// The ID of the project in which the policy will be created.
	ProjectId pulumi.StringPtrInput
	// Control whether the policy is enabled for the repository or the project. If `repositoryIds` not configured, the policy will be set to the project.
	RepositoryIds pulumi.StringArrayInput
}

func (RepositoryPolicyReservedNamesState) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryPolicyReservedNamesState)(nil)).Elem()
}

type repositoryPolicyReservedNamesArgs struct {
	// A flag indicating if the policy should be blocking. Defaults to `true`.
	Blocking *bool `pulumi:"blocking"`
	// A flag indicating if the policy should be enabled. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// The ID of the project in which the policy will be created.
	ProjectId string `pulumi:"projectId"`
	// Control whether the policy is enabled for the repository or the project. If `repositoryIds` not configured, the policy will be set to the project.
	RepositoryIds []string `pulumi:"repositoryIds"`
}

// The set of arguments for constructing a RepositoryPolicyReservedNames resource.
type RepositoryPolicyReservedNamesArgs struct {
	// A flag indicating if the policy should be blocking. Defaults to `true`.
	Blocking pulumi.BoolPtrInput
	// A flag indicating if the policy should be enabled. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// The ID of the project in which the policy will be created.
	ProjectId pulumi.StringInput
	// Control whether the policy is enabled for the repository or the project. If `repositoryIds` not configured, the policy will be set to the project.
	RepositoryIds pulumi.StringArrayInput
}

func (RepositoryPolicyReservedNamesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryPolicyReservedNamesArgs)(nil)).Elem()
}

type RepositoryPolicyReservedNamesInput interface {
	pulumi.Input

	ToRepositoryPolicyReservedNamesOutput() RepositoryPolicyReservedNamesOutput
	ToRepositoryPolicyReservedNamesOutputWithContext(ctx context.Context) RepositoryPolicyReservedNamesOutput
}

func (*RepositoryPolicyReservedNames) ElementType() reflect.Type {
	return reflect.TypeOf((*RepositoryPolicyReservedNames)(nil))
}

func (i *RepositoryPolicyReservedNames) ToRepositoryPolicyReservedNamesOutput() RepositoryPolicyReservedNamesOutput {
	return i.ToRepositoryPolicyReservedNamesOutputWithContext(context.Background())
}

func (i *RepositoryPolicyReservedNames) ToRepositoryPolicyReservedNamesOutputWithContext(ctx context.Context) RepositoryPolicyReservedNamesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryPolicyReservedNamesOutput)
}

func (i *RepositoryPolicyReservedNames) ToRepositoryPolicyReservedNamesPtrOutput() RepositoryPolicyReservedNamesPtrOutput {
	return i.ToRepositoryPolicyReservedNamesPtrOutputWithContext(context.Background())
}

func (i *RepositoryPolicyReservedNames) ToRepositoryPolicyReservedNamesPtrOutputWithContext(ctx context.Context) RepositoryPolicyReservedNamesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryPolicyReservedNamesPtrOutput)
}

type RepositoryPolicyReservedNamesPtrInput interface {
	pulumi.Input

	ToRepositoryPolicyReservedNamesPtrOutput() RepositoryPolicyReservedNamesPtrOutput
	ToRepositoryPolicyReservedNamesPtrOutputWithContext(ctx context.Context) RepositoryPolicyReservedNamesPtrOutput
}

type repositoryPolicyReservedNamesPtrType RepositoryPolicyReservedNamesArgs

func (*repositoryPolicyReservedNamesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryPolicyReservedNames)(nil))
}

func (i *repositoryPolicyReservedNamesPtrType) ToRepositoryPolicyReservedNamesPtrOutput() RepositoryPolicyReservedNamesPtrOutput {
	return i.ToRepositoryPolicyReservedNamesPtrOutputWithContext(context.Background())
}

func (i *repositoryPolicyReservedNamesPtrType) ToRepositoryPolicyReservedNamesPtrOutputWithContext(ctx context.Context) RepositoryPolicyReservedNamesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryPolicyReservedNamesPtrOutput)
}

// RepositoryPolicyReservedNamesArrayInput is an input type that accepts RepositoryPolicyReservedNamesArray and RepositoryPolicyReservedNamesArrayOutput values.
// You can construct a concrete instance of `RepositoryPolicyReservedNamesArrayInput` via:
//
//          RepositoryPolicyReservedNamesArray{ RepositoryPolicyReservedNamesArgs{...} }
type RepositoryPolicyReservedNamesArrayInput interface {
	pulumi.Input

	ToRepositoryPolicyReservedNamesArrayOutput() RepositoryPolicyReservedNamesArrayOutput
	ToRepositoryPolicyReservedNamesArrayOutputWithContext(context.Context) RepositoryPolicyReservedNamesArrayOutput
}

type RepositoryPolicyReservedNamesArray []RepositoryPolicyReservedNamesInput

func (RepositoryPolicyReservedNamesArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*RepositoryPolicyReservedNames)(nil))
}

func (i RepositoryPolicyReservedNamesArray) ToRepositoryPolicyReservedNamesArrayOutput() RepositoryPolicyReservedNamesArrayOutput {
	return i.ToRepositoryPolicyReservedNamesArrayOutputWithContext(context.Background())
}

func (i RepositoryPolicyReservedNamesArray) ToRepositoryPolicyReservedNamesArrayOutputWithContext(ctx context.Context) RepositoryPolicyReservedNamesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryPolicyReservedNamesArrayOutput)
}

// RepositoryPolicyReservedNamesMapInput is an input type that accepts RepositoryPolicyReservedNamesMap and RepositoryPolicyReservedNamesMapOutput values.
// You can construct a concrete instance of `RepositoryPolicyReservedNamesMapInput` via:
//
//          RepositoryPolicyReservedNamesMap{ "key": RepositoryPolicyReservedNamesArgs{...} }
type RepositoryPolicyReservedNamesMapInput interface {
	pulumi.Input

	ToRepositoryPolicyReservedNamesMapOutput() RepositoryPolicyReservedNamesMapOutput
	ToRepositoryPolicyReservedNamesMapOutputWithContext(context.Context) RepositoryPolicyReservedNamesMapOutput
}

type RepositoryPolicyReservedNamesMap map[string]RepositoryPolicyReservedNamesInput

func (RepositoryPolicyReservedNamesMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*RepositoryPolicyReservedNames)(nil))
}

func (i RepositoryPolicyReservedNamesMap) ToRepositoryPolicyReservedNamesMapOutput() RepositoryPolicyReservedNamesMapOutput {
	return i.ToRepositoryPolicyReservedNamesMapOutputWithContext(context.Background())
}

func (i RepositoryPolicyReservedNamesMap) ToRepositoryPolicyReservedNamesMapOutputWithContext(ctx context.Context) RepositoryPolicyReservedNamesMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryPolicyReservedNamesMapOutput)
}

type RepositoryPolicyReservedNamesOutput struct {
	*pulumi.OutputState
}

func (RepositoryPolicyReservedNamesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RepositoryPolicyReservedNames)(nil))
}

func (o RepositoryPolicyReservedNamesOutput) ToRepositoryPolicyReservedNamesOutput() RepositoryPolicyReservedNamesOutput {
	return o
}

func (o RepositoryPolicyReservedNamesOutput) ToRepositoryPolicyReservedNamesOutputWithContext(ctx context.Context) RepositoryPolicyReservedNamesOutput {
	return o
}

func (o RepositoryPolicyReservedNamesOutput) ToRepositoryPolicyReservedNamesPtrOutput() RepositoryPolicyReservedNamesPtrOutput {
	return o.ToRepositoryPolicyReservedNamesPtrOutputWithContext(context.Background())
}

func (o RepositoryPolicyReservedNamesOutput) ToRepositoryPolicyReservedNamesPtrOutputWithContext(ctx context.Context) RepositoryPolicyReservedNamesPtrOutput {
	return o.ApplyT(func(v RepositoryPolicyReservedNames) *RepositoryPolicyReservedNames {
		return &v
	}).(RepositoryPolicyReservedNamesPtrOutput)
}

type RepositoryPolicyReservedNamesPtrOutput struct {
	*pulumi.OutputState
}

func (RepositoryPolicyReservedNamesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryPolicyReservedNames)(nil))
}

func (o RepositoryPolicyReservedNamesPtrOutput) ToRepositoryPolicyReservedNamesPtrOutput() RepositoryPolicyReservedNamesPtrOutput {
	return o
}

func (o RepositoryPolicyReservedNamesPtrOutput) ToRepositoryPolicyReservedNamesPtrOutputWithContext(ctx context.Context) RepositoryPolicyReservedNamesPtrOutput {
	return o
}

type RepositoryPolicyReservedNamesArrayOutput struct{ *pulumi.OutputState }

func (RepositoryPolicyReservedNamesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RepositoryPolicyReservedNames)(nil))
}

func (o RepositoryPolicyReservedNamesArrayOutput) ToRepositoryPolicyReservedNamesArrayOutput() RepositoryPolicyReservedNamesArrayOutput {
	return o
}

func (o RepositoryPolicyReservedNamesArrayOutput) ToRepositoryPolicyReservedNamesArrayOutputWithContext(ctx context.Context) RepositoryPolicyReservedNamesArrayOutput {
	return o
}

func (o RepositoryPolicyReservedNamesArrayOutput) Index(i pulumi.IntInput) RepositoryPolicyReservedNamesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RepositoryPolicyReservedNames {
		return vs[0].([]RepositoryPolicyReservedNames)[vs[1].(int)]
	}).(RepositoryPolicyReservedNamesOutput)
}

type RepositoryPolicyReservedNamesMapOutput struct{ *pulumi.OutputState }

func (RepositoryPolicyReservedNamesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]RepositoryPolicyReservedNames)(nil))
}

func (o RepositoryPolicyReservedNamesMapOutput) ToRepositoryPolicyReservedNamesMapOutput() RepositoryPolicyReservedNamesMapOutput {
	return o
}

func (o RepositoryPolicyReservedNamesMapOutput) ToRepositoryPolicyReservedNamesMapOutputWithContext(ctx context.Context) RepositoryPolicyReservedNamesMapOutput {
	return o
}

func (o RepositoryPolicyReservedNamesMapOutput) MapIndex(k pulumi.StringInput) RepositoryPolicyReservedNamesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) RepositoryPolicyReservedNames {
		return vs[0].(map[string]RepositoryPolicyReservedNames)[vs[1].(string)]
	}).(RepositoryPolicyReservedNamesOutput)
}

func init() {
	pulumi.RegisterOutputType(RepositoryPolicyReservedNamesOutput{})
	pulumi.RegisterOutputType(RepositoryPolicyReservedNamesPtrOutput{})
	pulumi.RegisterOutputType(RepositoryPolicyReservedNamesArrayOutput{})
	pulumi.RegisterOutputType(RepositoryPolicyReservedNamesMapOutput{})
}
