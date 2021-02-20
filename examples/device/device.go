package main

import (
	"context"
	"fmt"
	"os"
	"time"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/runtime/logger"
	"github.com/go-openapi/strfmt"
	"github.com/t0mk/gometal/client"
	"github.com/t0mk/gometal/client/projects"
	"github.com/t0mk/gometal/types"
	"sigs.k8s.io/yaml"
)

func GetEnvStrict(envVarName string) string {
	ret := os.Getenv(envVarName)
	if len(ret) == 0 {
		panic(fmt.Errorf("You must set envvar %s", envVarName))
	}
	return ret
}

var (
	token     = GetEnvStrict("METAL_AUTH_TOKEN")
	p         = GetEnvStrict("METAL_PROJECT_ID")
	hostname  = "test"
	projectID = strfmt.UUID(p)
	plan      = "c3.small.x86"
	facility  = "da11"
	image     = "ubuntu_18_04"
	billing   = "hourly"
	// Remove device immediately
	terminate = time.Now().Local().Add(time.Second * time.Duration(1))

	trueBool  = true
	falseBool = false

	l = &logger.StandardLogger{}
)

func NewDebugClient(formats strfmt.Registry, cfg *client.TransportConfig) *client.MetalAPI {

	// ensure nullable parameters have default
	if cfg == nil {
		cfg = client.DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	transport.SetDebug(true)
	transport.SetLogger(l)
	return client.New(transport, formats)
}

func main() {
	c := NewDebugClient(nil, nil)
	auth := httptransport.APIKeyAuth("X-Auth-Token", "header", token)

	params := &projects.CreateDeviceParams{
		ID:      projectID,
		Context: context.Background(),
		Device: &types.DeviceCreateInput{
			Hostname: hostname,
			IPAddresses: []*types.DeviceCreateInputIPAddressesItems0{
				{
					Public:        &trueBool,
					Cidr:          31,
					AddressFamily: 4,
					// IPReservations: []string{},
				},
				{
					Public:        &trueBool,
					Cidr:          127,
					AddressFamily: 6,
					// IPReservations: []string{},
				},
				{ // TODO private address is not optional if ip_addresses is used
					// TODO: this must be set explicity in the serialized request.
					// --t0mk I think the false is put to the request, because
					// type of the "public" field is bool pointer which has null as default.
					Public:        &falseBool,
					Cidr:          28,
					AddressFamily: 4,
					// IPReservations: []string{},
				},
			},
			Plan:            &plan,
			Facility:        &facility,
			OperatingSystem: &image,
			BillingCycle:    billing,

			// TODO: go-swagger/go-swagger#2518
			TerminationTime: strfmt.DateTime(terminate),

			Tags:           []string{"test"},
			Features:       []string{}, // TODO: what would I put here? "Features are not supported on the c3.small.x86 plan: baremetal, backend_transfer, and layer_2"
			ProjectSSHKeys: []strfmt.UUID{},
		},
	}

	r, err := c.Projects.CreateDevice(params, auth)

	if err != nil {
		panic(err)
	}

	dev := r.GetPayload()

	s, err := yaml.Marshal(dev)
	if err != nil {
		panic(err)
	}

	l.Printf("%+v\n", string(s))
}
