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

// FindSpotMarketPrices200ResponseSpotMarketPrices struct for FindSpotMarketPrices200ResponseSpotMarketPrices
type FindSpotMarketPrices200ResponseSpotMarketPrices struct {
	Ams1 *FindSpotMarketPrices200ResponseSpotMarketPricesAms1 `json:"ams1,omitempty"`
	Atl1 *FindSpotMarketPrices200ResponseSpotMarketPricesAtl1 `json:"atl1,omitempty"`
	Dfw1 *FindSpotMarketPrices200ResponseSpotMarketPricesAtl1 `json:"dfw1,omitempty"`
	Ewr1 *FindSpotMarketPrices200ResponseSpotMarketPricesAms1 `json:"ewr1,omitempty"`
	Fra1 *FindSpotMarketPrices200ResponseSpotMarketPricesAtl1 `json:"fra1,omitempty"`
	Iad1 *FindSpotMarketPrices200ResponseSpotMarketPricesAtl1 `json:"iad1,omitempty"`
	Lax1 *FindSpotMarketPrices200ResponseSpotMarketPricesAtl1 `json:"lax1,omitempty"`
	Nrt1 *FindSpotMarketPrices200ResponseSpotMarketPricesAms1 `json:"nrt1,omitempty"`
	Ord1 *FindSpotMarketPrices200ResponseSpotMarketPricesAtl1 `json:"ord1,omitempty"`
	Sea1 *FindSpotMarketPrices200ResponseSpotMarketPricesAtl1 `json:"sea1,omitempty"`
	Sin1 *FindSpotMarketPrices200ResponseSpotMarketPricesAtl1 `json:"sin1,omitempty"`
	Sjc1 *FindSpotMarketPrices200ResponseSpotMarketPricesAms1 `json:"sjc1,omitempty"`
	Syd1 *FindSpotMarketPrices200ResponseSpotMarketPricesAtl1 `json:"syd1,omitempty"`
	Yyz1 *FindSpotMarketPrices200ResponseSpotMarketPricesAtl1 `json:"yyz1,omitempty"`
}

// NewFindSpotMarketPrices200ResponseSpotMarketPrices instantiates a new FindSpotMarketPrices200ResponseSpotMarketPrices object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFindSpotMarketPrices200ResponseSpotMarketPrices() *FindSpotMarketPrices200ResponseSpotMarketPrices {
	this := FindSpotMarketPrices200ResponseSpotMarketPrices{}
	return &this
}

// NewFindSpotMarketPrices200ResponseSpotMarketPricesWithDefaults instantiates a new FindSpotMarketPrices200ResponseSpotMarketPrices object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFindSpotMarketPrices200ResponseSpotMarketPricesWithDefaults() *FindSpotMarketPrices200ResponseSpotMarketPrices {
	this := FindSpotMarketPrices200ResponseSpotMarketPrices{}
	return &this
}

// GetAms1 returns the Ams1 field value if set, zero value otherwise.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) GetAms1() FindSpotMarketPrices200ResponseSpotMarketPricesAms1 {
	if o == nil || isNil(o.Ams1) {
		var ret FindSpotMarketPrices200ResponseSpotMarketPricesAms1
		return ret
	}
	return *o.Ams1
}

// GetAms1Ok returns a tuple with the Ams1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) GetAms1Ok() (*FindSpotMarketPrices200ResponseSpotMarketPricesAms1, bool) {
	if o == nil || isNil(o.Ams1) {
		return nil, false
	}
	return o.Ams1, true
}

// HasAms1 returns a boolean if a field has been set.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) HasAms1() bool {
	if o != nil && !isNil(o.Ams1) {
		return true
	}

	return false
}

// SetAms1 gets a reference to the given FindSpotMarketPrices200ResponseSpotMarketPricesAms1 and assigns it to the Ams1 field.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) SetAms1(v FindSpotMarketPrices200ResponseSpotMarketPricesAms1) {
	o.Ams1 = &v
}

// GetAtl1 returns the Atl1 field value if set, zero value otherwise.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) GetAtl1() FindSpotMarketPrices200ResponseSpotMarketPricesAtl1 {
	if o == nil || isNil(o.Atl1) {
		var ret FindSpotMarketPrices200ResponseSpotMarketPricesAtl1
		return ret
	}
	return *o.Atl1
}

// GetAtl1Ok returns a tuple with the Atl1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) GetAtl1Ok() (*FindSpotMarketPrices200ResponseSpotMarketPricesAtl1, bool) {
	if o == nil || isNil(o.Atl1) {
		return nil, false
	}
	return o.Atl1, true
}

// HasAtl1 returns a boolean if a field has been set.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) HasAtl1() bool {
	if o != nil && !isNil(o.Atl1) {
		return true
	}

	return false
}

// SetAtl1 gets a reference to the given FindSpotMarketPrices200ResponseSpotMarketPricesAtl1 and assigns it to the Atl1 field.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) SetAtl1(v FindSpotMarketPrices200ResponseSpotMarketPricesAtl1) {
	o.Atl1 = &v
}

// GetDfw1 returns the Dfw1 field value if set, zero value otherwise.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) GetDfw1() FindSpotMarketPrices200ResponseSpotMarketPricesAtl1 {
	if o == nil || isNil(o.Dfw1) {
		var ret FindSpotMarketPrices200ResponseSpotMarketPricesAtl1
		return ret
	}
	return *o.Dfw1
}

// GetDfw1Ok returns a tuple with the Dfw1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) GetDfw1Ok() (*FindSpotMarketPrices200ResponseSpotMarketPricesAtl1, bool) {
	if o == nil || isNil(o.Dfw1) {
		return nil, false
	}
	return o.Dfw1, true
}

// HasDfw1 returns a boolean if a field has been set.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) HasDfw1() bool {
	if o != nil && !isNil(o.Dfw1) {
		return true
	}

	return false
}

// SetDfw1 gets a reference to the given FindSpotMarketPrices200ResponseSpotMarketPricesAtl1 and assigns it to the Dfw1 field.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) SetDfw1(v FindSpotMarketPrices200ResponseSpotMarketPricesAtl1) {
	o.Dfw1 = &v
}

// GetEwr1 returns the Ewr1 field value if set, zero value otherwise.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) GetEwr1() FindSpotMarketPrices200ResponseSpotMarketPricesAms1 {
	if o == nil || isNil(o.Ewr1) {
		var ret FindSpotMarketPrices200ResponseSpotMarketPricesAms1
		return ret
	}
	return *o.Ewr1
}

// GetEwr1Ok returns a tuple with the Ewr1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) GetEwr1Ok() (*FindSpotMarketPrices200ResponseSpotMarketPricesAms1, bool) {
	if o == nil || isNil(o.Ewr1) {
		return nil, false
	}
	return o.Ewr1, true
}

// HasEwr1 returns a boolean if a field has been set.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) HasEwr1() bool {
	if o != nil && !isNil(o.Ewr1) {
		return true
	}

	return false
}

// SetEwr1 gets a reference to the given FindSpotMarketPrices200ResponseSpotMarketPricesAms1 and assigns it to the Ewr1 field.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) SetEwr1(v FindSpotMarketPrices200ResponseSpotMarketPricesAms1) {
	o.Ewr1 = &v
}

// GetFra1 returns the Fra1 field value if set, zero value otherwise.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) GetFra1() FindSpotMarketPrices200ResponseSpotMarketPricesAtl1 {
	if o == nil || isNil(o.Fra1) {
		var ret FindSpotMarketPrices200ResponseSpotMarketPricesAtl1
		return ret
	}
	return *o.Fra1
}

// GetFra1Ok returns a tuple with the Fra1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) GetFra1Ok() (*FindSpotMarketPrices200ResponseSpotMarketPricesAtl1, bool) {
	if o == nil || isNil(o.Fra1) {
		return nil, false
	}
	return o.Fra1, true
}

// HasFra1 returns a boolean if a field has been set.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) HasFra1() bool {
	if o != nil && !isNil(o.Fra1) {
		return true
	}

	return false
}

// SetFra1 gets a reference to the given FindSpotMarketPrices200ResponseSpotMarketPricesAtl1 and assigns it to the Fra1 field.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) SetFra1(v FindSpotMarketPrices200ResponseSpotMarketPricesAtl1) {
	o.Fra1 = &v
}

// GetIad1 returns the Iad1 field value if set, zero value otherwise.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) GetIad1() FindSpotMarketPrices200ResponseSpotMarketPricesAtl1 {
	if o == nil || isNil(o.Iad1) {
		var ret FindSpotMarketPrices200ResponseSpotMarketPricesAtl1
		return ret
	}
	return *o.Iad1
}

// GetIad1Ok returns a tuple with the Iad1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) GetIad1Ok() (*FindSpotMarketPrices200ResponseSpotMarketPricesAtl1, bool) {
	if o == nil || isNil(o.Iad1) {
		return nil, false
	}
	return o.Iad1, true
}

// HasIad1 returns a boolean if a field has been set.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) HasIad1() bool {
	if o != nil && !isNil(o.Iad1) {
		return true
	}

	return false
}

// SetIad1 gets a reference to the given FindSpotMarketPrices200ResponseSpotMarketPricesAtl1 and assigns it to the Iad1 field.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) SetIad1(v FindSpotMarketPrices200ResponseSpotMarketPricesAtl1) {
	o.Iad1 = &v
}

// GetLax1 returns the Lax1 field value if set, zero value otherwise.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) GetLax1() FindSpotMarketPrices200ResponseSpotMarketPricesAtl1 {
	if o == nil || isNil(o.Lax1) {
		var ret FindSpotMarketPrices200ResponseSpotMarketPricesAtl1
		return ret
	}
	return *o.Lax1
}

// GetLax1Ok returns a tuple with the Lax1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) GetLax1Ok() (*FindSpotMarketPrices200ResponseSpotMarketPricesAtl1, bool) {
	if o == nil || isNil(o.Lax1) {
		return nil, false
	}
	return o.Lax1, true
}

// HasLax1 returns a boolean if a field has been set.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) HasLax1() bool {
	if o != nil && !isNil(o.Lax1) {
		return true
	}

	return false
}

// SetLax1 gets a reference to the given FindSpotMarketPrices200ResponseSpotMarketPricesAtl1 and assigns it to the Lax1 field.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) SetLax1(v FindSpotMarketPrices200ResponseSpotMarketPricesAtl1) {
	o.Lax1 = &v
}

// GetNrt1 returns the Nrt1 field value if set, zero value otherwise.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) GetNrt1() FindSpotMarketPrices200ResponseSpotMarketPricesAms1 {
	if o == nil || isNil(o.Nrt1) {
		var ret FindSpotMarketPrices200ResponseSpotMarketPricesAms1
		return ret
	}
	return *o.Nrt1
}

// GetNrt1Ok returns a tuple with the Nrt1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) GetNrt1Ok() (*FindSpotMarketPrices200ResponseSpotMarketPricesAms1, bool) {
	if o == nil || isNil(o.Nrt1) {
		return nil, false
	}
	return o.Nrt1, true
}

// HasNrt1 returns a boolean if a field has been set.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) HasNrt1() bool {
	if o != nil && !isNil(o.Nrt1) {
		return true
	}

	return false
}

// SetNrt1 gets a reference to the given FindSpotMarketPrices200ResponseSpotMarketPricesAms1 and assigns it to the Nrt1 field.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) SetNrt1(v FindSpotMarketPrices200ResponseSpotMarketPricesAms1) {
	o.Nrt1 = &v
}

// GetOrd1 returns the Ord1 field value if set, zero value otherwise.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) GetOrd1() FindSpotMarketPrices200ResponseSpotMarketPricesAtl1 {
	if o == nil || isNil(o.Ord1) {
		var ret FindSpotMarketPrices200ResponseSpotMarketPricesAtl1
		return ret
	}
	return *o.Ord1
}

// GetOrd1Ok returns a tuple with the Ord1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) GetOrd1Ok() (*FindSpotMarketPrices200ResponseSpotMarketPricesAtl1, bool) {
	if o == nil || isNil(o.Ord1) {
		return nil, false
	}
	return o.Ord1, true
}

// HasOrd1 returns a boolean if a field has been set.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) HasOrd1() bool {
	if o != nil && !isNil(o.Ord1) {
		return true
	}

	return false
}

// SetOrd1 gets a reference to the given FindSpotMarketPrices200ResponseSpotMarketPricesAtl1 and assigns it to the Ord1 field.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) SetOrd1(v FindSpotMarketPrices200ResponseSpotMarketPricesAtl1) {
	o.Ord1 = &v
}

// GetSea1 returns the Sea1 field value if set, zero value otherwise.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) GetSea1() FindSpotMarketPrices200ResponseSpotMarketPricesAtl1 {
	if o == nil || isNil(o.Sea1) {
		var ret FindSpotMarketPrices200ResponseSpotMarketPricesAtl1
		return ret
	}
	return *o.Sea1
}

// GetSea1Ok returns a tuple with the Sea1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) GetSea1Ok() (*FindSpotMarketPrices200ResponseSpotMarketPricesAtl1, bool) {
	if o == nil || isNil(o.Sea1) {
		return nil, false
	}
	return o.Sea1, true
}

// HasSea1 returns a boolean if a field has been set.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) HasSea1() bool {
	if o != nil && !isNil(o.Sea1) {
		return true
	}

	return false
}

// SetSea1 gets a reference to the given FindSpotMarketPrices200ResponseSpotMarketPricesAtl1 and assigns it to the Sea1 field.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) SetSea1(v FindSpotMarketPrices200ResponseSpotMarketPricesAtl1) {
	o.Sea1 = &v
}

// GetSin1 returns the Sin1 field value if set, zero value otherwise.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) GetSin1() FindSpotMarketPrices200ResponseSpotMarketPricesAtl1 {
	if o == nil || isNil(o.Sin1) {
		var ret FindSpotMarketPrices200ResponseSpotMarketPricesAtl1
		return ret
	}
	return *o.Sin1
}

// GetSin1Ok returns a tuple with the Sin1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) GetSin1Ok() (*FindSpotMarketPrices200ResponseSpotMarketPricesAtl1, bool) {
	if o == nil || isNil(o.Sin1) {
		return nil, false
	}
	return o.Sin1, true
}

// HasSin1 returns a boolean if a field has been set.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) HasSin1() bool {
	if o != nil && !isNil(o.Sin1) {
		return true
	}

	return false
}

// SetSin1 gets a reference to the given FindSpotMarketPrices200ResponseSpotMarketPricesAtl1 and assigns it to the Sin1 field.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) SetSin1(v FindSpotMarketPrices200ResponseSpotMarketPricesAtl1) {
	o.Sin1 = &v
}

// GetSjc1 returns the Sjc1 field value if set, zero value otherwise.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) GetSjc1() FindSpotMarketPrices200ResponseSpotMarketPricesAms1 {
	if o == nil || isNil(o.Sjc1) {
		var ret FindSpotMarketPrices200ResponseSpotMarketPricesAms1
		return ret
	}
	return *o.Sjc1
}

// GetSjc1Ok returns a tuple with the Sjc1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) GetSjc1Ok() (*FindSpotMarketPrices200ResponseSpotMarketPricesAms1, bool) {
	if o == nil || isNil(o.Sjc1) {
		return nil, false
	}
	return o.Sjc1, true
}

// HasSjc1 returns a boolean if a field has been set.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) HasSjc1() bool {
	if o != nil && !isNil(o.Sjc1) {
		return true
	}

	return false
}

// SetSjc1 gets a reference to the given FindSpotMarketPrices200ResponseSpotMarketPricesAms1 and assigns it to the Sjc1 field.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) SetSjc1(v FindSpotMarketPrices200ResponseSpotMarketPricesAms1) {
	o.Sjc1 = &v
}

// GetSyd1 returns the Syd1 field value if set, zero value otherwise.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) GetSyd1() FindSpotMarketPrices200ResponseSpotMarketPricesAtl1 {
	if o == nil || isNil(o.Syd1) {
		var ret FindSpotMarketPrices200ResponseSpotMarketPricesAtl1
		return ret
	}
	return *o.Syd1
}

// GetSyd1Ok returns a tuple with the Syd1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) GetSyd1Ok() (*FindSpotMarketPrices200ResponseSpotMarketPricesAtl1, bool) {
	if o == nil || isNil(o.Syd1) {
		return nil, false
	}
	return o.Syd1, true
}

// HasSyd1 returns a boolean if a field has been set.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) HasSyd1() bool {
	if o != nil && !isNil(o.Syd1) {
		return true
	}

	return false
}

// SetSyd1 gets a reference to the given FindSpotMarketPrices200ResponseSpotMarketPricesAtl1 and assigns it to the Syd1 field.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) SetSyd1(v FindSpotMarketPrices200ResponseSpotMarketPricesAtl1) {
	o.Syd1 = &v
}

// GetYyz1 returns the Yyz1 field value if set, zero value otherwise.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) GetYyz1() FindSpotMarketPrices200ResponseSpotMarketPricesAtl1 {
	if o == nil || isNil(o.Yyz1) {
		var ret FindSpotMarketPrices200ResponseSpotMarketPricesAtl1
		return ret
	}
	return *o.Yyz1
}

// GetYyz1Ok returns a tuple with the Yyz1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) GetYyz1Ok() (*FindSpotMarketPrices200ResponseSpotMarketPricesAtl1, bool) {
	if o == nil || isNil(o.Yyz1) {
		return nil, false
	}
	return o.Yyz1, true
}

// HasYyz1 returns a boolean if a field has been set.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) HasYyz1() bool {
	if o != nil && !isNil(o.Yyz1) {
		return true
	}

	return false
}

// SetYyz1 gets a reference to the given FindSpotMarketPrices200ResponseSpotMarketPricesAtl1 and assigns it to the Yyz1 field.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) SetYyz1(v FindSpotMarketPrices200ResponseSpotMarketPricesAtl1) {
	o.Yyz1 = &v
}

func (o FindSpotMarketPrices200ResponseSpotMarketPrices) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Ams1) {
		toSerialize["ams1"] = o.Ams1
	}
	if !isNil(o.Atl1) {
		toSerialize["atl1"] = o.Atl1
	}
	if !isNil(o.Dfw1) {
		toSerialize["dfw1"] = o.Dfw1
	}
	if !isNil(o.Ewr1) {
		toSerialize["ewr1"] = o.Ewr1
	}
	if !isNil(o.Fra1) {
		toSerialize["fra1"] = o.Fra1
	}
	if !isNil(o.Iad1) {
		toSerialize["iad1"] = o.Iad1
	}
	if !isNil(o.Lax1) {
		toSerialize["lax1"] = o.Lax1
	}
	if !isNil(o.Nrt1) {
		toSerialize["nrt1"] = o.Nrt1
	}
	if !isNil(o.Ord1) {
		toSerialize["ord1"] = o.Ord1
	}
	if !isNil(o.Sea1) {
		toSerialize["sea1"] = o.Sea1
	}
	if !isNil(o.Sin1) {
		toSerialize["sin1"] = o.Sin1
	}
	if !isNil(o.Sjc1) {
		toSerialize["sjc1"] = o.Sjc1
	}
	if !isNil(o.Syd1) {
		toSerialize["syd1"] = o.Syd1
	}
	if !isNil(o.Yyz1) {
		toSerialize["yyz1"] = o.Yyz1
	}
	return json.Marshal(toSerialize)
}

type NullableFindSpotMarketPrices200ResponseSpotMarketPrices struct {
	value *FindSpotMarketPrices200ResponseSpotMarketPrices
	isSet bool
}

func (v NullableFindSpotMarketPrices200ResponseSpotMarketPrices) Get() *FindSpotMarketPrices200ResponseSpotMarketPrices {
	return v.value
}

func (v *NullableFindSpotMarketPrices200ResponseSpotMarketPrices) Set(val *FindSpotMarketPrices200ResponseSpotMarketPrices) {
	v.value = val
	v.isSet = true
}

func (v NullableFindSpotMarketPrices200ResponseSpotMarketPrices) IsSet() bool {
	return v.isSet
}

func (v *NullableFindSpotMarketPrices200ResponseSpotMarketPrices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFindSpotMarketPrices200ResponseSpotMarketPrices(val *FindSpotMarketPrices200ResponseSpotMarketPrices) *NullableFindSpotMarketPrices200ResponseSpotMarketPrices {
	return &NullableFindSpotMarketPrices200ResponseSpotMarketPrices{value: val, isSet: true}
}

func (v NullableFindSpotMarketPrices200ResponseSpotMarketPrices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFindSpotMarketPrices200ResponseSpotMarketPrices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
