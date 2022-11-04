# FindIPAddressById200ResponseOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addon** | Pointer to **bool** |  | [optional] 
**Address** | Pointer to **string** |  | [optional] 
**AddressFamily** | Pointer to **int32** |  | [optional] 
**Assignments** | Pointer to [**[]FindDeviceById200ResponseIpAddressesInner**](FindDeviceById200ResponseIpAddressesInner.md) |  | [optional] 
**Available** | Pointer to **string** |  | [optional] 
**Bill** | Pointer to **bool** |  | [optional] 
**Cidr** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Customdata** | Pointer to **map[string]interface{}** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Details** | Pointer to **string** |  | [optional] 
**Facility** | Pointer to [**FindIPAddressById200ResponseOneOfFacility**](FindIPAddressById200ResponseOneOfFacility.md) |  | [optional] 
**Gateway** | Pointer to **string** |  | [optional] 
**GlobalIp** | Pointer to **bool** |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Manageable** | Pointer to **bool** |  | [optional] 
**Management** | Pointer to **bool** |  | [optional] 
**MetalGateway** | Pointer to [**FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetworkMetalGatewaysInner**](FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetworkMetalGatewaysInner.md) |  | [optional] 
**Metro** | Pointer to [**FindIPAddressById200ResponseOneOfMetro**](FindIPAddressById200ResponseOneOfMetro.md) |  | [optional] 
**Netmask** | Pointer to **string** |  | [optional] 
**Network** | Pointer to **string** |  | [optional] 
**Project** | Pointer to [**GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1VrfProject**](GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1VrfProject.md) |  | [optional] 
**ProjectLite** | Pointer to [**FindBatchById200ResponseDevicesInner**](FindBatchById200ResponseDevicesInner.md) |  | [optional] 
**RequestedBy** | Pointer to [**FindBatchById200ResponseDevicesInner**](FindBatchById200ResponseDevicesInner.md) |  | [optional] 
**Public** | Pointer to **bool** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewFindIPAddressById200ResponseOneOf

`func NewFindIPAddressById200ResponseOneOf(type_ string, ) *FindIPAddressById200ResponseOneOf`

NewFindIPAddressById200ResponseOneOf instantiates a new FindIPAddressById200ResponseOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindIPAddressById200ResponseOneOfWithDefaults

`func NewFindIPAddressById200ResponseOneOfWithDefaults() *FindIPAddressById200ResponseOneOf`

NewFindIPAddressById200ResponseOneOfWithDefaults instantiates a new FindIPAddressById200ResponseOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddon

`func (o *FindIPAddressById200ResponseOneOf) GetAddon() bool`

GetAddon returns the Addon field if non-nil, zero value otherwise.

### GetAddonOk

`func (o *FindIPAddressById200ResponseOneOf) GetAddonOk() (*bool, bool)`

GetAddonOk returns a tuple with the Addon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddon

`func (o *FindIPAddressById200ResponseOneOf) SetAddon(v bool)`

SetAddon sets Addon field to given value.

### HasAddon

`func (o *FindIPAddressById200ResponseOneOf) HasAddon() bool`

HasAddon returns a boolean if a field has been set.

### GetAddress

`func (o *FindIPAddressById200ResponseOneOf) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *FindIPAddressById200ResponseOneOf) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *FindIPAddressById200ResponseOneOf) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *FindIPAddressById200ResponseOneOf) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAddressFamily

`func (o *FindIPAddressById200ResponseOneOf) GetAddressFamily() int32`

GetAddressFamily returns the AddressFamily field if non-nil, zero value otherwise.

### GetAddressFamilyOk

`func (o *FindIPAddressById200ResponseOneOf) GetAddressFamilyOk() (*int32, bool)`

GetAddressFamilyOk returns a tuple with the AddressFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressFamily

`func (o *FindIPAddressById200ResponseOneOf) SetAddressFamily(v int32)`

SetAddressFamily sets AddressFamily field to given value.

### HasAddressFamily

`func (o *FindIPAddressById200ResponseOneOf) HasAddressFamily() bool`

HasAddressFamily returns a boolean if a field has been set.

### GetAssignments

`func (o *FindIPAddressById200ResponseOneOf) GetAssignments() []FindDeviceById200ResponseIpAddressesInner`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *FindIPAddressById200ResponseOneOf) GetAssignmentsOk() (*[]FindDeviceById200ResponseIpAddressesInner, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignments

`func (o *FindIPAddressById200ResponseOneOf) SetAssignments(v []FindDeviceById200ResponseIpAddressesInner)`

SetAssignments sets Assignments field to given value.

### HasAssignments

`func (o *FindIPAddressById200ResponseOneOf) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### GetAvailable

`func (o *FindIPAddressById200ResponseOneOf) GetAvailable() string`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *FindIPAddressById200ResponseOneOf) GetAvailableOk() (*string, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *FindIPAddressById200ResponseOneOf) SetAvailable(v string)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *FindIPAddressById200ResponseOneOf) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetBill

`func (o *FindIPAddressById200ResponseOneOf) GetBill() bool`

GetBill returns the Bill field if non-nil, zero value otherwise.

### GetBillOk

`func (o *FindIPAddressById200ResponseOneOf) GetBillOk() (*bool, bool)`

GetBillOk returns a tuple with the Bill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBill

`func (o *FindIPAddressById200ResponseOneOf) SetBill(v bool)`

SetBill sets Bill field to given value.

### HasBill

`func (o *FindIPAddressById200ResponseOneOf) HasBill() bool`

HasBill returns a boolean if a field has been set.

### GetCidr

`func (o *FindIPAddressById200ResponseOneOf) GetCidr() int32`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *FindIPAddressById200ResponseOneOf) GetCidrOk() (*int32, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *FindIPAddressById200ResponseOneOf) SetCidr(v int32)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *FindIPAddressById200ResponseOneOf) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FindIPAddressById200ResponseOneOf) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FindIPAddressById200ResponseOneOf) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FindIPAddressById200ResponseOneOf) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FindIPAddressById200ResponseOneOf) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCustomdata

`func (o *FindIPAddressById200ResponseOneOf) GetCustomdata() map[string]interface{}`

GetCustomdata returns the Customdata field if non-nil, zero value otherwise.

### GetCustomdataOk

`func (o *FindIPAddressById200ResponseOneOf) GetCustomdataOk() (*map[string]interface{}, bool)`

GetCustomdataOk returns a tuple with the Customdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomdata

`func (o *FindIPAddressById200ResponseOneOf) SetCustomdata(v map[string]interface{})`

SetCustomdata sets Customdata field to given value.

### HasCustomdata

`func (o *FindIPAddressById200ResponseOneOf) HasCustomdata() bool`

HasCustomdata returns a boolean if a field has been set.

### GetEnabled

`func (o *FindIPAddressById200ResponseOneOf) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *FindIPAddressById200ResponseOneOf) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *FindIPAddressById200ResponseOneOf) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *FindIPAddressById200ResponseOneOf) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDetails

`func (o *FindIPAddressById200ResponseOneOf) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *FindIPAddressById200ResponseOneOf) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *FindIPAddressById200ResponseOneOf) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *FindIPAddressById200ResponseOneOf) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetFacility

`func (o *FindIPAddressById200ResponseOneOf) GetFacility() FindIPAddressById200ResponseOneOfFacility`

GetFacility returns the Facility field if non-nil, zero value otherwise.

### GetFacilityOk

`func (o *FindIPAddressById200ResponseOneOf) GetFacilityOk() (*FindIPAddressById200ResponseOneOfFacility, bool)`

GetFacilityOk returns a tuple with the Facility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacility

`func (o *FindIPAddressById200ResponseOneOf) SetFacility(v FindIPAddressById200ResponseOneOfFacility)`

SetFacility sets Facility field to given value.

### HasFacility

`func (o *FindIPAddressById200ResponseOneOf) HasFacility() bool`

HasFacility returns a boolean if a field has been set.

### GetGateway

`func (o *FindIPAddressById200ResponseOneOf) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *FindIPAddressById200ResponseOneOf) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *FindIPAddressById200ResponseOneOf) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *FindIPAddressById200ResponseOneOf) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetGlobalIp

`func (o *FindIPAddressById200ResponseOneOf) GetGlobalIp() bool`

GetGlobalIp returns the GlobalIp field if non-nil, zero value otherwise.

### GetGlobalIpOk

`func (o *FindIPAddressById200ResponseOneOf) GetGlobalIpOk() (*bool, bool)`

GetGlobalIpOk returns a tuple with the GlobalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalIp

`func (o *FindIPAddressById200ResponseOneOf) SetGlobalIp(v bool)`

SetGlobalIp sets GlobalIp field to given value.

### HasGlobalIp

`func (o *FindIPAddressById200ResponseOneOf) HasGlobalIp() bool`

HasGlobalIp returns a boolean if a field has been set.

### GetHref

`func (o *FindIPAddressById200ResponseOneOf) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *FindIPAddressById200ResponseOneOf) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *FindIPAddressById200ResponseOneOf) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *FindIPAddressById200ResponseOneOf) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *FindIPAddressById200ResponseOneOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FindIPAddressById200ResponseOneOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FindIPAddressById200ResponseOneOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FindIPAddressById200ResponseOneOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetManageable

`func (o *FindIPAddressById200ResponseOneOf) GetManageable() bool`

GetManageable returns the Manageable field if non-nil, zero value otherwise.

### GetManageableOk

`func (o *FindIPAddressById200ResponseOneOf) GetManageableOk() (*bool, bool)`

GetManageableOk returns a tuple with the Manageable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManageable

`func (o *FindIPAddressById200ResponseOneOf) SetManageable(v bool)`

SetManageable sets Manageable field to given value.

### HasManageable

`func (o *FindIPAddressById200ResponseOneOf) HasManageable() bool`

HasManageable returns a boolean if a field has been set.

### GetManagement

`func (o *FindIPAddressById200ResponseOneOf) GetManagement() bool`

GetManagement returns the Management field if non-nil, zero value otherwise.

### GetManagementOk

`func (o *FindIPAddressById200ResponseOneOf) GetManagementOk() (*bool, bool)`

GetManagementOk returns a tuple with the Management field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagement

`func (o *FindIPAddressById200ResponseOneOf) SetManagement(v bool)`

SetManagement sets Management field to given value.

### HasManagement

`func (o *FindIPAddressById200ResponseOneOf) HasManagement() bool`

HasManagement returns a boolean if a field has been set.

### GetMetalGateway

`func (o *FindIPAddressById200ResponseOneOf) GetMetalGateway() FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetworkMetalGatewaysInner`

GetMetalGateway returns the MetalGateway field if non-nil, zero value otherwise.

### GetMetalGatewayOk

`func (o *FindIPAddressById200ResponseOneOf) GetMetalGatewayOk() (*FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetworkMetalGatewaysInner, bool)`

GetMetalGatewayOk returns a tuple with the MetalGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetalGateway

`func (o *FindIPAddressById200ResponseOneOf) SetMetalGateway(v FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetworkMetalGatewaysInner)`

SetMetalGateway sets MetalGateway field to given value.

### HasMetalGateway

`func (o *FindIPAddressById200ResponseOneOf) HasMetalGateway() bool`

HasMetalGateway returns a boolean if a field has been set.

### GetMetro

`func (o *FindIPAddressById200ResponseOneOf) GetMetro() FindIPAddressById200ResponseOneOfMetro`

GetMetro returns the Metro field if non-nil, zero value otherwise.

### GetMetroOk

`func (o *FindIPAddressById200ResponseOneOf) GetMetroOk() (*FindIPAddressById200ResponseOneOfMetro, bool)`

GetMetroOk returns a tuple with the Metro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetro

`func (o *FindIPAddressById200ResponseOneOf) SetMetro(v FindIPAddressById200ResponseOneOfMetro)`

SetMetro sets Metro field to given value.

### HasMetro

`func (o *FindIPAddressById200ResponseOneOf) HasMetro() bool`

HasMetro returns a boolean if a field has been set.

### GetNetmask

`func (o *FindIPAddressById200ResponseOneOf) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *FindIPAddressById200ResponseOneOf) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *FindIPAddressById200ResponseOneOf) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *FindIPAddressById200ResponseOneOf) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.

### GetNetwork

`func (o *FindIPAddressById200ResponseOneOf) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *FindIPAddressById200ResponseOneOf) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *FindIPAddressById200ResponseOneOf) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *FindIPAddressById200ResponseOneOf) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetProject

`func (o *FindIPAddressById200ResponseOneOf) GetProject() GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1VrfProject`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *FindIPAddressById200ResponseOneOf) GetProjectOk() (*GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1VrfProject, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *FindIPAddressById200ResponseOneOf) SetProject(v GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1VrfProject)`

SetProject sets Project field to given value.

### HasProject

`func (o *FindIPAddressById200ResponseOneOf) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetProjectLite

`func (o *FindIPAddressById200ResponseOneOf) GetProjectLite() FindBatchById200ResponseDevicesInner`

GetProjectLite returns the ProjectLite field if non-nil, zero value otherwise.

### GetProjectLiteOk

`func (o *FindIPAddressById200ResponseOneOf) GetProjectLiteOk() (*FindBatchById200ResponseDevicesInner, bool)`

GetProjectLiteOk returns a tuple with the ProjectLite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectLite

`func (o *FindIPAddressById200ResponseOneOf) SetProjectLite(v FindBatchById200ResponseDevicesInner)`

SetProjectLite sets ProjectLite field to given value.

### HasProjectLite

`func (o *FindIPAddressById200ResponseOneOf) HasProjectLite() bool`

HasProjectLite returns a boolean if a field has been set.

### GetRequestedBy

`func (o *FindIPAddressById200ResponseOneOf) GetRequestedBy() FindBatchById200ResponseDevicesInner`

GetRequestedBy returns the RequestedBy field if non-nil, zero value otherwise.

### GetRequestedByOk

`func (o *FindIPAddressById200ResponseOneOf) GetRequestedByOk() (*FindBatchById200ResponseDevicesInner, bool)`

GetRequestedByOk returns a tuple with the RequestedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedBy

`func (o *FindIPAddressById200ResponseOneOf) SetRequestedBy(v FindBatchById200ResponseDevicesInner)`

SetRequestedBy sets RequestedBy field to given value.

### HasRequestedBy

`func (o *FindIPAddressById200ResponseOneOf) HasRequestedBy() bool`

HasRequestedBy returns a boolean if a field has been set.

### GetPublic

`func (o *FindIPAddressById200ResponseOneOf) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *FindIPAddressById200ResponseOneOf) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *FindIPAddressById200ResponseOneOf) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *FindIPAddressById200ResponseOneOf) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetState

`func (o *FindIPAddressById200ResponseOneOf) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *FindIPAddressById200ResponseOneOf) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *FindIPAddressById200ResponseOneOf) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *FindIPAddressById200ResponseOneOf) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTags

`func (o *FindIPAddressById200ResponseOneOf) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FindIPAddressById200ResponseOneOf) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FindIPAddressById200ResponseOneOf) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FindIPAddressById200ResponseOneOf) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *FindIPAddressById200ResponseOneOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FindIPAddressById200ResponseOneOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FindIPAddressById200ResponseOneOf) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


