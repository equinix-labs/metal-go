# Volume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 
**Locked** | Pointer to **bool** |  | [optional] 
**BillingCycle** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Project** | Pointer to [**Href**](Href.md) |  | [optional] 
**Facility** | Pointer to [**Href**](Href.md) |  | [optional] 
**SnapshotPolicies** | Pointer to [**[]Href**](Href.md) |  | [optional] 
**Attachments** | Pointer to [**[]Href**](Href.md) |  | [optional] 
**Plan** | Pointer to [**Plan**](Plan.md) |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**Customdata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewVolume

`func NewVolume() *Volume`

NewVolume instantiates a new Volume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeWithDefaults

`func NewVolumeWithDefaults() *Volume`

NewVolumeWithDefaults instantiates a new Volume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Volume) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Volume) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Volume) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Volume) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Volume) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Volume) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Volume) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Volume) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Volume) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Volume) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Volume) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Volume) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSize

`func (o *Volume) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Volume) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Volume) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *Volume) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetLocked

`func (o *Volume) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *Volume) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *Volume) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *Volume) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetBillingCycle

`func (o *Volume) GetBillingCycle() string`

GetBillingCycle returns the BillingCycle field if non-nil, zero value otherwise.

### GetBillingCycleOk

`func (o *Volume) GetBillingCycleOk() (*string, bool)`

GetBillingCycleOk returns a tuple with the BillingCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCycle

`func (o *Volume) SetBillingCycle(v string)`

SetBillingCycle sets BillingCycle field to given value.

### HasBillingCycle

`func (o *Volume) HasBillingCycle() bool`

HasBillingCycle returns a boolean if a field has been set.

### GetState

`func (o *Volume) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Volume) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Volume) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Volume) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Volume) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Volume) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Volume) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Volume) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Volume) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Volume) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Volume) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Volume) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetProject

`func (o *Volume) GetProject() Href`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *Volume) GetProjectOk() (*Href, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *Volume) SetProject(v Href)`

SetProject sets Project field to given value.

### HasProject

`func (o *Volume) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetFacility

`func (o *Volume) GetFacility() Href`

GetFacility returns the Facility field if non-nil, zero value otherwise.

### GetFacilityOk

`func (o *Volume) GetFacilityOk() (*Href, bool)`

GetFacilityOk returns a tuple with the Facility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacility

`func (o *Volume) SetFacility(v Href)`

SetFacility sets Facility field to given value.

### HasFacility

`func (o *Volume) HasFacility() bool`

HasFacility returns a boolean if a field has been set.

### GetSnapshotPolicies

`func (o *Volume) GetSnapshotPolicies() []Href`

GetSnapshotPolicies returns the SnapshotPolicies field if non-nil, zero value otherwise.

### GetSnapshotPoliciesOk

`func (o *Volume) GetSnapshotPoliciesOk() (*[]Href, bool)`

GetSnapshotPoliciesOk returns a tuple with the SnapshotPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotPolicies

`func (o *Volume) SetSnapshotPolicies(v []Href)`

SetSnapshotPolicies sets SnapshotPolicies field to given value.

### HasSnapshotPolicies

`func (o *Volume) HasSnapshotPolicies() bool`

HasSnapshotPolicies returns a boolean if a field has been set.

### GetAttachments

`func (o *Volume) GetAttachments() []Href`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *Volume) GetAttachmentsOk() (*[]Href, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *Volume) SetAttachments(v []Href)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *Volume) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetPlan

`func (o *Volume) GetPlan() Plan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *Volume) GetPlanOk() (*Plan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *Volume) SetPlan(v Plan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *Volume) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetHref

`func (o *Volume) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *Volume) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *Volume) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *Volume) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetCustomdata

`func (o *Volume) GetCustomdata() map[string]interface{}`

GetCustomdata returns the Customdata field if non-nil, zero value otherwise.

### GetCustomdataOk

`func (o *Volume) GetCustomdataOk() (*map[string]interface{}, bool)`

GetCustomdataOk returns a tuple with the Customdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomdata

`func (o *Volume) SetCustomdata(v map[string]interface{})`

SetCustomdata sets Customdata field to given value.

### HasCustomdata

`func (o *Volume) HasCustomdata() bool`

HasCustomdata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


