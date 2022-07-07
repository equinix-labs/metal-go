# RequestBgpConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asn** | **int32** |  | 
**DeploymentType** | **string** |  | 
**Md5** | Pointer to **string** |  | [optional] 
**UseCase** | Pointer to **string** |  | [optional] 

## Methods

### NewRequestBgpConfigRequest

`func NewRequestBgpConfigRequest(asn int32, deploymentType string, ) *RequestBgpConfigRequest`

NewRequestBgpConfigRequest instantiates a new RequestBgpConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestBgpConfigRequestWithDefaults

`func NewRequestBgpConfigRequestWithDefaults() *RequestBgpConfigRequest`

NewRequestBgpConfigRequestWithDefaults instantiates a new RequestBgpConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsn

`func (o *RequestBgpConfigRequest) GetAsn() int32`

GetAsn returns the Asn field if non-nil, zero value otherwise.

### GetAsnOk

`func (o *RequestBgpConfigRequest) GetAsnOk() (*int32, bool)`

GetAsnOk returns a tuple with the Asn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsn

`func (o *RequestBgpConfigRequest) SetAsn(v int32)`

SetAsn sets Asn field to given value.


### GetDeploymentType

`func (o *RequestBgpConfigRequest) GetDeploymentType() string`

GetDeploymentType returns the DeploymentType field if non-nil, zero value otherwise.

### GetDeploymentTypeOk

`func (o *RequestBgpConfigRequest) GetDeploymentTypeOk() (*string, bool)`

GetDeploymentTypeOk returns a tuple with the DeploymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentType

`func (o *RequestBgpConfigRequest) SetDeploymentType(v string)`

SetDeploymentType sets DeploymentType field to given value.


### GetMd5

`func (o *RequestBgpConfigRequest) GetMd5() string`

GetMd5 returns the Md5 field if non-nil, zero value otherwise.

### GetMd5Ok

`func (o *RequestBgpConfigRequest) GetMd5Ok() (*string, bool)`

GetMd5Ok returns a tuple with the Md5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5

`func (o *RequestBgpConfigRequest) SetMd5(v string)`

SetMd5 sets Md5 field to given value.

### HasMd5

`func (o *RequestBgpConfigRequest) HasMd5() bool`

HasMd5 returns a boolean if a field has been set.

### GetUseCase

`func (o *RequestBgpConfigRequest) GetUseCase() string`

GetUseCase returns the UseCase field if non-nil, zero value otherwise.

### GetUseCaseOk

`func (o *RequestBgpConfigRequest) GetUseCaseOk() (*string, bool)`

GetUseCaseOk returns a tuple with the UseCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCase

`func (o *RequestBgpConfigRequest) SetUseCase(v string)`

SetUseCase sets UseCase field to given value.

### HasUseCase

`func (o *RequestBgpConfigRequest) HasUseCase() bool`

HasUseCase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


