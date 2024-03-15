// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azuredevops/sdk/v3/go/azuredevops/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a user entitlement within Azure DevOps.
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
//			_, err := azuredevops.NewUser(ctx, "example", &azuredevops.UserArgs{
//				PrincipalName: pulumi.String("foo@contoso.com"),
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
// - [Azure DevOps Service REST API 7.0 - User Entitlements - Add](https://docs.microsoft.com/en-us/rest/api/azure/devops/memberentitlementmanagement/user-entitlements/add?view=azure-devops-rest-7.0)
// - [Programmatic mapping of access levels](https://docs.microsoft.com/en-us/azure/devops/organizations/security/access-levels?view=azure-devops#programmatic-mapping-of-access-levels)
//
// ## PAT Permissions Required
//
// - **Member Entitlement Management**: Read & Write
//
// ## Import
//
// The resources allows the import via the UUID of a user entitlement or by using the principal name of a user owning an entitlement.
type User struct {
	pulumi.CustomResourceState

	// Type of Account License. Valid values: `advanced`, `earlyAdopter`, `express`, `none`, `professional`, or `stakeholder`. Defaults to `express`. In addition the value `basic` is allowed which is an alias for `express` and reflects the name of the `express` license used in the Azure DevOps web interface.
	AccountLicenseType pulumi.StringPtrOutput `pulumi:"accountLicenseType"`
	// The descriptor is the primary way to reference the graph subject while the system is running. This field will uniquely identify the user graph subject.
	Descriptor pulumi.StringOutput `pulumi:"descriptor"`
	// The source of the licensing (e.g. Account. MSDN etc.) Valid values: `account` (Default), `auto`, `msdn`, `none`, `profile`, `trial`
	//
	// > **NOTE:** A user can only be referenced by it's `principalName` or by the combination of `originId` and `origin`.
	LicensingSource pulumi.StringPtrOutput `pulumi:"licensingSource"`
	// The type of source provider for the origin identifier.
	Origin pulumi.StringOutput `pulumi:"origin"`
	// The unique identifier from the system of origin. Typically a sid, object id or Guid. e.g. Used for member of other tenant on Azure Active Directory.
	OriginId pulumi.StringOutput `pulumi:"originId"`
	// The principal name is the PrincipalName of a graph member from the source provider. Usually, e-mail address.
	PrincipalName pulumi.StringOutput `pulumi:"principalName"`
}

// NewUser registers a new resource with the given unique name, arguments, and options.
func NewUser(ctx *pulumi.Context,
	name string, args *UserArgs, opts ...pulumi.ResourceOption) (*User, error) {
	if args == nil {
		args = &UserArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource User
	err := ctx.RegisterResource("azuredevops:index/user:User", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUser gets an existing User resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserState, opts ...pulumi.ResourceOption) (*User, error) {
	var resource User
	err := ctx.ReadResource("azuredevops:index/user:User", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering User resources.
type userState struct {
	// Type of Account License. Valid values: `advanced`, `earlyAdopter`, `express`, `none`, `professional`, or `stakeholder`. Defaults to `express`. In addition the value `basic` is allowed which is an alias for `express` and reflects the name of the `express` license used in the Azure DevOps web interface.
	AccountLicenseType *string `pulumi:"accountLicenseType"`
	// The descriptor is the primary way to reference the graph subject while the system is running. This field will uniquely identify the user graph subject.
	Descriptor *string `pulumi:"descriptor"`
	// The source of the licensing (e.g. Account. MSDN etc.) Valid values: `account` (Default), `auto`, `msdn`, `none`, `profile`, `trial`
	//
	// > **NOTE:** A user can only be referenced by it's `principalName` or by the combination of `originId` and `origin`.
	LicensingSource *string `pulumi:"licensingSource"`
	// The type of source provider for the origin identifier.
	Origin *string `pulumi:"origin"`
	// The unique identifier from the system of origin. Typically a sid, object id or Guid. e.g. Used for member of other tenant on Azure Active Directory.
	OriginId *string `pulumi:"originId"`
	// The principal name is the PrincipalName of a graph member from the source provider. Usually, e-mail address.
	PrincipalName *string `pulumi:"principalName"`
}

type UserState struct {
	// Type of Account License. Valid values: `advanced`, `earlyAdopter`, `express`, `none`, `professional`, or `stakeholder`. Defaults to `express`. In addition the value `basic` is allowed which is an alias for `express` and reflects the name of the `express` license used in the Azure DevOps web interface.
	AccountLicenseType pulumi.StringPtrInput
	// The descriptor is the primary way to reference the graph subject while the system is running. This field will uniquely identify the user graph subject.
	Descriptor pulumi.StringPtrInput
	// The source of the licensing (e.g. Account. MSDN etc.) Valid values: `account` (Default), `auto`, `msdn`, `none`, `profile`, `trial`
	//
	// > **NOTE:** A user can only be referenced by it's `principalName` or by the combination of `originId` and `origin`.
	LicensingSource pulumi.StringPtrInput
	// The type of source provider for the origin identifier.
	Origin pulumi.StringPtrInput
	// The unique identifier from the system of origin. Typically a sid, object id or Guid. e.g. Used for member of other tenant on Azure Active Directory.
	OriginId pulumi.StringPtrInput
	// The principal name is the PrincipalName of a graph member from the source provider. Usually, e-mail address.
	PrincipalName pulumi.StringPtrInput
}

func (UserState) ElementType() reflect.Type {
	return reflect.TypeOf((*userState)(nil)).Elem()
}

type userArgs struct {
	// Type of Account License. Valid values: `advanced`, `earlyAdopter`, `express`, `none`, `professional`, or `stakeholder`. Defaults to `express`. In addition the value `basic` is allowed which is an alias for `express` and reflects the name of the `express` license used in the Azure DevOps web interface.
	AccountLicenseType *string `pulumi:"accountLicenseType"`
	// The source of the licensing (e.g. Account. MSDN etc.) Valid values: `account` (Default), `auto`, `msdn`, `none`, `profile`, `trial`
	//
	// > **NOTE:** A user can only be referenced by it's `principalName` or by the combination of `originId` and `origin`.
	LicensingSource *string `pulumi:"licensingSource"`
	// The type of source provider for the origin identifier.
	Origin *string `pulumi:"origin"`
	// The unique identifier from the system of origin. Typically a sid, object id or Guid. e.g. Used for member of other tenant on Azure Active Directory.
	OriginId *string `pulumi:"originId"`
	// The principal name is the PrincipalName of a graph member from the source provider. Usually, e-mail address.
	PrincipalName *string `pulumi:"principalName"`
}

// The set of arguments for constructing a User resource.
type UserArgs struct {
	// Type of Account License. Valid values: `advanced`, `earlyAdopter`, `express`, `none`, `professional`, or `stakeholder`. Defaults to `express`. In addition the value `basic` is allowed which is an alias for `express` and reflects the name of the `express` license used in the Azure DevOps web interface.
	AccountLicenseType pulumi.StringPtrInput
	// The source of the licensing (e.g. Account. MSDN etc.) Valid values: `account` (Default), `auto`, `msdn`, `none`, `profile`, `trial`
	//
	// > **NOTE:** A user can only be referenced by it's `principalName` or by the combination of `originId` and `origin`.
	LicensingSource pulumi.StringPtrInput
	// The type of source provider for the origin identifier.
	Origin pulumi.StringPtrInput
	// The unique identifier from the system of origin. Typically a sid, object id or Guid. e.g. Used for member of other tenant on Azure Active Directory.
	OriginId pulumi.StringPtrInput
	// The principal name is the PrincipalName of a graph member from the source provider. Usually, e-mail address.
	PrincipalName pulumi.StringPtrInput
}

func (UserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userArgs)(nil)).Elem()
}

type UserInput interface {
	pulumi.Input

	ToUserOutput() UserOutput
	ToUserOutputWithContext(ctx context.Context) UserOutput
}

func (*User) ElementType() reflect.Type {
	return reflect.TypeOf((**User)(nil)).Elem()
}

func (i *User) ToUserOutput() UserOutput {
	return i.ToUserOutputWithContext(context.Background())
}

func (i *User) ToUserOutputWithContext(ctx context.Context) UserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserOutput)
}

// UserArrayInput is an input type that accepts UserArray and UserArrayOutput values.
// You can construct a concrete instance of `UserArrayInput` via:
//
//	UserArray{ UserArgs{...} }
type UserArrayInput interface {
	pulumi.Input

	ToUserArrayOutput() UserArrayOutput
	ToUserArrayOutputWithContext(context.Context) UserArrayOutput
}

type UserArray []UserInput

func (UserArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*User)(nil)).Elem()
}

func (i UserArray) ToUserArrayOutput() UserArrayOutput {
	return i.ToUserArrayOutputWithContext(context.Background())
}

func (i UserArray) ToUserArrayOutputWithContext(ctx context.Context) UserArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserArrayOutput)
}

// UserMapInput is an input type that accepts UserMap and UserMapOutput values.
// You can construct a concrete instance of `UserMapInput` via:
//
//	UserMap{ "key": UserArgs{...} }
type UserMapInput interface {
	pulumi.Input

	ToUserMapOutput() UserMapOutput
	ToUserMapOutputWithContext(context.Context) UserMapOutput
}

type UserMap map[string]UserInput

func (UserMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*User)(nil)).Elem()
}

func (i UserMap) ToUserMapOutput() UserMapOutput {
	return i.ToUserMapOutputWithContext(context.Background())
}

func (i UserMap) ToUserMapOutputWithContext(ctx context.Context) UserMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserMapOutput)
}

type UserOutput struct{ *pulumi.OutputState }

func (UserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**User)(nil)).Elem()
}

func (o UserOutput) ToUserOutput() UserOutput {
	return o
}

func (o UserOutput) ToUserOutputWithContext(ctx context.Context) UserOutput {
	return o
}

// Type of Account License. Valid values: `advanced`, `earlyAdopter`, `express`, `none`, `professional`, or `stakeholder`. Defaults to `express`. In addition the value `basic` is allowed which is an alias for `express` and reflects the name of the `express` license used in the Azure DevOps web interface.
func (o UserOutput) AccountLicenseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *User) pulumi.StringPtrOutput { return v.AccountLicenseType }).(pulumi.StringPtrOutput)
}

// The descriptor is the primary way to reference the graph subject while the system is running. This field will uniquely identify the user graph subject.
func (o UserOutput) Descriptor() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.Descriptor }).(pulumi.StringOutput)
}

// The source of the licensing (e.g. Account. MSDN etc.) Valid values: `account` (Default), `auto`, `msdn`, `none`, `profile`, `trial`
//
// > **NOTE:** A user can only be referenced by it's `principalName` or by the combination of `originId` and `origin`.
func (o UserOutput) LicensingSource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *User) pulumi.StringPtrOutput { return v.LicensingSource }).(pulumi.StringPtrOutput)
}

// The type of source provider for the origin identifier.
func (o UserOutput) Origin() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.Origin }).(pulumi.StringOutput)
}

// The unique identifier from the system of origin. Typically a sid, object id or Guid. e.g. Used for member of other tenant on Azure Active Directory.
func (o UserOutput) OriginId() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.OriginId }).(pulumi.StringOutput)
}

// The principal name is the PrincipalName of a graph member from the source provider. Usually, e-mail address.
func (o UserOutput) PrincipalName() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.PrincipalName }).(pulumi.StringOutput)
}

type UserArrayOutput struct{ *pulumi.OutputState }

func (UserArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*User)(nil)).Elem()
}

func (o UserArrayOutput) ToUserArrayOutput() UserArrayOutput {
	return o
}

func (o UserArrayOutput) ToUserArrayOutputWithContext(ctx context.Context) UserArrayOutput {
	return o
}

func (o UserArrayOutput) Index(i pulumi.IntInput) UserOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *User {
		return vs[0].([]*User)[vs[1].(int)]
	}).(UserOutput)
}

type UserMapOutput struct{ *pulumi.OutputState }

func (UserMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*User)(nil)).Elem()
}

func (o UserMapOutput) ToUserMapOutput() UserMapOutput {
	return o
}

func (o UserMapOutput) ToUserMapOutputWithContext(ctx context.Context) UserMapOutput {
	return o
}

func (o UserMapOutput) MapIndex(k pulumi.StringInput) UserOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *User {
		return vs[0].(map[string]*User)[vs[1].(string)]
	}).(UserOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserInput)(nil)).Elem(), &User{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserArrayInput)(nil)).Elem(), UserArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserMapInput)(nil)).Elem(), UserMap{})
	pulumi.RegisterOutputType(UserOutput{})
	pulumi.RegisterOutputType(UserArrayOutput{})
	pulumi.RegisterOutputType(UserMapOutput{})
}
