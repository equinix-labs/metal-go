package main

import (
	"log"
	"os"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/t0mk/gometal/client"
)

func main() {
	c := client.NewHTTPClient(nil)
	auth := httptransport.APIKeyAuth("X-Auth-Token", "header", os.Getenv("METAL_AUTH_TOKEN"))
	r, err := c.Projects.FindProjects(nil, auth)

	if err != nil {
		panic(err)
	}
	log.Println(r)
}
