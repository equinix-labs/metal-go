# FindOrganizations200ResponseOrganizationsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to [**FindDeviceById200ResponseFacilityAddress**](FindDeviceById200ResponseFacilityAddress.md) |  | [optional] 
**BillingAddress** | Pointer to [**FindDeviceById200ResponseFacilityAddress**](FindDeviceById200ResponseFacilityAddress.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**CreditAmount** | Pointer to **float32** |  | [optional] 
**Customdata** | Pointer to **map[string]interface{}** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Enforce2faAt** | Pointer to **time.Time** | Force to all members to have enabled the two factor authentication after that date, unless the value is null | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Logo** | Pointer to ***os.File** |  | [optional] 
**Members** | Pointer to [**[]FindBatchById200ResponseDevicesInner**](FindBatchById200ResponseDevicesInner.md) |  | [optional] 
**Memberships** | Pointer to [**[]FindBatchById200ResponseDevicesInner**](FindBatchById200ResponseDevicesInner.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Projects** | Pointer to [**[]FindBatchById200ResponseDevicesInner**](FindBatchById200ResponseDevicesInner.md) |  | [optional] 
**Terms** | Pointer to **int32** |  | [optional] 
**Twitter** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Website** | Pointer to **string** |  | [optional] 

## Methods

### NewFindOrganizations200ResponseOrganizationsInner

`func NewFindOrganizations200ResponseOrganizationsInner() *FindOrganizations200ResponseOrganizationsInner`

NewFindOrganizations200ResponseOrganizationsInner instantiates a new FindOrganizations200ResponseOrganizationsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindOrganizations200ResponseOrganizationsInnerWithDefaults

`func NewFindOrganizations200ResponseOrganizationsInnerWithDefaults() *FindOrganizations200ResponseOrganizationsInner`

NewFindOrganizations200ResponseOrganizationsInnerWithDefaults instantiates a new FindOrganizations200ResponseOrganizationsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *FindOrganizations200ResponseOrganizationsInner) GetAddress() FindDeviceById200ResponseFacilityAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *FindOrganizations200ResponseOrganizationsInner) GetAddressOk() (*FindDeviceById200ResponseFacilityAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *FindOrganizations200ResponseOrganizationsInner) SetAddress(v FindDeviceById200ResponseFacilityAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *FindOrganizations200ResponseOrganizationsInner) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetBillingAddress

`func (o *FindOrganizations200ResponseOrganizationsInner) GetBillingAddress() FindDeviceById200ResponseFacilityAddress`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *FindOrganizations200ResponseOrganizationsInner) GetBillingAddressOk() (*FindDeviceById200ResponseFacilityAddress, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *FindOrganizations200ResponseOrganizationsInner) SetBillingAddress(v FindDeviceById200ResponseFacilityAddress)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *FindOrganizations200ResponseOrganizationsInner) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FindOrganizations200ResponseOrganizationsInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FindOrganizations200ResponseOrganizationsInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FindOrganizations200ResponseOrganizationsInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FindOrganizations200ResponseOrganizationsInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreditAmount

`func (o *FindOrganizations200ResponseOrganizationsInner) GetCreditAmount() float32`

GetCreditAmount returns the CreditAmount field if non-nil, zero value otherwise.

### GetCreditAmountOk

`func (o *FindOrganizations200ResponseOrganizationsInner) GetCreditAmountOk() (*float32, bool)`

GetCreditAmountOk returns a tuple with the CreditAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditAmount

`func (o *FindOrganizations200ResponseOrganizationsInner) SetCreditAmount(v float32)`

SetCreditAmount sets CreditAmount field to given value.

### HasCreditAmount

`func (o *FindOrganizations200ResponseOrganizationsInner) HasCreditAmount() bool`

HasCreditAmount returns a boolean if a field has been set.

### GetCustomdata

`func (o *FindOrganizations200ResponseOrganizationsInner) GetCustomdata() map[string]interface{}`

GetCustomdata returns the Customdata field if non-nil, zero value otherwise.

### GetCustomdataOk

`func (o *FindOrganizations200ResponseOrganizationsInner) GetCustomdataOk() (*map[string]interface{}, bool)`

GetCustomdataOk returns a tuple with the Customdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomdata

`func (o *FindOrganizations200ResponseOrganizationsInner) SetCustomdata(v map[string]interface{})`

SetCustomdata sets Customdata field to given value.

### HasCustomdata

`func (o *FindOrganizations200ResponseOrganizationsInner) HasCustomdata() bool`

HasCustomdata returns a boolean if a field has been set.

### GetDescription

`func (o *FindOrganizations200ResponseOrganizationsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FindOrganizations200ResponseOrganizationsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FindOrganizations200ResponseOrganizationsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FindOrganizations200ResponseOrganizationsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnforce2faAt

`func (o *FindOrganizations200ResponseOrganizationsInner) GetEnforce2faAt() time.Time`

GetEnforce2faAt returns the Enforce2faAt field if non-nil, zero value otherwise.

### GetEnforce2faAtOk

`func (o *FindOrganizations200ResponseOrganizationsInner) GetEnforce2faAtOk() (*time.Time, bool)`

GetEnforce2faAtOk returns a tuple with the Enforce2faAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforce2faAt

`func (o *FindOrganizations200ResponseOrganizationsInner) SetEnforce2faAt(v time.Time)`

SetEnforce2faAt sets Enforce2faAt field to given value.

### HasEnforce2faAt

`func (o *FindOrganizations200ResponseOrganizationsInner) HasEnforce2faAt() bool`

HasEnforce2faAt returns a boolean if a field has been set.

### GetId

`func (o *FindOrganizations200ResponseOrganizationsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FindOrganizations200ResponseOrganizationsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FindOrganizations200ResponseOrganizationsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FindOrganizations200ResponseOrganizationsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLogo

`func (o *FindOrganizations200ResponseOrganizationsInner) GetLogo() *os.File`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *FindOrganizations200ResponseOrganizationsInner) GetLogoOk() (**os.File, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *FindOrganizations200ResponseOrganizationsInner) SetLogo(v *os.File)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *FindOrganizations200ResponseOrganizationsInner) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetMembers

`func (o *FindOrganizations200ResponseOrganizationsInner) GetMembers() []FindBatchById200ResponseDevicesInner`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *FindOrganizations200ResponseOrganizationsInner) GetMembersOk() (*[]FindBatchById200ResponseDevicesInner, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *FindOrganizations200ResponseOrganizationsInner) SetMembers(v []FindBatchById200ResponseDevicesInner)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *FindOrganizations200ResponseOrganizationsInner) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetMemberships

`func (o *FindOrganizations200ResponseOrganizationsInner) GetMemberships() []FindBatchById200ResponseDevicesInner`

GetMemberships returns the Memberships field if non-nil, zero value otherwise.

### GetMembershipsOk

`func (o *FindOrganizations200ResponseOrganizationsInner) GetMembershipsOk() (*[]FindBatchById200ResponseDevicesInner, bool)`

GetMembershipsOk returns a tuple with the Memberships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberships

`func (o *FindOrganizations200ResponseOrganizationsInner) SetMemberships(v []FindBatchById200ResponseDevicesInner)`

SetMemberships sets Memberships field to given value.

### HasMemberships

`func (o *FindOrganizations200ResponseOrganizationsInner) HasMemberships() bool`

HasMemberships returns a boolean if a field has been set.

### GetName

`func (o *FindOrganizations200ResponseOrganizationsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FindOrganizations200ResponseOrganizationsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FindOrganizations200ResponseOrganizationsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FindOrganizations200ResponseOrganizationsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProjects

`func (o *FindOrganizations200ResponseOrganizationsInner) GetProjects() []FindBatchById200ResponseDevicesInner`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *FindOrganizations200ResponseOrganizationsInner) GetProjectsOk() (*[]FindBatchById200ResponseDevicesInner, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *FindOrganizations200ResponseOrganizationsInner) SetProjects(v []FindBatchById200ResponseDevicesInner)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *FindOrganizations200ResponseOrganizationsInner) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### GetTerms

`func (o *FindOrganizations200ResponseOrganizationsInner) GetTerms() int32`

GetTerms returns the Terms field if non-nil, zero value otherwise.

### GetTermsOk

`func (o *FindOrganizations200ResponseOrganizationsInner) GetTermsOk() (*int32, bool)`

GetTermsOk returns a tuple with the Terms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerms

`func (o *FindOrganizations200ResponseOrganizationsInner) SetTerms(v int32)`

SetTerms sets Terms field to given value.

### HasTerms

`func (o *FindOrganizations200ResponseOrganizationsInner) HasTerms() bool`

HasTerms returns a boolean if a field has been set.

### GetTwitter

`func (o *FindOrganizations200ResponseOrganizationsInner) GetTwitter() string`

GetTwitter returns the Twitter field if non-nil, zero value otherwise.

### GetTwitterOk

`func (o *FindOrganizations200ResponseOrganizationsInner) GetTwitterOk() (*string, bool)`

GetTwitterOk returns a tuple with the Twitter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitter

`func (o *FindOrganizations200ResponseOrganizationsInner) SetTwitter(v string)`

SetTwitter sets Twitter field to given value.

### HasTwitter

`func (o *FindOrganizations200ResponseOrganizationsInner) HasTwitter() bool`

HasTwitter returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *FindOrganizations200ResponseOrganizationsInner) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FindOrganizations200ResponseOrganizationsInner) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FindOrganizations200ResponseOrganizationsInner) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FindOrganizations200ResponseOrganizationsInner) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetWebsite

`func (o *FindOrganizations200ResponseOrganizationsInner) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *FindOrganizations200ResponseOrganizationsInner) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *FindOrganizations200ResponseOrganizationsInner) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *FindOrganizations200ResponseOrganizationsInner) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


