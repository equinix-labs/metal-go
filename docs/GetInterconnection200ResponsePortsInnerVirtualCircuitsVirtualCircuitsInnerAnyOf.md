# GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bill** | **bool** | True if the Virtual Circuit is being billed. Currently, only Virtual Circuits of Fabric VCs (Metal Billed) will be billed. Usage will start the first time the Virtual Circuit becomes active, and will not stop until it is deleted from Metal. | [default to false]
**Description** | **string** |  | 
**Id** | **string** |  | 
**Name** | **string** |  | 
**NniVlan** | **int32** |  | 
**Port** | [**FindBatchById200ResponseDevicesInner**](FindBatchById200ResponseDevicesInner.md) |  | 
**Project** | [**FindBatchById200ResponseDevicesInner**](FindBatchById200ResponseDevicesInner.md) |  | 
**Speed** | Pointer to **int32** | For Virtual Circuits on shared and dedicated connections, this speed should match the one set on their Interconnection Ports. For Virtual Circuits on Fabric VCs (both Metal and Fabric Billed) that have found their corresponding Fabric connection, this is the actual speed of the interconnection that was configured when setting up the interconnection on the Fabric Portal. Details on Fabric VCs are included in the specification as a developer preview and is generally unavailable. Please contact our Support team for more details. | [optional] 
**Status** | **string** | The status of a Virtual Circuit is always &#39;Pending&#39; on creation. The status can turn to &#39;Waiting on Customer VLAN&#39; if a Metro VLAN was not set yet on the Virtual Circuit and is the last step needed for full activation. For Dedicated interconnections, as long as the Dedicated Port has been associated to the Virtual Circuit and a NNI VNID has been set, it will turn to &#39;Waiting on Customer VLAN&#39;. For Fabric VCs, it will only change to &#39;Waiting on Customer VLAN&#39; once the corresponding Fabric connection has been found on the Fabric side. Once a Metro VLAN is set on the Virtual Circuit (which for Fabric VCs, can be set on creation) and the necessary set up is done, it will turn into &#39;Activating&#39; status as it tries to activate the Virtual Circuit. Once the Virtual Circuit fully activates and is configured on the switch, it will turn to staus &#39;Active&#39;. For Fabric VCs (Metal Billed), we will start billing the moment the status of the Virtual Circuit turns to &#39;Active&#39;. If there are any changes to the VLAN after the Virtual Circuit is in an &#39;Active&#39; status, the status will show &#39;Changing VLAN&#39; if a new VLAN has been provided, or &#39;Deactivating&#39; if we are removing the VLAN. When a deletion request is issued for the Virtual Circuit, it will move to a &#39;deleting&#39; status until it is fully deleted. If the Virtual Circuit is on a Fabric VC, it can also change into an &#39;Expired&#39; status if the associated service token has expired. | 
**Tags** | **[]string** |  | 
**VirtualNetwork** | [**FindBatchById200ResponseDevicesInner**](FindBatchById200ResponseDevicesInner.md) |  | 
**Vnid** | **int32** |  | 

## Methods

### NewGetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf

`func NewGetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf(bill bool, description string, id string, name string, nniVlan int32, port FindBatchById200ResponseDevicesInner, project FindBatchById200ResponseDevicesInner, status string, tags []string, virtualNetwork FindBatchById200ResponseDevicesInner, vnid int32, ) *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf`

NewGetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf instantiates a new GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOfWithDefaults

`func NewGetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOfWithDefaults() *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf`

NewGetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOfWithDefaults instantiates a new GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBill

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf) GetBill() bool`

GetBill returns the Bill field if non-nil, zero value otherwise.

### GetBillOk

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf) GetBillOk() (*bool, bool)`

GetBillOk returns a tuple with the Bill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBill

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf) SetBill(v bool)`

SetBill sets Bill field to given value.


### GetDescription

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetId

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf) SetName(v string)`

SetName sets Name field to given value.


### GetNniVlan

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf) GetNniVlan() int32`

GetNniVlan returns the NniVlan field if non-nil, zero value otherwise.

### GetNniVlanOk

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf) GetNniVlanOk() (*int32, bool)`

GetNniVlanOk returns a tuple with the NniVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNniVlan

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf) SetNniVlan(v int32)`

SetNniVlan sets NniVlan field to given value.


### GetPort

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf) GetPort() FindBatchById200ResponseDevicesInner`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf) GetPortOk() (*FindBatchById200ResponseDevicesInner, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf) SetPort(v FindBatchById200ResponseDevicesInner)`

SetPort sets Port field to given value.


### GetProject

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf) GetProject() FindBatchById200ResponseDevicesInner`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf) GetProjectOk() (*FindBatchById200ResponseDevicesInner, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf) SetProject(v FindBatchById200ResponseDevicesInner)`

SetProject sets Project field to given value.


### GetSpeed

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf) GetSpeed() int32`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf) GetSpeedOk() (*int32, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf) SetSpeed(v int32)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetStatus

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTags

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetVirtualNetwork

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf) GetVirtualNetwork() FindBatchById200ResponseDevicesInner`

GetVirtualNetwork returns the VirtualNetwork field if non-nil, zero value otherwise.

### GetVirtualNetworkOk

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf) GetVirtualNetworkOk() (*FindBatchById200ResponseDevicesInner, bool)`

GetVirtualNetworkOk returns a tuple with the VirtualNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualNetwork

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf) SetVirtualNetwork(v FindBatchById200ResponseDevicesInner)`

SetVirtualNetwork sets VirtualNetwork field to given value.


### GetVnid

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf) GetVnid() int32`

GetVnid returns the Vnid field if non-nil, zero value otherwise.

### GetVnidOk

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf) GetVnidOk() (*int32, bool)`

GetVnidOk returns a tuple with the Vnid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnid

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf) SetVnid(v int32)`

SetVnid sets Vnid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


