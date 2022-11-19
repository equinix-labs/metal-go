# UpdateInterconnectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContactEmail** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Mode** | Pointer to **string** | The mode of the interconnection (only relevant to Dedicated Ports). Shared connections won&#39;t have this field. Can be either &#39;standard&#39; or &#39;tunnel&#39;.   The default mode of an interconnection on a Dedicated Port is &#39;standard&#39;. The mode can only be changed when there are no associated virtual circuits on the interconnection.   In tunnel mode, an 802.1q tunnel is added to a port to send/receive double tagged packets from server instances. | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Redundancy** | Pointer to **string** | Updating from &#39;redundant&#39; to &#39;primary&#39; will remove a secondary port, while updating from &#39;primary&#39; to &#39;redundant&#39; will add one. | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewUpdateInterconnectionRequest

`func NewUpdateInterconnectionRequest() *UpdateInterconnectionRequest`

NewUpdateInterconnectionRequest instantiates a new UpdateInterconnectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateInterconnectionRequestWithDefaults

`func NewUpdateInterconnectionRequestWithDefaults() *UpdateInterconnectionRequest`

NewUpdateInterconnectionRequestWithDefaults instantiates a new UpdateInterconnectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContactEmail

`func (o *UpdateInterconnectionRequest) GetContactEmail() string`

GetContactEmail returns the ContactEmail field if non-nil, zero value otherwise.

### GetContactEmailOk

`func (o *UpdateInterconnectionRequest) GetContactEmailOk() (*string, bool)`

GetContactEmailOk returns a tuple with the ContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactEmail

`func (o *UpdateInterconnectionRequest) SetContactEmail(v string)`

SetContactEmail sets ContactEmail field to given value.

### HasContactEmail

`func (o *UpdateInterconnectionRequest) HasContactEmail() bool`

HasContactEmail returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateInterconnectionRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateInterconnectionRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateInterconnectionRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateInterconnectionRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMode

`func (o *UpdateInterconnectionRequest) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *UpdateInterconnectionRequest) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *UpdateInterconnectionRequest) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *UpdateInterconnectionRequest) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetName

`func (o *UpdateInterconnectionRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateInterconnectionRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateInterconnectionRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateInterconnectionRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRedundancy

`func (o *UpdateInterconnectionRequest) GetRedundancy() string`

GetRedundancy returns the Redundancy field if non-nil, zero value otherwise.

### GetRedundancyOk

`func (o *UpdateInterconnectionRequest) GetRedundancyOk() (*string, bool)`

GetRedundancyOk returns a tuple with the Redundancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancy

`func (o *UpdateInterconnectionRequest) SetRedundancy(v string)`

SetRedundancy sets Redundancy field to given value.

### HasRedundancy

`func (o *UpdateInterconnectionRequest) HasRedundancy() bool`

HasRedundancy returns a boolean if a field has been set.

### GetTags

`func (o *UpdateInterconnectionRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateInterconnectionRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateInterconnectionRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateInterconnectionRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


