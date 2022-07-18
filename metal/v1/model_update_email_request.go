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

// UpdateEmailRequest struct for UpdateEmailRequest
type UpdateEmailRequest struct {
	Default *bool `json:"default,omitempty"`
}

// NewUpdateEmailRequest instantiates a new UpdateEmailRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateEmailRequest() *UpdateEmailRequest {
	this := UpdateEmailRequest{}
	return &this
}

// NewUpdateEmailRequestWithDefaults instantiates a new UpdateEmailRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateEmailRequestWithDefaults() *UpdateEmailRequest {
	this := UpdateEmailRequest{}
	return &this
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *UpdateEmailRequest) GetDefault() bool {
	if o == nil || o.Default == nil {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEmailRequest) GetDefaultOk() (*bool, bool) {
	if o == nil || o.Default == nil {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *UpdateEmailRequest) HasDefault() bool {
	if o != nil && o.Default != nil {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *UpdateEmailRequest) SetDefault(v bool) {
	o.Default = &v
}

func (o UpdateEmailRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Default != nil {
		toSerialize["default"] = o.Default
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateEmailRequest struct {
	value *UpdateEmailRequest
	isSet bool
}

func (v NullableUpdateEmailRequest) Get() *UpdateEmailRequest {
	return v.value
}

func (v *NullableUpdateEmailRequest) Set(val *UpdateEmailRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateEmailRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateEmailRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateEmailRequest(val *UpdateEmailRequest) *NullableUpdateEmailRequest {
	return &NullableUpdateEmailRequest{value: val, isSet: true}
}

func (v NullableUpdateEmailRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateEmailRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}