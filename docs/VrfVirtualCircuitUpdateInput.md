# VrfVirtualCircuitUpdateInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Speed** | Pointer to **string** | Speed can be changed only if it is an interconnection on a Dedicated Port | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewVrfVirtualCircuitUpdateInput

`func NewVrfVirtualCircuitUpdateInput() *VrfVirtualCircuitUpdateInput`

NewVrfVirtualCircuitUpdateInput instantiates a new VrfVirtualCircuitUpdateInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVrfVirtualCircuitUpdateInputWithDefaults

`func NewVrfVirtualCircuitUpdateInputWithDefaults() *VrfVirtualCircuitUpdateInput`

NewVrfVirtualCircuitUpdateInputWithDefaults instantiates a new VrfVirtualCircuitUpdateInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *VrfVirtualCircuitUpdateInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VrfVirtualCircuitUpdateInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VrfVirtualCircuitUpdateInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VrfVirtualCircuitUpdateInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *VrfVirtualCircuitUpdateInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VrfVirtualCircuitUpdateInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VrfVirtualCircuitUpdateInput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VrfVirtualCircuitUpdateInput) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSpeed

`func (o *VrfVirtualCircuitUpdateInput) GetSpeed() string`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *VrfVirtualCircuitUpdateInput) GetSpeedOk() (*string, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *VrfVirtualCircuitUpdateInput) SetSpeed(v string)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *VrfVirtualCircuitUpdateInput) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetTags

`func (o *VrfVirtualCircuitUpdateInput) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VrfVirtualCircuitUpdateInput) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VrfVirtualCircuitUpdateInput) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VrfVirtualCircuitUpdateInput) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


