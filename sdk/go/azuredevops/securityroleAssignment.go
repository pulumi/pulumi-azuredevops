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

// Manages assignment of security roles to various resources within Azure DevOps organization.
type SecurityroleAssignment struct {
	pulumi.CustomResourceState

	// The ID of the identity to authorize.
	IdentityId pulumi.StringOutput `pulumi:"identityId"`
	// The ID of the resource on which the role is to be assigned.
	ResourceId pulumi.StringOutput `pulumi:"resourceId"`
	// Name of the role to assign.
	RoleName pulumi.StringOutput `pulumi:"roleName"`
	// The scope in which this assignment should exist.
	Scope pulumi.StringOutput `pulumi:"scope"`
}

// NewSecurityroleAssignment registers a new resource with the given unique name, arguments, and options.
func NewSecurityroleAssignment(ctx *pulumi.Context,
	name string, args *SecurityroleAssignmentArgs, opts ...pulumi.ResourceOption) (*SecurityroleAssignment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IdentityId == nil {
		return nil, errors.New("invalid value for required argument 'IdentityId'")
	}
	if args.ResourceId == nil {
		return nil, errors.New("invalid value for required argument 'ResourceId'")
	}
	if args.RoleName == nil {
		return nil, errors.New("invalid value for required argument 'RoleName'")
	}
	if args.Scope == nil {
		return nil, errors.New("invalid value for required argument 'Scope'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SecurityroleAssignment
	err := ctx.RegisterResource("azuredevops:index/securityroleAssignment:SecurityroleAssignment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecurityroleAssignment gets an existing SecurityroleAssignment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecurityroleAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecurityroleAssignmentState, opts ...pulumi.ResourceOption) (*SecurityroleAssignment, error) {
	var resource SecurityroleAssignment
	err := ctx.ReadResource("azuredevops:index/securityroleAssignment:SecurityroleAssignment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecurityroleAssignment resources.
type securityroleAssignmentState struct {
	// The ID of the identity to authorize.
	IdentityId *string `pulumi:"identityId"`
	// The ID of the resource on which the role is to be assigned.
	ResourceId *string `pulumi:"resourceId"`
	// Name of the role to assign.
	RoleName *string `pulumi:"roleName"`
	// The scope in which this assignment should exist.
	Scope *string `pulumi:"scope"`
}

type SecurityroleAssignmentState struct {
	// The ID of the identity to authorize.
	IdentityId pulumi.StringPtrInput
	// The ID of the resource on which the role is to be assigned.
	ResourceId pulumi.StringPtrInput
	// Name of the role to assign.
	RoleName pulumi.StringPtrInput
	// The scope in which this assignment should exist.
	Scope pulumi.StringPtrInput
}

func (SecurityroleAssignmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*securityroleAssignmentState)(nil)).Elem()
}

type securityroleAssignmentArgs struct {
	// The ID of the identity to authorize.
	IdentityId string `pulumi:"identityId"`
	// The ID of the resource on which the role is to be assigned.
	ResourceId string `pulumi:"resourceId"`
	// Name of the role to assign.
	RoleName string `pulumi:"roleName"`
	// The scope in which this assignment should exist.
	Scope string `pulumi:"scope"`
}

// The set of arguments for constructing a SecurityroleAssignment resource.
type SecurityroleAssignmentArgs struct {
	// The ID of the identity to authorize.
	IdentityId pulumi.StringInput
	// The ID of the resource on which the role is to be assigned.
	ResourceId pulumi.StringInput
	// Name of the role to assign.
	RoleName pulumi.StringInput
	// The scope in which this assignment should exist.
	Scope pulumi.StringInput
}

func (SecurityroleAssignmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*securityroleAssignmentArgs)(nil)).Elem()
}

type SecurityroleAssignmentInput interface {
	pulumi.Input

	ToSecurityroleAssignmentOutput() SecurityroleAssignmentOutput
	ToSecurityroleAssignmentOutputWithContext(ctx context.Context) SecurityroleAssignmentOutput
}

func (*SecurityroleAssignment) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityroleAssignment)(nil)).Elem()
}

func (i *SecurityroleAssignment) ToSecurityroleAssignmentOutput() SecurityroleAssignmentOutput {
	return i.ToSecurityroleAssignmentOutputWithContext(context.Background())
}

func (i *SecurityroleAssignment) ToSecurityroleAssignmentOutputWithContext(ctx context.Context) SecurityroleAssignmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityroleAssignmentOutput)
}

// SecurityroleAssignmentArrayInput is an input type that accepts SecurityroleAssignmentArray and SecurityroleAssignmentArrayOutput values.
// You can construct a concrete instance of `SecurityroleAssignmentArrayInput` via:
//
//	SecurityroleAssignmentArray{ SecurityroleAssignmentArgs{...} }
type SecurityroleAssignmentArrayInput interface {
	pulumi.Input

	ToSecurityroleAssignmentArrayOutput() SecurityroleAssignmentArrayOutput
	ToSecurityroleAssignmentArrayOutputWithContext(context.Context) SecurityroleAssignmentArrayOutput
}

type SecurityroleAssignmentArray []SecurityroleAssignmentInput

func (SecurityroleAssignmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecurityroleAssignment)(nil)).Elem()
}

func (i SecurityroleAssignmentArray) ToSecurityroleAssignmentArrayOutput() SecurityroleAssignmentArrayOutput {
	return i.ToSecurityroleAssignmentArrayOutputWithContext(context.Background())
}

func (i SecurityroleAssignmentArray) ToSecurityroleAssignmentArrayOutputWithContext(ctx context.Context) SecurityroleAssignmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityroleAssignmentArrayOutput)
}

// SecurityroleAssignmentMapInput is an input type that accepts SecurityroleAssignmentMap and SecurityroleAssignmentMapOutput values.
// You can construct a concrete instance of `SecurityroleAssignmentMapInput` via:
//
//	SecurityroleAssignmentMap{ "key": SecurityroleAssignmentArgs{...} }
type SecurityroleAssignmentMapInput interface {
	pulumi.Input

	ToSecurityroleAssignmentMapOutput() SecurityroleAssignmentMapOutput
	ToSecurityroleAssignmentMapOutputWithContext(context.Context) SecurityroleAssignmentMapOutput
}

type SecurityroleAssignmentMap map[string]SecurityroleAssignmentInput

func (SecurityroleAssignmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecurityroleAssignment)(nil)).Elem()
}

func (i SecurityroleAssignmentMap) ToSecurityroleAssignmentMapOutput() SecurityroleAssignmentMapOutput {
	return i.ToSecurityroleAssignmentMapOutputWithContext(context.Background())
}

func (i SecurityroleAssignmentMap) ToSecurityroleAssignmentMapOutputWithContext(ctx context.Context) SecurityroleAssignmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityroleAssignmentMapOutput)
}

type SecurityroleAssignmentOutput struct{ *pulumi.OutputState }

func (SecurityroleAssignmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityroleAssignment)(nil)).Elem()
}

func (o SecurityroleAssignmentOutput) ToSecurityroleAssignmentOutput() SecurityroleAssignmentOutput {
	return o
}

func (o SecurityroleAssignmentOutput) ToSecurityroleAssignmentOutputWithContext(ctx context.Context) SecurityroleAssignmentOutput {
	return o
}

// The ID of the identity to authorize.
func (o SecurityroleAssignmentOutput) IdentityId() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityroleAssignment) pulumi.StringOutput { return v.IdentityId }).(pulumi.StringOutput)
}

// The ID of the resource on which the role is to be assigned.
func (o SecurityroleAssignmentOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityroleAssignment) pulumi.StringOutput { return v.ResourceId }).(pulumi.StringOutput)
}

// Name of the role to assign.
func (o SecurityroleAssignmentOutput) RoleName() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityroleAssignment) pulumi.StringOutput { return v.RoleName }).(pulumi.StringOutput)
}

// The scope in which this assignment should exist.
func (o SecurityroleAssignmentOutput) Scope() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityroleAssignment) pulumi.StringOutput { return v.Scope }).(pulumi.StringOutput)
}

type SecurityroleAssignmentArrayOutput struct{ *pulumi.OutputState }

func (SecurityroleAssignmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecurityroleAssignment)(nil)).Elem()
}

func (o SecurityroleAssignmentArrayOutput) ToSecurityroleAssignmentArrayOutput() SecurityroleAssignmentArrayOutput {
	return o
}

func (o SecurityroleAssignmentArrayOutput) ToSecurityroleAssignmentArrayOutputWithContext(ctx context.Context) SecurityroleAssignmentArrayOutput {
	return o
}

func (o SecurityroleAssignmentArrayOutput) Index(i pulumi.IntInput) SecurityroleAssignmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SecurityroleAssignment {
		return vs[0].([]*SecurityroleAssignment)[vs[1].(int)]
	}).(SecurityroleAssignmentOutput)
}

type SecurityroleAssignmentMapOutput struct{ *pulumi.OutputState }

func (SecurityroleAssignmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecurityroleAssignment)(nil)).Elem()
}

func (o SecurityroleAssignmentMapOutput) ToSecurityroleAssignmentMapOutput() SecurityroleAssignmentMapOutput {
	return o
}

func (o SecurityroleAssignmentMapOutput) ToSecurityroleAssignmentMapOutputWithContext(ctx context.Context) SecurityroleAssignmentMapOutput {
	return o
}

func (o SecurityroleAssignmentMapOutput) MapIndex(k pulumi.StringInput) SecurityroleAssignmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SecurityroleAssignment {
		return vs[0].(map[string]*SecurityroleAssignment)[vs[1].(string)]
	}).(SecurityroleAssignmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecurityroleAssignmentInput)(nil)).Elem(), &SecurityroleAssignment{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecurityroleAssignmentArrayInput)(nil)).Elem(), SecurityroleAssignmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecurityroleAssignmentMapInput)(nil)).Elem(), SecurityroleAssignmentMap{})
	pulumi.RegisterOutputType(SecurityroleAssignmentOutput{})
	pulumi.RegisterOutputType(SecurityroleAssignmentArrayOutput{})
	pulumi.RegisterOutputType(SecurityroleAssignmentMapOutput{})
}