# RequestIPReservationRequestOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comments** | Pointer to **string** |  | [optional] 
**Customdata** | Pointer to **map[string]interface{}** |  | [optional] 
**Details** | Pointer to **string** |  | [optional] 
**Facility** | Pointer to **string** |  | [optional] 
**FailOnApprovalRequired** | Pointer to **bool** |  | [optional] 
**Metro** | Pointer to **string** | The code of the metro you are requesting the IP reservation in. | [optional] 
**Quantity** | **int32** |  | 
**Tags** | Pointer to **[]string** |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewRequestIPReservationRequestOneOf

`func NewRequestIPReservationRequestOneOf(quantity int32, type_ string, ) *RequestIPReservationRequestOneOf`

NewRequestIPReservationRequestOneOf instantiates a new RequestIPReservationRequestOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestIPReservationRequestOneOfWithDefaults

`func NewRequestIPReservationRequestOneOfWithDefaults() *RequestIPReservationRequestOneOf`

NewRequestIPReservationRequestOneOfWithDefaults instantiates a new RequestIPReservationRequestOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComments

`func (o *RequestIPReservationRequestOneOf) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *RequestIPReservationRequestOneOf) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *RequestIPReservationRequestOneOf) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *RequestIPReservationRequestOneOf) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetCustomdata

`func (o *RequestIPReservationRequestOneOf) GetCustomdata() map[string]interface{}`

GetCustomdata returns the Customdata field if non-nil, zero value otherwise.

### GetCustomdataOk

`func (o *RequestIPReservationRequestOneOf) GetCustomdataOk() (*map[string]interface{}, bool)`

GetCustomdataOk returns a tuple with the Customdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomdata

`func (o *RequestIPReservationRequestOneOf) SetCustomdata(v map[string]interface{})`

SetCustomdata sets Customdata field to given value.

### HasCustomdata

`func (o *RequestIPReservationRequestOneOf) HasCustomdata() bool`

HasCustomdata returns a boolean if a field has been set.

### GetDetails

`func (o *RequestIPReservationRequestOneOf) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *RequestIPReservationRequestOneOf) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *RequestIPReservationRequestOneOf) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *RequestIPReservationRequestOneOf) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetFacility

`func (o *RequestIPReservationRequestOneOf) GetFacility() string`

GetFacility returns the Facility field if non-nil, zero value otherwise.

### GetFacilityOk

`func (o *RequestIPReservationRequestOneOf) GetFacilityOk() (*string, bool)`

GetFacilityOk returns a tuple with the Facility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacility

`func (o *RequestIPReservationRequestOneOf) SetFacility(v string)`

SetFacility sets Facility field to given value.

### HasFacility

`func (o *RequestIPReservationRequestOneOf) HasFacility() bool`

HasFacility returns a boolean if a field has been set.

### GetFailOnApprovalRequired

`func (o *RequestIPReservationRequestOneOf) GetFailOnApprovalRequired() bool`

GetFailOnApprovalRequired returns the FailOnApprovalRequired field if non-nil, zero value otherwise.

### GetFailOnApprovalRequiredOk

`func (o *RequestIPReservationRequestOneOf) GetFailOnApprovalRequiredOk() (*bool, bool)`

GetFailOnApprovalRequiredOk returns a tuple with the FailOnApprovalRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailOnApprovalRequired

`func (o *RequestIPReservationRequestOneOf) SetFailOnApprovalRequired(v bool)`

SetFailOnApprovalRequired sets FailOnApprovalRequired field to given value.

### HasFailOnApprovalRequired

`func (o *RequestIPReservationRequestOneOf) HasFailOnApprovalRequired() bool`

HasFailOnApprovalRequired returns a boolean if a field has been set.

### GetMetro

`func (o *RequestIPReservationRequestOneOf) GetMetro() string`

GetMetro returns the Metro field if non-nil, zero value otherwise.

### GetMetroOk

`func (o *RequestIPReservationRequestOneOf) GetMetroOk() (*string, bool)`

GetMetroOk returns a tuple with the Metro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetro

`func (o *RequestIPReservationRequestOneOf) SetMetro(v string)`

SetMetro sets Metro field to given value.

### HasMetro

`func (o *RequestIPReservationRequestOneOf) HasMetro() bool`

HasMetro returns a boolean if a field has been set.

### GetQuantity

`func (o *RequestIPReservationRequestOneOf) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *RequestIPReservationRequestOneOf) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *RequestIPReservationRequestOneOf) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetTags

`func (o *RequestIPReservationRequestOneOf) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RequestIPReservationRequestOneOf) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RequestIPReservationRequestOneOf) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *RequestIPReservationRequestOneOf) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *RequestIPReservationRequestOneOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RequestIPReservationRequestOneOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RequestIPReservationRequestOneOf) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


