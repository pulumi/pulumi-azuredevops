// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

// The url of the Azure DevOps instance which should be used.
func GetOrgServiceUrl(ctx *pulumi.Context) string {
	v, err := config.Try(ctx, "azuredevops:orgServiceUrl")
	if err == nil {
		return v
	}
	return getEnvOrDefault("", nil, "AZDO_ORG_SERVICE_URL").(string)
}

// The personal access token which should be used.
func GetPersonalAccessToken(ctx *pulumi.Context) string {
	return config.Get(ctx, "azuredevops:personalAccessToken")
}
