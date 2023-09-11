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
	"time"
)

// checks if the VlanVirtualCircuit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VlanVirtualCircuit{}

// VlanVirtualCircuit struct for VlanVirtualCircuit
type VlanVirtualCircuit struct {
	// True if the Virtual Circuit is being billed. Currently, only Virtual Circuits of Fabric VCs (Metal Billed) will be billed. Usage will start the first time the Virtual Circuit becomes active, and will not stop until it is deleted from Metal.
	Bill        bool   `json:"bill"`
	Description string `json:"description"`
	Id          string `json:"id"`
	Name        string `json:"name"`
	NniVlan     int32  `json:"nni_vlan"`
	Port        Href   `json:"port"`
	Project     Href   `json:"project"`
	// For Virtual Circuits on shared and dedicated connections, this speed should match the one set on their Interconnection Ports. For Virtual Circuits on Fabric VCs (both Metal and Fabric Billed) that have found their corresponding Fabric connection, this is the actual speed of the interconnection that was configured when setting up the interconnection on the Fabric Portal. Details on Fabric VCs are included in the specification as a developer preview and is generally unavailable. Please contact our Support team for more details.
	Speed *int32 `json:"speed,omitempty"`
	// The status of a Virtual Circuit is always 'pending' on creation. The status can turn to 'Waiting on Customer VLAN' if a Metro VLAN was not set yet on the Virtual Circuit and is the last step needed for full activation. For Dedicated interconnections, as long as the Dedicated Port has been associated to the Virtual Circuit and a NNI VNID has been set, it will turn to 'waiting_on_customer_vlan'. For Fabric VCs, it will only change to 'waiting_on_customer_vlan' once the corresponding Fabric connection has been found on the Fabric side. If the Fabric service token associated with the Virtual Circuit hasn't been redeemed on Fabric within the expiry time, it will change to an `expired` status. Once a Metro VLAN is set on the Virtual Circuit (which for Fabric VCs, can be set on creation of a Fabric VC) and the necessary set up is done, it will turn into 'Activating' status as it tries to activate the Virtual Circuit. Once the Virtual Circuit fully activates and is configured on the switch, it will turn to staus 'active'. For Fabric VCs (Metal Billed), we will start billing the moment the status of the Virtual Circuit turns to 'active'. If there are any changes to the VLAN after the Virtual Circuit is in an 'active' status, the status will show 'changing_vlan' if a new VLAN has been provided, or 'deactivating' if we are removing the VLAN. When a deletion request is issued for the Virtual Circuit, it will move to a 'deleting' status, and we will immediately unconfigure the switch for the Virtual Circuit and issue a deletion on any associated Fabric connections. Any associated Metro VLANs on the virtual circuit will also be unassociated after the switch has been successfully unconfigured. If there are any associated Fabric connections, we will only fully delete the Virtual Circuit once we have checked that the Fabric connection was fully deprovisioned on Fabric.
	Status               string     `json:"status"`
	Tags                 []string   `json:"tags"`
	VirtualNetwork       Href       `json:"virtual_network"`
	Vnid                 int32      `json:"vnid"`
	CreatedAt            *time.Time `json:"created_at,omitempty"`
	UpdatedAt            *time.Time `json:"updated_at,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VlanVirtualCircuit VlanVirtualCircuit

// NewVlanVirtualCircuit instantiates a new VlanVirtualCircuit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVlanVirtualCircuit(bill bool, description string, id string, name string, nniVlan int32, port Href, project Href, status string, tags []string, virtualNetwork Href, vnid int32) *VlanVirtualCircuit {
	this := VlanVirtualCircuit{}
	this.Bill = bill
	this.Description = description
	this.Id = id
	this.Name = name
	this.NniVlan = nniVlan
	this.Port = port
	this.Project = project
	this.Status = status
	this.Tags = tags
	this.VirtualNetwork = virtualNetwork
	this.Vnid = vnid
	return &this
}

// NewVlanVirtualCircuitWithDefaults instantiates a new VlanVirtualCircuit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVlanVirtualCircuitWithDefaults() *VlanVirtualCircuit {
	this := VlanVirtualCircuit{}
	var bill bool = false
	this.Bill = bill
	return &this
}

// GetBill returns the Bill field value
func (o *VlanVirtualCircuit) GetBill() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Bill
}

// GetBillOk returns a tuple with the Bill field value
// and a boolean to check if the value has been set.
func (o *VlanVirtualCircuit) GetBillOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bill, true
}

// SetBill sets field value
func (o *VlanVirtualCircuit) SetBill(v bool) {
	o.Bill = v
}

// GetDescription returns the Description field value
func (o *VlanVirtualCircuit) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *VlanVirtualCircuit) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *VlanVirtualCircuit) SetDescription(v string) {
	o.Description = v
}

// GetId returns the Id field value
func (o *VlanVirtualCircuit) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *VlanVirtualCircuit) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *VlanVirtualCircuit) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *VlanVirtualCircuit) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VlanVirtualCircuit) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VlanVirtualCircuit) SetName(v string) {
	o.Name = v
}

// GetNniVlan returns the NniVlan field value
func (o *VlanVirtualCircuit) GetNniVlan() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NniVlan
}

// GetNniVlanOk returns a tuple with the NniVlan field value
// and a boolean to check if the value has been set.
func (o *VlanVirtualCircuit) GetNniVlanOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NniVlan, true
}

// SetNniVlan sets field value
func (o *VlanVirtualCircuit) SetNniVlan(v int32) {
	o.NniVlan = v
}

// GetPort returns the Port field value
func (o *VlanVirtualCircuit) GetPort() Href {
	if o == nil {
		var ret Href
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *VlanVirtualCircuit) GetPortOk() (*Href, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *VlanVirtualCircuit) SetPort(v Href) {
	o.Port = v
}

// GetProject returns the Project field value
func (o *VlanVirtualCircuit) GetProject() Href {
	if o == nil {
		var ret Href
		return ret
	}

	return o.Project
}

// GetProjectOk returns a tuple with the Project field value
// and a boolean to check if the value has been set.
func (o *VlanVirtualCircuit) GetProjectOk() (*Href, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Project, true
}

// SetProject sets field value
func (o *VlanVirtualCircuit) SetProject(v Href) {
	o.Project = v
}

// GetSpeed returns the Speed field value if set, zero value otherwise.
func (o *VlanVirtualCircuit) GetSpeed() int32 {
	if o == nil || IsNil(o.Speed) {
		var ret int32
		return ret
	}
	return *o.Speed
}

// GetSpeedOk returns a tuple with the Speed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VlanVirtualCircuit) GetSpeedOk() (*int32, bool) {
	if o == nil || IsNil(o.Speed) {
		return nil, false
	}
	return o.Speed, true
}

// HasSpeed returns a boolean if a field has been set.
func (o *VlanVirtualCircuit) HasSpeed() bool {
	if o != nil && !IsNil(o.Speed) {
		return true
	}

	return false
}

// SetSpeed gets a reference to the given int32 and assigns it to the Speed field.
func (o *VlanVirtualCircuit) SetSpeed(v int32) {
	o.Speed = &v
}

// GetStatus returns the Status field value
func (o *VlanVirtualCircuit) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *VlanVirtualCircuit) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *VlanVirtualCircuit) SetStatus(v string) {
	o.Status = v
}

// GetTags returns the Tags field value
func (o *VlanVirtualCircuit) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *VlanVirtualCircuit) GetTagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *VlanVirtualCircuit) SetTags(v []string) {
	o.Tags = v
}

// GetVirtualNetwork returns the VirtualNetwork field value
func (o *VlanVirtualCircuit) GetVirtualNetwork() Href {
	if o == nil {
		var ret Href
		return ret
	}

	return o.VirtualNetwork
}

// GetVirtualNetworkOk returns a tuple with the VirtualNetwork field value
// and a boolean to check if the value has been set.
func (o *VlanVirtualCircuit) GetVirtualNetworkOk() (*Href, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VirtualNetwork, true
}

// SetVirtualNetwork sets field value
func (o *VlanVirtualCircuit) SetVirtualNetwork(v Href) {
	o.VirtualNetwork = v
}

// GetVnid returns the Vnid field value
func (o *VlanVirtualCircuit) GetVnid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Vnid
}

// GetVnidOk returns a tuple with the Vnid field value
// and a boolean to check if the value has been set.
func (o *VlanVirtualCircuit) GetVnidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vnid, true
}

// SetVnid sets field value
func (o *VlanVirtualCircuit) SetVnid(v int32) {
	o.Vnid = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *VlanVirtualCircuit) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VlanVirtualCircuit) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *VlanVirtualCircuit) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *VlanVirtualCircuit) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *VlanVirtualCircuit) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VlanVirtualCircuit) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *VlanVirtualCircuit) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *VlanVirtualCircuit) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o VlanVirtualCircuit) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VlanVirtualCircuit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bill"] = o.Bill
	toSerialize["description"] = o.Description
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["nni_vlan"] = o.NniVlan
	toSerialize["port"] = o.Port
	toSerialize["project"] = o.Project
	if !IsNil(o.Speed) {
		toSerialize["speed"] = o.Speed
	}
	toSerialize["status"] = o.Status
	toSerialize["tags"] = o.Tags
	toSerialize["virtual_network"] = o.VirtualNetwork
	toSerialize["vnid"] = o.Vnid
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VlanVirtualCircuit) UnmarshalJSON(bytes []byte) (err error) {
	varVlanVirtualCircuit := _VlanVirtualCircuit{}

	err = json.Unmarshal(bytes, &varVlanVirtualCircuit)

	if err != nil {
		return err
	}

	*o = VlanVirtualCircuit(varVlanVirtualCircuit)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "bill")
		delete(additionalProperties, "description")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "nni_vlan")
		delete(additionalProperties, "port")
		delete(additionalProperties, "project")
		delete(additionalProperties, "speed")
		delete(additionalProperties, "status")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "virtual_network")
		delete(additionalProperties, "vnid")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVlanVirtualCircuit struct {
	value *VlanVirtualCircuit
	isSet bool
}

func (v NullableVlanVirtualCircuit) Get() *VlanVirtualCircuit {
	return v.value
}

func (v *NullableVlanVirtualCircuit) Set(val *VlanVirtualCircuit) {
	v.value = val
	v.isSet = true
}

func (v NullableVlanVirtualCircuit) IsSet() bool {
	return v.isSet
}

func (v *NullableVlanVirtualCircuit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVlanVirtualCircuit(val *VlanVirtualCircuit) *NullableVlanVirtualCircuit {
	return &NullableVlanVirtualCircuit{value: val, isSet: true}
}

func (v NullableVlanVirtualCircuit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVlanVirtualCircuit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
