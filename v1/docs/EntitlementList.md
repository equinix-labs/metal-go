# EntitlementList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entitlements** | Pointer to [**[]Entitlement**](Entitlement.md) |  | [optional] 
**Meta** | Pointer to [**Meta**](Meta.md) |  | [optional] 

## Methods

### NewEntitlementList

`func NewEntitlementList() *EntitlementList`

NewEntitlementList instantiates a new EntitlementList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementListWithDefaults

`func NewEntitlementListWithDefaults() *EntitlementList`

NewEntitlementListWithDefaults instantiates a new EntitlementList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntitlements

`func (o *EntitlementList) GetEntitlements() []Entitlement`

GetEntitlements returns the Entitlements field if non-nil, zero value otherwise.

### GetEntitlementsOk

`func (o *EntitlementList) GetEntitlementsOk() (*[]Entitlement, bool)`

GetEntitlementsOk returns a tuple with the Entitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlements

`func (o *EntitlementList) SetEntitlements(v []Entitlement)`

SetEntitlements sets Entitlements field to given value.

### HasEntitlements

`func (o *EntitlementList) HasEntitlements() bool`

HasEntitlements returns a boolean if a field has been set.

### GetMeta

`func (o *EntitlementList) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *EntitlementList) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *EntitlementList) SetMeta(v Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *EntitlementList) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


