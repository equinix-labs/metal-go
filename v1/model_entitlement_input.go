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

// EntitlementInput struct for EntitlementInput
type EntitlementInput struct {
	Description *string `json:"description,omitempty"`
	Slug *string `json:"slug,omitempty"`
	Weight *int32 `json:"weight,omitempty"`
	InstanceQuota *map[string]interface{} `json:"instance_quota,omitempty"`
	ProjectQuota *int32 `json:"project_quota,omitempty"`
	VolumeQuota *map[string]interface{} `json:"volume_quota,omitempty"`
	FeatureAccess *map[string]interface{} `json:"feature_access,omitempty"`
}

// NewEntitlementInput instantiates a new EntitlementInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementInput() *EntitlementInput {
	this := EntitlementInput{}
	return &this
}

// NewEntitlementInputWithDefaults instantiates a new EntitlementInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementInputWithDefaults() *EntitlementInput {
	this := EntitlementInput{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EntitlementInput) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementInput) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EntitlementInput) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EntitlementInput) SetDescription(v string) {
	o.Description = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *EntitlementInput) GetSlug() string {
	if o == nil || o.Slug == nil {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementInput) GetSlugOk() (*string, bool) {
	if o == nil || o.Slug == nil {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *EntitlementInput) HasSlug() bool {
	if o != nil && o.Slug != nil {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *EntitlementInput) SetSlug(v string) {
	o.Slug = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *EntitlementInput) GetWeight() int32 {
	if o == nil || o.Weight == nil {
		var ret int32
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementInput) GetWeightOk() (*int32, bool) {
	if o == nil || o.Weight == nil {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *EntitlementInput) HasWeight() bool {
	if o != nil && o.Weight != nil {
		return true
	}

	return false
}

// SetWeight gets a reference to the given int32 and assigns it to the Weight field.
func (o *EntitlementInput) SetWeight(v int32) {
	o.Weight = &v
}

// GetInstanceQuota returns the InstanceQuota field value if set, zero value otherwise.
func (o *EntitlementInput) GetInstanceQuota() map[string]interface{} {
	if o == nil || o.InstanceQuota == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.InstanceQuota
}

// GetInstanceQuotaOk returns a tuple with the InstanceQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementInput) GetInstanceQuotaOk() (*map[string]interface{}, bool) {
	if o == nil || o.InstanceQuota == nil {
		return nil, false
	}
	return o.InstanceQuota, true
}

// HasInstanceQuota returns a boolean if a field has been set.
func (o *EntitlementInput) HasInstanceQuota() bool {
	if o != nil && o.InstanceQuota != nil {
		return true
	}

	return false
}

// SetInstanceQuota gets a reference to the given map[string]interface{} and assigns it to the InstanceQuota field.
func (o *EntitlementInput) SetInstanceQuota(v map[string]interface{}) {
	o.InstanceQuota = &v
}

// GetProjectQuota returns the ProjectQuota field value if set, zero value otherwise.
func (o *EntitlementInput) GetProjectQuota() int32 {
	if o == nil || o.ProjectQuota == nil {
		var ret int32
		return ret
	}
	return *o.ProjectQuota
}

// GetProjectQuotaOk returns a tuple with the ProjectQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementInput) GetProjectQuotaOk() (*int32, bool) {
	if o == nil || o.ProjectQuota == nil {
		return nil, false
	}
	return o.ProjectQuota, true
}

// HasProjectQuota returns a boolean if a field has been set.
func (o *EntitlementInput) HasProjectQuota() bool {
	if o != nil && o.ProjectQuota != nil {
		return true
	}

	return false
}

// SetProjectQuota gets a reference to the given int32 and assigns it to the ProjectQuota field.
func (o *EntitlementInput) SetProjectQuota(v int32) {
	o.ProjectQuota = &v
}

// GetVolumeQuota returns the VolumeQuota field value if set, zero value otherwise.
func (o *EntitlementInput) GetVolumeQuota() map[string]interface{} {
	if o == nil || o.VolumeQuota == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.VolumeQuota
}

// GetVolumeQuotaOk returns a tuple with the VolumeQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementInput) GetVolumeQuotaOk() (*map[string]interface{}, bool) {
	if o == nil || o.VolumeQuota == nil {
		return nil, false
	}
	return o.VolumeQuota, true
}

// HasVolumeQuota returns a boolean if a field has been set.
func (o *EntitlementInput) HasVolumeQuota() bool {
	if o != nil && o.VolumeQuota != nil {
		return true
	}

	return false
}

// SetVolumeQuota gets a reference to the given map[string]interface{} and assigns it to the VolumeQuota field.
func (o *EntitlementInput) SetVolumeQuota(v map[string]interface{}) {
	o.VolumeQuota = &v
}

// GetFeatureAccess returns the FeatureAccess field value if set, zero value otherwise.
func (o *EntitlementInput) GetFeatureAccess() map[string]interface{} {
	if o == nil || o.FeatureAccess == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.FeatureAccess
}

// GetFeatureAccessOk returns a tuple with the FeatureAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementInput) GetFeatureAccessOk() (*map[string]interface{}, bool) {
	if o == nil || o.FeatureAccess == nil {
		return nil, false
	}
	return o.FeatureAccess, true
}

// HasFeatureAccess returns a boolean if a field has been set.
func (o *EntitlementInput) HasFeatureAccess() bool {
	if o != nil && o.FeatureAccess != nil {
		return true
	}

	return false
}

// SetFeatureAccess gets a reference to the given map[string]interface{} and assigns it to the FeatureAccess field.
func (o *EntitlementInput) SetFeatureAccess(v map[string]interface{}) {
	o.FeatureAccess = &v
}

func (o EntitlementInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Slug != nil {
		toSerialize["slug"] = o.Slug
	}
	if o.Weight != nil {
		toSerialize["weight"] = o.Weight
	}
	if o.InstanceQuota != nil {
		toSerialize["instance_quota"] = o.InstanceQuota
	}
	if o.ProjectQuota != nil {
		toSerialize["project_quota"] = o.ProjectQuota
	}
	if o.VolumeQuota != nil {
		toSerialize["volume_quota"] = o.VolumeQuota
	}
	if o.FeatureAccess != nil {
		toSerialize["feature_access"] = o.FeatureAccess
	}
	return json.Marshal(toSerialize)
}

type NullableEntitlementInput struct {
	value *EntitlementInput
	isSet bool
}

func (v NullableEntitlementInput) Get() *EntitlementInput {
	return v.value
}

func (v *NullableEntitlementInput) Set(val *EntitlementInput) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementInput) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementInput(val *EntitlementInput) *NullableEntitlementInput {
	return &NullableEntitlementInput{value: val, isSet: true}
}

func (v NullableEntitlementInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


