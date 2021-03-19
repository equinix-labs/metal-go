# VolumeCreateInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Facility** | **string** | ams1, ewr1, nrt1, sjc1 | 
**Plan** | **string** | storage_1, storage_2 | 
**Size** | **int32** |  | 
**Locked** | Pointer to **bool** |  | [optional] 
**BillingCycle** | Pointer to **string** | hourly | [optional] 
**SnapshotPolicies** | Pointer to [**SnapshotPolicyInput**](SnapshotPolicyInput.md) |  | [optional] 
**Customdata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewVolumeCreateInput

`func NewVolumeCreateInput(facility string, plan string, size int32, ) *VolumeCreateInput`

NewVolumeCreateInput instantiates a new VolumeCreateInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeCreateInputWithDefaults

`func NewVolumeCreateInputWithDefaults() *VolumeCreateInput`

NewVolumeCreateInputWithDefaults instantiates a new VolumeCreateInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *VolumeCreateInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VolumeCreateInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VolumeCreateInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VolumeCreateInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFacility

`func (o *VolumeCreateInput) GetFacility() string`

GetFacility returns the Facility field if non-nil, zero value otherwise.

### GetFacilityOk

`func (o *VolumeCreateInput) GetFacilityOk() (*string, bool)`

GetFacilityOk returns a tuple with the Facility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacility

`func (o *VolumeCreateInput) SetFacility(v string)`

SetFacility sets Facility field to given value.


### GetPlan

`func (o *VolumeCreateInput) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *VolumeCreateInput) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *VolumeCreateInput) SetPlan(v string)`

SetPlan sets Plan field to given value.


### GetSize

`func (o *VolumeCreateInput) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *VolumeCreateInput) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *VolumeCreateInput) SetSize(v int32)`

SetSize sets Size field to given value.


### GetLocked

`func (o *VolumeCreateInput) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *VolumeCreateInput) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *VolumeCreateInput) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *VolumeCreateInput) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetBillingCycle

`func (o *VolumeCreateInput) GetBillingCycle() string`

GetBillingCycle returns the BillingCycle field if non-nil, zero value otherwise.

### GetBillingCycleOk

`func (o *VolumeCreateInput) GetBillingCycleOk() (*string, bool)`

GetBillingCycleOk returns a tuple with the BillingCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCycle

`func (o *VolumeCreateInput) SetBillingCycle(v string)`

SetBillingCycle sets BillingCycle field to given value.

### HasBillingCycle

`func (o *VolumeCreateInput) HasBillingCycle() bool`

HasBillingCycle returns a boolean if a field has been set.

### GetSnapshotPolicies

`func (o *VolumeCreateInput) GetSnapshotPolicies() SnapshotPolicyInput`

GetSnapshotPolicies returns the SnapshotPolicies field if non-nil, zero value otherwise.

### GetSnapshotPoliciesOk

`func (o *VolumeCreateInput) GetSnapshotPoliciesOk() (*SnapshotPolicyInput, bool)`

GetSnapshotPoliciesOk returns a tuple with the SnapshotPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotPolicies

`func (o *VolumeCreateInput) SetSnapshotPolicies(v SnapshotPolicyInput)`

SetSnapshotPolicies sets SnapshotPolicies field to given value.

### HasSnapshotPolicies

`func (o *VolumeCreateInput) HasSnapshotPolicies() bool`

HasSnapshotPolicies returns a boolean if a field has been set.

### GetCustomdata

`func (o *VolumeCreateInput) GetCustomdata() map[string]interface{}`

GetCustomdata returns the Customdata field if non-nil, zero value otherwise.

### GetCustomdataOk

`func (o *VolumeCreateInput) GetCustomdataOk() (*map[string]interface{}, bool)`

GetCustomdataOk returns a tuple with the Customdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomdata

`func (o *VolumeCreateInput) SetCustomdata(v map[string]interface{})`

SetCustomdata sets Customdata field to given value.

### HasCustomdata

`func (o *VolumeCreateInput) HasCustomdata() bool`

HasCustomdata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


