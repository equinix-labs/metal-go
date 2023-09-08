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

// checks if the VrfFabricVcCreateInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VrfFabricVcCreateInput{}

// VrfFabricVcCreateInput struct for VrfFabricVcCreateInput
type VrfFabricVcCreateInput struct {
	// The preferred email used for communication and notifications about the Equinix Fabric interconnection. Required when using a Project API key. Optional and defaults to the primary user email address when using a User API key.
	ContactEmail *string `json:"contact_email,omitempty"`
	Description  *string `json:"description,omitempty"`
	// A Metro ID or code. When creating Fabric VCs (Metal Billed), this is where interconnection will be originating from, as we pre-authorize the use of one of our shared ports as the origin of the interconnection using A-Side service tokens. We only allow local connections for Fabric VCs (Metal Billed), so the destination location must be the same as the origin. For Fabric VCs (Fabric Billed), or shared connections, this will be the destination of the interconnection. We allow remote connections for Fabric VCs (Fabric Billed), so the origin of the interconnection can be a different metro set here.
	Metro   string  `json:"metro"`
	Name    string  `json:"name"`
	Project *string `json:"project,omitempty"`
	// Either 'primary' or 'redundant'.
	Redundancy string `json:"redundancy"`
	// Either 'a_side' or 'z_side'. Setting this field to 'a_side' will create an interconnection with Fabric VCs (Metal Billed). Setting this field to 'z_side' will create an interconnection with Fabric VCs (Fabric Billed). This is required when the 'type' is 'shared', but this is not applicable when the 'type' is 'dedicated'. This parameter is included in the specification as a developer preview and is generally unavailable. Please contact our Support team for more details.
	ServiceTokenType string `json:"service_token_type"`
	// A interconnection speed, in bps, mbps, or gbps. For Fabric VCs, this represents the maximum speed of the interconnection. For Fabric VCs (Metal Billed), this can only be one of the following:  ''50mbps'', ''200mbps'', ''500mbps'', ''1gbps'', ''2gbps'', ''5gbps'' or ''10gbps'', and is required for creation. For Fabric VCs (Fabric Billed), this field will always default to ''10gbps'' even if it is not provided. For example, ''500000000'', ''50m'', or' ''500mbps'' will all work as valid inputs.
	Speed *int32   `json:"speed,omitempty"`
	Tags  []string `json:"tags,omitempty"`
	// When requesting for a Fabric VC, the value of this field should be 'shared'.
	Type string `json:"type"`
	// This field holds a list of VRF UUIDs that will be set automatically on the virtual circuits of Fabric VCs on creation, and can hold up to two UUIDs. Two UUIDs are required when requesting redundant Fabric VCs. The first UUID will be set on the primary virtual circuit, while the second UUID will be set on the secondary. The two UUIDs can be the same if both the primary and secondary virtual circuits will be in the same VRF. This parameter is included in the specification as a developer preview and is generally unavailable. Please contact our Support team for more details.
	Vrfs                 []string `json:"vrfs"`
	AdditionalProperties map[string]interface{}
}

type _VrfFabricVcCreateInput VrfFabricVcCreateInput

// NewVrfFabricVcCreateInput instantiates a new VrfFabricVcCreateInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVrfFabricVcCreateInput(metro string, name string, redundancy string, serviceTokenType string, type_ string, vrfs []string) *VrfFabricVcCreateInput {
	this := VrfFabricVcCreateInput{}
	this.Metro = metro
	this.Name = name
	this.Redundancy = redundancy
	this.ServiceTokenType = serviceTokenType
	this.Type = type_
	this.Vrfs = vrfs
	return &this
}

// NewVrfFabricVcCreateInputWithDefaults instantiates a new VrfFabricVcCreateInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVrfFabricVcCreateInputWithDefaults() *VrfFabricVcCreateInput {
	this := VrfFabricVcCreateInput{}
	return &this
}

// GetContactEmail returns the ContactEmail field value if set, zero value otherwise.
func (o *VrfFabricVcCreateInput) GetContactEmail() string {
	if o == nil || IsNil(o.ContactEmail) {
		var ret string
		return ret
	}
	return *o.ContactEmail
}

// GetContactEmailOk returns a tuple with the ContactEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfFabricVcCreateInput) GetContactEmailOk() (*string, bool) {
	if o == nil || IsNil(o.ContactEmail) {
		return nil, false
	}
	return o.ContactEmail, true
}

// HasContactEmail returns a boolean if a field has been set.
func (o *VrfFabricVcCreateInput) HasContactEmail() bool {
	if o != nil && !IsNil(o.ContactEmail) {
		return true
	}

	return false
}

// SetContactEmail gets a reference to the given string and assigns it to the ContactEmail field.
func (o *VrfFabricVcCreateInput) SetContactEmail(v string) {
	o.ContactEmail = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *VrfFabricVcCreateInput) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfFabricVcCreateInput) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *VrfFabricVcCreateInput) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *VrfFabricVcCreateInput) SetDescription(v string) {
	o.Description = &v
}

// GetMetro returns the Metro field value
func (o *VrfFabricVcCreateInput) GetMetro() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Metro
}

// GetMetroOk returns a tuple with the Metro field value
// and a boolean to check if the value has been set.
func (o *VrfFabricVcCreateInput) GetMetroOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metro, true
}

// SetMetro sets field value
func (o *VrfFabricVcCreateInput) SetMetro(v string) {
	o.Metro = v
}

// GetName returns the Name field value
func (o *VrfFabricVcCreateInput) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VrfFabricVcCreateInput) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VrfFabricVcCreateInput) SetName(v string) {
	o.Name = v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *VrfFabricVcCreateInput) GetProject() string {
	if o == nil || IsNil(o.Project) {
		var ret string
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfFabricVcCreateInput) GetProjectOk() (*string, bool) {
	if o == nil || IsNil(o.Project) {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *VrfFabricVcCreateInput) HasProject() bool {
	if o != nil && !IsNil(o.Project) {
		return true
	}

	return false
}

// SetProject gets a reference to the given string and assigns it to the Project field.
func (o *VrfFabricVcCreateInput) SetProject(v string) {
	o.Project = &v
}

// GetRedundancy returns the Redundancy field value
func (o *VrfFabricVcCreateInput) GetRedundancy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Redundancy
}

// GetRedundancyOk returns a tuple with the Redundancy field value
// and a boolean to check if the value has been set.
func (o *VrfFabricVcCreateInput) GetRedundancyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Redundancy, true
}

// SetRedundancy sets field value
func (o *VrfFabricVcCreateInput) SetRedundancy(v string) {
	o.Redundancy = v
}

// GetServiceTokenType returns the ServiceTokenType field value
func (o *VrfFabricVcCreateInput) GetServiceTokenType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceTokenType
}

// GetServiceTokenTypeOk returns a tuple with the ServiceTokenType field value
// and a boolean to check if the value has been set.
func (o *VrfFabricVcCreateInput) GetServiceTokenTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceTokenType, true
}

// SetServiceTokenType sets field value
func (o *VrfFabricVcCreateInput) SetServiceTokenType(v string) {
	o.ServiceTokenType = v
}

// GetSpeed returns the Speed field value if set, zero value otherwise.
func (o *VrfFabricVcCreateInput) GetSpeed() int32 {
	if o == nil || IsNil(o.Speed) {
		var ret int32
		return ret
	}
	return *o.Speed
}

// GetSpeedOk returns a tuple with the Speed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfFabricVcCreateInput) GetSpeedOk() (*int32, bool) {
	if o == nil || IsNil(o.Speed) {
		return nil, false
	}
	return o.Speed, true
}

// HasSpeed returns a boolean if a field has been set.
func (o *VrfFabricVcCreateInput) HasSpeed() bool {
	if o != nil && !IsNil(o.Speed) {
		return true
	}

	return false
}

// SetSpeed gets a reference to the given int32 and assigns it to the Speed field.
func (o *VrfFabricVcCreateInput) SetSpeed(v int32) {
	o.Speed = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *VrfFabricVcCreateInput) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfFabricVcCreateInput) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *VrfFabricVcCreateInput) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *VrfFabricVcCreateInput) SetTags(v []string) {
	o.Tags = v
}

// GetType returns the Type field value
func (o *VrfFabricVcCreateInput) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *VrfFabricVcCreateInput) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *VrfFabricVcCreateInput) SetType(v string) {
	o.Type = v
}

// GetVrfs returns the Vrfs field value
func (o *VrfFabricVcCreateInput) GetVrfs() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Vrfs
}

// GetVrfsOk returns a tuple with the Vrfs field value
// and a boolean to check if the value has been set.
func (o *VrfFabricVcCreateInput) GetVrfsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vrfs, true
}

// SetVrfs sets field value
func (o *VrfFabricVcCreateInput) SetVrfs(v []string) {
	o.Vrfs = v
}

func (o VrfFabricVcCreateInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VrfFabricVcCreateInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContactEmail) {
		toSerialize["contact_email"] = o.ContactEmail
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["metro"] = o.Metro
	toSerialize["name"] = o.Name
	if !IsNil(o.Project) {
		toSerialize["project"] = o.Project
	}
	toSerialize["redundancy"] = o.Redundancy
	toSerialize["service_token_type"] = o.ServiceTokenType
	if !IsNil(o.Speed) {
		toSerialize["speed"] = o.Speed
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	toSerialize["type"] = o.Type
	toSerialize["vrfs"] = o.Vrfs

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VrfFabricVcCreateInput) UnmarshalJSON(bytes []byte) (err error) {
	varVrfFabricVcCreateInput := _VrfFabricVcCreateInput{}

	if err = json.Unmarshal(bytes, &varVrfFabricVcCreateInput); err == nil {
		*o = VrfFabricVcCreateInput(varVrfFabricVcCreateInput)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "contact_email")
		delete(additionalProperties, "description")
		delete(additionalProperties, "metro")
		delete(additionalProperties, "name")
		delete(additionalProperties, "project")
		delete(additionalProperties, "redundancy")
		delete(additionalProperties, "service_token_type")
		delete(additionalProperties, "speed")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "type")
		delete(additionalProperties, "vrfs")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVrfFabricVcCreateInput struct {
	value *VrfFabricVcCreateInput
	isSet bool
}

func (v NullableVrfFabricVcCreateInput) Get() *VrfFabricVcCreateInput {
	return v.value
}

func (v *NullableVrfFabricVcCreateInput) Set(val *VrfFabricVcCreateInput) {
	v.value = val
	v.isSet = true
}

func (v NullableVrfFabricVcCreateInput) IsSet() bool {
	return v.isSet
}

func (v *NullableVrfFabricVcCreateInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVrfFabricVcCreateInput(val *VrfFabricVcCreateInput) *NullableVrfFabricVcCreateInput {
	return &NullableVrfFabricVcCreateInput{value: val, isSet: true}
}

func (v NullableVrfFabricVcCreateInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVrfFabricVcCreateInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}