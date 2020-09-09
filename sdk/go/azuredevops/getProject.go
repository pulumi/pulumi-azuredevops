// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to access information about an existing Project within Azure DevOps.
//
// ## Example Usage
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
// 		project, err := azuredevops.LookupProject(ctx, &azuredevops.LookupProjectArgs{
// 			ProjectName: "Sample Project",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("id", project.Id)
// 		ctx.Export("projectName", project.ProjectName)
// 		ctx.Export("visibility", project.Visibility)
// 		ctx.Export("versionControl", project.VersionControl)
// 		ctx.Export("workItemTemplate", project.WorkItemTemplate)
// 		ctx.Export("processTemplateId", project.ProcessTemplateId)
// 		return nil
// 	})
// }
// ```
// ## Relevant Links
//
// - [Azure DevOps Service REST API 5.1 - Projects - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/core/projects/get?view=azure-devops-rest-5.1)
func LookupProject(ctx *pulumi.Context, args *LookupProjectArgs, opts ...pulumi.InvokeOption) (*LookupProjectResult, error) {
	var rv LookupProjectResult
	err := ctx.Invoke("azuredevops:index/getProject:getProject", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProject.
type LookupProjectArgs struct {
	// Name of the Project.
	ProjectName string `pulumi:"projectName"`
}

// A collection of values returned by getProject.
type LookupProjectResult struct {
	Description string                 `pulumi:"description"`
	Features    map[string]interface{} `pulumi:"features"`
	// The provider-assigned unique ID for this managed resource.
	Id                string `pulumi:"id"`
	ProcessTemplateId string `pulumi:"processTemplateId"`
	ProjectName       string `pulumi:"projectName"`
	VersionControl    string `pulumi:"versionControl"`
	Visibility        string `pulumi:"visibility"`
	WorkItemTemplate  string `pulumi:"workItemTemplate"`
}
