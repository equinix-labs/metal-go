/*
Metal API

# Introduction Equinix Metal provides a RESTful HTTP API which can be reached at <https://api.equinix.com/metal/v1>. This document describes the API and how to use it.  The API allows you to programmatically interact with all of your Equinix Metal resources, including devices, networks, addresses, organizations, projects, and your user account. Every feature of the Equinix Metal web interface is accessible through the API.  The API docs are generated from the Equinix Metal OpenAPI specification and are officially hosted at <https://metal.equinix.com/developers/api>.  # Common Parameters  The Equinix Metal API uses a few methods to minimize network traffic and improve throughput. These parameters are not used in all API calls, but are used often enough to warrant their own section. Look for these parameters in the documentation for the API calls that support them.  ## Pagination  Pagination is used to limit the number of results returned in a single request. The API will return a maximum of 100 results per page. To retrieve additional results, you can use the `page` and `per_page` query parameters.  The `page` parameter is used to specify the page number. The first page is `1`. The `per_page` parameter is used to specify the number of results per page. The maximum number of results differs by resource type.  ## Sorting  Where offered, the API allows you to sort results by a specific field. To sort results use the `sort_by` query parameter with the root level field name as the value. The `sort_direction` parameter is used to specify the sort direction, either either `asc` (ascending) or `desc` (descending).  ## Filtering  Filtering is used to limit the results returned in a single request. The API supports filtering by certain fields in the response. To filter results, you can use the field as a query parameter.  For example, to filter the IP list to only return public IPv4 addresses, you can filter by the `type` field, as in the following request:  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/projects/id/ips?type=public_ipv4 ```  Only IP addresses with the `type` field set to `public_ipv4` will be returned.  ## Searching  Searching is used to find matching resources using multiple field comparissons. The API supports searching in resources that define this behavior. The fields available for search differ by resource, as does the search strategy.  To search resources you can use the `search` query parameter.  ## Include and Exclude  For resources that contain references to other resources, sucha as a Device that refers to the Project it resides in, the Equinix Metal API will returns `href` values (API links) to the associated resource.  ```json {   ...   \"project\": {     \"href\": \"/metal/v1/projects/f3f131c8-f302-49ef-8c44-9405022dc6dd\"   } } ```  If you're going need the project details, you can avoid a second API request.  Specify the contained `href` resources and collections that you'd like to have included in the response using the `include` query parameter.  For example:    ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=projects ```  The `include` parameter is generally accepted in `GET`, `POST`, `PUT`, and `PATCH` requests where `href` resources are presented.  To have multiple resources include, use a comma-separated list (e.g. `?include=emails,projects,memberships`).  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=emails,projects,memberships ```  You may also include nested associations up to three levels deep using dot notation (`?include=memberships.projects`):  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=memberships.projects ```  To exclude resources, and optimize response delivery, use the `exclude` query parameter. The `exclude` parameter is generally accepted in `GET`, `POST`, `PUT`, and `PATCH` requests for fields with nested object responses. When excluded, these fields will be replaced with an object that contains only an `href` field.

API version: 1.0.0
Contact: support@equinixmetal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
	"time"
)

// FindInterconnectionEvents200Response struct for FindInterconnectionEvents200Response
type FindInterconnectionEvents200Response struct {
	Body          *string                                `json:"body,omitempty"`
	CreatedAt     *time.Time                             `json:"created_at,omitempty"`
	Href          *string                                `json:"href,omitempty"`
	Id            *string                                `json:"id,omitempty"`
	Interpolated  *string                                `json:"interpolated,omitempty"`
	Relationships []FindBatchById200ResponseDevicesInner `json:"relationships,omitempty"`
	State         *string                                `json:"state,omitempty"`
	Type          *string                                `json:"type,omitempty"`
}

// NewFindInterconnectionEvents200Response instantiates a new FindInterconnectionEvents200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFindInterconnectionEvents200Response() *FindInterconnectionEvents200Response {
	this := FindInterconnectionEvents200Response{}
	return &this
}

// NewFindInterconnectionEvents200ResponseWithDefaults instantiates a new FindInterconnectionEvents200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFindInterconnectionEvents200ResponseWithDefaults() *FindInterconnectionEvents200Response {
	this := FindInterconnectionEvents200Response{}
	return &this
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *FindInterconnectionEvents200Response) GetBody() string {
	if o == nil || isNil(o.Body) {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindInterconnectionEvents200Response) GetBodyOk() (*string, bool) {
	if o == nil || isNil(o.Body) {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *FindInterconnectionEvents200Response) HasBody() bool {
	if o != nil && !isNil(o.Body) {
		return true
	}

	return false
}

// SetBody gets a reference to the given string and assigns it to the Body field.
func (o *FindInterconnectionEvents200Response) SetBody(v string) {
	o.Body = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *FindInterconnectionEvents200Response) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindInterconnectionEvents200Response) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *FindInterconnectionEvents200Response) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *FindInterconnectionEvents200Response) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *FindInterconnectionEvents200Response) GetHref() string {
	if o == nil || isNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindInterconnectionEvents200Response) GetHrefOk() (*string, bool) {
	if o == nil || isNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *FindInterconnectionEvents200Response) HasHref() bool {
	if o != nil && !isNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *FindInterconnectionEvents200Response) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FindInterconnectionEvents200Response) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindInterconnectionEvents200Response) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FindInterconnectionEvents200Response) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *FindInterconnectionEvents200Response) SetId(v string) {
	o.Id = &v
}

// GetInterpolated returns the Interpolated field value if set, zero value otherwise.
func (o *FindInterconnectionEvents200Response) GetInterpolated() string {
	if o == nil || isNil(o.Interpolated) {
		var ret string
		return ret
	}
	return *o.Interpolated
}

// GetInterpolatedOk returns a tuple with the Interpolated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindInterconnectionEvents200Response) GetInterpolatedOk() (*string, bool) {
	if o == nil || isNil(o.Interpolated) {
		return nil, false
	}
	return o.Interpolated, true
}

// HasInterpolated returns a boolean if a field has been set.
func (o *FindInterconnectionEvents200Response) HasInterpolated() bool {
	if o != nil && !isNil(o.Interpolated) {
		return true
	}

	return false
}

// SetInterpolated gets a reference to the given string and assigns it to the Interpolated field.
func (o *FindInterconnectionEvents200Response) SetInterpolated(v string) {
	o.Interpolated = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *FindInterconnectionEvents200Response) GetRelationships() []FindBatchById200ResponseDevicesInner {
	if o == nil || isNil(o.Relationships) {
		var ret []FindBatchById200ResponseDevicesInner
		return ret
	}
	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindInterconnectionEvents200Response) GetRelationshipsOk() ([]FindBatchById200ResponseDevicesInner, bool) {
	if o == nil || isNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *FindInterconnectionEvents200Response) HasRelationships() bool {
	if o != nil && !isNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given []FindBatchById200ResponseDevicesInner and assigns it to the Relationships field.
func (o *FindInterconnectionEvents200Response) SetRelationships(v []FindBatchById200ResponseDevicesInner) {
	o.Relationships = v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *FindInterconnectionEvents200Response) GetState() string {
	if o == nil || isNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindInterconnectionEvents200Response) GetStateOk() (*string, bool) {
	if o == nil || isNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *FindInterconnectionEvents200Response) HasState() bool {
	if o != nil && !isNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *FindInterconnectionEvents200Response) SetState(v string) {
	o.State = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FindInterconnectionEvents200Response) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindInterconnectionEvents200Response) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FindInterconnectionEvents200Response) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *FindInterconnectionEvents200Response) SetType(v string) {
	o.Type = &v
}

func (o FindInterconnectionEvents200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Body) {
		toSerialize["body"] = o.Body
	}
	if !isNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !isNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Interpolated) {
		toSerialize["interpolated"] = o.Interpolated
	}
	if !isNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	if !isNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableFindInterconnectionEvents200Response struct {
	value *FindInterconnectionEvents200Response
	isSet bool
}

func (v NullableFindInterconnectionEvents200Response) Get() *FindInterconnectionEvents200Response {
	return v.value
}

func (v *NullableFindInterconnectionEvents200Response) Set(val *FindInterconnectionEvents200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableFindInterconnectionEvents200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableFindInterconnectionEvents200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFindInterconnectionEvents200Response(val *FindInterconnectionEvents200Response) *NullableFindInterconnectionEvents200Response {
	return &NullableFindInterconnectionEvents200Response{value: val, isSet: true}
}

func (v NullableFindInterconnectionEvents200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFindInterconnectionEvents200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
