# Timeframe

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartedAt** | **time.Time** |  | 
**EndedAt** | **time.Time** |  | 

## Methods

### NewTimeframe

`func NewTimeframe(startedAt time.Time, endedAt time.Time, ) *Timeframe`

NewTimeframe instantiates a new Timeframe object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeframeWithDefaults

`func NewTimeframeWithDefaults() *Timeframe`

NewTimeframeWithDefaults instantiates a new Timeframe object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartedAt

`func (o *Timeframe) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *Timeframe) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *Timeframe) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetEndedAt

`func (o *Timeframe) GetEndedAt() time.Time`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *Timeframe) GetEndedAtOk() (*time.Time, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *Timeframe) SetEndedAt(v time.Time)`

SetEndedAt sets EndedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


