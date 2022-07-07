# MoveHardwareReservation201ResponseProject

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

### NewMoveHardwareReservation201ResponseProject

`func NewMoveHardwareReservation201ResponseProject() *MoveHardwareReservation201ResponseProject`

NewMoveHardwareReservation201ResponseProject instantiates a new MoveHardwareReservation201ResponseProject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoveHardwareReservation201ResponseProjectWithDefaults

`func NewMoveHardwareReservation201ResponseProjectWithDefaults() *MoveHardwareReservation201ResponseProject`

NewMoveHardwareReservation201ResponseProjectWithDefaults instantiates a new MoveHardwareReservation201ResponseProject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBgpConfig

`func (o *MoveHardwareReservation201ResponseProject) GetBgpConfig() FindBatchById200ResponseDevicesInner`

GetBgpConfig returns the BgpConfig field if non-nil, zero value otherwise.

### GetBgpConfigOk

`func (o *MoveHardwareReservation201ResponseProject) GetBgpConfigOk() (*FindBatchById200ResponseDevicesInner, bool)`

GetBgpConfigOk returns a tuple with the BgpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpConfig

`func (o *MoveHardwareReservation201ResponseProject) SetBgpConfig(v FindBatchById200ResponseDevicesInner)`

SetBgpConfig sets BgpConfig field to given value.

### HasBgpConfig

`func (o *MoveHardwareReservation201ResponseProject) HasBgpConfig() bool`

HasBgpConfig returns a boolean if a field has been set.

### GetCreatedAt

`func (o *MoveHardwareReservation201ResponseProject) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MoveHardwareReservation201ResponseProject) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MoveHardwareReservation201ResponseProject) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MoveHardwareReservation201ResponseProject) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCustomdata

`func (o *MoveHardwareReservation201ResponseProject) GetCustomdata() map[string]interface{}`

GetCustomdata returns the Customdata field if non-nil, zero value otherwise.

### GetCustomdataOk

`func (o *MoveHardwareReservation201ResponseProject) GetCustomdataOk() (*map[string]interface{}, bool)`

GetCustomdataOk returns a tuple with the Customdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomdata

`func (o *MoveHardwareReservation201ResponseProject) SetCustomdata(v map[string]interface{})`

SetCustomdata sets Customdata field to given value.

### HasCustomdata

`func (o *MoveHardwareReservation201ResponseProject) HasCustomdata() bool`

HasCustomdata returns a boolean if a field has been set.

### GetDevices

`func (o *MoveHardwareReservation201ResponseProject) GetDevices() []FindBatchById200ResponseDevicesInner`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *MoveHardwareReservation201ResponseProject) GetDevicesOk() (*[]FindBatchById200ResponseDevicesInner, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *MoveHardwareReservation201ResponseProject) SetDevices(v []FindBatchById200ResponseDevicesInner)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *MoveHardwareReservation201ResponseProject) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### GetId

`func (o *MoveHardwareReservation201ResponseProject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MoveHardwareReservation201ResponseProject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MoveHardwareReservation201ResponseProject) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MoveHardwareReservation201ResponseProject) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInvitations

`func (o *MoveHardwareReservation201ResponseProject) GetInvitations() []FindBatchById200ResponseDevicesInner`

GetInvitations returns the Invitations field if non-nil, zero value otherwise.

### GetInvitationsOk

`func (o *MoveHardwareReservation201ResponseProject) GetInvitationsOk() (*[]FindBatchById200ResponseDevicesInner, bool)`

GetInvitationsOk returns a tuple with the Invitations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitations

`func (o *MoveHardwareReservation201ResponseProject) SetInvitations(v []FindBatchById200ResponseDevicesInner)`

SetInvitations sets Invitations field to given value.

### HasInvitations

`func (o *MoveHardwareReservation201ResponseProject) HasInvitations() bool`

HasInvitations returns a boolean if a field has been set.

### GetMaxDevices

`func (o *MoveHardwareReservation201ResponseProject) GetMaxDevices() map[string]interface{}`

GetMaxDevices returns the MaxDevices field if non-nil, zero value otherwise.

### GetMaxDevicesOk

`func (o *MoveHardwareReservation201ResponseProject) GetMaxDevicesOk() (*map[string]interface{}, bool)`

GetMaxDevicesOk returns a tuple with the MaxDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDevices

`func (o *MoveHardwareReservation201ResponseProject) SetMaxDevices(v map[string]interface{})`

SetMaxDevices sets MaxDevices field to given value.

### HasMaxDevices

`func (o *MoveHardwareReservation201ResponseProject) HasMaxDevices() bool`

HasMaxDevices returns a boolean if a field has been set.

### GetMembers

`func (o *MoveHardwareReservation201ResponseProject) GetMembers() []FindBatchById200ResponseDevicesInner`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *MoveHardwareReservation201ResponseProject) GetMembersOk() (*[]FindBatchById200ResponseDevicesInner, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *MoveHardwareReservation201ResponseProject) SetMembers(v []FindBatchById200ResponseDevicesInner)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *MoveHardwareReservation201ResponseProject) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetMemberships

`func (o *MoveHardwareReservation201ResponseProject) GetMemberships() []FindBatchById200ResponseDevicesInner`

GetMemberships returns the Memberships field if non-nil, zero value otherwise.

### GetMembershipsOk

`func (o *MoveHardwareReservation201ResponseProject) GetMembershipsOk() (*[]FindBatchById200ResponseDevicesInner, bool)`

GetMembershipsOk returns a tuple with the Memberships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberships

`func (o *MoveHardwareReservation201ResponseProject) SetMemberships(v []FindBatchById200ResponseDevicesInner)`

SetMemberships sets Memberships field to given value.

### HasMemberships

`func (o *MoveHardwareReservation201ResponseProject) HasMemberships() bool`

HasMemberships returns a boolean if a field has been set.

### GetName

`func (o *MoveHardwareReservation201ResponseProject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MoveHardwareReservation201ResponseProject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MoveHardwareReservation201ResponseProject) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MoveHardwareReservation201ResponseProject) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkStatus

`func (o *MoveHardwareReservation201ResponseProject) GetNetworkStatus() map[string]interface{}`

GetNetworkStatus returns the NetworkStatus field if non-nil, zero value otherwise.

### GetNetworkStatusOk

`func (o *MoveHardwareReservation201ResponseProject) GetNetworkStatusOk() (*map[string]interface{}, bool)`

GetNetworkStatusOk returns a tuple with the NetworkStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkStatus

`func (o *MoveHardwareReservation201ResponseProject) SetNetworkStatus(v map[string]interface{})`

SetNetworkStatus sets NetworkStatus field to given value.

### HasNetworkStatus

`func (o *MoveHardwareReservation201ResponseProject) HasNetworkStatus() bool`

HasNetworkStatus returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *MoveHardwareReservation201ResponseProject) GetPaymentMethod() FindBatchById200ResponseDevicesInner`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *MoveHardwareReservation201ResponseProject) GetPaymentMethodOk() (*FindBatchById200ResponseDevicesInner, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *MoveHardwareReservation201ResponseProject) SetPaymentMethod(v FindBatchById200ResponseDevicesInner)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *MoveHardwareReservation201ResponseProject) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetSshKeys

`func (o *MoveHardwareReservation201ResponseProject) GetSshKeys() []FindBatchById200ResponseDevicesInner`

GetSshKeys returns the SshKeys field if non-nil, zero value otherwise.

### GetSshKeysOk

`func (o *MoveHardwareReservation201ResponseProject) GetSshKeysOk() (*[]FindBatchById200ResponseDevicesInner, bool)`

GetSshKeysOk returns a tuple with the SshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeys

`func (o *MoveHardwareReservation201ResponseProject) SetSshKeys(v []FindBatchById200ResponseDevicesInner)`

SetSshKeys sets SshKeys field to given value.

### HasSshKeys

`func (o *MoveHardwareReservation201ResponseProject) HasSshKeys() bool`

HasSshKeys returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *MoveHardwareReservation201ResponseProject) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MoveHardwareReservation201ResponseProject) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MoveHardwareReservation201ResponseProject) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *MoveHardwareReservation201ResponseProject) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVolumes

`func (o *MoveHardwareReservation201ResponseProject) GetVolumes() []FindBatchById200ResponseDevicesInner`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *MoveHardwareReservation201ResponseProject) GetVolumesOk() (*[]FindBatchById200ResponseDevicesInner, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *MoveHardwareReservation201ResponseProject) SetVolumes(v []FindBatchById200ResponseDevicesInner)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *MoveHardwareReservation201ResponseProject) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


