# ProjectOrHref

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | **string** |  | 
**BgpConfig** | Pointer to [**Href**](Href.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Customdata** | Pointer to **map[string]interface{}** |  | [optional] 
**Devices** | Pointer to [**[]Href**](Href.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Invitations** | Pointer to [**[]Href**](Href.md) |  | [optional] 
**MaxDevices** | Pointer to **map[string]interface{}** |  | [optional] 
**Members** | Pointer to [**[]Href**](Href.md) |  | [optional] 
**Memberships** | Pointer to [**[]Href**](Href.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NetworkStatus** | Pointer to **map[string]interface{}** |  | [optional] 
**PaymentMethod** | Pointer to [**Href**](Href.md) |  | [optional] 
**SshKeys** | Pointer to [**[]Href**](Href.md) |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Volumes** | Pointer to [**[]Href**](Href.md) |  | [optional] 

## Methods

### NewProjectOrHref

`func NewProjectOrHref(href string, ) *ProjectOrHref`

NewProjectOrHref instantiates a new ProjectOrHref object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectOrHrefWithDefaults

`func NewProjectOrHrefWithDefaults() *ProjectOrHref`

NewProjectOrHrefWithDefaults instantiates a new ProjectOrHref object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *ProjectOrHref) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ProjectOrHref) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ProjectOrHref) SetHref(v string)`

SetHref sets Href field to given value.


### GetBgpConfig

`func (o *ProjectOrHref) GetBgpConfig() Href`

GetBgpConfig returns the BgpConfig field if non-nil, zero value otherwise.

### GetBgpConfigOk

`func (o *ProjectOrHref) GetBgpConfigOk() (*Href, bool)`

GetBgpConfigOk returns a tuple with the BgpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpConfig

`func (o *ProjectOrHref) SetBgpConfig(v Href)`

SetBgpConfig sets BgpConfig field to given value.

### HasBgpConfig

`func (o *ProjectOrHref) HasBgpConfig() bool`

HasBgpConfig returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ProjectOrHref) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ProjectOrHref) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ProjectOrHref) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ProjectOrHref) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCustomdata

`func (o *ProjectOrHref) GetCustomdata() map[string]interface{}`

GetCustomdata returns the Customdata field if non-nil, zero value otherwise.

### GetCustomdataOk

`func (o *ProjectOrHref) GetCustomdataOk() (*map[string]interface{}, bool)`

GetCustomdataOk returns a tuple with the Customdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomdata

`func (o *ProjectOrHref) SetCustomdata(v map[string]interface{})`

SetCustomdata sets Customdata field to given value.

### HasCustomdata

`func (o *ProjectOrHref) HasCustomdata() bool`

HasCustomdata returns a boolean if a field has been set.

### GetDevices

`func (o *ProjectOrHref) GetDevices() []Href`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *ProjectOrHref) GetDevicesOk() (*[]Href, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *ProjectOrHref) SetDevices(v []Href)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *ProjectOrHref) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### GetId

`func (o *ProjectOrHref) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectOrHref) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectOrHref) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProjectOrHref) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInvitations

`func (o *ProjectOrHref) GetInvitations() []Href`

GetInvitations returns the Invitations field if non-nil, zero value otherwise.

### GetInvitationsOk

`func (o *ProjectOrHref) GetInvitationsOk() (*[]Href, bool)`

GetInvitationsOk returns a tuple with the Invitations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitations

`func (o *ProjectOrHref) SetInvitations(v []Href)`

SetInvitations sets Invitations field to given value.

### HasInvitations

`func (o *ProjectOrHref) HasInvitations() bool`

HasInvitations returns a boolean if a field has been set.

### GetMaxDevices

`func (o *ProjectOrHref) GetMaxDevices() map[string]interface{}`

GetMaxDevices returns the MaxDevices field if non-nil, zero value otherwise.

### GetMaxDevicesOk

`func (o *ProjectOrHref) GetMaxDevicesOk() (*map[string]interface{}, bool)`

GetMaxDevicesOk returns a tuple with the MaxDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDevices

`func (o *ProjectOrHref) SetMaxDevices(v map[string]interface{})`

SetMaxDevices sets MaxDevices field to given value.

### HasMaxDevices

`func (o *ProjectOrHref) HasMaxDevices() bool`

HasMaxDevices returns a boolean if a field has been set.

### GetMembers

`func (o *ProjectOrHref) GetMembers() []Href`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *ProjectOrHref) GetMembersOk() (*[]Href, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *ProjectOrHref) SetMembers(v []Href)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *ProjectOrHref) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetMemberships

`func (o *ProjectOrHref) GetMemberships() []Href`

GetMemberships returns the Memberships field if non-nil, zero value otherwise.

### GetMembershipsOk

`func (o *ProjectOrHref) GetMembershipsOk() (*[]Href, bool)`

GetMembershipsOk returns a tuple with the Memberships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberships

`func (o *ProjectOrHref) SetMemberships(v []Href)`

SetMemberships sets Memberships field to given value.

### HasMemberships

`func (o *ProjectOrHref) HasMemberships() bool`

HasMemberships returns a boolean if a field has been set.

### GetName

`func (o *ProjectOrHref) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectOrHref) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectOrHref) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProjectOrHref) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkStatus

`func (o *ProjectOrHref) GetNetworkStatus() map[string]interface{}`

GetNetworkStatus returns the NetworkStatus field if non-nil, zero value otherwise.

### GetNetworkStatusOk

`func (o *ProjectOrHref) GetNetworkStatusOk() (*map[string]interface{}, bool)`

GetNetworkStatusOk returns a tuple with the NetworkStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkStatus

`func (o *ProjectOrHref) SetNetworkStatus(v map[string]interface{})`

SetNetworkStatus sets NetworkStatus field to given value.

### HasNetworkStatus

`func (o *ProjectOrHref) HasNetworkStatus() bool`

HasNetworkStatus returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *ProjectOrHref) GetPaymentMethod() Href`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *ProjectOrHref) GetPaymentMethodOk() (*Href, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *ProjectOrHref) SetPaymentMethod(v Href)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *ProjectOrHref) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetSshKeys

`func (o *ProjectOrHref) GetSshKeys() []Href`

GetSshKeys returns the SshKeys field if non-nil, zero value otherwise.

### GetSshKeysOk

`func (o *ProjectOrHref) GetSshKeysOk() (*[]Href, bool)`

GetSshKeysOk returns a tuple with the SshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeys

`func (o *ProjectOrHref) SetSshKeys(v []Href)`

SetSshKeys sets SshKeys field to given value.

### HasSshKeys

`func (o *ProjectOrHref) HasSshKeys() bool`

HasSshKeys returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ProjectOrHref) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ProjectOrHref) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ProjectOrHref) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ProjectOrHref) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVolumes

`func (o *ProjectOrHref) GetVolumes() []Href`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *ProjectOrHref) GetVolumesOk() (*[]Href, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *ProjectOrHref) SetVolumes(v []Href)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *ProjectOrHref) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


