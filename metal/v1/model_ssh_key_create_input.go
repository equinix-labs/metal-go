/*
Metal API

# Introduction Equinix Metal provides a RESTful HTTP API which can be reached at <https://api.equinix.com/metal/v1>. This document describes the API and how to use it.  The API allows you to programmatically interact with all of your Equinix Metal resources, including devices, networks, addresses, organizations, projects, and your user account. Every feature of the Equinix Metal web interface is accessible through the API.  The API docs are generated from the Equinix Metal OpenAPI specification and are officially hosted at <https://metal.equinix.com/developers/api>.  # Common Parameters  The Equinix Metal API uses a few methods to minimize network traffic and improve throughput. These parameters are not used in all API calls, but are used often enough to warrant their own section. Look for these parameters in the documentation for the API calls that support them.  ## Pagination  Pagination is used to limit the number of results returned in a single request. The API will return a maximum of 100 results per page. To retrieve additional results, you can use the `page` and `per_page` query parameters.  The `page` parameter is used to specify the page number. The first page is `1`. The `per_page` parameter is used to specify the number of results per page. The maximum number of results differs by resource type.  ## Sorting  Where offered, the API allows you to sort results by a specific field. To sort results use the `sort_by` query parameter with the root level field name as the value. The `sort_direction` parameter is used to specify the sort direction, either either `asc` (ascending) or `desc` (descending).  ## Filtering  Filtering is used to limit the results returned in a single request. The API supports filtering by certain fields in the response. To filter results, you can use the field as a query parameter.  For example, to filter the IP list to only return public IPv4 addresses, you can filter by the `type` field, as in the following request:  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/projects/id/ips?type=public_ipv4 ```  Only IP addresses with the `type` field set to `public_ipv4` will be returned.  ## Searching  Searching is used to find matching resources using multiple field comparissons. The API supports searching in resources that define this behavior. Currently the search parameter is only available on devices, ssh_keys, api_keys and memberships endpoints.  To search resources you can use the `search` query parameter.  ## Include and Exclude  For resources that contain references to other resources, sucha as a Device that refers to the Project it resides in, the Equinix Metal API will returns `href` values (API links) to the associated resource.  ```json {   ...   \"project\": {     \"href\": \"/metal/v1/projects/f3f131c8-f302-49ef-8c44-9405022dc6dd\"   } } ```  If you're going need the project details, you can avoid a second API request.  Specify the contained `href` resources and collections that you'd like to have included in the response using the `include` query parameter.  For example:  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=projects ```  The `include` parameter is generally accepted in `GET`, `POST`, `PUT`, and `PATCH` requests where `href` resources are presented.  To have multiple resources include, use a comma-separated list (e.g. `?include=emails,projects,memberships`).  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=emails,projects,memberships ```  You may also include nested associations up to three levels deep using dot notation (`?include=memberships.projects`):  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=memberships.projects ```  To exclude resources, and optimize response delivery, use the `exclude` query parameter. The `exclude` parameter is generally accepted in `GET`, `POST`, `PUT`, and `PATCH` requests for fields with nested object responses. When excluded, these fields will be replaced with an object that contains only an `href` field.

API version: 1.0.0
Contact: support@equinixmetal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
)

// checks if the SSHKeyCreateInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SSHKeyCreateInput{}

// SSHKeyCreateInput struct for SSHKeyCreateInput
type SSHKeyCreateInput struct {
	// List of instance UUIDs to associate SSH key with, when empty array is sent all instances belonging       to entity will be included
	InstancesIds         []string `json:"instances_ids,omitempty"`
	Key                  *string  `json:"key,omitempty"`
	Label                *string  `json:"label,omitempty"`
	Tags                 []string `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SSHKeyCreateInput SSHKeyCreateInput

// NewSSHKeyCreateInput instantiates a new SSHKeyCreateInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSSHKeyCreateInput() *SSHKeyCreateInput {
	this := SSHKeyCreateInput{}
	return &this
}

// NewSSHKeyCreateInputWithDefaults instantiates a new SSHKeyCreateInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSSHKeyCreateInputWithDefaults() *SSHKeyCreateInput {
	this := SSHKeyCreateInput{}
	return &this
}

// GetInstancesIds returns the InstancesIds field value if set, zero value otherwise.
func (o *SSHKeyCreateInput) GetInstancesIds() []string {
	if o == nil || IsNil(o.InstancesIds) {
		var ret []string
		return ret
	}
	return o.InstancesIds
}

// GetInstancesIdsOk returns a tuple with the InstancesIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSHKeyCreateInput) GetInstancesIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.InstancesIds) {
		return nil, false
	}
	return o.InstancesIds, true
}

// HasInstancesIds returns a boolean if a field has been set.
func (o *SSHKeyCreateInput) HasInstancesIds() bool {
	if o != nil && !IsNil(o.InstancesIds) {
		return true
	}

	return false
}

// SetInstancesIds gets a reference to the given []string and assigns it to the InstancesIds field.
func (o *SSHKeyCreateInput) SetInstancesIds(v []string) {
	o.InstancesIds = v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *SSHKeyCreateInput) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSHKeyCreateInput) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *SSHKeyCreateInput) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *SSHKeyCreateInput) SetKey(v string) {
	o.Key = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *SSHKeyCreateInput) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSHKeyCreateInput) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *SSHKeyCreateInput) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *SSHKeyCreateInput) SetLabel(v string) {
	o.Label = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *SSHKeyCreateInput) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSHKeyCreateInput) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *SSHKeyCreateInput) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *SSHKeyCreateInput) SetTags(v []string) {
	o.Tags = v
}

func (o SSHKeyCreateInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SSHKeyCreateInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InstancesIds) {
		toSerialize["instances_ids"] = o.InstancesIds
	}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SSHKeyCreateInput) UnmarshalJSON(bytes []byte) (err error) {
	varSSHKeyCreateInput := _SSHKeyCreateInput{}

	err = json.Unmarshal(bytes, &varSSHKeyCreateInput)

	if err != nil {
		return err
	}

	*o = SSHKeyCreateInput(varSSHKeyCreateInput)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "instances_ids")
		delete(additionalProperties, "key")
		delete(additionalProperties, "label")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSSHKeyCreateInput struct {
	value *SSHKeyCreateInput
	isSet bool
}

func (v NullableSSHKeyCreateInput) Get() *SSHKeyCreateInput {
	return v.value
}

func (v *NullableSSHKeyCreateInput) Set(val *SSHKeyCreateInput) {
	v.value = val
	v.isSet = true
}

func (v NullableSSHKeyCreateInput) IsSet() bool {
	return v.isSet
}

func (v *NullableSSHKeyCreateInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSSHKeyCreateInput(val *SSHKeyCreateInput) *NullableSSHKeyCreateInput {
	return &NullableSSHKeyCreateInput{value: val, isSet: true}
}

func (v NullableSSHKeyCreateInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSSHKeyCreateInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
