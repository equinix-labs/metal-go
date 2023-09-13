/*
Metal API

Testing FirmwareSetsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package v1

import (
	"context"
	"testing"

	openapiclient "github.com/equinix-labs/metal-go/metal/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_v1_FirmwareSetsApiService(t *testing.T) {
	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test FirmwareSetsApiService GetOrganizationFirmwareSets", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.FirmwareSetsApi.GetOrganizationFirmwareSets(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test FirmwareSetsApiService GetProjectFirmwareSets", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.FirmwareSetsApi.GetProjectFirmwareSets(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})
}