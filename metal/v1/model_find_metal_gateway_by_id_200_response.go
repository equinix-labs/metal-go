/*
Metal API

# Introduction Equinix Metal provides a RESTful HTTP API which can be reached at <https://api.equinix.com/metal/v1>. This document describes the API and how to use it.  The API allows you to programmatically interact with all of your Equinix Metal resources, including devices, networks, addresses, organizations, projects, and your user account. Every feature of the Equinix Metal web interface is accessible through the API.  The API docs are generated from the Equinix Metal OpenAPI specification and are officially hosted at <https://metal.equinix.com/developers/api>.  # Common Parameters  The Equinix Metal API uses a few methods to minimize network traffic and improve throughput. These parameters are not used in all API calls, but are used often enough to warrant their own section. Look for these parameters in the documentation for the API calls that support them.  ## Pagination  Pagination is used to limit the number of results returned in a single request. The API will return a maximum of 100 results per page. To retrieve additional results, you can use the `page` and `per_page` query parameters.  The `page` parameter is used to specify the page number. The first page is `1`. The `per_page` parameter is used to specify the number of results per page. The maximum number of results differs by resource type.  ## Sorting  Where offered, the API allows you to sort results by a specific field. To sort results use the `sort_by` query parameter with the root level field name as the value. The `sort_direction` parameter is used to specify the sort direction, either either `asc` (ascending) or `desc` (descending).  ## Filtering  Filtering is used to limit the results returned in a single request. The API supports filtering by certain fields in the response. To filter results, you can use the field as a query parameter.  For example, to filter the IP list to only return public IPv4 addresses, you can filter by the `type` field, as in the following request:  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/projects/id/ips?type=public_ipv4 ```  Only IP addresses with the `type` field set to `public_ipv4` will be returned.  ## Searching  Searching is used to find matching resources using multiple field comparissons. The API supports searching in resources that define this behavior. The fields available for search differ by resource, as does the search strategy.  To search resources you can use the `search` query parameter.  ## Include and Exclude  For resources that contain references to other resources, sucha as a Device that refers to the Project it resides in, the Equinix Metal API will returns `href` values (API links) to the associated resource.  ```json {   ...   \"project\": {     \"href\": \"/metal/v1/projects/f3f131c8-f302-49ef-8c44-9405022dc6dd\"   } } ```  If you're going need the project details, you can avoid a second API request.  Specify the contained `href` resources and collections that you'd like to have included in the response using the `include` query parameter.  For example:    ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=projects ```  The `include` parameter is generally accepted in `GET`, `POST`, `PUT`, and `PATCH` requests where `href` resources are presented.  To have multiple resources include, use a comma-separated list (e.g. `?include=emails,projects,memberships`).  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=emails,projects,memberships ```  You may also include nested associations up to three levels deep using dot notation (`?include=memberships.projects`):  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=memberships.projects ```  To exclude resources, and optimize response delivery, use the `exclude` query parameter. The `exclude` parameter is generally accepted in `GET`, `POST`, `PUT`, and `PATCH` requests for fields with nested object responses. When excluded, these fields will be replaced with an object that contains only an `href` field.

API version: 1.0.0
Contact: support@equinixmetal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
	"fmt"
)

// FindMetalGatewayById200Response - struct for FindMetalGatewayById200Response
type FindMetalGatewayById200Response struct {
	FindMetalGatewayById200ResponseOneOf  *FindMetalGatewayById200ResponseOneOf
	FindMetalGatewayById200ResponseOneOf1 *FindMetalGatewayById200ResponseOneOf1
}

// FindMetalGatewayById200ResponseOneOfAsFindMetalGatewayById200Response is a convenience function that returns FindMetalGatewayById200ResponseOneOf wrapped in FindMetalGatewayById200Response
func FindMetalGatewayById200ResponseOneOfAsFindMetalGatewayById200Response(v *FindMetalGatewayById200ResponseOneOf) FindMetalGatewayById200Response {
	return FindMetalGatewayById200Response{
		FindMetalGatewayById200ResponseOneOf: v,
	}
}

// FindMetalGatewayById200ResponseOneOf1AsFindMetalGatewayById200Response is a convenience function that returns FindMetalGatewayById200ResponseOneOf1 wrapped in FindMetalGatewayById200Response
func FindMetalGatewayById200ResponseOneOf1AsFindMetalGatewayById200Response(v *FindMetalGatewayById200ResponseOneOf1) FindMetalGatewayById200Response {
	return FindMetalGatewayById200Response{
		FindMetalGatewayById200ResponseOneOf1: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *FindMetalGatewayById200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into FindMetalGatewayById200ResponseOneOf
	err = newStrictDecoder(data).Decode(&dst.FindMetalGatewayById200ResponseOneOf)
	if err == nil {
		jsonFindMetalGatewayById200ResponseOneOf, _ := json.Marshal(dst.FindMetalGatewayById200ResponseOneOf)
		if string(jsonFindMetalGatewayById200ResponseOneOf) == "{}" { // empty struct
			dst.FindMetalGatewayById200ResponseOneOf = nil
		} else {
			match++
		}
	} else {
		dst.FindMetalGatewayById200ResponseOneOf = nil
	}

	// try to unmarshal data into FindMetalGatewayById200ResponseOneOf1
	err = newStrictDecoder(data).Decode(&dst.FindMetalGatewayById200ResponseOneOf1)
	if err == nil {
		jsonFindMetalGatewayById200ResponseOneOf1, _ := json.Marshal(dst.FindMetalGatewayById200ResponseOneOf1)
		if string(jsonFindMetalGatewayById200ResponseOneOf1) == "{}" { // empty struct
			dst.FindMetalGatewayById200ResponseOneOf1 = nil
		} else {
			match++
		}
	} else {
		dst.FindMetalGatewayById200ResponseOneOf1 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.FindMetalGatewayById200ResponseOneOf = nil
		dst.FindMetalGatewayById200ResponseOneOf1 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(FindMetalGatewayById200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(FindMetalGatewayById200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src FindMetalGatewayById200Response) MarshalJSON() ([]byte, error) {
	if src.FindMetalGatewayById200ResponseOneOf != nil {
		return json.Marshal(&src.FindMetalGatewayById200ResponseOneOf)
	}

	if src.FindMetalGatewayById200ResponseOneOf1 != nil {
		return json.Marshal(&src.FindMetalGatewayById200ResponseOneOf1)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *FindMetalGatewayById200Response) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.FindMetalGatewayById200ResponseOneOf != nil {
		return obj.FindMetalGatewayById200ResponseOneOf
	}

	if obj.FindMetalGatewayById200ResponseOneOf1 != nil {
		return obj.FindMetalGatewayById200ResponseOneOf1
	}

	// all schemas are nil
	return nil
}

type NullableFindMetalGatewayById200Response struct {
	value *FindMetalGatewayById200Response
	isSet bool
}

func (v NullableFindMetalGatewayById200Response) Get() *FindMetalGatewayById200Response {
	return v.value
}

func (v *NullableFindMetalGatewayById200Response) Set(val *FindMetalGatewayById200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableFindMetalGatewayById200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableFindMetalGatewayById200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFindMetalGatewayById200Response(val *FindMetalGatewayById200Response) *NullableFindMetalGatewayById200Response {
	return &NullableFindMetalGatewayById200Response{value: val, isSet: true}
}

func (v NullableFindMetalGatewayById200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFindMetalGatewayById200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}