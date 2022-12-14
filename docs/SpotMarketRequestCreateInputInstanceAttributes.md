# SpotMarketRequestCreateInputInstanceAttributes

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

### NewSpotMarketRequestCreateInputInstanceAttributes

`func NewSpotMarketRequestCreateInputInstanceAttributes() *SpotMarketRequestCreateInputInstanceAttributes`

NewSpotMarketRequestCreateInputInstanceAttributes instantiates a new SpotMarketRequestCreateInputInstanceAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpotMarketRequestCreateInputInstanceAttributesWithDefaults

`func NewSpotMarketRequestCreateInputInstanceAttributesWithDefaults() *SpotMarketRequestCreateInputInstanceAttributes`

NewSpotMarketRequestCreateInputInstanceAttributesWithDefaults instantiates a new SpotMarketRequestCreateInputInstanceAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlwaysPxe

`func (o *SpotMarketRequestCreateInputInstanceAttributes) GetAlwaysPxe() bool`

GetAlwaysPxe returns the AlwaysPxe field if non-nil, zero value otherwise.

### GetAlwaysPxeOk

`func (o *SpotMarketRequestCreateInputInstanceAttributes) GetAlwaysPxeOk() (*bool, bool)`

GetAlwaysPxeOk returns a tuple with the AlwaysPxe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlwaysPxe

`func (o *SpotMarketRequestCreateInputInstanceAttributes) SetAlwaysPxe(v bool)`

SetAlwaysPxe sets AlwaysPxe field to given value.

### HasAlwaysPxe

`func (o *SpotMarketRequestCreateInputInstanceAttributes) HasAlwaysPxe() bool`

HasAlwaysPxe returns a boolean if a field has been set.

### GetBillingCycle

`func (o *SpotMarketRequestCreateInputInstanceAttributes) GetBillingCycle() string`

GetBillingCycle returns the BillingCycle field if non-nil, zero value otherwise.

### GetBillingCycleOk

`func (o *SpotMarketRequestCreateInputInstanceAttributes) GetBillingCycleOk() (*string, bool)`

GetBillingCycleOk returns a tuple with the BillingCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCycle

`func (o *SpotMarketRequestCreateInputInstanceAttributes) SetBillingCycle(v string)`

SetBillingCycle sets BillingCycle field to given value.

### HasBillingCycle

`func (o *SpotMarketRequestCreateInputInstanceAttributes) HasBillingCycle() bool`

HasBillingCycle returns a boolean if a field has been set.

### GetCustomdata

`func (o *SpotMarketRequestCreateInputInstanceAttributes) GetCustomdata() map[string]interface{}`

GetCustomdata returns the Customdata field if non-nil, zero value otherwise.

### GetCustomdataOk

`func (o *SpotMarketRequestCreateInputInstanceAttributes) GetCustomdataOk() (*map[string]interface{}, bool)`

GetCustomdataOk returns a tuple with the Customdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomdata

`func (o *SpotMarketRequestCreateInputInstanceAttributes) SetCustomdata(v map[string]interface{})`

SetCustomdata sets Customdata field to given value.

### HasCustomdata

`func (o *SpotMarketRequestCreateInputInstanceAttributes) HasCustomdata() bool`

HasCustomdata returns a boolean if a field has been set.

### GetDescription

`func (o *SpotMarketRequestCreateInputInstanceAttributes) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SpotMarketRequestCreateInputInstanceAttributes) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SpotMarketRequestCreateInputInstanceAttributes) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SpotMarketRequestCreateInputInstanceAttributes) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFeatures

`func (o *SpotMarketRequestCreateInputInstanceAttributes) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *SpotMarketRequestCreateInputInstanceAttributes) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *SpotMarketRequestCreateInputInstanceAttributes) SetFeatures(v []string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *SpotMarketRequestCreateInputInstanceAttributes) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetHostname

`func (o *SpotMarketRequestCreateInputInstanceAttributes) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *SpotMarketRequestCreateInputInstanceAttributes) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *SpotMarketRequestCreateInputInstanceAttributes) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *SpotMarketRequestCreateInputInstanceAttributes) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetHostnames

`func (o *SpotMarketRequestCreateInputInstanceAttributes) GetHostnames() []string`

GetHostnames returns the Hostnames field if non-nil, zero value otherwise.

### GetHostnamesOk

`func (o *SpotMarketRequestCreateInputInstanceAttributes) GetHostnamesOk() (*[]string, bool)`

GetHostnamesOk returns a tuple with the Hostnames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnames

`func (o *SpotMarketRequestCreateInputInstanceAttributes) SetHostnames(v []string)`

SetHostnames sets Hostnames field to given value.

### HasHostnames

`func (o *SpotMarketRequestCreateInputInstanceAttributes) HasHostnames() bool`

HasHostnames returns a boolean if a field has been set.

### GetLocked

`func (o *SpotMarketRequestCreateInputInstanceAttributes) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *SpotMarketRequestCreateInputInstanceAttributes) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *SpotMarketRequestCreateInputInstanceAttributes) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *SpotMarketRequestCreateInputInstanceAttributes) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetNoSshKeys

`func (o *SpotMarketRequestCreateInputInstanceAttributes) GetNoSshKeys() bool`

GetNoSshKeys returns the NoSshKeys field if non-nil, zero value otherwise.

### GetNoSshKeysOk

`func (o *SpotMarketRequestCreateInputInstanceAttributes) GetNoSshKeysOk() (*bool, bool)`

GetNoSshKeysOk returns a tuple with the NoSshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoSshKeys

`func (o *SpotMarketRequestCreateInputInstanceAttributes) SetNoSshKeys(v bool)`

SetNoSshKeys sets NoSshKeys field to given value.

### HasNoSshKeys

`func (o *SpotMarketRequestCreateInputInstanceAttributes) HasNoSshKeys() bool`

HasNoSshKeys returns a boolean if a field has been set.

### GetOperatingSystem

`func (o *SpotMarketRequestCreateInputInstanceAttributes) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *SpotMarketRequestCreateInputInstanceAttributes) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *SpotMarketRequestCreateInputInstanceAttributes) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *SpotMarketRequestCreateInputInstanceAttributes) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### GetPlan

`func (o *SpotMarketRequestCreateInputInstanceAttributes) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *SpotMarketRequestCreateInputInstanceAttributes) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *SpotMarketRequestCreateInputInstanceAttributes) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *SpotMarketRequestCreateInputInstanceAttributes) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetPrivateIpv4SubnetSize

`func (o *SpotMarketRequestCreateInputInstanceAttributes) GetPrivateIpv4SubnetSize() int32`

GetPrivateIpv4SubnetSize returns the PrivateIpv4SubnetSize field if non-nil, zero value otherwise.

### GetPrivateIpv4SubnetSizeOk

`func (o *SpotMarketRequestCreateInputInstanceAttributes) GetPrivateIpv4SubnetSizeOk() (*int32, bool)`

GetPrivateIpv4SubnetSizeOk returns a tuple with the PrivateIpv4SubnetSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIpv4SubnetSize

`func (o *SpotMarketRequestCreateInputInstanceAttributes) SetPrivateIpv4SubnetSize(v int32)`

SetPrivateIpv4SubnetSize sets PrivateIpv4SubnetSize field to given value.

### HasPrivateIpv4SubnetSize

`func (o *SpotMarketRequestCreateInputInstanceAttributes) HasPrivateIpv4SubnetSize() bool`

HasPrivateIpv4SubnetSize returns a boolean if a field has been set.

### GetProjectSshKeys

`func (o *SpotMarketRequestCreateInputInstanceAttributes) GetProjectSshKeys() []string`

GetProjectSshKeys returns the ProjectSshKeys field if non-nil, zero value otherwise.

### GetProjectSshKeysOk

`func (o *SpotMarketRequestCreateInputInstanceAttributes) GetProjectSshKeysOk() (*[]string, bool)`

GetProjectSshKeysOk returns a tuple with the ProjectSshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectSshKeys

`func (o *SpotMarketRequestCreateInputInstanceAttributes) SetProjectSshKeys(v []string)`

SetProjectSshKeys sets ProjectSshKeys field to given value.

### HasProjectSshKeys

`func (o *SpotMarketRequestCreateInputInstanceAttributes) HasProjectSshKeys() bool`

HasProjectSshKeys returns a boolean if a field has been set.

### GetPublicIpv4SubnetSize

`func (o *SpotMarketRequestCreateInputInstanceAttributes) GetPublicIpv4SubnetSize() int32`

GetPublicIpv4SubnetSize returns the PublicIpv4SubnetSize field if non-nil, zero value otherwise.

### GetPublicIpv4SubnetSizeOk

`func (o *SpotMarketRequestCreateInputInstanceAttributes) GetPublicIpv4SubnetSizeOk() (*int32, bool)`

GetPublicIpv4SubnetSizeOk returns a tuple with the PublicIpv4SubnetSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpv4SubnetSize

`func (o *SpotMarketRequestCreateInputInstanceAttributes) SetPublicIpv4SubnetSize(v int32)`

SetPublicIpv4SubnetSize sets PublicIpv4SubnetSize field to given value.

### HasPublicIpv4SubnetSize

`func (o *SpotMarketRequestCreateInputInstanceAttributes) HasPublicIpv4SubnetSize() bool`

HasPublicIpv4SubnetSize returns a boolean if a field has been set.

### GetTags

`func (o *SpotMarketRequestCreateInputInstanceAttributes) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SpotMarketRequestCreateInputInstanceAttributes) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SpotMarketRequestCreateInputInstanceAttributes) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SpotMarketRequestCreateInputInstanceAttributes) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTerminationTime

`func (o *SpotMarketRequestCreateInputInstanceAttributes) GetTerminationTime() time.Time`

GetTerminationTime returns the TerminationTime field if non-nil, zero value otherwise.

### GetTerminationTimeOk

`func (o *SpotMarketRequestCreateInputInstanceAttributes) GetTerminationTimeOk() (*time.Time, bool)`

GetTerminationTimeOk returns a tuple with the TerminationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationTime

`func (o *SpotMarketRequestCreateInputInstanceAttributes) SetTerminationTime(v time.Time)`

SetTerminationTime sets TerminationTime field to given value.

### HasTerminationTime

`func (o *SpotMarketRequestCreateInputInstanceAttributes) HasTerminationTime() bool`

HasTerminationTime returns a boolean if a field has been set.

### GetUserSshKeys

`func (o *SpotMarketRequestCreateInputInstanceAttributes) GetUserSshKeys() []string`

GetUserSshKeys returns the UserSshKeys field if non-nil, zero value otherwise.

### GetUserSshKeysOk

`func (o *SpotMarketRequestCreateInputInstanceAttributes) GetUserSshKeysOk() (*[]string, bool)`

GetUserSshKeysOk returns a tuple with the UserSshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSshKeys

`func (o *SpotMarketRequestCreateInputInstanceAttributes) SetUserSshKeys(v []string)`

SetUserSshKeys sets UserSshKeys field to given value.

### HasUserSshKeys

`func (o *SpotMarketRequestCreateInputInstanceAttributes) HasUserSshKeys() bool`

HasUserSshKeys returns a boolean if a field has been set.

### GetUserdata

`func (o *SpotMarketRequestCreateInputInstanceAttributes) GetUserdata() string`

GetUserdata returns the Userdata field if non-nil, zero value otherwise.

### GetUserdataOk

`func (o *SpotMarketRequestCreateInputInstanceAttributes) GetUserdataOk() (*string, bool)`

GetUserdataOk returns a tuple with the Userdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserdata

`func (o *SpotMarketRequestCreateInputInstanceAttributes) SetUserdata(v string)`

SetUserdata sets Userdata field to given value.

### HasUserdata

`func (o *SpotMarketRequestCreateInputInstanceAttributes) HasUserdata() bool`

HasUserdata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


