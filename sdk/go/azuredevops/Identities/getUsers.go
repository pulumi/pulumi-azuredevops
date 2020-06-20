// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package Identities

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to access information about an existing users within Azure DevOps.
//
//
// ## Relevant Links
//
// - [Azure DevOps Service REST API 5.1 - Graph Users API](https://docs.microsoft.com/en-us/rest/api/azure/devops/graph/users?view=azure-devops-rest-5.1)
func GetUsers(ctx *pulumi.Context, args *GetUsersArgs, opts ...pulumi.InvokeOption) (*GetUsersResult, error) {
	var rv GetUsersResult
	err := ctx.Invoke("azuredevops:Identities/getUsers:getUsers", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getUsers.
type GetUsersArgs struct {
	Origin        *string  `pulumi:"origin"`
	OriginId      *string  `pulumi:"originId"`
	PrincipalName *string  `pulumi:"principalName"`
	SubjectTypes  []string `pulumi:"subjectTypes"`
}

// A collection of values returned by getUsers.
type GetUsersResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id            string         `pulumi:"id"`
	Origin        *string        `pulumi:"origin"`
	OriginId      *string        `pulumi:"originId"`
	PrincipalName *string        `pulumi:"principalName"`
	SubjectTypes  []string       `pulumi:"subjectTypes"`
	Users         []GetUsersUser `pulumi:"users"`
}
