# FindVrfIpReservations200ResponseIpAddressesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to **[]string** |  | [optional] 
**AddressFamily** | Pointer to **int32** |  | [optional] 
**Cidr** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedBy** | Pointer to [**FindBatchById200ResponseDevicesInner**](FindBatchById200ResponseDevicesInner.md) |  | [optional] 
**Details** | Pointer to **string** |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**MetalGateway** | Pointer to [**FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway**](FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway.md) |  | [optional] 
**Netmask** | Pointer to **string** |  | [optional] 
**Network** | Pointer to **string** |  | [optional] 
**Project** | Pointer to [**MoveHardwareReservation201ResponseProject**](MoveHardwareReservation201ResponseProject.md) |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Vrf** | Pointer to [**FindVrfs200ResponseVrfsInner**](FindVrfs200ResponseVrfsInner.md) |  | [optional] 

## Methods

### NewFindVrfIpReservations200ResponseIpAddressesInner

`func NewFindVrfIpReservations200ResponseIpAddressesInner() *FindVrfIpReservations200ResponseIpAddressesInner`

NewFindVrfIpReservations200ResponseIpAddressesInner instantiates a new FindVrfIpReservations200ResponseIpAddressesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindVrfIpReservations200ResponseIpAddressesInnerWithDefaults

`func NewFindVrfIpReservations200ResponseIpAddressesInnerWithDefaults() *FindVrfIpReservations200ResponseIpAddressesInner`

NewFindVrfIpReservations200ResponseIpAddressesInnerWithDefaults instantiates a new FindVrfIpReservations200ResponseIpAddressesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *FindVrfIpReservations200ResponseIpAddressesInner) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FindVrfIpReservations200ResponseIpAddressesInner) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FindVrfIpReservations200ResponseIpAddressesInner) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FindVrfIpReservations200ResponseIpAddressesInner) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetAddressFamily

`func (o *FindVrfIpReservations200ResponseIpAddressesInner) GetAddressFamily() int32`

GetAddressFamily returns the AddressFamily field if non-nil, zero value otherwise.

### GetAddressFamilyOk

`func (o *FindVrfIpReservations200ResponseIpAddressesInner) GetAddressFamilyOk() (*int32, bool)`

GetAddressFamilyOk returns a tuple with the AddressFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressFamily

`func (o *FindVrfIpReservations200ResponseIpAddressesInner) SetAddressFamily(v int32)`

SetAddressFamily sets AddressFamily field to given value.

### HasAddressFamily

`func (o *FindVrfIpReservations200ResponseIpAddressesInner) HasAddressFamily() bool`

HasAddressFamily returns a boolean if a field has been set.

### GetCidr

`func (o *FindVrfIpReservations200ResponseIpAddressesInner) GetCidr() int32`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *FindVrfIpReservations200ResponseIpAddressesInner) GetCidrOk() (*int32, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *FindVrfIpReservations200ResponseIpAddressesInner) SetCidr(v int32)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *FindVrfIpReservations200ResponseIpAddressesInner) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FindVrfIpReservations200ResponseIpAddressesInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FindVrfIpReservations200ResponseIpAddressesInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FindVrfIpReservations200ResponseIpAddressesInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FindVrfIpReservations200ResponseIpAddressesInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *FindVrfIpReservations200ResponseIpAddressesInner) GetCreatedBy() FindBatchById200ResponseDevicesInner`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *FindVrfIpReservations200ResponseIpAddressesInner) GetCreatedByOk() (*FindBatchById200ResponseDevicesInner, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *FindVrfIpReservations200ResponseIpAddressesInner) SetCreatedBy(v FindBatchById200ResponseDevicesInner)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *FindVrfIpReservations200ResponseIpAddressesInner) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDetails

`func (o *FindVrfIpReservations200ResponseIpAddressesInner) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *FindVrfIpReservations200ResponseIpAddressesInner) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *FindVrfIpReservations200ResponseIpAddressesInner) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *FindVrfIpReservations200ResponseIpAddressesInner) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetHref

`func (o *FindVrfIpReservations200ResponseIpAddressesInner) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *FindVrfIpReservations200ResponseIpAddressesInner) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *FindVrfIpReservations200ResponseIpAddressesInner) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *FindVrfIpReservations200ResponseIpAddressesInner) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *FindVrfIpReservations200ResponseIpAddressesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FindVrfIpReservations200ResponseIpAddressesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FindVrfIpReservations200ResponseIpAddressesInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FindVrfIpReservations200ResponseIpAddressesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetalGateway

`func (o *FindVrfIpReservations200ResponseIpAddressesInner) GetMetalGateway() FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway`

GetMetalGateway returns the MetalGateway field if non-nil, zero value otherwise.

### GetMetalGatewayOk

`func (o *FindVrfIpReservations200ResponseIpAddressesInner) GetMetalGatewayOk() (*FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway, bool)`

GetMetalGatewayOk returns a tuple with the MetalGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetalGateway

`func (o *FindVrfIpReservations200ResponseIpAddressesInner) SetMetalGateway(v FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway)`

SetMetalGateway sets MetalGateway field to given value.

### HasMetalGateway

`func (o *FindVrfIpReservations200ResponseIpAddressesInner) HasMetalGateway() bool`

HasMetalGateway returns a boolean if a field has been set.

### GetNetmask

`func (o *FindVrfIpReservations200ResponseIpAddressesInner) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *FindVrfIpReservations200ResponseIpAddressesInner) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *FindVrfIpReservations200ResponseIpAddressesInner) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *FindVrfIpReservations200ResponseIpAddressesInner) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.

### GetNetwork

`func (o *FindVrfIpReservations200ResponseIpAddressesInner) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *FindVrfIpReservations200ResponseIpAddressesInner) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *FindVrfIpReservations200ResponseIpAddressesInner) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *FindVrfIpReservations200ResponseIpAddressesInner) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetProject

`func (o *FindVrfIpReservations200ResponseIpAddressesInner) GetProject() MoveHardwareReservation201ResponseProject`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *FindVrfIpReservations200ResponseIpAddressesInner) GetProjectOk() (*MoveHardwareReservation201ResponseProject, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *FindVrfIpReservations200ResponseIpAddressesInner) SetProject(v MoveHardwareReservation201ResponseProject)`

SetProject sets Project field to given value.

### HasProject

`func (o *FindVrfIpReservations200ResponseIpAddressesInner) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetState

`func (o *FindVrfIpReservations200ResponseIpAddressesInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *FindVrfIpReservations200ResponseIpAddressesInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *FindVrfIpReservations200ResponseIpAddressesInner) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *FindVrfIpReservations200ResponseIpAddressesInner) HasState() bool`

HasState returns a boolean if a field has been set.

### GetType

`func (o *FindVrfIpReservations200ResponseIpAddressesInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FindVrfIpReservations200ResponseIpAddressesInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FindVrfIpReservations200ResponseIpAddressesInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FindVrfIpReservations200ResponseIpAddressesInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVrf

`func (o *FindVrfIpReservations200ResponseIpAddressesInner) GetVrf() FindVrfs200ResponseVrfsInner`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *FindVrfIpReservations200ResponseIpAddressesInner) GetVrfOk() (*FindVrfs200ResponseVrfsInner, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *FindVrfIpReservations200ResponseIpAddressesInner) SetVrf(v FindVrfs200ResponseVrfsInner)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *FindVrfIpReservations200ResponseIpAddressesInner) HasVrf() bool`

HasVrf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


