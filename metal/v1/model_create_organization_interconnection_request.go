/*
Metal API

This is the API for Equinix Metal. The API allows you to programmatically interact with all of your Equinix Metal resources, including devices, networks, addresses, organizations, projects, and your user account.  The official API docs are hosted at <https://metal.equinix.com/developers/api>.

API version: 1.0.0
Contact: support@equinixmetal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
)

// CreateOrganizationInterconnectionRequest struct for CreateOrganizationInterconnectionRequest
type CreateOrganizationInterconnectionRequest struct {
	Tags         []string `json:"tags,omitempty"`
	ContactEmail *string  `json:"contact_email,omitempty"`
	Description  *string  `json:"description,omitempty"`
	// A Metro ID or code.
	Metro string `json:"metro"`
	// The mode of the connection (only relevant to dedicated connections). Shared connections won't have this field. Can be either 'standard' or 'tunnel'.   The default mode of a dedicated connection is 'standard'. The mode can only be changed when there are no associated virtual circuits on the connection.   In tunnel mode, an 802.1q tunnel is added to a port to send/receive double tagged packets from server instances.
	Mode    *string `json:"mode,omitempty"`
	Name    string  `json:"name"`
	Project *string `json:"project,omitempty"`
	// Either 'primary' or 'redundant'.
	Redundancy string `json:"redundancy"`
	// Can only be set to 'a_side' to create a shared connection with an A-side Fabric service token. This parameter is included in the specification as a developer preview and is generally unavailable. Please contact our Support team for more details.
	ServiceTokenType *string `json:"service_token_type,omitempty"`
	// A connection speed, in bps, mbps, or gbps. Ex: '100000000' or '100 mbps'.
	Speed *string `json:"speed,omitempty"`
	// Either 'shared' or 'dedicated'.
	Type string `json:"type"`
	// A list of one or two metro-based VLANs that will be set on the primary and/or secondary (if redundant) virtual circuits respectively when creating a shared connection. VLANs can also be set after the connection is created, but are required to activate the connection. This parameter is included in the specification as a developer preview and is generally unavailable. Please contact our Support team for more details.
	Vlans []int32 `json:"vlans,omitempty"`
}

// NewCreateOrganizationInterconnectionRequest instantiates a new CreateOrganizationInterconnectionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrganizationInterconnectionRequest(metro string, name string, redundancy string, type_ string) *CreateOrganizationInterconnectionRequest {
	this := CreateOrganizationInterconnectionRequest{}
	this.Metro = metro
	this.Name = name
	this.Redundancy = redundancy
	this.Type = type_
	return &this
}

// NewCreateOrganizationInterconnectionRequestWithDefaults instantiates a new CreateOrganizationInterconnectionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrganizationInterconnectionRequestWithDefaults() *CreateOrganizationInterconnectionRequest {
	this := CreateOrganizationInterconnectionRequest{}
	return &this
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *CreateOrganizationInterconnectionRequest) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationInterconnectionRequest) GetTagsOk() ([]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CreateOrganizationInterconnectionRequest) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *CreateOrganizationInterconnectionRequest) SetTags(v []string) {
	o.Tags = v
}

// GetContactEmail returns the ContactEmail field value if set, zero value otherwise.
func (o *CreateOrganizationInterconnectionRequest) GetContactEmail() string {
	if o == nil || o.ContactEmail == nil {
		var ret string
		return ret
	}
	return *o.ContactEmail
}

// GetContactEmailOk returns a tuple with the ContactEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationInterconnectionRequest) GetContactEmailOk() (*string, bool) {
	if o == nil || o.ContactEmail == nil {
		return nil, false
	}
	return o.ContactEmail, true
}

// HasContactEmail returns a boolean if a field has been set.
func (o *CreateOrganizationInterconnectionRequest) HasContactEmail() bool {
	if o != nil && o.ContactEmail != nil {
		return true
	}

	return false
}

// SetContactEmail gets a reference to the given string and assigns it to the ContactEmail field.
func (o *CreateOrganizationInterconnectionRequest) SetContactEmail(v string) {
	o.ContactEmail = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateOrganizationInterconnectionRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationInterconnectionRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateOrganizationInterconnectionRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateOrganizationInterconnectionRequest) SetDescription(v string) {
	o.Description = &v
}

// GetMetro returns the Metro field value
func (o *CreateOrganizationInterconnectionRequest) GetMetro() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Metro
}

// GetMetroOk returns a tuple with the Metro field value
// and a boolean to check if the value has been set.
func (o *CreateOrganizationInterconnectionRequest) GetMetroOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metro, true
}

// SetMetro sets field value
func (o *CreateOrganizationInterconnectionRequest) SetMetro(v string) {
	o.Metro = v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *CreateOrganizationInterconnectionRequest) GetMode() string {
	if o == nil || o.Mode == nil {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationInterconnectionRequest) GetModeOk() (*string, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *CreateOrganizationInterconnectionRequest) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *CreateOrganizationInterconnectionRequest) SetMode(v string) {
	o.Mode = &v
}

// GetName returns the Name field value
func (o *CreateOrganizationInterconnectionRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateOrganizationInterconnectionRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateOrganizationInterconnectionRequest) SetName(v string) {
	o.Name = v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *CreateOrganizationInterconnectionRequest) GetProject() string {
	if o == nil || o.Project == nil {
		var ret string
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationInterconnectionRequest) GetProjectOk() (*string, bool) {
	if o == nil || o.Project == nil {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *CreateOrganizationInterconnectionRequest) HasProject() bool {
	if o != nil && o.Project != nil {
		return true
	}

	return false
}

// SetProject gets a reference to the given string and assigns it to the Project field.
func (o *CreateOrganizationInterconnectionRequest) SetProject(v string) {
	o.Project = &v
}

// GetRedundancy returns the Redundancy field value
func (o *CreateOrganizationInterconnectionRequest) GetRedundancy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Redundancy
}

// GetRedundancyOk returns a tuple with the Redundancy field value
// and a boolean to check if the value has been set.
func (o *CreateOrganizationInterconnectionRequest) GetRedundancyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Redundancy, true
}

// SetRedundancy sets field value
func (o *CreateOrganizationInterconnectionRequest) SetRedundancy(v string) {
	o.Redundancy = v
}

// GetServiceTokenType returns the ServiceTokenType field value if set, zero value otherwise.
func (o *CreateOrganizationInterconnectionRequest) GetServiceTokenType() string {
	if o == nil || o.ServiceTokenType == nil {
		var ret string
		return ret
	}
	return *o.ServiceTokenType
}

// GetServiceTokenTypeOk returns a tuple with the ServiceTokenType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationInterconnectionRequest) GetServiceTokenTypeOk() (*string, bool) {
	if o == nil || o.ServiceTokenType == nil {
		return nil, false
	}
	return o.ServiceTokenType, true
}

// HasServiceTokenType returns a boolean if a field has been set.
func (o *CreateOrganizationInterconnectionRequest) HasServiceTokenType() bool {
	if o != nil && o.ServiceTokenType != nil {
		return true
	}

	return false
}

// SetServiceTokenType gets a reference to the given string and assigns it to the ServiceTokenType field.
func (o *CreateOrganizationInterconnectionRequest) SetServiceTokenType(v string) {
	o.ServiceTokenType = &v
}

// GetSpeed returns the Speed field value if set, zero value otherwise.
func (o *CreateOrganizationInterconnectionRequest) GetSpeed() string {
	if o == nil || o.Speed == nil {
		var ret string
		return ret
	}
	return *o.Speed
}

// GetSpeedOk returns a tuple with the Speed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationInterconnectionRequest) GetSpeedOk() (*string, bool) {
	if o == nil || o.Speed == nil {
		return nil, false
	}
	return o.Speed, true
}

// HasSpeed returns a boolean if a field has been set.
func (o *CreateOrganizationInterconnectionRequest) HasSpeed() bool {
	if o != nil && o.Speed != nil {
		return true
	}

	return false
}

// SetSpeed gets a reference to the given string and assigns it to the Speed field.
func (o *CreateOrganizationInterconnectionRequest) SetSpeed(v string) {
	o.Speed = &v
}

// GetType returns the Type field value
func (o *CreateOrganizationInterconnectionRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreateOrganizationInterconnectionRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CreateOrganizationInterconnectionRequest) SetType(v string) {
	o.Type = v
}

// GetVlans returns the Vlans field value if set, zero value otherwise.
func (o *CreateOrganizationInterconnectionRequest) GetVlans() []int32 {
	if o == nil || o.Vlans == nil {
		var ret []int32
		return ret
	}
	return o.Vlans
}

// GetVlansOk returns a tuple with the Vlans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationInterconnectionRequest) GetVlansOk() ([]int32, bool) {
	if o == nil || o.Vlans == nil {
		return nil, false
	}
	return o.Vlans, true
}

// HasVlans returns a boolean if a field has been set.
func (o *CreateOrganizationInterconnectionRequest) HasVlans() bool {
	if o != nil && o.Vlans != nil {
		return true
	}

	return false
}

// SetVlans gets a reference to the given []int32 and assigns it to the Vlans field.
func (o *CreateOrganizationInterconnectionRequest) SetVlans(v []int32) {
	o.Vlans = v
}

func (o CreateOrganizationInterconnectionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.ContactEmail != nil {
		toSerialize["contact_email"] = o.ContactEmail
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["metro"] = o.Metro
	}
	if o.Mode != nil {
		toSerialize["mode"] = o.Mode
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Project != nil {
		toSerialize["project"] = o.Project
	}
	if true {
		toSerialize["redundancy"] = o.Redundancy
	}
	if o.ServiceTokenType != nil {
		toSerialize["service_token_type"] = o.ServiceTokenType
	}
	if o.Speed != nil {
		toSerialize["speed"] = o.Speed
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.Vlans != nil {
		toSerialize["vlans"] = o.Vlans
	}
	return json.Marshal(toSerialize)
}

type NullableCreateOrganizationInterconnectionRequest struct {
	value *CreateOrganizationInterconnectionRequest
	isSet bool
}

func (v NullableCreateOrganizationInterconnectionRequest) Get() *CreateOrganizationInterconnectionRequest {
	return v.value
}

func (v *NullableCreateOrganizationInterconnectionRequest) Set(val *CreateOrganizationInterconnectionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrganizationInterconnectionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrganizationInterconnectionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrganizationInterconnectionRequest(val *CreateOrganizationInterconnectionRequest) *NullableCreateOrganizationInterconnectionRequest {
	return &NullableCreateOrganizationInterconnectionRequest{value: val, isSet: true}
}

func (v NullableCreateOrganizationInterconnectionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrganizationInterconnectionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}