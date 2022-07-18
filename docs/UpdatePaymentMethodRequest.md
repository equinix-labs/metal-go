# UpdatePaymentMethodRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingAddress** | Pointer to **map[string]interface{}** |  | [optional] 
**CardholderName** | Pointer to **string** |  | [optional] 
**Default** | Pointer to **bool** |  | [optional] 
**ExpirationMonth** | Pointer to **string** |  | [optional] 
**ExpirationYear** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdatePaymentMethodRequest

`func NewUpdatePaymentMethodRequest() *UpdatePaymentMethodRequest`

NewUpdatePaymentMethodRequest instantiates a new UpdatePaymentMethodRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePaymentMethodRequestWithDefaults

`func NewUpdatePaymentMethodRequestWithDefaults() *UpdatePaymentMethodRequest`

NewUpdatePaymentMethodRequestWithDefaults instantiates a new UpdatePaymentMethodRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingAddress

`func (o *UpdatePaymentMethodRequest) GetBillingAddress() map[string]interface{}`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *UpdatePaymentMethodRequest) GetBillingAddressOk() (*map[string]interface{}, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *UpdatePaymentMethodRequest) SetBillingAddress(v map[string]interface{})`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *UpdatePaymentMethodRequest) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### GetCardholderName

`func (o *UpdatePaymentMethodRequest) GetCardholderName() string`

GetCardholderName returns the CardholderName field if non-nil, zero value otherwise.

### GetCardholderNameOk

`func (o *UpdatePaymentMethodRequest) GetCardholderNameOk() (*string, bool)`

GetCardholderNameOk returns a tuple with the CardholderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardholderName

`func (o *UpdatePaymentMethodRequest) SetCardholderName(v string)`

SetCardholderName sets CardholderName field to given value.

### HasCardholderName

`func (o *UpdatePaymentMethodRequest) HasCardholderName() bool`

HasCardholderName returns a boolean if a field has been set.

### GetDefault

`func (o *UpdatePaymentMethodRequest) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *UpdatePaymentMethodRequest) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *UpdatePaymentMethodRequest) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *UpdatePaymentMethodRequest) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetExpirationMonth

`func (o *UpdatePaymentMethodRequest) GetExpirationMonth() string`

GetExpirationMonth returns the ExpirationMonth field if non-nil, zero value otherwise.

### GetExpirationMonthOk

`func (o *UpdatePaymentMethodRequest) GetExpirationMonthOk() (*string, bool)`

GetExpirationMonthOk returns a tuple with the ExpirationMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationMonth

`func (o *UpdatePaymentMethodRequest) SetExpirationMonth(v string)`

SetExpirationMonth sets ExpirationMonth field to given value.

### HasExpirationMonth

`func (o *UpdatePaymentMethodRequest) HasExpirationMonth() bool`

HasExpirationMonth returns a boolean if a field has been set.

### GetExpirationYear

`func (o *UpdatePaymentMethodRequest) GetExpirationYear() int32`

GetExpirationYear returns the ExpirationYear field if non-nil, zero value otherwise.

### GetExpirationYearOk

`func (o *UpdatePaymentMethodRequest) GetExpirationYearOk() (*int32, bool)`

GetExpirationYearOk returns a tuple with the ExpirationYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationYear

`func (o *UpdatePaymentMethodRequest) SetExpirationYear(v int32)`

SetExpirationYear sets ExpirationYear field to given value.

### HasExpirationYear

`func (o *UpdatePaymentMethodRequest) HasExpirationYear() bool`

HasExpirationYear returns a boolean if a field has been set.

### GetName

`func (o *UpdatePaymentMethodRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdatePaymentMethodRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdatePaymentMethodRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdatePaymentMethodRequest) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


