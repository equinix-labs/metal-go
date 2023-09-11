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
	"fmt"
)

// checks if the PlanSpecsNicsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlanSpecsNicsInner{}

// PlanSpecsNicsInnerType the model 'PlanSpecsNicsInnerType'
type PlanSpecsNicsInnerType string

// List of PlanSpecsNicsInnerType
const (
	PLANSPECSNICSINNER__1_GBPS  PlanSpecsNicsInnerType = "1Gbps"
	PLANSPECSNICSINNER__10_GBPS PlanSpecsNicsInnerType = "10Gbps"
	PLANSPECSNICSINNER__25_GBPS PlanSpecsNicsInnerType = "25Gbps"
)

// All allowed values of PlanSpecsNicsInnerType enum
var AllowedPlanSpecsNicsInnerTypeEnumValues = []PlanSpecsNicsInnerType{
	"1Gbps",
	"10Gbps",
	"25Gbps",
}

func (v *PlanSpecsNicsInnerType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PlanSpecsNicsInnerType(value)
	for _, existing := range AllowedPlanSpecsNicsInnerTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PlanSpecsNicsInnerType", value)
}

// NewPlanSpecsNicsInnerTypeFromValue returns a pointer to a valid PlanSpecsNicsInnerType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPlanSpecsNicsInnerTypeFromValue(v string) (*PlanSpecsNicsInnerType, error) {
	ev := PlanSpecsNicsInnerType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PlanSpecsNicsInnerType: valid values are %v", v, AllowedPlanSpecsNicsInnerTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PlanSpecsNicsInnerType) IsValid() bool {
	for _, existing := range AllowedPlanSpecsNicsInnerTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Type value
func (v PlanSpecsNicsInnerType) Ptr() *PlanSpecsNicsInnerType {
	return &v
}

type NullablePlanSpecsNicsInnerType struct {
	value *PlanSpecsNicsInnerType
	isSet bool
}

func (v NullablePlanSpecsNicsInnerType) Get() *PlanSpecsNicsInnerType {
	return v.value
}

func (v *NullablePlanSpecsNicsInnerType) Set(val *PlanSpecsNicsInnerType) {
	v.value = val
	v.isSet = true
}

func (v NullablePlanSpecsNicsInnerType) IsSet() bool {
	return v.isSet
}

func (v *NullablePlanSpecsNicsInnerType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlanSpecsNicsInnerType(val *PlanSpecsNicsInnerType) *NullablePlanSpecsNicsInnerType {
	return &NullablePlanSpecsNicsInnerType{value: val, isSet: true}
}

func (v NullablePlanSpecsNicsInnerType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlanSpecsNicsInnerType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

// PlanSpecsNicsInner struct for PlanSpecsNicsInner
type PlanSpecsNicsInner struct {
	Count                *int32  `json:"count,omitempty"`
	Type                 *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PlanSpecsNicsInner PlanSpecsNicsInner

// NewPlanSpecsNicsInner instantiates a new PlanSpecsNicsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlanSpecsNicsInner() *PlanSpecsNicsInner {
	this := PlanSpecsNicsInner{}
	return &this
}

// NewPlanSpecsNicsInnerWithDefaults instantiates a new PlanSpecsNicsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlanSpecsNicsInnerWithDefaults() *PlanSpecsNicsInner {
	this := PlanSpecsNicsInner{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *PlanSpecsNicsInner) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlanSpecsNicsInner) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *PlanSpecsNicsInner) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *PlanSpecsNicsInner) SetCount(v int32) {
	o.Count = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PlanSpecsNicsInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlanSpecsNicsInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PlanSpecsNicsInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PlanSpecsNicsInner) SetType(v string) {
	o.Type = &v
}

func (o PlanSpecsNicsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlanSpecsNicsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PlanSpecsNicsInner) UnmarshalJSON(bytes []byte) (err error) {
	varPlanSpecsNicsInner := _PlanSpecsNicsInner{}

	err = json.Unmarshal(bytes, &varPlanSpecsNicsInner)

	if err != nil {
		return err
	}

	*o = PlanSpecsNicsInner(varPlanSpecsNicsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "count")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePlanSpecsNicsInner struct {
	value *PlanSpecsNicsInner
	isSet bool
}

func (v NullablePlanSpecsNicsInner) Get() *PlanSpecsNicsInner {
	return v.value
}

func (v *NullablePlanSpecsNicsInner) Set(val *PlanSpecsNicsInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePlanSpecsNicsInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePlanSpecsNicsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlanSpecsNicsInner(val *PlanSpecsNicsInner) *NullablePlanSpecsNicsInner {
	return &NullablePlanSpecsNicsInner{value: val, isSet: true}
}

func (v NullablePlanSpecsNicsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlanSpecsNicsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
