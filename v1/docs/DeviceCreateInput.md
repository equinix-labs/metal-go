# DeviceCreateInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Facility** | **string** |  | 
**Plan** | **string** |  | 
**Hostname** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**BillingCycle** | Pointer to **string** |  | [optional] 
**OperatingSystem** | **string** |  | 
**AlwaysPxe** | Pointer to **bool** |  | [optional] 
**IpxeScriptUrl** | Pointer to **string** |  | [optional] 
**Userdata** | Pointer to **string** |  | [optional] 
**Locked** | Pointer to **bool** |  | [optional] 
**Customdata** | Pointer to **map[string]interface{}** |  | [optional] 
**Metro** | Pointer to **string** | Metro code or ID of where the instance should be provisioned in. | [optional] 
**HardwareReservationId** | Pointer to **string** |  | [optional] 
**SpotInstance** | Pointer to **bool** |  | [optional] 
**SpotPriceMax** | Pointer to **float32** |  | [optional] 
**TerminationTime** | Pointer to **time.Time** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**ProjectSshKeys** | Pointer to **[]string** |  | [optional] 
**UserSshKeys** | Pointer to **[]string** | The UUIDs of users whose SSH keys should be included on the provisioned device. | [optional] 
**Features** | Pointer to **[]string** |  | [optional] 
**PublicIpv4SubnetSize** | Pointer to **float32** | Deprecated. Use ip_addresses. Subnet range for addresses allocated to this device. Your project must have addresses available for a non-default request. | [optional] 
**PrivateIpv4SubnetSize** | Pointer to **float32** | Deprecated. Use ip_addresses. Subnet range for addresses allocated to this device. | [optional] 
**IpAddresses** | Pointer to [**[]DeviceCreateInputIpAddresses**](DeviceCreateInputIpAddresses.md) |  | [optional] 

## Methods

### NewDeviceCreateInput

`func NewDeviceCreateInput(facility string, plan string, operatingSystem string, ) *DeviceCreateInput`

NewDeviceCreateInput instantiates a new DeviceCreateInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceCreateInputWithDefaults

`func NewDeviceCreateInputWithDefaults() *DeviceCreateInput`

NewDeviceCreateInputWithDefaults instantiates a new DeviceCreateInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFacility

`func (o *DeviceCreateInput) GetFacility() string`

GetFacility returns the Facility field if non-nil, zero value otherwise.

### GetFacilityOk

`func (o *DeviceCreateInput) GetFacilityOk() (*string, bool)`

GetFacilityOk returns a tuple with the Facility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacility

`func (o *DeviceCreateInput) SetFacility(v string)`

SetFacility sets Facility field to given value.


### GetPlan

`func (o *DeviceCreateInput) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *DeviceCreateInput) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *DeviceCreateInput) SetPlan(v string)`

SetPlan sets Plan field to given value.


### GetHostname

`func (o *DeviceCreateInput) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *DeviceCreateInput) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *DeviceCreateInput) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *DeviceCreateInput) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetDescription

`func (o *DeviceCreateInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DeviceCreateInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DeviceCreateInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DeviceCreateInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetBillingCycle

`func (o *DeviceCreateInput) GetBillingCycle() string`

GetBillingCycle returns the BillingCycle field if non-nil, zero value otherwise.

### GetBillingCycleOk

`func (o *DeviceCreateInput) GetBillingCycleOk() (*string, bool)`

GetBillingCycleOk returns a tuple with the BillingCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCycle

`func (o *DeviceCreateInput) SetBillingCycle(v string)`

SetBillingCycle sets BillingCycle field to given value.

### HasBillingCycle

`func (o *DeviceCreateInput) HasBillingCycle() bool`

HasBillingCycle returns a boolean if a field has been set.

### GetOperatingSystem

`func (o *DeviceCreateInput) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *DeviceCreateInput) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *DeviceCreateInput) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.


### GetAlwaysPxe

`func (o *DeviceCreateInput) GetAlwaysPxe() bool`

GetAlwaysPxe returns the AlwaysPxe field if non-nil, zero value otherwise.

### GetAlwaysPxeOk

`func (o *DeviceCreateInput) GetAlwaysPxeOk() (*bool, bool)`

GetAlwaysPxeOk returns a tuple with the AlwaysPxe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlwaysPxe

`func (o *DeviceCreateInput) SetAlwaysPxe(v bool)`

SetAlwaysPxe sets AlwaysPxe field to given value.

### HasAlwaysPxe

`func (o *DeviceCreateInput) HasAlwaysPxe() bool`

HasAlwaysPxe returns a boolean if a field has been set.

### GetIpxeScriptUrl

`func (o *DeviceCreateInput) GetIpxeScriptUrl() string`

GetIpxeScriptUrl returns the IpxeScriptUrl field if non-nil, zero value otherwise.

### GetIpxeScriptUrlOk

`func (o *DeviceCreateInput) GetIpxeScriptUrlOk() (*string, bool)`

GetIpxeScriptUrlOk returns a tuple with the IpxeScriptUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpxeScriptUrl

`func (o *DeviceCreateInput) SetIpxeScriptUrl(v string)`

SetIpxeScriptUrl sets IpxeScriptUrl field to given value.

### HasIpxeScriptUrl

`func (o *DeviceCreateInput) HasIpxeScriptUrl() bool`

HasIpxeScriptUrl returns a boolean if a field has been set.

### GetUserdata

`func (o *DeviceCreateInput) GetUserdata() string`

GetUserdata returns the Userdata field if non-nil, zero value otherwise.

### GetUserdataOk

`func (o *DeviceCreateInput) GetUserdataOk() (*string, bool)`

GetUserdataOk returns a tuple with the Userdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserdata

`func (o *DeviceCreateInput) SetUserdata(v string)`

SetUserdata sets Userdata field to given value.

### HasUserdata

`func (o *DeviceCreateInput) HasUserdata() bool`

HasUserdata returns a boolean if a field has been set.

### GetLocked

`func (o *DeviceCreateInput) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *DeviceCreateInput) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *DeviceCreateInput) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *DeviceCreateInput) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetCustomdata

`func (o *DeviceCreateInput) GetCustomdata() map[string]interface{}`

GetCustomdata returns the Customdata field if non-nil, zero value otherwise.

### GetCustomdataOk

`func (o *DeviceCreateInput) GetCustomdataOk() (*map[string]interface{}, bool)`

GetCustomdataOk returns a tuple with the Customdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomdata

`func (o *DeviceCreateInput) SetCustomdata(v map[string]interface{})`

SetCustomdata sets Customdata field to given value.

### HasCustomdata

`func (o *DeviceCreateInput) HasCustomdata() bool`

HasCustomdata returns a boolean if a field has been set.

### GetMetro

`func (o *DeviceCreateInput) GetMetro() string`

GetMetro returns the Metro field if non-nil, zero value otherwise.

### GetMetroOk

`func (o *DeviceCreateInput) GetMetroOk() (*string, bool)`

GetMetroOk returns a tuple with the Metro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetro

`func (o *DeviceCreateInput) SetMetro(v string)`

SetMetro sets Metro field to given value.

### HasMetro

`func (o *DeviceCreateInput) HasMetro() bool`

HasMetro returns a boolean if a field has been set.

### GetHardwareReservationId

`func (o *DeviceCreateInput) GetHardwareReservationId() string`

GetHardwareReservationId returns the HardwareReservationId field if non-nil, zero value otherwise.

### GetHardwareReservationIdOk

`func (o *DeviceCreateInput) GetHardwareReservationIdOk() (*string, bool)`

GetHardwareReservationIdOk returns a tuple with the HardwareReservationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareReservationId

`func (o *DeviceCreateInput) SetHardwareReservationId(v string)`

SetHardwareReservationId sets HardwareReservationId field to given value.

### HasHardwareReservationId

`func (o *DeviceCreateInput) HasHardwareReservationId() bool`

HasHardwareReservationId returns a boolean if a field has been set.

### GetSpotInstance

`func (o *DeviceCreateInput) GetSpotInstance() bool`

GetSpotInstance returns the SpotInstance field if non-nil, zero value otherwise.

### GetSpotInstanceOk

`func (o *DeviceCreateInput) GetSpotInstanceOk() (*bool, bool)`

GetSpotInstanceOk returns a tuple with the SpotInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotInstance

`func (o *DeviceCreateInput) SetSpotInstance(v bool)`

SetSpotInstance sets SpotInstance field to given value.

### HasSpotInstance

`func (o *DeviceCreateInput) HasSpotInstance() bool`

HasSpotInstance returns a boolean if a field has been set.

### GetSpotPriceMax

`func (o *DeviceCreateInput) GetSpotPriceMax() float32`

GetSpotPriceMax returns the SpotPriceMax field if non-nil, zero value otherwise.

### GetSpotPriceMaxOk

`func (o *DeviceCreateInput) GetSpotPriceMaxOk() (*float32, bool)`

GetSpotPriceMaxOk returns a tuple with the SpotPriceMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotPriceMax

`func (o *DeviceCreateInput) SetSpotPriceMax(v float32)`

SetSpotPriceMax sets SpotPriceMax field to given value.

### HasSpotPriceMax

`func (o *DeviceCreateInput) HasSpotPriceMax() bool`

HasSpotPriceMax returns a boolean if a field has been set.

### GetTerminationTime

`func (o *DeviceCreateInput) GetTerminationTime() time.Time`

GetTerminationTime returns the TerminationTime field if non-nil, zero value otherwise.

### GetTerminationTimeOk

`func (o *DeviceCreateInput) GetTerminationTimeOk() (*time.Time, bool)`

GetTerminationTimeOk returns a tuple with the TerminationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationTime

`func (o *DeviceCreateInput) SetTerminationTime(v time.Time)`

SetTerminationTime sets TerminationTime field to given value.

### HasTerminationTime

`func (o *DeviceCreateInput) HasTerminationTime() bool`

HasTerminationTime returns a boolean if a field has been set.

### GetTags

`func (o *DeviceCreateInput) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DeviceCreateInput) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DeviceCreateInput) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DeviceCreateInput) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetProjectSshKeys

`func (o *DeviceCreateInput) GetProjectSshKeys() []string`

GetProjectSshKeys returns the ProjectSshKeys field if non-nil, zero value otherwise.

### GetProjectSshKeysOk

`func (o *DeviceCreateInput) GetProjectSshKeysOk() (*[]string, bool)`

GetProjectSshKeysOk returns a tuple with the ProjectSshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectSshKeys

`func (o *DeviceCreateInput) SetProjectSshKeys(v []string)`

SetProjectSshKeys sets ProjectSshKeys field to given value.

### HasProjectSshKeys

`func (o *DeviceCreateInput) HasProjectSshKeys() bool`

HasProjectSshKeys returns a boolean if a field has been set.

### GetUserSshKeys

`func (o *DeviceCreateInput) GetUserSshKeys() []string`

GetUserSshKeys returns the UserSshKeys field if non-nil, zero value otherwise.

### GetUserSshKeysOk

`func (o *DeviceCreateInput) GetUserSshKeysOk() (*[]string, bool)`

GetUserSshKeysOk returns a tuple with the UserSshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSshKeys

`func (o *DeviceCreateInput) SetUserSshKeys(v []string)`

SetUserSshKeys sets UserSshKeys field to given value.

### HasUserSshKeys

`func (o *DeviceCreateInput) HasUserSshKeys() bool`

HasUserSshKeys returns a boolean if a field has been set.

### GetFeatures

`func (o *DeviceCreateInput) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *DeviceCreateInput) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *DeviceCreateInput) SetFeatures(v []string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *DeviceCreateInput) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetPublicIpv4SubnetSize

`func (o *DeviceCreateInput) GetPublicIpv4SubnetSize() float32`

GetPublicIpv4SubnetSize returns the PublicIpv4SubnetSize field if non-nil, zero value otherwise.

### GetPublicIpv4SubnetSizeOk

`func (o *DeviceCreateInput) GetPublicIpv4SubnetSizeOk() (*float32, bool)`

GetPublicIpv4SubnetSizeOk returns a tuple with the PublicIpv4SubnetSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpv4SubnetSize

`func (o *DeviceCreateInput) SetPublicIpv4SubnetSize(v float32)`

SetPublicIpv4SubnetSize sets PublicIpv4SubnetSize field to given value.

### HasPublicIpv4SubnetSize

`func (o *DeviceCreateInput) HasPublicIpv4SubnetSize() bool`

HasPublicIpv4SubnetSize returns a boolean if a field has been set.

### GetPrivateIpv4SubnetSize

`func (o *DeviceCreateInput) GetPrivateIpv4SubnetSize() float32`

GetPrivateIpv4SubnetSize returns the PrivateIpv4SubnetSize field if non-nil, zero value otherwise.

### GetPrivateIpv4SubnetSizeOk

`func (o *DeviceCreateInput) GetPrivateIpv4SubnetSizeOk() (*float32, bool)`

GetPrivateIpv4SubnetSizeOk returns a tuple with the PrivateIpv4SubnetSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIpv4SubnetSize

`func (o *DeviceCreateInput) SetPrivateIpv4SubnetSize(v float32)`

SetPrivateIpv4SubnetSize sets PrivateIpv4SubnetSize field to given value.

### HasPrivateIpv4SubnetSize

`func (o *DeviceCreateInput) HasPrivateIpv4SubnetSize() bool`

HasPrivateIpv4SubnetSize returns a boolean if a field has been set.

### GetIpAddresses

`func (o *DeviceCreateInput) GetIpAddresses() []DeviceCreateInputIpAddresses`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *DeviceCreateInput) GetIpAddressesOk() (*[]DeviceCreateInputIpAddresses, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *DeviceCreateInput) SetIpAddresses(v []DeviceCreateInputIpAddresses)`

SetIpAddresses sets IpAddresses field to given value.

### HasIpAddresses

`func (o *DeviceCreateInput) HasIpAddresses() bool`

HasIpAddresses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


