# VolumeAttachment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Volume** | Pointer to [**Href**](Href.md) |  | [optional] 
**Device** | Pointer to [**Href**](Href.md) |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 

## Methods

### NewVolumeAttachment

`func NewVolumeAttachment() *VolumeAttachment`

NewVolumeAttachment instantiates a new VolumeAttachment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeAttachmentWithDefaults

`func NewVolumeAttachmentWithDefaults() *VolumeAttachment`

NewVolumeAttachmentWithDefaults instantiates a new VolumeAttachment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VolumeAttachment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VolumeAttachment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VolumeAttachment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VolumeAttachment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *VolumeAttachment) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VolumeAttachment) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VolumeAttachment) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *VolumeAttachment) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetVolume

`func (o *VolumeAttachment) GetVolume() Href`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *VolumeAttachment) GetVolumeOk() (*Href, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *VolumeAttachment) SetVolume(v Href)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *VolumeAttachment) HasVolume() bool`

HasVolume returns a boolean if a field has been set.

### GetDevice

`func (o *VolumeAttachment) GetDevice() Href`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *VolumeAttachment) GetDeviceOk() (*Href, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *VolumeAttachment) SetDevice(v Href)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *VolumeAttachment) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetHref

`func (o *VolumeAttachment) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *VolumeAttachment) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *VolumeAttachment) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *VolumeAttachment) HasHref() bool`

HasHref returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


