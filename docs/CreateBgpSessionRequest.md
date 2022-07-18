# CreateBgpSessionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressFamily** | Pointer to **string** | Address family for BGP session. | [optional] 
**DefaultRoute** | Pointer to **bool** | Set the default route policy. | [optional] [default to false]

## Methods

### NewCreateBgpSessionRequest

`func NewCreateBgpSessionRequest() *CreateBgpSessionRequest`

NewCreateBgpSessionRequest instantiates a new CreateBgpSessionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBgpSessionRequestWithDefaults

`func NewCreateBgpSessionRequestWithDefaults() *CreateBgpSessionRequest`

NewCreateBgpSessionRequestWithDefaults instantiates a new CreateBgpSessionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressFamily

`func (o *CreateBgpSessionRequest) GetAddressFamily() string`

GetAddressFamily returns the AddressFamily field if non-nil, zero value otherwise.

### GetAddressFamilyOk

`func (o *CreateBgpSessionRequest) GetAddressFamilyOk() (*string, bool)`

GetAddressFamilyOk returns a tuple with the AddressFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressFamily

`func (o *CreateBgpSessionRequest) SetAddressFamily(v string)`

SetAddressFamily sets AddressFamily field to given value.

### HasAddressFamily

`func (o *CreateBgpSessionRequest) HasAddressFamily() bool`

HasAddressFamily returns a boolean if a field has been set.

### GetDefaultRoute

`func (o *CreateBgpSessionRequest) GetDefaultRoute() bool`

GetDefaultRoute returns the DefaultRoute field if non-nil, zero value otherwise.

### GetDefaultRouteOk

`func (o *CreateBgpSessionRequest) GetDefaultRouteOk() (*bool, bool)`

GetDefaultRouteOk returns a tuple with the DefaultRoute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRoute

`func (o *CreateBgpSessionRequest) SetDefaultRoute(v bool)`

SetDefaultRoute sets DefaultRoute field to given value.

### HasDefaultRoute

`func (o *CreateBgpSessionRequest) HasDefaultRoute() bool`

HasDefaultRoute returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


