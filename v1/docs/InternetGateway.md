# InternetGateway

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**VirtualNetwork** | Pointer to [**Href**](Href.md) |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**IpReservations** | Pointer to [**[]Href**](Href.md) |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 

## Methods

### NewInternetGateway

`func NewInternetGateway() *InternetGateway`

NewInternetGateway instantiates a new InternetGateway object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternetGatewayWithDefaults

`func NewInternetGatewayWithDefaults() *InternetGateway`

NewInternetGatewayWithDefaults instantiates a new InternetGateway object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InternetGateway) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InternetGateway) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InternetGateway) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InternetGateway) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *InternetGateway) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InternetGateway) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InternetGateway) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *InternetGateway) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetVirtualNetwork

`func (o *InternetGateway) GetVirtualNetwork() Href`

GetVirtualNetwork returns the VirtualNetwork field if non-nil, zero value otherwise.

### GetVirtualNetworkOk

`func (o *InternetGateway) GetVirtualNetworkOk() (*Href, bool)`

GetVirtualNetworkOk returns a tuple with the VirtualNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualNetwork

`func (o *InternetGateway) SetVirtualNetwork(v Href)`

SetVirtualNetwork sets VirtualNetwork field to given value.

### HasVirtualNetwork

`func (o *InternetGateway) HasVirtualNetwork() bool`

HasVirtualNetwork returns a boolean if a field has been set.

### GetCreatedBy

`func (o *InternetGateway) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *InternetGateway) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *InternetGateway) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *InternetGateway) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetIpReservations

`func (o *InternetGateway) GetIpReservations() []Href`

GetIpReservations returns the IpReservations field if non-nil, zero value otherwise.

### GetIpReservationsOk

`func (o *InternetGateway) GetIpReservationsOk() (*[]Href, bool)`

GetIpReservationsOk returns a tuple with the IpReservations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpReservations

`func (o *InternetGateway) SetIpReservations(v []Href)`

SetIpReservations sets IpReservations field to given value.

### HasIpReservations

`func (o *InternetGateway) HasIpReservations() bool`

HasIpReservations returns a boolean if a field has been set.

### GetHref

`func (o *InternetGateway) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *InternetGateway) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *InternetGateway) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *InternetGateway) HasHref() bool`

HasHref returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


