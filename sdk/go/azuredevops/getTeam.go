// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing Team in a Project within Azure DevOps.
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
// 		exampleProject, err := azuredevops.NewProject(ctx, "exampleProject", &azuredevops.ProjectArgs{
// 			WorkItemTemplate: pulumi.String("Agile"),
// 			VersionControl:   pulumi.String("Git"),
// 			Visibility:       pulumi.String("private"),
// 			Description:      pulumi.String("Managed by Terraform"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_ = azuredevops.LookupTeamOutput(ctx, GetTeamOutputArgs{
// 			ProjectId: exampleProject.ID(),
// 			Name:      pulumi.String("Example Project Team"),
// 		}, nil)
// 		return nil
// 	})
// }
// ```
// ## Relevant Links
//
// - [Azure DevOps Service REST API 6.0 - Teams - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/core/teams/get?view=azure-devops-rest-6.0)
//
// ## PAT Permissions Required
//
// - **vso.project**:	Grants the ability to read projects and teams.
func LookupTeam(ctx *pulumi.Context, args *LookupTeamArgs, opts ...pulumi.InvokeOption) (*LookupTeamResult, error) {
	var rv LookupTeamResult
	err := ctx.Invoke("azuredevops:index/getTeam:getTeam", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTeam.
type LookupTeamArgs struct {
	// The name of the Team.
	Name string `pulumi:"name"`
	// The Project ID.
	ProjectId string `pulumi:"projectId"`
}

// A collection of values returned by getTeam.
type LookupTeamResult struct {
	// List of subject descriptors for `administrators` of the team.
	Administrators []string `pulumi:"administrators"`
	// Team description.
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// List of subject descriptors for `members` of the team.
	Members   []string `pulumi:"members"`
	Name      string   `pulumi:"name"`
	ProjectId string   `pulumi:"projectId"`
}

func LookupTeamOutput(ctx *pulumi.Context, args LookupTeamOutputArgs, opts ...pulumi.InvokeOption) LookupTeamResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTeamResult, error) {
			args := v.(LookupTeamArgs)
			r, err := LookupTeam(ctx, &args, opts...)
			var s LookupTeamResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTeamResultOutput)
}

// A collection of arguments for invoking getTeam.
type LookupTeamOutputArgs struct {
	// The name of the Team.
	Name pulumi.StringInput `pulumi:"name"`
	// The Project ID.
	ProjectId pulumi.StringInput `pulumi:"projectId"`
}

func (LookupTeamOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTeamArgs)(nil)).Elem()
}

// A collection of values returned by getTeam.
type LookupTeamResultOutput struct{ *pulumi.OutputState }

func (LookupTeamResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTeamResult)(nil)).Elem()
}

func (o LookupTeamResultOutput) ToLookupTeamResultOutput() LookupTeamResultOutput {
	return o
}

func (o LookupTeamResultOutput) ToLookupTeamResultOutputWithContext(ctx context.Context) LookupTeamResultOutput {
	return o
}

// List of subject descriptors for `administrators` of the team.
func (o LookupTeamResultOutput) Administrators() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupTeamResult) []string { return v.Administrators }).(pulumi.StringArrayOutput)
}

// Team description.
func (o LookupTeamResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTeamResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupTeamResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTeamResult) string { return v.Id }).(pulumi.StringOutput)
}

// List of subject descriptors for `members` of the team.
func (o LookupTeamResultOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupTeamResult) []string { return v.Members }).(pulumi.StringArrayOutput)
}

func (o LookupTeamResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTeamResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupTeamResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTeamResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTeamResultOutput{})
}
