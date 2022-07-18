# CreateOrganizationInterconnectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to **[]string** |  | [optional] 
**ContactEmail** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Metro** | **string** | A Metro ID or code. | 
**Mode** | Pointer to **string** | The mode of the connection (only relevant to dedicated connections). Shared connections won&#39;t have this field. Can be either &#39;standard&#39; or &#39;tunnel&#39;.   The default mode of a dedicated connection is &#39;standard&#39;. The mode can only be changed when there are no associated virtual circuits on the connection.   In tunnel mode, an 802.1q tunnel is added to a port to send/receive double tagged packets from server instances. | [optional] 
**Name** | **string** |  | 
**Project** | Pointer to **string** |  | [optional] 
**Redundancy** | **string** | Either &#39;primary&#39; or &#39;redundant&#39;. | 
**ServiceTokenType** | Pointer to **string** | Can only be set to &#39;a_side&#39; to create a shared connection with an A-side Fabric service token. This parameter is included in the specification as a developer preview and is generally unavailable. Please contact our Support team for more details. | [optional] 
**Speed** | Pointer to **string** | A connection speed, in bps, mbps, or gbps. Ex: &#39;100000000&#39; or &#39;100 mbps&#39;. | [optional] 
**Type** | **string** | Either &#39;shared&#39; or &#39;dedicated&#39;. | 
**Vlans** | Pointer to **[]int32** | A list of one or two metro-based VLANs that will be set on the primary and/or secondary (if redundant) virtual circuits respectively when creating a shared connection. VLANs can also be set after the connection is created, but are required to activate the connection. This parameter is included in the specification as a developer preview and is generally unavailable. Please contact our Support team for more details. | [optional] 

## Methods

### NewCreateOrganizationInterconnectionRequest

`func NewCreateOrganizationInterconnectionRequest(metro string, name string, redundancy string, type_ string, ) *CreateOrganizationInterconnectionRequest`

NewCreateOrganizationInterconnectionRequest instantiates a new CreateOrganizationInterconnectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrganizationInterconnectionRequestWithDefaults

`func NewCreateOrganizationInterconnectionRequestWithDefaults() *CreateOrganizationInterconnectionRequest`

NewCreateOrganizationInterconnectionRequestWithDefaults instantiates a new CreateOrganizationInterconnectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *CreateOrganizationInterconnectionRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateOrganizationInterconnectionRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateOrganizationInterconnectionRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateOrganizationInterconnectionRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetContactEmail

`func (o *CreateOrganizationInterconnectionRequest) GetContactEmail() string`

GetContactEmail returns the ContactEmail field if non-nil, zero value otherwise.

### GetContactEmailOk

`func (o *CreateOrganizationInterconnectionRequest) GetContactEmailOk() (*string, bool)`

GetContactEmailOk returns a tuple with the ContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactEmail

`func (o *CreateOrganizationInterconnectionRequest) SetContactEmail(v string)`

SetContactEmail sets ContactEmail field to given value.

### HasContactEmail

`func (o *CreateOrganizationInterconnectionRequest) HasContactEmail() bool`

HasContactEmail returns a boolean if a field has been set.

### GetDescription

`func (o *CreateOrganizationInterconnectionRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateOrganizationInterconnectionRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateOrganizationInterconnectionRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateOrganizationInterconnectionRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetro

`func (o *CreateOrganizationInterconnectionRequest) GetMetro() string`

GetMetro returns the Metro field if non-nil, zero value otherwise.

### GetMetroOk

`func (o *CreateOrganizationInterconnectionRequest) GetMetroOk() (*string, bool)`

GetMetroOk returns a tuple with the Metro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetro

`func (o *CreateOrganizationInterconnectionRequest) SetMetro(v string)`

SetMetro sets Metro field to given value.


### GetMode

`func (o *CreateOrganizationInterconnectionRequest) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *CreateOrganizationInterconnectionRequest) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *CreateOrganizationInterconnectionRequest) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *CreateOrganizationInterconnectionRequest) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetName

`func (o *CreateOrganizationInterconnectionRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateOrganizationInterconnectionRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateOrganizationInterconnectionRequest) SetName(v string)`

SetName sets Name field to given value.


### GetProject

`func (o *CreateOrganizationInterconnectionRequest) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *CreateOrganizationInterconnectionRequest) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *CreateOrganizationInterconnectionRequest) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *CreateOrganizationInterconnectionRequest) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetRedundancy

`func (o *CreateOrganizationInterconnectionRequest) GetRedundancy() string`

GetRedundancy returns the Redundancy field if non-nil, zero value otherwise.

### GetRedundancyOk

`func (o *CreateOrganizationInterconnectionRequest) GetRedundancyOk() (*string, bool)`

GetRedundancyOk returns a tuple with the Redundancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancy

`func (o *CreateOrganizationInterconnectionRequest) SetRedundancy(v string)`

SetRedundancy sets Redundancy field to given value.


### GetServiceTokenType

`func (o *CreateOrganizationInterconnectionRequest) GetServiceTokenType() string`

GetServiceTokenType returns the ServiceTokenType field if non-nil, zero value otherwise.

### GetServiceTokenTypeOk

`func (o *CreateOrganizationInterconnectionRequest) GetServiceTokenTypeOk() (*string, bool)`

GetServiceTokenTypeOk returns a tuple with the ServiceTokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTokenType

`func (o *CreateOrganizationInterconnectionRequest) SetServiceTokenType(v string)`

SetServiceTokenType sets ServiceTokenType field to given value.

### HasServiceTokenType

`func (o *CreateOrganizationInterconnectionRequest) HasServiceTokenType() bool`

HasServiceTokenType returns a boolean if a field has been set.

### GetSpeed

`func (o *CreateOrganizationInterconnectionRequest) GetSpeed() string`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *CreateOrganizationInterconnectionRequest) GetSpeedOk() (*string, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *CreateOrganizationInterconnectionRequest) SetSpeed(v string)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *CreateOrganizationInterconnectionRequest) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetType

`func (o *CreateOrganizationInterconnectionRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateOrganizationInterconnectionRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateOrganizationInterconnectionRequest) SetType(v string)`

SetType sets Type field to given value.


### GetVlans

`func (o *CreateOrganizationInterconnectionRequest) GetVlans() []int32`

GetVlans returns the Vlans field if non-nil, zero value otherwise.

### GetVlansOk

`func (o *CreateOrganizationInterconnectionRequest) GetVlansOk() (*[]int32, bool)`

GetVlansOk returns a tuple with the Vlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlans

`func (o *CreateOrganizationInterconnectionRequest) SetVlans(v []int32)`

SetVlans sets Vlans field to given value.

### HasVlans

`func (o *CreateOrganizationInterconnectionRequest) HasVlans() bool`

HasVlans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


