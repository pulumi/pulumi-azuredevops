// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azuredevops/sdk/v3/go/azuredevops/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing Group within Azure DevOps
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
//			exampleGetGroup, err := azuredevops.LookupGroup(ctx, &azuredevops.LookupGroupArgs{
//				ProjectId: pulumi.StringRef(example.Id),
//				Name:      "Example Group",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("groupId", exampleGetGroup.Id)
//			ctx.Export("groupDescriptor", exampleGetGroup.Descriptor)
//			_, err = azuredevops.LookupGroup(ctx, &azuredevops.LookupGroupArgs{
//				Name: "Project Collection Administrators",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("collectionGroupId", exampleGetGroup.Id)
//			ctx.Export("collectionGroupDescriptor", exampleGetGroup.Descriptor)
//			return nil
//		})
//	}
//
// ```
//
// ## Relevant Links
//
// - [Azure DevOps Service REST API 7.0 - Groups - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/graph/groups/get?view=azure-devops-rest-7.0)
//
// ## PAT Permissions Required
//
// - **Graph**: Read
// - **Work Items**: Read
func LookupGroup(ctx *pulumi.Context, args *LookupGroupArgs, opts ...pulumi.InvokeOption) (*LookupGroupResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupGroupResult
	err := ctx.Invoke("azuredevops:index/getGroup:getGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGroup.
type LookupGroupArgs struct {
	// The Name of the Group.
	Name string `pulumi:"name"`
	// The ID of the Project. If `projectId` is not specified the project collection groups will be searched.
	ProjectId *string `pulumi:"projectId"`
}

// A collection of values returned by getGroup.
type LookupGroupResult struct {
	// The Descriptor is the primary way to reference the graph subject. This field will uniquely identify the same graph subject across both Accounts and Organizations. In format of `vssgp.xxxxxxxxxxxxxxxxxxx`
	Descriptor string `pulumi:"descriptor"`
	// The ID of the group.
	GroupId string `pulumi:"groupId"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// The type of source provider for the origin identifier (ex:AD, AAD, MSA)
	Origin string `pulumi:"origin"`
	// The unique identifier from the system of origin. Typically a sid, object id or Guid. Linking and unlinking operations can cause this value to change for a user because the user is not backed by a different provider and has a different unique id in the new provider.
	OriginId  string  `pulumi:"originId"`
	ProjectId *string `pulumi:"projectId"`
}

func LookupGroupOutput(ctx *pulumi.Context, args LookupGroupOutputArgs, opts ...pulumi.InvokeOption) LookupGroupResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupGroupResultOutput, error) {
			args := v.(LookupGroupArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azuredevops:index/getGroup:getGroup", args, LookupGroupResultOutput{}, options).(LookupGroupResultOutput), nil
		}).(LookupGroupResultOutput)
}

// A collection of arguments for invoking getGroup.
type LookupGroupOutputArgs struct {
	// The Name of the Group.
	Name pulumi.StringInput `pulumi:"name"`
	// The ID of the Project. If `projectId` is not specified the project collection groups will be searched.
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
}

func (LookupGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGroupArgs)(nil)).Elem()
}

// A collection of values returned by getGroup.
type LookupGroupResultOutput struct{ *pulumi.OutputState }

func (LookupGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGroupResult)(nil)).Elem()
}

func (o LookupGroupResultOutput) ToLookupGroupResultOutput() LookupGroupResultOutput {
	return o
}

func (o LookupGroupResultOutput) ToLookupGroupResultOutputWithContext(ctx context.Context) LookupGroupResultOutput {
	return o
}

// The Descriptor is the primary way to reference the graph subject. This field will uniquely identify the same graph subject across both Accounts and Organizations. In format of `vssgp.xxxxxxxxxxxxxxxxxxx`
func (o LookupGroupResultOutput) Descriptor() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.Descriptor }).(pulumi.StringOutput)
}

// The ID of the group.
func (o LookupGroupResultOutput) GroupId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.GroupId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

// The type of source provider for the origin identifier (ex:AD, AAD, MSA)
func (o LookupGroupResultOutput) Origin() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.Origin }).(pulumi.StringOutput)
}

// The unique identifier from the system of origin. Typically a sid, object id or Guid. Linking and unlinking operations can cause this value to change for a user because the user is not backed by a different provider and has a different unique id in the new provider.
func (o LookupGroupResultOutput) OriginId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.OriginId }).(pulumi.StringOutput)
}

func (o LookupGroupResultOutput) ProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGroupResult) *string { return v.ProjectId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGroupResultOutput{})
}
