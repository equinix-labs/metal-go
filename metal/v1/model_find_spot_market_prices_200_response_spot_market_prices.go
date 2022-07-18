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
	if o == nil || o.Ams1 == nil {
		var ret FindSpotMarketPrices200ResponseSpotMarketPricesAms1
		return ret
	}
	return *o.Ams1
}

// GetAms1Ok returns a tuple with the Ams1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) GetAms1Ok() (*FindSpotMarketPrices200ResponseSpotMarketPricesAms1, bool) {
	if o == nil || o.Ams1 == nil {
		return nil, false
	}
	return o.Ams1, true
}

// HasAms1 returns a boolean if a field has been set.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) HasAms1() bool {
	if o != nil && o.Ams1 != nil {
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
	if o == nil || o.Atl1 == nil {
		var ret FindSpotMarketPrices200ResponseSpotMarketPricesAtl1
		return ret
	}
	return *o.Atl1
}

// GetAtl1Ok returns a tuple with the Atl1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) GetAtl1Ok() (*FindSpotMarketPrices200ResponseSpotMarketPricesAtl1, bool) {
	if o == nil || o.Atl1 == nil {
		return nil, false
	}
	return o.Atl1, true
}

// HasAtl1 returns a boolean if a field has been set.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) HasAtl1() bool {
	if o != nil && o.Atl1 != nil {
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
	if o == nil || o.Dfw1 == nil {
		var ret FindSpotMarketPrices200ResponseSpotMarketPricesAtl1
		return ret
	}
	return *o.Dfw1
}

// GetDfw1Ok returns a tuple with the Dfw1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) GetDfw1Ok() (*FindSpotMarketPrices200ResponseSpotMarketPricesAtl1, bool) {
	if o == nil || o.Dfw1 == nil {
		return nil, false
	}
	return o.Dfw1, true
}

// HasDfw1 returns a boolean if a field has been set.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) HasDfw1() bool {
	if o != nil && o.Dfw1 != nil {
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
	if o == nil || o.Ewr1 == nil {
		var ret FindSpotMarketPrices200ResponseSpotMarketPricesAms1
		return ret
	}
	return *o.Ewr1
}

// GetEwr1Ok returns a tuple with the Ewr1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) GetEwr1Ok() (*FindSpotMarketPrices200ResponseSpotMarketPricesAms1, bool) {
	if o == nil || o.Ewr1 == nil {
		return nil, false
	}
	return o.Ewr1, true
}

// HasEwr1 returns a boolean if a field has been set.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) HasEwr1() bool {
	if o != nil && o.Ewr1 != nil {
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
	if o == nil || o.Fra1 == nil {
		var ret FindSpotMarketPrices200ResponseSpotMarketPricesAtl1
		return ret
	}
	return *o.Fra1
}

// GetFra1Ok returns a tuple with the Fra1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) GetFra1Ok() (*FindSpotMarketPrices200ResponseSpotMarketPricesAtl1, bool) {
	if o == nil || o.Fra1 == nil {
		return nil, false
	}
	return o.Fra1, true
}

// HasFra1 returns a boolean if a field has been set.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) HasFra1() bool {
	if o != nil && o.Fra1 != nil {
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
	if o == nil || o.Iad1 == nil {
		var ret FindSpotMarketPrices200ResponseSpotMarketPricesAtl1
		return ret
	}
	return *o.Iad1
}

// GetIad1Ok returns a tuple with the Iad1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) GetIad1Ok() (*FindSpotMarketPrices200ResponseSpotMarketPricesAtl1, bool) {
	if o == nil || o.Iad1 == nil {
		return nil, false
	}
	return o.Iad1, true
}

// HasIad1 returns a boolean if a field has been set.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) HasIad1() bool {
	if o != nil && o.Iad1 != nil {
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
	if o == nil || o.Lax1 == nil {
		var ret FindSpotMarketPrices200ResponseSpotMarketPricesAtl1
		return ret
	}
	return *o.Lax1
}

// GetLax1Ok returns a tuple with the Lax1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) GetLax1Ok() (*FindSpotMarketPrices200ResponseSpotMarketPricesAtl1, bool) {
	if o == nil || o.Lax1 == nil {
		return nil, false
	}
	return o.Lax1, true
}

// HasLax1 returns a boolean if a field has been set.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) HasLax1() bool {
	if o != nil && o.Lax1 != nil {
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
	if o == nil || o.Nrt1 == nil {
		var ret FindSpotMarketPrices200ResponseSpotMarketPricesAms1
		return ret
	}
	return *o.Nrt1
}

// GetNrt1Ok returns a tuple with the Nrt1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) GetNrt1Ok() (*FindSpotMarketPrices200ResponseSpotMarketPricesAms1, bool) {
	if o == nil || o.Nrt1 == nil {
		return nil, false
	}
	return o.Nrt1, true
}

// HasNrt1 returns a boolean if a field has been set.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) HasNrt1() bool {
	if o != nil && o.Nrt1 != nil {
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
	if o == nil || o.Ord1 == nil {
		var ret FindSpotMarketPrices200ResponseSpotMarketPricesAtl1
		return ret
	}
	return *o.Ord1
}

// GetOrd1Ok returns a tuple with the Ord1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) GetOrd1Ok() (*FindSpotMarketPrices200ResponseSpotMarketPricesAtl1, bool) {
	if o == nil || o.Ord1 == nil {
		return nil, false
	}
	return o.Ord1, true
}

// HasOrd1 returns a boolean if a field has been set.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) HasOrd1() bool {
	if o != nil && o.Ord1 != nil {
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
	if o == nil || o.Sea1 == nil {
		var ret FindSpotMarketPrices200ResponseSpotMarketPricesAtl1
		return ret
	}
	return *o.Sea1
}

// GetSea1Ok returns a tuple with the Sea1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) GetSea1Ok() (*FindSpotMarketPrices200ResponseSpotMarketPricesAtl1, bool) {
	if o == nil || o.Sea1 == nil {
		return nil, false
	}
	return o.Sea1, true
}

// HasSea1 returns a boolean if a field has been set.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) HasSea1() bool {
	if o != nil && o.Sea1 != nil {
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
	if o == nil || o.Sin1 == nil {
		var ret FindSpotMarketPrices200ResponseSpotMarketPricesAtl1
		return ret
	}
	return *o.Sin1
}

// GetSin1Ok returns a tuple with the Sin1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) GetSin1Ok() (*FindSpotMarketPrices200ResponseSpotMarketPricesAtl1, bool) {
	if o == nil || o.Sin1 == nil {
		return nil, false
	}
	return o.Sin1, true
}

// HasSin1 returns a boolean if a field has been set.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) HasSin1() bool {
	if o != nil && o.Sin1 != nil {
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
	if o == nil || o.Sjc1 == nil {
		var ret FindSpotMarketPrices200ResponseSpotMarketPricesAms1
		return ret
	}
	return *o.Sjc1
}

// GetSjc1Ok returns a tuple with the Sjc1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) GetSjc1Ok() (*FindSpotMarketPrices200ResponseSpotMarketPricesAms1, bool) {
	if o == nil || o.Sjc1 == nil {
		return nil, false
	}
	return o.Sjc1, true
}

// HasSjc1 returns a boolean if a field has been set.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) HasSjc1() bool {
	if o != nil && o.Sjc1 != nil {
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
	if o == nil || o.Syd1 == nil {
		var ret FindSpotMarketPrices200ResponseSpotMarketPricesAtl1
		return ret
	}
	return *o.Syd1
}

// GetSyd1Ok returns a tuple with the Syd1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) GetSyd1Ok() (*FindSpotMarketPrices200ResponseSpotMarketPricesAtl1, bool) {
	if o == nil || o.Syd1 == nil {
		return nil, false
	}
	return o.Syd1, true
}

// HasSyd1 returns a boolean if a field has been set.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) HasSyd1() bool {
	if o != nil && o.Syd1 != nil {
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
	if o == nil || o.Yyz1 == nil {
		var ret FindSpotMarketPrices200ResponseSpotMarketPricesAtl1
		return ret
	}
	return *o.Yyz1
}

// GetYyz1Ok returns a tuple with the Yyz1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) GetYyz1Ok() (*FindSpotMarketPrices200ResponseSpotMarketPricesAtl1, bool) {
	if o == nil || o.Yyz1 == nil {
		return nil, false
	}
	return o.Yyz1, true
}

// HasYyz1 returns a boolean if a field has been set.
func (o *FindSpotMarketPrices200ResponseSpotMarketPrices) HasYyz1() bool {
	if o != nil && o.Yyz1 != nil {
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
	if o.Ams1 != nil {
		toSerialize["ams1"] = o.Ams1
	}
	if o.Atl1 != nil {
		toSerialize["atl1"] = o.Atl1
	}
	if o.Dfw1 != nil {
		toSerialize["dfw1"] = o.Dfw1
	}
	if o.Ewr1 != nil {
		toSerialize["ewr1"] = o.Ewr1
	}
	if o.Fra1 != nil {
		toSerialize["fra1"] = o.Fra1
	}
	if o.Iad1 != nil {
		toSerialize["iad1"] = o.Iad1
	}
	if o.Lax1 != nil {
		toSerialize["lax1"] = o.Lax1
	}
	if o.Nrt1 != nil {
		toSerialize["nrt1"] = o.Nrt1
	}
	if o.Ord1 != nil {
		toSerialize["ord1"] = o.Ord1
	}
	if o.Sea1 != nil {
		toSerialize["sea1"] = o.Sea1
	}
	if o.Sin1 != nil {
		toSerialize["sin1"] = o.Sin1
	}
	if o.Sjc1 != nil {
		toSerialize["sjc1"] = o.Sjc1
	}
	if o.Syd1 != nil {
		toSerialize["syd1"] = o.Syd1
	}
	if o.Yyz1 != nil {
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