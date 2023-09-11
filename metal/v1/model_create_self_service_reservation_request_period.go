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

// checks if the CreateSelfServiceReservationRequestPeriod type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSelfServiceReservationRequestPeriod{}

// CreateSelfServiceReservationRequestPeriodCount the model 'CreateSelfServiceReservationRequestPeriodCount'
type CreateSelfServiceReservationRequestPeriodCount int32

// List of CreateSelfServiceReservationRequestPeriodCount
const (
	CREATESELFSERVICERESERVATIONREQUESTPERIOD__12 CreateSelfServiceReservationRequestPeriodCount = 12
	CREATESELFSERVICERESERVATIONREQUESTPERIOD__36 CreateSelfServiceReservationRequestPeriodCount = 36
)

// All allowed values of CreateSelfServiceReservationRequestPeriodCount enum
var AllowedCreateSelfServiceReservationRequestPeriodCountEnumValues = []CreateSelfServiceReservationRequestPeriodCount{
	12,
	36,
}

func (v *CreateSelfServiceReservationRequestPeriodCount) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CreateSelfServiceReservationRequestPeriodCount(value)
	for _, existing := range AllowedCreateSelfServiceReservationRequestPeriodCountEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CreateSelfServiceReservationRequestPeriodCount", value)
}

// NewCreateSelfServiceReservationRequestPeriodCountFromValue returns a pointer to a valid CreateSelfServiceReservationRequestPeriodCount
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreateSelfServiceReservationRequestPeriodCountFromValue(v int32) (*CreateSelfServiceReservationRequestPeriodCount, error) {
	ev := CreateSelfServiceReservationRequestPeriodCount(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreateSelfServiceReservationRequestPeriodCount: valid values are %v", v, AllowedCreateSelfServiceReservationRequestPeriodCountEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreateSelfServiceReservationRequestPeriodCount) IsValid() bool {
	for _, existing := range AllowedCreateSelfServiceReservationRequestPeriodCountEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Count value
func (v CreateSelfServiceReservationRequestPeriodCount) Ptr() *CreateSelfServiceReservationRequestPeriodCount {
	return &v
}

type NullableCreateSelfServiceReservationRequestPeriodCount struct {
	value *CreateSelfServiceReservationRequestPeriodCount
	isSet bool
}

func (v NullableCreateSelfServiceReservationRequestPeriodCount) Get() *CreateSelfServiceReservationRequestPeriodCount {
	return v.value
}

func (v *NullableCreateSelfServiceReservationRequestPeriodCount) Set(val *CreateSelfServiceReservationRequestPeriodCount) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSelfServiceReservationRequestPeriodCount) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSelfServiceReservationRequestPeriodCount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSelfServiceReservationRequestPeriodCount(val *CreateSelfServiceReservationRequestPeriodCount) *NullableCreateSelfServiceReservationRequestPeriodCount {
	return &NullableCreateSelfServiceReservationRequestPeriodCount{value: val, isSet: true}
}

func (v NullableCreateSelfServiceReservationRequestPeriodCount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSelfServiceReservationRequestPeriodCount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

// CreateSelfServiceReservationRequestPeriodUnit the model 'CreateSelfServiceReservationRequestPeriodUnit'
type CreateSelfServiceReservationRequestPeriodUnit string

// List of CreateSelfServiceReservationRequestPeriodUnit
const (
	CREATESELFSERVICERESERVATIONREQUESTPERIOD_MONTHLY CreateSelfServiceReservationRequestPeriodUnit = "monthly"
)

// All allowed values of CreateSelfServiceReservationRequestPeriodUnit enum
var AllowedCreateSelfServiceReservationRequestPeriodUnitEnumValues = []CreateSelfServiceReservationRequestPeriodUnit{
	"monthly",
}

func (v *CreateSelfServiceReservationRequestPeriodUnit) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CreateSelfServiceReservationRequestPeriodUnit(value)
	for _, existing := range AllowedCreateSelfServiceReservationRequestPeriodUnitEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CreateSelfServiceReservationRequestPeriodUnit", value)
}

// NewCreateSelfServiceReservationRequestPeriodUnitFromValue returns a pointer to a valid CreateSelfServiceReservationRequestPeriodUnit
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreateSelfServiceReservationRequestPeriodUnitFromValue(v string) (*CreateSelfServiceReservationRequestPeriodUnit, error) {
	ev := CreateSelfServiceReservationRequestPeriodUnit(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreateSelfServiceReservationRequestPeriodUnit: valid values are %v", v, AllowedCreateSelfServiceReservationRequestPeriodUnitEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreateSelfServiceReservationRequestPeriodUnit) IsValid() bool {
	for _, existing := range AllowedCreateSelfServiceReservationRequestPeriodUnitEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Unit value
func (v CreateSelfServiceReservationRequestPeriodUnit) Ptr() *CreateSelfServiceReservationRequestPeriodUnit {
	return &v
}

type NullableCreateSelfServiceReservationRequestPeriodUnit struct {
	value *CreateSelfServiceReservationRequestPeriodUnit
	isSet bool
}

func (v NullableCreateSelfServiceReservationRequestPeriodUnit) Get() *CreateSelfServiceReservationRequestPeriodUnit {
	return v.value
}

func (v *NullableCreateSelfServiceReservationRequestPeriodUnit) Set(val *CreateSelfServiceReservationRequestPeriodUnit) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSelfServiceReservationRequestPeriodUnit) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSelfServiceReservationRequestPeriodUnit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSelfServiceReservationRequestPeriodUnit(val *CreateSelfServiceReservationRequestPeriodUnit) *NullableCreateSelfServiceReservationRequestPeriodUnit {
	return &NullableCreateSelfServiceReservationRequestPeriodUnit{value: val, isSet: true}
}

func (v NullableCreateSelfServiceReservationRequestPeriodUnit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSelfServiceReservationRequestPeriodUnit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

// CreateSelfServiceReservationRequestPeriod struct for CreateSelfServiceReservationRequestPeriod
type CreateSelfServiceReservationRequestPeriod struct {
	Count                *int32  `json:"count,omitempty"`
	Unit                 *string `json:"unit,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateSelfServiceReservationRequestPeriod CreateSelfServiceReservationRequestPeriod

// NewCreateSelfServiceReservationRequestPeriod instantiates a new CreateSelfServiceReservationRequestPeriod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSelfServiceReservationRequestPeriod() *CreateSelfServiceReservationRequestPeriod {
	this := CreateSelfServiceReservationRequestPeriod{}
	return &this
}

// NewCreateSelfServiceReservationRequestPeriodWithDefaults instantiates a new CreateSelfServiceReservationRequestPeriod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSelfServiceReservationRequestPeriodWithDefaults() *CreateSelfServiceReservationRequestPeriod {
	this := CreateSelfServiceReservationRequestPeriod{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *CreateSelfServiceReservationRequestPeriod) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSelfServiceReservationRequestPeriod) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *CreateSelfServiceReservationRequestPeriod) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *CreateSelfServiceReservationRequestPeriod) SetCount(v int32) {
	o.Count = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *CreateSelfServiceReservationRequestPeriod) GetUnit() string {
	if o == nil || IsNil(o.Unit) {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSelfServiceReservationRequestPeriod) GetUnitOk() (*string, bool) {
	if o == nil || IsNil(o.Unit) {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *CreateSelfServiceReservationRequestPeriod) HasUnit() bool {
	if o != nil && !IsNil(o.Unit) {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *CreateSelfServiceReservationRequestPeriod) SetUnit(v string) {
	o.Unit = &v
}

func (o CreateSelfServiceReservationRequestPeriod) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateSelfServiceReservationRequestPeriod) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Unit) {
		toSerialize["unit"] = o.Unit
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateSelfServiceReservationRequestPeriod) UnmarshalJSON(bytes []byte) (err error) {
	varCreateSelfServiceReservationRequestPeriod := _CreateSelfServiceReservationRequestPeriod{}

	err = json.Unmarshal(bytes, &varCreateSelfServiceReservationRequestPeriod)

	if err != nil {
		return err
	}

	*o = CreateSelfServiceReservationRequestPeriod(varCreateSelfServiceReservationRequestPeriod)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "count")
		delete(additionalProperties, "unit")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateSelfServiceReservationRequestPeriod struct {
	value *CreateSelfServiceReservationRequestPeriod
	isSet bool
}

func (v NullableCreateSelfServiceReservationRequestPeriod) Get() *CreateSelfServiceReservationRequestPeriod {
	return v.value
}

func (v *NullableCreateSelfServiceReservationRequestPeriod) Set(val *CreateSelfServiceReservationRequestPeriod) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSelfServiceReservationRequestPeriod) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSelfServiceReservationRequestPeriod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSelfServiceReservationRequestPeriod(val *CreateSelfServiceReservationRequestPeriod) *NullableCreateSelfServiceReservationRequestPeriod {
	return &NullableCreateSelfServiceReservationRequestPeriod{value: val, isSet: true}
}

func (v NullableCreateSelfServiceReservationRequestPeriod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSelfServiceReservationRequestPeriod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
