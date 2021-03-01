package main

import (
	"log"
	"os"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/t0mk/gometal/client"
	"github.com/t0mk/gometal/client/projects"
	"github.com/t0mk/gometal/types"
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

	page := int64(1)
	params := projects.NewFindProjectsParams()
	params.SetDefaults()
	params.SetPage(&page)

	projs := []*types.Project{}

	for {
		next, err := c.Projects.FindProjects(params, auth)
		if err != nil {
			log.Fatal(err)
		}

		payload := next.GetPayload()
		projs = append(projs, payload.Projects...)

		if payload.Meta.Next == nil || payload.Meta.Next.Href == nil {
			break
		}

		page++
	}

	s, err := yaml.Marshal(projs)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(string(s))
}
