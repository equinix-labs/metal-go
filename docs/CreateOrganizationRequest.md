# CreateOrganizationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to [**FindDeviceById200ResponseFacilityAddress**](FindDeviceById200ResponseFacilityAddress.md) |  | [optional] 
**BillingAddress** | Pointer to [**FindDeviceById200ResponseFacilityAddress**](FindDeviceById200ResponseFacilityAddress.md) |  | [optional] 
**Customdata** | Pointer to **map[string]interface{}** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Enforce2faAt** | Pointer to **time.Time** | Force to all members to have enabled the two factor authentication after that date, unless the value is null | [optional] 
**Logo** | Pointer to ***os.File** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Twitter** | Pointer to **string** |  | [optional] 
**Website** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateOrganizationRequest

`func NewCreateOrganizationRequest() *CreateOrganizationRequest`

NewCreateOrganizationRequest instantiates a new CreateOrganizationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrganizationRequestWithDefaults

`func NewCreateOrganizationRequestWithDefaults() *CreateOrganizationRequest`

NewCreateOrganizationRequestWithDefaults instantiates a new CreateOrganizationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *CreateOrganizationRequest) GetAddress() FindDeviceById200ResponseFacilityAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CreateOrganizationRequest) GetAddressOk() (*FindDeviceById200ResponseFacilityAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CreateOrganizationRequest) SetAddress(v FindDeviceById200ResponseFacilityAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *CreateOrganizationRequest) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetBillingAddress

`func (o *CreateOrganizationRequest) GetBillingAddress() FindDeviceById200ResponseFacilityAddress`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *CreateOrganizationRequest) GetBillingAddressOk() (*FindDeviceById200ResponseFacilityAddress, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *CreateOrganizationRequest) SetBillingAddress(v FindDeviceById200ResponseFacilityAddress)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *CreateOrganizationRequest) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### GetCustomdata

`func (o *CreateOrganizationRequest) GetCustomdata() map[string]interface{}`

GetCustomdata returns the Customdata field if non-nil, zero value otherwise.

### GetCustomdataOk

`func (o *CreateOrganizationRequest) GetCustomdataOk() (*map[string]interface{}, bool)`

GetCustomdataOk returns a tuple with the Customdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomdata

`func (o *CreateOrganizationRequest) SetCustomdata(v map[string]interface{})`

SetCustomdata sets Customdata field to given value.

### HasCustomdata

`func (o *CreateOrganizationRequest) HasCustomdata() bool`

HasCustomdata returns a boolean if a field has been set.

### GetDescription

`func (o *CreateOrganizationRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateOrganizationRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateOrganizationRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateOrganizationRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnforce2faAt

`func (o *CreateOrganizationRequest) GetEnforce2faAt() time.Time`

GetEnforce2faAt returns the Enforce2faAt field if non-nil, zero value otherwise.

### GetEnforce2faAtOk

`func (o *CreateOrganizationRequest) GetEnforce2faAtOk() (*time.Time, bool)`

GetEnforce2faAtOk returns a tuple with the Enforce2faAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforce2faAt

`func (o *CreateOrganizationRequest) SetEnforce2faAt(v time.Time)`

SetEnforce2faAt sets Enforce2faAt field to given value.

### HasEnforce2faAt

`func (o *CreateOrganizationRequest) HasEnforce2faAt() bool`

HasEnforce2faAt returns a boolean if a field has been set.

### GetLogo

`func (o *CreateOrganizationRequest) GetLogo() *os.File`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *CreateOrganizationRequest) GetLogoOk() (**os.File, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *CreateOrganizationRequest) SetLogo(v *os.File)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *CreateOrganizationRequest) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetName

`func (o *CreateOrganizationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateOrganizationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateOrganizationRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateOrganizationRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTwitter

`func (o *CreateOrganizationRequest) GetTwitter() string`

GetTwitter returns the Twitter field if non-nil, zero value otherwise.

### GetTwitterOk

`func (o *CreateOrganizationRequest) GetTwitterOk() (*string, bool)`

GetTwitterOk returns a tuple with the Twitter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitter

`func (o *CreateOrganizationRequest) SetTwitter(v string)`

SetTwitter sets Twitter field to given value.

### HasTwitter

`func (o *CreateOrganizationRequest) HasTwitter() bool`

HasTwitter returns a boolean if a field has been set.

### GetWebsite

`func (o *CreateOrganizationRequest) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *CreateOrganizationRequest) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *CreateOrganizationRequest) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *CreateOrganizationRequest) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


