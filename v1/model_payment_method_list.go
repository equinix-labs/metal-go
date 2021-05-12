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

// PaymentMethodList struct for PaymentMethodList
type PaymentMethodList struct {
	PaymentMethods *[]PaymentMethod `json:"payment_methods,omitempty"`
}

// NewPaymentMethodList instantiates a new PaymentMethodList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMethodList() *PaymentMethodList {
	this := PaymentMethodList{}
	return &this
}

// NewPaymentMethodListWithDefaults instantiates a new PaymentMethodList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMethodListWithDefaults() *PaymentMethodList {
	this := PaymentMethodList{}
	return &this
}

// GetPaymentMethods returns the PaymentMethods field value if set, zero value otherwise.
func (o *PaymentMethodList) GetPaymentMethods() []PaymentMethod {
	if o == nil || o.PaymentMethods == nil {
		var ret []PaymentMethod
		return ret
	}
	return *o.PaymentMethods
}

// GetPaymentMethodsOk returns a tuple with the PaymentMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodList) GetPaymentMethodsOk() (*[]PaymentMethod, bool) {
	if o == nil || o.PaymentMethods == nil {
		return nil, false
	}
	return o.PaymentMethods, true
}

// HasPaymentMethods returns a boolean if a field has been set.
func (o *PaymentMethodList) HasPaymentMethods() bool {
	if o != nil && o.PaymentMethods != nil {
		return true
	}

	return false
}

// SetPaymentMethods gets a reference to the given []PaymentMethod and assigns it to the PaymentMethods field.
func (o *PaymentMethodList) SetPaymentMethods(v []PaymentMethod) {
	o.PaymentMethods = &v
}

func (o PaymentMethodList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PaymentMethods != nil {
		toSerialize["payment_methods"] = o.PaymentMethods
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentMethodList struct {
	value *PaymentMethodList
	isSet bool
}

func (v NullablePaymentMethodList) Get() *PaymentMethodList {
	return v.value
}

func (v *NullablePaymentMethodList) Set(val *PaymentMethodList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethodList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethodList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethodList(val *PaymentMethodList) *NullablePaymentMethodList {
	return &NullablePaymentMethodList{value: val, isSet: true}
}

func (v NullablePaymentMethodList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethodList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


