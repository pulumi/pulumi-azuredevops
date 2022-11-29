// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package identities

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages group membership within Azure DevOps.
//
// ## Example Usage
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
//			exampleUser, err := azuredevops.NewUser(ctx, "exampleUser", &azuredevops.UserArgs{
//				PrincipalName: pulumi.String("foo@contoso.com"),
//			})
//			if err != nil {
//				return err
//			}
//			exampleGroup := azuredevops.LookupGroupOutput(ctx, GetGroupOutputArgs{
//				ProjectId: exampleProject.ID(),
//				Name:      pulumi.String("Build Administrators"),
//			}, nil)
//			_, err = azuredevops.NewGroupMembership(ctx, "exampleGroupMembership", &azuredevops.GroupMembershipArgs{
//				Group: exampleGroup.ApplyT(func(exampleGroup GetGroupResult) (string, error) {
//					return exampleGroup.Descriptor, nil
//				}).(pulumi.StringOutput),
//				Members: pulumi.StringArray{
//					exampleUser.Descriptor,
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
// ## Relevant Links
//
// - [Azure DevOps Service REST API 6.0 - Memberships](https://docs.microsoft.com/en-us/rest/api/azure/devops/graph/memberships?view=azure-devops-rest-6.0)
//
// ## PAT Permissions Required
//
// - **Deployment Groups**: Read & Manage
//
// ## Import
//
// Not supported.
//
// Deprecated: azuredevops.identities.GroupMembership has been deprecated in favor of azuredevops.GroupMembership
type GroupMembership struct {
	pulumi.CustomResourceState

	// The descriptor of the group being managed.
	Group pulumi.StringOutput `pulumi:"group"`
	// A list of user or group descriptors that will become members of the group.
	// > NOTE: It's possible to define group members both within the `GroupMembership resource` via the members block and by using the `Group` resource. However it's not possible to use both methods to manage group members, since there'll be conflicts.
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// The mode how the resource manages group members.
	// - `mode == add`: the resource will ensure that all specified members will be part of the referenced group
	// - `mode == overwrite`: the resource will replace all existing members with the members specified within the `members` block
	// > NOTE: To clear all members from a group, specify an empty list of descriptors in the `members` attribute and set the `mode` member to `overwrite`.
	Mode pulumi.StringPtrOutput `pulumi:"mode"`
}

// NewGroupMembership registers a new resource with the given unique name, arguments, and options.
func NewGroupMembership(ctx *pulumi.Context,
	name string, args *GroupMembershipArgs, opts ...pulumi.ResourceOption) (*GroupMembership, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Group == nil {
		return nil, errors.New("invalid value for required argument 'Group'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	var resource GroupMembership
	err := ctx.RegisterResource("azuredevops:Identities/groupMembership:GroupMembership", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroupMembership gets an existing GroupMembership resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroupMembership(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupMembershipState, opts ...pulumi.ResourceOption) (*GroupMembership, error) {
	var resource GroupMembership
	err := ctx.ReadResource("azuredevops:Identities/groupMembership:GroupMembership", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GroupMembership resources.
type groupMembershipState struct {
	// The descriptor of the group being managed.
	Group *string `pulumi:"group"`
	// A list of user or group descriptors that will become members of the group.
	// > NOTE: It's possible to define group members both within the `GroupMembership resource` via the members block and by using the `Group` resource. However it's not possible to use both methods to manage group members, since there'll be conflicts.
	Members []string `pulumi:"members"`
	// The mode how the resource manages group members.
	// - `mode == add`: the resource will ensure that all specified members will be part of the referenced group
	// - `mode == overwrite`: the resource will replace all existing members with the members specified within the `members` block
	// > NOTE: To clear all members from a group, specify an empty list of descriptors in the `members` attribute and set the `mode` member to `overwrite`.
	Mode *string `pulumi:"mode"`
}

type GroupMembershipState struct {
	// The descriptor of the group being managed.
	Group pulumi.StringPtrInput
	// A list of user or group descriptors that will become members of the group.
	// > NOTE: It's possible to define group members both within the `GroupMembership resource` via the members block and by using the `Group` resource. However it's not possible to use both methods to manage group members, since there'll be conflicts.
	Members pulumi.StringArrayInput
	// The mode how the resource manages group members.
	// - `mode == add`: the resource will ensure that all specified members will be part of the referenced group
	// - `mode == overwrite`: the resource will replace all existing members with the members specified within the `members` block
	// > NOTE: To clear all members from a group, specify an empty list of descriptors in the `members` attribute and set the `mode` member to `overwrite`.
	Mode pulumi.StringPtrInput
}

func (GroupMembershipState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupMembershipState)(nil)).Elem()
}

type groupMembershipArgs struct {
	// The descriptor of the group being managed.
	Group string `pulumi:"group"`
	// A list of user or group descriptors that will become members of the group.
	// > NOTE: It's possible to define group members both within the `GroupMembership resource` via the members block and by using the `Group` resource. However it's not possible to use both methods to manage group members, since there'll be conflicts.
	Members []string `pulumi:"members"`
	// The mode how the resource manages group members.
	// - `mode == add`: the resource will ensure that all specified members will be part of the referenced group
	// - `mode == overwrite`: the resource will replace all existing members with the members specified within the `members` block
	// > NOTE: To clear all members from a group, specify an empty list of descriptors in the `members` attribute and set the `mode` member to `overwrite`.
	Mode *string `pulumi:"mode"`
}

// The set of arguments for constructing a GroupMembership resource.
type GroupMembershipArgs struct {
	// The descriptor of the group being managed.
	Group pulumi.StringInput
	// A list of user or group descriptors that will become members of the group.
	// > NOTE: It's possible to define group members both within the `GroupMembership resource` via the members block and by using the `Group` resource. However it's not possible to use both methods to manage group members, since there'll be conflicts.
	Members pulumi.StringArrayInput
	// The mode how the resource manages group members.
	// - `mode == add`: the resource will ensure that all specified members will be part of the referenced group
	// - `mode == overwrite`: the resource will replace all existing members with the members specified within the `members` block
	// > NOTE: To clear all members from a group, specify an empty list of descriptors in the `members` attribute and set the `mode` member to `overwrite`.
	Mode pulumi.StringPtrInput
}

func (GroupMembershipArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupMembershipArgs)(nil)).Elem()
}

type GroupMembershipInput interface {
	pulumi.Input

	ToGroupMembershipOutput() GroupMembershipOutput
	ToGroupMembershipOutputWithContext(ctx context.Context) GroupMembershipOutput
}

func (*GroupMembership) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupMembership)(nil)).Elem()
}

func (i *GroupMembership) ToGroupMembershipOutput() GroupMembershipOutput {
	return i.ToGroupMembershipOutputWithContext(context.Background())
}

func (i *GroupMembership) ToGroupMembershipOutputWithContext(ctx context.Context) GroupMembershipOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupMembershipOutput)
}

// GroupMembershipArrayInput is an input type that accepts GroupMembershipArray and GroupMembershipArrayOutput values.
// You can construct a concrete instance of `GroupMembershipArrayInput` via:
//
//	GroupMembershipArray{ GroupMembershipArgs{...} }
type GroupMembershipArrayInput interface {
	pulumi.Input

	ToGroupMembershipArrayOutput() GroupMembershipArrayOutput
	ToGroupMembershipArrayOutputWithContext(context.Context) GroupMembershipArrayOutput
}

type GroupMembershipArray []GroupMembershipInput

func (GroupMembershipArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GroupMembership)(nil)).Elem()
}

func (i GroupMembershipArray) ToGroupMembershipArrayOutput() GroupMembershipArrayOutput {
	return i.ToGroupMembershipArrayOutputWithContext(context.Background())
}

func (i GroupMembershipArray) ToGroupMembershipArrayOutputWithContext(ctx context.Context) GroupMembershipArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupMembershipArrayOutput)
}

// GroupMembershipMapInput is an input type that accepts GroupMembershipMap and GroupMembershipMapOutput values.
// You can construct a concrete instance of `GroupMembershipMapInput` via:
//
//	GroupMembershipMap{ "key": GroupMembershipArgs{...} }
type GroupMembershipMapInput interface {
	pulumi.Input

	ToGroupMembershipMapOutput() GroupMembershipMapOutput
	ToGroupMembershipMapOutputWithContext(context.Context) GroupMembershipMapOutput
}

type GroupMembershipMap map[string]GroupMembershipInput

func (GroupMembershipMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GroupMembership)(nil)).Elem()
}

func (i GroupMembershipMap) ToGroupMembershipMapOutput() GroupMembershipMapOutput {
	return i.ToGroupMembershipMapOutputWithContext(context.Background())
}

func (i GroupMembershipMap) ToGroupMembershipMapOutputWithContext(ctx context.Context) GroupMembershipMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupMembershipMapOutput)
}

type GroupMembershipOutput struct{ *pulumi.OutputState }

func (GroupMembershipOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupMembership)(nil)).Elem()
}

func (o GroupMembershipOutput) ToGroupMembershipOutput() GroupMembershipOutput {
	return o
}

func (o GroupMembershipOutput) ToGroupMembershipOutputWithContext(ctx context.Context) GroupMembershipOutput {
	return o
}

// The descriptor of the group being managed.
func (o GroupMembershipOutput) Group() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupMembership) pulumi.StringOutput { return v.Group }).(pulumi.StringOutput)
}

// A list of user or group descriptors that will become members of the group.
// > NOTE: It's possible to define group members both within the `GroupMembership resource` via the members block and by using the `Group` resource. However it's not possible to use both methods to manage group members, since there'll be conflicts.
func (o GroupMembershipOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GroupMembership) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

// The mode how the resource manages group members.
// - `mode == add`: the resource will ensure that all specified members will be part of the referenced group
// - `mode == overwrite`: the resource will replace all existing members with the members specified within the `members` block
// > NOTE: To clear all members from a group, specify an empty list of descriptors in the `members` attribute and set the `mode` member to `overwrite`.
func (o GroupMembershipOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GroupMembership) pulumi.StringPtrOutput { return v.Mode }).(pulumi.StringPtrOutput)
}

type GroupMembershipArrayOutput struct{ *pulumi.OutputState }

func (GroupMembershipArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GroupMembership)(nil)).Elem()
}

func (o GroupMembershipArrayOutput) ToGroupMembershipArrayOutput() GroupMembershipArrayOutput {
	return o
}

func (o GroupMembershipArrayOutput) ToGroupMembershipArrayOutputWithContext(ctx context.Context) GroupMembershipArrayOutput {
	return o
}

func (o GroupMembershipArrayOutput) Index(i pulumi.IntInput) GroupMembershipOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GroupMembership {
		return vs[0].([]*GroupMembership)[vs[1].(int)]
	}).(GroupMembershipOutput)
}

type GroupMembershipMapOutput struct{ *pulumi.OutputState }

func (GroupMembershipMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GroupMembership)(nil)).Elem()
}

func (o GroupMembershipMapOutput) ToGroupMembershipMapOutput() GroupMembershipMapOutput {
	return o
}

func (o GroupMembershipMapOutput) ToGroupMembershipMapOutputWithContext(ctx context.Context) GroupMembershipMapOutput {
	return o
}

func (o GroupMembershipMapOutput) MapIndex(k pulumi.StringInput) GroupMembershipOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GroupMembership {
		return vs[0].(map[string]*GroupMembership)[vs[1].(string)]
	}).(GroupMembershipOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GroupMembershipInput)(nil)).Elem(), &GroupMembership{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupMembershipArrayInput)(nil)).Elem(), GroupMembershipArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupMembershipMapInput)(nil)).Elem(), GroupMembershipMap{})
	pulumi.RegisterOutputType(GroupMembershipOutput{})
	pulumi.RegisterOutputType(GroupMembershipArrayOutput{})
	pulumi.RegisterOutputType(GroupMembershipMapOutput{})
}
