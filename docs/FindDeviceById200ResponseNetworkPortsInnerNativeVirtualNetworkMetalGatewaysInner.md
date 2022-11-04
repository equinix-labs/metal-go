# FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetworkMetalGatewaysInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**GatewayAddress** | Pointer to **string** | The gateway address with subnet CIDR value for this Metal Gateway. For example, a Metal Gateway using an IP reservation with block 10.1.2.0/27 would have a gateway address of 10.1.2.1/27. | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** | The current state of the Metal Gateway. &#39;Ready&#39; indicates the gateway record has been configured, but is currently not active on the network. &#39;Active&#39; indicates the gateway has been configured on the network. &#39;Deleting&#39; is a temporary state used to indicate that the gateway is in the process of being un-configured from the network, after which the gateway record will be deleted. | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Vlan** | Pointer to **int32** | The VLAN id of the Virtual Network record associated to this Metal Gateway. | [optional] 

## Methods

### NewFindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetworkMetalGatewaysInner

`func NewFindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetworkMetalGatewaysInner() *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetworkMetalGatewaysInner`

NewFindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetworkMetalGatewaysInner instantiates a new FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetworkMetalGatewaysInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetworkMetalGatewaysInnerWithDefaults

`func NewFindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetworkMetalGatewaysInnerWithDefaults() *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetworkMetalGatewaysInner`

NewFindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetworkMetalGatewaysInnerWithDefaults instantiates a new FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetworkMetalGatewaysInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetworkMetalGatewaysInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetworkMetalGatewaysInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetworkMetalGatewaysInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetworkMetalGatewaysInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetGatewayAddress

`func (o *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetworkMetalGatewaysInner) GetGatewayAddress() string`

GetGatewayAddress returns the GatewayAddress field if non-nil, zero value otherwise.

### GetGatewayAddressOk

`func (o *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetworkMetalGatewaysInner) GetGatewayAddressOk() (*string, bool)`

GetGatewayAddressOk returns a tuple with the GatewayAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayAddress

`func (o *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetworkMetalGatewaysInner) SetGatewayAddress(v string)`

SetGatewayAddress sets GatewayAddress field to given value.

### HasGatewayAddress

`func (o *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetworkMetalGatewaysInner) HasGatewayAddress() bool`

HasGatewayAddress returns a boolean if a field has been set.

### GetHref

`func (o *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetworkMetalGatewaysInner) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetworkMetalGatewaysInner) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetworkMetalGatewaysInner) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetworkMetalGatewaysInner) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetworkMetalGatewaysInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetworkMetalGatewaysInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetworkMetalGatewaysInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetworkMetalGatewaysInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetState

`func (o *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetworkMetalGatewaysInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetworkMetalGatewaysInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetworkMetalGatewaysInner) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetworkMetalGatewaysInner) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetworkMetalGatewaysInner) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetworkMetalGatewaysInner) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetworkMetalGatewaysInner) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetworkMetalGatewaysInner) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVlan

`func (o *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetworkMetalGatewaysInner) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetworkMetalGatewaysInner) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetworkMetalGatewaysInner) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetworkMetalGatewaysInner) HasVlan() bool`

HasVlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


