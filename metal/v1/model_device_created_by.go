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

// checks if the DeviceCreatedBy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceCreatedBy{}

// DeviceCreatedBy struct for DeviceCreatedBy
type DeviceCreatedBy struct {
	// Avatar thumbnail URL of the User
	AvatarThumbUrl *string `json:"avatar_thumb_url,omitempty"`
	// When the user was created
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Primary email address of the User
	Email *string `json:"email,omitempty"`
	// First name of the User
	FirstName *string `json:"first_name,omitempty"`
	// Full name of the User
	FullName *string `json:"full_name,omitempty"`
	// API URL uniquely representing the User
	Href *string `json:"href,omitempty"`
	// ID of the User
	Id string `json:"id"`
	// Last name of the User
	LastName *string `json:"last_name,omitempty"`
	// Short ID of the User
	ShortId string `json:"short_id"`
	// When the user details were last updated
	UpdatedAt            *time.Time `json:"updated_at,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeviceCreatedBy DeviceCreatedBy

// NewDeviceCreatedBy instantiates a new DeviceCreatedBy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceCreatedBy(id string, shortId string) *DeviceCreatedBy {
	this := DeviceCreatedBy{}
	this.Id = id
	this.ShortId = shortId
	return &this
}

// NewDeviceCreatedByWithDefaults instantiates a new DeviceCreatedBy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceCreatedByWithDefaults() *DeviceCreatedBy {
	this := DeviceCreatedBy{}
	return &this
}

// GetAvatarThumbUrl returns the AvatarThumbUrl field value if set, zero value otherwise.
func (o *DeviceCreatedBy) GetAvatarThumbUrl() string {
	if o == nil || IsNil(o.AvatarThumbUrl) {
		var ret string
		return ret
	}
	return *o.AvatarThumbUrl
}

// GetAvatarThumbUrlOk returns a tuple with the AvatarThumbUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceCreatedBy) GetAvatarThumbUrlOk() (*string, bool) {
	if o == nil || IsNil(o.AvatarThumbUrl) {
		return nil, false
	}
	return o.AvatarThumbUrl, true
}

// HasAvatarThumbUrl returns a boolean if a field has been set.
func (o *DeviceCreatedBy) HasAvatarThumbUrl() bool {
	if o != nil && !IsNil(o.AvatarThumbUrl) {
		return true
	}

	return false
}

// SetAvatarThumbUrl gets a reference to the given string and assigns it to the AvatarThumbUrl field.
func (o *DeviceCreatedBy) SetAvatarThumbUrl(v string) {
	o.AvatarThumbUrl = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *DeviceCreatedBy) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceCreatedBy) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *DeviceCreatedBy) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *DeviceCreatedBy) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *DeviceCreatedBy) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceCreatedBy) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *DeviceCreatedBy) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *DeviceCreatedBy) SetEmail(v string) {
	o.Email = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *DeviceCreatedBy) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceCreatedBy) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *DeviceCreatedBy) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *DeviceCreatedBy) SetFirstName(v string) {
	o.FirstName = &v
}

// GetFullName returns the FullName field value if set, zero value otherwise.
func (o *DeviceCreatedBy) GetFullName() string {
	if o == nil || IsNil(o.FullName) {
		var ret string
		return ret
	}
	return *o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceCreatedBy) GetFullNameOk() (*string, bool) {
	if o == nil || IsNil(o.FullName) {
		return nil, false
	}
	return o.FullName, true
}

// HasFullName returns a boolean if a field has been set.
func (o *DeviceCreatedBy) HasFullName() bool {
	if o != nil && !IsNil(o.FullName) {
		return true
	}

	return false
}

// SetFullName gets a reference to the given string and assigns it to the FullName field.
func (o *DeviceCreatedBy) SetFullName(v string) {
	o.FullName = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *DeviceCreatedBy) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceCreatedBy) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *DeviceCreatedBy) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *DeviceCreatedBy) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value
func (o *DeviceCreatedBy) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DeviceCreatedBy) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DeviceCreatedBy) SetId(v string) {
	o.Id = v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *DeviceCreatedBy) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceCreatedBy) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *DeviceCreatedBy) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *DeviceCreatedBy) SetLastName(v string) {
	o.LastName = &v
}

// GetShortId returns the ShortId field value
func (o *DeviceCreatedBy) GetShortId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ShortId
}

// GetShortIdOk returns a tuple with the ShortId field value
// and a boolean to check if the value has been set.
func (o *DeviceCreatedBy) GetShortIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShortId, true
}

// SetShortId sets field value
func (o *DeviceCreatedBy) SetShortId(v string) {
	o.ShortId = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *DeviceCreatedBy) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceCreatedBy) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *DeviceCreatedBy) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *DeviceCreatedBy) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o DeviceCreatedBy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceCreatedBy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AvatarThumbUrl) {
		toSerialize["avatar_thumb_url"] = o.AvatarThumbUrl
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.FirstName) {
		toSerialize["first_name"] = o.FirstName
	}
	if !IsNil(o.FullName) {
		toSerialize["full_name"] = o.FullName
	}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	toSerialize["id"] = o.Id
	if !IsNil(o.LastName) {
		toSerialize["last_name"] = o.LastName
	}
	toSerialize["short_id"] = o.ShortId
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeviceCreatedBy) UnmarshalJSON(bytes []byte) (err error) {
	requiredProperties := []string{
		"id",
		"short_id",
	}

	allProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &allProperties); err == nil {
		for _, requiredProperty := range requiredProperties {
			if _, exists := allProperties[requiredProperty]; !exists {
				return MissingRequiredFieldError(requiredProperty)
			}
		}
	}

	varDeviceCreatedBy := _DeviceCreatedBy{}

	err = json.Unmarshal(bytes, &varDeviceCreatedBy)

	if err != nil {
		return err
	}

	*o = DeviceCreatedBy(varDeviceCreatedBy)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "avatar_thumb_url")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "email")
		delete(additionalProperties, "first_name")
		delete(additionalProperties, "full_name")
		delete(additionalProperties, "href")
		delete(additionalProperties, "id")
		delete(additionalProperties, "last_name")
		delete(additionalProperties, "short_id")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeviceCreatedBy struct {
	value *DeviceCreatedBy
	isSet bool
}

func (v NullableDeviceCreatedBy) Get() *DeviceCreatedBy {
	return v.value
}

func (v *NullableDeviceCreatedBy) Set(val *DeviceCreatedBy) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceCreatedBy) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceCreatedBy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceCreatedBy(val *DeviceCreatedBy) *NullableDeviceCreatedBy {
	return &NullableDeviceCreatedBy{value: val, isSet: true}
}

func (v NullableDeviceCreatedBy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceCreatedBy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
