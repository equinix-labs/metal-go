# VolumeList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Volumes** | Pointer to [**[]Volume**](Volume.md) |  | [optional] 
**Meta** | Pointer to [**Meta**](Meta.md) |  | [optional] 

## Methods

### NewVolumeList

`func NewVolumeList() *VolumeList`

NewVolumeList instantiates a new VolumeList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeListWithDefaults

`func NewVolumeListWithDefaults() *VolumeList`

NewVolumeListWithDefaults instantiates a new VolumeList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVolumes

`func (o *VolumeList) GetVolumes() []Volume`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *VolumeList) GetVolumesOk() (*[]Volume, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *VolumeList) SetVolumes(v []Volume)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *VolumeList) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetMeta

`func (o *VolumeList) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *VolumeList) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *VolumeList) SetMeta(v Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *VolumeList) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


