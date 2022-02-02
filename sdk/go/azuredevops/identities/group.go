// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package identities

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a group within Azure DevOps.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azuredevops/sdk/v2/go/azuredevops"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		project, err := azuredevops.NewProject(ctx, "project", nil)
// 		if err != nil {
// 			return err
// 		}
// 		tf_project_readers := azuredevops.LookupGroupOutput(ctx, GetGroupOutputArgs{
// 			ProjectId: project.ID(),
// 			Name:      pulumi.String("Readers"),
// 		}, nil)
// 		tf_project_contributors := azuredevops.LookupGroupOutput(ctx, GetGroupOutputArgs{
// 			ProjectId: project.ID(),
// 			Name:      pulumi.String("Contributors"),
// 		}, nil)
// 		_, err = azuredevops.NewGroup(ctx, "group", &azuredevops.GroupArgs{
// 			Scope:       project.ID(),
// 			DisplayName: pulumi.String("Test group"),
// 			Description: pulumi.String("Test description"),
// 			Members: pulumi.StringArray{
// 				tf_project_readers.ApplyT(func(tf_project_readers GetGroupResult) (string, error) {
// 					return tf_project_readers.Descriptor, nil
// 				}).(pulumi.StringOutput),
// 				tf_project_contributors.ApplyT(func(tf_project_contributors GetGroupResult) (string, error) {
// 					return tf_project_contributors.Descriptor, nil
// 				}).(pulumi.StringOutput),
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
// - [Azure DevOps Service REST API 5.1 - Groups](https://docs.microsoft.com/en-us/rest/api/azure/devops/graph/groups?view=azure-devops-rest-5.1)
//
// ## PAT Permissions Required
//
// - **Project & Team**: Read, Write, & Manage
//
// ## Import
//
// Azure DevOps groups can be imported using the group identity descriptor, e.g.
//
// ```sh
//  $ pulumi import azuredevops:Identities/group:Group id aadgp.Uy0xLTktMTU1MTM3NDI0NS0xMjA0NDAwOTY5LTI0MDI5ODY0MTMtMjE3OTQwODYxNi0zLTIxNjc2NjQyNTMtMzI1Nzg0NDI4OS0yMjU4MjcwOTc0LTI2MDYxODY2NDU
// ```
//
// Deprecated: azuredevops.identities.Group has been deprecated in favor of azuredevops.Group
type Group struct {
	pulumi.CustomResourceState

	// The Description of the Project.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The identity (subject) descriptor of the Group.
	Descriptor pulumi.StringOutput `pulumi:"descriptor"`
	// The name of a new Azure DevOps group that is not backed by an external provider. The `originId` and `mail` arguments cannot be used simultaneously with `displayName`.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// This represents the name of the container of origin for a graph member.
	Domain pulumi.StringOutput `pulumi:"domain"`
	// The mail address as a reference to an existing group from an external AD or AAD backed provider. The `scope`, `originId` and `displayName` arguments cannot be used simultaneously with `mail`.
	Mail pulumi.StringOutput `pulumi:"mail"`
	// > NOTE: It's possible to define group members both within the `Group` resource via the members block and by using the `GroupMembership` resource. However it's not possible to use both methods to manage group members, since there'll be conflicts.
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// The type of source provider for the origin identifier (ex:AD, AAD, MSA)
	Origin pulumi.StringOutput `pulumi:"origin"`
	// The OriginID as a reference to a group from an external AD or AAD backed provider. The `scope`, `mail` and `displayName` arguments cannot be used simultaneously with `originId`.
	OriginId pulumi.StringOutput `pulumi:"originId"`
	// This is the PrincipalName of this graph member from the source provider.
	PrincipalName pulumi.StringOutput `pulumi:"principalName"`
	// The scope of the group. A descriptor referencing the scope (collection, project) in which the group should be created. If omitted, will be created in the scope of the enclosing account or organization.x
	Scope pulumi.StringPtrOutput `pulumi:"scope"`
	// This field identifies the type of the graph subject (ex: Group, Scope, User).
	SubjectKind pulumi.StringOutput `pulumi:"subjectKind"`
	// This url is the full route to the source resource of this graph subject.
	Url pulumi.StringOutput `pulumi:"url"`
}

// NewGroup registers a new resource with the given unique name, arguments, and options.
func NewGroup(ctx *pulumi.Context,
	name string, args *GroupArgs, opts ...pulumi.ResourceOption) (*Group, error) {
	if args == nil {
		args = &GroupArgs{}
	}

	var resource Group
	err := ctx.RegisterResource("azuredevops:Identities/group:Group", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroup gets an existing Group resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupState, opts ...pulumi.ResourceOption) (*Group, error) {
	var resource Group
	err := ctx.ReadResource("azuredevops:Identities/group:Group", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Group resources.
type groupState struct {
	// The Description of the Project.
	Description *string `pulumi:"description"`
	// The identity (subject) descriptor of the Group.
	Descriptor *string `pulumi:"descriptor"`
	// The name of a new Azure DevOps group that is not backed by an external provider. The `originId` and `mail` arguments cannot be used simultaneously with `displayName`.
	DisplayName *string `pulumi:"displayName"`
	// This represents the name of the container of origin for a graph member.
	Domain *string `pulumi:"domain"`
	// The mail address as a reference to an existing group from an external AD or AAD backed provider. The `scope`, `originId` and `displayName` arguments cannot be used simultaneously with `mail`.
	Mail *string `pulumi:"mail"`
	// > NOTE: It's possible to define group members both within the `Group` resource via the members block and by using the `GroupMembership` resource. However it's not possible to use both methods to manage group members, since there'll be conflicts.
	Members []string `pulumi:"members"`
	// The type of source provider for the origin identifier (ex:AD, AAD, MSA)
	Origin *string `pulumi:"origin"`
	// The OriginID as a reference to a group from an external AD or AAD backed provider. The `scope`, `mail` and `displayName` arguments cannot be used simultaneously with `originId`.
	OriginId *string `pulumi:"originId"`
	// This is the PrincipalName of this graph member from the source provider.
	PrincipalName *string `pulumi:"principalName"`
	// The scope of the group. A descriptor referencing the scope (collection, project) in which the group should be created. If omitted, will be created in the scope of the enclosing account or organization.x
	Scope *string `pulumi:"scope"`
	// This field identifies the type of the graph subject (ex: Group, Scope, User).
	SubjectKind *string `pulumi:"subjectKind"`
	// This url is the full route to the source resource of this graph subject.
	Url *string `pulumi:"url"`
}

type GroupState struct {
	// The Description of the Project.
	Description pulumi.StringPtrInput
	// The identity (subject) descriptor of the Group.
	Descriptor pulumi.StringPtrInput
	// The name of a new Azure DevOps group that is not backed by an external provider. The `originId` and `mail` arguments cannot be used simultaneously with `displayName`.
	DisplayName pulumi.StringPtrInput
	// This represents the name of the container of origin for a graph member.
	Domain pulumi.StringPtrInput
	// The mail address as a reference to an existing group from an external AD or AAD backed provider. The `scope`, `originId` and `displayName` arguments cannot be used simultaneously with `mail`.
	Mail pulumi.StringPtrInput
	// > NOTE: It's possible to define group members both within the `Group` resource via the members block and by using the `GroupMembership` resource. However it's not possible to use both methods to manage group members, since there'll be conflicts.
	Members pulumi.StringArrayInput
	// The type of source provider for the origin identifier (ex:AD, AAD, MSA)
	Origin pulumi.StringPtrInput
	// The OriginID as a reference to a group from an external AD or AAD backed provider. The `scope`, `mail` and `displayName` arguments cannot be used simultaneously with `originId`.
	OriginId pulumi.StringPtrInput
	// This is the PrincipalName of this graph member from the source provider.
	PrincipalName pulumi.StringPtrInput
	// The scope of the group. A descriptor referencing the scope (collection, project) in which the group should be created. If omitted, will be created in the scope of the enclosing account or organization.x
	Scope pulumi.StringPtrInput
	// This field identifies the type of the graph subject (ex: Group, Scope, User).
	SubjectKind pulumi.StringPtrInput
	// This url is the full route to the source resource of this graph subject.
	Url pulumi.StringPtrInput
}

func (GroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupState)(nil)).Elem()
}

type groupArgs struct {
	// The Description of the Project.
	Description *string `pulumi:"description"`
	// The name of a new Azure DevOps group that is not backed by an external provider. The `originId` and `mail` arguments cannot be used simultaneously with `displayName`.
	DisplayName *string `pulumi:"displayName"`
	// The mail address as a reference to an existing group from an external AD or AAD backed provider. The `scope`, `originId` and `displayName` arguments cannot be used simultaneously with `mail`.
	Mail *string `pulumi:"mail"`
	// > NOTE: It's possible to define group members both within the `Group` resource via the members block and by using the `GroupMembership` resource. However it's not possible to use both methods to manage group members, since there'll be conflicts.
	Members []string `pulumi:"members"`
	// The OriginID as a reference to a group from an external AD or AAD backed provider. The `scope`, `mail` and `displayName` arguments cannot be used simultaneously with `originId`.
	OriginId *string `pulumi:"originId"`
	// The scope of the group. A descriptor referencing the scope (collection, project) in which the group should be created. If omitted, will be created in the scope of the enclosing account or organization.x
	Scope *string `pulumi:"scope"`
}

// The set of arguments for constructing a Group resource.
type GroupArgs struct {
	// The Description of the Project.
	Description pulumi.StringPtrInput
	// The name of a new Azure DevOps group that is not backed by an external provider. The `originId` and `mail` arguments cannot be used simultaneously with `displayName`.
	DisplayName pulumi.StringPtrInput
	// The mail address as a reference to an existing group from an external AD or AAD backed provider. The `scope`, `originId` and `displayName` arguments cannot be used simultaneously with `mail`.
	Mail pulumi.StringPtrInput
	// > NOTE: It's possible to define group members both within the `Group` resource via the members block and by using the `GroupMembership` resource. However it's not possible to use both methods to manage group members, since there'll be conflicts.
	Members pulumi.StringArrayInput
	// The OriginID as a reference to a group from an external AD or AAD backed provider. The `scope`, `mail` and `displayName` arguments cannot be used simultaneously with `originId`.
	OriginId pulumi.StringPtrInput
	// The scope of the group. A descriptor referencing the scope (collection, project) in which the group should be created. If omitted, will be created in the scope of the enclosing account or organization.x
	Scope pulumi.StringPtrInput
}

func (GroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupArgs)(nil)).Elem()
}

type GroupInput interface {
	pulumi.Input

	ToGroupOutput() GroupOutput
	ToGroupOutputWithContext(ctx context.Context) GroupOutput
}

func (*Group) ElementType() reflect.Type {
	return reflect.TypeOf((**Group)(nil)).Elem()
}

func (i *Group) ToGroupOutput() GroupOutput {
	return i.ToGroupOutputWithContext(context.Background())
}

func (i *Group) ToGroupOutputWithContext(ctx context.Context) GroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupOutput)
}

// GroupArrayInput is an input type that accepts GroupArray and GroupArrayOutput values.
// You can construct a concrete instance of `GroupArrayInput` via:
//
//          GroupArray{ GroupArgs{...} }
type GroupArrayInput interface {
	pulumi.Input

	ToGroupArrayOutput() GroupArrayOutput
	ToGroupArrayOutputWithContext(context.Context) GroupArrayOutput
}

type GroupArray []GroupInput

func (GroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Group)(nil)).Elem()
}

func (i GroupArray) ToGroupArrayOutput() GroupArrayOutput {
	return i.ToGroupArrayOutputWithContext(context.Background())
}

func (i GroupArray) ToGroupArrayOutputWithContext(ctx context.Context) GroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupArrayOutput)
}

// GroupMapInput is an input type that accepts GroupMap and GroupMapOutput values.
// You can construct a concrete instance of `GroupMapInput` via:
//
//          GroupMap{ "key": GroupArgs{...} }
type GroupMapInput interface {
	pulumi.Input

	ToGroupMapOutput() GroupMapOutput
	ToGroupMapOutputWithContext(context.Context) GroupMapOutput
}

type GroupMap map[string]GroupInput

func (GroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Group)(nil)).Elem()
}

func (i GroupMap) ToGroupMapOutput() GroupMapOutput {
	return i.ToGroupMapOutputWithContext(context.Background())
}

func (i GroupMap) ToGroupMapOutputWithContext(ctx context.Context) GroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupMapOutput)
}

type GroupOutput struct{ *pulumi.OutputState }

func (GroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Group)(nil)).Elem()
}

func (o GroupOutput) ToGroupOutput() GroupOutput {
	return o
}

func (o GroupOutput) ToGroupOutputWithContext(ctx context.Context) GroupOutput {
	return o
}

type GroupArrayOutput struct{ *pulumi.OutputState }

func (GroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Group)(nil)).Elem()
}

func (o GroupArrayOutput) ToGroupArrayOutput() GroupArrayOutput {
	return o
}

func (o GroupArrayOutput) ToGroupArrayOutputWithContext(ctx context.Context) GroupArrayOutput {
	return o
}

func (o GroupArrayOutput) Index(i pulumi.IntInput) GroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Group {
		return vs[0].([]*Group)[vs[1].(int)]
	}).(GroupOutput)
}

type GroupMapOutput struct{ *pulumi.OutputState }

func (GroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Group)(nil)).Elem()
}

func (o GroupMapOutput) ToGroupMapOutput() GroupMapOutput {
	return o
}

func (o GroupMapOutput) ToGroupMapOutputWithContext(ctx context.Context) GroupMapOutput {
	return o
}

func (o GroupMapOutput) MapIndex(k pulumi.StringInput) GroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Group {
		return vs[0].(map[string]*Group)[vs[1].(string)]
	}).(GroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GroupInput)(nil)).Elem(), &Group{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupArrayInput)(nil)).Elem(), GroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupMapInput)(nil)).Elem(), GroupMap{})
	pulumi.RegisterOutputType(GroupOutput{})
	pulumi.RegisterOutputType(GroupArrayOutput{})
	pulumi.RegisterOutputType(GroupMapOutput{})
}
