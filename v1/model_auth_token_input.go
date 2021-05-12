/*
 * Metal API
 *
 * This is the API for Equinix Metal. The API allows you to programmatically interact with all of your Equinix Metal resources, including devices, networks, addresses, organizations, projects, and your user account.  The official API docs are hosted at <https://metal.equinix.com/developers/api>. 
 *
 * API version: 1.0.0
 * Contact: support@equinixmetal.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
)

// AuthTokenInput struct for AuthTokenInput
type AuthTokenInput struct {
	Description *string `json:"description,omitempty"`
	ReadOnly *bool `json:"read_only,omitempty"`
}

// NewAuthTokenInput instantiates a new AuthTokenInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthTokenInput() *AuthTokenInput {
	this := AuthTokenInput{}
	return &this
}

// NewAuthTokenInputWithDefaults instantiates a new AuthTokenInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthTokenInputWithDefaults() *AuthTokenInput {
	this := AuthTokenInput{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AuthTokenInput) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthTokenInput) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AuthTokenInput) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AuthTokenInput) SetDescription(v string) {
	o.Description = &v
}

// GetReadOnly returns the ReadOnly field value if set, zero value otherwise.
func (o *AuthTokenInput) GetReadOnly() bool {
	if o == nil || o.ReadOnly == nil {
		var ret bool
		return ret
	}
	return *o.ReadOnly
}

// GetReadOnlyOk returns a tuple with the ReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthTokenInput) GetReadOnlyOk() (*bool, bool) {
	if o == nil || o.ReadOnly == nil {
		return nil, false
	}
	return o.ReadOnly, true
}

// HasReadOnly returns a boolean if a field has been set.
func (o *AuthTokenInput) HasReadOnly() bool {
	if o != nil && o.ReadOnly != nil {
		return true
	}

	return false
}

// SetReadOnly gets a reference to the given bool and assigns it to the ReadOnly field.
func (o *AuthTokenInput) SetReadOnly(v bool) {
	o.ReadOnly = &v
}

func (o AuthTokenInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.ReadOnly != nil {
		toSerialize["read_only"] = o.ReadOnly
	}
	return json.Marshal(toSerialize)
}

type NullableAuthTokenInput struct {
	value *AuthTokenInput
	isSet bool
}

func (v NullableAuthTokenInput) Get() *AuthTokenInput {
	return v.value
}

func (v *NullableAuthTokenInput) Set(val *AuthTokenInput) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthTokenInput) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthTokenInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthTokenInput(val *AuthTokenInput) *NullableAuthTokenInput {
	return &NullableAuthTokenInput{value: val, isSet: true}
}

func (v NullableAuthTokenInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthTokenInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


