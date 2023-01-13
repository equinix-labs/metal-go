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
)

// InterconnectionPort struct for InterconnectionPort
type InterconnectionPort struct {
	Id           *string `json:"id,omitempty"`
	Organization *Href   `json:"organization,omitempty"`
	// Either 'primary' or 'secondary'.
	Role *string `json:"role,omitempty"`
	// For both Fabric VCs and Dedicated Ports, this will be 'requested' on creation and 'deleting' on deletion. Once the Fabric VC has found its corresponding Fabric connection, this will turn to 'active'. For Dedicated Ports, once the dedicated port is associated, this will also turn to 'active'. For Fabric VCs, this can turn into an 'expired' state if the service token associated is expired.
	Status *string `json:"status,omitempty"`
	// A switch 'short ID'
	SwitchId        *string             `json:"switch_id,omitempty"`
	VirtualCircuits *VirtualCircuitList `json:"virtual_circuits,omitempty"`
}

// NewInterconnectionPort instantiates a new InterconnectionPort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInterconnectionPort() *InterconnectionPort {
	this := InterconnectionPort{}
	return &this
}

// NewInterconnectionPortWithDefaults instantiates a new InterconnectionPort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInterconnectionPortWithDefaults() *InterconnectionPort {
	this := InterconnectionPort{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InterconnectionPort) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterconnectionPort) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InterconnectionPort) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InterconnectionPort) SetId(v string) {
	o.Id = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *InterconnectionPort) GetOrganization() Href {
	if o == nil || isNil(o.Organization) {
		var ret Href
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterconnectionPort) GetOrganizationOk() (*Href, bool) {
	if o == nil || isNil(o.Organization) {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *InterconnectionPort) HasOrganization() bool {
	if o != nil && !isNil(o.Organization) {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given Href and assigns it to the Organization field.
func (o *InterconnectionPort) SetOrganization(v Href) {
	o.Organization = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *InterconnectionPort) GetRole() string {
	if o == nil || isNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterconnectionPort) GetRoleOk() (*string, bool) {
	if o == nil || isNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *InterconnectionPort) HasRole() bool {
	if o != nil && !isNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *InterconnectionPort) SetRole(v string) {
	o.Role = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InterconnectionPort) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterconnectionPort) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InterconnectionPort) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *InterconnectionPort) SetStatus(v string) {
	o.Status = &v
}

// GetSwitchId returns the SwitchId field value if set, zero value otherwise.
func (o *InterconnectionPort) GetSwitchId() string {
	if o == nil || isNil(o.SwitchId) {
		var ret string
		return ret
	}
	return *o.SwitchId
}

// GetSwitchIdOk returns a tuple with the SwitchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterconnectionPort) GetSwitchIdOk() (*string, bool) {
	if o == nil || isNil(o.SwitchId) {
		return nil, false
	}
	return o.SwitchId, true
}

// HasSwitchId returns a boolean if a field has been set.
func (o *InterconnectionPort) HasSwitchId() bool {
	if o != nil && !isNil(o.SwitchId) {
		return true
	}

	return false
}

// SetSwitchId gets a reference to the given string and assigns it to the SwitchId field.
func (o *InterconnectionPort) SetSwitchId(v string) {
	o.SwitchId = &v
}

// GetVirtualCircuits returns the VirtualCircuits field value if set, zero value otherwise.
func (o *InterconnectionPort) GetVirtualCircuits() VirtualCircuitList {
	if o == nil || isNil(o.VirtualCircuits) {
		var ret VirtualCircuitList
		return ret
	}
	return *o.VirtualCircuits
}

// GetVirtualCircuitsOk returns a tuple with the VirtualCircuits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterconnectionPort) GetVirtualCircuitsOk() (*VirtualCircuitList, bool) {
	if o == nil || isNil(o.VirtualCircuits) {
		return nil, false
	}
	return o.VirtualCircuits, true
}

// HasVirtualCircuits returns a boolean if a field has been set.
func (o *InterconnectionPort) HasVirtualCircuits() bool {
	if o != nil && !isNil(o.VirtualCircuits) {
		return true
	}

	return false
}

// SetVirtualCircuits gets a reference to the given VirtualCircuitList and assigns it to the VirtualCircuits field.
func (o *InterconnectionPort) SetVirtualCircuits(v VirtualCircuitList) {
	o.VirtualCircuits = &v
}

func (o InterconnectionPort) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Organization) {
		toSerialize["organization"] = o.Organization
	}
	if !isNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.SwitchId) {
		toSerialize["switch_id"] = o.SwitchId
	}
	if !isNil(o.VirtualCircuits) {
		toSerialize["virtual_circuits"] = o.VirtualCircuits
	}
	return json.Marshal(toSerialize)
}

type NullableInterconnectionPort struct {
	value *InterconnectionPort
	isSet bool
}

func (v NullableInterconnectionPort) Get() *InterconnectionPort {
	return v.value
}

func (v *NullableInterconnectionPort) Set(val *InterconnectionPort) {
	v.value = val
	v.isSet = true
}

func (v NullableInterconnectionPort) IsSet() bool {
	return v.isSet
}

func (v *NullableInterconnectionPort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterconnectionPort(val *InterconnectionPort) *NullableInterconnectionPort {
	return &NullableInterconnectionPort{value: val, isSet: true}
}

func (v NullableInterconnectionPort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterconnectionPort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
