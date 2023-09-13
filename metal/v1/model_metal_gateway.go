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
	"fmt"
	"time"
)

// checks if the MetalGateway type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetalGateway{}

// MetalGatewayState The current state of the Metal Gateway. 'Ready' indicates the gateway record has been configured, but is currently not active on the network. 'Active' indicates the gateway has been configured on the network. 'Deleting' is a temporary state used to indicate that the gateway is in the process of being un-configured from the network, after which the gateway record will be deleted.
type MetalGatewayState string

// List of MetalGatewayState
const (
	METALGATEWAY_READY    MetalGatewayState = "ready"
	METALGATEWAY_ACTIVE   MetalGatewayState = "active"
	METALGATEWAY_DELETING MetalGatewayState = "deleting"
)

// All allowed values of MetalGatewayState enum
var AllowedMetalGatewayStateEnumValues = []MetalGatewayState{
	"ready",
	"active",
	"deleting",
}

func (v *MetalGatewayState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MetalGatewayState(value)
	for _, existing := range AllowedMetalGatewayStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MetalGatewayState", value)
}

// NewMetalGatewayStateFromValue returns a pointer to a valid MetalGatewayState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMetalGatewayStateFromValue(v string) (*MetalGatewayState, error) {
	ev := MetalGatewayState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MetalGatewayState: valid values are %v", v, AllowedMetalGatewayStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MetalGatewayState) IsValid() bool {
	for _, existing := range AllowedMetalGatewayStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to State value
func (v MetalGatewayState) Ptr() *MetalGatewayState {
	return &v
}

type NullableMetalGatewayState struct {
	value *MetalGatewayState
	isSet bool
}

func (v NullableMetalGatewayState) Get() *MetalGatewayState {
	return v.value
}

func (v *NullableMetalGatewayState) Set(val *MetalGatewayState) {
	v.value = val
	v.isSet = true
}

func (v NullableMetalGatewayState) IsSet() bool {
	return v.isSet
}

func (v *NullableMetalGatewayState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetalGatewayState(val *MetalGatewayState) *NullableMetalGatewayState {
	return &NullableMetalGatewayState{value: val, isSet: true}
}

func (v NullableMetalGatewayState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetalGatewayState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

// MetalGateway struct for MetalGateway
type MetalGateway struct {
	CreatedAt     *time.Time     `json:"created_at,omitempty"`
	CreatedBy     *Href          `json:"created_by,omitempty"`
	Href          *string        `json:"href,omitempty"`
	Id            *string        `json:"id,omitempty"`
	IpReservation *IPReservation `json:"ip_reservation,omitempty"`
	Project       *Project       `json:"project,omitempty"`
	// The current state of the Metal Gateway. 'Ready' indicates the gateway record has been configured, but is currently not active on the network. 'Active' indicates the gateway has been configured on the network. 'Deleting' is a temporary state used to indicate that the gateway is in the process of being un-configured from the network, after which the gateway record will be deleted.
	State                *MetalGatewayState `json:"state,omitempty"`
	UpdatedAt            *time.Time         `json:"updated_at,omitempty"`
	VirtualNetwork       *VirtualNetwork    `json:"virtual_network,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MetalGateway MetalGateway

// NewMetalGateway instantiates a new MetalGateway object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetalGateway() *MetalGateway {
	this := MetalGateway{}
	return &this
}

// NewMetalGatewayWithDefaults instantiates a new MetalGateway object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetalGatewayWithDefaults() *MetalGateway {
	this := MetalGateway{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *MetalGateway) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetalGateway) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *MetalGateway) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *MetalGateway) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *MetalGateway) GetCreatedBy() Href {
	if o == nil || IsNil(o.CreatedBy) {
		var ret Href
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetalGateway) GetCreatedByOk() (*Href, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *MetalGateway) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given Href and assigns it to the CreatedBy field.
func (o *MetalGateway) SetCreatedBy(v Href) {
	o.CreatedBy = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *MetalGateway) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetalGateway) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *MetalGateway) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *MetalGateway) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MetalGateway) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetalGateway) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MetalGateway) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MetalGateway) SetId(v string) {
	o.Id = &v
}

// GetIpReservation returns the IpReservation field value if set, zero value otherwise.
func (o *MetalGateway) GetIpReservation() IPReservation {
	if o == nil || IsNil(o.IpReservation) {
		var ret IPReservation
		return ret
	}
	return *o.IpReservation
}

// GetIpReservationOk returns a tuple with the IpReservation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetalGateway) GetIpReservationOk() (*IPReservation, bool) {
	if o == nil || IsNil(o.IpReservation) {
		return nil, false
	}
	return o.IpReservation, true
}

// HasIpReservation returns a boolean if a field has been set.
func (o *MetalGateway) HasIpReservation() bool {
	if o != nil && !IsNil(o.IpReservation) {
		return true
	}

	return false
}

// SetIpReservation gets a reference to the given IPReservation and assigns it to the IpReservation field.
func (o *MetalGateway) SetIpReservation(v IPReservation) {
	o.IpReservation = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *MetalGateway) GetProject() Project {
	if o == nil || IsNil(o.Project) {
		var ret Project
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetalGateway) GetProjectOk() (*Project, bool) {
	if o == nil || IsNil(o.Project) {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *MetalGateway) HasProject() bool {
	if o != nil && !IsNil(o.Project) {
		return true
	}

	return false
}

// SetProject gets a reference to the given Project and assigns it to the Project field.
func (o *MetalGateway) SetProject(v Project) {
	o.Project = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *MetalGateway) GetState() MetalGatewayState {
	if o == nil || IsNil(o.State) {
		var ret MetalGatewayState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetalGateway) GetStateOk() (*MetalGatewayState, bool) {
	if o == nil || IsNil(o.State) {
		return o.State, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *MetalGateway) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *MetalGateway) SetState(v MetalGatewayState) {
	o.State = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *MetalGateway) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetalGateway) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *MetalGateway) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *MetalGateway) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetVirtualNetwork returns the VirtualNetwork field value if set, zero value otherwise.
func (o *MetalGateway) GetVirtualNetwork() VirtualNetwork {
	if o == nil || IsNil(o.VirtualNetwork) {
		var ret VirtualNetwork
		return ret
	}
	return *o.VirtualNetwork
}

// GetVirtualNetworkOk returns a tuple with the VirtualNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetalGateway) GetVirtualNetworkOk() (*VirtualNetwork, bool) {
	if o == nil || IsNil(o.VirtualNetwork) {
		return nil, false
	}
	return o.VirtualNetwork, true
}

// HasVirtualNetwork returns a boolean if a field has been set.
func (o *MetalGateway) HasVirtualNetwork() bool {
	if o != nil && !IsNil(o.VirtualNetwork) {
		return true
	}

	return false
}

// SetVirtualNetwork gets a reference to the given VirtualNetwork and assigns it to the VirtualNetwork field.
func (o *MetalGateway) SetVirtualNetwork(v VirtualNetwork) {
	o.VirtualNetwork = &v
}

func (o MetalGateway) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetalGateway) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.CreatedBy) {
		toSerialize["created_by"] = o.CreatedBy
	}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.IpReservation) {
		toSerialize["ip_reservation"] = o.IpReservation
	}
	if !IsNil(o.Project) {
		toSerialize["project"] = o.Project
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.VirtualNetwork) {
		toSerialize["virtual_network"] = o.VirtualNetwork
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MetalGateway) UnmarshalJSON(bytes []byte) (err error) {
	varMetalGateway := _MetalGateway{}

	err = json.Unmarshal(bytes, &varMetalGateway)

	if err != nil {
		return err
	}

	*o = MetalGateway(varMetalGateway)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "created_by")
		delete(additionalProperties, "href")
		delete(additionalProperties, "id")
		delete(additionalProperties, "ip_reservation")
		delete(additionalProperties, "project")
		delete(additionalProperties, "state")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "virtual_network")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMetalGateway struct {
	value *MetalGateway
	isSet bool
}

func (v NullableMetalGateway) Get() *MetalGateway {
	return v.value
}

func (v *NullableMetalGateway) Set(val *MetalGateway) {
	v.value = val
	v.isSet = true
}

func (v NullableMetalGateway) IsSet() bool {
	return v.isSet
}

func (v *NullableMetalGateway) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetalGateway(val *MetalGateway) *NullableMetalGateway {
	return &NullableMetalGateway{value: val, isSet: true}
}

func (v NullableMetalGateway) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetalGateway) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
