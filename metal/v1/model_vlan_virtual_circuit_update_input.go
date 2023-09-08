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

// checks if the VlanVirtualCircuitUpdateInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VlanVirtualCircuitUpdateInput{}

// VlanVirtualCircuitUpdateInput struct for VlanVirtualCircuitUpdateInput
type VlanVirtualCircuitUpdateInput struct {
	Description *string `json:"description,omitempty"`
	Name        *string `json:"name,omitempty"`
	// Speed can be changed only if it is an interconnection on a Dedicated Port
	Speed *string  `json:"speed,omitempty"`
	Tags  []string `json:"tags,omitempty"`
	// A Virtual Network record UUID or the VNID of a Metro Virtual Network in your project.
	Vnid                 *string `json:"vnid,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VlanVirtualCircuitUpdateInput VlanVirtualCircuitUpdateInput

// NewVlanVirtualCircuitUpdateInput instantiates a new VlanVirtualCircuitUpdateInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVlanVirtualCircuitUpdateInput() *VlanVirtualCircuitUpdateInput {
	this := VlanVirtualCircuitUpdateInput{}
	return &this
}

// NewVlanVirtualCircuitUpdateInputWithDefaults instantiates a new VlanVirtualCircuitUpdateInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVlanVirtualCircuitUpdateInputWithDefaults() *VlanVirtualCircuitUpdateInput {
	this := VlanVirtualCircuitUpdateInput{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *VlanVirtualCircuitUpdateInput) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VlanVirtualCircuitUpdateInput) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *VlanVirtualCircuitUpdateInput) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *VlanVirtualCircuitUpdateInput) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VlanVirtualCircuitUpdateInput) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VlanVirtualCircuitUpdateInput) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VlanVirtualCircuitUpdateInput) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VlanVirtualCircuitUpdateInput) SetName(v string) {
	o.Name = &v
}

// GetSpeed returns the Speed field value if set, zero value otherwise.
func (o *VlanVirtualCircuitUpdateInput) GetSpeed() string {
	if o == nil || IsNil(o.Speed) {
		var ret string
		return ret
	}
	return *o.Speed
}

// GetSpeedOk returns a tuple with the Speed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VlanVirtualCircuitUpdateInput) GetSpeedOk() (*string, bool) {
	if o == nil || IsNil(o.Speed) {
		return nil, false
	}
	return o.Speed, true
}

// HasSpeed returns a boolean if a field has been set.
func (o *VlanVirtualCircuitUpdateInput) HasSpeed() bool {
	if o != nil && !IsNil(o.Speed) {
		return true
	}

	return false
}

// SetSpeed gets a reference to the given string and assigns it to the Speed field.
func (o *VlanVirtualCircuitUpdateInput) SetSpeed(v string) {
	o.Speed = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *VlanVirtualCircuitUpdateInput) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VlanVirtualCircuitUpdateInput) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *VlanVirtualCircuitUpdateInput) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *VlanVirtualCircuitUpdateInput) SetTags(v []string) {
	o.Tags = v
}

// GetVnid returns the Vnid field value if set, zero value otherwise.
func (o *VlanVirtualCircuitUpdateInput) GetVnid() string {
	if o == nil || IsNil(o.Vnid) {
		var ret string
		return ret
	}
	return *o.Vnid
}

// GetVnidOk returns a tuple with the Vnid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VlanVirtualCircuitUpdateInput) GetVnidOk() (*string, bool) {
	if o == nil || IsNil(o.Vnid) {
		return nil, false
	}
	return o.Vnid, true
}

// HasVnid returns a boolean if a field has been set.
func (o *VlanVirtualCircuitUpdateInput) HasVnid() bool {
	if o != nil && !IsNil(o.Vnid) {
		return true
	}

	return false
}

// SetVnid gets a reference to the given string and assigns it to the Vnid field.
func (o *VlanVirtualCircuitUpdateInput) SetVnid(v string) {
	o.Vnid = &v
}

func (o VlanVirtualCircuitUpdateInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VlanVirtualCircuitUpdateInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Speed) {
		toSerialize["speed"] = o.Speed
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Vnid) {
		toSerialize["vnid"] = o.Vnid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VlanVirtualCircuitUpdateInput) UnmarshalJSON(bytes []byte) (err error) {
	varVlanVirtualCircuitUpdateInput := _VlanVirtualCircuitUpdateInput{}

	if err = json.Unmarshal(bytes, &varVlanVirtualCircuitUpdateInput); err == nil {
		*o = VlanVirtualCircuitUpdateInput(varVlanVirtualCircuitUpdateInput)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "name")
		delete(additionalProperties, "speed")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "vnid")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVlanVirtualCircuitUpdateInput struct {
	value *VlanVirtualCircuitUpdateInput
	isSet bool
}

func (v NullableVlanVirtualCircuitUpdateInput) Get() *VlanVirtualCircuitUpdateInput {
	return v.value
}

func (v *NullableVlanVirtualCircuitUpdateInput) Set(val *VlanVirtualCircuitUpdateInput) {
	v.value = val
	v.isSet = true
}

func (v NullableVlanVirtualCircuitUpdateInput) IsSet() bool {
	return v.isSet
}

func (v *NullableVlanVirtualCircuitUpdateInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVlanVirtualCircuitUpdateInput(val *VlanVirtualCircuitUpdateInput) *NullableVlanVirtualCircuitUpdateInput {
	return &NullableVlanVirtualCircuitUpdateInput{value: val, isSet: true}
}

func (v NullableVlanVirtualCircuitUpdateInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVlanVirtualCircuitUpdateInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}