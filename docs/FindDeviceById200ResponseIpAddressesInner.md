# FindDeviceById200ResponseIpAddressesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** |  | [optional] 
**AddressFamily** | Pointer to **int32** |  | [optional] 
**AssignedTo** | Pointer to [**FindBatchById200ResponseDevicesInner**](FindBatchById200ResponseDevicesInner.md) |  | [optional] 
**Cidr** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Gateway** | Pointer to **string** |  | [optional] 
**GlobalIp** | Pointer to **bool** |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Manageable** | Pointer to **bool** |  | [optional] 
**Management** | Pointer to **bool** |  | [optional] 
**Metro** | Pointer to [**FindDeviceById200ResponseIpAddressesInnerMetro**](FindDeviceById200ResponseIpAddressesInnerMetro.md) |  | [optional] 
**Netmask** | Pointer to **string** |  | [optional] 
**Network** | Pointer to **string** |  | [optional] 
**ParentBlock** | Pointer to [**FindDeviceById200ResponseIpAddressesInnerParentBlock**](FindDeviceById200ResponseIpAddressesInnerParentBlock.md) |  | [optional] 
**Public** | Pointer to **bool** |  | [optional] 

## Methods

### NewFindDeviceById200ResponseIpAddressesInner

`func NewFindDeviceById200ResponseIpAddressesInner() *FindDeviceById200ResponseIpAddressesInner`

NewFindDeviceById200ResponseIpAddressesInner instantiates a new FindDeviceById200ResponseIpAddressesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindDeviceById200ResponseIpAddressesInnerWithDefaults

`func NewFindDeviceById200ResponseIpAddressesInnerWithDefaults() *FindDeviceById200ResponseIpAddressesInner`

NewFindDeviceById200ResponseIpAddressesInnerWithDefaults instantiates a new FindDeviceById200ResponseIpAddressesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *FindDeviceById200ResponseIpAddressesInner) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *FindDeviceById200ResponseIpAddressesInner) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *FindDeviceById200ResponseIpAddressesInner) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *FindDeviceById200ResponseIpAddressesInner) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAddressFamily

`func (o *FindDeviceById200ResponseIpAddressesInner) GetAddressFamily() int32`

GetAddressFamily returns the AddressFamily field if non-nil, zero value otherwise.

### GetAddressFamilyOk

`func (o *FindDeviceById200ResponseIpAddressesInner) GetAddressFamilyOk() (*int32, bool)`

GetAddressFamilyOk returns a tuple with the AddressFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressFamily

`func (o *FindDeviceById200ResponseIpAddressesInner) SetAddressFamily(v int32)`

SetAddressFamily sets AddressFamily field to given value.

### HasAddressFamily

`func (o *FindDeviceById200ResponseIpAddressesInner) HasAddressFamily() bool`

HasAddressFamily returns a boolean if a field has been set.

### GetAssignedTo

`func (o *FindDeviceById200ResponseIpAddressesInner) GetAssignedTo() FindBatchById200ResponseDevicesInner`

GetAssignedTo returns the AssignedTo field if non-nil, zero value otherwise.

### GetAssignedToOk

`func (o *FindDeviceById200ResponseIpAddressesInner) GetAssignedToOk() (*FindBatchById200ResponseDevicesInner, bool)`

GetAssignedToOk returns a tuple with the AssignedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedTo

`func (o *FindDeviceById200ResponseIpAddressesInner) SetAssignedTo(v FindBatchById200ResponseDevicesInner)`

SetAssignedTo sets AssignedTo field to given value.

### HasAssignedTo

`func (o *FindDeviceById200ResponseIpAddressesInner) HasAssignedTo() bool`

HasAssignedTo returns a boolean if a field has been set.

### GetCidr

`func (o *FindDeviceById200ResponseIpAddressesInner) GetCidr() int32`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *FindDeviceById200ResponseIpAddressesInner) GetCidrOk() (*int32, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *FindDeviceById200ResponseIpAddressesInner) SetCidr(v int32)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *FindDeviceById200ResponseIpAddressesInner) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FindDeviceById200ResponseIpAddressesInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FindDeviceById200ResponseIpAddressesInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FindDeviceById200ResponseIpAddressesInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FindDeviceById200ResponseIpAddressesInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEnabled

`func (o *FindDeviceById200ResponseIpAddressesInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *FindDeviceById200ResponseIpAddressesInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *FindDeviceById200ResponseIpAddressesInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *FindDeviceById200ResponseIpAddressesInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetGateway

`func (o *FindDeviceById200ResponseIpAddressesInner) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *FindDeviceById200ResponseIpAddressesInner) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *FindDeviceById200ResponseIpAddressesInner) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *FindDeviceById200ResponseIpAddressesInner) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetGlobalIp

`func (o *FindDeviceById200ResponseIpAddressesInner) GetGlobalIp() bool`

GetGlobalIp returns the GlobalIp field if non-nil, zero value otherwise.

### GetGlobalIpOk

`func (o *FindDeviceById200ResponseIpAddressesInner) GetGlobalIpOk() (*bool, bool)`

GetGlobalIpOk returns a tuple with the GlobalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalIp

`func (o *FindDeviceById200ResponseIpAddressesInner) SetGlobalIp(v bool)`

SetGlobalIp sets GlobalIp field to given value.

### HasGlobalIp

`func (o *FindDeviceById200ResponseIpAddressesInner) HasGlobalIp() bool`

HasGlobalIp returns a boolean if a field has been set.

### GetHref

`func (o *FindDeviceById200ResponseIpAddressesInner) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *FindDeviceById200ResponseIpAddressesInner) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *FindDeviceById200ResponseIpAddressesInner) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *FindDeviceById200ResponseIpAddressesInner) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *FindDeviceById200ResponseIpAddressesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FindDeviceById200ResponseIpAddressesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FindDeviceById200ResponseIpAddressesInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FindDeviceById200ResponseIpAddressesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetManageable

`func (o *FindDeviceById200ResponseIpAddressesInner) GetManageable() bool`

GetManageable returns the Manageable field if non-nil, zero value otherwise.

### GetManageableOk

`func (o *FindDeviceById200ResponseIpAddressesInner) GetManageableOk() (*bool, bool)`

GetManageableOk returns a tuple with the Manageable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManageable

`func (o *FindDeviceById200ResponseIpAddressesInner) SetManageable(v bool)`

SetManageable sets Manageable field to given value.

### HasManageable

`func (o *FindDeviceById200ResponseIpAddressesInner) HasManageable() bool`

HasManageable returns a boolean if a field has been set.

### GetManagement

`func (o *FindDeviceById200ResponseIpAddressesInner) GetManagement() bool`

GetManagement returns the Management field if non-nil, zero value otherwise.

### GetManagementOk

`func (o *FindDeviceById200ResponseIpAddressesInner) GetManagementOk() (*bool, bool)`

GetManagementOk returns a tuple with the Management field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagement

`func (o *FindDeviceById200ResponseIpAddressesInner) SetManagement(v bool)`

SetManagement sets Management field to given value.

### HasManagement

`func (o *FindDeviceById200ResponseIpAddressesInner) HasManagement() bool`

HasManagement returns a boolean if a field has been set.

### GetMetro

`func (o *FindDeviceById200ResponseIpAddressesInner) GetMetro() FindDeviceById200ResponseIpAddressesInnerMetro`

GetMetro returns the Metro field if non-nil, zero value otherwise.

### GetMetroOk

`func (o *FindDeviceById200ResponseIpAddressesInner) GetMetroOk() (*FindDeviceById200ResponseIpAddressesInnerMetro, bool)`

GetMetroOk returns a tuple with the Metro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetro

`func (o *FindDeviceById200ResponseIpAddressesInner) SetMetro(v FindDeviceById200ResponseIpAddressesInnerMetro)`

SetMetro sets Metro field to given value.

### HasMetro

`func (o *FindDeviceById200ResponseIpAddressesInner) HasMetro() bool`

HasMetro returns a boolean if a field has been set.

### GetNetmask

`func (o *FindDeviceById200ResponseIpAddressesInner) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *FindDeviceById200ResponseIpAddressesInner) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *FindDeviceById200ResponseIpAddressesInner) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *FindDeviceById200ResponseIpAddressesInner) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.

### GetNetwork

`func (o *FindDeviceById200ResponseIpAddressesInner) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *FindDeviceById200ResponseIpAddressesInner) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *FindDeviceById200ResponseIpAddressesInner) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *FindDeviceById200ResponseIpAddressesInner) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetParentBlock

`func (o *FindDeviceById200ResponseIpAddressesInner) GetParentBlock() FindDeviceById200ResponseIpAddressesInnerParentBlock`

GetParentBlock returns the ParentBlock field if non-nil, zero value otherwise.

### GetParentBlockOk

`func (o *FindDeviceById200ResponseIpAddressesInner) GetParentBlockOk() (*FindDeviceById200ResponseIpAddressesInnerParentBlock, bool)`

GetParentBlockOk returns a tuple with the ParentBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentBlock

`func (o *FindDeviceById200ResponseIpAddressesInner) SetParentBlock(v FindDeviceById200ResponseIpAddressesInnerParentBlock)`

SetParentBlock sets ParentBlock field to given value.

### HasParentBlock

`func (o *FindDeviceById200ResponseIpAddressesInner) HasParentBlock() bool`

HasParentBlock returns a boolean if a field has been set.

### GetPublic

`func (o *FindDeviceById200ResponseIpAddressesInner) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *FindDeviceById200ResponseIpAddressesInner) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *FindDeviceById200ResponseIpAddressesInner) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *FindDeviceById200ResponseIpAddressesInner) HasPublic() bool`

HasPublic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


