module github.com/pulumi/pulumi-azuredevops/provider

go 1.15

replace github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible

require (
	github.com/microsoft/terraform-provider-azuredevops v0.1.1
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.16.0
	github.com/pulumi/pulumi/sdk/v2 v2.16.0
)
