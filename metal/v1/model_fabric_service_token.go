/*
Metal API

# Introduction Equinix Metal provides a RESTful HTTP API which can be reached at <https://api.equinix.com/metal/v1>. This document describes the API and how to use it.  The API allows you to programmatically interact with all of your Equinix Metal resources, including devices, networks, addresses, organizations, projects, and your user account. Every feature of the Equinix Metal web interface is accessible through the API.  The API docs are generated from the Equinix Metal OpenAPI specification and are officially hosted at <https://metal.equinix.com/developers/api>.  # Common Parameters  The Equinix Metal API uses a few methods to minimize network traffic and improve throughput. These parameters are not used in all API calls, but are used often enough to warrant their own section. Look for these parameters in the documentation for the API calls that support them.  ## Pagination  Pagination is used to limit the number of results returned in a single request. The API will return a maximum of 100 results per page. To retrieve additional results, you can use the `page` and `per_page` query parameters.  The `page` parameter is used to specify the page number. The first page is `1`. The `per_page` parameter is used to specify the number of results per page. The maximum number of results differs by resource type.  ## Sorting  Where offered, the API allows you to sort results by a specific field. To sort results use the `sort_by` query parameter with the root level field name as the value. The `sort_direction` parameter is used to specify the sort direction, either either `asc` (ascending) or `desc` (descending).  ## Filtering  Filtering is used to limit the results returned in a single request. The API supports filtering by certain fields in the response. To filter results, you can use the field as a query parameter.  For example, to filter the IP list to only return public IPv4 addresses, you can filter by the `type` field, as in the following request:  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/projects/id/ips?type=public_ipv4 ```  Only IP addresses with the `type` field set to `public_ipv4` will be returned.  ## Searching  Searching is used to find matching resources using multiple field comparisons. The API supports searching in resources that define this behavior. Currently the search parameter is only available on devices, ssh_keys, api_keys and memberships endpoints.  To search resources you can use the `search` query parameter.  ## Include and Exclude  For resources that contain references to other resources, such as a Device that refers to the Project it resides in, the Equinix Metal API will returns `href` values (API links) to the associated resource.  ```json {   ...   \"project\": {     \"href\": \"/metal/v1/projects/f3f131c8-f302-49ef-8c44-9405022dc6dd\"   } } ```  If you're going need the project details, you can avoid a second API request.  Specify the contained `href` resources and collections that you'd like to have included in the response using the `include` query parameter.  For example:  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=projects ```  The `include` parameter is generally accepted in `GET`, `POST`, `PUT`, and `PATCH` requests where `href` resources are presented.  To have multiple resources include, use a comma-separated list (e.g. `?include=emails,projects,memberships`).  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=emails,projects,memberships ```  You may also include nested associations up to three levels deep using dot notation (`?include=memberships.projects`):  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=memberships.projects ```  To exclude resources, and optimize response delivery, use the `exclude` query parameter. The `exclude` parameter is generally accepted in `GET`, `POST`, `PUT`, and `PATCH` requests for fields with nested object responses. When excluded, these fields will be replaced with an object that contains only an `href` field.

API version: 1.0.0
Contact: support@equinixmetal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
	"time"
)

// checks if the FabricServiceToken type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FabricServiceToken{}

// FabricServiceToken struct for FabricServiceToken
type FabricServiceToken struct {
	// The expiration date and time of the Fabric service token. Once a service token is expired, it is no longer redeemable.
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
	// The UUID that can be used on the Fabric Portal to redeem either an A-Side or Z-Side Service Token. For Fabric VCs (Metal Billed), this UUID will represent an A-Side Service Token, which will allow interconnections to be made from Equinix Metal to other Service Providers on Fabric. For Fabric VCs (Fabric Billed), this UUID will represent a Z-Side Service Token, which will allow interconnections to be made to connect an owned Fabric Port or  Virtual Device to Equinix Metal.
	Id *string `json:"id,omitempty"`
	// The maximum speed that can be selected on the Fabric Portal when configuring a interconnection with either  an A-Side or Z-Side Service Token. For Fabric VCs (Metal Billed), this is what the billing is based off of, and can be one of the following options, '50mbps', '200mbps', '500mbps', '1gbps', '2gbps', '5gbps' or '10gbps'. For Fabric VCs (Fabric Billed), this will default to 10Gbps.
	MaxAllowedSpeed      *int32                              `json:"max_allowed_speed,omitempty"`
	Role                 *FabricServiceTokenRole             `json:"role,omitempty"`
	ServiceTokenType     *FabricServiceTokenServiceTokenType `json:"service_token_type,omitempty"`
	State                *FabricServiceTokenState            `json:"state,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricServiceToken FabricServiceToken

// NewFabricServiceToken instantiates a new FabricServiceToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricServiceToken() *FabricServiceToken {
	this := FabricServiceToken{}
	return &this
}

// NewFabricServiceTokenWithDefaults instantiates a new FabricServiceToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricServiceTokenWithDefaults() *FabricServiceToken {
	this := FabricServiceToken{}
	return &this
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *FabricServiceToken) GetExpiresAt() time.Time {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricServiceToken) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *FabricServiceToken) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *FabricServiceToken) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FabricServiceToken) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricServiceToken) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FabricServiceToken) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *FabricServiceToken) SetId(v string) {
	o.Id = &v
}

// GetMaxAllowedSpeed returns the MaxAllowedSpeed field value if set, zero value otherwise.
func (o *FabricServiceToken) GetMaxAllowedSpeed() int32 {
	if o == nil || IsNil(o.MaxAllowedSpeed) {
		var ret int32
		return ret
	}
	return *o.MaxAllowedSpeed
}

// GetMaxAllowedSpeedOk returns a tuple with the MaxAllowedSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricServiceToken) GetMaxAllowedSpeedOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxAllowedSpeed) {
		return nil, false
	}
	return o.MaxAllowedSpeed, true
}

// HasMaxAllowedSpeed returns a boolean if a field has been set.
func (o *FabricServiceToken) HasMaxAllowedSpeed() bool {
	if o != nil && !IsNil(o.MaxAllowedSpeed) {
		return true
	}

	return false
}

// SetMaxAllowedSpeed gets a reference to the given int32 and assigns it to the MaxAllowedSpeed field.
func (o *FabricServiceToken) SetMaxAllowedSpeed(v int32) {
	o.MaxAllowedSpeed = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *FabricServiceToken) GetRole() FabricServiceTokenRole {
	if o == nil || IsNil(o.Role) {
		var ret FabricServiceTokenRole
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricServiceToken) GetRoleOk() (*FabricServiceTokenRole, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *FabricServiceToken) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given FabricServiceTokenRole and assigns it to the Role field.
func (o *FabricServiceToken) SetRole(v FabricServiceTokenRole) {
	o.Role = &v
}

// GetServiceTokenType returns the ServiceTokenType field value if set, zero value otherwise.
func (o *FabricServiceToken) GetServiceTokenType() FabricServiceTokenServiceTokenType {
	if o == nil || IsNil(o.ServiceTokenType) {
		var ret FabricServiceTokenServiceTokenType
		return ret
	}
	return *o.ServiceTokenType
}

// GetServiceTokenTypeOk returns a tuple with the ServiceTokenType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricServiceToken) GetServiceTokenTypeOk() (*FabricServiceTokenServiceTokenType, bool) {
	if o == nil || IsNil(o.ServiceTokenType) {
		return nil, false
	}
	return o.ServiceTokenType, true
}

// HasServiceTokenType returns a boolean if a field has been set.
func (o *FabricServiceToken) HasServiceTokenType() bool {
	if o != nil && !IsNil(o.ServiceTokenType) {
		return true
	}

	return false
}

// SetServiceTokenType gets a reference to the given FabricServiceTokenServiceTokenType and assigns it to the ServiceTokenType field.
func (o *FabricServiceToken) SetServiceTokenType(v FabricServiceTokenServiceTokenType) {
	o.ServiceTokenType = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *FabricServiceToken) GetState() FabricServiceTokenState {
	if o == nil || IsNil(o.State) {
		var ret FabricServiceTokenState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricServiceToken) GetStateOk() (*FabricServiceTokenState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *FabricServiceToken) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given FabricServiceTokenState and assigns it to the State field.
func (o *FabricServiceToken) SetState(v FabricServiceTokenState) {
	o.State = &v
}

func (o FabricServiceToken) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FabricServiceToken) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExpiresAt) {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.MaxAllowedSpeed) {
		toSerialize["max_allowed_speed"] = o.MaxAllowedSpeed
	}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	if !IsNil(o.ServiceTokenType) {
		toSerialize["service_token_type"] = o.ServiceTokenType
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FabricServiceToken) UnmarshalJSON(bytes []byte) (err error) {
	varFabricServiceToken := _FabricServiceToken{}

	err = json.Unmarshal(bytes, &varFabricServiceToken)

	if err != nil {
		return err
	}

	*o = FabricServiceToken(varFabricServiceToken)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "expires_at")
		delete(additionalProperties, "id")
		delete(additionalProperties, "max_allowed_speed")
		delete(additionalProperties, "role")
		delete(additionalProperties, "service_token_type")
		delete(additionalProperties, "state")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFabricServiceToken struct {
	value *FabricServiceToken
	isSet bool
}

func (v NullableFabricServiceToken) Get() *FabricServiceToken {
	return v.value
}

func (v *NullableFabricServiceToken) Set(val *FabricServiceToken) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricServiceToken) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricServiceToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricServiceToken(val *FabricServiceToken) *NullableFabricServiceToken {
	return &NullableFabricServiceToken{value: val, isSet: true}
}

func (v NullableFabricServiceToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricServiceToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
