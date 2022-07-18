# RequestSuppertRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | Pointer to **string** |  | [optional] 
**Message** | **string** |  | 
**Priority** | Pointer to **string** |  | [optional] 
**ProjectId** | Pointer to **string** |  | [optional] 
**Subject** | **string** |  | 

## Methods

### NewRequestSuppertRequest

`func NewRequestSuppertRequest(message string, subject string, ) *RequestSuppertRequest`

NewRequestSuppertRequest instantiates a new RequestSuppertRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestSuppertRequestWithDefaults

`func NewRequestSuppertRequestWithDefaults() *RequestSuppertRequest`

NewRequestSuppertRequestWithDefaults instantiates a new RequestSuppertRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *RequestSuppertRequest) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *RequestSuppertRequest) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *RequestSuppertRequest) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *RequestSuppertRequest) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetMessage

`func (o *RequestSuppertRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *RequestSuppertRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *RequestSuppertRequest) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetPriority

`func (o *RequestSuppertRequest) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *RequestSuppertRequest) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *RequestSuppertRequest) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *RequestSuppertRequest) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetProjectId

`func (o *RequestSuppertRequest) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *RequestSuppertRequest) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *RequestSuppertRequest) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *RequestSuppertRequest) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetSubject

`func (o *RequestSuppertRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *RequestSuppertRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *RequestSuppertRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


