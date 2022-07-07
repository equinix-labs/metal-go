# FindDeviceById200ResponseCreatedBy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvatarThumbUrl** | Pointer to **string** | Avatar thumbnail URL of the User | [optional] 
**CreatedAt** | Pointer to **time.Time** | When the user was created | [optional] 
**Email** | Pointer to **string** | Primary email address of the User | [optional] 
**FirstName** | Pointer to **string** | First name of the User | [optional] 
**FullName** | Pointer to **string** | Full name of the User | [optional] 
**Href** | Pointer to **string** | API URL uniquely representing the User | [optional] 
**Id** | **string** | ID of the User | 
**LastName** | Pointer to **string** | Last name of the User | [optional] 
**ShortId** | **string** | Short ID of the User | 
**UpdatedAt** | Pointer to **time.Time** | When the user details were last updated | [optional] 

## Methods

### NewFindDeviceById200ResponseCreatedBy

`func NewFindDeviceById200ResponseCreatedBy(id string, shortId string, ) *FindDeviceById200ResponseCreatedBy`

NewFindDeviceById200ResponseCreatedBy instantiates a new FindDeviceById200ResponseCreatedBy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindDeviceById200ResponseCreatedByWithDefaults

`func NewFindDeviceById200ResponseCreatedByWithDefaults() *FindDeviceById200ResponseCreatedBy`

NewFindDeviceById200ResponseCreatedByWithDefaults instantiates a new FindDeviceById200ResponseCreatedBy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvatarThumbUrl

`func (o *FindDeviceById200ResponseCreatedBy) GetAvatarThumbUrl() string`

GetAvatarThumbUrl returns the AvatarThumbUrl field if non-nil, zero value otherwise.

### GetAvatarThumbUrlOk

`func (o *FindDeviceById200ResponseCreatedBy) GetAvatarThumbUrlOk() (*string, bool)`

GetAvatarThumbUrlOk returns a tuple with the AvatarThumbUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarThumbUrl

`func (o *FindDeviceById200ResponseCreatedBy) SetAvatarThumbUrl(v string)`

SetAvatarThumbUrl sets AvatarThumbUrl field to given value.

### HasAvatarThumbUrl

`func (o *FindDeviceById200ResponseCreatedBy) HasAvatarThumbUrl() bool`

HasAvatarThumbUrl returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FindDeviceById200ResponseCreatedBy) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FindDeviceById200ResponseCreatedBy) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FindDeviceById200ResponseCreatedBy) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FindDeviceById200ResponseCreatedBy) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEmail

`func (o *FindDeviceById200ResponseCreatedBy) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *FindDeviceById200ResponseCreatedBy) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *FindDeviceById200ResponseCreatedBy) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *FindDeviceById200ResponseCreatedBy) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFirstName

`func (o *FindDeviceById200ResponseCreatedBy) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *FindDeviceById200ResponseCreatedBy) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *FindDeviceById200ResponseCreatedBy) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *FindDeviceById200ResponseCreatedBy) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetFullName

`func (o *FindDeviceById200ResponseCreatedBy) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *FindDeviceById200ResponseCreatedBy) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *FindDeviceById200ResponseCreatedBy) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *FindDeviceById200ResponseCreatedBy) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetHref

`func (o *FindDeviceById200ResponseCreatedBy) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *FindDeviceById200ResponseCreatedBy) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *FindDeviceById200ResponseCreatedBy) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *FindDeviceById200ResponseCreatedBy) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *FindDeviceById200ResponseCreatedBy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FindDeviceById200ResponseCreatedBy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FindDeviceById200ResponseCreatedBy) SetId(v string)`

SetId sets Id field to given value.


### GetLastName

`func (o *FindDeviceById200ResponseCreatedBy) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *FindDeviceById200ResponseCreatedBy) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *FindDeviceById200ResponseCreatedBy) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *FindDeviceById200ResponseCreatedBy) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetShortId

`func (o *FindDeviceById200ResponseCreatedBy) GetShortId() string`

GetShortId returns the ShortId field if non-nil, zero value otherwise.

### GetShortIdOk

`func (o *FindDeviceById200ResponseCreatedBy) GetShortIdOk() (*string, bool)`

GetShortIdOk returns a tuple with the ShortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortId

`func (o *FindDeviceById200ResponseCreatedBy) SetShortId(v string)`

SetShortId sets ShortId field to given value.


### GetUpdatedAt

`func (o *FindDeviceById200ResponseCreatedBy) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FindDeviceById200ResponseCreatedBy) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FindDeviceById200ResponseCreatedBy) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FindDeviceById200ResponseCreatedBy) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


