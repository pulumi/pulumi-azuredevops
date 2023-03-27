// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package core

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about the Azure DevOps organization configured for the provider.
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
//			example, err := azuredevops.GetClientConfig(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("orgUrl", example.OrganizationUrl)
//			return nil
//		})
//	}
//
// ```
//
// Deprecated: azuredevops.core.getClientConfig has been deprecated in favor of azuredevops.getClientConfig
func GetClientConfig(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetClientConfigResult, error) {
	var rv GetClientConfigResult
	err := ctx.Invoke("azuredevops:Core/getClientConfig:getClientConfig", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getClientConfig.
type GetClientConfigResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id              string `pulumi:"id"`
	OrganizationUrl string `pulumi:"organizationUrl"`
}
