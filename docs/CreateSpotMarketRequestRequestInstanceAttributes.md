# CreateSpotMarketRequestRequestInstanceAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlwaysPxe** | Pointer to **bool** |  | [optional] 
**BillingCycle** | Pointer to **string** |  | [optional] 
**Customdata** | Pointer to **map[string]interface{}** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Features** | Pointer to **[]string** |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**Hostnames** | Pointer to **[]string** |  | [optional] 
**Locked** | Pointer to **bool** |  | [optional] 
**NoSshKeys** | Pointer to **bool** |  | [optional] 
**OperatingSystem** | Pointer to **string** |  | [optional] 
**Plan** | Pointer to **string** |  | [optional] 
**PrivateIpv4SubnetSize** | Pointer to **int32** |  | [optional] 
**ProjectSshKeys** | Pointer to **[]string** |  | [optional] 
**PublicIpv4SubnetSize** | Pointer to **int32** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**TerminationTime** | Pointer to **time.Time** |  | [optional] 
**UserSshKeys** | Pointer to **[]string** | The UUIDs of users whose SSH keys should be included on the provisioned device. | [optional] 
**Userdata** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateSpotMarketRequestRequestInstanceAttributes

`func NewCreateSpotMarketRequestRequestInstanceAttributes() *CreateSpotMarketRequestRequestInstanceAttributes`

NewCreateSpotMarketRequestRequestInstanceAttributes instantiates a new CreateSpotMarketRequestRequestInstanceAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSpotMarketRequestRequestInstanceAttributesWithDefaults

`func NewCreateSpotMarketRequestRequestInstanceAttributesWithDefaults() *CreateSpotMarketRequestRequestInstanceAttributes`

NewCreateSpotMarketRequestRequestInstanceAttributesWithDefaults instantiates a new CreateSpotMarketRequestRequestInstanceAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlwaysPxe

`func (o *CreateSpotMarketRequestRequestInstanceAttributes) GetAlwaysPxe() bool`

GetAlwaysPxe returns the AlwaysPxe field if non-nil, zero value otherwise.

### GetAlwaysPxeOk

`func (o *CreateSpotMarketRequestRequestInstanceAttributes) GetAlwaysPxeOk() (*bool, bool)`

GetAlwaysPxeOk returns a tuple with the AlwaysPxe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlwaysPxe

`func (o *CreateSpotMarketRequestRequestInstanceAttributes) SetAlwaysPxe(v bool)`

SetAlwaysPxe sets AlwaysPxe field to given value.

### HasAlwaysPxe

`func (o *CreateSpotMarketRequestRequestInstanceAttributes) HasAlwaysPxe() bool`

HasAlwaysPxe returns a boolean if a field has been set.

### GetBillingCycle

`func (o *CreateSpotMarketRequestRequestInstanceAttributes) GetBillingCycle() string`

GetBillingCycle returns the BillingCycle field if non-nil, zero value otherwise.

### GetBillingCycleOk

`func (o *CreateSpotMarketRequestRequestInstanceAttributes) GetBillingCycleOk() (*string, bool)`

GetBillingCycleOk returns a tuple with the BillingCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCycle

`func (o *CreateSpotMarketRequestRequestInstanceAttributes) SetBillingCycle(v string)`

SetBillingCycle sets BillingCycle field to given value.

### HasBillingCycle

`func (o *CreateSpotMarketRequestRequestInstanceAttributes) HasBillingCycle() bool`

HasBillingCycle returns a boolean if a field has been set.

### GetCustomdata

`func (o *CreateSpotMarketRequestRequestInstanceAttributes) GetCustomdata() map[string]interface{}`

GetCustomdata returns the Customdata field if non-nil, zero value otherwise.

### GetCustomdataOk

`func (o *CreateSpotMarketRequestRequestInstanceAttributes) GetCustomdataOk() (*map[string]interface{}, bool)`

GetCustomdataOk returns a tuple with the Customdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomdata

`func (o *CreateSpotMarketRequestRequestInstanceAttributes) SetCustomdata(v map[string]interface{})`

SetCustomdata sets Customdata field to given value.

### HasCustomdata

`func (o *CreateSpotMarketRequestRequestInstanceAttributes) HasCustomdata() bool`

HasCustomdata returns a boolean if a field has been set.

### GetDescription

`func (o *CreateSpotMarketRequestRequestInstanceAttributes) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateSpotMarketRequestRequestInstanceAttributes) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateSpotMarketRequestRequestInstanceAttributes) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateSpotMarketRequestRequestInstanceAttributes) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFeatures

`func (o *CreateSpotMarketRequestRequestInstanceAttributes) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *CreateSpotMarketRequestRequestInstanceAttributes) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *CreateSpotMarketRequestRequestInstanceAttributes) SetFeatures(v []string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *CreateSpotMarketRequestRequestInstanceAttributes) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetHostname

`func (o *CreateSpotMarketRequestRequestInstanceAttributes) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *CreateSpotMarketRequestRequestInstanceAttributes) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *CreateSpotMarketRequestRequestInstanceAttributes) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *CreateSpotMarketRequestRequestInstanceAttributes) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetHostnames

`func (o *CreateSpotMarketRequestRequestInstanceAttributes) GetHostnames() []string`

GetHostnames returns the Hostnames field if non-nil, zero value otherwise.

### GetHostnamesOk

`func (o *CreateSpotMarketRequestRequestInstanceAttributes) GetHostnamesOk() (*[]string, bool)`

GetHostnamesOk returns a tuple with the Hostnames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnames

`func (o *CreateSpotMarketRequestRequestInstanceAttributes) SetHostnames(v []string)`

SetHostnames sets Hostnames field to given value.

### HasHostnames

`func (o *CreateSpotMarketRequestRequestInstanceAttributes) HasHostnames() bool`

HasHostnames returns a boolean if a field has been set.

### GetLocked

`func (o *CreateSpotMarketRequestRequestInstanceAttributes) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *CreateSpotMarketRequestRequestInstanceAttributes) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *CreateSpotMarketRequestRequestInstanceAttributes) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *CreateSpotMarketRequestRequestInstanceAttributes) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetNoSshKeys

`func (o *CreateSpotMarketRequestRequestInstanceAttributes) GetNoSshKeys() bool`

GetNoSshKeys returns the NoSshKeys field if non-nil, zero value otherwise.

### GetNoSshKeysOk

`func (o *CreateSpotMarketRequestRequestInstanceAttributes) GetNoSshKeysOk() (*bool, bool)`

GetNoSshKeysOk returns a tuple with the NoSshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoSshKeys

`func (o *CreateSpotMarketRequestRequestInstanceAttributes) SetNoSshKeys(v bool)`

SetNoSshKeys sets NoSshKeys field to given value.

### HasNoSshKeys

`func (o *CreateSpotMarketRequestRequestInstanceAttributes) HasNoSshKeys() bool`

HasNoSshKeys returns a boolean if a field has been set.

### GetOperatingSystem

`func (o *CreateSpotMarketRequestRequestInstanceAttributes) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *CreateSpotMarketRequestRequestInstanceAttributes) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *CreateSpotMarketRequestRequestInstanceAttributes) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *CreateSpotMarketRequestRequestInstanceAttributes) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### GetPlan

`func (o *CreateSpotMarketRequestRequestInstanceAttributes) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *CreateSpotMarketRequestRequestInstanceAttributes) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *CreateSpotMarketRequestRequestInstanceAttributes) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *CreateSpotMarketRequestRequestInstanceAttributes) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetPrivateIpv4SubnetSize

`func (o *CreateSpotMarketRequestRequestInstanceAttributes) GetPrivateIpv4SubnetSize() int32`

GetPrivateIpv4SubnetSize returns the PrivateIpv4SubnetSize field if non-nil, zero value otherwise.

### GetPrivateIpv4SubnetSizeOk

`func (o *CreateSpotMarketRequestRequestInstanceAttributes) GetPrivateIpv4SubnetSizeOk() (*int32, bool)`

GetPrivateIpv4SubnetSizeOk returns a tuple with the PrivateIpv4SubnetSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIpv4SubnetSize

`func (o *CreateSpotMarketRequestRequestInstanceAttributes) SetPrivateIpv4SubnetSize(v int32)`

SetPrivateIpv4SubnetSize sets PrivateIpv4SubnetSize field to given value.

### HasPrivateIpv4SubnetSize

`func (o *CreateSpotMarketRequestRequestInstanceAttributes) HasPrivateIpv4SubnetSize() bool`

HasPrivateIpv4SubnetSize returns a boolean if a field has been set.

### GetProjectSshKeys

`func (o *CreateSpotMarketRequestRequestInstanceAttributes) GetProjectSshKeys() []string`

GetProjectSshKeys returns the ProjectSshKeys field if non-nil, zero value otherwise.

### GetProjectSshKeysOk

`func (o *CreateSpotMarketRequestRequestInstanceAttributes) GetProjectSshKeysOk() (*[]string, bool)`

GetProjectSshKeysOk returns a tuple with the ProjectSshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectSshKeys

`func (o *CreateSpotMarketRequestRequestInstanceAttributes) SetProjectSshKeys(v []string)`

SetProjectSshKeys sets ProjectSshKeys field to given value.

### HasProjectSshKeys

`func (o *CreateSpotMarketRequestRequestInstanceAttributes) HasProjectSshKeys() bool`

HasProjectSshKeys returns a boolean if a field has been set.

### GetPublicIpv4SubnetSize

`func (o *CreateSpotMarketRequestRequestInstanceAttributes) GetPublicIpv4SubnetSize() int32`

GetPublicIpv4SubnetSize returns the PublicIpv4SubnetSize field if non-nil, zero value otherwise.

### GetPublicIpv4SubnetSizeOk

`func (o *CreateSpotMarketRequestRequestInstanceAttributes) GetPublicIpv4SubnetSizeOk() (*int32, bool)`

GetPublicIpv4SubnetSizeOk returns a tuple with the PublicIpv4SubnetSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpv4SubnetSize

`func (o *CreateSpotMarketRequestRequestInstanceAttributes) SetPublicIpv4SubnetSize(v int32)`

SetPublicIpv4SubnetSize sets PublicIpv4SubnetSize field to given value.

### HasPublicIpv4SubnetSize

`func (o *CreateSpotMarketRequestRequestInstanceAttributes) HasPublicIpv4SubnetSize() bool`

HasPublicIpv4SubnetSize returns a boolean if a field has been set.

### GetTags

`func (o *CreateSpotMarketRequestRequestInstanceAttributes) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateSpotMarketRequestRequestInstanceAttributes) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateSpotMarketRequestRequestInstanceAttributes) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateSpotMarketRequestRequestInstanceAttributes) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTerminationTime

`func (o *CreateSpotMarketRequestRequestInstanceAttributes) GetTerminationTime() time.Time`

GetTerminationTime returns the TerminationTime field if non-nil, zero value otherwise.

### GetTerminationTimeOk

`func (o *CreateSpotMarketRequestRequestInstanceAttributes) GetTerminationTimeOk() (*time.Time, bool)`

GetTerminationTimeOk returns a tuple with the TerminationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationTime

`func (o *CreateSpotMarketRequestRequestInstanceAttributes) SetTerminationTime(v time.Time)`

SetTerminationTime sets TerminationTime field to given value.

### HasTerminationTime

`func (o *CreateSpotMarketRequestRequestInstanceAttributes) HasTerminationTime() bool`

HasTerminationTime returns a boolean if a field has been set.

### GetUserSshKeys

`func (o *CreateSpotMarketRequestRequestInstanceAttributes) GetUserSshKeys() []string`

GetUserSshKeys returns the UserSshKeys field if non-nil, zero value otherwise.

### GetUserSshKeysOk

`func (o *CreateSpotMarketRequestRequestInstanceAttributes) GetUserSshKeysOk() (*[]string, bool)`

GetUserSshKeysOk returns a tuple with the UserSshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSshKeys

`func (o *CreateSpotMarketRequestRequestInstanceAttributes) SetUserSshKeys(v []string)`

SetUserSshKeys sets UserSshKeys field to given value.

### HasUserSshKeys

`func (o *CreateSpotMarketRequestRequestInstanceAttributes) HasUserSshKeys() bool`

HasUserSshKeys returns a boolean if a field has been set.

### GetUserdata

`func (o *CreateSpotMarketRequestRequestInstanceAttributes) GetUserdata() string`

GetUserdata returns the Userdata field if non-nil, zero value otherwise.

### GetUserdataOk

`func (o *CreateSpotMarketRequestRequestInstanceAttributes) GetUserdataOk() (*string, bool)`

GetUserdataOk returns a tuple with the Userdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserdata

`func (o *CreateSpotMarketRequestRequestInstanceAttributes) SetUserdata(v string)`

SetUserdata sets Userdata field to given value.

### HasUserdata

`func (o *CreateSpotMarketRequestRequestInstanceAttributes) HasUserdata() bool`

HasUserdata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


