# UpdateVrfRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**IpRanges** | Pointer to **[]string** | A list of CIDR network addresses. Like [\&quot;10.0.0.0/16\&quot;, \&quot;2001:d78::/56\&quot;]. IPv4 blocks must be between /8 and /29 in size. IPv6 blocks must be between /56 and /64. A VRF\\&#39;s IP ranges must be defined in order to create VRF IP Reservations, which can then be used for Metal Gateways or Virtual Circuits. Adding a new CIDR address to the list will result in the creation of a new IP Range for this VRF. Removal of an existing CIDR address from the list will result in the deletion of an existing IP Range for this VRF. Deleting an IP Range will result in the deletion of any VRF IP Reservations contained within the IP Range, as well as the VRF IP Reservation\\&#39;s associated Metal Gateways or Virtual Circuits. If you do not wish to add or remove IP Ranges, either include the full existing list of IP Ranges in the update request, or do not specify the &#x60;ip_ranges&#x60; field in the update request. Specifying a value of &#x60;[]&#x60; will remove all existing IP Ranges from the VRF. | [optional] 
**LocalAsn** | Pointer to **int32** | The new &#x60;local_asn&#x60; value for the VRF. This field cannot be updated when there are active Interconnection Virtual Circuits associated to the VRF. | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateVrfRequest

`func NewUpdateVrfRequest() *UpdateVrfRequest`

NewUpdateVrfRequest instantiates a new UpdateVrfRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateVrfRequestWithDefaults

`func NewUpdateVrfRequestWithDefaults() *UpdateVrfRequest`

NewUpdateVrfRequestWithDefaults instantiates a new UpdateVrfRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UpdateVrfRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateVrfRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateVrfRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateVrfRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIpRanges

`func (o *UpdateVrfRequest) GetIpRanges() []string`

GetIpRanges returns the IpRanges field if non-nil, zero value otherwise.

### GetIpRangesOk

`func (o *UpdateVrfRequest) GetIpRangesOk() (*[]string, bool)`

GetIpRangesOk returns a tuple with the IpRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRanges

`func (o *UpdateVrfRequest) SetIpRanges(v []string)`

SetIpRanges sets IpRanges field to given value.

### HasIpRanges

`func (o *UpdateVrfRequest) HasIpRanges() bool`

HasIpRanges returns a boolean if a field has been set.

### GetLocalAsn

`func (o *UpdateVrfRequest) GetLocalAsn() int32`

GetLocalAsn returns the LocalAsn field if non-nil, zero value otherwise.

### GetLocalAsnOk

`func (o *UpdateVrfRequest) GetLocalAsnOk() (*int32, bool)`

GetLocalAsnOk returns a tuple with the LocalAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAsn

`func (o *UpdateVrfRequest) SetLocalAsn(v int32)`

SetLocalAsn sets LocalAsn field to given value.

### HasLocalAsn

`func (o *UpdateVrfRequest) HasLocalAsn() bool`

HasLocalAsn returns a boolean if a field has been set.

### GetName

`func (o *UpdateVrfRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateVrfRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateVrfRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateVrfRequest) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


