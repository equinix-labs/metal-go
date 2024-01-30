package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	metal "github.com/equinix-labs/metal-go/metal/v1"
)

func main() {
	var deviceId string
	flag.StringVar(&deviceId, "deviceId", "", "")
	flag.Parse()

	include := []string{}
	exclude := []string{}
	configuration := metal.NewConfiguration()
	configuration.AddDefaultHeader("X-Auth-Token", os.Getenv("METAL_AUTH_TOKEN"))
	api_client := metal.NewAPIClient(configuration)
	device, r, err := api_client.DevicesApi.FindDeviceById(context.Background(), deviceId).Include(include).Exclude(exclude).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.FindDeviceById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	} else {
		fmt.Printf("FindDeviceById, Device ID %v, project ID %v", *device.Id, device.Project.Id)

		r, err = api_client.HrefApi.FindByHref(context.Background(), device.Project).Execute()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `api_client.FindByHref``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		} else {
			fmt.Printf("FindByHref, Device ID %v, project ID %v", *device.Id, *device.Project.Id)
		}
	}
}
