module github.com/pulumi/pulumi-azuredevops/provider

go 1.14

replace github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible

require (
	github.com/hashicorp/terraform-plugin-sdk v1.10.0
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.6.0
	github.com/pulumi/pulumi/sdk/v2 v2.7.1
	github.com/terraform-providers/terraform-provider-azuredevops v0.0.1
)
