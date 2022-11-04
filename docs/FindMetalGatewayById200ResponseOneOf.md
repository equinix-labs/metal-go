# FindMetalGatewayById200ResponseOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedBy** | Pointer to [**FindBatchById200ResponseDevicesInner**](FindBatchById200ResponseDevicesInner.md) |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IpReservation** | Pointer to [**FindIPAddressById200ResponseOneOf**](FindIPAddressById200ResponseOneOf.md) |  | [optional] 
**Project** | Pointer to [**GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1VrfProject**](GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1VrfProject.md) |  | [optional] 
**State** | Pointer to **string** | The current state of the Metal Gateway. &#39;Ready&#39; indicates the gateway record has been configured, but is currently not active on the network. &#39;Active&#39; indicates the gateway has been configured on the network. &#39;Deleting&#39; is a temporary state used to indicate that the gateway is in the process of being un-configured from the network, after which the gateway record will be deleted. | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**VirtualNetwork** | Pointer to [**FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork**](FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork.md) |  | [optional] 

## Methods

### NewFindMetalGatewayById200ResponseOneOf

`func NewFindMetalGatewayById200ResponseOneOf() *FindMetalGatewayById200ResponseOneOf`

NewFindMetalGatewayById200ResponseOneOf instantiates a new FindMetalGatewayById200ResponseOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindMetalGatewayById200ResponseOneOfWithDefaults

`func NewFindMetalGatewayById200ResponseOneOfWithDefaults() *FindMetalGatewayById200ResponseOneOf`

NewFindMetalGatewayById200ResponseOneOfWithDefaults instantiates a new FindMetalGatewayById200ResponseOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *FindMetalGatewayById200ResponseOneOf) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FindMetalGatewayById200ResponseOneOf) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FindMetalGatewayById200ResponseOneOf) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FindMetalGatewayById200ResponseOneOf) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *FindMetalGatewayById200ResponseOneOf) GetCreatedBy() FindBatchById200ResponseDevicesInner`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *FindMetalGatewayById200ResponseOneOf) GetCreatedByOk() (*FindBatchById200ResponseDevicesInner, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *FindMetalGatewayById200ResponseOneOf) SetCreatedBy(v FindBatchById200ResponseDevicesInner)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *FindMetalGatewayById200ResponseOneOf) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetHref

`func (o *FindMetalGatewayById200ResponseOneOf) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *FindMetalGatewayById200ResponseOneOf) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *FindMetalGatewayById200ResponseOneOf) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *FindMetalGatewayById200ResponseOneOf) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *FindMetalGatewayById200ResponseOneOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FindMetalGatewayById200ResponseOneOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FindMetalGatewayById200ResponseOneOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FindMetalGatewayById200ResponseOneOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIpReservation

`func (o *FindMetalGatewayById200ResponseOneOf) GetIpReservation() FindIPAddressById200ResponseOneOf`

GetIpReservation returns the IpReservation field if non-nil, zero value otherwise.

### GetIpReservationOk

`func (o *FindMetalGatewayById200ResponseOneOf) GetIpReservationOk() (*FindIPAddressById200ResponseOneOf, bool)`

GetIpReservationOk returns a tuple with the IpReservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpReservation

`func (o *FindMetalGatewayById200ResponseOneOf) SetIpReservation(v FindIPAddressById200ResponseOneOf)`

SetIpReservation sets IpReservation field to given value.

### HasIpReservation

`func (o *FindMetalGatewayById200ResponseOneOf) HasIpReservation() bool`

HasIpReservation returns a boolean if a field has been set.

### GetProject

`func (o *FindMetalGatewayById200ResponseOneOf) GetProject() GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1VrfProject`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *FindMetalGatewayById200ResponseOneOf) GetProjectOk() (*GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1VrfProject, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *FindMetalGatewayById200ResponseOneOf) SetProject(v GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1VrfProject)`

SetProject sets Project field to given value.

### HasProject

`func (o *FindMetalGatewayById200ResponseOneOf) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetState

`func (o *FindMetalGatewayById200ResponseOneOf) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *FindMetalGatewayById200ResponseOneOf) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *FindMetalGatewayById200ResponseOneOf) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *FindMetalGatewayById200ResponseOneOf) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *FindMetalGatewayById200ResponseOneOf) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FindMetalGatewayById200ResponseOneOf) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FindMetalGatewayById200ResponseOneOf) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FindMetalGatewayById200ResponseOneOf) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVirtualNetwork

`func (o *FindMetalGatewayById200ResponseOneOf) GetVirtualNetwork() FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork`

GetVirtualNetwork returns the VirtualNetwork field if non-nil, zero value otherwise.

### GetVirtualNetworkOk

`func (o *FindMetalGatewayById200ResponseOneOf) GetVirtualNetworkOk() (*FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork, bool)`

GetVirtualNetworkOk returns a tuple with the VirtualNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualNetwork

`func (o *FindMetalGatewayById200ResponseOneOf) SetVirtualNetwork(v FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork)`

SetVirtualNetwork sets VirtualNetwork field to given value.

### HasVirtualNetwork

`func (o *FindMetalGatewayById200ResponseOneOf) HasVirtualNetwork() bool`

HasVirtualNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


