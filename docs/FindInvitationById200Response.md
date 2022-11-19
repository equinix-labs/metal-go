# FindInvitationById200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Invitation** | Pointer to [**FindBatchById200ResponseDevicesInner**](FindBatchById200ResponseDevicesInner.md) |  | [optional] 
**InvitedBy** | Pointer to [**FindBatchById200ResponseDevicesInner**](FindBatchById200ResponseDevicesInner.md) |  | [optional] 
**Invitee** | Pointer to **string** |  | [optional] 
**Nonce** | Pointer to **string** |  | [optional] 
**Organization** | Pointer to [**FindBatchById200ResponseDevicesInner**](FindBatchById200ResponseDevicesInner.md) |  | [optional] 
**Projects** | Pointer to [**[]FindBatchById200ResponseDevicesInner**](FindBatchById200ResponseDevicesInner.md) |  | [optional] 
**Roles** | Pointer to **[]string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewFindInvitationById200Response

`func NewFindInvitationById200Response() *FindInvitationById200Response`

NewFindInvitationById200Response instantiates a new FindInvitationById200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindInvitationById200ResponseWithDefaults

`func NewFindInvitationById200ResponseWithDefaults() *FindInvitationById200Response`

NewFindInvitationById200ResponseWithDefaults instantiates a new FindInvitationById200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *FindInvitationById200Response) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FindInvitationById200Response) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FindInvitationById200Response) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FindInvitationById200Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetHref

`func (o *FindInvitationById200Response) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *FindInvitationById200Response) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *FindInvitationById200Response) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *FindInvitationById200Response) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *FindInvitationById200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FindInvitationById200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FindInvitationById200Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FindInvitationById200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInvitation

`func (o *FindInvitationById200Response) GetInvitation() FindBatchById200ResponseDevicesInner`

GetInvitation returns the Invitation field if non-nil, zero value otherwise.

### GetInvitationOk

`func (o *FindInvitationById200Response) GetInvitationOk() (*FindBatchById200ResponseDevicesInner, bool)`

GetInvitationOk returns a tuple with the Invitation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitation

`func (o *FindInvitationById200Response) SetInvitation(v FindBatchById200ResponseDevicesInner)`

SetInvitation sets Invitation field to given value.

### HasInvitation

`func (o *FindInvitationById200Response) HasInvitation() bool`

HasInvitation returns a boolean if a field has been set.

### GetInvitedBy

`func (o *FindInvitationById200Response) GetInvitedBy() FindBatchById200ResponseDevicesInner`

GetInvitedBy returns the InvitedBy field if non-nil, zero value otherwise.

### GetInvitedByOk

`func (o *FindInvitationById200Response) GetInvitedByOk() (*FindBatchById200ResponseDevicesInner, bool)`

GetInvitedByOk returns a tuple with the InvitedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitedBy

`func (o *FindInvitationById200Response) SetInvitedBy(v FindBatchById200ResponseDevicesInner)`

SetInvitedBy sets InvitedBy field to given value.

### HasInvitedBy

`func (o *FindInvitationById200Response) HasInvitedBy() bool`

HasInvitedBy returns a boolean if a field has been set.

### GetInvitee

`func (o *FindInvitationById200Response) GetInvitee() string`

GetInvitee returns the Invitee field if non-nil, zero value otherwise.

### GetInviteeOk

`func (o *FindInvitationById200Response) GetInviteeOk() (*string, bool)`

GetInviteeOk returns a tuple with the Invitee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitee

`func (o *FindInvitationById200Response) SetInvitee(v string)`

SetInvitee sets Invitee field to given value.

### HasInvitee

`func (o *FindInvitationById200Response) HasInvitee() bool`

HasInvitee returns a boolean if a field has been set.

### GetNonce

`func (o *FindInvitationById200Response) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *FindInvitationById200Response) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *FindInvitationById200Response) SetNonce(v string)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *FindInvitationById200Response) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

### GetOrganization

`func (o *FindInvitationById200Response) GetOrganization() FindBatchById200ResponseDevicesInner`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *FindInvitationById200Response) GetOrganizationOk() (*FindBatchById200ResponseDevicesInner, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *FindInvitationById200Response) SetOrganization(v FindBatchById200ResponseDevicesInner)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *FindInvitationById200Response) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetProjects

`func (o *FindInvitationById200Response) GetProjects() []FindBatchById200ResponseDevicesInner`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *FindInvitationById200Response) GetProjectsOk() (*[]FindBatchById200ResponseDevicesInner, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *FindInvitationById200Response) SetProjects(v []FindBatchById200ResponseDevicesInner)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *FindInvitationById200Response) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### GetRoles

`func (o *FindInvitationById200Response) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *FindInvitationById200Response) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *FindInvitationById200Response) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *FindInvitationById200Response) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *FindInvitationById200Response) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FindInvitationById200Response) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FindInvitationById200Response) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FindInvitationById200Response) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


