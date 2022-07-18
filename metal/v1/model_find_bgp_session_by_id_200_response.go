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
	"time"
)

// FindBgpSessionById200Response struct for FindBgpSessionById200Response
type FindBgpSessionById200Response struct {
	AddressFamily string                                `json:"address_family"`
	CreatedAt     *time.Time                            `json:"created_at,omitempty"`
	DefaultRoute  *bool                                 `json:"default_route,omitempty"`
	Device        *FindBatchById200ResponseDevicesInner `json:"device,omitempty"`
	Href          *string                               `json:"href,omitempty"`
	Id            *string                               `json:"id,omitempty"`
	LearnedRoutes []string                              `json:"learned_routes,omitempty"`
	//  The status of the BGP Session. Multiple status values may be reported when the device is connected to multiple switches, one value per switch. Each status will start with \"unknown\" and progress to \"up\" or \"down\" depending on the connected device. Subsequent \"unknown\" values indicate a problem acquiring status from the switch.
	Status    *string    `json:"status,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewFindBgpSessionById200Response instantiates a new FindBgpSessionById200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFindBgpSessionById200Response(addressFamily string) *FindBgpSessionById200Response {
	this := FindBgpSessionById200Response{}
	this.AddressFamily = addressFamily
	return &this
}

// NewFindBgpSessionById200ResponseWithDefaults instantiates a new FindBgpSessionById200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFindBgpSessionById200ResponseWithDefaults() *FindBgpSessionById200Response {
	this := FindBgpSessionById200Response{}
	return &this
}

// GetAddressFamily returns the AddressFamily field value
func (o *FindBgpSessionById200Response) GetAddressFamily() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AddressFamily
}

// GetAddressFamilyOk returns a tuple with the AddressFamily field value
// and a boolean to check if the value has been set.
func (o *FindBgpSessionById200Response) GetAddressFamilyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AddressFamily, true
}

// SetAddressFamily sets field value
func (o *FindBgpSessionById200Response) SetAddressFamily(v string) {
	o.AddressFamily = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *FindBgpSessionById200Response) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindBgpSessionById200Response) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *FindBgpSessionById200Response) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *FindBgpSessionById200Response) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDefaultRoute returns the DefaultRoute field value if set, zero value otherwise.
func (o *FindBgpSessionById200Response) GetDefaultRoute() bool {
	if o == nil || o.DefaultRoute == nil {
		var ret bool
		return ret
	}
	return *o.DefaultRoute
}

// GetDefaultRouteOk returns a tuple with the DefaultRoute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindBgpSessionById200Response) GetDefaultRouteOk() (*bool, bool) {
	if o == nil || o.DefaultRoute == nil {
		return nil, false
	}
	return o.DefaultRoute, true
}

// HasDefaultRoute returns a boolean if a field has been set.
func (o *FindBgpSessionById200Response) HasDefaultRoute() bool {
	if o != nil && o.DefaultRoute != nil {
		return true
	}

	return false
}

// SetDefaultRoute gets a reference to the given bool and assigns it to the DefaultRoute field.
func (o *FindBgpSessionById200Response) SetDefaultRoute(v bool) {
	o.DefaultRoute = &v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *FindBgpSessionById200Response) GetDevice() FindBatchById200ResponseDevicesInner {
	if o == nil || o.Device == nil {
		var ret FindBatchById200ResponseDevicesInner
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindBgpSessionById200Response) GetDeviceOk() (*FindBatchById200ResponseDevicesInner, bool) {
	if o == nil || o.Device == nil {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *FindBgpSessionById200Response) HasDevice() bool {
	if o != nil && o.Device != nil {
		return true
	}

	return false
}

// SetDevice gets a reference to the given FindBatchById200ResponseDevicesInner and assigns it to the Device field.
func (o *FindBgpSessionById200Response) SetDevice(v FindBatchById200ResponseDevicesInner) {
	o.Device = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *FindBgpSessionById200Response) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindBgpSessionById200Response) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *FindBgpSessionById200Response) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *FindBgpSessionById200Response) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FindBgpSessionById200Response) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindBgpSessionById200Response) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FindBgpSessionById200Response) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *FindBgpSessionById200Response) SetId(v string) {
	o.Id = &v
}

// GetLearnedRoutes returns the LearnedRoutes field value if set, zero value otherwise.
func (o *FindBgpSessionById200Response) GetLearnedRoutes() []string {
	if o == nil || o.LearnedRoutes == nil {
		var ret []string
		return ret
	}
	return o.LearnedRoutes
}

// GetLearnedRoutesOk returns a tuple with the LearnedRoutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindBgpSessionById200Response) GetLearnedRoutesOk() ([]string, bool) {
	if o == nil || o.LearnedRoutes == nil {
		return nil, false
	}
	return o.LearnedRoutes, true
}

// HasLearnedRoutes returns a boolean if a field has been set.
func (o *FindBgpSessionById200Response) HasLearnedRoutes() bool {
	if o != nil && o.LearnedRoutes != nil {
		return true
	}

	return false
}

// SetLearnedRoutes gets a reference to the given []string and assigns it to the LearnedRoutes field.
func (o *FindBgpSessionById200Response) SetLearnedRoutes(v []string) {
	o.LearnedRoutes = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *FindBgpSessionById200Response) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindBgpSessionById200Response) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *FindBgpSessionById200Response) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *FindBgpSessionById200Response) SetStatus(v string) {
	o.Status = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *FindBgpSessionById200Response) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindBgpSessionById200Response) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *FindBgpSessionById200Response) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *FindBgpSessionById200Response) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o FindBgpSessionById200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["address_family"] = o.AddressFamily
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.DefaultRoute != nil {
		toSerialize["default_route"] = o.DefaultRoute
	}
	if o.Device != nil {
		toSerialize["device"] = o.Device
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LearnedRoutes != nil {
		toSerialize["learned_routes"] = o.LearnedRoutes
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableFindBgpSessionById200Response struct {
	value *FindBgpSessionById200Response
	isSet bool
}

func (v NullableFindBgpSessionById200Response) Get() *FindBgpSessionById200Response {
	return v.value
}

func (v *NullableFindBgpSessionById200Response) Set(val *FindBgpSessionById200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableFindBgpSessionById200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableFindBgpSessionById200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFindBgpSessionById200Response(val *FindBgpSessionById200Response) *NullableFindBgpSessionById200Response {
	return &NullableFindBgpSessionById200Response{value: val, isSet: true}
}

func (v NullableFindBgpSessionById200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFindBgpSessionById200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}