# FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork

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
**MetalGateways** | Pointer to [**[]FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetworkMetalGatewaysInner**](FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetworkMetalGatewaysInner.md) | A list of metal gateways currently associated to this Virtual Network. | [optional] 
**Metro** | Pointer to [**FindBatchById200ResponseDevicesInner**](FindBatchById200ResponseDevicesInner.md) |  | [optional] 
**MetroCode** | Pointer to **string** | The Metro code of the metro in which this Virtual Network is defined. | [optional] 
**Vxlan** | Pointer to **int32** |  | [optional] 

## Methods

### NewFindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork

`func NewFindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork() *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork`

NewFindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork instantiates a new FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetworkWithDefaults

`func NewFindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetworkWithDefaults() *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork`

NewFindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetworkWithDefaults instantiates a new FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignedTo

`func (o *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork) GetAssignedTo() FindBatchById200ResponseDevicesInner`

GetAssignedTo returns the AssignedTo field if non-nil, zero value otherwise.

### GetAssignedToOk

`func (o *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork) GetAssignedToOk() (*FindBatchById200ResponseDevicesInner, bool)`

GetAssignedToOk returns a tuple with the AssignedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedTo

`func (o *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork) SetAssignedTo(v FindBatchById200ResponseDevicesInner)`

SetAssignedTo sets AssignedTo field to given value.

### HasAssignedTo

`func (o *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork) HasAssignedTo() bool`

HasAssignedTo returns a boolean if a field has been set.

### GetAssignedToVirtualCircuit

`func (o *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork) GetAssignedToVirtualCircuit() bool`

GetAssignedToVirtualCircuit returns the AssignedToVirtualCircuit field if non-nil, zero value otherwise.

### GetAssignedToVirtualCircuitOk

`func (o *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork) GetAssignedToVirtualCircuitOk() (*bool, bool)`

GetAssignedToVirtualCircuitOk returns a tuple with the AssignedToVirtualCircuit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedToVirtualCircuit

`func (o *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork) SetAssignedToVirtualCircuit(v bool)`

SetAssignedToVirtualCircuit sets AssignedToVirtualCircuit field to given value.

### HasAssignedToVirtualCircuit

`func (o *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork) HasAssignedToVirtualCircuit() bool`

HasAssignedToVirtualCircuit returns a boolean if a field has been set.

### GetDescription

`func (o *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFacility

`func (o *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork) GetFacility() FindBatchById200ResponseDevicesInner`

GetFacility returns the Facility field if non-nil, zero value otherwise.

### GetFacilityOk

`func (o *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork) GetFacilityOk() (*FindBatchById200ResponseDevicesInner, bool)`

GetFacilityOk returns a tuple with the Facility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacility

`func (o *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork) SetFacility(v FindBatchById200ResponseDevicesInner)`

SetFacility sets Facility field to given value.

### HasFacility

`func (o *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork) HasFacility() bool`

HasFacility returns a boolean if a field has been set.

### GetHref

`func (o *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInstances

`func (o *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork) GetInstances() []FindBatchById200ResponseDevicesInner`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork) GetInstancesOk() (*[]FindBatchById200ResponseDevicesInner, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork) SetInstances(v []FindBatchById200ResponseDevicesInner)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

### GetMetalGateways

`func (o *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork) GetMetalGateways() []FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetworkMetalGatewaysInner`

GetMetalGateways returns the MetalGateways field if non-nil, zero value otherwise.

### GetMetalGatewaysOk

`func (o *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork) GetMetalGatewaysOk() (*[]FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetworkMetalGatewaysInner, bool)`

GetMetalGatewaysOk returns a tuple with the MetalGateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetalGateways

`func (o *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork) SetMetalGateways(v []FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetworkMetalGatewaysInner)`

SetMetalGateways sets MetalGateways field to given value.

### HasMetalGateways

`func (o *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork) HasMetalGateways() bool`

HasMetalGateways returns a boolean if a field has been set.

### GetMetro

`func (o *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork) GetMetro() FindBatchById200ResponseDevicesInner`

GetMetro returns the Metro field if non-nil, zero value otherwise.

### GetMetroOk

`func (o *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork) GetMetroOk() (*FindBatchById200ResponseDevicesInner, bool)`

GetMetroOk returns a tuple with the Metro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetro

`func (o *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork) SetMetro(v FindBatchById200ResponseDevicesInner)`

SetMetro sets Metro field to given value.

### HasMetro

`func (o *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork) HasMetro() bool`

HasMetro returns a boolean if a field has been set.

### GetMetroCode

`func (o *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork) GetMetroCode() string`

GetMetroCode returns the MetroCode field if non-nil, zero value otherwise.

### GetMetroCodeOk

`func (o *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork) GetMetroCodeOk() (*string, bool)`

GetMetroCodeOk returns a tuple with the MetroCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetroCode

`func (o *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork) SetMetroCode(v string)`

SetMetroCode sets MetroCode field to given value.

### HasMetroCode

`func (o *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork) HasMetroCode() bool`

HasMetroCode returns a boolean if a field has been set.

### GetVxlan

`func (o *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork) GetVxlan() int32`

GetVxlan returns the Vxlan field if non-nil, zero value otherwise.

### GetVxlanOk

`func (o *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork) GetVxlanOk() (*int32, bool)`

GetVxlanOk returns a tuple with the Vxlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVxlan

`func (o *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork) SetVxlan(v int32)`

SetVxlan sets Vxlan field to given value.

### HasVxlan

`func (o *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork) HasVxlan() bool`

HasVxlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


