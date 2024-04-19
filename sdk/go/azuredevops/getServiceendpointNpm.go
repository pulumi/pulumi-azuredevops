// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azuredevops/sdk/v3/go/azuredevops/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing NPM Service Endpoint.
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
//			example, err := azuredevops.GetServiceendpointNpm(ctx, &azuredevops.GetServiceendpointNpmArgs{
//				ProjectId:           exampleAzuredevopsProject.Id,
//				ServiceEndpointName: pulumi.StringRef("Example npm"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("serviceEndpointId", example.Id)
//			return nil
//		})
//	}
//
// ```
func GetServiceendpointNpm(ctx *pulumi.Context, args *GetServiceendpointNpmArgs, opts ...pulumi.InvokeOption) (*GetServiceendpointNpmResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetServiceendpointNpmResult
	err := ctx.Invoke("azuredevops:index/getServiceendpointNpm:getServiceendpointNpm", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getServiceendpointNpm.
type GetServiceendpointNpmArgs struct {
	// The ID of the project.
	ProjectId string `pulumi:"projectId"`
	// the ID of the Service Endpoint.
	ServiceEndpointId *string `pulumi:"serviceEndpointId"`
	// the Name of the Service Endpoint.
	//
	// > **NOTE:** One of either `serviceEndpointId` or `serviceEndpointName` must be specified.
	ServiceEndpointName *string `pulumi:"serviceEndpointName"`
}

// A collection of values returned by getServiceendpointNpm.
type GetServiceendpointNpmResult struct {
	// Specifies the Authorization Scheme Map.
	Authorization map[string]string `pulumi:"authorization"`
	// Specifies the description of the Service Endpoint.
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id                  string `pulumi:"id"`
	ProjectId           string `pulumi:"projectId"`
	ServiceEndpointId   string `pulumi:"serviceEndpointId"`
	ServiceEndpointName string `pulumi:"serviceEndpointName"`
	// Specifies the URL of the npm registry to connect with.
	Url string `pulumi:"url"`
}

func GetServiceendpointNpmOutput(ctx *pulumi.Context, args GetServiceendpointNpmOutputArgs, opts ...pulumi.InvokeOption) GetServiceendpointNpmResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetServiceendpointNpmResult, error) {
			args := v.(GetServiceendpointNpmArgs)
			r, err := GetServiceendpointNpm(ctx, &args, opts...)
			var s GetServiceendpointNpmResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetServiceendpointNpmResultOutput)
}

// A collection of arguments for invoking getServiceendpointNpm.
type GetServiceendpointNpmOutputArgs struct {
	// The ID of the project.
	ProjectId pulumi.StringInput `pulumi:"projectId"`
	// the ID of the Service Endpoint.
	ServiceEndpointId pulumi.StringPtrInput `pulumi:"serviceEndpointId"`
	// the Name of the Service Endpoint.
	//
	// > **NOTE:** One of either `serviceEndpointId` or `serviceEndpointName` must be specified.
	ServiceEndpointName pulumi.StringPtrInput `pulumi:"serviceEndpointName"`
}

func (GetServiceendpointNpmOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServiceendpointNpmArgs)(nil)).Elem()
}

// A collection of values returned by getServiceendpointNpm.
type GetServiceendpointNpmResultOutput struct{ *pulumi.OutputState }

func (GetServiceendpointNpmResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServiceendpointNpmResult)(nil)).Elem()
}

func (o GetServiceendpointNpmResultOutput) ToGetServiceendpointNpmResultOutput() GetServiceendpointNpmResultOutput {
	return o
}

func (o GetServiceendpointNpmResultOutput) ToGetServiceendpointNpmResultOutputWithContext(ctx context.Context) GetServiceendpointNpmResultOutput {
	return o
}

// Specifies the Authorization Scheme Map.
func (o GetServiceendpointNpmResultOutput) Authorization() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetServiceendpointNpmResult) map[string]string { return v.Authorization }).(pulumi.StringMapOutput)
}

// Specifies the description of the Service Endpoint.
func (o GetServiceendpointNpmResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceendpointNpmResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetServiceendpointNpmResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceendpointNpmResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetServiceendpointNpmResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceendpointNpmResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

func (o GetServiceendpointNpmResultOutput) ServiceEndpointId() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceendpointNpmResult) string { return v.ServiceEndpointId }).(pulumi.StringOutput)
}

func (o GetServiceendpointNpmResultOutput) ServiceEndpointName() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceendpointNpmResult) string { return v.ServiceEndpointName }).(pulumi.StringOutput)
}

// Specifies the URL of the npm registry to connect with.
func (o GetServiceendpointNpmResultOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceendpointNpmResult) string { return v.Url }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetServiceendpointNpmResultOutput{})
}
