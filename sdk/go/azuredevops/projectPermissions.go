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

// Manages permissions for a AzureDevOps project
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
//				Visibility:       pulumi.String("private"),
//				VersionControl:   pulumi.String("Git"),
//				WorkItemTemplate: pulumi.String("Agile"),
//				Description:      pulumi.String("Managed by Terraform"),
//			})
//			if err != nil {
//				return err
//			}
//			example_readers := azuredevops.LookupGroupOutput(ctx, azuredevops.GetGroupOutputArgs{
//				ProjectId: example.ID(),
//				Name:      pulumi.String("Readers"),
//			}, nil)
//			_, err = azuredevops.NewProjectPermissions(ctx, "example-permission", &azuredevops.ProjectPermissionsArgs{
//				ProjectId: example.ID(),
//				Principal: pulumi.String(example_readers.ApplyT(func(example_readers azuredevops.GetGroupResult) (*string, error) {
//					return &example_readers.Id, nil
//				}).(pulumi.StringPtrOutput)),
//				Permissions: pulumi.StringMap{
//					"DELETE":              pulumi.String("Deny"),
//					"EDIT_BUILD_STATUS":   pulumi.String("NotSet"),
//					"WORK_ITEM_MOVE":      pulumi.String("Allow"),
//					"DELETE_TEST_RESULTS": pulumi.String("Deny"),
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
type ProjectPermissions struct {
	pulumi.CustomResourceState

	// the permissions to assign. The following permissions are available
	//
	// | Permission                   | Description                                  |
	// |------------------------------|----------------------------------------------|
	// | GENERIC_READ                 | View project-level information               |
	// | GENERIC_WRITE                | Edit project-level information               |
	// | DELETE                       | Delete team project                          |
	// | PUBLISH_TEST_RESULTS         | Create test runs                             |
	// | ADMINISTER_BUILD             | Administer a build                           |
	// | START_BUILD                  | Start a build                                |
	// | EDIT_BUILD_STATUS            | Edit build quality                           |
	// | UPDATE_BUILD                 | Write to build operational store             |
	// | DELETE_TEST_RESULTS          | Delete test runs                             |
	// | VIEW_TEST_RESULTS            | View test runs                               |
	// | MANAGE_TEST_ENVIRONMENTS     | Manage test environments                     |
	// | MANAGE_TEST_CONFIGURATIONS   | Manage test configurations                   |
	// | WORK_ITEM_DELETE             | Delete and restore work items                |
	// | WORK_ITEM_MOVE               | Move work items out of this project          |
	// | WORK_ITEM_PERMANENTLY_DELETE | Permanently delete work items                |
	// | RENAME                       | Rename team project                          |
	// | MANAGE_PROPERTIES            | Manage project properties                    |
	// | MANAGE_SYSTEM_PROPERTIES     | Manage system project properties             |
	// | BYPASS_PROPERTY_CACHE        | Bypass project property cache                |
	// | BYPASS_RULES                 | Bypass rules on work item updates            |
	// | SUPPRESS_NOTIFICATIONS       | Suppress notifications for work item updates |
	// | UPDATE_VISIBILITY            | Update project visibility                    |
	// | CHANGE_PROCESS               | Change process of team project.              |
	// | AGILETOOLS_BACKLOG           | Agile backlog management.                    |
	// | AGILETOOLS_PLANS             | Agile plans.                                 |
	Permissions pulumi.StringMapOutput `pulumi:"permissions"`
	// The **group** principal to assign the permissions.
	Principal pulumi.StringOutput `pulumi:"principal"`
	// The ID of the project to assign the permissions.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Replace (`true`) or merge (`false`) the permissions. Default: `true`
	Replace pulumi.BoolPtrOutput `pulumi:"replace"`
}

// NewProjectPermissions registers a new resource with the given unique name, arguments, and options.
func NewProjectPermissions(ctx *pulumi.Context,
	name string, args *ProjectPermissionsArgs, opts ...pulumi.ResourceOption) (*ProjectPermissions, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
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
	var resource ProjectPermissions
	err := ctx.RegisterResource("azuredevops:index/projectPermissions:ProjectPermissions", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProjectPermissions gets an existing ProjectPermissions resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProjectPermissions(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectPermissionsState, opts ...pulumi.ResourceOption) (*ProjectPermissions, error) {
	var resource ProjectPermissions
	err := ctx.ReadResource("azuredevops:index/projectPermissions:ProjectPermissions", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProjectPermissions resources.
type projectPermissionsState struct {
	// the permissions to assign. The following permissions are available
	//
	// | Permission                   | Description                                  |
	// |------------------------------|----------------------------------------------|
	// | GENERIC_READ                 | View project-level information               |
	// | GENERIC_WRITE                | Edit project-level information               |
	// | DELETE                       | Delete team project                          |
	// | PUBLISH_TEST_RESULTS         | Create test runs                             |
	// | ADMINISTER_BUILD             | Administer a build                           |
	// | START_BUILD                  | Start a build                                |
	// | EDIT_BUILD_STATUS            | Edit build quality                           |
	// | UPDATE_BUILD                 | Write to build operational store             |
	// | DELETE_TEST_RESULTS          | Delete test runs                             |
	// | VIEW_TEST_RESULTS            | View test runs                               |
	// | MANAGE_TEST_ENVIRONMENTS     | Manage test environments                     |
	// | MANAGE_TEST_CONFIGURATIONS   | Manage test configurations                   |
	// | WORK_ITEM_DELETE             | Delete and restore work items                |
	// | WORK_ITEM_MOVE               | Move work items out of this project          |
	// | WORK_ITEM_PERMANENTLY_DELETE | Permanently delete work items                |
	// | RENAME                       | Rename team project                          |
	// | MANAGE_PROPERTIES            | Manage project properties                    |
	// | MANAGE_SYSTEM_PROPERTIES     | Manage system project properties             |
	// | BYPASS_PROPERTY_CACHE        | Bypass project property cache                |
	// | BYPASS_RULES                 | Bypass rules on work item updates            |
	// | SUPPRESS_NOTIFICATIONS       | Suppress notifications for work item updates |
	// | UPDATE_VISIBILITY            | Update project visibility                    |
	// | CHANGE_PROCESS               | Change process of team project.              |
	// | AGILETOOLS_BACKLOG           | Agile backlog management.                    |
	// | AGILETOOLS_PLANS             | Agile plans.                                 |
	Permissions map[string]string `pulumi:"permissions"`
	// The **group** principal to assign the permissions.
	Principal *string `pulumi:"principal"`
	// The ID of the project to assign the permissions.
	ProjectId *string `pulumi:"projectId"`
	// Replace (`true`) or merge (`false`) the permissions. Default: `true`
	Replace *bool `pulumi:"replace"`
}

type ProjectPermissionsState struct {
	// the permissions to assign. The following permissions are available
	//
	// | Permission                   | Description                                  |
	// |------------------------------|----------------------------------------------|
	// | GENERIC_READ                 | View project-level information               |
	// | GENERIC_WRITE                | Edit project-level information               |
	// | DELETE                       | Delete team project                          |
	// | PUBLISH_TEST_RESULTS         | Create test runs                             |
	// | ADMINISTER_BUILD             | Administer a build                           |
	// | START_BUILD                  | Start a build                                |
	// | EDIT_BUILD_STATUS            | Edit build quality                           |
	// | UPDATE_BUILD                 | Write to build operational store             |
	// | DELETE_TEST_RESULTS          | Delete test runs                             |
	// | VIEW_TEST_RESULTS            | View test runs                               |
	// | MANAGE_TEST_ENVIRONMENTS     | Manage test environments                     |
	// | MANAGE_TEST_CONFIGURATIONS   | Manage test configurations                   |
	// | WORK_ITEM_DELETE             | Delete and restore work items                |
	// | WORK_ITEM_MOVE               | Move work items out of this project          |
	// | WORK_ITEM_PERMANENTLY_DELETE | Permanently delete work items                |
	// | RENAME                       | Rename team project                          |
	// | MANAGE_PROPERTIES            | Manage project properties                    |
	// | MANAGE_SYSTEM_PROPERTIES     | Manage system project properties             |
	// | BYPASS_PROPERTY_CACHE        | Bypass project property cache                |
	// | BYPASS_RULES                 | Bypass rules on work item updates            |
	// | SUPPRESS_NOTIFICATIONS       | Suppress notifications for work item updates |
	// | UPDATE_VISIBILITY            | Update project visibility                    |
	// | CHANGE_PROCESS               | Change process of team project.              |
	// | AGILETOOLS_BACKLOG           | Agile backlog management.                    |
	// | AGILETOOLS_PLANS             | Agile plans.                                 |
	Permissions pulumi.StringMapInput
	// The **group** principal to assign the permissions.
	Principal pulumi.StringPtrInput
	// The ID of the project to assign the permissions.
	ProjectId pulumi.StringPtrInput
	// Replace (`true`) or merge (`false`) the permissions. Default: `true`
	Replace pulumi.BoolPtrInput
}

func (ProjectPermissionsState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectPermissionsState)(nil)).Elem()
}

type projectPermissionsArgs struct {
	// the permissions to assign. The following permissions are available
	//
	// | Permission                   | Description                                  |
	// |------------------------------|----------------------------------------------|
	// | GENERIC_READ                 | View project-level information               |
	// | GENERIC_WRITE                | Edit project-level information               |
	// | DELETE                       | Delete team project                          |
	// | PUBLISH_TEST_RESULTS         | Create test runs                             |
	// | ADMINISTER_BUILD             | Administer a build                           |
	// | START_BUILD                  | Start a build                                |
	// | EDIT_BUILD_STATUS            | Edit build quality                           |
	// | UPDATE_BUILD                 | Write to build operational store             |
	// | DELETE_TEST_RESULTS          | Delete test runs                             |
	// | VIEW_TEST_RESULTS            | View test runs                               |
	// | MANAGE_TEST_ENVIRONMENTS     | Manage test environments                     |
	// | MANAGE_TEST_CONFIGURATIONS   | Manage test configurations                   |
	// | WORK_ITEM_DELETE             | Delete and restore work items                |
	// | WORK_ITEM_MOVE               | Move work items out of this project          |
	// | WORK_ITEM_PERMANENTLY_DELETE | Permanently delete work items                |
	// | RENAME                       | Rename team project                          |
	// | MANAGE_PROPERTIES            | Manage project properties                    |
	// | MANAGE_SYSTEM_PROPERTIES     | Manage system project properties             |
	// | BYPASS_PROPERTY_CACHE        | Bypass project property cache                |
	// | BYPASS_RULES                 | Bypass rules on work item updates            |
	// | SUPPRESS_NOTIFICATIONS       | Suppress notifications for work item updates |
	// | UPDATE_VISIBILITY            | Update project visibility                    |
	// | CHANGE_PROCESS               | Change process of team project.              |
	// | AGILETOOLS_BACKLOG           | Agile backlog management.                    |
	// | AGILETOOLS_PLANS             | Agile plans.                                 |
	Permissions map[string]string `pulumi:"permissions"`
	// The **group** principal to assign the permissions.
	Principal string `pulumi:"principal"`
	// The ID of the project to assign the permissions.
	ProjectId string `pulumi:"projectId"`
	// Replace (`true`) or merge (`false`) the permissions. Default: `true`
	Replace *bool `pulumi:"replace"`
}

// The set of arguments for constructing a ProjectPermissions resource.
type ProjectPermissionsArgs struct {
	// the permissions to assign. The following permissions are available
	//
	// | Permission                   | Description                                  |
	// |------------------------------|----------------------------------------------|
	// | GENERIC_READ                 | View project-level information               |
	// | GENERIC_WRITE                | Edit project-level information               |
	// | DELETE                       | Delete team project                          |
	// | PUBLISH_TEST_RESULTS         | Create test runs                             |
	// | ADMINISTER_BUILD             | Administer a build                           |
	// | START_BUILD                  | Start a build                                |
	// | EDIT_BUILD_STATUS            | Edit build quality                           |
	// | UPDATE_BUILD                 | Write to build operational store             |
	// | DELETE_TEST_RESULTS          | Delete test runs                             |
	// | VIEW_TEST_RESULTS            | View test runs                               |
	// | MANAGE_TEST_ENVIRONMENTS     | Manage test environments                     |
	// | MANAGE_TEST_CONFIGURATIONS   | Manage test configurations                   |
	// | WORK_ITEM_DELETE             | Delete and restore work items                |
	// | WORK_ITEM_MOVE               | Move work items out of this project          |
	// | WORK_ITEM_PERMANENTLY_DELETE | Permanently delete work items                |
	// | RENAME                       | Rename team project                          |
	// | MANAGE_PROPERTIES            | Manage project properties                    |
	// | MANAGE_SYSTEM_PROPERTIES     | Manage system project properties             |
	// | BYPASS_PROPERTY_CACHE        | Bypass project property cache                |
	// | BYPASS_RULES                 | Bypass rules on work item updates            |
	// | SUPPRESS_NOTIFICATIONS       | Suppress notifications for work item updates |
	// | UPDATE_VISIBILITY            | Update project visibility                    |
	// | CHANGE_PROCESS               | Change process of team project.              |
	// | AGILETOOLS_BACKLOG           | Agile backlog management.                    |
	// | AGILETOOLS_PLANS             | Agile plans.                                 |
	Permissions pulumi.StringMapInput
	// The **group** principal to assign the permissions.
	Principal pulumi.StringInput
	// The ID of the project to assign the permissions.
	ProjectId pulumi.StringInput
	// Replace (`true`) or merge (`false`) the permissions. Default: `true`
	Replace pulumi.BoolPtrInput
}

func (ProjectPermissionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectPermissionsArgs)(nil)).Elem()
}

type ProjectPermissionsInput interface {
	pulumi.Input

	ToProjectPermissionsOutput() ProjectPermissionsOutput
	ToProjectPermissionsOutputWithContext(ctx context.Context) ProjectPermissionsOutput
}

func (*ProjectPermissions) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectPermissions)(nil)).Elem()
}

func (i *ProjectPermissions) ToProjectPermissionsOutput() ProjectPermissionsOutput {
	return i.ToProjectPermissionsOutputWithContext(context.Background())
}

func (i *ProjectPermissions) ToProjectPermissionsOutputWithContext(ctx context.Context) ProjectPermissionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectPermissionsOutput)
}

// ProjectPermissionsArrayInput is an input type that accepts ProjectPermissionsArray and ProjectPermissionsArrayOutput values.
// You can construct a concrete instance of `ProjectPermissionsArrayInput` via:
//
//	ProjectPermissionsArray{ ProjectPermissionsArgs{...} }
type ProjectPermissionsArrayInput interface {
	pulumi.Input

	ToProjectPermissionsArrayOutput() ProjectPermissionsArrayOutput
	ToProjectPermissionsArrayOutputWithContext(context.Context) ProjectPermissionsArrayOutput
}

type ProjectPermissionsArray []ProjectPermissionsInput

func (ProjectPermissionsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectPermissions)(nil)).Elem()
}

func (i ProjectPermissionsArray) ToProjectPermissionsArrayOutput() ProjectPermissionsArrayOutput {
	return i.ToProjectPermissionsArrayOutputWithContext(context.Background())
}

func (i ProjectPermissionsArray) ToProjectPermissionsArrayOutputWithContext(ctx context.Context) ProjectPermissionsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectPermissionsArrayOutput)
}

// ProjectPermissionsMapInput is an input type that accepts ProjectPermissionsMap and ProjectPermissionsMapOutput values.
// You can construct a concrete instance of `ProjectPermissionsMapInput` via:
//
//	ProjectPermissionsMap{ "key": ProjectPermissionsArgs{...} }
type ProjectPermissionsMapInput interface {
	pulumi.Input

	ToProjectPermissionsMapOutput() ProjectPermissionsMapOutput
	ToProjectPermissionsMapOutputWithContext(context.Context) ProjectPermissionsMapOutput
}

type ProjectPermissionsMap map[string]ProjectPermissionsInput

func (ProjectPermissionsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectPermissions)(nil)).Elem()
}

func (i ProjectPermissionsMap) ToProjectPermissionsMapOutput() ProjectPermissionsMapOutput {
	return i.ToProjectPermissionsMapOutputWithContext(context.Background())
}

func (i ProjectPermissionsMap) ToProjectPermissionsMapOutputWithContext(ctx context.Context) ProjectPermissionsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectPermissionsMapOutput)
}

type ProjectPermissionsOutput struct{ *pulumi.OutputState }

func (ProjectPermissionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectPermissions)(nil)).Elem()
}

func (o ProjectPermissionsOutput) ToProjectPermissionsOutput() ProjectPermissionsOutput {
	return o
}

func (o ProjectPermissionsOutput) ToProjectPermissionsOutputWithContext(ctx context.Context) ProjectPermissionsOutput {
	return o
}

// the permissions to assign. The following permissions are available
//
// | Permission                   | Description                                  |
// |------------------------------|----------------------------------------------|
// | GENERIC_READ                 | View project-level information               |
// | GENERIC_WRITE                | Edit project-level information               |
// | DELETE                       | Delete team project                          |
// | PUBLISH_TEST_RESULTS         | Create test runs                             |
// | ADMINISTER_BUILD             | Administer a build                           |
// | START_BUILD                  | Start a build                                |
// | EDIT_BUILD_STATUS            | Edit build quality                           |
// | UPDATE_BUILD                 | Write to build operational store             |
// | DELETE_TEST_RESULTS          | Delete test runs                             |
// | VIEW_TEST_RESULTS            | View test runs                               |
// | MANAGE_TEST_ENVIRONMENTS     | Manage test environments                     |
// | MANAGE_TEST_CONFIGURATIONS   | Manage test configurations                   |
// | WORK_ITEM_DELETE             | Delete and restore work items                |
// | WORK_ITEM_MOVE               | Move work items out of this project          |
// | WORK_ITEM_PERMANENTLY_DELETE | Permanently delete work items                |
// | RENAME                       | Rename team project                          |
// | MANAGE_PROPERTIES            | Manage project properties                    |
// | MANAGE_SYSTEM_PROPERTIES     | Manage system project properties             |
// | BYPASS_PROPERTY_CACHE        | Bypass project property cache                |
// | BYPASS_RULES                 | Bypass rules on work item updates            |
// | SUPPRESS_NOTIFICATIONS       | Suppress notifications for work item updates |
// | UPDATE_VISIBILITY            | Update project visibility                    |
// | CHANGE_PROCESS               | Change process of team project.              |
// | AGILETOOLS_BACKLOG           | Agile backlog management.                    |
// | AGILETOOLS_PLANS             | Agile plans.                                 |
func (o ProjectPermissionsOutput) Permissions() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ProjectPermissions) pulumi.StringMapOutput { return v.Permissions }).(pulumi.StringMapOutput)
}

// The **group** principal to assign the permissions.
func (o ProjectPermissionsOutput) Principal() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectPermissions) pulumi.StringOutput { return v.Principal }).(pulumi.StringOutput)
}

// The ID of the project to assign the permissions.
func (o ProjectPermissionsOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectPermissions) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// Replace (`true`) or merge (`false`) the permissions. Default: `true`
func (o ProjectPermissionsOutput) Replace() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ProjectPermissions) pulumi.BoolPtrOutput { return v.Replace }).(pulumi.BoolPtrOutput)
}

type ProjectPermissionsArrayOutput struct{ *pulumi.OutputState }

func (ProjectPermissionsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectPermissions)(nil)).Elem()
}

func (o ProjectPermissionsArrayOutput) ToProjectPermissionsArrayOutput() ProjectPermissionsArrayOutput {
	return o
}

func (o ProjectPermissionsArrayOutput) ToProjectPermissionsArrayOutputWithContext(ctx context.Context) ProjectPermissionsArrayOutput {
	return o
}

func (o ProjectPermissionsArrayOutput) Index(i pulumi.IntInput) ProjectPermissionsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ProjectPermissions {
		return vs[0].([]*ProjectPermissions)[vs[1].(int)]
	}).(ProjectPermissionsOutput)
}

type ProjectPermissionsMapOutput struct{ *pulumi.OutputState }

func (ProjectPermissionsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectPermissions)(nil)).Elem()
}

func (o ProjectPermissionsMapOutput) ToProjectPermissionsMapOutput() ProjectPermissionsMapOutput {
	return o
}

func (o ProjectPermissionsMapOutput) ToProjectPermissionsMapOutputWithContext(ctx context.Context) ProjectPermissionsMapOutput {
	return o
}

func (o ProjectPermissionsMapOutput) MapIndex(k pulumi.StringInput) ProjectPermissionsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ProjectPermissions {
		return vs[0].(map[string]*ProjectPermissions)[vs[1].(string)]
	}).(ProjectPermissionsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectPermissionsInput)(nil)).Elem(), &ProjectPermissions{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectPermissionsArrayInput)(nil)).Elem(), ProjectPermissionsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectPermissionsMapInput)(nil)).Elem(), ProjectPermissionsMap{})
	pulumi.RegisterOutputType(ProjectPermissionsOutput{})
	pulumi.RegisterOutputType(ProjectPermissionsArrayOutput{})
	pulumi.RegisterOutputType(ProjectPermissionsMapOutput{})
}
