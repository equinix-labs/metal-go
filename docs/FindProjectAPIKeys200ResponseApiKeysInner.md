# FindProjectAPIKeys200ResponseApiKeysInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** | Available only for API keys | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Project** | Pointer to [**FindProjectAPIKeys200ResponseApiKeysInnerProject**](FindProjectAPIKeys200ResponseApiKeysInnerProject.md) |  | [optional] 
**ReadOnly** | Pointer to **bool** |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**User** | Pointer to [**FindProjectAPIKeys200ResponseApiKeysInnerUser**](FindProjectAPIKeys200ResponseApiKeysInnerUser.md) |  | [optional] 

## Methods

### NewFindProjectAPIKeys200ResponseApiKeysInner

`func NewFindProjectAPIKeys200ResponseApiKeysInner() *FindProjectAPIKeys200ResponseApiKeysInner`

NewFindProjectAPIKeys200ResponseApiKeysInner instantiates a new FindProjectAPIKeys200ResponseApiKeysInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindProjectAPIKeys200ResponseApiKeysInnerWithDefaults

`func NewFindProjectAPIKeys200ResponseApiKeysInnerWithDefaults() *FindProjectAPIKeys200ResponseApiKeysInner`

NewFindProjectAPIKeys200ResponseApiKeysInnerWithDefaults instantiates a new FindProjectAPIKeys200ResponseApiKeysInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *FindProjectAPIKeys200ResponseApiKeysInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FindProjectAPIKeys200ResponseApiKeysInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FindProjectAPIKeys200ResponseApiKeysInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FindProjectAPIKeys200ResponseApiKeysInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *FindProjectAPIKeys200ResponseApiKeysInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FindProjectAPIKeys200ResponseApiKeysInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FindProjectAPIKeys200ResponseApiKeysInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FindProjectAPIKeys200ResponseApiKeysInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *FindProjectAPIKeys200ResponseApiKeysInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FindProjectAPIKeys200ResponseApiKeysInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FindProjectAPIKeys200ResponseApiKeysInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FindProjectAPIKeys200ResponseApiKeysInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProject

`func (o *FindProjectAPIKeys200ResponseApiKeysInner) GetProject() FindProjectAPIKeys200ResponseApiKeysInnerProject`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *FindProjectAPIKeys200ResponseApiKeysInner) GetProjectOk() (*FindProjectAPIKeys200ResponseApiKeysInnerProject, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *FindProjectAPIKeys200ResponseApiKeysInner) SetProject(v FindProjectAPIKeys200ResponseApiKeysInnerProject)`

SetProject sets Project field to given value.

### HasProject

`func (o *FindProjectAPIKeys200ResponseApiKeysInner) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetReadOnly

`func (o *FindProjectAPIKeys200ResponseApiKeysInner) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *FindProjectAPIKeys200ResponseApiKeysInner) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *FindProjectAPIKeys200ResponseApiKeysInner) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *FindProjectAPIKeys200ResponseApiKeysInner) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetToken

`func (o *FindProjectAPIKeys200ResponseApiKeysInner) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *FindProjectAPIKeys200ResponseApiKeysInner) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *FindProjectAPIKeys200ResponseApiKeysInner) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *FindProjectAPIKeys200ResponseApiKeysInner) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *FindProjectAPIKeys200ResponseApiKeysInner) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FindProjectAPIKeys200ResponseApiKeysInner) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FindProjectAPIKeys200ResponseApiKeysInner) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FindProjectAPIKeys200ResponseApiKeysInner) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUser

`func (o *FindProjectAPIKeys200ResponseApiKeysInner) GetUser() FindProjectAPIKeys200ResponseApiKeysInnerUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *FindProjectAPIKeys200ResponseApiKeysInner) GetUserOk() (*FindProjectAPIKeys200ResponseApiKeysInnerUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *FindProjectAPIKeys200ResponseApiKeysInner) SetUser(v FindProjectAPIKeys200ResponseApiKeysInnerUser)`

SetUser sets User field to given value.

### HasUser

`func (o *FindProjectAPIKeys200ResponseApiKeysInner) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


