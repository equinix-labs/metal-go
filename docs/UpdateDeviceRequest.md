# UpdateDeviceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to **[]string** |  | [optional] 
**AlwaysPxe** | Pointer to **bool** |  | [optional] 
**BillingCycle** | Pointer to **string** |  | [optional] 
**Customdata** | Pointer to **map[string]interface{}** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**IpxeScriptUrl** | Pointer to **string** |  | [optional] 
**Locked** | Pointer to **bool** |  | [optional] 
**NetworkFrozen** | Pointer to **bool** | If true, this instance can not be converted to a different network type. | [optional] 
**SpotInstance** | Pointer to **bool** |  | [optional] 
**Userdata** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateDeviceRequest

`func NewUpdateDeviceRequest() *UpdateDeviceRequest`

NewUpdateDeviceRequest instantiates a new UpdateDeviceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDeviceRequestWithDefaults

`func NewUpdateDeviceRequestWithDefaults() *UpdateDeviceRequest`

NewUpdateDeviceRequestWithDefaults instantiates a new UpdateDeviceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *UpdateDeviceRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateDeviceRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateDeviceRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateDeviceRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetAlwaysPxe

`func (o *UpdateDeviceRequest) GetAlwaysPxe() bool`

GetAlwaysPxe returns the AlwaysPxe field if non-nil, zero value otherwise.

### GetAlwaysPxeOk

`func (o *UpdateDeviceRequest) GetAlwaysPxeOk() (*bool, bool)`

GetAlwaysPxeOk returns a tuple with the AlwaysPxe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlwaysPxe

`func (o *UpdateDeviceRequest) SetAlwaysPxe(v bool)`

SetAlwaysPxe sets AlwaysPxe field to given value.

### HasAlwaysPxe

`func (o *UpdateDeviceRequest) HasAlwaysPxe() bool`

HasAlwaysPxe returns a boolean if a field has been set.

### GetBillingCycle

`func (o *UpdateDeviceRequest) GetBillingCycle() string`

GetBillingCycle returns the BillingCycle field if non-nil, zero value otherwise.

### GetBillingCycleOk

`func (o *UpdateDeviceRequest) GetBillingCycleOk() (*string, bool)`

GetBillingCycleOk returns a tuple with the BillingCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCycle

`func (o *UpdateDeviceRequest) SetBillingCycle(v string)`

SetBillingCycle sets BillingCycle field to given value.

### HasBillingCycle

`func (o *UpdateDeviceRequest) HasBillingCycle() bool`

HasBillingCycle returns a boolean if a field has been set.

### GetCustomdata

`func (o *UpdateDeviceRequest) GetCustomdata() map[string]interface{}`

GetCustomdata returns the Customdata field if non-nil, zero value otherwise.

### GetCustomdataOk

`func (o *UpdateDeviceRequest) GetCustomdataOk() (*map[string]interface{}, bool)`

GetCustomdataOk returns a tuple with the Customdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomdata

`func (o *UpdateDeviceRequest) SetCustomdata(v map[string]interface{})`

SetCustomdata sets Customdata field to given value.

### HasCustomdata

`func (o *UpdateDeviceRequest) HasCustomdata() bool`

HasCustomdata returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateDeviceRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateDeviceRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateDeviceRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateDeviceRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHostname

`func (o *UpdateDeviceRequest) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *UpdateDeviceRequest) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *UpdateDeviceRequest) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *UpdateDeviceRequest) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetIpxeScriptUrl

`func (o *UpdateDeviceRequest) GetIpxeScriptUrl() string`

GetIpxeScriptUrl returns the IpxeScriptUrl field if non-nil, zero value otherwise.

### GetIpxeScriptUrlOk

`func (o *UpdateDeviceRequest) GetIpxeScriptUrlOk() (*string, bool)`

GetIpxeScriptUrlOk returns a tuple with the IpxeScriptUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpxeScriptUrl

`func (o *UpdateDeviceRequest) SetIpxeScriptUrl(v string)`

SetIpxeScriptUrl sets IpxeScriptUrl field to given value.

### HasIpxeScriptUrl

`func (o *UpdateDeviceRequest) HasIpxeScriptUrl() bool`

HasIpxeScriptUrl returns a boolean if a field has been set.

### GetLocked

`func (o *UpdateDeviceRequest) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *UpdateDeviceRequest) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *UpdateDeviceRequest) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *UpdateDeviceRequest) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetNetworkFrozen

`func (o *UpdateDeviceRequest) GetNetworkFrozen() bool`

GetNetworkFrozen returns the NetworkFrozen field if non-nil, zero value otherwise.

### GetNetworkFrozenOk

`func (o *UpdateDeviceRequest) GetNetworkFrozenOk() (*bool, bool)`

GetNetworkFrozenOk returns a tuple with the NetworkFrozen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkFrozen

`func (o *UpdateDeviceRequest) SetNetworkFrozen(v bool)`

SetNetworkFrozen sets NetworkFrozen field to given value.

### HasNetworkFrozen

`func (o *UpdateDeviceRequest) HasNetworkFrozen() bool`

HasNetworkFrozen returns a boolean if a field has been set.

### GetSpotInstance

`func (o *UpdateDeviceRequest) GetSpotInstance() bool`

GetSpotInstance returns the SpotInstance field if non-nil, zero value otherwise.

### GetSpotInstanceOk

`func (o *UpdateDeviceRequest) GetSpotInstanceOk() (*bool, bool)`

GetSpotInstanceOk returns a tuple with the SpotInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotInstance

`func (o *UpdateDeviceRequest) SetSpotInstance(v bool)`

SetSpotInstance sets SpotInstance field to given value.

### HasSpotInstance

`func (o *UpdateDeviceRequest) HasSpotInstance() bool`

HasSpotInstance returns a boolean if a field has been set.

### GetUserdata

`func (o *UpdateDeviceRequest) GetUserdata() string`

GetUserdata returns the Userdata field if non-nil, zero value otherwise.

### GetUserdataOk

`func (o *UpdateDeviceRequest) GetUserdataOk() (*string, bool)`

GetUserdataOk returns a tuple with the Userdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserdata

`func (o *UpdateDeviceRequest) SetUserdata(v string)`

SetUserdata sets Userdata field to given value.

### HasUserdata

`func (o *UpdateDeviceRequest) HasUserdata() bool`

HasUserdata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


