# VolumeUpdateInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 
**Locked** | Pointer to **bool** |  | [optional] 
**BillingCycle** | Pointer to **string** | hourly | [optional] 
**Customdata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewVolumeUpdateInput

`func NewVolumeUpdateInput() *VolumeUpdateInput`

NewVolumeUpdateInput instantiates a new VolumeUpdateInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeUpdateInputWithDefaults

`func NewVolumeUpdateInputWithDefaults() *VolumeUpdateInput`

NewVolumeUpdateInputWithDefaults instantiates a new VolumeUpdateInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *VolumeUpdateInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VolumeUpdateInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VolumeUpdateInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VolumeUpdateInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSize

`func (o *VolumeUpdateInput) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *VolumeUpdateInput) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *VolumeUpdateInput) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *VolumeUpdateInput) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetLocked

`func (o *VolumeUpdateInput) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *VolumeUpdateInput) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *VolumeUpdateInput) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *VolumeUpdateInput) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetBillingCycle

`func (o *VolumeUpdateInput) GetBillingCycle() string`

GetBillingCycle returns the BillingCycle field if non-nil, zero value otherwise.

### GetBillingCycleOk

`func (o *VolumeUpdateInput) GetBillingCycleOk() (*string, bool)`

GetBillingCycleOk returns a tuple with the BillingCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCycle

`func (o *VolumeUpdateInput) SetBillingCycle(v string)`

SetBillingCycle sets BillingCycle field to given value.

### HasBillingCycle

`func (o *VolumeUpdateInput) HasBillingCycle() bool`

HasBillingCycle returns a boolean if a field has been set.

### GetCustomdata

`func (o *VolumeUpdateInput) GetCustomdata() map[string]interface{}`

GetCustomdata returns the Customdata field if non-nil, zero value otherwise.

### GetCustomdataOk

`func (o *VolumeUpdateInput) GetCustomdataOk() (*map[string]interface{}, bool)`

GetCustomdataOk returns a tuple with the Customdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomdata

`func (o *VolumeUpdateInput) SetCustomdata(v map[string]interface{})`

SetCustomdata sets Customdata field to given value.

### HasCustomdata

`func (o *VolumeUpdateInput) HasCustomdata() bool`

HasCustomdata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


