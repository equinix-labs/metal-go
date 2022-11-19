# GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerIp** | Pointer to **string** | An IP address from the subnet that will be used on the Customer side. This parameter is optional, but if supplied, we will use the other usable IP address in the subnet as the Metal IP. By default, the last usable IP address in the subnet will be used. | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Md5** | Pointer to **string** | The MD5 password for the BGP peering in plaintext (not a checksum). | [optional] 
**MetalIp** | Pointer to **string** | An IP address from the subnet that will be used on the Metal side. This parameter is optional, but if supplied, we will use the other usable IP address in the subnet as the Customer IP. By default, the first usable IP address in the subnet will be used. | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Port** | Pointer to [**FindBatchById200ResponseDevicesInner**](FindBatchById200ResponseDevicesInner.md) |  | [optional] 
**NniVlan** | Pointer to **int32** |  | [optional] 
**PeerAsn** | Pointer to **int32** | The peer ASN that will be used with the VRF on the Virtual Circuit. | [optional] 
**Project** | Pointer to [**FindBatchById200ResponseDevicesInner**](FindBatchById200ResponseDevicesInner.md) |  | [optional] 
**Speed** | Pointer to **int32** | integer representing bps speed | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Subnet** | Pointer to **string** | The /30 or /31 subnet of one of the VRF IP Blocks that will be used with the VRF for the Virtual Circuit. This subnet does not have to be an existing VRF IP reservation, as we will create the VRF IP reservation on creation if it does not exist. The Metal IP and Customer IP must be IPs from this subnet. For /30 subnets, the network and broadcast IPs cannot be used as the Metal or Customer IP. | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Vrf** | Pointer to [**GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1Vrf**](GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1Vrf.md) |  | [optional] 

## Methods

### NewGetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1

`func NewGetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1() *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1`

NewGetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1 instantiates a new GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1WithDefaults

`func NewGetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1WithDefaults() *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1`

NewGetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1WithDefaults instantiates a new GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerIp

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1) GetCustomerIp() string`

GetCustomerIp returns the CustomerIp field if non-nil, zero value otherwise.

### GetCustomerIpOk

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1) GetCustomerIpOk() (*string, bool)`

GetCustomerIpOk returns a tuple with the CustomerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerIp

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1) SetCustomerIp(v string)`

SetCustomerIp sets CustomerIp field to given value.

### HasCustomerIp

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1) HasCustomerIp() bool`

HasCustomerIp returns a boolean if a field has been set.

### GetDescription

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMd5

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1) GetMd5() string`

GetMd5 returns the Md5 field if non-nil, zero value otherwise.

### GetMd5Ok

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1) GetMd5Ok() (*string, bool)`

GetMd5Ok returns a tuple with the Md5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1) SetMd5(v string)`

SetMd5 sets Md5 field to given value.

### HasMd5

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1) HasMd5() bool`

HasMd5 returns a boolean if a field has been set.

### GetMetalIp

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1) GetMetalIp() string`

GetMetalIp returns the MetalIp field if non-nil, zero value otherwise.

### GetMetalIpOk

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1) GetMetalIpOk() (*string, bool)`

GetMetalIpOk returns a tuple with the MetalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetalIp

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1) SetMetalIp(v string)`

SetMetalIp sets MetalIp field to given value.

### HasMetalIp

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1) HasMetalIp() bool`

HasMetalIp returns a boolean if a field has been set.

### GetName

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPort

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1) GetPort() FindBatchById200ResponseDevicesInner`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1) GetPortOk() (*FindBatchById200ResponseDevicesInner, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1) SetPort(v FindBatchById200ResponseDevicesInner)`

SetPort sets Port field to given value.

### HasPort

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetNniVlan

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1) GetNniVlan() int32`

GetNniVlan returns the NniVlan field if non-nil, zero value otherwise.

### GetNniVlanOk

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1) GetNniVlanOk() (*int32, bool)`

GetNniVlanOk returns a tuple with the NniVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNniVlan

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1) SetNniVlan(v int32)`

SetNniVlan sets NniVlan field to given value.

### HasNniVlan

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1) HasNniVlan() bool`

HasNniVlan returns a boolean if a field has been set.

### GetPeerAsn

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1) GetPeerAsn() int32`

GetPeerAsn returns the PeerAsn field if non-nil, zero value otherwise.

### GetPeerAsnOk

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1) GetPeerAsnOk() (*int32, bool)`

GetPeerAsnOk returns a tuple with the PeerAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerAsn

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1) SetPeerAsn(v int32)`

SetPeerAsn sets PeerAsn field to given value.

### HasPeerAsn

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1) HasPeerAsn() bool`

HasPeerAsn returns a boolean if a field has been set.

### GetProject

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1) GetProject() FindBatchById200ResponseDevicesInner`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1) GetProjectOk() (*FindBatchById200ResponseDevicesInner, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1) SetProject(v FindBatchById200ResponseDevicesInner)`

SetProject sets Project field to given value.

### HasProject

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetSpeed

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1) GetSpeed() int32`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1) GetSpeedOk() (*int32, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1) SetSpeed(v int32)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetStatus

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubnet

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetTags

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVrf

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1) GetVrf() GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1Vrf`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1) GetVrfOk() (*GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1Vrf, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1) SetVrf(v GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1Vrf)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1) HasVrf() bool`

HasVrf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


