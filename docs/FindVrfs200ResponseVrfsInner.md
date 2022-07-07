# FindVrfs200ResponseVrfsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedBy** | Pointer to [**FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf**](FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf.md) |  | [optional] 
**Description** | Pointer to **string** | Optional field that can be set to describe the VRF | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IpRanges** | Pointer to **[]string** | A list of CIDR network addresses. Like [\&quot;10.0.0.0/16\&quot;, \&quot;2001:d78::/56\&quot;]. | [optional] 
**LocalAsn** | Pointer to **int32** | A 4-byte ASN associated with the VRF. | [optional] 
**Metro** | Pointer to [**GetInterconnection200ResponseMetroAllOf**](GetInterconnection200ResponseMetroAllOf.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Project** | Pointer to [**MoveHardwareReservation201ResponseProject**](MoveHardwareReservation201ResponseProject.md) |  | [optional] 

## Methods

### NewFindVrfs200ResponseVrfsInner

`func NewFindVrfs200ResponseVrfsInner() *FindVrfs200ResponseVrfsInner`

NewFindVrfs200ResponseVrfsInner instantiates a new FindVrfs200ResponseVrfsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindVrfs200ResponseVrfsInnerWithDefaults

`func NewFindVrfs200ResponseVrfsInnerWithDefaults() *FindVrfs200ResponseVrfsInner`

NewFindVrfs200ResponseVrfsInnerWithDefaults instantiates a new FindVrfs200ResponseVrfsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedBy

`func (o *FindVrfs200ResponseVrfsInner) GetCreatedBy() FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *FindVrfs200ResponseVrfsInner) GetCreatedByOk() (*FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *FindVrfs200ResponseVrfsInner) SetCreatedBy(v FindProjectAPIKeys200ResponseApiKeysInnerUserAllOf)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *FindVrfs200ResponseVrfsInner) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDescription

`func (o *FindVrfs200ResponseVrfsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FindVrfs200ResponseVrfsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FindVrfs200ResponseVrfsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FindVrfs200ResponseVrfsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHref

`func (o *FindVrfs200ResponseVrfsInner) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *FindVrfs200ResponseVrfsInner) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *FindVrfs200ResponseVrfsInner) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *FindVrfs200ResponseVrfsInner) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *FindVrfs200ResponseVrfsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FindVrfs200ResponseVrfsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FindVrfs200ResponseVrfsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FindVrfs200ResponseVrfsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIpRanges

`func (o *FindVrfs200ResponseVrfsInner) GetIpRanges() []string`

GetIpRanges returns the IpRanges field if non-nil, zero value otherwise.

### GetIpRangesOk

`func (o *FindVrfs200ResponseVrfsInner) GetIpRangesOk() (*[]string, bool)`

GetIpRangesOk returns a tuple with the IpRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRanges

`func (o *FindVrfs200ResponseVrfsInner) SetIpRanges(v []string)`

SetIpRanges sets IpRanges field to given value.

### HasIpRanges

`func (o *FindVrfs200ResponseVrfsInner) HasIpRanges() bool`

HasIpRanges returns a boolean if a field has been set.

### GetLocalAsn

`func (o *FindVrfs200ResponseVrfsInner) GetLocalAsn() int32`

GetLocalAsn returns the LocalAsn field if non-nil, zero value otherwise.

### GetLocalAsnOk

`func (o *FindVrfs200ResponseVrfsInner) GetLocalAsnOk() (*int32, bool)`

GetLocalAsnOk returns a tuple with the LocalAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAsn

`func (o *FindVrfs200ResponseVrfsInner) SetLocalAsn(v int32)`

SetLocalAsn sets LocalAsn field to given value.

### HasLocalAsn

`func (o *FindVrfs200ResponseVrfsInner) HasLocalAsn() bool`

HasLocalAsn returns a boolean if a field has been set.

### GetMetro

`func (o *FindVrfs200ResponseVrfsInner) GetMetro() GetInterconnection200ResponseMetroAllOf`

GetMetro returns the Metro field if non-nil, zero value otherwise.

### GetMetroOk

`func (o *FindVrfs200ResponseVrfsInner) GetMetroOk() (*GetInterconnection200ResponseMetroAllOf, bool)`

GetMetroOk returns a tuple with the Metro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetro

`func (o *FindVrfs200ResponseVrfsInner) SetMetro(v GetInterconnection200ResponseMetroAllOf)`

SetMetro sets Metro field to given value.

### HasMetro

`func (o *FindVrfs200ResponseVrfsInner) HasMetro() bool`

HasMetro returns a boolean if a field has been set.

### GetName

`func (o *FindVrfs200ResponseVrfsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FindVrfs200ResponseVrfsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FindVrfs200ResponseVrfsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FindVrfs200ResponseVrfsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProject

`func (o *FindVrfs200ResponseVrfsInner) GetProject() MoveHardwareReservation201ResponseProject`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *FindVrfs200ResponseVrfsInner) GetProjectOk() (*MoveHardwareReservation201ResponseProject, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *FindVrfs200ResponseVrfsInner) SetProject(v MoveHardwareReservation201ResponseProject)`

SetProject sets Project field to given value.

### HasProject

`func (o *FindVrfs200ResponseVrfsInner) HasProject() bool`

HasProject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


