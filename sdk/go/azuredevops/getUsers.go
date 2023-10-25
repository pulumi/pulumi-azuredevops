// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azuredevops/sdk/v2/go/azuredevops/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Use this data source to access information about an existing users within Azure DevOps.
//
// ## Relevant Links
//
// - [Azure DevOps Service REST API 7.0 - Graph Users API](https://docs.microsoft.com/en-us/rest/api/azure/devops/graph/users?view=azure-devops-rest-7.0)
func GetUsers(ctx *pulumi.Context, args *GetUsersArgs, opts ...pulumi.InvokeOption) (*GetUsersResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetUsersResult
	err := ctx.Invoke("azuredevops:index/getUsers:getUsers", args, &rv, opts...)
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
	//
	// DataSource without specifying any arguments will return all users inside an organization.
	//
	// List of possible subject types
	//
	// List of possible origins
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
	// A set of existing users in your Azure DevOps Organization with details about every single user which includes:
	Users []GetUsersUser `pulumi:"users"`
}

func GetUsersOutput(ctx *pulumi.Context, args GetUsersOutputArgs, opts ...pulumi.InvokeOption) GetUsersResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetUsersResult, error) {
			args := v.(GetUsersArgs)
			r, err := GetUsers(ctx, &args, opts...)
			var s GetUsersResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetUsersResultOutput)
}

// A collection of arguments for invoking getUsers.
type GetUsersOutputArgs struct {
	// The type of source provider for the `originId` parameter (ex:AD, AAD, MSA) The supported origins are listed below.
	Origin pulumi.StringPtrInput `pulumi:"origin"`
	// The unique identifier from the system of origin.
	//
	// DataSource without specifying any arguments will return all users inside an organization.
	//
	// List of possible subject types
	//
	// List of possible origins
	OriginId pulumi.StringPtrInput `pulumi:"originId"`
	// The PrincipalName of this graph member from the source provider.
	PrincipalName pulumi.StringPtrInput `pulumi:"principalName"`
	// A list of user subject subtypes to reduce the retrieved results, e.g. `msa`, `aad`, `svc` (service identity), `imp` (imported identity), etc. The supported subject types are listed below.
	SubjectTypes pulumi.StringArrayInput `pulumi:"subjectTypes"`
}

func (GetUsersOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetUsersArgs)(nil)).Elem()
}

// A collection of values returned by getUsers.
type GetUsersResultOutput struct{ *pulumi.OutputState }

func (GetUsersResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetUsersResult)(nil)).Elem()
}

func (o GetUsersResultOutput) ToGetUsersResultOutput() GetUsersResultOutput {
	return o
}

func (o GetUsersResultOutput) ToGetUsersResultOutputWithContext(ctx context.Context) GetUsersResultOutput {
	return o
}

func (o GetUsersResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetUsersResult] {
	return pulumix.Output[GetUsersResult]{
		OutputState: o.OutputState,
	}
}

// The provider-assigned unique ID for this managed resource.
func (o GetUsersResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetUsersResult) string { return v.Id }).(pulumi.StringOutput)
}

// The type of source provider for the origin identifier (ex:AD, AAD, MSA)
func (o GetUsersResultOutput) Origin() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetUsersResult) *string { return v.Origin }).(pulumi.StringPtrOutput)
}

// The unique identifier from the system of origin. Typically a sid, object id or Guid. Linking and unlinking operations can cause this value to change for a user because the user is not backed by a different provider and has a different unique id in the new provider.
func (o GetUsersResultOutput) OriginId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetUsersResult) *string { return v.OriginId }).(pulumi.StringPtrOutput)
}

// This is the PrincipalName of this graph member from the source provider. The source provider may change this field over time and it is not guaranteed to be immutable for the life of the graph member by VSTS.
func (o GetUsersResultOutput) PrincipalName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetUsersResult) *string { return v.PrincipalName }).(pulumi.StringPtrOutput)
}

func (o GetUsersResultOutput) SubjectTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetUsersResult) []string { return v.SubjectTypes }).(pulumi.StringArrayOutput)
}

// A set of existing users in your Azure DevOps Organization with details about every single user which includes:
func (o GetUsersResultOutput) Users() GetUsersUserArrayOutput {
	return o.ApplyT(func(v GetUsersResult) []GetUsersUser { return v.Users }).(GetUsersUserArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetUsersResultOutput{})
}
