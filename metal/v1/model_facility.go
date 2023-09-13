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

// checks if the Facility type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Facility{}

// FacilityFeatures the model 'FacilityFeatures'
type FacilityFeatures string

// List of FacilityFeatures
const (
	FACILITY_BAREMETAL        FacilityFeatures = "baremetal"
	FACILITY_BACKEND_TRANSFER FacilityFeatures = "backend_transfer"
	FACILITY_LAYER_2          FacilityFeatures = "layer_2"
	FACILITY_GLOBAL_IPV4      FacilityFeatures = "global_ipv4"
	FACILITY_IBX              FacilityFeatures = "ibx"
)

// All allowed values of FacilityFeatures enum
var AllowedFacilityFeaturesEnumValues = []FacilityFeatures{
	"baremetal",
	"backend_transfer",
	"layer_2",
	"global_ipv4",
	"ibx",
}

func (v *FacilityFeatures) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FacilityFeatures(value)
	for _, existing := range AllowedFacilityFeaturesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FacilityFeatures", value)
}

// NewFacilityFeaturesFromValue returns a pointer to a valid FacilityFeatures
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFacilityFeaturesFromValue(v string) (*FacilityFeatures, error) {
	ev := FacilityFeatures(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FacilityFeatures: valid values are %v", v, AllowedFacilityFeaturesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FacilityFeatures) IsValid() bool {
	for _, existing := range AllowedFacilityFeaturesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Features value
func (v FacilityFeatures) Ptr() *FacilityFeatures {
	return &v
}

type NullableFacilityFeatures struct {
	value *FacilityFeatures
	isSet bool
}

func (v NullableFacilityFeatures) Get() *FacilityFeatures {
	return v.value
}

func (v *NullableFacilityFeatures) Set(val *FacilityFeatures) {
	v.value = val
	v.isSet = true
}

func (v NullableFacilityFeatures) IsSet() bool {
	return v.isSet
}

func (v *NullableFacilityFeatures) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFacilityFeatures(val *FacilityFeatures) *NullableFacilityFeatures {
	return &NullableFacilityFeatures{value: val, isSet: true}
}

func (v NullableFacilityFeatures) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFacilityFeatures) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

// Facility struct for Facility
type Facility struct {
	Address  *Address         `json:"address,omitempty"`
	Code     *string          `json:"code,omitempty"`
	Features FacilityFeatures `json:"features,omitempty"`
	Id       *string          `json:"id,omitempty"`
	// IP ranges registered in facility. Can be used for GeoIP location
	IpRanges             []string     `json:"ip_ranges,omitempty"`
	Metro                *DeviceMetro `json:"metro,omitempty"`
	Name                 *string      `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Facility Facility

// NewFacility instantiates a new Facility object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFacility() *Facility {
	this := Facility{}
	return &this
}

// NewFacilityWithDefaults instantiates a new Facility object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFacilityWithDefaults() *Facility {
	this := Facility{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *Facility) GetAddress() Address {
	if o == nil || IsNil(o.Address) {
		var ret Address
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Facility) GetAddressOk() (*Address, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *Facility) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given Address and assigns it to the Address field.
func (o *Facility) SetAddress(v Address) {
	o.Address = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *Facility) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Facility) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *Facility) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *Facility) SetCode(v string) {
	o.Code = &v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *Facility) GetFeatures() FacilityFeatures {
	if o == nil || IsNil(o.Features) {
		var ret FacilityFeatures
		return ret
	}
	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Facility) GetFeaturesOk() (FacilityFeatures, bool) {
	if o == nil || IsNil(o.Features) {
		return o.Features, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *Facility) HasFeatures() bool {
	if o != nil && !IsNil(o.Features) {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given []string and assigns it to the Features field.
func (o *Facility) SetFeatures(v FacilityFeatures) {
	o.Features = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Facility) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Facility) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Facility) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Facility) SetId(v string) {
	o.Id = &v
}

// GetIpRanges returns the IpRanges field value if set, zero value otherwise.
func (o *Facility) GetIpRanges() []string {
	if o == nil || IsNil(o.IpRanges) {
		var ret []string
		return ret
	}
	return o.IpRanges
}

// GetIpRangesOk returns a tuple with the IpRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Facility) GetIpRangesOk() ([]string, bool) {
	if o == nil || IsNil(o.IpRanges) {
		return nil, false
	}
	return o.IpRanges, true
}

// HasIpRanges returns a boolean if a field has been set.
func (o *Facility) HasIpRanges() bool {
	if o != nil && !IsNil(o.IpRanges) {
		return true
	}

	return false
}

// SetIpRanges gets a reference to the given []string and assigns it to the IpRanges field.
func (o *Facility) SetIpRanges(v []string) {
	o.IpRanges = v
}

// GetMetro returns the Metro field value if set, zero value otherwise.
func (o *Facility) GetMetro() DeviceMetro {
	if o == nil || IsNil(o.Metro) {
		var ret DeviceMetro
		return ret
	}
	return *o.Metro
}

// GetMetroOk returns a tuple with the Metro field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Facility) GetMetroOk() (*DeviceMetro, bool) {
	if o == nil || IsNil(o.Metro) {
		return nil, false
	}
	return o.Metro, true
}

// HasMetro returns a boolean if a field has been set.
func (o *Facility) HasMetro() bool {
	if o != nil && !IsNil(o.Metro) {
		return true
	}

	return false
}

// SetMetro gets a reference to the given DeviceMetro and assigns it to the Metro field.
func (o *Facility) SetMetro(v DeviceMetro) {
	o.Metro = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Facility) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Facility) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Facility) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Facility) SetName(v string) {
	o.Name = &v
}

func (o Facility) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Facility) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Features) {
		toSerialize["features"] = o.Features
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.IpRanges) {
		toSerialize["ip_ranges"] = o.IpRanges
	}
	if !IsNil(o.Metro) {
		toSerialize["metro"] = o.Metro
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Facility) UnmarshalJSON(bytes []byte) (err error) {
	varFacility := _Facility{}

	err = json.Unmarshal(bytes, &varFacility)

	if err != nil {
		return err
	}

	*o = Facility(varFacility)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "address")
		delete(additionalProperties, "code")
		delete(additionalProperties, "features")
		delete(additionalProperties, "id")
		delete(additionalProperties, "ip_ranges")
		delete(additionalProperties, "metro")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFacility struct {
	value *Facility
	isSet bool
}

func (v NullableFacility) Get() *Facility {
	return v.value
}

func (v *NullableFacility) Set(val *Facility) {
	v.value = val
	v.isSet = true
}

func (v NullableFacility) IsSet() bool {
	return v.isSet
}

func (v *NullableFacility) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFacility(val *Facility) *NullableFacility {
	return &NullableFacility{value: val, isSet: true}
}

func (v NullableFacility) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFacility) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
