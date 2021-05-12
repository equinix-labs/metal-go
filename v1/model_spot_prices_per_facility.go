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

// SpotPricesPerFacility struct for SpotPricesPerFacility
type SpotPricesPerFacility struct {
	Baremetal2a *SpotPricesPerBaremetal `json:"baremetal_2a,omitempty"`
	Baremetal2a2 *SpotPricesPerBaremetal `json:"baremetal_2a2,omitempty"`
	Baremetal1 *SpotPricesPerBaremetal `json:"baremetal_1,omitempty"`
	Baremetal3 *SpotPricesPerBaremetal `json:"baremetal_3,omitempty"`
	C2MediumX86 *SpotPricesPerBaremetal `json:"c2.medium.x86,omitempty"`
	Baremetal2 *SpotPricesPerBaremetal `json:"baremetal_2,omitempty"`
	M2XlargeX86 *SpotPricesPerBaremetal `json:"m2.xlarge.x86,omitempty"`
	BaremetalS *SpotPricesPerBaremetal `json:"baremetal_s,omitempty"`
	Baremetal0 *SpotPricesPerBaremetal `json:"baremetal_0,omitempty"`
}

// NewSpotPricesPerFacility instantiates a new SpotPricesPerFacility object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpotPricesPerFacility() *SpotPricesPerFacility {
	this := SpotPricesPerFacility{}
	return &this
}

// NewSpotPricesPerFacilityWithDefaults instantiates a new SpotPricesPerFacility object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpotPricesPerFacilityWithDefaults() *SpotPricesPerFacility {
	this := SpotPricesPerFacility{}
	return &this
}

// GetBaremetal2a returns the Baremetal2a field value if set, zero value otherwise.
func (o *SpotPricesPerFacility) GetBaremetal2a() SpotPricesPerBaremetal {
	if o == nil || o.Baremetal2a == nil {
		var ret SpotPricesPerBaremetal
		return ret
	}
	return *o.Baremetal2a
}

// GetBaremetal2aOk returns a tuple with the Baremetal2a field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotPricesPerFacility) GetBaremetal2aOk() (*SpotPricesPerBaremetal, bool) {
	if o == nil || o.Baremetal2a == nil {
		return nil, false
	}
	return o.Baremetal2a, true
}

// HasBaremetal2a returns a boolean if a field has been set.
func (o *SpotPricesPerFacility) HasBaremetal2a() bool {
	if o != nil && o.Baremetal2a != nil {
		return true
	}

	return false
}

// SetBaremetal2a gets a reference to the given SpotPricesPerBaremetal and assigns it to the Baremetal2a field.
func (o *SpotPricesPerFacility) SetBaremetal2a(v SpotPricesPerBaremetal) {
	o.Baremetal2a = &v
}

// GetBaremetal2a2 returns the Baremetal2a2 field value if set, zero value otherwise.
func (o *SpotPricesPerFacility) GetBaremetal2a2() SpotPricesPerBaremetal {
	if o == nil || o.Baremetal2a2 == nil {
		var ret SpotPricesPerBaremetal
		return ret
	}
	return *o.Baremetal2a2
}

// GetBaremetal2a2Ok returns a tuple with the Baremetal2a2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotPricesPerFacility) GetBaremetal2a2Ok() (*SpotPricesPerBaremetal, bool) {
	if o == nil || o.Baremetal2a2 == nil {
		return nil, false
	}
	return o.Baremetal2a2, true
}

// HasBaremetal2a2 returns a boolean if a field has been set.
func (o *SpotPricesPerFacility) HasBaremetal2a2() bool {
	if o != nil && o.Baremetal2a2 != nil {
		return true
	}

	return false
}

// SetBaremetal2a2 gets a reference to the given SpotPricesPerBaremetal and assigns it to the Baremetal2a2 field.
func (o *SpotPricesPerFacility) SetBaremetal2a2(v SpotPricesPerBaremetal) {
	o.Baremetal2a2 = &v
}

// GetBaremetal1 returns the Baremetal1 field value if set, zero value otherwise.
func (o *SpotPricesPerFacility) GetBaremetal1() SpotPricesPerBaremetal {
	if o == nil || o.Baremetal1 == nil {
		var ret SpotPricesPerBaremetal
		return ret
	}
	return *o.Baremetal1
}

// GetBaremetal1Ok returns a tuple with the Baremetal1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotPricesPerFacility) GetBaremetal1Ok() (*SpotPricesPerBaremetal, bool) {
	if o == nil || o.Baremetal1 == nil {
		return nil, false
	}
	return o.Baremetal1, true
}

// HasBaremetal1 returns a boolean if a field has been set.
func (o *SpotPricesPerFacility) HasBaremetal1() bool {
	if o != nil && o.Baremetal1 != nil {
		return true
	}

	return false
}

// SetBaremetal1 gets a reference to the given SpotPricesPerBaremetal and assigns it to the Baremetal1 field.
func (o *SpotPricesPerFacility) SetBaremetal1(v SpotPricesPerBaremetal) {
	o.Baremetal1 = &v
}

// GetBaremetal3 returns the Baremetal3 field value if set, zero value otherwise.
func (o *SpotPricesPerFacility) GetBaremetal3() SpotPricesPerBaremetal {
	if o == nil || o.Baremetal3 == nil {
		var ret SpotPricesPerBaremetal
		return ret
	}
	return *o.Baremetal3
}

// GetBaremetal3Ok returns a tuple with the Baremetal3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotPricesPerFacility) GetBaremetal3Ok() (*SpotPricesPerBaremetal, bool) {
	if o == nil || o.Baremetal3 == nil {
		return nil, false
	}
	return o.Baremetal3, true
}

// HasBaremetal3 returns a boolean if a field has been set.
func (o *SpotPricesPerFacility) HasBaremetal3() bool {
	if o != nil && o.Baremetal3 != nil {
		return true
	}

	return false
}

// SetBaremetal3 gets a reference to the given SpotPricesPerBaremetal and assigns it to the Baremetal3 field.
func (o *SpotPricesPerFacility) SetBaremetal3(v SpotPricesPerBaremetal) {
	o.Baremetal3 = &v
}

// GetC2MediumX86 returns the C2MediumX86 field value if set, zero value otherwise.
func (o *SpotPricesPerFacility) GetC2MediumX86() SpotPricesPerBaremetal {
	if o == nil || o.C2MediumX86 == nil {
		var ret SpotPricesPerBaremetal
		return ret
	}
	return *o.C2MediumX86
}

// GetC2MediumX86Ok returns a tuple with the C2MediumX86 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotPricesPerFacility) GetC2MediumX86Ok() (*SpotPricesPerBaremetal, bool) {
	if o == nil || o.C2MediumX86 == nil {
		return nil, false
	}
	return o.C2MediumX86, true
}

// HasC2MediumX86 returns a boolean if a field has been set.
func (o *SpotPricesPerFacility) HasC2MediumX86() bool {
	if o != nil && o.C2MediumX86 != nil {
		return true
	}

	return false
}

// SetC2MediumX86 gets a reference to the given SpotPricesPerBaremetal and assigns it to the C2MediumX86 field.
func (o *SpotPricesPerFacility) SetC2MediumX86(v SpotPricesPerBaremetal) {
	o.C2MediumX86 = &v
}

// GetBaremetal2 returns the Baremetal2 field value if set, zero value otherwise.
func (o *SpotPricesPerFacility) GetBaremetal2() SpotPricesPerBaremetal {
	if o == nil || o.Baremetal2 == nil {
		var ret SpotPricesPerBaremetal
		return ret
	}
	return *o.Baremetal2
}

// GetBaremetal2Ok returns a tuple with the Baremetal2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotPricesPerFacility) GetBaremetal2Ok() (*SpotPricesPerBaremetal, bool) {
	if o == nil || o.Baremetal2 == nil {
		return nil, false
	}
	return o.Baremetal2, true
}

// HasBaremetal2 returns a boolean if a field has been set.
func (o *SpotPricesPerFacility) HasBaremetal2() bool {
	if o != nil && o.Baremetal2 != nil {
		return true
	}

	return false
}

// SetBaremetal2 gets a reference to the given SpotPricesPerBaremetal and assigns it to the Baremetal2 field.
func (o *SpotPricesPerFacility) SetBaremetal2(v SpotPricesPerBaremetal) {
	o.Baremetal2 = &v
}

// GetM2XlargeX86 returns the M2XlargeX86 field value if set, zero value otherwise.
func (o *SpotPricesPerFacility) GetM2XlargeX86() SpotPricesPerBaremetal {
	if o == nil || o.M2XlargeX86 == nil {
		var ret SpotPricesPerBaremetal
		return ret
	}
	return *o.M2XlargeX86
}

// GetM2XlargeX86Ok returns a tuple with the M2XlargeX86 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotPricesPerFacility) GetM2XlargeX86Ok() (*SpotPricesPerBaremetal, bool) {
	if o == nil || o.M2XlargeX86 == nil {
		return nil, false
	}
	return o.M2XlargeX86, true
}

// HasM2XlargeX86 returns a boolean if a field has been set.
func (o *SpotPricesPerFacility) HasM2XlargeX86() bool {
	if o != nil && o.M2XlargeX86 != nil {
		return true
	}

	return false
}

// SetM2XlargeX86 gets a reference to the given SpotPricesPerBaremetal and assigns it to the M2XlargeX86 field.
func (o *SpotPricesPerFacility) SetM2XlargeX86(v SpotPricesPerBaremetal) {
	o.M2XlargeX86 = &v
}

// GetBaremetalS returns the BaremetalS field value if set, zero value otherwise.
func (o *SpotPricesPerFacility) GetBaremetalS() SpotPricesPerBaremetal {
	if o == nil || o.BaremetalS == nil {
		var ret SpotPricesPerBaremetal
		return ret
	}
	return *o.BaremetalS
}

// GetBaremetalSOk returns a tuple with the BaremetalS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotPricesPerFacility) GetBaremetalSOk() (*SpotPricesPerBaremetal, bool) {
	if o == nil || o.BaremetalS == nil {
		return nil, false
	}
	return o.BaremetalS, true
}

// HasBaremetalS returns a boolean if a field has been set.
func (o *SpotPricesPerFacility) HasBaremetalS() bool {
	if o != nil && o.BaremetalS != nil {
		return true
	}

	return false
}

// SetBaremetalS gets a reference to the given SpotPricesPerBaremetal and assigns it to the BaremetalS field.
func (o *SpotPricesPerFacility) SetBaremetalS(v SpotPricesPerBaremetal) {
	o.BaremetalS = &v
}

// GetBaremetal0 returns the Baremetal0 field value if set, zero value otherwise.
func (o *SpotPricesPerFacility) GetBaremetal0() SpotPricesPerBaremetal {
	if o == nil || o.Baremetal0 == nil {
		var ret SpotPricesPerBaremetal
		return ret
	}
	return *o.Baremetal0
}

// GetBaremetal0Ok returns a tuple with the Baremetal0 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotPricesPerFacility) GetBaremetal0Ok() (*SpotPricesPerBaremetal, bool) {
	if o == nil || o.Baremetal0 == nil {
		return nil, false
	}
	return o.Baremetal0, true
}

// HasBaremetal0 returns a boolean if a field has been set.
func (o *SpotPricesPerFacility) HasBaremetal0() bool {
	if o != nil && o.Baremetal0 != nil {
		return true
	}

	return false
}

// SetBaremetal0 gets a reference to the given SpotPricesPerBaremetal and assigns it to the Baremetal0 field.
func (o *SpotPricesPerFacility) SetBaremetal0(v SpotPricesPerBaremetal) {
	o.Baremetal0 = &v
}

func (o SpotPricesPerFacility) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Baremetal2a != nil {
		toSerialize["baremetal_2a"] = o.Baremetal2a
	}
	if o.Baremetal2a2 != nil {
		toSerialize["baremetal_2a2"] = o.Baremetal2a2
	}
	if o.Baremetal1 != nil {
		toSerialize["baremetal_1"] = o.Baremetal1
	}
	if o.Baremetal3 != nil {
		toSerialize["baremetal_3"] = o.Baremetal3
	}
	if o.C2MediumX86 != nil {
		toSerialize["c2.medium.x86"] = o.C2MediumX86
	}
	if o.Baremetal2 != nil {
		toSerialize["baremetal_2"] = o.Baremetal2
	}
	if o.M2XlargeX86 != nil {
		toSerialize["m2.xlarge.x86"] = o.M2XlargeX86
	}
	if o.BaremetalS != nil {
		toSerialize["baremetal_s"] = o.BaremetalS
	}
	if o.Baremetal0 != nil {
		toSerialize["baremetal_0"] = o.Baremetal0
	}
	return json.Marshal(toSerialize)
}

type NullableSpotPricesPerFacility struct {
	value *SpotPricesPerFacility
	isSet bool
}

func (v NullableSpotPricesPerFacility) Get() *SpotPricesPerFacility {
	return v.value
}

func (v *NullableSpotPricesPerFacility) Set(val *SpotPricesPerFacility) {
	v.value = val
	v.isSet = true
}

func (v NullableSpotPricesPerFacility) IsSet() bool {
	return v.isSet
}

func (v *NullableSpotPricesPerFacility) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpotPricesPerFacility(val *SpotPricesPerFacility) *NullableSpotPricesPerFacility {
	return &NullableSpotPricesPerFacility{value: val, isSet: true}
}

func (v NullableSpotPricesPerFacility) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpotPricesPerFacility) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


