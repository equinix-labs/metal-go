# FindPortVlanAssignments200ResponseVlanAssignmentsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Native** | Pointer to **bool** |  | [optional] 
**Port** | Pointer to [**FindBatchById200ResponseDevicesInner**](FindBatchById200ResponseDevicesInner.md) |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**VirtualNetwork** | Pointer to [**FindBatchById200ResponseDevicesInner**](FindBatchById200ResponseDevicesInner.md) |  | [optional] 
**Vlan** | Pointer to **int32** |  | [optional] 

## Methods

### NewFindPortVlanAssignments200ResponseVlanAssignmentsInner

`func NewFindPortVlanAssignments200ResponseVlanAssignmentsInner() *FindPortVlanAssignments200ResponseVlanAssignmentsInner`

NewFindPortVlanAssignments200ResponseVlanAssignmentsInner instantiates a new FindPortVlanAssignments200ResponseVlanAssignmentsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindPortVlanAssignments200ResponseVlanAssignmentsInnerWithDefaults

`func NewFindPortVlanAssignments200ResponseVlanAssignmentsInnerWithDefaults() *FindPortVlanAssignments200ResponseVlanAssignmentsInner`

NewFindPortVlanAssignments200ResponseVlanAssignmentsInnerWithDefaults instantiates a new FindPortVlanAssignments200ResponseVlanAssignmentsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *FindPortVlanAssignments200ResponseVlanAssignmentsInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FindPortVlanAssignments200ResponseVlanAssignmentsInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FindPortVlanAssignments200ResponseVlanAssignmentsInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FindPortVlanAssignments200ResponseVlanAssignmentsInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *FindPortVlanAssignments200ResponseVlanAssignmentsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FindPortVlanAssignments200ResponseVlanAssignmentsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FindPortVlanAssignments200ResponseVlanAssignmentsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FindPortVlanAssignments200ResponseVlanAssignmentsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNative

`func (o *FindPortVlanAssignments200ResponseVlanAssignmentsInner) GetNative() bool`

GetNative returns the Native field if non-nil, zero value otherwise.

### GetNativeOk

`func (o *FindPortVlanAssignments200ResponseVlanAssignmentsInner) GetNativeOk() (*bool, bool)`

GetNativeOk returns a tuple with the Native field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNative

`func (o *FindPortVlanAssignments200ResponseVlanAssignmentsInner) SetNative(v bool)`

SetNative sets Native field to given value.

### HasNative

`func (o *FindPortVlanAssignments200ResponseVlanAssignmentsInner) HasNative() bool`

HasNative returns a boolean if a field has been set.

### GetPort

`func (o *FindPortVlanAssignments200ResponseVlanAssignmentsInner) GetPort() FindBatchById200ResponseDevicesInner`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *FindPortVlanAssignments200ResponseVlanAssignmentsInner) GetPortOk() (*FindBatchById200ResponseDevicesInner, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *FindPortVlanAssignments200ResponseVlanAssignmentsInner) SetPort(v FindBatchById200ResponseDevicesInner)`

SetPort sets Port field to given value.

### HasPort

`func (o *FindPortVlanAssignments200ResponseVlanAssignmentsInner) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetState

`func (o *FindPortVlanAssignments200ResponseVlanAssignmentsInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *FindPortVlanAssignments200ResponseVlanAssignmentsInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *FindPortVlanAssignments200ResponseVlanAssignmentsInner) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *FindPortVlanAssignments200ResponseVlanAssignmentsInner) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *FindPortVlanAssignments200ResponseVlanAssignmentsInner) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FindPortVlanAssignments200ResponseVlanAssignmentsInner) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FindPortVlanAssignments200ResponseVlanAssignmentsInner) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FindPortVlanAssignments200ResponseVlanAssignmentsInner) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVirtualNetwork

`func (o *FindPortVlanAssignments200ResponseVlanAssignmentsInner) GetVirtualNetwork() FindBatchById200ResponseDevicesInner`

GetVirtualNetwork returns the VirtualNetwork field if non-nil, zero value otherwise.

### GetVirtualNetworkOk

`func (o *FindPortVlanAssignments200ResponseVlanAssignmentsInner) GetVirtualNetworkOk() (*FindBatchById200ResponseDevicesInner, bool)`

GetVirtualNetworkOk returns a tuple with the VirtualNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualNetwork

`func (o *FindPortVlanAssignments200ResponseVlanAssignmentsInner) SetVirtualNetwork(v FindBatchById200ResponseDevicesInner)`

SetVirtualNetwork sets VirtualNetwork field to given value.

### HasVirtualNetwork

`func (o *FindPortVlanAssignments200ResponseVlanAssignmentsInner) HasVirtualNetwork() bool`

HasVirtualNetwork returns a boolean if a field has been set.

### GetVlan

`func (o *FindPortVlanAssignments200ResponseVlanAssignmentsInner) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *FindPortVlanAssignments200ResponseVlanAssignmentsInner) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *FindPortVlanAssignments200ResponseVlanAssignmentsInner) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *FindPortVlanAssignments200ResponseVlanAssignmentsInner) HasVlan() bool`

HasVlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


