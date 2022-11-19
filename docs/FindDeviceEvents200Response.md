# FindDeviceEvents200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | Pointer to [**[]FindInterconnectionEvents200Response**](FindInterconnectionEvents200Response.md) |  | [optional] 
**Meta** | Pointer to [**FindDeviceEvents200ResponseMeta**](FindDeviceEvents200ResponseMeta.md) |  | [optional] 

## Methods

### NewFindDeviceEvents200Response

`func NewFindDeviceEvents200Response() *FindDeviceEvents200Response`

NewFindDeviceEvents200Response instantiates a new FindDeviceEvents200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindDeviceEvents200ResponseWithDefaults

`func NewFindDeviceEvents200ResponseWithDefaults() *FindDeviceEvents200Response`

NewFindDeviceEvents200ResponseWithDefaults instantiates a new FindDeviceEvents200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *FindDeviceEvents200Response) GetEvents() []FindInterconnectionEvents200Response`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *FindDeviceEvents200Response) GetEventsOk() (*[]FindInterconnectionEvents200Response, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *FindDeviceEvents200Response) SetEvents(v []FindInterconnectionEvents200Response)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *FindDeviceEvents200Response) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetMeta

`func (o *FindDeviceEvents200Response) GetMeta() FindDeviceEvents200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *FindDeviceEvents200Response) GetMetaOk() (*FindDeviceEvents200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *FindDeviceEvents200Response) SetMeta(v FindDeviceEvents200ResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *FindDeviceEvents200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


