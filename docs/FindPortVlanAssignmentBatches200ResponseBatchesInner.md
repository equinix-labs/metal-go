# FindPortVlanAssignmentBatches200ResponseBatchesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**ErrorMessages** | Pointer to **[]string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Port** | Pointer to [**FindDeviceById200ResponseNetworkPortsInner**](FindDeviceById200ResponseNetworkPortsInner.md) |  | [optional] 
**Quantity** | Pointer to **int32** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**VlanAssignments** | Pointer to [**[]FindPortVlanAssignmentBatches200ResponseBatchesInnerVlanAssignmentsInner**](FindPortVlanAssignmentBatches200ResponseBatchesInnerVlanAssignmentsInner.md) |  | [optional] 

## Methods

### NewFindPortVlanAssignmentBatches200ResponseBatchesInner

`func NewFindPortVlanAssignmentBatches200ResponseBatchesInner() *FindPortVlanAssignmentBatches200ResponseBatchesInner`

NewFindPortVlanAssignmentBatches200ResponseBatchesInner instantiates a new FindPortVlanAssignmentBatches200ResponseBatchesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindPortVlanAssignmentBatches200ResponseBatchesInnerWithDefaults

`func NewFindPortVlanAssignmentBatches200ResponseBatchesInnerWithDefaults() *FindPortVlanAssignmentBatches200ResponseBatchesInner`

NewFindPortVlanAssignmentBatches200ResponseBatchesInnerWithDefaults instantiates a new FindPortVlanAssignmentBatches200ResponseBatchesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *FindPortVlanAssignmentBatches200ResponseBatchesInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FindPortVlanAssignmentBatches200ResponseBatchesInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FindPortVlanAssignmentBatches200ResponseBatchesInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FindPortVlanAssignmentBatches200ResponseBatchesInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetErrorMessages

`func (o *FindPortVlanAssignmentBatches200ResponseBatchesInner) GetErrorMessages() []string`

GetErrorMessages returns the ErrorMessages field if non-nil, zero value otherwise.

### GetErrorMessagesOk

`func (o *FindPortVlanAssignmentBatches200ResponseBatchesInner) GetErrorMessagesOk() (*[]string, bool)`

GetErrorMessagesOk returns a tuple with the ErrorMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessages

`func (o *FindPortVlanAssignmentBatches200ResponseBatchesInner) SetErrorMessages(v []string)`

SetErrorMessages sets ErrorMessages field to given value.

### HasErrorMessages

`func (o *FindPortVlanAssignmentBatches200ResponseBatchesInner) HasErrorMessages() bool`

HasErrorMessages returns a boolean if a field has been set.

### GetId

`func (o *FindPortVlanAssignmentBatches200ResponseBatchesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FindPortVlanAssignmentBatches200ResponseBatchesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FindPortVlanAssignmentBatches200ResponseBatchesInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FindPortVlanAssignmentBatches200ResponseBatchesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPort

`func (o *FindPortVlanAssignmentBatches200ResponseBatchesInner) GetPort() FindDeviceById200ResponseNetworkPortsInner`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *FindPortVlanAssignmentBatches200ResponseBatchesInner) GetPortOk() (*FindDeviceById200ResponseNetworkPortsInner, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *FindPortVlanAssignmentBatches200ResponseBatchesInner) SetPort(v FindDeviceById200ResponseNetworkPortsInner)`

SetPort sets Port field to given value.

### HasPort

`func (o *FindPortVlanAssignmentBatches200ResponseBatchesInner) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetQuantity

`func (o *FindPortVlanAssignmentBatches200ResponseBatchesInner) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *FindPortVlanAssignmentBatches200ResponseBatchesInner) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *FindPortVlanAssignmentBatches200ResponseBatchesInner) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *FindPortVlanAssignmentBatches200ResponseBatchesInner) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetState

`func (o *FindPortVlanAssignmentBatches200ResponseBatchesInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *FindPortVlanAssignmentBatches200ResponseBatchesInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *FindPortVlanAssignmentBatches200ResponseBatchesInner) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *FindPortVlanAssignmentBatches200ResponseBatchesInner) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *FindPortVlanAssignmentBatches200ResponseBatchesInner) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FindPortVlanAssignmentBatches200ResponseBatchesInner) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FindPortVlanAssignmentBatches200ResponseBatchesInner) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FindPortVlanAssignmentBatches200ResponseBatchesInner) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVlanAssignments

`func (o *FindPortVlanAssignmentBatches200ResponseBatchesInner) GetVlanAssignments() []FindPortVlanAssignmentBatches200ResponseBatchesInnerVlanAssignmentsInner`

GetVlanAssignments returns the VlanAssignments field if non-nil, zero value otherwise.

### GetVlanAssignmentsOk

`func (o *FindPortVlanAssignmentBatches200ResponseBatchesInner) GetVlanAssignmentsOk() (*[]FindPortVlanAssignmentBatches200ResponseBatchesInnerVlanAssignmentsInner, bool)`

GetVlanAssignmentsOk returns a tuple with the VlanAssignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanAssignments

`func (o *FindPortVlanAssignmentBatches200ResponseBatchesInner) SetVlanAssignments(v []FindPortVlanAssignmentBatches200ResponseBatchesInnerVlanAssignmentsInner)`

SetVlanAssignments sets VlanAssignments field to given value.

### HasVlanAssignments

`func (o *FindPortVlanAssignmentBatches200ResponseBatchesInner) HasVlanAssignments() bool`

HasVlanAssignments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


