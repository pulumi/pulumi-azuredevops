// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azuredevops/sdk/v3/go/azuredevops/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing Azure Container Registry Service Endpoint.
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
//			example, err := azuredevops.GetServiceendpointAzurecr(ctx, &azuredevops.GetServiceendpointAzurecrArgs{
//				ProjectId:           exampleAzuredevopsProject.Id,
//				ServiceEndpointName: pulumi.StringRef("Example Azure Container Registry"),
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
func GetServiceendpointAzurecr(ctx *pulumi.Context, args *GetServiceendpointAzurecrArgs, opts ...pulumi.InvokeOption) (*GetServiceendpointAzurecrResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetServiceendpointAzurecrResult
	err := ctx.Invoke("azuredevops:index/getServiceendpointAzurecr:getServiceendpointAzurecr", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getServiceendpointAzurecr.
type GetServiceendpointAzurecrArgs struct {
	// The ID of the project.
	ProjectId string `pulumi:"projectId"`
	// the ID of the Service Endpoint.
	ServiceEndpointId *string `pulumi:"serviceEndpointId"`
	// the Name of the Service Endpoint.
	//
	// > **NOTE:** One of either `serviceEndpointId` or `serviceEndpointName` must be specified.
	ServiceEndpointName *string `pulumi:"serviceEndpointName"`
}

// A collection of values returned by getServiceendpointAzurecr.
type GetServiceendpointAzurecrResult struct {
	// The Object ID of the Service Principal.
	AppObjectId string `pulumi:"appObjectId"`
	// Specifies the Authorization Scheme Map.
	Authorization map[string]string `pulumi:"authorization"`
	// The ID of Service Principal Role Assignment.
	AzSpnRoleAssignmentId string `pulumi:"azSpnRoleAssignmentId"`
	// The Service Principal Role Permissions.
	AzSpnRolePermissions string `pulumi:"azSpnRolePermissions"`
	// The Azure Container Registry name.
	AzurecrName string `pulumi:"azurecrName"`
	// The Tenant ID of the service principal.
	AzurecrSpnTenantid string `pulumi:"azurecrSpnTenantid"`
	// The Subscription ID of the Azure targets.
	AzurecrSubscriptionId string `pulumi:"azurecrSubscriptionId"`
	// The Subscription Name of the Azure targets.
	AzurecrSubscriptionName string `pulumi:"azurecrSubscriptionName"`
	// The Service Endpoint description.
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id        string `pulumi:"id"`
	ProjectId string `pulumi:"projectId"`
	// The Resource Group to which the Container Registry belongs.
	ResourceGroup       string `pulumi:"resourceGroup"`
	ServiceEndpointId   string `pulumi:"serviceEndpointId"`
	ServiceEndpointName string `pulumi:"serviceEndpointName"`
	// The Application(Client) ID of the Service Principal.
	ServicePrincipalId string `pulumi:"servicePrincipalId"`
	// The ID of the Service Principal.
	SpnObjectId string `pulumi:"spnObjectId"`
}

func GetServiceendpointAzurecrOutput(ctx *pulumi.Context, args GetServiceendpointAzurecrOutputArgs, opts ...pulumi.InvokeOption) GetServiceendpointAzurecrResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetServiceendpointAzurecrResultOutput, error) {
			args := v.(GetServiceendpointAzurecrArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azuredevops:index/getServiceendpointAzurecr:getServiceendpointAzurecr", args, GetServiceendpointAzurecrResultOutput{}, options).(GetServiceendpointAzurecrResultOutput), nil
		}).(GetServiceendpointAzurecrResultOutput)
}

// A collection of arguments for invoking getServiceendpointAzurecr.
type GetServiceendpointAzurecrOutputArgs struct {
	// The ID of the project.
	ProjectId pulumi.StringInput `pulumi:"projectId"`
	// the ID of the Service Endpoint.
	ServiceEndpointId pulumi.StringPtrInput `pulumi:"serviceEndpointId"`
	// the Name of the Service Endpoint.
	//
	// > **NOTE:** One of either `serviceEndpointId` or `serviceEndpointName` must be specified.
	ServiceEndpointName pulumi.StringPtrInput `pulumi:"serviceEndpointName"`
}

func (GetServiceendpointAzurecrOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServiceendpointAzurecrArgs)(nil)).Elem()
}

// A collection of values returned by getServiceendpointAzurecr.
type GetServiceendpointAzurecrResultOutput struct{ *pulumi.OutputState }

func (GetServiceendpointAzurecrResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServiceendpointAzurecrResult)(nil)).Elem()
}

func (o GetServiceendpointAzurecrResultOutput) ToGetServiceendpointAzurecrResultOutput() GetServiceendpointAzurecrResultOutput {
	return o
}

func (o GetServiceendpointAzurecrResultOutput) ToGetServiceendpointAzurecrResultOutputWithContext(ctx context.Context) GetServiceendpointAzurecrResultOutput {
	return o
}

// The Object ID of the Service Principal.
func (o GetServiceendpointAzurecrResultOutput) AppObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceendpointAzurecrResult) string { return v.AppObjectId }).(pulumi.StringOutput)
}

// Specifies the Authorization Scheme Map.
func (o GetServiceendpointAzurecrResultOutput) Authorization() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetServiceendpointAzurecrResult) map[string]string { return v.Authorization }).(pulumi.StringMapOutput)
}

// The ID of Service Principal Role Assignment.
func (o GetServiceendpointAzurecrResultOutput) AzSpnRoleAssignmentId() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceendpointAzurecrResult) string { return v.AzSpnRoleAssignmentId }).(pulumi.StringOutput)
}

// The Service Principal Role Permissions.
func (o GetServiceendpointAzurecrResultOutput) AzSpnRolePermissions() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceendpointAzurecrResult) string { return v.AzSpnRolePermissions }).(pulumi.StringOutput)
}

// The Azure Container Registry name.
func (o GetServiceendpointAzurecrResultOutput) AzurecrName() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceendpointAzurecrResult) string { return v.AzurecrName }).(pulumi.StringOutput)
}

// The Tenant ID of the service principal.
func (o GetServiceendpointAzurecrResultOutput) AzurecrSpnTenantid() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceendpointAzurecrResult) string { return v.AzurecrSpnTenantid }).(pulumi.StringOutput)
}

// The Subscription ID of the Azure targets.
func (o GetServiceendpointAzurecrResultOutput) AzurecrSubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceendpointAzurecrResult) string { return v.AzurecrSubscriptionId }).(pulumi.StringOutput)
}

// The Subscription Name of the Azure targets.
func (o GetServiceendpointAzurecrResultOutput) AzurecrSubscriptionName() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceendpointAzurecrResult) string { return v.AzurecrSubscriptionName }).(pulumi.StringOutput)
}

// The Service Endpoint description.
func (o GetServiceendpointAzurecrResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceendpointAzurecrResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetServiceendpointAzurecrResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceendpointAzurecrResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetServiceendpointAzurecrResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceendpointAzurecrResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

// The Resource Group to which the Container Registry belongs.
func (o GetServiceendpointAzurecrResultOutput) ResourceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceendpointAzurecrResult) string { return v.ResourceGroup }).(pulumi.StringOutput)
}

func (o GetServiceendpointAzurecrResultOutput) ServiceEndpointId() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceendpointAzurecrResult) string { return v.ServiceEndpointId }).(pulumi.StringOutput)
}

func (o GetServiceendpointAzurecrResultOutput) ServiceEndpointName() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceendpointAzurecrResult) string { return v.ServiceEndpointName }).(pulumi.StringOutput)
}

// The Application(Client) ID of the Service Principal.
func (o GetServiceendpointAzurecrResultOutput) ServicePrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceendpointAzurecrResult) string { return v.ServicePrincipalId }).(pulumi.StringOutput)
}

// The ID of the Service Principal.
func (o GetServiceendpointAzurecrResultOutput) SpnObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceendpointAzurecrResult) string { return v.SpnObjectId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetServiceendpointAzurecrResultOutput{})
}
