# Invitation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Roles** | Pointer to **[]string** |  | [optional] 
**Invitee** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**InvitedBy** | Pointer to [**Href**](Href.md) |  | [optional] 
**Organization** | Pointer to [**Href**](Href.md) |  | [optional] 
**ProjectsIds** | Pointer to **[]string** |  | [optional] 
**Invitation** | Pointer to [**Href**](Href.md) |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 

## Methods

### NewInvitation

`func NewInvitation() *Invitation`

NewInvitation instantiates a new Invitation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvitationWithDefaults

`func NewInvitationWithDefaults() *Invitation`

NewInvitationWithDefaults instantiates a new Invitation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Invitation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Invitation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Invitation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Invitation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRoles

`func (o *Invitation) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *Invitation) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *Invitation) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *Invitation) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetInvitee

`func (o *Invitation) GetInvitee() string`

GetInvitee returns the Invitee field if non-nil, zero value otherwise.

### GetInviteeOk

`func (o *Invitation) GetInviteeOk() (*string, bool)`

GetInviteeOk returns a tuple with the Invitee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitee

`func (o *Invitation) SetInvitee(v string)`

SetInvitee sets Invitee field to given value.

### HasInvitee

`func (o *Invitation) HasInvitee() bool`

HasInvitee returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Invitation) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Invitation) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Invitation) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Invitation) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Invitation) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Invitation) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Invitation) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Invitation) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetInvitedBy

`func (o *Invitation) GetInvitedBy() Href`

GetInvitedBy returns the InvitedBy field if non-nil, zero value otherwise.

### GetInvitedByOk

`func (o *Invitation) GetInvitedByOk() (*Href, bool)`

GetInvitedByOk returns a tuple with the InvitedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitedBy

`func (o *Invitation) SetInvitedBy(v Href)`

SetInvitedBy sets InvitedBy field to given value.

### HasInvitedBy

`func (o *Invitation) HasInvitedBy() bool`

HasInvitedBy returns a boolean if a field has been set.

### GetOrganization

`func (o *Invitation) GetOrganization() Href`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *Invitation) GetOrganizationOk() (*Href, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *Invitation) SetOrganization(v Href)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *Invitation) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetProjectsIds

`func (o *Invitation) GetProjectsIds() []string`

GetProjectsIds returns the ProjectsIds field if non-nil, zero value otherwise.

### GetProjectsIdsOk

`func (o *Invitation) GetProjectsIdsOk() (*[]string, bool)`

GetProjectsIdsOk returns a tuple with the ProjectsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectsIds

`func (o *Invitation) SetProjectsIds(v []string)`

SetProjectsIds sets ProjectsIds field to given value.

### HasProjectsIds

`func (o *Invitation) HasProjectsIds() bool`

HasProjectsIds returns a boolean if a field has been set.

### GetInvitation

`func (o *Invitation) GetInvitation() Href`

GetInvitation returns the Invitation field if non-nil, zero value otherwise.

### GetInvitationOk

`func (o *Invitation) GetInvitationOk() (*Href, bool)`

GetInvitationOk returns a tuple with the Invitation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitation

`func (o *Invitation) SetInvitation(v Href)`

SetInvitation sets Invitation field to given value.

### HasInvitation

`func (o *Invitation) HasInvitation() bool`

HasInvitation returns a boolean if a field has been set.

### GetHref

`func (o *Invitation) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *Invitation) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *Invitation) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *Invitation) HasHref() bool`

HasHref returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


