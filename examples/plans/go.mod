module github.com/t0mk/gometal/examples/plans

go 1.15

replace github.com/t0mk/gometal => ../..

require (
	github.com/go-openapi/runtime v0.19.26
	github.com/go-openapi/strfmt v0.20.0
	github.com/t0mk/gometal v0.0.0-20210124095926-212875b1b106
	sigs.k8s.io/yaml v1.2.0
)
