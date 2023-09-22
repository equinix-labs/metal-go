/*
Metal API

# Introduction Equinix Metal provides a RESTful HTTP API which can be reached at <https://api.equinix.com/metal/v1>. This document describes the API and how to use it.  The API allows you to programmatically interact with all of your Equinix Metal resources, including devices, networks, addresses, organizations, projects, and your user account. Every feature of the Equinix Metal web interface is accessible through the API.  The API docs are generated from the Equinix Metal OpenAPI specification and are officially hosted at <https://metal.equinix.com/developers/api>.  # Common Parameters  The Equinix Metal API uses a few methods to minimize network traffic and improve throughput. These parameters are not used in all API calls, but are used often enough to warrant their own section. Look for these parameters in the documentation for the API calls that support them.  ## Pagination  Pagination is used to limit the number of results returned in a single request. The API will return a maximum of 100 results per page. To retrieve additional results, you can use the `page` and `per_page` query parameters.  The `page` parameter is used to specify the page number. The first page is `1`. The `per_page` parameter is used to specify the number of results per page. The maximum number of results differs by resource type.  ## Sorting  Where offered, the API allows you to sort results by a specific field. To sort results use the `sort_by` query parameter with the root level field name as the value. The `sort_direction` parameter is used to specify the sort direction, either either `asc` (ascending) or `desc` (descending).  ## Filtering  Filtering is used to limit the results returned in a single request. The API supports filtering by certain fields in the response. To filter results, you can use the field as a query parameter.  For example, to filter the IP list to only return public IPv4 addresses, you can filter by the `type` field, as in the following request:  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/projects/id/ips?type=public_ipv4 ```  Only IP addresses with the `type` field set to `public_ipv4` will be returned.  ## Searching  Searching is used to find matching resources using multiple field comparisons. The API supports searching in resources that define this behavior. Currently the search parameter is only available on devices, ssh_keys, api_keys and memberships endpoints.  To search resources you can use the `search` query parameter.  ## Include and Exclude  For resources that contain references to other resources, such as a Device that refers to the Project it resides in, the Equinix Metal API will returns `href` values (API links) to the associated resource.  ```json {   ...   \"project\": {     \"href\": \"/metal/v1/projects/f3f131c8-f302-49ef-8c44-9405022dc6dd\"   } } ```  If you're going need the project details, you can avoid a second API request.  Specify the contained `href` resources and collections that you'd like to have included in the response using the `include` query parameter.  For example:  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=projects ```  The `include` parameter is generally accepted in `GET`, `POST`, `PUT`, and `PATCH` requests where `href` resources are presented.  To have multiple resources include, use a comma-separated list (e.g. `?include=emails,projects,memberships`).  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=emails,projects,memberships ```  You may also include nested associations up to three levels deep using dot notation (`?include=memberships.projects`):  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=memberships.projects ```  To exclude resources, and optimize response delivery, use the `exclude` query parameter. The `exclude` parameter is generally accepted in `GET`, `POST`, `PUT`, and `PATCH` requests for fields with nested object responses. When excluded, these fields will be replaced with an object that contains only an `href` field.

API version: 1.0.0
Contact: support@equinixmetal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
)

// checks if the Disk type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Disk{}

// Disk struct for Disk
type Disk struct {
	Device               *string     `json:"device,omitempty"`
	WipeTable            *bool       `json:"wipeTable,omitempty"`
	Partitions           []Partition `json:"partitions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Disk Disk

// NewDisk instantiates a new Disk object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDisk() *Disk {
	this := Disk{}
	return &this
}

// NewDiskWithDefaults instantiates a new Disk object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiskWithDefaults() *Disk {
	this := Disk{}
	return &this
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *Disk) GetDevice() string {
	if o == nil || IsNil(o.Device) {
		var ret string
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Disk) GetDeviceOk() (*string, bool) {
	if o == nil || IsNil(o.Device) {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *Disk) HasDevice() bool {
	if o != nil && !IsNil(o.Device) {
		return true
	}

	return false
}

// SetDevice gets a reference to the given string and assigns it to the Device field.
func (o *Disk) SetDevice(v string) {
	o.Device = &v
}

// GetWipeTable returns the WipeTable field value if set, zero value otherwise.
func (o *Disk) GetWipeTable() bool {
	if o == nil || IsNil(o.WipeTable) {
		var ret bool
		return ret
	}
	return *o.WipeTable
}

// GetWipeTableOk returns a tuple with the WipeTable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Disk) GetWipeTableOk() (*bool, bool) {
	if o == nil || IsNil(o.WipeTable) {
		return nil, false
	}
	return o.WipeTable, true
}

// HasWipeTable returns a boolean if a field has been set.
func (o *Disk) HasWipeTable() bool {
	if o != nil && !IsNil(o.WipeTable) {
		return true
	}

	return false
}

// SetWipeTable gets a reference to the given bool and assigns it to the WipeTable field.
func (o *Disk) SetWipeTable(v bool) {
	o.WipeTable = &v
}

// GetPartitions returns the Partitions field value if set, zero value otherwise.
func (o *Disk) GetPartitions() []Partition {
	if o == nil || IsNil(o.Partitions) {
		var ret []Partition
		return ret
	}
	return o.Partitions
}

// GetPartitionsOk returns a tuple with the Partitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Disk) GetPartitionsOk() ([]Partition, bool) {
	if o == nil || IsNil(o.Partitions) {
		return nil, false
	}
	return o.Partitions, true
}

// HasPartitions returns a boolean if a field has been set.
func (o *Disk) HasPartitions() bool {
	if o != nil && !IsNil(o.Partitions) {
		return true
	}

	return false
}

// SetPartitions gets a reference to the given []Partition and assigns it to the Partitions field.
func (o *Disk) SetPartitions(v []Partition) {
	o.Partitions = v
}

func (o Disk) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Disk) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Device) {
		toSerialize["device"] = o.Device
	}
	if !IsNil(o.WipeTable) {
		toSerialize["wipeTable"] = o.WipeTable
	}
	if !IsNil(o.Partitions) {
		toSerialize["partitions"] = o.Partitions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Disk) UnmarshalJSON(bytes []byte) (err error) {
	varDisk := _Disk{}

	err = json.Unmarshal(bytes, &varDisk)

	if err != nil {
		return err
	}

	*o = Disk(varDisk)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "device")
		delete(additionalProperties, "wipeTable")
		delete(additionalProperties, "partitions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDisk struct {
	value *Disk
	isSet bool
}

func (v NullableDisk) Get() *Disk {
	return v.value
}

func (v *NullableDisk) Set(val *Disk) {
	v.value = val
	v.isSet = true
}

func (v NullableDisk) IsSet() bool {
	return v.isSet
}

func (v *NullableDisk) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDisk(val *Disk) *NullableDisk {
	return &NullableDisk{value: val, isSet: true}
}

func (v NullableDisk) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDisk) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
