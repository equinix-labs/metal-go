# CreateOrganizationProjectRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Customdata** | Pointer to **map[string]interface{}** |  | [optional] 
**Name** | **string** |  | 
**PaymentMethodId** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateOrganizationProjectRequest

`func NewCreateOrganizationProjectRequest(name string, ) *CreateOrganizationProjectRequest`

NewCreateOrganizationProjectRequest instantiates a new CreateOrganizationProjectRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrganizationProjectRequestWithDefaults

`func NewCreateOrganizationProjectRequestWithDefaults() *CreateOrganizationProjectRequest`

NewCreateOrganizationProjectRequestWithDefaults instantiates a new CreateOrganizationProjectRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomdata

`func (o *CreateOrganizationProjectRequest) GetCustomdata() map[string]interface{}`

GetCustomdata returns the Customdata field if non-nil, zero value otherwise.

### GetCustomdataOk

`func (o *CreateOrganizationProjectRequest) GetCustomdataOk() (*map[string]interface{}, bool)`

GetCustomdataOk returns a tuple with the Customdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomdata

`func (o *CreateOrganizationProjectRequest) SetCustomdata(v map[string]interface{})`

SetCustomdata sets Customdata field to given value.

### HasCustomdata

`func (o *CreateOrganizationProjectRequest) HasCustomdata() bool`

HasCustomdata returns a boolean if a field has been set.

### GetName

`func (o *CreateOrganizationProjectRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateOrganizationProjectRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateOrganizationProjectRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPaymentMethodId

`func (o *CreateOrganizationProjectRequest) GetPaymentMethodId() string`

GetPaymentMethodId returns the PaymentMethodId field if non-nil, zero value otherwise.

### GetPaymentMethodIdOk

`func (o *CreateOrganizationProjectRequest) GetPaymentMethodIdOk() (*string, bool)`

GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodId

`func (o *CreateOrganizationProjectRequest) SetPaymentMethodId(v string)`

SetPaymentMethodId sets PaymentMethodId field to given value.

### HasPaymentMethodId

`func (o *CreateOrganizationProjectRequest) HasPaymentMethodId() bool`

HasPaymentMethodId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


