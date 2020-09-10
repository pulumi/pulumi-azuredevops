// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a git repository within Azure DevOps.
//
// ## Example Usage
// ### Create Git repository
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
// 		project, err := azuredevops.NewProject(ctx, "project", &azuredevops.ProjectArgs{
// 			Visibility:       pulumi.String("private"),
// 			VersionControl:   pulumi.String("Git"),
// 			WorkItemTemplate: pulumi.String("Agile"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = azuredevops.NewGit(ctx, "repo", &azuredevops.GitArgs{
// 			ProjectId: project.ID(),
// 			Initialization: &azuredevops.GitInitializationArgs{
// 				InitType: pulumi.String("Clean"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Create Fork of another Azure DevOps Git repository
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
// 		_, err := azuredevops.NewGit(ctx, "repo", &azuredevops.GitArgs{
// 			ProjectId:          pulumi.Any(azuredevops_project.Project.Id),
// 			ParentRepositoryId: pulumi.Any(azuredevops_git_repository.Parent.Id),
// 			Initialization: &azuredevops.GitInitializationArgs{
// 				InitType: pulumi.String("Clean"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Create Import from another Git repository
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
// 		_, err := azuredevops.NewGit(ctx, "repo", &azuredevops.GitArgs{
// 			ProjectId: pulumi.Any(azuredevops_project.Project.Id),
// 			Initialization: &azuredevops.GitInitializationArgs{
// 				InitType:   pulumi.String("Import"),
// 				SourceType: pulumi.String("Git"),
// 				SourceUrl:  pulumi.String("https://github.com/microsoft/terraform-provider-azuredevops.git"),
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
// - [Azure DevOps Service REST API 5.1 - Agent Pools](https://docs.microsoft.com/en-us/rest/api/azure/devops/git/repositories?view=azure-devops-rest-5.1)
type Git struct {
	pulumi.CustomResourceState

	// The ref of the default branch. Will be used as the branch name for initialized repositories.
	DefaultBranch pulumi.StringOutput `pulumi:"defaultBranch"`
	// An `initialization` block as documented below.
	Initialization GitInitializationOutput `pulumi:"initialization"`
	// True if the repository was created as a fork.
	IsFork pulumi.BoolOutput `pulumi:"isFork"`
	// The name of the git repository.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of a Git project from which a fork is to be created.
	ParentRepositoryId pulumi.StringPtrOutput `pulumi:"parentRepositoryId"`
	// The project ID or project name.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Git HTTPS URL of the repository
	RemoteUrl pulumi.StringOutput `pulumi:"remoteUrl"`
	// Size in bytes.
	Size pulumi.IntOutput `pulumi:"size"`
	// Git SSH URL of the repository.
	SshUrl pulumi.StringOutput `pulumi:"sshUrl"`
	// REST API URL of the repository.
	Url pulumi.StringOutput `pulumi:"url"`
	// Web link to the repository.
	WebUrl pulumi.StringOutput `pulumi:"webUrl"`
}

// NewGit registers a new resource with the given unique name, arguments, and options.
func NewGit(ctx *pulumi.Context,
	name string, args *GitArgs, opts ...pulumi.ResourceOption) (*Git, error) {
	if args == nil || args.Initialization == nil {
		return nil, errors.New("missing required argument 'Initialization'")
	}
	if args == nil || args.ProjectId == nil {
		return nil, errors.New("missing required argument 'ProjectId'")
	}
	if args == nil {
		args = &GitArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azuredevops:Repository/git:Git"),
		},
	})
	opts = append(opts, aliases)
	var resource Git
	err := ctx.RegisterResource("azuredevops:index/git:Git", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGit gets an existing Git resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGit(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GitState, opts ...pulumi.ResourceOption) (*Git, error) {
	var resource Git
	err := ctx.ReadResource("azuredevops:index/git:Git", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Git resources.
type gitState struct {
	// The ref of the default branch. Will be used as the branch name for initialized repositories.
	DefaultBranch *string `pulumi:"defaultBranch"`
	// An `initialization` block as documented below.
	Initialization *GitInitialization `pulumi:"initialization"`
	// True if the repository was created as a fork.
	IsFork *bool `pulumi:"isFork"`
	// The name of the git repository.
	Name *string `pulumi:"name"`
	// The ID of a Git project from which a fork is to be created.
	ParentRepositoryId *string `pulumi:"parentRepositoryId"`
	// The project ID or project name.
	ProjectId *string `pulumi:"projectId"`
	// Git HTTPS URL of the repository
	RemoteUrl *string `pulumi:"remoteUrl"`
	// Size in bytes.
	Size *int `pulumi:"size"`
	// Git SSH URL of the repository.
	SshUrl *string `pulumi:"sshUrl"`
	// REST API URL of the repository.
	Url *string `pulumi:"url"`
	// Web link to the repository.
	WebUrl *string `pulumi:"webUrl"`
}

type GitState struct {
	// The ref of the default branch. Will be used as the branch name for initialized repositories.
	DefaultBranch pulumi.StringPtrInput
	// An `initialization` block as documented below.
	Initialization GitInitializationPtrInput
	// True if the repository was created as a fork.
	IsFork pulumi.BoolPtrInput
	// The name of the git repository.
	Name pulumi.StringPtrInput
	// The ID of a Git project from which a fork is to be created.
	ParentRepositoryId pulumi.StringPtrInput
	// The project ID or project name.
	ProjectId pulumi.StringPtrInput
	// Git HTTPS URL of the repository
	RemoteUrl pulumi.StringPtrInput
	// Size in bytes.
	Size pulumi.IntPtrInput
	// Git SSH URL of the repository.
	SshUrl pulumi.StringPtrInput
	// REST API URL of the repository.
	Url pulumi.StringPtrInput
	// Web link to the repository.
	WebUrl pulumi.StringPtrInput
}

func (GitState) ElementType() reflect.Type {
	return reflect.TypeOf((*gitState)(nil)).Elem()
}

type gitArgs struct {
	// The ref of the default branch. Will be used as the branch name for initialized repositories.
	DefaultBranch *string `pulumi:"defaultBranch"`
	// An `initialization` block as documented below.
	Initialization GitInitialization `pulumi:"initialization"`
	// The name of the git repository.
	Name *string `pulumi:"name"`
	// The ID of a Git project from which a fork is to be created.
	ParentRepositoryId *string `pulumi:"parentRepositoryId"`
	// The project ID or project name.
	ProjectId string `pulumi:"projectId"`
}

// The set of arguments for constructing a Git resource.
type GitArgs struct {
	// The ref of the default branch. Will be used as the branch name for initialized repositories.
	DefaultBranch pulumi.StringPtrInput
	// An `initialization` block as documented below.
	Initialization GitInitializationInput
	// The name of the git repository.
	Name pulumi.StringPtrInput
	// The ID of a Git project from which a fork is to be created.
	ParentRepositoryId pulumi.StringPtrInput
	// The project ID or project name.
	ProjectId pulumi.StringInput
}

func (GitArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gitArgs)(nil)).Elem()
}
