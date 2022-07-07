# UpdateProjectRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackendTransferEnabled** | Pointer to **bool** |  | [optional] 
**Customdata** | Pointer to **map[string]interface{}** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PaymentMethodId** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateProjectRequest

`func NewUpdateProjectRequest() *UpdateProjectRequest`

NewUpdateProjectRequest instantiates a new UpdateProjectRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateProjectRequestWithDefaults

`func NewUpdateProjectRequestWithDefaults() *UpdateProjectRequest`

NewUpdateProjectRequestWithDefaults instantiates a new UpdateProjectRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackendTransferEnabled

`func (o *UpdateProjectRequest) GetBackendTransferEnabled() bool`

GetBackendTransferEnabled returns the BackendTransferEnabled field if non-nil, zero value otherwise.

### GetBackendTransferEnabledOk

`func (o *UpdateProjectRequest) GetBackendTransferEnabledOk() (*bool, bool)`

GetBackendTransferEnabledOk returns a tuple with the BackendTransferEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendTransferEnabled

`func (o *UpdateProjectRequest) SetBackendTransferEnabled(v bool)`

SetBackendTransferEnabled sets BackendTransferEnabled field to given value.

### HasBackendTransferEnabled

`func (o *UpdateProjectRequest) HasBackendTransferEnabled() bool`

HasBackendTransferEnabled returns a boolean if a field has been set.

### GetCustomdata

`func (o *UpdateProjectRequest) GetCustomdata() map[string]interface{}`

GetCustomdata returns the Customdata field if non-nil, zero value otherwise.

### GetCustomdataOk

`func (o *UpdateProjectRequest) GetCustomdataOk() (*map[string]interface{}, bool)`

GetCustomdataOk returns a tuple with the Customdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomdata

`func (o *UpdateProjectRequest) SetCustomdata(v map[string]interface{})`

SetCustomdata sets Customdata field to given value.

### HasCustomdata

`func (o *UpdateProjectRequest) HasCustomdata() bool`

HasCustomdata returns a boolean if a field has been set.

### GetName

`func (o *UpdateProjectRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateProjectRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateProjectRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateProjectRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPaymentMethodId

`func (o *UpdateProjectRequest) GetPaymentMethodId() string`

GetPaymentMethodId returns the PaymentMethodId field if non-nil, zero value otherwise.

### GetPaymentMethodIdOk

`func (o *UpdateProjectRequest) GetPaymentMethodIdOk() (*string, bool)`

GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodId

`func (o *UpdateProjectRequest) SetPaymentMethodId(v string)`

SetPaymentMethodId sets PaymentMethodId field to given value.

### HasPaymentMethodId

`func (o *UpdateProjectRequest) HasPaymentMethodId() bool`

HasPaymentMethodId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


