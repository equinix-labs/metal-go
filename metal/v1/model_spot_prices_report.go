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

// checks if the SpotPricesReport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpotPricesReport{}

// SpotPricesReport struct for SpotPricesReport
type SpotPricesReport struct {
	Ams1                 *SpotPricesPerFacility    `json:"ams1,omitempty"`
	Atl1                 *SpotPricesPerNewFacility `json:"atl1,omitempty"`
	Dfw1                 *SpotPricesPerNewFacility `json:"dfw1,omitempty"`
	Ewr1                 *SpotPricesPerFacility    `json:"ewr1,omitempty"`
	Fra1                 *SpotPricesPerNewFacility `json:"fra1,omitempty"`
	Iad1                 *SpotPricesPerNewFacility `json:"iad1,omitempty"`
	Lax1                 *SpotPricesPerNewFacility `json:"lax1,omitempty"`
	Nrt1                 *SpotPricesPerFacility    `json:"nrt1,omitempty"`
	Ord1                 *SpotPricesPerNewFacility `json:"ord1,omitempty"`
	Sea1                 *SpotPricesPerNewFacility `json:"sea1,omitempty"`
	Sin1                 *SpotPricesPerNewFacility `json:"sin1,omitempty"`
	Sjc1                 *SpotPricesPerFacility    `json:"sjc1,omitempty"`
	Syd1                 *SpotPricesPerNewFacility `json:"syd1,omitempty"`
	Yyz1                 *SpotPricesPerNewFacility `json:"yyz1,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SpotPricesReport SpotPricesReport

// NewSpotPricesReport instantiates a new SpotPricesReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpotPricesReport() *SpotPricesReport {
	this := SpotPricesReport{}
	return &this
}

// NewSpotPricesReportWithDefaults instantiates a new SpotPricesReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpotPricesReportWithDefaults() *SpotPricesReport {
	this := SpotPricesReport{}
	return &this
}

// GetAms1 returns the Ams1 field value if set, zero value otherwise.
func (o *SpotPricesReport) GetAms1() SpotPricesPerFacility {
	if o == nil || IsNil(o.Ams1) {
		var ret SpotPricesPerFacility
		return ret
	}
	return *o.Ams1
}

// GetAms1Ok returns a tuple with the Ams1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotPricesReport) GetAms1Ok() (*SpotPricesPerFacility, bool) {
	if o == nil || IsNil(o.Ams1) {
		return nil, false
	}
	return o.Ams1, true
}

// HasAms1 returns a boolean if a field has been set.
func (o *SpotPricesReport) HasAms1() bool {
	if o != nil && !IsNil(o.Ams1) {
		return true
	}

	return false
}

// SetAms1 gets a reference to the given SpotPricesPerFacility and assigns it to the Ams1 field.
func (o *SpotPricesReport) SetAms1(v SpotPricesPerFacility) {
	o.Ams1 = &v
}

// GetAtl1 returns the Atl1 field value if set, zero value otherwise.
func (o *SpotPricesReport) GetAtl1() SpotPricesPerNewFacility {
	if o == nil || IsNil(o.Atl1) {
		var ret SpotPricesPerNewFacility
		return ret
	}
	return *o.Atl1
}

// GetAtl1Ok returns a tuple with the Atl1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotPricesReport) GetAtl1Ok() (*SpotPricesPerNewFacility, bool) {
	if o == nil || IsNil(o.Atl1) {
		return nil, false
	}
	return o.Atl1, true
}

// HasAtl1 returns a boolean if a field has been set.
func (o *SpotPricesReport) HasAtl1() bool {
	if o != nil && !IsNil(o.Atl1) {
		return true
	}

	return false
}

// SetAtl1 gets a reference to the given SpotPricesPerNewFacility and assigns it to the Atl1 field.
func (o *SpotPricesReport) SetAtl1(v SpotPricesPerNewFacility) {
	o.Atl1 = &v
}

// GetDfw1 returns the Dfw1 field value if set, zero value otherwise.
func (o *SpotPricesReport) GetDfw1() SpotPricesPerNewFacility {
	if o == nil || IsNil(o.Dfw1) {
		var ret SpotPricesPerNewFacility
		return ret
	}
	return *o.Dfw1
}

// GetDfw1Ok returns a tuple with the Dfw1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotPricesReport) GetDfw1Ok() (*SpotPricesPerNewFacility, bool) {
	if o == nil || IsNil(o.Dfw1) {
		return nil, false
	}
	return o.Dfw1, true
}

// HasDfw1 returns a boolean if a field has been set.
func (o *SpotPricesReport) HasDfw1() bool {
	if o != nil && !IsNil(o.Dfw1) {
		return true
	}

	return false
}

// SetDfw1 gets a reference to the given SpotPricesPerNewFacility and assigns it to the Dfw1 field.
func (o *SpotPricesReport) SetDfw1(v SpotPricesPerNewFacility) {
	o.Dfw1 = &v
}

// GetEwr1 returns the Ewr1 field value if set, zero value otherwise.
func (o *SpotPricesReport) GetEwr1() SpotPricesPerFacility {
	if o == nil || IsNil(o.Ewr1) {
		var ret SpotPricesPerFacility
		return ret
	}
	return *o.Ewr1
}

// GetEwr1Ok returns a tuple with the Ewr1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotPricesReport) GetEwr1Ok() (*SpotPricesPerFacility, bool) {
	if o == nil || IsNil(o.Ewr1) {
		return nil, false
	}
	return o.Ewr1, true
}

// HasEwr1 returns a boolean if a field has been set.
func (o *SpotPricesReport) HasEwr1() bool {
	if o != nil && !IsNil(o.Ewr1) {
		return true
	}

	return false
}

// SetEwr1 gets a reference to the given SpotPricesPerFacility and assigns it to the Ewr1 field.
func (o *SpotPricesReport) SetEwr1(v SpotPricesPerFacility) {
	o.Ewr1 = &v
}

// GetFra1 returns the Fra1 field value if set, zero value otherwise.
func (o *SpotPricesReport) GetFra1() SpotPricesPerNewFacility {
	if o == nil || IsNil(o.Fra1) {
		var ret SpotPricesPerNewFacility
		return ret
	}
	return *o.Fra1
}

// GetFra1Ok returns a tuple with the Fra1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotPricesReport) GetFra1Ok() (*SpotPricesPerNewFacility, bool) {
	if o == nil || IsNil(o.Fra1) {
		return nil, false
	}
	return o.Fra1, true
}

// HasFra1 returns a boolean if a field has been set.
func (o *SpotPricesReport) HasFra1() bool {
	if o != nil && !IsNil(o.Fra1) {
		return true
	}

	return false
}

// SetFra1 gets a reference to the given SpotPricesPerNewFacility and assigns it to the Fra1 field.
func (o *SpotPricesReport) SetFra1(v SpotPricesPerNewFacility) {
	o.Fra1 = &v
}

// GetIad1 returns the Iad1 field value if set, zero value otherwise.
func (o *SpotPricesReport) GetIad1() SpotPricesPerNewFacility {
	if o == nil || IsNil(o.Iad1) {
		var ret SpotPricesPerNewFacility
		return ret
	}
	return *o.Iad1
}

// GetIad1Ok returns a tuple with the Iad1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotPricesReport) GetIad1Ok() (*SpotPricesPerNewFacility, bool) {
	if o == nil || IsNil(o.Iad1) {
		return nil, false
	}
	return o.Iad1, true
}

// HasIad1 returns a boolean if a field has been set.
func (o *SpotPricesReport) HasIad1() bool {
	if o != nil && !IsNil(o.Iad1) {
		return true
	}

	return false
}

// SetIad1 gets a reference to the given SpotPricesPerNewFacility and assigns it to the Iad1 field.
func (o *SpotPricesReport) SetIad1(v SpotPricesPerNewFacility) {
	o.Iad1 = &v
}

// GetLax1 returns the Lax1 field value if set, zero value otherwise.
func (o *SpotPricesReport) GetLax1() SpotPricesPerNewFacility {
	if o == nil || IsNil(o.Lax1) {
		var ret SpotPricesPerNewFacility
		return ret
	}
	return *o.Lax1
}

// GetLax1Ok returns a tuple with the Lax1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotPricesReport) GetLax1Ok() (*SpotPricesPerNewFacility, bool) {
	if o == nil || IsNil(o.Lax1) {
		return nil, false
	}
	return o.Lax1, true
}

// HasLax1 returns a boolean if a field has been set.
func (o *SpotPricesReport) HasLax1() bool {
	if o != nil && !IsNil(o.Lax1) {
		return true
	}

	return false
}

// SetLax1 gets a reference to the given SpotPricesPerNewFacility and assigns it to the Lax1 field.
func (o *SpotPricesReport) SetLax1(v SpotPricesPerNewFacility) {
	o.Lax1 = &v
}

// GetNrt1 returns the Nrt1 field value if set, zero value otherwise.
func (o *SpotPricesReport) GetNrt1() SpotPricesPerFacility {
	if o == nil || IsNil(o.Nrt1) {
		var ret SpotPricesPerFacility
		return ret
	}
	return *o.Nrt1
}

// GetNrt1Ok returns a tuple with the Nrt1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotPricesReport) GetNrt1Ok() (*SpotPricesPerFacility, bool) {
	if o == nil || IsNil(o.Nrt1) {
		return nil, false
	}
	return o.Nrt1, true
}

// HasNrt1 returns a boolean if a field has been set.
func (o *SpotPricesReport) HasNrt1() bool {
	if o != nil && !IsNil(o.Nrt1) {
		return true
	}

	return false
}

// SetNrt1 gets a reference to the given SpotPricesPerFacility and assigns it to the Nrt1 field.
func (o *SpotPricesReport) SetNrt1(v SpotPricesPerFacility) {
	o.Nrt1 = &v
}

// GetOrd1 returns the Ord1 field value if set, zero value otherwise.
func (o *SpotPricesReport) GetOrd1() SpotPricesPerNewFacility {
	if o == nil || IsNil(o.Ord1) {
		var ret SpotPricesPerNewFacility
		return ret
	}
	return *o.Ord1
}

// GetOrd1Ok returns a tuple with the Ord1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotPricesReport) GetOrd1Ok() (*SpotPricesPerNewFacility, bool) {
	if o == nil || IsNil(o.Ord1) {
		return nil, false
	}
	return o.Ord1, true
}

// HasOrd1 returns a boolean if a field has been set.
func (o *SpotPricesReport) HasOrd1() bool {
	if o != nil && !IsNil(o.Ord1) {
		return true
	}

	return false
}

// SetOrd1 gets a reference to the given SpotPricesPerNewFacility and assigns it to the Ord1 field.
func (o *SpotPricesReport) SetOrd1(v SpotPricesPerNewFacility) {
	o.Ord1 = &v
}

// GetSea1 returns the Sea1 field value if set, zero value otherwise.
func (o *SpotPricesReport) GetSea1() SpotPricesPerNewFacility {
	if o == nil || IsNil(o.Sea1) {
		var ret SpotPricesPerNewFacility
		return ret
	}
	return *o.Sea1
}

// GetSea1Ok returns a tuple with the Sea1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotPricesReport) GetSea1Ok() (*SpotPricesPerNewFacility, bool) {
	if o == nil || IsNil(o.Sea1) {
		return nil, false
	}
	return o.Sea1, true
}

// HasSea1 returns a boolean if a field has been set.
func (o *SpotPricesReport) HasSea1() bool {
	if o != nil && !IsNil(o.Sea1) {
		return true
	}

	return false
}

// SetSea1 gets a reference to the given SpotPricesPerNewFacility and assigns it to the Sea1 field.
func (o *SpotPricesReport) SetSea1(v SpotPricesPerNewFacility) {
	o.Sea1 = &v
}

// GetSin1 returns the Sin1 field value if set, zero value otherwise.
func (o *SpotPricesReport) GetSin1() SpotPricesPerNewFacility {
	if o == nil || IsNil(o.Sin1) {
		var ret SpotPricesPerNewFacility
		return ret
	}
	return *o.Sin1
}

// GetSin1Ok returns a tuple with the Sin1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotPricesReport) GetSin1Ok() (*SpotPricesPerNewFacility, bool) {
	if o == nil || IsNil(o.Sin1) {
		return nil, false
	}
	return o.Sin1, true
}

// HasSin1 returns a boolean if a field has been set.
func (o *SpotPricesReport) HasSin1() bool {
	if o != nil && !IsNil(o.Sin1) {
		return true
	}

	return false
}

// SetSin1 gets a reference to the given SpotPricesPerNewFacility and assigns it to the Sin1 field.
func (o *SpotPricesReport) SetSin1(v SpotPricesPerNewFacility) {
	o.Sin1 = &v
}

// GetSjc1 returns the Sjc1 field value if set, zero value otherwise.
func (o *SpotPricesReport) GetSjc1() SpotPricesPerFacility {
	if o == nil || IsNil(o.Sjc1) {
		var ret SpotPricesPerFacility
		return ret
	}
	return *o.Sjc1
}

// GetSjc1Ok returns a tuple with the Sjc1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotPricesReport) GetSjc1Ok() (*SpotPricesPerFacility, bool) {
	if o == nil || IsNil(o.Sjc1) {
		return nil, false
	}
	return o.Sjc1, true
}

// HasSjc1 returns a boolean if a field has been set.
func (o *SpotPricesReport) HasSjc1() bool {
	if o != nil && !IsNil(o.Sjc1) {
		return true
	}

	return false
}

// SetSjc1 gets a reference to the given SpotPricesPerFacility and assigns it to the Sjc1 field.
func (o *SpotPricesReport) SetSjc1(v SpotPricesPerFacility) {
	o.Sjc1 = &v
}

// GetSyd1 returns the Syd1 field value if set, zero value otherwise.
func (o *SpotPricesReport) GetSyd1() SpotPricesPerNewFacility {
	if o == nil || IsNil(o.Syd1) {
		var ret SpotPricesPerNewFacility
		return ret
	}
	return *o.Syd1
}

// GetSyd1Ok returns a tuple with the Syd1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotPricesReport) GetSyd1Ok() (*SpotPricesPerNewFacility, bool) {
	if o == nil || IsNil(o.Syd1) {
		return nil, false
	}
	return o.Syd1, true
}

// HasSyd1 returns a boolean if a field has been set.
func (o *SpotPricesReport) HasSyd1() bool {
	if o != nil && !IsNil(o.Syd1) {
		return true
	}

	return false
}

// SetSyd1 gets a reference to the given SpotPricesPerNewFacility and assigns it to the Syd1 field.
func (o *SpotPricesReport) SetSyd1(v SpotPricesPerNewFacility) {
	o.Syd1 = &v
}

// GetYyz1 returns the Yyz1 field value if set, zero value otherwise.
func (o *SpotPricesReport) GetYyz1() SpotPricesPerNewFacility {
	if o == nil || IsNil(o.Yyz1) {
		var ret SpotPricesPerNewFacility
		return ret
	}
	return *o.Yyz1
}

// GetYyz1Ok returns a tuple with the Yyz1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotPricesReport) GetYyz1Ok() (*SpotPricesPerNewFacility, bool) {
	if o == nil || IsNil(o.Yyz1) {
		return nil, false
	}
	return o.Yyz1, true
}

// HasYyz1 returns a boolean if a field has been set.
func (o *SpotPricesReport) HasYyz1() bool {
	if o != nil && !IsNil(o.Yyz1) {
		return true
	}

	return false
}

// SetYyz1 gets a reference to the given SpotPricesPerNewFacility and assigns it to the Yyz1 field.
func (o *SpotPricesReport) SetYyz1(v SpotPricesPerNewFacility) {
	o.Yyz1 = &v
}

func (o SpotPricesReport) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpotPricesReport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ams1) {
		toSerialize["ams1"] = o.Ams1
	}
	if !IsNil(o.Atl1) {
		toSerialize["atl1"] = o.Atl1
	}
	if !IsNil(o.Dfw1) {
		toSerialize["dfw1"] = o.Dfw1
	}
	if !IsNil(o.Ewr1) {
		toSerialize["ewr1"] = o.Ewr1
	}
	if !IsNil(o.Fra1) {
		toSerialize["fra1"] = o.Fra1
	}
	if !IsNil(o.Iad1) {
		toSerialize["iad1"] = o.Iad1
	}
	if !IsNil(o.Lax1) {
		toSerialize["lax1"] = o.Lax1
	}
	if !IsNil(o.Nrt1) {
		toSerialize["nrt1"] = o.Nrt1
	}
	if !IsNil(o.Ord1) {
		toSerialize["ord1"] = o.Ord1
	}
	if !IsNil(o.Sea1) {
		toSerialize["sea1"] = o.Sea1
	}
	if !IsNil(o.Sin1) {
		toSerialize["sin1"] = o.Sin1
	}
	if !IsNil(o.Sjc1) {
		toSerialize["sjc1"] = o.Sjc1
	}
	if !IsNil(o.Syd1) {
		toSerialize["syd1"] = o.Syd1
	}
	if !IsNil(o.Yyz1) {
		toSerialize["yyz1"] = o.Yyz1
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SpotPricesReport) UnmarshalJSON(bytes []byte) (err error) {
	varSpotPricesReport := _SpotPricesReport{}

	err = json.Unmarshal(bytes, &varSpotPricesReport)

	if err != nil {
		return err
	}

	*o = SpotPricesReport(varSpotPricesReport)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ams1")
		delete(additionalProperties, "atl1")
		delete(additionalProperties, "dfw1")
		delete(additionalProperties, "ewr1")
		delete(additionalProperties, "fra1")
		delete(additionalProperties, "iad1")
		delete(additionalProperties, "lax1")
		delete(additionalProperties, "nrt1")
		delete(additionalProperties, "ord1")
		delete(additionalProperties, "sea1")
		delete(additionalProperties, "sin1")
		delete(additionalProperties, "sjc1")
		delete(additionalProperties, "syd1")
		delete(additionalProperties, "yyz1")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSpotPricesReport struct {
	value *SpotPricesReport
	isSet bool
}

func (v NullableSpotPricesReport) Get() *SpotPricesReport {
	return v.value
}

func (v *NullableSpotPricesReport) Set(val *SpotPricesReport) {
	v.value = val
	v.isSet = true
}

func (v NullableSpotPricesReport) IsSet() bool {
	return v.isSet
}

func (v *NullableSpotPricesReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpotPricesReport(val *SpotPricesReport) *NullableSpotPricesReport {
	return &NullableSpotPricesReport{value: val, isSet: true}
}

func (v NullableSpotPricesReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpotPricesReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
