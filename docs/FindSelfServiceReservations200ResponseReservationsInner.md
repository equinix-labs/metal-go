# FindSelfServiceReservations200ResponseReservationsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Item** | Pointer to [**[]FindSelfServiceReservations200ResponseReservationsInnerItemInner**](FindSelfServiceReservations200ResponseReservationsInnerItemInner.md) |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**Organization** | Pointer to **string** |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**Period** | Pointer to [**FindSelfServiceReservations200ResponseReservationsInnerPeriod**](FindSelfServiceReservations200ResponseReservationsInnerPeriod.md) |  | [optional] 
**Project** | Pointer to **string** |  | [optional] 
**ProjectId** | Pointer to **string** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TotalCost** | Pointer to **int32** |  | [optional] 

## Methods

### NewFindSelfServiceReservations200ResponseReservationsInner

`func NewFindSelfServiceReservations200ResponseReservationsInner() *FindSelfServiceReservations200ResponseReservationsInner`

NewFindSelfServiceReservations200ResponseReservationsInner instantiates a new FindSelfServiceReservations200ResponseReservationsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindSelfServiceReservations200ResponseReservationsInnerWithDefaults

`func NewFindSelfServiceReservations200ResponseReservationsInnerWithDefaults() *FindSelfServiceReservations200ResponseReservationsInner`

NewFindSelfServiceReservations200ResponseReservationsInnerWithDefaults instantiates a new FindSelfServiceReservations200ResponseReservationsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *FindSelfServiceReservations200ResponseReservationsInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FindSelfServiceReservations200ResponseReservationsInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FindSelfServiceReservations200ResponseReservationsInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FindSelfServiceReservations200ResponseReservationsInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetItem

`func (o *FindSelfServiceReservations200ResponseReservationsInner) GetItem() []FindSelfServiceReservations200ResponseReservationsInnerItemInner`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *FindSelfServiceReservations200ResponseReservationsInner) GetItemOk() (*[]FindSelfServiceReservations200ResponseReservationsInnerItemInner, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *FindSelfServiceReservations200ResponseReservationsInner) SetItem(v []FindSelfServiceReservations200ResponseReservationsInnerItemInner)`

SetItem sets Item field to given value.

### HasItem

`func (o *FindSelfServiceReservations200ResponseReservationsInner) HasItem() bool`

HasItem returns a boolean if a field has been set.

### GetNotes

`func (o *FindSelfServiceReservations200ResponseReservationsInner) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *FindSelfServiceReservations200ResponseReservationsInner) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *FindSelfServiceReservations200ResponseReservationsInner) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *FindSelfServiceReservations200ResponseReservationsInner) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetOrganization

`func (o *FindSelfServiceReservations200ResponseReservationsInner) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *FindSelfServiceReservations200ResponseReservationsInner) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *FindSelfServiceReservations200ResponseReservationsInner) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *FindSelfServiceReservations200ResponseReservationsInner) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetOrganizationId

`func (o *FindSelfServiceReservations200ResponseReservationsInner) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *FindSelfServiceReservations200ResponseReservationsInner) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *FindSelfServiceReservations200ResponseReservationsInner) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *FindSelfServiceReservations200ResponseReservationsInner) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetPeriod

`func (o *FindSelfServiceReservations200ResponseReservationsInner) GetPeriod() FindSelfServiceReservations200ResponseReservationsInnerPeriod`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *FindSelfServiceReservations200ResponseReservationsInner) GetPeriodOk() (*FindSelfServiceReservations200ResponseReservationsInnerPeriod, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *FindSelfServiceReservations200ResponseReservationsInner) SetPeriod(v FindSelfServiceReservations200ResponseReservationsInnerPeriod)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *FindSelfServiceReservations200ResponseReservationsInner) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetProject

`func (o *FindSelfServiceReservations200ResponseReservationsInner) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *FindSelfServiceReservations200ResponseReservationsInner) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *FindSelfServiceReservations200ResponseReservationsInner) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *FindSelfServiceReservations200ResponseReservationsInner) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetProjectId

`func (o *FindSelfServiceReservations200ResponseReservationsInner) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *FindSelfServiceReservations200ResponseReservationsInner) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *FindSelfServiceReservations200ResponseReservationsInner) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *FindSelfServiceReservations200ResponseReservationsInner) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetStartDate

`func (o *FindSelfServiceReservations200ResponseReservationsInner) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *FindSelfServiceReservations200ResponseReservationsInner) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *FindSelfServiceReservations200ResponseReservationsInner) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *FindSelfServiceReservations200ResponseReservationsInner) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetStatus

`func (o *FindSelfServiceReservations200ResponseReservationsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FindSelfServiceReservations200ResponseReservationsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FindSelfServiceReservations200ResponseReservationsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FindSelfServiceReservations200ResponseReservationsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTotalCost

`func (o *FindSelfServiceReservations200ResponseReservationsInner) GetTotalCost() int32`

GetTotalCost returns the TotalCost field if non-nil, zero value otherwise.

### GetTotalCostOk

`func (o *FindSelfServiceReservations200ResponseReservationsInner) GetTotalCostOk() (*int32, bool)`

GetTotalCostOk returns a tuple with the TotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCost

`func (o *FindSelfServiceReservations200ResponseReservationsInner) SetTotalCost(v int32)`

SetTotalCost sets TotalCost field to given value.

### HasTotalCost

`func (o *FindSelfServiceReservations200ResponseReservationsInner) HasTotalCost() bool`

HasTotalCost returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


