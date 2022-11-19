# GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInner

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
**Speed** | Pointer to **int32** | integer representing bps speed | [optional] 
**Status** | **string** |  | 
**Tags** | **[]string** |  | 
**VirtualNetwork** | [**FindBatchById200ResponseDevicesInner**](FindBatchById200ResponseDevicesInner.md) |  | 
**Vnid** | **int32** |  | 
**CustomerIp** | Pointer to **string** | An IP address from the subnet that will be used on the Customer side. This parameter is optional, but if supplied, we will use the other usable IP address in the subnet as the Metal IP. By default, the last usable IP address in the subnet will be used. | [optional] 
**Md5** | Pointer to **string** | The MD5 password for the BGP peering in plaintext (not a checksum). | [optional] 
**MetalIp** | Pointer to **string** | An IP address from the subnet that will be used on the Metal side. This parameter is optional, but if supplied, we will use the other usable IP address in the subnet as the Customer IP. By default, the first usable IP address in the subnet will be used. | [optional] 
**PeerAsn** | Pointer to **int32** | The peer ASN that will be used with the VRF on the Virtual Circuit. | [optional] 
**Subnet** | Pointer to **string** | The /30 or /31 subnet of one of the VRF IP Blocks that will be used with the VRF for the Virtual Circuit. This subnet does not have to be an existing VRF IP reservation, as we will create the VRF IP reservation on creation if it does not exist. The Metal IP and Customer IP must be IPs from this subnet. For /30 subnets, the network and broadcast IPs cannot be used as the Metal or Customer IP. | [optional] 
**Vrf** | Pointer to [**GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1Vrf**](GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1Vrf.md) |  | [optional] 

## Methods

### NewGetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInner

`func NewGetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInner(bill bool, description string, id string, name string, nniVlan int32, port FindBatchById200ResponseDevicesInner, project FindBatchById200ResponseDevicesInner, status string, tags []string, virtualNetwork FindBatchById200ResponseDevicesInner, vnid int32, ) *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInner`

NewGetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInner instantiates a new GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerWithDefaults

`func NewGetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerWithDefaults() *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInner`

NewGetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerWithDefaults instantiates a new GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBill

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInner) GetBill() bool`

GetBill returns the Bill field if non-nil, zero value otherwise.

### GetBillOk

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInner) GetBillOk() (*bool, bool)`

GetBillOk returns a tuple with the Bill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBill

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInner) SetBill(v bool)`

SetBill sets Bill field to given value.


### GetDescription

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInner) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetId

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInner) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInner) SetName(v string)`

SetName sets Name field to given value.


### GetNniVlan

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInner) GetNniVlan() int32`

GetNniVlan returns the NniVlan field if non-nil, zero value otherwise.

### GetNniVlanOk

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInner) GetNniVlanOk() (*int32, bool)`

GetNniVlanOk returns a tuple with the NniVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNniVlan

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInner) SetNniVlan(v int32)`

SetNniVlan sets NniVlan field to given value.


### GetPort

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInner) GetPort() FindBatchById200ResponseDevicesInner`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInner) GetPortOk() (*FindBatchById200ResponseDevicesInner, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInner) SetPort(v FindBatchById200ResponseDevicesInner)`

SetPort sets Port field to given value.


### GetProject

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInner) GetProject() FindBatchById200ResponseDevicesInner`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInner) GetProjectOk() (*FindBatchById200ResponseDevicesInner, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInner) SetProject(v FindBatchById200ResponseDevicesInner)`

SetProject sets Project field to given value.


### GetSpeed

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInner) GetSpeed() int32`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInner) GetSpeedOk() (*int32, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInner) SetSpeed(v int32)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInner) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetStatus

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInner) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTags

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInner) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInner) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInner) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetVirtualNetwork

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInner) GetVirtualNetwork() FindBatchById200ResponseDevicesInner`

GetVirtualNetwork returns the VirtualNetwork field if non-nil, zero value otherwise.

### GetVirtualNetworkOk

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInner) GetVirtualNetworkOk() (*FindBatchById200ResponseDevicesInner, bool)`

GetVirtualNetworkOk returns a tuple with the VirtualNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualNetwork

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInner) SetVirtualNetwork(v FindBatchById200ResponseDevicesInner)`

SetVirtualNetwork sets VirtualNetwork field to given value.


### GetVnid

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInner) GetVnid() int32`

GetVnid returns the Vnid field if non-nil, zero value otherwise.

### GetVnidOk

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInner) GetVnidOk() (*int32, bool)`

GetVnidOk returns a tuple with the Vnid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnid

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInner) SetVnid(v int32)`

SetVnid sets Vnid field to given value.


### GetCustomerIp

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInner) GetCustomerIp() string`

GetCustomerIp returns the CustomerIp field if non-nil, zero value otherwise.

### GetCustomerIpOk

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInner) GetCustomerIpOk() (*string, bool)`

GetCustomerIpOk returns a tuple with the CustomerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerIp

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInner) SetCustomerIp(v string)`

SetCustomerIp sets CustomerIp field to given value.

### HasCustomerIp

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInner) HasCustomerIp() bool`

HasCustomerIp returns a boolean if a field has been set.

### GetMd5

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInner) GetMd5() string`

GetMd5 returns the Md5 field if non-nil, zero value otherwise.

### GetMd5Ok

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInner) GetMd5Ok() (*string, bool)`

GetMd5Ok returns a tuple with the Md5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInner) SetMd5(v string)`

SetMd5 sets Md5 field to given value.

### HasMd5

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInner) HasMd5() bool`

HasMd5 returns a boolean if a field has been set.

### GetMetalIp

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInner) GetMetalIp() string`

GetMetalIp returns the MetalIp field if non-nil, zero value otherwise.

### GetMetalIpOk

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInner) GetMetalIpOk() (*string, bool)`

GetMetalIpOk returns a tuple with the MetalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetalIp

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInner) SetMetalIp(v string)`

SetMetalIp sets MetalIp field to given value.

### HasMetalIp

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInner) HasMetalIp() bool`

HasMetalIp returns a boolean if a field has been set.

### GetPeerAsn

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInner) GetPeerAsn() int32`

GetPeerAsn returns the PeerAsn field if non-nil, zero value otherwise.

### GetPeerAsnOk

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInner) GetPeerAsnOk() (*int32, bool)`

GetPeerAsnOk returns a tuple with the PeerAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerAsn

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInner) SetPeerAsn(v int32)`

SetPeerAsn sets PeerAsn field to given value.

### HasPeerAsn

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInner) HasPeerAsn() bool`

HasPeerAsn returns a boolean if a field has been set.

### GetSubnet

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInner) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInner) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInner) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInner) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetVrf

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInner) GetVrf() GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1Vrf`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInner) GetVrfOk() (*GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1Vrf, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInner) SetVrf(v GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1Vrf)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInner) HasVrf() bool`

HasVrf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


