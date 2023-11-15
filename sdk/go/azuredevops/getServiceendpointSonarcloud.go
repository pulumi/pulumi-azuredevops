// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azuredevops/sdk/v2/go/azuredevops/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing Sonar Cloud Service Endpoint.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-azuredevops/sdk/v2/go/azuredevops"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			example, err := azuredevops.GetServiceendpointSonarcloud(ctx, &azuredevops.GetServiceendpointSonarcloudArgs{
//				ProjectId:           azuredevops_project.Example.Id,
//				ServiceEndpointName: pulumi.StringRef("Example Sonar Cloud"),
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
func GetServiceendpointSonarcloud(ctx *pulumi.Context, args *GetServiceendpointSonarcloudArgs, opts ...pulumi.InvokeOption) (*GetServiceendpointSonarcloudResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetServiceendpointSonarcloudResult
	err := ctx.Invoke("azuredevops:index/getServiceendpointSonarcloud:getServiceendpointSonarcloud", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getServiceendpointSonarcloud.
type GetServiceendpointSonarcloudArgs struct {
	// The ID of the project.
	ProjectId string `pulumi:"projectId"`
	// the ID of the Service Endpoint.
	ServiceEndpointId *string `pulumi:"serviceEndpointId"`
	// the Name of the Service Endpoint.
	//
	// > **NOTE:** One of either `serviceEndpointId` or `serviceEndpointName` must be specified.
	ServiceEndpointName *string `pulumi:"serviceEndpointName"`
}

// A collection of values returned by getServiceendpointSonarcloud.
type GetServiceendpointSonarcloudResult struct {
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

func GetServiceendpointSonarcloudOutput(ctx *pulumi.Context, args GetServiceendpointSonarcloudOutputArgs, opts ...pulumi.InvokeOption) GetServiceendpointSonarcloudResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetServiceendpointSonarcloudResult, error) {
			args := v.(GetServiceendpointSonarcloudArgs)
			r, err := GetServiceendpointSonarcloud(ctx, &args, opts...)
			var s GetServiceendpointSonarcloudResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetServiceendpointSonarcloudResultOutput)
}

// A collection of arguments for invoking getServiceendpointSonarcloud.
type GetServiceendpointSonarcloudOutputArgs struct {
	// The ID of the project.
	ProjectId pulumi.StringInput `pulumi:"projectId"`
	// the ID of the Service Endpoint.
	ServiceEndpointId pulumi.StringPtrInput `pulumi:"serviceEndpointId"`
	// the Name of the Service Endpoint.
	//
	// > **NOTE:** One of either `serviceEndpointId` or `serviceEndpointName` must be specified.
	ServiceEndpointName pulumi.StringPtrInput `pulumi:"serviceEndpointName"`
}

func (GetServiceendpointSonarcloudOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServiceendpointSonarcloudArgs)(nil)).Elem()
}

// A collection of values returned by getServiceendpointSonarcloud.
type GetServiceendpointSonarcloudResultOutput struct{ *pulumi.OutputState }

func (GetServiceendpointSonarcloudResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServiceendpointSonarcloudResult)(nil)).Elem()
}

func (o GetServiceendpointSonarcloudResultOutput) ToGetServiceendpointSonarcloudResultOutput() GetServiceendpointSonarcloudResultOutput {
	return o
}

func (o GetServiceendpointSonarcloudResultOutput) ToGetServiceendpointSonarcloudResultOutputWithContext(ctx context.Context) GetServiceendpointSonarcloudResultOutput {
	return o
}

// Specifies the Authorization Scheme Map.
func (o GetServiceendpointSonarcloudResultOutput) Authorization() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetServiceendpointSonarcloudResult) map[string]string { return v.Authorization }).(pulumi.StringMapOutput)
}

// Specifies the description of the Service Endpoint.
func (o GetServiceendpointSonarcloudResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceendpointSonarcloudResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetServiceendpointSonarcloudResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceendpointSonarcloudResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetServiceendpointSonarcloudResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceendpointSonarcloudResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

func (o GetServiceendpointSonarcloudResultOutput) ServiceEndpointId() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceendpointSonarcloudResult) string { return v.ServiceEndpointId }).(pulumi.StringOutput)
}

func (o GetServiceendpointSonarcloudResultOutput) ServiceEndpointName() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceendpointSonarcloudResult) string { return v.ServiceEndpointName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetServiceendpointSonarcloudResultOutput{})
}
