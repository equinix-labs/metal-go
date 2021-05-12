/*
 * Metal API
 *
 * This is the API for Equinix Metal. The API allows you to programmatically interact with all of your Equinix Metal resources, including devices, networks, addresses, organizations, projects, and your user account.  The official API docs are hosted at <https://metal.equinix.com/developers/api>. 
 *
 * API version: 1.0.0
 * Contact: support@equinixmetal.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
)

// VirtualNetwork struct for VirtualNetwork
type VirtualNetwork struct {
	Id *string `json:"id,omitempty"`
	Description *string `json:"description,omitempty"`
	Vxlan *int32 `json:"vxlan,omitempty"`
	Facility *Href `json:"facility,omitempty"`
	// A list of instances with ports currently associated to this Virtual Network.
	Instances *[]Href `json:"instances,omitempty"`
	// The Metro code of the metro in which this Virtual Network is defined.
	MetroCode *string `json:"metro_code,omitempty"`
	Metro *Href `json:"metro,omitempty"`
	// True if the virtual network is attached to a virtual circuit. False if not.
	AssignedToVirtualCircuit *bool `json:"assigned_to_virtual_circuit,omitempty"`
	AssignedTo *Href `json:"assigned_to,omitempty"`
	Href *string `json:"href,omitempty"`
}

// NewVirtualNetwork instantiates a new VirtualNetwork object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualNetwork() *VirtualNetwork {
	this := VirtualNetwork{}
	return &this
}

// NewVirtualNetworkWithDefaults instantiates a new VirtualNetwork object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualNetworkWithDefaults() *VirtualNetwork {
	this := VirtualNetwork{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VirtualNetwork) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualNetwork) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VirtualNetwork) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *VirtualNetwork) SetId(v string) {
	o.Id = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *VirtualNetwork) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualNetwork) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *VirtualNetwork) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *VirtualNetwork) SetDescription(v string) {
	o.Description = &v
}

// GetVxlan returns the Vxlan field value if set, zero value otherwise.
func (o *VirtualNetwork) GetVxlan() int32 {
	if o == nil || o.Vxlan == nil {
		var ret int32
		return ret
	}
	return *o.Vxlan
}

// GetVxlanOk returns a tuple with the Vxlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualNetwork) GetVxlanOk() (*int32, bool) {
	if o == nil || o.Vxlan == nil {
		return nil, false
	}
	return o.Vxlan, true
}

// HasVxlan returns a boolean if a field has been set.
func (o *VirtualNetwork) HasVxlan() bool {
	if o != nil && o.Vxlan != nil {
		return true
	}

	return false
}

// SetVxlan gets a reference to the given int32 and assigns it to the Vxlan field.
func (o *VirtualNetwork) SetVxlan(v int32) {
	o.Vxlan = &v
}

// GetFacility returns the Facility field value if set, zero value otherwise.
func (o *VirtualNetwork) GetFacility() Href {
	if o == nil || o.Facility == nil {
		var ret Href
		return ret
	}
	return *o.Facility
}

// GetFacilityOk returns a tuple with the Facility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualNetwork) GetFacilityOk() (*Href, bool) {
	if o == nil || o.Facility == nil {
		return nil, false
	}
	return o.Facility, true
}

// HasFacility returns a boolean if a field has been set.
func (o *VirtualNetwork) HasFacility() bool {
	if o != nil && o.Facility != nil {
		return true
	}

	return false
}

// SetFacility gets a reference to the given Href and assigns it to the Facility field.
func (o *VirtualNetwork) SetFacility(v Href) {
	o.Facility = &v
}

// GetInstances returns the Instances field value if set, zero value otherwise.
func (o *VirtualNetwork) GetInstances() []Href {
	if o == nil || o.Instances == nil {
		var ret []Href
		return ret
	}
	return *o.Instances
}

// GetInstancesOk returns a tuple with the Instances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualNetwork) GetInstancesOk() (*[]Href, bool) {
	if o == nil || o.Instances == nil {
		return nil, false
	}
	return o.Instances, true
}

// HasInstances returns a boolean if a field has been set.
func (o *VirtualNetwork) HasInstances() bool {
	if o != nil && o.Instances != nil {
		return true
	}

	return false
}

// SetInstances gets a reference to the given []Href and assigns it to the Instances field.
func (o *VirtualNetwork) SetInstances(v []Href) {
	o.Instances = &v
}

// GetMetroCode returns the MetroCode field value if set, zero value otherwise.
func (o *VirtualNetwork) GetMetroCode() string {
	if o == nil || o.MetroCode == nil {
		var ret string
		return ret
	}
	return *o.MetroCode
}

// GetMetroCodeOk returns a tuple with the MetroCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualNetwork) GetMetroCodeOk() (*string, bool) {
	if o == nil || o.MetroCode == nil {
		return nil, false
	}
	return o.MetroCode, true
}

// HasMetroCode returns a boolean if a field has been set.
func (o *VirtualNetwork) HasMetroCode() bool {
	if o != nil && o.MetroCode != nil {
		return true
	}

	return false
}

// SetMetroCode gets a reference to the given string and assigns it to the MetroCode field.
func (o *VirtualNetwork) SetMetroCode(v string) {
	o.MetroCode = &v
}

// GetMetro returns the Metro field value if set, zero value otherwise.
func (o *VirtualNetwork) GetMetro() Href {
	if o == nil || o.Metro == nil {
		var ret Href
		return ret
	}
	return *o.Metro
}

// GetMetroOk returns a tuple with the Metro field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualNetwork) GetMetroOk() (*Href, bool) {
	if o == nil || o.Metro == nil {
		return nil, false
	}
	return o.Metro, true
}

// HasMetro returns a boolean if a field has been set.
func (o *VirtualNetwork) HasMetro() bool {
	if o != nil && o.Metro != nil {
		return true
	}

	return false
}

// SetMetro gets a reference to the given Href and assigns it to the Metro field.
func (o *VirtualNetwork) SetMetro(v Href) {
	o.Metro = &v
}

// GetAssignedToVirtualCircuit returns the AssignedToVirtualCircuit field value if set, zero value otherwise.
func (o *VirtualNetwork) GetAssignedToVirtualCircuit() bool {
	if o == nil || o.AssignedToVirtualCircuit == nil {
		var ret bool
		return ret
	}
	return *o.AssignedToVirtualCircuit
}

// GetAssignedToVirtualCircuitOk returns a tuple with the AssignedToVirtualCircuit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualNetwork) GetAssignedToVirtualCircuitOk() (*bool, bool) {
	if o == nil || o.AssignedToVirtualCircuit == nil {
		return nil, false
	}
	return o.AssignedToVirtualCircuit, true
}

// HasAssignedToVirtualCircuit returns a boolean if a field has been set.
func (o *VirtualNetwork) HasAssignedToVirtualCircuit() bool {
	if o != nil && o.AssignedToVirtualCircuit != nil {
		return true
	}

	return false
}

// SetAssignedToVirtualCircuit gets a reference to the given bool and assigns it to the AssignedToVirtualCircuit field.
func (o *VirtualNetwork) SetAssignedToVirtualCircuit(v bool) {
	o.AssignedToVirtualCircuit = &v
}

// GetAssignedTo returns the AssignedTo field value if set, zero value otherwise.
func (o *VirtualNetwork) GetAssignedTo() Href {
	if o == nil || o.AssignedTo == nil {
		var ret Href
		return ret
	}
	return *o.AssignedTo
}

// GetAssignedToOk returns a tuple with the AssignedTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualNetwork) GetAssignedToOk() (*Href, bool) {
	if o == nil || o.AssignedTo == nil {
		return nil, false
	}
	return o.AssignedTo, true
}

// HasAssignedTo returns a boolean if a field has been set.
func (o *VirtualNetwork) HasAssignedTo() bool {
	if o != nil && o.AssignedTo != nil {
		return true
	}

	return false
}

// SetAssignedTo gets a reference to the given Href and assigns it to the AssignedTo field.
func (o *VirtualNetwork) SetAssignedTo(v Href) {
	o.AssignedTo = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *VirtualNetwork) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualNetwork) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *VirtualNetwork) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *VirtualNetwork) SetHref(v string) {
	o.Href = &v
}

func (o VirtualNetwork) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Vxlan != nil {
		toSerialize["vxlan"] = o.Vxlan
	}
	if o.Facility != nil {
		toSerialize["facility"] = o.Facility
	}
	if o.Instances != nil {
		toSerialize["instances"] = o.Instances
	}
	if o.MetroCode != nil {
		toSerialize["metro_code"] = o.MetroCode
	}
	if o.Metro != nil {
		toSerialize["metro"] = o.Metro
	}
	if o.AssignedToVirtualCircuit != nil {
		toSerialize["assigned_to_virtual_circuit"] = o.AssignedToVirtualCircuit
	}
	if o.AssignedTo != nil {
		toSerialize["assigned_to"] = o.AssignedTo
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	return json.Marshal(toSerialize)
}

type NullableVirtualNetwork struct {
	value *VirtualNetwork
	isSet bool
}

func (v NullableVirtualNetwork) Get() *VirtualNetwork {
	return v.value
}

func (v *NullableVirtualNetwork) Set(val *VirtualNetwork) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualNetwork(val *VirtualNetwork) *NullableVirtualNetwork {
	return &NullableVirtualNetwork{value: val, isSet: true}
}

func (v NullableVirtualNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


