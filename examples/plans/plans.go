package main

import (
	"fmt"
	"os"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/t0mk/gometal/client"
	"sigs.k8s.io/yaml"
)

func NewDebugClient(formats strfmt.Registry, cfg *client.TransportConfig) *client.MetalAPI {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = client.DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	transport.SetDebug(true)
	return client.New(transport, formats)
}

func main() {
	c := NewDebugClient(nil, nil)
	auth := httptransport.APIKeyAuth("X-Auth-Token", "header", os.Getenv("METAL_AUTH_TOKEN"))

	r, err := c.Plans.FindPlans(nil, auth)

	if err != nil {
		panic(err)
	}

	s, err := yaml.Marshal(r.GetPayload().Plans)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(s))
}
