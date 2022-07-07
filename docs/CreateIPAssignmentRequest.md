# CreateIPAssignmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** |  | 
**Customdata** | Pointer to **map[string]interface{}** |  | [optional] 
**Manageable** | Pointer to **bool** |  | [optional] 

## Methods

### NewCreateIPAssignmentRequest

`func NewCreateIPAssignmentRequest(address string, ) *CreateIPAssignmentRequest`

NewCreateIPAssignmentRequest instantiates a new CreateIPAssignmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateIPAssignmentRequestWithDefaults

`func NewCreateIPAssignmentRequestWithDefaults() *CreateIPAssignmentRequest`

NewCreateIPAssignmentRequestWithDefaults instantiates a new CreateIPAssignmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *CreateIPAssignmentRequest) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CreateIPAssignmentRequest) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CreateIPAssignmentRequest) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetCustomdata

`func (o *CreateIPAssignmentRequest) GetCustomdata() map[string]interface{}`

GetCustomdata returns the Customdata field if non-nil, zero value otherwise.

### GetCustomdataOk

`func (o *CreateIPAssignmentRequest) GetCustomdataOk() (*map[string]interface{}, bool)`

GetCustomdataOk returns a tuple with the Customdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomdata

`func (o *CreateIPAssignmentRequest) SetCustomdata(v map[string]interface{})`

SetCustomdata sets Customdata field to given value.

### HasCustomdata

`func (o *CreateIPAssignmentRequest) HasCustomdata() bool`

HasCustomdata returns a boolean if a field has been set.

### GetManageable

`func (o *CreateIPAssignmentRequest) GetManageable() bool`

GetManageable returns the Manageable field if non-nil, zero value otherwise.

### GetManageableOk

`func (o *CreateIPAssignmentRequest) GetManageableOk() (*bool, bool)`

GetManageableOk returns a tuple with the Manageable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManageable

`func (o *CreateIPAssignmentRequest) SetManageable(v bool)`

SetManageable sets Manageable field to given value.

### HasManageable

`func (o *CreateIPAssignmentRequest) HasManageable() bool`

HasManageable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


