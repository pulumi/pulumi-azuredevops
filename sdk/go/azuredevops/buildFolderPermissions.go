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

// Manages permissions for a Build Folder
//
// > **Note** Permissions can be assigned to group principals and not to single user principals.
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
//			example_readers := azuredevops.LookupGroupOutput(ctx, azuredevops.GetGroupOutputArgs{
//				ProjectId: exampleProject.ID(),
//				Name:      pulumi.String("Readers"),
//			}, nil)
//			_, err = azuredevops.NewBuildFolder(ctx, "exampleBuildFolder", &azuredevops.BuildFolderArgs{
//				ProjectId:   exampleProject.ID(),
//				Path:        pulumi.String("\\ExampleFolder"),
//				Description: pulumi.String("ExampleFolder description"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuredevops.NewBuildFolderPermissions(ctx, "exampleBuildFolderPermissions", &azuredevops.BuildFolderPermissionsArgs{
//				ProjectId: exampleProject.ID(),
//				Path:      pulumi.String("\\ExampleFolder"),
//				Principal: example_readers.ApplyT(func(example_readers azuredevops.GetGroupResult) (*string, error) {
//					return &example_readers.Id, nil
//				}).(pulumi.StringPtrOutput),
//				Permissions: pulumi.StringMap{
//					"ViewBuilds":                 pulumi.String("Allow"),
//					"EditBuildQuality":           pulumi.String("Allow"),
//					"RetainIndefinitely":         pulumi.String("Allow"),
//					"DeleteBuilds":               pulumi.String("Deny"),
//					"ManageBuildQualities":       pulumi.String("Deny"),
//					"DestroyBuilds":              pulumi.String("Deny"),
//					"UpdateBuildInformation":     pulumi.String("Deny"),
//					"QueueBuilds":                pulumi.String("Allow"),
//					"ManageBuildQueue":           pulumi.String("Deny"),
//					"StopBuilds":                 pulumi.String("Allow"),
//					"ViewBuildDefinition":        pulumi.String("Allow"),
//					"EditBuildDefinition":        pulumi.String("Deny"),
//					"DeleteBuildDefinition":      pulumi.String("Deny"),
//					"AdministerBuildPermissions": pulumi.String("NotSet"),
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
// * [Azure DevOps Service REST API 6.0 - Security](https://docs.microsoft.com/en-us/rest/api/azure/devops/security/?view=azure-devops-rest-6.0)
//
// ## PAT Permissions Required
//
// - **Project & Team**: vso.security_manage - Grants the ability to read, write, and manage security permissions.
//
// ## Import
//
// The resource does not support import.
type BuildFolderPermissions struct {
	pulumi.CustomResourceState

	// The folder path to assign the permissions.
	Path pulumi.StringOutput `pulumi:"path"`
	// the permissions to assign. The following permissions are available.
	//
	// | Permission                     | Description                           |
	// |--------------------------------|---------------------------------------|
	// | ViewBuilds                     | View builds                           |
	// | EditBuildQuality               | Edit build quality                    |
	// | RetainIndefinitely             | Retain indefinitely                   |
	// | DeleteBuilds                   | Delete builds                         |
	// | ManageBuildQualities           | Manage build qualities                |
	// | DestroyBuilds                  | Destroy builds                        |
	// | UpdateBuildInformation         | Update build information              |
	// | QueueBuilds                    | Queue builds                          |
	// | ManageBuildQueue               | Manage build queue                    |
	// | StopBuilds                     | Stop builds                           |
	// | ViewBuildDefinition            | View build pipeline                   |
	// | EditBuildDefinition            | Edit build pipeline                   |
	// | DeleteBuildDefinition          | Delete build pipeline                 |
	// | OverrideBuildCheckInValidation | Override check-in validation by build |
	// | AdministerBuildPermissions     | Administer build permissions          |
	Permissions pulumi.StringMapOutput `pulumi:"permissions"`
	// The **group** principal to assign the permissions.
	Principal pulumi.StringOutput `pulumi:"principal"`
	// The ID of the project to assign the permissions.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Replace (`true`) or merge (`false`) the permissions. Default: `true`.
	Replace pulumi.BoolPtrOutput `pulumi:"replace"`
}

// NewBuildFolderPermissions registers a new resource with the given unique name, arguments, and options.
func NewBuildFolderPermissions(ctx *pulumi.Context,
	name string, args *BuildFolderPermissionsArgs, opts ...pulumi.ResourceOption) (*BuildFolderPermissions, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Path == nil {
		return nil, errors.New("invalid value for required argument 'Path'")
	}
	if args.Permissions == nil {
		return nil, errors.New("invalid value for required argument 'Permissions'")
	}
	if args.Principal == nil {
		return nil, errors.New("invalid value for required argument 'Principal'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource BuildFolderPermissions
	err := ctx.RegisterResource("azuredevops:index/buildFolderPermissions:BuildFolderPermissions", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBuildFolderPermissions gets an existing BuildFolderPermissions resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBuildFolderPermissions(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BuildFolderPermissionsState, opts ...pulumi.ResourceOption) (*BuildFolderPermissions, error) {
	var resource BuildFolderPermissions
	err := ctx.ReadResource("azuredevops:index/buildFolderPermissions:BuildFolderPermissions", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BuildFolderPermissions resources.
type buildFolderPermissionsState struct {
	// The folder path to assign the permissions.
	Path *string `pulumi:"path"`
	// the permissions to assign. The following permissions are available.
	//
	// | Permission                     | Description                           |
	// |--------------------------------|---------------------------------------|
	// | ViewBuilds                     | View builds                           |
	// | EditBuildQuality               | Edit build quality                    |
	// | RetainIndefinitely             | Retain indefinitely                   |
	// | DeleteBuilds                   | Delete builds                         |
	// | ManageBuildQualities           | Manage build qualities                |
	// | DestroyBuilds                  | Destroy builds                        |
	// | UpdateBuildInformation         | Update build information              |
	// | QueueBuilds                    | Queue builds                          |
	// | ManageBuildQueue               | Manage build queue                    |
	// | StopBuilds                     | Stop builds                           |
	// | ViewBuildDefinition            | View build pipeline                   |
	// | EditBuildDefinition            | Edit build pipeline                   |
	// | DeleteBuildDefinition          | Delete build pipeline                 |
	// | OverrideBuildCheckInValidation | Override check-in validation by build |
	// | AdministerBuildPermissions     | Administer build permissions          |
	Permissions map[string]string `pulumi:"permissions"`
	// The **group** principal to assign the permissions.
	Principal *string `pulumi:"principal"`
	// The ID of the project to assign the permissions.
	ProjectId *string `pulumi:"projectId"`
	// Replace (`true`) or merge (`false`) the permissions. Default: `true`.
	Replace *bool `pulumi:"replace"`
}

type BuildFolderPermissionsState struct {
	// The folder path to assign the permissions.
	Path pulumi.StringPtrInput
	// the permissions to assign. The following permissions are available.
	//
	// | Permission                     | Description                           |
	// |--------------------------------|---------------------------------------|
	// | ViewBuilds                     | View builds                           |
	// | EditBuildQuality               | Edit build quality                    |
	// | RetainIndefinitely             | Retain indefinitely                   |
	// | DeleteBuilds                   | Delete builds                         |
	// | ManageBuildQualities           | Manage build qualities                |
	// | DestroyBuilds                  | Destroy builds                        |
	// | UpdateBuildInformation         | Update build information              |
	// | QueueBuilds                    | Queue builds                          |
	// | ManageBuildQueue               | Manage build queue                    |
	// | StopBuilds                     | Stop builds                           |
	// | ViewBuildDefinition            | View build pipeline                   |
	// | EditBuildDefinition            | Edit build pipeline                   |
	// | DeleteBuildDefinition          | Delete build pipeline                 |
	// | OverrideBuildCheckInValidation | Override check-in validation by build |
	// | AdministerBuildPermissions     | Administer build permissions          |
	Permissions pulumi.StringMapInput
	// The **group** principal to assign the permissions.
	Principal pulumi.StringPtrInput
	// The ID of the project to assign the permissions.
	ProjectId pulumi.StringPtrInput
	// Replace (`true`) or merge (`false`) the permissions. Default: `true`.
	Replace pulumi.BoolPtrInput
}

func (BuildFolderPermissionsState) ElementType() reflect.Type {
	return reflect.TypeOf((*buildFolderPermissionsState)(nil)).Elem()
}

type buildFolderPermissionsArgs struct {
	// The folder path to assign the permissions.
	Path string `pulumi:"path"`
	// the permissions to assign. The following permissions are available.
	//
	// | Permission                     | Description                           |
	// |--------------------------------|---------------------------------------|
	// | ViewBuilds                     | View builds                           |
	// | EditBuildQuality               | Edit build quality                    |
	// | RetainIndefinitely             | Retain indefinitely                   |
	// | DeleteBuilds                   | Delete builds                         |
	// | ManageBuildQualities           | Manage build qualities                |
	// | DestroyBuilds                  | Destroy builds                        |
	// | UpdateBuildInformation         | Update build information              |
	// | QueueBuilds                    | Queue builds                          |
	// | ManageBuildQueue               | Manage build queue                    |
	// | StopBuilds                     | Stop builds                           |
	// | ViewBuildDefinition            | View build pipeline                   |
	// | EditBuildDefinition            | Edit build pipeline                   |
	// | DeleteBuildDefinition          | Delete build pipeline                 |
	// | OverrideBuildCheckInValidation | Override check-in validation by build |
	// | AdministerBuildPermissions     | Administer build permissions          |
	Permissions map[string]string `pulumi:"permissions"`
	// The **group** principal to assign the permissions.
	Principal string `pulumi:"principal"`
	// The ID of the project to assign the permissions.
	ProjectId string `pulumi:"projectId"`
	// Replace (`true`) or merge (`false`) the permissions. Default: `true`.
	Replace *bool `pulumi:"replace"`
}

// The set of arguments for constructing a BuildFolderPermissions resource.
type BuildFolderPermissionsArgs struct {
	// The folder path to assign the permissions.
	Path pulumi.StringInput
	// the permissions to assign. The following permissions are available.
	//
	// | Permission                     | Description                           |
	// |--------------------------------|---------------------------------------|
	// | ViewBuilds                     | View builds                           |
	// | EditBuildQuality               | Edit build quality                    |
	// | RetainIndefinitely             | Retain indefinitely                   |
	// | DeleteBuilds                   | Delete builds                         |
	// | ManageBuildQualities           | Manage build qualities                |
	// | DestroyBuilds                  | Destroy builds                        |
	// | UpdateBuildInformation         | Update build information              |
	// | QueueBuilds                    | Queue builds                          |
	// | ManageBuildQueue               | Manage build queue                    |
	// | StopBuilds                     | Stop builds                           |
	// | ViewBuildDefinition            | View build pipeline                   |
	// | EditBuildDefinition            | Edit build pipeline                   |
	// | DeleteBuildDefinition          | Delete build pipeline                 |
	// | OverrideBuildCheckInValidation | Override check-in validation by build |
	// | AdministerBuildPermissions     | Administer build permissions          |
	Permissions pulumi.StringMapInput
	// The **group** principal to assign the permissions.
	Principal pulumi.StringInput
	// The ID of the project to assign the permissions.
	ProjectId pulumi.StringInput
	// Replace (`true`) or merge (`false`) the permissions. Default: `true`.
	Replace pulumi.BoolPtrInput
}

func (BuildFolderPermissionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*buildFolderPermissionsArgs)(nil)).Elem()
}

type BuildFolderPermissionsInput interface {
	pulumi.Input

	ToBuildFolderPermissionsOutput() BuildFolderPermissionsOutput
	ToBuildFolderPermissionsOutputWithContext(ctx context.Context) BuildFolderPermissionsOutput
}

func (*BuildFolderPermissions) ElementType() reflect.Type {
	return reflect.TypeOf((**BuildFolderPermissions)(nil)).Elem()
}

func (i *BuildFolderPermissions) ToBuildFolderPermissionsOutput() BuildFolderPermissionsOutput {
	return i.ToBuildFolderPermissionsOutputWithContext(context.Background())
}

func (i *BuildFolderPermissions) ToBuildFolderPermissionsOutputWithContext(ctx context.Context) BuildFolderPermissionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BuildFolderPermissionsOutput)
}

// BuildFolderPermissionsArrayInput is an input type that accepts BuildFolderPermissionsArray and BuildFolderPermissionsArrayOutput values.
// You can construct a concrete instance of `BuildFolderPermissionsArrayInput` via:
//
//	BuildFolderPermissionsArray{ BuildFolderPermissionsArgs{...} }
type BuildFolderPermissionsArrayInput interface {
	pulumi.Input

	ToBuildFolderPermissionsArrayOutput() BuildFolderPermissionsArrayOutput
	ToBuildFolderPermissionsArrayOutputWithContext(context.Context) BuildFolderPermissionsArrayOutput
}

type BuildFolderPermissionsArray []BuildFolderPermissionsInput

func (BuildFolderPermissionsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BuildFolderPermissions)(nil)).Elem()
}

func (i BuildFolderPermissionsArray) ToBuildFolderPermissionsArrayOutput() BuildFolderPermissionsArrayOutput {
	return i.ToBuildFolderPermissionsArrayOutputWithContext(context.Background())
}

func (i BuildFolderPermissionsArray) ToBuildFolderPermissionsArrayOutputWithContext(ctx context.Context) BuildFolderPermissionsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BuildFolderPermissionsArrayOutput)
}

// BuildFolderPermissionsMapInput is an input type that accepts BuildFolderPermissionsMap and BuildFolderPermissionsMapOutput values.
// You can construct a concrete instance of `BuildFolderPermissionsMapInput` via:
//
//	BuildFolderPermissionsMap{ "key": BuildFolderPermissionsArgs{...} }
type BuildFolderPermissionsMapInput interface {
	pulumi.Input

	ToBuildFolderPermissionsMapOutput() BuildFolderPermissionsMapOutput
	ToBuildFolderPermissionsMapOutputWithContext(context.Context) BuildFolderPermissionsMapOutput
}

type BuildFolderPermissionsMap map[string]BuildFolderPermissionsInput

func (BuildFolderPermissionsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BuildFolderPermissions)(nil)).Elem()
}

func (i BuildFolderPermissionsMap) ToBuildFolderPermissionsMapOutput() BuildFolderPermissionsMapOutput {
	return i.ToBuildFolderPermissionsMapOutputWithContext(context.Background())
}

func (i BuildFolderPermissionsMap) ToBuildFolderPermissionsMapOutputWithContext(ctx context.Context) BuildFolderPermissionsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BuildFolderPermissionsMapOutput)
}

type BuildFolderPermissionsOutput struct{ *pulumi.OutputState }

func (BuildFolderPermissionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BuildFolderPermissions)(nil)).Elem()
}

func (o BuildFolderPermissionsOutput) ToBuildFolderPermissionsOutput() BuildFolderPermissionsOutput {
	return o
}

func (o BuildFolderPermissionsOutput) ToBuildFolderPermissionsOutputWithContext(ctx context.Context) BuildFolderPermissionsOutput {
	return o
}

// The folder path to assign the permissions.
func (o BuildFolderPermissionsOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v *BuildFolderPermissions) pulumi.StringOutput { return v.Path }).(pulumi.StringOutput)
}

// the permissions to assign. The following permissions are available.
//
// | Permission                     | Description                           |
// |--------------------------------|---------------------------------------|
// | ViewBuilds                     | View builds                           |
// | EditBuildQuality               | Edit build quality                    |
// | RetainIndefinitely             | Retain indefinitely                   |
// | DeleteBuilds                   | Delete builds                         |
// | ManageBuildQualities           | Manage build qualities                |
// | DestroyBuilds                  | Destroy builds                        |
// | UpdateBuildInformation         | Update build information              |
// | QueueBuilds                    | Queue builds                          |
// | ManageBuildQueue               | Manage build queue                    |
// | StopBuilds                     | Stop builds                           |
// | ViewBuildDefinition            | View build pipeline                   |
// | EditBuildDefinition            | Edit build pipeline                   |
// | DeleteBuildDefinition          | Delete build pipeline                 |
// | OverrideBuildCheckInValidation | Override check-in validation by build |
// | AdministerBuildPermissions     | Administer build permissions          |
func (o BuildFolderPermissionsOutput) Permissions() pulumi.StringMapOutput {
	return o.ApplyT(func(v *BuildFolderPermissions) pulumi.StringMapOutput { return v.Permissions }).(pulumi.StringMapOutput)
}

// The **group** principal to assign the permissions.
func (o BuildFolderPermissionsOutput) Principal() pulumi.StringOutput {
	return o.ApplyT(func(v *BuildFolderPermissions) pulumi.StringOutput { return v.Principal }).(pulumi.StringOutput)
}

// The ID of the project to assign the permissions.
func (o BuildFolderPermissionsOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *BuildFolderPermissions) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// Replace (`true`) or merge (`false`) the permissions. Default: `true`.
func (o BuildFolderPermissionsOutput) Replace() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BuildFolderPermissions) pulumi.BoolPtrOutput { return v.Replace }).(pulumi.BoolPtrOutput)
}

type BuildFolderPermissionsArrayOutput struct{ *pulumi.OutputState }

func (BuildFolderPermissionsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BuildFolderPermissions)(nil)).Elem()
}

func (o BuildFolderPermissionsArrayOutput) ToBuildFolderPermissionsArrayOutput() BuildFolderPermissionsArrayOutput {
	return o
}

func (o BuildFolderPermissionsArrayOutput) ToBuildFolderPermissionsArrayOutputWithContext(ctx context.Context) BuildFolderPermissionsArrayOutput {
	return o
}

func (o BuildFolderPermissionsArrayOutput) Index(i pulumi.IntInput) BuildFolderPermissionsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BuildFolderPermissions {
		return vs[0].([]*BuildFolderPermissions)[vs[1].(int)]
	}).(BuildFolderPermissionsOutput)
}

type BuildFolderPermissionsMapOutput struct{ *pulumi.OutputState }

func (BuildFolderPermissionsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BuildFolderPermissions)(nil)).Elem()
}

func (o BuildFolderPermissionsMapOutput) ToBuildFolderPermissionsMapOutput() BuildFolderPermissionsMapOutput {
	return o
}

func (o BuildFolderPermissionsMapOutput) ToBuildFolderPermissionsMapOutputWithContext(ctx context.Context) BuildFolderPermissionsMapOutput {
	return o
}

func (o BuildFolderPermissionsMapOutput) MapIndex(k pulumi.StringInput) BuildFolderPermissionsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BuildFolderPermissions {
		return vs[0].(map[string]*BuildFolderPermissions)[vs[1].(string)]
	}).(BuildFolderPermissionsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BuildFolderPermissionsInput)(nil)).Elem(), &BuildFolderPermissions{})
	pulumi.RegisterInputType(reflect.TypeOf((*BuildFolderPermissionsArrayInput)(nil)).Elem(), BuildFolderPermissionsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BuildFolderPermissionsMapInput)(nil)).Elem(), BuildFolderPermissionsMap{})
	pulumi.RegisterOutputType(BuildFolderPermissionsOutput{})
	pulumi.RegisterOutputType(BuildFolderPermissionsArrayOutput{})
	pulumi.RegisterOutputType(BuildFolderPermissionsMapOutput{})
}
