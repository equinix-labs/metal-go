# FindMetalGatewaysByProject200ResponseMetalGatewaysInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedBy** | Pointer to [**FindBatchById200ResponseDevicesInner**](FindBatchById200ResponseDevicesInner.md) |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IpReservation** | Pointer to [**FindIPAddressById200ResponseOneOf1**](FindIPAddressById200ResponseOneOf1.md) |  | [optional] 
**Project** | Pointer to [**GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1VrfProject**](GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1VrfProject.md) |  | [optional] 
**State** | Pointer to **string** | The current state of the Metal Gateway. &#39;Ready&#39; indicates the gateway record has been configured, but is currently not active on the network. &#39;Active&#39; indicates the gateway has been configured on the network. &#39;Deleting&#39; is a temporary state used to indicate that the gateway is in the process of being un-configured from the network, after which the gateway record will be deleted. | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**VirtualNetwork** | Pointer to [**FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork**](FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork.md) |  | [optional] 
**Vrf** | Pointer to [**GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1Vrf**](GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1Vrf.md) |  | [optional] 

## Methods

### NewFindMetalGatewaysByProject200ResponseMetalGatewaysInner

`func NewFindMetalGatewaysByProject200ResponseMetalGatewaysInner() *FindMetalGatewaysByProject200ResponseMetalGatewaysInner`

NewFindMetalGatewaysByProject200ResponseMetalGatewaysInner instantiates a new FindMetalGatewaysByProject200ResponseMetalGatewaysInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindMetalGatewaysByProject200ResponseMetalGatewaysInnerWithDefaults

`func NewFindMetalGatewaysByProject200ResponseMetalGatewaysInnerWithDefaults() *FindMetalGatewaysByProject200ResponseMetalGatewaysInner`

NewFindMetalGatewaysByProject200ResponseMetalGatewaysInnerWithDefaults instantiates a new FindMetalGatewaysByProject200ResponseMetalGatewaysInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *FindMetalGatewaysByProject200ResponseMetalGatewaysInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FindMetalGatewaysByProject200ResponseMetalGatewaysInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FindMetalGatewaysByProject200ResponseMetalGatewaysInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FindMetalGatewaysByProject200ResponseMetalGatewaysInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *FindMetalGatewaysByProject200ResponseMetalGatewaysInner) GetCreatedBy() FindBatchById200ResponseDevicesInner`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *FindMetalGatewaysByProject200ResponseMetalGatewaysInner) GetCreatedByOk() (*FindBatchById200ResponseDevicesInner, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *FindMetalGatewaysByProject200ResponseMetalGatewaysInner) SetCreatedBy(v FindBatchById200ResponseDevicesInner)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *FindMetalGatewaysByProject200ResponseMetalGatewaysInner) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetHref

`func (o *FindMetalGatewaysByProject200ResponseMetalGatewaysInner) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *FindMetalGatewaysByProject200ResponseMetalGatewaysInner) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *FindMetalGatewaysByProject200ResponseMetalGatewaysInner) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *FindMetalGatewaysByProject200ResponseMetalGatewaysInner) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *FindMetalGatewaysByProject200ResponseMetalGatewaysInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FindMetalGatewaysByProject200ResponseMetalGatewaysInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FindMetalGatewaysByProject200ResponseMetalGatewaysInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FindMetalGatewaysByProject200ResponseMetalGatewaysInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIpReservation

`func (o *FindMetalGatewaysByProject200ResponseMetalGatewaysInner) GetIpReservation() FindIPAddressById200ResponseOneOf1`

GetIpReservation returns the IpReservation field if non-nil, zero value otherwise.

### GetIpReservationOk

`func (o *FindMetalGatewaysByProject200ResponseMetalGatewaysInner) GetIpReservationOk() (*FindIPAddressById200ResponseOneOf1, bool)`

GetIpReservationOk returns a tuple with the IpReservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpReservation

`func (o *FindMetalGatewaysByProject200ResponseMetalGatewaysInner) SetIpReservation(v FindIPAddressById200ResponseOneOf1)`

SetIpReservation sets IpReservation field to given value.

### HasIpReservation

`func (o *FindMetalGatewaysByProject200ResponseMetalGatewaysInner) HasIpReservation() bool`

HasIpReservation returns a boolean if a field has been set.

### GetProject

`func (o *FindMetalGatewaysByProject200ResponseMetalGatewaysInner) GetProject() GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1VrfProject`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *FindMetalGatewaysByProject200ResponseMetalGatewaysInner) GetProjectOk() (*GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1VrfProject, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *FindMetalGatewaysByProject200ResponseMetalGatewaysInner) SetProject(v GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1VrfProject)`

SetProject sets Project field to given value.

### HasProject

`func (o *FindMetalGatewaysByProject200ResponseMetalGatewaysInner) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetState

`func (o *FindMetalGatewaysByProject200ResponseMetalGatewaysInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *FindMetalGatewaysByProject200ResponseMetalGatewaysInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *FindMetalGatewaysByProject200ResponseMetalGatewaysInner) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *FindMetalGatewaysByProject200ResponseMetalGatewaysInner) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *FindMetalGatewaysByProject200ResponseMetalGatewaysInner) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FindMetalGatewaysByProject200ResponseMetalGatewaysInner) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FindMetalGatewaysByProject200ResponseMetalGatewaysInner) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FindMetalGatewaysByProject200ResponseMetalGatewaysInner) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVirtualNetwork

`func (o *FindMetalGatewaysByProject200ResponseMetalGatewaysInner) GetVirtualNetwork() FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork`

GetVirtualNetwork returns the VirtualNetwork field if non-nil, zero value otherwise.

### GetVirtualNetworkOk

`func (o *FindMetalGatewaysByProject200ResponseMetalGatewaysInner) GetVirtualNetworkOk() (*FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork, bool)`

GetVirtualNetworkOk returns a tuple with the VirtualNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualNetwork

`func (o *FindMetalGatewaysByProject200ResponseMetalGatewaysInner) SetVirtualNetwork(v FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork)`

SetVirtualNetwork sets VirtualNetwork field to given value.

### HasVirtualNetwork

`func (o *FindMetalGatewaysByProject200ResponseMetalGatewaysInner) HasVirtualNetwork() bool`

HasVirtualNetwork returns a boolean if a field has been set.

### GetVrf

`func (o *FindMetalGatewaysByProject200ResponseMetalGatewaysInner) GetVrf() GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1Vrf`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *FindMetalGatewaysByProject200ResponseMetalGatewaysInner) GetVrfOk() (*GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1Vrf, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *FindMetalGatewaysByProject200ResponseMetalGatewaysInner) SetVrf(v GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1Vrf)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *FindMetalGatewaysByProject200ResponseMetalGatewaysInner) HasVrf() bool`

HasVrf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


