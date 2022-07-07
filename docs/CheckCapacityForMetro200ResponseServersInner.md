# CheckCapacityForMetro200ResponseServersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **bool** | Returns true if there is enough capacity in the metro to fulfill the quantity set. Returns false if there is not enough. | [optional] 
**Metro** | Pointer to **string** | The metro ID or code sent to check capacity. | [optional] 
**Plan** | Pointer to **string** | The plan ID or slug sent to check capacity. | [optional] 
**Quantity** | Pointer to **string** | The number of servers sent to check capacity. | [optional] 

## Methods

### NewCheckCapacityForMetro200ResponseServersInner

`func NewCheckCapacityForMetro200ResponseServersInner() *CheckCapacityForMetro200ResponseServersInner`

NewCheckCapacityForMetro200ResponseServersInner instantiates a new CheckCapacityForMetro200ResponseServersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckCapacityForMetro200ResponseServersInnerWithDefaults

`func NewCheckCapacityForMetro200ResponseServersInnerWithDefaults() *CheckCapacityForMetro200ResponseServersInner`

NewCheckCapacityForMetro200ResponseServersInnerWithDefaults instantiates a new CheckCapacityForMetro200ResponseServersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *CheckCapacityForMetro200ResponseServersInner) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *CheckCapacityForMetro200ResponseServersInner) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *CheckCapacityForMetro200ResponseServersInner) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *CheckCapacityForMetro200ResponseServersInner) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetMetro

`func (o *CheckCapacityForMetro200ResponseServersInner) GetMetro() string`

GetMetro returns the Metro field if non-nil, zero value otherwise.

### GetMetroOk

`func (o *CheckCapacityForMetro200ResponseServersInner) GetMetroOk() (*string, bool)`

GetMetroOk returns a tuple with the Metro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetro

`func (o *CheckCapacityForMetro200ResponseServersInner) SetMetro(v string)`

SetMetro sets Metro field to given value.

### HasMetro

`func (o *CheckCapacityForMetro200ResponseServersInner) HasMetro() bool`

HasMetro returns a boolean if a field has been set.

### GetPlan

`func (o *CheckCapacityForMetro200ResponseServersInner) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *CheckCapacityForMetro200ResponseServersInner) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *CheckCapacityForMetro200ResponseServersInner) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *CheckCapacityForMetro200ResponseServersInner) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetQuantity

`func (o *CheckCapacityForMetro200ResponseServersInner) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CheckCapacityForMetro200ResponseServersInner) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CheckCapacityForMetro200ResponseServersInner) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *CheckCapacityForMetro200ResponseServersInner) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


