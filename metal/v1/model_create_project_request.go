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

// CreateProjectRequest struct for CreateProjectRequest
type CreateProjectRequest struct {
	Customdata      map[string]interface{} `json:"customdata,omitempty"`
	Name            string                 `json:"name"`
	OrganizationId  *string                `json:"organization_id,omitempty"`
	PaymentMethodId *string                `json:"payment_method_id,omitempty"`
}

// NewCreateProjectRequest instantiates a new CreateProjectRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateProjectRequest(name string) *CreateProjectRequest {
	this := CreateProjectRequest{}
	this.Name = name
	return &this
}

// NewCreateProjectRequestWithDefaults instantiates a new CreateProjectRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateProjectRequestWithDefaults() *CreateProjectRequest {
	this := CreateProjectRequest{}
	return &this
}

// GetCustomdata returns the Customdata field value if set, zero value otherwise.
func (o *CreateProjectRequest) GetCustomdata() map[string]interface{} {
	if o == nil || o.Customdata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Customdata
}

// GetCustomdataOk returns a tuple with the Customdata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProjectRequest) GetCustomdataOk() (map[string]interface{}, bool) {
	if o == nil || o.Customdata == nil {
		return nil, false
	}
	return o.Customdata, true
}

// HasCustomdata returns a boolean if a field has been set.
func (o *CreateProjectRequest) HasCustomdata() bool {
	if o != nil && o.Customdata != nil {
		return true
	}

	return false
}

// SetCustomdata gets a reference to the given map[string]interface{} and assigns it to the Customdata field.
func (o *CreateProjectRequest) SetCustomdata(v map[string]interface{}) {
	o.Customdata = v
}

// GetName returns the Name field value
func (o *CreateProjectRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateProjectRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateProjectRequest) SetName(v string) {
	o.Name = v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *CreateProjectRequest) GetOrganizationId() string {
	if o == nil || o.OrganizationId == nil {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProjectRequest) GetOrganizationIdOk() (*string, bool) {
	if o == nil || o.OrganizationId == nil {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *CreateProjectRequest) HasOrganizationId() bool {
	if o != nil && o.OrganizationId != nil {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *CreateProjectRequest) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetPaymentMethodId returns the PaymentMethodId field value if set, zero value otherwise.
func (o *CreateProjectRequest) GetPaymentMethodId() string {
	if o == nil || o.PaymentMethodId == nil {
		var ret string
		return ret
	}
	return *o.PaymentMethodId
}

// GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProjectRequest) GetPaymentMethodIdOk() (*string, bool) {
	if o == nil || o.PaymentMethodId == nil {
		return nil, false
	}
	return o.PaymentMethodId, true
}

// HasPaymentMethodId returns a boolean if a field has been set.
func (o *CreateProjectRequest) HasPaymentMethodId() bool {
	if o != nil && o.PaymentMethodId != nil {
		return true
	}

	return false
}

// SetPaymentMethodId gets a reference to the given string and assigns it to the PaymentMethodId field.
func (o *CreateProjectRequest) SetPaymentMethodId(v string) {
	o.PaymentMethodId = &v
}

func (o CreateProjectRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Customdata != nil {
		toSerialize["customdata"] = o.Customdata
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.OrganizationId != nil {
		toSerialize["organization_id"] = o.OrganizationId
	}
	if o.PaymentMethodId != nil {
		toSerialize["payment_method_id"] = o.PaymentMethodId
	}
	return json.Marshal(toSerialize)
}

type NullableCreateProjectRequest struct {
	value *CreateProjectRequest
	isSet bool
}

func (v NullableCreateProjectRequest) Get() *CreateProjectRequest {
	return v.value
}

func (v *NullableCreateProjectRequest) Set(val *CreateProjectRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateProjectRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateProjectRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateProjectRequest(val *CreateProjectRequest) *NullableCreateProjectRequest {
	return &NullableCreateProjectRequest{value: val, isSet: true}
}

func (v NullableCreateProjectRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateProjectRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}