# FindTrafficRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndedAt** | **time.Time** |  | 
**StartedAt** | **time.Time** |  | 

## Methods

### NewFindTrafficRequest

`func NewFindTrafficRequest(endedAt time.Time, startedAt time.Time, ) *FindTrafficRequest`

NewFindTrafficRequest instantiates a new FindTrafficRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindTrafficRequestWithDefaults

`func NewFindTrafficRequestWithDefaults() *FindTrafficRequest`

NewFindTrafficRequestWithDefaults instantiates a new FindTrafficRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndedAt

`func (o *FindTrafficRequest) GetEndedAt() time.Time`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *FindTrafficRequest) GetEndedAtOk() (*time.Time, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *FindTrafficRequest) SetEndedAt(v time.Time)`

SetEndedAt sets EndedAt field to given value.


### GetStartedAt

`func (o *FindTrafficRequest) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *FindTrafficRequest) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *FindTrafficRequest) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


