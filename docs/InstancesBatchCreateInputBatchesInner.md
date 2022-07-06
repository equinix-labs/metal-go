# InstancesBatchCreateInputBatchesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Plan** | Pointer to **string** |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**Hostnames** | Pointer to **[]string** |  | [optional] 
**Facility** | Pointer to **[]string** | Array of facility codes the batch can use for provisioning. This param also takes a string if you want the batch to be fulfilled in only one facility. Cannot be set if the metro is already set. | [optional] 
**Metro** | Pointer to **string** | The metro ID or code the batch can use for provisioning. Cannot be set if the facility is already set. | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**BillingCycle** | Pointer to **string** |  | [optional] 
**OperatingSystem** | Pointer to **string** |  | [optional] 
**AlwaysPxe** | Pointer to **bool** |  | [optional] 
**Userdata** | Pointer to **string** |  | [optional] 
**Locked** | Pointer to **bool** |  | [optional] 
**TerminationTime** | Pointer to **time.Time** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**ProjectSshKeys** | Pointer to **[]string** |  | [optional] 
**UserSshKeys** | Pointer to **[]string** | The UUIDs of users whose SSH keys should be included on the provisioned device. | [optional] 
**NoSshKeys** | Pointer to **bool** |  | [optional] 
**Features** | Pointer to **[]string** |  | [optional] 
**Customdata** | Pointer to **map[string]interface{}** |  | [optional] 
**IpAddresses** | Pointer to [**[]InstancesBatchCreateInputBatchesInnerIpAddressesInner**](InstancesBatchCreateInputBatchesInnerIpAddressesInner.md) |  | [optional] 

## Methods

### NewInstancesBatchCreateInputBatchesInner

`func NewInstancesBatchCreateInputBatchesInner() *InstancesBatchCreateInputBatchesInner`

NewInstancesBatchCreateInputBatchesInner instantiates a new InstancesBatchCreateInputBatchesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstancesBatchCreateInputBatchesInnerWithDefaults

`func NewInstancesBatchCreateInputBatchesInnerWithDefaults() *InstancesBatchCreateInputBatchesInner`

NewInstancesBatchCreateInputBatchesInnerWithDefaults instantiates a new InstancesBatchCreateInputBatchesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlan

`func (o *InstancesBatchCreateInputBatchesInner) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *InstancesBatchCreateInputBatchesInner) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *InstancesBatchCreateInputBatchesInner) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *InstancesBatchCreateInputBatchesInner) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetHostname

`func (o *InstancesBatchCreateInputBatchesInner) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *InstancesBatchCreateInputBatchesInner) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *InstancesBatchCreateInputBatchesInner) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *InstancesBatchCreateInputBatchesInner) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetHostnames

`func (o *InstancesBatchCreateInputBatchesInner) GetHostnames() []string`

GetHostnames returns the Hostnames field if non-nil, zero value otherwise.

### GetHostnamesOk

`func (o *InstancesBatchCreateInputBatchesInner) GetHostnamesOk() (*[]string, bool)`

GetHostnamesOk returns a tuple with the Hostnames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnames

`func (o *InstancesBatchCreateInputBatchesInner) SetHostnames(v []string)`

SetHostnames sets Hostnames field to given value.

### HasHostnames

`func (o *InstancesBatchCreateInputBatchesInner) HasHostnames() bool`

HasHostnames returns a boolean if a field has been set.

### GetFacility

`func (o *InstancesBatchCreateInputBatchesInner) GetFacility() []string`

GetFacility returns the Facility field if non-nil, zero value otherwise.

### GetFacilityOk

`func (o *InstancesBatchCreateInputBatchesInner) GetFacilityOk() (*[]string, bool)`

GetFacilityOk returns a tuple with the Facility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacility

`func (o *InstancesBatchCreateInputBatchesInner) SetFacility(v []string)`

SetFacility sets Facility field to given value.

### HasFacility

`func (o *InstancesBatchCreateInputBatchesInner) HasFacility() bool`

HasFacility returns a boolean if a field has been set.

### GetMetro

`func (o *InstancesBatchCreateInputBatchesInner) GetMetro() string`

GetMetro returns the Metro field if non-nil, zero value otherwise.

### GetMetroOk

`func (o *InstancesBatchCreateInputBatchesInner) GetMetroOk() (*string, bool)`

GetMetroOk returns a tuple with the Metro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetro

`func (o *InstancesBatchCreateInputBatchesInner) SetMetro(v string)`

SetMetro sets Metro field to given value.

### HasMetro

`func (o *InstancesBatchCreateInputBatchesInner) HasMetro() bool`

HasMetro returns a boolean if a field has been set.

### GetDescription

`func (o *InstancesBatchCreateInputBatchesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InstancesBatchCreateInputBatchesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InstancesBatchCreateInputBatchesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InstancesBatchCreateInputBatchesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetBillingCycle

`func (o *InstancesBatchCreateInputBatchesInner) GetBillingCycle() string`

GetBillingCycle returns the BillingCycle field if non-nil, zero value otherwise.

### GetBillingCycleOk

`func (o *InstancesBatchCreateInputBatchesInner) GetBillingCycleOk() (*string, bool)`

GetBillingCycleOk returns a tuple with the BillingCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCycle

`func (o *InstancesBatchCreateInputBatchesInner) SetBillingCycle(v string)`

SetBillingCycle sets BillingCycle field to given value.

### HasBillingCycle

`func (o *InstancesBatchCreateInputBatchesInner) HasBillingCycle() bool`

HasBillingCycle returns a boolean if a field has been set.

### GetOperatingSystem

`func (o *InstancesBatchCreateInputBatchesInner) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *InstancesBatchCreateInputBatchesInner) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *InstancesBatchCreateInputBatchesInner) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *InstancesBatchCreateInputBatchesInner) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### GetAlwaysPxe

`func (o *InstancesBatchCreateInputBatchesInner) GetAlwaysPxe() bool`

GetAlwaysPxe returns the AlwaysPxe field if non-nil, zero value otherwise.

### GetAlwaysPxeOk

`func (o *InstancesBatchCreateInputBatchesInner) GetAlwaysPxeOk() (*bool, bool)`

GetAlwaysPxeOk returns a tuple with the AlwaysPxe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlwaysPxe

`func (o *InstancesBatchCreateInputBatchesInner) SetAlwaysPxe(v bool)`

SetAlwaysPxe sets AlwaysPxe field to given value.

### HasAlwaysPxe

`func (o *InstancesBatchCreateInputBatchesInner) HasAlwaysPxe() bool`

HasAlwaysPxe returns a boolean if a field has been set.

### GetUserdata

`func (o *InstancesBatchCreateInputBatchesInner) GetUserdata() string`

GetUserdata returns the Userdata field if non-nil, zero value otherwise.

### GetUserdataOk

`func (o *InstancesBatchCreateInputBatchesInner) GetUserdataOk() (*string, bool)`

GetUserdataOk returns a tuple with the Userdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserdata

`func (o *InstancesBatchCreateInputBatchesInner) SetUserdata(v string)`

SetUserdata sets Userdata field to given value.

### HasUserdata

`func (o *InstancesBatchCreateInputBatchesInner) HasUserdata() bool`

HasUserdata returns a boolean if a field has been set.

### GetLocked

`func (o *InstancesBatchCreateInputBatchesInner) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *InstancesBatchCreateInputBatchesInner) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *InstancesBatchCreateInputBatchesInner) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *InstancesBatchCreateInputBatchesInner) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetTerminationTime

`func (o *InstancesBatchCreateInputBatchesInner) GetTerminationTime() time.Time`

GetTerminationTime returns the TerminationTime field if non-nil, zero value otherwise.

### GetTerminationTimeOk

`func (o *InstancesBatchCreateInputBatchesInner) GetTerminationTimeOk() (*time.Time, bool)`

GetTerminationTimeOk returns a tuple with the TerminationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationTime

`func (o *InstancesBatchCreateInputBatchesInner) SetTerminationTime(v time.Time)`

SetTerminationTime sets TerminationTime field to given value.

### HasTerminationTime

`func (o *InstancesBatchCreateInputBatchesInner) HasTerminationTime() bool`

HasTerminationTime returns a boolean if a field has been set.

### GetTags

`func (o *InstancesBatchCreateInputBatchesInner) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InstancesBatchCreateInputBatchesInner) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InstancesBatchCreateInputBatchesInner) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InstancesBatchCreateInputBatchesInner) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetProjectSshKeys

`func (o *InstancesBatchCreateInputBatchesInner) GetProjectSshKeys() []string`

GetProjectSshKeys returns the ProjectSshKeys field if non-nil, zero value otherwise.

### GetProjectSshKeysOk

`func (o *InstancesBatchCreateInputBatchesInner) GetProjectSshKeysOk() (*[]string, bool)`

GetProjectSshKeysOk returns a tuple with the ProjectSshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectSshKeys

`func (o *InstancesBatchCreateInputBatchesInner) SetProjectSshKeys(v []string)`

SetProjectSshKeys sets ProjectSshKeys field to given value.

### HasProjectSshKeys

`func (o *InstancesBatchCreateInputBatchesInner) HasProjectSshKeys() bool`

HasProjectSshKeys returns a boolean if a field has been set.

### GetUserSshKeys

`func (o *InstancesBatchCreateInputBatchesInner) GetUserSshKeys() []string`

GetUserSshKeys returns the UserSshKeys field if non-nil, zero value otherwise.

### GetUserSshKeysOk

`func (o *InstancesBatchCreateInputBatchesInner) GetUserSshKeysOk() (*[]string, bool)`

GetUserSshKeysOk returns a tuple with the UserSshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSshKeys

`func (o *InstancesBatchCreateInputBatchesInner) SetUserSshKeys(v []string)`

SetUserSshKeys sets UserSshKeys field to given value.

### HasUserSshKeys

`func (o *InstancesBatchCreateInputBatchesInner) HasUserSshKeys() bool`

HasUserSshKeys returns a boolean if a field has been set.

### GetNoSshKeys

`func (o *InstancesBatchCreateInputBatchesInner) GetNoSshKeys() bool`

GetNoSshKeys returns the NoSshKeys field if non-nil, zero value otherwise.

### GetNoSshKeysOk

`func (o *InstancesBatchCreateInputBatchesInner) GetNoSshKeysOk() (*bool, bool)`

GetNoSshKeysOk returns a tuple with the NoSshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoSshKeys

`func (o *InstancesBatchCreateInputBatchesInner) SetNoSshKeys(v bool)`

SetNoSshKeys sets NoSshKeys field to given value.

### HasNoSshKeys

`func (o *InstancesBatchCreateInputBatchesInner) HasNoSshKeys() bool`

HasNoSshKeys returns a boolean if a field has been set.

### GetFeatures

`func (o *InstancesBatchCreateInputBatchesInner) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *InstancesBatchCreateInputBatchesInner) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *InstancesBatchCreateInputBatchesInner) SetFeatures(v []string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *InstancesBatchCreateInputBatchesInner) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetCustomdata

`func (o *InstancesBatchCreateInputBatchesInner) GetCustomdata() map[string]interface{}`

GetCustomdata returns the Customdata field if non-nil, zero value otherwise.

### GetCustomdataOk

`func (o *InstancesBatchCreateInputBatchesInner) GetCustomdataOk() (*map[string]interface{}, bool)`

GetCustomdataOk returns a tuple with the Customdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomdata

`func (o *InstancesBatchCreateInputBatchesInner) SetCustomdata(v map[string]interface{})`

SetCustomdata sets Customdata field to given value.

### HasCustomdata

`func (o *InstancesBatchCreateInputBatchesInner) HasCustomdata() bool`

HasCustomdata returns a boolean if a field has been set.

### GetIpAddresses

`func (o *InstancesBatchCreateInputBatchesInner) GetIpAddresses() []InstancesBatchCreateInputBatchesInnerIpAddressesInner`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *InstancesBatchCreateInputBatchesInner) GetIpAddressesOk() (*[]InstancesBatchCreateInputBatchesInnerIpAddressesInner, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *InstancesBatchCreateInputBatchesInner) SetIpAddresses(v []InstancesBatchCreateInputBatchesInnerIpAddressesInner)`

SetIpAddresses sets IpAddresses field to given value.

### HasIpAddresses

`func (o *InstancesBatchCreateInputBatchesInner) HasIpAddresses() bool`

HasIpAddresses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


