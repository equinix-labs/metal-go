# VrfRouteMetalGateway

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | **string** |  | 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedBy** | Pointer to [**Href**](Href.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IpReservation** | Pointer to [**IPReservation**](IPReservation.md) |  | [optional] 
**Project** | Pointer to [**Project**](Project.md) |  | [optional] 
**State** | Pointer to **string** | The current state of the Metal Gateway. &#39;Ready&#39; indicates the gateway record has been configured, but is currently not active on the network. &#39;Active&#39; indicates the gateway has been configured on the network. &#39;Deleting&#39; is a temporary state used to indicate that the gateway is in the process of being un-configured from the network, after which the gateway record will be deleted. | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**VirtualNetwork** | Pointer to [**VirtualNetwork**](VirtualNetwork.md) |  | [optional] 

## Methods

### NewVrfRouteMetalGateway

`func NewVrfRouteMetalGateway(href string, ) *VrfRouteMetalGateway`

NewVrfRouteMetalGateway instantiates a new VrfRouteMetalGateway object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVrfRouteMetalGatewayWithDefaults

`func NewVrfRouteMetalGatewayWithDefaults() *VrfRouteMetalGateway`

NewVrfRouteMetalGatewayWithDefaults instantiates a new VrfRouteMetalGateway object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *VrfRouteMetalGateway) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *VrfRouteMetalGateway) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *VrfRouteMetalGateway) SetHref(v string)`

SetHref sets Href field to given value.


### GetCreatedAt

`func (o *VrfRouteMetalGateway) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VrfRouteMetalGateway) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VrfRouteMetalGateway) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *VrfRouteMetalGateway) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *VrfRouteMetalGateway) GetCreatedBy() Href`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *VrfRouteMetalGateway) GetCreatedByOk() (*Href, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *VrfRouteMetalGateway) SetCreatedBy(v Href)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *VrfRouteMetalGateway) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetId

`func (o *VrfRouteMetalGateway) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VrfRouteMetalGateway) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VrfRouteMetalGateway) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VrfRouteMetalGateway) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIpReservation

`func (o *VrfRouteMetalGateway) GetIpReservation() IPReservation`

GetIpReservation returns the IpReservation field if non-nil, zero value otherwise.

### GetIpReservationOk

`func (o *VrfRouteMetalGateway) GetIpReservationOk() (*IPReservation, bool)`

GetIpReservationOk returns a tuple with the IpReservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpReservation

`func (o *VrfRouteMetalGateway) SetIpReservation(v IPReservation)`

SetIpReservation sets IpReservation field to given value.

### HasIpReservation

`func (o *VrfRouteMetalGateway) HasIpReservation() bool`

HasIpReservation returns a boolean if a field has been set.

### GetProject

`func (o *VrfRouteMetalGateway) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *VrfRouteMetalGateway) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *VrfRouteMetalGateway) SetProject(v Project)`

SetProject sets Project field to given value.

### HasProject

`func (o *VrfRouteMetalGateway) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetState

`func (o *VrfRouteMetalGateway) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VrfRouteMetalGateway) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *VrfRouteMetalGateway) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *VrfRouteMetalGateway) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *VrfRouteMetalGateway) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *VrfRouteMetalGateway) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *VrfRouteMetalGateway) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *VrfRouteMetalGateway) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVirtualNetwork

`func (o *VrfRouteMetalGateway) GetVirtualNetwork() VirtualNetwork`

GetVirtualNetwork returns the VirtualNetwork field if non-nil, zero value otherwise.

### GetVirtualNetworkOk

`func (o *VrfRouteMetalGateway) GetVirtualNetworkOk() (*VirtualNetwork, bool)`

GetVirtualNetworkOk returns a tuple with the VirtualNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualNetwork

`func (o *VrfRouteMetalGateway) SetVirtualNetwork(v VirtualNetwork)`

SetVirtualNetwork sets VirtualNetwork field to given value.

### HasVirtualNetwork

`func (o *VrfRouteMetalGateway) HasVirtualNetwork() bool`

HasVirtualNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


