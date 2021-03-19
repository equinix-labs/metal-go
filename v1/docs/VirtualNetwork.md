# VirtualNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Vxlan** | Pointer to **int32** |  | [optional] 
**Facility** | Pointer to [**Href**](Href.md) |  | [optional] 
**Instances** | Pointer to [**[]Href**](Href.md) | A list of instances with ports currently associated to this Virtual Network. | [optional] 
**AssignedTo** | Pointer to [**Href**](Href.md) |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 

## Methods

### NewVirtualNetwork

`func NewVirtualNetwork() *VirtualNetwork`

NewVirtualNetwork instantiates a new VirtualNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualNetworkWithDefaults

`func NewVirtualNetworkWithDefaults() *VirtualNetwork`

NewVirtualNetworkWithDefaults instantiates a new VirtualNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VirtualNetwork) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VirtualNetwork) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VirtualNetwork) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VirtualNetwork) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescription

`func (o *VirtualNetwork) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VirtualNetwork) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VirtualNetwork) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VirtualNetwork) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVxlan

`func (o *VirtualNetwork) GetVxlan() int32`

GetVxlan returns the Vxlan field if non-nil, zero value otherwise.

### GetVxlanOk

`func (o *VirtualNetwork) GetVxlanOk() (*int32, bool)`

GetVxlanOk returns a tuple with the Vxlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVxlan

`func (o *VirtualNetwork) SetVxlan(v int32)`

SetVxlan sets Vxlan field to given value.

### HasVxlan

`func (o *VirtualNetwork) HasVxlan() bool`

HasVxlan returns a boolean if a field has been set.

### GetFacility

`func (o *VirtualNetwork) GetFacility() Href`

GetFacility returns the Facility field if non-nil, zero value otherwise.

### GetFacilityOk

`func (o *VirtualNetwork) GetFacilityOk() (*Href, bool)`

GetFacilityOk returns a tuple with the Facility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacility

`func (o *VirtualNetwork) SetFacility(v Href)`

SetFacility sets Facility field to given value.

### HasFacility

`func (o *VirtualNetwork) HasFacility() bool`

HasFacility returns a boolean if a field has been set.

### GetInstances

`func (o *VirtualNetwork) GetInstances() []Href`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *VirtualNetwork) GetInstancesOk() (*[]Href, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *VirtualNetwork) SetInstances(v []Href)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *VirtualNetwork) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

### GetAssignedTo

`func (o *VirtualNetwork) GetAssignedTo() Href`

GetAssignedTo returns the AssignedTo field if non-nil, zero value otherwise.

### GetAssignedToOk

`func (o *VirtualNetwork) GetAssignedToOk() (*Href, bool)`

GetAssignedToOk returns a tuple with the AssignedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedTo

`func (o *VirtualNetwork) SetAssignedTo(v Href)`

SetAssignedTo sets AssignedTo field to given value.

### HasAssignedTo

`func (o *VirtualNetwork) HasAssignedTo() bool`

HasAssignedTo returns a boolean if a field has been set.

### GetHref

`func (o *VirtualNetwork) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *VirtualNetwork) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *VirtualNetwork) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *VirtualNetwork) HasHref() bool`

HasHref returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


