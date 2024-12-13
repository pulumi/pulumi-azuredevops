// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azuredevops/sdk/v3/go/azuredevops/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing Bitbucket service Endpoint.
//
// ## Example Usage
//
// ### By Service Endpoint ID
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
//			exampleGetServiceendpointBitbucket, err := azuredevops.GetServiceendpointBitbucket(ctx, &azuredevops.GetServiceendpointBitbucketArgs{
//				ProjectId:         example.Id,
//				ServiceEndpointId: pulumi.StringRef("00000000-0000-0000-0000-000000000000"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("serviceEndpointName", exampleGetServiceendpointBitbucket.ServiceEndpointName)
//			return nil
//		})
//	}
//
// ```
//
// ### By Service Endpoint Name
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
//			exampleGetServiceendpointBitbucket, err := azuredevops.GetServiceendpointBitbucket(ctx, &azuredevops.GetServiceendpointBitbucketArgs{
//				ProjectId:           example.Id,
//				ServiceEndpointName: pulumi.StringRef("Example"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("serviceEndpointId", exampleGetServiceendpointBitbucket.Id)
//			return nil
//		})
//	}
//
// ```
//
// ## PAT Permissions Required
//
// - **vso.serviceendpoint**: Grants the ability to read service endpoints.
func GetServiceendpointBitbucket(ctx *pulumi.Context, args *GetServiceendpointBitbucketArgs, opts ...pulumi.InvokeOption) (*GetServiceendpointBitbucketResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetServiceendpointBitbucketResult
	err := ctx.Invoke("azuredevops:index/getServiceendpointBitbucket:getServiceendpointBitbucket", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getServiceendpointBitbucket.
type GetServiceendpointBitbucketArgs struct {
	// The ID of the project.
	ProjectId string `pulumi:"projectId"`
	// the ID of the Service Endpoint.
	ServiceEndpointId *string `pulumi:"serviceEndpointId"`
	// the Name of the Service Endpoint.
	//
	// > **NOTE:** One of either `serviceEndpointId` or `serviceEndpointName` must be specified.
	// **NOTE:** When supplying `serviceEndpointName`, take care to ensure that this is a unique name.
	ServiceEndpointName *string `pulumi:"serviceEndpointName"`
}

// A collection of values returned by getServiceendpointBitbucket.
type GetServiceendpointBitbucketResult struct {
	// Specifies the Authorization Scheme Map.
	Authorization map[string]string `pulumi:"authorization"`
	// Specifies the description of the Service Endpoint.
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id                  string `pulumi:"id"`
	ProjectId           string `pulumi:"projectId"`
	ServiceEndpointId   string `pulumi:"serviceEndpointId"`
	ServiceEndpointName string `pulumi:"serviceEndpointName"`
}

func GetServiceendpointBitbucketOutput(ctx *pulumi.Context, args GetServiceendpointBitbucketOutputArgs, opts ...pulumi.InvokeOption) GetServiceendpointBitbucketResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetServiceendpointBitbucketResultOutput, error) {
			args := v.(GetServiceendpointBitbucketArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azuredevops:index/getServiceendpointBitbucket:getServiceendpointBitbucket", args, GetServiceendpointBitbucketResultOutput{}, options).(GetServiceendpointBitbucketResultOutput), nil
		}).(GetServiceendpointBitbucketResultOutput)
}

// A collection of arguments for invoking getServiceendpointBitbucket.
type GetServiceendpointBitbucketOutputArgs struct {
	// The ID of the project.
	ProjectId pulumi.StringInput `pulumi:"projectId"`
	// the ID of the Service Endpoint.
	ServiceEndpointId pulumi.StringPtrInput `pulumi:"serviceEndpointId"`
	// the Name of the Service Endpoint.
	//
	// > **NOTE:** One of either `serviceEndpointId` or `serviceEndpointName` must be specified.
	// **NOTE:** When supplying `serviceEndpointName`, take care to ensure that this is a unique name.
	ServiceEndpointName pulumi.StringPtrInput `pulumi:"serviceEndpointName"`
}

func (GetServiceendpointBitbucketOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServiceendpointBitbucketArgs)(nil)).Elem()
}

// A collection of values returned by getServiceendpointBitbucket.
type GetServiceendpointBitbucketResultOutput struct{ *pulumi.OutputState }

func (GetServiceendpointBitbucketResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServiceendpointBitbucketResult)(nil)).Elem()
}

func (o GetServiceendpointBitbucketResultOutput) ToGetServiceendpointBitbucketResultOutput() GetServiceendpointBitbucketResultOutput {
	return o
}

func (o GetServiceendpointBitbucketResultOutput) ToGetServiceendpointBitbucketResultOutputWithContext(ctx context.Context) GetServiceendpointBitbucketResultOutput {
	return o
}

// Specifies the Authorization Scheme Map.
func (o GetServiceendpointBitbucketResultOutput) Authorization() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetServiceendpointBitbucketResult) map[string]string { return v.Authorization }).(pulumi.StringMapOutput)
}

// Specifies the description of the Service Endpoint.
func (o GetServiceendpointBitbucketResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceendpointBitbucketResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetServiceendpointBitbucketResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceendpointBitbucketResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetServiceendpointBitbucketResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceendpointBitbucketResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

func (o GetServiceendpointBitbucketResultOutput) ServiceEndpointId() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceendpointBitbucketResult) string { return v.ServiceEndpointId }).(pulumi.StringOutput)
}

func (o GetServiceendpointBitbucketResultOutput) ServiceEndpointName() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceendpointBitbucketResult) string { return v.ServiceEndpointName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetServiceendpointBitbucketResultOutput{})
}
