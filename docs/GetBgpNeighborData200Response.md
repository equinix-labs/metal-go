# GetBgpNeighborData200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BgpNeighbors** | Pointer to [**[]GetBgpNeighborData200ResponseBgpNeighborsInner**](GetBgpNeighborData200ResponseBgpNeighborsInner.md) | A list of BGP session neighbor data | [optional] 

## Methods

### NewGetBgpNeighborData200Response

`func NewGetBgpNeighborData200Response() *GetBgpNeighborData200Response`

NewGetBgpNeighborData200Response instantiates a new GetBgpNeighborData200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBgpNeighborData200ResponseWithDefaults

`func NewGetBgpNeighborData200ResponseWithDefaults() *GetBgpNeighborData200Response`

NewGetBgpNeighborData200ResponseWithDefaults instantiates a new GetBgpNeighborData200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBgpNeighbors

`func (o *GetBgpNeighborData200Response) GetBgpNeighbors() []GetBgpNeighborData200ResponseBgpNeighborsInner`

GetBgpNeighbors returns the BgpNeighbors field if non-nil, zero value otherwise.

### GetBgpNeighborsOk

`func (o *GetBgpNeighborData200Response) GetBgpNeighborsOk() (*[]GetBgpNeighborData200ResponseBgpNeighborsInner, bool)`

GetBgpNeighborsOk returns a tuple with the BgpNeighbors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpNeighbors

`func (o *GetBgpNeighborData200Response) SetBgpNeighbors(v []GetBgpNeighborData200ResponseBgpNeighborsInner)`

SetBgpNeighbors sets BgpNeighbors field to given value.

### HasBgpNeighbors

`func (o *GetBgpNeighborData200Response) HasBgpNeighbors() bool`

HasBgpNeighbors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


