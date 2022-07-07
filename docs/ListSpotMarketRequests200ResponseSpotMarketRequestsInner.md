# ListSpotMarketRequests200ResponseSpotMarketRequestsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**DevicesMax** | Pointer to **int32** |  | [optional] 
**DevicesMin** | Pointer to **int32** |  | [optional] 
**EndAt** | Pointer to **time.Time** |  | [optional] 
**Facilities** | Pointer to [**FindBatchById200ResponseDevicesInner**](FindBatchById200ResponseDevicesInner.md) |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Instances** | Pointer to [**FindBatchById200ResponseDevicesInner**](FindBatchById200ResponseDevicesInner.md) |  | [optional] 
**MaxBidPrice** | Pointer to **float32** |  | [optional] 
**Metro** | Pointer to [**ListSpotMarketRequests200ResponseSpotMarketRequestsInnerMetro**](ListSpotMarketRequests200ResponseSpotMarketRequestsInnerMetro.md) |  | [optional] 
**Project** | Pointer to [**FindBatchById200ResponseDevicesInner**](FindBatchById200ResponseDevicesInner.md) |  | [optional] 

## Methods

### NewListSpotMarketRequests200ResponseSpotMarketRequestsInner

`func NewListSpotMarketRequests200ResponseSpotMarketRequestsInner() *ListSpotMarketRequests200ResponseSpotMarketRequestsInner`

NewListSpotMarketRequests200ResponseSpotMarketRequestsInner instantiates a new ListSpotMarketRequests200ResponseSpotMarketRequestsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSpotMarketRequests200ResponseSpotMarketRequestsInnerWithDefaults

`func NewListSpotMarketRequests200ResponseSpotMarketRequestsInnerWithDefaults() *ListSpotMarketRequests200ResponseSpotMarketRequestsInner`

NewListSpotMarketRequests200ResponseSpotMarketRequestsInnerWithDefaults instantiates a new ListSpotMarketRequests200ResponseSpotMarketRequestsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ListSpotMarketRequests200ResponseSpotMarketRequestsInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListSpotMarketRequests200ResponseSpotMarketRequestsInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListSpotMarketRequests200ResponseSpotMarketRequestsInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ListSpotMarketRequests200ResponseSpotMarketRequestsInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDevicesMax

`func (o *ListSpotMarketRequests200ResponseSpotMarketRequestsInner) GetDevicesMax() int32`

GetDevicesMax returns the DevicesMax field if non-nil, zero value otherwise.

### GetDevicesMaxOk

`func (o *ListSpotMarketRequests200ResponseSpotMarketRequestsInner) GetDevicesMaxOk() (*int32, bool)`

GetDevicesMaxOk returns a tuple with the DevicesMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicesMax

`func (o *ListSpotMarketRequests200ResponseSpotMarketRequestsInner) SetDevicesMax(v int32)`

SetDevicesMax sets DevicesMax field to given value.

### HasDevicesMax

`func (o *ListSpotMarketRequests200ResponseSpotMarketRequestsInner) HasDevicesMax() bool`

HasDevicesMax returns a boolean if a field has been set.

### GetDevicesMin

`func (o *ListSpotMarketRequests200ResponseSpotMarketRequestsInner) GetDevicesMin() int32`

GetDevicesMin returns the DevicesMin field if non-nil, zero value otherwise.

### GetDevicesMinOk

`func (o *ListSpotMarketRequests200ResponseSpotMarketRequestsInner) GetDevicesMinOk() (*int32, bool)`

GetDevicesMinOk returns a tuple with the DevicesMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicesMin

`func (o *ListSpotMarketRequests200ResponseSpotMarketRequestsInner) SetDevicesMin(v int32)`

SetDevicesMin sets DevicesMin field to given value.

### HasDevicesMin

`func (o *ListSpotMarketRequests200ResponseSpotMarketRequestsInner) HasDevicesMin() bool`

HasDevicesMin returns a boolean if a field has been set.

### GetEndAt

`func (o *ListSpotMarketRequests200ResponseSpotMarketRequestsInner) GetEndAt() time.Time`

GetEndAt returns the EndAt field if non-nil, zero value otherwise.

### GetEndAtOk

`func (o *ListSpotMarketRequests200ResponseSpotMarketRequestsInner) GetEndAtOk() (*time.Time, bool)`

GetEndAtOk returns a tuple with the EndAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAt

`func (o *ListSpotMarketRequests200ResponseSpotMarketRequestsInner) SetEndAt(v time.Time)`

SetEndAt sets EndAt field to given value.

### HasEndAt

`func (o *ListSpotMarketRequests200ResponseSpotMarketRequestsInner) HasEndAt() bool`

HasEndAt returns a boolean if a field has been set.

### GetFacilities

`func (o *ListSpotMarketRequests200ResponseSpotMarketRequestsInner) GetFacilities() FindBatchById200ResponseDevicesInner`

GetFacilities returns the Facilities field if non-nil, zero value otherwise.

### GetFacilitiesOk

`func (o *ListSpotMarketRequests200ResponseSpotMarketRequestsInner) GetFacilitiesOk() (*FindBatchById200ResponseDevicesInner, bool)`

GetFacilitiesOk returns a tuple with the Facilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilities

`func (o *ListSpotMarketRequests200ResponseSpotMarketRequestsInner) SetFacilities(v FindBatchById200ResponseDevicesInner)`

SetFacilities sets Facilities field to given value.

### HasFacilities

`func (o *ListSpotMarketRequests200ResponseSpotMarketRequestsInner) HasFacilities() bool`

HasFacilities returns a boolean if a field has been set.

### GetHref

`func (o *ListSpotMarketRequests200ResponseSpotMarketRequestsInner) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ListSpotMarketRequests200ResponseSpotMarketRequestsInner) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ListSpotMarketRequests200ResponseSpotMarketRequestsInner) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ListSpotMarketRequests200ResponseSpotMarketRequestsInner) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *ListSpotMarketRequests200ResponseSpotMarketRequestsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListSpotMarketRequests200ResponseSpotMarketRequestsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListSpotMarketRequests200ResponseSpotMarketRequestsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ListSpotMarketRequests200ResponseSpotMarketRequestsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInstances

`func (o *ListSpotMarketRequests200ResponseSpotMarketRequestsInner) GetInstances() FindBatchById200ResponseDevicesInner`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *ListSpotMarketRequests200ResponseSpotMarketRequestsInner) GetInstancesOk() (*FindBatchById200ResponseDevicesInner, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *ListSpotMarketRequests200ResponseSpotMarketRequestsInner) SetInstances(v FindBatchById200ResponseDevicesInner)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *ListSpotMarketRequests200ResponseSpotMarketRequestsInner) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

### GetMaxBidPrice

`func (o *ListSpotMarketRequests200ResponseSpotMarketRequestsInner) GetMaxBidPrice() float32`

GetMaxBidPrice returns the MaxBidPrice field if non-nil, zero value otherwise.

### GetMaxBidPriceOk

`func (o *ListSpotMarketRequests200ResponseSpotMarketRequestsInner) GetMaxBidPriceOk() (*float32, bool)`

GetMaxBidPriceOk returns a tuple with the MaxBidPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBidPrice

`func (o *ListSpotMarketRequests200ResponseSpotMarketRequestsInner) SetMaxBidPrice(v float32)`

SetMaxBidPrice sets MaxBidPrice field to given value.

### HasMaxBidPrice

`func (o *ListSpotMarketRequests200ResponseSpotMarketRequestsInner) HasMaxBidPrice() bool`

HasMaxBidPrice returns a boolean if a field has been set.

### GetMetro

`func (o *ListSpotMarketRequests200ResponseSpotMarketRequestsInner) GetMetro() ListSpotMarketRequests200ResponseSpotMarketRequestsInnerMetro`

GetMetro returns the Metro field if non-nil, zero value otherwise.

### GetMetroOk

`func (o *ListSpotMarketRequests200ResponseSpotMarketRequestsInner) GetMetroOk() (*ListSpotMarketRequests200ResponseSpotMarketRequestsInnerMetro, bool)`

GetMetroOk returns a tuple with the Metro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetro

`func (o *ListSpotMarketRequests200ResponseSpotMarketRequestsInner) SetMetro(v ListSpotMarketRequests200ResponseSpotMarketRequestsInnerMetro)`

SetMetro sets Metro field to given value.

### HasMetro

`func (o *ListSpotMarketRequests200ResponseSpotMarketRequestsInner) HasMetro() bool`

HasMetro returns a boolean if a field has been set.

### GetProject

`func (o *ListSpotMarketRequests200ResponseSpotMarketRequestsInner) GetProject() FindBatchById200ResponseDevicesInner`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ListSpotMarketRequests200ResponseSpotMarketRequestsInner) GetProjectOk() (*FindBatchById200ResponseDevicesInner, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ListSpotMarketRequests200ResponseSpotMarketRequestsInner) SetProject(v FindBatchById200ResponseDevicesInner)`

SetProject sets Project field to given value.

### HasProject

`func (o *ListSpotMarketRequests200ResponseSpotMarketRequestsInner) HasProject() bool`

HasProject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


