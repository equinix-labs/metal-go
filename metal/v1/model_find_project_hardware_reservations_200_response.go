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

// FindProjectHardwareReservations200Response struct for FindProjectHardwareReservations200Response
type FindProjectHardwareReservations200Response struct {
	HardwareReservations []MoveHardwareReservation201Response `json:"hardware_reservations,omitempty"`
	Meta                 *FindDeviceEvents200ResponseMeta     `json:"meta,omitempty"`
}

// NewFindProjectHardwareReservations200Response instantiates a new FindProjectHardwareReservations200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFindProjectHardwareReservations200Response() *FindProjectHardwareReservations200Response {
	this := FindProjectHardwareReservations200Response{}
	return &this
}

// NewFindProjectHardwareReservations200ResponseWithDefaults instantiates a new FindProjectHardwareReservations200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFindProjectHardwareReservations200ResponseWithDefaults() *FindProjectHardwareReservations200Response {
	this := FindProjectHardwareReservations200Response{}
	return &this
}

// GetHardwareReservations returns the HardwareReservations field value if set, zero value otherwise.
func (o *FindProjectHardwareReservations200Response) GetHardwareReservations() []MoveHardwareReservation201Response {
	if o == nil || o.HardwareReservations == nil {
		var ret []MoveHardwareReservation201Response
		return ret
	}
	return o.HardwareReservations
}

// GetHardwareReservationsOk returns a tuple with the HardwareReservations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindProjectHardwareReservations200Response) GetHardwareReservationsOk() ([]MoveHardwareReservation201Response, bool) {
	if o == nil || o.HardwareReservations == nil {
		return nil, false
	}
	return o.HardwareReservations, true
}

// HasHardwareReservations returns a boolean if a field has been set.
func (o *FindProjectHardwareReservations200Response) HasHardwareReservations() bool {
	if o != nil && o.HardwareReservations != nil {
		return true
	}

	return false
}

// SetHardwareReservations gets a reference to the given []MoveHardwareReservation201Response and assigns it to the HardwareReservations field.
func (o *FindProjectHardwareReservations200Response) SetHardwareReservations(v []MoveHardwareReservation201Response) {
	o.HardwareReservations = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *FindProjectHardwareReservations200Response) GetMeta() FindDeviceEvents200ResponseMeta {
	if o == nil || o.Meta == nil {
		var ret FindDeviceEvents200ResponseMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindProjectHardwareReservations200Response) GetMetaOk() (*FindDeviceEvents200ResponseMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *FindProjectHardwareReservations200Response) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given FindDeviceEvents200ResponseMeta and assigns it to the Meta field.
func (o *FindProjectHardwareReservations200Response) SetMeta(v FindDeviceEvents200ResponseMeta) {
	o.Meta = &v
}

func (o FindProjectHardwareReservations200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.HardwareReservations != nil {
		toSerialize["hardware_reservations"] = o.HardwareReservations
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableFindProjectHardwareReservations200Response struct {
	value *FindProjectHardwareReservations200Response
	isSet bool
}

func (v NullableFindProjectHardwareReservations200Response) Get() *FindProjectHardwareReservations200Response {
	return v.value
}

func (v *NullableFindProjectHardwareReservations200Response) Set(val *FindProjectHardwareReservations200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableFindProjectHardwareReservations200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableFindProjectHardwareReservations200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFindProjectHardwareReservations200Response(val *FindProjectHardwareReservations200Response) *NullableFindProjectHardwareReservations200Response {
	return &NullableFindProjectHardwareReservations200Response{value: val, isSet: true}
}

func (v NullableFindProjectHardwareReservations200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFindProjectHardwareReservations200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}