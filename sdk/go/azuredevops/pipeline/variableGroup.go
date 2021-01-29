// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package pipeline

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages variable groups within Azure DevOps.
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
// 		_, err = azuredevops.NewVariableGroup(ctx, "variablegroup", &azuredevops.VariableGroupArgs{
// 			ProjectId:   project.ID(),
// 			Description: pulumi.String("Test Variable Group Description"),
// 			AllowAccess: pulumi.Bool(true),
// 			Variables: azuredevops.VariableGroupVariableArray{
// 				&azuredevops.VariableGroupVariableArgs{
// 					Name:  pulumi.String("key"),
// 					Value: pulumi.String("value"),
// 				},
// 				&azuredevops.VariableGroupVariableArgs{
// 					Name:        pulumi.String("Account Password"),
// 					SecretValue: pulumi.String("p@ssword123"),
// 					IsSecret:    pulumi.Bool(true),
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
// - [Azure DevOps Service REST API 5.1 - Variable Groups](https://docs.microsoft.com/en-us/rest/api/azure/devops/distributedtask/variablegroups?view=azure-devops-rest-5.1)
// - [Azure DevOps Service REST API 5.1 - Authorized Resources](https://docs.microsoft.com/en-us/rest/api/azure/devops/build/authorizedresources?view=azure-devops-rest-5.1)
//
// ## PAT Permissions Required
//
// - **Variable Groups**: Read, Create, & Manage
//
// ## Import
//
// **Variable groups containing secret values cannot be imported.** Azure DevOps Variable groups can be imported using the project name/variable group ID or by the project Guid/variable group ID, e.g.
//
// ```sh
//  $ pulumi import azuredevops:Pipeline/variableGroup:VariableGroup variablegroup "Test Project/10"
// ```
//
//  or
//
// ```sh
//  $ pulumi import azuredevops:Pipeline/variableGroup:VariableGroup variablegroup 00000000-0000-0000-0000-000000000000/0
// ```
//
//  _Note that for secret variables, the import command retrieve blank value in the tfstate._
//
// Deprecated: azuredevops.pipeline.VariableGroup has been deprecated in favor of azuredevops.VariableGroup
type VariableGroup struct {
	pulumi.CustomResourceState

	// Boolean that indicate if this variable group is shared by all pipelines of this project.
	AllowAccess pulumi.BoolPtrOutput `pulumi:"allowAccess"`
	// The description of the Variable Group.
	Description pulumi.StringPtrOutput         `pulumi:"description"`
	KeyVault    VariableGroupKeyVaultPtrOutput `pulumi:"keyVault"`
	// The name of the Variable Group.
	Name pulumi.StringOutput `pulumi:"name"`
	// The project ID or project name.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// One or more `variable` blocks as documented below.
	Variables VariableGroupVariableArrayOutput `pulumi:"variables"`
}

// NewVariableGroup registers a new resource with the given unique name, arguments, and options.
func NewVariableGroup(ctx *pulumi.Context,
	name string, args *VariableGroupArgs, opts ...pulumi.ResourceOption) (*VariableGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.Variables == nil {
		return nil, errors.New("invalid value for required argument 'Variables'")
	}
	var resource VariableGroup
	err := ctx.RegisterResource("azuredevops:Pipeline/variableGroup:VariableGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVariableGroup gets an existing VariableGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVariableGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VariableGroupState, opts ...pulumi.ResourceOption) (*VariableGroup, error) {
	var resource VariableGroup
	err := ctx.ReadResource("azuredevops:Pipeline/variableGroup:VariableGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VariableGroup resources.
type variableGroupState struct {
	// Boolean that indicate if this variable group is shared by all pipelines of this project.
	AllowAccess *bool `pulumi:"allowAccess"`
	// The description of the Variable Group.
	Description *string                `pulumi:"description"`
	KeyVault    *VariableGroupKeyVault `pulumi:"keyVault"`
	// The name of the Variable Group.
	Name *string `pulumi:"name"`
	// The project ID or project name.
	ProjectId *string `pulumi:"projectId"`
	// One or more `variable` blocks as documented below.
	Variables []VariableGroupVariable `pulumi:"variables"`
}

type VariableGroupState struct {
	// Boolean that indicate if this variable group is shared by all pipelines of this project.
	AllowAccess pulumi.BoolPtrInput
	// The description of the Variable Group.
	Description pulumi.StringPtrInput
	KeyVault    VariableGroupKeyVaultPtrInput
	// The name of the Variable Group.
	Name pulumi.StringPtrInput
	// The project ID or project name.
	ProjectId pulumi.StringPtrInput
	// One or more `variable` blocks as documented below.
	Variables VariableGroupVariableArrayInput
}

func (VariableGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*variableGroupState)(nil)).Elem()
}

type variableGroupArgs struct {
	// Boolean that indicate if this variable group is shared by all pipelines of this project.
	AllowAccess *bool `pulumi:"allowAccess"`
	// The description of the Variable Group.
	Description *string                `pulumi:"description"`
	KeyVault    *VariableGroupKeyVault `pulumi:"keyVault"`
	// The name of the Variable Group.
	Name *string `pulumi:"name"`
	// The project ID or project name.
	ProjectId string `pulumi:"projectId"`
	// One or more `variable` blocks as documented below.
	Variables []VariableGroupVariable `pulumi:"variables"`
}

// The set of arguments for constructing a VariableGroup resource.
type VariableGroupArgs struct {
	// Boolean that indicate if this variable group is shared by all pipelines of this project.
	AllowAccess pulumi.BoolPtrInput
	// The description of the Variable Group.
	Description pulumi.StringPtrInput
	KeyVault    VariableGroupKeyVaultPtrInput
	// The name of the Variable Group.
	Name pulumi.StringPtrInput
	// The project ID or project name.
	ProjectId pulumi.StringInput
	// One or more `variable` blocks as documented below.
	Variables VariableGroupVariableArrayInput
}

func (VariableGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*variableGroupArgs)(nil)).Elem()
}

type VariableGroupInput interface {
	pulumi.Input

	ToVariableGroupOutput() VariableGroupOutput
	ToVariableGroupOutputWithContext(ctx context.Context) VariableGroupOutput
}

func (*VariableGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*VariableGroup)(nil))
}

func (i *VariableGroup) ToVariableGroupOutput() VariableGroupOutput {
	return i.ToVariableGroupOutputWithContext(context.Background())
}

func (i *VariableGroup) ToVariableGroupOutputWithContext(ctx context.Context) VariableGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VariableGroupOutput)
}

type VariableGroupOutput struct {
	*pulumi.OutputState
}

func (VariableGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VariableGroup)(nil))
}

func (o VariableGroupOutput) ToVariableGroupOutput() VariableGroupOutput {
	return o
}

func (o VariableGroupOutput) ToVariableGroupOutputWithContext(ctx context.Context) VariableGroupOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(VariableGroupOutput{})
}
