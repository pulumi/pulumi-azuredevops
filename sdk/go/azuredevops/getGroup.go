// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to access information about an existing Group within Azure DevOps
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
// 		opt0 := "contoso-project"
// 		project, err := azuredevops.LookupProject(ctx, &azuredevops.LookupProjectArgs{
// 			Name: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		opt1 := project.Id
// 		test, err := azuredevops.LookupGroup(ctx, &azuredevops.LookupGroupArgs{
// 			ProjectId: &opt1,
// 			Name:      "Test Group",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("groupId", test.Id)
// 		ctx.Export("groupDescriptor", test.Descriptor)
// 		test_collection_group, err := azuredevops.LookupGroup(ctx, &azuredevops.LookupGroupArgs{
// 			Name: "Project Collection Administrators",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("collectionGroupId", test_collection_group.Id)
// 		ctx.Export("collectionGroupDescriptor", test_collection_group.Descriptor)
// 		return nil
// 	})
// }
// ```
// ## Relevant Links
//
// - [Azure DevOps Service REST API 5.1 - Groups - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/graph/groups/get?view=azure-devops-rest-5.1)
func LookupGroup(ctx *pulumi.Context, args *LookupGroupArgs, opts ...pulumi.InvokeOption) (*LookupGroupResult, error) {
	var rv LookupGroupResult
	err := ctx.Invoke("azuredevops:index/getGroup:getGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGroup.
type LookupGroupArgs struct {
	// The Group Name.
	Name string `pulumi:"name"`
	// The Project ID. If no project ID is specified the project collection groups will be searched.
	ProjectId *string `pulumi:"projectId"`
}

// A collection of values returned by getGroup.
type LookupGroupResult struct {
	// The Descriptor is the primary way to reference the graph subject. This field will uniquely identify the same graph subject across both Accounts and Organizations.
	Descriptor string `pulumi:"descriptor"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// The type of source provider for the origin identifier (ex:AD, AAD, MSA)
	Origin string `pulumi:"origin"`
	// The unique identifier from the system of origin. Typically a sid, object id or Guid. Linking and unlinking operations can cause this value to change for a user because the user is not backed by a different provider and has a different unique id in the new provider.
	OriginId  string  `pulumi:"originId"`
	ProjectId *string `pulumi:"projectId"`
}
