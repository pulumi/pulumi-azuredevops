// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package repository

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type GitInitialization struct {
	// The type of repository to create. Valid values: `Uninitialized`, `Clean` or `Import`.
	InitType string `pulumi:"initType"`
	// The id of service connection used to authenticate to a private repository for import initialization.
	ServiceConnectionId *string `pulumi:"serviceConnectionId"`
	// Type of the source repository. Used if the `initType` is `Import`. Valid values: `Git`.
	SourceType *string `pulumi:"sourceType"`
	// The URL of the source repository. Used if the `initType` is `Import`.
	SourceUrl *string `pulumi:"sourceUrl"`
}

// GitInitializationInput is an input type that accepts GitInitializationArgs and GitInitializationOutput values.
// You can construct a concrete instance of `GitInitializationInput` via:
//
//	GitInitializationArgs{...}
type GitInitializationInput interface {
	pulumi.Input

	ToGitInitializationOutput() GitInitializationOutput
	ToGitInitializationOutputWithContext(context.Context) GitInitializationOutput
}

type GitInitializationArgs struct {
	// The type of repository to create. Valid values: `Uninitialized`, `Clean` or `Import`.
	InitType pulumi.StringInput `pulumi:"initType"`
	// The id of service connection used to authenticate to a private repository for import initialization.
	ServiceConnectionId pulumi.StringPtrInput `pulumi:"serviceConnectionId"`
	// Type of the source repository. Used if the `initType` is `Import`. Valid values: `Git`.
	SourceType pulumi.StringPtrInput `pulumi:"sourceType"`
	// The URL of the source repository. Used if the `initType` is `Import`.
	SourceUrl pulumi.StringPtrInput `pulumi:"sourceUrl"`
}

func (GitInitializationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GitInitialization)(nil)).Elem()
}

func (i GitInitializationArgs) ToGitInitializationOutput() GitInitializationOutput {
	return i.ToGitInitializationOutputWithContext(context.Background())
}

func (i GitInitializationArgs) ToGitInitializationOutputWithContext(ctx context.Context) GitInitializationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitInitializationOutput)
}

func (i GitInitializationArgs) ToGitInitializationPtrOutput() GitInitializationPtrOutput {
	return i.ToGitInitializationPtrOutputWithContext(context.Background())
}

func (i GitInitializationArgs) ToGitInitializationPtrOutputWithContext(ctx context.Context) GitInitializationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitInitializationOutput).ToGitInitializationPtrOutputWithContext(ctx)
}

// GitInitializationPtrInput is an input type that accepts GitInitializationArgs, GitInitializationPtr and GitInitializationPtrOutput values.
// You can construct a concrete instance of `GitInitializationPtrInput` via:
//
//	        GitInitializationArgs{...}
//
//	or:
//
//	        nil
type GitInitializationPtrInput interface {
	pulumi.Input

	ToGitInitializationPtrOutput() GitInitializationPtrOutput
	ToGitInitializationPtrOutputWithContext(context.Context) GitInitializationPtrOutput
}

type gitInitializationPtrType GitInitializationArgs

func GitInitializationPtr(v *GitInitializationArgs) GitInitializationPtrInput {
	return (*gitInitializationPtrType)(v)
}

func (*gitInitializationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GitInitialization)(nil)).Elem()
}

func (i *gitInitializationPtrType) ToGitInitializationPtrOutput() GitInitializationPtrOutput {
	return i.ToGitInitializationPtrOutputWithContext(context.Background())
}

func (i *gitInitializationPtrType) ToGitInitializationPtrOutputWithContext(ctx context.Context) GitInitializationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitInitializationPtrOutput)
}

type GitInitializationOutput struct{ *pulumi.OutputState }

func (GitInitializationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GitInitialization)(nil)).Elem()
}

func (o GitInitializationOutput) ToGitInitializationOutput() GitInitializationOutput {
	return o
}

func (o GitInitializationOutput) ToGitInitializationOutputWithContext(ctx context.Context) GitInitializationOutput {
	return o
}

func (o GitInitializationOutput) ToGitInitializationPtrOutput() GitInitializationPtrOutput {
	return o.ToGitInitializationPtrOutputWithContext(context.Background())
}

func (o GitInitializationOutput) ToGitInitializationPtrOutputWithContext(ctx context.Context) GitInitializationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GitInitialization) *GitInitialization {
		return &v
	}).(GitInitializationPtrOutput)
}

// The type of repository to create. Valid values: `Uninitialized`, `Clean` or `Import`.
func (o GitInitializationOutput) InitType() pulumi.StringOutput {
	return o.ApplyT(func(v GitInitialization) string { return v.InitType }).(pulumi.StringOutput)
}

// The id of service connection used to authenticate to a private repository for import initialization.
func (o GitInitializationOutput) ServiceConnectionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitInitialization) *string { return v.ServiceConnectionId }).(pulumi.StringPtrOutput)
}

// Type of the source repository. Used if the `initType` is `Import`. Valid values: `Git`.
func (o GitInitializationOutput) SourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitInitialization) *string { return v.SourceType }).(pulumi.StringPtrOutput)
}

// The URL of the source repository. Used if the `initType` is `Import`.
func (o GitInitializationOutput) SourceUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitInitialization) *string { return v.SourceUrl }).(pulumi.StringPtrOutput)
}

type GitInitializationPtrOutput struct{ *pulumi.OutputState }

func (GitInitializationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GitInitialization)(nil)).Elem()
}

func (o GitInitializationPtrOutput) ToGitInitializationPtrOutput() GitInitializationPtrOutput {
	return o
}

func (o GitInitializationPtrOutput) ToGitInitializationPtrOutputWithContext(ctx context.Context) GitInitializationPtrOutput {
	return o
}

func (o GitInitializationPtrOutput) Elem() GitInitializationOutput {
	return o.ApplyT(func(v *GitInitialization) GitInitialization {
		if v != nil {
			return *v
		}
		var ret GitInitialization
		return ret
	}).(GitInitializationOutput)
}

// The type of repository to create. Valid values: `Uninitialized`, `Clean` or `Import`.
func (o GitInitializationPtrOutput) InitType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitInitialization) *string {
		if v == nil {
			return nil
		}
		return &v.InitType
	}).(pulumi.StringPtrOutput)
}

// The id of service connection used to authenticate to a private repository for import initialization.
func (o GitInitializationPtrOutput) ServiceConnectionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitInitialization) *string {
		if v == nil {
			return nil
		}
		return v.ServiceConnectionId
	}).(pulumi.StringPtrOutput)
}

// Type of the source repository. Used if the `initType` is `Import`. Valid values: `Git`.
func (o GitInitializationPtrOutput) SourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitInitialization) *string {
		if v == nil {
			return nil
		}
		return v.SourceType
	}).(pulumi.StringPtrOutput)
}

// The URL of the source repository. Used if the `initType` is `Import`.
func (o GitInitializationPtrOutput) SourceUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitInitialization) *string {
		if v == nil {
			return nil
		}
		return v.SourceUrl
	}).(pulumi.StringPtrOutput)
}

type GetRepositoriesRepository struct {
	// The ref of the default branch.
	DefaultBranch string `pulumi:"defaultBranch"`
	// Git repository identifier.
	Id string `pulumi:"id"`
	// Name of the Git repository to retrieve; requires `projectId` to be specified as well
	Name string `pulumi:"name"`
	// ID of project to list Git repositories
	ProjectId string `pulumi:"projectId"`
	// HTTPS Url to clone the Git repository
	RemoteUrl string `pulumi:"remoteUrl"`
	// Compressed size (bytes) of the repository.
	Size int `pulumi:"size"`
	// SSH Url to clone the Git repository
	SshUrl string `pulumi:"sshUrl"`
	// Details REST API endpoint for the Git Repository.
	Url string `pulumi:"url"`
	// Url of the Git repository web view
	WebUrl string `pulumi:"webUrl"`
}

// GetRepositoriesRepositoryInput is an input type that accepts GetRepositoriesRepositoryArgs and GetRepositoriesRepositoryOutput values.
// You can construct a concrete instance of `GetRepositoriesRepositoryInput` via:
//
//	GetRepositoriesRepositoryArgs{...}
type GetRepositoriesRepositoryInput interface {
	pulumi.Input

	ToGetRepositoriesRepositoryOutput() GetRepositoriesRepositoryOutput
	ToGetRepositoriesRepositoryOutputWithContext(context.Context) GetRepositoriesRepositoryOutput
}

type GetRepositoriesRepositoryArgs struct {
	// The ref of the default branch.
	DefaultBranch pulumi.StringInput `pulumi:"defaultBranch"`
	// Git repository identifier.
	Id pulumi.StringInput `pulumi:"id"`
	// Name of the Git repository to retrieve; requires `projectId` to be specified as well
	Name pulumi.StringInput `pulumi:"name"`
	// ID of project to list Git repositories
	ProjectId pulumi.StringInput `pulumi:"projectId"`
	// HTTPS Url to clone the Git repository
	RemoteUrl pulumi.StringInput `pulumi:"remoteUrl"`
	// Compressed size (bytes) of the repository.
	Size pulumi.IntInput `pulumi:"size"`
	// SSH Url to clone the Git repository
	SshUrl pulumi.StringInput `pulumi:"sshUrl"`
	// Details REST API endpoint for the Git Repository.
	Url pulumi.StringInput `pulumi:"url"`
	// Url of the Git repository web view
	WebUrl pulumi.StringInput `pulumi:"webUrl"`
}

func (GetRepositoriesRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRepositoriesRepository)(nil)).Elem()
}

func (i GetRepositoriesRepositoryArgs) ToGetRepositoriesRepositoryOutput() GetRepositoriesRepositoryOutput {
	return i.ToGetRepositoriesRepositoryOutputWithContext(context.Background())
}

func (i GetRepositoriesRepositoryArgs) ToGetRepositoriesRepositoryOutputWithContext(ctx context.Context) GetRepositoriesRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetRepositoriesRepositoryOutput)
}

// GetRepositoriesRepositoryArrayInput is an input type that accepts GetRepositoriesRepositoryArray and GetRepositoriesRepositoryArrayOutput values.
// You can construct a concrete instance of `GetRepositoriesRepositoryArrayInput` via:
//
//	GetRepositoriesRepositoryArray{ GetRepositoriesRepositoryArgs{...} }
type GetRepositoriesRepositoryArrayInput interface {
	pulumi.Input

	ToGetRepositoriesRepositoryArrayOutput() GetRepositoriesRepositoryArrayOutput
	ToGetRepositoriesRepositoryArrayOutputWithContext(context.Context) GetRepositoriesRepositoryArrayOutput
}

type GetRepositoriesRepositoryArray []GetRepositoriesRepositoryInput

func (GetRepositoriesRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetRepositoriesRepository)(nil)).Elem()
}

func (i GetRepositoriesRepositoryArray) ToGetRepositoriesRepositoryArrayOutput() GetRepositoriesRepositoryArrayOutput {
	return i.ToGetRepositoriesRepositoryArrayOutputWithContext(context.Background())
}

func (i GetRepositoriesRepositoryArray) ToGetRepositoriesRepositoryArrayOutputWithContext(ctx context.Context) GetRepositoriesRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetRepositoriesRepositoryArrayOutput)
}

type GetRepositoriesRepositoryOutput struct{ *pulumi.OutputState }

func (GetRepositoriesRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRepositoriesRepository)(nil)).Elem()
}

func (o GetRepositoriesRepositoryOutput) ToGetRepositoriesRepositoryOutput() GetRepositoriesRepositoryOutput {
	return o
}

func (o GetRepositoriesRepositoryOutput) ToGetRepositoriesRepositoryOutputWithContext(ctx context.Context) GetRepositoriesRepositoryOutput {
	return o
}

// The ref of the default branch.
func (o GetRepositoriesRepositoryOutput) DefaultBranch() pulumi.StringOutput {
	return o.ApplyT(func(v GetRepositoriesRepository) string { return v.DefaultBranch }).(pulumi.StringOutput)
}

// Git repository identifier.
func (o GetRepositoriesRepositoryOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetRepositoriesRepository) string { return v.Id }).(pulumi.StringOutput)
}

// Name of the Git repository to retrieve; requires `projectId` to be specified as well
func (o GetRepositoriesRepositoryOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetRepositoriesRepository) string { return v.Name }).(pulumi.StringOutput)
}

// ID of project to list Git repositories
func (o GetRepositoriesRepositoryOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v GetRepositoriesRepository) string { return v.ProjectId }).(pulumi.StringOutput)
}

// HTTPS Url to clone the Git repository
func (o GetRepositoriesRepositoryOutput) RemoteUrl() pulumi.StringOutput {
	return o.ApplyT(func(v GetRepositoriesRepository) string { return v.RemoteUrl }).(pulumi.StringOutput)
}

// Compressed size (bytes) of the repository.
func (o GetRepositoriesRepositoryOutput) Size() pulumi.IntOutput {
	return o.ApplyT(func(v GetRepositoriesRepository) int { return v.Size }).(pulumi.IntOutput)
}

// SSH Url to clone the Git repository
func (o GetRepositoriesRepositoryOutput) SshUrl() pulumi.StringOutput {
	return o.ApplyT(func(v GetRepositoriesRepository) string { return v.SshUrl }).(pulumi.StringOutput)
}

// Details REST API endpoint for the Git Repository.
func (o GetRepositoriesRepositoryOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v GetRepositoriesRepository) string { return v.Url }).(pulumi.StringOutput)
}

// Url of the Git repository web view
func (o GetRepositoriesRepositoryOutput) WebUrl() pulumi.StringOutput {
	return o.ApplyT(func(v GetRepositoriesRepository) string { return v.WebUrl }).(pulumi.StringOutput)
}

type GetRepositoriesRepositoryArrayOutput struct{ *pulumi.OutputState }

func (GetRepositoriesRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetRepositoriesRepository)(nil)).Elem()
}

func (o GetRepositoriesRepositoryArrayOutput) ToGetRepositoriesRepositoryArrayOutput() GetRepositoriesRepositoryArrayOutput {
	return o
}

func (o GetRepositoriesRepositoryArrayOutput) ToGetRepositoriesRepositoryArrayOutputWithContext(ctx context.Context) GetRepositoriesRepositoryArrayOutput {
	return o
}

func (o GetRepositoriesRepositoryArrayOutput) Index(i pulumi.IntInput) GetRepositoriesRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetRepositoriesRepository {
		return vs[0].([]GetRepositoriesRepository)[vs[1].(int)]
	}).(GetRepositoriesRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GitInitializationInput)(nil)).Elem(), GitInitializationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GitInitializationPtrInput)(nil)).Elem(), GitInitializationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetRepositoriesRepositoryInput)(nil)).Elem(), GetRepositoriesRepositoryArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetRepositoriesRepositoryArrayInput)(nil)).Elem(), GetRepositoriesRepositoryArray{})
	pulumi.RegisterOutputType(GitInitializationOutput{})
	pulumi.RegisterOutputType(GitInitializationPtrOutput{})
	pulumi.RegisterOutputType(GetRepositoriesRepositoryOutput{})
	pulumi.RegisterOutputType(GetRepositoriesRepositoryArrayOutput{})
}
