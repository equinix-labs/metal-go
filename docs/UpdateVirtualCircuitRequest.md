# UpdateVirtualCircuitRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Speed** | Pointer to **string** | Speed can be changed only if it is an interconnection on a Dedicated Port | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Vnid** | Pointer to **string** | A Virtual Network record UUID or the VNID of a Metro Virtual Network in your project. | [optional] 
**CustomerIp** | Pointer to **string** | An IP address from the subnet that will be used on the Customer side. This parameter is optional, but if supplied, we will use the other usable IP address in the subnet as the Metal IP. By default, the last usable IP address in the subnet will be used. | [optional] 
**Md5** | Pointer to **string** | The MD5 password for the BGP peering in plaintext (not a checksum). | [optional] 
**MetalIp** | Pointer to **string** | An IP address from the subnet that will be used on the Metal side. This parameter is optional, but if supplied, we will use the other usable IP address in the subnet as the Customer IP. By default, the first usable IP address in the subnet will be used. | [optional] 
**PeerAsn** | Pointer to **int32** | The peer ASN that will be used with the VRF on the Virtual Circuit. | [optional] 
**Subnet** | Pointer to **string** | The /30 or /31 subnet of one of the VRF IP Blocks that will be used with the VRF for the Virtual Circuit. This subnet does not have to be an existing VRF IP reservation, as we will create the VRF IP reservation on creation if it does not exist. The Metal IP and Customer IP must be IPs from this subnet. For /30 subnets, the network and broadcast IPs cannot be used as the Metal or Customer IP. | [optional] 

## Methods

### NewUpdateVirtualCircuitRequest

`func NewUpdateVirtualCircuitRequest() *UpdateVirtualCircuitRequest`

NewUpdateVirtualCircuitRequest instantiates a new UpdateVirtualCircuitRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateVirtualCircuitRequestWithDefaults

`func NewUpdateVirtualCircuitRequestWithDefaults() *UpdateVirtualCircuitRequest`

NewUpdateVirtualCircuitRequestWithDefaults instantiates a new UpdateVirtualCircuitRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UpdateVirtualCircuitRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateVirtualCircuitRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateVirtualCircuitRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateVirtualCircuitRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *UpdateVirtualCircuitRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateVirtualCircuitRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateVirtualCircuitRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateVirtualCircuitRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSpeed

`func (o *UpdateVirtualCircuitRequest) GetSpeed() string`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *UpdateVirtualCircuitRequest) GetSpeedOk() (*string, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *UpdateVirtualCircuitRequest) SetSpeed(v string)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *UpdateVirtualCircuitRequest) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetTags

`func (o *UpdateVirtualCircuitRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateVirtualCircuitRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateVirtualCircuitRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateVirtualCircuitRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVnid

`func (o *UpdateVirtualCircuitRequest) GetVnid() string`

GetVnid returns the Vnid field if non-nil, zero value otherwise.

### GetVnidOk

`func (o *UpdateVirtualCircuitRequest) GetVnidOk() (*string, bool)`

GetVnidOk returns a tuple with the Vnid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnid

`func (o *UpdateVirtualCircuitRequest) SetVnid(v string)`

SetVnid sets Vnid field to given value.

### HasVnid

`func (o *UpdateVirtualCircuitRequest) HasVnid() bool`

HasVnid returns a boolean if a field has been set.

### GetCustomerIp

`func (o *UpdateVirtualCircuitRequest) GetCustomerIp() string`

GetCustomerIp returns the CustomerIp field if non-nil, zero value otherwise.

### GetCustomerIpOk

`func (o *UpdateVirtualCircuitRequest) GetCustomerIpOk() (*string, bool)`

GetCustomerIpOk returns a tuple with the CustomerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerIp

`func (o *UpdateVirtualCircuitRequest) SetCustomerIp(v string)`

SetCustomerIp sets CustomerIp field to given value.

### HasCustomerIp

`func (o *UpdateVirtualCircuitRequest) HasCustomerIp() bool`

HasCustomerIp returns a boolean if a field has been set.

### GetMd5

`func (o *UpdateVirtualCircuitRequest) GetMd5() string`

GetMd5 returns the Md5 field if non-nil, zero value otherwise.

### GetMd5Ok

`func (o *UpdateVirtualCircuitRequest) GetMd5Ok() (*string, bool)`

GetMd5Ok returns a tuple with the Md5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5

`func (o *UpdateVirtualCircuitRequest) SetMd5(v string)`

SetMd5 sets Md5 field to given value.

### HasMd5

`func (o *UpdateVirtualCircuitRequest) HasMd5() bool`

HasMd5 returns a boolean if a field has been set.

### GetMetalIp

`func (o *UpdateVirtualCircuitRequest) GetMetalIp() string`

GetMetalIp returns the MetalIp field if non-nil, zero value otherwise.

### GetMetalIpOk

`func (o *UpdateVirtualCircuitRequest) GetMetalIpOk() (*string, bool)`

GetMetalIpOk returns a tuple with the MetalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetalIp

`func (o *UpdateVirtualCircuitRequest) SetMetalIp(v string)`

SetMetalIp sets MetalIp field to given value.

### HasMetalIp

`func (o *UpdateVirtualCircuitRequest) HasMetalIp() bool`

HasMetalIp returns a boolean if a field has been set.

### GetPeerAsn

`func (o *UpdateVirtualCircuitRequest) GetPeerAsn() int32`

GetPeerAsn returns the PeerAsn field if non-nil, zero value otherwise.

### GetPeerAsnOk

`func (o *UpdateVirtualCircuitRequest) GetPeerAsnOk() (*int32, bool)`

GetPeerAsnOk returns a tuple with the PeerAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerAsn

`func (o *UpdateVirtualCircuitRequest) SetPeerAsn(v int32)`

SetPeerAsn sets PeerAsn field to given value.

### HasPeerAsn

`func (o *UpdateVirtualCircuitRequest) HasPeerAsn() bool`

HasPeerAsn returns a boolean if a field has been set.

### GetSubnet

`func (o *UpdateVirtualCircuitRequest) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *UpdateVirtualCircuitRequest) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *UpdateVirtualCircuitRequest) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *UpdateVirtualCircuitRequest) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


