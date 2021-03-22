# VirtualNetworkCreateInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Facility** | Pointer to **string** | The UUID (or facility code) for the Facility in which to create this Virtual network. | [optional] 

## Methods

### NewVirtualNetworkCreateInput

`func NewVirtualNetworkCreateInput() *VirtualNetworkCreateInput`

NewVirtualNetworkCreateInput instantiates a new VirtualNetworkCreateInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualNetworkCreateInputWithDefaults

`func NewVirtualNetworkCreateInputWithDefaults() *VirtualNetworkCreateInput`

NewVirtualNetworkCreateInputWithDefaults instantiates a new VirtualNetworkCreateInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *VirtualNetworkCreateInput) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *VirtualNetworkCreateInput) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *VirtualNetworkCreateInput) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *VirtualNetworkCreateInput) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetDescription

`func (o *VirtualNetworkCreateInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VirtualNetworkCreateInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VirtualNetworkCreateInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VirtualNetworkCreateInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFacility

`func (o *VirtualNetworkCreateInput) GetFacility() string`

GetFacility returns the Facility field if non-nil, zero value otherwise.

### GetFacilityOk

`func (o *VirtualNetworkCreateInput) GetFacilityOk() (*string, bool)`

GetFacilityOk returns a tuple with the Facility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacility

`func (o *VirtualNetworkCreateInput) SetFacility(v string)`

SetFacility sets Facility field to given value.

### HasFacility

`func (o *VirtualNetworkCreateInput) HasFacility() bool`

HasFacility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


