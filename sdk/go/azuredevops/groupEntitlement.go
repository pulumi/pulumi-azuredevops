// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azuredevops/sdk/v2/go/azuredevops/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Manages a group entitlement within Azure DevOps.
//
// ## Example Usage
// ## Relevant Links
//
// - [Azure DevOps Service REST API 7.0 - Group Entitlements](https://learn.microsoft.com/en-us/rest/api/azure/devops/memberentitlementmanagement/group-entitlements?view=azure-devops-rest-7.1)
// - [Programmatic mapping of access levels](https://docs.microsoft.com/en-us/azure/devops/organizations/security/access-levels?view=azure-devops#programmatic-mapping-of-access-levels)
//
// ## PAT Permissions Required
//
// - **Member Entitlement Management**: Read & Write
//
// ## Import
//
// The resource allows the import via the ID of a group entitlement, which is a UUID.
//
// ```sh
//
//	$ pulumi import azuredevops:index/groupEntitlement:GroupEntitlement example 00000000-0000-0000-0000-000000000000
//
// ```
type GroupEntitlement struct {
	pulumi.CustomResourceState

	// Type of Account License. Valid values: `advanced`, `earlyAdopter`, `express`, `none`, `professional`, or `stakeholder`. Defaults to `express`. In addition, the value `basic` is allowed which is an alias for `express` and reflects the name of the `express` license used in the Azure DevOps web interface.
	AccountLicenseType pulumi.StringPtrOutput `pulumi:"accountLicenseType"`
	// The descriptor is the primary way to reference the graph subject while the system is running. This field will uniquely identify the group graph subject.
	Descriptor pulumi.StringOutput `pulumi:"descriptor"`
	// The display name is the name used in Azure DevOps UI. Cannot be set together with `originId` and `origin`.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The source of the licensing (e.g. Account. MSDN etc.) Valid values: `account` (Default), `auto`, `msdn`, `none`, `profile`, `trial`
	//
	// > **NOTE:** A existing group in Azure AD can only be referenced by the combination of `originId` and `origin`.
	LicensingSource pulumi.StringPtrOutput `pulumi:"licensingSource"`
	// The type of source provider for the origin identifier.
	Origin pulumi.StringOutput `pulumi:"origin"`
	// The unique identifier from the system of origin. Typically, a sid, object id or Guid. e.g. Used for member of other tenant on Azure Active Directory.
	OriginId pulumi.StringOutput `pulumi:"originId"`
	// The principal name of a graph member on Azure DevOps
	PrincipalName pulumi.StringOutput `pulumi:"principalName"`
}

// NewGroupEntitlement registers a new resource with the given unique name, arguments, and options.
func NewGroupEntitlement(ctx *pulumi.Context,
	name string, args *GroupEntitlementArgs, opts ...pulumi.ResourceOption) (*GroupEntitlement, error) {
	if args == nil {
		args = &GroupEntitlementArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource GroupEntitlement
	err := ctx.RegisterResource("azuredevops:index/groupEntitlement:GroupEntitlement", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroupEntitlement gets an existing GroupEntitlement resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroupEntitlement(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupEntitlementState, opts ...pulumi.ResourceOption) (*GroupEntitlement, error) {
	var resource GroupEntitlement
	err := ctx.ReadResource("azuredevops:index/groupEntitlement:GroupEntitlement", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GroupEntitlement resources.
type groupEntitlementState struct {
	// Type of Account License. Valid values: `advanced`, `earlyAdopter`, `express`, `none`, `professional`, or `stakeholder`. Defaults to `express`. In addition, the value `basic` is allowed which is an alias for `express` and reflects the name of the `express` license used in the Azure DevOps web interface.
	AccountLicenseType *string `pulumi:"accountLicenseType"`
	// The descriptor is the primary way to reference the graph subject while the system is running. This field will uniquely identify the group graph subject.
	Descriptor *string `pulumi:"descriptor"`
	// The display name is the name used in Azure DevOps UI. Cannot be set together with `originId` and `origin`.
	DisplayName *string `pulumi:"displayName"`
	// The source of the licensing (e.g. Account. MSDN etc.) Valid values: `account` (Default), `auto`, `msdn`, `none`, `profile`, `trial`
	//
	// > **NOTE:** A existing group in Azure AD can only be referenced by the combination of `originId` and `origin`.
	LicensingSource *string `pulumi:"licensingSource"`
	// The type of source provider for the origin identifier.
	Origin *string `pulumi:"origin"`
	// The unique identifier from the system of origin. Typically, a sid, object id or Guid. e.g. Used for member of other tenant on Azure Active Directory.
	OriginId *string `pulumi:"originId"`
	// The principal name of a graph member on Azure DevOps
	PrincipalName *string `pulumi:"principalName"`
}

type GroupEntitlementState struct {
	// Type of Account License. Valid values: `advanced`, `earlyAdopter`, `express`, `none`, `professional`, or `stakeholder`. Defaults to `express`. In addition, the value `basic` is allowed which is an alias for `express` and reflects the name of the `express` license used in the Azure DevOps web interface.
	AccountLicenseType pulumi.StringPtrInput
	// The descriptor is the primary way to reference the graph subject while the system is running. This field will uniquely identify the group graph subject.
	Descriptor pulumi.StringPtrInput
	// The display name is the name used in Azure DevOps UI. Cannot be set together with `originId` and `origin`.
	DisplayName pulumi.StringPtrInput
	// The source of the licensing (e.g. Account. MSDN etc.) Valid values: `account` (Default), `auto`, `msdn`, `none`, `profile`, `trial`
	//
	// > **NOTE:** A existing group in Azure AD can only be referenced by the combination of `originId` and `origin`.
	LicensingSource pulumi.StringPtrInput
	// The type of source provider for the origin identifier.
	Origin pulumi.StringPtrInput
	// The unique identifier from the system of origin. Typically, a sid, object id or Guid. e.g. Used for member of other tenant on Azure Active Directory.
	OriginId pulumi.StringPtrInput
	// The principal name of a graph member on Azure DevOps
	PrincipalName pulumi.StringPtrInput
}

func (GroupEntitlementState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupEntitlementState)(nil)).Elem()
}

type groupEntitlementArgs struct {
	// Type of Account License. Valid values: `advanced`, `earlyAdopter`, `express`, `none`, `professional`, or `stakeholder`. Defaults to `express`. In addition, the value `basic` is allowed which is an alias for `express` and reflects the name of the `express` license used in the Azure DevOps web interface.
	AccountLicenseType *string `pulumi:"accountLicenseType"`
	// The display name is the name used in Azure DevOps UI. Cannot be set together with `originId` and `origin`.
	DisplayName *string `pulumi:"displayName"`
	// The source of the licensing (e.g. Account. MSDN etc.) Valid values: `account` (Default), `auto`, `msdn`, `none`, `profile`, `trial`
	//
	// > **NOTE:** A existing group in Azure AD can only be referenced by the combination of `originId` and `origin`.
	LicensingSource *string `pulumi:"licensingSource"`
	// The type of source provider for the origin identifier.
	Origin *string `pulumi:"origin"`
	// The unique identifier from the system of origin. Typically, a sid, object id or Guid. e.g. Used for member of other tenant on Azure Active Directory.
	OriginId *string `pulumi:"originId"`
}

// The set of arguments for constructing a GroupEntitlement resource.
type GroupEntitlementArgs struct {
	// Type of Account License. Valid values: `advanced`, `earlyAdopter`, `express`, `none`, `professional`, or `stakeholder`. Defaults to `express`. In addition, the value `basic` is allowed which is an alias for `express` and reflects the name of the `express` license used in the Azure DevOps web interface.
	AccountLicenseType pulumi.StringPtrInput
	// The display name is the name used in Azure DevOps UI. Cannot be set together with `originId` and `origin`.
	DisplayName pulumi.StringPtrInput
	// The source of the licensing (e.g. Account. MSDN etc.) Valid values: `account` (Default), `auto`, `msdn`, `none`, `profile`, `trial`
	//
	// > **NOTE:** A existing group in Azure AD can only be referenced by the combination of `originId` and `origin`.
	LicensingSource pulumi.StringPtrInput
	// The type of source provider for the origin identifier.
	Origin pulumi.StringPtrInput
	// The unique identifier from the system of origin. Typically, a sid, object id or Guid. e.g. Used for member of other tenant on Azure Active Directory.
	OriginId pulumi.StringPtrInput
}

func (GroupEntitlementArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupEntitlementArgs)(nil)).Elem()
}

type GroupEntitlementInput interface {
	pulumi.Input

	ToGroupEntitlementOutput() GroupEntitlementOutput
	ToGroupEntitlementOutputWithContext(ctx context.Context) GroupEntitlementOutput
}

func (*GroupEntitlement) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupEntitlement)(nil)).Elem()
}

func (i *GroupEntitlement) ToGroupEntitlementOutput() GroupEntitlementOutput {
	return i.ToGroupEntitlementOutputWithContext(context.Background())
}

func (i *GroupEntitlement) ToGroupEntitlementOutputWithContext(ctx context.Context) GroupEntitlementOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupEntitlementOutput)
}

func (i *GroupEntitlement) ToOutput(ctx context.Context) pulumix.Output[*GroupEntitlement] {
	return pulumix.Output[*GroupEntitlement]{
		OutputState: i.ToGroupEntitlementOutputWithContext(ctx).OutputState,
	}
}

// GroupEntitlementArrayInput is an input type that accepts GroupEntitlementArray and GroupEntitlementArrayOutput values.
// You can construct a concrete instance of `GroupEntitlementArrayInput` via:
//
//	GroupEntitlementArray{ GroupEntitlementArgs{...} }
type GroupEntitlementArrayInput interface {
	pulumi.Input

	ToGroupEntitlementArrayOutput() GroupEntitlementArrayOutput
	ToGroupEntitlementArrayOutputWithContext(context.Context) GroupEntitlementArrayOutput
}

type GroupEntitlementArray []GroupEntitlementInput

func (GroupEntitlementArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GroupEntitlement)(nil)).Elem()
}

func (i GroupEntitlementArray) ToGroupEntitlementArrayOutput() GroupEntitlementArrayOutput {
	return i.ToGroupEntitlementArrayOutputWithContext(context.Background())
}

func (i GroupEntitlementArray) ToGroupEntitlementArrayOutputWithContext(ctx context.Context) GroupEntitlementArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupEntitlementArrayOutput)
}

func (i GroupEntitlementArray) ToOutput(ctx context.Context) pulumix.Output[[]*GroupEntitlement] {
	return pulumix.Output[[]*GroupEntitlement]{
		OutputState: i.ToGroupEntitlementArrayOutputWithContext(ctx).OutputState,
	}
}

// GroupEntitlementMapInput is an input type that accepts GroupEntitlementMap and GroupEntitlementMapOutput values.
// You can construct a concrete instance of `GroupEntitlementMapInput` via:
//
//	GroupEntitlementMap{ "key": GroupEntitlementArgs{...} }
type GroupEntitlementMapInput interface {
	pulumi.Input

	ToGroupEntitlementMapOutput() GroupEntitlementMapOutput
	ToGroupEntitlementMapOutputWithContext(context.Context) GroupEntitlementMapOutput
}

type GroupEntitlementMap map[string]GroupEntitlementInput

func (GroupEntitlementMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GroupEntitlement)(nil)).Elem()
}

func (i GroupEntitlementMap) ToGroupEntitlementMapOutput() GroupEntitlementMapOutput {
	return i.ToGroupEntitlementMapOutputWithContext(context.Background())
}

func (i GroupEntitlementMap) ToGroupEntitlementMapOutputWithContext(ctx context.Context) GroupEntitlementMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupEntitlementMapOutput)
}

func (i GroupEntitlementMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*GroupEntitlement] {
	return pulumix.Output[map[string]*GroupEntitlement]{
		OutputState: i.ToGroupEntitlementMapOutputWithContext(ctx).OutputState,
	}
}

type GroupEntitlementOutput struct{ *pulumi.OutputState }

func (GroupEntitlementOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupEntitlement)(nil)).Elem()
}

func (o GroupEntitlementOutput) ToGroupEntitlementOutput() GroupEntitlementOutput {
	return o
}

func (o GroupEntitlementOutput) ToGroupEntitlementOutputWithContext(ctx context.Context) GroupEntitlementOutput {
	return o
}

func (o GroupEntitlementOutput) ToOutput(ctx context.Context) pulumix.Output[*GroupEntitlement] {
	return pulumix.Output[*GroupEntitlement]{
		OutputState: o.OutputState,
	}
}

// Type of Account License. Valid values: `advanced`, `earlyAdopter`, `express`, `none`, `professional`, or `stakeholder`. Defaults to `express`. In addition, the value `basic` is allowed which is an alias for `express` and reflects the name of the `express` license used in the Azure DevOps web interface.
func (o GroupEntitlementOutput) AccountLicenseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GroupEntitlement) pulumi.StringPtrOutput { return v.AccountLicenseType }).(pulumi.StringPtrOutput)
}

// The descriptor is the primary way to reference the graph subject while the system is running. This field will uniquely identify the group graph subject.
func (o GroupEntitlementOutput) Descriptor() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupEntitlement) pulumi.StringOutput { return v.Descriptor }).(pulumi.StringOutput)
}

// The display name is the name used in Azure DevOps UI. Cannot be set together with `originId` and `origin`.
func (o GroupEntitlementOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupEntitlement) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// The source of the licensing (e.g. Account. MSDN etc.) Valid values: `account` (Default), `auto`, `msdn`, `none`, `profile`, `trial`
//
// > **NOTE:** A existing group in Azure AD can only be referenced by the combination of `originId` and `origin`.
func (o GroupEntitlementOutput) LicensingSource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GroupEntitlement) pulumi.StringPtrOutput { return v.LicensingSource }).(pulumi.StringPtrOutput)
}

// The type of source provider for the origin identifier.
func (o GroupEntitlementOutput) Origin() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupEntitlement) pulumi.StringOutput { return v.Origin }).(pulumi.StringOutput)
}

// The unique identifier from the system of origin. Typically, a sid, object id or Guid. e.g. Used for member of other tenant on Azure Active Directory.
func (o GroupEntitlementOutput) OriginId() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupEntitlement) pulumi.StringOutput { return v.OriginId }).(pulumi.StringOutput)
}

// The principal name of a graph member on Azure DevOps
func (o GroupEntitlementOutput) PrincipalName() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupEntitlement) pulumi.StringOutput { return v.PrincipalName }).(pulumi.StringOutput)
}

type GroupEntitlementArrayOutput struct{ *pulumi.OutputState }

func (GroupEntitlementArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GroupEntitlement)(nil)).Elem()
}

func (o GroupEntitlementArrayOutput) ToGroupEntitlementArrayOutput() GroupEntitlementArrayOutput {
	return o
}

func (o GroupEntitlementArrayOutput) ToGroupEntitlementArrayOutputWithContext(ctx context.Context) GroupEntitlementArrayOutput {
	return o
}

func (o GroupEntitlementArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*GroupEntitlement] {
	return pulumix.Output[[]*GroupEntitlement]{
		OutputState: o.OutputState,
	}
}

func (o GroupEntitlementArrayOutput) Index(i pulumi.IntInput) GroupEntitlementOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GroupEntitlement {
		return vs[0].([]*GroupEntitlement)[vs[1].(int)]
	}).(GroupEntitlementOutput)
}

type GroupEntitlementMapOutput struct{ *pulumi.OutputState }

func (GroupEntitlementMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GroupEntitlement)(nil)).Elem()
}

func (o GroupEntitlementMapOutput) ToGroupEntitlementMapOutput() GroupEntitlementMapOutput {
	return o
}

func (o GroupEntitlementMapOutput) ToGroupEntitlementMapOutputWithContext(ctx context.Context) GroupEntitlementMapOutput {
	return o
}

func (o GroupEntitlementMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*GroupEntitlement] {
	return pulumix.Output[map[string]*GroupEntitlement]{
		OutputState: o.OutputState,
	}
}

func (o GroupEntitlementMapOutput) MapIndex(k pulumi.StringInput) GroupEntitlementOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GroupEntitlement {
		return vs[0].(map[string]*GroupEntitlement)[vs[1].(string)]
	}).(GroupEntitlementOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GroupEntitlementInput)(nil)).Elem(), &GroupEntitlement{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupEntitlementArrayInput)(nil)).Elem(), GroupEntitlementArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupEntitlementMapInput)(nil)).Elem(), GroupEntitlementMap{})
	pulumi.RegisterOutputType(GroupEntitlementOutput{})
	pulumi.RegisterOutputType(GroupEntitlementArrayOutput{})
	pulumi.RegisterOutputType(GroupEntitlementMapOutput{})
}
