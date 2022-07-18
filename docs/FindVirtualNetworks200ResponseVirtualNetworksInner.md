# FindVirtualNetworks200ResponseVirtualNetworksInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignedTo** | Pointer to [**FindBatchById200ResponseDevicesInner**](FindBatchById200ResponseDevicesInner.md) |  | [optional] 
**AssignedToVirtualCircuit** | Pointer to **bool** | True if the virtual network is attached to a virtual circuit. False if not. | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Facility** | Pointer to [**FindBatchById200ResponseDevicesInner**](FindBatchById200ResponseDevicesInner.md) |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Instances** | Pointer to [**[]FindBatchById200ResponseDevicesInner**](FindBatchById200ResponseDevicesInner.md) | A list of instances with ports currently associated to this Virtual Network. | [optional] 
**MetalGateway** | Pointer to [**FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway**](FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway.md) |  | [optional] 
**Metro** | Pointer to [**FindBatchById200ResponseDevicesInner**](FindBatchById200ResponseDevicesInner.md) |  | [optional] 
**MetroCode** | Pointer to **string** | The Metro code of the metro in which this Virtual Network is defined. | [optional] 
**Vxlan** | Pointer to **int32** |  | [optional] 

## Methods

### NewFindVirtualNetworks200ResponseVirtualNetworksInner

`func NewFindVirtualNetworks200ResponseVirtualNetworksInner() *FindVirtualNetworks200ResponseVirtualNetworksInner`

NewFindVirtualNetworks200ResponseVirtualNetworksInner instantiates a new FindVirtualNetworks200ResponseVirtualNetworksInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindVirtualNetworks200ResponseVirtualNetworksInnerWithDefaults

`func NewFindVirtualNetworks200ResponseVirtualNetworksInnerWithDefaults() *FindVirtualNetworks200ResponseVirtualNetworksInner`

NewFindVirtualNetworks200ResponseVirtualNetworksInnerWithDefaults instantiates a new FindVirtualNetworks200ResponseVirtualNetworksInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignedTo

`func (o *FindVirtualNetworks200ResponseVirtualNetworksInner) GetAssignedTo() FindBatchById200ResponseDevicesInner`

GetAssignedTo returns the AssignedTo field if non-nil, zero value otherwise.

### GetAssignedToOk

`func (o *FindVirtualNetworks200ResponseVirtualNetworksInner) GetAssignedToOk() (*FindBatchById200ResponseDevicesInner, bool)`

GetAssignedToOk returns a tuple with the AssignedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedTo

`func (o *FindVirtualNetworks200ResponseVirtualNetworksInner) SetAssignedTo(v FindBatchById200ResponseDevicesInner)`

SetAssignedTo sets AssignedTo field to given value.

### HasAssignedTo

`func (o *FindVirtualNetworks200ResponseVirtualNetworksInner) HasAssignedTo() bool`

HasAssignedTo returns a boolean if a field has been set.

### GetAssignedToVirtualCircuit

`func (o *FindVirtualNetworks200ResponseVirtualNetworksInner) GetAssignedToVirtualCircuit() bool`

GetAssignedToVirtualCircuit returns the AssignedToVirtualCircuit field if non-nil, zero value otherwise.

### GetAssignedToVirtualCircuitOk

`func (o *FindVirtualNetworks200ResponseVirtualNetworksInner) GetAssignedToVirtualCircuitOk() (*bool, bool)`

GetAssignedToVirtualCircuitOk returns a tuple with the AssignedToVirtualCircuit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedToVirtualCircuit

`func (o *FindVirtualNetworks200ResponseVirtualNetworksInner) SetAssignedToVirtualCircuit(v bool)`

SetAssignedToVirtualCircuit sets AssignedToVirtualCircuit field to given value.

### HasAssignedToVirtualCircuit

`func (o *FindVirtualNetworks200ResponseVirtualNetworksInner) HasAssignedToVirtualCircuit() bool`

HasAssignedToVirtualCircuit returns a boolean if a field has been set.

### GetDescription

`func (o *FindVirtualNetworks200ResponseVirtualNetworksInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FindVirtualNetworks200ResponseVirtualNetworksInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FindVirtualNetworks200ResponseVirtualNetworksInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FindVirtualNetworks200ResponseVirtualNetworksInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFacility

`func (o *FindVirtualNetworks200ResponseVirtualNetworksInner) GetFacility() FindBatchById200ResponseDevicesInner`

GetFacility returns the Facility field if non-nil, zero value otherwise.

### GetFacilityOk

`func (o *FindVirtualNetworks200ResponseVirtualNetworksInner) GetFacilityOk() (*FindBatchById200ResponseDevicesInner, bool)`

GetFacilityOk returns a tuple with the Facility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacility

`func (o *FindVirtualNetworks200ResponseVirtualNetworksInner) SetFacility(v FindBatchById200ResponseDevicesInner)`

SetFacility sets Facility field to given value.

### HasFacility

`func (o *FindVirtualNetworks200ResponseVirtualNetworksInner) HasFacility() bool`

HasFacility returns a boolean if a field has been set.

### GetHref

`func (o *FindVirtualNetworks200ResponseVirtualNetworksInner) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *FindVirtualNetworks200ResponseVirtualNetworksInner) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *FindVirtualNetworks200ResponseVirtualNetworksInner) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *FindVirtualNetworks200ResponseVirtualNetworksInner) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *FindVirtualNetworks200ResponseVirtualNetworksInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FindVirtualNetworks200ResponseVirtualNetworksInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FindVirtualNetworks200ResponseVirtualNetworksInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FindVirtualNetworks200ResponseVirtualNetworksInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInstances

`func (o *FindVirtualNetworks200ResponseVirtualNetworksInner) GetInstances() []FindBatchById200ResponseDevicesInner`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *FindVirtualNetworks200ResponseVirtualNetworksInner) GetInstancesOk() (*[]FindBatchById200ResponseDevicesInner, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *FindVirtualNetworks200ResponseVirtualNetworksInner) SetInstances(v []FindBatchById200ResponseDevicesInner)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *FindVirtualNetworks200ResponseVirtualNetworksInner) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

### GetMetalGateway

`func (o *FindVirtualNetworks200ResponseVirtualNetworksInner) GetMetalGateway() FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway`

GetMetalGateway returns the MetalGateway field if non-nil, zero value otherwise.

### GetMetalGatewayOk

`func (o *FindVirtualNetworks200ResponseVirtualNetworksInner) GetMetalGatewayOk() (*FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway, bool)`

GetMetalGatewayOk returns a tuple with the MetalGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetalGateway

`func (o *FindVirtualNetworks200ResponseVirtualNetworksInner) SetMetalGateway(v FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway)`

SetMetalGateway sets MetalGateway field to given value.

### HasMetalGateway

`func (o *FindVirtualNetworks200ResponseVirtualNetworksInner) HasMetalGateway() bool`

HasMetalGateway returns a boolean if a field has been set.

### GetMetro

`func (o *FindVirtualNetworks200ResponseVirtualNetworksInner) GetMetro() FindBatchById200ResponseDevicesInner`

GetMetro returns the Metro field if non-nil, zero value otherwise.

### GetMetroOk

`func (o *FindVirtualNetworks200ResponseVirtualNetworksInner) GetMetroOk() (*FindBatchById200ResponseDevicesInner, bool)`

GetMetroOk returns a tuple with the Metro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetro

`func (o *FindVirtualNetworks200ResponseVirtualNetworksInner) SetMetro(v FindBatchById200ResponseDevicesInner)`

SetMetro sets Metro field to given value.

### HasMetro

`func (o *FindVirtualNetworks200ResponseVirtualNetworksInner) HasMetro() bool`

HasMetro returns a boolean if a field has been set.

### GetMetroCode

`func (o *FindVirtualNetworks200ResponseVirtualNetworksInner) GetMetroCode() string`

GetMetroCode returns the MetroCode field if non-nil, zero value otherwise.

### GetMetroCodeOk

`func (o *FindVirtualNetworks200ResponseVirtualNetworksInner) GetMetroCodeOk() (*string, bool)`

GetMetroCodeOk returns a tuple with the MetroCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetroCode

`func (o *FindVirtualNetworks200ResponseVirtualNetworksInner) SetMetroCode(v string)`

SetMetroCode sets MetroCode field to given value.

### HasMetroCode

`func (o *FindVirtualNetworks200ResponseVirtualNetworksInner) HasMetroCode() bool`

HasMetroCode returns a boolean if a field has been set.

### GetVxlan

`func (o *FindVirtualNetworks200ResponseVirtualNetworksInner) GetVxlan() int32`

GetVxlan returns the Vxlan field if non-nil, zero value otherwise.

### GetVxlanOk

`func (o *FindVirtualNetworks200ResponseVirtualNetworksInner) GetVxlanOk() (*int32, bool)`

GetVxlanOk returns a tuple with the Vxlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVxlan

`func (o *FindVirtualNetworks200ResponseVirtualNetworksInner) SetVxlan(v int32)`

SetVxlan sets Vxlan field to given value.

### HasVxlan

`func (o *FindVirtualNetworks200ResponseVirtualNetworksInner) HasVxlan() bool`

HasVxlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


