# DeviceCreateInFacilityInputAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Facility** | [**NullableAnyOfarraystring**](anyOf&lt;array,string&gt;.md) | The datacenter where the device should be created.  Either metro or facility must be provided.  The API will accept either a single facility &#x60;{ \&quot;facility\&quot;: \&quot;f1\&quot; }&#x60;, or it can be instructed to create the device in the best available datacenter &#x60;{ \&quot;facility\&quot;: \&quot;any\&quot; }&#x60;.  Additionally it is possible to set a prioritized location selection. For example &#x60;{ \&quot;facility\&quot;: [\&quot;f3\&quot;, \&quot;f2\&quot;, \&quot;any\&quot;] }&#x60; can be used to prioritize &#x60;f3&#x60; and then &#x60;f2&#x60; before accepting &#x60;any&#x60; facility. If none of the facilities provided have availability for the requested device the request will fail. | 

## Methods

### NewDeviceCreateInFacilityInputAllOf

`func NewDeviceCreateInFacilityInputAllOf(facility NullableAnyOfarraystring, ) *DeviceCreateInFacilityInputAllOf`

NewDeviceCreateInFacilityInputAllOf instantiates a new DeviceCreateInFacilityInputAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceCreateInFacilityInputAllOfWithDefaults

`func NewDeviceCreateInFacilityInputAllOfWithDefaults() *DeviceCreateInFacilityInputAllOf`

NewDeviceCreateInFacilityInputAllOfWithDefaults instantiates a new DeviceCreateInFacilityInputAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFacility

`func (o *DeviceCreateInFacilityInputAllOf) GetFacility() AnyOfarraystring`

GetFacility returns the Facility field if non-nil, zero value otherwise.

### GetFacilityOk

`func (o *DeviceCreateInFacilityInputAllOf) GetFacilityOk() (*AnyOfarraystring, bool)`

GetFacilityOk returns a tuple with the Facility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacility

`func (o *DeviceCreateInFacilityInputAllOf) SetFacility(v AnyOfarraystring)`

SetFacility sets Facility field to given value.


### SetFacilityNil

`func (o *DeviceCreateInFacilityInputAllOf) SetFacilityNil(b bool)`

 SetFacilityNil sets the value for Facility to be an explicit nil

### UnsetFacility
`func (o *DeviceCreateInFacilityInputAllOf) UnsetFacility()`

UnsetFacility ensures that no value is present for Facility, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


