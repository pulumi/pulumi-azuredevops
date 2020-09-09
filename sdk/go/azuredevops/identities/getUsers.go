// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package identities

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to access information about an existing users within Azure DevOps.
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
// 		opt0 := "contoso-user@contoso.onmicrosoft.com"
// 		_, err := azuredevops.GetUsers(ctx, &azuredevops.GetUsersArgs{
// 			PrincipalName: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = azuredevops.GetUsers(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		opt1 := "aad"
// 		_, err = azuredevops.GetUsers(ctx, &azuredevops.GetUsersArgs{
// 			Origin: &opt1,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = azuredevops.GetUsers(ctx, &azuredevops.GetUsersArgs{
// 			SubjectTypes: []string{
// 				"aad",
// 				"msa",
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		opt2 := "aad"
// 		opt3 := "a7ead982-8438-4cd2-b9e3-c3aa51a7b675"
// 		_, err = azuredevops.GetUsers(ctx, &azuredevops.GetUsersArgs{
// 			Origin:   &opt2,
// 			OriginId: &opt3,
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
// - [Azure DevOps Service REST API 5.1 - Graph Users API](https://docs.microsoft.com/en-us/rest/api/azure/devops/graph/users?view=azure-devops-rest-5.1)
//
// Deprecated: azuredevops.identities.getUsers has been deprecated in favor of azuredevops.getUsers
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
	// The type of source provider for the `originId` parameter (ex:AD, AAD, MSA) The supported origins are listed below.
	Origin *string `pulumi:"origin"`
	// The unique identifier from the system of origin.
	OriginId *string `pulumi:"originId"`
	// The PrincipalName of this graph member from the source provider.
	PrincipalName *string `pulumi:"principalName"`
	// A list of user subject subtypes to reduce the retrieved results, e.g. `msa`, `aad`, `svc` (service identity), `imp` (imported identity), etc. The supported subject types are listed below.
	SubjectTypes []string `pulumi:"subjectTypes"`
}

// A collection of values returned by getUsers.
type GetUsersResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The type of source provider for the origin identifier (ex:AD, AAD, MSA)
	Origin *string `pulumi:"origin"`
	// The unique identifier from the system of origin. Typically a sid, object id or Guid. Linking and unlinking operations can cause this value to change for a user because the user is not backed by a different provider and has a different unique id in the new provider.
	OriginId *string `pulumi:"originId"`
	// This is the PrincipalName of this graph member from the source provider. The source provider may change this field over time and it is not guaranteed to be immutable for the life of the graph member by VSTS.
	PrincipalName *string  `pulumi:"principalName"`
	SubjectTypes  []string `pulumi:"subjectTypes"`
	// A list of existing users in your Azure DevOps Organization with details about every single user which includes:
	Users []GetUsersUser `pulumi:"users"`
}
