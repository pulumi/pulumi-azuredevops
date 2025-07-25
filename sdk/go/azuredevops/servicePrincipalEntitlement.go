// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azuredevops/sdk/v3/go/azuredevops/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Service Principal Entitlement.
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
//			_, err := azuredevops.NewServicePrincipalEntitlement(ctx, "example", &azuredevops.ServicePrincipalEntitlementArgs{
//				OriginId: pulumi.String("00000000-0000-0000-0000-000000000000"),
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
// ## Import
//
// Service Principal Entitlements can be imported using the `resource id`.
//
// The `resource id` can be found using DEV Tools in the `Users` section of the ADO organization.
//
// ```sh
// $ pulumi import azuredevops:index/servicePrincipalEntitlement:ServicePrincipalEntitlement example 8480c6eb-ce60-47e9-88df-eca3c801638b
// ```
type ServicePrincipalEntitlement struct {
	pulumi.CustomResourceState

	// Type of Account License. Possible values are: `advanced`, `earlyAdopter`, `express`, `none`, `professional`, or `stakeholder`. Defaults to `express`. In addition the value `basic` is allowed which is an alias for `express` and reflects the name of the `express` license used in the Azure DevOps web interface.
	AccountLicenseType pulumi.StringPtrOutput `pulumi:"accountLicenseType"`
	// The descriptor is the primary way to reference the graph subject while the system is running. This field will uniquely identify the user graph subject.
	Descriptor pulumi.StringOutput `pulumi:"descriptor"`
	// The display name of service principal.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The source of the licensing (e.g. Account. MSDN etc.) Possible values are: `account` (Default), `auto`, `msdn`, `none`, `profile`, `trial`
	LicensingSource pulumi.StringPtrOutput `pulumi:"licensingSource"`
	// The type of source provider for the origin identifier.
	Origin pulumi.StringOutput `pulumi:"origin"`
	// The Object ID of the service principal in Entra ID. Changing this forces a new Service Principal Entitlement to be created.
	OriginId pulumi.StringOutput `pulumi:"originId"`
}

// NewServicePrincipalEntitlement registers a new resource with the given unique name, arguments, and options.
func NewServicePrincipalEntitlement(ctx *pulumi.Context,
	name string, args *ServicePrincipalEntitlementArgs, opts ...pulumi.ResourceOption) (*ServicePrincipalEntitlement, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.OriginId == nil {
		return nil, errors.New("invalid value for required argument 'OriginId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ServicePrincipalEntitlement
	err := ctx.RegisterResource("azuredevops:index/servicePrincipalEntitlement:ServicePrincipalEntitlement", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServicePrincipalEntitlement gets an existing ServicePrincipalEntitlement resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServicePrincipalEntitlement(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServicePrincipalEntitlementState, opts ...pulumi.ResourceOption) (*ServicePrincipalEntitlement, error) {
	var resource ServicePrincipalEntitlement
	err := ctx.ReadResource("azuredevops:index/servicePrincipalEntitlement:ServicePrincipalEntitlement", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServicePrincipalEntitlement resources.
type servicePrincipalEntitlementState struct {
	// Type of Account License. Possible values are: `advanced`, `earlyAdopter`, `express`, `none`, `professional`, or `stakeholder`. Defaults to `express`. In addition the value `basic` is allowed which is an alias for `express` and reflects the name of the `express` license used in the Azure DevOps web interface.
	AccountLicenseType *string `pulumi:"accountLicenseType"`
	// The descriptor is the primary way to reference the graph subject while the system is running. This field will uniquely identify the user graph subject.
	Descriptor *string `pulumi:"descriptor"`
	// The display name of service principal.
	DisplayName *string `pulumi:"displayName"`
	// The source of the licensing (e.g. Account. MSDN etc.) Possible values are: `account` (Default), `auto`, `msdn`, `none`, `profile`, `trial`
	LicensingSource *string `pulumi:"licensingSource"`
	// The type of source provider for the origin identifier.
	Origin *string `pulumi:"origin"`
	// The Object ID of the service principal in Entra ID. Changing this forces a new Service Principal Entitlement to be created.
	OriginId *string `pulumi:"originId"`
}

type ServicePrincipalEntitlementState struct {
	// Type of Account License. Possible values are: `advanced`, `earlyAdopter`, `express`, `none`, `professional`, or `stakeholder`. Defaults to `express`. In addition the value `basic` is allowed which is an alias for `express` and reflects the name of the `express` license used in the Azure DevOps web interface.
	AccountLicenseType pulumi.StringPtrInput
	// The descriptor is the primary way to reference the graph subject while the system is running. This field will uniquely identify the user graph subject.
	Descriptor pulumi.StringPtrInput
	// The display name of service principal.
	DisplayName pulumi.StringPtrInput
	// The source of the licensing (e.g. Account. MSDN etc.) Possible values are: `account` (Default), `auto`, `msdn`, `none`, `profile`, `trial`
	LicensingSource pulumi.StringPtrInput
	// The type of source provider for the origin identifier.
	Origin pulumi.StringPtrInput
	// The Object ID of the service principal in Entra ID. Changing this forces a new Service Principal Entitlement to be created.
	OriginId pulumi.StringPtrInput
}

func (ServicePrincipalEntitlementState) ElementType() reflect.Type {
	return reflect.TypeOf((*servicePrincipalEntitlementState)(nil)).Elem()
}

type servicePrincipalEntitlementArgs struct {
	// Type of Account License. Possible values are: `advanced`, `earlyAdopter`, `express`, `none`, `professional`, or `stakeholder`. Defaults to `express`. In addition the value `basic` is allowed which is an alias for `express` and reflects the name of the `express` license used in the Azure DevOps web interface.
	AccountLicenseType *string `pulumi:"accountLicenseType"`
	// The source of the licensing (e.g. Account. MSDN etc.) Possible values are: `account` (Default), `auto`, `msdn`, `none`, `profile`, `trial`
	LicensingSource *string `pulumi:"licensingSource"`
	// The type of source provider for the origin identifier.
	Origin *string `pulumi:"origin"`
	// The Object ID of the service principal in Entra ID. Changing this forces a new Service Principal Entitlement to be created.
	OriginId string `pulumi:"originId"`
}

// The set of arguments for constructing a ServicePrincipalEntitlement resource.
type ServicePrincipalEntitlementArgs struct {
	// Type of Account License. Possible values are: `advanced`, `earlyAdopter`, `express`, `none`, `professional`, or `stakeholder`. Defaults to `express`. In addition the value `basic` is allowed which is an alias for `express` and reflects the name of the `express` license used in the Azure DevOps web interface.
	AccountLicenseType pulumi.StringPtrInput
	// The source of the licensing (e.g. Account. MSDN etc.) Possible values are: `account` (Default), `auto`, `msdn`, `none`, `profile`, `trial`
	LicensingSource pulumi.StringPtrInput
	// The type of source provider for the origin identifier.
	Origin pulumi.StringPtrInput
	// The Object ID of the service principal in Entra ID. Changing this forces a new Service Principal Entitlement to be created.
	OriginId pulumi.StringInput
}

func (ServicePrincipalEntitlementArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*servicePrincipalEntitlementArgs)(nil)).Elem()
}

type ServicePrincipalEntitlementInput interface {
	pulumi.Input

	ToServicePrincipalEntitlementOutput() ServicePrincipalEntitlementOutput
	ToServicePrincipalEntitlementOutputWithContext(ctx context.Context) ServicePrincipalEntitlementOutput
}

func (*ServicePrincipalEntitlement) ElementType() reflect.Type {
	return reflect.TypeOf((**ServicePrincipalEntitlement)(nil)).Elem()
}

func (i *ServicePrincipalEntitlement) ToServicePrincipalEntitlementOutput() ServicePrincipalEntitlementOutput {
	return i.ToServicePrincipalEntitlementOutputWithContext(context.Background())
}

func (i *ServicePrincipalEntitlement) ToServicePrincipalEntitlementOutputWithContext(ctx context.Context) ServicePrincipalEntitlementOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePrincipalEntitlementOutput)
}

// ServicePrincipalEntitlementArrayInput is an input type that accepts ServicePrincipalEntitlementArray and ServicePrincipalEntitlementArrayOutput values.
// You can construct a concrete instance of `ServicePrincipalEntitlementArrayInput` via:
//
//	ServicePrincipalEntitlementArray{ ServicePrincipalEntitlementArgs{...} }
type ServicePrincipalEntitlementArrayInput interface {
	pulumi.Input

	ToServicePrincipalEntitlementArrayOutput() ServicePrincipalEntitlementArrayOutput
	ToServicePrincipalEntitlementArrayOutputWithContext(context.Context) ServicePrincipalEntitlementArrayOutput
}

type ServicePrincipalEntitlementArray []ServicePrincipalEntitlementInput

func (ServicePrincipalEntitlementArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServicePrincipalEntitlement)(nil)).Elem()
}

func (i ServicePrincipalEntitlementArray) ToServicePrincipalEntitlementArrayOutput() ServicePrincipalEntitlementArrayOutput {
	return i.ToServicePrincipalEntitlementArrayOutputWithContext(context.Background())
}

func (i ServicePrincipalEntitlementArray) ToServicePrincipalEntitlementArrayOutputWithContext(ctx context.Context) ServicePrincipalEntitlementArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePrincipalEntitlementArrayOutput)
}

// ServicePrincipalEntitlementMapInput is an input type that accepts ServicePrincipalEntitlementMap and ServicePrincipalEntitlementMapOutput values.
// You can construct a concrete instance of `ServicePrincipalEntitlementMapInput` via:
//
//	ServicePrincipalEntitlementMap{ "key": ServicePrincipalEntitlementArgs{...} }
type ServicePrincipalEntitlementMapInput interface {
	pulumi.Input

	ToServicePrincipalEntitlementMapOutput() ServicePrincipalEntitlementMapOutput
	ToServicePrincipalEntitlementMapOutputWithContext(context.Context) ServicePrincipalEntitlementMapOutput
}

type ServicePrincipalEntitlementMap map[string]ServicePrincipalEntitlementInput

func (ServicePrincipalEntitlementMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServicePrincipalEntitlement)(nil)).Elem()
}

func (i ServicePrincipalEntitlementMap) ToServicePrincipalEntitlementMapOutput() ServicePrincipalEntitlementMapOutput {
	return i.ToServicePrincipalEntitlementMapOutputWithContext(context.Background())
}

func (i ServicePrincipalEntitlementMap) ToServicePrincipalEntitlementMapOutputWithContext(ctx context.Context) ServicePrincipalEntitlementMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePrincipalEntitlementMapOutput)
}

type ServicePrincipalEntitlementOutput struct{ *pulumi.OutputState }

func (ServicePrincipalEntitlementOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServicePrincipalEntitlement)(nil)).Elem()
}

func (o ServicePrincipalEntitlementOutput) ToServicePrincipalEntitlementOutput() ServicePrincipalEntitlementOutput {
	return o
}

func (o ServicePrincipalEntitlementOutput) ToServicePrincipalEntitlementOutputWithContext(ctx context.Context) ServicePrincipalEntitlementOutput {
	return o
}

// Type of Account License. Possible values are: `advanced`, `earlyAdopter`, `express`, `none`, `professional`, or `stakeholder`. Defaults to `express`. In addition the value `basic` is allowed which is an alias for `express` and reflects the name of the `express` license used in the Azure DevOps web interface.
func (o ServicePrincipalEntitlementOutput) AccountLicenseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServicePrincipalEntitlement) pulumi.StringPtrOutput { return v.AccountLicenseType }).(pulumi.StringPtrOutput)
}

// The descriptor is the primary way to reference the graph subject while the system is running. This field will uniquely identify the user graph subject.
func (o ServicePrincipalEntitlementOutput) Descriptor() pulumi.StringOutput {
	return o.ApplyT(func(v *ServicePrincipalEntitlement) pulumi.StringOutput { return v.Descriptor }).(pulumi.StringOutput)
}

// The display name of service principal.
func (o ServicePrincipalEntitlementOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *ServicePrincipalEntitlement) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// The source of the licensing (e.g. Account. MSDN etc.) Possible values are: `account` (Default), `auto`, `msdn`, `none`, `profile`, `trial`
func (o ServicePrincipalEntitlementOutput) LicensingSource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServicePrincipalEntitlement) pulumi.StringPtrOutput { return v.LicensingSource }).(pulumi.StringPtrOutput)
}

// The type of source provider for the origin identifier.
func (o ServicePrincipalEntitlementOutput) Origin() pulumi.StringOutput {
	return o.ApplyT(func(v *ServicePrincipalEntitlement) pulumi.StringOutput { return v.Origin }).(pulumi.StringOutput)
}

// The Object ID of the service principal in Entra ID. Changing this forces a new Service Principal Entitlement to be created.
func (o ServicePrincipalEntitlementOutput) OriginId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServicePrincipalEntitlement) pulumi.StringOutput { return v.OriginId }).(pulumi.StringOutput)
}

type ServicePrincipalEntitlementArrayOutput struct{ *pulumi.OutputState }

func (ServicePrincipalEntitlementArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServicePrincipalEntitlement)(nil)).Elem()
}

func (o ServicePrincipalEntitlementArrayOutput) ToServicePrincipalEntitlementArrayOutput() ServicePrincipalEntitlementArrayOutput {
	return o
}

func (o ServicePrincipalEntitlementArrayOutput) ToServicePrincipalEntitlementArrayOutputWithContext(ctx context.Context) ServicePrincipalEntitlementArrayOutput {
	return o
}

func (o ServicePrincipalEntitlementArrayOutput) Index(i pulumi.IntInput) ServicePrincipalEntitlementOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServicePrincipalEntitlement {
		return vs[0].([]*ServicePrincipalEntitlement)[vs[1].(int)]
	}).(ServicePrincipalEntitlementOutput)
}

type ServicePrincipalEntitlementMapOutput struct{ *pulumi.OutputState }

func (ServicePrincipalEntitlementMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServicePrincipalEntitlement)(nil)).Elem()
}

func (o ServicePrincipalEntitlementMapOutput) ToServicePrincipalEntitlementMapOutput() ServicePrincipalEntitlementMapOutput {
	return o
}

func (o ServicePrincipalEntitlementMapOutput) ToServicePrincipalEntitlementMapOutputWithContext(ctx context.Context) ServicePrincipalEntitlementMapOutput {
	return o
}

func (o ServicePrincipalEntitlementMapOutput) MapIndex(k pulumi.StringInput) ServicePrincipalEntitlementOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServicePrincipalEntitlement {
		return vs[0].(map[string]*ServicePrincipalEntitlement)[vs[1].(string)]
	}).(ServicePrincipalEntitlementOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServicePrincipalEntitlementInput)(nil)).Elem(), &ServicePrincipalEntitlement{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServicePrincipalEntitlementArrayInput)(nil)).Elem(), ServicePrincipalEntitlementArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServicePrincipalEntitlementMapInput)(nil)).Elem(), ServicePrincipalEntitlementMap{})
	pulumi.RegisterOutputType(ServicePrincipalEntitlementOutput{})
	pulumi.RegisterOutputType(ServicePrincipalEntitlementArrayOutput{})
	pulumi.RegisterOutputType(ServicePrincipalEntitlementMapOutput{})
}
