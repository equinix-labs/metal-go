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

// UpdateDeviceRequest struct for UpdateDeviceRequest
type UpdateDeviceRequest struct {
	Tags          []string               `json:"tags,omitempty"`
	AlwaysPxe     *bool                  `json:"always_pxe,omitempty"`
	BillingCycle  *string                `json:"billing_cycle,omitempty"`
	Customdata    map[string]interface{} `json:"customdata,omitempty"`
	Description   *string                `json:"description,omitempty"`
	Hostname      *string                `json:"hostname,omitempty"`
	IpxeScriptUrl *string                `json:"ipxe_script_url,omitempty"`
	Locked        *bool                  `json:"locked,omitempty"`
	// If true, this instance can not be converted to a different network type.
	NetworkFrozen *bool   `json:"network_frozen,omitempty"`
	SpotInstance  *bool   `json:"spot_instance,omitempty"`
	Userdata      *string `json:"userdata,omitempty"`
}

// NewUpdateDeviceRequest instantiates a new UpdateDeviceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDeviceRequest() *UpdateDeviceRequest {
	this := UpdateDeviceRequest{}
	return &this
}

// NewUpdateDeviceRequestWithDefaults instantiates a new UpdateDeviceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDeviceRequestWithDefaults() *UpdateDeviceRequest {
	this := UpdateDeviceRequest{}
	return &this
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *UpdateDeviceRequest) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceRequest) GetTagsOk() ([]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *UpdateDeviceRequest) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *UpdateDeviceRequest) SetTags(v []string) {
	o.Tags = v
}

// GetAlwaysPxe returns the AlwaysPxe field value if set, zero value otherwise.
func (o *UpdateDeviceRequest) GetAlwaysPxe() bool {
	if o == nil || o.AlwaysPxe == nil {
		var ret bool
		return ret
	}
	return *o.AlwaysPxe
}

// GetAlwaysPxeOk returns a tuple with the AlwaysPxe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceRequest) GetAlwaysPxeOk() (*bool, bool) {
	if o == nil || o.AlwaysPxe == nil {
		return nil, false
	}
	return o.AlwaysPxe, true
}

// HasAlwaysPxe returns a boolean if a field has been set.
func (o *UpdateDeviceRequest) HasAlwaysPxe() bool {
	if o != nil && o.AlwaysPxe != nil {
		return true
	}

	return false
}

// SetAlwaysPxe gets a reference to the given bool and assigns it to the AlwaysPxe field.
func (o *UpdateDeviceRequest) SetAlwaysPxe(v bool) {
	o.AlwaysPxe = &v
}

// GetBillingCycle returns the BillingCycle field value if set, zero value otherwise.
func (o *UpdateDeviceRequest) GetBillingCycle() string {
	if o == nil || o.BillingCycle == nil {
		var ret string
		return ret
	}
	return *o.BillingCycle
}

// GetBillingCycleOk returns a tuple with the BillingCycle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceRequest) GetBillingCycleOk() (*string, bool) {
	if o == nil || o.BillingCycle == nil {
		return nil, false
	}
	return o.BillingCycle, true
}

// HasBillingCycle returns a boolean if a field has been set.
func (o *UpdateDeviceRequest) HasBillingCycle() bool {
	if o != nil && o.BillingCycle != nil {
		return true
	}

	return false
}

// SetBillingCycle gets a reference to the given string and assigns it to the BillingCycle field.
func (o *UpdateDeviceRequest) SetBillingCycle(v string) {
	o.BillingCycle = &v
}

// GetCustomdata returns the Customdata field value if set, zero value otherwise.
func (o *UpdateDeviceRequest) GetCustomdata() map[string]interface{} {
	if o == nil || o.Customdata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Customdata
}

// GetCustomdataOk returns a tuple with the Customdata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceRequest) GetCustomdataOk() (map[string]interface{}, bool) {
	if o == nil || o.Customdata == nil {
		return nil, false
	}
	return o.Customdata, true
}

// HasCustomdata returns a boolean if a field has been set.
func (o *UpdateDeviceRequest) HasCustomdata() bool {
	if o != nil && o.Customdata != nil {
		return true
	}

	return false
}

// SetCustomdata gets a reference to the given map[string]interface{} and assigns it to the Customdata field.
func (o *UpdateDeviceRequest) SetCustomdata(v map[string]interface{}) {
	o.Customdata = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateDeviceRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateDeviceRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateDeviceRequest) SetDescription(v string) {
	o.Description = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *UpdateDeviceRequest) GetHostname() string {
	if o == nil || o.Hostname == nil {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceRequest) GetHostnameOk() (*string, bool) {
	if o == nil || o.Hostname == nil {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *UpdateDeviceRequest) HasHostname() bool {
	if o != nil && o.Hostname != nil {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *UpdateDeviceRequest) SetHostname(v string) {
	o.Hostname = &v
}

// GetIpxeScriptUrl returns the IpxeScriptUrl field value if set, zero value otherwise.
func (o *UpdateDeviceRequest) GetIpxeScriptUrl() string {
	if o == nil || o.IpxeScriptUrl == nil {
		var ret string
		return ret
	}
	return *o.IpxeScriptUrl
}

// GetIpxeScriptUrlOk returns a tuple with the IpxeScriptUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceRequest) GetIpxeScriptUrlOk() (*string, bool) {
	if o == nil || o.IpxeScriptUrl == nil {
		return nil, false
	}
	return o.IpxeScriptUrl, true
}

// HasIpxeScriptUrl returns a boolean if a field has been set.
func (o *UpdateDeviceRequest) HasIpxeScriptUrl() bool {
	if o != nil && o.IpxeScriptUrl != nil {
		return true
	}

	return false
}

// SetIpxeScriptUrl gets a reference to the given string and assigns it to the IpxeScriptUrl field.
func (o *UpdateDeviceRequest) SetIpxeScriptUrl(v string) {
	o.IpxeScriptUrl = &v
}

// GetLocked returns the Locked field value if set, zero value otherwise.
func (o *UpdateDeviceRequest) GetLocked() bool {
	if o == nil || o.Locked == nil {
		var ret bool
		return ret
	}
	return *o.Locked
}

// GetLockedOk returns a tuple with the Locked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceRequest) GetLockedOk() (*bool, bool) {
	if o == nil || o.Locked == nil {
		return nil, false
	}
	return o.Locked, true
}

// HasLocked returns a boolean if a field has been set.
func (o *UpdateDeviceRequest) HasLocked() bool {
	if o != nil && o.Locked != nil {
		return true
	}

	return false
}

// SetLocked gets a reference to the given bool and assigns it to the Locked field.
func (o *UpdateDeviceRequest) SetLocked(v bool) {
	o.Locked = &v
}

// GetNetworkFrozen returns the NetworkFrozen field value if set, zero value otherwise.
func (o *UpdateDeviceRequest) GetNetworkFrozen() bool {
	if o == nil || o.NetworkFrozen == nil {
		var ret bool
		return ret
	}
	return *o.NetworkFrozen
}

// GetNetworkFrozenOk returns a tuple with the NetworkFrozen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceRequest) GetNetworkFrozenOk() (*bool, bool) {
	if o == nil || o.NetworkFrozen == nil {
		return nil, false
	}
	return o.NetworkFrozen, true
}

// HasNetworkFrozen returns a boolean if a field has been set.
func (o *UpdateDeviceRequest) HasNetworkFrozen() bool {
	if o != nil && o.NetworkFrozen != nil {
		return true
	}

	return false
}

// SetNetworkFrozen gets a reference to the given bool and assigns it to the NetworkFrozen field.
func (o *UpdateDeviceRequest) SetNetworkFrozen(v bool) {
	o.NetworkFrozen = &v
}

// GetSpotInstance returns the SpotInstance field value if set, zero value otherwise.
func (o *UpdateDeviceRequest) GetSpotInstance() bool {
	if o == nil || o.SpotInstance == nil {
		var ret bool
		return ret
	}
	return *o.SpotInstance
}

// GetSpotInstanceOk returns a tuple with the SpotInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceRequest) GetSpotInstanceOk() (*bool, bool) {
	if o == nil || o.SpotInstance == nil {
		return nil, false
	}
	return o.SpotInstance, true
}

// HasSpotInstance returns a boolean if a field has been set.
func (o *UpdateDeviceRequest) HasSpotInstance() bool {
	if o != nil && o.SpotInstance != nil {
		return true
	}

	return false
}

// SetSpotInstance gets a reference to the given bool and assigns it to the SpotInstance field.
func (o *UpdateDeviceRequest) SetSpotInstance(v bool) {
	o.SpotInstance = &v
}

// GetUserdata returns the Userdata field value if set, zero value otherwise.
func (o *UpdateDeviceRequest) GetUserdata() string {
	if o == nil || o.Userdata == nil {
		var ret string
		return ret
	}
	return *o.Userdata
}

// GetUserdataOk returns a tuple with the Userdata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceRequest) GetUserdataOk() (*string, bool) {
	if o == nil || o.Userdata == nil {
		return nil, false
	}
	return o.Userdata, true
}

// HasUserdata returns a boolean if a field has been set.
func (o *UpdateDeviceRequest) HasUserdata() bool {
	if o != nil && o.Userdata != nil {
		return true
	}

	return false
}

// SetUserdata gets a reference to the given string and assigns it to the Userdata field.
func (o *UpdateDeviceRequest) SetUserdata(v string) {
	o.Userdata = &v
}

func (o UpdateDeviceRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.AlwaysPxe != nil {
		toSerialize["always_pxe"] = o.AlwaysPxe
	}
	if o.BillingCycle != nil {
		toSerialize["billing_cycle"] = o.BillingCycle
	}
	if o.Customdata != nil {
		toSerialize["customdata"] = o.Customdata
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Hostname != nil {
		toSerialize["hostname"] = o.Hostname
	}
	if o.IpxeScriptUrl != nil {
		toSerialize["ipxe_script_url"] = o.IpxeScriptUrl
	}
	if o.Locked != nil {
		toSerialize["locked"] = o.Locked
	}
	if o.NetworkFrozen != nil {
		toSerialize["network_frozen"] = o.NetworkFrozen
	}
	if o.SpotInstance != nil {
		toSerialize["spot_instance"] = o.SpotInstance
	}
	if o.Userdata != nil {
		toSerialize["userdata"] = o.Userdata
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateDeviceRequest struct {
	value *UpdateDeviceRequest
	isSet bool
}

func (v NullableUpdateDeviceRequest) Get() *UpdateDeviceRequest {
	return v.value
}

func (v *NullableUpdateDeviceRequest) Set(val *UpdateDeviceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDeviceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDeviceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDeviceRequest(val *UpdateDeviceRequest) *NullableUpdateDeviceRequest {
	return &NullableUpdateDeviceRequest{value: val, isSet: true}
}

func (v NullableUpdateDeviceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDeviceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}