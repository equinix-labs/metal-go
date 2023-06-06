/*
Metal API

# Introduction Equinix Metal provides a RESTful HTTP API which can be reached at <https://api.equinix.com/metal/v1>. This document describes the API and how to use it.  The API allows you to programmatically interact with all of your Equinix Metal resources, including devices, networks, addresses, organizations, projects, and your user account. Every feature of the Equinix Metal web interface is accessible through the API.  The API docs are generated from the Equinix Metal OpenAPI specification and are officially hosted at <https://metal.equinix.com/developers/api>.  # Common Parameters  The Equinix Metal API uses a few methods to minimize network traffic and improve throughput. These parameters are not used in all API calls, but are used often enough to warrant their own section. Look for these parameters in the documentation for the API calls that support them.  ## Pagination  Pagination is used to limit the number of results returned in a single request. The API will return a maximum of 100 results per page. To retrieve additional results, you can use the `page` and `per_page` query parameters.  The `page` parameter is used to specify the page number. The first page is `1`. The `per_page` parameter is used to specify the number of results per page. The maximum number of results differs by resource type.  ## Sorting  Where offered, the API allows you to sort results by a specific field. To sort results use the `sort_by` query parameter with the root level field name as the value. The `sort_direction` parameter is used to specify the sort direction, either either `asc` (ascending) or `desc` (descending).  ## Filtering  Filtering is used to limit the results returned in a single request. The API supports filtering by certain fields in the response. To filter results, you can use the field as a query parameter.  For example, to filter the IP list to only return public IPv4 addresses, you can filter by the `type` field, as in the following request:  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/projects/id/ips?type=public_ipv4 ```  Only IP addresses with the `type` field set to `public_ipv4` will be returned.  ## Searching  Searching is used to find matching resources using multiple field comparissons. The API supports searching in resources that define this behavior. The fields available for search differ by resource, as does the search strategy.  To search resources you can use the `search` query parameter.  ## Include and Exclude  For resources that contain references to other resources, sucha as a Device that refers to the Project it resides in, the Equinix Metal API will returns `href` values (API links) to the associated resource.  ```json {   ...   \"project\": {     \"href\": \"/metal/v1/projects/f3f131c8-f302-49ef-8c44-9405022dc6dd\"   } } ```  If you're going need the project details, you can avoid a second API request.  Specify the contained `href` resources and collections that you'd like to have included in the response using the `include` query parameter.  For example:  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=projects ```  The `include` parameter is generally accepted in `GET`, `POST`, `PUT`, and `PATCH` requests where `href` resources are presented.  To have multiple resources include, use a comma-separated list (e.g. `?include=emails,projects,memberships`).  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=emails,projects,memberships ```  You may also include nested associations up to three levels deep using dot notation (`?include=memberships.projects`):  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=memberships.projects ```  To exclude resources, and optimize response delivery, use the `exclude` query parameter. The `exclude` parameter is generally accepted in `GET`, `POST`, `PUT`, and `PATCH` requests for fields with nested object responses. When excluded, these fields will be replaced with an object that contains only an `href` field.

API version: 1.0.0
Contact: support@equinixmetal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
)

// checks if the Entitlement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Entitlement{}

// Entitlement struct for Entitlement
type Entitlement struct {
	Description          *string                `json:"description,omitempty"`
	FeatureAccess        map[string]interface{} `json:"feature_access,omitempty"`
	Href                 *string                `json:"href,omitempty"`
	Id                   string                 `json:"id"`
	InstanceQuota        map[string]interface{} `json:"instance_quota,omitempty"`
	IpQuota              map[string]interface{} `json:"ip_quota,omitempty"`
	Name                 *string                `json:"name,omitempty"`
	ProjectQuota         *int32                 `json:"project_quota,omitempty"`
	Slug                 string                 `json:"slug"`
	VolumeLimits         map[string]interface{} `json:"volume_limits,omitempty"`
	VolumeQuota          map[string]interface{} `json:"volume_quota,omitempty"`
	Weight               int32                  `json:"weight"`
	AdditionalProperties map[string]interface{}
}

type _Entitlement Entitlement

// NewEntitlement instantiates a new Entitlement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlement(id string, slug string, weight int32) *Entitlement {
	this := Entitlement{}
	this.Id = id
	var projectQuota int32 = 0
	this.ProjectQuota = &projectQuota
	this.Slug = slug
	this.Weight = weight
	return &this
}

// NewEntitlementWithDefaults instantiates a new Entitlement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementWithDefaults() *Entitlement {
	this := Entitlement{}
	var projectQuota int32 = 0
	this.ProjectQuota = &projectQuota
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Entitlement) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Entitlement) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Entitlement) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Entitlement) SetDescription(v string) {
	o.Description = &v
}

// GetFeatureAccess returns the FeatureAccess field value if set, zero value otherwise.
func (o *Entitlement) GetFeatureAccess() map[string]interface{} {
	if o == nil || IsNil(o.FeatureAccess) {
		var ret map[string]interface{}
		return ret
	}
	return o.FeatureAccess
}

// GetFeatureAccessOk returns a tuple with the FeatureAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Entitlement) GetFeatureAccessOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.FeatureAccess) {
		return map[string]interface{}{}, false
	}
	return o.FeatureAccess, true
}

// HasFeatureAccess returns a boolean if a field has been set.
func (o *Entitlement) HasFeatureAccess() bool {
	if o != nil && !IsNil(o.FeatureAccess) {
		return true
	}

	return false
}

// SetFeatureAccess gets a reference to the given map[string]interface{} and assigns it to the FeatureAccess field.
func (o *Entitlement) SetFeatureAccess(v map[string]interface{}) {
	o.FeatureAccess = v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *Entitlement) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Entitlement) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *Entitlement) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *Entitlement) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value
func (o *Entitlement) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Entitlement) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Entitlement) SetId(v string) {
	o.Id = v
}

// GetInstanceQuota returns the InstanceQuota field value if set, zero value otherwise.
func (o *Entitlement) GetInstanceQuota() map[string]interface{} {
	if o == nil || IsNil(o.InstanceQuota) {
		var ret map[string]interface{}
		return ret
	}
	return o.InstanceQuota
}

// GetInstanceQuotaOk returns a tuple with the InstanceQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Entitlement) GetInstanceQuotaOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.InstanceQuota) {
		return map[string]interface{}{}, false
	}
	return o.InstanceQuota, true
}

// HasInstanceQuota returns a boolean if a field has been set.
func (o *Entitlement) HasInstanceQuota() bool {
	if o != nil && !IsNil(o.InstanceQuota) {
		return true
	}

	return false
}

// SetInstanceQuota gets a reference to the given map[string]interface{} and assigns it to the InstanceQuota field.
func (o *Entitlement) SetInstanceQuota(v map[string]interface{}) {
	o.InstanceQuota = v
}

// GetIpQuota returns the IpQuota field value if set, zero value otherwise.
func (o *Entitlement) GetIpQuota() map[string]interface{} {
	if o == nil || IsNil(o.IpQuota) {
		var ret map[string]interface{}
		return ret
	}
	return o.IpQuota
}

// GetIpQuotaOk returns a tuple with the IpQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Entitlement) GetIpQuotaOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.IpQuota) {
		return map[string]interface{}{}, false
	}
	return o.IpQuota, true
}

// HasIpQuota returns a boolean if a field has been set.
func (o *Entitlement) HasIpQuota() bool {
	if o != nil && !IsNil(o.IpQuota) {
		return true
	}

	return false
}

// SetIpQuota gets a reference to the given map[string]interface{} and assigns it to the IpQuota field.
func (o *Entitlement) SetIpQuota(v map[string]interface{}) {
	o.IpQuota = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Entitlement) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Entitlement) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Entitlement) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Entitlement) SetName(v string) {
	o.Name = &v
}

// GetProjectQuota returns the ProjectQuota field value if set, zero value otherwise.
func (o *Entitlement) GetProjectQuota() int32 {
	if o == nil || IsNil(o.ProjectQuota) {
		var ret int32
		return ret
	}
	return *o.ProjectQuota
}

// GetProjectQuotaOk returns a tuple with the ProjectQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Entitlement) GetProjectQuotaOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectQuota) {
		return nil, false
	}
	return o.ProjectQuota, true
}

// HasProjectQuota returns a boolean if a field has been set.
func (o *Entitlement) HasProjectQuota() bool {
	if o != nil && !IsNil(o.ProjectQuota) {
		return true
	}

	return false
}

// SetProjectQuota gets a reference to the given int32 and assigns it to the ProjectQuota field.
func (o *Entitlement) SetProjectQuota(v int32) {
	o.ProjectQuota = &v
}

// GetSlug returns the Slug field value
func (o *Entitlement) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *Entitlement) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *Entitlement) SetSlug(v string) {
	o.Slug = v
}

// GetVolumeLimits returns the VolumeLimits field value if set, zero value otherwise.
func (o *Entitlement) GetVolumeLimits() map[string]interface{} {
	if o == nil || IsNil(o.VolumeLimits) {
		var ret map[string]interface{}
		return ret
	}
	return o.VolumeLimits
}

// GetVolumeLimitsOk returns a tuple with the VolumeLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Entitlement) GetVolumeLimitsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.VolumeLimits) {
		return map[string]interface{}{}, false
	}
	return o.VolumeLimits, true
}

// HasVolumeLimits returns a boolean if a field has been set.
func (o *Entitlement) HasVolumeLimits() bool {
	if o != nil && !IsNil(o.VolumeLimits) {
		return true
	}

	return false
}

// SetVolumeLimits gets a reference to the given map[string]interface{} and assigns it to the VolumeLimits field.
func (o *Entitlement) SetVolumeLimits(v map[string]interface{}) {
	o.VolumeLimits = v
}

// GetVolumeQuota returns the VolumeQuota field value if set, zero value otherwise.
func (o *Entitlement) GetVolumeQuota() map[string]interface{} {
	if o == nil || IsNil(o.VolumeQuota) {
		var ret map[string]interface{}
		return ret
	}
	return o.VolumeQuota
}

// GetVolumeQuotaOk returns a tuple with the VolumeQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Entitlement) GetVolumeQuotaOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.VolumeQuota) {
		return map[string]interface{}{}, false
	}
	return o.VolumeQuota, true
}

// HasVolumeQuota returns a boolean if a field has been set.
func (o *Entitlement) HasVolumeQuota() bool {
	if o != nil && !IsNil(o.VolumeQuota) {
		return true
	}

	return false
}

// SetVolumeQuota gets a reference to the given map[string]interface{} and assigns it to the VolumeQuota field.
func (o *Entitlement) SetVolumeQuota(v map[string]interface{}) {
	o.VolumeQuota = v
}

// GetWeight returns the Weight field value
func (o *Entitlement) GetWeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Weight
}

// GetWeightOk returns a tuple with the Weight field value
// and a boolean to check if the value has been set.
func (o *Entitlement) GetWeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Weight, true
}

// SetWeight sets field value
func (o *Entitlement) SetWeight(v int32) {
	o.Weight = v
}

func (o Entitlement) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Entitlement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.FeatureAccess) {
		toSerialize["feature_access"] = o.FeatureAccess
	}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	toSerialize["id"] = o.Id
	if !IsNil(o.InstanceQuota) {
		toSerialize["instance_quota"] = o.InstanceQuota
	}
	if !IsNil(o.IpQuota) {
		toSerialize["ip_quota"] = o.IpQuota
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ProjectQuota) {
		toSerialize["project_quota"] = o.ProjectQuota
	}
	toSerialize["slug"] = o.Slug
	if !IsNil(o.VolumeLimits) {
		toSerialize["volume_limits"] = o.VolumeLimits
	}
	if !IsNil(o.VolumeQuota) {
		toSerialize["volume_quota"] = o.VolumeQuota
	}
	toSerialize["weight"] = o.Weight

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Entitlement) UnmarshalJSON(bytes []byte) (err error) {
	varEntitlement := _Entitlement{}

	if err = json.Unmarshal(bytes, &varEntitlement); err == nil {
		*o = Entitlement(varEntitlement)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "feature_access")
		delete(additionalProperties, "href")
		delete(additionalProperties, "id")
		delete(additionalProperties, "instance_quota")
		delete(additionalProperties, "ip_quota")
		delete(additionalProperties, "name")
		delete(additionalProperties, "project_quota")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "volume_limits")
		delete(additionalProperties, "volume_quota")
		delete(additionalProperties, "weight")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEntitlement struct {
	value *Entitlement
	isSet bool
}

func (v NullableEntitlement) Get() *Entitlement {
	return v.value
}

func (v *NullableEntitlement) Set(val *Entitlement) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlement) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlement(val *Entitlement) *NullableEntitlement {
	return &NullableEntitlement{value: val, isSet: true}
}

func (v NullableEntitlement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
