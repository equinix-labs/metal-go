package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	metal "github.com/equinix-labs/metal-go/metal/v1"
	"sigs.k8s.io/yaml"
)

func main() {
	var projectId string
	flag.StringVar(&projectId, "projectId", "", "")
	flag.Parse()

	configuration := metal.NewConfiguration()
	configuration.Debug = true
	configuration.AddDefaultHeader("X-Auth-Token", os.Getenv("METAL_AUTH_TOKEN"))
	api_client := metal.NewAPIClient(configuration)

	findDevicesRequest := api_client.DevicesApi.FindProjectDevices(context.Background(), projectId)
	findDevicesRequest = findDevicesRequest.Exclude([]string{"ssh_keys"})
	devices, err := findDevicesRequest.ExecuteWithPagination()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.FindProjectDevices` with automatic page navigation: %v\n", err)
	}

	// response from `FindProjectDevices`: Device
	s, err := yaml.Marshal(devices)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(os.Stdout, "Response from `DevicesApi.FindProjectDevices`: %v\n", string(s))
}
