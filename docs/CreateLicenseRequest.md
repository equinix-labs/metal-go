# CreateLicenseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**LicenseeProductId** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **float32** |  | [optional] 

## Methods

### NewCreateLicenseRequest

`func NewCreateLicenseRequest() *CreateLicenseRequest`

NewCreateLicenseRequest instantiates a new CreateLicenseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLicenseRequestWithDefaults

`func NewCreateLicenseRequestWithDefaults() *CreateLicenseRequest`

NewCreateLicenseRequestWithDefaults instantiates a new CreateLicenseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateLicenseRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateLicenseRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateLicenseRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateLicenseRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLicenseeProductId

`func (o *CreateLicenseRequest) GetLicenseeProductId() string`

GetLicenseeProductId returns the LicenseeProductId field if non-nil, zero value otherwise.

### GetLicenseeProductIdOk

`func (o *CreateLicenseRequest) GetLicenseeProductIdOk() (*string, bool)`

GetLicenseeProductIdOk returns a tuple with the LicenseeProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseeProductId

`func (o *CreateLicenseRequest) SetLicenseeProductId(v string)`

SetLicenseeProductId sets LicenseeProductId field to given value.

### HasLicenseeProductId

`func (o *CreateLicenseRequest) HasLicenseeProductId() bool`

HasLicenseeProductId returns a boolean if a field has been set.

### GetSize

`func (o *CreateLicenseRequest) GetSize() float32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *CreateLicenseRequest) GetSizeOk() (*float32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *CreateLicenseRequest) SetSize(v float32)`

SetSize sets Size field to given value.

### HasSize

`func (o *CreateLicenseRequest) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


