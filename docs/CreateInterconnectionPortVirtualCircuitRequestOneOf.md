# CreateInterconnectionPortVirtualCircuitRequestOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NniVlan** | Pointer to **int32** |  | [optional] 
**Project** | Pointer to **string** |  | [optional] 
**Speed** | Pointer to **int32** | speed can be passed as integer number representing bps speed or string (e.g. &#39;52m&#39; or &#39;100g&#39; or &#39;4 gbps&#39;) | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Vnid** | Pointer to **string** | A Virtual Network record UUID or the VNID of a Metro Virtual Network in your project (sent as integer). | [optional] 

## Methods

### NewCreateInterconnectionPortVirtualCircuitRequestOneOf

`func NewCreateInterconnectionPortVirtualCircuitRequestOneOf() *CreateInterconnectionPortVirtualCircuitRequestOneOf`

NewCreateInterconnectionPortVirtualCircuitRequestOneOf instantiates a new CreateInterconnectionPortVirtualCircuitRequestOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateInterconnectionPortVirtualCircuitRequestOneOfWithDefaults

`func NewCreateInterconnectionPortVirtualCircuitRequestOneOfWithDefaults() *CreateInterconnectionPortVirtualCircuitRequestOneOf`

NewCreateInterconnectionPortVirtualCircuitRequestOneOfWithDefaults instantiates a new CreateInterconnectionPortVirtualCircuitRequestOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateInterconnectionPortVirtualCircuitRequestOneOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateInterconnectionPortVirtualCircuitRequestOneOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateInterconnectionPortVirtualCircuitRequestOneOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateInterconnectionPortVirtualCircuitRequestOneOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *CreateInterconnectionPortVirtualCircuitRequestOneOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateInterconnectionPortVirtualCircuitRequestOneOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateInterconnectionPortVirtualCircuitRequestOneOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateInterconnectionPortVirtualCircuitRequestOneOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNniVlan

`func (o *CreateInterconnectionPortVirtualCircuitRequestOneOf) GetNniVlan() int32`

GetNniVlan returns the NniVlan field if non-nil, zero value otherwise.

### GetNniVlanOk

`func (o *CreateInterconnectionPortVirtualCircuitRequestOneOf) GetNniVlanOk() (*int32, bool)`

GetNniVlanOk returns a tuple with the NniVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNniVlan

`func (o *CreateInterconnectionPortVirtualCircuitRequestOneOf) SetNniVlan(v int32)`

SetNniVlan sets NniVlan field to given value.

### HasNniVlan

`func (o *CreateInterconnectionPortVirtualCircuitRequestOneOf) HasNniVlan() bool`

HasNniVlan returns a boolean if a field has been set.

### GetProject

`func (o *CreateInterconnectionPortVirtualCircuitRequestOneOf) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *CreateInterconnectionPortVirtualCircuitRequestOneOf) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *CreateInterconnectionPortVirtualCircuitRequestOneOf) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *CreateInterconnectionPortVirtualCircuitRequestOneOf) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetSpeed

`func (o *CreateInterconnectionPortVirtualCircuitRequestOneOf) GetSpeed() int32`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *CreateInterconnectionPortVirtualCircuitRequestOneOf) GetSpeedOk() (*int32, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *CreateInterconnectionPortVirtualCircuitRequestOneOf) SetSpeed(v int32)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *CreateInterconnectionPortVirtualCircuitRequestOneOf) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetTags

`func (o *CreateInterconnectionPortVirtualCircuitRequestOneOf) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateInterconnectionPortVirtualCircuitRequestOneOf) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateInterconnectionPortVirtualCircuitRequestOneOf) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateInterconnectionPortVirtualCircuitRequestOneOf) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVnid

`func (o *CreateInterconnectionPortVirtualCircuitRequestOneOf) GetVnid() string`

GetVnid returns the Vnid field if non-nil, zero value otherwise.

### GetVnidOk

`func (o *CreateInterconnectionPortVirtualCircuitRequestOneOf) GetVnidOk() (*string, bool)`

GetVnidOk returns a tuple with the Vnid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnid

`func (o *CreateInterconnectionPortVirtualCircuitRequestOneOf) SetVnid(v string)`

SetVnid sets Vnid field to given value.

### HasVnid

`func (o *CreateInterconnectionPortVirtualCircuitRequestOneOf) HasVnid() bool`

HasVnid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


