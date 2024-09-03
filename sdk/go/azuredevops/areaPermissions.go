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

// Manages permissions for an Area (Component)
//
// > **Note** Permissions can be assigned to group principals and not to single user principals.
//
// ## Permission levels
//
// Permission for Areas within Azure DevOps can be applied on two different levels.
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
//			example_project_readers := azuredevops.LookupGroupOutput(ctx, azuredevops.GetGroupOutputArgs{
//				ProjectId: example.ID(),
//				Name:      pulumi.String("Readers"),
//			}, nil)
//			_, err = azuredevops.NewAreaPermissions(ctx, "example-root-permissions", &azuredevops.AreaPermissionsArgs{
//				ProjectId: example.ID(),
//				Principal: pulumi.String(example_project_readers.ApplyT(func(example_project_readers azuredevops.GetGroupResult) (*string, error) {
//					return &example_project_readers.Id, nil
//				}).(pulumi.StringPtrOutput)),
//				Path: pulumi.String("/"),
//				Permissions: pulumi.StringMap{
//					"CREATE_CHILDREN": pulumi.String("Deny"),
//					"GENERIC_READ":    pulumi.String("Allow"),
//					"DELETE":          pulumi.String("Deny"),
//					"WORK_ITEM_READ":  pulumi.String("Allow"),
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
type AreaPermissions struct {
	pulumi.CustomResourceState

	// The name of the branch to assign the permissions.
	Path pulumi.StringPtrOutput `pulumi:"path"`
	// the permissions to assign. The following permissions are available.
	Permissions pulumi.StringMapOutput `pulumi:"permissions"`
	// The **group** principal to assign the permissions.
	Principal pulumi.StringOutput `pulumi:"principal"`
	// The ID of the project to assign the permissions.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Replace (`true`) or merge (`false`) the permissions. Default: `true`.
	//
	// | Permission |           Description           |
	// |------------|---------------------------------|
	// | GENERIC_   | View permissions for this       |
	// | GENERIC_   | Edit this                       |
	// | CREATE_    | Create child                    |
	// | DELETE     | Delete this                     |
	// | WORK_      | View work items in this         |
	// | WORK_      | Edit work items in this         |
	// | MANAGE_    | Manage test                     |
	// | MANAGE_    | Manage test                     |
	// | WORK_      | Edit work item comments in this |
	Replace pulumi.BoolPtrOutput `pulumi:"replace"`
}

// NewAreaPermissions registers a new resource with the given unique name, arguments, and options.
func NewAreaPermissions(ctx *pulumi.Context,
	name string, args *AreaPermissionsArgs, opts ...pulumi.ResourceOption) (*AreaPermissions, error) {
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
	var resource AreaPermissions
	err := ctx.RegisterResource("azuredevops:index/areaPermissions:AreaPermissions", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAreaPermissions gets an existing AreaPermissions resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAreaPermissions(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AreaPermissionsState, opts ...pulumi.ResourceOption) (*AreaPermissions, error) {
	var resource AreaPermissions
	err := ctx.ReadResource("azuredevops:index/areaPermissions:AreaPermissions", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AreaPermissions resources.
type areaPermissionsState struct {
	// The name of the branch to assign the permissions.
	Path *string `pulumi:"path"`
	// the permissions to assign. The following permissions are available.
	Permissions map[string]string `pulumi:"permissions"`
	// The **group** principal to assign the permissions.
	Principal *string `pulumi:"principal"`
	// The ID of the project to assign the permissions.
	ProjectId *string `pulumi:"projectId"`
	// Replace (`true`) or merge (`false`) the permissions. Default: `true`.
	//
	// | Permission |           Description           |
	// |------------|---------------------------------|
	// | GENERIC_   | View permissions for this       |
	// | GENERIC_   | Edit this                       |
	// | CREATE_    | Create child                    |
	// | DELETE     | Delete this                     |
	// | WORK_      | View work items in this         |
	// | WORK_      | Edit work items in this         |
	// | MANAGE_    | Manage test                     |
	// | MANAGE_    | Manage test                     |
	// | WORK_      | Edit work item comments in this |
	Replace *bool `pulumi:"replace"`
}

type AreaPermissionsState struct {
	// The name of the branch to assign the permissions.
	Path pulumi.StringPtrInput
	// the permissions to assign. The following permissions are available.
	Permissions pulumi.StringMapInput
	// The **group** principal to assign the permissions.
	Principal pulumi.StringPtrInput
	// The ID of the project to assign the permissions.
	ProjectId pulumi.StringPtrInput
	// Replace (`true`) or merge (`false`) the permissions. Default: `true`.
	//
	// | Permission |           Description           |
	// |------------|---------------------------------|
	// | GENERIC_   | View permissions for this       |
	// | GENERIC_   | Edit this                       |
	// | CREATE_    | Create child                    |
	// | DELETE     | Delete this                     |
	// | WORK_      | View work items in this         |
	// | WORK_      | Edit work items in this         |
	// | MANAGE_    | Manage test                     |
	// | MANAGE_    | Manage test                     |
	// | WORK_      | Edit work item comments in this |
	Replace pulumi.BoolPtrInput
}

func (AreaPermissionsState) ElementType() reflect.Type {
	return reflect.TypeOf((*areaPermissionsState)(nil)).Elem()
}

type areaPermissionsArgs struct {
	// The name of the branch to assign the permissions.
	Path *string `pulumi:"path"`
	// the permissions to assign. The following permissions are available.
	Permissions map[string]string `pulumi:"permissions"`
	// The **group** principal to assign the permissions.
	Principal string `pulumi:"principal"`
	// The ID of the project to assign the permissions.
	ProjectId string `pulumi:"projectId"`
	// Replace (`true`) or merge (`false`) the permissions. Default: `true`.
	//
	// | Permission |           Description           |
	// |------------|---------------------------------|
	// | GENERIC_   | View permissions for this       |
	// | GENERIC_   | Edit this                       |
	// | CREATE_    | Create child                    |
	// | DELETE     | Delete this                     |
	// | WORK_      | View work items in this         |
	// | WORK_      | Edit work items in this         |
	// | MANAGE_    | Manage test                     |
	// | MANAGE_    | Manage test                     |
	// | WORK_      | Edit work item comments in this |
	Replace *bool `pulumi:"replace"`
}

// The set of arguments for constructing a AreaPermissions resource.
type AreaPermissionsArgs struct {
	// The name of the branch to assign the permissions.
	Path pulumi.StringPtrInput
	// the permissions to assign. The following permissions are available.
	Permissions pulumi.StringMapInput
	// The **group** principal to assign the permissions.
	Principal pulumi.StringInput
	// The ID of the project to assign the permissions.
	ProjectId pulumi.StringInput
	// Replace (`true`) or merge (`false`) the permissions. Default: `true`.
	//
	// | Permission |           Description           |
	// |------------|---------------------------------|
	// | GENERIC_   | View permissions for this       |
	// | GENERIC_   | Edit this                       |
	// | CREATE_    | Create child                    |
	// | DELETE     | Delete this                     |
	// | WORK_      | View work items in this         |
	// | WORK_      | Edit work items in this         |
	// | MANAGE_    | Manage test                     |
	// | MANAGE_    | Manage test                     |
	// | WORK_      | Edit work item comments in this |
	Replace pulumi.BoolPtrInput
}

func (AreaPermissionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*areaPermissionsArgs)(nil)).Elem()
}

type AreaPermissionsInput interface {
	pulumi.Input

	ToAreaPermissionsOutput() AreaPermissionsOutput
	ToAreaPermissionsOutputWithContext(ctx context.Context) AreaPermissionsOutput
}

func (*AreaPermissions) ElementType() reflect.Type {
	return reflect.TypeOf((**AreaPermissions)(nil)).Elem()
}

func (i *AreaPermissions) ToAreaPermissionsOutput() AreaPermissionsOutput {
	return i.ToAreaPermissionsOutputWithContext(context.Background())
}

func (i *AreaPermissions) ToAreaPermissionsOutputWithContext(ctx context.Context) AreaPermissionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AreaPermissionsOutput)
}

// AreaPermissionsArrayInput is an input type that accepts AreaPermissionsArray and AreaPermissionsArrayOutput values.
// You can construct a concrete instance of `AreaPermissionsArrayInput` via:
//
//	AreaPermissionsArray{ AreaPermissionsArgs{...} }
type AreaPermissionsArrayInput interface {
	pulumi.Input

	ToAreaPermissionsArrayOutput() AreaPermissionsArrayOutput
	ToAreaPermissionsArrayOutputWithContext(context.Context) AreaPermissionsArrayOutput
}

type AreaPermissionsArray []AreaPermissionsInput

func (AreaPermissionsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AreaPermissions)(nil)).Elem()
}

func (i AreaPermissionsArray) ToAreaPermissionsArrayOutput() AreaPermissionsArrayOutput {
	return i.ToAreaPermissionsArrayOutputWithContext(context.Background())
}

func (i AreaPermissionsArray) ToAreaPermissionsArrayOutputWithContext(ctx context.Context) AreaPermissionsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AreaPermissionsArrayOutput)
}

// AreaPermissionsMapInput is an input type that accepts AreaPermissionsMap and AreaPermissionsMapOutput values.
// You can construct a concrete instance of `AreaPermissionsMapInput` via:
//
//	AreaPermissionsMap{ "key": AreaPermissionsArgs{...} }
type AreaPermissionsMapInput interface {
	pulumi.Input

	ToAreaPermissionsMapOutput() AreaPermissionsMapOutput
	ToAreaPermissionsMapOutputWithContext(context.Context) AreaPermissionsMapOutput
}

type AreaPermissionsMap map[string]AreaPermissionsInput

func (AreaPermissionsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AreaPermissions)(nil)).Elem()
}

func (i AreaPermissionsMap) ToAreaPermissionsMapOutput() AreaPermissionsMapOutput {
	return i.ToAreaPermissionsMapOutputWithContext(context.Background())
}

func (i AreaPermissionsMap) ToAreaPermissionsMapOutputWithContext(ctx context.Context) AreaPermissionsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AreaPermissionsMapOutput)
}

type AreaPermissionsOutput struct{ *pulumi.OutputState }

func (AreaPermissionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AreaPermissions)(nil)).Elem()
}

func (o AreaPermissionsOutput) ToAreaPermissionsOutput() AreaPermissionsOutput {
	return o
}

func (o AreaPermissionsOutput) ToAreaPermissionsOutputWithContext(ctx context.Context) AreaPermissionsOutput {
	return o
}

// The name of the branch to assign the permissions.
func (o AreaPermissionsOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AreaPermissions) pulumi.StringPtrOutput { return v.Path }).(pulumi.StringPtrOutput)
}

// the permissions to assign. The following permissions are available.
func (o AreaPermissionsOutput) Permissions() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AreaPermissions) pulumi.StringMapOutput { return v.Permissions }).(pulumi.StringMapOutput)
}

// The **group** principal to assign the permissions.
func (o AreaPermissionsOutput) Principal() pulumi.StringOutput {
	return o.ApplyT(func(v *AreaPermissions) pulumi.StringOutput { return v.Principal }).(pulumi.StringOutput)
}

// The ID of the project to assign the permissions.
func (o AreaPermissionsOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *AreaPermissions) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// Replace (`true`) or merge (`false`) the permissions. Default: `true`.
//
// | Permission |           Description           |
// |------------|---------------------------------|
// | GENERIC_   | View permissions for this       |
// | GENERIC_   | Edit this                       |
// | CREATE_    | Create child                    |
// | DELETE     | Delete this                     |
// | WORK_      | View work items in this         |
// | WORK_      | Edit work items in this         |
// | MANAGE_    | Manage test                     |
// | MANAGE_    | Manage test                     |
// | WORK_      | Edit work item comments in this |
func (o AreaPermissionsOutput) Replace() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AreaPermissions) pulumi.BoolPtrOutput { return v.Replace }).(pulumi.BoolPtrOutput)
}

type AreaPermissionsArrayOutput struct{ *pulumi.OutputState }

func (AreaPermissionsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AreaPermissions)(nil)).Elem()
}

func (o AreaPermissionsArrayOutput) ToAreaPermissionsArrayOutput() AreaPermissionsArrayOutput {
	return o
}

func (o AreaPermissionsArrayOutput) ToAreaPermissionsArrayOutputWithContext(ctx context.Context) AreaPermissionsArrayOutput {
	return o
}

func (o AreaPermissionsArrayOutput) Index(i pulumi.IntInput) AreaPermissionsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AreaPermissions {
		return vs[0].([]*AreaPermissions)[vs[1].(int)]
	}).(AreaPermissionsOutput)
}

type AreaPermissionsMapOutput struct{ *pulumi.OutputState }

func (AreaPermissionsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AreaPermissions)(nil)).Elem()
}

func (o AreaPermissionsMapOutput) ToAreaPermissionsMapOutput() AreaPermissionsMapOutput {
	return o
}

func (o AreaPermissionsMapOutput) ToAreaPermissionsMapOutputWithContext(ctx context.Context) AreaPermissionsMapOutput {
	return o
}

func (o AreaPermissionsMapOutput) MapIndex(k pulumi.StringInput) AreaPermissionsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AreaPermissions {
		return vs[0].(map[string]*AreaPermissions)[vs[1].(string)]
	}).(AreaPermissionsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AreaPermissionsInput)(nil)).Elem(), &AreaPermissions{})
	pulumi.RegisterInputType(reflect.TypeOf((*AreaPermissionsArrayInput)(nil)).Elem(), AreaPermissionsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AreaPermissionsMapInput)(nil)).Elem(), AreaPermissionsMap{})
	pulumi.RegisterOutputType(AreaPermissionsOutput{})
	pulumi.RegisterOutputType(AreaPermissionsArrayOutput{})
	pulumi.RegisterOutputType(AreaPermissionsMapOutput{})
}
