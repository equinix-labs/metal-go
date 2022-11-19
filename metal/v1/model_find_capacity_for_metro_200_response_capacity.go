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

// FindCapacityForMetro200ResponseCapacity struct for FindCapacityForMetro200ResponseCapacity
type FindCapacityForMetro200ResponseCapacity struct {
	Am *FindCapacityForFacility200ResponseCapacityAms1 `json:"am,omitempty"`
	At *FindCapacityForFacility200ResponseCapacityAms1 `json:"at,omitempty"`
	Ch *FindCapacityForFacility200ResponseCapacityAms1 `json:"ch,omitempty"`
	Da *FindCapacityForFacility200ResponseCapacityAms1 `json:"da,omitempty"`
	Dc *FindCapacityForFacility200ResponseCapacityAms1 `json:"dc,omitempty"`
	Fr *FindCapacityForFacility200ResponseCapacityAms1 `json:"fr,omitempty"`
	Hk *FindCapacityForFacility200ResponseCapacityAms1 `json:"hk,omitempty"`
	La *FindCapacityForFacility200ResponseCapacityAms1 `json:"la,omitempty"`
	Ld *FindCapacityForFacility200ResponseCapacityAms1 `json:"ld,omitempty"`
	Md *FindCapacityForFacility200ResponseCapacityAms1 `json:"md,omitempty"`
	Ny *FindCapacityForFacility200ResponseCapacityAms1 `json:"ny,omitempty"`
	Pa *FindCapacityForFacility200ResponseCapacityAms1 `json:"pa,omitempty"`
	Se *FindCapacityForFacility200ResponseCapacityAms1 `json:"se,omitempty"`
	Sg *FindCapacityForFacility200ResponseCapacityAms1 `json:"sg,omitempty"`
	Sl *FindCapacityForFacility200ResponseCapacityAms1 `json:"sl,omitempty"`
	Sp *FindCapacityForFacility200ResponseCapacityAms1 `json:"sp,omitempty"`
	Sv *FindCapacityForFacility200ResponseCapacityAms1 `json:"sv,omitempty"`
	Sy *FindCapacityForFacility200ResponseCapacityAms1 `json:"sy,omitempty"`
	Tr *FindCapacityForFacility200ResponseCapacityAms1 `json:"tr,omitempty"`
	Ty *FindCapacityForFacility200ResponseCapacityAms1 `json:"ty,omitempty"`
}

// NewFindCapacityForMetro200ResponseCapacity instantiates a new FindCapacityForMetro200ResponseCapacity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFindCapacityForMetro200ResponseCapacity() *FindCapacityForMetro200ResponseCapacity {
	this := FindCapacityForMetro200ResponseCapacity{}
	return &this
}

// NewFindCapacityForMetro200ResponseCapacityWithDefaults instantiates a new FindCapacityForMetro200ResponseCapacity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFindCapacityForMetro200ResponseCapacityWithDefaults() *FindCapacityForMetro200ResponseCapacity {
	this := FindCapacityForMetro200ResponseCapacity{}
	return &this
}

// GetAm returns the Am field value if set, zero value otherwise.
func (o *FindCapacityForMetro200ResponseCapacity) GetAm() FindCapacityForFacility200ResponseCapacityAms1 {
	if o == nil || isNil(o.Am) {
		var ret FindCapacityForFacility200ResponseCapacityAms1
		return ret
	}
	return *o.Am
}

// GetAmOk returns a tuple with the Am field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindCapacityForMetro200ResponseCapacity) GetAmOk() (*FindCapacityForFacility200ResponseCapacityAms1, bool) {
	if o == nil || isNil(o.Am) {
		return nil, false
	}
	return o.Am, true
}

// HasAm returns a boolean if a field has been set.
func (o *FindCapacityForMetro200ResponseCapacity) HasAm() bool {
	if o != nil && !isNil(o.Am) {
		return true
	}

	return false
}

// SetAm gets a reference to the given FindCapacityForFacility200ResponseCapacityAms1 and assigns it to the Am field.
func (o *FindCapacityForMetro200ResponseCapacity) SetAm(v FindCapacityForFacility200ResponseCapacityAms1) {
	o.Am = &v
}

// GetAt returns the At field value if set, zero value otherwise.
func (o *FindCapacityForMetro200ResponseCapacity) GetAt() FindCapacityForFacility200ResponseCapacityAms1 {
	if o == nil || isNil(o.At) {
		var ret FindCapacityForFacility200ResponseCapacityAms1
		return ret
	}
	return *o.At
}

// GetAtOk returns a tuple with the At field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindCapacityForMetro200ResponseCapacity) GetAtOk() (*FindCapacityForFacility200ResponseCapacityAms1, bool) {
	if o == nil || isNil(o.At) {
		return nil, false
	}
	return o.At, true
}

// HasAt returns a boolean if a field has been set.
func (o *FindCapacityForMetro200ResponseCapacity) HasAt() bool {
	if o != nil && !isNil(o.At) {
		return true
	}

	return false
}

// SetAt gets a reference to the given FindCapacityForFacility200ResponseCapacityAms1 and assigns it to the At field.
func (o *FindCapacityForMetro200ResponseCapacity) SetAt(v FindCapacityForFacility200ResponseCapacityAms1) {
	o.At = &v
}

// GetCh returns the Ch field value if set, zero value otherwise.
func (o *FindCapacityForMetro200ResponseCapacity) GetCh() FindCapacityForFacility200ResponseCapacityAms1 {
	if o == nil || isNil(o.Ch) {
		var ret FindCapacityForFacility200ResponseCapacityAms1
		return ret
	}
	return *o.Ch
}

// GetChOk returns a tuple with the Ch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindCapacityForMetro200ResponseCapacity) GetChOk() (*FindCapacityForFacility200ResponseCapacityAms1, bool) {
	if o == nil || isNil(o.Ch) {
		return nil, false
	}
	return o.Ch, true
}

// HasCh returns a boolean if a field has been set.
func (o *FindCapacityForMetro200ResponseCapacity) HasCh() bool {
	if o != nil && !isNil(o.Ch) {
		return true
	}

	return false
}

// SetCh gets a reference to the given FindCapacityForFacility200ResponseCapacityAms1 and assigns it to the Ch field.
func (o *FindCapacityForMetro200ResponseCapacity) SetCh(v FindCapacityForFacility200ResponseCapacityAms1) {
	o.Ch = &v
}

// GetDa returns the Da field value if set, zero value otherwise.
func (o *FindCapacityForMetro200ResponseCapacity) GetDa() FindCapacityForFacility200ResponseCapacityAms1 {
	if o == nil || isNil(o.Da) {
		var ret FindCapacityForFacility200ResponseCapacityAms1
		return ret
	}
	return *o.Da
}

// GetDaOk returns a tuple with the Da field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindCapacityForMetro200ResponseCapacity) GetDaOk() (*FindCapacityForFacility200ResponseCapacityAms1, bool) {
	if o == nil || isNil(o.Da) {
		return nil, false
	}
	return o.Da, true
}

// HasDa returns a boolean if a field has been set.
func (o *FindCapacityForMetro200ResponseCapacity) HasDa() bool {
	if o != nil && !isNil(o.Da) {
		return true
	}

	return false
}

// SetDa gets a reference to the given FindCapacityForFacility200ResponseCapacityAms1 and assigns it to the Da field.
func (o *FindCapacityForMetro200ResponseCapacity) SetDa(v FindCapacityForFacility200ResponseCapacityAms1) {
	o.Da = &v
}

// GetDc returns the Dc field value if set, zero value otherwise.
func (o *FindCapacityForMetro200ResponseCapacity) GetDc() FindCapacityForFacility200ResponseCapacityAms1 {
	if o == nil || isNil(o.Dc) {
		var ret FindCapacityForFacility200ResponseCapacityAms1
		return ret
	}
	return *o.Dc
}

// GetDcOk returns a tuple with the Dc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindCapacityForMetro200ResponseCapacity) GetDcOk() (*FindCapacityForFacility200ResponseCapacityAms1, bool) {
	if o == nil || isNil(o.Dc) {
		return nil, false
	}
	return o.Dc, true
}

// HasDc returns a boolean if a field has been set.
func (o *FindCapacityForMetro200ResponseCapacity) HasDc() bool {
	if o != nil && !isNil(o.Dc) {
		return true
	}

	return false
}

// SetDc gets a reference to the given FindCapacityForFacility200ResponseCapacityAms1 and assigns it to the Dc field.
func (o *FindCapacityForMetro200ResponseCapacity) SetDc(v FindCapacityForFacility200ResponseCapacityAms1) {
	o.Dc = &v
}

// GetFr returns the Fr field value if set, zero value otherwise.
func (o *FindCapacityForMetro200ResponseCapacity) GetFr() FindCapacityForFacility200ResponseCapacityAms1 {
	if o == nil || isNil(o.Fr) {
		var ret FindCapacityForFacility200ResponseCapacityAms1
		return ret
	}
	return *o.Fr
}

// GetFrOk returns a tuple with the Fr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindCapacityForMetro200ResponseCapacity) GetFrOk() (*FindCapacityForFacility200ResponseCapacityAms1, bool) {
	if o == nil || isNil(o.Fr) {
		return nil, false
	}
	return o.Fr, true
}

// HasFr returns a boolean if a field has been set.
func (o *FindCapacityForMetro200ResponseCapacity) HasFr() bool {
	if o != nil && !isNil(o.Fr) {
		return true
	}

	return false
}

// SetFr gets a reference to the given FindCapacityForFacility200ResponseCapacityAms1 and assigns it to the Fr field.
func (o *FindCapacityForMetro200ResponseCapacity) SetFr(v FindCapacityForFacility200ResponseCapacityAms1) {
	o.Fr = &v
}

// GetHk returns the Hk field value if set, zero value otherwise.
func (o *FindCapacityForMetro200ResponseCapacity) GetHk() FindCapacityForFacility200ResponseCapacityAms1 {
	if o == nil || isNil(o.Hk) {
		var ret FindCapacityForFacility200ResponseCapacityAms1
		return ret
	}
	return *o.Hk
}

// GetHkOk returns a tuple with the Hk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindCapacityForMetro200ResponseCapacity) GetHkOk() (*FindCapacityForFacility200ResponseCapacityAms1, bool) {
	if o == nil || isNil(o.Hk) {
		return nil, false
	}
	return o.Hk, true
}

// HasHk returns a boolean if a field has been set.
func (o *FindCapacityForMetro200ResponseCapacity) HasHk() bool {
	if o != nil && !isNil(o.Hk) {
		return true
	}

	return false
}

// SetHk gets a reference to the given FindCapacityForFacility200ResponseCapacityAms1 and assigns it to the Hk field.
func (o *FindCapacityForMetro200ResponseCapacity) SetHk(v FindCapacityForFacility200ResponseCapacityAms1) {
	o.Hk = &v
}

// GetLa returns the La field value if set, zero value otherwise.
func (o *FindCapacityForMetro200ResponseCapacity) GetLa() FindCapacityForFacility200ResponseCapacityAms1 {
	if o == nil || isNil(o.La) {
		var ret FindCapacityForFacility200ResponseCapacityAms1
		return ret
	}
	return *o.La
}

// GetLaOk returns a tuple with the La field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindCapacityForMetro200ResponseCapacity) GetLaOk() (*FindCapacityForFacility200ResponseCapacityAms1, bool) {
	if o == nil || isNil(o.La) {
		return nil, false
	}
	return o.La, true
}

// HasLa returns a boolean if a field has been set.
func (o *FindCapacityForMetro200ResponseCapacity) HasLa() bool {
	if o != nil && !isNil(o.La) {
		return true
	}

	return false
}

// SetLa gets a reference to the given FindCapacityForFacility200ResponseCapacityAms1 and assigns it to the La field.
func (o *FindCapacityForMetro200ResponseCapacity) SetLa(v FindCapacityForFacility200ResponseCapacityAms1) {
	o.La = &v
}

// GetLd returns the Ld field value if set, zero value otherwise.
func (o *FindCapacityForMetro200ResponseCapacity) GetLd() FindCapacityForFacility200ResponseCapacityAms1 {
	if o == nil || isNil(o.Ld) {
		var ret FindCapacityForFacility200ResponseCapacityAms1
		return ret
	}
	return *o.Ld
}

// GetLdOk returns a tuple with the Ld field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindCapacityForMetro200ResponseCapacity) GetLdOk() (*FindCapacityForFacility200ResponseCapacityAms1, bool) {
	if o == nil || isNil(o.Ld) {
		return nil, false
	}
	return o.Ld, true
}

// HasLd returns a boolean if a field has been set.
func (o *FindCapacityForMetro200ResponseCapacity) HasLd() bool {
	if o != nil && !isNil(o.Ld) {
		return true
	}

	return false
}

// SetLd gets a reference to the given FindCapacityForFacility200ResponseCapacityAms1 and assigns it to the Ld field.
func (o *FindCapacityForMetro200ResponseCapacity) SetLd(v FindCapacityForFacility200ResponseCapacityAms1) {
	o.Ld = &v
}

// GetMd returns the Md field value if set, zero value otherwise.
func (o *FindCapacityForMetro200ResponseCapacity) GetMd() FindCapacityForFacility200ResponseCapacityAms1 {
	if o == nil || isNil(o.Md) {
		var ret FindCapacityForFacility200ResponseCapacityAms1
		return ret
	}
	return *o.Md
}

// GetMdOk returns a tuple with the Md field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindCapacityForMetro200ResponseCapacity) GetMdOk() (*FindCapacityForFacility200ResponseCapacityAms1, bool) {
	if o == nil || isNil(o.Md) {
		return nil, false
	}
	return o.Md, true
}

// HasMd returns a boolean if a field has been set.
func (o *FindCapacityForMetro200ResponseCapacity) HasMd() bool {
	if o != nil && !isNil(o.Md) {
		return true
	}

	return false
}

// SetMd gets a reference to the given FindCapacityForFacility200ResponseCapacityAms1 and assigns it to the Md field.
func (o *FindCapacityForMetro200ResponseCapacity) SetMd(v FindCapacityForFacility200ResponseCapacityAms1) {
	o.Md = &v
}

// GetNy returns the Ny field value if set, zero value otherwise.
func (o *FindCapacityForMetro200ResponseCapacity) GetNy() FindCapacityForFacility200ResponseCapacityAms1 {
	if o == nil || isNil(o.Ny) {
		var ret FindCapacityForFacility200ResponseCapacityAms1
		return ret
	}
	return *o.Ny
}

// GetNyOk returns a tuple with the Ny field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindCapacityForMetro200ResponseCapacity) GetNyOk() (*FindCapacityForFacility200ResponseCapacityAms1, bool) {
	if o == nil || isNil(o.Ny) {
		return nil, false
	}
	return o.Ny, true
}

// HasNy returns a boolean if a field has been set.
func (o *FindCapacityForMetro200ResponseCapacity) HasNy() bool {
	if o != nil && !isNil(o.Ny) {
		return true
	}

	return false
}

// SetNy gets a reference to the given FindCapacityForFacility200ResponseCapacityAms1 and assigns it to the Ny field.
func (o *FindCapacityForMetro200ResponseCapacity) SetNy(v FindCapacityForFacility200ResponseCapacityAms1) {
	o.Ny = &v
}

// GetPa returns the Pa field value if set, zero value otherwise.
func (o *FindCapacityForMetro200ResponseCapacity) GetPa() FindCapacityForFacility200ResponseCapacityAms1 {
	if o == nil || isNil(o.Pa) {
		var ret FindCapacityForFacility200ResponseCapacityAms1
		return ret
	}
	return *o.Pa
}

// GetPaOk returns a tuple with the Pa field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindCapacityForMetro200ResponseCapacity) GetPaOk() (*FindCapacityForFacility200ResponseCapacityAms1, bool) {
	if o == nil || isNil(o.Pa) {
		return nil, false
	}
	return o.Pa, true
}

// HasPa returns a boolean if a field has been set.
func (o *FindCapacityForMetro200ResponseCapacity) HasPa() bool {
	if o != nil && !isNil(o.Pa) {
		return true
	}

	return false
}

// SetPa gets a reference to the given FindCapacityForFacility200ResponseCapacityAms1 and assigns it to the Pa field.
func (o *FindCapacityForMetro200ResponseCapacity) SetPa(v FindCapacityForFacility200ResponseCapacityAms1) {
	o.Pa = &v
}

// GetSe returns the Se field value if set, zero value otherwise.
func (o *FindCapacityForMetro200ResponseCapacity) GetSe() FindCapacityForFacility200ResponseCapacityAms1 {
	if o == nil || isNil(o.Se) {
		var ret FindCapacityForFacility200ResponseCapacityAms1
		return ret
	}
	return *o.Se
}

// GetSeOk returns a tuple with the Se field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindCapacityForMetro200ResponseCapacity) GetSeOk() (*FindCapacityForFacility200ResponseCapacityAms1, bool) {
	if o == nil || isNil(o.Se) {
		return nil, false
	}
	return o.Se, true
}

// HasSe returns a boolean if a field has been set.
func (o *FindCapacityForMetro200ResponseCapacity) HasSe() bool {
	if o != nil && !isNil(o.Se) {
		return true
	}

	return false
}

// SetSe gets a reference to the given FindCapacityForFacility200ResponseCapacityAms1 and assigns it to the Se field.
func (o *FindCapacityForMetro200ResponseCapacity) SetSe(v FindCapacityForFacility200ResponseCapacityAms1) {
	o.Se = &v
}

// GetSg returns the Sg field value if set, zero value otherwise.
func (o *FindCapacityForMetro200ResponseCapacity) GetSg() FindCapacityForFacility200ResponseCapacityAms1 {
	if o == nil || isNil(o.Sg) {
		var ret FindCapacityForFacility200ResponseCapacityAms1
		return ret
	}
	return *o.Sg
}

// GetSgOk returns a tuple with the Sg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindCapacityForMetro200ResponseCapacity) GetSgOk() (*FindCapacityForFacility200ResponseCapacityAms1, bool) {
	if o == nil || isNil(o.Sg) {
		return nil, false
	}
	return o.Sg, true
}

// HasSg returns a boolean if a field has been set.
func (o *FindCapacityForMetro200ResponseCapacity) HasSg() bool {
	if o != nil && !isNil(o.Sg) {
		return true
	}

	return false
}

// SetSg gets a reference to the given FindCapacityForFacility200ResponseCapacityAms1 and assigns it to the Sg field.
func (o *FindCapacityForMetro200ResponseCapacity) SetSg(v FindCapacityForFacility200ResponseCapacityAms1) {
	o.Sg = &v
}

// GetSl returns the Sl field value if set, zero value otherwise.
func (o *FindCapacityForMetro200ResponseCapacity) GetSl() FindCapacityForFacility200ResponseCapacityAms1 {
	if o == nil || isNil(o.Sl) {
		var ret FindCapacityForFacility200ResponseCapacityAms1
		return ret
	}
	return *o.Sl
}

// GetSlOk returns a tuple with the Sl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindCapacityForMetro200ResponseCapacity) GetSlOk() (*FindCapacityForFacility200ResponseCapacityAms1, bool) {
	if o == nil || isNil(o.Sl) {
		return nil, false
	}
	return o.Sl, true
}

// HasSl returns a boolean if a field has been set.
func (o *FindCapacityForMetro200ResponseCapacity) HasSl() bool {
	if o != nil && !isNil(o.Sl) {
		return true
	}

	return false
}

// SetSl gets a reference to the given FindCapacityForFacility200ResponseCapacityAms1 and assigns it to the Sl field.
func (o *FindCapacityForMetro200ResponseCapacity) SetSl(v FindCapacityForFacility200ResponseCapacityAms1) {
	o.Sl = &v
}

// GetSp returns the Sp field value if set, zero value otherwise.
func (o *FindCapacityForMetro200ResponseCapacity) GetSp() FindCapacityForFacility200ResponseCapacityAms1 {
	if o == nil || isNil(o.Sp) {
		var ret FindCapacityForFacility200ResponseCapacityAms1
		return ret
	}
	return *o.Sp
}

// GetSpOk returns a tuple with the Sp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindCapacityForMetro200ResponseCapacity) GetSpOk() (*FindCapacityForFacility200ResponseCapacityAms1, bool) {
	if o == nil || isNil(o.Sp) {
		return nil, false
	}
	return o.Sp, true
}

// HasSp returns a boolean if a field has been set.
func (o *FindCapacityForMetro200ResponseCapacity) HasSp() bool {
	if o != nil && !isNil(o.Sp) {
		return true
	}

	return false
}

// SetSp gets a reference to the given FindCapacityForFacility200ResponseCapacityAms1 and assigns it to the Sp field.
func (o *FindCapacityForMetro200ResponseCapacity) SetSp(v FindCapacityForFacility200ResponseCapacityAms1) {
	o.Sp = &v
}

// GetSv returns the Sv field value if set, zero value otherwise.
func (o *FindCapacityForMetro200ResponseCapacity) GetSv() FindCapacityForFacility200ResponseCapacityAms1 {
	if o == nil || isNil(o.Sv) {
		var ret FindCapacityForFacility200ResponseCapacityAms1
		return ret
	}
	return *o.Sv
}

// GetSvOk returns a tuple with the Sv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindCapacityForMetro200ResponseCapacity) GetSvOk() (*FindCapacityForFacility200ResponseCapacityAms1, bool) {
	if o == nil || isNil(o.Sv) {
		return nil, false
	}
	return o.Sv, true
}

// HasSv returns a boolean if a field has been set.
func (o *FindCapacityForMetro200ResponseCapacity) HasSv() bool {
	if o != nil && !isNil(o.Sv) {
		return true
	}

	return false
}

// SetSv gets a reference to the given FindCapacityForFacility200ResponseCapacityAms1 and assigns it to the Sv field.
func (o *FindCapacityForMetro200ResponseCapacity) SetSv(v FindCapacityForFacility200ResponseCapacityAms1) {
	o.Sv = &v
}

// GetSy returns the Sy field value if set, zero value otherwise.
func (o *FindCapacityForMetro200ResponseCapacity) GetSy() FindCapacityForFacility200ResponseCapacityAms1 {
	if o == nil || isNil(o.Sy) {
		var ret FindCapacityForFacility200ResponseCapacityAms1
		return ret
	}
	return *o.Sy
}

// GetSyOk returns a tuple with the Sy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindCapacityForMetro200ResponseCapacity) GetSyOk() (*FindCapacityForFacility200ResponseCapacityAms1, bool) {
	if o == nil || isNil(o.Sy) {
		return nil, false
	}
	return o.Sy, true
}

// HasSy returns a boolean if a field has been set.
func (o *FindCapacityForMetro200ResponseCapacity) HasSy() bool {
	if o != nil && !isNil(o.Sy) {
		return true
	}

	return false
}

// SetSy gets a reference to the given FindCapacityForFacility200ResponseCapacityAms1 and assigns it to the Sy field.
func (o *FindCapacityForMetro200ResponseCapacity) SetSy(v FindCapacityForFacility200ResponseCapacityAms1) {
	o.Sy = &v
}

// GetTr returns the Tr field value if set, zero value otherwise.
func (o *FindCapacityForMetro200ResponseCapacity) GetTr() FindCapacityForFacility200ResponseCapacityAms1 {
	if o == nil || isNil(o.Tr) {
		var ret FindCapacityForFacility200ResponseCapacityAms1
		return ret
	}
	return *o.Tr
}

// GetTrOk returns a tuple with the Tr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindCapacityForMetro200ResponseCapacity) GetTrOk() (*FindCapacityForFacility200ResponseCapacityAms1, bool) {
	if o == nil || isNil(o.Tr) {
		return nil, false
	}
	return o.Tr, true
}

// HasTr returns a boolean if a field has been set.
func (o *FindCapacityForMetro200ResponseCapacity) HasTr() bool {
	if o != nil && !isNil(o.Tr) {
		return true
	}

	return false
}

// SetTr gets a reference to the given FindCapacityForFacility200ResponseCapacityAms1 and assigns it to the Tr field.
func (o *FindCapacityForMetro200ResponseCapacity) SetTr(v FindCapacityForFacility200ResponseCapacityAms1) {
	o.Tr = &v
}

// GetTy returns the Ty field value if set, zero value otherwise.
func (o *FindCapacityForMetro200ResponseCapacity) GetTy() FindCapacityForFacility200ResponseCapacityAms1 {
	if o == nil || isNil(o.Ty) {
		var ret FindCapacityForFacility200ResponseCapacityAms1
		return ret
	}
	return *o.Ty
}

// GetTyOk returns a tuple with the Ty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindCapacityForMetro200ResponseCapacity) GetTyOk() (*FindCapacityForFacility200ResponseCapacityAms1, bool) {
	if o == nil || isNil(o.Ty) {
		return nil, false
	}
	return o.Ty, true
}

// HasTy returns a boolean if a field has been set.
func (o *FindCapacityForMetro200ResponseCapacity) HasTy() bool {
	if o != nil && !isNil(o.Ty) {
		return true
	}

	return false
}

// SetTy gets a reference to the given FindCapacityForFacility200ResponseCapacityAms1 and assigns it to the Ty field.
func (o *FindCapacityForMetro200ResponseCapacity) SetTy(v FindCapacityForFacility200ResponseCapacityAms1) {
	o.Ty = &v
}

func (o FindCapacityForMetro200ResponseCapacity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Am) {
		toSerialize["am"] = o.Am
	}
	if !isNil(o.At) {
		toSerialize["at"] = o.At
	}
	if !isNil(o.Ch) {
		toSerialize["ch"] = o.Ch
	}
	if !isNil(o.Da) {
		toSerialize["da"] = o.Da
	}
	if !isNil(o.Dc) {
		toSerialize["dc"] = o.Dc
	}
	if !isNil(o.Fr) {
		toSerialize["fr"] = o.Fr
	}
	if !isNil(o.Hk) {
		toSerialize["hk"] = o.Hk
	}
	if !isNil(o.La) {
		toSerialize["la"] = o.La
	}
	if !isNil(o.Ld) {
		toSerialize["ld"] = o.Ld
	}
	if !isNil(o.Md) {
		toSerialize["md"] = o.Md
	}
	if !isNil(o.Ny) {
		toSerialize["ny"] = o.Ny
	}
	if !isNil(o.Pa) {
		toSerialize["pa"] = o.Pa
	}
	if !isNil(o.Se) {
		toSerialize["se"] = o.Se
	}
	if !isNil(o.Sg) {
		toSerialize["sg"] = o.Sg
	}
	if !isNil(o.Sl) {
		toSerialize["sl"] = o.Sl
	}
	if !isNil(o.Sp) {
		toSerialize["sp"] = o.Sp
	}
	if !isNil(o.Sv) {
		toSerialize["sv"] = o.Sv
	}
	if !isNil(o.Sy) {
		toSerialize["sy"] = o.Sy
	}
	if !isNil(o.Tr) {
		toSerialize["tr"] = o.Tr
	}
	if !isNil(o.Ty) {
		toSerialize["ty"] = o.Ty
	}
	return json.Marshal(toSerialize)
}

type NullableFindCapacityForMetro200ResponseCapacity struct {
	value *FindCapacityForMetro200ResponseCapacity
	isSet bool
}

func (v NullableFindCapacityForMetro200ResponseCapacity) Get() *FindCapacityForMetro200ResponseCapacity {
	return v.value
}

func (v *NullableFindCapacityForMetro200ResponseCapacity) Set(val *FindCapacityForMetro200ResponseCapacity) {
	v.value = val
	v.isSet = true
}

func (v NullableFindCapacityForMetro200ResponseCapacity) IsSet() bool {
	return v.isSet
}

func (v *NullableFindCapacityForMetro200ResponseCapacity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFindCapacityForMetro200ResponseCapacity(val *FindCapacityForMetro200ResponseCapacity) *NullableFindCapacityForMetro200ResponseCapacity {
	return &NullableFindCapacityForMetro200ResponseCapacity{value: val, isSet: true}
}

func (v NullableFindCapacityForMetro200ResponseCapacity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFindCapacityForMetro200ResponseCapacity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
