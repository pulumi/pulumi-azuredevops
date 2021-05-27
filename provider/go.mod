module github.com/pulumi/pulumi-azuredevops/provider/v2

go 1.16

replace (
	github.com/hashicorp/go-getter v1.5.0 => github.com/hashicorp/go-getter v1.4.0
	github.com/hashicorp/terraform-plugin-sdk/v2 => github.com/pulumi/terraform-plugin-sdk/v2 v2.0.0-20201218231525-9cca98608a5e
	github.com/microsoft/terraform-provider-azuredevops => github.com/pulumi/terraform-provider-azuredevops v0.1.3-0.20210418192049-8caffc4f5c73
)

require (
	github.com/microsoft/terraform-provider-azuredevops v0.1.2
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.2.1
	github.com/pulumi/pulumi/pkg/v3 v3.3.2-0.20210526172205-85142462c7ed
	github.com/pulumi/pulumi/sdk/v3 v3.3.2-0.20210526172205-85142462c7ed
)
