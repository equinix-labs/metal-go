# GetInterconnection200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to **[]string** |  | [optional] 
**ContactEmail** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Facility** | Pointer to [**FindBatchById200ResponseDevicesInner**](FindBatchById200ResponseDevicesInner.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Metro** | Pointer to [**GetInterconnection200ResponseMetro**](GetInterconnection200ResponseMetro.md) |  | [optional] 
**Mode** | Pointer to **string** | The mode of the connection (only relevant to dedicated connections). Shared connections won&#39;t have this field. Can be either &#39;standard&#39; or &#39;tunnel&#39;.   The default mode of a dedicated connection is &#39;standard&#39;. The mode can only be changed when there are no associated virtual circuits on the connection.   In tunnel mode, an 802.1q tunnel is added to a port to send/receive double tagged packets from server instances. | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Organization** | Pointer to [**FindBatchById200ResponseDevicesInner**](FindBatchById200ResponseDevicesInner.md) |  | [optional] 
**Ports** | Pointer to [**[]GetInterconnection200ResponsePortsInner**](GetInterconnection200ResponsePortsInner.md) |  | [optional] 
**Redundancy** | Pointer to **string** |  | [optional] 
**ServiceTokens** | Pointer to [**[]GetInterconnection200ResponseServiceTokensInner**](GetInterconnection200ResponseServiceTokensInner.md) |  | [optional] 
**Speed** | Pointer to **int32** | The connection&#39;s speed in bps. | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewGetInterconnection200Response

`func NewGetInterconnection200Response() *GetInterconnection200Response`

NewGetInterconnection200Response instantiates a new GetInterconnection200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInterconnection200ResponseWithDefaults

`func NewGetInterconnection200ResponseWithDefaults() *GetInterconnection200Response`

NewGetInterconnection200ResponseWithDefaults instantiates a new GetInterconnection200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *GetInterconnection200Response) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GetInterconnection200Response) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GetInterconnection200Response) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GetInterconnection200Response) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetContactEmail

`func (o *GetInterconnection200Response) GetContactEmail() string`

GetContactEmail returns the ContactEmail field if non-nil, zero value otherwise.

### GetContactEmailOk

`func (o *GetInterconnection200Response) GetContactEmailOk() (*string, bool)`

GetContactEmailOk returns a tuple with the ContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactEmail

`func (o *GetInterconnection200Response) SetContactEmail(v string)`

SetContactEmail sets ContactEmail field to given value.

### HasContactEmail

`func (o *GetInterconnection200Response) HasContactEmail() bool`

HasContactEmail returns a boolean if a field has been set.

### GetDescription

`func (o *GetInterconnection200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetInterconnection200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetInterconnection200Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetInterconnection200Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFacility

`func (o *GetInterconnection200Response) GetFacility() FindBatchById200ResponseDevicesInner`

GetFacility returns the Facility field if non-nil, zero value otherwise.

### GetFacilityOk

`func (o *GetInterconnection200Response) GetFacilityOk() (*FindBatchById200ResponseDevicesInner, bool)`

GetFacilityOk returns a tuple with the Facility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacility

`func (o *GetInterconnection200Response) SetFacility(v FindBatchById200ResponseDevicesInner)`

SetFacility sets Facility field to given value.

### HasFacility

`func (o *GetInterconnection200Response) HasFacility() bool`

HasFacility returns a boolean if a field has been set.

### GetId

`func (o *GetInterconnection200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetInterconnection200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetInterconnection200Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetInterconnection200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetro

`func (o *GetInterconnection200Response) GetMetro() GetInterconnection200ResponseMetro`

GetMetro returns the Metro field if non-nil, zero value otherwise.

### GetMetroOk

`func (o *GetInterconnection200Response) GetMetroOk() (*GetInterconnection200ResponseMetro, bool)`

GetMetroOk returns a tuple with the Metro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetro

`func (o *GetInterconnection200Response) SetMetro(v GetInterconnection200ResponseMetro)`

SetMetro sets Metro field to given value.

### HasMetro

`func (o *GetInterconnection200Response) HasMetro() bool`

HasMetro returns a boolean if a field has been set.

### GetMode

`func (o *GetInterconnection200Response) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *GetInterconnection200Response) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *GetInterconnection200Response) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *GetInterconnection200Response) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetName

`func (o *GetInterconnection200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetInterconnection200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetInterconnection200Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetInterconnection200Response) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrganization

`func (o *GetInterconnection200Response) GetOrganization() FindBatchById200ResponseDevicesInner`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *GetInterconnection200Response) GetOrganizationOk() (*FindBatchById200ResponseDevicesInner, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *GetInterconnection200Response) SetOrganization(v FindBatchById200ResponseDevicesInner)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *GetInterconnection200Response) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetPorts

`func (o *GetInterconnection200Response) GetPorts() []GetInterconnection200ResponsePortsInner`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *GetInterconnection200Response) GetPortsOk() (*[]GetInterconnection200ResponsePortsInner, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *GetInterconnection200Response) SetPorts(v []GetInterconnection200ResponsePortsInner)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *GetInterconnection200Response) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetRedundancy

`func (o *GetInterconnection200Response) GetRedundancy() string`

GetRedundancy returns the Redundancy field if non-nil, zero value otherwise.

### GetRedundancyOk

`func (o *GetInterconnection200Response) GetRedundancyOk() (*string, bool)`

GetRedundancyOk returns a tuple with the Redundancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancy

`func (o *GetInterconnection200Response) SetRedundancy(v string)`

SetRedundancy sets Redundancy field to given value.

### HasRedundancy

`func (o *GetInterconnection200Response) HasRedundancy() bool`

HasRedundancy returns a boolean if a field has been set.

### GetServiceTokens

`func (o *GetInterconnection200Response) GetServiceTokens() []GetInterconnection200ResponseServiceTokensInner`

GetServiceTokens returns the ServiceTokens field if non-nil, zero value otherwise.

### GetServiceTokensOk

`func (o *GetInterconnection200Response) GetServiceTokensOk() (*[]GetInterconnection200ResponseServiceTokensInner, bool)`

GetServiceTokensOk returns a tuple with the ServiceTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTokens

`func (o *GetInterconnection200Response) SetServiceTokens(v []GetInterconnection200ResponseServiceTokensInner)`

SetServiceTokens sets ServiceTokens field to given value.

### HasServiceTokens

`func (o *GetInterconnection200Response) HasServiceTokens() bool`

HasServiceTokens returns a boolean if a field has been set.

### GetSpeed

`func (o *GetInterconnection200Response) GetSpeed() int32`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *GetInterconnection200Response) GetSpeedOk() (*int32, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *GetInterconnection200Response) SetSpeed(v int32)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *GetInterconnection200Response) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetStatus

`func (o *GetInterconnection200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetInterconnection200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetInterconnection200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetInterconnection200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *GetInterconnection200Response) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetInterconnection200Response) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetInterconnection200Response) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetInterconnection200Response) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


