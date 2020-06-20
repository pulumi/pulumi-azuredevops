# Azure DevOps Resource Provider

The Azure DevOps resource provider for Pulumi lets you manage Azure DevOps
resources in your cloud programs. To use this package, please [install the
Pulumi CLI first](https://pulumi.io/).

## Installing

This package is available in many languages in the standard packaging formats.

### Node.js (Java/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

    $ npm install @pulumi/azuredevops

or `yarn`:

    $ yarn add @pulumi/azuredevops

### Python

To use from Python, install using `pip`:

    $ pip install pulumi_azuredevops

### Go

To use from Go, use `go get` to grab the latest version of the library

    $ go get github.com/pulumi/pulumi-azuredevops/sdk/go/...

## Configuration

The following configuration points are available:

* azuredevops:orgServiceUrl - (Required) This is the Azure DevOps organization url. It can also be sourced from the 
  `AZDO_ORG_SERVICE_URL` environment variable.
* azuredevops:personalAccessToken - (Required) This is the Azure DevOps organization personal access token. The account
  corresponding to the token will need "owner" privileges for this organization. It can also be sourced from the 
  `AZDO_PERSONAL_ACCESS_TOKEN` environment variable.

## Reference

For further information, please visit [the AzureDevOps provider docs](https://www.pulumi.com/docs/intro/cloud-providers/azuredevops) 
or for detailed reference documentation, please visit [the API docs](https://www.pulumi.com/docs/reference/pkg/azuredevops).
