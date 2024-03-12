[![Actions Status](https://github.com/pulumi/pulumi-azuredevops/workflows/master/badge.svg)](https://github.com/pulumi/pulumi-azuredevops/actions)
[![Slack](http://www.pulumi.com/images/docs/badges/slack.svg)](https://slack.pulumi.com)
[![NPM version](https://badge.fury.io/js/%40pulumi%2Fazuredevops.svg)](https://www.npmjs.com/package/@pulumi/azuredevops)
[![Python version](https://badge.fury.io/py/pulumi-azuredevops.svg)](https://pypi.org/project/pulumi-azuredevops)
[![NuGet version](https://badge.fury.io/nu/pulumi.azuredevops.svg)](https://badge.fury.io/nu/pulumi.azuredevops)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/pulumi/pulumi-azuredevops/sdk/v3/go)](https://pkg.go.dev/github.com/pulumi/pulumi-azuredevops/sdk/v3/go)
[![License](https://img.shields.io/npm/l/%40pulumi%2Fpulumi.svg)](https://github.com/pulumi/pulumi-azuredevops/blob/master/LICENSE)

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

    $ go get github.com/pulumi/pulumi-azuredevops/sdk/v3

## Configuration

The following configuration points are available:

* `azuredevops:orgServiceUrl` - (Required) This is the Azure DevOps organization url. It can also be sourced from the `AZDO_ORG_SERVICE_URL` environment variable.
* `azuredevops:personalAccessToken` - (Required) This is the Azure DevOps organization personal access token. The account corresponding to the token will need "owner" privileges for this organization. It can also be sourced from the `AZDO_PERSONAL_ACCESS_TOKEN` environment variable.

## Reference

For further information, please visit [the AzureDevOps provider docs](https://www.pulumi.com/registry/packages/azuredevops/) 
or for detailed reference documentation, please visit [the API docs](https://www.pulumi.com/registry/packages/azuredevops/api-docs/).
