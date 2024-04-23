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
//				WorkItemTemplate: pulumi.String("Agile"),
//				VersionControl:   pulumi.String("Git"),
//				Visibility:       pulumi.String("private"),
//				Description:      pulumi.String("Managed by Terraform"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuredevops.NewVariableGroup(ctx, "example", &azuredevops.VariableGroupArgs{
//				ProjectId:   example.ID(),
//				Name:        pulumi.String("Example Variable Group"),
//				Description: pulumi.String("Example Variable Group Description"),
//				AllowAccess: pulumi.Bool(true),
//				Variables: azuredevops.VariableGroupVariableArray{
//					&azuredevops.VariableGroupVariableArgs{
//						Name:  pulumi.String("key1"),
//						Value: pulumi.String("val1"),
//					},
//					&azuredevops.VariableGroupVariableArgs{
//						Name:        pulumi.String("key2"),
//						SecretValue: pulumi.String("val2"),
//						IsSecret:    pulumi.Bool(true),
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
// ### With AzureRM Key Vault
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
//				WorkItemTemplate: pulumi.String("Agile"),
//				VersionControl:   pulumi.String("Git"),
//				Visibility:       pulumi.String("private"),
//				Description:      pulumi.String("Managed by Terraform"),
//			})
//			if err != nil {
//				return err
//			}
//			exampleServiceEndpointAzureRM, err := azuredevops.NewServiceEndpointAzureRM(ctx, "example", &azuredevops.ServiceEndpointAzureRMArgs{
//				ProjectId:           example.ID(),
//				ServiceEndpointName: pulumi.String("Example AzureRM"),
//				Description:         pulumi.String("Managed by Terraform"),
//				Credentials: &azuredevops.ServiceEndpointAzureRMCredentialsArgs{
//					Serviceprincipalid:  pulumi.String("00000000-0000-0000-0000-000000000000"),
//					Serviceprincipalkey: pulumi.String("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
//				},
//				AzurermSpnTenantid:      pulumi.String("00000000-0000-0000-0000-000000000000"),
//				AzurermSubscriptionId:   pulumi.String("00000000-0000-0000-0000-000000000000"),
//				AzurermSubscriptionName: pulumi.String("Example Subscription Name"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuredevops.NewVariableGroup(ctx, "example", &azuredevops.VariableGroupArgs{
//				ProjectId:   example.ID(),
//				Name:        pulumi.String("Example Variable Group"),
//				Description: pulumi.String("Example Variable Group Description"),
//				AllowAccess: pulumi.Bool(true),
//				KeyVault: &azuredevops.VariableGroupKeyVaultArgs{
//					Name:              pulumi.String("example-kv"),
//					ServiceEndpointId: exampleServiceEndpointAzureRM.ID(),
//				},
//				Variables: azuredevops.VariableGroupVariableArray{
//					&azuredevops.VariableGroupVariableArgs{
//						Name: pulumi.String("key1"),
//					},
//					&azuredevops.VariableGroupVariableArgs{
//						Name: pulumi.String("key2"),
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
// - [Azure DevOps Service REST API 7.0 - Variable Groups](https://docs.microsoft.com/en-us/rest/api/azure/devops/distributedtask/variablegroups?view=azure-devops-rest-7.0)
// - [Azure DevOps Service REST API 7.0 - Authorized Resources](https://docs.microsoft.com/en-us/rest/api/azure/devops/build/authorizedresources?view=azure-devops-rest-7.0)
//
// ## PAT Permissions Required
//
// - **Variable Groups**: Read, Create, & Manage
// - **Build**: Read & execute
// - **Project and Team**: Read
// - **Token Administration**: Read & manage
// - **Tokens**: Read & manage
// - **Work Items**: Read
//
// ## Import
//
// **Variable groups containing secret values cannot be imported.**
//
// Azure DevOps Variable groups can be imported using the project name/variable group ID or by the project Guid/variable group ID, e.g.
//
// ```sh
// $ pulumi import azuredevops:index/variableGroup:VariableGroup example "Example Project/10"
// ```
//
// or
//
// ```sh
// $ pulumi import azuredevops:index/variableGroup:VariableGroup example 00000000-0000-0000-0000-000000000000/0
// ```
//
// _Note that for secret variables, the import command retrieve blank value in the tfstate._
type VariableGroup struct {
	pulumi.CustomResourceState

	// Boolean that indicate if this variable group is shared by all pipelines of this project.
	AllowAccess pulumi.BoolPtrOutput `pulumi:"allowAccess"`
	// The description of the Variable Group.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A list of `keyVault` blocks as documented below.
	KeyVault VariableGroupKeyVaultPtrOutput `pulumi:"keyVault"`
	// The name of the Variable Group.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project.
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
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VariableGroup
	err := ctx.RegisterResource("azuredevops:index/variableGroup:VariableGroup", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azuredevops:index/variableGroup:VariableGroup", name, id, state, &resource, opts...)
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
	Description *string `pulumi:"description"`
	// A list of `keyVault` blocks as documented below.
	KeyVault *VariableGroupKeyVault `pulumi:"keyVault"`
	// The name of the Variable Group.
	Name *string `pulumi:"name"`
	// The ID of the project.
	ProjectId *string `pulumi:"projectId"`
	// One or more `variable` blocks as documented below.
	Variables []VariableGroupVariable `pulumi:"variables"`
}

type VariableGroupState struct {
	// Boolean that indicate if this variable group is shared by all pipelines of this project.
	AllowAccess pulumi.BoolPtrInput
	// The description of the Variable Group.
	Description pulumi.StringPtrInput
	// A list of `keyVault` blocks as documented below.
	KeyVault VariableGroupKeyVaultPtrInput
	// The name of the Variable Group.
	Name pulumi.StringPtrInput
	// The ID of the project.
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
	Description *string `pulumi:"description"`
	// A list of `keyVault` blocks as documented below.
	KeyVault *VariableGroupKeyVault `pulumi:"keyVault"`
	// The name of the Variable Group.
	Name *string `pulumi:"name"`
	// The ID of the project.
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
	// A list of `keyVault` blocks as documented below.
	KeyVault VariableGroupKeyVaultPtrInput
	// The name of the Variable Group.
	Name pulumi.StringPtrInput
	// The ID of the project.
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
	return reflect.TypeOf((**VariableGroup)(nil)).Elem()
}

func (i *VariableGroup) ToVariableGroupOutput() VariableGroupOutput {
	return i.ToVariableGroupOutputWithContext(context.Background())
}

func (i *VariableGroup) ToVariableGroupOutputWithContext(ctx context.Context) VariableGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VariableGroupOutput)
}

// VariableGroupArrayInput is an input type that accepts VariableGroupArray and VariableGroupArrayOutput values.
// You can construct a concrete instance of `VariableGroupArrayInput` via:
//
//	VariableGroupArray{ VariableGroupArgs{...} }
type VariableGroupArrayInput interface {
	pulumi.Input

	ToVariableGroupArrayOutput() VariableGroupArrayOutput
	ToVariableGroupArrayOutputWithContext(context.Context) VariableGroupArrayOutput
}

type VariableGroupArray []VariableGroupInput

func (VariableGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VariableGroup)(nil)).Elem()
}

func (i VariableGroupArray) ToVariableGroupArrayOutput() VariableGroupArrayOutput {
	return i.ToVariableGroupArrayOutputWithContext(context.Background())
}

func (i VariableGroupArray) ToVariableGroupArrayOutputWithContext(ctx context.Context) VariableGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VariableGroupArrayOutput)
}

// VariableGroupMapInput is an input type that accepts VariableGroupMap and VariableGroupMapOutput values.
// You can construct a concrete instance of `VariableGroupMapInput` via:
//
//	VariableGroupMap{ "key": VariableGroupArgs{...} }
type VariableGroupMapInput interface {
	pulumi.Input

	ToVariableGroupMapOutput() VariableGroupMapOutput
	ToVariableGroupMapOutputWithContext(context.Context) VariableGroupMapOutput
}

type VariableGroupMap map[string]VariableGroupInput

func (VariableGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VariableGroup)(nil)).Elem()
}

func (i VariableGroupMap) ToVariableGroupMapOutput() VariableGroupMapOutput {
	return i.ToVariableGroupMapOutputWithContext(context.Background())
}

func (i VariableGroupMap) ToVariableGroupMapOutputWithContext(ctx context.Context) VariableGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VariableGroupMapOutput)
}

type VariableGroupOutput struct{ *pulumi.OutputState }

func (VariableGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VariableGroup)(nil)).Elem()
}

func (o VariableGroupOutput) ToVariableGroupOutput() VariableGroupOutput {
	return o
}

func (o VariableGroupOutput) ToVariableGroupOutputWithContext(ctx context.Context) VariableGroupOutput {
	return o
}

// Boolean that indicate if this variable group is shared by all pipelines of this project.
func (o VariableGroupOutput) AllowAccess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VariableGroup) pulumi.BoolPtrOutput { return v.AllowAccess }).(pulumi.BoolPtrOutput)
}

// The description of the Variable Group.
func (o VariableGroupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VariableGroup) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// A list of `keyVault` blocks as documented below.
func (o VariableGroupOutput) KeyVault() VariableGroupKeyVaultPtrOutput {
	return o.ApplyT(func(v *VariableGroup) VariableGroupKeyVaultPtrOutput { return v.KeyVault }).(VariableGroupKeyVaultPtrOutput)
}

// The name of the Variable Group.
func (o VariableGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VariableGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the project.
func (o VariableGroupOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *VariableGroup) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// One or more `variable` blocks as documented below.
func (o VariableGroupOutput) Variables() VariableGroupVariableArrayOutput {
	return o.ApplyT(func(v *VariableGroup) VariableGroupVariableArrayOutput { return v.Variables }).(VariableGroupVariableArrayOutput)
}

type VariableGroupArrayOutput struct{ *pulumi.OutputState }

func (VariableGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VariableGroup)(nil)).Elem()
}

func (o VariableGroupArrayOutput) ToVariableGroupArrayOutput() VariableGroupArrayOutput {
	return o
}

func (o VariableGroupArrayOutput) ToVariableGroupArrayOutputWithContext(ctx context.Context) VariableGroupArrayOutput {
	return o
}

func (o VariableGroupArrayOutput) Index(i pulumi.IntInput) VariableGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VariableGroup {
		return vs[0].([]*VariableGroup)[vs[1].(int)]
	}).(VariableGroupOutput)
}

type VariableGroupMapOutput struct{ *pulumi.OutputState }

func (VariableGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VariableGroup)(nil)).Elem()
}

func (o VariableGroupMapOutput) ToVariableGroupMapOutput() VariableGroupMapOutput {
	return o
}

func (o VariableGroupMapOutput) ToVariableGroupMapOutputWithContext(ctx context.Context) VariableGroupMapOutput {
	return o
}

func (o VariableGroupMapOutput) MapIndex(k pulumi.StringInput) VariableGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VariableGroup {
		return vs[0].(map[string]*VariableGroup)[vs[1].(string)]
	}).(VariableGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VariableGroupInput)(nil)).Elem(), &VariableGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*VariableGroupArrayInput)(nil)).Elem(), VariableGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VariableGroupMapInput)(nil)).Elem(), VariableGroupMap{})
	pulumi.RegisterOutputType(VariableGroupOutput{})
	pulumi.RegisterOutputType(VariableGroupArrayOutput{})
	pulumi.RegisterOutputType(VariableGroupMapOutput{})
}
