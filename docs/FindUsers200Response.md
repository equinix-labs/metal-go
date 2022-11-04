# FindUsers200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**FindDeviceEvents200ResponseMeta**](FindDeviceEvents200ResponseMeta.md) |  | [optional] 
**Users** | Pointer to [**[]GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1VrfCreatedBy**](GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1VrfCreatedBy.md) |  | [optional] 

## Methods

### NewFindUsers200Response

`func NewFindUsers200Response() *FindUsers200Response`

NewFindUsers200Response instantiates a new FindUsers200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindUsers200ResponseWithDefaults

`func NewFindUsers200ResponseWithDefaults() *FindUsers200Response`

NewFindUsers200ResponseWithDefaults instantiates a new FindUsers200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *FindUsers200Response) GetMeta() FindDeviceEvents200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *FindUsers200Response) GetMetaOk() (*FindDeviceEvents200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *FindUsers200Response) SetMeta(v FindDeviceEvents200ResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *FindUsers200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUsers

`func (o *FindUsers200Response) GetUsers() []GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1VrfCreatedBy`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *FindUsers200Response) GetUsersOk() (*[]GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1VrfCreatedBy, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *FindUsers200Response) SetUsers(v []GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1VrfCreatedBy)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *FindUsers200Response) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


