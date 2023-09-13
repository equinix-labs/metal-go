/*
Metal API

# Introduction Equinix Metal provides a RESTful HTTP API which can be reached at <https://api.equinix.com/metal/v1>. This document describes the API and how to use it.  The API allows you to programmatically interact with all of your Equinix Metal resources, including devices, networks, addresses, organizations, projects, and your user account. Every feature of the Equinix Metal web interface is accessible through the API.  The API docs are generated from the Equinix Metal OpenAPI specification and are officially hosted at <https://metal.equinix.com/developers/api>.  # Common Parameters  The Equinix Metal API uses a few methods to minimize network traffic and improve throughput. These parameters are not used in all API calls, but are used often enough to warrant their own section. Look for these parameters in the documentation for the API calls that support them.  ## Pagination  Pagination is used to limit the number of results returned in a single request. The API will return a maximum of 100 results per page. To retrieve additional results, you can use the `page` and `per_page` query parameters.  The `page` parameter is used to specify the page number. The first page is `1`. The `per_page` parameter is used to specify the number of results per page. The maximum number of results differs by resource type.  ## Sorting  Where offered, the API allows you to sort results by a specific field. To sort results use the `sort_by` query parameter with the root level field name as the value. The `sort_direction` parameter is used to specify the sort direction, either either `asc` (ascending) or `desc` (descending).  ## Filtering  Filtering is used to limit the results returned in a single request. The API supports filtering by certain fields in the response. To filter results, you can use the field as a query parameter.  For example, to filter the IP list to only return public IPv4 addresses, you can filter by the `type` field, as in the following request:  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/projects/id/ips?type=public_ipv4 ```  Only IP addresses with the `type` field set to `public_ipv4` will be returned.  ## Searching  Searching is used to find matching resources using multiple field comparissons. The API supports searching in resources that define this behavior. Currently the search parameter is only available on devices, ssh_keys, api_keys and memberships endpoints.  To search resources you can use the `search` query parameter.  ## Include and Exclude  For resources that contain references to other resources, sucha as a Device that refers to the Project it resides in, the Equinix Metal API will returns `href` values (API links) to the associated resource.  ```json {   ...   \"project\": {     \"href\": \"/metal/v1/projects/f3f131c8-f302-49ef-8c44-9405022dc6dd\"   } } ```  If you're going need the project details, you can avoid a second API request.  Specify the contained `href` resources and collections that you'd like to have included in the response using the `include` query parameter.  For example:  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=projects ```  The `include` parameter is generally accepted in `GET`, `POST`, `PUT`, and `PATCH` requests where `href` resources are presented.  To have multiple resources include, use a comma-separated list (e.g. `?include=emails,projects,memberships`).  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=emails,projects,memberships ```  You may also include nested associations up to three levels deep using dot notation (`?include=memberships.projects`):  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=memberships.projects ```  To exclude resources, and optimize response delivery, use the `exclude` query parameter. The `exclude` parameter is generally accepted in `GET`, `POST`, `PUT`, and `PATCH` requests for fields with nested object responses. When excluded, these fields will be replaced with an object that contains only an `href` field.

API version: 1.0.0
Contact: support@equinixmetal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
	"time"
)

// checks if the PortVlanAssignment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PortVlanAssignment{}

// PortVlanAssignment struct for PortVlanAssignment
type PortVlanAssignment struct {
	CreatedAt            *time.Time `json:"created_at,omitempty"`
	Id                   *string    `json:"id,omitempty"`
	Native               *bool      `json:"native,omitempty"`
	Port                 *Href      `json:"port,omitempty"`
	State                *string    `json:"state,omitempty"`
	UpdatedAt            *time.Time `json:"updated_at,omitempty"`
	VirtualNetwork       *Href      `json:"virtual_network,omitempty"`
	Vlan                 *int32     `json:"vlan,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PortVlanAssignment PortVlanAssignment

// NewPortVlanAssignment instantiates a new PortVlanAssignment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortVlanAssignment() *PortVlanAssignment {
	this := PortVlanAssignment{}
	return &this
}

// NewPortVlanAssignmentWithDefaults instantiates a new PortVlanAssignment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortVlanAssignmentWithDefaults() *PortVlanAssignment {
	this := PortVlanAssignment{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PortVlanAssignment) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortVlanAssignment) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PortVlanAssignment) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *PortVlanAssignment) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PortVlanAssignment) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortVlanAssignment) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PortVlanAssignment) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PortVlanAssignment) SetId(v string) {
	o.Id = &v
}

// GetNative returns the Native field value if set, zero value otherwise.
func (o *PortVlanAssignment) GetNative() bool {
	if o == nil || IsNil(o.Native) {
		var ret bool
		return ret
	}
	return *o.Native
}

// GetNativeOk returns a tuple with the Native field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortVlanAssignment) GetNativeOk() (*bool, bool) {
	if o == nil || IsNil(o.Native) {
		return nil, false
	}
	return o.Native, true
}

// HasNative returns a boolean if a field has been set.
func (o *PortVlanAssignment) HasNative() bool {
	if o != nil && !IsNil(o.Native) {
		return true
	}

	return false
}

// SetNative gets a reference to the given bool and assigns it to the Native field.
func (o *PortVlanAssignment) SetNative(v bool) {
	o.Native = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *PortVlanAssignment) GetPort() Href {
	if o == nil || IsNil(o.Port) {
		var ret Href
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortVlanAssignment) GetPortOk() (*Href, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *PortVlanAssignment) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given Href and assigns it to the Port field.
func (o *PortVlanAssignment) SetPort(v Href) {
	o.Port = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *PortVlanAssignment) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortVlanAssignment) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *PortVlanAssignment) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *PortVlanAssignment) SetState(v string) {
	o.State = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *PortVlanAssignment) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortVlanAssignment) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *PortVlanAssignment) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *PortVlanAssignment) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetVirtualNetwork returns the VirtualNetwork field value if set, zero value otherwise.
func (o *PortVlanAssignment) GetVirtualNetwork() Href {
	if o == nil || IsNil(o.VirtualNetwork) {
		var ret Href
		return ret
	}
	return *o.VirtualNetwork
}

// GetVirtualNetworkOk returns a tuple with the VirtualNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortVlanAssignment) GetVirtualNetworkOk() (*Href, bool) {
	if o == nil || IsNil(o.VirtualNetwork) {
		return nil, false
	}
	return o.VirtualNetwork, true
}

// HasVirtualNetwork returns a boolean if a field has been set.
func (o *PortVlanAssignment) HasVirtualNetwork() bool {
	if o != nil && !IsNil(o.VirtualNetwork) {
		return true
	}

	return false
}

// SetVirtualNetwork gets a reference to the given Href and assigns it to the VirtualNetwork field.
func (o *PortVlanAssignment) SetVirtualNetwork(v Href) {
	o.VirtualNetwork = &v
}

// GetVlan returns the Vlan field value if set, zero value otherwise.
func (o *PortVlanAssignment) GetVlan() int32 {
	if o == nil || IsNil(o.Vlan) {
		var ret int32
		return ret
	}
	return *o.Vlan
}

// GetVlanOk returns a tuple with the Vlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortVlanAssignment) GetVlanOk() (*int32, bool) {
	if o == nil || IsNil(o.Vlan) {
		return nil, false
	}
	return o.Vlan, true
}

// HasVlan returns a boolean if a field has been set.
func (o *PortVlanAssignment) HasVlan() bool {
	if o != nil && !IsNil(o.Vlan) {
		return true
	}

	return false
}

// SetVlan gets a reference to the given int32 and assigns it to the Vlan field.
func (o *PortVlanAssignment) SetVlan(v int32) {
	o.Vlan = &v
}

func (o PortVlanAssignment) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PortVlanAssignment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Native) {
		toSerialize["native"] = o.Native
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.VirtualNetwork) {
		toSerialize["virtual_network"] = o.VirtualNetwork
	}
	if !IsNil(o.Vlan) {
		toSerialize["vlan"] = o.Vlan
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PortVlanAssignment) UnmarshalJSON(bytes []byte) (err error) {
	varPortVlanAssignment := _PortVlanAssignment{}

	err = json.Unmarshal(bytes, &varPortVlanAssignment)

	if err != nil {
		return err
	}

	*o = PortVlanAssignment(varPortVlanAssignment)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "id")
		delete(additionalProperties, "native")
		delete(additionalProperties, "port")
		delete(additionalProperties, "state")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "virtual_network")
		delete(additionalProperties, "vlan")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePortVlanAssignment struct {
	value *PortVlanAssignment
	isSet bool
}

func (v NullablePortVlanAssignment) Get() *PortVlanAssignment {
	return v.value
}

func (v *NullablePortVlanAssignment) Set(val *PortVlanAssignment) {
	v.value = val
	v.isSet = true
}

func (v NullablePortVlanAssignment) IsSet() bool {
	return v.isSet
}

func (v *NullablePortVlanAssignment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortVlanAssignment(val *PortVlanAssignment) *NullablePortVlanAssignment {
	return &NullablePortVlanAssignment{value: val, isSet: true}
}

func (v NullablePortVlanAssignment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortVlanAssignment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
