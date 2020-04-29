# Azure DevOps Resource Provider

The Azure DevOps resource provider for Pulumi lets you manage Azure DevOps
resources in your cloud programs. To use this package, please [install the
Pulumi CLI first](https://pulumi.io/).

## Installing

> Currently the provider is not distributed via the Pulumi infrastructure, thus
> you've to build and install it locally.
>  
> Please refer to the corresponding [document](BuildLocally.md) for accomplish this.

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

* azuredevops:orgServiceUrl - (Required) 
* azuredevops:personalAccessToken - (Required) 

## Reference

For detailed reference documentation, please visit [the API docs](https://pulumi.io/reference/pkg/nodejs/@pulumi/azuredevops/index.html).
