// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manage a credentials check repository policy within Azure DevOps project. Block pushes that introduce files, folders, or branch names that include platform reserved names or incompatible characters.
//
// > If both project and project policy are enabled, the project policy has high priority.
//
// ## Example Usage
//
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
//			exampleProject, err := azuredevops.NewProject(ctx, "exampleProject", &azuredevops.ProjectArgs{
//				Visibility:       pulumi.String("private"),
//				VersionControl:   pulumi.String("Git"),
//				WorkItemTemplate: pulumi.String("Agile"),
//				Description:      pulumi.String("Managed by Terraform"),
//			})
//			if err != nil {
//				return err
//			}
//			exampleGit, err := azuredevops.NewGit(ctx, "exampleGit", &azuredevops.GitArgs{
//				ProjectId: exampleProject.ID(),
//				Initialization: &GitInitializationArgs{
//					InitType: pulumi.String("Clean"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuredevops.NewRepositoryPolicyCheckCredentials(ctx, "exampleRepositoryPolicyCheckCredentials", &azuredevops.RepositoryPolicyCheckCredentialsArgs{
//				ProjectId: exampleProject.ID(),
//				Enabled:   pulumi.Bool(true),
//				Blocking:  pulumi.Bool(true),
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
//	"github.com/pulumi/pulumi-azuredevops/sdk/v2/go/azuredevops"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleProject, err := azuredevops.NewProject(ctx, "exampleProject", &azuredevops.ProjectArgs{
//				Visibility:       pulumi.String("private"),
//				VersionControl:   pulumi.String("Git"),
//				WorkItemTemplate: pulumi.String("Agile"),
//				Description:      pulumi.String("Managed by Terraform"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuredevops.NewRepositoryPolicyCheckCredentials(ctx, "exampleRepositoryPolicyCheckCredentials", &azuredevops.RepositoryPolicyCheckCredentialsArgs{
//				ProjectId: exampleProject.ID(),
//				Enabled:   pulumi.Bool(true),
//				Blocking:  pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ## Relevant Links
//
// - [Azure DevOps Service REST API 6.0 - Policy Configurations](https://docs.microsoft.com/en-us/rest/api/azure/devops/policy/configurations?view=azure-devops-rest-6.0)
//
// ## Import
//
// Azure DevOps repository policies can be imported using the projectID/policyID or projectName/policyID
//
// ```sh
//
//	$ pulumi import azuredevops:index/repositoryPolicyCheckCredentials:RepositoryPolicyCheckCredentials example 00000000-0000-0000-0000-000000000000/0
//
// ```
type RepositoryPolicyCheckCredentials struct {
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

// NewRepositoryPolicyCheckCredentials registers a new resource with the given unique name, arguments, and options.
func NewRepositoryPolicyCheckCredentials(ctx *pulumi.Context,
	name string, args *RepositoryPolicyCheckCredentialsArgs, opts ...pulumi.ResourceOption) (*RepositoryPolicyCheckCredentials, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	var resource RepositoryPolicyCheckCredentials
	err := ctx.RegisterResource("azuredevops:index/repositoryPolicyCheckCredentials:RepositoryPolicyCheckCredentials", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRepositoryPolicyCheckCredentials gets an existing RepositoryPolicyCheckCredentials resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRepositoryPolicyCheckCredentials(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RepositoryPolicyCheckCredentialsState, opts ...pulumi.ResourceOption) (*RepositoryPolicyCheckCredentials, error) {
	var resource RepositoryPolicyCheckCredentials
	err := ctx.ReadResource("azuredevops:index/repositoryPolicyCheckCredentials:RepositoryPolicyCheckCredentials", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RepositoryPolicyCheckCredentials resources.
type repositoryPolicyCheckCredentialsState struct {
	// A flag indicating if the policy should be blocking. Defaults to `true`.
	Blocking *bool `pulumi:"blocking"`
	// A flag indicating if the policy should be enabled. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// The ID of the project in which the policy will be created.
	ProjectId *string `pulumi:"projectId"`
	// Control whether the policy is enabled for the repository or the project. If `repositoryIds` not configured, the policy will be set to the project.
	RepositoryIds []string `pulumi:"repositoryIds"`
}

type RepositoryPolicyCheckCredentialsState struct {
	// A flag indicating if the policy should be blocking. Defaults to `true`.
	Blocking pulumi.BoolPtrInput
	// A flag indicating if the policy should be enabled. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// The ID of the project in which the policy will be created.
	ProjectId pulumi.StringPtrInput
	// Control whether the policy is enabled for the repository or the project. If `repositoryIds` not configured, the policy will be set to the project.
	RepositoryIds pulumi.StringArrayInput
}

func (RepositoryPolicyCheckCredentialsState) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryPolicyCheckCredentialsState)(nil)).Elem()
}

type repositoryPolicyCheckCredentialsArgs struct {
	// A flag indicating if the policy should be blocking. Defaults to `true`.
	Blocking *bool `pulumi:"blocking"`
	// A flag indicating if the policy should be enabled. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// The ID of the project in which the policy will be created.
	ProjectId string `pulumi:"projectId"`
	// Control whether the policy is enabled for the repository or the project. If `repositoryIds` not configured, the policy will be set to the project.
	RepositoryIds []string `pulumi:"repositoryIds"`
}

// The set of arguments for constructing a RepositoryPolicyCheckCredentials resource.
type RepositoryPolicyCheckCredentialsArgs struct {
	// A flag indicating if the policy should be blocking. Defaults to `true`.
	Blocking pulumi.BoolPtrInput
	// A flag indicating if the policy should be enabled. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// The ID of the project in which the policy will be created.
	ProjectId pulumi.StringInput
	// Control whether the policy is enabled for the repository or the project. If `repositoryIds` not configured, the policy will be set to the project.
	RepositoryIds pulumi.StringArrayInput
}

func (RepositoryPolicyCheckCredentialsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryPolicyCheckCredentialsArgs)(nil)).Elem()
}

type RepositoryPolicyCheckCredentialsInput interface {
	pulumi.Input

	ToRepositoryPolicyCheckCredentialsOutput() RepositoryPolicyCheckCredentialsOutput
	ToRepositoryPolicyCheckCredentialsOutputWithContext(ctx context.Context) RepositoryPolicyCheckCredentialsOutput
}

func (*RepositoryPolicyCheckCredentials) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryPolicyCheckCredentials)(nil)).Elem()
}

func (i *RepositoryPolicyCheckCredentials) ToRepositoryPolicyCheckCredentialsOutput() RepositoryPolicyCheckCredentialsOutput {
	return i.ToRepositoryPolicyCheckCredentialsOutputWithContext(context.Background())
}

func (i *RepositoryPolicyCheckCredentials) ToRepositoryPolicyCheckCredentialsOutputWithContext(ctx context.Context) RepositoryPolicyCheckCredentialsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryPolicyCheckCredentialsOutput)
}

// RepositoryPolicyCheckCredentialsArrayInput is an input type that accepts RepositoryPolicyCheckCredentialsArray and RepositoryPolicyCheckCredentialsArrayOutput values.
// You can construct a concrete instance of `RepositoryPolicyCheckCredentialsArrayInput` via:
//
//	RepositoryPolicyCheckCredentialsArray{ RepositoryPolicyCheckCredentialsArgs{...} }
type RepositoryPolicyCheckCredentialsArrayInput interface {
	pulumi.Input

	ToRepositoryPolicyCheckCredentialsArrayOutput() RepositoryPolicyCheckCredentialsArrayOutput
	ToRepositoryPolicyCheckCredentialsArrayOutputWithContext(context.Context) RepositoryPolicyCheckCredentialsArrayOutput
}

type RepositoryPolicyCheckCredentialsArray []RepositoryPolicyCheckCredentialsInput

func (RepositoryPolicyCheckCredentialsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RepositoryPolicyCheckCredentials)(nil)).Elem()
}

func (i RepositoryPolicyCheckCredentialsArray) ToRepositoryPolicyCheckCredentialsArrayOutput() RepositoryPolicyCheckCredentialsArrayOutput {
	return i.ToRepositoryPolicyCheckCredentialsArrayOutputWithContext(context.Background())
}

func (i RepositoryPolicyCheckCredentialsArray) ToRepositoryPolicyCheckCredentialsArrayOutputWithContext(ctx context.Context) RepositoryPolicyCheckCredentialsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryPolicyCheckCredentialsArrayOutput)
}

// RepositoryPolicyCheckCredentialsMapInput is an input type that accepts RepositoryPolicyCheckCredentialsMap and RepositoryPolicyCheckCredentialsMapOutput values.
// You can construct a concrete instance of `RepositoryPolicyCheckCredentialsMapInput` via:
//
//	RepositoryPolicyCheckCredentialsMap{ "key": RepositoryPolicyCheckCredentialsArgs{...} }
type RepositoryPolicyCheckCredentialsMapInput interface {
	pulumi.Input

	ToRepositoryPolicyCheckCredentialsMapOutput() RepositoryPolicyCheckCredentialsMapOutput
	ToRepositoryPolicyCheckCredentialsMapOutputWithContext(context.Context) RepositoryPolicyCheckCredentialsMapOutput
}

type RepositoryPolicyCheckCredentialsMap map[string]RepositoryPolicyCheckCredentialsInput

func (RepositoryPolicyCheckCredentialsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RepositoryPolicyCheckCredentials)(nil)).Elem()
}

func (i RepositoryPolicyCheckCredentialsMap) ToRepositoryPolicyCheckCredentialsMapOutput() RepositoryPolicyCheckCredentialsMapOutput {
	return i.ToRepositoryPolicyCheckCredentialsMapOutputWithContext(context.Background())
}

func (i RepositoryPolicyCheckCredentialsMap) ToRepositoryPolicyCheckCredentialsMapOutputWithContext(ctx context.Context) RepositoryPolicyCheckCredentialsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryPolicyCheckCredentialsMapOutput)
}

type RepositoryPolicyCheckCredentialsOutput struct{ *pulumi.OutputState }

func (RepositoryPolicyCheckCredentialsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryPolicyCheckCredentials)(nil)).Elem()
}

func (o RepositoryPolicyCheckCredentialsOutput) ToRepositoryPolicyCheckCredentialsOutput() RepositoryPolicyCheckCredentialsOutput {
	return o
}

func (o RepositoryPolicyCheckCredentialsOutput) ToRepositoryPolicyCheckCredentialsOutputWithContext(ctx context.Context) RepositoryPolicyCheckCredentialsOutput {
	return o
}

// A flag indicating if the policy should be blocking. Defaults to `true`.
func (o RepositoryPolicyCheckCredentialsOutput) Blocking() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RepositoryPolicyCheckCredentials) pulumi.BoolPtrOutput { return v.Blocking }).(pulumi.BoolPtrOutput)
}

// A flag indicating if the policy should be enabled. Defaults to `true`.
func (o RepositoryPolicyCheckCredentialsOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RepositoryPolicyCheckCredentials) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// The ID of the project in which the policy will be created.
func (o RepositoryPolicyCheckCredentialsOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryPolicyCheckCredentials) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// Control whether the policy is enabled for the repository or the project. If `repositoryIds` not configured, the policy will be set to the project.
func (o RepositoryPolicyCheckCredentialsOutput) RepositoryIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RepositoryPolicyCheckCredentials) pulumi.StringArrayOutput { return v.RepositoryIds }).(pulumi.StringArrayOutput)
}

type RepositoryPolicyCheckCredentialsArrayOutput struct{ *pulumi.OutputState }

func (RepositoryPolicyCheckCredentialsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RepositoryPolicyCheckCredentials)(nil)).Elem()
}

func (o RepositoryPolicyCheckCredentialsArrayOutput) ToRepositoryPolicyCheckCredentialsArrayOutput() RepositoryPolicyCheckCredentialsArrayOutput {
	return o
}

func (o RepositoryPolicyCheckCredentialsArrayOutput) ToRepositoryPolicyCheckCredentialsArrayOutputWithContext(ctx context.Context) RepositoryPolicyCheckCredentialsArrayOutput {
	return o
}

func (o RepositoryPolicyCheckCredentialsArrayOutput) Index(i pulumi.IntInput) RepositoryPolicyCheckCredentialsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RepositoryPolicyCheckCredentials {
		return vs[0].([]*RepositoryPolicyCheckCredentials)[vs[1].(int)]
	}).(RepositoryPolicyCheckCredentialsOutput)
}

type RepositoryPolicyCheckCredentialsMapOutput struct{ *pulumi.OutputState }

func (RepositoryPolicyCheckCredentialsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RepositoryPolicyCheckCredentials)(nil)).Elem()
}

func (o RepositoryPolicyCheckCredentialsMapOutput) ToRepositoryPolicyCheckCredentialsMapOutput() RepositoryPolicyCheckCredentialsMapOutput {
	return o
}

func (o RepositoryPolicyCheckCredentialsMapOutput) ToRepositoryPolicyCheckCredentialsMapOutputWithContext(ctx context.Context) RepositoryPolicyCheckCredentialsMapOutput {
	return o
}

func (o RepositoryPolicyCheckCredentialsMapOutput) MapIndex(k pulumi.StringInput) RepositoryPolicyCheckCredentialsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RepositoryPolicyCheckCredentials {
		return vs[0].(map[string]*RepositoryPolicyCheckCredentials)[vs[1].(string)]
	}).(RepositoryPolicyCheckCredentialsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryPolicyCheckCredentialsInput)(nil)).Elem(), &RepositoryPolicyCheckCredentials{})
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryPolicyCheckCredentialsArrayInput)(nil)).Elem(), RepositoryPolicyCheckCredentialsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryPolicyCheckCredentialsMapInput)(nil)).Elem(), RepositoryPolicyCheckCredentialsMap{})
	pulumi.RegisterOutputType(RepositoryPolicyCheckCredentialsOutput{})
	pulumi.RegisterOutputType(RepositoryPolicyCheckCredentialsArrayOutput{})
	pulumi.RegisterOutputType(RepositoryPolicyCheckCredentialsMapOutput{})
}
