// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package entitlement

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a user entitlement within Azure DevOps.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azuredevops/sdk/go/azuredevops/Entitlement"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Entitlement.NewUser(ctx, "user", &Entitlement.UserArgs{
// 			PrincipalName: pulumi.String("foo@contoso.com"),
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
// * [Azure DevOps Service REST API 5.1 - User Entitlements - Add](https://docs.microsoft.com/en-us/rest/api/azure/devops/memberentitlementmanagement/user%20entitlements/add?view=azure-devops-rest-5.1)
//
// ## PAT Permissions Required
//
// - **Member Entitlement Management**: Read & Write
type User struct {
	pulumi.CustomResourceState

	// Type of Account License. Valid values: `advanced`, `earlyAdopter`, `express`, `none`, `professional`, or `stakeholder`. Defaults to `express`. In addition the value `basic` is allowed which is an alias for `express` and reflects the name of the `express` license used in the Azure DevOps web interface.
	AccountLicenseType pulumi.StringPtrOutput `pulumi:"accountLicenseType"`
	// The descriptor is the primary way to reference the graph subject while the system is running. This field will uniquely identify the user graph subject.
	Descriptor pulumi.StringOutput `pulumi:"descriptor"`
	// The source of the licensing (e.g. Account. MSDN etc.) Valid values: `account` (Default), `auto`, `msdn`, `none`, `profile`, `trail`
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
	var resource User
	err := ctx.RegisterResource("azuredevops:Entitlement/user:User", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azuredevops:Entitlement/user:User", name, id, state, &resource, opts...)
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
	// The source of the licensing (e.g. Account. MSDN etc.) Valid values: `account` (Default), `auto`, `msdn`, `none`, `profile`, `trail`
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
	// The source of the licensing (e.g. Account. MSDN etc.) Valid values: `account` (Default), `auto`, `msdn`, `none`, `profile`, `trail`
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
	// The source of the licensing (e.g. Account. MSDN etc.) Valid values: `account` (Default), `auto`, `msdn`, `none`, `profile`, `trail`
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
	// The source of the licensing (e.g. Account. MSDN etc.) Valid values: `account` (Default), `auto`, `msdn`, `none`, `profile`, `trail`
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
