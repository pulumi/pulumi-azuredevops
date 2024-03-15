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

// Manages permissions for tagging
//
// ## Permission levels
//
// Permissions for tagging within Azure DevOps can be applied only on Organizational and Project level.
// The project level is reflected by specifying the argument `projectId`, otherwise the permissions are set on the organizational level.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
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
//			_, err = azuredevops.NewTaggingPermissions(ctx, "example-permissions", &azuredevops.TaggingPermissionsArgs{
//				ProjectId: example.ID(),
//				Principal: example_readers.ApplyT(func(example_readers azuredevops.GetGroupResult) (*string, error) {
//					return &example_readers.Id, nil
//				}).(pulumi.StringPtrOutput),
//				Permissions: pulumi.StringMap{
//					"Enumerate": pulumi.String("allow"),
//					"Create":    pulumi.String("allow"),
//					"Update":    pulumi.String("allow"),
//					"Delete":    pulumi.String("allow"),
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
// <!--End PulumiCodeChooser -->
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
type TaggingPermissions struct {
	pulumi.CustomResourceState

	// the permissions to assign. The following permissions are available.
	Permissions pulumi.StringMapOutput `pulumi:"permissions"`
	// The **group or user** principal to assign the permissions.
	Principal pulumi.StringOutput `pulumi:"principal"`
	// The ID of the project to assign the permissions. If omitted, organization wide permissions for tagging are managed.
	ProjectId pulumi.StringPtrOutput `pulumi:"projectId"`
	// Replace (`true`) or merge (`false`) the permissions. Default: `true`
	//
	// | Name               | Permission Description     |
	// | ------------------ | -------------------------- |
	// | Enumerate          | Enumerate tag definitions  |
	// | Create             | Create tag definition      |
	// | Update             | Update tag definition      |
	// | Delete             | Delete tag definition      |
	Replace pulumi.BoolPtrOutput `pulumi:"replace"`
}

// NewTaggingPermissions registers a new resource with the given unique name, arguments, and options.
func NewTaggingPermissions(ctx *pulumi.Context,
	name string, args *TaggingPermissionsArgs, opts ...pulumi.ResourceOption) (*TaggingPermissions, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Permissions == nil {
		return nil, errors.New("invalid value for required argument 'Permissions'")
	}
	if args.Principal == nil {
		return nil, errors.New("invalid value for required argument 'Principal'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource TaggingPermissions
	err := ctx.RegisterResource("azuredevops:index/taggingPermissions:TaggingPermissions", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTaggingPermissions gets an existing TaggingPermissions resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTaggingPermissions(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TaggingPermissionsState, opts ...pulumi.ResourceOption) (*TaggingPermissions, error) {
	var resource TaggingPermissions
	err := ctx.ReadResource("azuredevops:index/taggingPermissions:TaggingPermissions", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TaggingPermissions resources.
type taggingPermissionsState struct {
	// the permissions to assign. The following permissions are available.
	Permissions map[string]string `pulumi:"permissions"`
	// The **group or user** principal to assign the permissions.
	Principal *string `pulumi:"principal"`
	// The ID of the project to assign the permissions. If omitted, organization wide permissions for tagging are managed.
	ProjectId *string `pulumi:"projectId"`
	// Replace (`true`) or merge (`false`) the permissions. Default: `true`
	//
	// | Name               | Permission Description     |
	// | ------------------ | -------------------------- |
	// | Enumerate          | Enumerate tag definitions  |
	// | Create             | Create tag definition      |
	// | Update             | Update tag definition      |
	// | Delete             | Delete tag definition      |
	Replace *bool `pulumi:"replace"`
}

type TaggingPermissionsState struct {
	// the permissions to assign. The following permissions are available.
	Permissions pulumi.StringMapInput
	// The **group or user** principal to assign the permissions.
	Principal pulumi.StringPtrInput
	// The ID of the project to assign the permissions. If omitted, organization wide permissions for tagging are managed.
	ProjectId pulumi.StringPtrInput
	// Replace (`true`) or merge (`false`) the permissions. Default: `true`
	//
	// | Name               | Permission Description     |
	// | ------------------ | -------------------------- |
	// | Enumerate          | Enumerate tag definitions  |
	// | Create             | Create tag definition      |
	// | Update             | Update tag definition      |
	// | Delete             | Delete tag definition      |
	Replace pulumi.BoolPtrInput
}

func (TaggingPermissionsState) ElementType() reflect.Type {
	return reflect.TypeOf((*taggingPermissionsState)(nil)).Elem()
}

type taggingPermissionsArgs struct {
	// the permissions to assign. The following permissions are available.
	Permissions map[string]string `pulumi:"permissions"`
	// The **group or user** principal to assign the permissions.
	Principal string `pulumi:"principal"`
	// The ID of the project to assign the permissions. If omitted, organization wide permissions for tagging are managed.
	ProjectId *string `pulumi:"projectId"`
	// Replace (`true`) or merge (`false`) the permissions. Default: `true`
	//
	// | Name               | Permission Description     |
	// | ------------------ | -------------------------- |
	// | Enumerate          | Enumerate tag definitions  |
	// | Create             | Create tag definition      |
	// | Update             | Update tag definition      |
	// | Delete             | Delete tag definition      |
	Replace *bool `pulumi:"replace"`
}

// The set of arguments for constructing a TaggingPermissions resource.
type TaggingPermissionsArgs struct {
	// the permissions to assign. The following permissions are available.
	Permissions pulumi.StringMapInput
	// The **group or user** principal to assign the permissions.
	Principal pulumi.StringInput
	// The ID of the project to assign the permissions. If omitted, organization wide permissions for tagging are managed.
	ProjectId pulumi.StringPtrInput
	// Replace (`true`) or merge (`false`) the permissions. Default: `true`
	//
	// | Name               | Permission Description     |
	// | ------------------ | -------------------------- |
	// | Enumerate          | Enumerate tag definitions  |
	// | Create             | Create tag definition      |
	// | Update             | Update tag definition      |
	// | Delete             | Delete tag definition      |
	Replace pulumi.BoolPtrInput
}

func (TaggingPermissionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*taggingPermissionsArgs)(nil)).Elem()
}

type TaggingPermissionsInput interface {
	pulumi.Input

	ToTaggingPermissionsOutput() TaggingPermissionsOutput
	ToTaggingPermissionsOutputWithContext(ctx context.Context) TaggingPermissionsOutput
}

func (*TaggingPermissions) ElementType() reflect.Type {
	return reflect.TypeOf((**TaggingPermissions)(nil)).Elem()
}

func (i *TaggingPermissions) ToTaggingPermissionsOutput() TaggingPermissionsOutput {
	return i.ToTaggingPermissionsOutputWithContext(context.Background())
}

func (i *TaggingPermissions) ToTaggingPermissionsOutputWithContext(ctx context.Context) TaggingPermissionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TaggingPermissionsOutput)
}

// TaggingPermissionsArrayInput is an input type that accepts TaggingPermissionsArray and TaggingPermissionsArrayOutput values.
// You can construct a concrete instance of `TaggingPermissionsArrayInput` via:
//
//	TaggingPermissionsArray{ TaggingPermissionsArgs{...} }
type TaggingPermissionsArrayInput interface {
	pulumi.Input

	ToTaggingPermissionsArrayOutput() TaggingPermissionsArrayOutput
	ToTaggingPermissionsArrayOutputWithContext(context.Context) TaggingPermissionsArrayOutput
}

type TaggingPermissionsArray []TaggingPermissionsInput

func (TaggingPermissionsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TaggingPermissions)(nil)).Elem()
}

func (i TaggingPermissionsArray) ToTaggingPermissionsArrayOutput() TaggingPermissionsArrayOutput {
	return i.ToTaggingPermissionsArrayOutputWithContext(context.Background())
}

func (i TaggingPermissionsArray) ToTaggingPermissionsArrayOutputWithContext(ctx context.Context) TaggingPermissionsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TaggingPermissionsArrayOutput)
}

// TaggingPermissionsMapInput is an input type that accepts TaggingPermissionsMap and TaggingPermissionsMapOutput values.
// You can construct a concrete instance of `TaggingPermissionsMapInput` via:
//
//	TaggingPermissionsMap{ "key": TaggingPermissionsArgs{...} }
type TaggingPermissionsMapInput interface {
	pulumi.Input

	ToTaggingPermissionsMapOutput() TaggingPermissionsMapOutput
	ToTaggingPermissionsMapOutputWithContext(context.Context) TaggingPermissionsMapOutput
}

type TaggingPermissionsMap map[string]TaggingPermissionsInput

func (TaggingPermissionsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TaggingPermissions)(nil)).Elem()
}

func (i TaggingPermissionsMap) ToTaggingPermissionsMapOutput() TaggingPermissionsMapOutput {
	return i.ToTaggingPermissionsMapOutputWithContext(context.Background())
}

func (i TaggingPermissionsMap) ToTaggingPermissionsMapOutputWithContext(ctx context.Context) TaggingPermissionsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TaggingPermissionsMapOutput)
}

type TaggingPermissionsOutput struct{ *pulumi.OutputState }

func (TaggingPermissionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TaggingPermissions)(nil)).Elem()
}

func (o TaggingPermissionsOutput) ToTaggingPermissionsOutput() TaggingPermissionsOutput {
	return o
}

func (o TaggingPermissionsOutput) ToTaggingPermissionsOutputWithContext(ctx context.Context) TaggingPermissionsOutput {
	return o
}

// the permissions to assign. The following permissions are available.
func (o TaggingPermissionsOutput) Permissions() pulumi.StringMapOutput {
	return o.ApplyT(func(v *TaggingPermissions) pulumi.StringMapOutput { return v.Permissions }).(pulumi.StringMapOutput)
}

// The **group or user** principal to assign the permissions.
func (o TaggingPermissionsOutput) Principal() pulumi.StringOutput {
	return o.ApplyT(func(v *TaggingPermissions) pulumi.StringOutput { return v.Principal }).(pulumi.StringOutput)
}

// The ID of the project to assign the permissions. If omitted, organization wide permissions for tagging are managed.
func (o TaggingPermissionsOutput) ProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TaggingPermissions) pulumi.StringPtrOutput { return v.ProjectId }).(pulumi.StringPtrOutput)
}

// Replace (`true`) or merge (`false`) the permissions. Default: `true`
//
// | Name               | Permission Description     |
// | ------------------ | -------------------------- |
// | Enumerate          | Enumerate tag definitions  |
// | Create             | Create tag definition      |
// | Update             | Update tag definition      |
// | Delete             | Delete tag definition      |
func (o TaggingPermissionsOutput) Replace() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *TaggingPermissions) pulumi.BoolPtrOutput { return v.Replace }).(pulumi.BoolPtrOutput)
}

type TaggingPermissionsArrayOutput struct{ *pulumi.OutputState }

func (TaggingPermissionsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TaggingPermissions)(nil)).Elem()
}

func (o TaggingPermissionsArrayOutput) ToTaggingPermissionsArrayOutput() TaggingPermissionsArrayOutput {
	return o
}

func (o TaggingPermissionsArrayOutput) ToTaggingPermissionsArrayOutputWithContext(ctx context.Context) TaggingPermissionsArrayOutput {
	return o
}

func (o TaggingPermissionsArrayOutput) Index(i pulumi.IntInput) TaggingPermissionsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TaggingPermissions {
		return vs[0].([]*TaggingPermissions)[vs[1].(int)]
	}).(TaggingPermissionsOutput)
}

type TaggingPermissionsMapOutput struct{ *pulumi.OutputState }

func (TaggingPermissionsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TaggingPermissions)(nil)).Elem()
}

func (o TaggingPermissionsMapOutput) ToTaggingPermissionsMapOutput() TaggingPermissionsMapOutput {
	return o
}

func (o TaggingPermissionsMapOutput) ToTaggingPermissionsMapOutputWithContext(ctx context.Context) TaggingPermissionsMapOutput {
	return o
}

func (o TaggingPermissionsMapOutput) MapIndex(k pulumi.StringInput) TaggingPermissionsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TaggingPermissions {
		return vs[0].(map[string]*TaggingPermissions)[vs[1].(string)]
	}).(TaggingPermissionsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TaggingPermissionsInput)(nil)).Elem(), &TaggingPermissions{})
	pulumi.RegisterInputType(reflect.TypeOf((*TaggingPermissionsArrayInput)(nil)).Elem(), TaggingPermissionsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TaggingPermissionsMapInput)(nil)).Elem(), TaggingPermissionsMap{})
	pulumi.RegisterOutputType(TaggingPermissionsOutput{})
	pulumi.RegisterOutputType(TaggingPermissionsArrayOutput{})
	pulumi.RegisterOutputType(TaggingPermissionsMapOutput{})
}
