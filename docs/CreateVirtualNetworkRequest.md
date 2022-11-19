# CreateVirtualNetworkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Facility** | Pointer to **string** | The UUID (or facility code) for the Facility in which to create this Virtual network. | [optional] 
**Metro** | Pointer to **string** | The UUID (or metro code) for the Metro in which to create this Virtual Network. | [optional] 
**Vxlan** | Pointer to **int32** | VLAN ID between 2-3999. Must be unique for the project within the Metro in which this Virtual Network is being created. If no value is specified, the next-available VLAN ID in the range 1000-1999 will be automatically selected. | [optional] 

## Methods

### NewCreateVirtualNetworkRequest

`func NewCreateVirtualNetworkRequest() *CreateVirtualNetworkRequest`

NewCreateVirtualNetworkRequest instantiates a new CreateVirtualNetworkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVirtualNetworkRequestWithDefaults

`func NewCreateVirtualNetworkRequestWithDefaults() *CreateVirtualNetworkRequest`

NewCreateVirtualNetworkRequestWithDefaults instantiates a new CreateVirtualNetworkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateVirtualNetworkRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateVirtualNetworkRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateVirtualNetworkRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateVirtualNetworkRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFacility

`func (o *CreateVirtualNetworkRequest) GetFacility() string`

GetFacility returns the Facility field if non-nil, zero value otherwise.

### GetFacilityOk

`func (o *CreateVirtualNetworkRequest) GetFacilityOk() (*string, bool)`

GetFacilityOk returns a tuple with the Facility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacility

`func (o *CreateVirtualNetworkRequest) SetFacility(v string)`

SetFacility sets Facility field to given value.

### HasFacility

`func (o *CreateVirtualNetworkRequest) HasFacility() bool`

HasFacility returns a boolean if a field has been set.

### GetMetro

`func (o *CreateVirtualNetworkRequest) GetMetro() string`

GetMetro returns the Metro field if non-nil, zero value otherwise.

### GetMetroOk

`func (o *CreateVirtualNetworkRequest) GetMetroOk() (*string, bool)`

GetMetroOk returns a tuple with the Metro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetro

`func (o *CreateVirtualNetworkRequest) SetMetro(v string)`

SetMetro sets Metro field to given value.

### HasMetro

`func (o *CreateVirtualNetworkRequest) HasMetro() bool`

HasMetro returns a boolean if a field has been set.

### GetVxlan

`func (o *CreateVirtualNetworkRequest) GetVxlan() int32`

GetVxlan returns the Vxlan field if non-nil, zero value otherwise.

### GetVxlanOk

`func (o *CreateVirtualNetworkRequest) GetVxlanOk() (*int32, bool)`

GetVxlanOk returns a tuple with the Vxlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVxlan

`func (o *CreateVirtualNetworkRequest) SetVxlan(v int32)`

SetVxlan sets Vxlan field to given value.

### HasVxlan

`func (o *CreateVirtualNetworkRequest) HasVxlan() bool`

HasVxlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


