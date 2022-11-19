# FindIPReservations200ResponseIpAddressesInner

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
**CreatedBy** | Pointer to [**FindBatchById200ResponseDevicesInner**](FindBatchById200ResponseDevicesInner.md) |  | [optional] 
**Vrf** | [**GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1Vrf**](GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1Vrf.md) |  | 

## Methods

### NewFindIPReservations200ResponseIpAddressesInner

`func NewFindIPReservations200ResponseIpAddressesInner(type_ string, vrf GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1Vrf, ) *FindIPReservations200ResponseIpAddressesInner`

NewFindIPReservations200ResponseIpAddressesInner instantiates a new FindIPReservations200ResponseIpAddressesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindIPReservations200ResponseIpAddressesInnerWithDefaults

`func NewFindIPReservations200ResponseIpAddressesInnerWithDefaults() *FindIPReservations200ResponseIpAddressesInner`

NewFindIPReservations200ResponseIpAddressesInnerWithDefaults instantiates a new FindIPReservations200ResponseIpAddressesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddon

`func (o *FindIPReservations200ResponseIpAddressesInner) GetAddon() bool`

GetAddon returns the Addon field if non-nil, zero value otherwise.

### GetAddonOk

`func (o *FindIPReservations200ResponseIpAddressesInner) GetAddonOk() (*bool, bool)`

GetAddonOk returns a tuple with the Addon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddon

`func (o *FindIPReservations200ResponseIpAddressesInner) SetAddon(v bool)`

SetAddon sets Addon field to given value.

### HasAddon

`func (o *FindIPReservations200ResponseIpAddressesInner) HasAddon() bool`

HasAddon returns a boolean if a field has been set.

### GetAddress

`func (o *FindIPReservations200ResponseIpAddressesInner) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *FindIPReservations200ResponseIpAddressesInner) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *FindIPReservations200ResponseIpAddressesInner) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *FindIPReservations200ResponseIpAddressesInner) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAddressFamily

`func (o *FindIPReservations200ResponseIpAddressesInner) GetAddressFamily() int32`

GetAddressFamily returns the AddressFamily field if non-nil, zero value otherwise.

### GetAddressFamilyOk

`func (o *FindIPReservations200ResponseIpAddressesInner) GetAddressFamilyOk() (*int32, bool)`

GetAddressFamilyOk returns a tuple with the AddressFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressFamily

`func (o *FindIPReservations200ResponseIpAddressesInner) SetAddressFamily(v int32)`

SetAddressFamily sets AddressFamily field to given value.

### HasAddressFamily

`func (o *FindIPReservations200ResponseIpAddressesInner) HasAddressFamily() bool`

HasAddressFamily returns a boolean if a field has been set.

### GetAssignments

`func (o *FindIPReservations200ResponseIpAddressesInner) GetAssignments() []FindDeviceById200ResponseIpAddressesInner`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *FindIPReservations200ResponseIpAddressesInner) GetAssignmentsOk() (*[]FindDeviceById200ResponseIpAddressesInner, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignments

`func (o *FindIPReservations200ResponseIpAddressesInner) SetAssignments(v []FindDeviceById200ResponseIpAddressesInner)`

SetAssignments sets Assignments field to given value.

### HasAssignments

`func (o *FindIPReservations200ResponseIpAddressesInner) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### GetAvailable

`func (o *FindIPReservations200ResponseIpAddressesInner) GetAvailable() string`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *FindIPReservations200ResponseIpAddressesInner) GetAvailableOk() (*string, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *FindIPReservations200ResponseIpAddressesInner) SetAvailable(v string)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *FindIPReservations200ResponseIpAddressesInner) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetBill

`func (o *FindIPReservations200ResponseIpAddressesInner) GetBill() bool`

GetBill returns the Bill field if non-nil, zero value otherwise.

### GetBillOk

`func (o *FindIPReservations200ResponseIpAddressesInner) GetBillOk() (*bool, bool)`

GetBillOk returns a tuple with the Bill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBill

`func (o *FindIPReservations200ResponseIpAddressesInner) SetBill(v bool)`

SetBill sets Bill field to given value.

### HasBill

`func (o *FindIPReservations200ResponseIpAddressesInner) HasBill() bool`

HasBill returns a boolean if a field has been set.

### GetCidr

`func (o *FindIPReservations200ResponseIpAddressesInner) GetCidr() int32`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *FindIPReservations200ResponseIpAddressesInner) GetCidrOk() (*int32, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *FindIPReservations200ResponseIpAddressesInner) SetCidr(v int32)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *FindIPReservations200ResponseIpAddressesInner) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FindIPReservations200ResponseIpAddressesInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FindIPReservations200ResponseIpAddressesInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FindIPReservations200ResponseIpAddressesInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FindIPReservations200ResponseIpAddressesInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCustomdata

`func (o *FindIPReservations200ResponseIpAddressesInner) GetCustomdata() map[string]interface{}`

GetCustomdata returns the Customdata field if non-nil, zero value otherwise.

### GetCustomdataOk

`func (o *FindIPReservations200ResponseIpAddressesInner) GetCustomdataOk() (*map[string]interface{}, bool)`

GetCustomdataOk returns a tuple with the Customdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomdata

`func (o *FindIPReservations200ResponseIpAddressesInner) SetCustomdata(v map[string]interface{})`

SetCustomdata sets Customdata field to given value.

### HasCustomdata

`func (o *FindIPReservations200ResponseIpAddressesInner) HasCustomdata() bool`

HasCustomdata returns a boolean if a field has been set.

### GetEnabled

`func (o *FindIPReservations200ResponseIpAddressesInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *FindIPReservations200ResponseIpAddressesInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *FindIPReservations200ResponseIpAddressesInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *FindIPReservations200ResponseIpAddressesInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDetails

`func (o *FindIPReservations200ResponseIpAddressesInner) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *FindIPReservations200ResponseIpAddressesInner) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *FindIPReservations200ResponseIpAddressesInner) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *FindIPReservations200ResponseIpAddressesInner) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetFacility

`func (o *FindIPReservations200ResponseIpAddressesInner) GetFacility() FindIPAddressById200ResponseOneOfFacility`

GetFacility returns the Facility field if non-nil, zero value otherwise.

### GetFacilityOk

`func (o *FindIPReservations200ResponseIpAddressesInner) GetFacilityOk() (*FindIPAddressById200ResponseOneOfFacility, bool)`

GetFacilityOk returns a tuple with the Facility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacility

`func (o *FindIPReservations200ResponseIpAddressesInner) SetFacility(v FindIPAddressById200ResponseOneOfFacility)`

SetFacility sets Facility field to given value.

### HasFacility

`func (o *FindIPReservations200ResponseIpAddressesInner) HasFacility() bool`

HasFacility returns a boolean if a field has been set.

### GetGateway

`func (o *FindIPReservations200ResponseIpAddressesInner) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *FindIPReservations200ResponseIpAddressesInner) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *FindIPReservations200ResponseIpAddressesInner) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *FindIPReservations200ResponseIpAddressesInner) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetGlobalIp

`func (o *FindIPReservations200ResponseIpAddressesInner) GetGlobalIp() bool`

GetGlobalIp returns the GlobalIp field if non-nil, zero value otherwise.

### GetGlobalIpOk

`func (o *FindIPReservations200ResponseIpAddressesInner) GetGlobalIpOk() (*bool, bool)`

GetGlobalIpOk returns a tuple with the GlobalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalIp

`func (o *FindIPReservations200ResponseIpAddressesInner) SetGlobalIp(v bool)`

SetGlobalIp sets GlobalIp field to given value.

### HasGlobalIp

`func (o *FindIPReservations200ResponseIpAddressesInner) HasGlobalIp() bool`

HasGlobalIp returns a boolean if a field has been set.

### GetHref

`func (o *FindIPReservations200ResponseIpAddressesInner) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *FindIPReservations200ResponseIpAddressesInner) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *FindIPReservations200ResponseIpAddressesInner) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *FindIPReservations200ResponseIpAddressesInner) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *FindIPReservations200ResponseIpAddressesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FindIPReservations200ResponseIpAddressesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FindIPReservations200ResponseIpAddressesInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FindIPReservations200ResponseIpAddressesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetManageable

`func (o *FindIPReservations200ResponseIpAddressesInner) GetManageable() bool`

GetManageable returns the Manageable field if non-nil, zero value otherwise.

### GetManageableOk

`func (o *FindIPReservations200ResponseIpAddressesInner) GetManageableOk() (*bool, bool)`

GetManageableOk returns a tuple with the Manageable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManageable

`func (o *FindIPReservations200ResponseIpAddressesInner) SetManageable(v bool)`

SetManageable sets Manageable field to given value.

### HasManageable

`func (o *FindIPReservations200ResponseIpAddressesInner) HasManageable() bool`

HasManageable returns a boolean if a field has been set.

### GetManagement

`func (o *FindIPReservations200ResponseIpAddressesInner) GetManagement() bool`

GetManagement returns the Management field if non-nil, zero value otherwise.

### GetManagementOk

`func (o *FindIPReservations200ResponseIpAddressesInner) GetManagementOk() (*bool, bool)`

GetManagementOk returns a tuple with the Management field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagement

`func (o *FindIPReservations200ResponseIpAddressesInner) SetManagement(v bool)`

SetManagement sets Management field to given value.

### HasManagement

`func (o *FindIPReservations200ResponseIpAddressesInner) HasManagement() bool`

HasManagement returns a boolean if a field has been set.

### GetMetalGateway

`func (o *FindIPReservations200ResponseIpAddressesInner) GetMetalGateway() FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetworkMetalGatewaysInner`

GetMetalGateway returns the MetalGateway field if non-nil, zero value otherwise.

### GetMetalGatewayOk

`func (o *FindIPReservations200ResponseIpAddressesInner) GetMetalGatewayOk() (*FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetworkMetalGatewaysInner, bool)`

GetMetalGatewayOk returns a tuple with the MetalGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetalGateway

`func (o *FindIPReservations200ResponseIpAddressesInner) SetMetalGateway(v FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetworkMetalGatewaysInner)`

SetMetalGateway sets MetalGateway field to given value.

### HasMetalGateway

`func (o *FindIPReservations200ResponseIpAddressesInner) HasMetalGateway() bool`

HasMetalGateway returns a boolean if a field has been set.

### GetMetro

`func (o *FindIPReservations200ResponseIpAddressesInner) GetMetro() FindIPAddressById200ResponseOneOfMetro`

GetMetro returns the Metro field if non-nil, zero value otherwise.

### GetMetroOk

`func (o *FindIPReservations200ResponseIpAddressesInner) GetMetroOk() (*FindIPAddressById200ResponseOneOfMetro, bool)`

GetMetroOk returns a tuple with the Metro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetro

`func (o *FindIPReservations200ResponseIpAddressesInner) SetMetro(v FindIPAddressById200ResponseOneOfMetro)`

SetMetro sets Metro field to given value.

### HasMetro

`func (o *FindIPReservations200ResponseIpAddressesInner) HasMetro() bool`

HasMetro returns a boolean if a field has been set.

### GetNetmask

`func (o *FindIPReservations200ResponseIpAddressesInner) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *FindIPReservations200ResponseIpAddressesInner) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *FindIPReservations200ResponseIpAddressesInner) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *FindIPReservations200ResponseIpAddressesInner) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.

### GetNetwork

`func (o *FindIPReservations200ResponseIpAddressesInner) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *FindIPReservations200ResponseIpAddressesInner) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *FindIPReservations200ResponseIpAddressesInner) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *FindIPReservations200ResponseIpAddressesInner) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetProject

`func (o *FindIPReservations200ResponseIpAddressesInner) GetProject() GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1VrfProject`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *FindIPReservations200ResponseIpAddressesInner) GetProjectOk() (*GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1VrfProject, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *FindIPReservations200ResponseIpAddressesInner) SetProject(v GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1VrfProject)`

SetProject sets Project field to given value.

### HasProject

`func (o *FindIPReservations200ResponseIpAddressesInner) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetProjectLite

`func (o *FindIPReservations200ResponseIpAddressesInner) GetProjectLite() FindBatchById200ResponseDevicesInner`

GetProjectLite returns the ProjectLite field if non-nil, zero value otherwise.

### GetProjectLiteOk

`func (o *FindIPReservations200ResponseIpAddressesInner) GetProjectLiteOk() (*FindBatchById200ResponseDevicesInner, bool)`

GetProjectLiteOk returns a tuple with the ProjectLite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectLite

`func (o *FindIPReservations200ResponseIpAddressesInner) SetProjectLite(v FindBatchById200ResponseDevicesInner)`

SetProjectLite sets ProjectLite field to given value.

### HasProjectLite

`func (o *FindIPReservations200ResponseIpAddressesInner) HasProjectLite() bool`

HasProjectLite returns a boolean if a field has been set.

### GetRequestedBy

`func (o *FindIPReservations200ResponseIpAddressesInner) GetRequestedBy() FindBatchById200ResponseDevicesInner`

GetRequestedBy returns the RequestedBy field if non-nil, zero value otherwise.

### GetRequestedByOk

`func (o *FindIPReservations200ResponseIpAddressesInner) GetRequestedByOk() (*FindBatchById200ResponseDevicesInner, bool)`

GetRequestedByOk returns a tuple with the RequestedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedBy

`func (o *FindIPReservations200ResponseIpAddressesInner) SetRequestedBy(v FindBatchById200ResponseDevicesInner)`

SetRequestedBy sets RequestedBy field to given value.

### HasRequestedBy

`func (o *FindIPReservations200ResponseIpAddressesInner) HasRequestedBy() bool`

HasRequestedBy returns a boolean if a field has been set.

### GetPublic

`func (o *FindIPReservations200ResponseIpAddressesInner) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *FindIPReservations200ResponseIpAddressesInner) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *FindIPReservations200ResponseIpAddressesInner) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *FindIPReservations200ResponseIpAddressesInner) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetState

`func (o *FindIPReservations200ResponseIpAddressesInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *FindIPReservations200ResponseIpAddressesInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *FindIPReservations200ResponseIpAddressesInner) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *FindIPReservations200ResponseIpAddressesInner) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTags

`func (o *FindIPReservations200ResponseIpAddressesInner) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FindIPReservations200ResponseIpAddressesInner) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FindIPReservations200ResponseIpAddressesInner) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FindIPReservations200ResponseIpAddressesInner) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *FindIPReservations200ResponseIpAddressesInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FindIPReservations200ResponseIpAddressesInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FindIPReservations200ResponseIpAddressesInner) SetType(v string)`

SetType sets Type field to given value.


### GetCreatedBy

`func (o *FindIPReservations200ResponseIpAddressesInner) GetCreatedBy() FindBatchById200ResponseDevicesInner`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *FindIPReservations200ResponseIpAddressesInner) GetCreatedByOk() (*FindBatchById200ResponseDevicesInner, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *FindIPReservations200ResponseIpAddressesInner) SetCreatedBy(v FindBatchById200ResponseDevicesInner)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *FindIPReservations200ResponseIpAddressesInner) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetVrf

`func (o *FindIPReservations200ResponseIpAddressesInner) GetVrf() GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1Vrf`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *FindIPReservations200ResponseIpAddressesInner) GetVrfOk() (*GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1Vrf, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *FindIPReservations200ResponseIpAddressesInner) SetVrf(v GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1Vrf)`

SetVrf sets Vrf field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


