# UpdateVirtualCircuitRequestOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Speed** | Pointer to **string** | Speed can be changed only if it is an interconnection on a Dedicated Port | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Vnid** | Pointer to **string** | A Virtual Network record UUID or the VNID of a Virtual Network in your project. | [optional] 

## Methods

### NewUpdateVirtualCircuitRequestOneOf

`func NewUpdateVirtualCircuitRequestOneOf() *UpdateVirtualCircuitRequestOneOf`

NewUpdateVirtualCircuitRequestOneOf instantiates a new UpdateVirtualCircuitRequestOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateVirtualCircuitRequestOneOfWithDefaults

`func NewUpdateVirtualCircuitRequestOneOfWithDefaults() *UpdateVirtualCircuitRequestOneOf`

NewUpdateVirtualCircuitRequestOneOfWithDefaults instantiates a new UpdateVirtualCircuitRequestOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UpdateVirtualCircuitRequestOneOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateVirtualCircuitRequestOneOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateVirtualCircuitRequestOneOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateVirtualCircuitRequestOneOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *UpdateVirtualCircuitRequestOneOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateVirtualCircuitRequestOneOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateVirtualCircuitRequestOneOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateVirtualCircuitRequestOneOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSpeed

`func (o *UpdateVirtualCircuitRequestOneOf) GetSpeed() string`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *UpdateVirtualCircuitRequestOneOf) GetSpeedOk() (*string, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *UpdateVirtualCircuitRequestOneOf) SetSpeed(v string)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *UpdateVirtualCircuitRequestOneOf) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetTags

`func (o *UpdateVirtualCircuitRequestOneOf) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateVirtualCircuitRequestOneOf) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateVirtualCircuitRequestOneOf) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateVirtualCircuitRequestOneOf) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVnid

`func (o *UpdateVirtualCircuitRequestOneOf) GetVnid() string`

GetVnid returns the Vnid field if non-nil, zero value otherwise.

### GetVnidOk

`func (o *UpdateVirtualCircuitRequestOneOf) GetVnidOk() (*string, bool)`

GetVnidOk returns a tuple with the Vnid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnid

`func (o *UpdateVirtualCircuitRequestOneOf) SetVnid(v string)`

SetVnid sets Vnid field to given value.

### HasVnid

`func (o *UpdateVirtualCircuitRequestOneOf) HasVnid() bool`

HasVnid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


