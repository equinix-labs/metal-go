# VrfVirtualCircuitsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Md5** | Pointer to **string** | The MD5 password for the BGP peering in plaintext (not a checksum). | [optional] 
**PeerAsn** | Pointer to **int32** | The peer ASN that will be used with the VRF on the Virtual Circuit. | [optional] 
**Subnet** | Pointer to **string** | The /30 or /31 subnet of one of the VRF IP Blocks that will be used with the VRF for the Virtual Circuit. This subnet does not have to be an existing VRF IP reservation, as we will create the VRF IP reservation on creation if it does not exist. The Metal IP and Customer IP must be IPs from this subnet. For /30 subnets, the network and broadcast IPs cannot be used as the Metal or Customer IP. | [optional] 
**MetalIp** | Pointer to **string** | An IP address from the subnet that will be used on the Metal side. This parameter is optional, but if supplied, we will use the other usable IP address in the subnet as the Customer IP. By default, the first usable IP address in the subnet will be used. | [optional] 
**CustomerIp** | Pointer to **string** | An IP address from the subnet that will be used on the Customer side. This parameter is optional, but if supplied, we will use the other usable IP address in the subnet as the Metal IP. By default, the last usable IP address in the subnet will be used. | [optional] 
**Port** | Pointer to [**Href**](Href.md) |  | [optional] 

## Methods

### NewVrfVirtualCircuitsInner

`func NewVrfVirtualCircuitsInner() *VrfVirtualCircuitsInner`

NewVrfVirtualCircuitsInner instantiates a new VrfVirtualCircuitsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVrfVirtualCircuitsInnerWithDefaults

`func NewVrfVirtualCircuitsInnerWithDefaults() *VrfVirtualCircuitsInner`

NewVrfVirtualCircuitsInnerWithDefaults instantiates a new VrfVirtualCircuitsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VrfVirtualCircuitsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VrfVirtualCircuitsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VrfVirtualCircuitsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VrfVirtualCircuitsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *VrfVirtualCircuitsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VrfVirtualCircuitsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VrfVirtualCircuitsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VrfVirtualCircuitsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *VrfVirtualCircuitsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VrfVirtualCircuitsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VrfVirtualCircuitsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VrfVirtualCircuitsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMd5

`func (o *VrfVirtualCircuitsInner) GetMd5() string`

GetMd5 returns the Md5 field if non-nil, zero value otherwise.

### GetMd5Ok

`func (o *VrfVirtualCircuitsInner) GetMd5Ok() (*string, bool)`

GetMd5Ok returns a tuple with the Md5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5

`func (o *VrfVirtualCircuitsInner) SetMd5(v string)`

SetMd5 sets Md5 field to given value.

### HasMd5

`func (o *VrfVirtualCircuitsInner) HasMd5() bool`

HasMd5 returns a boolean if a field has been set.

### GetPeerAsn

`func (o *VrfVirtualCircuitsInner) GetPeerAsn() int32`

GetPeerAsn returns the PeerAsn field if non-nil, zero value otherwise.

### GetPeerAsnOk

`func (o *VrfVirtualCircuitsInner) GetPeerAsnOk() (*int32, bool)`

GetPeerAsnOk returns a tuple with the PeerAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerAsn

`func (o *VrfVirtualCircuitsInner) SetPeerAsn(v int32)`

SetPeerAsn sets PeerAsn field to given value.

### HasPeerAsn

`func (o *VrfVirtualCircuitsInner) HasPeerAsn() bool`

HasPeerAsn returns a boolean if a field has been set.

### GetSubnet

`func (o *VrfVirtualCircuitsInner) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *VrfVirtualCircuitsInner) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *VrfVirtualCircuitsInner) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *VrfVirtualCircuitsInner) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetMetalIp

`func (o *VrfVirtualCircuitsInner) GetMetalIp() string`

GetMetalIp returns the MetalIp field if non-nil, zero value otherwise.

### GetMetalIpOk

`func (o *VrfVirtualCircuitsInner) GetMetalIpOk() (*string, bool)`

GetMetalIpOk returns a tuple with the MetalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetalIp

`func (o *VrfVirtualCircuitsInner) SetMetalIp(v string)`

SetMetalIp sets MetalIp field to given value.

### HasMetalIp

`func (o *VrfVirtualCircuitsInner) HasMetalIp() bool`

HasMetalIp returns a boolean if a field has been set.

### GetCustomerIp

`func (o *VrfVirtualCircuitsInner) GetCustomerIp() string`

GetCustomerIp returns the CustomerIp field if non-nil, zero value otherwise.

### GetCustomerIpOk

`func (o *VrfVirtualCircuitsInner) GetCustomerIpOk() (*string, bool)`

GetCustomerIpOk returns a tuple with the CustomerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerIp

`func (o *VrfVirtualCircuitsInner) SetCustomerIp(v string)`

SetCustomerIp sets CustomerIp field to given value.

### HasCustomerIp

`func (o *VrfVirtualCircuitsInner) HasCustomerIp() bool`

HasCustomerIp returns a boolean if a field has been set.

### GetPort

`func (o *VrfVirtualCircuitsInner) GetPort() Href`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *VrfVirtualCircuitsInner) GetPortOk() (*Href, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *VrfVirtualCircuitsInner) SetPort(v Href)`

SetPort sets Port field to given value.

### HasPort

`func (o *VrfVirtualCircuitsInner) HasPort() bool`

HasPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


