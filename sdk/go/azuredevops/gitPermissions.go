// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages permissions for Git repositories.
//
// > **Note** Permissions can be assigned to group principals and not to single user principals.
//
// ## Permission levels
//
// Permission for Git Repositories within Azure DevOps can be applied on three different levels.
// Those levels are reflected by specifying (or omitting) values for the arguments `projectId`, `repositoryId` and `branchName`.
//
// ### Project level
//
// Permissions for all Git Repositories inside a project (existing or newly created ones) are specified, if only the argument `projectId` has a value.
//
// #### Example usage
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
//			example, err := azuredevops.NewProject(ctx, "example", &azuredevops.ProjectArgs{
//				WorkItemTemplate: pulumi.String("Agile"),
//				VersionControl:   pulumi.String("Git"),
//				Visibility:       pulumi.String("private"),
//				Description:      pulumi.String("Managed by Terraform"),
//			})
//			if err != nil {
//				return err
//			}
//			example_readers := azuredevops.LookupGroupOutput(ctx, azuredevops.GetGroupOutputArgs{
//				ProjectId: example.ID(),
//				Name:      pulumi.String("Readers"),
//			}, nil)
//			_, err = azuredevops.NewGitPermissions(ctx, "example-permissions", &azuredevops.GitPermissionsArgs{
//				ProjectId: example.ID(),
//				Principal: example_readers.ApplyT(func(example_readers azuredevops.GetGroupResult) (*string, error) {
//					return &example_readers.Id, nil
//				}).(pulumi.StringPtrOutput),
//				Permissions: pulumi.StringMap{
//					"CreateRepository": pulumi.String("Deny"),
//					"DeleteRepository": pulumi.String("Deny"),
//					"RenameRepository": pulumi.String("NotSet"),
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
// ### Repository level
//
// Permissions for a specific Git Repository and all existing or newly created branches are specified if the arguments `projectId` and `repositoryId` are set.
//
// #### Example usage
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
//			example_group, err := azuredevops.LookupGroup(ctx, &azuredevops.LookupGroupArgs{
//				Name: "Project Collection Administrators",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			exampleGit, err := azuredevops.NewGit(ctx, "exampleGit", &azuredevops.GitArgs{
//				ProjectId: exampleProject.ID(),
//				Initialization: &azuredevops.GitInitializationArgs{
//					InitType: pulumi.String("Clean"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuredevops.NewGitPermissions(ctx, "example-permissions", &azuredevops.GitPermissionsArgs{
//				ProjectId:    exampleGit.ProjectId,
//				RepositoryId: exampleGit.ID(),
//				Principal:    *pulumi.String(example_group.Id),
//				Permissions: pulumi.StringMap{
//					"RemoveOthersLocks": pulumi.String("Allow"),
//					"ManagePermissions": pulumi.String("Deny"),
//					"CreateTag":         pulumi.String("Deny"),
//					"CreateBranch":      pulumi.String("NotSet"),
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
// ### Branch level
//
// Permissions for a specific branch inside a Git Repository are specified if all above mentioned the arguments are set.
//
// #### Example usage
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
//			exampleGit, err := azuredevops.NewGit(ctx, "exampleGit", &azuredevops.GitArgs{
//				ProjectId: exampleProject.ID(),
//				Initialization: &azuredevops.GitInitializationArgs{
//					InitType: pulumi.String("Clean"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			example_group, err := azuredevops.LookupGroup(ctx, &azuredevops.LookupGroupArgs{
//				Name: "Project Collection Administrators",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = azuredevops.NewGitPermissions(ctx, "example-permissions", &azuredevops.GitPermissionsArgs{
//				ProjectId:    exampleGit.ProjectId,
//				RepositoryId: exampleGit.ID(),
//				BranchName:   pulumi.String("refs/heads/master"),
//				Principal:    *pulumi.String(example_group.Id),
//				Permissions: pulumi.StringMap{
//					"RemoveOthersLocks": pulumi.String("Allow"),
//					"ForcePush":         pulumi.String("Deny"),
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
//				Visibility:       pulumi.String("private"),
//				VersionControl:   pulumi.String("Git"),
//				WorkItemTemplate: pulumi.String("Agile"),
//				Description:      pulumi.String("Managed by Terraform"),
//			})
//			if err != nil {
//				return err
//			}
//			example_project_readers := azuredevops.LookupGroupOutput(ctx, azuredevops.GetGroupOutputArgs{
//				ProjectId: exampleProject.ID(),
//				Name:      pulumi.String("Readers"),
//			}, nil)
//			example_project_contributors := azuredevops.LookupGroupOutput(ctx, azuredevops.GetGroupOutputArgs{
//				ProjectId: exampleProject.ID(),
//				Name:      pulumi.String("Contributors"),
//			}, nil)
//			example_project_administrators := azuredevops.LookupGroupOutput(ctx, azuredevops.GetGroupOutputArgs{
//				ProjectId: exampleProject.ID(),
//				Name:      pulumi.String("Project administrators"),
//			}, nil)
//			_, err = azuredevops.NewGitPermissions(ctx, "example-permissions", &azuredevops.GitPermissionsArgs{
//				ProjectId: exampleProject.ID(),
//				Principal: example_project_readers.ApplyT(func(example_project_readers azuredevops.GetGroupResult) (*string, error) {
//					return &example_project_readers.Id, nil
//				}).(pulumi.StringPtrOutput),
//				Permissions: pulumi.StringMap{
//					"CreateRepository": pulumi.String("Deny"),
//					"DeleteRepository": pulumi.String("Deny"),
//					"RenameRepository": pulumi.String("NotSet"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			exampleGit, err := azuredevops.NewGit(ctx, "exampleGit", &azuredevops.GitArgs{
//				ProjectId:     exampleProject.ID(),
//				DefaultBranch: pulumi.String("refs/heads/master"),
//				Initialization: &azuredevops.GitInitializationArgs{
//					InitType: pulumi.String("Clean"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuredevops.NewGitPermissions(ctx, "example-repo-permissions", &azuredevops.GitPermissionsArgs{
//				ProjectId:    exampleGit.ProjectId,
//				RepositoryId: exampleGit.ID(),
//				Principal: example_project_administrators.ApplyT(func(example_project_administrators azuredevops.GetGroupResult) (*string, error) {
//					return &example_project_administrators.Id, nil
//				}).(pulumi.StringPtrOutput),
//				Permissions: pulumi.StringMap{
//					"RemoveOthersLocks": pulumi.String("Allow"),
//					"ManagePermissions": pulumi.String("Deny"),
//					"CreateTag":         pulumi.String("Deny"),
//					"CreateBranch":      pulumi.String("NotSet"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuredevops.NewGitPermissions(ctx, "example-branch-permissions", &azuredevops.GitPermissionsArgs{
//				ProjectId:    exampleGit.ProjectId,
//				RepositoryId: exampleGit.ID(),
//				BranchName:   pulumi.String("master"),
//				Principal: example_project_contributors.ApplyT(func(example_project_contributors azuredevops.GetGroupResult) (*string, error) {
//					return &example_project_contributors.Id, nil
//				}).(pulumi.StringPtrOutput),
//				Permissions: pulumi.StringMap{
//					"RemoveOthersLocks": pulumi.String("Allow"),
//					"ForcePush":         pulumi.String("Deny"),
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
type GitPermissions struct {
	pulumi.CustomResourceState

	// The name of the branch to assign the permissions.
	BranchName pulumi.StringPtrOutput `pulumi:"branchName"`
	// the permissions to assign. The follwing permissions are available
	Permissions pulumi.StringMapOutput `pulumi:"permissions"`
	// The **group** principal to assign the permissions.
	Principal pulumi.StringOutput `pulumi:"principal"`
	// The ID of the project to assign the permissions.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Replace (`true`) or merge (`false`) the permissions. Default: `true`
	Replace pulumi.BoolPtrOutput `pulumi:"replace"`
	// The ID of the GIT repository to assign the permissions
	RepositoryId pulumi.StringPtrOutput `pulumi:"repositoryId"`
}

// NewGitPermissions registers a new resource with the given unique name, arguments, and options.
func NewGitPermissions(ctx *pulumi.Context,
	name string, args *GitPermissionsArgs, opts ...pulumi.ResourceOption) (*GitPermissions, error) {
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
	var resource GitPermissions
	err := ctx.RegisterResource("azuredevops:index/gitPermissions:GitPermissions", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGitPermissions gets an existing GitPermissions resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGitPermissions(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GitPermissionsState, opts ...pulumi.ResourceOption) (*GitPermissions, error) {
	var resource GitPermissions
	err := ctx.ReadResource("azuredevops:index/gitPermissions:GitPermissions", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GitPermissions resources.
type gitPermissionsState struct {
	// The name of the branch to assign the permissions.
	BranchName *string `pulumi:"branchName"`
	// the permissions to assign. The follwing permissions are available
	Permissions map[string]string `pulumi:"permissions"`
	// The **group** principal to assign the permissions.
	Principal *string `pulumi:"principal"`
	// The ID of the project to assign the permissions.
	ProjectId *string `pulumi:"projectId"`
	// Replace (`true`) or merge (`false`) the permissions. Default: `true`
	Replace *bool `pulumi:"replace"`
	// The ID of the GIT repository to assign the permissions
	RepositoryId *string `pulumi:"repositoryId"`
}

type GitPermissionsState struct {
	// The name of the branch to assign the permissions.
	BranchName pulumi.StringPtrInput
	// the permissions to assign. The follwing permissions are available
	Permissions pulumi.StringMapInput
	// The **group** principal to assign the permissions.
	Principal pulumi.StringPtrInput
	// The ID of the project to assign the permissions.
	ProjectId pulumi.StringPtrInput
	// Replace (`true`) or merge (`false`) the permissions. Default: `true`
	Replace pulumi.BoolPtrInput
	// The ID of the GIT repository to assign the permissions
	RepositoryId pulumi.StringPtrInput
}

func (GitPermissionsState) ElementType() reflect.Type {
	return reflect.TypeOf((*gitPermissionsState)(nil)).Elem()
}

type gitPermissionsArgs struct {
	// The name of the branch to assign the permissions.
	BranchName *string `pulumi:"branchName"`
	// the permissions to assign. The follwing permissions are available
	Permissions map[string]string `pulumi:"permissions"`
	// The **group** principal to assign the permissions.
	Principal string `pulumi:"principal"`
	// The ID of the project to assign the permissions.
	ProjectId string `pulumi:"projectId"`
	// Replace (`true`) or merge (`false`) the permissions. Default: `true`
	Replace *bool `pulumi:"replace"`
	// The ID of the GIT repository to assign the permissions
	RepositoryId *string `pulumi:"repositoryId"`
}

// The set of arguments for constructing a GitPermissions resource.
type GitPermissionsArgs struct {
	// The name of the branch to assign the permissions.
	BranchName pulumi.StringPtrInput
	// the permissions to assign. The follwing permissions are available
	Permissions pulumi.StringMapInput
	// The **group** principal to assign the permissions.
	Principal pulumi.StringInput
	// The ID of the project to assign the permissions.
	ProjectId pulumi.StringInput
	// Replace (`true`) or merge (`false`) the permissions. Default: `true`
	Replace pulumi.BoolPtrInput
	// The ID of the GIT repository to assign the permissions
	RepositoryId pulumi.StringPtrInput
}

func (GitPermissionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gitPermissionsArgs)(nil)).Elem()
}

type GitPermissionsInput interface {
	pulumi.Input

	ToGitPermissionsOutput() GitPermissionsOutput
	ToGitPermissionsOutputWithContext(ctx context.Context) GitPermissionsOutput
}

func (*GitPermissions) ElementType() reflect.Type {
	return reflect.TypeOf((**GitPermissions)(nil)).Elem()
}

func (i *GitPermissions) ToGitPermissionsOutput() GitPermissionsOutput {
	return i.ToGitPermissionsOutputWithContext(context.Background())
}

func (i *GitPermissions) ToGitPermissionsOutputWithContext(ctx context.Context) GitPermissionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitPermissionsOutput)
}

// GitPermissionsArrayInput is an input type that accepts GitPermissionsArray and GitPermissionsArrayOutput values.
// You can construct a concrete instance of `GitPermissionsArrayInput` via:
//
//	GitPermissionsArray{ GitPermissionsArgs{...} }
type GitPermissionsArrayInput interface {
	pulumi.Input

	ToGitPermissionsArrayOutput() GitPermissionsArrayOutput
	ToGitPermissionsArrayOutputWithContext(context.Context) GitPermissionsArrayOutput
}

type GitPermissionsArray []GitPermissionsInput

func (GitPermissionsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GitPermissions)(nil)).Elem()
}

func (i GitPermissionsArray) ToGitPermissionsArrayOutput() GitPermissionsArrayOutput {
	return i.ToGitPermissionsArrayOutputWithContext(context.Background())
}

func (i GitPermissionsArray) ToGitPermissionsArrayOutputWithContext(ctx context.Context) GitPermissionsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitPermissionsArrayOutput)
}

// GitPermissionsMapInput is an input type that accepts GitPermissionsMap and GitPermissionsMapOutput values.
// You can construct a concrete instance of `GitPermissionsMapInput` via:
//
//	GitPermissionsMap{ "key": GitPermissionsArgs{...} }
type GitPermissionsMapInput interface {
	pulumi.Input

	ToGitPermissionsMapOutput() GitPermissionsMapOutput
	ToGitPermissionsMapOutputWithContext(context.Context) GitPermissionsMapOutput
}

type GitPermissionsMap map[string]GitPermissionsInput

func (GitPermissionsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GitPermissions)(nil)).Elem()
}

func (i GitPermissionsMap) ToGitPermissionsMapOutput() GitPermissionsMapOutput {
	return i.ToGitPermissionsMapOutputWithContext(context.Background())
}

func (i GitPermissionsMap) ToGitPermissionsMapOutputWithContext(ctx context.Context) GitPermissionsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitPermissionsMapOutput)
}

type GitPermissionsOutput struct{ *pulumi.OutputState }

func (GitPermissionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GitPermissions)(nil)).Elem()
}

func (o GitPermissionsOutput) ToGitPermissionsOutput() GitPermissionsOutput {
	return o
}

func (o GitPermissionsOutput) ToGitPermissionsOutputWithContext(ctx context.Context) GitPermissionsOutput {
	return o
}

// The name of the branch to assign the permissions.
func (o GitPermissionsOutput) BranchName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitPermissions) pulumi.StringPtrOutput { return v.BranchName }).(pulumi.StringPtrOutput)
}

// the permissions to assign. The follwing permissions are available
func (o GitPermissionsOutput) Permissions() pulumi.StringMapOutput {
	return o.ApplyT(func(v *GitPermissions) pulumi.StringMapOutput { return v.Permissions }).(pulumi.StringMapOutput)
}

// The **group** principal to assign the permissions.
func (o GitPermissionsOutput) Principal() pulumi.StringOutput {
	return o.ApplyT(func(v *GitPermissions) pulumi.StringOutput { return v.Principal }).(pulumi.StringOutput)
}

// The ID of the project to assign the permissions.
func (o GitPermissionsOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *GitPermissions) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// Replace (`true`) or merge (`false`) the permissions. Default: `true`
func (o GitPermissionsOutput) Replace() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GitPermissions) pulumi.BoolPtrOutput { return v.Replace }).(pulumi.BoolPtrOutput)
}

// The ID of the GIT repository to assign the permissions
func (o GitPermissionsOutput) RepositoryId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitPermissions) pulumi.StringPtrOutput { return v.RepositoryId }).(pulumi.StringPtrOutput)
}

type GitPermissionsArrayOutput struct{ *pulumi.OutputState }

func (GitPermissionsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GitPermissions)(nil)).Elem()
}

func (o GitPermissionsArrayOutput) ToGitPermissionsArrayOutput() GitPermissionsArrayOutput {
	return o
}

func (o GitPermissionsArrayOutput) ToGitPermissionsArrayOutputWithContext(ctx context.Context) GitPermissionsArrayOutput {
	return o
}

func (o GitPermissionsArrayOutput) Index(i pulumi.IntInput) GitPermissionsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GitPermissions {
		return vs[0].([]*GitPermissions)[vs[1].(int)]
	}).(GitPermissionsOutput)
}

type GitPermissionsMapOutput struct{ *pulumi.OutputState }

func (GitPermissionsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GitPermissions)(nil)).Elem()
}

func (o GitPermissionsMapOutput) ToGitPermissionsMapOutput() GitPermissionsMapOutput {
	return o
}

func (o GitPermissionsMapOutput) ToGitPermissionsMapOutputWithContext(ctx context.Context) GitPermissionsMapOutput {
	return o
}

func (o GitPermissionsMapOutput) MapIndex(k pulumi.StringInput) GitPermissionsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GitPermissions {
		return vs[0].(map[string]*GitPermissions)[vs[1].(string)]
	}).(GitPermissionsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GitPermissionsInput)(nil)).Elem(), &GitPermissions{})
	pulumi.RegisterInputType(reflect.TypeOf((*GitPermissionsArrayInput)(nil)).Elem(), GitPermissionsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GitPermissionsMapInput)(nil)).Elem(), GitPermissionsMap{})
	pulumi.RegisterOutputType(GitPermissionsOutput{})
	pulumi.RegisterOutputType(GitPermissionsArrayOutput{})
	pulumi.RegisterOutputType(GitPermissionsMapOutput{})
}
