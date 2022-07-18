# VirtualCircuit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | **[]string** |  | 
**Bill** | **bool** | True if the Virtual Circuit is being billed. Currently, only Virtual Circuits that are created with A-side service tokens will be billed. Usage will start the first time the Virtual Circuit becomes active, and will not stop until it is deleted. | [default to false]
**Description** | **string** |  | 
**Id** | **string** |  | 
**Name** | **string** |  | 
**NniVlan** | **int32** |  | 
**Port** | [**FindBatchById200ResponseDevicesInner**](FindBatchById200ResponseDevicesInner.md) |  | 
**Project** | [**FindBatchById200ResponseDevicesInner**](FindBatchById200ResponseDevicesInner.md) |  | 
**Speed** | Pointer to **int32** | integer representing bps speed | [optional] 
**Status** | **string** |  | 
**VirtualNetwork** | [**FindBatchById200ResponseDevicesInner**](FindBatchById200ResponseDevicesInner.md) |  | 
**Vnid** | **int32** |  | 

## Methods

### NewVirtualCircuit

`func NewVirtualCircuit(tags []string, bill bool, description string, id string, name string, nniVlan int32, port FindBatchById200ResponseDevicesInner, project FindBatchById200ResponseDevicesInner, status string, virtualNetwork FindBatchById200ResponseDevicesInner, vnid int32, ) *VirtualCircuit`

NewVirtualCircuit instantiates a new VirtualCircuit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualCircuitWithDefaults

`func NewVirtualCircuitWithDefaults() *VirtualCircuit`

NewVirtualCircuitWithDefaults instantiates a new VirtualCircuit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *VirtualCircuit) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VirtualCircuit) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VirtualCircuit) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetBill

`func (o *VirtualCircuit) GetBill() bool`

GetBill returns the Bill field if non-nil, zero value otherwise.

### GetBillOk

`func (o *VirtualCircuit) GetBillOk() (*bool, bool)`

GetBillOk returns a tuple with the Bill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBill

`func (o *VirtualCircuit) SetBill(v bool)`

SetBill sets Bill field to given value.


### GetDescription

`func (o *VirtualCircuit) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VirtualCircuit) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VirtualCircuit) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetId

`func (o *VirtualCircuit) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VirtualCircuit) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VirtualCircuit) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *VirtualCircuit) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualCircuit) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualCircuit) SetName(v string)`

SetName sets Name field to given value.


### GetNniVlan

`func (o *VirtualCircuit) GetNniVlan() int32`

GetNniVlan returns the NniVlan field if non-nil, zero value otherwise.

### GetNniVlanOk

`func (o *VirtualCircuit) GetNniVlanOk() (*int32, bool)`

GetNniVlanOk returns a tuple with the NniVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNniVlan

`func (o *VirtualCircuit) SetNniVlan(v int32)`

SetNniVlan sets NniVlan field to given value.


### GetPort

`func (o *VirtualCircuit) GetPort() FindBatchById200ResponseDevicesInner`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *VirtualCircuit) GetPortOk() (*FindBatchById200ResponseDevicesInner, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *VirtualCircuit) SetPort(v FindBatchById200ResponseDevicesInner)`

SetPort sets Port field to given value.


### GetProject

`func (o *VirtualCircuit) GetProject() FindBatchById200ResponseDevicesInner`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *VirtualCircuit) GetProjectOk() (*FindBatchById200ResponseDevicesInner, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *VirtualCircuit) SetProject(v FindBatchById200ResponseDevicesInner)`

SetProject sets Project field to given value.


### GetSpeed

`func (o *VirtualCircuit) GetSpeed() int32`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *VirtualCircuit) GetSpeedOk() (*int32, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *VirtualCircuit) SetSpeed(v int32)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *VirtualCircuit) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetStatus

`func (o *VirtualCircuit) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VirtualCircuit) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VirtualCircuit) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetVirtualNetwork

`func (o *VirtualCircuit) GetVirtualNetwork() FindBatchById200ResponseDevicesInner`

GetVirtualNetwork returns the VirtualNetwork field if non-nil, zero value otherwise.

### GetVirtualNetworkOk

`func (o *VirtualCircuit) GetVirtualNetworkOk() (*FindBatchById200ResponseDevicesInner, bool)`

GetVirtualNetworkOk returns a tuple with the VirtualNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualNetwork

`func (o *VirtualCircuit) SetVirtualNetwork(v FindBatchById200ResponseDevicesInner)`

SetVirtualNetwork sets VirtualNetwork field to given value.


### GetVnid

`func (o *VirtualCircuit) GetVnid() int32`

GetVnid returns the Vnid field if non-nil, zero value otherwise.

### GetVnidOk

`func (o *VirtualCircuit) GetVnidOk() (*int32, bool)`

GetVnidOk returns a tuple with the Vnid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnid

`func (o *VirtualCircuit) SetVnid(v int32)`

SetVnid sets Vnid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


