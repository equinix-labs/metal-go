# CreateDeviceBatchRequestBatchesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to **[]string** |  | [optional] 
**AlwaysPxe** | Pointer to **bool** |  | [optional] 
**BillingCycle** | Pointer to **string** |  | [optional] 
**Customdata** | Pointer to **map[string]interface{}** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Facility** | Pointer to **[]string** | Array of facility codes the batch can use for provisioning. This param also takes a string if you want the batch to be fulfilled in only one facility. Cannot be set if the metro is already set. | [optional] 
**Features** | Pointer to **[]string** |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**Hostnames** | Pointer to **[]string** |  | [optional] 
**IpAddresses** | Pointer to [**[]CreateDeviceBatchRequestBatchesInnerIpAddressesInner**](CreateDeviceBatchRequestBatchesInnerIpAddressesInner.md) |  | [optional] 
**Locked** | Pointer to **bool** |  | [optional] 
**Metro** | Pointer to **string** | The metro ID or code the batch can use for provisioning. Cannot be set if the facility is already set. | [optional] 
**NoSshKeys** | Pointer to **bool** |  | [optional] 
**OperatingSystem** | Pointer to **string** |  | [optional] 
**Plan** | Pointer to **string** |  | [optional] 
**ProjectSshKeys** | Pointer to **[]string** |  | [optional] 
**TerminationTime** | Pointer to **time.Time** |  | [optional] 
**UserSshKeys** | Pointer to **[]string** | The UUIDs of users whose SSH keys should be included on the provisioned device. | [optional] 
**Userdata** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateDeviceBatchRequestBatchesInner

`func NewCreateDeviceBatchRequestBatchesInner() *CreateDeviceBatchRequestBatchesInner`

NewCreateDeviceBatchRequestBatchesInner instantiates a new CreateDeviceBatchRequestBatchesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDeviceBatchRequestBatchesInnerWithDefaults

`func NewCreateDeviceBatchRequestBatchesInnerWithDefaults() *CreateDeviceBatchRequestBatchesInner`

NewCreateDeviceBatchRequestBatchesInnerWithDefaults instantiates a new CreateDeviceBatchRequestBatchesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *CreateDeviceBatchRequestBatchesInner) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateDeviceBatchRequestBatchesInner) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateDeviceBatchRequestBatchesInner) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateDeviceBatchRequestBatchesInner) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetAlwaysPxe

`func (o *CreateDeviceBatchRequestBatchesInner) GetAlwaysPxe() bool`

GetAlwaysPxe returns the AlwaysPxe field if non-nil, zero value otherwise.

### GetAlwaysPxeOk

`func (o *CreateDeviceBatchRequestBatchesInner) GetAlwaysPxeOk() (*bool, bool)`

GetAlwaysPxeOk returns a tuple with the AlwaysPxe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlwaysPxe

`func (o *CreateDeviceBatchRequestBatchesInner) SetAlwaysPxe(v bool)`

SetAlwaysPxe sets AlwaysPxe field to given value.

### HasAlwaysPxe

`func (o *CreateDeviceBatchRequestBatchesInner) HasAlwaysPxe() bool`

HasAlwaysPxe returns a boolean if a field has been set.

### GetBillingCycle

`func (o *CreateDeviceBatchRequestBatchesInner) GetBillingCycle() string`

GetBillingCycle returns the BillingCycle field if non-nil, zero value otherwise.

### GetBillingCycleOk

`func (o *CreateDeviceBatchRequestBatchesInner) GetBillingCycleOk() (*string, bool)`

GetBillingCycleOk returns a tuple with the BillingCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCycle

`func (o *CreateDeviceBatchRequestBatchesInner) SetBillingCycle(v string)`

SetBillingCycle sets BillingCycle field to given value.

### HasBillingCycle

`func (o *CreateDeviceBatchRequestBatchesInner) HasBillingCycle() bool`

HasBillingCycle returns a boolean if a field has been set.

### GetCustomdata

`func (o *CreateDeviceBatchRequestBatchesInner) GetCustomdata() map[string]interface{}`

GetCustomdata returns the Customdata field if non-nil, zero value otherwise.

### GetCustomdataOk

`func (o *CreateDeviceBatchRequestBatchesInner) GetCustomdataOk() (*map[string]interface{}, bool)`

GetCustomdataOk returns a tuple with the Customdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomdata

`func (o *CreateDeviceBatchRequestBatchesInner) SetCustomdata(v map[string]interface{})`

SetCustomdata sets Customdata field to given value.

### HasCustomdata

`func (o *CreateDeviceBatchRequestBatchesInner) HasCustomdata() bool`

HasCustomdata returns a boolean if a field has been set.

### GetDescription

`func (o *CreateDeviceBatchRequestBatchesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateDeviceBatchRequestBatchesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateDeviceBatchRequestBatchesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateDeviceBatchRequestBatchesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFacility

`func (o *CreateDeviceBatchRequestBatchesInner) GetFacility() []string`

GetFacility returns the Facility field if non-nil, zero value otherwise.

### GetFacilityOk

`func (o *CreateDeviceBatchRequestBatchesInner) GetFacilityOk() (*[]string, bool)`

GetFacilityOk returns a tuple with the Facility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacility

`func (o *CreateDeviceBatchRequestBatchesInner) SetFacility(v []string)`

SetFacility sets Facility field to given value.

### HasFacility

`func (o *CreateDeviceBatchRequestBatchesInner) HasFacility() bool`

HasFacility returns a boolean if a field has been set.

### GetFeatures

`func (o *CreateDeviceBatchRequestBatchesInner) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *CreateDeviceBatchRequestBatchesInner) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *CreateDeviceBatchRequestBatchesInner) SetFeatures(v []string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *CreateDeviceBatchRequestBatchesInner) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetHostname

`func (o *CreateDeviceBatchRequestBatchesInner) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *CreateDeviceBatchRequestBatchesInner) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *CreateDeviceBatchRequestBatchesInner) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *CreateDeviceBatchRequestBatchesInner) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetHostnames

`func (o *CreateDeviceBatchRequestBatchesInner) GetHostnames() []string`

GetHostnames returns the Hostnames field if non-nil, zero value otherwise.

### GetHostnamesOk

`func (o *CreateDeviceBatchRequestBatchesInner) GetHostnamesOk() (*[]string, bool)`

GetHostnamesOk returns a tuple with the Hostnames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnames

`func (o *CreateDeviceBatchRequestBatchesInner) SetHostnames(v []string)`

SetHostnames sets Hostnames field to given value.

### HasHostnames

`func (o *CreateDeviceBatchRequestBatchesInner) HasHostnames() bool`

HasHostnames returns a boolean if a field has been set.

### GetIpAddresses

`func (o *CreateDeviceBatchRequestBatchesInner) GetIpAddresses() []CreateDeviceBatchRequestBatchesInnerIpAddressesInner`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *CreateDeviceBatchRequestBatchesInner) GetIpAddressesOk() (*[]CreateDeviceBatchRequestBatchesInnerIpAddressesInner, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *CreateDeviceBatchRequestBatchesInner) SetIpAddresses(v []CreateDeviceBatchRequestBatchesInnerIpAddressesInner)`

SetIpAddresses sets IpAddresses field to given value.

### HasIpAddresses

`func (o *CreateDeviceBatchRequestBatchesInner) HasIpAddresses() bool`

HasIpAddresses returns a boolean if a field has been set.

### GetLocked

`func (o *CreateDeviceBatchRequestBatchesInner) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *CreateDeviceBatchRequestBatchesInner) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *CreateDeviceBatchRequestBatchesInner) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *CreateDeviceBatchRequestBatchesInner) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetMetro

`func (o *CreateDeviceBatchRequestBatchesInner) GetMetro() string`

GetMetro returns the Metro field if non-nil, zero value otherwise.

### GetMetroOk

`func (o *CreateDeviceBatchRequestBatchesInner) GetMetroOk() (*string, bool)`

GetMetroOk returns a tuple with the Metro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetro

`func (o *CreateDeviceBatchRequestBatchesInner) SetMetro(v string)`

SetMetro sets Metro field to given value.

### HasMetro

`func (o *CreateDeviceBatchRequestBatchesInner) HasMetro() bool`

HasMetro returns a boolean if a field has been set.

### GetNoSshKeys

`func (o *CreateDeviceBatchRequestBatchesInner) GetNoSshKeys() bool`

GetNoSshKeys returns the NoSshKeys field if non-nil, zero value otherwise.

### GetNoSshKeysOk

`func (o *CreateDeviceBatchRequestBatchesInner) GetNoSshKeysOk() (*bool, bool)`

GetNoSshKeysOk returns a tuple with the NoSshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoSshKeys

`func (o *CreateDeviceBatchRequestBatchesInner) SetNoSshKeys(v bool)`

SetNoSshKeys sets NoSshKeys field to given value.

### HasNoSshKeys

`func (o *CreateDeviceBatchRequestBatchesInner) HasNoSshKeys() bool`

HasNoSshKeys returns a boolean if a field has been set.

### GetOperatingSystem

`func (o *CreateDeviceBatchRequestBatchesInner) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *CreateDeviceBatchRequestBatchesInner) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *CreateDeviceBatchRequestBatchesInner) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *CreateDeviceBatchRequestBatchesInner) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### GetPlan

`func (o *CreateDeviceBatchRequestBatchesInner) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *CreateDeviceBatchRequestBatchesInner) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *CreateDeviceBatchRequestBatchesInner) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *CreateDeviceBatchRequestBatchesInner) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetProjectSshKeys

`func (o *CreateDeviceBatchRequestBatchesInner) GetProjectSshKeys() []string`

GetProjectSshKeys returns the ProjectSshKeys field if non-nil, zero value otherwise.

### GetProjectSshKeysOk

`func (o *CreateDeviceBatchRequestBatchesInner) GetProjectSshKeysOk() (*[]string, bool)`

GetProjectSshKeysOk returns a tuple with the ProjectSshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectSshKeys

`func (o *CreateDeviceBatchRequestBatchesInner) SetProjectSshKeys(v []string)`

SetProjectSshKeys sets ProjectSshKeys field to given value.

### HasProjectSshKeys

`func (o *CreateDeviceBatchRequestBatchesInner) HasProjectSshKeys() bool`

HasProjectSshKeys returns a boolean if a field has been set.

### GetTerminationTime

`func (o *CreateDeviceBatchRequestBatchesInner) GetTerminationTime() time.Time`

GetTerminationTime returns the TerminationTime field if non-nil, zero value otherwise.

### GetTerminationTimeOk

`func (o *CreateDeviceBatchRequestBatchesInner) GetTerminationTimeOk() (*time.Time, bool)`

GetTerminationTimeOk returns a tuple with the TerminationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationTime

`func (o *CreateDeviceBatchRequestBatchesInner) SetTerminationTime(v time.Time)`

SetTerminationTime sets TerminationTime field to given value.

### HasTerminationTime

`func (o *CreateDeviceBatchRequestBatchesInner) HasTerminationTime() bool`

HasTerminationTime returns a boolean if a field has been set.

### GetUserSshKeys

`func (o *CreateDeviceBatchRequestBatchesInner) GetUserSshKeys() []string`

GetUserSshKeys returns the UserSshKeys field if non-nil, zero value otherwise.

### GetUserSshKeysOk

`func (o *CreateDeviceBatchRequestBatchesInner) GetUserSshKeysOk() (*[]string, bool)`

GetUserSshKeysOk returns a tuple with the UserSshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSshKeys

`func (o *CreateDeviceBatchRequestBatchesInner) SetUserSshKeys(v []string)`

SetUserSshKeys sets UserSshKeys field to given value.

### HasUserSshKeys

`func (o *CreateDeviceBatchRequestBatchesInner) HasUserSshKeys() bool`

HasUserSshKeys returns a boolean if a field has been set.

### GetUserdata

`func (o *CreateDeviceBatchRequestBatchesInner) GetUserdata() string`

GetUserdata returns the Userdata field if non-nil, zero value otherwise.

### GetUserdataOk

`func (o *CreateDeviceBatchRequestBatchesInner) GetUserdataOk() (*string, bool)`

GetUserdataOk returns a tuple with the Userdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserdata

`func (o *CreateDeviceBatchRequestBatchesInner) SetUserdata(v string)`

SetUserdata sets Userdata field to given value.

### HasUserdata

`func (o *CreateDeviceBatchRequestBatchesInner) HasUserdata() bool`

HasUserdata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


