# Golang client for Equinix Metal generated with go-swagger

This is experimental. Don't use it.

This repo contains Golang code generated from customized [equinix-metal.swagger.json](equinix-metal.swagger.json) from the Equinix Metal API from <https://api.equinix.com/metal/v1/api-docs>.

Contents:
- `Makefile` includes tasks to fetch the API spec, apply patches, and generate a client
- `equinix-metal.swagger.json` the latest fetched swagger spec
- `patches/*.patch` patch files to apply against the Equinix Metal API spec
- `clients/` generated clients
- `types/` generated types
- `examples/` hand crafted examples to demonstrate usage


## Build

To build the client, run `make`.

## Examples

You can see usage of the generated code in the `examples` directory. In order to try, export PACKET_AUTH token and execute the code, e.g.

```
$ PACKET_AUTH_TOKEN=... go run examples/projects/projects.go
[...]
$ cd examples/plans; PACKET_AUTH_TOKEN=... go run .
```

## Generated code

The generated functions look like (project listing):

```go
FindProjects(params *FindProjectsParams, authInfo runtime.ClientAuthInfoWriter) (*FindProjectsOK, error)
```

.. where authInfo is used for HTTP auth (auth token to header in case of Equinix Metal) and params are sth like

```go
type FindProjectsParams struct {
    Include *string
    Page *int64
    PerPage *int64
    timeout    time.Duration
    Context    context.Context
    HTTPClient *http.Client
}
```

There's an `HTTPClient` in the params and the client can maybe used for isolated HTTP traffic logging.



## Notes

- authentication struct must be passed to every function. All the generated functions have authWriter as an argument. I'm not sure if it's possible to set the auth header globally so that all the func use it. And I'm not sure if it's possible to remove the authWrite from the generated code.
