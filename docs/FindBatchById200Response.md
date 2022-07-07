# FindBatchById200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Devices** | Pointer to [**[]FindBatchById200ResponseDevicesInner**](FindBatchById200ResponseDevicesInner.md) |  | [optional] 
**ErrorMessages** | Pointer to **[]string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Project** | Pointer to [**FindBatchById200ResponseDevicesInner**](FindBatchById200ResponseDevicesInner.md) |  | [optional] 
**Quantity** | Pointer to **int32** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewFindBatchById200Response

`func NewFindBatchById200Response() *FindBatchById200Response`

NewFindBatchById200Response instantiates a new FindBatchById200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindBatchById200ResponseWithDefaults

`func NewFindBatchById200ResponseWithDefaults() *FindBatchById200Response`

NewFindBatchById200ResponseWithDefaults instantiates a new FindBatchById200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *FindBatchById200Response) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FindBatchById200Response) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FindBatchById200Response) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FindBatchById200Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDevices

`func (o *FindBatchById200Response) GetDevices() []FindBatchById200ResponseDevicesInner`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *FindBatchById200Response) GetDevicesOk() (*[]FindBatchById200ResponseDevicesInner, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *FindBatchById200Response) SetDevices(v []FindBatchById200ResponseDevicesInner)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *FindBatchById200Response) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### GetErrorMessages

`func (o *FindBatchById200Response) GetErrorMessages() []string`

GetErrorMessages returns the ErrorMessages field if non-nil, zero value otherwise.

### GetErrorMessagesOk

`func (o *FindBatchById200Response) GetErrorMessagesOk() (*[]string, bool)`

GetErrorMessagesOk returns a tuple with the ErrorMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessages

`func (o *FindBatchById200Response) SetErrorMessages(v []string)`

SetErrorMessages sets ErrorMessages field to given value.

### HasErrorMessages

`func (o *FindBatchById200Response) HasErrorMessages() bool`

HasErrorMessages returns a boolean if a field has been set.

### GetId

`func (o *FindBatchById200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FindBatchById200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FindBatchById200Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FindBatchById200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProject

`func (o *FindBatchById200Response) GetProject() FindBatchById200ResponseDevicesInner`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *FindBatchById200Response) GetProjectOk() (*FindBatchById200ResponseDevicesInner, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *FindBatchById200Response) SetProject(v FindBatchById200ResponseDevicesInner)`

SetProject sets Project field to given value.

### HasProject

`func (o *FindBatchById200Response) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetQuantity

`func (o *FindBatchById200Response) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *FindBatchById200Response) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *FindBatchById200Response) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *FindBatchById200Response) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetState

`func (o *FindBatchById200Response) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *FindBatchById200Response) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *FindBatchById200Response) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *FindBatchById200Response) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *FindBatchById200Response) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FindBatchById200Response) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FindBatchById200Response) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FindBatchById200Response) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


