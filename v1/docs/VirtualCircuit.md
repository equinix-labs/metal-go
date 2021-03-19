# VirtualCircuit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Vnid** | Pointer to **int32** |  | [optional] 
**NniVlan** | Pointer to **int32** |  | [optional] 
**Speed** | Pointer to **int32** | integer representing bps speed | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Project** | Pointer to [**Href**](Href.md) |  | [optional] 
**VirtualNetwork** | Pointer to [**Href**](Href.md) |  | [optional] 

## Methods

### NewVirtualCircuit

`func NewVirtualCircuit() *VirtualCircuit`

NewVirtualCircuit instantiates a new VirtualCircuit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualCircuitWithDefaults

`func NewVirtualCircuitWithDefaults() *VirtualCircuit`

NewVirtualCircuitWithDefaults instantiates a new VirtualCircuit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VirtualCircuit) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VirtualCircuit) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VirtualCircuit) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VirtualCircuit) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *VirtualCircuit) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualCircuit) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualCircuit) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VirtualCircuit) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *VirtualCircuit) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VirtualCircuit) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VirtualCircuit) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VirtualCircuit) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *VirtualCircuit) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VirtualCircuit) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VirtualCircuit) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VirtualCircuit) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVnid

`func (o *VirtualCircuit) GetVnid() int32`

GetVnid returns the Vnid field if non-nil, zero value otherwise.

### GetVnidOk

`func (o *VirtualCircuit) GetVnidOk() (*int32, bool)`

GetVnidOk returns a tuple with the Vnid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnid

`func (o *VirtualCircuit) SetVnid(v int32)`

SetVnid sets Vnid field to given value.

### HasVnid

`func (o *VirtualCircuit) HasVnid() bool`

HasVnid returns a boolean if a field has been set.

### GetNniVlan

`func (o *VirtualCircuit) GetNniVlan() int32`

GetNniVlan returns the NniVlan field if non-nil, zero value otherwise.

### GetNniVlanOk

`func (o *VirtualCircuit) GetNniVlanOk() (*int32, bool)`

GetNniVlanOk returns a tuple with the NniVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNniVlan

`func (o *VirtualCircuit) SetNniVlan(v int32)`

SetNniVlan sets NniVlan field to given value.

### HasNniVlan

`func (o *VirtualCircuit) HasNniVlan() bool`

HasNniVlan returns a boolean if a field has been set.

### GetSpeed

`func (o *VirtualCircuit) GetSpeed() int32`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *VirtualCircuit) GetSpeedOk() (*int32, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *VirtualCircuit) SetSpeed(v int32)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *VirtualCircuit) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetTags

`func (o *VirtualCircuit) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VirtualCircuit) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VirtualCircuit) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VirtualCircuit) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetProject

`func (o *VirtualCircuit) GetProject() Href`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *VirtualCircuit) GetProjectOk() (*Href, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *VirtualCircuit) SetProject(v Href)`

SetProject sets Project field to given value.

### HasProject

`func (o *VirtualCircuit) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetVirtualNetwork

`func (o *VirtualCircuit) GetVirtualNetwork() Href`

GetVirtualNetwork returns the VirtualNetwork field if non-nil, zero value otherwise.

### GetVirtualNetworkOk

`func (o *VirtualCircuit) GetVirtualNetworkOk() (*Href, bool)`

GetVirtualNetworkOk returns a tuple with the VirtualNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualNetwork

`func (o *VirtualCircuit) SetVirtualNetwork(v Href)`

SetVirtualNetwork sets VirtualNetwork field to given value.

### HasVirtualNetwork

`func (o *VirtualCircuit) HasVirtualNetwork() bool`

HasVirtualNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


