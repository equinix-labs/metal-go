# IPReservation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**AddressFamily** | Pointer to **int32** |  | [optional] 
**Netmask** | Pointer to **string** |  | [optional] 
**Public** | Pointer to **bool** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Cidr** | Pointer to **int32** |  | [optional] 
**Management** | Pointer to **bool** |  | [optional] 
**Manageable** | Pointer to **bool** |  | [optional] 
**Addon** | Pointer to **bool** |  | [optional] 
**Bill** | Pointer to **bool** |  | [optional] 
**Assignments** | Pointer to [**[]IPAssignment**](IPAssignment.md) |  | [optional] 
**Network** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Facility** | Pointer to [**Facility**](Facility.md) |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Metro** | Pointer to [**Metro**](Metro.md) |  | [optional] 

## Methods

### NewIPReservation

`func NewIPReservation() *IPReservation`

NewIPReservation instantiates a new IPReservation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPReservationWithDefaults

`func NewIPReservationWithDefaults() *IPReservation`

NewIPReservationWithDefaults instantiates a new IPReservation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IPReservation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IPReservation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IPReservation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IPReservation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAddressFamily

`func (o *IPReservation) GetAddressFamily() int32`

GetAddressFamily returns the AddressFamily field if non-nil, zero value otherwise.

### GetAddressFamilyOk

`func (o *IPReservation) GetAddressFamilyOk() (*int32, bool)`

GetAddressFamilyOk returns a tuple with the AddressFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressFamily

`func (o *IPReservation) SetAddressFamily(v int32)`

SetAddressFamily sets AddressFamily field to given value.

### HasAddressFamily

`func (o *IPReservation) HasAddressFamily() bool`

HasAddressFamily returns a boolean if a field has been set.

### GetNetmask

`func (o *IPReservation) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *IPReservation) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *IPReservation) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *IPReservation) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.

### GetPublic

`func (o *IPReservation) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *IPReservation) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *IPReservation) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *IPReservation) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetEnabled

`func (o *IPReservation) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *IPReservation) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *IPReservation) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *IPReservation) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetCidr

`func (o *IPReservation) GetCidr() int32`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *IPReservation) GetCidrOk() (*int32, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *IPReservation) SetCidr(v int32)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *IPReservation) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetManagement

`func (o *IPReservation) GetManagement() bool`

GetManagement returns the Management field if non-nil, zero value otherwise.

### GetManagementOk

`func (o *IPReservation) GetManagementOk() (*bool, bool)`

GetManagementOk returns a tuple with the Management field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagement

`func (o *IPReservation) SetManagement(v bool)`

SetManagement sets Management field to given value.

### HasManagement

`func (o *IPReservation) HasManagement() bool`

HasManagement returns a boolean if a field has been set.

### GetManageable

`func (o *IPReservation) GetManageable() bool`

GetManageable returns the Manageable field if non-nil, zero value otherwise.

### GetManageableOk

`func (o *IPReservation) GetManageableOk() (*bool, bool)`

GetManageableOk returns a tuple with the Manageable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManageable

`func (o *IPReservation) SetManageable(v bool)`

SetManageable sets Manageable field to given value.

### HasManageable

`func (o *IPReservation) HasManageable() bool`

HasManageable returns a boolean if a field has been set.

### GetAddon

`func (o *IPReservation) GetAddon() bool`

GetAddon returns the Addon field if non-nil, zero value otherwise.

### GetAddonOk

`func (o *IPReservation) GetAddonOk() (*bool, bool)`

GetAddonOk returns a tuple with the Addon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddon

`func (o *IPReservation) SetAddon(v bool)`

SetAddon sets Addon field to given value.

### HasAddon

`func (o *IPReservation) HasAddon() bool`

HasAddon returns a boolean if a field has been set.

### GetBill

`func (o *IPReservation) GetBill() bool`

GetBill returns the Bill field if non-nil, zero value otherwise.

### GetBillOk

`func (o *IPReservation) GetBillOk() (*bool, bool)`

GetBillOk returns a tuple with the Bill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBill

`func (o *IPReservation) SetBill(v bool)`

SetBill sets Bill field to given value.

### HasBill

`func (o *IPReservation) HasBill() bool`

HasBill returns a boolean if a field has been set.

### GetAssignments

`func (o *IPReservation) GetAssignments() []IPAssignment`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *IPReservation) GetAssignmentsOk() (*[]IPAssignment, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignments

`func (o *IPReservation) SetAssignments(v []IPAssignment)`

SetAssignments sets Assignments field to given value.

### HasAssignments

`func (o *IPReservation) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### GetNetwork

`func (o *IPReservation) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *IPReservation) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *IPReservation) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *IPReservation) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetCreatedAt

`func (o *IPReservation) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IPReservation) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IPReservation) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IPReservation) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetFacility

`func (o *IPReservation) GetFacility() Facility`

GetFacility returns the Facility field if non-nil, zero value otherwise.

### GetFacilityOk

`func (o *IPReservation) GetFacilityOk() (*Facility, bool)`

GetFacilityOk returns a tuple with the Facility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacility

`func (o *IPReservation) SetFacility(v Facility)`

SetFacility sets Facility field to given value.

### HasFacility

`func (o *IPReservation) HasFacility() bool`

HasFacility returns a boolean if a field has been set.

### GetHref

`func (o *IPReservation) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *IPReservation) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *IPReservation) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *IPReservation) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetTags

`func (o *IPReservation) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *IPReservation) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *IPReservation) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *IPReservation) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetState

`func (o *IPReservation) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *IPReservation) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *IPReservation) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *IPReservation) HasState() bool`

HasState returns a boolean if a field has been set.

### GetMetro

`func (o *IPReservation) GetMetro() Metro`

GetMetro returns the Metro field if non-nil, zero value otherwise.

### GetMetroOk

`func (o *IPReservation) GetMetroOk() (*Metro, bool)`

GetMetroOk returns a tuple with the Metro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetro

`func (o *IPReservation) SetMetro(v Metro)`

SetMetro sets Metro field to given value.

### HasMetro

`func (o *IPReservation) HasMetro() bool`

HasMetro returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


