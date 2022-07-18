# FindOrganizations200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**FindDeviceEvents200ResponseMeta**](FindDeviceEvents200ResponseMeta.md) |  | [optional] 
**Organizations** | Pointer to [**[]FindOrganizations200ResponseOrganizationsInner**](FindOrganizations200ResponseOrganizationsInner.md) |  | [optional] 

## Methods

### NewFindOrganizations200Response

`func NewFindOrganizations200Response() *FindOrganizations200Response`

NewFindOrganizations200Response instantiates a new FindOrganizations200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindOrganizations200ResponseWithDefaults

`func NewFindOrganizations200ResponseWithDefaults() *FindOrganizations200Response`

NewFindOrganizations200ResponseWithDefaults instantiates a new FindOrganizations200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *FindOrganizations200Response) GetMeta() FindDeviceEvents200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *FindOrganizations200Response) GetMetaOk() (*FindDeviceEvents200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *FindOrganizations200Response) SetMeta(v FindDeviceEvents200ResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *FindOrganizations200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetOrganizations

`func (o *FindOrganizations200Response) GetOrganizations() []FindOrganizations200ResponseOrganizationsInner`

GetOrganizations returns the Organizations field if non-nil, zero value otherwise.

### GetOrganizationsOk

`func (o *FindOrganizations200Response) GetOrganizationsOk() (*[]FindOrganizations200ResponseOrganizationsInner, bool)`

GetOrganizationsOk returns a tuple with the Organizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizations

`func (o *FindOrganizations200Response) SetOrganizations(v []FindOrganizations200ResponseOrganizationsInner)`

SetOrganizations sets Organizations field to given value.

### HasOrganizations

`func (o *FindOrganizations200Response) HasOrganizations() bool`

HasOrganizations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


