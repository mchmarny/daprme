# daprme

[![Go Report Card](https://goreportcard.com/badge/github.com/dapr-templates/daprme)](https://goreportcard.com/report/github.com/dapr-templates/daprme) ![Test](https://github.com/dapr-templates/daprme/workflows/Test/badge.svg) ![Release](https://github.com/dapr-templates/daprme/workflows/Release/badge.svg) [![codecov](https://codecov.io/gh/dapr-templates/daprme/branch/master/graph/badge.svg)](https://codecov.io/gh/dapr-templates/daprme)

This new app wizard CLI guides you through the creation of a new Dapr project. Including the Dapr component selection and scaffolding of a new application.

Currently supported application types:

* Command-line (e.g. CLI)
* HTTP Service 
* gRPC Service 

> Each one of these applications supports options have client to call into Dapr API

Currently supported development languages: 

* Go 

> The `daprme` wizard is template driven, PRs with support for other languages welcomed


## Installation

The `daprme` CLI includes support for OSes (architecture support coming soon)

1. Download the latest [daprme](https://github.com/dapr-templates/daprme/releases/latest) release for your OS
2. Move it to your desired in PATH location
   * For Linux/MacOS - `/usr/local/bin`
   * For Windows, create a directory and add this to your System PATH

## Usage 

> Assuming the `daprme` CLI is already in PATH (see Installation above)

Run `daprme` and follow the prompts

```shell
daprme
````