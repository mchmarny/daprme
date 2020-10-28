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
* Node.js (under development)

> The `daprme` wizard is template driven, PRs with support for other languages welcomed


## Installation

### Mac OS using Homebrew

```sh
brew tap dapr-templates/daprme
brew install dapr-templates/daprme/daprme
```

### Other

1. Download the latest [daprme](https://github.com/dapr-templates/daprme/releases/latest) release for your OS
2. Move it to your desired in PATH location (e.g. `/usr/local/bin`)

## Usage 

> Assuming the `daprme` CLI is already in PATH (see Installation above)

Run `daprme` and follow the prompts

```shell
daprme
```

To specify the output directory (defaults to current), provide the `--out` flag, for example:

```shell
daprme --out ~/Downloads
```

To re-use an existing app manifest provide the `--file` flag, for example:

```shell
daprme --file ~/dapr-templates/my-common-app.yaml
```

Here is an example of an app template for a gRPC service application type in Go with a couple of components:

> The `daprme` prompt will guide you through template definition and output resulting manifest at the end, so no need to write any YAML by hand. 

```yaml
Meta:
  Name: demo
  Type: gRPC
  Lang: go
  Main: main.go
  Port: 50050
  UsesClient: true
  Owner: mchmarny
PubSubs:
- Type: pubsub.redis
  Name: redis-pubsub
  Topic: messages
Bindings:
- Type: bindings.cron
  Name: cron-binding
Services:
- Name: myService
- Name: myOtherService
Components:
- Type: secretstores.local.env
  Name: localenv-secret
- Type: state.redis
  Name: redis-store
```

## Adding Language Support 

To learn about ways you can contribute and how to setup your development environment check the [CONTRIBUTING.md](./CONTRIBUTING.md) doc. 

The best place to start is adding support for additional languages. `daprme` is uses Go templating, so adding addition language support is as simple as providing language specific templates in the [template](./template) directory. In the template, you can use any value from the context `daprme` passes to these templates.

In addition, you will need to implement the language specific provider interface `Configurable` in [pkg/lang](./pkg/lang) package. It lists the templates and provides language specific configuration. 

> When possible, aim for runnable project vs advanced features that require users to perform additional "plumbing" steps. 

## Code of Conduct

Please refer to our included [Code of Conduct](./CODE_OF_CONDUCT.md)