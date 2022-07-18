# FindOrganizationPaymentMethods200ResponsePaymentMethodsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingAddress** | Pointer to [**FindOrganizationPaymentMethods200ResponsePaymentMethodsInnerBillingAddress**](FindOrganizationPaymentMethods200ResponsePaymentMethodsInnerBillingAddress.md) |  | [optional] 
**CardType** | Pointer to **string** |  | [optional] 
**CardholderName** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedByUser** | Pointer to [**FindBatchById200ResponseDevicesInner**](FindBatchById200ResponseDevicesInner.md) |  | [optional] 
**Default** | Pointer to **bool** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**ExpirationMonth** | Pointer to **string** |  | [optional] 
**ExpirationYear** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Organization** | Pointer to [**FindBatchById200ResponseDevicesInner**](FindBatchById200ResponseDevicesInner.md) |  | [optional] 
**Projects** | Pointer to [**[]FindBatchById200ResponseDevicesInner**](FindBatchById200ResponseDevicesInner.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewFindOrganizationPaymentMethods200ResponsePaymentMethodsInner

`func NewFindOrganizationPaymentMethods200ResponsePaymentMethodsInner() *FindOrganizationPaymentMethods200ResponsePaymentMethodsInner`

NewFindOrganizationPaymentMethods200ResponsePaymentMethodsInner instantiates a new FindOrganizationPaymentMethods200ResponsePaymentMethodsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindOrganizationPaymentMethods200ResponsePaymentMethodsInnerWithDefaults

`func NewFindOrganizationPaymentMethods200ResponsePaymentMethodsInnerWithDefaults() *FindOrganizationPaymentMethods200ResponsePaymentMethodsInner`

NewFindOrganizationPaymentMethods200ResponsePaymentMethodsInnerWithDefaults instantiates a new FindOrganizationPaymentMethods200ResponsePaymentMethodsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingAddress

`func (o *FindOrganizationPaymentMethods200ResponsePaymentMethodsInner) GetBillingAddress() FindOrganizationPaymentMethods200ResponsePaymentMethodsInnerBillingAddress`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *FindOrganizationPaymentMethods200ResponsePaymentMethodsInner) GetBillingAddressOk() (*FindOrganizationPaymentMethods200ResponsePaymentMethodsInnerBillingAddress, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *FindOrganizationPaymentMethods200ResponsePaymentMethodsInner) SetBillingAddress(v FindOrganizationPaymentMethods200ResponsePaymentMethodsInnerBillingAddress)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *FindOrganizationPaymentMethods200ResponsePaymentMethodsInner) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### GetCardType

`func (o *FindOrganizationPaymentMethods200ResponsePaymentMethodsInner) GetCardType() string`

GetCardType returns the CardType field if non-nil, zero value otherwise.

### GetCardTypeOk

`func (o *FindOrganizationPaymentMethods200ResponsePaymentMethodsInner) GetCardTypeOk() (*string, bool)`

GetCardTypeOk returns a tuple with the CardType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardType

`func (o *FindOrganizationPaymentMethods200ResponsePaymentMethodsInner) SetCardType(v string)`

SetCardType sets CardType field to given value.

### HasCardType

`func (o *FindOrganizationPaymentMethods200ResponsePaymentMethodsInner) HasCardType() bool`

HasCardType returns a boolean if a field has been set.

### GetCardholderName

`func (o *FindOrganizationPaymentMethods200ResponsePaymentMethodsInner) GetCardholderName() string`

GetCardholderName returns the CardholderName field if non-nil, zero value otherwise.

### GetCardholderNameOk

`func (o *FindOrganizationPaymentMethods200ResponsePaymentMethodsInner) GetCardholderNameOk() (*string, bool)`

GetCardholderNameOk returns a tuple with the CardholderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardholderName

`func (o *FindOrganizationPaymentMethods200ResponsePaymentMethodsInner) SetCardholderName(v string)`

SetCardholderName sets CardholderName field to given value.

### HasCardholderName

`func (o *FindOrganizationPaymentMethods200ResponsePaymentMethodsInner) HasCardholderName() bool`

HasCardholderName returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FindOrganizationPaymentMethods200ResponsePaymentMethodsInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FindOrganizationPaymentMethods200ResponsePaymentMethodsInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FindOrganizationPaymentMethods200ResponsePaymentMethodsInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FindOrganizationPaymentMethods200ResponsePaymentMethodsInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedByUser

`func (o *FindOrganizationPaymentMethods200ResponsePaymentMethodsInner) GetCreatedByUser() FindBatchById200ResponseDevicesInner`

GetCreatedByUser returns the CreatedByUser field if non-nil, zero value otherwise.

### GetCreatedByUserOk

`func (o *FindOrganizationPaymentMethods200ResponsePaymentMethodsInner) GetCreatedByUserOk() (*FindBatchById200ResponseDevicesInner, bool)`

GetCreatedByUserOk returns a tuple with the CreatedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUser

`func (o *FindOrganizationPaymentMethods200ResponsePaymentMethodsInner) SetCreatedByUser(v FindBatchById200ResponseDevicesInner)`

SetCreatedByUser sets CreatedByUser field to given value.

### HasCreatedByUser

`func (o *FindOrganizationPaymentMethods200ResponsePaymentMethodsInner) HasCreatedByUser() bool`

HasCreatedByUser returns a boolean if a field has been set.

### GetDefault

`func (o *FindOrganizationPaymentMethods200ResponsePaymentMethodsInner) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *FindOrganizationPaymentMethods200ResponsePaymentMethodsInner) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *FindOrganizationPaymentMethods200ResponsePaymentMethodsInner) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *FindOrganizationPaymentMethods200ResponsePaymentMethodsInner) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetEmail

`func (o *FindOrganizationPaymentMethods200ResponsePaymentMethodsInner) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *FindOrganizationPaymentMethods200ResponsePaymentMethodsInner) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *FindOrganizationPaymentMethods200ResponsePaymentMethodsInner) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *FindOrganizationPaymentMethods200ResponsePaymentMethodsInner) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetExpirationMonth

`func (o *FindOrganizationPaymentMethods200ResponsePaymentMethodsInner) GetExpirationMonth() string`

GetExpirationMonth returns the ExpirationMonth field if non-nil, zero value otherwise.

### GetExpirationMonthOk

`func (o *FindOrganizationPaymentMethods200ResponsePaymentMethodsInner) GetExpirationMonthOk() (*string, bool)`

GetExpirationMonthOk returns a tuple with the ExpirationMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationMonth

`func (o *FindOrganizationPaymentMethods200ResponsePaymentMethodsInner) SetExpirationMonth(v string)`

SetExpirationMonth sets ExpirationMonth field to given value.

### HasExpirationMonth

`func (o *FindOrganizationPaymentMethods200ResponsePaymentMethodsInner) HasExpirationMonth() bool`

HasExpirationMonth returns a boolean if a field has been set.

### GetExpirationYear

`func (o *FindOrganizationPaymentMethods200ResponsePaymentMethodsInner) GetExpirationYear() string`

GetExpirationYear returns the ExpirationYear field if non-nil, zero value otherwise.

### GetExpirationYearOk

`func (o *FindOrganizationPaymentMethods200ResponsePaymentMethodsInner) GetExpirationYearOk() (*string, bool)`

GetExpirationYearOk returns a tuple with the ExpirationYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationYear

`func (o *FindOrganizationPaymentMethods200ResponsePaymentMethodsInner) SetExpirationYear(v string)`

SetExpirationYear sets ExpirationYear field to given value.

### HasExpirationYear

`func (o *FindOrganizationPaymentMethods200ResponsePaymentMethodsInner) HasExpirationYear() bool`

HasExpirationYear returns a boolean if a field has been set.

### GetId

`func (o *FindOrganizationPaymentMethods200ResponsePaymentMethodsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FindOrganizationPaymentMethods200ResponsePaymentMethodsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FindOrganizationPaymentMethods200ResponsePaymentMethodsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FindOrganizationPaymentMethods200ResponsePaymentMethodsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *FindOrganizationPaymentMethods200ResponsePaymentMethodsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FindOrganizationPaymentMethods200ResponsePaymentMethodsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FindOrganizationPaymentMethods200ResponsePaymentMethodsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FindOrganizationPaymentMethods200ResponsePaymentMethodsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrganization

`func (o *FindOrganizationPaymentMethods200ResponsePaymentMethodsInner) GetOrganization() FindBatchById200ResponseDevicesInner`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *FindOrganizationPaymentMethods200ResponsePaymentMethodsInner) GetOrganizationOk() (*FindBatchById200ResponseDevicesInner, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *FindOrganizationPaymentMethods200ResponsePaymentMethodsInner) SetOrganization(v FindBatchById200ResponseDevicesInner)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *FindOrganizationPaymentMethods200ResponsePaymentMethodsInner) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetProjects

`func (o *FindOrganizationPaymentMethods200ResponsePaymentMethodsInner) GetProjects() []FindBatchById200ResponseDevicesInner`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *FindOrganizationPaymentMethods200ResponsePaymentMethodsInner) GetProjectsOk() (*[]FindBatchById200ResponseDevicesInner, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *FindOrganizationPaymentMethods200ResponsePaymentMethodsInner) SetProjects(v []FindBatchById200ResponseDevicesInner)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *FindOrganizationPaymentMethods200ResponsePaymentMethodsInner) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### GetType

`func (o *FindOrganizationPaymentMethods200ResponsePaymentMethodsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FindOrganizationPaymentMethods200ResponsePaymentMethodsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FindOrganizationPaymentMethods200ResponsePaymentMethodsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FindOrganizationPaymentMethods200ResponsePaymentMethodsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *FindOrganizationPaymentMethods200ResponsePaymentMethodsInner) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FindOrganizationPaymentMethods200ResponsePaymentMethodsInner) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FindOrganizationPaymentMethods200ResponsePaymentMethodsInner) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FindOrganizationPaymentMethods200ResponsePaymentMethodsInner) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


