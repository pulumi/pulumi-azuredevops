// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages group membership within Azure DevOps.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azuredevops/sdk/go/azuredevops"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		project, err := azuredevops.NewProject(ctx, "project", nil)
// 		if err != nil {
// 			return err
// 		}
// 		user, err := azuredevops.NewUser(ctx, "user", &azuredevops.UserArgs{
// 			PrincipalName: pulumi.String("foo@contoso.com"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = azuredevops.NewGroupMembership(ctx, "membership", &azuredevops.GroupMembershipArgs{
// 			Group: group.ApplyT(func(group azuredevops.LookupGroupResult) (string, error) {
// 				return group.Descriptor, nil
// 			}).(pulumi.StringOutput),
// 			Members: pulumi.StringArray{
// 				user.Descriptor,
// 			},
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
// - [Azure DevOps Service REST API 5.1 - Memberships](https://docs.microsoft.com/en-us/rest/api/azure/devops/graph/memberships?view=azure-devops-rest-5.0)
//
// ## PAT Permissions Required
//
// - **Deployment Groups**: Read & Manage
//
// ## Import
//
// Not supported.
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
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azuredevops:Identities/groupMembership:GroupMembership"),
		},
	})
	opts = append(opts, aliases)
	var resource GroupMembership
	err := ctx.RegisterResource("azuredevops:index/groupMembership:GroupMembership", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azuredevops:index/groupMembership:GroupMembership", name, id, state, &resource, opts...)
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

func (GroupMembership) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupMembership)(nil)).Elem()
}

func (i GroupMembership) ToGroupMembershipOutput() GroupMembershipOutput {
	return i.ToGroupMembershipOutputWithContext(context.Background())
}

func (i GroupMembership) ToGroupMembershipOutputWithContext(ctx context.Context) GroupMembershipOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupMembershipOutput)
}

type GroupMembershipOutput struct {
	*pulumi.OutputState
}

func (GroupMembershipOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupMembershipOutput)(nil)).Elem()
}

func (o GroupMembershipOutput) ToGroupMembershipOutput() GroupMembershipOutput {
	return o
}

func (o GroupMembershipOutput) ToGroupMembershipOutputWithContext(ctx context.Context) GroupMembershipOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(GroupMembershipOutput{})
}
