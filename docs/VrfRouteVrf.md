# VrfRouteVrf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | **string** |  | 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** | Optional field that can be set to describe the VRF | [optional] 
**LocalAsn** | Pointer to **int32** | A 4-byte ASN associated with the VRF. | [optional] 
**IpRanges** | Pointer to **[]string** | A list of CIDR network addresses. Like [\&quot;10.0.0.0/16\&quot;, \&quot;2001:d78::/56\&quot;]. | [optional] 
**Project** | Pointer to [**Project**](Project.md) |  | [optional] 
**Metro** | Pointer to [**Metro**](Metro.md) |  | [optional] 
**CreatedBy** | Pointer to [**User**](User.md) |  | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


