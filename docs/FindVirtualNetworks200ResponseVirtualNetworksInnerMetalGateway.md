# FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**GatewayAddress** | Pointer to **string** | The gateway address with subnet CIDR value for this Metal Gateway. For example, a Metal Gateway using an IP reservation with block 10.1.2.0/27 would have a gateway address of 10.1.2.1/27. | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** | The current state of the Metal Gateway. &#39;Ready&#39; indicates the gateway record has been configured, but is currently not active on the network. &#39;Active&#39; indicates the gateway has been configured on the network. &#39;Deleting&#39; is a temporary state used to indicate that the gateway is in the process of being un-configured from the network, after which the gateway record will be deleted. | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Vlan** | Pointer to **float32** | The VLAN id of the Virtual Network record associated to this Metal Gateway. Example: 1001. | [optional] 

## Methods

### NewFindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway

`func NewFindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway() *FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway`

NewFindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway instantiates a new FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindVirtualNetworks200ResponseVirtualNetworksInnerMetalGatewayWithDefaults

`func NewFindVirtualNetworks200ResponseVirtualNetworksInnerMetalGatewayWithDefaults() *FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway`

NewFindVirtualNetworks200ResponseVirtualNetworksInnerMetalGatewayWithDefaults instantiates a new FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetGatewayAddress

`func (o *FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway) GetGatewayAddress() string`

GetGatewayAddress returns the GatewayAddress field if non-nil, zero value otherwise.

### GetGatewayAddressOk

`func (o *FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway) GetGatewayAddressOk() (*string, bool)`

GetGatewayAddressOk returns a tuple with the GatewayAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayAddress

`func (o *FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway) SetGatewayAddress(v string)`

SetGatewayAddress sets GatewayAddress field to given value.

### HasGatewayAddress

`func (o *FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway) HasGatewayAddress() bool`

HasGatewayAddress returns a boolean if a field has been set.

### GetHref

`func (o *FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway) HasId() bool`

HasId returns a boolean if a field has been set.

### GetState

`func (o *FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVlan

`func (o *FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway) GetVlan() float32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway) GetVlanOk() (*float32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway) SetVlan(v float32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway) HasVlan() bool`

HasVlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


