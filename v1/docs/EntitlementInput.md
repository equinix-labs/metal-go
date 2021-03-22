# EntitlementInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**Weight** | Pointer to **int32** |  | [optional] 
**InstanceQuota** | Pointer to **map[string]interface{}** |  | [optional] 
**ProjectQuota** | Pointer to **int32** |  | [optional] 
**VolumeQuota** | Pointer to **map[string]interface{}** |  | [optional] 
**FeatureAccess** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewEntitlementInput

`func NewEntitlementInput() *EntitlementInput`

NewEntitlementInput instantiates a new EntitlementInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementInputWithDefaults

`func NewEntitlementInputWithDefaults() *EntitlementInput`

NewEntitlementInputWithDefaults instantiates a new EntitlementInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *EntitlementInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EntitlementInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EntitlementInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EntitlementInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSlug

`func (o *EntitlementInput) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *EntitlementInput) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *EntitlementInput) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *EntitlementInput) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetWeight

`func (o *EntitlementInput) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *EntitlementInput) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *EntitlementInput) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *EntitlementInput) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetInstanceQuota

`func (o *EntitlementInput) GetInstanceQuota() map[string]interface{}`

GetInstanceQuota returns the InstanceQuota field if non-nil, zero value otherwise.

### GetInstanceQuotaOk

`func (o *EntitlementInput) GetInstanceQuotaOk() (*map[string]interface{}, bool)`

GetInstanceQuotaOk returns a tuple with the InstanceQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceQuota

`func (o *EntitlementInput) SetInstanceQuota(v map[string]interface{})`

SetInstanceQuota sets InstanceQuota field to given value.

### HasInstanceQuota

`func (o *EntitlementInput) HasInstanceQuota() bool`

HasInstanceQuota returns a boolean if a field has been set.

### GetProjectQuota

`func (o *EntitlementInput) GetProjectQuota() int32`

GetProjectQuota returns the ProjectQuota field if non-nil, zero value otherwise.

### GetProjectQuotaOk

`func (o *EntitlementInput) GetProjectQuotaOk() (*int32, bool)`

GetProjectQuotaOk returns a tuple with the ProjectQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectQuota

`func (o *EntitlementInput) SetProjectQuota(v int32)`

SetProjectQuota sets ProjectQuota field to given value.

### HasProjectQuota

`func (o *EntitlementInput) HasProjectQuota() bool`

HasProjectQuota returns a boolean if a field has been set.

### GetVolumeQuota

`func (o *EntitlementInput) GetVolumeQuota() map[string]interface{}`

GetVolumeQuota returns the VolumeQuota field if non-nil, zero value otherwise.

### GetVolumeQuotaOk

`func (o *EntitlementInput) GetVolumeQuotaOk() (*map[string]interface{}, bool)`

GetVolumeQuotaOk returns a tuple with the VolumeQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeQuota

`func (o *EntitlementInput) SetVolumeQuota(v map[string]interface{})`

SetVolumeQuota sets VolumeQuota field to given value.

### HasVolumeQuota

`func (o *EntitlementInput) HasVolumeQuota() bool`

HasVolumeQuota returns a boolean if a field has been set.

### GetFeatureAccess

`func (o *EntitlementInput) GetFeatureAccess() map[string]interface{}`

GetFeatureAccess returns the FeatureAccess field if non-nil, zero value otherwise.

### GetFeatureAccessOk

`func (o *EntitlementInput) GetFeatureAccessOk() (*map[string]interface{}, bool)`

GetFeatureAccessOk returns a tuple with the FeatureAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureAccess

`func (o *EntitlementInput) SetFeatureAccess(v map[string]interface{})`

SetFeatureAccess sets FeatureAccess field to given value.

### HasFeatureAccess

`func (o *EntitlementInput) HasFeatureAccess() bool`

HasFeatureAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


