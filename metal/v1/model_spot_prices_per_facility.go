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

// checks if the SpotPricesPerFacility type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpotPricesPerFacility{}

// SpotPricesPerFacility struct for SpotPricesPerFacility
type SpotPricesPerFacility struct {
	Baremetal0           *SpotPricesPerBaremetal `json:"baremetal_0,omitempty"`
	Baremetal1           *SpotPricesPerBaremetal `json:"baremetal_1,omitempty"`
	Baremetal2           *SpotPricesPerBaremetal `json:"baremetal_2,omitempty"`
	Baremetal2a          *SpotPricesPerBaremetal `json:"baremetal_2a,omitempty"`
	Baremetal2a2         *SpotPricesPerBaremetal `json:"baremetal_2a2,omitempty"`
	Baremetal3           *SpotPricesPerBaremetal `json:"baremetal_3,omitempty"`
	BaremetalS           *SpotPricesPerBaremetal `json:"baremetal_s,omitempty"`
	C2MediumX86          *SpotPricesPerBaremetal `json:"c2.medium.x86,omitempty"`
	M2XlargeX86          *SpotPricesPerBaremetal `json:"m2.xlarge.x86,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SpotPricesPerFacility SpotPricesPerFacility

// NewSpotPricesPerFacility instantiates a new SpotPricesPerFacility object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpotPricesPerFacility() *SpotPricesPerFacility {
	this := SpotPricesPerFacility{}
	return &this
}

// NewSpotPricesPerFacilityWithDefaults instantiates a new SpotPricesPerFacility object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpotPricesPerFacilityWithDefaults() *SpotPricesPerFacility {
	this := SpotPricesPerFacility{}
	return &this
}

// GetBaremetal0 returns the Baremetal0 field value if set, zero value otherwise.
func (o *SpotPricesPerFacility) GetBaremetal0() SpotPricesPerBaremetal {
	if o == nil || IsNil(o.Baremetal0) {
		var ret SpotPricesPerBaremetal
		return ret
	}
	return *o.Baremetal0
}

// GetBaremetal0Ok returns a tuple with the Baremetal0 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotPricesPerFacility) GetBaremetal0Ok() (*SpotPricesPerBaremetal, bool) {
	if o == nil || IsNil(o.Baremetal0) {
		return nil, false
	}
	return o.Baremetal0, true
}

// HasBaremetal0 returns a boolean if a field has been set.
func (o *SpotPricesPerFacility) HasBaremetal0() bool {
	if o != nil && !IsNil(o.Baremetal0) {
		return true
	}

	return false
}

// SetBaremetal0 gets a reference to the given SpotPricesPerBaremetal and assigns it to the Baremetal0 field.
func (o *SpotPricesPerFacility) SetBaremetal0(v SpotPricesPerBaremetal) {
	o.Baremetal0 = &v
}

// GetBaremetal1 returns the Baremetal1 field value if set, zero value otherwise.
func (o *SpotPricesPerFacility) GetBaremetal1() SpotPricesPerBaremetal {
	if o == nil || IsNil(o.Baremetal1) {
		var ret SpotPricesPerBaremetal
		return ret
	}
	return *o.Baremetal1
}

// GetBaremetal1Ok returns a tuple with the Baremetal1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotPricesPerFacility) GetBaremetal1Ok() (*SpotPricesPerBaremetal, bool) {
	if o == nil || IsNil(o.Baremetal1) {
		return nil, false
	}
	return o.Baremetal1, true
}

// HasBaremetal1 returns a boolean if a field has been set.
func (o *SpotPricesPerFacility) HasBaremetal1() bool {
	if o != nil && !IsNil(o.Baremetal1) {
		return true
	}

	return false
}

// SetBaremetal1 gets a reference to the given SpotPricesPerBaremetal and assigns it to the Baremetal1 field.
func (o *SpotPricesPerFacility) SetBaremetal1(v SpotPricesPerBaremetal) {
	o.Baremetal1 = &v
}

// GetBaremetal2 returns the Baremetal2 field value if set, zero value otherwise.
func (o *SpotPricesPerFacility) GetBaremetal2() SpotPricesPerBaremetal {
	if o == nil || IsNil(o.Baremetal2) {
		var ret SpotPricesPerBaremetal
		return ret
	}
	return *o.Baremetal2
}

// GetBaremetal2Ok returns a tuple with the Baremetal2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotPricesPerFacility) GetBaremetal2Ok() (*SpotPricesPerBaremetal, bool) {
	if o == nil || IsNil(o.Baremetal2) {
		return nil, false
	}
	return o.Baremetal2, true
}

// HasBaremetal2 returns a boolean if a field has been set.
func (o *SpotPricesPerFacility) HasBaremetal2() bool {
	if o != nil && !IsNil(o.Baremetal2) {
		return true
	}

	return false
}

// SetBaremetal2 gets a reference to the given SpotPricesPerBaremetal and assigns it to the Baremetal2 field.
func (o *SpotPricesPerFacility) SetBaremetal2(v SpotPricesPerBaremetal) {
	o.Baremetal2 = &v
}

// GetBaremetal2a returns the Baremetal2a field value if set, zero value otherwise.
func (o *SpotPricesPerFacility) GetBaremetal2a() SpotPricesPerBaremetal {
	if o == nil || IsNil(o.Baremetal2a) {
		var ret SpotPricesPerBaremetal
		return ret
	}
	return *o.Baremetal2a
}

// GetBaremetal2aOk returns a tuple with the Baremetal2a field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotPricesPerFacility) GetBaremetal2aOk() (*SpotPricesPerBaremetal, bool) {
	if o == nil || IsNil(o.Baremetal2a) {
		return nil, false
	}
	return o.Baremetal2a, true
}

// HasBaremetal2a returns a boolean if a field has been set.
func (o *SpotPricesPerFacility) HasBaremetal2a() bool {
	if o != nil && !IsNil(o.Baremetal2a) {
		return true
	}

	return false
}

// SetBaremetal2a gets a reference to the given SpotPricesPerBaremetal and assigns it to the Baremetal2a field.
func (o *SpotPricesPerFacility) SetBaremetal2a(v SpotPricesPerBaremetal) {
	o.Baremetal2a = &v
}

// GetBaremetal2a2 returns the Baremetal2a2 field value if set, zero value otherwise.
func (o *SpotPricesPerFacility) GetBaremetal2a2() SpotPricesPerBaremetal {
	if o == nil || IsNil(o.Baremetal2a2) {
		var ret SpotPricesPerBaremetal
		return ret
	}
	return *o.Baremetal2a2
}

// GetBaremetal2a2Ok returns a tuple with the Baremetal2a2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotPricesPerFacility) GetBaremetal2a2Ok() (*SpotPricesPerBaremetal, bool) {
	if o == nil || IsNil(o.Baremetal2a2) {
		return nil, false
	}
	return o.Baremetal2a2, true
}

// HasBaremetal2a2 returns a boolean if a field has been set.
func (o *SpotPricesPerFacility) HasBaremetal2a2() bool {
	if o != nil && !IsNil(o.Baremetal2a2) {
		return true
	}

	return false
}

// SetBaremetal2a2 gets a reference to the given SpotPricesPerBaremetal and assigns it to the Baremetal2a2 field.
func (o *SpotPricesPerFacility) SetBaremetal2a2(v SpotPricesPerBaremetal) {
	o.Baremetal2a2 = &v
}

// GetBaremetal3 returns the Baremetal3 field value if set, zero value otherwise.
func (o *SpotPricesPerFacility) GetBaremetal3() SpotPricesPerBaremetal {
	if o == nil || IsNil(o.Baremetal3) {
		var ret SpotPricesPerBaremetal
		return ret
	}
	return *o.Baremetal3
}

// GetBaremetal3Ok returns a tuple with the Baremetal3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotPricesPerFacility) GetBaremetal3Ok() (*SpotPricesPerBaremetal, bool) {
	if o == nil || IsNil(o.Baremetal3) {
		return nil, false
	}
	return o.Baremetal3, true
}

// HasBaremetal3 returns a boolean if a field has been set.
func (o *SpotPricesPerFacility) HasBaremetal3() bool {
	if o != nil && !IsNil(o.Baremetal3) {
		return true
	}

	return false
}

// SetBaremetal3 gets a reference to the given SpotPricesPerBaremetal and assigns it to the Baremetal3 field.
func (o *SpotPricesPerFacility) SetBaremetal3(v SpotPricesPerBaremetal) {
	o.Baremetal3 = &v
}

// GetBaremetalS returns the BaremetalS field value if set, zero value otherwise.
func (o *SpotPricesPerFacility) GetBaremetalS() SpotPricesPerBaremetal {
	if o == nil || IsNil(o.BaremetalS) {
		var ret SpotPricesPerBaremetal
		return ret
	}
	return *o.BaremetalS
}

// GetBaremetalSOk returns a tuple with the BaremetalS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotPricesPerFacility) GetBaremetalSOk() (*SpotPricesPerBaremetal, bool) {
	if o == nil || IsNil(o.BaremetalS) {
		return nil, false
	}
	return o.BaremetalS, true
}

// HasBaremetalS returns a boolean if a field has been set.
func (o *SpotPricesPerFacility) HasBaremetalS() bool {
	if o != nil && !IsNil(o.BaremetalS) {
		return true
	}

	return false
}

// SetBaremetalS gets a reference to the given SpotPricesPerBaremetal and assigns it to the BaremetalS field.
func (o *SpotPricesPerFacility) SetBaremetalS(v SpotPricesPerBaremetal) {
	o.BaremetalS = &v
}

// GetC2MediumX86 returns the C2MediumX86 field value if set, zero value otherwise.
func (o *SpotPricesPerFacility) GetC2MediumX86() SpotPricesPerBaremetal {
	if o == nil || IsNil(o.C2MediumX86) {
		var ret SpotPricesPerBaremetal
		return ret
	}
	return *o.C2MediumX86
}

// GetC2MediumX86Ok returns a tuple with the C2MediumX86 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotPricesPerFacility) GetC2MediumX86Ok() (*SpotPricesPerBaremetal, bool) {
	if o == nil || IsNil(o.C2MediumX86) {
		return nil, false
	}
	return o.C2MediumX86, true
}

// HasC2MediumX86 returns a boolean if a field has been set.
func (o *SpotPricesPerFacility) HasC2MediumX86() bool {
	if o != nil && !IsNil(o.C2MediumX86) {
		return true
	}

	return false
}

// SetC2MediumX86 gets a reference to the given SpotPricesPerBaremetal and assigns it to the C2MediumX86 field.
func (o *SpotPricesPerFacility) SetC2MediumX86(v SpotPricesPerBaremetal) {
	o.C2MediumX86 = &v
}

// GetM2XlargeX86 returns the M2XlargeX86 field value if set, zero value otherwise.
func (o *SpotPricesPerFacility) GetM2XlargeX86() SpotPricesPerBaremetal {
	if o == nil || IsNil(o.M2XlargeX86) {
		var ret SpotPricesPerBaremetal
		return ret
	}
	return *o.M2XlargeX86
}

// GetM2XlargeX86Ok returns a tuple with the M2XlargeX86 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotPricesPerFacility) GetM2XlargeX86Ok() (*SpotPricesPerBaremetal, bool) {
	if o == nil || IsNil(o.M2XlargeX86) {
		return nil, false
	}
	return o.M2XlargeX86, true
}

// HasM2XlargeX86 returns a boolean if a field has been set.
func (o *SpotPricesPerFacility) HasM2XlargeX86() bool {
	if o != nil && !IsNil(o.M2XlargeX86) {
		return true
	}

	return false
}

// SetM2XlargeX86 gets a reference to the given SpotPricesPerBaremetal and assigns it to the M2XlargeX86 field.
func (o *SpotPricesPerFacility) SetM2XlargeX86(v SpotPricesPerBaremetal) {
	o.M2XlargeX86 = &v
}

func (o SpotPricesPerFacility) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpotPricesPerFacility) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Baremetal0) {
		toSerialize["baremetal_0"] = o.Baremetal0
	}
	if !IsNil(o.Baremetal1) {
		toSerialize["baremetal_1"] = o.Baremetal1
	}
	if !IsNil(o.Baremetal2) {
		toSerialize["baremetal_2"] = o.Baremetal2
	}
	if !IsNil(o.Baremetal2a) {
		toSerialize["baremetal_2a"] = o.Baremetal2a
	}
	if !IsNil(o.Baremetal2a2) {
		toSerialize["baremetal_2a2"] = o.Baremetal2a2
	}
	if !IsNil(o.Baremetal3) {
		toSerialize["baremetal_3"] = o.Baremetal3
	}
	if !IsNil(o.BaremetalS) {
		toSerialize["baremetal_s"] = o.BaremetalS
	}
	if !IsNil(o.C2MediumX86) {
		toSerialize["c2.medium.x86"] = o.C2MediumX86
	}
	if !IsNil(o.M2XlargeX86) {
		toSerialize["m2.xlarge.x86"] = o.M2XlargeX86
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SpotPricesPerFacility) UnmarshalJSON(bytes []byte) (err error) {
	varSpotPricesPerFacility := _SpotPricesPerFacility{}

	err = json.Unmarshal(bytes, &varSpotPricesPerFacility)

	if err != nil {
		return err
	}

	*o = SpotPricesPerFacility(varSpotPricesPerFacility)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "baremetal_0")
		delete(additionalProperties, "baremetal_1")
		delete(additionalProperties, "baremetal_2")
		delete(additionalProperties, "baremetal_2a")
		delete(additionalProperties, "baremetal_2a2")
		delete(additionalProperties, "baremetal_3")
		delete(additionalProperties, "baremetal_s")
		delete(additionalProperties, "c2.medium.x86")
		delete(additionalProperties, "m2.xlarge.x86")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSpotPricesPerFacility struct {
	value *SpotPricesPerFacility
	isSet bool
}

func (v NullableSpotPricesPerFacility) Get() *SpotPricesPerFacility {
	return v.value
}

func (v *NullableSpotPricesPerFacility) Set(val *SpotPricesPerFacility) {
	v.value = val
	v.isSet = true
}

func (v NullableSpotPricesPerFacility) IsSet() bool {
	return v.isSet
}

func (v *NullableSpotPricesPerFacility) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpotPricesPerFacility(val *SpotPricesPerFacility) *NullableSpotPricesPerFacility {
	return &NullableSpotPricesPerFacility{value: val, isSet: true}
}

func (v NullableSpotPricesPerFacility) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpotPricesPerFacility) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
