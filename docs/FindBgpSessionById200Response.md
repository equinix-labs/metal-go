# FindBgpSessionById200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressFamily** | **string** |  | 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**DefaultRoute** | Pointer to **bool** |  | [optional] 
**Device** | Pointer to [**FindBatchById200ResponseDevicesInner**](FindBatchById200ResponseDevicesInner.md) |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LearnedRoutes** | Pointer to **[]string** |  | [optional] 
**Status** | Pointer to **string** |  The status of the BGP Session. Multiple status values may be reported when the device is connected to multiple switches, one value per switch. Each status will start with \&quot;unknown\&quot; and progress to \&quot;up\&quot; or \&quot;down\&quot; depending on the connected device. Subsequent \&quot;unknown\&quot; values indicate a problem acquiring status from the switch.  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewFindBgpSessionById200Response

`func NewFindBgpSessionById200Response(addressFamily string, ) *FindBgpSessionById200Response`

NewFindBgpSessionById200Response instantiates a new FindBgpSessionById200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindBgpSessionById200ResponseWithDefaults

`func NewFindBgpSessionById200ResponseWithDefaults() *FindBgpSessionById200Response`

NewFindBgpSessionById200ResponseWithDefaults instantiates a new FindBgpSessionById200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressFamily

`func (o *FindBgpSessionById200Response) GetAddressFamily() string`

GetAddressFamily returns the AddressFamily field if non-nil, zero value otherwise.

### GetAddressFamilyOk

`func (o *FindBgpSessionById200Response) GetAddressFamilyOk() (*string, bool)`

GetAddressFamilyOk returns a tuple with the AddressFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressFamily

`func (o *FindBgpSessionById200Response) SetAddressFamily(v string)`

SetAddressFamily sets AddressFamily field to given value.


### GetCreatedAt

`func (o *FindBgpSessionById200Response) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FindBgpSessionById200Response) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FindBgpSessionById200Response) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FindBgpSessionById200Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDefaultRoute

`func (o *FindBgpSessionById200Response) GetDefaultRoute() bool`

GetDefaultRoute returns the DefaultRoute field if non-nil, zero value otherwise.

### GetDefaultRouteOk

`func (o *FindBgpSessionById200Response) GetDefaultRouteOk() (*bool, bool)`

GetDefaultRouteOk returns a tuple with the DefaultRoute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRoute

`func (o *FindBgpSessionById200Response) SetDefaultRoute(v bool)`

SetDefaultRoute sets DefaultRoute field to given value.

### HasDefaultRoute

`func (o *FindBgpSessionById200Response) HasDefaultRoute() bool`

HasDefaultRoute returns a boolean if a field has been set.

### GetDevice

`func (o *FindBgpSessionById200Response) GetDevice() FindBatchById200ResponseDevicesInner`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *FindBgpSessionById200Response) GetDeviceOk() (*FindBatchById200ResponseDevicesInner, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *FindBgpSessionById200Response) SetDevice(v FindBatchById200ResponseDevicesInner)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *FindBgpSessionById200Response) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetHref

`func (o *FindBgpSessionById200Response) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *FindBgpSessionById200Response) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *FindBgpSessionById200Response) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *FindBgpSessionById200Response) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *FindBgpSessionById200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FindBgpSessionById200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FindBgpSessionById200Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FindBgpSessionById200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLearnedRoutes

`func (o *FindBgpSessionById200Response) GetLearnedRoutes() []string`

GetLearnedRoutes returns the LearnedRoutes field if non-nil, zero value otherwise.

### GetLearnedRoutesOk

`func (o *FindBgpSessionById200Response) GetLearnedRoutesOk() (*[]string, bool)`

GetLearnedRoutesOk returns a tuple with the LearnedRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLearnedRoutes

`func (o *FindBgpSessionById200Response) SetLearnedRoutes(v []string)`

SetLearnedRoutes sets LearnedRoutes field to given value.

### HasLearnedRoutes

`func (o *FindBgpSessionById200Response) HasLearnedRoutes() bool`

HasLearnedRoutes returns a boolean if a field has been set.

### GetStatus

`func (o *FindBgpSessionById200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FindBgpSessionById200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FindBgpSessionById200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FindBgpSessionById200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *FindBgpSessionById200Response) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FindBgpSessionById200Response) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FindBgpSessionById200Response) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FindBgpSessionById200Response) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


