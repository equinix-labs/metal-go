# FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf

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

### NewFindProjectAPIKeys200ResponseApiKeysInnerUserAllOf

`func NewFindProjectAPIKeys200ResponseApiKeysInnerUserAllOf() *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf`

NewFindProjectAPIKeys200ResponseApiKeysInnerUserAllOf instantiates a new FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindProjectAPIKeys200ResponseApiKeysInnerUserAllOfWithDefaults

`func NewFindProjectAPIKeys200ResponseApiKeysInnerUserAllOfWithDefaults() *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf`

NewFindProjectAPIKeys200ResponseApiKeysInnerUserAllOfWithDefaults instantiates a new FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvatarThumbUrl

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) GetAvatarThumbUrl() string`

GetAvatarThumbUrl returns the AvatarThumbUrl field if non-nil, zero value otherwise.

### GetAvatarThumbUrlOk

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) GetAvatarThumbUrlOk() (*string, bool)`

GetAvatarThumbUrlOk returns a tuple with the AvatarThumbUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarThumbUrl

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) SetAvatarThumbUrl(v string)`

SetAvatarThumbUrl sets AvatarThumbUrl field to given value.

### HasAvatarThumbUrl

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) HasAvatarThumbUrl() bool`

HasAvatarThumbUrl returns a boolean if a field has been set.

### GetAvatarUrl

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCustomdata

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) GetCustomdata() map[string]interface{}`

GetCustomdata returns the Customdata field if non-nil, zero value otherwise.

### GetCustomdataOk

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) GetCustomdataOk() (*map[string]interface{}, bool)`

GetCustomdataOk returns a tuple with the Customdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomdata

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) SetCustomdata(v map[string]interface{})`

SetCustomdata sets Customdata field to given value.

### HasCustomdata

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) HasCustomdata() bool`

HasCustomdata returns a boolean if a field has been set.

### GetEmail

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEmails

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) GetEmails() []FindBatchById200ResponseDevicesInner`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) GetEmailsOk() (*[]FindBatchById200ResponseDevicesInner, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) SetEmails(v []FindBatchById200ResponseDevicesInner)`

SetEmails sets Emails field to given value.

### HasEmails

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) HasEmails() bool`

HasEmails returns a boolean if a field has been set.

### GetFirstName

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetFraudScore

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) GetFraudScore() string`

GetFraudScore returns the FraudScore field if non-nil, zero value otherwise.

### GetFraudScoreOk

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) GetFraudScoreOk() (*string, bool)`

GetFraudScoreOk returns a tuple with the FraudScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFraudScore

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) SetFraudScore(v string)`

SetFraudScore sets FraudScore field to given value.

### HasFraudScore

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) HasFraudScore() bool`

HasFraudScore returns a boolean if a field has been set.

### GetFullName

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetHref

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastLoginAt

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) GetLastLoginAt() time.Time`

GetLastLoginAt returns the LastLoginAt field if non-nil, zero value otherwise.

### GetLastLoginAtOk

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) GetLastLoginAtOk() (*time.Time, bool)`

GetLastLoginAtOk returns a tuple with the LastLoginAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginAt

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) SetLastLoginAt(v time.Time)`

SetLastLoginAt sets LastLoginAt field to given value.

### HasLastLoginAt

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) HasLastLoginAt() bool`

HasLastLoginAt returns a boolean if a field has been set.

### GetLastName

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetMaxOrganizations

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) GetMaxOrganizations() int32`

GetMaxOrganizations returns the MaxOrganizations field if non-nil, zero value otherwise.

### GetMaxOrganizationsOk

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) GetMaxOrganizationsOk() (*int32, bool)`

GetMaxOrganizationsOk returns a tuple with the MaxOrganizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxOrganizations

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) SetMaxOrganizations(v int32)`

SetMaxOrganizations sets MaxOrganizations field to given value.

### HasMaxOrganizations

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) HasMaxOrganizations() bool`

HasMaxOrganizations returns a boolean if a field has been set.

### GetMaxProjects

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) GetMaxProjects() int32`

GetMaxProjects returns the MaxProjects field if non-nil, zero value otherwise.

### GetMaxProjectsOk

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) GetMaxProjectsOk() (*int32, bool)`

GetMaxProjectsOk returns a tuple with the MaxProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxProjects

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) SetMaxProjects(v int32)`

SetMaxProjects sets MaxProjects field to given value.

### HasMaxProjects

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) HasMaxProjects() bool`

HasMaxProjects returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetShortId

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) GetShortId() string`

GetShortId returns the ShortId field if non-nil, zero value otherwise.

### GetShortIdOk

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) GetShortIdOk() (*string, bool)`

GetShortIdOk returns a tuple with the ShortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortId

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) SetShortId(v string)`

SetShortId sets ShortId field to given value.

### HasShortId

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) HasShortId() bool`

HasShortId returns a boolean if a field has been set.

### GetTimezone

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetTwoFactorAuth

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) GetTwoFactorAuth() string`

GetTwoFactorAuth returns the TwoFactorAuth field if non-nil, zero value otherwise.

### GetTwoFactorAuthOk

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) GetTwoFactorAuthOk() (*string, bool)`

GetTwoFactorAuthOk returns a tuple with the TwoFactorAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFactorAuth

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) SetTwoFactorAuth(v string)`

SetTwoFactorAuth sets TwoFactorAuth field to given value.

### HasTwoFactorAuth

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) HasTwoFactorAuth() bool`

HasTwoFactorAuth returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


