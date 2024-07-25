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

// Manages permissions for an Iteration (Sprint)
//
// > **Note** Permissions can be assigned to group principals and not to single user principals.
//
// ## Permission levels
//
// Permission for Iterations within Azure DevOps can be applied on two different levels.
// Those levels are reflected by specifying (or omitting) values for the arguments `projectId` and `path`.
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
//				WorkItemTemplate: pulumi.String("Agile"),
//				VersionControl:   pulumi.String("Git"),
//				Visibility:       pulumi.String("private"),
//				Description:      pulumi.String("Managed by Terraform"),
//			})
//			if err != nil {
//				return err
//			}
//			example_readers := azuredevops.LookupGroupOutput(ctx, azuredevops.GetGroupOutputArgs{
//				ProjectId: example.ID(),
//				Name:      pulumi.String("Readers"),
//			}, nil)
//			_, err = azuredevops.NewIterativePermissions(ctx, "example-root-permissions", &azuredevops.IterativePermissionsArgs{
//				ProjectId: example.ID(),
//				Principal: pulumi.String(example_readers.ApplyT(func(example_readers azuredevops.GetGroupResult) (*string, error) {
//					return &example_readers.Id, nil
//				}).(pulumi.StringPtrOutput)),
//				Permissions: pulumi.StringMap{
//					"CREATE_CHILDREN": pulumi.String("Deny"),
//					"GENERIC_READ":    pulumi.String("NotSet"),
//					"DELETE":          pulumi.String("Deny"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuredevops.NewIterativePermissions(ctx, "example-iteration-permissions", &azuredevops.IterativePermissionsArgs{
//				ProjectId: example.ID(),
//				Principal: pulumi.String(example_readers.ApplyT(func(example_readers azuredevops.GetGroupResult) (*string, error) {
//					return &example_readers.Id, nil
//				}).(pulumi.StringPtrOutput)),
//				Path: pulumi.String("Iteration 1"),
//				Permissions: pulumi.StringMap{
//					"CREATE_CHILDREN": pulumi.String("Allow"),
//					"GENERIC_READ":    pulumi.String("NotSet"),
//					"DELETE":          pulumi.String("Allow"),
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
// * [Azure DevOps Service REST API 7.0 - Security](https://docs.microsoft.com/en-us/rest/api/azure/devops/security/?view=azure-devops-rest-7.0)
//
// ## PAT Permissions Required
//
// - **Project & Team**: vso.security_manage - Grants the ability to read, write, and manage security permissions.
//
// ## Import
//
// The resource does not support import.
type IterativePermissions struct {
	pulumi.CustomResourceState

	// The name of the branch to assign the permissions.
	Path pulumi.StringPtrOutput `pulumi:"path"`
	// the permissions to assign. The following permissions are available.
	Permissions pulumi.StringMapOutput `pulumi:"permissions"`
	// The **group** principal to assign the permissions.
	Principal pulumi.StringOutput `pulumi:"principal"`
	// The ID of the project to assign the permissions.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Replace (`true`) or merge (`false`) the permissions. Default: `true`
	//
	// | Permission      | Description                    |
	// |-----------------|--------------------------------|
	// | GENERIC_READ    | View permissions for this node |
	// | GENERIC_WRITE   | Edit this node                 |
	// | CREATE_CHILDREN | Create child nodes             |
	// | DELETE          | Delete this node               |
	Replace pulumi.BoolPtrOutput `pulumi:"replace"`
}

// NewIterativePermissions registers a new resource with the given unique name, arguments, and options.
func NewIterativePermissions(ctx *pulumi.Context,
	name string, args *IterativePermissionsArgs, opts ...pulumi.ResourceOption) (*IterativePermissions, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Permissions == nil {
		return nil, errors.New("invalid value for required argument 'Permissions'")
	}
	if args.Principal == nil {
		return nil, errors.New("invalid value for required argument 'Principal'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource IterativePermissions
	err := ctx.RegisterResource("azuredevops:index/iterativePermissions:IterativePermissions", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIterativePermissions gets an existing IterativePermissions resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIterativePermissions(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IterativePermissionsState, opts ...pulumi.ResourceOption) (*IterativePermissions, error) {
	var resource IterativePermissions
	err := ctx.ReadResource("azuredevops:index/iterativePermissions:IterativePermissions", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IterativePermissions resources.
type iterativePermissionsState struct {
	// The name of the branch to assign the permissions.
	Path *string `pulumi:"path"`
	// the permissions to assign. The following permissions are available.
	Permissions map[string]string `pulumi:"permissions"`
	// The **group** principal to assign the permissions.
	Principal *string `pulumi:"principal"`
	// The ID of the project to assign the permissions.
	ProjectId *string `pulumi:"projectId"`
	// Replace (`true`) or merge (`false`) the permissions. Default: `true`
	//
	// | Permission      | Description                    |
	// |-----------------|--------------------------------|
	// | GENERIC_READ    | View permissions for this node |
	// | GENERIC_WRITE   | Edit this node                 |
	// | CREATE_CHILDREN | Create child nodes             |
	// | DELETE          | Delete this node               |
	Replace *bool `pulumi:"replace"`
}

type IterativePermissionsState struct {
	// The name of the branch to assign the permissions.
	Path pulumi.StringPtrInput
	// the permissions to assign. The following permissions are available.
	Permissions pulumi.StringMapInput
	// The **group** principal to assign the permissions.
	Principal pulumi.StringPtrInput
	// The ID of the project to assign the permissions.
	ProjectId pulumi.StringPtrInput
	// Replace (`true`) or merge (`false`) the permissions. Default: `true`
	//
	// | Permission      | Description                    |
	// |-----------------|--------------------------------|
	// | GENERIC_READ    | View permissions for this node |
	// | GENERIC_WRITE   | Edit this node                 |
	// | CREATE_CHILDREN | Create child nodes             |
	// | DELETE          | Delete this node               |
	Replace pulumi.BoolPtrInput
}

func (IterativePermissionsState) ElementType() reflect.Type {
	return reflect.TypeOf((*iterativePermissionsState)(nil)).Elem()
}

type iterativePermissionsArgs struct {
	// The name of the branch to assign the permissions.
	Path *string `pulumi:"path"`
	// the permissions to assign. The following permissions are available.
	Permissions map[string]string `pulumi:"permissions"`
	// The **group** principal to assign the permissions.
	Principal string `pulumi:"principal"`
	// The ID of the project to assign the permissions.
	ProjectId string `pulumi:"projectId"`
	// Replace (`true`) or merge (`false`) the permissions. Default: `true`
	//
	// | Permission      | Description                    |
	// |-----------------|--------------------------------|
	// | GENERIC_READ    | View permissions for this node |
	// | GENERIC_WRITE   | Edit this node                 |
	// | CREATE_CHILDREN | Create child nodes             |
	// | DELETE          | Delete this node               |
	Replace *bool `pulumi:"replace"`
}

// The set of arguments for constructing a IterativePermissions resource.
type IterativePermissionsArgs struct {
	// The name of the branch to assign the permissions.
	Path pulumi.StringPtrInput
	// the permissions to assign. The following permissions are available.
	Permissions pulumi.StringMapInput
	// The **group** principal to assign the permissions.
	Principal pulumi.StringInput
	// The ID of the project to assign the permissions.
	ProjectId pulumi.StringInput
	// Replace (`true`) or merge (`false`) the permissions. Default: `true`
	//
	// | Permission      | Description                    |
	// |-----------------|--------------------------------|
	// | GENERIC_READ    | View permissions for this node |
	// | GENERIC_WRITE   | Edit this node                 |
	// | CREATE_CHILDREN | Create child nodes             |
	// | DELETE          | Delete this node               |
	Replace pulumi.BoolPtrInput
}

func (IterativePermissionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iterativePermissionsArgs)(nil)).Elem()
}

type IterativePermissionsInput interface {
	pulumi.Input

	ToIterativePermissionsOutput() IterativePermissionsOutput
	ToIterativePermissionsOutputWithContext(ctx context.Context) IterativePermissionsOutput
}

func (*IterativePermissions) ElementType() reflect.Type {
	return reflect.TypeOf((**IterativePermissions)(nil)).Elem()
}

func (i *IterativePermissions) ToIterativePermissionsOutput() IterativePermissionsOutput {
	return i.ToIterativePermissionsOutputWithContext(context.Background())
}

func (i *IterativePermissions) ToIterativePermissionsOutputWithContext(ctx context.Context) IterativePermissionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IterativePermissionsOutput)
}

// IterativePermissionsArrayInput is an input type that accepts IterativePermissionsArray and IterativePermissionsArrayOutput values.
// You can construct a concrete instance of `IterativePermissionsArrayInput` via:
//
//	IterativePermissionsArray{ IterativePermissionsArgs{...} }
type IterativePermissionsArrayInput interface {
	pulumi.Input

	ToIterativePermissionsArrayOutput() IterativePermissionsArrayOutput
	ToIterativePermissionsArrayOutputWithContext(context.Context) IterativePermissionsArrayOutput
}

type IterativePermissionsArray []IterativePermissionsInput

func (IterativePermissionsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IterativePermissions)(nil)).Elem()
}

func (i IterativePermissionsArray) ToIterativePermissionsArrayOutput() IterativePermissionsArrayOutput {
	return i.ToIterativePermissionsArrayOutputWithContext(context.Background())
}

func (i IterativePermissionsArray) ToIterativePermissionsArrayOutputWithContext(ctx context.Context) IterativePermissionsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IterativePermissionsArrayOutput)
}

// IterativePermissionsMapInput is an input type that accepts IterativePermissionsMap and IterativePermissionsMapOutput values.
// You can construct a concrete instance of `IterativePermissionsMapInput` via:
//
//	IterativePermissionsMap{ "key": IterativePermissionsArgs{...} }
type IterativePermissionsMapInput interface {
	pulumi.Input

	ToIterativePermissionsMapOutput() IterativePermissionsMapOutput
	ToIterativePermissionsMapOutputWithContext(context.Context) IterativePermissionsMapOutput
}

type IterativePermissionsMap map[string]IterativePermissionsInput

func (IterativePermissionsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IterativePermissions)(nil)).Elem()
}

func (i IterativePermissionsMap) ToIterativePermissionsMapOutput() IterativePermissionsMapOutput {
	return i.ToIterativePermissionsMapOutputWithContext(context.Background())
}

func (i IterativePermissionsMap) ToIterativePermissionsMapOutputWithContext(ctx context.Context) IterativePermissionsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IterativePermissionsMapOutput)
}

type IterativePermissionsOutput struct{ *pulumi.OutputState }

func (IterativePermissionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IterativePermissions)(nil)).Elem()
}

func (o IterativePermissionsOutput) ToIterativePermissionsOutput() IterativePermissionsOutput {
	return o
}

func (o IterativePermissionsOutput) ToIterativePermissionsOutputWithContext(ctx context.Context) IterativePermissionsOutput {
	return o
}

// The name of the branch to assign the permissions.
func (o IterativePermissionsOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IterativePermissions) pulumi.StringPtrOutput { return v.Path }).(pulumi.StringPtrOutput)
}

// the permissions to assign. The following permissions are available.
func (o IterativePermissionsOutput) Permissions() pulumi.StringMapOutput {
	return o.ApplyT(func(v *IterativePermissions) pulumi.StringMapOutput { return v.Permissions }).(pulumi.StringMapOutput)
}

// The **group** principal to assign the permissions.
func (o IterativePermissionsOutput) Principal() pulumi.StringOutput {
	return o.ApplyT(func(v *IterativePermissions) pulumi.StringOutput { return v.Principal }).(pulumi.StringOutput)
}

// The ID of the project to assign the permissions.
func (o IterativePermissionsOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *IterativePermissions) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// Replace (`true`) or merge (`false`) the permissions. Default: `true`
//
// | Permission      | Description                    |
// |-----------------|--------------------------------|
// | GENERIC_READ    | View permissions for this node |
// | GENERIC_WRITE   | Edit this node                 |
// | CREATE_CHILDREN | Create child nodes             |
// | DELETE          | Delete this node               |
func (o IterativePermissionsOutput) Replace() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *IterativePermissions) pulumi.BoolPtrOutput { return v.Replace }).(pulumi.BoolPtrOutput)
}

type IterativePermissionsArrayOutput struct{ *pulumi.OutputState }

func (IterativePermissionsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IterativePermissions)(nil)).Elem()
}

func (o IterativePermissionsArrayOutput) ToIterativePermissionsArrayOutput() IterativePermissionsArrayOutput {
	return o
}

func (o IterativePermissionsArrayOutput) ToIterativePermissionsArrayOutputWithContext(ctx context.Context) IterativePermissionsArrayOutput {
	return o
}

func (o IterativePermissionsArrayOutput) Index(i pulumi.IntInput) IterativePermissionsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IterativePermissions {
		return vs[0].([]*IterativePermissions)[vs[1].(int)]
	}).(IterativePermissionsOutput)
}

type IterativePermissionsMapOutput struct{ *pulumi.OutputState }

func (IterativePermissionsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IterativePermissions)(nil)).Elem()
}

func (o IterativePermissionsMapOutput) ToIterativePermissionsMapOutput() IterativePermissionsMapOutput {
	return o
}

func (o IterativePermissionsMapOutput) ToIterativePermissionsMapOutputWithContext(ctx context.Context) IterativePermissionsMapOutput {
	return o
}

func (o IterativePermissionsMapOutput) MapIndex(k pulumi.StringInput) IterativePermissionsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IterativePermissions {
		return vs[0].(map[string]*IterativePermissions)[vs[1].(string)]
	}).(IterativePermissionsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IterativePermissionsInput)(nil)).Elem(), &IterativePermissions{})
	pulumi.RegisterInputType(reflect.TypeOf((*IterativePermissionsArrayInput)(nil)).Elem(), IterativePermissionsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IterativePermissionsMapInput)(nil)).Elem(), IterativePermissionsMap{})
	pulumi.RegisterOutputType(IterativePermissionsOutput{})
	pulumi.RegisterOutputType(IterativePermissionsArrayOutput{})
	pulumi.RegisterOutputType(IterativePermissionsMapOutput{})
}
