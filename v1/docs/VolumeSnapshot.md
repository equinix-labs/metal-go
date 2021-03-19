# VolumeSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Volume** | Pointer to [**Href**](Href.md) |  | [optional] 

## Methods

### NewVolumeSnapshot

`func NewVolumeSnapshot() *VolumeSnapshot`

NewVolumeSnapshot instantiates a new VolumeSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeSnapshotWithDefaults

`func NewVolumeSnapshotWithDefaults() *VolumeSnapshot`

NewVolumeSnapshotWithDefaults instantiates a new VolumeSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VolumeSnapshot) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VolumeSnapshot) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VolumeSnapshot) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VolumeSnapshot) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *VolumeSnapshot) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VolumeSnapshot) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VolumeSnapshot) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VolumeSnapshot) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *VolumeSnapshot) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VolumeSnapshot) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VolumeSnapshot) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *VolumeSnapshot) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetTimestamp

`func (o *VolumeSnapshot) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *VolumeSnapshot) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *VolumeSnapshot) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *VolumeSnapshot) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetVolume

`func (o *VolumeSnapshot) GetVolume() Href`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *VolumeSnapshot) GetVolumeOk() (*Href, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *VolumeSnapshot) SetVolume(v Href)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *VolumeSnapshot) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


