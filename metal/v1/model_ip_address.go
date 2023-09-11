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

// checks if the IPAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IPAddress{}

// IPAddressAddressFamily Address Family for IP Address
type IPAddressAddressFamily int32

// List of IPAddressAddressFamily
const (
	IPADDRESS__4 IPAddressAddressFamily = 4
	IPADDRESS__6 IPAddressAddressFamily = 6
)

// All allowed values of IPAddressAddressFamily enum
var AllowedIPAddressAddressFamilyEnumValues = []IPAddressAddressFamily{
	4,
	6,
}

func (v *IPAddressAddressFamily) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IPAddressAddressFamily(value)
	for _, existing := range AllowedIPAddressAddressFamilyEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IPAddressAddressFamily", value)
}

// NewIPAddressAddressFamilyFromValue returns a pointer to a valid IPAddressAddressFamily
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIPAddressAddressFamilyFromValue(v int32) (*IPAddressAddressFamily, error) {
	ev := IPAddressAddressFamily(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IPAddressAddressFamily: valid values are %v", v, AllowedIPAddressAddressFamilyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IPAddressAddressFamily) IsValid() bool {
	for _, existing := range AllowedIPAddressAddressFamilyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AddressFamily value
func (v IPAddressAddressFamily) Ptr() *IPAddressAddressFamily {
	return &v
}

type NullableIPAddressAddressFamily struct {
	value *IPAddressAddressFamily
	isSet bool
}

func (v NullableIPAddressAddressFamily) Get() *IPAddressAddressFamily {
	return v.value
}

func (v *NullableIPAddressAddressFamily) Set(val *IPAddressAddressFamily) {
	v.value = val
	v.isSet = true
}

func (v NullableIPAddressAddressFamily) IsSet() bool {
	return v.isSet
}

func (v *NullableIPAddressAddressFamily) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIPAddressAddressFamily(val *IPAddressAddressFamily) *NullableIPAddressAddressFamily {
	return &NullableIPAddressAddressFamily{value: val, isSet: true}
}

func (v NullableIPAddressAddressFamily) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIPAddressAddressFamily) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

// IPAddress struct for IPAddress
type IPAddress struct {
	// Address Family for IP Address
	AddressFamily *int32 `json:"address_family,omitempty"`
	// Cidr Size for the IP Block created. Valid values depends on the operating system being provisioned. (28..32 for IPv4 addresses, 124..127 for IPv6 addresses)
	Cidr *int32 `json:"cidr,omitempty"`
	// UUIDs of any IP reservations to use when assigning IPs
	IpReservations []string `json:"ip_reservations,omitempty"`
	// Address Type for IP Address
	Public               *bool `json:"public,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IPAddress IPAddress

// NewIPAddress instantiates a new IPAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIPAddress() *IPAddress {
	this := IPAddress{}
	var public bool = true
	this.Public = &public
	return &this
}

// NewIPAddressWithDefaults instantiates a new IPAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIPAddressWithDefaults() *IPAddress {
	this := IPAddress{}
	var public bool = true
	this.Public = &public
	return &this
}

// GetAddressFamily returns the AddressFamily field value if set, zero value otherwise.
func (o *IPAddress) GetAddressFamily() int32 {
	if o == nil || IsNil(o.AddressFamily) {
		var ret int32
		return ret
	}
	return *o.AddressFamily
}

// GetAddressFamilyOk returns a tuple with the AddressFamily field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPAddress) GetAddressFamilyOk() (*int32, bool) {
	if o == nil || IsNil(o.AddressFamily) {
		return nil, false
	}
	return o.AddressFamily, true
}

// HasAddressFamily returns a boolean if a field has been set.
func (o *IPAddress) HasAddressFamily() bool {
	if o != nil && !IsNil(o.AddressFamily) {
		return true
	}

	return false
}

// SetAddressFamily gets a reference to the given int32 and assigns it to the AddressFamily field.
func (o *IPAddress) SetAddressFamily(v int32) {
	o.AddressFamily = &v
}

// GetCidr returns the Cidr field value if set, zero value otherwise.
func (o *IPAddress) GetCidr() int32 {
	if o == nil || IsNil(o.Cidr) {
		var ret int32
		return ret
	}
	return *o.Cidr
}

// GetCidrOk returns a tuple with the Cidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPAddress) GetCidrOk() (*int32, bool) {
	if o == nil || IsNil(o.Cidr) {
		return nil, false
	}
	return o.Cidr, true
}

// HasCidr returns a boolean if a field has been set.
func (o *IPAddress) HasCidr() bool {
	if o != nil && !IsNil(o.Cidr) {
		return true
	}

	return false
}

// SetCidr gets a reference to the given int32 and assigns it to the Cidr field.
func (o *IPAddress) SetCidr(v int32) {
	o.Cidr = &v
}

// GetIpReservations returns the IpReservations field value if set, zero value otherwise.
func (o *IPAddress) GetIpReservations() []string {
	if o == nil || IsNil(o.IpReservations) {
		var ret []string
		return ret
	}
	return o.IpReservations
}

// GetIpReservationsOk returns a tuple with the IpReservations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPAddress) GetIpReservationsOk() ([]string, bool) {
	if o == nil || IsNil(o.IpReservations) {
		return nil, false
	}
	return o.IpReservations, true
}

// HasIpReservations returns a boolean if a field has been set.
func (o *IPAddress) HasIpReservations() bool {
	if o != nil && !IsNil(o.IpReservations) {
		return true
	}

	return false
}

// SetIpReservations gets a reference to the given []string and assigns it to the IpReservations field.
func (o *IPAddress) SetIpReservations(v []string) {
	o.IpReservations = v
}

// GetPublic returns the Public field value if set, zero value otherwise.
func (o *IPAddress) GetPublic() bool {
	if o == nil || IsNil(o.Public) {
		var ret bool
		return ret
	}
	return *o.Public
}

// GetPublicOk returns a tuple with the Public field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPAddress) GetPublicOk() (*bool, bool) {
	if o == nil || IsNil(o.Public) {
		return nil, false
	}
	return o.Public, true
}

// HasPublic returns a boolean if a field has been set.
func (o *IPAddress) HasPublic() bool {
	if o != nil && !IsNil(o.Public) {
		return true
	}

	return false
}

// SetPublic gets a reference to the given bool and assigns it to the Public field.
func (o *IPAddress) SetPublic(v bool) {
	o.Public = &v
}

func (o IPAddress) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IPAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AddressFamily) {
		toSerialize["address_family"] = o.AddressFamily
	}
	if !IsNil(o.Cidr) {
		toSerialize["cidr"] = o.Cidr
	}
	if !IsNil(o.IpReservations) {
		toSerialize["ip_reservations"] = o.IpReservations
	}
	if !IsNil(o.Public) {
		toSerialize["public"] = o.Public
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IPAddress) UnmarshalJSON(bytes []byte) (err error) {
	varIPAddress := _IPAddress{}

	err = json.Unmarshal(bytes, &varIPAddress)

	if err != nil {
		return err
	}

	*o = IPAddress(varIPAddress)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "address_family")
		delete(additionalProperties, "cidr")
		delete(additionalProperties, "ip_reservations")
		delete(additionalProperties, "public")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIPAddress struct {
	value *IPAddress
	isSet bool
}

func (v NullableIPAddress) Get() *IPAddress {
	return v.value
}

func (v *NullableIPAddress) Set(val *IPAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableIPAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableIPAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIPAddress(val *IPAddress) *NullableIPAddress {
	return &NullableIPAddress{value: val, isSet: true}
}

func (v NullableIPAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIPAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
