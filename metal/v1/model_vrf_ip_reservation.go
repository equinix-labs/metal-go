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

// checks if the VrfIpReservation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VrfIpReservation{}

// VrfIpReservationType the model 'VrfIpReservationType'
type VrfIpReservationType string

// List of VrfIpReservationType
const (
	VRFIPRESERVATION_VRF VrfIpReservationType = "vrf"
)

// All allowed values of VrfIpReservationType enum
var AllowedVrfIpReservationTypeEnumValues = []VrfIpReservationType{
	"vrf",
}

func (v *VrfIpReservationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VrfIpReservationType(value)
	for _, existing := range AllowedVrfIpReservationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VrfIpReservationType", value)
}

// NewVrfIpReservationTypeFromValue returns a pointer to a valid VrfIpReservationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVrfIpReservationTypeFromValue(v string) (*VrfIpReservationType, error) {
	ev := VrfIpReservationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VrfIpReservationType: valid values are %v", v, AllowedVrfIpReservationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VrfIpReservationType) IsValid() bool {
	for _, existing := range AllowedVrfIpReservationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Type value
func (v VrfIpReservationType) Ptr() *VrfIpReservationType {
	return &v
}

type NullableVrfIpReservationType struct {
	value *VrfIpReservationType
	isSet bool
}

func (v NullableVrfIpReservationType) Get() *VrfIpReservationType {
	return v.value
}

func (v *NullableVrfIpReservationType) Set(val *VrfIpReservationType) {
	v.value = val
	v.isSet = true
}

func (v NullableVrfIpReservationType) IsSet() bool {
	return v.isSet
}

func (v *NullableVrfIpReservationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVrfIpReservationType(val *VrfIpReservationType) *NullableVrfIpReservationType {
	return &NullableVrfIpReservationType{value: val, isSet: true}
}

func (v NullableVrfIpReservationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVrfIpReservationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

// VrfIpReservation struct for VrfIpReservation
type VrfIpReservation struct {
	AddressFamily        *int32                 `json:"address_family,omitempty"`
	Cidr                 *int32                 `json:"cidr,omitempty"`
	CreatedAt            *time.Time             `json:"created_at,omitempty"`
	CreatedBy            *Href                  `json:"created_by,omitempty"`
	Details              *string                `json:"details,omitempty"`
	Href                 *string                `json:"href,omitempty"`
	Id                   *string                `json:"id,omitempty"`
	MetalGateway         *MetalGatewayLite      `json:"metal_gateway,omitempty"`
	Netmask              *string                `json:"netmask,omitempty"`
	Network              *string                `json:"network,omitempty"`
	Project              *Project               `json:"project,omitempty"`
	State                *string                `json:"state,omitempty"`
	Tags                 []string               `json:"tags,omitempty"`
	Type                 VrfIpReservationType   `json:"type"`
	Vrf                  Vrf                    `json:"vrf"`
	Public               *bool                  `json:"public,omitempty"`
	Management           *bool                  `json:"management,omitempty"`
	Manageable           *bool                  `json:"manageable,omitempty"`
	Customdata           map[string]interface{} `json:"customdata,omitempty"`
	Bill                 *bool                  `json:"bill,omitempty"`
	ProjectLite          *Project               `json:"project_lite,omitempty"`
	Address              *string                `json:"address,omitempty"`
	Gateway              *string                `json:"gateway,omitempty"`
	Metro                *Metro                 `json:"metro,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VrfIpReservation VrfIpReservation

// NewVrfIpReservation instantiates a new VrfIpReservation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVrfIpReservation(type_ VrfIpReservationType, vrf Vrf) *VrfIpReservation {
	this := VrfIpReservation{}
	this.Type = type_
	this.Vrf = vrf
	return &this
}

// NewVrfIpReservationWithDefaults instantiates a new VrfIpReservation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVrfIpReservationWithDefaults() *VrfIpReservation {
	this := VrfIpReservation{}
	return &this
}

// GetAddressFamily returns the AddressFamily field value if set, zero value otherwise.
func (o *VrfIpReservation) GetAddressFamily() int32 {
	if o == nil || IsNil(o.AddressFamily) {
		var ret int32
		return ret
	}
	return *o.AddressFamily
}

// GetAddressFamilyOk returns a tuple with the AddressFamily field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfIpReservation) GetAddressFamilyOk() (*int32, bool) {
	if o == nil || IsNil(o.AddressFamily) {
		return nil, false
	}
	return o.AddressFamily, true
}

// HasAddressFamily returns a boolean if a field has been set.
func (o *VrfIpReservation) HasAddressFamily() bool {
	if o != nil && !IsNil(o.AddressFamily) {
		return true
	}

	return false
}

// SetAddressFamily gets a reference to the given int32 and assigns it to the AddressFamily field.
func (o *VrfIpReservation) SetAddressFamily(v int32) {
	o.AddressFamily = &v
}

// GetCidr returns the Cidr field value if set, zero value otherwise.
func (o *VrfIpReservation) GetCidr() int32 {
	if o == nil || IsNil(o.Cidr) {
		var ret int32
		return ret
	}
	return *o.Cidr
}

// GetCidrOk returns a tuple with the Cidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfIpReservation) GetCidrOk() (*int32, bool) {
	if o == nil || IsNil(o.Cidr) {
		return nil, false
	}
	return o.Cidr, true
}

// HasCidr returns a boolean if a field has been set.
func (o *VrfIpReservation) HasCidr() bool {
	if o != nil && !IsNil(o.Cidr) {
		return true
	}

	return false
}

// SetCidr gets a reference to the given int32 and assigns it to the Cidr field.
func (o *VrfIpReservation) SetCidr(v int32) {
	o.Cidr = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *VrfIpReservation) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfIpReservation) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *VrfIpReservation) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *VrfIpReservation) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *VrfIpReservation) GetCreatedBy() Href {
	if o == nil || IsNil(o.CreatedBy) {
		var ret Href
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfIpReservation) GetCreatedByOk() (*Href, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *VrfIpReservation) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given Href and assigns it to the CreatedBy field.
func (o *VrfIpReservation) SetCreatedBy(v Href) {
	o.CreatedBy = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *VrfIpReservation) GetDetails() string {
	if o == nil || IsNil(o.Details) {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfIpReservation) GetDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *VrfIpReservation) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *VrfIpReservation) SetDetails(v string) {
	o.Details = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *VrfIpReservation) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfIpReservation) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *VrfIpReservation) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *VrfIpReservation) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VrfIpReservation) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfIpReservation) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VrfIpReservation) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *VrfIpReservation) SetId(v string) {
	o.Id = &v
}

// GetMetalGateway returns the MetalGateway field value if set, zero value otherwise.
func (o *VrfIpReservation) GetMetalGateway() MetalGatewayLite {
	if o == nil || IsNil(o.MetalGateway) {
		var ret MetalGatewayLite
		return ret
	}
	return *o.MetalGateway
}

// GetMetalGatewayOk returns a tuple with the MetalGateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfIpReservation) GetMetalGatewayOk() (*MetalGatewayLite, bool) {
	if o == nil || IsNil(o.MetalGateway) {
		return nil, false
	}
	return o.MetalGateway, true
}

// HasMetalGateway returns a boolean if a field has been set.
func (o *VrfIpReservation) HasMetalGateway() bool {
	if o != nil && !IsNil(o.MetalGateway) {
		return true
	}

	return false
}

// SetMetalGateway gets a reference to the given MetalGatewayLite and assigns it to the MetalGateway field.
func (o *VrfIpReservation) SetMetalGateway(v MetalGatewayLite) {
	o.MetalGateway = &v
}

// GetNetmask returns the Netmask field value if set, zero value otherwise.
func (o *VrfIpReservation) GetNetmask() string {
	if o == nil || IsNil(o.Netmask) {
		var ret string
		return ret
	}
	return *o.Netmask
}

// GetNetmaskOk returns a tuple with the Netmask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfIpReservation) GetNetmaskOk() (*string, bool) {
	if o == nil || IsNil(o.Netmask) {
		return nil, false
	}
	return o.Netmask, true
}

// HasNetmask returns a boolean if a field has been set.
func (o *VrfIpReservation) HasNetmask() bool {
	if o != nil && !IsNil(o.Netmask) {
		return true
	}

	return false
}

// SetNetmask gets a reference to the given string and assigns it to the Netmask field.
func (o *VrfIpReservation) SetNetmask(v string) {
	o.Netmask = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *VrfIpReservation) GetNetwork() string {
	if o == nil || IsNil(o.Network) {
		var ret string
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfIpReservation) GetNetworkOk() (*string, bool) {
	if o == nil || IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *VrfIpReservation) HasNetwork() bool {
	if o != nil && !IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given string and assigns it to the Network field.
func (o *VrfIpReservation) SetNetwork(v string) {
	o.Network = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *VrfIpReservation) GetProject() Project {
	if o == nil || IsNil(o.Project) {
		var ret Project
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfIpReservation) GetProjectOk() (*Project, bool) {
	if o == nil || IsNil(o.Project) {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *VrfIpReservation) HasProject() bool {
	if o != nil && !IsNil(o.Project) {
		return true
	}

	return false
}

// SetProject gets a reference to the given Project and assigns it to the Project field.
func (o *VrfIpReservation) SetProject(v Project) {
	o.Project = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *VrfIpReservation) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfIpReservation) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *VrfIpReservation) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *VrfIpReservation) SetState(v string) {
	o.State = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *VrfIpReservation) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfIpReservation) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *VrfIpReservation) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *VrfIpReservation) SetTags(v []string) {
	o.Tags = v
}

// GetType returns the Type field value
func (o *VrfIpReservation) GetType() VrfIpReservationType {
	if o == nil {
		var ret VrfIpReservationType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *VrfIpReservation) GetTypeOk() (*VrfIpReservationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *VrfIpReservation) SetType(v VrfIpReservationType) {
	o.Type = v
}

// GetVrf returns the Vrf field value
func (o *VrfIpReservation) GetVrf() Vrf {
	if o == nil {
		var ret Vrf
		return ret
	}

	return o.Vrf
}

// GetVrfOk returns a tuple with the Vrf field value
// and a boolean to check if the value has been set.
func (o *VrfIpReservation) GetVrfOk() (*Vrf, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vrf, true
}

// SetVrf sets field value
func (o *VrfIpReservation) SetVrf(v Vrf) {
	o.Vrf = v
}

// GetPublic returns the Public field value if set, zero value otherwise.
func (o *VrfIpReservation) GetPublic() bool {
	if o == nil || IsNil(o.Public) {
		var ret bool
		return ret
	}
	return *o.Public
}

// GetPublicOk returns a tuple with the Public field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfIpReservation) GetPublicOk() (*bool, bool) {
	if o == nil || IsNil(o.Public) {
		return nil, false
	}
	return o.Public, true
}

// HasPublic returns a boolean if a field has been set.
func (o *VrfIpReservation) HasPublic() bool {
	if o != nil && !IsNil(o.Public) {
		return true
	}

	return false
}

// SetPublic gets a reference to the given bool and assigns it to the Public field.
func (o *VrfIpReservation) SetPublic(v bool) {
	o.Public = &v
}

// GetManagement returns the Management field value if set, zero value otherwise.
func (o *VrfIpReservation) GetManagement() bool {
	if o == nil || IsNil(o.Management) {
		var ret bool
		return ret
	}
	return *o.Management
}

// GetManagementOk returns a tuple with the Management field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfIpReservation) GetManagementOk() (*bool, bool) {
	if o == nil || IsNil(o.Management) {
		return nil, false
	}
	return o.Management, true
}

// HasManagement returns a boolean if a field has been set.
func (o *VrfIpReservation) HasManagement() bool {
	if o != nil && !IsNil(o.Management) {
		return true
	}

	return false
}

// SetManagement gets a reference to the given bool and assigns it to the Management field.
func (o *VrfIpReservation) SetManagement(v bool) {
	o.Management = &v
}

// GetManageable returns the Manageable field value if set, zero value otherwise.
func (o *VrfIpReservation) GetManageable() bool {
	if o == nil || IsNil(o.Manageable) {
		var ret bool
		return ret
	}
	return *o.Manageable
}

// GetManageableOk returns a tuple with the Manageable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfIpReservation) GetManageableOk() (*bool, bool) {
	if o == nil || IsNil(o.Manageable) {
		return nil, false
	}
	return o.Manageable, true
}

// HasManageable returns a boolean if a field has been set.
func (o *VrfIpReservation) HasManageable() bool {
	if o != nil && !IsNil(o.Manageable) {
		return true
	}

	return false
}

// SetManageable gets a reference to the given bool and assigns it to the Manageable field.
func (o *VrfIpReservation) SetManageable(v bool) {
	o.Manageable = &v
}

// GetCustomdata returns the Customdata field value if set, zero value otherwise.
func (o *VrfIpReservation) GetCustomdata() map[string]interface{} {
	if o == nil || IsNil(o.Customdata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Customdata
}

// GetCustomdataOk returns a tuple with the Customdata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfIpReservation) GetCustomdataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Customdata) {
		return map[string]interface{}{}, false
	}
	return o.Customdata, true
}

// HasCustomdata returns a boolean if a field has been set.
func (o *VrfIpReservation) HasCustomdata() bool {
	if o != nil && !IsNil(o.Customdata) {
		return true
	}

	return false
}

// SetCustomdata gets a reference to the given map[string]interface{} and assigns it to the Customdata field.
func (o *VrfIpReservation) SetCustomdata(v map[string]interface{}) {
	o.Customdata = v
}

// GetBill returns the Bill field value if set, zero value otherwise.
func (o *VrfIpReservation) GetBill() bool {
	if o == nil || IsNil(o.Bill) {
		var ret bool
		return ret
	}
	return *o.Bill
}

// GetBillOk returns a tuple with the Bill field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfIpReservation) GetBillOk() (*bool, bool) {
	if o == nil || IsNil(o.Bill) {
		return nil, false
	}
	return o.Bill, true
}

// HasBill returns a boolean if a field has been set.
func (o *VrfIpReservation) HasBill() bool {
	if o != nil && !IsNil(o.Bill) {
		return true
	}

	return false
}

// SetBill gets a reference to the given bool and assigns it to the Bill field.
func (o *VrfIpReservation) SetBill(v bool) {
	o.Bill = &v
}

// GetProjectLite returns the ProjectLite field value if set, zero value otherwise.
func (o *VrfIpReservation) GetProjectLite() Project {
	if o == nil || IsNil(o.ProjectLite) {
		var ret Project
		return ret
	}
	return *o.ProjectLite
}

// GetProjectLiteOk returns a tuple with the ProjectLite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfIpReservation) GetProjectLiteOk() (*Project, bool) {
	if o == nil || IsNil(o.ProjectLite) {
		return nil, false
	}
	return o.ProjectLite, true
}

// HasProjectLite returns a boolean if a field has been set.
func (o *VrfIpReservation) HasProjectLite() bool {
	if o != nil && !IsNil(o.ProjectLite) {
		return true
	}

	return false
}

// SetProjectLite gets a reference to the given Project and assigns it to the ProjectLite field.
func (o *VrfIpReservation) SetProjectLite(v Project) {
	o.ProjectLite = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *VrfIpReservation) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfIpReservation) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *VrfIpReservation) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *VrfIpReservation) SetAddress(v string) {
	o.Address = &v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *VrfIpReservation) GetGateway() string {
	if o == nil || IsNil(o.Gateway) {
		var ret string
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfIpReservation) GetGatewayOk() (*string, bool) {
	if o == nil || IsNil(o.Gateway) {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *VrfIpReservation) HasGateway() bool {
	if o != nil && !IsNil(o.Gateway) {
		return true
	}

	return false
}

// SetGateway gets a reference to the given string and assigns it to the Gateway field.
func (o *VrfIpReservation) SetGateway(v string) {
	o.Gateway = &v
}

// GetMetro returns the Metro field value if set, zero value otherwise.
func (o *VrfIpReservation) GetMetro() Metro {
	if o == nil || IsNil(o.Metro) {
		var ret Metro
		return ret
	}
	return *o.Metro
}

// GetMetroOk returns a tuple with the Metro field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfIpReservation) GetMetroOk() (*Metro, bool) {
	if o == nil || IsNil(o.Metro) {
		return nil, false
	}
	return o.Metro, true
}

// HasMetro returns a boolean if a field has been set.
func (o *VrfIpReservation) HasMetro() bool {
	if o != nil && !IsNil(o.Metro) {
		return true
	}

	return false
}

// SetMetro gets a reference to the given Metro and assigns it to the Metro field.
func (o *VrfIpReservation) SetMetro(v Metro) {
	o.Metro = &v
}

func (o VrfIpReservation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VrfIpReservation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AddressFamily) {
		toSerialize["address_family"] = o.AddressFamily
	}
	if !IsNil(o.Cidr) {
		toSerialize["cidr"] = o.Cidr
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.CreatedBy) {
		toSerialize["created_by"] = o.CreatedBy
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.MetalGateway) {
		toSerialize["metal_gateway"] = o.MetalGateway
	}
	if !IsNil(o.Netmask) {
		toSerialize["netmask"] = o.Netmask
	}
	if !IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !IsNil(o.Project) {
		toSerialize["project"] = o.Project
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	toSerialize["type"] = o.Type
	toSerialize["vrf"] = o.Vrf
	if !IsNil(o.Public) {
		toSerialize["public"] = o.Public
	}
	if !IsNil(o.Management) {
		toSerialize["management"] = o.Management
	}
	if !IsNil(o.Manageable) {
		toSerialize["manageable"] = o.Manageable
	}
	if !IsNil(o.Customdata) {
		toSerialize["customdata"] = o.Customdata
	}
	if !IsNil(o.Bill) {
		toSerialize["bill"] = o.Bill
	}
	if !IsNil(o.ProjectLite) {
		toSerialize["project_lite"] = o.ProjectLite
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.Gateway) {
		toSerialize["gateway"] = o.Gateway
	}
	if !IsNil(o.Metro) {
		toSerialize["metro"] = o.Metro
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VrfIpReservation) UnmarshalJSON(bytes []byte) (err error) {
	varVrfIpReservation := _VrfIpReservation{}

	err = json.Unmarshal(bytes, &varVrfIpReservation)

	if err != nil {
		return err
	}

	*o = VrfIpReservation(varVrfIpReservation)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "address_family")
		delete(additionalProperties, "cidr")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "created_by")
		delete(additionalProperties, "details")
		delete(additionalProperties, "href")
		delete(additionalProperties, "id")
		delete(additionalProperties, "metal_gateway")
		delete(additionalProperties, "netmask")
		delete(additionalProperties, "network")
		delete(additionalProperties, "project")
		delete(additionalProperties, "state")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "type")
		delete(additionalProperties, "vrf")
		delete(additionalProperties, "public")
		delete(additionalProperties, "management")
		delete(additionalProperties, "manageable")
		delete(additionalProperties, "customdata")
		delete(additionalProperties, "bill")
		delete(additionalProperties, "project_lite")
		delete(additionalProperties, "address")
		delete(additionalProperties, "gateway")
		delete(additionalProperties, "metro")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVrfIpReservation struct {
	value *VrfIpReservation
	isSet bool
}

func (v NullableVrfIpReservation) Get() *VrfIpReservation {
	return v.value
}

func (v *NullableVrfIpReservation) Set(val *VrfIpReservation) {
	v.value = val
	v.isSet = true
}

func (v NullableVrfIpReservation) IsSet() bool {
	return v.isSet
}

func (v *NullableVrfIpReservation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVrfIpReservation(val *VrfIpReservation) *NullableVrfIpReservation {
	return &NullableVrfIpReservation{value: val, isSet: true}
}

func (v NullableVrfIpReservation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVrfIpReservation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
