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
)

// CheckCapacityForMetro200ResponseServersInner struct for CheckCapacityForMetro200ResponseServersInner
type CheckCapacityForMetro200ResponseServersInner struct {
	// Returns true if there is enough capacity in the metro to fulfill the quantity set. Returns false if there is not enough.
	Available *bool `json:"available,omitempty"`
	// The metro ID or code sent to check capacity.
	Metro *string `json:"metro,omitempty"`
	// The plan ID or slug sent to check capacity.
	Plan *string `json:"plan,omitempty"`
	// The number of servers sent to check capacity.
	Quantity *string `json:"quantity,omitempty"`
}

// NewCheckCapacityForMetro200ResponseServersInner instantiates a new CheckCapacityForMetro200ResponseServersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckCapacityForMetro200ResponseServersInner() *CheckCapacityForMetro200ResponseServersInner {
	this := CheckCapacityForMetro200ResponseServersInner{}
	return &this
}

// NewCheckCapacityForMetro200ResponseServersInnerWithDefaults instantiates a new CheckCapacityForMetro200ResponseServersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckCapacityForMetro200ResponseServersInnerWithDefaults() *CheckCapacityForMetro200ResponseServersInner {
	this := CheckCapacityForMetro200ResponseServersInner{}
	return &this
}

// GetAvailable returns the Available field value if set, zero value otherwise.
func (o *CheckCapacityForMetro200ResponseServersInner) GetAvailable() bool {
	if o == nil || isNil(o.Available) {
		var ret bool
		return ret
	}
	return *o.Available
}

// GetAvailableOk returns a tuple with the Available field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckCapacityForMetro200ResponseServersInner) GetAvailableOk() (*bool, bool) {
	if o == nil || isNil(o.Available) {
		return nil, false
	}
	return o.Available, true
}

// HasAvailable returns a boolean if a field has been set.
func (o *CheckCapacityForMetro200ResponseServersInner) HasAvailable() bool {
	if o != nil && !isNil(o.Available) {
		return true
	}

	return false
}

// SetAvailable gets a reference to the given bool and assigns it to the Available field.
func (o *CheckCapacityForMetro200ResponseServersInner) SetAvailable(v bool) {
	o.Available = &v
}

// GetMetro returns the Metro field value if set, zero value otherwise.
func (o *CheckCapacityForMetro200ResponseServersInner) GetMetro() string {
	if o == nil || isNil(o.Metro) {
		var ret string
		return ret
	}
	return *o.Metro
}

// GetMetroOk returns a tuple with the Metro field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckCapacityForMetro200ResponseServersInner) GetMetroOk() (*string, bool) {
	if o == nil || isNil(o.Metro) {
		return nil, false
	}
	return o.Metro, true
}

// HasMetro returns a boolean if a field has been set.
func (o *CheckCapacityForMetro200ResponseServersInner) HasMetro() bool {
	if o != nil && !isNil(o.Metro) {
		return true
	}

	return false
}

// SetMetro gets a reference to the given string and assigns it to the Metro field.
func (o *CheckCapacityForMetro200ResponseServersInner) SetMetro(v string) {
	o.Metro = &v
}

// GetPlan returns the Plan field value if set, zero value otherwise.
func (o *CheckCapacityForMetro200ResponseServersInner) GetPlan() string {
	if o == nil || isNil(o.Plan) {
		var ret string
		return ret
	}
	return *o.Plan
}

// GetPlanOk returns a tuple with the Plan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckCapacityForMetro200ResponseServersInner) GetPlanOk() (*string, bool) {
	if o == nil || isNil(o.Plan) {
		return nil, false
	}
	return o.Plan, true
}

// HasPlan returns a boolean if a field has been set.
func (o *CheckCapacityForMetro200ResponseServersInner) HasPlan() bool {
	if o != nil && !isNil(o.Plan) {
		return true
	}

	return false
}

// SetPlan gets a reference to the given string and assigns it to the Plan field.
func (o *CheckCapacityForMetro200ResponseServersInner) SetPlan(v string) {
	o.Plan = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *CheckCapacityForMetro200ResponseServersInner) GetQuantity() string {
	if o == nil || isNil(o.Quantity) {
		var ret string
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckCapacityForMetro200ResponseServersInner) GetQuantityOk() (*string, bool) {
	if o == nil || isNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *CheckCapacityForMetro200ResponseServersInner) HasQuantity() bool {
	if o != nil && !isNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given string and assigns it to the Quantity field.
func (o *CheckCapacityForMetro200ResponseServersInner) SetQuantity(v string) {
	o.Quantity = &v
}

func (o CheckCapacityForMetro200ResponseServersInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Available) {
		toSerialize["available"] = o.Available
	}
	if !isNil(o.Metro) {
		toSerialize["metro"] = o.Metro
	}
	if !isNil(o.Plan) {
		toSerialize["plan"] = o.Plan
	}
	if !isNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	return json.Marshal(toSerialize)
}

type NullableCheckCapacityForMetro200ResponseServersInner struct {
	value *CheckCapacityForMetro200ResponseServersInner
	isSet bool
}

func (v NullableCheckCapacityForMetro200ResponseServersInner) Get() *CheckCapacityForMetro200ResponseServersInner {
	return v.value
}

func (v *NullableCheckCapacityForMetro200ResponseServersInner) Set(val *CheckCapacityForMetro200ResponseServersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckCapacityForMetro200ResponseServersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckCapacityForMetro200ResponseServersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckCapacityForMetro200ResponseServersInner(val *CheckCapacityForMetro200ResponseServersInner) *NullableCheckCapacityForMetro200ResponseServersInner {
	return &NullableCheckCapacityForMetro200ResponseServersInner{value: val, isSet: true}
}

func (v NullableCheckCapacityForMetro200ResponseServersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckCapacityForMetro200ResponseServersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
