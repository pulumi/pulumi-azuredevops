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

// Manages administrators of a team within a project in a Azure DevOps organization.
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
//			exampleTeam, err := azuredevops.NewTeam(ctx, "exampleTeam", &azuredevops.TeamArgs{
//				ProjectId: exampleProject.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuredevops.NewTeamAdministrators(ctx, "example-team-administrators", &azuredevops.TeamAdministratorsArgs{
//				ProjectId: exampleTeam.ProjectId,
//				TeamId:    exampleTeam.ID(),
//				Mode:      pulumi.String("overwrite"),
//				Administrators: pulumi.StringArray{
//					example_project_contributors.ApplyT(func(example_project_contributors azuredevops.GetGroupResult) (*string, error) {
//						return &example_project_contributors.Descriptor, nil
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
// - [Azure DevOps Service REST API 7.0 - Teams - Update](https://docs.microsoft.com/en-us/rest/api/azure/devops/core/teams/update?view=azure-devops-rest-7.0)
//
// ## PAT Permissions Required
//
// - **vso.project_write**:	Grants the ability to read and update projects and teams.
//
// ## Import
//
// The resource does not support import.
type TeamAdministrators struct {
	pulumi.CustomResourceState

	// List of subject descriptors to define adminitrators of the team.
	//
	// > NOTE: It's possible to define team administrators both within the
	// `Team` resource via the `administrators` block and by using the
	// `TeamAdministrators` resource. However it's not possible to use
	// both methods to manage team administrators, since there'll be conflicts.
	Administrators pulumi.StringArrayOutput `pulumi:"administrators"`
	// The mode how the resource manages team administrators.
	// - `mode == add`: the resource will ensure that all specified administrators will be part of the referenced team
	// - `mode == overwrite`: the resource will replace all existing administrators with the administrators specified within the `administrators` block
	Mode pulumi.StringPtrOutput `pulumi:"mode"`
	// The Project ID.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The ID of the Team.
	TeamId pulumi.StringOutput `pulumi:"teamId"`
}

// NewTeamAdministrators registers a new resource with the given unique name, arguments, and options.
func NewTeamAdministrators(ctx *pulumi.Context,
	name string, args *TeamAdministratorsArgs, opts ...pulumi.ResourceOption) (*TeamAdministrators, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Administrators == nil {
		return nil, errors.New("invalid value for required argument 'Administrators'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.TeamId == nil {
		return nil, errors.New("invalid value for required argument 'TeamId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource TeamAdministrators
	err := ctx.RegisterResource("azuredevops:index/teamAdministrators:TeamAdministrators", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTeamAdministrators gets an existing TeamAdministrators resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTeamAdministrators(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TeamAdministratorsState, opts ...pulumi.ResourceOption) (*TeamAdministrators, error) {
	var resource TeamAdministrators
	err := ctx.ReadResource("azuredevops:index/teamAdministrators:TeamAdministrators", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TeamAdministrators resources.
type teamAdministratorsState struct {
	// List of subject descriptors to define adminitrators of the team.
	//
	// > NOTE: It's possible to define team administrators both within the
	// `Team` resource via the `administrators` block and by using the
	// `TeamAdministrators` resource. However it's not possible to use
	// both methods to manage team administrators, since there'll be conflicts.
	Administrators []string `pulumi:"administrators"`
	// The mode how the resource manages team administrators.
	// - `mode == add`: the resource will ensure that all specified administrators will be part of the referenced team
	// - `mode == overwrite`: the resource will replace all existing administrators with the administrators specified within the `administrators` block
	Mode *string `pulumi:"mode"`
	// The Project ID.
	ProjectId *string `pulumi:"projectId"`
	// The ID of the Team.
	TeamId *string `pulumi:"teamId"`
}

type TeamAdministratorsState struct {
	// List of subject descriptors to define adminitrators of the team.
	//
	// > NOTE: It's possible to define team administrators both within the
	// `Team` resource via the `administrators` block and by using the
	// `TeamAdministrators` resource. However it's not possible to use
	// both methods to manage team administrators, since there'll be conflicts.
	Administrators pulumi.StringArrayInput
	// The mode how the resource manages team administrators.
	// - `mode == add`: the resource will ensure that all specified administrators will be part of the referenced team
	// - `mode == overwrite`: the resource will replace all existing administrators with the administrators specified within the `administrators` block
	Mode pulumi.StringPtrInput
	// The Project ID.
	ProjectId pulumi.StringPtrInput
	// The ID of the Team.
	TeamId pulumi.StringPtrInput
}

func (TeamAdministratorsState) ElementType() reflect.Type {
	return reflect.TypeOf((*teamAdministratorsState)(nil)).Elem()
}

type teamAdministratorsArgs struct {
	// List of subject descriptors to define adminitrators of the team.
	//
	// > NOTE: It's possible to define team administrators both within the
	// `Team` resource via the `administrators` block and by using the
	// `TeamAdministrators` resource. However it's not possible to use
	// both methods to manage team administrators, since there'll be conflicts.
	Administrators []string `pulumi:"administrators"`
	// The mode how the resource manages team administrators.
	// - `mode == add`: the resource will ensure that all specified administrators will be part of the referenced team
	// - `mode == overwrite`: the resource will replace all existing administrators with the administrators specified within the `administrators` block
	Mode *string `pulumi:"mode"`
	// The Project ID.
	ProjectId string `pulumi:"projectId"`
	// The ID of the Team.
	TeamId string `pulumi:"teamId"`
}

// The set of arguments for constructing a TeamAdministrators resource.
type TeamAdministratorsArgs struct {
	// List of subject descriptors to define adminitrators of the team.
	//
	// > NOTE: It's possible to define team administrators both within the
	// `Team` resource via the `administrators` block and by using the
	// `TeamAdministrators` resource. However it's not possible to use
	// both methods to manage team administrators, since there'll be conflicts.
	Administrators pulumi.StringArrayInput
	// The mode how the resource manages team administrators.
	// - `mode == add`: the resource will ensure that all specified administrators will be part of the referenced team
	// - `mode == overwrite`: the resource will replace all existing administrators with the administrators specified within the `administrators` block
	Mode pulumi.StringPtrInput
	// The Project ID.
	ProjectId pulumi.StringInput
	// The ID of the Team.
	TeamId pulumi.StringInput
}

func (TeamAdministratorsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*teamAdministratorsArgs)(nil)).Elem()
}

type TeamAdministratorsInput interface {
	pulumi.Input

	ToTeamAdministratorsOutput() TeamAdministratorsOutput
	ToTeamAdministratorsOutputWithContext(ctx context.Context) TeamAdministratorsOutput
}

func (*TeamAdministrators) ElementType() reflect.Type {
	return reflect.TypeOf((**TeamAdministrators)(nil)).Elem()
}

func (i *TeamAdministrators) ToTeamAdministratorsOutput() TeamAdministratorsOutput {
	return i.ToTeamAdministratorsOutputWithContext(context.Background())
}

func (i *TeamAdministrators) ToTeamAdministratorsOutputWithContext(ctx context.Context) TeamAdministratorsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TeamAdministratorsOutput)
}

// TeamAdministratorsArrayInput is an input type that accepts TeamAdministratorsArray and TeamAdministratorsArrayOutput values.
// You can construct a concrete instance of `TeamAdministratorsArrayInput` via:
//
//	TeamAdministratorsArray{ TeamAdministratorsArgs{...} }
type TeamAdministratorsArrayInput interface {
	pulumi.Input

	ToTeamAdministratorsArrayOutput() TeamAdministratorsArrayOutput
	ToTeamAdministratorsArrayOutputWithContext(context.Context) TeamAdministratorsArrayOutput
}

type TeamAdministratorsArray []TeamAdministratorsInput

func (TeamAdministratorsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TeamAdministrators)(nil)).Elem()
}

func (i TeamAdministratorsArray) ToTeamAdministratorsArrayOutput() TeamAdministratorsArrayOutput {
	return i.ToTeamAdministratorsArrayOutputWithContext(context.Background())
}

func (i TeamAdministratorsArray) ToTeamAdministratorsArrayOutputWithContext(ctx context.Context) TeamAdministratorsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TeamAdministratorsArrayOutput)
}

// TeamAdministratorsMapInput is an input type that accepts TeamAdministratorsMap and TeamAdministratorsMapOutput values.
// You can construct a concrete instance of `TeamAdministratorsMapInput` via:
//
//	TeamAdministratorsMap{ "key": TeamAdministratorsArgs{...} }
type TeamAdministratorsMapInput interface {
	pulumi.Input

	ToTeamAdministratorsMapOutput() TeamAdministratorsMapOutput
	ToTeamAdministratorsMapOutputWithContext(context.Context) TeamAdministratorsMapOutput
}

type TeamAdministratorsMap map[string]TeamAdministratorsInput

func (TeamAdministratorsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TeamAdministrators)(nil)).Elem()
}

func (i TeamAdministratorsMap) ToTeamAdministratorsMapOutput() TeamAdministratorsMapOutput {
	return i.ToTeamAdministratorsMapOutputWithContext(context.Background())
}

func (i TeamAdministratorsMap) ToTeamAdministratorsMapOutputWithContext(ctx context.Context) TeamAdministratorsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TeamAdministratorsMapOutput)
}

type TeamAdministratorsOutput struct{ *pulumi.OutputState }

func (TeamAdministratorsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TeamAdministrators)(nil)).Elem()
}

func (o TeamAdministratorsOutput) ToTeamAdministratorsOutput() TeamAdministratorsOutput {
	return o
}

func (o TeamAdministratorsOutput) ToTeamAdministratorsOutputWithContext(ctx context.Context) TeamAdministratorsOutput {
	return o
}

// List of subject descriptors to define adminitrators of the team.
//
// > NOTE: It's possible to define team administrators both within the
// `Team` resource via the `administrators` block and by using the
// `TeamAdministrators` resource. However it's not possible to use
// both methods to manage team administrators, since there'll be conflicts.
func (o TeamAdministratorsOutput) Administrators() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TeamAdministrators) pulumi.StringArrayOutput { return v.Administrators }).(pulumi.StringArrayOutput)
}

// The mode how the resource manages team administrators.
// - `mode == add`: the resource will ensure that all specified administrators will be part of the referenced team
// - `mode == overwrite`: the resource will replace all existing administrators with the administrators specified within the `administrators` block
func (o TeamAdministratorsOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TeamAdministrators) pulumi.StringPtrOutput { return v.Mode }).(pulumi.StringPtrOutput)
}

// The Project ID.
func (o TeamAdministratorsOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *TeamAdministrators) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The ID of the Team.
func (o TeamAdministratorsOutput) TeamId() pulumi.StringOutput {
	return o.ApplyT(func(v *TeamAdministrators) pulumi.StringOutput { return v.TeamId }).(pulumi.StringOutput)
}

type TeamAdministratorsArrayOutput struct{ *pulumi.OutputState }

func (TeamAdministratorsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TeamAdministrators)(nil)).Elem()
}

func (o TeamAdministratorsArrayOutput) ToTeamAdministratorsArrayOutput() TeamAdministratorsArrayOutput {
	return o
}

func (o TeamAdministratorsArrayOutput) ToTeamAdministratorsArrayOutputWithContext(ctx context.Context) TeamAdministratorsArrayOutput {
	return o
}

func (o TeamAdministratorsArrayOutput) Index(i pulumi.IntInput) TeamAdministratorsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TeamAdministrators {
		return vs[0].([]*TeamAdministrators)[vs[1].(int)]
	}).(TeamAdministratorsOutput)
}

type TeamAdministratorsMapOutput struct{ *pulumi.OutputState }

func (TeamAdministratorsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TeamAdministrators)(nil)).Elem()
}

func (o TeamAdministratorsMapOutput) ToTeamAdministratorsMapOutput() TeamAdministratorsMapOutput {
	return o
}

func (o TeamAdministratorsMapOutput) ToTeamAdministratorsMapOutputWithContext(ctx context.Context) TeamAdministratorsMapOutput {
	return o
}

func (o TeamAdministratorsMapOutput) MapIndex(k pulumi.StringInput) TeamAdministratorsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TeamAdministrators {
		return vs[0].(map[string]*TeamAdministrators)[vs[1].(string)]
	}).(TeamAdministratorsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TeamAdministratorsInput)(nil)).Elem(), &TeamAdministrators{})
	pulumi.RegisterInputType(reflect.TypeOf((*TeamAdministratorsArrayInput)(nil)).Elem(), TeamAdministratorsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TeamAdministratorsMapInput)(nil)).Elem(), TeamAdministratorsMap{})
	pulumi.RegisterOutputType(TeamAdministratorsOutput{})
	pulumi.RegisterOutputType(TeamAdministratorsArrayOutput{})
	pulumi.RegisterOutputType(TeamAdministratorsMapOutput{})
}
