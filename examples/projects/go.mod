module github.com/t0mk/gometal/examples/device

go 1.15

replace github.com/t0mk/gometal => ../../

require (
	github.com/go-openapi/runtime v0.19.26
	github.com/t0mk/gometal v0.0.0-20210218083440-d503aa5a375a
	sigs.k8s.io/yaml v1.2.0
)
