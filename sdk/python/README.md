[![Actions Status](https://github.com/pulumi/pulumi-external/workflows/master/badge.svg)](https://github.com/pulumi/pulumi-external/actions)
[![NPM version](https://img.shields.io/npm/v/@pulumi/external)](https://www.npmjs.com/package/@pulumi/external)
[![Python version](https://img.shields.io/pypi/v/pulumi_external)](https://pypi.org/project/pulumi_external)
[![NuGet version](https://img.shields.io/nuget/v/Pulumi.External)](https://www.nuget.org/packages/Pulumi.External)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/pulumi/pulumi-external/sdk/go)](https://pkg.go.dev/github.com/pulumi/pulumi-external/sdk/go)
[![License](https://img.shields.io/github/license/pulumi/pulumi-external)](https://github.com/pulumi/pulumi-external/blob/master/LICENSE)

# External Resource Provider

This provider is mainly used for ease of converting terraform programs to Pulumi.
When using standard Pulumi programs, you would not need to use this provider. 

The External resource provider for Pulumi lets you use External resources in your cloud programs.
To use this package, please [install the Pulumi CLI first](https://www.pulumi.com/docs/install/).

## Installing

This package is available in many languages in the standard packaging formats.

### Node.js (Java/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

    $ npm install @pulumi/external

or `yarn`:

    $ yarn add @pulumi/external

### Python

To use from Python, install using `pip`:

    $ pip install pulumi_external

### Go

To use from Go, use `go get` to grab the latest version of the library:

    $ go get github.com/pulumi/pulumi-external/sdk

### .NET

To use from .NET, install using `dotnet add package`:

    $ dotnet add package Pulumi.External

<!-- If your provider has configuration, remove this comment and the comment tags below, updating the documentation. -->
<!--

## Configuration

The following Pulumi configuration can be used:

- `external:token` - (Required) The API token to use with External. When not set, the provider will use the `EXTERNAL_TOKEN` environment variable.

-->

<!-- If your provider has reference material available elsewhere, remove this comment and the comment tags below, updating the documentation. -->
<!--

## Reference

For further information, please visit [External reference documentation](https://example.com/external).

-->
