# FindDeviceMetadataByID200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Class** | Pointer to **string** |  | [optional] 
**Customdata** | Pointer to **map[string]interface{}** |  | [optional] 
**Facility** | Pointer to **string** | The facility code of the instance | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Iqn** | Pointer to **string** |  | [optional] 
**Metro** | Pointer to **string** | The metro code of the instance | [optional] 
**Network** | Pointer to [**FindDeviceMetadataByID200ResponseNetwork**](FindDeviceMetadataByID200ResponseNetwork.md) |  | [optional] 
**OperatingSystem** | Pointer to **map[string]interface{}** |  | [optional] 
**Plan** | Pointer to **string** | The plan slug of the instance | [optional] 
**PrivateSubnets** | Pointer to **[]string** | An array of the private subnets | [optional] 
**Reserved** | Pointer to **bool** |  | [optional] 
**Specs** | Pointer to **map[string]interface{}** | The specs of the plan version of the instance | [optional] 
**SshKeys** | Pointer to **[]string** |  | [optional] 
**SwitchShortId** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Volumes** | Pointer to **[]string** |  | [optional] 

## Methods

### NewFindDeviceMetadataByID200Response

`func NewFindDeviceMetadataByID200Response() *FindDeviceMetadataByID200Response`

NewFindDeviceMetadataByID200Response instantiates a new FindDeviceMetadataByID200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindDeviceMetadataByID200ResponseWithDefaults

`func NewFindDeviceMetadataByID200ResponseWithDefaults() *FindDeviceMetadataByID200Response`

NewFindDeviceMetadataByID200ResponseWithDefaults instantiates a new FindDeviceMetadataByID200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClass

`func (o *FindDeviceMetadataByID200Response) GetClass() string`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *FindDeviceMetadataByID200Response) GetClassOk() (*string, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *FindDeviceMetadataByID200Response) SetClass(v string)`

SetClass sets Class field to given value.

### HasClass

`func (o *FindDeviceMetadataByID200Response) HasClass() bool`

HasClass returns a boolean if a field has been set.

### GetCustomdata

`func (o *FindDeviceMetadataByID200Response) GetCustomdata() map[string]interface{}`

GetCustomdata returns the Customdata field if non-nil, zero value otherwise.

### GetCustomdataOk

`func (o *FindDeviceMetadataByID200Response) GetCustomdataOk() (*map[string]interface{}, bool)`

GetCustomdataOk returns a tuple with the Customdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomdata

`func (o *FindDeviceMetadataByID200Response) SetCustomdata(v map[string]interface{})`

SetCustomdata sets Customdata field to given value.

### HasCustomdata

`func (o *FindDeviceMetadataByID200Response) HasCustomdata() bool`

HasCustomdata returns a boolean if a field has been set.

### GetFacility

`func (o *FindDeviceMetadataByID200Response) GetFacility() string`

GetFacility returns the Facility field if non-nil, zero value otherwise.

### GetFacilityOk

`func (o *FindDeviceMetadataByID200Response) GetFacilityOk() (*string, bool)`

GetFacilityOk returns a tuple with the Facility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacility

`func (o *FindDeviceMetadataByID200Response) SetFacility(v string)`

SetFacility sets Facility field to given value.

### HasFacility

`func (o *FindDeviceMetadataByID200Response) HasFacility() bool`

HasFacility returns a boolean if a field has been set.

### GetHostname

`func (o *FindDeviceMetadataByID200Response) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *FindDeviceMetadataByID200Response) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *FindDeviceMetadataByID200Response) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *FindDeviceMetadataByID200Response) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetId

`func (o *FindDeviceMetadataByID200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FindDeviceMetadataByID200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FindDeviceMetadataByID200Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FindDeviceMetadataByID200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIqn

`func (o *FindDeviceMetadataByID200Response) GetIqn() string`

GetIqn returns the Iqn field if non-nil, zero value otherwise.

### GetIqnOk

`func (o *FindDeviceMetadataByID200Response) GetIqnOk() (*string, bool)`

GetIqnOk returns a tuple with the Iqn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIqn

`func (o *FindDeviceMetadataByID200Response) SetIqn(v string)`

SetIqn sets Iqn field to given value.

### HasIqn

`func (o *FindDeviceMetadataByID200Response) HasIqn() bool`

HasIqn returns a boolean if a field has been set.

### GetMetro

`func (o *FindDeviceMetadataByID200Response) GetMetro() string`

GetMetro returns the Metro field if non-nil, zero value otherwise.

### GetMetroOk

`func (o *FindDeviceMetadataByID200Response) GetMetroOk() (*string, bool)`

GetMetroOk returns a tuple with the Metro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetro

`func (o *FindDeviceMetadataByID200Response) SetMetro(v string)`

SetMetro sets Metro field to given value.

### HasMetro

`func (o *FindDeviceMetadataByID200Response) HasMetro() bool`

HasMetro returns a boolean if a field has been set.

### GetNetwork

`func (o *FindDeviceMetadataByID200Response) GetNetwork() FindDeviceMetadataByID200ResponseNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *FindDeviceMetadataByID200Response) GetNetworkOk() (*FindDeviceMetadataByID200ResponseNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *FindDeviceMetadataByID200Response) SetNetwork(v FindDeviceMetadataByID200ResponseNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *FindDeviceMetadataByID200Response) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetOperatingSystem

`func (o *FindDeviceMetadataByID200Response) GetOperatingSystem() map[string]interface{}`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *FindDeviceMetadataByID200Response) GetOperatingSystemOk() (*map[string]interface{}, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *FindDeviceMetadataByID200Response) SetOperatingSystem(v map[string]interface{})`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *FindDeviceMetadataByID200Response) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### GetPlan

`func (o *FindDeviceMetadataByID200Response) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *FindDeviceMetadataByID200Response) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *FindDeviceMetadataByID200Response) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *FindDeviceMetadataByID200Response) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetPrivateSubnets

`func (o *FindDeviceMetadataByID200Response) GetPrivateSubnets() []string`

GetPrivateSubnets returns the PrivateSubnets field if non-nil, zero value otherwise.

### GetPrivateSubnetsOk

`func (o *FindDeviceMetadataByID200Response) GetPrivateSubnetsOk() (*[]string, bool)`

GetPrivateSubnetsOk returns a tuple with the PrivateSubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateSubnets

`func (o *FindDeviceMetadataByID200Response) SetPrivateSubnets(v []string)`

SetPrivateSubnets sets PrivateSubnets field to given value.

### HasPrivateSubnets

`func (o *FindDeviceMetadataByID200Response) HasPrivateSubnets() bool`

HasPrivateSubnets returns a boolean if a field has been set.

### GetReserved

`func (o *FindDeviceMetadataByID200Response) GetReserved() bool`

GetReserved returns the Reserved field if non-nil, zero value otherwise.

### GetReservedOk

`func (o *FindDeviceMetadataByID200Response) GetReservedOk() (*bool, bool)`

GetReservedOk returns a tuple with the Reserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserved

`func (o *FindDeviceMetadataByID200Response) SetReserved(v bool)`

SetReserved sets Reserved field to given value.

### HasReserved

`func (o *FindDeviceMetadataByID200Response) HasReserved() bool`

HasReserved returns a boolean if a field has been set.

### GetSpecs

`func (o *FindDeviceMetadataByID200Response) GetSpecs() map[string]interface{}`

GetSpecs returns the Specs field if non-nil, zero value otherwise.

### GetSpecsOk

`func (o *FindDeviceMetadataByID200Response) GetSpecsOk() (*map[string]interface{}, bool)`

GetSpecsOk returns a tuple with the Specs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecs

`func (o *FindDeviceMetadataByID200Response) SetSpecs(v map[string]interface{})`

SetSpecs sets Specs field to given value.

### HasSpecs

`func (o *FindDeviceMetadataByID200Response) HasSpecs() bool`

HasSpecs returns a boolean if a field has been set.

### GetSshKeys

`func (o *FindDeviceMetadataByID200Response) GetSshKeys() []string`

GetSshKeys returns the SshKeys field if non-nil, zero value otherwise.

### GetSshKeysOk

`func (o *FindDeviceMetadataByID200Response) GetSshKeysOk() (*[]string, bool)`

GetSshKeysOk returns a tuple with the SshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeys

`func (o *FindDeviceMetadataByID200Response) SetSshKeys(v []string)`

SetSshKeys sets SshKeys field to given value.

### HasSshKeys

`func (o *FindDeviceMetadataByID200Response) HasSshKeys() bool`

HasSshKeys returns a boolean if a field has been set.

### GetSwitchShortId

`func (o *FindDeviceMetadataByID200Response) GetSwitchShortId() string`

GetSwitchShortId returns the SwitchShortId field if non-nil, zero value otherwise.

### GetSwitchShortIdOk

`func (o *FindDeviceMetadataByID200Response) GetSwitchShortIdOk() (*string, bool)`

GetSwitchShortIdOk returns a tuple with the SwitchShortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchShortId

`func (o *FindDeviceMetadataByID200Response) SetSwitchShortId(v string)`

SetSwitchShortId sets SwitchShortId field to given value.

### HasSwitchShortId

`func (o *FindDeviceMetadataByID200Response) HasSwitchShortId() bool`

HasSwitchShortId returns a boolean if a field has been set.

### GetTags

`func (o *FindDeviceMetadataByID200Response) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FindDeviceMetadataByID200Response) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FindDeviceMetadataByID200Response) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FindDeviceMetadataByID200Response) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVolumes

`func (o *FindDeviceMetadataByID200Response) GetVolumes() []string`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *FindDeviceMetadataByID200Response) GetVolumesOk() (*[]string, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *FindDeviceMetadataByID200Response) SetVolumes(v []string)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *FindDeviceMetadataByID200Response) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


