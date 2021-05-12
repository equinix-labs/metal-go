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

// SSHKeyInput struct for SSHKeyInput
type SSHKeyInput struct {
	Label *string `json:"label,omitempty"`
	Key *string `json:"key,omitempty"`
}

// NewSSHKeyInput instantiates a new SSHKeyInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSSHKeyInput() *SSHKeyInput {
	this := SSHKeyInput{}
	return &this
}

// NewSSHKeyInputWithDefaults instantiates a new SSHKeyInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSSHKeyInputWithDefaults() *SSHKeyInput {
	this := SSHKeyInput{}
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *SSHKeyInput) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSHKeyInput) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *SSHKeyInput) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *SSHKeyInput) SetLabel(v string) {
	o.Label = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *SSHKeyInput) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSHKeyInput) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *SSHKeyInput) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *SSHKeyInput) SetKey(v string) {
	o.Key = &v
}

func (o SSHKeyInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	return json.Marshal(toSerialize)
}

type NullableSSHKeyInput struct {
	value *SSHKeyInput
	isSet bool
}

func (v NullableSSHKeyInput) Get() *SSHKeyInput {
	return v.value
}

func (v *NullableSSHKeyInput) Set(val *SSHKeyInput) {
	v.value = val
	v.isSet = true
}

func (v NullableSSHKeyInput) IsSet() bool {
	return v.isSet
}

func (v *NullableSSHKeyInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSSHKeyInput(val *SSHKeyInput) *NullableSSHKeyInput {
	return &NullableSSHKeyInput{value: val, isSet: true}
}

func (v NullableSSHKeyInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSSHKeyInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


