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

// BgpConfig struct for BgpConfig
type BgpConfig struct {
	// Autonomous System Number. ASN is required with Global BGP. With Local BGP the private ASN, 65000, is assigned.
	Asn       *int32     `json:"asn,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// In a Local BGP deployment, a customer uses an internal ASN to control routes within a single Equinix Metal datacenter. This means that the routes are never advertised to the global Internet. Global BGP, on the other hand, requires a customer to have a registered ASN and IP space.
	DeploymentType *string `json:"deployment_type,omitempty"`
	Href           *string `json:"href,omitempty"`
	Id             *string `json:"id,omitempty"`
	// The maximum number of route filters allowed per server
	MaxPrefix *int32 `json:"max_prefix,omitempty"`
	// (Optional) Password for BGP session in plaintext (not a checksum)
	Md5     NullableString                        `json:"md5,omitempty"`
	Project *FindBatchById200ResponseDevicesInner `json:"project,omitempty"`
	// The IP block ranges associated to the ASN (Populated in Global BGP only)
	Ranges      []FindBgpConfigByProject200ResponseRangesInner `json:"ranges,omitempty"`
	RequestedAt *time.Time                                     `json:"requested_at,omitempty"`
	// Specifies AS-MACRO (aka AS-SET) to use when building client route filters
	RouteObject *string `json:"route_object,omitempty"`
	// The direct connections between neighboring routers that want to exchange routing information.
	Sessions []FindBgpSessionById200Response `json:"sessions,omitempty"`
	// Status of the BGP Config. Status \"requested\" is valid only with the \"global\" deployment_type.
	Status *string `json:"status,omitempty"`
}

// NewBgpConfig instantiates a new BgpConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBgpConfig() *BgpConfig {
	this := BgpConfig{}
	return &this
}

// NewBgpConfigWithDefaults instantiates a new BgpConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBgpConfigWithDefaults() *BgpConfig {
	this := BgpConfig{}
	return &this
}

// GetAsn returns the Asn field value if set, zero value otherwise.
func (o *BgpConfig) GetAsn() int32 {
	if o == nil || o.Asn == nil {
		var ret int32
		return ret
	}
	return *o.Asn
}

// GetAsnOk returns a tuple with the Asn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpConfig) GetAsnOk() (*int32, bool) {
	if o == nil || o.Asn == nil {
		return nil, false
	}
	return o.Asn, true
}

// HasAsn returns a boolean if a field has been set.
func (o *BgpConfig) HasAsn() bool {
	if o != nil && o.Asn != nil {
		return true
	}

	return false
}

// SetAsn gets a reference to the given int32 and assigns it to the Asn field.
func (o *BgpConfig) SetAsn(v int32) {
	o.Asn = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *BgpConfig) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpConfig) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *BgpConfig) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *BgpConfig) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDeploymentType returns the DeploymentType field value if set, zero value otherwise.
func (o *BgpConfig) GetDeploymentType() string {
	if o == nil || o.DeploymentType == nil {
		var ret string
		return ret
	}
	return *o.DeploymentType
}

// GetDeploymentTypeOk returns a tuple with the DeploymentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpConfig) GetDeploymentTypeOk() (*string, bool) {
	if o == nil || o.DeploymentType == nil {
		return nil, false
	}
	return o.DeploymentType, true
}

// HasDeploymentType returns a boolean if a field has been set.
func (o *BgpConfig) HasDeploymentType() bool {
	if o != nil && o.DeploymentType != nil {
		return true
	}

	return false
}

// SetDeploymentType gets a reference to the given string and assigns it to the DeploymentType field.
func (o *BgpConfig) SetDeploymentType(v string) {
	o.DeploymentType = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *BgpConfig) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpConfig) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *BgpConfig) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *BgpConfig) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BgpConfig) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpConfig) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BgpConfig) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BgpConfig) SetId(v string) {
	o.Id = &v
}

// GetMaxPrefix returns the MaxPrefix field value if set, zero value otherwise.
func (o *BgpConfig) GetMaxPrefix() int32 {
	if o == nil || o.MaxPrefix == nil {
		var ret int32
		return ret
	}
	return *o.MaxPrefix
}

// GetMaxPrefixOk returns a tuple with the MaxPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpConfig) GetMaxPrefixOk() (*int32, bool) {
	if o == nil || o.MaxPrefix == nil {
		return nil, false
	}
	return o.MaxPrefix, true
}

// HasMaxPrefix returns a boolean if a field has been set.
func (o *BgpConfig) HasMaxPrefix() bool {
	if o != nil && o.MaxPrefix != nil {
		return true
	}

	return false
}

// SetMaxPrefix gets a reference to the given int32 and assigns it to the MaxPrefix field.
func (o *BgpConfig) SetMaxPrefix(v int32) {
	o.MaxPrefix = &v
}

// GetMd5 returns the Md5 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BgpConfig) GetMd5() string {
	if o == nil || o.Md5.Get() == nil {
		var ret string
		return ret
	}
	return *o.Md5.Get()
}

// GetMd5Ok returns a tuple with the Md5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BgpConfig) GetMd5Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Md5.Get(), o.Md5.IsSet()
}

// HasMd5 returns a boolean if a field has been set.
func (o *BgpConfig) HasMd5() bool {
	if o != nil && o.Md5.IsSet() {
		return true
	}

	return false
}

// SetMd5 gets a reference to the given NullableString and assigns it to the Md5 field.
func (o *BgpConfig) SetMd5(v string) {
	o.Md5.Set(&v)
}

// SetMd5Nil sets the value for Md5 to be an explicit nil
func (o *BgpConfig) SetMd5Nil() {
	o.Md5.Set(nil)
}

// UnsetMd5 ensures that no value is present for Md5, not even an explicit nil
func (o *BgpConfig) UnsetMd5() {
	o.Md5.Unset()
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *BgpConfig) GetProject() FindBatchById200ResponseDevicesInner {
	if o == nil || o.Project == nil {
		var ret FindBatchById200ResponseDevicesInner
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpConfig) GetProjectOk() (*FindBatchById200ResponseDevicesInner, bool) {
	if o == nil || o.Project == nil {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *BgpConfig) HasProject() bool {
	if o != nil && o.Project != nil {
		return true
	}

	return false
}

// SetProject gets a reference to the given FindBatchById200ResponseDevicesInner and assigns it to the Project field.
func (o *BgpConfig) SetProject(v FindBatchById200ResponseDevicesInner) {
	o.Project = &v
}

// GetRanges returns the Ranges field value if set, zero value otherwise.
func (o *BgpConfig) GetRanges() []FindBgpConfigByProject200ResponseRangesInner {
	if o == nil || o.Ranges == nil {
		var ret []FindBgpConfigByProject200ResponseRangesInner
		return ret
	}
	return o.Ranges
}

// GetRangesOk returns a tuple with the Ranges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpConfig) GetRangesOk() ([]FindBgpConfigByProject200ResponseRangesInner, bool) {
	if o == nil || o.Ranges == nil {
		return nil, false
	}
	return o.Ranges, true
}

// HasRanges returns a boolean if a field has been set.
func (o *BgpConfig) HasRanges() bool {
	if o != nil && o.Ranges != nil {
		return true
	}

	return false
}

// SetRanges gets a reference to the given []FindBgpConfigByProject200ResponseRangesInner and assigns it to the Ranges field.
func (o *BgpConfig) SetRanges(v []FindBgpConfigByProject200ResponseRangesInner) {
	o.Ranges = v
}

// GetRequestedAt returns the RequestedAt field value if set, zero value otherwise.
func (o *BgpConfig) GetRequestedAt() time.Time {
	if o == nil || o.RequestedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.RequestedAt
}

// GetRequestedAtOk returns a tuple with the RequestedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpConfig) GetRequestedAtOk() (*time.Time, bool) {
	if o == nil || o.RequestedAt == nil {
		return nil, false
	}
	return o.RequestedAt, true
}

// HasRequestedAt returns a boolean if a field has been set.
func (o *BgpConfig) HasRequestedAt() bool {
	if o != nil && o.RequestedAt != nil {
		return true
	}

	return false
}

// SetRequestedAt gets a reference to the given time.Time and assigns it to the RequestedAt field.
func (o *BgpConfig) SetRequestedAt(v time.Time) {
	o.RequestedAt = &v
}

// GetRouteObject returns the RouteObject field value if set, zero value otherwise.
func (o *BgpConfig) GetRouteObject() string {
	if o == nil || o.RouteObject == nil {
		var ret string
		return ret
	}
	return *o.RouteObject
}

// GetRouteObjectOk returns a tuple with the RouteObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpConfig) GetRouteObjectOk() (*string, bool) {
	if o == nil || o.RouteObject == nil {
		return nil, false
	}
	return o.RouteObject, true
}

// HasRouteObject returns a boolean if a field has been set.
func (o *BgpConfig) HasRouteObject() bool {
	if o != nil && o.RouteObject != nil {
		return true
	}

	return false
}

// SetRouteObject gets a reference to the given string and assigns it to the RouteObject field.
func (o *BgpConfig) SetRouteObject(v string) {
	o.RouteObject = &v
}

// GetSessions returns the Sessions field value if set, zero value otherwise.
func (o *BgpConfig) GetSessions() []FindBgpSessionById200Response {
	if o == nil || o.Sessions == nil {
		var ret []FindBgpSessionById200Response
		return ret
	}
	return o.Sessions
}

// GetSessionsOk returns a tuple with the Sessions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpConfig) GetSessionsOk() ([]FindBgpSessionById200Response, bool) {
	if o == nil || o.Sessions == nil {
		return nil, false
	}
	return o.Sessions, true
}

// HasSessions returns a boolean if a field has been set.
func (o *BgpConfig) HasSessions() bool {
	if o != nil && o.Sessions != nil {
		return true
	}

	return false
}

// SetSessions gets a reference to the given []FindBgpSessionById200Response and assigns it to the Sessions field.
func (o *BgpConfig) SetSessions(v []FindBgpSessionById200Response) {
	o.Sessions = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BgpConfig) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpConfig) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BgpConfig) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *BgpConfig) SetStatus(v string) {
	o.Status = &v
}

func (o BgpConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Asn != nil {
		toSerialize["asn"] = o.Asn
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.DeploymentType != nil {
		toSerialize["deployment_type"] = o.DeploymentType
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.MaxPrefix != nil {
		toSerialize["max_prefix"] = o.MaxPrefix
	}
	if o.Md5.IsSet() {
		toSerialize["md5"] = o.Md5.Get()
	}
	if o.Project != nil {
		toSerialize["project"] = o.Project
	}
	if o.Ranges != nil {
		toSerialize["ranges"] = o.Ranges
	}
	if o.RequestedAt != nil {
		toSerialize["requested_at"] = o.RequestedAt
	}
	if o.RouteObject != nil {
		toSerialize["route_object"] = o.RouteObject
	}
	if o.Sessions != nil {
		toSerialize["sessions"] = o.Sessions
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableBgpConfig struct {
	value *BgpConfig
	isSet bool
}

func (v NullableBgpConfig) Get() *BgpConfig {
	return v.value
}

func (v *NullableBgpConfig) Set(val *BgpConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableBgpConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableBgpConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBgpConfig(val *BgpConfig) *NullableBgpConfig {
	return &NullableBgpConfig{value: val, isSet: true}
}

func (v NullableBgpConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBgpConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}