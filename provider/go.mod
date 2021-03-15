module github.com/pulumi/pulumi-azuredevops/provider

go 1.16

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
	github.com/hashicorp/go-getter v1.5.0 => github.com/hashicorp/go-getter v1.4.0
	github.com/hashicorp/terraform-plugin-sdk/v2 => github.com/pulumi/terraform-plugin-sdk/v2 v2.0.0-20201218231525-9cca98608a5e
	github.com/microsoft/terraform-provider-azuredevops => github.com/pulumi/terraform-provider-azuredevops v0.1.3-0.20210210180928-1e7cebb0ffb7
)

require (
	github.com/microsoft/terraform-provider-azuredevops v0.1.2
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.21.0
	github.com/pulumi/pulumi/sdk/v2 v2.22.1-0.20210310211618-1f16423ede4c
)
