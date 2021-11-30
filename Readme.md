> This fork removes all POC code and fixes a cyclic depedency go mod creates with `github.com/danielvladco/go-proto-gql/pb`

Protoc plugins for generating graphql schema and go graphql code

If you use micro-service architecture with grpc for back-end and graphql api gateway for front-end, you will find yourself
repeating a lot of code for translating from one transport layer to another (which many times may be a source of bugs)

This repository aims to simplify working with grpc trough protocol buffers and graphql by generating code.

## Install:

```sh
go install github.com/togatv/go-proto-gql/protoc-gen-gql
```

## Usage Examples:

The protoc compiler expects to find plugins named `proto-gen-<PLUGIN_NAME>` on the execution `$PATH`. So first:

```sh
export PATH=${PATH}:${GOPATH}/bin
```

---

`--gql_out` plugin will generate graphql files with extension `.graphql`
rather than go code which means it can be further used for any other language or framework.

Example:

```sh
protoc --gql_out=paths=source_relative:. -I=. -I=./example/ ./example/*.proto
```

---

See `/example` folder for more examples.

## Community:

Will be very glad for any contributions so feel free to create issues, forks and PRs.

## License:

`go-proto-gql` is released under the Apache 2.0 license. See the LICENSE file for details.
