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

// Manages a Git Repository Branch.
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
//			})
//			if err != nil {
//				return err
//			}
//			exampleGit, err := azuredevops.NewGit(ctx, "example", &azuredevops.GitArgs{
//				ProjectId: example.ID(),
//				Name:      pulumi.String("Example Git Repository"),
//				Initialization: &azuredevops.GitInitializationArgs{
//					InitType: pulumi.String("Clean"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			exampleGitRepositoryBranch, err := azuredevops.NewGitRepositoryBranch(ctx, "example", &azuredevops.GitRepositoryBranchArgs{
//				RepositoryId: exampleGit.ID(),
//				Name:         pulumi.String("example-branch-name"),
//				RefBranch:    exampleGit.DefaultBranch,
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuredevops.NewGitRepositoryBranch(ctx, "example_from_commit_id", &azuredevops.GitRepositoryBranchArgs{
//				RepositoryId: exampleGit.ID(),
//				Name:         pulumi.String("example-from-commit-id"),
//				RefCommitId:  exampleGitRepositoryBranch.LastCommitId,
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type GitRepositoryBranch struct {
	pulumi.CustomResourceState

	// The commit object ID of last commit on the branch.
	LastCommitId pulumi.StringOutput `pulumi:"lastCommitId"`
	// The name of the branch in short format not prefixed with `refs/heads/`.
	Name pulumi.StringOutput `pulumi:"name"`
	// The reference to the source branch to create the branch from, in `<name>` or `refs/heads/<name>` format. Conflict with `refTag`, `refCommitId`.
	RefBranch pulumi.StringPtrOutput `pulumi:"refBranch"`
	// The commit object ID to create the branch from. Conflict with `refBranch`, `refTag`.
	RefCommitId pulumi.StringPtrOutput `pulumi:"refCommitId"`
	// The reference to the tag to create the branch from, in `<name>` or `refs/tags/<name>` format. Conflict with `refBranch`, `refCommitId`.
	RefTag pulumi.StringPtrOutput `pulumi:"refTag"`
	// The ID of the repository the branch is created against.
	RepositoryId pulumi.StringOutput `pulumi:"repositoryId"`
}

// NewGitRepositoryBranch registers a new resource with the given unique name, arguments, and options.
func NewGitRepositoryBranch(ctx *pulumi.Context,
	name string, args *GitRepositoryBranchArgs, opts ...pulumi.ResourceOption) (*GitRepositoryBranch, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RepositoryId == nil {
		return nil, errors.New("invalid value for required argument 'RepositoryId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource GitRepositoryBranch
	err := ctx.RegisterResource("azuredevops:index/gitRepositoryBranch:GitRepositoryBranch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGitRepositoryBranch gets an existing GitRepositoryBranch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGitRepositoryBranch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GitRepositoryBranchState, opts ...pulumi.ResourceOption) (*GitRepositoryBranch, error) {
	var resource GitRepositoryBranch
	err := ctx.ReadResource("azuredevops:index/gitRepositoryBranch:GitRepositoryBranch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GitRepositoryBranch resources.
type gitRepositoryBranchState struct {
	// The commit object ID of last commit on the branch.
	LastCommitId *string `pulumi:"lastCommitId"`
	// The name of the branch in short format not prefixed with `refs/heads/`.
	Name *string `pulumi:"name"`
	// The reference to the source branch to create the branch from, in `<name>` or `refs/heads/<name>` format. Conflict with `refTag`, `refCommitId`.
	RefBranch *string `pulumi:"refBranch"`
	// The commit object ID to create the branch from. Conflict with `refBranch`, `refTag`.
	RefCommitId *string `pulumi:"refCommitId"`
	// The reference to the tag to create the branch from, in `<name>` or `refs/tags/<name>` format. Conflict with `refBranch`, `refCommitId`.
	RefTag *string `pulumi:"refTag"`
	// The ID of the repository the branch is created against.
	RepositoryId *string `pulumi:"repositoryId"`
}

type GitRepositoryBranchState struct {
	// The commit object ID of last commit on the branch.
	LastCommitId pulumi.StringPtrInput
	// The name of the branch in short format not prefixed with `refs/heads/`.
	Name pulumi.StringPtrInput
	// The reference to the source branch to create the branch from, in `<name>` or `refs/heads/<name>` format. Conflict with `refTag`, `refCommitId`.
	RefBranch pulumi.StringPtrInput
	// The commit object ID to create the branch from. Conflict with `refBranch`, `refTag`.
	RefCommitId pulumi.StringPtrInput
	// The reference to the tag to create the branch from, in `<name>` or `refs/tags/<name>` format. Conflict with `refBranch`, `refCommitId`.
	RefTag pulumi.StringPtrInput
	// The ID of the repository the branch is created against.
	RepositoryId pulumi.StringPtrInput
}

func (GitRepositoryBranchState) ElementType() reflect.Type {
	return reflect.TypeOf((*gitRepositoryBranchState)(nil)).Elem()
}

type gitRepositoryBranchArgs struct {
	// The name of the branch in short format not prefixed with `refs/heads/`.
	Name *string `pulumi:"name"`
	// The reference to the source branch to create the branch from, in `<name>` or `refs/heads/<name>` format. Conflict with `refTag`, `refCommitId`.
	RefBranch *string `pulumi:"refBranch"`
	// The commit object ID to create the branch from. Conflict with `refBranch`, `refTag`.
	RefCommitId *string `pulumi:"refCommitId"`
	// The reference to the tag to create the branch from, in `<name>` or `refs/tags/<name>` format. Conflict with `refBranch`, `refCommitId`.
	RefTag *string `pulumi:"refTag"`
	// The ID of the repository the branch is created against.
	RepositoryId string `pulumi:"repositoryId"`
}

// The set of arguments for constructing a GitRepositoryBranch resource.
type GitRepositoryBranchArgs struct {
	// The name of the branch in short format not prefixed with `refs/heads/`.
	Name pulumi.StringPtrInput
	// The reference to the source branch to create the branch from, in `<name>` or `refs/heads/<name>` format. Conflict with `refTag`, `refCommitId`.
	RefBranch pulumi.StringPtrInput
	// The commit object ID to create the branch from. Conflict with `refBranch`, `refTag`.
	RefCommitId pulumi.StringPtrInput
	// The reference to the tag to create the branch from, in `<name>` or `refs/tags/<name>` format. Conflict with `refBranch`, `refCommitId`.
	RefTag pulumi.StringPtrInput
	// The ID of the repository the branch is created against.
	RepositoryId pulumi.StringInput
}

func (GitRepositoryBranchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gitRepositoryBranchArgs)(nil)).Elem()
}

type GitRepositoryBranchInput interface {
	pulumi.Input

	ToGitRepositoryBranchOutput() GitRepositoryBranchOutput
	ToGitRepositoryBranchOutputWithContext(ctx context.Context) GitRepositoryBranchOutput
}

func (*GitRepositoryBranch) ElementType() reflect.Type {
	return reflect.TypeOf((**GitRepositoryBranch)(nil)).Elem()
}

func (i *GitRepositoryBranch) ToGitRepositoryBranchOutput() GitRepositoryBranchOutput {
	return i.ToGitRepositoryBranchOutputWithContext(context.Background())
}

func (i *GitRepositoryBranch) ToGitRepositoryBranchOutputWithContext(ctx context.Context) GitRepositoryBranchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitRepositoryBranchOutput)
}

// GitRepositoryBranchArrayInput is an input type that accepts GitRepositoryBranchArray and GitRepositoryBranchArrayOutput values.
// You can construct a concrete instance of `GitRepositoryBranchArrayInput` via:
//
//	GitRepositoryBranchArray{ GitRepositoryBranchArgs{...} }
type GitRepositoryBranchArrayInput interface {
	pulumi.Input

	ToGitRepositoryBranchArrayOutput() GitRepositoryBranchArrayOutput
	ToGitRepositoryBranchArrayOutputWithContext(context.Context) GitRepositoryBranchArrayOutput
}

type GitRepositoryBranchArray []GitRepositoryBranchInput

func (GitRepositoryBranchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GitRepositoryBranch)(nil)).Elem()
}

func (i GitRepositoryBranchArray) ToGitRepositoryBranchArrayOutput() GitRepositoryBranchArrayOutput {
	return i.ToGitRepositoryBranchArrayOutputWithContext(context.Background())
}

func (i GitRepositoryBranchArray) ToGitRepositoryBranchArrayOutputWithContext(ctx context.Context) GitRepositoryBranchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitRepositoryBranchArrayOutput)
}

// GitRepositoryBranchMapInput is an input type that accepts GitRepositoryBranchMap and GitRepositoryBranchMapOutput values.
// You can construct a concrete instance of `GitRepositoryBranchMapInput` via:
//
//	GitRepositoryBranchMap{ "key": GitRepositoryBranchArgs{...} }
type GitRepositoryBranchMapInput interface {
	pulumi.Input

	ToGitRepositoryBranchMapOutput() GitRepositoryBranchMapOutput
	ToGitRepositoryBranchMapOutputWithContext(context.Context) GitRepositoryBranchMapOutput
}

type GitRepositoryBranchMap map[string]GitRepositoryBranchInput

func (GitRepositoryBranchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GitRepositoryBranch)(nil)).Elem()
}

func (i GitRepositoryBranchMap) ToGitRepositoryBranchMapOutput() GitRepositoryBranchMapOutput {
	return i.ToGitRepositoryBranchMapOutputWithContext(context.Background())
}

func (i GitRepositoryBranchMap) ToGitRepositoryBranchMapOutputWithContext(ctx context.Context) GitRepositoryBranchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitRepositoryBranchMapOutput)
}

type GitRepositoryBranchOutput struct{ *pulumi.OutputState }

func (GitRepositoryBranchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GitRepositoryBranch)(nil)).Elem()
}

func (o GitRepositoryBranchOutput) ToGitRepositoryBranchOutput() GitRepositoryBranchOutput {
	return o
}

func (o GitRepositoryBranchOutput) ToGitRepositoryBranchOutputWithContext(ctx context.Context) GitRepositoryBranchOutput {
	return o
}

// The commit object ID of last commit on the branch.
func (o GitRepositoryBranchOutput) LastCommitId() pulumi.StringOutput {
	return o.ApplyT(func(v *GitRepositoryBranch) pulumi.StringOutput { return v.LastCommitId }).(pulumi.StringOutput)
}

// The name of the branch in short format not prefixed with `refs/heads/`.
func (o GitRepositoryBranchOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *GitRepositoryBranch) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The reference to the source branch to create the branch from, in `<name>` or `refs/heads/<name>` format. Conflict with `refTag`, `refCommitId`.
func (o GitRepositoryBranchOutput) RefBranch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitRepositoryBranch) pulumi.StringPtrOutput { return v.RefBranch }).(pulumi.StringPtrOutput)
}

// The commit object ID to create the branch from. Conflict with `refBranch`, `refTag`.
func (o GitRepositoryBranchOutput) RefCommitId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitRepositoryBranch) pulumi.StringPtrOutput { return v.RefCommitId }).(pulumi.StringPtrOutput)
}

// The reference to the tag to create the branch from, in `<name>` or `refs/tags/<name>` format. Conflict with `refBranch`, `refCommitId`.
func (o GitRepositoryBranchOutput) RefTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitRepositoryBranch) pulumi.StringPtrOutput { return v.RefTag }).(pulumi.StringPtrOutput)
}

// The ID of the repository the branch is created against.
func (o GitRepositoryBranchOutput) RepositoryId() pulumi.StringOutput {
	return o.ApplyT(func(v *GitRepositoryBranch) pulumi.StringOutput { return v.RepositoryId }).(pulumi.StringOutput)
}

type GitRepositoryBranchArrayOutput struct{ *pulumi.OutputState }

func (GitRepositoryBranchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GitRepositoryBranch)(nil)).Elem()
}

func (o GitRepositoryBranchArrayOutput) ToGitRepositoryBranchArrayOutput() GitRepositoryBranchArrayOutput {
	return o
}

func (o GitRepositoryBranchArrayOutput) ToGitRepositoryBranchArrayOutputWithContext(ctx context.Context) GitRepositoryBranchArrayOutput {
	return o
}

func (o GitRepositoryBranchArrayOutput) Index(i pulumi.IntInput) GitRepositoryBranchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GitRepositoryBranch {
		return vs[0].([]*GitRepositoryBranch)[vs[1].(int)]
	}).(GitRepositoryBranchOutput)
}

type GitRepositoryBranchMapOutput struct{ *pulumi.OutputState }

func (GitRepositoryBranchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GitRepositoryBranch)(nil)).Elem()
}

func (o GitRepositoryBranchMapOutput) ToGitRepositoryBranchMapOutput() GitRepositoryBranchMapOutput {
	return o
}

func (o GitRepositoryBranchMapOutput) ToGitRepositoryBranchMapOutputWithContext(ctx context.Context) GitRepositoryBranchMapOutput {
	return o
}

func (o GitRepositoryBranchMapOutput) MapIndex(k pulumi.StringInput) GitRepositoryBranchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GitRepositoryBranch {
		return vs[0].(map[string]*GitRepositoryBranch)[vs[1].(string)]
	}).(GitRepositoryBranchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GitRepositoryBranchInput)(nil)).Elem(), &GitRepositoryBranch{})
	pulumi.RegisterInputType(reflect.TypeOf((*GitRepositoryBranchArrayInput)(nil)).Elem(), GitRepositoryBranchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GitRepositoryBranchMapInput)(nil)).Elem(), GitRepositoryBranchMap{})
	pulumi.RegisterOutputType(GitRepositoryBranchOutput{})
	pulumi.RegisterOutputType(GitRepositoryBranchArrayOutput{})
	pulumi.RegisterOutputType(GitRepositoryBranchMapOutput{})
}
