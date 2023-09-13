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
)

// checks if the Plan type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Plan{}

// Plan struct for Plan
type Plan struct {
	// Shows which facilities the plan is available in, and the facility-based price if it is different from the default price.
	AvailableIn []PlanAvailableInInner `json:"available_in,omitempty"`
	// Shows which metros the plan is available in, and the metro-based price if it is different from the default price.
	AvailableInMetros []PlanAvailableInMetrosInner `json:"available_in_metros,omitempty"`
	// Categories of the plan, like compute or storage. A Plan can belong to multiple categories.
	Categories      []string `json:"categories,omitempty"`
	Class           *string  `json:"class,omitempty"`
	Description     *string  `json:"description,omitempty"`
	DeploymentTypes []string `json:"deployment_types,omitempty"`
	Id              *string  `json:"id,omitempty"`
	// Deprecated. Always return false
	Legacy  *bool                  `json:"legacy,omitempty"`
	Line    *string                `json:"line,omitempty"`
	Name    *string                `json:"name,omitempty"`
	Pricing map[string]interface{} `json:"pricing,omitempty"`
	Slug    *string                `json:"slug,omitempty"`
	Specs   *PlanSpecs             `json:"specs,omitempty"`
	// The plan type
	Type                 *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Plan Plan

// NewPlan instantiates a new Plan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlan() *Plan {
	this := Plan{}
	return &this
}

// NewPlanWithDefaults instantiates a new Plan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlanWithDefaults() *Plan {
	this := Plan{}
	return &this
}

// GetAvailableIn returns the AvailableIn field value if set, zero value otherwise.
func (o *Plan) GetAvailableIn() []PlanAvailableInInner {
	if o == nil || IsNil(o.AvailableIn) {
		var ret []PlanAvailableInInner
		return ret
	}
	return o.AvailableIn
}

// GetAvailableInOk returns a tuple with the AvailableIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plan) GetAvailableInOk() ([]PlanAvailableInInner, bool) {
	if o == nil || IsNil(o.AvailableIn) {
		return nil, false
	}
	return o.AvailableIn, true
}

// HasAvailableIn returns a boolean if a field has been set.
func (o *Plan) HasAvailableIn() bool {
	if o != nil && !IsNil(o.AvailableIn) {
		return true
	}

	return false
}

// SetAvailableIn gets a reference to the given []PlanAvailableInInner and assigns it to the AvailableIn field.
func (o *Plan) SetAvailableIn(v []PlanAvailableInInner) {
	o.AvailableIn = v
}

// GetAvailableInMetros returns the AvailableInMetros field value if set, zero value otherwise.
func (o *Plan) GetAvailableInMetros() []PlanAvailableInMetrosInner {
	if o == nil || IsNil(o.AvailableInMetros) {
		var ret []PlanAvailableInMetrosInner
		return ret
	}
	return o.AvailableInMetros
}

// GetAvailableInMetrosOk returns a tuple with the AvailableInMetros field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plan) GetAvailableInMetrosOk() ([]PlanAvailableInMetrosInner, bool) {
	if o == nil || IsNil(o.AvailableInMetros) {
		return nil, false
	}
	return o.AvailableInMetros, true
}

// HasAvailableInMetros returns a boolean if a field has been set.
func (o *Plan) HasAvailableInMetros() bool {
	if o != nil && !IsNil(o.AvailableInMetros) {
		return true
	}

	return false
}

// SetAvailableInMetros gets a reference to the given []PlanAvailableInMetrosInner and assigns it to the AvailableInMetros field.
func (o *Plan) SetAvailableInMetros(v []PlanAvailableInMetrosInner) {
	o.AvailableInMetros = v
}

// GetCategories returns the Categories field value if set, zero value otherwise.
func (o *Plan) GetCategories() []string {
	if o == nil || IsNil(o.Categories) {
		var ret []string
		return ret
	}
	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plan) GetCategoriesOk() ([]string, bool) {
	if o == nil || IsNil(o.Categories) {
		return nil, false
	}
	return o.Categories, true
}

// HasCategories returns a boolean if a field has been set.
func (o *Plan) HasCategories() bool {
	if o != nil && !IsNil(o.Categories) {
		return true
	}

	return false
}

// SetCategories gets a reference to the given []string and assigns it to the Categories field.
func (o *Plan) SetCategories(v []string) {
	o.Categories = v
}

// GetClass returns the Class field value if set, zero value otherwise.
func (o *Plan) GetClass() string {
	if o == nil || IsNil(o.Class) {
		var ret string
		return ret
	}
	return *o.Class
}

// GetClassOk returns a tuple with the Class field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plan) GetClassOk() (*string, bool) {
	if o == nil || IsNil(o.Class) {
		return nil, false
	}
	return o.Class, true
}

// HasClass returns a boolean if a field has been set.
func (o *Plan) HasClass() bool {
	if o != nil && !IsNil(o.Class) {
		return true
	}

	return false
}

// SetClass gets a reference to the given string and assigns it to the Class field.
func (o *Plan) SetClass(v string) {
	o.Class = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Plan) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plan) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Plan) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Plan) SetDescription(v string) {
	o.Description = &v
}

// GetDeploymentTypes returns the DeploymentTypes field value if set, zero value otherwise.
func (o *Plan) GetDeploymentTypes() []string {
	if o == nil || IsNil(o.DeploymentTypes) {
		var ret []string
		return ret
	}
	return o.DeploymentTypes
}

// GetDeploymentTypesOk returns a tuple with the DeploymentTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plan) GetDeploymentTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.DeploymentTypes) {
		return nil, false
	}
	return o.DeploymentTypes, true
}

// HasDeploymentTypes returns a boolean if a field has been set.
func (o *Plan) HasDeploymentTypes() bool {
	if o != nil && !IsNil(o.DeploymentTypes) {
		return true
	}

	return false
}

// SetDeploymentTypes gets a reference to the given []string and assigns it to the DeploymentTypes field.
func (o *Plan) SetDeploymentTypes(v []string) {
	o.DeploymentTypes = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Plan) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plan) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Plan) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Plan) SetId(v string) {
	o.Id = &v
}

// GetLegacy returns the Legacy field value if set, zero value otherwise.
func (o *Plan) GetLegacy() bool {
	if o == nil || IsNil(o.Legacy) {
		var ret bool
		return ret
	}
	return *o.Legacy
}

// GetLegacyOk returns a tuple with the Legacy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plan) GetLegacyOk() (*bool, bool) {
	if o == nil || IsNil(o.Legacy) {
		return nil, false
	}
	return o.Legacy, true
}

// HasLegacy returns a boolean if a field has been set.
func (o *Plan) HasLegacy() bool {
	if o != nil && !IsNil(o.Legacy) {
		return true
	}

	return false
}

// SetLegacy gets a reference to the given bool and assigns it to the Legacy field.
func (o *Plan) SetLegacy(v bool) {
	o.Legacy = &v
}

// GetLine returns the Line field value if set, zero value otherwise.
func (o *Plan) GetLine() string {
	if o == nil || IsNil(o.Line) {
		var ret string
		return ret
	}
	return *o.Line
}

// GetLineOk returns a tuple with the Line field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plan) GetLineOk() (*string, bool) {
	if o == nil || IsNil(o.Line) {
		return nil, false
	}
	return o.Line, true
}

// HasLine returns a boolean if a field has been set.
func (o *Plan) HasLine() bool {
	if o != nil && !IsNil(o.Line) {
		return true
	}

	return false
}

// SetLine gets a reference to the given string and assigns it to the Line field.
func (o *Plan) SetLine(v string) {
	o.Line = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Plan) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plan) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Plan) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Plan) SetName(v string) {
	o.Name = &v
}

// GetPricing returns the Pricing field value if set, zero value otherwise.
func (o *Plan) GetPricing() map[string]interface{} {
	if o == nil || IsNil(o.Pricing) {
		var ret map[string]interface{}
		return ret
	}
	return o.Pricing
}

// GetPricingOk returns a tuple with the Pricing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plan) GetPricingOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Pricing) {
		return map[string]interface{}{}, false
	}
	return o.Pricing, true
}

// HasPricing returns a boolean if a field has been set.
func (o *Plan) HasPricing() bool {
	if o != nil && !IsNil(o.Pricing) {
		return true
	}

	return false
}

// SetPricing gets a reference to the given map[string]interface{} and assigns it to the Pricing field.
func (o *Plan) SetPricing(v map[string]interface{}) {
	o.Pricing = v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *Plan) GetSlug() string {
	if o == nil || IsNil(o.Slug) {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plan) GetSlugOk() (*string, bool) {
	if o == nil || IsNil(o.Slug) {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *Plan) HasSlug() bool {
	if o != nil && !IsNil(o.Slug) {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *Plan) SetSlug(v string) {
	o.Slug = &v
}

// GetSpecs returns the Specs field value if set, zero value otherwise.
func (o *Plan) GetSpecs() PlanSpecs {
	if o == nil || IsNil(o.Specs) {
		var ret PlanSpecs
		return ret
	}
	return *o.Specs
}

// GetSpecsOk returns a tuple with the Specs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plan) GetSpecsOk() (*PlanSpecs, bool) {
	if o == nil || IsNil(o.Specs) {
		return nil, false
	}
	return o.Specs, true
}

// HasSpecs returns a boolean if a field has been set.
func (o *Plan) HasSpecs() bool {
	if o != nil && !IsNil(o.Specs) {
		return true
	}

	return false
}

// SetSpecs gets a reference to the given PlanSpecs and assigns it to the Specs field.
func (o *Plan) SetSpecs(v PlanSpecs) {
	o.Specs = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Plan) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plan) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Plan) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Plan) SetType(v string) {
	o.Type = &v
}

func (o Plan) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Plan) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AvailableIn) {
		toSerialize["available_in"] = o.AvailableIn
	}
	if !IsNil(o.AvailableInMetros) {
		toSerialize["available_in_metros"] = o.AvailableInMetros
	}
	if !IsNil(o.Categories) {
		toSerialize["categories"] = o.Categories
	}
	if !IsNil(o.Class) {
		toSerialize["class"] = o.Class
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.DeploymentTypes) {
		toSerialize["deployment_types"] = o.DeploymentTypes
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Legacy) {
		toSerialize["legacy"] = o.Legacy
	}
	if !IsNil(o.Line) {
		toSerialize["line"] = o.Line
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Pricing) {
		toSerialize["pricing"] = o.Pricing
	}
	if !IsNil(o.Slug) {
		toSerialize["slug"] = o.Slug
	}
	if !IsNil(o.Specs) {
		toSerialize["specs"] = o.Specs
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Plan) UnmarshalJSON(bytes []byte) (err error) {
	varPlan := _Plan{}

	err = json.Unmarshal(bytes, &varPlan)

	if err != nil {
		return err
	}

	*o = Plan(varPlan)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "available_in")
		delete(additionalProperties, "available_in_metros")
		delete(additionalProperties, "categories")
		delete(additionalProperties, "class")
		delete(additionalProperties, "description")
		delete(additionalProperties, "deployment_types")
		delete(additionalProperties, "id")
		delete(additionalProperties, "legacy")
		delete(additionalProperties, "line")
		delete(additionalProperties, "name")
		delete(additionalProperties, "pricing")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "specs")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePlan struct {
	value *Plan
	isSet bool
}

func (v NullablePlan) Get() *Plan {
	return v.value
}

func (v *NullablePlan) Set(val *Plan) {
	v.value = val
	v.isSet = true
}

func (v NullablePlan) IsSet() bool {
	return v.isSet
}

func (v *NullablePlan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlan(val *Plan) *NullablePlan {
	return &NullablePlan{value: val, isSet: true}
}

func (v NullablePlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
