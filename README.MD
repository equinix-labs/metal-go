# Golang client for Equinix Metal generated with go-swagger

This repo contains Golang code generated from customized [api-docs.json](api-docs.json) from the Equinix Metal API from https://api.equinix.com/metal/v1/api-docs.

I tried to put it together and in [client.go](client.go) there's code that should list projects. However, there probably some inconsitencies in the API definition, for instance the `customdata` parameter doesn't work out-of-the-box:

```
tomk@xps ~/gometal » ./gometal 
panic: json: cannot unmarshal object into Go struct field Project.projects.customdata of type string

goroutine 1 [running]:
main.main()
/home/tomk/gometal/client.go:17 +0x145
```

(customdata is in most of the resrouce types).


## Notes

- authentication struct must be passed to every function. All the generated functions have authWriter as an argument. I'm not sure if it's possible to set the auth header globally so that all the func use it. And I'm not sure if it's possible to remove the authWrite from the generated code.

