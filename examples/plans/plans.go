package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/t0mk/gometal/client"
)

func main() {
	c := client.NewHTTPClient(nil)
	auth := httptransport.APIKeyAuth("X-Auth-Token", "header", os.Getenv("PACKET_AUTH_TOKEN"))
	r, err := c.Plans.FindPlans(nil, auth)

	if err != nil {
		panic(err)
	}

	s, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}

	fmt.Printf(string(s))
	log.Println(r)
}
