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

// PaymentMethodUpdateInput struct for PaymentMethodUpdateInput
type PaymentMethodUpdateInput struct {
	Name *string `json:"name,omitempty"`
	Default *bool `json:"default,omitempty"`
	CardholderName *string `json:"cardholder_name,omitempty"`
	ExpirationMonth *string `json:"expiration_month,omitempty"`
	ExpirationYear *int32 `json:"expiration_year,omitempty"`
	BillingAddress map[string]interface{} `json:"billing_address,omitempty"`
}

// NewPaymentMethodUpdateInput instantiates a new PaymentMethodUpdateInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMethodUpdateInput() *PaymentMethodUpdateInput {
	this := PaymentMethodUpdateInput{}
	return &this
}

// NewPaymentMethodUpdateInputWithDefaults instantiates a new PaymentMethodUpdateInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMethodUpdateInputWithDefaults() *PaymentMethodUpdateInput {
	this := PaymentMethodUpdateInput{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PaymentMethodUpdateInput) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodUpdateInput) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PaymentMethodUpdateInput) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PaymentMethodUpdateInput) SetName(v string) {
	o.Name = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *PaymentMethodUpdateInput) GetDefault() bool {
	if o == nil || o.Default == nil {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodUpdateInput) GetDefaultOk() (*bool, bool) {
	if o == nil || o.Default == nil {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *PaymentMethodUpdateInput) HasDefault() bool {
	if o != nil && o.Default != nil {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *PaymentMethodUpdateInput) SetDefault(v bool) {
	o.Default = &v
}

// GetCardholderName returns the CardholderName field value if set, zero value otherwise.
func (o *PaymentMethodUpdateInput) GetCardholderName() string {
	if o == nil || o.CardholderName == nil {
		var ret string
		return ret
	}
	return *o.CardholderName
}

// GetCardholderNameOk returns a tuple with the CardholderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodUpdateInput) GetCardholderNameOk() (*string, bool) {
	if o == nil || o.CardholderName == nil {
		return nil, false
	}
	return o.CardholderName, true
}

// HasCardholderName returns a boolean if a field has been set.
func (o *PaymentMethodUpdateInput) HasCardholderName() bool {
	if o != nil && o.CardholderName != nil {
		return true
	}

	return false
}

// SetCardholderName gets a reference to the given string and assigns it to the CardholderName field.
func (o *PaymentMethodUpdateInput) SetCardholderName(v string) {
	o.CardholderName = &v
}

// GetExpirationMonth returns the ExpirationMonth field value if set, zero value otherwise.
func (o *PaymentMethodUpdateInput) GetExpirationMonth() string {
	if o == nil || o.ExpirationMonth == nil {
		var ret string
		return ret
	}
	return *o.ExpirationMonth
}

// GetExpirationMonthOk returns a tuple with the ExpirationMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodUpdateInput) GetExpirationMonthOk() (*string, bool) {
	if o == nil || o.ExpirationMonth == nil {
		return nil, false
	}
	return o.ExpirationMonth, true
}

// HasExpirationMonth returns a boolean if a field has been set.
func (o *PaymentMethodUpdateInput) HasExpirationMonth() bool {
	if o != nil && o.ExpirationMonth != nil {
		return true
	}

	return false
}

// SetExpirationMonth gets a reference to the given string and assigns it to the ExpirationMonth field.
func (o *PaymentMethodUpdateInput) SetExpirationMonth(v string) {
	o.ExpirationMonth = &v
}

// GetExpirationYear returns the ExpirationYear field value if set, zero value otherwise.
func (o *PaymentMethodUpdateInput) GetExpirationYear() int32 {
	if o == nil || o.ExpirationYear == nil {
		var ret int32
		return ret
	}
	return *o.ExpirationYear
}

// GetExpirationYearOk returns a tuple with the ExpirationYear field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodUpdateInput) GetExpirationYearOk() (*int32, bool) {
	if o == nil || o.ExpirationYear == nil {
		return nil, false
	}
	return o.ExpirationYear, true
}

// HasExpirationYear returns a boolean if a field has been set.
func (o *PaymentMethodUpdateInput) HasExpirationYear() bool {
	if o != nil && o.ExpirationYear != nil {
		return true
	}

	return false
}

// SetExpirationYear gets a reference to the given int32 and assigns it to the ExpirationYear field.
func (o *PaymentMethodUpdateInput) SetExpirationYear(v int32) {
	o.ExpirationYear = &v
}

// GetBillingAddress returns the BillingAddress field value if set, zero value otherwise.
func (o *PaymentMethodUpdateInput) GetBillingAddress() map[string]interface{} {
	if o == nil || o.BillingAddress == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.BillingAddress
}

// GetBillingAddressOk returns a tuple with the BillingAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodUpdateInput) GetBillingAddressOk() (map[string]interface{}, bool) {
	if o == nil || o.BillingAddress == nil {
		return nil, false
	}
	return o.BillingAddress, true
}

// HasBillingAddress returns a boolean if a field has been set.
func (o *PaymentMethodUpdateInput) HasBillingAddress() bool {
	if o != nil && o.BillingAddress != nil {
		return true
	}

	return false
}

// SetBillingAddress gets a reference to the given map[string]interface{} and assigns it to the BillingAddress field.
func (o *PaymentMethodUpdateInput) SetBillingAddress(v map[string]interface{}) {
	o.BillingAddress = v
}

func (o PaymentMethodUpdateInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Default != nil {
		toSerialize["default"] = o.Default
	}
	if o.CardholderName != nil {
		toSerialize["cardholder_name"] = o.CardholderName
	}
	if o.ExpirationMonth != nil {
		toSerialize["expiration_month"] = o.ExpirationMonth
	}
	if o.ExpirationYear != nil {
		toSerialize["expiration_year"] = o.ExpirationYear
	}
	if o.BillingAddress != nil {
		toSerialize["billing_address"] = o.BillingAddress
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentMethodUpdateInput struct {
	value *PaymentMethodUpdateInput
	isSet bool
}

func (v NullablePaymentMethodUpdateInput) Get() *PaymentMethodUpdateInput {
	return v.value
}

func (v *NullablePaymentMethodUpdateInput) Set(val *PaymentMethodUpdateInput) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethodUpdateInput) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethodUpdateInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethodUpdateInput(val *PaymentMethodUpdateInput) *NullablePaymentMethodUpdateInput {
	return &NullablePaymentMethodUpdateInput{value: val, isSet: true}
}

func (v NullablePaymentMethodUpdateInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethodUpdateInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


