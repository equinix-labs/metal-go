# OrganizationList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**FindDeviceEvents200ResponseMeta**](FindDeviceEvents200ResponseMeta.md) |  | [optional] 
**Organizations** | Pointer to [**[]FindOrganizations200ResponseOrganizationsInner**](FindOrganizations200ResponseOrganizationsInner.md) |  | [optional] 

## Methods

### NewOrganizationList

`func NewOrganizationList() *OrganizationList`

NewOrganizationList instantiates a new OrganizationList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationListWithDefaults

`func NewOrganizationListWithDefaults() *OrganizationList`

NewOrganizationListWithDefaults instantiates a new OrganizationList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *OrganizationList) GetMeta() FindDeviceEvents200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *OrganizationList) GetMetaOk() (*FindDeviceEvents200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *OrganizationList) SetMeta(v FindDeviceEvents200ResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *OrganizationList) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetOrganizations

`func (o *OrganizationList) GetOrganizations() []FindOrganizations200ResponseOrganizationsInner`

GetOrganizations returns the Organizations field if non-nil, zero value otherwise.

### GetOrganizationsOk

`func (o *OrganizationList) GetOrganizationsOk() (*[]FindOrganizations200ResponseOrganizationsInner, bool)`

GetOrganizationsOk returns a tuple with the Organizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizations

`func (o *OrganizationList) SetOrganizations(v []FindOrganizations200ResponseOrganizationsInner)`

SetOrganizations sets Organizations field to given value.

### HasOrganizations

`func (o *OrganizationList) HasOrganizations() bool`

HasOrganizations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


