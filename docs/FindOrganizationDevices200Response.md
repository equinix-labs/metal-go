# FindOrganizationDevices200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Devices** | Pointer to [**[]FindDeviceById200Response**](FindDeviceById200Response.md) |  | [optional] 
**Meta** | Pointer to [**FindDeviceEvents200ResponseMeta**](FindDeviceEvents200ResponseMeta.md) |  | [optional] 

## Methods

### NewFindOrganizationDevices200Response

`func NewFindOrganizationDevices200Response() *FindOrganizationDevices200Response`

NewFindOrganizationDevices200Response instantiates a new FindOrganizationDevices200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindOrganizationDevices200ResponseWithDefaults

`func NewFindOrganizationDevices200ResponseWithDefaults() *FindOrganizationDevices200Response`

NewFindOrganizationDevices200ResponseWithDefaults instantiates a new FindOrganizationDevices200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevices

`func (o *FindOrganizationDevices200Response) GetDevices() []FindDeviceById200Response`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *FindOrganizationDevices200Response) GetDevicesOk() (*[]FindDeviceById200Response, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *FindOrganizationDevices200Response) SetDevices(v []FindDeviceById200Response)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *FindOrganizationDevices200Response) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### GetMeta

`func (o *FindOrganizationDevices200Response) GetMeta() FindDeviceEvents200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *FindOrganizationDevices200Response) GetMetaOk() (*FindDeviceEvents200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *FindOrganizationDevices200Response) SetMeta(v FindDeviceEvents200ResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *FindOrganizationDevices200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


