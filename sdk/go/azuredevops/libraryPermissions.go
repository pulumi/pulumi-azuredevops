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

// Manages permissions for a Library
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
//			project, err := azuredevops.NewProject(ctx, "project", &azuredevops.ProjectArgs{
//				Name:             pulumi.String("Testing"),
//				Description:      pulumi.String("Testing-description"),
//				Visibility:       pulumi.String("private"),
//				VersionControl:   pulumi.String("Git"),
//				WorkItemTemplate: pulumi.String("Agile"),
//			})
//			if err != nil {
//				return err
//			}
//			tf_project_readers := azuredevops.LookupGroupOutput(ctx, azuredevops.GetGroupOutputArgs{
//				ProjectId: project.ID(),
//				Name:      pulumi.String("Readers"),
//			}, nil)
//			_, err = azuredevops.NewLibraryPermissions(ctx, "permissions", &azuredevops.LibraryPermissionsArgs{
//				ProjectId: project.ID(),
//				Principal: pulumi.String(tf_project_readers.ApplyT(func(tf_project_readers azuredevops.GetGroupResult) (*string, error) {
//					return &tf_project_readers.Id, nil
//				}).(pulumi.StringPtrOutput)),
//				Permissions: pulumi.StringMap{
//					"View":       pulumi.String("allow"),
//					"Administer": pulumi.String("allow"),
//					"Use":        pulumi.String("allow"),
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
// ## Roles
//
// The Azure DevOps UI uses roles to assign permissions for the Library.
//
// | Role          | Allowed Permissions    |
// | ------------- | ---------------------- |
// | Reader        | View                   |
// | Creator       | View, Create           |
// | User          | View, Use              |
// | Administrator | View, Use, Administer  |
//
// ## Relevant Links
//
// * [Azure DevOps Service REST API 6.0 - Security](https://docs.microsoft.com/en-us/rest/api/azure/devops/security/?view=azure-devops-rest-6.0)
//
// ## PAT Permissions Required
//
// - **Project & Team**: vso.security_manage - Grants the ability to read, write, and manage security permissions.
//
// ## Import
//
// The resource does not support import.
type LibraryPermissions struct {
	pulumi.CustomResourceState

	// the permissions to assign. The following permissions are available.
	Permissions pulumi.StringMapOutput `pulumi:"permissions"`
	// The **group** principal to assign the permissions.
	Principal pulumi.StringOutput `pulumi:"principal"`
	// The ID of the project.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Replace (`true`) or merge (`false`) the permissions. Default: `true`
	//
	// | Permission        | Description                         |
	// | ----------------- | ----------------------------------- |
	// | View              | View library item                   |
	// | Administer        | Administer library item             |
	// | Create            | Create library item                 |
	// | ViewSecrets       | View library item secrets           |
	// | Use               | Use library item                    |
	// | Owner             | Owner library item                  |
	Replace pulumi.BoolPtrOutput `pulumi:"replace"`
}

// NewLibraryPermissions registers a new resource with the given unique name, arguments, and options.
func NewLibraryPermissions(ctx *pulumi.Context,
	name string, args *LibraryPermissionsArgs, opts ...pulumi.ResourceOption) (*LibraryPermissions, error) {
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
	var resource LibraryPermissions
	err := ctx.RegisterResource("azuredevops:index/libraryPermissions:LibraryPermissions", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLibraryPermissions gets an existing LibraryPermissions resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLibraryPermissions(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LibraryPermissionsState, opts ...pulumi.ResourceOption) (*LibraryPermissions, error) {
	var resource LibraryPermissions
	err := ctx.ReadResource("azuredevops:index/libraryPermissions:LibraryPermissions", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LibraryPermissions resources.
type libraryPermissionsState struct {
	// the permissions to assign. The following permissions are available.
	Permissions map[string]string `pulumi:"permissions"`
	// The **group** principal to assign the permissions.
	Principal *string `pulumi:"principal"`
	// The ID of the project.
	ProjectId *string `pulumi:"projectId"`
	// Replace (`true`) or merge (`false`) the permissions. Default: `true`
	//
	// | Permission        | Description                         |
	// | ----------------- | ----------------------------------- |
	// | View              | View library item                   |
	// | Administer        | Administer library item             |
	// | Create            | Create library item                 |
	// | ViewSecrets       | View library item secrets           |
	// | Use               | Use library item                    |
	// | Owner             | Owner library item                  |
	Replace *bool `pulumi:"replace"`
}

type LibraryPermissionsState struct {
	// the permissions to assign. The following permissions are available.
	Permissions pulumi.StringMapInput
	// The **group** principal to assign the permissions.
	Principal pulumi.StringPtrInput
	// The ID of the project.
	ProjectId pulumi.StringPtrInput
	// Replace (`true`) or merge (`false`) the permissions. Default: `true`
	//
	// | Permission        | Description                         |
	// | ----------------- | ----------------------------------- |
	// | View              | View library item                   |
	// | Administer        | Administer library item             |
	// | Create            | Create library item                 |
	// | ViewSecrets       | View library item secrets           |
	// | Use               | Use library item                    |
	// | Owner             | Owner library item                  |
	Replace pulumi.BoolPtrInput
}

func (LibraryPermissionsState) ElementType() reflect.Type {
	return reflect.TypeOf((*libraryPermissionsState)(nil)).Elem()
}

type libraryPermissionsArgs struct {
	// the permissions to assign. The following permissions are available.
	Permissions map[string]string `pulumi:"permissions"`
	// The **group** principal to assign the permissions.
	Principal string `pulumi:"principal"`
	// The ID of the project.
	ProjectId string `pulumi:"projectId"`
	// Replace (`true`) or merge (`false`) the permissions. Default: `true`
	//
	// | Permission        | Description                         |
	// | ----------------- | ----------------------------------- |
	// | View              | View library item                   |
	// | Administer        | Administer library item             |
	// | Create            | Create library item                 |
	// | ViewSecrets       | View library item secrets           |
	// | Use               | Use library item                    |
	// | Owner             | Owner library item                  |
	Replace *bool `pulumi:"replace"`
}

// The set of arguments for constructing a LibraryPermissions resource.
type LibraryPermissionsArgs struct {
	// the permissions to assign. The following permissions are available.
	Permissions pulumi.StringMapInput
	// The **group** principal to assign the permissions.
	Principal pulumi.StringInput
	// The ID of the project.
	ProjectId pulumi.StringInput
	// Replace (`true`) or merge (`false`) the permissions. Default: `true`
	//
	// | Permission        | Description                         |
	// | ----------------- | ----------------------------------- |
	// | View              | View library item                   |
	// | Administer        | Administer library item             |
	// | Create            | Create library item                 |
	// | ViewSecrets       | View library item secrets           |
	// | Use               | Use library item                    |
	// | Owner             | Owner library item                  |
	Replace pulumi.BoolPtrInput
}

func (LibraryPermissionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*libraryPermissionsArgs)(nil)).Elem()
}

type LibraryPermissionsInput interface {
	pulumi.Input

	ToLibraryPermissionsOutput() LibraryPermissionsOutput
	ToLibraryPermissionsOutputWithContext(ctx context.Context) LibraryPermissionsOutput
}

func (*LibraryPermissions) ElementType() reflect.Type {
	return reflect.TypeOf((**LibraryPermissions)(nil)).Elem()
}

func (i *LibraryPermissions) ToLibraryPermissionsOutput() LibraryPermissionsOutput {
	return i.ToLibraryPermissionsOutputWithContext(context.Background())
}

func (i *LibraryPermissions) ToLibraryPermissionsOutputWithContext(ctx context.Context) LibraryPermissionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LibraryPermissionsOutput)
}

// LibraryPermissionsArrayInput is an input type that accepts LibraryPermissionsArray and LibraryPermissionsArrayOutput values.
// You can construct a concrete instance of `LibraryPermissionsArrayInput` via:
//
//	LibraryPermissionsArray{ LibraryPermissionsArgs{...} }
type LibraryPermissionsArrayInput interface {
	pulumi.Input

	ToLibraryPermissionsArrayOutput() LibraryPermissionsArrayOutput
	ToLibraryPermissionsArrayOutputWithContext(context.Context) LibraryPermissionsArrayOutput
}

type LibraryPermissionsArray []LibraryPermissionsInput

func (LibraryPermissionsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LibraryPermissions)(nil)).Elem()
}

func (i LibraryPermissionsArray) ToLibraryPermissionsArrayOutput() LibraryPermissionsArrayOutput {
	return i.ToLibraryPermissionsArrayOutputWithContext(context.Background())
}

func (i LibraryPermissionsArray) ToLibraryPermissionsArrayOutputWithContext(ctx context.Context) LibraryPermissionsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LibraryPermissionsArrayOutput)
}

// LibraryPermissionsMapInput is an input type that accepts LibraryPermissionsMap and LibraryPermissionsMapOutput values.
// You can construct a concrete instance of `LibraryPermissionsMapInput` via:
//
//	LibraryPermissionsMap{ "key": LibraryPermissionsArgs{...} }
type LibraryPermissionsMapInput interface {
	pulumi.Input

	ToLibraryPermissionsMapOutput() LibraryPermissionsMapOutput
	ToLibraryPermissionsMapOutputWithContext(context.Context) LibraryPermissionsMapOutput
}

type LibraryPermissionsMap map[string]LibraryPermissionsInput

func (LibraryPermissionsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LibraryPermissions)(nil)).Elem()
}

func (i LibraryPermissionsMap) ToLibraryPermissionsMapOutput() LibraryPermissionsMapOutput {
	return i.ToLibraryPermissionsMapOutputWithContext(context.Background())
}

func (i LibraryPermissionsMap) ToLibraryPermissionsMapOutputWithContext(ctx context.Context) LibraryPermissionsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LibraryPermissionsMapOutput)
}

type LibraryPermissionsOutput struct{ *pulumi.OutputState }

func (LibraryPermissionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LibraryPermissions)(nil)).Elem()
}

func (o LibraryPermissionsOutput) ToLibraryPermissionsOutput() LibraryPermissionsOutput {
	return o
}

func (o LibraryPermissionsOutput) ToLibraryPermissionsOutputWithContext(ctx context.Context) LibraryPermissionsOutput {
	return o
}

// the permissions to assign. The following permissions are available.
func (o LibraryPermissionsOutput) Permissions() pulumi.StringMapOutput {
	return o.ApplyT(func(v *LibraryPermissions) pulumi.StringMapOutput { return v.Permissions }).(pulumi.StringMapOutput)
}

// The **group** principal to assign the permissions.
func (o LibraryPermissionsOutput) Principal() pulumi.StringOutput {
	return o.ApplyT(func(v *LibraryPermissions) pulumi.StringOutput { return v.Principal }).(pulumi.StringOutput)
}

// The ID of the project.
func (o LibraryPermissionsOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *LibraryPermissions) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// Replace (`true`) or merge (`false`) the permissions. Default: `true`
//
// | Permission        | Description                         |
// | ----------------- | ----------------------------------- |
// | View              | View library item                   |
// | Administer        | Administer library item             |
// | Create            | Create library item                 |
// | ViewSecrets       | View library item secrets           |
// | Use               | Use library item                    |
// | Owner             | Owner library item                  |
func (o LibraryPermissionsOutput) Replace() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LibraryPermissions) pulumi.BoolPtrOutput { return v.Replace }).(pulumi.BoolPtrOutput)
}

type LibraryPermissionsArrayOutput struct{ *pulumi.OutputState }

func (LibraryPermissionsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LibraryPermissions)(nil)).Elem()
}

func (o LibraryPermissionsArrayOutput) ToLibraryPermissionsArrayOutput() LibraryPermissionsArrayOutput {
	return o
}

func (o LibraryPermissionsArrayOutput) ToLibraryPermissionsArrayOutputWithContext(ctx context.Context) LibraryPermissionsArrayOutput {
	return o
}

func (o LibraryPermissionsArrayOutput) Index(i pulumi.IntInput) LibraryPermissionsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LibraryPermissions {
		return vs[0].([]*LibraryPermissions)[vs[1].(int)]
	}).(LibraryPermissionsOutput)
}

type LibraryPermissionsMapOutput struct{ *pulumi.OutputState }

func (LibraryPermissionsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LibraryPermissions)(nil)).Elem()
}

func (o LibraryPermissionsMapOutput) ToLibraryPermissionsMapOutput() LibraryPermissionsMapOutput {
	return o
}

func (o LibraryPermissionsMapOutput) ToLibraryPermissionsMapOutputWithContext(ctx context.Context) LibraryPermissionsMapOutput {
	return o
}

func (o LibraryPermissionsMapOutput) MapIndex(k pulumi.StringInput) LibraryPermissionsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LibraryPermissions {
		return vs[0].(map[string]*LibraryPermissions)[vs[1].(string)]
	}).(LibraryPermissionsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LibraryPermissionsInput)(nil)).Elem(), &LibraryPermissions{})
	pulumi.RegisterInputType(reflect.TypeOf((*LibraryPermissionsArrayInput)(nil)).Elem(), LibraryPermissionsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LibraryPermissionsMapInput)(nil)).Elem(), LibraryPermissionsMap{})
	pulumi.RegisterOutputType(LibraryPermissionsOutput{})
	pulumi.RegisterOutputType(LibraryPermissionsArrayOutput{})
	pulumi.RegisterOutputType(LibraryPermissionsMapOutput{})
}
