# VlanVirtualCircuit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bill** | **bool** | True if the Virtual Circuit is being billed. Currently, only Virtual Circuits of Fabric VCs (Metal Billed) will be billed. Usage will start the first time the Virtual Circuit becomes active, and will not stop until it is deleted from Metal. | [default to false]
**Description** | **string** |  | 
**Id** | **string** |  | 
**Name** | **string** |  | 
**NniVlan** | **int32** |  | 
**Port** | [**Href**](Href.md) |  | 
**Project** | [**Href**](Href.md) |  | 
**Speed** | Pointer to **int32** | For Virtual Circuits on shared and dedicated connections, this speed should match the one set on their Interconnection Ports. For Virtual Circuits on Fabric VCs (both Metal and Fabric Billed) that have found their corresponding Fabric connection, this is the actual speed of the interconnection that was configured when setting up the interconnection on the Fabric Portal. Details on Fabric VCs are included in the specification as a developer preview and is generally unavailable. Please contact our Support team for more details. | [optional] 
**Status** | **string** | The status of a Virtual Circuit is always &#39;pending&#39; on creation. The status can turn to &#39;Waiting on Customer VLAN&#39; if a Metro VLAN was not set yet on the Virtual Circuit and is the last step needed for full activation. For Dedicated interconnections, as long as the Dedicated Port has been associated to the Virtual Circuit and a NNI VNID has been set, it will turn to &#39;waiting_on_customer_vlan&#39;. For Fabric VCs, it will only change to &#39;waiting_on_customer_vlan&#39; once the corresponding Fabric connection has been found on the Fabric side. If the Fabric service token associated with the Virtual Circuit hasn&#39;t been redeemed on Fabric within the expiry time, it will change to an &#x60;expired&#x60; status. Once a Metro VLAN is set on the Virtual Circuit (which for Fabric VCs, can be set on creation of a Fabric VC) and the necessary set up is done, it will turn into &#39;Activating&#39; status as it tries to activate the Virtual Circuit. Once the Virtual Circuit fully activates and is configured on the switch, it will turn to staus &#39;active&#39;. For Fabric VCs (Metal Billed), we will start billing the moment the status of the Virtual Circuit turns to &#39;active&#39;. If there are any changes to the VLAN after the Virtual Circuit is in an &#39;active&#39; status, the status will show &#39;changing_vlan&#39; if a new VLAN has been provided, or &#39;deactivating&#39; if we are removing the VLAN. When a deletion request is issued for the Virtual Circuit, it will move to a &#39;deleting&#39; status, and we will immediately unconfigure the switch for the Virtual Circuit and issue a deletion on any associated Fabric connections. Any associated Metro VLANs on the virtual circuit will also be unassociated after the switch has been successfully unconfigured. If there are any associated Fabric connections, we will only fully delete the Virtual Circuit once we have checked that the Fabric connection was fully deprovisioned on Fabric. | 
**Tags** | **[]string** |  | 
**VirtualNetwork** | [**Href**](Href.md) |  | 
**Vnid** | **int32** |  | 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewVlanVirtualCircuit

`func NewVlanVirtualCircuit(bill bool, description string, id string, name string, nniVlan int32, port Href, project Href, status string, tags []string, virtualNetwork Href, vnid int32, ) *VlanVirtualCircuit`

NewVlanVirtualCircuit instantiates a new VlanVirtualCircuit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVlanVirtualCircuitWithDefaults

`func NewVlanVirtualCircuitWithDefaults() *VlanVirtualCircuit`

NewVlanVirtualCircuitWithDefaults instantiates a new VlanVirtualCircuit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBill

`func (o *VlanVirtualCircuit) GetBill() bool`

GetBill returns the Bill field if non-nil, zero value otherwise.

### GetBillOk

`func (o *VlanVirtualCircuit) GetBillOk() (*bool, bool)`

GetBillOk returns a tuple with the Bill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBill

`func (o *VlanVirtualCircuit) SetBill(v bool)`

SetBill sets Bill field to given value.


### GetDescription

`func (o *VlanVirtualCircuit) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VlanVirtualCircuit) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VlanVirtualCircuit) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetId

`func (o *VlanVirtualCircuit) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VlanVirtualCircuit) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VlanVirtualCircuit) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *VlanVirtualCircuit) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VlanVirtualCircuit) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VlanVirtualCircuit) SetName(v string)`

SetName sets Name field to given value.


### GetNniVlan

`func (o *VlanVirtualCircuit) GetNniVlan() int32`

GetNniVlan returns the NniVlan field if non-nil, zero value otherwise.

### GetNniVlanOk

`func (o *VlanVirtualCircuit) GetNniVlanOk() (*int32, bool)`

GetNniVlanOk returns a tuple with the NniVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNniVlan

`func (o *VlanVirtualCircuit) SetNniVlan(v int32)`

SetNniVlan sets NniVlan field to given value.


### GetPort

`func (o *VlanVirtualCircuit) GetPort() Href`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *VlanVirtualCircuit) GetPortOk() (*Href, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *VlanVirtualCircuit) SetPort(v Href)`

SetPort sets Port field to given value.


### GetProject

`func (o *VlanVirtualCircuit) GetProject() Href`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *VlanVirtualCircuit) GetProjectOk() (*Href, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *VlanVirtualCircuit) SetProject(v Href)`

SetProject sets Project field to given value.


### GetSpeed

`func (o *VlanVirtualCircuit) GetSpeed() int32`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *VlanVirtualCircuit) GetSpeedOk() (*int32, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *VlanVirtualCircuit) SetSpeed(v int32)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *VlanVirtualCircuit) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetStatus

`func (o *VlanVirtualCircuit) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VlanVirtualCircuit) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VlanVirtualCircuit) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTags

`func (o *VlanVirtualCircuit) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VlanVirtualCircuit) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VlanVirtualCircuit) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetVirtualNetwork

`func (o *VlanVirtualCircuit) GetVirtualNetwork() Href`

GetVirtualNetwork returns the VirtualNetwork field if non-nil, zero value otherwise.

### GetVirtualNetworkOk

`func (o *VlanVirtualCircuit) GetVirtualNetworkOk() (*Href, bool)`

GetVirtualNetworkOk returns a tuple with the VirtualNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualNetwork

`func (o *VlanVirtualCircuit) SetVirtualNetwork(v Href)`

SetVirtualNetwork sets VirtualNetwork field to given value.


### GetVnid

`func (o *VlanVirtualCircuit) GetVnid() int32`

GetVnid returns the Vnid field if non-nil, zero value otherwise.

### GetVnidOk

`func (o *VlanVirtualCircuit) GetVnidOk() (*int32, bool)`

GetVnidOk returns a tuple with the Vnid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnid

`func (o *VlanVirtualCircuit) SetVnid(v int32)`

SetVnid sets Vnid field to given value.


### GetCreatedAt

`func (o *VlanVirtualCircuit) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VlanVirtualCircuit) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VlanVirtualCircuit) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *VlanVirtualCircuit) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *VlanVirtualCircuit) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *VlanVirtualCircuit) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *VlanVirtualCircuit) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *VlanVirtualCircuit) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


