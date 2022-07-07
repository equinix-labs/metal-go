# FindProjectAPIKeys200ResponseApiKeysInnerUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvatarThumbUrl** | Pointer to **string** |  | [optional] 
**AvatarUrl** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Customdata** | Pointer to **map[string]interface{}** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Emails** | Pointer to [**[]FindBatchById200ResponseDevicesInner**](FindBatchById200ResponseDevicesInner.md) |  | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**FraudScore** | Pointer to **string** |  | [optional] 
**FullName** | Pointer to **string** |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LastLoginAt** | Pointer to **time.Time** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**MaxOrganizations** | Pointer to **int32** |  | [optional] 
**MaxProjects** | Pointer to **int32** |  | [optional] 
**PhoneNumber** | Pointer to **string** |  | [optional] 
**ShortId** | Pointer to **string** |  | [optional] 
**Timezone** | Pointer to **string** |  | [optional] 
**TwoFactorAuth** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewFindProjectAPIKeys200ResponseApiKeysInnerUser

`func NewFindProjectAPIKeys200ResponseApiKeysInnerUser() *FindProjectAPIKeys200ResponseApiKeysInnerUser`

NewFindProjectAPIKeys200ResponseApiKeysInnerUser instantiates a new FindProjectAPIKeys200ResponseApiKeysInnerUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindProjectAPIKeys200ResponseApiKeysInnerUserWithDefaults

`func NewFindProjectAPIKeys200ResponseApiKeysInnerUserWithDefaults() *FindProjectAPIKeys200ResponseApiKeysInnerUser`

NewFindProjectAPIKeys200ResponseApiKeysInnerUserWithDefaults instantiates a new FindProjectAPIKeys200ResponseApiKeysInnerUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvatarThumbUrl

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) GetAvatarThumbUrl() string`

GetAvatarThumbUrl returns the AvatarThumbUrl field if non-nil, zero value otherwise.

### GetAvatarThumbUrlOk

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) GetAvatarThumbUrlOk() (*string, bool)`

GetAvatarThumbUrlOk returns a tuple with the AvatarThumbUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarThumbUrl

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) SetAvatarThumbUrl(v string)`

SetAvatarThumbUrl sets AvatarThumbUrl field to given value.

### HasAvatarThumbUrl

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) HasAvatarThumbUrl() bool`

HasAvatarThumbUrl returns a boolean if a field has been set.

### GetAvatarUrl

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCustomdata

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) GetCustomdata() map[string]interface{}`

GetCustomdata returns the Customdata field if non-nil, zero value otherwise.

### GetCustomdataOk

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) GetCustomdataOk() (*map[string]interface{}, bool)`

GetCustomdataOk returns a tuple with the Customdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomdata

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) SetCustomdata(v map[string]interface{})`

SetCustomdata sets Customdata field to given value.

### HasCustomdata

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) HasCustomdata() bool`

HasCustomdata returns a boolean if a field has been set.

### GetEmail

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEmails

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) GetEmails() []FindBatchById200ResponseDevicesInner`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) GetEmailsOk() (*[]FindBatchById200ResponseDevicesInner, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) SetEmails(v []FindBatchById200ResponseDevicesInner)`

SetEmails sets Emails field to given value.

### HasEmails

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) HasEmails() bool`

HasEmails returns a boolean if a field has been set.

### GetFirstName

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetFraudScore

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) GetFraudScore() string`

GetFraudScore returns the FraudScore field if non-nil, zero value otherwise.

### GetFraudScoreOk

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) GetFraudScoreOk() (*string, bool)`

GetFraudScoreOk returns a tuple with the FraudScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFraudScore

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) SetFraudScore(v string)`

SetFraudScore sets FraudScore field to given value.

### HasFraudScore

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) HasFraudScore() bool`

HasFraudScore returns a boolean if a field has been set.

### GetFullName

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetHref

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastLoginAt

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) GetLastLoginAt() time.Time`

GetLastLoginAt returns the LastLoginAt field if non-nil, zero value otherwise.

### GetLastLoginAtOk

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) GetLastLoginAtOk() (*time.Time, bool)`

GetLastLoginAtOk returns a tuple with the LastLoginAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginAt

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) SetLastLoginAt(v time.Time)`

SetLastLoginAt sets LastLoginAt field to given value.

### HasLastLoginAt

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) HasLastLoginAt() bool`

HasLastLoginAt returns a boolean if a field has been set.

### GetLastName

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetMaxOrganizations

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) GetMaxOrganizations() int32`

GetMaxOrganizations returns the MaxOrganizations field if non-nil, zero value otherwise.

### GetMaxOrganizationsOk

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) GetMaxOrganizationsOk() (*int32, bool)`

GetMaxOrganizationsOk returns a tuple with the MaxOrganizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxOrganizations

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) SetMaxOrganizations(v int32)`

SetMaxOrganizations sets MaxOrganizations field to given value.

### HasMaxOrganizations

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) HasMaxOrganizations() bool`

HasMaxOrganizations returns a boolean if a field has been set.

### GetMaxProjects

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) GetMaxProjects() int32`

GetMaxProjects returns the MaxProjects field if non-nil, zero value otherwise.

### GetMaxProjectsOk

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) GetMaxProjectsOk() (*int32, bool)`

GetMaxProjectsOk returns a tuple with the MaxProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxProjects

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) SetMaxProjects(v int32)`

SetMaxProjects sets MaxProjects field to given value.

### HasMaxProjects

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) HasMaxProjects() bool`

HasMaxProjects returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetShortId

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) GetShortId() string`

GetShortId returns the ShortId field if non-nil, zero value otherwise.

### GetShortIdOk

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) GetShortIdOk() (*string, bool)`

GetShortIdOk returns a tuple with the ShortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortId

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) SetShortId(v string)`

SetShortId sets ShortId field to given value.

### HasShortId

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) HasShortId() bool`

HasShortId returns a boolean if a field has been set.

### GetTimezone

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetTwoFactorAuth

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) GetTwoFactorAuth() string`

GetTwoFactorAuth returns the TwoFactorAuth field if non-nil, zero value otherwise.

### GetTwoFactorAuthOk

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) GetTwoFactorAuthOk() (*string, bool)`

GetTwoFactorAuthOk returns a tuple with the TwoFactorAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFactorAuth

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) SetTwoFactorAuth(v string)`

SetTwoFactorAuth sets TwoFactorAuth field to given value.

### HasTwoFactorAuth

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) HasTwoFactorAuth() bool`

HasTwoFactorAuth returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUser) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


