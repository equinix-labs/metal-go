# CreateSpotMarketRequestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DevicesMax** | Pointer to **int32** |  | [optional] 
**DevicesMin** | Pointer to **int32** |  | [optional] 
**EndAt** | Pointer to **time.Time** |  | [optional] 
**Facilities** | Pointer to **[]string** |  | [optional] 
**InstanceAttributes** | Pointer to [**CreateSpotMarketRequestRequestInstanceAttributes**](CreateSpotMarketRequestRequestInstanceAttributes.md) |  | [optional] 
**MaxBidPrice** | Pointer to **float32** |  | [optional] 
**Metro** | Pointer to **string** | The metro ID or code the spot market request will be created in. | [optional] 

## Methods

### NewCreateSpotMarketRequestRequest

`func NewCreateSpotMarketRequestRequest() *CreateSpotMarketRequestRequest`

NewCreateSpotMarketRequestRequest instantiates a new CreateSpotMarketRequestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSpotMarketRequestRequestWithDefaults

`func NewCreateSpotMarketRequestRequestWithDefaults() *CreateSpotMarketRequestRequest`

NewCreateSpotMarketRequestRequestWithDefaults instantiates a new CreateSpotMarketRequestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevicesMax

`func (o *CreateSpotMarketRequestRequest) GetDevicesMax() int32`

GetDevicesMax returns the DevicesMax field if non-nil, zero value otherwise.

### GetDevicesMaxOk

`func (o *CreateSpotMarketRequestRequest) GetDevicesMaxOk() (*int32, bool)`

GetDevicesMaxOk returns a tuple with the DevicesMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicesMax

`func (o *CreateSpotMarketRequestRequest) SetDevicesMax(v int32)`

SetDevicesMax sets DevicesMax field to given value.

### HasDevicesMax

`func (o *CreateSpotMarketRequestRequest) HasDevicesMax() bool`

HasDevicesMax returns a boolean if a field has been set.

### GetDevicesMin

`func (o *CreateSpotMarketRequestRequest) GetDevicesMin() int32`

GetDevicesMin returns the DevicesMin field if non-nil, zero value otherwise.

### GetDevicesMinOk

`func (o *CreateSpotMarketRequestRequest) GetDevicesMinOk() (*int32, bool)`

GetDevicesMinOk returns a tuple with the DevicesMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicesMin

`func (o *CreateSpotMarketRequestRequest) SetDevicesMin(v int32)`

SetDevicesMin sets DevicesMin field to given value.

### HasDevicesMin

`func (o *CreateSpotMarketRequestRequest) HasDevicesMin() bool`

HasDevicesMin returns a boolean if a field has been set.

### GetEndAt

`func (o *CreateSpotMarketRequestRequest) GetEndAt() time.Time`

GetEndAt returns the EndAt field if non-nil, zero value otherwise.

### GetEndAtOk

`func (o *CreateSpotMarketRequestRequest) GetEndAtOk() (*time.Time, bool)`

GetEndAtOk returns a tuple with the EndAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAt

`func (o *CreateSpotMarketRequestRequest) SetEndAt(v time.Time)`

SetEndAt sets EndAt field to given value.

### HasEndAt

`func (o *CreateSpotMarketRequestRequest) HasEndAt() bool`

HasEndAt returns a boolean if a field has been set.

### GetFacilities

`func (o *CreateSpotMarketRequestRequest) GetFacilities() []string`

GetFacilities returns the Facilities field if non-nil, zero value otherwise.

### GetFacilitiesOk

`func (o *CreateSpotMarketRequestRequest) GetFacilitiesOk() (*[]string, bool)`

GetFacilitiesOk returns a tuple with the Facilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilities

`func (o *CreateSpotMarketRequestRequest) SetFacilities(v []string)`

SetFacilities sets Facilities field to given value.

### HasFacilities

`func (o *CreateSpotMarketRequestRequest) HasFacilities() bool`

HasFacilities returns a boolean if a field has been set.

### GetInstanceAttributes

`func (o *CreateSpotMarketRequestRequest) GetInstanceAttributes() CreateSpotMarketRequestRequestInstanceAttributes`

GetInstanceAttributes returns the InstanceAttributes field if non-nil, zero value otherwise.

### GetInstanceAttributesOk

`func (o *CreateSpotMarketRequestRequest) GetInstanceAttributesOk() (*CreateSpotMarketRequestRequestInstanceAttributes, bool)`

GetInstanceAttributesOk returns a tuple with the InstanceAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceAttributes

`func (o *CreateSpotMarketRequestRequest) SetInstanceAttributes(v CreateSpotMarketRequestRequestInstanceAttributes)`

SetInstanceAttributes sets InstanceAttributes field to given value.

### HasInstanceAttributes

`func (o *CreateSpotMarketRequestRequest) HasInstanceAttributes() bool`

HasInstanceAttributes returns a boolean if a field has been set.

### GetMaxBidPrice

`func (o *CreateSpotMarketRequestRequest) GetMaxBidPrice() float32`

GetMaxBidPrice returns the MaxBidPrice field if non-nil, zero value otherwise.

### GetMaxBidPriceOk

`func (o *CreateSpotMarketRequestRequest) GetMaxBidPriceOk() (*float32, bool)`

GetMaxBidPriceOk returns a tuple with the MaxBidPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBidPrice

`func (o *CreateSpotMarketRequestRequest) SetMaxBidPrice(v float32)`

SetMaxBidPrice sets MaxBidPrice field to given value.

### HasMaxBidPrice

`func (o *CreateSpotMarketRequestRequest) HasMaxBidPrice() bool`

HasMaxBidPrice returns a boolean if a field has been set.

### GetMetro

`func (o *CreateSpotMarketRequestRequest) GetMetro() string`

GetMetro returns the Metro field if non-nil, zero value otherwise.

### GetMetroOk

`func (o *CreateSpotMarketRequestRequest) GetMetroOk() (*string, bool)`

GetMetroOk returns a tuple with the Metro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetro

`func (o *CreateSpotMarketRequestRequest) SetMetro(v string)`

SetMetro sets Metro field to given value.

### HasMetro

`func (o *CreateSpotMarketRequestRequest) HasMetro() bool`

HasMetro returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


