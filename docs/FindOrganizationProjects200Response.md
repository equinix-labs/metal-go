# FindOrganizationProjects200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**FindDeviceEvents200ResponseMeta**](FindDeviceEvents200ResponseMeta.md) |  | [optional] 
**Projects** | Pointer to [**[]GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1VrfProject**](GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1VrfProject.md) |  | [optional] 

## Methods

### NewFindOrganizationProjects200Response

`func NewFindOrganizationProjects200Response() *FindOrganizationProjects200Response`

NewFindOrganizationProjects200Response instantiates a new FindOrganizationProjects200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindOrganizationProjects200ResponseWithDefaults

`func NewFindOrganizationProjects200ResponseWithDefaults() *FindOrganizationProjects200Response`

NewFindOrganizationProjects200ResponseWithDefaults instantiates a new FindOrganizationProjects200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *FindOrganizationProjects200Response) GetMeta() FindDeviceEvents200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *FindOrganizationProjects200Response) GetMetaOk() (*FindDeviceEvents200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *FindOrganizationProjects200Response) SetMeta(v FindDeviceEvents200ResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *FindOrganizationProjects200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetProjects

`func (o *FindOrganizationProjects200Response) GetProjects() []GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1VrfProject`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *FindOrganizationProjects200Response) GetProjectsOk() (*[]GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1VrfProject, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *FindOrganizationProjects200Response) SetProjects(v []GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1VrfProject)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *FindOrganizationProjects200Response) HasProjects() bool`

HasProjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


