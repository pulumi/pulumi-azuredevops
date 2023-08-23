// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azuredevops/sdk/v2/go/azuredevops/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Approval Check.
//
// ## Example Usage
// ### Protect an environment
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-azuredevops/sdk/v2/go/azuredevops"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleProject, err := azuredevops.NewProject(ctx, "exampleProject", nil)
//			if err != nil {
//				return err
//			}
//			exampleEnvironment, err := azuredevops.NewEnvironment(ctx, "exampleEnvironment", &azuredevops.EnvironmentArgs{
//				ProjectId: exampleProject.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			exampleGroup, err := azuredevops.NewGroup(ctx, "exampleGroup", &azuredevops.GroupArgs{
//				DisplayName: pulumi.String("some-azdo-group"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuredevops.NewCheckApproval(ctx, "exampleCheckApproval", &azuredevops.CheckApprovalArgs{
//				ProjectId:           exampleProject.ID(),
//				TargetResourceId:    exampleEnvironment.ID(),
//				TargetResourceType:  pulumi.String("environment"),
//				RequesterCanApprove: pulumi.Bool(true),
//				Approvers: pulumi.StringArray{
//					exampleGroup.OriginId,
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
// ## Import
//
// Importing this resource is not supported.
type CheckApproval struct {
	pulumi.CustomResourceState

	// Specifies a list of approver IDs.
	Approvers pulumi.StringArrayOutput `pulumi:"approvers"`
	// The instructions for the approvers.
	Instructions pulumi.StringPtrOutput `pulumi:"instructions"`
	// The minimum number of approvers. This property is applicable when there is more than 1 approver.
	MinimumRequiredApprovers pulumi.IntPtrOutput `pulumi:"minimumRequiredApprovers"`
	// The project ID. Changing this forces a new Approval Check to be created.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Can the requestor approve? Defaults to `false`.
	RequesterCanApprove pulumi.BoolPtrOutput `pulumi:"requesterCanApprove"`
	// The ID of the resource being protected by the check. Changing this forces a new Approval Check to be created.
	TargetResourceId pulumi.StringOutput `pulumi:"targetResourceId"`
	// The type of resource being protected by the check. Valid values: `endpoint`, `environment`, `queue`, `repository`, `securefile`, `variablegroup`. Changing this forces a new Approval Check to be created.
	TargetResourceType pulumi.StringOutput `pulumi:"targetResourceType"`
	// The timeout in minutes for the approval.  Defaults to `43200`.
	Timeout pulumi.IntPtrOutput `pulumi:"timeout"`
}

// NewCheckApproval registers a new resource with the given unique name, arguments, and options.
func NewCheckApproval(ctx *pulumi.Context,
	name string, args *CheckApprovalArgs, opts ...pulumi.ResourceOption) (*CheckApproval, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Approvers == nil {
		return nil, errors.New("invalid value for required argument 'Approvers'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.TargetResourceId == nil {
		return nil, errors.New("invalid value for required argument 'TargetResourceId'")
	}
	if args.TargetResourceType == nil {
		return nil, errors.New("invalid value for required argument 'TargetResourceType'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CheckApproval
	err := ctx.RegisterResource("azuredevops:index/checkApproval:CheckApproval", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCheckApproval gets an existing CheckApproval resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCheckApproval(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CheckApprovalState, opts ...pulumi.ResourceOption) (*CheckApproval, error) {
	var resource CheckApproval
	err := ctx.ReadResource("azuredevops:index/checkApproval:CheckApproval", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CheckApproval resources.
type checkApprovalState struct {
	// Specifies a list of approver IDs.
	Approvers []string `pulumi:"approvers"`
	// The instructions for the approvers.
	Instructions *string `pulumi:"instructions"`
	// The minimum number of approvers. This property is applicable when there is more than 1 approver.
	MinimumRequiredApprovers *int `pulumi:"minimumRequiredApprovers"`
	// The project ID. Changing this forces a new Approval Check to be created.
	ProjectId *string `pulumi:"projectId"`
	// Can the requestor approve? Defaults to `false`.
	RequesterCanApprove *bool `pulumi:"requesterCanApprove"`
	// The ID of the resource being protected by the check. Changing this forces a new Approval Check to be created.
	TargetResourceId *string `pulumi:"targetResourceId"`
	// The type of resource being protected by the check. Valid values: `endpoint`, `environment`, `queue`, `repository`, `securefile`, `variablegroup`. Changing this forces a new Approval Check to be created.
	TargetResourceType *string `pulumi:"targetResourceType"`
	// The timeout in minutes for the approval.  Defaults to `43200`.
	Timeout *int `pulumi:"timeout"`
}

type CheckApprovalState struct {
	// Specifies a list of approver IDs.
	Approvers pulumi.StringArrayInput
	// The instructions for the approvers.
	Instructions pulumi.StringPtrInput
	// The minimum number of approvers. This property is applicable when there is more than 1 approver.
	MinimumRequiredApprovers pulumi.IntPtrInput
	// The project ID. Changing this forces a new Approval Check to be created.
	ProjectId pulumi.StringPtrInput
	// Can the requestor approve? Defaults to `false`.
	RequesterCanApprove pulumi.BoolPtrInput
	// The ID of the resource being protected by the check. Changing this forces a new Approval Check to be created.
	TargetResourceId pulumi.StringPtrInput
	// The type of resource being protected by the check. Valid values: `endpoint`, `environment`, `queue`, `repository`, `securefile`, `variablegroup`. Changing this forces a new Approval Check to be created.
	TargetResourceType pulumi.StringPtrInput
	// The timeout in minutes for the approval.  Defaults to `43200`.
	Timeout pulumi.IntPtrInput
}

func (CheckApprovalState) ElementType() reflect.Type {
	return reflect.TypeOf((*checkApprovalState)(nil)).Elem()
}

type checkApprovalArgs struct {
	// Specifies a list of approver IDs.
	Approvers []string `pulumi:"approvers"`
	// The instructions for the approvers.
	Instructions *string `pulumi:"instructions"`
	// The minimum number of approvers. This property is applicable when there is more than 1 approver.
	MinimumRequiredApprovers *int `pulumi:"minimumRequiredApprovers"`
	// The project ID. Changing this forces a new Approval Check to be created.
	ProjectId string `pulumi:"projectId"`
	// Can the requestor approve? Defaults to `false`.
	RequesterCanApprove *bool `pulumi:"requesterCanApprove"`
	// The ID of the resource being protected by the check. Changing this forces a new Approval Check to be created.
	TargetResourceId string `pulumi:"targetResourceId"`
	// The type of resource being protected by the check. Valid values: `endpoint`, `environment`, `queue`, `repository`, `securefile`, `variablegroup`. Changing this forces a new Approval Check to be created.
	TargetResourceType string `pulumi:"targetResourceType"`
	// The timeout in minutes for the approval.  Defaults to `43200`.
	Timeout *int `pulumi:"timeout"`
}

// The set of arguments for constructing a CheckApproval resource.
type CheckApprovalArgs struct {
	// Specifies a list of approver IDs.
	Approvers pulumi.StringArrayInput
	// The instructions for the approvers.
	Instructions pulumi.StringPtrInput
	// The minimum number of approvers. This property is applicable when there is more than 1 approver.
	MinimumRequiredApprovers pulumi.IntPtrInput
	// The project ID. Changing this forces a new Approval Check to be created.
	ProjectId pulumi.StringInput
	// Can the requestor approve? Defaults to `false`.
	RequesterCanApprove pulumi.BoolPtrInput
	// The ID of the resource being protected by the check. Changing this forces a new Approval Check to be created.
	TargetResourceId pulumi.StringInput
	// The type of resource being protected by the check. Valid values: `endpoint`, `environment`, `queue`, `repository`, `securefile`, `variablegroup`. Changing this forces a new Approval Check to be created.
	TargetResourceType pulumi.StringInput
	// The timeout in minutes for the approval.  Defaults to `43200`.
	Timeout pulumi.IntPtrInput
}

func (CheckApprovalArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*checkApprovalArgs)(nil)).Elem()
}

type CheckApprovalInput interface {
	pulumi.Input

	ToCheckApprovalOutput() CheckApprovalOutput
	ToCheckApprovalOutputWithContext(ctx context.Context) CheckApprovalOutput
}

func (*CheckApproval) ElementType() reflect.Type {
	return reflect.TypeOf((**CheckApproval)(nil)).Elem()
}

func (i *CheckApproval) ToCheckApprovalOutput() CheckApprovalOutput {
	return i.ToCheckApprovalOutputWithContext(context.Background())
}

func (i *CheckApproval) ToCheckApprovalOutputWithContext(ctx context.Context) CheckApprovalOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CheckApprovalOutput)
}

// CheckApprovalArrayInput is an input type that accepts CheckApprovalArray and CheckApprovalArrayOutput values.
// You can construct a concrete instance of `CheckApprovalArrayInput` via:
//
//	CheckApprovalArray{ CheckApprovalArgs{...} }
type CheckApprovalArrayInput interface {
	pulumi.Input

	ToCheckApprovalArrayOutput() CheckApprovalArrayOutput
	ToCheckApprovalArrayOutputWithContext(context.Context) CheckApprovalArrayOutput
}

type CheckApprovalArray []CheckApprovalInput

func (CheckApprovalArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CheckApproval)(nil)).Elem()
}

func (i CheckApprovalArray) ToCheckApprovalArrayOutput() CheckApprovalArrayOutput {
	return i.ToCheckApprovalArrayOutputWithContext(context.Background())
}

func (i CheckApprovalArray) ToCheckApprovalArrayOutputWithContext(ctx context.Context) CheckApprovalArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CheckApprovalArrayOutput)
}

// CheckApprovalMapInput is an input type that accepts CheckApprovalMap and CheckApprovalMapOutput values.
// You can construct a concrete instance of `CheckApprovalMapInput` via:
//
//	CheckApprovalMap{ "key": CheckApprovalArgs{...} }
type CheckApprovalMapInput interface {
	pulumi.Input

	ToCheckApprovalMapOutput() CheckApprovalMapOutput
	ToCheckApprovalMapOutputWithContext(context.Context) CheckApprovalMapOutput
}

type CheckApprovalMap map[string]CheckApprovalInput

func (CheckApprovalMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CheckApproval)(nil)).Elem()
}

func (i CheckApprovalMap) ToCheckApprovalMapOutput() CheckApprovalMapOutput {
	return i.ToCheckApprovalMapOutputWithContext(context.Background())
}

func (i CheckApprovalMap) ToCheckApprovalMapOutputWithContext(ctx context.Context) CheckApprovalMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CheckApprovalMapOutput)
}

type CheckApprovalOutput struct{ *pulumi.OutputState }

func (CheckApprovalOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CheckApproval)(nil)).Elem()
}

func (o CheckApprovalOutput) ToCheckApprovalOutput() CheckApprovalOutput {
	return o
}

func (o CheckApprovalOutput) ToCheckApprovalOutputWithContext(ctx context.Context) CheckApprovalOutput {
	return o
}

// Specifies a list of approver IDs.
func (o CheckApprovalOutput) Approvers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CheckApproval) pulumi.StringArrayOutput { return v.Approvers }).(pulumi.StringArrayOutput)
}

// The instructions for the approvers.
func (o CheckApprovalOutput) Instructions() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CheckApproval) pulumi.StringPtrOutput { return v.Instructions }).(pulumi.StringPtrOutput)
}

// The minimum number of approvers. This property is applicable when there is more than 1 approver.
func (o CheckApprovalOutput) MinimumRequiredApprovers() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CheckApproval) pulumi.IntPtrOutput { return v.MinimumRequiredApprovers }).(pulumi.IntPtrOutput)
}

// The project ID. Changing this forces a new Approval Check to be created.
func (o CheckApprovalOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *CheckApproval) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// Can the requestor approve? Defaults to `false`.
func (o CheckApprovalOutput) RequesterCanApprove() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CheckApproval) pulumi.BoolPtrOutput { return v.RequesterCanApprove }).(pulumi.BoolPtrOutput)
}

// The ID of the resource being protected by the check. Changing this forces a new Approval Check to be created.
func (o CheckApprovalOutput) TargetResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *CheckApproval) pulumi.StringOutput { return v.TargetResourceId }).(pulumi.StringOutput)
}

// The type of resource being protected by the check. Valid values: `endpoint`, `environment`, `queue`, `repository`, `securefile`, `variablegroup`. Changing this forces a new Approval Check to be created.
func (o CheckApprovalOutput) TargetResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *CheckApproval) pulumi.StringOutput { return v.TargetResourceType }).(pulumi.StringOutput)
}

// The timeout in minutes for the approval.  Defaults to `43200`.
func (o CheckApprovalOutput) Timeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CheckApproval) pulumi.IntPtrOutput { return v.Timeout }).(pulumi.IntPtrOutput)
}

type CheckApprovalArrayOutput struct{ *pulumi.OutputState }

func (CheckApprovalArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CheckApproval)(nil)).Elem()
}

func (o CheckApprovalArrayOutput) ToCheckApprovalArrayOutput() CheckApprovalArrayOutput {
	return o
}

func (o CheckApprovalArrayOutput) ToCheckApprovalArrayOutputWithContext(ctx context.Context) CheckApprovalArrayOutput {
	return o
}

func (o CheckApprovalArrayOutput) Index(i pulumi.IntInput) CheckApprovalOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CheckApproval {
		return vs[0].([]*CheckApproval)[vs[1].(int)]
	}).(CheckApprovalOutput)
}

type CheckApprovalMapOutput struct{ *pulumi.OutputState }

func (CheckApprovalMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CheckApproval)(nil)).Elem()
}

func (o CheckApprovalMapOutput) ToCheckApprovalMapOutput() CheckApprovalMapOutput {
	return o
}

func (o CheckApprovalMapOutput) ToCheckApprovalMapOutputWithContext(ctx context.Context) CheckApprovalMapOutput {
	return o
}

func (o CheckApprovalMapOutput) MapIndex(k pulumi.StringInput) CheckApprovalOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CheckApproval {
		return vs[0].(map[string]*CheckApproval)[vs[1].(string)]
	}).(CheckApprovalOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CheckApprovalInput)(nil)).Elem(), &CheckApproval{})
	pulumi.RegisterInputType(reflect.TypeOf((*CheckApprovalArrayInput)(nil)).Elem(), CheckApprovalArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CheckApprovalMapInput)(nil)).Elem(), CheckApprovalMap{})
	pulumi.RegisterOutputType(CheckApprovalOutput{})
	pulumi.RegisterOutputType(CheckApprovalArrayOutput{})
	pulumi.RegisterOutputType(CheckApprovalMapOutput{})
}