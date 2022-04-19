// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manage files within an Azure DevOps Git repository.
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
// 		project, err := azuredevops.NewProject(ctx, "project", &azuredevops.ProjectArgs{
// 			Visibility:       pulumi.String("private"),
// 			VersionControl:   pulumi.String("Git"),
// 			WorkItemTemplate: pulumi.String("Agile"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		repo, err := azuredevops.NewGit(ctx, "repo", &azuredevops.GitArgs{
// 			ProjectId: project.ID(),
// 			Initialization: &GitInitializationArgs{
// 				InitType: pulumi.String("Clean"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = azuredevops.NewGitRepositoryFile(ctx, "repoFile", &azuredevops.GitRepositoryFileArgs{
// 			RepositoryId:      repo.ID(),
// 			File:              pulumi.String(".gitignore"),
// 			Content:           pulumi.String("**/*.tfstate"),
// 			Branch:            pulumi.String("refs/heads/master"),
// 			CommitMessage:     pulumi.String("First commit"),
// 			OverwriteOnCreate: pulumi.Bool(false),
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
// - [Azure DevOps Service REST API 5.1 - Git API](https://docs.microsoft.com/en-us/rest/api/azure/devops/git/?view=azure-devops-rest-5.1)
//
// ## Import
//
// Repository files can be imported using a combination of the `repositroy ID` and `file`, e.g.
//
// ```sh
//  $ pulumi import azuredevops:index/gitRepositoryFile:GitRepositoryFile repo_file 00000000-0000-0000-0000-000000000000/.gitignore
// ```
//
//  To import a file from a branch other than `master`, append `:` and the branch name, e.g.
//
// ```sh
//  $ pulumi import azuredevops:index/gitRepositoryFile:GitRepositoryFile repo_file 00000000-0000-0000-0000-000000000000/.gitignore:refs/heads/master
// ```
type GitRepositoryFile struct {
	pulumi.CustomResourceState

	// Git branch (defaults to `refs/heads/master`). The branch must already exist, it will not be created if it
	// does not already exist.
	Branch pulumi.StringPtrOutput `pulumi:"branch"`
	// Commit message when adding or updating the managed file.
	CommitMessage pulumi.StringOutput `pulumi:"commitMessage"`
	// The file content.
	Content pulumi.StringOutput `pulumi:"content"`
	// The path of the file to manage.
	File pulumi.StringOutput `pulumi:"file"`
	// Enable overwriting existing files (defaults to `false`).
	OverwriteOnCreate pulumi.BoolPtrOutput `pulumi:"overwriteOnCreate"`
	// The ID of the Git repository.
	RepositoryId pulumi.StringOutput `pulumi:"repositoryId"`
}

// NewGitRepositoryFile registers a new resource with the given unique name, arguments, and options.
func NewGitRepositoryFile(ctx *pulumi.Context,
	name string, args *GitRepositoryFileArgs, opts ...pulumi.ResourceOption) (*GitRepositoryFile, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Content == nil {
		return nil, errors.New("invalid value for required argument 'Content'")
	}
	if args.File == nil {
		return nil, errors.New("invalid value for required argument 'File'")
	}
	if args.RepositoryId == nil {
		return nil, errors.New("invalid value for required argument 'RepositoryId'")
	}
	var resource GitRepositoryFile
	err := ctx.RegisterResource("azuredevops:index/gitRepositoryFile:GitRepositoryFile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGitRepositoryFile gets an existing GitRepositoryFile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGitRepositoryFile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GitRepositoryFileState, opts ...pulumi.ResourceOption) (*GitRepositoryFile, error) {
	var resource GitRepositoryFile
	err := ctx.ReadResource("azuredevops:index/gitRepositoryFile:GitRepositoryFile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GitRepositoryFile resources.
type gitRepositoryFileState struct {
	// Git branch (defaults to `refs/heads/master`). The branch must already exist, it will not be created if it
	// does not already exist.
	Branch *string `pulumi:"branch"`
	// Commit message when adding or updating the managed file.
	CommitMessage *string `pulumi:"commitMessage"`
	// The file content.
	Content *string `pulumi:"content"`
	// The path of the file to manage.
	File *string `pulumi:"file"`
	// Enable overwriting existing files (defaults to `false`).
	OverwriteOnCreate *bool `pulumi:"overwriteOnCreate"`
	// The ID of the Git repository.
	RepositoryId *string `pulumi:"repositoryId"`
}

type GitRepositoryFileState struct {
	// Git branch (defaults to `refs/heads/master`). The branch must already exist, it will not be created if it
	// does not already exist.
	Branch pulumi.StringPtrInput
	// Commit message when adding or updating the managed file.
	CommitMessage pulumi.StringPtrInput
	// The file content.
	Content pulumi.StringPtrInput
	// The path of the file to manage.
	File pulumi.StringPtrInput
	// Enable overwriting existing files (defaults to `false`).
	OverwriteOnCreate pulumi.BoolPtrInput
	// The ID of the Git repository.
	RepositoryId pulumi.StringPtrInput
}

func (GitRepositoryFileState) ElementType() reflect.Type {
	return reflect.TypeOf((*gitRepositoryFileState)(nil)).Elem()
}

type gitRepositoryFileArgs struct {
	// Git branch (defaults to `refs/heads/master`). The branch must already exist, it will not be created if it
	// does not already exist.
	Branch *string `pulumi:"branch"`
	// Commit message when adding or updating the managed file.
	CommitMessage *string `pulumi:"commitMessage"`
	// The file content.
	Content string `pulumi:"content"`
	// The path of the file to manage.
	File string `pulumi:"file"`
	// Enable overwriting existing files (defaults to `false`).
	OverwriteOnCreate *bool `pulumi:"overwriteOnCreate"`
	// The ID of the Git repository.
	RepositoryId string `pulumi:"repositoryId"`
}

// The set of arguments for constructing a GitRepositoryFile resource.
type GitRepositoryFileArgs struct {
	// Git branch (defaults to `refs/heads/master`). The branch must already exist, it will not be created if it
	// does not already exist.
	Branch pulumi.StringPtrInput
	// Commit message when adding or updating the managed file.
	CommitMessage pulumi.StringPtrInput
	// The file content.
	Content pulumi.StringInput
	// The path of the file to manage.
	File pulumi.StringInput
	// Enable overwriting existing files (defaults to `false`).
	OverwriteOnCreate pulumi.BoolPtrInput
	// The ID of the Git repository.
	RepositoryId pulumi.StringInput
}

func (GitRepositoryFileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gitRepositoryFileArgs)(nil)).Elem()
}

type GitRepositoryFileInput interface {
	pulumi.Input

	ToGitRepositoryFileOutput() GitRepositoryFileOutput
	ToGitRepositoryFileOutputWithContext(ctx context.Context) GitRepositoryFileOutput
}

func (*GitRepositoryFile) ElementType() reflect.Type {
	return reflect.TypeOf((**GitRepositoryFile)(nil)).Elem()
}

func (i *GitRepositoryFile) ToGitRepositoryFileOutput() GitRepositoryFileOutput {
	return i.ToGitRepositoryFileOutputWithContext(context.Background())
}

func (i *GitRepositoryFile) ToGitRepositoryFileOutputWithContext(ctx context.Context) GitRepositoryFileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitRepositoryFileOutput)
}

// GitRepositoryFileArrayInput is an input type that accepts GitRepositoryFileArray and GitRepositoryFileArrayOutput values.
// You can construct a concrete instance of `GitRepositoryFileArrayInput` via:
//
//          GitRepositoryFileArray{ GitRepositoryFileArgs{...} }
type GitRepositoryFileArrayInput interface {
	pulumi.Input

	ToGitRepositoryFileArrayOutput() GitRepositoryFileArrayOutput
	ToGitRepositoryFileArrayOutputWithContext(context.Context) GitRepositoryFileArrayOutput
}

type GitRepositoryFileArray []GitRepositoryFileInput

func (GitRepositoryFileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GitRepositoryFile)(nil)).Elem()
}

func (i GitRepositoryFileArray) ToGitRepositoryFileArrayOutput() GitRepositoryFileArrayOutput {
	return i.ToGitRepositoryFileArrayOutputWithContext(context.Background())
}

func (i GitRepositoryFileArray) ToGitRepositoryFileArrayOutputWithContext(ctx context.Context) GitRepositoryFileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitRepositoryFileArrayOutput)
}

// GitRepositoryFileMapInput is an input type that accepts GitRepositoryFileMap and GitRepositoryFileMapOutput values.
// You can construct a concrete instance of `GitRepositoryFileMapInput` via:
//
//          GitRepositoryFileMap{ "key": GitRepositoryFileArgs{...} }
type GitRepositoryFileMapInput interface {
	pulumi.Input

	ToGitRepositoryFileMapOutput() GitRepositoryFileMapOutput
	ToGitRepositoryFileMapOutputWithContext(context.Context) GitRepositoryFileMapOutput
}

type GitRepositoryFileMap map[string]GitRepositoryFileInput

func (GitRepositoryFileMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GitRepositoryFile)(nil)).Elem()
}

func (i GitRepositoryFileMap) ToGitRepositoryFileMapOutput() GitRepositoryFileMapOutput {
	return i.ToGitRepositoryFileMapOutputWithContext(context.Background())
}

func (i GitRepositoryFileMap) ToGitRepositoryFileMapOutputWithContext(ctx context.Context) GitRepositoryFileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitRepositoryFileMapOutput)
}

type GitRepositoryFileOutput struct{ *pulumi.OutputState }

func (GitRepositoryFileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GitRepositoryFile)(nil)).Elem()
}

func (o GitRepositoryFileOutput) ToGitRepositoryFileOutput() GitRepositoryFileOutput {
	return o
}

func (o GitRepositoryFileOutput) ToGitRepositoryFileOutputWithContext(ctx context.Context) GitRepositoryFileOutput {
	return o
}

type GitRepositoryFileArrayOutput struct{ *pulumi.OutputState }

func (GitRepositoryFileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GitRepositoryFile)(nil)).Elem()
}

func (o GitRepositoryFileArrayOutput) ToGitRepositoryFileArrayOutput() GitRepositoryFileArrayOutput {
	return o
}

func (o GitRepositoryFileArrayOutput) ToGitRepositoryFileArrayOutputWithContext(ctx context.Context) GitRepositoryFileArrayOutput {
	return o
}

func (o GitRepositoryFileArrayOutput) Index(i pulumi.IntInput) GitRepositoryFileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GitRepositoryFile {
		return vs[0].([]*GitRepositoryFile)[vs[1].(int)]
	}).(GitRepositoryFileOutput)
}

type GitRepositoryFileMapOutput struct{ *pulumi.OutputState }

func (GitRepositoryFileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GitRepositoryFile)(nil)).Elem()
}

func (o GitRepositoryFileMapOutput) ToGitRepositoryFileMapOutput() GitRepositoryFileMapOutput {
	return o
}

func (o GitRepositoryFileMapOutput) ToGitRepositoryFileMapOutputWithContext(ctx context.Context) GitRepositoryFileMapOutput {
	return o
}

func (o GitRepositoryFileMapOutput) MapIndex(k pulumi.StringInput) GitRepositoryFileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GitRepositoryFile {
		return vs[0].(map[string]*GitRepositoryFile)[vs[1].(string)]
	}).(GitRepositoryFileOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GitRepositoryFileInput)(nil)).Elem(), &GitRepositoryFile{})
	pulumi.RegisterInputType(reflect.TypeOf((*GitRepositoryFileArrayInput)(nil)).Elem(), GitRepositoryFileArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GitRepositoryFileMapInput)(nil)).Elem(), GitRepositoryFileMap{})
	pulumi.RegisterOutputType(GitRepositoryFileOutput{})
	pulumi.RegisterOutputType(GitRepositoryFileArrayOutput{})
	pulumi.RegisterOutputType(GitRepositoryFileMapOutput{})
}
