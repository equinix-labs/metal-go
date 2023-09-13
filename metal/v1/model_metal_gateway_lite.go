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

// checks if the MetalGatewayLite type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetalGatewayLite{}

// MetalGatewayLite struct for MetalGatewayLite
type MetalGatewayLite struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The gateway address with subnet CIDR value for this Metal Gateway. For example, a Metal Gateway using an IP reservation with block 10.1.2.0/27 would have a gateway address of 10.1.2.1/27.
	GatewayAddress *string `json:"gateway_address,omitempty"`
	Href           *string `json:"href,omitempty"`
	Id             *string `json:"id,omitempty"`
	// The current state of the Metal Gateway. 'Ready' indicates the gateway record has been configured, but is currently not active on the network. 'Active' indicates the gateway has been configured on the network. 'Deleting' is a temporary state used to indicate that the gateway is in the process of being un-configured from the network, after which the gateway record will be deleted.
	State     *string    `json:"state,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// The VLAN id of the Virtual Network record associated to this Metal Gateway.
	Vlan                 *int32 `json:"vlan,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MetalGatewayLite MetalGatewayLite

// NewMetalGatewayLite instantiates a new MetalGatewayLite object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetalGatewayLite() *MetalGatewayLite {
	this := MetalGatewayLite{}
	return &this
}

// NewMetalGatewayLiteWithDefaults instantiates a new MetalGatewayLite object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetalGatewayLiteWithDefaults() *MetalGatewayLite {
	this := MetalGatewayLite{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *MetalGatewayLite) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetalGatewayLite) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *MetalGatewayLite) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *MetalGatewayLite) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetGatewayAddress returns the GatewayAddress field value if set, zero value otherwise.
func (o *MetalGatewayLite) GetGatewayAddress() string {
	if o == nil || IsNil(o.GatewayAddress) {
		var ret string
		return ret
	}
	return *o.GatewayAddress
}

// GetGatewayAddressOk returns a tuple with the GatewayAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetalGatewayLite) GetGatewayAddressOk() (*string, bool) {
	if o == nil || IsNil(o.GatewayAddress) {
		return nil, false
	}
	return o.GatewayAddress, true
}

// HasGatewayAddress returns a boolean if a field has been set.
func (o *MetalGatewayLite) HasGatewayAddress() bool {
	if o != nil && !IsNil(o.GatewayAddress) {
		return true
	}

	return false
}

// SetGatewayAddress gets a reference to the given string and assigns it to the GatewayAddress field.
func (o *MetalGatewayLite) SetGatewayAddress(v string) {
	o.GatewayAddress = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *MetalGatewayLite) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetalGatewayLite) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *MetalGatewayLite) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *MetalGatewayLite) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MetalGatewayLite) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetalGatewayLite) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MetalGatewayLite) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MetalGatewayLite) SetId(v string) {
	o.Id = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *MetalGatewayLite) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetalGatewayLite) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *MetalGatewayLite) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *MetalGatewayLite) SetState(v string) {
	o.State = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *MetalGatewayLite) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetalGatewayLite) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *MetalGatewayLite) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *MetalGatewayLite) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetVlan returns the Vlan field value if set, zero value otherwise.
func (o *MetalGatewayLite) GetVlan() int32 {
	if o == nil || IsNil(o.Vlan) {
		var ret int32
		return ret
	}
	return *o.Vlan
}

// GetVlanOk returns a tuple with the Vlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetalGatewayLite) GetVlanOk() (*int32, bool) {
	if o == nil || IsNil(o.Vlan) {
		return nil, false
	}
	return o.Vlan, true
}

// HasVlan returns a boolean if a field has been set.
func (o *MetalGatewayLite) HasVlan() bool {
	if o != nil && !IsNil(o.Vlan) {
		return true
	}

	return false
}

// SetVlan gets a reference to the given int32 and assigns it to the Vlan field.
func (o *MetalGatewayLite) SetVlan(v int32) {
	o.Vlan = &v
}

func (o MetalGatewayLite) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetalGatewayLite) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.GatewayAddress) {
		toSerialize["gateway_address"] = o.GatewayAddress
	}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.Vlan) {
		toSerialize["vlan"] = o.Vlan
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MetalGatewayLite) UnmarshalJSON(bytes []byte) (err error) {
	varMetalGatewayLite := _MetalGatewayLite{}

	err = json.Unmarshal(bytes, &varMetalGatewayLite)

	if err != nil {
		return err
	}

	*o = MetalGatewayLite(varMetalGatewayLite)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "gateway_address")
		delete(additionalProperties, "href")
		delete(additionalProperties, "id")
		delete(additionalProperties, "state")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "vlan")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMetalGatewayLite struct {
	value *MetalGatewayLite
	isSet bool
}

func (v NullableMetalGatewayLite) Get() *MetalGatewayLite {
	return v.value
}

func (v *NullableMetalGatewayLite) Set(val *MetalGatewayLite) {
	v.value = val
	v.isSet = true
}

func (v NullableMetalGatewayLite) IsSet() bool {
	return v.isSet
}

func (v *NullableMetalGatewayLite) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetalGatewayLite(val *MetalGatewayLite) *NullableMetalGatewayLite {
	return &NullableMetalGatewayLite{value: val, isSet: true}
}

func (v NullableMetalGatewayLite) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetalGatewayLite) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
