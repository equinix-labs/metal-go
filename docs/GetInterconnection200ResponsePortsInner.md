# GetInterconnection200ResponsePortsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Organization** | Pointer to [**FindBatchById200ResponseDevicesInner**](FindBatchById200ResponseDevicesInner.md) |  | [optional] 
**Role** | Pointer to **string** | Either &#39;primary&#39; or &#39;secondary&#39;. | [optional] 
**Status** | Pointer to **string** | For both Fabric VCs and Dedicated Ports, this will be &#39;requested&#39; on creation and &#39;deleting&#39; on deletion. Once the Fabric VC has found its corresponding Fabric connection, this will turn to &#39;active&#39;. For Dedicated Ports, once the dedicated port is associated, this will also turn to &#39;active&#39;. For Fabric VCs, this can turn into an &#39;expired&#39; state if the service token associated is expired. | [optional] 
**SwitchId** | Pointer to **string** | A switch &#39;short ID&#39; | [optional] 
**VirtualCircuits** | Pointer to [**GetInterconnection200ResponsePortsInnerVirtualCircuits**](GetInterconnection200ResponsePortsInnerVirtualCircuits.md) |  | [optional] 

## Methods

### NewGetInterconnection200ResponsePortsInner

`func NewGetInterconnection200ResponsePortsInner() *GetInterconnection200ResponsePortsInner`

NewGetInterconnection200ResponsePortsInner instantiates a new GetInterconnection200ResponsePortsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInterconnection200ResponsePortsInnerWithDefaults

`func NewGetInterconnection200ResponsePortsInnerWithDefaults() *GetInterconnection200ResponsePortsInner`

NewGetInterconnection200ResponsePortsInnerWithDefaults instantiates a new GetInterconnection200ResponsePortsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetInterconnection200ResponsePortsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetInterconnection200ResponsePortsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetInterconnection200ResponsePortsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetInterconnection200ResponsePortsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrganization

`func (o *GetInterconnection200ResponsePortsInner) GetOrganization() FindBatchById200ResponseDevicesInner`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *GetInterconnection200ResponsePortsInner) GetOrganizationOk() (*FindBatchById200ResponseDevicesInner, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *GetInterconnection200ResponsePortsInner) SetOrganization(v FindBatchById200ResponseDevicesInner)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *GetInterconnection200ResponsePortsInner) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetRole

`func (o *GetInterconnection200ResponsePortsInner) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *GetInterconnection200ResponsePortsInner) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *GetInterconnection200ResponsePortsInner) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *GetInterconnection200ResponsePortsInner) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetStatus

`func (o *GetInterconnection200ResponsePortsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetInterconnection200ResponsePortsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetInterconnection200ResponsePortsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetInterconnection200ResponsePortsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSwitchId

`func (o *GetInterconnection200ResponsePortsInner) GetSwitchId() string`

GetSwitchId returns the SwitchId field if non-nil, zero value otherwise.

### GetSwitchIdOk

`func (o *GetInterconnection200ResponsePortsInner) GetSwitchIdOk() (*string, bool)`

GetSwitchIdOk returns a tuple with the SwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchId

`func (o *GetInterconnection200ResponsePortsInner) SetSwitchId(v string)`

SetSwitchId sets SwitchId field to given value.

### HasSwitchId

`func (o *GetInterconnection200ResponsePortsInner) HasSwitchId() bool`

HasSwitchId returns a boolean if a field has been set.

### GetVirtualCircuits

`func (o *GetInterconnection200ResponsePortsInner) GetVirtualCircuits() GetInterconnection200ResponsePortsInnerVirtualCircuits`

GetVirtualCircuits returns the VirtualCircuits field if non-nil, zero value otherwise.

### GetVirtualCircuitsOk

`func (o *GetInterconnection200ResponsePortsInner) GetVirtualCircuitsOk() (*GetInterconnection200ResponsePortsInnerVirtualCircuits, bool)`

GetVirtualCircuitsOk returns a tuple with the VirtualCircuits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualCircuits

`func (o *GetInterconnection200ResponsePortsInner) SetVirtualCircuits(v GetInterconnection200ResponsePortsInnerVirtualCircuits)`

SetVirtualCircuits sets VirtualCircuits field to given value.

### HasVirtualCircuits

`func (o *GetInterconnection200ResponsePortsInner) HasVirtualCircuits() bool`

HasVirtualCircuits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


