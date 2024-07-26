// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azuredevops/sdk/v3/go/azuredevops/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing AzureRM service Endpoint.
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
//			serviceendpoint, err := azuredevops.LookupServiceEndpointAzureRM(ctx, &azuredevops.LookupServiceEndpointAzureRMArgs{
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
//			serviceendpoint, err := azuredevops.LookupServiceEndpointAzureRM(ctx, &azuredevops.LookupServiceEndpointAzureRMArgs{
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
func LookupServiceEndpointAzureRM(ctx *pulumi.Context, args *LookupServiceEndpointAzureRMArgs, opts ...pulumi.InvokeOption) (*LookupServiceEndpointAzureRMResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupServiceEndpointAzureRMResult
	err := ctx.Invoke("azuredevops:index/getServiceEndpointAzureRM:getServiceEndpointAzureRM", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getServiceEndpointAzureRM.
type LookupServiceEndpointAzureRMArgs struct {
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

// A collection of values returned by getServiceEndpointAzureRM.
type LookupServiceEndpointAzureRMResult struct {
	// Specifies the Authorization Scheme Map.
	Authorization map[string]string `pulumi:"authorization"`
	// Specified the Management Group ID of the Service Endpoint is target, if available.
	AzurermManagementGroupId string `pulumi:"azurermManagementGroupId"`
	// Specified the Management Group Name of the Service Endpoint target, if available.
	AzurermManagementGroupName string `pulumi:"azurermManagementGroupName"`
	// Specifies the Tenant ID of the Azure targets.
	AzurermSpnTenantid string `pulumi:"azurermSpnTenantid"`
	// Specifies the Subscription ID of the Service Endpoint target, if available.
	AzurermSubscriptionId string `pulumi:"azurermSubscriptionId"`
	// Specifies the Subscription Name of the Service Endpoint target, if available.
	AzurermSubscriptionName string `pulumi:"azurermSubscriptionName"`
	// Specifies the description of the Service Endpoint.
	Description string `pulumi:"description"`
	// The Cloud Environment. Possible values are `AzureCloud`, `AzureChinaCloud`, `AzureUSGovernment`, and `AzureGermanCloud`.
	Environment string `pulumi:"environment"`
	// The provider-assigned unique ID for this managed resource.
	Id        string `pulumi:"id"`
	ProjectId string `pulumi:"projectId"`
	// Specifies the Resource Group of the Service Endpoint target, if available.
	ResourceGroup string `pulumi:"resourceGroup"`
	// Specifies the authentication scheme of azurerm endpoint, either `WorkloadIdentityFederation`, `ManagedServiceIdentity` or `ServicePrincipal`.
	ServiceEndpointAuthenticationScheme string `pulumi:"serviceEndpointAuthenticationScheme"`
	ServiceEndpointId                   string `pulumi:"serviceEndpointId"`
	ServiceEndpointName                 string `pulumi:"serviceEndpointName"`
	// The issuer if `serviceEndpointAuthenticationScheme` is set to `WorkloadIdentityFederation`. This looks like `https://vstoken.dev.azure.com/f66a4bc2-08ad-4ec0-a25e-e769d6b3b294`, where the GUID is the Organization ID of your Azure DevOps Organisation.
	WorkloadIdentityFederationIssuer string `pulumi:"workloadIdentityFederationIssuer"`
	// The subject if `serviceEndpointAuthenticationScheme` is set to `WorkloadIdentityFederation`. This looks like `sc://my-organisation/my-project/my-service-connection-name`.
	WorkloadIdentityFederationSubject string `pulumi:"workloadIdentityFederationSubject"`
}

func LookupServiceEndpointAzureRMOutput(ctx *pulumi.Context, args LookupServiceEndpointAzureRMOutputArgs, opts ...pulumi.InvokeOption) LookupServiceEndpointAzureRMResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServiceEndpointAzureRMResult, error) {
			args := v.(LookupServiceEndpointAzureRMArgs)
			r, err := LookupServiceEndpointAzureRM(ctx, &args, opts...)
			var s LookupServiceEndpointAzureRMResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupServiceEndpointAzureRMResultOutput)
}

// A collection of arguments for invoking getServiceEndpointAzureRM.
type LookupServiceEndpointAzureRMOutputArgs struct {
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

func (LookupServiceEndpointAzureRMOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceEndpointAzureRMArgs)(nil)).Elem()
}

// A collection of values returned by getServiceEndpointAzureRM.
type LookupServiceEndpointAzureRMResultOutput struct{ *pulumi.OutputState }

func (LookupServiceEndpointAzureRMResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceEndpointAzureRMResult)(nil)).Elem()
}

func (o LookupServiceEndpointAzureRMResultOutput) ToLookupServiceEndpointAzureRMResultOutput() LookupServiceEndpointAzureRMResultOutput {
	return o
}

func (o LookupServiceEndpointAzureRMResultOutput) ToLookupServiceEndpointAzureRMResultOutputWithContext(ctx context.Context) LookupServiceEndpointAzureRMResultOutput {
	return o
}

// Specifies the Authorization Scheme Map.
func (o LookupServiceEndpointAzureRMResultOutput) Authorization() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupServiceEndpointAzureRMResult) map[string]string { return v.Authorization }).(pulumi.StringMapOutput)
}

// Specified the Management Group ID of the Service Endpoint is target, if available.
func (o LookupServiceEndpointAzureRMResultOutput) AzurermManagementGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceEndpointAzureRMResult) string { return v.AzurermManagementGroupId }).(pulumi.StringOutput)
}

// Specified the Management Group Name of the Service Endpoint target, if available.
func (o LookupServiceEndpointAzureRMResultOutput) AzurermManagementGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceEndpointAzureRMResult) string { return v.AzurermManagementGroupName }).(pulumi.StringOutput)
}

// Specifies the Tenant ID of the Azure targets.
func (o LookupServiceEndpointAzureRMResultOutput) AzurermSpnTenantid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceEndpointAzureRMResult) string { return v.AzurermSpnTenantid }).(pulumi.StringOutput)
}

// Specifies the Subscription ID of the Service Endpoint target, if available.
func (o LookupServiceEndpointAzureRMResultOutput) AzurermSubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceEndpointAzureRMResult) string { return v.AzurermSubscriptionId }).(pulumi.StringOutput)
}

// Specifies the Subscription Name of the Service Endpoint target, if available.
func (o LookupServiceEndpointAzureRMResultOutput) AzurermSubscriptionName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceEndpointAzureRMResult) string { return v.AzurermSubscriptionName }).(pulumi.StringOutput)
}

// Specifies the description of the Service Endpoint.
func (o LookupServiceEndpointAzureRMResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceEndpointAzureRMResult) string { return v.Description }).(pulumi.StringOutput)
}

// The Cloud Environment. Possible values are `AzureCloud`, `AzureChinaCloud`, `AzureUSGovernment`, and `AzureGermanCloud`.
func (o LookupServiceEndpointAzureRMResultOutput) Environment() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceEndpointAzureRMResult) string { return v.Environment }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupServiceEndpointAzureRMResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceEndpointAzureRMResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupServiceEndpointAzureRMResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceEndpointAzureRMResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

// Specifies the Resource Group of the Service Endpoint target, if available.
func (o LookupServiceEndpointAzureRMResultOutput) ResourceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceEndpointAzureRMResult) string { return v.ResourceGroup }).(pulumi.StringOutput)
}

// Specifies the authentication scheme of azurerm endpoint, either `WorkloadIdentityFederation`, `ManagedServiceIdentity` or `ServicePrincipal`.
func (o LookupServiceEndpointAzureRMResultOutput) ServiceEndpointAuthenticationScheme() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceEndpointAzureRMResult) string { return v.ServiceEndpointAuthenticationScheme }).(pulumi.StringOutput)
}

func (o LookupServiceEndpointAzureRMResultOutput) ServiceEndpointId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceEndpointAzureRMResult) string { return v.ServiceEndpointId }).(pulumi.StringOutput)
}

func (o LookupServiceEndpointAzureRMResultOutput) ServiceEndpointName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceEndpointAzureRMResult) string { return v.ServiceEndpointName }).(pulumi.StringOutput)
}

// The issuer if `serviceEndpointAuthenticationScheme` is set to `WorkloadIdentityFederation`. This looks like `https://vstoken.dev.azure.com/f66a4bc2-08ad-4ec0-a25e-e769d6b3b294`, where the GUID is the Organization ID of your Azure DevOps Organisation.
func (o LookupServiceEndpointAzureRMResultOutput) WorkloadIdentityFederationIssuer() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceEndpointAzureRMResult) string { return v.WorkloadIdentityFederationIssuer }).(pulumi.StringOutput)
}

// The subject if `serviceEndpointAuthenticationScheme` is set to `WorkloadIdentityFederation`. This looks like `sc://my-organisation/my-project/my-service-connection-name`.
func (o LookupServiceEndpointAzureRMResultOutput) WorkloadIdentityFederationSubject() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceEndpointAzureRMResult) string { return v.WorkloadIdentityFederationSubject }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServiceEndpointAzureRMResultOutput{})
}
