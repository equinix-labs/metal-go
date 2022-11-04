# CreateInterconnectionPortVirtualCircuitRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NniVlan** | **int32** |  | 
**Project** | **string** |  | 
**Speed** | Pointer to **int32** | speed can be passed as integer number representing bps speed or string (e.g. &#39;52m&#39; or &#39;100g&#39; or &#39;4 gbps&#39;) | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Vnid** | Pointer to **string** | A Virtual Network record UUID or the VNID of a Virtual Network in your project (sent as integer). | [optional] 
**CustomerIp** | Pointer to **string** | An IP address from the subnet that will be used on the Customer side. This parameter is optional, but if supplied, we will use the other usable IP address in the subnet as the Metal IP. By default, the last usable IP address in the subnet will be used. | [optional] 
**Md5** | Pointer to **NullableString** | The MD5 password for the BGP peering in plaintext (not a checksum). | [optional] 
**MetalIp** | Pointer to **string** | An IP address from the subnet that will be used on the Metal side. This parameter is optional, but if supplied, we will use the other usable IP address in the subnet as the Customer IP. By default, the first usable IP address in the subnet will be used. | [optional] 
**PeerAsn** | **int32** | The peer ASN that will be used with the VRF on the Virtual Circuit. | 
**Subnet** | **string** | The /30 or /31 subnet of one of the VRF IP Blocks that will be used with the VRF for the Virtual Circuit. This subnet does not have to be an existing VRF IP reservation, as we will create the VRF IP reservation on creation if it does not exist. The Metal IP and Customer IP must be IPs from this subnet. For /30 subnets, the network and broadcast IPs cannot be used as the Metal or Customer IP. The subnet specified must be contained within an already-defined IP Range for the VRF. | 
**Vrf** | **string** | The UUID of the VRF that will be associated with the Virtual Circuit. | 

## Methods

### NewCreateInterconnectionPortVirtualCircuitRequest

`func NewCreateInterconnectionPortVirtualCircuitRequest(nniVlan int32, project string, peerAsn int32, subnet string, vrf string, ) *CreateInterconnectionPortVirtualCircuitRequest`

NewCreateInterconnectionPortVirtualCircuitRequest instantiates a new CreateInterconnectionPortVirtualCircuitRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateInterconnectionPortVirtualCircuitRequestWithDefaults

`func NewCreateInterconnectionPortVirtualCircuitRequestWithDefaults() *CreateInterconnectionPortVirtualCircuitRequest`

NewCreateInterconnectionPortVirtualCircuitRequestWithDefaults instantiates a new CreateInterconnectionPortVirtualCircuitRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateInterconnectionPortVirtualCircuitRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateInterconnectionPortVirtualCircuitRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateInterconnectionPortVirtualCircuitRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateInterconnectionPortVirtualCircuitRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *CreateInterconnectionPortVirtualCircuitRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateInterconnectionPortVirtualCircuitRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateInterconnectionPortVirtualCircuitRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateInterconnectionPortVirtualCircuitRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNniVlan

`func (o *CreateInterconnectionPortVirtualCircuitRequest) GetNniVlan() int32`

GetNniVlan returns the NniVlan field if non-nil, zero value otherwise.

### GetNniVlanOk

`func (o *CreateInterconnectionPortVirtualCircuitRequest) GetNniVlanOk() (*int32, bool)`

GetNniVlanOk returns a tuple with the NniVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNniVlan

`func (o *CreateInterconnectionPortVirtualCircuitRequest) SetNniVlan(v int32)`

SetNniVlan sets NniVlan field to given value.


### GetProject

`func (o *CreateInterconnectionPortVirtualCircuitRequest) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *CreateInterconnectionPortVirtualCircuitRequest) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *CreateInterconnectionPortVirtualCircuitRequest) SetProject(v string)`

SetProject sets Project field to given value.


### GetSpeed

`func (o *CreateInterconnectionPortVirtualCircuitRequest) GetSpeed() int32`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *CreateInterconnectionPortVirtualCircuitRequest) GetSpeedOk() (*int32, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *CreateInterconnectionPortVirtualCircuitRequest) SetSpeed(v int32)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *CreateInterconnectionPortVirtualCircuitRequest) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetTags

`func (o *CreateInterconnectionPortVirtualCircuitRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateInterconnectionPortVirtualCircuitRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateInterconnectionPortVirtualCircuitRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateInterconnectionPortVirtualCircuitRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVnid

`func (o *CreateInterconnectionPortVirtualCircuitRequest) GetVnid() string`

GetVnid returns the Vnid field if non-nil, zero value otherwise.

### GetVnidOk

`func (o *CreateInterconnectionPortVirtualCircuitRequest) GetVnidOk() (*string, bool)`

GetVnidOk returns a tuple with the Vnid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnid

`func (o *CreateInterconnectionPortVirtualCircuitRequest) SetVnid(v string)`

SetVnid sets Vnid field to given value.

### HasVnid

`func (o *CreateInterconnectionPortVirtualCircuitRequest) HasVnid() bool`

HasVnid returns a boolean if a field has been set.

### GetCustomerIp

`func (o *CreateInterconnectionPortVirtualCircuitRequest) GetCustomerIp() string`

GetCustomerIp returns the CustomerIp field if non-nil, zero value otherwise.

### GetCustomerIpOk

`func (o *CreateInterconnectionPortVirtualCircuitRequest) GetCustomerIpOk() (*string, bool)`

GetCustomerIpOk returns a tuple with the CustomerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerIp

`func (o *CreateInterconnectionPortVirtualCircuitRequest) SetCustomerIp(v string)`

SetCustomerIp sets CustomerIp field to given value.

### HasCustomerIp

`func (o *CreateInterconnectionPortVirtualCircuitRequest) HasCustomerIp() bool`

HasCustomerIp returns a boolean if a field has been set.

### GetMd5

`func (o *CreateInterconnectionPortVirtualCircuitRequest) GetMd5() string`

GetMd5 returns the Md5 field if non-nil, zero value otherwise.

### GetMd5Ok

`func (o *CreateInterconnectionPortVirtualCircuitRequest) GetMd5Ok() (*string, bool)`

GetMd5Ok returns a tuple with the Md5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5

`func (o *CreateInterconnectionPortVirtualCircuitRequest) SetMd5(v string)`

SetMd5 sets Md5 field to given value.

### HasMd5

`func (o *CreateInterconnectionPortVirtualCircuitRequest) HasMd5() bool`

HasMd5 returns a boolean if a field has been set.

### SetMd5Nil

`func (o *CreateInterconnectionPortVirtualCircuitRequest) SetMd5Nil(b bool)`

 SetMd5Nil sets the value for Md5 to be an explicit nil

### UnsetMd5
`func (o *CreateInterconnectionPortVirtualCircuitRequest) UnsetMd5()`

UnsetMd5 ensures that no value is present for Md5, not even an explicit nil
### GetMetalIp

`func (o *CreateInterconnectionPortVirtualCircuitRequest) GetMetalIp() string`

GetMetalIp returns the MetalIp field if non-nil, zero value otherwise.

### GetMetalIpOk

`func (o *CreateInterconnectionPortVirtualCircuitRequest) GetMetalIpOk() (*string, bool)`

GetMetalIpOk returns a tuple with the MetalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetalIp

`func (o *CreateInterconnectionPortVirtualCircuitRequest) SetMetalIp(v string)`

SetMetalIp sets MetalIp field to given value.

### HasMetalIp

`func (o *CreateInterconnectionPortVirtualCircuitRequest) HasMetalIp() bool`

HasMetalIp returns a boolean if a field has been set.

### GetPeerAsn

`func (o *CreateInterconnectionPortVirtualCircuitRequest) GetPeerAsn() int32`

GetPeerAsn returns the PeerAsn field if non-nil, zero value otherwise.

### GetPeerAsnOk

`func (o *CreateInterconnectionPortVirtualCircuitRequest) GetPeerAsnOk() (*int32, bool)`

GetPeerAsnOk returns a tuple with the PeerAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerAsn

`func (o *CreateInterconnectionPortVirtualCircuitRequest) SetPeerAsn(v int32)`

SetPeerAsn sets PeerAsn field to given value.


### GetSubnet

`func (o *CreateInterconnectionPortVirtualCircuitRequest) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *CreateInterconnectionPortVirtualCircuitRequest) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *CreateInterconnectionPortVirtualCircuitRequest) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.


### GetVrf

`func (o *CreateInterconnectionPortVirtualCircuitRequest) GetVrf() string`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *CreateInterconnectionPortVirtualCircuitRequest) GetVrfOk() (*string, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *CreateInterconnectionPortVirtualCircuitRequest) SetVrf(v string)`

SetVrf sets Vrf field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


