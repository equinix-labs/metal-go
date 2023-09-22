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
	"time"
)

// checks if the BgpSession type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BgpSession{}

// BgpSession struct for BgpSession
type BgpSession struct {
	AddressFamily        BgpSessionAddressFamily `json:"address_family"`
	CreatedAt            *time.Time              `json:"created_at,omitempty"`
	DefaultRoute         *bool                   `json:"default_route,omitempty"`
	Device               *Href                   `json:"device,omitempty"`
	Href                 *string                 `json:"href,omitempty"`
	Id                   *string                 `json:"id,omitempty"`
	LearnedRoutes        []string                `json:"learned_routes,omitempty"`
	Status               *BgpSessionStatus       `json:"status,omitempty"`
	UpdatedAt            *time.Time              `json:"updated_at,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BgpSession BgpSession

// NewBgpSession instantiates a new BgpSession object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBgpSession(addressFamily BgpSessionAddressFamily) *BgpSession {
	this := BgpSession{}
	this.AddressFamily = addressFamily
	return &this
}

// NewBgpSessionWithDefaults instantiates a new BgpSession object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBgpSessionWithDefaults() *BgpSession {
	this := BgpSession{}
	return &this
}

// GetAddressFamily returns the AddressFamily field value
func (o *BgpSession) GetAddressFamily() BgpSessionAddressFamily {
	if o == nil {
		var ret BgpSessionAddressFamily
		return ret
	}

	return o.AddressFamily
}

// GetAddressFamilyOk returns a tuple with the AddressFamily field value
// and a boolean to check if the value has been set.
func (o *BgpSession) GetAddressFamilyOk() (*BgpSessionAddressFamily, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AddressFamily, true
}

// SetAddressFamily sets field value
func (o *BgpSession) SetAddressFamily(v BgpSessionAddressFamily) {
	o.AddressFamily = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *BgpSession) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpSession) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *BgpSession) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *BgpSession) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDefaultRoute returns the DefaultRoute field value if set, zero value otherwise.
func (o *BgpSession) GetDefaultRoute() bool {
	if o == nil || IsNil(o.DefaultRoute) {
		var ret bool
		return ret
	}
	return *o.DefaultRoute
}

// GetDefaultRouteOk returns a tuple with the DefaultRoute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpSession) GetDefaultRouteOk() (*bool, bool) {
	if o == nil || IsNil(o.DefaultRoute) {
		return nil, false
	}
	return o.DefaultRoute, true
}

// HasDefaultRoute returns a boolean if a field has been set.
func (o *BgpSession) HasDefaultRoute() bool {
	if o != nil && !IsNil(o.DefaultRoute) {
		return true
	}

	return false
}

// SetDefaultRoute gets a reference to the given bool and assigns it to the DefaultRoute field.
func (o *BgpSession) SetDefaultRoute(v bool) {
	o.DefaultRoute = &v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *BgpSession) GetDevice() Href {
	if o == nil || IsNil(o.Device) {
		var ret Href
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpSession) GetDeviceOk() (*Href, bool) {
	if o == nil || IsNil(o.Device) {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *BgpSession) HasDevice() bool {
	if o != nil && !IsNil(o.Device) {
		return true
	}

	return false
}

// SetDevice gets a reference to the given Href and assigns it to the Device field.
func (o *BgpSession) SetDevice(v Href) {
	o.Device = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *BgpSession) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpSession) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *BgpSession) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *BgpSession) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BgpSession) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpSession) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BgpSession) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BgpSession) SetId(v string) {
	o.Id = &v
}

// GetLearnedRoutes returns the LearnedRoutes field value if set, zero value otherwise.
func (o *BgpSession) GetLearnedRoutes() []string {
	if o == nil || IsNil(o.LearnedRoutes) {
		var ret []string
		return ret
	}
	return o.LearnedRoutes
}

// GetLearnedRoutesOk returns a tuple with the LearnedRoutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpSession) GetLearnedRoutesOk() ([]string, bool) {
	if o == nil || IsNil(o.LearnedRoutes) {
		return nil, false
	}
	return o.LearnedRoutes, true
}

// HasLearnedRoutes returns a boolean if a field has been set.
func (o *BgpSession) HasLearnedRoutes() bool {
	if o != nil && !IsNil(o.LearnedRoutes) {
		return true
	}

	return false
}

// SetLearnedRoutes gets a reference to the given []string and assigns it to the LearnedRoutes field.
func (o *BgpSession) SetLearnedRoutes(v []string) {
	o.LearnedRoutes = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BgpSession) GetStatus() BgpSessionStatus {
	if o == nil || IsNil(o.Status) {
		var ret BgpSessionStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpSession) GetStatusOk() (*BgpSessionStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BgpSession) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given BgpSessionStatus and assigns it to the Status field.
func (o *BgpSession) SetStatus(v BgpSessionStatus) {
	o.Status = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *BgpSession) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpSession) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *BgpSession) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *BgpSession) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o BgpSession) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BgpSession) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["address_family"] = o.AddressFamily
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.DefaultRoute) {
		toSerialize["default_route"] = o.DefaultRoute
	}
	if !IsNil(o.Device) {
		toSerialize["device"] = o.Device
	}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.LearnedRoutes) {
		toSerialize["learned_routes"] = o.LearnedRoutes
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BgpSession) UnmarshalJSON(bytes []byte) (err error) {
	varBgpSession := _BgpSession{}

	err = json.Unmarshal(bytes, &varBgpSession)

	if err != nil {
		return err
	}

	*o = BgpSession(varBgpSession)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "address_family")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "default_route")
		delete(additionalProperties, "device")
		delete(additionalProperties, "href")
		delete(additionalProperties, "id")
		delete(additionalProperties, "learned_routes")
		delete(additionalProperties, "status")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBgpSession struct {
	value *BgpSession
	isSet bool
}

func (v NullableBgpSession) Get() *BgpSession {
	return v.value
}

func (v *NullableBgpSession) Set(val *BgpSession) {
	v.value = val
	v.isSet = true
}

func (v NullableBgpSession) IsSet() bool {
	return v.isSet
}

func (v *NullableBgpSession) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBgpSession(val *BgpSession) *NullableBgpSession {
	return &NullableBgpSession{value: val, isSet: true}
}

func (v NullableBgpSession) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBgpSession) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
