# CreateUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Avatar** | Pointer to ***os.File** |  | [optional] 
**CompanyName** | Pointer to **string** |  | [optional] 
**CompanyUrl** | Pointer to **string** |  | [optional] 
**Customdata** | Pointer to **map[string]interface{}** |  | [optional] 
**Emails** | [**[]CreateUserRequestEmailsInner**](CreateUserRequestEmailsInner.md) |  | 
**FirstName** | **string** |  | 
**LastName** | **string** |  | 
**Level** | Pointer to **string** |  | [optional] 
**Locked** | Pointer to **bool** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**PhoneNumber** | Pointer to **string** |  | [optional] 
**SocialAccounts** | Pointer to **map[string]interface{}** |  | [optional] 
**Timezone** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**TwoFactorAuth** | Pointer to **string** |  | [optional] 
**VerifiedAt** | Pointer to **time.Time** |  | [optional] 
**InvitationId** | Pointer to **string** |  | [optional] 
**Nonce** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateUserRequest

`func NewCreateUserRequest(emails []CreateUserRequestEmailsInner, firstName string, lastName string, ) *CreateUserRequest`

NewCreateUserRequest instantiates a new CreateUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUserRequestWithDefaults

`func NewCreateUserRequestWithDefaults() *CreateUserRequest`

NewCreateUserRequestWithDefaults instantiates a new CreateUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvatar

`func (o *CreateUserRequest) GetAvatar() *os.File`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *CreateUserRequest) GetAvatarOk() (**os.File, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *CreateUserRequest) SetAvatar(v *os.File)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *CreateUserRequest) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetCompanyName

`func (o *CreateUserRequest) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *CreateUserRequest) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *CreateUserRequest) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *CreateUserRequest) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetCompanyUrl

`func (o *CreateUserRequest) GetCompanyUrl() string`

GetCompanyUrl returns the CompanyUrl field if non-nil, zero value otherwise.

### GetCompanyUrlOk

`func (o *CreateUserRequest) GetCompanyUrlOk() (*string, bool)`

GetCompanyUrlOk returns a tuple with the CompanyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyUrl

`func (o *CreateUserRequest) SetCompanyUrl(v string)`

SetCompanyUrl sets CompanyUrl field to given value.

### HasCompanyUrl

`func (o *CreateUserRequest) HasCompanyUrl() bool`

HasCompanyUrl returns a boolean if a field has been set.

### GetCustomdata

`func (o *CreateUserRequest) GetCustomdata() map[string]interface{}`

GetCustomdata returns the Customdata field if non-nil, zero value otherwise.

### GetCustomdataOk

`func (o *CreateUserRequest) GetCustomdataOk() (*map[string]interface{}, bool)`

GetCustomdataOk returns a tuple with the Customdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomdata

`func (o *CreateUserRequest) SetCustomdata(v map[string]interface{})`

SetCustomdata sets Customdata field to given value.

### HasCustomdata

`func (o *CreateUserRequest) HasCustomdata() bool`

HasCustomdata returns a boolean if a field has been set.

### GetEmails

`func (o *CreateUserRequest) GetEmails() []CreateUserRequestEmailsInner`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *CreateUserRequest) GetEmailsOk() (*[]CreateUserRequestEmailsInner, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *CreateUserRequest) SetEmails(v []CreateUserRequestEmailsInner)`

SetEmails sets Emails field to given value.


### GetFirstName

`func (o *CreateUserRequest) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *CreateUserRequest) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *CreateUserRequest) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *CreateUserRequest) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *CreateUserRequest) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *CreateUserRequest) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetLevel

`func (o *CreateUserRequest) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *CreateUserRequest) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *CreateUserRequest) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *CreateUserRequest) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetLocked

`func (o *CreateUserRequest) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *CreateUserRequest) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *CreateUserRequest) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *CreateUserRequest) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetPassword

`func (o *CreateUserRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateUserRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateUserRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CreateUserRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *CreateUserRequest) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *CreateUserRequest) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *CreateUserRequest) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *CreateUserRequest) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetSocialAccounts

`func (o *CreateUserRequest) GetSocialAccounts() map[string]interface{}`

GetSocialAccounts returns the SocialAccounts field if non-nil, zero value otherwise.

### GetSocialAccountsOk

`func (o *CreateUserRequest) GetSocialAccountsOk() (*map[string]interface{}, bool)`

GetSocialAccountsOk returns a tuple with the SocialAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialAccounts

`func (o *CreateUserRequest) SetSocialAccounts(v map[string]interface{})`

SetSocialAccounts sets SocialAccounts field to given value.

### HasSocialAccounts

`func (o *CreateUserRequest) HasSocialAccounts() bool`

HasSocialAccounts returns a boolean if a field has been set.

### GetTimezone

`func (o *CreateUserRequest) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *CreateUserRequest) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *CreateUserRequest) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *CreateUserRequest) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetTitle

`func (o *CreateUserRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateUserRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateUserRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CreateUserRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTwoFactorAuth

`func (o *CreateUserRequest) GetTwoFactorAuth() string`

GetTwoFactorAuth returns the TwoFactorAuth field if non-nil, zero value otherwise.

### GetTwoFactorAuthOk

`func (o *CreateUserRequest) GetTwoFactorAuthOk() (*string, bool)`

GetTwoFactorAuthOk returns a tuple with the TwoFactorAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFactorAuth

`func (o *CreateUserRequest) SetTwoFactorAuth(v string)`

SetTwoFactorAuth sets TwoFactorAuth field to given value.

### HasTwoFactorAuth

`func (o *CreateUserRequest) HasTwoFactorAuth() bool`

HasTwoFactorAuth returns a boolean if a field has been set.

### GetVerifiedAt

`func (o *CreateUserRequest) GetVerifiedAt() time.Time`

GetVerifiedAt returns the VerifiedAt field if non-nil, zero value otherwise.

### GetVerifiedAtOk

`func (o *CreateUserRequest) GetVerifiedAtOk() (*time.Time, bool)`

GetVerifiedAtOk returns a tuple with the VerifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedAt

`func (o *CreateUserRequest) SetVerifiedAt(v time.Time)`

SetVerifiedAt sets VerifiedAt field to given value.

### HasVerifiedAt

`func (o *CreateUserRequest) HasVerifiedAt() bool`

HasVerifiedAt returns a boolean if a field has been set.

### GetInvitationId

`func (o *CreateUserRequest) GetInvitationId() string`

GetInvitationId returns the InvitationId field if non-nil, zero value otherwise.

### GetInvitationIdOk

`func (o *CreateUserRequest) GetInvitationIdOk() (*string, bool)`

GetInvitationIdOk returns a tuple with the InvitationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitationId

`func (o *CreateUserRequest) SetInvitationId(v string)`

SetInvitationId sets InvitationId field to given value.

### HasInvitationId

`func (o *CreateUserRequest) HasInvitationId() bool`

HasInvitationId returns a boolean if a field has been set.

### GetNonce

`func (o *CreateUserRequest) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *CreateUserRequest) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *CreateUserRequest) SetNonce(v string)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *CreateUserRequest) HasNonce() bool`

HasNonce returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


