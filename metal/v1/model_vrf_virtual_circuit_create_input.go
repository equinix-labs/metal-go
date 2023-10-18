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
)

// checks if the VrfVirtualCircuitCreateInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VrfVirtualCircuitCreateInput{}

// VrfVirtualCircuitCreateInput struct for VrfVirtualCircuitCreateInput
type VrfVirtualCircuitCreateInput struct {
	// An IP address from the subnet that will be used on the Customer side. This parameter is optional, but if supplied, we will use the other usable IP address in the subnet as the Metal IP. By default, the last usable IP address in the subnet will be used.
	CustomerIp  *string `json:"customer_ip,omitempty"`
	Description *string `json:"description,omitempty"`
	// The plaintext BGP peering password shared by neighbors as an MD5 checksum: * must be 10-20 characters long * may not include punctuation * must be a combination of numbers and letters * must contain at least one lowercase, uppercase, and digit character
	Md5 NullableString `json:"md5,omitempty"`
	// An IP address from the subnet that will be used on the Metal side. This parameter is optional, but if supplied, we will use the other usable IP address in the subnet as the Customer IP. By default, the first usable IP address in the subnet will be used.
	MetalIp *string `json:"metal_ip,omitempty"`
	Name    *string `json:"name,omitempty"`
	NniVlan int32   `json:"nni_vlan"`
	// The peer ASN that will be used with the VRF on the Virtual Circuit.
	PeerAsn   int32  `json:"peer_asn"`
	ProjectId string `json:"project_id"`
	// speed can be passed as integer number representing bps speed or string (e.g. '52m' or '100g' or '4 gbps')
	Speed *int32 `json:"speed,omitempty"`
	// The /30 or /31 subnet of one of the VRF IP Blocks that will be used with the VRF for the Virtual Circuit. This subnet does not have to be an existing VRF IP reservation, as we will create the VRF IP reservation on creation if it does not exist. The Metal IP and Customer IP must be IPs from this subnet. For /30 subnets, the network and broadcast IPs cannot be used as the Metal or Customer IP. The subnet specified must be contained within an already-defined IP Range for the VRF.
	Subnet string   `json:"subnet"`
	Tags   []string `json:"tags,omitempty"`
	// The UUID of the VRF that will be associated with the Virtual Circuit.
	Vrf                  string `json:"vrf"`
	AdditionalProperties map[string]interface{}
}

type _VrfVirtualCircuitCreateInput VrfVirtualCircuitCreateInput

// NewVrfVirtualCircuitCreateInput instantiates a new VrfVirtualCircuitCreateInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVrfVirtualCircuitCreateInput(nniVlan int32, peerAsn int32, projectId string, subnet string, vrf string) *VrfVirtualCircuitCreateInput {
	this := VrfVirtualCircuitCreateInput{}
	this.NniVlan = nniVlan
	this.PeerAsn = peerAsn
	this.ProjectId = projectId
	this.Subnet = subnet
	this.Vrf = vrf
	return &this
}

// NewVrfVirtualCircuitCreateInputWithDefaults instantiates a new VrfVirtualCircuitCreateInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVrfVirtualCircuitCreateInputWithDefaults() *VrfVirtualCircuitCreateInput {
	this := VrfVirtualCircuitCreateInput{}
	return &this
}

// GetCustomerIp returns the CustomerIp field value if set, zero value otherwise.
func (o *VrfVirtualCircuitCreateInput) GetCustomerIp() string {
	if o == nil || IsNil(o.CustomerIp) {
		var ret string
		return ret
	}
	return *o.CustomerIp
}

// GetCustomerIpOk returns a tuple with the CustomerIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfVirtualCircuitCreateInput) GetCustomerIpOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerIp) {
		return nil, false
	}
	return o.CustomerIp, true
}

// HasCustomerIp returns a boolean if a field has been set.
func (o *VrfVirtualCircuitCreateInput) HasCustomerIp() bool {
	if o != nil && !IsNil(o.CustomerIp) {
		return true
	}

	return false
}

// SetCustomerIp gets a reference to the given string and assigns it to the CustomerIp field.
func (o *VrfVirtualCircuitCreateInput) SetCustomerIp(v string) {
	o.CustomerIp = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *VrfVirtualCircuitCreateInput) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfVirtualCircuitCreateInput) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *VrfVirtualCircuitCreateInput) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *VrfVirtualCircuitCreateInput) SetDescription(v string) {
	o.Description = &v
}

// GetMd5 returns the Md5 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VrfVirtualCircuitCreateInput) GetMd5() string {
	if o == nil || IsNil(o.Md5.Get()) {
		var ret string
		return ret
	}
	return *o.Md5.Get()
}

// GetMd5Ok returns a tuple with the Md5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VrfVirtualCircuitCreateInput) GetMd5Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Md5.Get(), o.Md5.IsSet()
}

// HasMd5 returns a boolean if a field has been set.
func (o *VrfVirtualCircuitCreateInput) HasMd5() bool {
	if o != nil && o.Md5.IsSet() {
		return true
	}

	return false
}

// SetMd5 gets a reference to the given NullableString and assigns it to the Md5 field.
func (o *VrfVirtualCircuitCreateInput) SetMd5(v string) {
	o.Md5.Set(&v)
}

// SetMd5Nil sets the value for Md5 to be an explicit nil
func (o *VrfVirtualCircuitCreateInput) SetMd5Nil() {
	o.Md5.Set(nil)
}

// UnsetMd5 ensures that no value is present for Md5, not even an explicit nil
func (o *VrfVirtualCircuitCreateInput) UnsetMd5() {
	o.Md5.Unset()
}

// GetMetalIp returns the MetalIp field value if set, zero value otherwise.
func (o *VrfVirtualCircuitCreateInput) GetMetalIp() string {
	if o == nil || IsNil(o.MetalIp) {
		var ret string
		return ret
	}
	return *o.MetalIp
}

// GetMetalIpOk returns a tuple with the MetalIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfVirtualCircuitCreateInput) GetMetalIpOk() (*string, bool) {
	if o == nil || IsNil(o.MetalIp) {
		return nil, false
	}
	return o.MetalIp, true
}

// HasMetalIp returns a boolean if a field has been set.
func (o *VrfVirtualCircuitCreateInput) HasMetalIp() bool {
	if o != nil && !IsNil(o.MetalIp) {
		return true
	}

	return false
}

// SetMetalIp gets a reference to the given string and assigns it to the MetalIp field.
func (o *VrfVirtualCircuitCreateInput) SetMetalIp(v string) {
	o.MetalIp = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VrfVirtualCircuitCreateInput) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfVirtualCircuitCreateInput) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VrfVirtualCircuitCreateInput) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VrfVirtualCircuitCreateInput) SetName(v string) {
	o.Name = &v
}

// GetNniVlan returns the NniVlan field value
func (o *VrfVirtualCircuitCreateInput) GetNniVlan() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NniVlan
}

// GetNniVlanOk returns a tuple with the NniVlan field value
// and a boolean to check if the value has been set.
func (o *VrfVirtualCircuitCreateInput) GetNniVlanOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NniVlan, true
}

// SetNniVlan sets field value
func (o *VrfVirtualCircuitCreateInput) SetNniVlan(v int32) {
	o.NniVlan = v
}

// GetPeerAsn returns the PeerAsn field value
func (o *VrfVirtualCircuitCreateInput) GetPeerAsn() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PeerAsn
}

// GetPeerAsnOk returns a tuple with the PeerAsn field value
// and a boolean to check if the value has been set.
func (o *VrfVirtualCircuitCreateInput) GetPeerAsnOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PeerAsn, true
}

// SetPeerAsn sets field value
func (o *VrfVirtualCircuitCreateInput) SetPeerAsn(v int32) {
	o.PeerAsn = v
}

// GetProjectId returns the ProjectId field value
func (o *VrfVirtualCircuitCreateInput) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *VrfVirtualCircuitCreateInput) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *VrfVirtualCircuitCreateInput) SetProjectId(v string) {
	o.ProjectId = v
}

// GetSpeed returns the Speed field value if set, zero value otherwise.
func (o *VrfVirtualCircuitCreateInput) GetSpeed() int32 {
	if o == nil || IsNil(o.Speed) {
		var ret int32
		return ret
	}
	return *o.Speed
}

// GetSpeedOk returns a tuple with the Speed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfVirtualCircuitCreateInput) GetSpeedOk() (*int32, bool) {
	if o == nil || IsNil(o.Speed) {
		return nil, false
	}
	return o.Speed, true
}

// HasSpeed returns a boolean if a field has been set.
func (o *VrfVirtualCircuitCreateInput) HasSpeed() bool {
	if o != nil && !IsNil(o.Speed) {
		return true
	}

	return false
}

// SetSpeed gets a reference to the given int32 and assigns it to the Speed field.
func (o *VrfVirtualCircuitCreateInput) SetSpeed(v int32) {
	o.Speed = &v
}

// GetSubnet returns the Subnet field value
func (o *VrfVirtualCircuitCreateInput) GetSubnet() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subnet
}

// GetSubnetOk returns a tuple with the Subnet field value
// and a boolean to check if the value has been set.
func (o *VrfVirtualCircuitCreateInput) GetSubnetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subnet, true
}

// SetSubnet sets field value
func (o *VrfVirtualCircuitCreateInput) SetSubnet(v string) {
	o.Subnet = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *VrfVirtualCircuitCreateInput) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfVirtualCircuitCreateInput) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *VrfVirtualCircuitCreateInput) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *VrfVirtualCircuitCreateInput) SetTags(v []string) {
	o.Tags = v
}

// GetVrf returns the Vrf field value
func (o *VrfVirtualCircuitCreateInput) GetVrf() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Vrf
}

// GetVrfOk returns a tuple with the Vrf field value
// and a boolean to check if the value has been set.
func (o *VrfVirtualCircuitCreateInput) GetVrfOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vrf, true
}

// SetVrf sets field value
func (o *VrfVirtualCircuitCreateInput) SetVrf(v string) {
	o.Vrf = v
}

func (o VrfVirtualCircuitCreateInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VrfVirtualCircuitCreateInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CustomerIp) {
		toSerialize["customer_ip"] = o.CustomerIp
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if o.Md5.IsSet() {
		toSerialize["md5"] = o.Md5.Get()
	}
	if !IsNil(o.MetalIp) {
		toSerialize["metal_ip"] = o.MetalIp
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["nni_vlan"] = o.NniVlan
	toSerialize["peer_asn"] = o.PeerAsn
	toSerialize["project_id"] = o.ProjectId
	if !IsNil(o.Speed) {
		toSerialize["speed"] = o.Speed
	}
	toSerialize["subnet"] = o.Subnet
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	toSerialize["vrf"] = o.Vrf

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VrfVirtualCircuitCreateInput) UnmarshalJSON(bytes []byte) (err error) {
	requiredProperties := []string{
		"nni_vlan",
		"peer_asn",
		"project_id",
		"subnet",
		"vrf",
	}

	allProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &allProperties); err == nil {
		for _, requiredProperty := range requiredProperties {
			if _, exists := allProperties[requiredProperty]; !exists {
				return MissingRequiredFieldError(requiredProperty)
			}
		}
	}

	varVrfVirtualCircuitCreateInput := _VrfVirtualCircuitCreateInput{}

	err = json.Unmarshal(bytes, &varVrfVirtualCircuitCreateInput)

	if err != nil {
		return err
	}

	*o = VrfVirtualCircuitCreateInput(varVrfVirtualCircuitCreateInput)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "customer_ip")
		delete(additionalProperties, "description")
		delete(additionalProperties, "md5")
		delete(additionalProperties, "metal_ip")
		delete(additionalProperties, "name")
		delete(additionalProperties, "nni_vlan")
		delete(additionalProperties, "peer_asn")
		delete(additionalProperties, "project_id")
		delete(additionalProperties, "speed")
		delete(additionalProperties, "subnet")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "vrf")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVrfVirtualCircuitCreateInput struct {
	value *VrfVirtualCircuitCreateInput
	isSet bool
}

func (v NullableVrfVirtualCircuitCreateInput) Get() *VrfVirtualCircuitCreateInput {
	return v.value
}

func (v *NullableVrfVirtualCircuitCreateInput) Set(val *VrfVirtualCircuitCreateInput) {
	v.value = val
	v.isSet = true
}

func (v NullableVrfVirtualCircuitCreateInput) IsSet() bool {
	return v.isSet
}

func (v *NullableVrfVirtualCircuitCreateInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVrfVirtualCircuitCreateInput(val *VrfVirtualCircuitCreateInput) *NullableVrfVirtualCircuitCreateInput {
	return &NullableVrfVirtualCircuitCreateInput{value: val, isSet: true}
}

func (v NullableVrfVirtualCircuitCreateInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVrfVirtualCircuitCreateInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
