// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azuredevops/sdk/v3/go/azuredevops/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing GitHub service Endpoint.
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
//			sample, err := azuredevops.LookupProject(ctx, &azuredevops.LookupProjectArgs{
//				Name: pulumi.StringRef("Sample Project"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			serviceendpoint, err := azuredevops.GetServiceEndpointGithub(ctx, &azuredevops.GetServiceEndpointGithubArgs{
//				ProjectId:         sample.Id,
//				ServiceEndpointId: pulumi.StringRef("00000000-0000-0000-0000-000000000000"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("serviceEndpointName", serviceendpoint.ServiceEndpointName)
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
//			sample, err := azuredevops.LookupProject(ctx, &azuredevops.LookupProjectArgs{
//				Name: pulumi.StringRef("Sample Project"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			serviceendpoint, err := azuredevops.GetServiceEndpointGithub(ctx, &azuredevops.GetServiceEndpointGithubArgs{
//				ProjectId:           sample.Id,
//				ServiceEndpointName: pulumi.StringRef("Example-Service-Endpoint"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("serviceEndpointId", serviceendpoint.Id)
//			return nil
//		})
//	}
//
// ```
func GetServiceEndpointGithub(ctx *pulumi.Context, args *GetServiceEndpointGithubArgs, opts ...pulumi.InvokeOption) (*GetServiceEndpointGithubResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetServiceEndpointGithubResult
	err := ctx.Invoke("azuredevops:index/getServiceEndpointGithub:getServiceEndpointGithub", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getServiceEndpointGithub.
type GetServiceEndpointGithubArgs struct {
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

// A collection of values returned by getServiceEndpointGithub.
type GetServiceEndpointGithubResult struct {
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

func GetServiceEndpointGithubOutput(ctx *pulumi.Context, args GetServiceEndpointGithubOutputArgs, opts ...pulumi.InvokeOption) GetServiceEndpointGithubResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetServiceEndpointGithubResult, error) {
			args := v.(GetServiceEndpointGithubArgs)
			r, err := GetServiceEndpointGithub(ctx, &args, opts...)
			var s GetServiceEndpointGithubResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetServiceEndpointGithubResultOutput)
}

// A collection of arguments for invoking getServiceEndpointGithub.
type GetServiceEndpointGithubOutputArgs struct {
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

func (GetServiceEndpointGithubOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServiceEndpointGithubArgs)(nil)).Elem()
}

// A collection of values returned by getServiceEndpointGithub.
type GetServiceEndpointGithubResultOutput struct{ *pulumi.OutputState }

func (GetServiceEndpointGithubResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServiceEndpointGithubResult)(nil)).Elem()
}

func (o GetServiceEndpointGithubResultOutput) ToGetServiceEndpointGithubResultOutput() GetServiceEndpointGithubResultOutput {
	return o
}

func (o GetServiceEndpointGithubResultOutput) ToGetServiceEndpointGithubResultOutputWithContext(ctx context.Context) GetServiceEndpointGithubResultOutput {
	return o
}

// Specifies the Authorization Scheme Map.
func (o GetServiceEndpointGithubResultOutput) Authorization() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetServiceEndpointGithubResult) map[string]string { return v.Authorization }).(pulumi.StringMapOutput)
}

// Specifies the description of the Service Endpoint.
func (o GetServiceEndpointGithubResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceEndpointGithubResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetServiceEndpointGithubResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceEndpointGithubResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetServiceEndpointGithubResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceEndpointGithubResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

func (o GetServiceEndpointGithubResultOutput) ServiceEndpointId() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceEndpointGithubResult) string { return v.ServiceEndpointId }).(pulumi.StringOutput)
}

func (o GetServiceEndpointGithubResultOutput) ServiceEndpointName() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceEndpointGithubResult) string { return v.ServiceEndpointName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetServiceEndpointGithubResultOutput{})
}
