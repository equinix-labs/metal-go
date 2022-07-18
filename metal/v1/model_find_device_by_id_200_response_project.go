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

// FindDeviceById200ResponseProject struct for FindDeviceById200ResponseProject
type FindDeviceById200ResponseProject struct {
	Href string `json:"href"`
}

// NewFindDeviceById200ResponseProject instantiates a new FindDeviceById200ResponseProject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFindDeviceById200ResponseProject(href string) *FindDeviceById200ResponseProject {
	this := FindDeviceById200ResponseProject{}
	this.Href = href
	return &this
}

// NewFindDeviceById200ResponseProjectWithDefaults instantiates a new FindDeviceById200ResponseProject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFindDeviceById200ResponseProjectWithDefaults() *FindDeviceById200ResponseProject {
	this := FindDeviceById200ResponseProject{}
	return &this
}

// GetHref returns the Href field value
func (o *FindDeviceById200ResponseProject) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *FindDeviceById200ResponseProject) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *FindDeviceById200ResponseProject) SetHref(v string) {
	o.Href = v
}

func (o FindDeviceById200ResponseProject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["href"] = o.Href
	}
	return json.Marshal(toSerialize)
}

type NullableFindDeviceById200ResponseProject struct {
	value *FindDeviceById200ResponseProject
	isSet bool
}

func (v NullableFindDeviceById200ResponseProject) Get() *FindDeviceById200ResponseProject {
	return v.value
}

func (v *NullableFindDeviceById200ResponseProject) Set(val *FindDeviceById200ResponseProject) {
	v.value = val
	v.isSet = true
}

func (v NullableFindDeviceById200ResponseProject) IsSet() bool {
	return v.isSet
}

func (v *NullableFindDeviceById200ResponseProject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFindDeviceById200ResponseProject(val *FindDeviceById200ResponseProject) *NullableFindDeviceById200ResponseProject {
	return &NullableFindDeviceById200ResponseProject{value: val, isSet: true}
}

func (v NullableFindDeviceById200ResponseProject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFindDeviceById200ResponseProject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}