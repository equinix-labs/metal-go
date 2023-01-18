# VrfRouteVirtualNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | **string** |  | 
**AssignedTo** | Pointer to [**Href**](Href.md) |  | [optional] 
**AssignedToVirtualCircuit** | Pointer to **bool** | True if the virtual network is attached to a virtual circuit. False if not. | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Facility** | Pointer to [**Href**](Href.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Instances** | Pointer to [**[]Href**](Href.md) | A list of instances with ports currently associated to this Virtual Network. | [optional] 
**MetalGateways** | Pointer to [**[]MetalGatewayLite**](MetalGatewayLite.md) | A list of metal gateways currently associated to this Virtual Network. | [optional] 
**Metro** | Pointer to [**Href**](Href.md) |  | [optional] 
**MetroCode** | Pointer to **string** | The Metro code of the metro in which this Virtual Network is defined. | [optional] 
**Vxlan** | Pointer to **int32** |  | [optional] 

## Methods

### NewVrfRouteVirtualNetwork

`func NewVrfRouteVirtualNetwork(href string, ) *VrfRouteVirtualNetwork`

NewVrfRouteVirtualNetwork instantiates a new VrfRouteVirtualNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVrfRouteVirtualNetworkWithDefaults

`func NewVrfRouteVirtualNetworkWithDefaults() *VrfRouteVirtualNetwork`

NewVrfRouteVirtualNetworkWithDefaults instantiates a new VrfRouteVirtualNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *VrfRouteVirtualNetwork) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *VrfRouteVirtualNetwork) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *VrfRouteVirtualNetwork) SetHref(v string)`

SetHref sets Href field to given value.


### GetAssignedTo

`func (o *VrfRouteVirtualNetwork) GetAssignedTo() Href`

GetAssignedTo returns the AssignedTo field if non-nil, zero value otherwise.

### GetAssignedToOk

`func (o *VrfRouteVirtualNetwork) GetAssignedToOk() (*Href, bool)`

GetAssignedToOk returns a tuple with the AssignedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedTo

`func (o *VrfRouteVirtualNetwork) SetAssignedTo(v Href)`

SetAssignedTo sets AssignedTo field to given value.

### HasAssignedTo

`func (o *VrfRouteVirtualNetwork) HasAssignedTo() bool`

HasAssignedTo returns a boolean if a field has been set.

### GetAssignedToVirtualCircuit

`func (o *VrfRouteVirtualNetwork) GetAssignedToVirtualCircuit() bool`

GetAssignedToVirtualCircuit returns the AssignedToVirtualCircuit field if non-nil, zero value otherwise.

### GetAssignedToVirtualCircuitOk

`func (o *VrfRouteVirtualNetwork) GetAssignedToVirtualCircuitOk() (*bool, bool)`

GetAssignedToVirtualCircuitOk returns a tuple with the AssignedToVirtualCircuit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedToVirtualCircuit

`func (o *VrfRouteVirtualNetwork) SetAssignedToVirtualCircuit(v bool)`

SetAssignedToVirtualCircuit sets AssignedToVirtualCircuit field to given value.

### HasAssignedToVirtualCircuit

`func (o *VrfRouteVirtualNetwork) HasAssignedToVirtualCircuit() bool`

HasAssignedToVirtualCircuit returns a boolean if a field has been set.

### GetDescription

`func (o *VrfRouteVirtualNetwork) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VrfRouteVirtualNetwork) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VrfRouteVirtualNetwork) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VrfRouteVirtualNetwork) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFacility

`func (o *VrfRouteVirtualNetwork) GetFacility() Href`

GetFacility returns the Facility field if non-nil, zero value otherwise.

### GetFacilityOk

`func (o *VrfRouteVirtualNetwork) GetFacilityOk() (*Href, bool)`

GetFacilityOk returns a tuple with the Facility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacility

`func (o *VrfRouteVirtualNetwork) SetFacility(v Href)`

SetFacility sets Facility field to given value.

### HasFacility

`func (o *VrfRouteVirtualNetwork) HasFacility() bool`

HasFacility returns a boolean if a field has been set.

### GetId

`func (o *VrfRouteVirtualNetwork) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VrfRouteVirtualNetwork) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VrfRouteVirtualNetwork) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VrfRouteVirtualNetwork) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInstances

`func (o *VrfRouteVirtualNetwork) GetInstances() []Href`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *VrfRouteVirtualNetwork) GetInstancesOk() (*[]Href, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *VrfRouteVirtualNetwork) SetInstances(v []Href)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *VrfRouteVirtualNetwork) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

### GetMetalGateways

`func (o *VrfRouteVirtualNetwork) GetMetalGateways() []MetalGatewayLite`

GetMetalGateways returns the MetalGateways field if non-nil, zero value otherwise.

### GetMetalGatewaysOk

`func (o *VrfRouteVirtualNetwork) GetMetalGatewaysOk() (*[]MetalGatewayLite, bool)`

GetMetalGatewaysOk returns a tuple with the MetalGateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetalGateways

`func (o *VrfRouteVirtualNetwork) SetMetalGateways(v []MetalGatewayLite)`

SetMetalGateways sets MetalGateways field to given value.

### HasMetalGateways

`func (o *VrfRouteVirtualNetwork) HasMetalGateways() bool`

HasMetalGateways returns a boolean if a field has been set.

### GetMetro

`func (o *VrfRouteVirtualNetwork) GetMetro() Href`

GetMetro returns the Metro field if non-nil, zero value otherwise.

### GetMetroOk

`func (o *VrfRouteVirtualNetwork) GetMetroOk() (*Href, bool)`

GetMetroOk returns a tuple with the Metro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetro

`func (o *VrfRouteVirtualNetwork) SetMetro(v Href)`

SetMetro sets Metro field to given value.

### HasMetro

`func (o *VrfRouteVirtualNetwork) HasMetro() bool`

HasMetro returns a boolean if a field has been set.

### GetMetroCode

`func (o *VrfRouteVirtualNetwork) GetMetroCode() string`

GetMetroCode returns the MetroCode field if non-nil, zero value otherwise.

### GetMetroCodeOk

`func (o *VrfRouteVirtualNetwork) GetMetroCodeOk() (*string, bool)`

GetMetroCodeOk returns a tuple with the MetroCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetroCode

`func (o *VrfRouteVirtualNetwork) SetMetroCode(v string)`

SetMetroCode sets MetroCode field to given value.

### HasMetroCode

`func (o *VrfRouteVirtualNetwork) HasMetroCode() bool`

HasMetroCode returns a boolean if a field has been set.

### GetVxlan

`func (o *VrfRouteVirtualNetwork) GetVxlan() int32`

GetVxlan returns the Vxlan field if non-nil, zero value otherwise.

### GetVxlanOk

`func (o *VrfRouteVirtualNetwork) GetVxlanOk() (*int32, bool)`

GetVxlanOk returns a tuple with the Vxlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVxlan

`func (o *VrfRouteVirtualNetwork) SetVxlan(v int32)`

SetVxlan sets Vxlan field to given value.

### HasVxlan

`func (o *VrfRouteVirtualNetwork) HasVxlan() bool`

HasVxlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


