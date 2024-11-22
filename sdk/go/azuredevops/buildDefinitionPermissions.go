// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azuredevops/sdk/v3/go/azuredevops/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages permissions for a Build Definition
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
//	"github.com/pulumi/pulumi-azuredevops/sdk/v3/go/azuredevops"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			example, err := azuredevops.NewProject(ctx, "example", &azuredevops.ProjectArgs{
//				Name:             pulumi.String("Example Project"),
//				WorkItemTemplate: pulumi.String("Agile"),
//				VersionControl:   pulumi.String("Git"),
//				Visibility:       pulumi.String("private"),
//				Description:      pulumi.String("Managed by Pulumi"),
//			})
//			if err != nil {
//				return err
//			}
//			example_readers := azuredevops.LookupGroupOutput(ctx, azuredevops.GetGroupOutputArgs{
//				ProjectId: example.ID(),
//				Name:      pulumi.String("Readers"),
//			}, nil)
//			exampleGit, err := azuredevops.NewGit(ctx, "example", &azuredevops.GitArgs{
//				ProjectId: example.ID(),
//				Name:      pulumi.String("Example Repository"),
//				Initialization: &azuredevops.GitInitializationArgs{
//					InitType: pulumi.String("Clean"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			exampleBuildDefinition, err := azuredevops.NewBuildDefinition(ctx, "example", &azuredevops.BuildDefinitionArgs{
//				ProjectId: example.ID(),
//				Name:      pulumi.String("Example Build Definition"),
//				Path:      pulumi.String("\\ExampleFolder"),
//				CiTrigger: &azuredevops.BuildDefinitionCiTriggerArgs{
//					UseYaml: pulumi.Bool(true),
//				},
//				Repository: &azuredevops.BuildDefinitionRepositoryArgs{
//					RepoType:   pulumi.String("TfsGit"),
//					RepoId:     exampleGit.ID(),
//					BranchName: exampleGit.DefaultBranch,
//					YmlPath:    pulumi.String("azure-pipelines.yml"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuredevops.NewBuildDefinitionPermissions(ctx, "example", &azuredevops.BuildDefinitionPermissionsArgs{
//				ProjectId: example.ID(),
//				Principal: pulumi.String(example_readers.ApplyT(func(example_readers azuredevops.GetGroupResult) (*string, error) {
//					return &example_readers.Id, nil
//				}).(pulumi.StringPtrOutput)),
//				BuildDefinitionId: exampleBuildDefinition.ID(),
//				Permissions: pulumi.StringMap{
//					"ViewBuilds":       pulumi.String("Allow"),
//					"EditBuildQuality": pulumi.String("Deny"),
//					"DeleteBuilds":     pulumi.String("Deny"),
//					"StopBuilds":       pulumi.String("Allow"),
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
// ## Relevant Links
//
// * [Azure DevOps Service REST API 7.0 - Security](https://docs.microsoft.com/en-us/rest/api/azure/devops/security/?view=azure-devops-rest-7.0)
//
// ## PAT Permissions Required
//
// - **Project & Team**: vso.security_manage - Grants the ability to read, write, and manage security permissions.
//
// ## Import
//
// The resource does not support import.
type BuildDefinitionPermissions struct {
	pulumi.CustomResourceState

	// The id of the build definition to assign the permissions.
	BuildDefinitionId pulumi.StringOutput `pulumi:"buildDefinitionId"`
	// the permissions to assign. The following permissions are available.
	//
	// | Permission                               | Description                           |
	// |------------------------------------------|---------------------------------------|
	// | ViewBuilds                               | View builds                           |
	// | EditBuildQuality                         | Edit build quality                    |
	// | RetainIndefinitely                       | Retain indefinitely                   |
	// | DeleteBuilds                             | Delete builds                         |
	// | ManageBuildQualities                     | Manage build qualities                |
	// | DestroyBuilds                            | Destroy builds                        |
	// | UpdateBuildInformation                   | Update build information              |
	// | QueueBuilds                              | Queue builds                          |
	// | ManageBuildQueue                         | Manage build queue                    |
	// | StopBuilds                               | Stop builds                           |
	// | ViewBuildDefinition                      | View build pipeline                   |
	// | EditBuildDefinition                      | Edit build pipeline                   |
	// | DeleteBuildDefinition                    | Delete build pipeline                 |
	// | OverrideBuildCheckInValidation           | Override check-in validation by build |
	// | AdministerBuildPermissions               | Administer build permissions          |
	// | CreateBuildDefinition                    | Create build pipeline                 |
	// | EditPipelineQueueConfigurationPermission | Edit queue build configuration        |
	Permissions pulumi.StringMapOutput `pulumi:"permissions"`
	// The **group** principal to assign the permissions.
	Principal pulumi.StringOutput `pulumi:"principal"`
	// The ID of the project to assign the permissions.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Replace (`true`) or merge (`false`) the permissions. Default: `true`.
	Replace pulumi.BoolPtrOutput `pulumi:"replace"`
}

// NewBuildDefinitionPermissions registers a new resource with the given unique name, arguments, and options.
func NewBuildDefinitionPermissions(ctx *pulumi.Context,
	name string, args *BuildDefinitionPermissionsArgs, opts ...pulumi.ResourceOption) (*BuildDefinitionPermissions, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BuildDefinitionId == nil {
		return nil, errors.New("invalid value for required argument 'BuildDefinitionId'")
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
	var resource BuildDefinitionPermissions
	err := ctx.RegisterResource("azuredevops:index/buildDefinitionPermissions:BuildDefinitionPermissions", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBuildDefinitionPermissions gets an existing BuildDefinitionPermissions resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBuildDefinitionPermissions(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BuildDefinitionPermissionsState, opts ...pulumi.ResourceOption) (*BuildDefinitionPermissions, error) {
	var resource BuildDefinitionPermissions
	err := ctx.ReadResource("azuredevops:index/buildDefinitionPermissions:BuildDefinitionPermissions", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BuildDefinitionPermissions resources.
type buildDefinitionPermissionsState struct {
	// The id of the build definition to assign the permissions.
	BuildDefinitionId *string `pulumi:"buildDefinitionId"`
	// the permissions to assign. The following permissions are available.
	//
	// | Permission                               | Description                           |
	// |------------------------------------------|---------------------------------------|
	// | ViewBuilds                               | View builds                           |
	// | EditBuildQuality                         | Edit build quality                    |
	// | RetainIndefinitely                       | Retain indefinitely                   |
	// | DeleteBuilds                             | Delete builds                         |
	// | ManageBuildQualities                     | Manage build qualities                |
	// | DestroyBuilds                            | Destroy builds                        |
	// | UpdateBuildInformation                   | Update build information              |
	// | QueueBuilds                              | Queue builds                          |
	// | ManageBuildQueue                         | Manage build queue                    |
	// | StopBuilds                               | Stop builds                           |
	// | ViewBuildDefinition                      | View build pipeline                   |
	// | EditBuildDefinition                      | Edit build pipeline                   |
	// | DeleteBuildDefinition                    | Delete build pipeline                 |
	// | OverrideBuildCheckInValidation           | Override check-in validation by build |
	// | AdministerBuildPermissions               | Administer build permissions          |
	// | CreateBuildDefinition                    | Create build pipeline                 |
	// | EditPipelineQueueConfigurationPermission | Edit queue build configuration        |
	Permissions map[string]string `pulumi:"permissions"`
	// The **group** principal to assign the permissions.
	Principal *string `pulumi:"principal"`
	// The ID of the project to assign the permissions.
	ProjectId *string `pulumi:"projectId"`
	// Replace (`true`) or merge (`false`) the permissions. Default: `true`.
	Replace *bool `pulumi:"replace"`
}

type BuildDefinitionPermissionsState struct {
	// The id of the build definition to assign the permissions.
	BuildDefinitionId pulumi.StringPtrInput
	// the permissions to assign. The following permissions are available.
	//
	// | Permission                               | Description                           |
	// |------------------------------------------|---------------------------------------|
	// | ViewBuilds                               | View builds                           |
	// | EditBuildQuality                         | Edit build quality                    |
	// | RetainIndefinitely                       | Retain indefinitely                   |
	// | DeleteBuilds                             | Delete builds                         |
	// | ManageBuildQualities                     | Manage build qualities                |
	// | DestroyBuilds                            | Destroy builds                        |
	// | UpdateBuildInformation                   | Update build information              |
	// | QueueBuilds                              | Queue builds                          |
	// | ManageBuildQueue                         | Manage build queue                    |
	// | StopBuilds                               | Stop builds                           |
	// | ViewBuildDefinition                      | View build pipeline                   |
	// | EditBuildDefinition                      | Edit build pipeline                   |
	// | DeleteBuildDefinition                    | Delete build pipeline                 |
	// | OverrideBuildCheckInValidation           | Override check-in validation by build |
	// | AdministerBuildPermissions               | Administer build permissions          |
	// | CreateBuildDefinition                    | Create build pipeline                 |
	// | EditPipelineQueueConfigurationPermission | Edit queue build configuration        |
	Permissions pulumi.StringMapInput
	// The **group** principal to assign the permissions.
	Principal pulumi.StringPtrInput
	// The ID of the project to assign the permissions.
	ProjectId pulumi.StringPtrInput
	// Replace (`true`) or merge (`false`) the permissions. Default: `true`.
	Replace pulumi.BoolPtrInput
}

func (BuildDefinitionPermissionsState) ElementType() reflect.Type {
	return reflect.TypeOf((*buildDefinitionPermissionsState)(nil)).Elem()
}

type buildDefinitionPermissionsArgs struct {
	// The id of the build definition to assign the permissions.
	BuildDefinitionId string `pulumi:"buildDefinitionId"`
	// the permissions to assign. The following permissions are available.
	//
	// | Permission                               | Description                           |
	// |------------------------------------------|---------------------------------------|
	// | ViewBuilds                               | View builds                           |
	// | EditBuildQuality                         | Edit build quality                    |
	// | RetainIndefinitely                       | Retain indefinitely                   |
	// | DeleteBuilds                             | Delete builds                         |
	// | ManageBuildQualities                     | Manage build qualities                |
	// | DestroyBuilds                            | Destroy builds                        |
	// | UpdateBuildInformation                   | Update build information              |
	// | QueueBuilds                              | Queue builds                          |
	// | ManageBuildQueue                         | Manage build queue                    |
	// | StopBuilds                               | Stop builds                           |
	// | ViewBuildDefinition                      | View build pipeline                   |
	// | EditBuildDefinition                      | Edit build pipeline                   |
	// | DeleteBuildDefinition                    | Delete build pipeline                 |
	// | OverrideBuildCheckInValidation           | Override check-in validation by build |
	// | AdministerBuildPermissions               | Administer build permissions          |
	// | CreateBuildDefinition                    | Create build pipeline                 |
	// | EditPipelineQueueConfigurationPermission | Edit queue build configuration        |
	Permissions map[string]string `pulumi:"permissions"`
	// The **group** principal to assign the permissions.
	Principal string `pulumi:"principal"`
	// The ID of the project to assign the permissions.
	ProjectId string `pulumi:"projectId"`
	// Replace (`true`) or merge (`false`) the permissions. Default: `true`.
	Replace *bool `pulumi:"replace"`
}

// The set of arguments for constructing a BuildDefinitionPermissions resource.
type BuildDefinitionPermissionsArgs struct {
	// The id of the build definition to assign the permissions.
	BuildDefinitionId pulumi.StringInput
	// the permissions to assign. The following permissions are available.
	//
	// | Permission                               | Description                           |
	// |------------------------------------------|---------------------------------------|
	// | ViewBuilds                               | View builds                           |
	// | EditBuildQuality                         | Edit build quality                    |
	// | RetainIndefinitely                       | Retain indefinitely                   |
	// | DeleteBuilds                             | Delete builds                         |
	// | ManageBuildQualities                     | Manage build qualities                |
	// | DestroyBuilds                            | Destroy builds                        |
	// | UpdateBuildInformation                   | Update build information              |
	// | QueueBuilds                              | Queue builds                          |
	// | ManageBuildQueue                         | Manage build queue                    |
	// | StopBuilds                               | Stop builds                           |
	// | ViewBuildDefinition                      | View build pipeline                   |
	// | EditBuildDefinition                      | Edit build pipeline                   |
	// | DeleteBuildDefinition                    | Delete build pipeline                 |
	// | OverrideBuildCheckInValidation           | Override check-in validation by build |
	// | AdministerBuildPermissions               | Administer build permissions          |
	// | CreateBuildDefinition                    | Create build pipeline                 |
	// | EditPipelineQueueConfigurationPermission | Edit queue build configuration        |
	Permissions pulumi.StringMapInput
	// The **group** principal to assign the permissions.
	Principal pulumi.StringInput
	// The ID of the project to assign the permissions.
	ProjectId pulumi.StringInput
	// Replace (`true`) or merge (`false`) the permissions. Default: `true`.
	Replace pulumi.BoolPtrInput
}

func (BuildDefinitionPermissionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*buildDefinitionPermissionsArgs)(nil)).Elem()
}

type BuildDefinitionPermissionsInput interface {
	pulumi.Input

	ToBuildDefinitionPermissionsOutput() BuildDefinitionPermissionsOutput
	ToBuildDefinitionPermissionsOutputWithContext(ctx context.Context) BuildDefinitionPermissionsOutput
}

func (*BuildDefinitionPermissions) ElementType() reflect.Type {
	return reflect.TypeOf((**BuildDefinitionPermissions)(nil)).Elem()
}

func (i *BuildDefinitionPermissions) ToBuildDefinitionPermissionsOutput() BuildDefinitionPermissionsOutput {
	return i.ToBuildDefinitionPermissionsOutputWithContext(context.Background())
}

func (i *BuildDefinitionPermissions) ToBuildDefinitionPermissionsOutputWithContext(ctx context.Context) BuildDefinitionPermissionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BuildDefinitionPermissionsOutput)
}

// BuildDefinitionPermissionsArrayInput is an input type that accepts BuildDefinitionPermissionsArray and BuildDefinitionPermissionsArrayOutput values.
// You can construct a concrete instance of `BuildDefinitionPermissionsArrayInput` via:
//
//	BuildDefinitionPermissionsArray{ BuildDefinitionPermissionsArgs{...} }
type BuildDefinitionPermissionsArrayInput interface {
	pulumi.Input

	ToBuildDefinitionPermissionsArrayOutput() BuildDefinitionPermissionsArrayOutput
	ToBuildDefinitionPermissionsArrayOutputWithContext(context.Context) BuildDefinitionPermissionsArrayOutput
}

type BuildDefinitionPermissionsArray []BuildDefinitionPermissionsInput

func (BuildDefinitionPermissionsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BuildDefinitionPermissions)(nil)).Elem()
}

func (i BuildDefinitionPermissionsArray) ToBuildDefinitionPermissionsArrayOutput() BuildDefinitionPermissionsArrayOutput {
	return i.ToBuildDefinitionPermissionsArrayOutputWithContext(context.Background())
}

func (i BuildDefinitionPermissionsArray) ToBuildDefinitionPermissionsArrayOutputWithContext(ctx context.Context) BuildDefinitionPermissionsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BuildDefinitionPermissionsArrayOutput)
}

// BuildDefinitionPermissionsMapInput is an input type that accepts BuildDefinitionPermissionsMap and BuildDefinitionPermissionsMapOutput values.
// You can construct a concrete instance of `BuildDefinitionPermissionsMapInput` via:
//
//	BuildDefinitionPermissionsMap{ "key": BuildDefinitionPermissionsArgs{...} }
type BuildDefinitionPermissionsMapInput interface {
	pulumi.Input

	ToBuildDefinitionPermissionsMapOutput() BuildDefinitionPermissionsMapOutput
	ToBuildDefinitionPermissionsMapOutputWithContext(context.Context) BuildDefinitionPermissionsMapOutput
}

type BuildDefinitionPermissionsMap map[string]BuildDefinitionPermissionsInput

func (BuildDefinitionPermissionsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BuildDefinitionPermissions)(nil)).Elem()
}

func (i BuildDefinitionPermissionsMap) ToBuildDefinitionPermissionsMapOutput() BuildDefinitionPermissionsMapOutput {
	return i.ToBuildDefinitionPermissionsMapOutputWithContext(context.Background())
}

func (i BuildDefinitionPermissionsMap) ToBuildDefinitionPermissionsMapOutputWithContext(ctx context.Context) BuildDefinitionPermissionsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BuildDefinitionPermissionsMapOutput)
}

type BuildDefinitionPermissionsOutput struct{ *pulumi.OutputState }

func (BuildDefinitionPermissionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BuildDefinitionPermissions)(nil)).Elem()
}

func (o BuildDefinitionPermissionsOutput) ToBuildDefinitionPermissionsOutput() BuildDefinitionPermissionsOutput {
	return o
}

func (o BuildDefinitionPermissionsOutput) ToBuildDefinitionPermissionsOutputWithContext(ctx context.Context) BuildDefinitionPermissionsOutput {
	return o
}

// The id of the build definition to assign the permissions.
func (o BuildDefinitionPermissionsOutput) BuildDefinitionId() pulumi.StringOutput {
	return o.ApplyT(func(v *BuildDefinitionPermissions) pulumi.StringOutput { return v.BuildDefinitionId }).(pulumi.StringOutput)
}

// the permissions to assign. The following permissions are available.
//
// | Permission                               | Description                           |
// |------------------------------------------|---------------------------------------|
// | ViewBuilds                               | View builds                           |
// | EditBuildQuality                         | Edit build quality                    |
// | RetainIndefinitely                       | Retain indefinitely                   |
// | DeleteBuilds                             | Delete builds                         |
// | ManageBuildQualities                     | Manage build qualities                |
// | DestroyBuilds                            | Destroy builds                        |
// | UpdateBuildInformation                   | Update build information              |
// | QueueBuilds                              | Queue builds                          |
// | ManageBuildQueue                         | Manage build queue                    |
// | StopBuilds                               | Stop builds                           |
// | ViewBuildDefinition                      | View build pipeline                   |
// | EditBuildDefinition                      | Edit build pipeline                   |
// | DeleteBuildDefinition                    | Delete build pipeline                 |
// | OverrideBuildCheckInValidation           | Override check-in validation by build |
// | AdministerBuildPermissions               | Administer build permissions          |
// | CreateBuildDefinition                    | Create build pipeline                 |
// | EditPipelineQueueConfigurationPermission | Edit queue build configuration        |
func (o BuildDefinitionPermissionsOutput) Permissions() pulumi.StringMapOutput {
	return o.ApplyT(func(v *BuildDefinitionPermissions) pulumi.StringMapOutput { return v.Permissions }).(pulumi.StringMapOutput)
}

// The **group** principal to assign the permissions.
func (o BuildDefinitionPermissionsOutput) Principal() pulumi.StringOutput {
	return o.ApplyT(func(v *BuildDefinitionPermissions) pulumi.StringOutput { return v.Principal }).(pulumi.StringOutput)
}

// The ID of the project to assign the permissions.
func (o BuildDefinitionPermissionsOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *BuildDefinitionPermissions) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// Replace (`true`) or merge (`false`) the permissions. Default: `true`.
func (o BuildDefinitionPermissionsOutput) Replace() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BuildDefinitionPermissions) pulumi.BoolPtrOutput { return v.Replace }).(pulumi.BoolPtrOutput)
}

type BuildDefinitionPermissionsArrayOutput struct{ *pulumi.OutputState }

func (BuildDefinitionPermissionsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BuildDefinitionPermissions)(nil)).Elem()
}

func (o BuildDefinitionPermissionsArrayOutput) ToBuildDefinitionPermissionsArrayOutput() BuildDefinitionPermissionsArrayOutput {
	return o
}

func (o BuildDefinitionPermissionsArrayOutput) ToBuildDefinitionPermissionsArrayOutputWithContext(ctx context.Context) BuildDefinitionPermissionsArrayOutput {
	return o
}

func (o BuildDefinitionPermissionsArrayOutput) Index(i pulumi.IntInput) BuildDefinitionPermissionsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BuildDefinitionPermissions {
		return vs[0].([]*BuildDefinitionPermissions)[vs[1].(int)]
	}).(BuildDefinitionPermissionsOutput)
}

type BuildDefinitionPermissionsMapOutput struct{ *pulumi.OutputState }

func (BuildDefinitionPermissionsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BuildDefinitionPermissions)(nil)).Elem()
}

func (o BuildDefinitionPermissionsMapOutput) ToBuildDefinitionPermissionsMapOutput() BuildDefinitionPermissionsMapOutput {
	return o
}

func (o BuildDefinitionPermissionsMapOutput) ToBuildDefinitionPermissionsMapOutputWithContext(ctx context.Context) BuildDefinitionPermissionsMapOutput {
	return o
}

func (o BuildDefinitionPermissionsMapOutput) MapIndex(k pulumi.StringInput) BuildDefinitionPermissionsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BuildDefinitionPermissions {
		return vs[0].(map[string]*BuildDefinitionPermissions)[vs[1].(string)]
	}).(BuildDefinitionPermissionsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BuildDefinitionPermissionsInput)(nil)).Elem(), &BuildDefinitionPermissions{})
	pulumi.RegisterInputType(reflect.TypeOf((*BuildDefinitionPermissionsArrayInput)(nil)).Elem(), BuildDefinitionPermissionsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BuildDefinitionPermissionsMapInput)(nil)).Elem(), BuildDefinitionPermissionsMap{})
	pulumi.RegisterOutputType(BuildDefinitionPermissionsOutput{})
	pulumi.RegisterOutputType(BuildDefinitionPermissionsArrayOutput{})
	pulumi.RegisterOutputType(BuildDefinitionPermissionsMapOutput{})
}
