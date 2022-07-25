// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package repository

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about **multiple** existing Git Repositories within Azure DevOps.
// To read informations about a **single** Git Repository use the data source `Git`
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
// 		example, err := azuredevops.LookupProject(ctx, &GetProjectArgs{
// 			Name: pulumi.StringRef("Example Project"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = azuredevops.GetRepositories(ctx, &GetRepositoriesArgs{
// 			ProjectId:     pulumi.StringRef(example.Id),
// 			IncludeHidden: pulumi.BoolRef(true),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = azuredevops.GetRepositories(ctx, &GetRepositoriesArgs{
// 			ProjectId: pulumi.StringRef(example.Id),
// 			Name:      pulumi.StringRef("Example Repository"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## Relevant Links
//
// - [Azure DevOps Service REST API 6.0 - Git API](https://docs.microsoft.com/en-us/rest/api/azure/devops/git/?view=azure-devops-rest-6.0)
//
// Deprecated: azuredevops.repository.getRepositories has been deprecated in favor of azuredevops.getRepositories
func GetRepositories(ctx *pulumi.Context, args *GetRepositoriesArgs, opts ...pulumi.InvokeOption) (*GetRepositoriesResult, error) {
	var rv GetRepositoriesResult
	err := ctx.Invoke("azuredevops:Repository/getRepositories:getRepositories", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRepositories.
type GetRepositoriesArgs struct {
	IncludeHidden *bool `pulumi:"includeHidden"`
	// Name of the Git repository to retrieve; requires `projectId` to be specified as well
	Name *string `pulumi:"name"`
	// ID of project to list Git repositories
	ProjectId *string `pulumi:"projectId"`
}

// A collection of values returned by getRepositories.
type GetRepositoriesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id            string `pulumi:"id"`
	IncludeHidden *bool  `pulumi:"includeHidden"`
	// Git repository name.
	Name *string `pulumi:"name"`
	// Project identifier to which the Git repository belongs.
	ProjectId *string `pulumi:"projectId"`
	// A list of existing projects in your Azure DevOps Organization with details about every project which includes:
	Repositories []GetRepositoriesRepository `pulumi:"repositories"`
}

func GetRepositoriesOutput(ctx *pulumi.Context, args GetRepositoriesOutputArgs, opts ...pulumi.InvokeOption) GetRepositoriesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetRepositoriesResult, error) {
			args := v.(GetRepositoriesArgs)
			r, err := GetRepositories(ctx, &args, opts...)
			var s GetRepositoriesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetRepositoriesResultOutput)
}

// A collection of arguments for invoking getRepositories.
type GetRepositoriesOutputArgs struct {
	IncludeHidden pulumi.BoolPtrInput `pulumi:"includeHidden"`
	// Name of the Git repository to retrieve; requires `projectId` to be specified as well
	Name pulumi.StringPtrInput `pulumi:"name"`
	// ID of project to list Git repositories
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
}

func (GetRepositoriesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRepositoriesArgs)(nil)).Elem()
}

// A collection of values returned by getRepositories.
type GetRepositoriesResultOutput struct{ *pulumi.OutputState }

func (GetRepositoriesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRepositoriesResult)(nil)).Elem()
}

func (o GetRepositoriesResultOutput) ToGetRepositoriesResultOutput() GetRepositoriesResultOutput {
	return o
}

func (o GetRepositoriesResultOutput) ToGetRepositoriesResultOutputWithContext(ctx context.Context) GetRepositoriesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetRepositoriesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetRepositoriesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetRepositoriesResultOutput) IncludeHidden() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetRepositoriesResult) *bool { return v.IncludeHidden }).(pulumi.BoolPtrOutput)
}

// Git repository name.
func (o GetRepositoriesResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRepositoriesResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// Project identifier to which the Git repository belongs.
func (o GetRepositoriesResultOutput) ProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRepositoriesResult) *string { return v.ProjectId }).(pulumi.StringPtrOutput)
}

// A list of existing projects in your Azure DevOps Organization with details about every project which includes:
func (o GetRepositoriesResultOutput) Repositories() GetRepositoriesRepositoryArrayOutput {
	return o.ApplyT(func(v GetRepositoriesResult) []GetRepositoriesRepository { return v.Repositories }).(GetRepositoriesRepositoryArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetRepositoriesResultOutput{})
}
