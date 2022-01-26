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

// SelfServiceReservationList struct for SelfServiceReservationList
type SelfServiceReservationList struct {
	Reservations []SelfServiceReservationResponse `json:"reservations,omitempty"`
}

// NewSelfServiceReservationList instantiates a new SelfServiceReservationList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSelfServiceReservationList() *SelfServiceReservationList {
	this := SelfServiceReservationList{}
	return &this
}

// NewSelfServiceReservationListWithDefaults instantiates a new SelfServiceReservationList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSelfServiceReservationListWithDefaults() *SelfServiceReservationList {
	this := SelfServiceReservationList{}
	return &this
}

// GetReservations returns the Reservations field value if set, zero value otherwise.
func (o *SelfServiceReservationList) GetReservations() []SelfServiceReservationResponse {
	if o == nil || o.Reservations == nil {
		var ret []SelfServiceReservationResponse
		return ret
	}
	return o.Reservations
}

// GetReservationsOk returns a tuple with the Reservations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceReservationList) GetReservationsOk() ([]SelfServiceReservationResponse, bool) {
	if o == nil || o.Reservations == nil {
		return nil, false
	}
	return o.Reservations, true
}

// HasReservations returns a boolean if a field has been set.
func (o *SelfServiceReservationList) HasReservations() bool {
	if o != nil && o.Reservations != nil {
		return true
	}

	return false
}

// SetReservations gets a reference to the given []SelfServiceReservationResponse and assigns it to the Reservations field.
func (o *SelfServiceReservationList) SetReservations(v []SelfServiceReservationResponse) {
	o.Reservations = v
}

func (o SelfServiceReservationList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Reservations != nil {
		toSerialize["reservations"] = o.Reservations
	}
	return json.Marshal(toSerialize)
}

type NullableSelfServiceReservationList struct {
	value *SelfServiceReservationList
	isSet bool
}

func (v NullableSelfServiceReservationList) Get() *SelfServiceReservationList {
	return v.value
}

func (v *NullableSelfServiceReservationList) Set(val *SelfServiceReservationList) {
	v.value = val
	v.isSet = true
}

func (v NullableSelfServiceReservationList) IsSet() bool {
	return v.isSet
}

func (v *NullableSelfServiceReservationList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelfServiceReservationList(val *SelfServiceReservationList) *NullableSelfServiceReservationList {
	return &NullableSelfServiceReservationList{value: val, isSet: true}
}

func (v NullableSelfServiceReservationList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSelfServiceReservationList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


