# FindIPAddressById200ResponseOneOf1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressFamily** | Pointer to **int32** |  | [optional] 
**Cidr** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedBy** | Pointer to [**FindBatchById200ResponseDevicesInner**](FindBatchById200ResponseDevicesInner.md) |  | [optional] 
**Details** | Pointer to **string** |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**MetalGateway** | Pointer to [**FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetworkMetalGatewaysInner**](FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetworkMetalGatewaysInner.md) |  | [optional] 
**Netmask** | Pointer to **string** |  | [optional] 
**Network** | Pointer to **string** |  | [optional] 
**Project** | Pointer to [**GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1VrfProject**](GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1VrfProject.md) |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Type** | **string** |  | 
**Vrf** | [**GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1Vrf**](GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1Vrf.md) |  | 

## Methods

### NewFindIPAddressById200ResponseOneOf1

`func NewFindIPAddressById200ResponseOneOf1(type_ string, vrf GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1Vrf, ) *FindIPAddressById200ResponseOneOf1`

NewFindIPAddressById200ResponseOneOf1 instantiates a new FindIPAddressById200ResponseOneOf1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindIPAddressById200ResponseOneOf1WithDefaults

`func NewFindIPAddressById200ResponseOneOf1WithDefaults() *FindIPAddressById200ResponseOneOf1`

NewFindIPAddressById200ResponseOneOf1WithDefaults instantiates a new FindIPAddressById200ResponseOneOf1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressFamily

`func (o *FindIPAddressById200ResponseOneOf1) GetAddressFamily() int32`

GetAddressFamily returns the AddressFamily field if non-nil, zero value otherwise.

### GetAddressFamilyOk

`func (o *FindIPAddressById200ResponseOneOf1) GetAddressFamilyOk() (*int32, bool)`

GetAddressFamilyOk returns a tuple with the AddressFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressFamily

`func (o *FindIPAddressById200ResponseOneOf1) SetAddressFamily(v int32)`

SetAddressFamily sets AddressFamily field to given value.

### HasAddressFamily

`func (o *FindIPAddressById200ResponseOneOf1) HasAddressFamily() bool`

HasAddressFamily returns a boolean if a field has been set.

### GetCidr

`func (o *FindIPAddressById200ResponseOneOf1) GetCidr() int32`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *FindIPAddressById200ResponseOneOf1) GetCidrOk() (*int32, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *FindIPAddressById200ResponseOneOf1) SetCidr(v int32)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *FindIPAddressById200ResponseOneOf1) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FindIPAddressById200ResponseOneOf1) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FindIPAddressById200ResponseOneOf1) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FindIPAddressById200ResponseOneOf1) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FindIPAddressById200ResponseOneOf1) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *FindIPAddressById200ResponseOneOf1) GetCreatedBy() FindBatchById200ResponseDevicesInner`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *FindIPAddressById200ResponseOneOf1) GetCreatedByOk() (*FindBatchById200ResponseDevicesInner, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *FindIPAddressById200ResponseOneOf1) SetCreatedBy(v FindBatchById200ResponseDevicesInner)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *FindIPAddressById200ResponseOneOf1) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDetails

`func (o *FindIPAddressById200ResponseOneOf1) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *FindIPAddressById200ResponseOneOf1) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *FindIPAddressById200ResponseOneOf1) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *FindIPAddressById200ResponseOneOf1) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetHref

`func (o *FindIPAddressById200ResponseOneOf1) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *FindIPAddressById200ResponseOneOf1) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *FindIPAddressById200ResponseOneOf1) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *FindIPAddressById200ResponseOneOf1) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *FindIPAddressById200ResponseOneOf1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FindIPAddressById200ResponseOneOf1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FindIPAddressById200ResponseOneOf1) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FindIPAddressById200ResponseOneOf1) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetalGateway

`func (o *FindIPAddressById200ResponseOneOf1) GetMetalGateway() FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetworkMetalGatewaysInner`

GetMetalGateway returns the MetalGateway field if non-nil, zero value otherwise.

### GetMetalGatewayOk

`func (o *FindIPAddressById200ResponseOneOf1) GetMetalGatewayOk() (*FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetworkMetalGatewaysInner, bool)`

GetMetalGatewayOk returns a tuple with the MetalGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetalGateway

`func (o *FindIPAddressById200ResponseOneOf1) SetMetalGateway(v FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetworkMetalGatewaysInner)`

SetMetalGateway sets MetalGateway field to given value.

### HasMetalGateway

`func (o *FindIPAddressById200ResponseOneOf1) HasMetalGateway() bool`

HasMetalGateway returns a boolean if a field has been set.

### GetNetmask

`func (o *FindIPAddressById200ResponseOneOf1) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *FindIPAddressById200ResponseOneOf1) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *FindIPAddressById200ResponseOneOf1) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *FindIPAddressById200ResponseOneOf1) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.

### GetNetwork

`func (o *FindIPAddressById200ResponseOneOf1) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *FindIPAddressById200ResponseOneOf1) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *FindIPAddressById200ResponseOneOf1) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *FindIPAddressById200ResponseOneOf1) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetProject

`func (o *FindIPAddressById200ResponseOneOf1) GetProject() GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1VrfProject`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *FindIPAddressById200ResponseOneOf1) GetProjectOk() (*GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1VrfProject, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *FindIPAddressById200ResponseOneOf1) SetProject(v GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1VrfProject)`

SetProject sets Project field to given value.

### HasProject

`func (o *FindIPAddressById200ResponseOneOf1) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetState

`func (o *FindIPAddressById200ResponseOneOf1) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *FindIPAddressById200ResponseOneOf1) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *FindIPAddressById200ResponseOneOf1) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *FindIPAddressById200ResponseOneOf1) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTags

`func (o *FindIPAddressById200ResponseOneOf1) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FindIPAddressById200ResponseOneOf1) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FindIPAddressById200ResponseOneOf1) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FindIPAddressById200ResponseOneOf1) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *FindIPAddressById200ResponseOneOf1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FindIPAddressById200ResponseOneOf1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FindIPAddressById200ResponseOneOf1) SetType(v string)`

SetType sets Type field to given value.


### GetVrf

`func (o *FindIPAddressById200ResponseOneOf1) GetVrf() GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1Vrf`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *FindIPAddressById200ResponseOneOf1) GetVrfOk() (*GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1Vrf, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *FindIPAddressById200ResponseOneOf1) SetVrf(v GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1Vrf)`

SetVrf sets Vrf field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


