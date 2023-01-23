# UpdateIPAddress200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** |  | [optional] 
**AddressFamily** | Pointer to **int32** |  | [optional] 
**AssignedTo** | Pointer to [**Href**](Href.md) |  | [optional] 
**Cidr** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Gateway** | Pointer to **string** |  | [optional] 
**GlobalIp** | Pointer to **bool** |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Manageable** | Pointer to **bool** |  | [optional] 
**Management** | Pointer to **bool** |  | [optional] 
**Metro** | Pointer to [**Metro**](Metro.md) |  | [optional] 
**Netmask** | Pointer to **string** |  | [optional] 
**Network** | Pointer to **string** |  | [optional] 
**ParentBlock** | Pointer to [**ParentBlock**](ParentBlock.md) |  | [optional] 
**Public** | Pointer to **bool** |  | [optional] 
**Addon** | Pointer to **bool** |  | [optional] 
**Assignments** | Pointer to [**[]IPAssignment**](IPAssignment.md) |  | [optional] 
**Available** | Pointer to **string** |  | [optional] 
**Bill** | Pointer to **bool** |  | [optional] 
**Customdata** | Pointer to **map[string]interface{}** |  | [optional] 
**Details** | Pointer to **string** |  | [optional] 
**Facility** | Pointer to [**IPReservationFacility**](IPReservationFacility.md) |  | [optional] 
**MetalGateway** | Pointer to [**MetalGatewayLite**](MetalGatewayLite.md) |  | [optional] 
**Project** | Pointer to [**Project**](Project.md) |  | [optional] 
**ProjectLite** | Pointer to [**Project**](Project.md) |  | [optional] 
**RequestedBy** | Pointer to [**Href**](Href.md) |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Type** | **string** |  | 
**CreatedBy** | Pointer to [**Href**](Href.md) |  | [optional] 
**Vrf** | [**Vrf**](Vrf.md) |  | 

## Methods

### NewUpdateIPAddress200Response

`func NewUpdateIPAddress200Response(type_ string, vrf Vrf, ) *UpdateIPAddress200Response`

NewUpdateIPAddress200Response instantiates a new UpdateIPAddress200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateIPAddress200ResponseWithDefaults

`func NewUpdateIPAddress200ResponseWithDefaults() *UpdateIPAddress200Response`

NewUpdateIPAddress200ResponseWithDefaults instantiates a new UpdateIPAddress200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *UpdateIPAddress200Response) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *UpdateIPAddress200Response) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *UpdateIPAddress200Response) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *UpdateIPAddress200Response) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAddressFamily

`func (o *UpdateIPAddress200Response) GetAddressFamily() int32`

GetAddressFamily returns the AddressFamily field if non-nil, zero value otherwise.

### GetAddressFamilyOk

`func (o *UpdateIPAddress200Response) GetAddressFamilyOk() (*int32, bool)`

GetAddressFamilyOk returns a tuple with the AddressFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressFamily

`func (o *UpdateIPAddress200Response) SetAddressFamily(v int32)`

SetAddressFamily sets AddressFamily field to given value.

### HasAddressFamily

`func (o *UpdateIPAddress200Response) HasAddressFamily() bool`

HasAddressFamily returns a boolean if a field has been set.

### GetAssignedTo

`func (o *UpdateIPAddress200Response) GetAssignedTo() Href`

GetAssignedTo returns the AssignedTo field if non-nil, zero value otherwise.

### GetAssignedToOk

`func (o *UpdateIPAddress200Response) GetAssignedToOk() (*Href, bool)`

GetAssignedToOk returns a tuple with the AssignedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedTo

`func (o *UpdateIPAddress200Response) SetAssignedTo(v Href)`

SetAssignedTo sets AssignedTo field to given value.

### HasAssignedTo

`func (o *UpdateIPAddress200Response) HasAssignedTo() bool`

HasAssignedTo returns a boolean if a field has been set.

### GetCidr

`func (o *UpdateIPAddress200Response) GetCidr() int32`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *UpdateIPAddress200Response) GetCidrOk() (*int32, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *UpdateIPAddress200Response) SetCidr(v int32)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *UpdateIPAddress200Response) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetCreatedAt

`func (o *UpdateIPAddress200Response) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UpdateIPAddress200Response) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UpdateIPAddress200Response) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *UpdateIPAddress200Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEnabled

`func (o *UpdateIPAddress200Response) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateIPAddress200Response) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateIPAddress200Response) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateIPAddress200Response) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetGateway

`func (o *UpdateIPAddress200Response) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *UpdateIPAddress200Response) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *UpdateIPAddress200Response) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *UpdateIPAddress200Response) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetGlobalIp

`func (o *UpdateIPAddress200Response) GetGlobalIp() bool`

GetGlobalIp returns the GlobalIp field if non-nil, zero value otherwise.

### GetGlobalIpOk

`func (o *UpdateIPAddress200Response) GetGlobalIpOk() (*bool, bool)`

GetGlobalIpOk returns a tuple with the GlobalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalIp

`func (o *UpdateIPAddress200Response) SetGlobalIp(v bool)`

SetGlobalIp sets GlobalIp field to given value.

### HasGlobalIp

`func (o *UpdateIPAddress200Response) HasGlobalIp() bool`

HasGlobalIp returns a boolean if a field has been set.

### GetHref

`func (o *UpdateIPAddress200Response) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *UpdateIPAddress200Response) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *UpdateIPAddress200Response) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *UpdateIPAddress200Response) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *UpdateIPAddress200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateIPAddress200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateIPAddress200Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateIPAddress200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetManageable

`func (o *UpdateIPAddress200Response) GetManageable() bool`

GetManageable returns the Manageable field if non-nil, zero value otherwise.

### GetManageableOk

`func (o *UpdateIPAddress200Response) GetManageableOk() (*bool, bool)`

GetManageableOk returns a tuple with the Manageable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManageable

`func (o *UpdateIPAddress200Response) SetManageable(v bool)`

SetManageable sets Manageable field to given value.

### HasManageable

`func (o *UpdateIPAddress200Response) HasManageable() bool`

HasManageable returns a boolean if a field has been set.

### GetManagement

`func (o *UpdateIPAddress200Response) GetManagement() bool`

GetManagement returns the Management field if non-nil, zero value otherwise.

### GetManagementOk

`func (o *UpdateIPAddress200Response) GetManagementOk() (*bool, bool)`

GetManagementOk returns a tuple with the Management field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagement

`func (o *UpdateIPAddress200Response) SetManagement(v bool)`

SetManagement sets Management field to given value.

### HasManagement

`func (o *UpdateIPAddress200Response) HasManagement() bool`

HasManagement returns a boolean if a field has been set.

### GetMetro

`func (o *UpdateIPAddress200Response) GetMetro() Metro`

GetMetro returns the Metro field if non-nil, zero value otherwise.

### GetMetroOk

`func (o *UpdateIPAddress200Response) GetMetroOk() (*Metro, bool)`

GetMetroOk returns a tuple with the Metro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetro

`func (o *UpdateIPAddress200Response) SetMetro(v Metro)`

SetMetro sets Metro field to given value.

### HasMetro

`func (o *UpdateIPAddress200Response) HasMetro() bool`

HasMetro returns a boolean if a field has been set.

### GetNetmask

`func (o *UpdateIPAddress200Response) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *UpdateIPAddress200Response) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *UpdateIPAddress200Response) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *UpdateIPAddress200Response) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.

### GetNetwork

`func (o *UpdateIPAddress200Response) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *UpdateIPAddress200Response) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *UpdateIPAddress200Response) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *UpdateIPAddress200Response) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetParentBlock

`func (o *UpdateIPAddress200Response) GetParentBlock() ParentBlock`

GetParentBlock returns the ParentBlock field if non-nil, zero value otherwise.

### GetParentBlockOk

`func (o *UpdateIPAddress200Response) GetParentBlockOk() (*ParentBlock, bool)`

GetParentBlockOk returns a tuple with the ParentBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentBlock

`func (o *UpdateIPAddress200Response) SetParentBlock(v ParentBlock)`

SetParentBlock sets ParentBlock field to given value.

### HasParentBlock

`func (o *UpdateIPAddress200Response) HasParentBlock() bool`

HasParentBlock returns a boolean if a field has been set.

### GetPublic

`func (o *UpdateIPAddress200Response) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *UpdateIPAddress200Response) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *UpdateIPAddress200Response) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *UpdateIPAddress200Response) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetAddon

`func (o *UpdateIPAddress200Response) GetAddon() bool`

GetAddon returns the Addon field if non-nil, zero value otherwise.

### GetAddonOk

`func (o *UpdateIPAddress200Response) GetAddonOk() (*bool, bool)`

GetAddonOk returns a tuple with the Addon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddon

`func (o *UpdateIPAddress200Response) SetAddon(v bool)`

SetAddon sets Addon field to given value.

### HasAddon

`func (o *UpdateIPAddress200Response) HasAddon() bool`

HasAddon returns a boolean if a field has been set.

### GetAssignments

`func (o *UpdateIPAddress200Response) GetAssignments() []IPAssignment`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *UpdateIPAddress200Response) GetAssignmentsOk() (*[]IPAssignment, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignments

`func (o *UpdateIPAddress200Response) SetAssignments(v []IPAssignment)`

SetAssignments sets Assignments field to given value.

### HasAssignments

`func (o *UpdateIPAddress200Response) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### GetAvailable

`func (o *UpdateIPAddress200Response) GetAvailable() string`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *UpdateIPAddress200Response) GetAvailableOk() (*string, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *UpdateIPAddress200Response) SetAvailable(v string)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *UpdateIPAddress200Response) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetBill

`func (o *UpdateIPAddress200Response) GetBill() bool`

GetBill returns the Bill field if non-nil, zero value otherwise.

### GetBillOk

`func (o *UpdateIPAddress200Response) GetBillOk() (*bool, bool)`

GetBillOk returns a tuple with the Bill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBill

`func (o *UpdateIPAddress200Response) SetBill(v bool)`

SetBill sets Bill field to given value.

### HasBill

`func (o *UpdateIPAddress200Response) HasBill() bool`

HasBill returns a boolean if a field has been set.

### GetCustomdata

`func (o *UpdateIPAddress200Response) GetCustomdata() map[string]interface{}`

GetCustomdata returns the Customdata field if non-nil, zero value otherwise.

### GetCustomdataOk

`func (o *UpdateIPAddress200Response) GetCustomdataOk() (*map[string]interface{}, bool)`

GetCustomdataOk returns a tuple with the Customdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomdata

`func (o *UpdateIPAddress200Response) SetCustomdata(v map[string]interface{})`

SetCustomdata sets Customdata field to given value.

### HasCustomdata

`func (o *UpdateIPAddress200Response) HasCustomdata() bool`

HasCustomdata returns a boolean if a field has been set.

### GetDetails

`func (o *UpdateIPAddress200Response) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *UpdateIPAddress200Response) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *UpdateIPAddress200Response) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *UpdateIPAddress200Response) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetFacility

`func (o *UpdateIPAddress200Response) GetFacility() IPReservationFacility`

GetFacility returns the Facility field if non-nil, zero value otherwise.

### GetFacilityOk

`func (o *UpdateIPAddress200Response) GetFacilityOk() (*IPReservationFacility, bool)`

GetFacilityOk returns a tuple with the Facility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacility

`func (o *UpdateIPAddress200Response) SetFacility(v IPReservationFacility)`

SetFacility sets Facility field to given value.

### HasFacility

`func (o *UpdateIPAddress200Response) HasFacility() bool`

HasFacility returns a boolean if a field has been set.

### GetMetalGateway

`func (o *UpdateIPAddress200Response) GetMetalGateway() MetalGatewayLite`

GetMetalGateway returns the MetalGateway field if non-nil, zero value otherwise.

### GetMetalGatewayOk

`func (o *UpdateIPAddress200Response) GetMetalGatewayOk() (*MetalGatewayLite, bool)`

GetMetalGatewayOk returns a tuple with the MetalGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetalGateway

`func (o *UpdateIPAddress200Response) SetMetalGateway(v MetalGatewayLite)`

SetMetalGateway sets MetalGateway field to given value.

### HasMetalGateway

`func (o *UpdateIPAddress200Response) HasMetalGateway() bool`

HasMetalGateway returns a boolean if a field has been set.

### GetProject

`func (o *UpdateIPAddress200Response) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *UpdateIPAddress200Response) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *UpdateIPAddress200Response) SetProject(v Project)`

SetProject sets Project field to given value.

### HasProject

`func (o *UpdateIPAddress200Response) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetProjectLite

`func (o *UpdateIPAddress200Response) GetProjectLite() Project`

GetProjectLite returns the ProjectLite field if non-nil, zero value otherwise.

### GetProjectLiteOk

`func (o *UpdateIPAddress200Response) GetProjectLiteOk() (*Project, bool)`

GetProjectLiteOk returns a tuple with the ProjectLite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectLite

`func (o *UpdateIPAddress200Response) SetProjectLite(v Project)`

SetProjectLite sets ProjectLite field to given value.

### HasProjectLite

`func (o *UpdateIPAddress200Response) HasProjectLite() bool`

HasProjectLite returns a boolean if a field has been set.

### GetRequestedBy

`func (o *UpdateIPAddress200Response) GetRequestedBy() Href`

GetRequestedBy returns the RequestedBy field if non-nil, zero value otherwise.

### GetRequestedByOk

`func (o *UpdateIPAddress200Response) GetRequestedByOk() (*Href, bool)`

GetRequestedByOk returns a tuple with the RequestedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedBy

`func (o *UpdateIPAddress200Response) SetRequestedBy(v Href)`

SetRequestedBy sets RequestedBy field to given value.

### HasRequestedBy

`func (o *UpdateIPAddress200Response) HasRequestedBy() bool`

HasRequestedBy returns a boolean if a field has been set.

### GetState

`func (o *UpdateIPAddress200Response) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *UpdateIPAddress200Response) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *UpdateIPAddress200Response) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *UpdateIPAddress200Response) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTags

`func (o *UpdateIPAddress200Response) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateIPAddress200Response) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateIPAddress200Response) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateIPAddress200Response) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *UpdateIPAddress200Response) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateIPAddress200Response) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateIPAddress200Response) SetType(v string)`

SetType sets Type field to given value.


### GetCreatedBy

`func (o *UpdateIPAddress200Response) GetCreatedBy() Href`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *UpdateIPAddress200Response) GetCreatedByOk() (*Href, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *UpdateIPAddress200Response) SetCreatedBy(v Href)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *UpdateIPAddress200Response) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetVrf

`func (o *UpdateIPAddress200Response) GetVrf() Vrf`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *UpdateIPAddress200Response) GetVrfOk() (*Vrf, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *UpdateIPAddress200Response) SetVrf(v Vrf)`

SetVrf sets Vrf field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


