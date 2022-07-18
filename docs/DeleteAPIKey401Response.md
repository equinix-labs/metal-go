# DeleteAPIKey401Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** | A description of the error that caused the request to fail. | [optional] 
**Errors** | Pointer to **[]string** | A list of errors that contributed to the request failing. | [optional] 

## Methods

### NewDeleteAPIKey401Response

`func NewDeleteAPIKey401Response() *DeleteAPIKey401Response`

NewDeleteAPIKey401Response instantiates a new DeleteAPIKey401Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteAPIKey401ResponseWithDefaults

`func NewDeleteAPIKey401ResponseWithDefaults() *DeleteAPIKey401Response`

NewDeleteAPIKey401ResponseWithDefaults instantiates a new DeleteAPIKey401Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *DeleteAPIKey401Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *DeleteAPIKey401Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *DeleteAPIKey401Response) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *DeleteAPIKey401Response) HasError() bool`

HasError returns a boolean if a field has been set.

### GetErrors

`func (o *DeleteAPIKey401Response) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *DeleteAPIKey401Response) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *DeleteAPIKey401Response) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *DeleteAPIKey401Response) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


