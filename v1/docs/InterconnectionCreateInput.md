# InterconnectionCreateInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Facility** | **string** | A Facility ID or code. | 
**Metro** | Pointer to **string** | A Metro ID or code. Required for creating a connection, unless creating with facility. | [optional] 
**Type** | **string** | Either &#39;shared&#39; or &#39;dedicated&#39;. | 
**Redundancy** | **string** | Either &#39;primary&#39; or &#39;redundant&#39;. | 
**ContactEmail** | Pointer to **string** |  | [optional] 
**Project** | Pointer to **string** |  | [optional] 
**Speed** | Pointer to **string** | A connection speed, in bps, mbps, or gbps. Ex: &#39;100000000&#39; or &#39;100 mbps&#39;. | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewInterconnectionCreateInput

`func NewInterconnectionCreateInput(name string, facility string, type_ string, redundancy string, ) *InterconnectionCreateInput`

NewInterconnectionCreateInput instantiates a new InterconnectionCreateInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInterconnectionCreateInputWithDefaults

`func NewInterconnectionCreateInputWithDefaults() *InterconnectionCreateInput`

NewInterconnectionCreateInputWithDefaults instantiates a new InterconnectionCreateInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InterconnectionCreateInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InterconnectionCreateInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InterconnectionCreateInput) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *InterconnectionCreateInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InterconnectionCreateInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InterconnectionCreateInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InterconnectionCreateInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFacility

`func (o *InterconnectionCreateInput) GetFacility() string`

GetFacility returns the Facility field if non-nil, zero value otherwise.

### GetFacilityOk

`func (o *InterconnectionCreateInput) GetFacilityOk() (*string, bool)`

GetFacilityOk returns a tuple with the Facility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacility

`func (o *InterconnectionCreateInput) SetFacility(v string)`

SetFacility sets Facility field to given value.


### GetMetro

`func (o *InterconnectionCreateInput) GetMetro() string`

GetMetro returns the Metro field if non-nil, zero value otherwise.

### GetMetroOk

`func (o *InterconnectionCreateInput) GetMetroOk() (*string, bool)`

GetMetroOk returns a tuple with the Metro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetro

`func (o *InterconnectionCreateInput) SetMetro(v string)`

SetMetro sets Metro field to given value.

### HasMetro

`func (o *InterconnectionCreateInput) HasMetro() bool`

HasMetro returns a boolean if a field has been set.

### GetType

`func (o *InterconnectionCreateInput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InterconnectionCreateInput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InterconnectionCreateInput) SetType(v string)`

SetType sets Type field to given value.


### GetRedundancy

`func (o *InterconnectionCreateInput) GetRedundancy() string`

GetRedundancy returns the Redundancy field if non-nil, zero value otherwise.

### GetRedundancyOk

`func (o *InterconnectionCreateInput) GetRedundancyOk() (*string, bool)`

GetRedundancyOk returns a tuple with the Redundancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancy

`func (o *InterconnectionCreateInput) SetRedundancy(v string)`

SetRedundancy sets Redundancy field to given value.


### GetContactEmail

`func (o *InterconnectionCreateInput) GetContactEmail() string`

GetContactEmail returns the ContactEmail field if non-nil, zero value otherwise.

### GetContactEmailOk

`func (o *InterconnectionCreateInput) GetContactEmailOk() (*string, bool)`

GetContactEmailOk returns a tuple with the ContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactEmail

`func (o *InterconnectionCreateInput) SetContactEmail(v string)`

SetContactEmail sets ContactEmail field to given value.

### HasContactEmail

`func (o *InterconnectionCreateInput) HasContactEmail() bool`

HasContactEmail returns a boolean if a field has been set.

### GetProject

`func (o *InterconnectionCreateInput) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *InterconnectionCreateInput) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *InterconnectionCreateInput) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *InterconnectionCreateInput) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetSpeed

`func (o *InterconnectionCreateInput) GetSpeed() string`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *InterconnectionCreateInput) GetSpeedOk() (*string, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *InterconnectionCreateInput) SetSpeed(v string)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *InterconnectionCreateInput) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetTags

`func (o *InterconnectionCreateInput) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InterconnectionCreateInput) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InterconnectionCreateInput) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InterconnectionCreateInput) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


