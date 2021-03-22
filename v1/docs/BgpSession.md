# BgpSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**LearnedRoutes** | Pointer to **[]string** |  | [optional] 
**AddressFamily** | Pointer to **string** |  | [optional] 
**Device** | Pointer to [**Href**](Href.md) |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**DefaultRoute** | Pointer to **bool** |  | [optional] 

## Methods

### NewBgpSession

`func NewBgpSession() *BgpSession`

NewBgpSession instantiates a new BgpSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBgpSessionWithDefaults

`func NewBgpSessionWithDefaults() *BgpSession`

NewBgpSessionWithDefaults instantiates a new BgpSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BgpSession) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BgpSession) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BgpSession) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BgpSession) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *BgpSession) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BgpSession) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BgpSession) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BgpSession) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLearnedRoutes

`func (o *BgpSession) GetLearnedRoutes() []string`

GetLearnedRoutes returns the LearnedRoutes field if non-nil, zero value otherwise.

### GetLearnedRoutesOk

`func (o *BgpSession) GetLearnedRoutesOk() (*[]string, bool)`

GetLearnedRoutesOk returns a tuple with the LearnedRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLearnedRoutes

`func (o *BgpSession) SetLearnedRoutes(v []string)`

SetLearnedRoutes sets LearnedRoutes field to given value.

### HasLearnedRoutes

`func (o *BgpSession) HasLearnedRoutes() bool`

HasLearnedRoutes returns a boolean if a field has been set.

### GetAddressFamily

`func (o *BgpSession) GetAddressFamily() string`

GetAddressFamily returns the AddressFamily field if non-nil, zero value otherwise.

### GetAddressFamilyOk

`func (o *BgpSession) GetAddressFamilyOk() (*string, bool)`

GetAddressFamilyOk returns a tuple with the AddressFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressFamily

`func (o *BgpSession) SetAddressFamily(v string)`

SetAddressFamily sets AddressFamily field to given value.

### HasAddressFamily

`func (o *BgpSession) HasAddressFamily() bool`

HasAddressFamily returns a boolean if a field has been set.

### GetDevice

`func (o *BgpSession) GetDevice() Href`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *BgpSession) GetDeviceOk() (*Href, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *BgpSession) SetDevice(v Href)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *BgpSession) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetHref

`func (o *BgpSession) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BgpSession) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BgpSession) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *BgpSession) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetDefaultRoute

`func (o *BgpSession) GetDefaultRoute() bool`

GetDefaultRoute returns the DefaultRoute field if non-nil, zero value otherwise.

### GetDefaultRouteOk

`func (o *BgpSession) GetDefaultRouteOk() (*bool, bool)`

GetDefaultRouteOk returns a tuple with the DefaultRoute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRoute

`func (o *BgpSession) SetDefaultRoute(v bool)`

SetDefaultRoute sets DefaultRoute field to given value.

### HasDefaultRoute

`func (o *BgpSession) HasDefaultRoute() bool`

HasDefaultRoute returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


