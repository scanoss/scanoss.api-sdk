# scanoss.api-sdk (Go types)

Generated Go types for the [SCANOSS REST API](https://github.com/scanoss/scanoss.api).
These are the request/response DTOs the API is documented with in its
`docs/openapi.yaml`, produced with [`oapi-codegen`](https://github.com/oapi-codegen/oapi-codegen).
Import them so your service (un)marshals against the exact structs the API ships.

> **Generated — do not edit.** `types.gen.go` is produced from the API's OpenAPI
> spec and published here automatically on every API release. Hand edits will be
> overwritten. File issues and spec changes against
> [`scanoss/scanoss.api`](https://github.com/scanoss/scanoss.api).

## Install

```sh
go get github.com/scanoss/scanoss.api-sdk@vX.Y.Z
```

## Versioning

This package is **published on every `vX.Y.Z` tag of the API, with the same
version**: `scanoss.api-sdk vX.Y.Z` contains the types generated from
`scanoss.api vX.Y.Z`. Pin the SDK to the API version your service targets.

## Usage

```go
import scanossapi "github.com/scanoss/scanoss.api-sdk"

var env scanossapi.ScanEnvelope
_ = json.Unmarshal(body, &env)
```

## License

[MIT](./LICENSE) © SCANOSS
