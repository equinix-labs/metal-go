/*
Metal API

This is the API for Equinix Metal. The API allows you to programmatically interact with all of your Equinix Metal resources, including devices, networks, addresses, organizations, projects, and your user account.  The official API docs are hosted at <https://metal.equinix.com/developers/api>.

API version: 1.0.0
Contact: support@equinixmetal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
)

// GetBgpNeighborData200ResponseBgpNeighborsInnerRoutesOutInner struct for GetBgpNeighborData200ResponseBgpNeighborsInnerRoutesOutInner
type GetBgpNeighborData200ResponseBgpNeighborsInnerRoutesOutInner struct {
	Exact *bool   `json:"exact,omitempty"`
	Route *string `json:"route,omitempty"`
}

// NewGetBgpNeighborData200ResponseBgpNeighborsInnerRoutesOutInner instantiates a new GetBgpNeighborData200ResponseBgpNeighborsInnerRoutesOutInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBgpNeighborData200ResponseBgpNeighborsInnerRoutesOutInner() *GetBgpNeighborData200ResponseBgpNeighborsInnerRoutesOutInner {
	this := GetBgpNeighborData200ResponseBgpNeighborsInnerRoutesOutInner{}
	return &this
}

// NewGetBgpNeighborData200ResponseBgpNeighborsInnerRoutesOutInnerWithDefaults instantiates a new GetBgpNeighborData200ResponseBgpNeighborsInnerRoutesOutInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBgpNeighborData200ResponseBgpNeighborsInnerRoutesOutInnerWithDefaults() *GetBgpNeighborData200ResponseBgpNeighborsInnerRoutesOutInner {
	this := GetBgpNeighborData200ResponseBgpNeighborsInnerRoutesOutInner{}
	return &this
}

// GetExact returns the Exact field value if set, zero value otherwise.
func (o *GetBgpNeighborData200ResponseBgpNeighborsInnerRoutesOutInner) GetExact() bool {
	if o == nil || o.Exact == nil {
		var ret bool
		return ret
	}
	return *o.Exact
}

// GetExactOk returns a tuple with the Exact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBgpNeighborData200ResponseBgpNeighborsInnerRoutesOutInner) GetExactOk() (*bool, bool) {
	if o == nil || o.Exact == nil {
		return nil, false
	}
	return o.Exact, true
}

// HasExact returns a boolean if a field has been set.
func (o *GetBgpNeighborData200ResponseBgpNeighborsInnerRoutesOutInner) HasExact() bool {
	if o != nil && o.Exact != nil {
		return true
	}

	return false
}

// SetExact gets a reference to the given bool and assigns it to the Exact field.
func (o *GetBgpNeighborData200ResponseBgpNeighborsInnerRoutesOutInner) SetExact(v bool) {
	o.Exact = &v
}

// GetRoute returns the Route field value if set, zero value otherwise.
func (o *GetBgpNeighborData200ResponseBgpNeighborsInnerRoutesOutInner) GetRoute() string {
	if o == nil || o.Route == nil {
		var ret string
		return ret
	}
	return *o.Route
}

// GetRouteOk returns a tuple with the Route field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBgpNeighborData200ResponseBgpNeighborsInnerRoutesOutInner) GetRouteOk() (*string, bool) {
	if o == nil || o.Route == nil {
		return nil, false
	}
	return o.Route, true
}

// HasRoute returns a boolean if a field has been set.
func (o *GetBgpNeighborData200ResponseBgpNeighborsInnerRoutesOutInner) HasRoute() bool {
	if o != nil && o.Route != nil {
		return true
	}

	return false
}

// SetRoute gets a reference to the given string and assigns it to the Route field.
func (o *GetBgpNeighborData200ResponseBgpNeighborsInnerRoutesOutInner) SetRoute(v string) {
	o.Route = &v
}

func (o GetBgpNeighborData200ResponseBgpNeighborsInnerRoutesOutInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Exact != nil {
		toSerialize["exact"] = o.Exact
	}
	if o.Route != nil {
		toSerialize["route"] = o.Route
	}
	return json.Marshal(toSerialize)
}

type NullableGetBgpNeighborData200ResponseBgpNeighborsInnerRoutesOutInner struct {
	value *GetBgpNeighborData200ResponseBgpNeighborsInnerRoutesOutInner
	isSet bool
}

func (v NullableGetBgpNeighborData200ResponseBgpNeighborsInnerRoutesOutInner) Get() *GetBgpNeighborData200ResponseBgpNeighborsInnerRoutesOutInner {
	return v.value
}

func (v *NullableGetBgpNeighborData200ResponseBgpNeighborsInnerRoutesOutInner) Set(val *GetBgpNeighborData200ResponseBgpNeighborsInnerRoutesOutInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBgpNeighborData200ResponseBgpNeighborsInnerRoutesOutInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBgpNeighborData200ResponseBgpNeighborsInnerRoutesOutInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBgpNeighborData200ResponseBgpNeighborsInnerRoutesOutInner(val *GetBgpNeighborData200ResponseBgpNeighborsInnerRoutesOutInner) *NullableGetBgpNeighborData200ResponseBgpNeighborsInnerRoutesOutInner {
	return &NullableGetBgpNeighborData200ResponseBgpNeighborsInnerRoutesOutInner{value: val, isSet: true}
}

func (v NullableGetBgpNeighborData200ResponseBgpNeighborsInnerRoutesOutInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBgpNeighborData200ResponseBgpNeighborsInnerRoutesOutInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}