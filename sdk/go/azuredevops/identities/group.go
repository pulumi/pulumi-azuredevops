// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package identities

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a group within Azure DevOps.
//
// ## Relevant Links
//
// * [Azure DevOps Service REST API 5.1 - Groups](https://docs.microsoft.com/en-us/rest/api/azure/devops/graph/groups?view=azure-devops-rest-5.1)
//
// ## PAT Permissions Required
//
// - **Project & Team**: Read, Write, & Manage
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
	// > NOTE: It's possible to define group members both within the `Identities.Group` resource via the members block and by using the   `Identities.GroupMembership` resource. However it's not possible to use both methods to manage group members, since there'll be conflicts.
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
	// > NOTE: It's possible to define group members both within the `Identities.Group` resource via the members block and by using the   `Identities.GroupMembership` resource. However it's not possible to use both methods to manage group members, since there'll be conflicts.
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
	// > NOTE: It's possible to define group members both within the `Identities.Group` resource via the members block and by using the   `Identities.GroupMembership` resource. However it's not possible to use both methods to manage group members, since there'll be conflicts.
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
	// > NOTE: It's possible to define group members both within the `Identities.Group` resource via the members block and by using the   `Identities.GroupMembership` resource. However it's not possible to use both methods to manage group members, since there'll be conflicts.
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
	// > NOTE: It's possible to define group members both within the `Identities.Group` resource via the members block and by using the   `Identities.GroupMembership` resource. However it's not possible to use both methods to manage group members, since there'll be conflicts.
	Members pulumi.StringArrayInput
	// The OriginID as a reference to a group from an external AD or AAD backed provider. The `scope`, `mail` and `displayName` arguments cannot be used simultaneously with `originId`.
	OriginId pulumi.StringPtrInput
	// The scope of the group. A descriptor referencing the scope (collection, project) in which the group should be created. If omitted, will be created in the scope of the enclosing account or organization.x
	Scope pulumi.StringPtrInput
}

func (GroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupArgs)(nil)).Elem()
}
