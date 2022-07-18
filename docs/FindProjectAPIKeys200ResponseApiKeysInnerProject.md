# FindProjectAPIKeys200ResponseApiKeysInnerProject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BgpConfig** | Pointer to [**FindBatchById200ResponseDevicesInner**](FindBatchById200ResponseDevicesInner.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Customdata** | Pointer to **map[string]interface{}** |  | [optional] 
**Devices** | Pointer to [**[]FindBatchById200ResponseDevicesInner**](FindBatchById200ResponseDevicesInner.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Invitations** | Pointer to [**[]FindBatchById200ResponseDevicesInner**](FindBatchById200ResponseDevicesInner.md) |  | [optional] 
**MaxDevices** | Pointer to **map[string]interface{}** |  | [optional] 
**Members** | Pointer to [**[]FindBatchById200ResponseDevicesInner**](FindBatchById200ResponseDevicesInner.md) |  | [optional] 
**Memberships** | Pointer to [**[]FindBatchById200ResponseDevicesInner**](FindBatchById200ResponseDevicesInner.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NetworkStatus** | Pointer to **map[string]interface{}** |  | [optional] 
**PaymentMethod** | Pointer to [**FindBatchById200ResponseDevicesInner**](FindBatchById200ResponseDevicesInner.md) |  | [optional] 
**SshKeys** | Pointer to [**[]FindBatchById200ResponseDevicesInner**](FindBatchById200ResponseDevicesInner.md) |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Volumes** | Pointer to [**[]FindBatchById200ResponseDevicesInner**](FindBatchById200ResponseDevicesInner.md) |  | [optional] 

## Methods

### NewFindProjectAPIKeys200ResponseApiKeysInnerProject

`func NewFindProjectAPIKeys200ResponseApiKeysInnerProject() *FindProjectAPIKeys200ResponseApiKeysInnerProject`

NewFindProjectAPIKeys200ResponseApiKeysInnerProject instantiates a new FindProjectAPIKeys200ResponseApiKeysInnerProject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindProjectAPIKeys200ResponseApiKeysInnerProjectWithDefaults

`func NewFindProjectAPIKeys200ResponseApiKeysInnerProjectWithDefaults() *FindProjectAPIKeys200ResponseApiKeysInnerProject`

NewFindProjectAPIKeys200ResponseApiKeysInnerProjectWithDefaults instantiates a new FindProjectAPIKeys200ResponseApiKeysInnerProject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBgpConfig

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) GetBgpConfig() FindBatchById200ResponseDevicesInner`

GetBgpConfig returns the BgpConfig field if non-nil, zero value otherwise.

### GetBgpConfigOk

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) GetBgpConfigOk() (*FindBatchById200ResponseDevicesInner, bool)`

GetBgpConfigOk returns a tuple with the BgpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpConfig

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) SetBgpConfig(v FindBatchById200ResponseDevicesInner)`

SetBgpConfig sets BgpConfig field to given value.

### HasBgpConfig

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) HasBgpConfig() bool`

HasBgpConfig returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCustomdata

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) GetCustomdata() map[string]interface{}`

GetCustomdata returns the Customdata field if non-nil, zero value otherwise.

### GetCustomdataOk

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) GetCustomdataOk() (*map[string]interface{}, bool)`

GetCustomdataOk returns a tuple with the Customdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomdata

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) SetCustomdata(v map[string]interface{})`

SetCustomdata sets Customdata field to given value.

### HasCustomdata

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) HasCustomdata() bool`

HasCustomdata returns a boolean if a field has been set.

### GetDevices

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) GetDevices() []FindBatchById200ResponseDevicesInner`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) GetDevicesOk() (*[]FindBatchById200ResponseDevicesInner, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) SetDevices(v []FindBatchById200ResponseDevicesInner)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### GetId

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInvitations

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) GetInvitations() []FindBatchById200ResponseDevicesInner`

GetInvitations returns the Invitations field if non-nil, zero value otherwise.

### GetInvitationsOk

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) GetInvitationsOk() (*[]FindBatchById200ResponseDevicesInner, bool)`

GetInvitationsOk returns a tuple with the Invitations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitations

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) SetInvitations(v []FindBatchById200ResponseDevicesInner)`

SetInvitations sets Invitations field to given value.

### HasInvitations

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) HasInvitations() bool`

HasInvitations returns a boolean if a field has been set.

### GetMaxDevices

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) GetMaxDevices() map[string]interface{}`

GetMaxDevices returns the MaxDevices field if non-nil, zero value otherwise.

### GetMaxDevicesOk

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) GetMaxDevicesOk() (*map[string]interface{}, bool)`

GetMaxDevicesOk returns a tuple with the MaxDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDevices

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) SetMaxDevices(v map[string]interface{})`

SetMaxDevices sets MaxDevices field to given value.

### HasMaxDevices

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) HasMaxDevices() bool`

HasMaxDevices returns a boolean if a field has been set.

### GetMembers

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) GetMembers() []FindBatchById200ResponseDevicesInner`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) GetMembersOk() (*[]FindBatchById200ResponseDevicesInner, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) SetMembers(v []FindBatchById200ResponseDevicesInner)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetMemberships

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) GetMemberships() []FindBatchById200ResponseDevicesInner`

GetMemberships returns the Memberships field if non-nil, zero value otherwise.

### GetMembershipsOk

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) GetMembershipsOk() (*[]FindBatchById200ResponseDevicesInner, bool)`

GetMembershipsOk returns a tuple with the Memberships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberships

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) SetMemberships(v []FindBatchById200ResponseDevicesInner)`

SetMemberships sets Memberships field to given value.

### HasMemberships

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) HasMemberships() bool`

HasMemberships returns a boolean if a field has been set.

### GetName

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkStatus

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) GetNetworkStatus() map[string]interface{}`

GetNetworkStatus returns the NetworkStatus field if non-nil, zero value otherwise.

### GetNetworkStatusOk

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) GetNetworkStatusOk() (*map[string]interface{}, bool)`

GetNetworkStatusOk returns a tuple with the NetworkStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkStatus

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) SetNetworkStatus(v map[string]interface{})`

SetNetworkStatus sets NetworkStatus field to given value.

### HasNetworkStatus

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) HasNetworkStatus() bool`

HasNetworkStatus returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) GetPaymentMethod() FindBatchById200ResponseDevicesInner`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) GetPaymentMethodOk() (*FindBatchById200ResponseDevicesInner, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) SetPaymentMethod(v FindBatchById200ResponseDevicesInner)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetSshKeys

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) GetSshKeys() []FindBatchById200ResponseDevicesInner`

GetSshKeys returns the SshKeys field if non-nil, zero value otherwise.

### GetSshKeysOk

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) GetSshKeysOk() (*[]FindBatchById200ResponseDevicesInner, bool)`

GetSshKeysOk returns a tuple with the SshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeys

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) SetSshKeys(v []FindBatchById200ResponseDevicesInner)`

SetSshKeys sets SshKeys field to given value.

### HasSshKeys

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) HasSshKeys() bool`

HasSshKeys returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVolumes

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) GetVolumes() []FindBatchById200ResponseDevicesInner`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) GetVolumesOk() (*[]FindBatchById200ResponseDevicesInner, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) SetVolumes(v []FindBatchById200ResponseDevicesInner)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


