# ListConnectionPorts200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ports** | Pointer to [**[]GetInterconnection200ResponsePortsInner**](GetInterconnection200ResponsePortsInner.md) |  | [optional] 

## Methods

### NewListConnectionPorts200Response

`func NewListConnectionPorts200Response() *ListConnectionPorts200Response`

NewListConnectionPorts200Response instantiates a new ListConnectionPorts200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListConnectionPorts200ResponseWithDefaults

`func NewListConnectionPorts200ResponseWithDefaults() *ListConnectionPorts200Response`

NewListConnectionPorts200ResponseWithDefaults instantiates a new ListConnectionPorts200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPorts

`func (o *ListConnectionPorts200Response) GetPorts() []GetInterconnection200ResponsePortsInner`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *ListConnectionPorts200Response) GetPortsOk() (*[]GetInterconnection200ResponsePortsInner, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *ListConnectionPorts200Response) SetPorts(v []GetInterconnection200ResponsePortsInner)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *ListConnectionPorts200Response) HasPorts() bool`

HasPorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


