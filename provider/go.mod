module github.com/pulumi/pulumi-azuredevops/provider

go 1.14

replace github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible

require (
	github.com/microsoft/terraform-provider-azuredevops v0.0.2-0.20200910101109-763483ebe715
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.11.0
	github.com/pulumi/pulumi/sdk/v2 v2.12.0
)
