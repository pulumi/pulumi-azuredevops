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

// Manages members of a team within a project in a Azure DevOps organization.
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
//			exampleProject, err := azuredevops.NewProject(ctx, "exampleProject", &azuredevops.ProjectArgs{
//				WorkItemTemplate: pulumi.String("Agile"),
//				VersionControl:   pulumi.String("Git"),
//				Visibility:       pulumi.String("private"),
//				Description:      pulumi.String("Managed by Terraform"),
//			})
//			if err != nil {
//				return err
//			}
//			example_project_readers := azuredevops.LookupGroupOutput(ctx, azuredevops.GetGroupOutputArgs{
//				ProjectId: exampleProject.ID(),
//				Name:      pulumi.String("Readers"),
//			}, nil)
//			exampleTeam, err := azuredevops.NewTeam(ctx, "exampleTeam", &azuredevops.TeamArgs{
//				ProjectId: exampleProject.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuredevops.NewTeamMembers(ctx, "example-team-members", &azuredevops.TeamMembersArgs{
//				ProjectId: exampleTeam.ProjectId,
//				TeamId:    exampleTeam.ID(),
//				Mode:      pulumi.String("overwrite"),
//				Members: pulumi.StringArray{
//					example_project_readers.ApplyT(func(example_project_readers azuredevops.GetGroupResult) (*string, error) {
//						return &example_project_readers.Descriptor, nil
//					}).(pulumi.StringPtrOutput),
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
// - [Azure DevOps Service REST API 6.0 - Teams - Update](https://docs.microsoft.com/en-us/rest/api/azure/devops/core/teams/update?view=azure-devops-rest-6.0)
//
// ## PAT Permissions Required
//
// - **vso.project_write**:	Grants the ability to read and update projects and teams.
//
// ## Import
//
// The resource does not support import.
type TeamMembers struct {
	pulumi.CustomResourceState

	// List of subject descriptors to define members of the team.
	//
	// > NOTE: It's possible to define team members both within the
	// `Team` resource via the `members` block and by using the
	// `TeamMembers` resource. However it's not possible to use
	// both methods to manage team members, since there'll be conflicts.
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// The mode how the resource manages team members.
	// - `mode == add`: the resource will ensure that all specified members will be part of the referenced team
	// - `mode == overwrite`: the resource will replace all existing members with the members specified within the `members` block
	Mode pulumi.StringPtrOutput `pulumi:"mode"`
	// The Project ID.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The ID of the Team.
	TeamId pulumi.StringOutput `pulumi:"teamId"`
}

// NewTeamMembers registers a new resource with the given unique name, arguments, and options.
func NewTeamMembers(ctx *pulumi.Context,
	name string, args *TeamMembersArgs, opts ...pulumi.ResourceOption) (*TeamMembers, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.TeamId == nil {
		return nil, errors.New("invalid value for required argument 'TeamId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource TeamMembers
	err := ctx.RegisterResource("azuredevops:index/teamMembers:TeamMembers", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTeamMembers gets an existing TeamMembers resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTeamMembers(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TeamMembersState, opts ...pulumi.ResourceOption) (*TeamMembers, error) {
	var resource TeamMembers
	err := ctx.ReadResource("azuredevops:index/teamMembers:TeamMembers", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TeamMembers resources.
type teamMembersState struct {
	// List of subject descriptors to define members of the team.
	//
	// > NOTE: It's possible to define team members both within the
	// `Team` resource via the `members` block and by using the
	// `TeamMembers` resource. However it's not possible to use
	// both methods to manage team members, since there'll be conflicts.
	Members []string `pulumi:"members"`
	// The mode how the resource manages team members.
	// - `mode == add`: the resource will ensure that all specified members will be part of the referenced team
	// - `mode == overwrite`: the resource will replace all existing members with the members specified within the `members` block
	Mode *string `pulumi:"mode"`
	// The Project ID.
	ProjectId *string `pulumi:"projectId"`
	// The ID of the Team.
	TeamId *string `pulumi:"teamId"`
}

type TeamMembersState struct {
	// List of subject descriptors to define members of the team.
	//
	// > NOTE: It's possible to define team members both within the
	// `Team` resource via the `members` block and by using the
	// `TeamMembers` resource. However it's not possible to use
	// both methods to manage team members, since there'll be conflicts.
	Members pulumi.StringArrayInput
	// The mode how the resource manages team members.
	// - `mode == add`: the resource will ensure that all specified members will be part of the referenced team
	// - `mode == overwrite`: the resource will replace all existing members with the members specified within the `members` block
	Mode pulumi.StringPtrInput
	// The Project ID.
	ProjectId pulumi.StringPtrInput
	// The ID of the Team.
	TeamId pulumi.StringPtrInput
}

func (TeamMembersState) ElementType() reflect.Type {
	return reflect.TypeOf((*teamMembersState)(nil)).Elem()
}

type teamMembersArgs struct {
	// List of subject descriptors to define members of the team.
	//
	// > NOTE: It's possible to define team members both within the
	// `Team` resource via the `members` block and by using the
	// `TeamMembers` resource. However it's not possible to use
	// both methods to manage team members, since there'll be conflicts.
	Members []string `pulumi:"members"`
	// The mode how the resource manages team members.
	// - `mode == add`: the resource will ensure that all specified members will be part of the referenced team
	// - `mode == overwrite`: the resource will replace all existing members with the members specified within the `members` block
	Mode *string `pulumi:"mode"`
	// The Project ID.
	ProjectId string `pulumi:"projectId"`
	// The ID of the Team.
	TeamId string `pulumi:"teamId"`
}

// The set of arguments for constructing a TeamMembers resource.
type TeamMembersArgs struct {
	// List of subject descriptors to define members of the team.
	//
	// > NOTE: It's possible to define team members both within the
	// `Team` resource via the `members` block and by using the
	// `TeamMembers` resource. However it's not possible to use
	// both methods to manage team members, since there'll be conflicts.
	Members pulumi.StringArrayInput
	// The mode how the resource manages team members.
	// - `mode == add`: the resource will ensure that all specified members will be part of the referenced team
	// - `mode == overwrite`: the resource will replace all existing members with the members specified within the `members` block
	Mode pulumi.StringPtrInput
	// The Project ID.
	ProjectId pulumi.StringInput
	// The ID of the Team.
	TeamId pulumi.StringInput
}

func (TeamMembersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*teamMembersArgs)(nil)).Elem()
}

type TeamMembersInput interface {
	pulumi.Input

	ToTeamMembersOutput() TeamMembersOutput
	ToTeamMembersOutputWithContext(ctx context.Context) TeamMembersOutput
}

func (*TeamMembers) ElementType() reflect.Type {
	return reflect.TypeOf((**TeamMembers)(nil)).Elem()
}

func (i *TeamMembers) ToTeamMembersOutput() TeamMembersOutput {
	return i.ToTeamMembersOutputWithContext(context.Background())
}

func (i *TeamMembers) ToTeamMembersOutputWithContext(ctx context.Context) TeamMembersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TeamMembersOutput)
}

// TeamMembersArrayInput is an input type that accepts TeamMembersArray and TeamMembersArrayOutput values.
// You can construct a concrete instance of `TeamMembersArrayInput` via:
//
//	TeamMembersArray{ TeamMembersArgs{...} }
type TeamMembersArrayInput interface {
	pulumi.Input

	ToTeamMembersArrayOutput() TeamMembersArrayOutput
	ToTeamMembersArrayOutputWithContext(context.Context) TeamMembersArrayOutput
}

type TeamMembersArray []TeamMembersInput

func (TeamMembersArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TeamMembers)(nil)).Elem()
}

func (i TeamMembersArray) ToTeamMembersArrayOutput() TeamMembersArrayOutput {
	return i.ToTeamMembersArrayOutputWithContext(context.Background())
}

func (i TeamMembersArray) ToTeamMembersArrayOutputWithContext(ctx context.Context) TeamMembersArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TeamMembersArrayOutput)
}

// TeamMembersMapInput is an input type that accepts TeamMembersMap and TeamMembersMapOutput values.
// You can construct a concrete instance of `TeamMembersMapInput` via:
//
//	TeamMembersMap{ "key": TeamMembersArgs{...} }
type TeamMembersMapInput interface {
	pulumi.Input

	ToTeamMembersMapOutput() TeamMembersMapOutput
	ToTeamMembersMapOutputWithContext(context.Context) TeamMembersMapOutput
}

type TeamMembersMap map[string]TeamMembersInput

func (TeamMembersMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TeamMembers)(nil)).Elem()
}

func (i TeamMembersMap) ToTeamMembersMapOutput() TeamMembersMapOutput {
	return i.ToTeamMembersMapOutputWithContext(context.Background())
}

func (i TeamMembersMap) ToTeamMembersMapOutputWithContext(ctx context.Context) TeamMembersMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TeamMembersMapOutput)
}

type TeamMembersOutput struct{ *pulumi.OutputState }

func (TeamMembersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TeamMembers)(nil)).Elem()
}

func (o TeamMembersOutput) ToTeamMembersOutput() TeamMembersOutput {
	return o
}

func (o TeamMembersOutput) ToTeamMembersOutputWithContext(ctx context.Context) TeamMembersOutput {
	return o
}

// List of subject descriptors to define members of the team.
//
// > NOTE: It's possible to define team members both within the
// `Team` resource via the `members` block and by using the
// `TeamMembers` resource. However it's not possible to use
// both methods to manage team members, since there'll be conflicts.
func (o TeamMembersOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TeamMembers) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

// The mode how the resource manages team members.
// - `mode == add`: the resource will ensure that all specified members will be part of the referenced team
// - `mode == overwrite`: the resource will replace all existing members with the members specified within the `members` block
func (o TeamMembersOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TeamMembers) pulumi.StringPtrOutput { return v.Mode }).(pulumi.StringPtrOutput)
}

// The Project ID.
func (o TeamMembersOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *TeamMembers) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The ID of the Team.
func (o TeamMembersOutput) TeamId() pulumi.StringOutput {
	return o.ApplyT(func(v *TeamMembers) pulumi.StringOutput { return v.TeamId }).(pulumi.StringOutput)
}

type TeamMembersArrayOutput struct{ *pulumi.OutputState }

func (TeamMembersArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TeamMembers)(nil)).Elem()
}

func (o TeamMembersArrayOutput) ToTeamMembersArrayOutput() TeamMembersArrayOutput {
	return o
}

func (o TeamMembersArrayOutput) ToTeamMembersArrayOutputWithContext(ctx context.Context) TeamMembersArrayOutput {
	return o
}

func (o TeamMembersArrayOutput) Index(i pulumi.IntInput) TeamMembersOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TeamMembers {
		return vs[0].([]*TeamMembers)[vs[1].(int)]
	}).(TeamMembersOutput)
}

type TeamMembersMapOutput struct{ *pulumi.OutputState }

func (TeamMembersMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TeamMembers)(nil)).Elem()
}

func (o TeamMembersMapOutput) ToTeamMembersMapOutput() TeamMembersMapOutput {
	return o
}

func (o TeamMembersMapOutput) ToTeamMembersMapOutputWithContext(ctx context.Context) TeamMembersMapOutput {
	return o
}

func (o TeamMembersMapOutput) MapIndex(k pulumi.StringInput) TeamMembersOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TeamMembers {
		return vs[0].(map[string]*TeamMembers)[vs[1].(string)]
	}).(TeamMembersOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TeamMembersInput)(nil)).Elem(), &TeamMembers{})
	pulumi.RegisterInputType(reflect.TypeOf((*TeamMembersArrayInput)(nil)).Elem(), TeamMembersArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TeamMembersMapInput)(nil)).Elem(), TeamMembersMap{})
	pulumi.RegisterOutputType(TeamMembersOutput{})
	pulumi.RegisterOutputType(TeamMembersArrayOutput{})
	pulumi.RegisterOutputType(TeamMembersMapOutput{})
}
