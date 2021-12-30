module github.com/pulumi/pulumi-azuredevops/provider/v2

go 1.16

replace (
	github.com/hashicorp/go-getter v1.5.0 => github.com/hashicorp/go-getter v1.4.0
	github.com/hashicorp/terraform-plugin-sdk/v2 => github.com/pulumi/terraform-plugin-sdk/v2 v2.0.0-20201218231525-9cca98608a5e
	github.com/microsoft/terraform-provider-azuredevops => github.com/pulumi/terraform-provider-azuredevops v0.1.3-0.20211229165113-8b59b8c3e421
)

require (
	github.com/microsoft/terraform-provider-azuredevops v0.1.8
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.11.0
	github.com/pulumi/pulumi/sdk/v3 v3.17.0
)
