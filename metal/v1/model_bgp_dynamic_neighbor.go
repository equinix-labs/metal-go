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
	"time"
)

// checks if the BgpDynamicNeighbor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BgpDynamicNeighbor{}

// BgpDynamicNeighbor struct for BgpDynamicNeighbor
type BgpDynamicNeighbor struct {
	// The unique identifier for the resource
	Id *string `json:"id,omitempty"`
	// The ASN of the dynamic BGP neighbor
	BgpNeighborAsn *int32 `json:"bgp_neighbor_asn,omitempty"`
	// Network range of the dynamic BGP neighbor in CIDR format
	BgpNeighborRange     *string          `json:"bgp_neighbor_range,omitempty"`
	MetalGateway         *VrfMetalGateway `json:"metal_gateway,omitempty"`
	State                *string          `json:"state,omitempty"`
	Href                 *string          `json:"href,omitempty"`
	CreatedAt            *time.Time       `json:"created_at,omitempty"`
	CreatedBy            *UserLimited     `json:"created_by,omitempty"`
	UpdatedAt            *time.Time       `json:"updated_at,omitempty"`
	Tags                 []string         `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BgpDynamicNeighbor BgpDynamicNeighbor

// NewBgpDynamicNeighbor instantiates a new BgpDynamicNeighbor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBgpDynamicNeighbor() *BgpDynamicNeighbor {
	this := BgpDynamicNeighbor{}
	return &this
}

// NewBgpDynamicNeighborWithDefaults instantiates a new BgpDynamicNeighbor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBgpDynamicNeighborWithDefaults() *BgpDynamicNeighbor {
	this := BgpDynamicNeighbor{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BgpDynamicNeighbor) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpDynamicNeighbor) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BgpDynamicNeighbor) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BgpDynamicNeighbor) SetId(v string) {
	o.Id = &v
}

// GetBgpNeighborAsn returns the BgpNeighborAsn field value if set, zero value otherwise.
func (o *BgpDynamicNeighbor) GetBgpNeighborAsn() int32 {
	if o == nil || IsNil(o.BgpNeighborAsn) {
		var ret int32
		return ret
	}
	return *o.BgpNeighborAsn
}

// GetBgpNeighborAsnOk returns a tuple with the BgpNeighborAsn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpDynamicNeighbor) GetBgpNeighborAsnOk() (*int32, bool) {
	if o == nil || IsNil(o.BgpNeighborAsn) {
		return nil, false
	}
	return o.BgpNeighborAsn, true
}

// HasBgpNeighborAsn returns a boolean if a field has been set.
func (o *BgpDynamicNeighbor) HasBgpNeighborAsn() bool {
	if o != nil && !IsNil(o.BgpNeighborAsn) {
		return true
	}

	return false
}

// SetBgpNeighborAsn gets a reference to the given int32 and assigns it to the BgpNeighborAsn field.
func (o *BgpDynamicNeighbor) SetBgpNeighborAsn(v int32) {
	o.BgpNeighborAsn = &v
}

// GetBgpNeighborRange returns the BgpNeighborRange field value if set, zero value otherwise.
func (o *BgpDynamicNeighbor) GetBgpNeighborRange() string {
	if o == nil || IsNil(o.BgpNeighborRange) {
		var ret string
		return ret
	}
	return *o.BgpNeighborRange
}

// GetBgpNeighborRangeOk returns a tuple with the BgpNeighborRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpDynamicNeighbor) GetBgpNeighborRangeOk() (*string, bool) {
	if o == nil || IsNil(o.BgpNeighborRange) {
		return nil, false
	}
	return o.BgpNeighborRange, true
}

// HasBgpNeighborRange returns a boolean if a field has been set.
func (o *BgpDynamicNeighbor) HasBgpNeighborRange() bool {
	if o != nil && !IsNil(o.BgpNeighborRange) {
		return true
	}

	return false
}

// SetBgpNeighborRange gets a reference to the given string and assigns it to the BgpNeighborRange field.
func (o *BgpDynamicNeighbor) SetBgpNeighborRange(v string) {
	o.BgpNeighborRange = &v
}

// GetMetalGateway returns the MetalGateway field value if set, zero value otherwise.
func (o *BgpDynamicNeighbor) GetMetalGateway() VrfMetalGateway {
	if o == nil || IsNil(o.MetalGateway) {
		var ret VrfMetalGateway
		return ret
	}
	return *o.MetalGateway
}

// GetMetalGatewayOk returns a tuple with the MetalGateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpDynamicNeighbor) GetMetalGatewayOk() (*VrfMetalGateway, bool) {
	if o == nil || IsNil(o.MetalGateway) {
		return nil, false
	}
	return o.MetalGateway, true
}

// HasMetalGateway returns a boolean if a field has been set.
func (o *BgpDynamicNeighbor) HasMetalGateway() bool {
	if o != nil && !IsNil(o.MetalGateway) {
		return true
	}

	return false
}

// SetMetalGateway gets a reference to the given VrfMetalGateway and assigns it to the MetalGateway field.
func (o *BgpDynamicNeighbor) SetMetalGateway(v VrfMetalGateway) {
	o.MetalGateway = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *BgpDynamicNeighbor) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpDynamicNeighbor) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *BgpDynamicNeighbor) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *BgpDynamicNeighbor) SetState(v string) {
	o.State = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *BgpDynamicNeighbor) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpDynamicNeighbor) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *BgpDynamicNeighbor) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *BgpDynamicNeighbor) SetHref(v string) {
	o.Href = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *BgpDynamicNeighbor) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpDynamicNeighbor) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *BgpDynamicNeighbor) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *BgpDynamicNeighbor) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *BgpDynamicNeighbor) GetCreatedBy() UserLimited {
	if o == nil || IsNil(o.CreatedBy) {
		var ret UserLimited
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpDynamicNeighbor) GetCreatedByOk() (*UserLimited, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *BgpDynamicNeighbor) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given UserLimited and assigns it to the CreatedBy field.
func (o *BgpDynamicNeighbor) SetCreatedBy(v UserLimited) {
	o.CreatedBy = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *BgpDynamicNeighbor) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpDynamicNeighbor) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *BgpDynamicNeighbor) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *BgpDynamicNeighbor) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *BgpDynamicNeighbor) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpDynamicNeighbor) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *BgpDynamicNeighbor) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *BgpDynamicNeighbor) SetTags(v []string) {
	o.Tags = v
}

func (o BgpDynamicNeighbor) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BgpDynamicNeighbor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.BgpNeighborAsn) {
		toSerialize["bgp_neighbor_asn"] = o.BgpNeighborAsn
	}
	if !IsNil(o.BgpNeighborRange) {
		toSerialize["bgp_neighbor_range"] = o.BgpNeighborRange
	}
	if !IsNil(o.MetalGateway) {
		toSerialize["metal_gateway"] = o.MetalGateway
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.CreatedBy) {
		toSerialize["created_by"] = o.CreatedBy
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BgpDynamicNeighbor) UnmarshalJSON(bytes []byte) (err error) {
	varBgpDynamicNeighbor := _BgpDynamicNeighbor{}

	err = json.Unmarshal(bytes, &varBgpDynamicNeighbor)

	if err != nil {
		return err
	}

	*o = BgpDynamicNeighbor(varBgpDynamicNeighbor)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "bgp_neighbor_asn")
		delete(additionalProperties, "bgp_neighbor_range")
		delete(additionalProperties, "metal_gateway")
		delete(additionalProperties, "state")
		delete(additionalProperties, "href")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "created_by")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBgpDynamicNeighbor struct {
	value *BgpDynamicNeighbor
	isSet bool
}

func (v NullableBgpDynamicNeighbor) Get() *BgpDynamicNeighbor {
	return v.value
}

func (v *NullableBgpDynamicNeighbor) Set(val *BgpDynamicNeighbor) {
	v.value = val
	v.isSet = true
}

func (v NullableBgpDynamicNeighbor) IsSet() bool {
	return v.isSet
}

func (v *NullableBgpDynamicNeighbor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBgpDynamicNeighbor(val *BgpDynamicNeighbor) *NullableBgpDynamicNeighbor {
	return &NullableBgpDynamicNeighbor{value: val, isSet: true}
}

func (v NullableBgpDynamicNeighbor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBgpDynamicNeighbor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
