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

// VirtualNetworkList struct for VirtualNetworkList
type VirtualNetworkList struct {
	VirtualNetworks []FindVirtualNetworks200ResponseVirtualNetworksInner `json:"virtual_networks,omitempty"`
}

// NewVirtualNetworkList instantiates a new VirtualNetworkList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualNetworkList() *VirtualNetworkList {
	this := VirtualNetworkList{}
	return &this
}

// NewVirtualNetworkListWithDefaults instantiates a new VirtualNetworkList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualNetworkListWithDefaults() *VirtualNetworkList {
	this := VirtualNetworkList{}
	return &this
}

// GetVirtualNetworks returns the VirtualNetworks field value if set, zero value otherwise.
func (o *VirtualNetworkList) GetVirtualNetworks() []FindVirtualNetworks200ResponseVirtualNetworksInner {
	if o == nil || o.VirtualNetworks == nil {
		var ret []FindVirtualNetworks200ResponseVirtualNetworksInner
		return ret
	}
	return o.VirtualNetworks
}

// GetVirtualNetworksOk returns a tuple with the VirtualNetworks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualNetworkList) GetVirtualNetworksOk() ([]FindVirtualNetworks200ResponseVirtualNetworksInner, bool) {
	if o == nil || o.VirtualNetworks == nil {
		return nil, false
	}
	return o.VirtualNetworks, true
}

// HasVirtualNetworks returns a boolean if a field has been set.
func (o *VirtualNetworkList) HasVirtualNetworks() bool {
	if o != nil && o.VirtualNetworks != nil {
		return true
	}

	return false
}

// SetVirtualNetworks gets a reference to the given []FindVirtualNetworks200ResponseVirtualNetworksInner and assigns it to the VirtualNetworks field.
func (o *VirtualNetworkList) SetVirtualNetworks(v []FindVirtualNetworks200ResponseVirtualNetworksInner) {
	o.VirtualNetworks = v
}

func (o VirtualNetworkList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.VirtualNetworks != nil {
		toSerialize["virtual_networks"] = o.VirtualNetworks
	}
	return json.Marshal(toSerialize)
}

type NullableVirtualNetworkList struct {
	value *VirtualNetworkList
	isSet bool
}

func (v NullableVirtualNetworkList) Get() *VirtualNetworkList {
	return v.value
}

func (v *NullableVirtualNetworkList) Set(val *VirtualNetworkList) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualNetworkList) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualNetworkList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualNetworkList(val *VirtualNetworkList) *NullableVirtualNetworkList {
	return &NullableVirtualNetworkList{value: val, isSet: true}
}

func (v NullableVirtualNetworkList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualNetworkList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}