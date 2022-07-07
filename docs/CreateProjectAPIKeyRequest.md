# CreateProjectAPIKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**ReadOnly** | Pointer to **bool** |  | [optional] 

## Methods

### NewCreateProjectAPIKeyRequest

`func NewCreateProjectAPIKeyRequest() *CreateProjectAPIKeyRequest`

NewCreateProjectAPIKeyRequest instantiates a new CreateProjectAPIKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProjectAPIKeyRequestWithDefaults

`func NewCreateProjectAPIKeyRequestWithDefaults() *CreateProjectAPIKeyRequest`

NewCreateProjectAPIKeyRequestWithDefaults instantiates a new CreateProjectAPIKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateProjectAPIKeyRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateProjectAPIKeyRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateProjectAPIKeyRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateProjectAPIKeyRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetReadOnly

`func (o *CreateProjectAPIKeyRequest) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *CreateProjectAPIKeyRequest) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *CreateProjectAPIKeyRequest) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *CreateProjectAPIKeyRequest) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


