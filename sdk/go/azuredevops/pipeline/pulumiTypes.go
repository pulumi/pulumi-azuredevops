// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package pipeline

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type VariableGroupKeyVault struct {
	// The name of the Variable Group.
	Name              string `pulumi:"name"`
	ServiceEndpointId string `pulumi:"serviceEndpointId"`
}

// VariableGroupKeyVaultInput is an input type that accepts VariableGroupKeyVaultArgs and VariableGroupKeyVaultOutput values.
// You can construct a concrete instance of `VariableGroupKeyVaultInput` via:
//
//          VariableGroupKeyVaultArgs{...}
type VariableGroupKeyVaultInput interface {
	pulumi.Input

	ToVariableGroupKeyVaultOutput() VariableGroupKeyVaultOutput
	ToVariableGroupKeyVaultOutputWithContext(context.Context) VariableGroupKeyVaultOutput
}

type VariableGroupKeyVaultArgs struct {
	// The name of the Variable Group.
	Name              pulumi.StringInput `pulumi:"name"`
	ServiceEndpointId pulumi.StringInput `pulumi:"serviceEndpointId"`
}

func (VariableGroupKeyVaultArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VariableGroupKeyVault)(nil)).Elem()
}

func (i VariableGroupKeyVaultArgs) ToVariableGroupKeyVaultOutput() VariableGroupKeyVaultOutput {
	return i.ToVariableGroupKeyVaultOutputWithContext(context.Background())
}

func (i VariableGroupKeyVaultArgs) ToVariableGroupKeyVaultOutputWithContext(ctx context.Context) VariableGroupKeyVaultOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VariableGroupKeyVaultOutput)
}

func (i VariableGroupKeyVaultArgs) ToVariableGroupKeyVaultPtrOutput() VariableGroupKeyVaultPtrOutput {
	return i.ToVariableGroupKeyVaultPtrOutputWithContext(context.Background())
}

func (i VariableGroupKeyVaultArgs) ToVariableGroupKeyVaultPtrOutputWithContext(ctx context.Context) VariableGroupKeyVaultPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VariableGroupKeyVaultOutput).ToVariableGroupKeyVaultPtrOutputWithContext(ctx)
}

// VariableGroupKeyVaultPtrInput is an input type that accepts VariableGroupKeyVaultArgs, VariableGroupKeyVaultPtr and VariableGroupKeyVaultPtrOutput values.
// You can construct a concrete instance of `VariableGroupKeyVaultPtrInput` via:
//
//          VariableGroupKeyVaultArgs{...}
//
//  or:
//
//          nil
type VariableGroupKeyVaultPtrInput interface {
	pulumi.Input

	ToVariableGroupKeyVaultPtrOutput() VariableGroupKeyVaultPtrOutput
	ToVariableGroupKeyVaultPtrOutputWithContext(context.Context) VariableGroupKeyVaultPtrOutput
}

type variableGroupKeyVaultPtrType VariableGroupKeyVaultArgs

func VariableGroupKeyVaultPtr(v *VariableGroupKeyVaultArgs) VariableGroupKeyVaultPtrInput {
	return (*variableGroupKeyVaultPtrType)(v)
}

func (*variableGroupKeyVaultPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VariableGroupKeyVault)(nil)).Elem()
}

func (i *variableGroupKeyVaultPtrType) ToVariableGroupKeyVaultPtrOutput() VariableGroupKeyVaultPtrOutput {
	return i.ToVariableGroupKeyVaultPtrOutputWithContext(context.Background())
}

func (i *variableGroupKeyVaultPtrType) ToVariableGroupKeyVaultPtrOutputWithContext(ctx context.Context) VariableGroupKeyVaultPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VariableGroupKeyVaultPtrOutput)
}

type VariableGroupKeyVaultOutput struct{ *pulumi.OutputState }

func (VariableGroupKeyVaultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VariableGroupKeyVault)(nil)).Elem()
}

func (o VariableGroupKeyVaultOutput) ToVariableGroupKeyVaultOutput() VariableGroupKeyVaultOutput {
	return o
}

func (o VariableGroupKeyVaultOutput) ToVariableGroupKeyVaultOutputWithContext(ctx context.Context) VariableGroupKeyVaultOutput {
	return o
}

func (o VariableGroupKeyVaultOutput) ToVariableGroupKeyVaultPtrOutput() VariableGroupKeyVaultPtrOutput {
	return o.ToVariableGroupKeyVaultPtrOutputWithContext(context.Background())
}

func (o VariableGroupKeyVaultOutput) ToVariableGroupKeyVaultPtrOutputWithContext(ctx context.Context) VariableGroupKeyVaultPtrOutput {
	return o.ApplyT(func(v VariableGroupKeyVault) *VariableGroupKeyVault {
		return &v
	}).(VariableGroupKeyVaultPtrOutput)
}

// The name of the Variable Group.
func (o VariableGroupKeyVaultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v VariableGroupKeyVault) string { return v.Name }).(pulumi.StringOutput)
}

func (o VariableGroupKeyVaultOutput) ServiceEndpointId() pulumi.StringOutput {
	return o.ApplyT(func(v VariableGroupKeyVault) string { return v.ServiceEndpointId }).(pulumi.StringOutput)
}

type VariableGroupKeyVaultPtrOutput struct{ *pulumi.OutputState }

func (VariableGroupKeyVaultPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VariableGroupKeyVault)(nil)).Elem()
}

func (o VariableGroupKeyVaultPtrOutput) ToVariableGroupKeyVaultPtrOutput() VariableGroupKeyVaultPtrOutput {
	return o
}

func (o VariableGroupKeyVaultPtrOutput) ToVariableGroupKeyVaultPtrOutputWithContext(ctx context.Context) VariableGroupKeyVaultPtrOutput {
	return o
}

func (o VariableGroupKeyVaultPtrOutput) Elem() VariableGroupKeyVaultOutput {
	return o.ApplyT(func(v *VariableGroupKeyVault) VariableGroupKeyVault { return *v }).(VariableGroupKeyVaultOutput)
}

// The name of the Variable Group.
func (o VariableGroupKeyVaultPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VariableGroupKeyVault) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o VariableGroupKeyVaultPtrOutput) ServiceEndpointId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VariableGroupKeyVault) *string {
		if v == nil {
			return nil
		}
		return &v.ServiceEndpointId
	}).(pulumi.StringPtrOutput)
}

type VariableGroupVariable struct {
	ContentType *string `pulumi:"contentType"`
	Enabled     *bool   `pulumi:"enabled"`
	Expires     *string `pulumi:"expires"`
	// A boolean flag describing if the variable value is sensitive. Defaults to `false`.
	IsSecret *bool `pulumi:"isSecret"`
	// The key value used for the variable. Must be unique within the Variable Group.
	Name string `pulumi:"name"`
	// The secret value of the variable. If omitted, it will default to empty string. Used when `isSecret` set to `true`.
	SecretValue *string `pulumi:"secretValue"`
	// The value of the variable. If omitted, it will default to empty string.
	Value *string `pulumi:"value"`
}

// VariableGroupVariableInput is an input type that accepts VariableGroupVariableArgs and VariableGroupVariableOutput values.
// You can construct a concrete instance of `VariableGroupVariableInput` via:
//
//          VariableGroupVariableArgs{...}
type VariableGroupVariableInput interface {
	pulumi.Input

	ToVariableGroupVariableOutput() VariableGroupVariableOutput
	ToVariableGroupVariableOutputWithContext(context.Context) VariableGroupVariableOutput
}

type VariableGroupVariableArgs struct {
	ContentType pulumi.StringPtrInput `pulumi:"contentType"`
	Enabled     pulumi.BoolPtrInput   `pulumi:"enabled"`
	Expires     pulumi.StringPtrInput `pulumi:"expires"`
	// A boolean flag describing if the variable value is sensitive. Defaults to `false`.
	IsSecret pulumi.BoolPtrInput `pulumi:"isSecret"`
	// The key value used for the variable. Must be unique within the Variable Group.
	Name pulumi.StringInput `pulumi:"name"`
	// The secret value of the variable. If omitted, it will default to empty string. Used when `isSecret` set to `true`.
	SecretValue pulumi.StringPtrInput `pulumi:"secretValue"`
	// The value of the variable. If omitted, it will default to empty string.
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (VariableGroupVariableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VariableGroupVariable)(nil)).Elem()
}

func (i VariableGroupVariableArgs) ToVariableGroupVariableOutput() VariableGroupVariableOutput {
	return i.ToVariableGroupVariableOutputWithContext(context.Background())
}

func (i VariableGroupVariableArgs) ToVariableGroupVariableOutputWithContext(ctx context.Context) VariableGroupVariableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VariableGroupVariableOutput)
}

// VariableGroupVariableArrayInput is an input type that accepts VariableGroupVariableArray and VariableGroupVariableArrayOutput values.
// You can construct a concrete instance of `VariableGroupVariableArrayInput` via:
//
//          VariableGroupVariableArray{ VariableGroupVariableArgs{...} }
type VariableGroupVariableArrayInput interface {
	pulumi.Input

	ToVariableGroupVariableArrayOutput() VariableGroupVariableArrayOutput
	ToVariableGroupVariableArrayOutputWithContext(context.Context) VariableGroupVariableArrayOutput
}

type VariableGroupVariableArray []VariableGroupVariableInput

func (VariableGroupVariableArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VariableGroupVariable)(nil)).Elem()
}

func (i VariableGroupVariableArray) ToVariableGroupVariableArrayOutput() VariableGroupVariableArrayOutput {
	return i.ToVariableGroupVariableArrayOutputWithContext(context.Background())
}

func (i VariableGroupVariableArray) ToVariableGroupVariableArrayOutputWithContext(ctx context.Context) VariableGroupVariableArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VariableGroupVariableArrayOutput)
}

type VariableGroupVariableOutput struct{ *pulumi.OutputState }

func (VariableGroupVariableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VariableGroupVariable)(nil)).Elem()
}

func (o VariableGroupVariableOutput) ToVariableGroupVariableOutput() VariableGroupVariableOutput {
	return o
}

func (o VariableGroupVariableOutput) ToVariableGroupVariableOutputWithContext(ctx context.Context) VariableGroupVariableOutput {
	return o
}

func (o VariableGroupVariableOutput) ContentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VariableGroupVariable) *string { return v.ContentType }).(pulumi.StringPtrOutput)
}

func (o VariableGroupVariableOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VariableGroupVariable) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o VariableGroupVariableOutput) Expires() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VariableGroupVariable) *string { return v.Expires }).(pulumi.StringPtrOutput)
}

// A boolean flag describing if the variable value is sensitive. Defaults to `false`.
func (o VariableGroupVariableOutput) IsSecret() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VariableGroupVariable) *bool { return v.IsSecret }).(pulumi.BoolPtrOutput)
}

// The key value used for the variable. Must be unique within the Variable Group.
func (o VariableGroupVariableOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v VariableGroupVariable) string { return v.Name }).(pulumi.StringOutput)
}

// The secret value of the variable. If omitted, it will default to empty string. Used when `isSecret` set to `true`.
func (o VariableGroupVariableOutput) SecretValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VariableGroupVariable) *string { return v.SecretValue }).(pulumi.StringPtrOutput)
}

// The value of the variable. If omitted, it will default to empty string.
func (o VariableGroupVariableOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VariableGroupVariable) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type VariableGroupVariableArrayOutput struct{ *pulumi.OutputState }

func (VariableGroupVariableArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VariableGroupVariable)(nil)).Elem()
}

func (o VariableGroupVariableArrayOutput) ToVariableGroupVariableArrayOutput() VariableGroupVariableArrayOutput {
	return o
}

func (o VariableGroupVariableArrayOutput) ToVariableGroupVariableArrayOutputWithContext(ctx context.Context) VariableGroupVariableArrayOutput {
	return o
}

func (o VariableGroupVariableArrayOutput) Index(i pulumi.IntInput) VariableGroupVariableOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VariableGroupVariable {
		return vs[0].([]VariableGroupVariable)[vs[1].(int)]
	}).(VariableGroupVariableOutput)
}

func init() {
	pulumi.RegisterOutputType(VariableGroupKeyVaultOutput{})
	pulumi.RegisterOutputType(VariableGroupKeyVaultPtrOutput{})
	pulumi.RegisterOutputType(VariableGroupVariableOutput{})
	pulumi.RegisterOutputType(VariableGroupVariableArrayOutput{})
}
