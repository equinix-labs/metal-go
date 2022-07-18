# CreateProjectRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Customdata** | Pointer to **map[string]interface{}** |  | [optional] 
**Name** | **string** |  | 
**OrganizationId** | Pointer to **string** |  | [optional] 
**PaymentMethodId** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateProjectRequest

`func NewCreateProjectRequest(name string, ) *CreateProjectRequest`

NewCreateProjectRequest instantiates a new CreateProjectRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProjectRequestWithDefaults

`func NewCreateProjectRequestWithDefaults() *CreateProjectRequest`

NewCreateProjectRequestWithDefaults instantiates a new CreateProjectRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomdata

`func (o *CreateProjectRequest) GetCustomdata() map[string]interface{}`

GetCustomdata returns the Customdata field if non-nil, zero value otherwise.

### GetCustomdataOk

`func (o *CreateProjectRequest) GetCustomdataOk() (*map[string]interface{}, bool)`

GetCustomdataOk returns a tuple with the Customdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomdata

`func (o *CreateProjectRequest) SetCustomdata(v map[string]interface{})`

SetCustomdata sets Customdata field to given value.

### HasCustomdata

`func (o *CreateProjectRequest) HasCustomdata() bool`

HasCustomdata returns a boolean if a field has been set.

### GetName

`func (o *CreateProjectRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateProjectRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateProjectRequest) SetName(v string)`

SetName sets Name field to given value.


### GetOrganizationId

`func (o *CreateProjectRequest) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *CreateProjectRequest) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *CreateProjectRequest) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *CreateProjectRequest) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetPaymentMethodId

`func (o *CreateProjectRequest) GetPaymentMethodId() string`

GetPaymentMethodId returns the PaymentMethodId field if non-nil, zero value otherwise.

### GetPaymentMethodIdOk

`func (o *CreateProjectRequest) GetPaymentMethodIdOk() (*string, bool)`

GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodId

`func (o *CreateProjectRequest) SetPaymentMethodId(v string)`

SetPaymentMethodId sets PaymentMethodId field to given value.

### HasPaymentMethodId

`func (o *CreateProjectRequest) HasPaymentMethodId() bool`

HasPaymentMethodId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


