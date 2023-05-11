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
	ctx := context.Background()
	metro := "da"

	createVrfRequest := metal.VrfCreateInput{
		Name:     "my-test-vrf-shared",
		Metro:    metro,
		IpRanges: []string{"10.0.0.0/16"},
	}

	gateways, r, err := api_client.MetalGatewaysApi.FindMetalGatewaysByProject(ctx, projectId).Execute()

	for _, gateway := range gateways.MetalGateways {
		api_client.MetalGatewaysApi.DeleteMetalGateway(ctx, gateway.MetalGateway.GetId()).Execute()
	}

	vrfs, r, err := api_client.VRFsApi.FindVrfs(ctx, projectId).Execute()

	for _, vrf := range vrfs.Vrfs {
		api_client.VRFsApi.DeleteVrf(ctx, vrf.GetId()).Execute()
	}

	vrf, r, err := api_client.VRFsApi.CreateVrf(ctx, projectId).VrfCreateInput(createVrfRequest).Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VRFsApi.CreateVrf`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		panic(err)
	}

	ipReservationRequest := metal.RequestIPReservationRequest{}
	ipReservationRequest.VrfIpReservationCreateInput = &metal.VrfIpReservationCreateInput{
		Type:    "vrf",
		Cidr:    24,
		Network: "10.0.0.0",
		VrfId:   vrf.GetId(),
	}

	block, r, err := api_client.IPAddressesApi.RequestIPReservation(ctx, projectId).RequestIPReservationRequest(ipReservationRequest).Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPAddressesApi.RequestIPReservation`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		panic(err)
	}

	createVlanRequest := metal.VirtualNetworkCreateInput{
		Metro: &metro,
	}

	vlan, r, err := api_client.VLANsApi.CreateVirtualNetwork(ctx, projectId).VirtualNetworkCreateInput(createVlanRequest).Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VLANsApi.CreateVirtualNetwork`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		panic(err)
	}

	createGatewayRequest := metal.CreateMetalGatewayRequest{}
	createGatewayRequest.VrfMetalGatewayCreateInput = &metal.VrfMetalGatewayCreateInput{
		IpReservationId:  block.VrfIpReservation.GetId(),
		VirtualNetworkId: vlan.GetId(),
	}

	_, r, err = api_client.MetalGatewaysApi.CreateMetalGateway(ctx, projectId).CreateMetalGatewayRequest(createGatewayRequest).Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetalGatewaysApi.CreateMetalGateway`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		panic(err)
	}

	var speed int32 = 500000000
	createConnectionRequest := metal.CreateOrganizationInterconnectionRequest{}
	createConnectionRequest.VrfFabricVcCreateInput = &metal.VrfFabricVcCreateInput{
		Metro:            metro,
		Name:             "metal-go-example",
		Redundancy:       "redundant",
		ServiceTokenType: "a_side",
		Type:             "shared",
		Vrfs:             []string{vrf.GetId(), vrf.GetId()},
		Speed:            &speed,
	}

	connection, r, err := api_client.InterconnectionsApi.CreateProjectInterconnection(ctx, projectId).CreateOrganizationInterconnectionRequest(createConnectionRequest).Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InterconnectionsApi.CreateOrganizationInterconnection`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		panic(err)
	}

	subnet := "10.0.0.2/31"
	metalIp := "10.0.0.2"
	customerIp := "10.0.0.3"
	peerAsn := int32(100)
	virtualCircuitUpdateRequest := metal.VirtualCircuitUpdateInput{}
	virtualCircuitUpdateRequest.VrfVirtualCircuitUpdateInput = &metal.VrfVirtualCircuitUpdateInput{
		PeerAsn:    &peerAsn,
		Subnet:     &subnet,
		MetalIp:    &metalIp,
		CustomerIp: &customerIp,
	}

	for _, port := range connection.Ports {
		for _, vc := range port.VirtualCircuits {
			vcId := vc.VrfVirtualCircuit.GetId()
			_, r, err := api_client.InterconnectionsApi.UpdateVirtualCircuit(ctx, vcId).VirtualCircuitUpdateInput(virtualCircuitUpdateRequest).Execute()

			if err != nil {
				fmt.Fprintf(os.Stderr, "Error when calling `InterconnectionsApi.UpdateVirtualCircuit`: %v\n", err)
				fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
				panic(err)
			}
		}
	}

	createVRFRouteRequest := metal.VrfRouteCreateInput{
		Prefix:  "10.0.5.0/0",
		NextHop: "10.0.5.2",
	}

	vrfRoute, r, err := api_client.VRFsApi.CreateVrfRoute(ctx, vrf.GetId()).VrfRouteCreateInput(createVRFRouteRequest).Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VRFsApi.CreateVrfRoute`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		panic(err)
	}

	s, err := yaml.Marshal(vrf)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(os.Stdout, "Response from `VRFsApi.CreateVrf`: \n%v\n", string(s))

	s, err = yaml.Marshal(vrfRoute)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(os.Stdout, "Response from `VRFsApi.CreateVrfRoute`: \n%v\n", string(s))
}
