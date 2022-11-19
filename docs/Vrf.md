# Vrf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** | Optional field that can be set to describe the VRF | [optional] 
**LocalAsn** | Pointer to **int32** | A 4-byte ASN associated with the VRF. | [optional] 
**IpRanges** | Pointer to **[]string** | A list of CIDR network addresses. Like [\&quot;10.0.0.0/16\&quot;, \&quot;2001:d78::/56\&quot;]. | [optional] 
**Project** | Pointer to [**GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1VrfProject**](GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1VrfProject.md) |  | [optional] 
**Metro** | Pointer to [**GetInterconnection200ResponseMetroAllOf**](GetInterconnection200ResponseMetroAllOf.md) |  | [optional] 
**CreatedBy** | Pointer to [**GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1VrfCreatedBy**](GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1VrfCreatedBy.md) |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 

## Methods

### NewVrf

`func NewVrf() *Vrf`

NewVrf instantiates a new Vrf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVrfWithDefaults

`func NewVrfWithDefaults() *Vrf`

NewVrfWithDefaults instantiates a new Vrf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Vrf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Vrf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Vrf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Vrf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Vrf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Vrf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Vrf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Vrf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Vrf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Vrf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Vrf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Vrf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLocalAsn

`func (o *Vrf) GetLocalAsn() int32`

GetLocalAsn returns the LocalAsn field if non-nil, zero value otherwise.

### GetLocalAsnOk

`func (o *Vrf) GetLocalAsnOk() (*int32, bool)`

GetLocalAsnOk returns a tuple with the LocalAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAsn

`func (o *Vrf) SetLocalAsn(v int32)`

SetLocalAsn sets LocalAsn field to given value.

### HasLocalAsn

`func (o *Vrf) HasLocalAsn() bool`

HasLocalAsn returns a boolean if a field has been set.

### GetIpRanges

`func (o *Vrf) GetIpRanges() []string`

GetIpRanges returns the IpRanges field if non-nil, zero value otherwise.

### GetIpRangesOk

`func (o *Vrf) GetIpRangesOk() (*[]string, bool)`

GetIpRangesOk returns a tuple with the IpRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRanges

`func (o *Vrf) SetIpRanges(v []string)`

SetIpRanges sets IpRanges field to given value.

### HasIpRanges

`func (o *Vrf) HasIpRanges() bool`

HasIpRanges returns a boolean if a field has been set.

### GetProject

`func (o *Vrf) GetProject() GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1VrfProject`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *Vrf) GetProjectOk() (*GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1VrfProject, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *Vrf) SetProject(v GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1VrfProject)`

SetProject sets Project field to given value.

### HasProject

`func (o *Vrf) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetMetro

`func (o *Vrf) GetMetro() GetInterconnection200ResponseMetroAllOf`

GetMetro returns the Metro field if non-nil, zero value otherwise.

### GetMetroOk

`func (o *Vrf) GetMetroOk() (*GetInterconnection200ResponseMetroAllOf, bool)`

GetMetroOk returns a tuple with the Metro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetro

`func (o *Vrf) SetMetro(v GetInterconnection200ResponseMetroAllOf)`

SetMetro sets Metro field to given value.

### HasMetro

`func (o *Vrf) HasMetro() bool`

HasMetro returns a boolean if a field has been set.

### GetCreatedBy

`func (o *Vrf) GetCreatedBy() GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1VrfCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Vrf) GetCreatedByOk() (*GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1VrfCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Vrf) SetCreatedBy(v GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1VrfCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Vrf) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetHref

`func (o *Vrf) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *Vrf) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *Vrf) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *Vrf) HasHref() bool`

HasHref returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


