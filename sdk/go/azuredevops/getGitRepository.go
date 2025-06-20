// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azuredevops/sdk/v3/go/azuredevops/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about a **single** (existing) Git Repository within Azure DevOps.
// To read information about **multiple** Git Repositories use the data source `getRepositories`
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
//			example, err := azuredevops.LookupProject(ctx, &azuredevops.LookupProjectArgs{
//				Name: pulumi.StringRef("Example Project"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			// Load a specific Git repository by name
//			_, err = azuredevops.GetGitRepository(ctx, &azuredevops.GetGitRepositoryArgs{
//				ProjectId: example.Id,
//				Name:      "Example Repository",
//			}, nil)
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
// - [Azure DevOps Service REST API 7.0 - Git API](https://docs.microsoft.com/en-us/rest/api/azure/devops/git/?view=azure-devops-rest-7.0)
func GetGitRepository(ctx *pulumi.Context, args *GetGitRepositoryArgs, opts ...pulumi.InvokeOption) (*GetGitRepositoryResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetGitRepositoryResult
	err := ctx.Invoke("azuredevops:index/getGitRepository:getGitRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGitRepository.
type GetGitRepositoryArgs struct {
	// The Name of the Git repository to retrieve
	Name string `pulumi:"name"`
	// The ID of project to list Git repositories
	ProjectId string `pulumi:"projectId"`
}

// A collection of values returned by getGitRepository.
type GetGitRepositoryResult struct {
	// The ref of the default branch.
	DefaultBranch string `pulumi:"defaultBranch"`
	// Indicates whether the repository is disabled.
	Disabled bool `pulumi:"disabled"`
	// The provider-assigned unique ID for this managed resource.
	Id     string `pulumi:"id"`
	IsFork bool   `pulumi:"isFork"`
	// The name of the Git repository.
	Name string `pulumi:"name"`
	// Project identifier to which the Git repository belongs.
	ProjectId string `pulumi:"projectId"`
	// HTTPS Url to clone the Git repository
	RemoteUrl string `pulumi:"remoteUrl"`
	// Compressed size (bytes) of the repository.
	Size int `pulumi:"size"`
	// SSH Url to clone the Git repository
	SshUrl string `pulumi:"sshUrl"`
	// The details REST API endpoint for the Git Repository.
	Url string `pulumi:"url"`
	// The Url of the Git repository web view
	WebUrl string `pulumi:"webUrl"`
}

func GetGitRepositoryOutput(ctx *pulumi.Context, args GetGitRepositoryOutputArgs, opts ...pulumi.InvokeOption) GetGitRepositoryResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetGitRepositoryResultOutput, error) {
			args := v.(GetGitRepositoryArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azuredevops:index/getGitRepository:getGitRepository", args, GetGitRepositoryResultOutput{}, options).(GetGitRepositoryResultOutput), nil
		}).(GetGitRepositoryResultOutput)
}

// A collection of arguments for invoking getGitRepository.
type GetGitRepositoryOutputArgs struct {
	// The Name of the Git repository to retrieve
	Name pulumi.StringInput `pulumi:"name"`
	// The ID of project to list Git repositories
	ProjectId pulumi.StringInput `pulumi:"projectId"`
}

func (GetGitRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGitRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getGitRepository.
type GetGitRepositoryResultOutput struct{ *pulumi.OutputState }

func (GetGitRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGitRepositoryResult)(nil)).Elem()
}

func (o GetGitRepositoryResultOutput) ToGetGitRepositoryResultOutput() GetGitRepositoryResultOutput {
	return o
}

func (o GetGitRepositoryResultOutput) ToGetGitRepositoryResultOutputWithContext(ctx context.Context) GetGitRepositoryResultOutput {
	return o
}

// The ref of the default branch.
func (o GetGitRepositoryResultOutput) DefaultBranch() pulumi.StringOutput {
	return o.ApplyT(func(v GetGitRepositoryResult) string { return v.DefaultBranch }).(pulumi.StringOutput)
}

// Indicates whether the repository is disabled.
func (o GetGitRepositoryResultOutput) Disabled() pulumi.BoolOutput {
	return o.ApplyT(func(v GetGitRepositoryResult) bool { return v.Disabled }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetGitRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetGitRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetGitRepositoryResultOutput) IsFork() pulumi.BoolOutput {
	return o.ApplyT(func(v GetGitRepositoryResult) bool { return v.IsFork }).(pulumi.BoolOutput)
}

// The name of the Git repository.
func (o GetGitRepositoryResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetGitRepositoryResult) string { return v.Name }).(pulumi.StringOutput)
}

// Project identifier to which the Git repository belongs.
func (o GetGitRepositoryResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v GetGitRepositoryResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

// HTTPS Url to clone the Git repository
func (o GetGitRepositoryResultOutput) RemoteUrl() pulumi.StringOutput {
	return o.ApplyT(func(v GetGitRepositoryResult) string { return v.RemoteUrl }).(pulumi.StringOutput)
}

// Compressed size (bytes) of the repository.
func (o GetGitRepositoryResultOutput) Size() pulumi.IntOutput {
	return o.ApplyT(func(v GetGitRepositoryResult) int { return v.Size }).(pulumi.IntOutput)
}

// SSH Url to clone the Git repository
func (o GetGitRepositoryResultOutput) SshUrl() pulumi.StringOutput {
	return o.ApplyT(func(v GetGitRepositoryResult) string { return v.SshUrl }).(pulumi.StringOutput)
}

// The details REST API endpoint for the Git Repository.
func (o GetGitRepositoryResultOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v GetGitRepositoryResult) string { return v.Url }).(pulumi.StringOutput)
}

// The Url of the Git repository web view
func (o GetGitRepositoryResultOutput) WebUrl() pulumi.StringOutput {
	return o.ApplyT(func(v GetGitRepositoryResult) string { return v.WebUrl }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetGitRepositoryResultOutput{})
}
