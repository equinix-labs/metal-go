# AssignPortRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vnid** | Pointer to **string** | Virtual Network ID. May be the UUID of the Virtual Network record, or the VLAN value itself.  | [optional] 

## Methods

### NewAssignPortRequest

`func NewAssignPortRequest() *AssignPortRequest`

NewAssignPortRequest instantiates a new AssignPortRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssignPortRequestWithDefaults

`func NewAssignPortRequestWithDefaults() *AssignPortRequest`

NewAssignPortRequestWithDefaults instantiates a new AssignPortRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVnid

`func (o *AssignPortRequest) GetVnid() string`

GetVnid returns the Vnid field if non-nil, zero value otherwise.

### GetVnidOk

`func (o *AssignPortRequest) GetVnidOk() (*string, bool)`

GetVnidOk returns a tuple with the Vnid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnid

`func (o *AssignPortRequest) SetVnid(v string)`

SetVnid sets Vnid field to given value.

### HasVnid

`func (o *AssignPortRequest) HasVnid() bool`

HasVnid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


