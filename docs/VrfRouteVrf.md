# VrfRouteVrf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | **string** |  | 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** | Optional field that can be set to describe the VRF | [optional] 
**Bill** | Pointer to **bool** | True if the VRF is being billed. Usage will start when the first VRF Virtual Circuit is active, and will only stop when the VRF has been deleted. | [optional] [default to false]
**BgpDynamicNeighborsEnabled** | Pointer to **bool** | Toggle to enable the dynamic bgp neighbors feature on the VRF | [optional] 
**BgpDynamicNeighborsExportRouteMap** | Pointer to **bool** | Toggle to export the VRF route-map to the dynamic bgp neighbors | [optional] 
**BgpDynamicNeighborsBfdEnabled** | Pointer to **bool** | Toggle BFD on dynamic bgp neighbors sessions | [optional] 
**LocalAsn** | Pointer to **int32** | A 4-byte ASN associated with the VRF. | [optional] 
**VirtualCircuits** | Pointer to [**[]VrfVirtualCircuitsInner**](VrfVirtualCircuitsInner.md) | Virtual circuits that are in the VRF | [optional] 
**IpRanges** | Pointer to **[]string** | A list of CIDR network addresses. Like [\&quot;10.0.0.0/16\&quot;, \&quot;2001:d78::/56\&quot;]. | [optional] 
**Project** | Pointer to [**Project**](Project.md) |  | [optional] 
**Metro** | Pointer to [**Metro**](Metro.md) |  | [optional] 
**CreatedBy** | Pointer to [**User**](User.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewVrfRouteVrf

`func NewVrfRouteVrf(href string, ) *VrfRouteVrf`

NewVrfRouteVrf instantiates a new VrfRouteVrf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVrfRouteVrfWithDefaults

`func NewVrfRouteVrfWithDefaults() *VrfRouteVrf`

NewVrfRouteVrfWithDefaults instantiates a new VrfRouteVrf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *VrfRouteVrf) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *VrfRouteVrf) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *VrfRouteVrf) SetHref(v string)`

SetHref sets Href field to given value.


### GetId

`func (o *VrfRouteVrf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VrfRouteVrf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VrfRouteVrf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VrfRouteVrf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *VrfRouteVrf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VrfRouteVrf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VrfRouteVrf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VrfRouteVrf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *VrfRouteVrf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VrfRouteVrf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VrfRouteVrf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VrfRouteVrf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetBill

`func (o *VrfRouteVrf) GetBill() bool`

GetBill returns the Bill field if non-nil, zero value otherwise.

### GetBillOk

`func (o *VrfRouteVrf) GetBillOk() (*bool, bool)`

GetBillOk returns a tuple with the Bill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBill

`func (o *VrfRouteVrf) SetBill(v bool)`

SetBill sets Bill field to given value.

### HasBill

`func (o *VrfRouteVrf) HasBill() bool`

HasBill returns a boolean if a field has been set.

### GetBgpDynamicNeighborsEnabled

`func (o *VrfRouteVrf) GetBgpDynamicNeighborsEnabled() bool`

GetBgpDynamicNeighborsEnabled returns the BgpDynamicNeighborsEnabled field if non-nil, zero value otherwise.

### GetBgpDynamicNeighborsEnabledOk

`func (o *VrfRouteVrf) GetBgpDynamicNeighborsEnabledOk() (*bool, bool)`

GetBgpDynamicNeighborsEnabledOk returns a tuple with the BgpDynamicNeighborsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpDynamicNeighborsEnabled

`func (o *VrfRouteVrf) SetBgpDynamicNeighborsEnabled(v bool)`

SetBgpDynamicNeighborsEnabled sets BgpDynamicNeighborsEnabled field to given value.

### HasBgpDynamicNeighborsEnabled

`func (o *VrfRouteVrf) HasBgpDynamicNeighborsEnabled() bool`

HasBgpDynamicNeighborsEnabled returns a boolean if a field has been set.

### GetBgpDynamicNeighborsExportRouteMap

`func (o *VrfRouteVrf) GetBgpDynamicNeighborsExportRouteMap() bool`

GetBgpDynamicNeighborsExportRouteMap returns the BgpDynamicNeighborsExportRouteMap field if non-nil, zero value otherwise.

### GetBgpDynamicNeighborsExportRouteMapOk

`func (o *VrfRouteVrf) GetBgpDynamicNeighborsExportRouteMapOk() (*bool, bool)`

GetBgpDynamicNeighborsExportRouteMapOk returns a tuple with the BgpDynamicNeighborsExportRouteMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpDynamicNeighborsExportRouteMap

`func (o *VrfRouteVrf) SetBgpDynamicNeighborsExportRouteMap(v bool)`

SetBgpDynamicNeighborsExportRouteMap sets BgpDynamicNeighborsExportRouteMap field to given value.

### HasBgpDynamicNeighborsExportRouteMap

`func (o *VrfRouteVrf) HasBgpDynamicNeighborsExportRouteMap() bool`

HasBgpDynamicNeighborsExportRouteMap returns a boolean if a field has been set.

### GetBgpDynamicNeighborsBfdEnabled

`func (o *VrfRouteVrf) GetBgpDynamicNeighborsBfdEnabled() bool`

GetBgpDynamicNeighborsBfdEnabled returns the BgpDynamicNeighborsBfdEnabled field if non-nil, zero value otherwise.

### GetBgpDynamicNeighborsBfdEnabledOk

`func (o *VrfRouteVrf) GetBgpDynamicNeighborsBfdEnabledOk() (*bool, bool)`

GetBgpDynamicNeighborsBfdEnabledOk returns a tuple with the BgpDynamicNeighborsBfdEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpDynamicNeighborsBfdEnabled

`func (o *VrfRouteVrf) SetBgpDynamicNeighborsBfdEnabled(v bool)`

SetBgpDynamicNeighborsBfdEnabled sets BgpDynamicNeighborsBfdEnabled field to given value.

### HasBgpDynamicNeighborsBfdEnabled

`func (o *VrfRouteVrf) HasBgpDynamicNeighborsBfdEnabled() bool`

HasBgpDynamicNeighborsBfdEnabled returns a boolean if a field has been set.

### GetLocalAsn

`func (o *VrfRouteVrf) GetLocalAsn() int32`

GetLocalAsn returns the LocalAsn field if non-nil, zero value otherwise.

### GetLocalAsnOk

`func (o *VrfRouteVrf) GetLocalAsnOk() (*int32, bool)`

GetLocalAsnOk returns a tuple with the LocalAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAsn

`func (o *VrfRouteVrf) SetLocalAsn(v int32)`

SetLocalAsn sets LocalAsn field to given value.

### HasLocalAsn

`func (o *VrfRouteVrf) HasLocalAsn() bool`

HasLocalAsn returns a boolean if a field has been set.

### GetVirtualCircuits

`func (o *VrfRouteVrf) GetVirtualCircuits() []VrfVirtualCircuitsInner`

GetVirtualCircuits returns the VirtualCircuits field if non-nil, zero value otherwise.

### GetVirtualCircuitsOk

`func (o *VrfRouteVrf) GetVirtualCircuitsOk() (*[]VrfVirtualCircuitsInner, bool)`

GetVirtualCircuitsOk returns a tuple with the VirtualCircuits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualCircuits

`func (o *VrfRouteVrf) SetVirtualCircuits(v []VrfVirtualCircuitsInner)`

SetVirtualCircuits sets VirtualCircuits field to given value.

### HasVirtualCircuits

`func (o *VrfRouteVrf) HasVirtualCircuits() bool`

HasVirtualCircuits returns a boolean if a field has been set.

### GetIpRanges

`func (o *VrfRouteVrf) GetIpRanges() []string`

GetIpRanges returns the IpRanges field if non-nil, zero value otherwise.

### GetIpRangesOk

`func (o *VrfRouteVrf) GetIpRangesOk() (*[]string, bool)`

GetIpRangesOk returns a tuple with the IpRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRanges

`func (o *VrfRouteVrf) SetIpRanges(v []string)`

SetIpRanges sets IpRanges field to given value.

### HasIpRanges

`func (o *VrfRouteVrf) HasIpRanges() bool`

HasIpRanges returns a boolean if a field has been set.

### GetProject

`func (o *VrfRouteVrf) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *VrfRouteVrf) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *VrfRouteVrf) SetProject(v Project)`

SetProject sets Project field to given value.

### HasProject

`func (o *VrfRouteVrf) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetMetro

`func (o *VrfRouteVrf) GetMetro() Metro`

GetMetro returns the Metro field if non-nil, zero value otherwise.

### GetMetroOk

`func (o *VrfRouteVrf) GetMetroOk() (*Metro, bool)`

GetMetroOk returns a tuple with the Metro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetro

`func (o *VrfRouteVrf) SetMetro(v Metro)`

SetMetro sets Metro field to given value.

### HasMetro

`func (o *VrfRouteVrf) HasMetro() bool`

HasMetro returns a boolean if a field has been set.

### GetCreatedBy

`func (o *VrfRouteVrf) GetCreatedBy() User`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *VrfRouteVrf) GetCreatedByOk() (*User, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *VrfRouteVrf) SetCreatedBy(v User)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *VrfRouteVrf) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedAt

`func (o *VrfRouteVrf) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VrfRouteVrf) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VrfRouteVrf) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *VrfRouteVrf) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *VrfRouteVrf) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *VrfRouteVrf) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *VrfRouteVrf) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *VrfRouteVrf) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


