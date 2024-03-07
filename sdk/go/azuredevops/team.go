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

// Manages a team within a project in a Azure DevOps organization.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
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
//			example_project_contributors := azuredevops.LookupGroupOutput(ctx, azuredevops.GetGroupOutputArgs{
//				ProjectId: exampleProject.ID(),
//				Name:      pulumi.String("Contributors"),
//			}, nil)
//			example_project_readers := azuredevops.LookupGroupOutput(ctx, azuredevops.GetGroupOutputArgs{
//				ProjectId: exampleProject.ID(),
//				Name:      pulumi.String("Readers"),
//			}, nil)
//			_, err = azuredevops.NewTeam(ctx, "exampleTeam", &azuredevops.TeamArgs{
//				ProjectId: exampleProject.ID(),
//				Administrators: pulumi.StringArray{
//					example_project_contributors.ApplyT(func(example_project_contributors azuredevops.GetGroupResult) (*string, error) {
//						return &example_project_contributors.Descriptor, nil
//					}).(pulumi.StringPtrOutput),
//				},
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
// <!--End PulumiCodeChooser -->
//
// ## Relevant Links
//
// - [Azure DevOps Service REST API 7.0 - Teams - Create](https://docs.microsoft.com/en-us/rest/api/azure/devops/core/teams/create?view=azure-devops-rest-7.0)
//
// ## PAT Permissions Required
//
// - **vso.project_manage**:	Grants the ability to create, read, update, and delete projects and teams.
//
// ## Import
//
// Azure DevOps teams can be imported using the complete resource id `<project_id>/<team_id>` e.g.
//
// ```sh
// $ pulumi import azuredevops:index/team:Team example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
// ```
type Team struct {
	pulumi.CustomResourceState

	// List of subject descriptors to define administrators of the team.
	//
	// > NOTE: It's possible to define team administrators both within the
	// `Team` resource via the `administrators` block and by using the
	// `TeamAdministrators` resource. However it's not possible to use
	// both methods to manage team administrators, since there'll be conflicts.
	Administrators pulumi.StringArrayOutput `pulumi:"administrators"`
	// The description of the Team.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The descriptor of the Team.
	Descriptor pulumi.StringOutput `pulumi:"descriptor"`
	// List of subject descriptors to define members of the team.
	//
	// > NOTE: It's possible to define team members both within the
	// `Team` resource via the `members` block and by using the
	// `TeamMembers` resource. However it's not possible to use
	// both methods to manage team members, since there'll be conflicts.
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// The name of the Team.
	Name pulumi.StringOutput `pulumi:"name"`
	// The Project ID.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
}

// NewTeam registers a new resource with the given unique name, arguments, and options.
func NewTeam(ctx *pulumi.Context,
	name string, args *TeamArgs, opts ...pulumi.ResourceOption) (*Team, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Team
	err := ctx.RegisterResource("azuredevops:index/team:Team", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTeam gets an existing Team resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTeam(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TeamState, opts ...pulumi.ResourceOption) (*Team, error) {
	var resource Team
	err := ctx.ReadResource("azuredevops:index/team:Team", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Team resources.
type teamState struct {
	// List of subject descriptors to define administrators of the team.
	//
	// > NOTE: It's possible to define team administrators both within the
	// `Team` resource via the `administrators` block and by using the
	// `TeamAdministrators` resource. However it's not possible to use
	// both methods to manage team administrators, since there'll be conflicts.
	Administrators []string `pulumi:"administrators"`
	// The description of the Team.
	Description *string `pulumi:"description"`
	// The descriptor of the Team.
	Descriptor *string `pulumi:"descriptor"`
	// List of subject descriptors to define members of the team.
	//
	// > NOTE: It's possible to define team members both within the
	// `Team` resource via the `members` block and by using the
	// `TeamMembers` resource. However it's not possible to use
	// both methods to manage team members, since there'll be conflicts.
	Members []string `pulumi:"members"`
	// The name of the Team.
	Name *string `pulumi:"name"`
	// The Project ID.
	ProjectId *string `pulumi:"projectId"`
}

type TeamState struct {
	// List of subject descriptors to define administrators of the team.
	//
	// > NOTE: It's possible to define team administrators both within the
	// `Team` resource via the `administrators` block and by using the
	// `TeamAdministrators` resource. However it's not possible to use
	// both methods to manage team administrators, since there'll be conflicts.
	Administrators pulumi.StringArrayInput
	// The description of the Team.
	Description pulumi.StringPtrInput
	// The descriptor of the Team.
	Descriptor pulumi.StringPtrInput
	// List of subject descriptors to define members of the team.
	//
	// > NOTE: It's possible to define team members both within the
	// `Team` resource via the `members` block and by using the
	// `TeamMembers` resource. However it's not possible to use
	// both methods to manage team members, since there'll be conflicts.
	Members pulumi.StringArrayInput
	// The name of the Team.
	Name pulumi.StringPtrInput
	// The Project ID.
	ProjectId pulumi.StringPtrInput
}

func (TeamState) ElementType() reflect.Type {
	return reflect.TypeOf((*teamState)(nil)).Elem()
}

type teamArgs struct {
	// List of subject descriptors to define administrators of the team.
	//
	// > NOTE: It's possible to define team administrators both within the
	// `Team` resource via the `administrators` block and by using the
	// `TeamAdministrators` resource. However it's not possible to use
	// both methods to manage team administrators, since there'll be conflicts.
	Administrators []string `pulumi:"administrators"`
	// The description of the Team.
	Description *string `pulumi:"description"`
	// List of subject descriptors to define members of the team.
	//
	// > NOTE: It's possible to define team members both within the
	// `Team` resource via the `members` block and by using the
	// `TeamMembers` resource. However it's not possible to use
	// both methods to manage team members, since there'll be conflicts.
	Members []string `pulumi:"members"`
	// The name of the Team.
	Name *string `pulumi:"name"`
	// The Project ID.
	ProjectId string `pulumi:"projectId"`
}

// The set of arguments for constructing a Team resource.
type TeamArgs struct {
	// List of subject descriptors to define administrators of the team.
	//
	// > NOTE: It's possible to define team administrators both within the
	// `Team` resource via the `administrators` block and by using the
	// `TeamAdministrators` resource. However it's not possible to use
	// both methods to manage team administrators, since there'll be conflicts.
	Administrators pulumi.StringArrayInput
	// The description of the Team.
	Description pulumi.StringPtrInput
	// List of subject descriptors to define members of the team.
	//
	// > NOTE: It's possible to define team members both within the
	// `Team` resource via the `members` block and by using the
	// `TeamMembers` resource. However it's not possible to use
	// both methods to manage team members, since there'll be conflicts.
	Members pulumi.StringArrayInput
	// The name of the Team.
	Name pulumi.StringPtrInput
	// The Project ID.
	ProjectId pulumi.StringInput
}

func (TeamArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*teamArgs)(nil)).Elem()
}

type TeamInput interface {
	pulumi.Input

	ToTeamOutput() TeamOutput
	ToTeamOutputWithContext(ctx context.Context) TeamOutput
}

func (*Team) ElementType() reflect.Type {
	return reflect.TypeOf((**Team)(nil)).Elem()
}

func (i *Team) ToTeamOutput() TeamOutput {
	return i.ToTeamOutputWithContext(context.Background())
}

func (i *Team) ToTeamOutputWithContext(ctx context.Context) TeamOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TeamOutput)
}

// TeamArrayInput is an input type that accepts TeamArray and TeamArrayOutput values.
// You can construct a concrete instance of `TeamArrayInput` via:
//
//	TeamArray{ TeamArgs{...} }
type TeamArrayInput interface {
	pulumi.Input

	ToTeamArrayOutput() TeamArrayOutput
	ToTeamArrayOutputWithContext(context.Context) TeamArrayOutput
}

type TeamArray []TeamInput

func (TeamArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Team)(nil)).Elem()
}

func (i TeamArray) ToTeamArrayOutput() TeamArrayOutput {
	return i.ToTeamArrayOutputWithContext(context.Background())
}

func (i TeamArray) ToTeamArrayOutputWithContext(ctx context.Context) TeamArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TeamArrayOutput)
}

// TeamMapInput is an input type that accepts TeamMap and TeamMapOutput values.
// You can construct a concrete instance of `TeamMapInput` via:
//
//	TeamMap{ "key": TeamArgs{...} }
type TeamMapInput interface {
	pulumi.Input

	ToTeamMapOutput() TeamMapOutput
	ToTeamMapOutputWithContext(context.Context) TeamMapOutput
}

type TeamMap map[string]TeamInput

func (TeamMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Team)(nil)).Elem()
}

func (i TeamMap) ToTeamMapOutput() TeamMapOutput {
	return i.ToTeamMapOutputWithContext(context.Background())
}

func (i TeamMap) ToTeamMapOutputWithContext(ctx context.Context) TeamMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TeamMapOutput)
}

type TeamOutput struct{ *pulumi.OutputState }

func (TeamOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Team)(nil)).Elem()
}

func (o TeamOutput) ToTeamOutput() TeamOutput {
	return o
}

func (o TeamOutput) ToTeamOutputWithContext(ctx context.Context) TeamOutput {
	return o
}

// List of subject descriptors to define administrators of the team.
//
// > NOTE: It's possible to define team administrators both within the
// `Team` resource via the `administrators` block and by using the
// `TeamAdministrators` resource. However it's not possible to use
// both methods to manage team administrators, since there'll be conflicts.
func (o TeamOutput) Administrators() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Team) pulumi.StringArrayOutput { return v.Administrators }).(pulumi.StringArrayOutput)
}

// The description of the Team.
func (o TeamOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Team) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The descriptor of the Team.
func (o TeamOutput) Descriptor() pulumi.StringOutput {
	return o.ApplyT(func(v *Team) pulumi.StringOutput { return v.Descriptor }).(pulumi.StringOutput)
}

// List of subject descriptors to define members of the team.
//
// > NOTE: It's possible to define team members both within the
// `Team` resource via the `members` block and by using the
// `TeamMembers` resource. However it's not possible to use
// both methods to manage team members, since there'll be conflicts.
func (o TeamOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Team) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

// The name of the Team.
func (o TeamOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Team) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The Project ID.
func (o TeamOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *Team) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

type TeamArrayOutput struct{ *pulumi.OutputState }

func (TeamArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Team)(nil)).Elem()
}

func (o TeamArrayOutput) ToTeamArrayOutput() TeamArrayOutput {
	return o
}

func (o TeamArrayOutput) ToTeamArrayOutputWithContext(ctx context.Context) TeamArrayOutput {
	return o
}

func (o TeamArrayOutput) Index(i pulumi.IntInput) TeamOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Team {
		return vs[0].([]*Team)[vs[1].(int)]
	}).(TeamOutput)
}

type TeamMapOutput struct{ *pulumi.OutputState }

func (TeamMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Team)(nil)).Elem()
}

func (o TeamMapOutput) ToTeamMapOutput() TeamMapOutput {
	return o
}

func (o TeamMapOutput) ToTeamMapOutputWithContext(ctx context.Context) TeamMapOutput {
	return o
}

func (o TeamMapOutput) MapIndex(k pulumi.StringInput) TeamOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Team {
		return vs[0].(map[string]*Team)[vs[1].(string)]
	}).(TeamOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TeamInput)(nil)).Elem(), &Team{})
	pulumi.RegisterInputType(reflect.TypeOf((*TeamArrayInput)(nil)).Elem(), TeamArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TeamMapInput)(nil)).Elem(), TeamMap{})
	pulumi.RegisterOutputType(TeamOutput{})
	pulumi.RegisterOutputType(TeamArrayOutput{})
	pulumi.RegisterOutputType(TeamMapOutput{})
}
