module github.com/pulumi/pulumi-azuredevops/provider/v2

go 1.16

replace (
	github.com/hashicorp/go-getter v1.5.0 => github.com/hashicorp/go-getter v1.4.0
	github.com/hashicorp/terraform-plugin-sdk/v2 => github.com/pulumi/terraform-plugin-sdk/v2 v2.0.0-20201218231525-9cca98608a5e
	github.com/microsoft/terraform-provider-azuredevops => github.com/pulumi/terraform-provider-azuredevops v0.1.3-0.20220420211229-578149558283
	// Fixes a bug in codegen. Can be removed when the bridge references pulumi/pulumi >= 3.29.1:
	github.com/pulumi/pulumi/pkg/v3 => github.com/pulumi/pulumi/pkg/v3 v3.29.1
)

require (
	github.com/microsoft/terraform-provider-azuredevops v0.2.0
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.20.0
	github.com/pulumi/pulumi/sdk/v3 v3.29.1
)
