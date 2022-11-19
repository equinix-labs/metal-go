# InterconnectionPort

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

### NewInterconnectionPort

`func NewInterconnectionPort() *InterconnectionPort`

NewInterconnectionPort instantiates a new InterconnectionPort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInterconnectionPortWithDefaults

`func NewInterconnectionPortWithDefaults() *InterconnectionPort`

NewInterconnectionPortWithDefaults instantiates a new InterconnectionPort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InterconnectionPort) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InterconnectionPort) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InterconnectionPort) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InterconnectionPort) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrganization

`func (o *InterconnectionPort) GetOrganization() FindBatchById200ResponseDevicesInner`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *InterconnectionPort) GetOrganizationOk() (*FindBatchById200ResponseDevicesInner, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *InterconnectionPort) SetOrganization(v FindBatchById200ResponseDevicesInner)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *InterconnectionPort) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetRole

`func (o *InterconnectionPort) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *InterconnectionPort) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *InterconnectionPort) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *InterconnectionPort) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetStatus

`func (o *InterconnectionPort) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InterconnectionPort) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InterconnectionPort) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InterconnectionPort) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSwitchId

`func (o *InterconnectionPort) GetSwitchId() string`

GetSwitchId returns the SwitchId field if non-nil, zero value otherwise.

### GetSwitchIdOk

`func (o *InterconnectionPort) GetSwitchIdOk() (*string, bool)`

GetSwitchIdOk returns a tuple with the SwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchId

`func (o *InterconnectionPort) SetSwitchId(v string)`

SetSwitchId sets SwitchId field to given value.

### HasSwitchId

`func (o *InterconnectionPort) HasSwitchId() bool`

HasSwitchId returns a boolean if a field has been set.

### GetVirtualCircuits

`func (o *InterconnectionPort) GetVirtualCircuits() GetInterconnection200ResponsePortsInnerVirtualCircuits`

GetVirtualCircuits returns the VirtualCircuits field if non-nil, zero value otherwise.

### GetVirtualCircuitsOk

`func (o *InterconnectionPort) GetVirtualCircuitsOk() (*GetInterconnection200ResponsePortsInnerVirtualCircuits, bool)`

GetVirtualCircuitsOk returns a tuple with the VirtualCircuits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualCircuits

`func (o *InterconnectionPort) SetVirtualCircuits(v GetInterconnection200ResponsePortsInnerVirtualCircuits)`

SetVirtualCircuits sets VirtualCircuits field to given value.

### HasVirtualCircuits

`func (o *InterconnectionPort) HasVirtualCircuits() bool`

HasVirtualCircuits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


