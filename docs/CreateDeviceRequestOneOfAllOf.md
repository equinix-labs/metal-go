# CreateDeviceRequestOneOfAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metro** | **string** | Metro code or ID of where the instance should be provisioned in.  Either metro or facility must be provided. | 

## Methods

### NewCreateDeviceRequestOneOfAllOf

`func NewCreateDeviceRequestOneOfAllOf(metro string, ) *CreateDeviceRequestOneOfAllOf`

NewCreateDeviceRequestOneOfAllOf instantiates a new CreateDeviceRequestOneOfAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDeviceRequestOneOfAllOfWithDefaults

`func NewCreateDeviceRequestOneOfAllOfWithDefaults() *CreateDeviceRequestOneOfAllOf`

NewCreateDeviceRequestOneOfAllOfWithDefaults instantiates a new CreateDeviceRequestOneOfAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetro

`func (o *CreateDeviceRequestOneOfAllOf) GetMetro() string`

GetMetro returns the Metro field if non-nil, zero value otherwise.

### GetMetroOk

`func (o *CreateDeviceRequestOneOfAllOf) GetMetroOk() (*string, bool)`

GetMetroOk returns a tuple with the Metro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetro

`func (o *CreateDeviceRequestOneOfAllOf) SetMetro(v string)`

SetMetro sets Metro field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


