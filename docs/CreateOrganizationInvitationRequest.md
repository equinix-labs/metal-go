# CreateOrganizationInvitationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Invitee** | **string** |  | 
**Message** | Pointer to **string** |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**ProjectsIds** | Pointer to **[]string** |  | [optional] 
**Roles** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCreateOrganizationInvitationRequest

`func NewCreateOrganizationInvitationRequest(invitee string, ) *CreateOrganizationInvitationRequest`

NewCreateOrganizationInvitationRequest instantiates a new CreateOrganizationInvitationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrganizationInvitationRequestWithDefaults

`func NewCreateOrganizationInvitationRequestWithDefaults() *CreateOrganizationInvitationRequest`

NewCreateOrganizationInvitationRequestWithDefaults instantiates a new CreateOrganizationInvitationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvitee

`func (o *CreateOrganizationInvitationRequest) GetInvitee() string`

GetInvitee returns the Invitee field if non-nil, zero value otherwise.

### GetInviteeOk

`func (o *CreateOrganizationInvitationRequest) GetInviteeOk() (*string, bool)`

GetInviteeOk returns a tuple with the Invitee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitee

`func (o *CreateOrganizationInvitationRequest) SetInvitee(v string)`

SetInvitee sets Invitee field to given value.


### GetMessage

`func (o *CreateOrganizationInvitationRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CreateOrganizationInvitationRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CreateOrganizationInvitationRequest) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CreateOrganizationInvitationRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetOrganizationId

`func (o *CreateOrganizationInvitationRequest) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *CreateOrganizationInvitationRequest) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *CreateOrganizationInvitationRequest) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *CreateOrganizationInvitationRequest) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetProjectsIds

`func (o *CreateOrganizationInvitationRequest) GetProjectsIds() []string`

GetProjectsIds returns the ProjectsIds field if non-nil, zero value otherwise.

### GetProjectsIdsOk

`func (o *CreateOrganizationInvitationRequest) GetProjectsIdsOk() (*[]string, bool)`

GetProjectsIdsOk returns a tuple with the ProjectsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectsIds

`func (o *CreateOrganizationInvitationRequest) SetProjectsIds(v []string)`

SetProjectsIds sets ProjectsIds field to given value.

### HasProjectsIds

`func (o *CreateOrganizationInvitationRequest) HasProjectsIds() bool`

HasProjectsIds returns a boolean if a field has been set.

### GetRoles

`func (o *CreateOrganizationInvitationRequest) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *CreateOrganizationInvitationRequest) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *CreateOrganizationInvitationRequest) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *CreateOrganizationInvitationRequest) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


