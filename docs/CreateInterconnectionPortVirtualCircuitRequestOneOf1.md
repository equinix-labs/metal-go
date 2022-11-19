# CreateInterconnectionPortVirtualCircuitRequestOneOf1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerIp** | Pointer to **string** | An IP address from the subnet that will be used on the Customer side. This parameter is optional, but if supplied, we will use the other usable IP address in the subnet as the Metal IP. By default, the last usable IP address in the subnet will be used. | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Md5** | Pointer to **NullableString** | The MD5 password for the BGP peering in plaintext (not a checksum). | [optional] 
**MetalIp** | Pointer to **string** | An IP address from the subnet that will be used on the Metal side. This parameter is optional, but if supplied, we will use the other usable IP address in the subnet as the Customer IP. By default, the first usable IP address in the subnet will be used. | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NniVlan** | **int32** |  | 
**PeerAsn** | **int32** | The peer ASN that will be used with the VRF on the Virtual Circuit. | 
**Project** | **string** |  | 
**Speed** | Pointer to **int32** | speed can be passed as integer number representing bps speed or string (e.g. &#39;52m&#39; or &#39;100g&#39; or &#39;4 gbps&#39;) | [optional] 
**Subnet** | **string** | The /30 or /31 subnet of one of the VRF IP Blocks that will be used with the VRF for the Virtual Circuit. This subnet does not have to be an existing VRF IP reservation, as we will create the VRF IP reservation on creation if it does not exist. The Metal IP and Customer IP must be IPs from this subnet. For /30 subnets, the network and broadcast IPs cannot be used as the Metal or Customer IP. The subnet specified must be contained within an already-defined IP Range for the VRF. | 
**Tags** | Pointer to **[]string** |  | [optional] 
**Vrf** | **string** | The UUID of the VRF that will be associated with the Virtual Circuit. | 

## Methods

### NewCreateInterconnectionPortVirtualCircuitRequestOneOf1

`func NewCreateInterconnectionPortVirtualCircuitRequestOneOf1(nniVlan int32, peerAsn int32, project string, subnet string, vrf string, ) *CreateInterconnectionPortVirtualCircuitRequestOneOf1`

NewCreateInterconnectionPortVirtualCircuitRequestOneOf1 instantiates a new CreateInterconnectionPortVirtualCircuitRequestOneOf1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateInterconnectionPortVirtualCircuitRequestOneOf1WithDefaults

`func NewCreateInterconnectionPortVirtualCircuitRequestOneOf1WithDefaults() *CreateInterconnectionPortVirtualCircuitRequestOneOf1`

NewCreateInterconnectionPortVirtualCircuitRequestOneOf1WithDefaults instantiates a new CreateInterconnectionPortVirtualCircuitRequestOneOf1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerIp

`func (o *CreateInterconnectionPortVirtualCircuitRequestOneOf1) GetCustomerIp() string`

GetCustomerIp returns the CustomerIp field if non-nil, zero value otherwise.

### GetCustomerIpOk

`func (o *CreateInterconnectionPortVirtualCircuitRequestOneOf1) GetCustomerIpOk() (*string, bool)`

GetCustomerIpOk returns a tuple with the CustomerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerIp

`func (o *CreateInterconnectionPortVirtualCircuitRequestOneOf1) SetCustomerIp(v string)`

SetCustomerIp sets CustomerIp field to given value.

### HasCustomerIp

`func (o *CreateInterconnectionPortVirtualCircuitRequestOneOf1) HasCustomerIp() bool`

HasCustomerIp returns a boolean if a field has been set.

### GetDescription

`func (o *CreateInterconnectionPortVirtualCircuitRequestOneOf1) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateInterconnectionPortVirtualCircuitRequestOneOf1) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateInterconnectionPortVirtualCircuitRequestOneOf1) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateInterconnectionPortVirtualCircuitRequestOneOf1) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMd5

`func (o *CreateInterconnectionPortVirtualCircuitRequestOneOf1) GetMd5() string`

GetMd5 returns the Md5 field if non-nil, zero value otherwise.

### GetMd5Ok

`func (o *CreateInterconnectionPortVirtualCircuitRequestOneOf1) GetMd5Ok() (*string, bool)`

GetMd5Ok returns a tuple with the Md5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5

`func (o *CreateInterconnectionPortVirtualCircuitRequestOneOf1) SetMd5(v string)`

SetMd5 sets Md5 field to given value.

### HasMd5

`func (o *CreateInterconnectionPortVirtualCircuitRequestOneOf1) HasMd5() bool`

HasMd5 returns a boolean if a field has been set.

### SetMd5Nil

`func (o *CreateInterconnectionPortVirtualCircuitRequestOneOf1) SetMd5Nil(b bool)`

 SetMd5Nil sets the value for Md5 to be an explicit nil

### UnsetMd5
`func (o *CreateInterconnectionPortVirtualCircuitRequestOneOf1) UnsetMd5()`

UnsetMd5 ensures that no value is present for Md5, not even an explicit nil
### GetMetalIp

`func (o *CreateInterconnectionPortVirtualCircuitRequestOneOf1) GetMetalIp() string`

GetMetalIp returns the MetalIp field if non-nil, zero value otherwise.

### GetMetalIpOk

`func (o *CreateInterconnectionPortVirtualCircuitRequestOneOf1) GetMetalIpOk() (*string, bool)`

GetMetalIpOk returns a tuple with the MetalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetalIp

`func (o *CreateInterconnectionPortVirtualCircuitRequestOneOf1) SetMetalIp(v string)`

SetMetalIp sets MetalIp field to given value.

### HasMetalIp

`func (o *CreateInterconnectionPortVirtualCircuitRequestOneOf1) HasMetalIp() bool`

HasMetalIp returns a boolean if a field has been set.

### GetName

`func (o *CreateInterconnectionPortVirtualCircuitRequestOneOf1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateInterconnectionPortVirtualCircuitRequestOneOf1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateInterconnectionPortVirtualCircuitRequestOneOf1) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateInterconnectionPortVirtualCircuitRequestOneOf1) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNniVlan

`func (o *CreateInterconnectionPortVirtualCircuitRequestOneOf1) GetNniVlan() int32`

GetNniVlan returns the NniVlan field if non-nil, zero value otherwise.

### GetNniVlanOk

`func (o *CreateInterconnectionPortVirtualCircuitRequestOneOf1) GetNniVlanOk() (*int32, bool)`

GetNniVlanOk returns a tuple with the NniVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNniVlan

`func (o *CreateInterconnectionPortVirtualCircuitRequestOneOf1) SetNniVlan(v int32)`

SetNniVlan sets NniVlan field to given value.


### GetPeerAsn

`func (o *CreateInterconnectionPortVirtualCircuitRequestOneOf1) GetPeerAsn() int32`

GetPeerAsn returns the PeerAsn field if non-nil, zero value otherwise.

### GetPeerAsnOk

`func (o *CreateInterconnectionPortVirtualCircuitRequestOneOf1) GetPeerAsnOk() (*int32, bool)`

GetPeerAsnOk returns a tuple with the PeerAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerAsn

`func (o *CreateInterconnectionPortVirtualCircuitRequestOneOf1) SetPeerAsn(v int32)`

SetPeerAsn sets PeerAsn field to given value.


### GetProject

`func (o *CreateInterconnectionPortVirtualCircuitRequestOneOf1) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *CreateInterconnectionPortVirtualCircuitRequestOneOf1) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *CreateInterconnectionPortVirtualCircuitRequestOneOf1) SetProject(v string)`

SetProject sets Project field to given value.


### GetSpeed

`func (o *CreateInterconnectionPortVirtualCircuitRequestOneOf1) GetSpeed() int32`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *CreateInterconnectionPortVirtualCircuitRequestOneOf1) GetSpeedOk() (*int32, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *CreateInterconnectionPortVirtualCircuitRequestOneOf1) SetSpeed(v int32)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *CreateInterconnectionPortVirtualCircuitRequestOneOf1) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetSubnet

`func (o *CreateInterconnectionPortVirtualCircuitRequestOneOf1) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *CreateInterconnectionPortVirtualCircuitRequestOneOf1) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *CreateInterconnectionPortVirtualCircuitRequestOneOf1) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.


### GetTags

`func (o *CreateInterconnectionPortVirtualCircuitRequestOneOf1) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateInterconnectionPortVirtualCircuitRequestOneOf1) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateInterconnectionPortVirtualCircuitRequestOneOf1) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateInterconnectionPortVirtualCircuitRequestOneOf1) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVrf

`func (o *CreateInterconnectionPortVirtualCircuitRequestOneOf1) GetVrf() string`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *CreateInterconnectionPortVirtualCircuitRequestOneOf1) GetVrfOk() (*string, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *CreateInterconnectionPortVirtualCircuitRequestOneOf1) SetVrf(v string)`

SetVrf sets Vrf field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

