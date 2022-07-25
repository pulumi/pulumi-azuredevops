// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about existing Teams in a Project or globally within an Azure DevOps organization
//
// ## Relevant Links
//
// - [Azure DevOps Service REST API 6.0 - Teams - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/core/teams/get?view=azure-devops-rest-6.0)
//
// ## PAT Permissions Required
//
// - **vso.project**:	Grants the ability to read projects and teams.
func GetTeams(ctx *pulumi.Context, args *GetTeamsArgs, opts ...pulumi.InvokeOption) (*GetTeamsResult, error) {
	var rv GetTeamsResult
	err := ctx.Invoke("azuredevops:index/getTeams:getTeams", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTeams.
type GetTeamsArgs struct {
	// The Project ID. If no project ID all teams of the organization will be returned.
	ProjectId *string `pulumi:"projectId"`
}

// A collection of values returned by getTeams.
type GetTeamsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Project identifier.
	// - `id - Team identifier
	ProjectId *string `pulumi:"projectId"`
	// A list of existing projects in your Azure DevOps Organization with details about every project which includes:
	Teams []GetTeamsTeam `pulumi:"teams"`
}

func GetTeamsOutput(ctx *pulumi.Context, args GetTeamsOutputArgs, opts ...pulumi.InvokeOption) GetTeamsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetTeamsResult, error) {
			args := v.(GetTeamsArgs)
			r, err := GetTeams(ctx, &args, opts...)
			var s GetTeamsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetTeamsResultOutput)
}

// A collection of arguments for invoking getTeams.
type GetTeamsOutputArgs struct {
	// The Project ID. If no project ID all teams of the organization will be returned.
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
}

func (GetTeamsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTeamsArgs)(nil)).Elem()
}

// A collection of values returned by getTeams.
type GetTeamsResultOutput struct{ *pulumi.OutputState }

func (GetTeamsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTeamsResult)(nil)).Elem()
}

func (o GetTeamsResultOutput) ToGetTeamsResultOutput() GetTeamsResultOutput {
	return o
}

func (o GetTeamsResultOutput) ToGetTeamsResultOutputWithContext(ctx context.Context) GetTeamsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetTeamsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetTeamsResult) string { return v.Id }).(pulumi.StringOutput)
}

// Project identifier.
// - `id - Team identifier
func (o GetTeamsResultOutput) ProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTeamsResult) *string { return v.ProjectId }).(pulumi.StringPtrOutput)
}

// A list of existing projects in your Azure DevOps Organization with details about every project which includes:
func (o GetTeamsResultOutput) Teams() GetTeamsTeamArrayOutput {
	return o.ApplyT(func(v GetTeamsResult) []GetTeamsTeam { return v.Teams }).(GetTeamsTeamArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetTeamsResultOutput{})
}
