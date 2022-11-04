# VirtualCircuitCreateInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NniVlan** | Pointer to **int32** |  | [optional] 
**Project** | Pointer to **string** |  | [optional] 
**Speed** | Pointer to **int32** | speed can be passed as integer number representing bps speed or string (e.g. &#39;52m&#39; or &#39;100g&#39; or &#39;4 gbps&#39;) | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Vnid** | Pointer to **string** | A Virtual Network record UUID or the VNID of a Virtual Network in your project (sent as integer). | [optional] 

## Methods

### NewVirtualCircuitCreateInput

`func NewVirtualCircuitCreateInput() *VirtualCircuitCreateInput`

NewVirtualCircuitCreateInput instantiates a new VirtualCircuitCreateInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualCircuitCreateInputWithDefaults

`func NewVirtualCircuitCreateInputWithDefaults() *VirtualCircuitCreateInput`

NewVirtualCircuitCreateInputWithDefaults instantiates a new VirtualCircuitCreateInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *VirtualCircuitCreateInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VirtualCircuitCreateInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VirtualCircuitCreateInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VirtualCircuitCreateInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *VirtualCircuitCreateInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualCircuitCreateInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualCircuitCreateInput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VirtualCircuitCreateInput) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNniVlan

`func (o *VirtualCircuitCreateInput) GetNniVlan() int32`

GetNniVlan returns the NniVlan field if non-nil, zero value otherwise.

### GetNniVlanOk

`func (o *VirtualCircuitCreateInput) GetNniVlanOk() (*int32, bool)`

GetNniVlanOk returns a tuple with the NniVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNniVlan

`func (o *VirtualCircuitCreateInput) SetNniVlan(v int32)`

SetNniVlan sets NniVlan field to given value.

### HasNniVlan

`func (o *VirtualCircuitCreateInput) HasNniVlan() bool`

HasNniVlan returns a boolean if a field has been set.

### GetProject

`func (o *VirtualCircuitCreateInput) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *VirtualCircuitCreateInput) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *VirtualCircuitCreateInput) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *VirtualCircuitCreateInput) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetSpeed

`func (o *VirtualCircuitCreateInput) GetSpeed() int32`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *VirtualCircuitCreateInput) GetSpeedOk() (*int32, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *VirtualCircuitCreateInput) SetSpeed(v int32)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *VirtualCircuitCreateInput) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetTags

`func (o *VirtualCircuitCreateInput) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VirtualCircuitCreateInput) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VirtualCircuitCreateInput) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VirtualCircuitCreateInput) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVnid

`func (o *VirtualCircuitCreateInput) GetVnid() string`

GetVnid returns the Vnid field if non-nil, zero value otherwise.

### GetVnidOk

`func (o *VirtualCircuitCreateInput) GetVnidOk() (*string, bool)`

GetVnidOk returns a tuple with the Vnid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnid

`func (o *VirtualCircuitCreateInput) SetVnid(v string)`

SetVnid sets Vnid field to given value.

### HasVnid

`func (o *VirtualCircuitCreateInput) HasVnid() bool`

HasVnid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


