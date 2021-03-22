# OperatingSystem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Distro** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**ProvisionableOn** | Pointer to **[]string** |  | [optional] 

## Methods

### NewOperatingSystem

`func NewOperatingSystem() *OperatingSystem`

NewOperatingSystem instantiates a new OperatingSystem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperatingSystemWithDefaults

`func NewOperatingSystemWithDefaults() *OperatingSystem`

NewOperatingSystemWithDefaults instantiates a new OperatingSystem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OperatingSystem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OperatingSystem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OperatingSystem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OperatingSystem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSlug

`func (o *OperatingSystem) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *OperatingSystem) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *OperatingSystem) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *OperatingSystem) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetName

`func (o *OperatingSystem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OperatingSystem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OperatingSystem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OperatingSystem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDistro

`func (o *OperatingSystem) GetDistro() string`

GetDistro returns the Distro field if non-nil, zero value otherwise.

### GetDistroOk

`func (o *OperatingSystem) GetDistroOk() (*string, bool)`

GetDistroOk returns a tuple with the Distro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistro

`func (o *OperatingSystem) SetDistro(v string)`

SetDistro sets Distro field to given value.

### HasDistro

`func (o *OperatingSystem) HasDistro() bool`

HasDistro returns a boolean if a field has been set.

### GetVersion

`func (o *OperatingSystem) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *OperatingSystem) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *OperatingSystem) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *OperatingSystem) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetProvisionableOn

`func (o *OperatingSystem) GetProvisionableOn() []string`

GetProvisionableOn returns the ProvisionableOn field if non-nil, zero value otherwise.

### GetProvisionableOnOk

`func (o *OperatingSystem) GetProvisionableOnOk() (*[]string, bool)`

GetProvisionableOnOk returns a tuple with the ProvisionableOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionableOn

`func (o *OperatingSystem) SetProvisionableOn(v []string)`

SetProvisionableOn sets ProvisionableOn field to given value.

### HasProvisionableOn

`func (o *OperatingSystem) HasProvisionableOn() bool`

HasProvisionableOn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


