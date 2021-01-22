# Golang client for Equinix Metal generated with go-swagger

This is experimental. Don't use it.

This repo contains Golang code generated from customized [equinix-metal.swagger.json](equinix-metal.swagger.json) from the Equinix Metal API from <https://api.equinix.com/metal/v1/api-docs>.

Contents:
- `gen.sh a script to fetch the API spec, apply patches, and generate a client
- `equinix-metal.swagger.json` the latest fetched swagger spec
- `\*.patch` patch files to apply against the Equinix Metal API spec
- `clients/` generated clients
- `types/` generated types
- `examples/` hand crafted examples to demonstrate usage


I tried to put it together and in [examples/projects/projects.go](examples/projects/projects.go) there's code that should list projects. However, there probably some inconsitencies in the API definition, for instance the `customdata` parameter doesn't work out-of-the-box:

```
tomk@xps ~/gometal » PACKET_AUTH_TOKEN=... go run examples/projects/projects.go
panic: json: cannot unmarshal object into Go struct field Project.projects.customdata of type string

goroutine 1 [running]:
main.main()
/home/tomk/gometal/client.go:17 +0x145
```

(customdata is in most of the resrouce types).

Another example in [examples/plans/plans.go](examples/plans/plans.go) runs without error.

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

There's HTTPCleint in the params and the client can maybe used for HTTP traffic logging, but it wouldn't be trivial.



## Notes

- authentication struct must be passed to every function. All the generated functions have authWriter as an argument. I'm not sure if it's possible to set the auth header globally so that all the func use it. And I'm not sure if it's possible to remove the authWrite from the generated code.
- there seems to be no simple way to print HTTP requests and responses in the generated client:
  https://github.com/swagger-api/swagger-codegen/issues/5703
