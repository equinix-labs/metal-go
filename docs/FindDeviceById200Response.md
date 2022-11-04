# FindDeviceById200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlwaysPxe** | Pointer to **bool** |  | [optional] 
**BillingCycle** | Pointer to **string** |  | [optional] 
**BondingMode** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedBy** | Pointer to [**FindDeviceById200ResponseCreatedBy**](FindDeviceById200ResponseCreatedBy.md) |  | [optional] 
**Customdata** | Pointer to **map[string]interface{}** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Facility** | Pointer to [**FindDeviceById200ResponseFacility**](FindDeviceById200ResponseFacility.md) |  | [optional] 
**HardwareReservation** | Pointer to [**FindBatchById200ResponseDevicesInner**](FindBatchById200ResponseDevicesInner.md) |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**ImageUrl** | Pointer to **string** |  | [optional] 
**IpAddresses** | Pointer to [**[]FindDeviceById200ResponseIpAddressesInner**](FindDeviceById200ResponseIpAddressesInner.md) |  | [optional] 
**IpxeScriptUrl** | Pointer to **string** |  | [optional] 
**Iqn** | Pointer to **string** |  | [optional] 
**Locked** | Pointer to **bool** |  | [optional] 
**Metro** | Pointer to [**FindDeviceById200ResponseFacilityMetro**](FindDeviceById200ResponseFacilityMetro.md) |  | [optional] 
**NetworkPorts** | Pointer to [**[]FindDeviceById200ResponseNetworkPortsInner**](FindDeviceById200ResponseNetworkPortsInner.md) | By default, servers at Equinix Metal are configured in a “bonded” mode using LACP (Link Aggregation Control Protocol). Each 2-NIC server is configured with a single bond (namely bond0) with both interfaces eth0 and eth1 as members of the bond in a default Layer 3 mode. Some device plans may have a different number of ports and bonds available. | [optional] 
**OperatingSystem** | Pointer to [**FindDeviceById200ResponseOperatingSystem**](FindDeviceById200ResponseOperatingSystem.md) |  | [optional] 
**Actions** | Pointer to [**[]FindDeviceById200ResponseActionsInner**](FindDeviceById200ResponseActionsInner.md) | Actions supported by the device instance. | [optional] 
**Plan** | Pointer to [**FindDeviceById200ResponsePlan**](FindDeviceById200ResponsePlan.md) |  | [optional] 
**Project** | Pointer to [**FindDeviceById200ResponseProject**](FindDeviceById200ResponseProject.md) |  | [optional] 
**ProjectLite** | Pointer to [**FindDeviceById200ResponseProjectLite**](FindDeviceById200ResponseProjectLite.md) |  | [optional] 
**ProvisioningEvents** | Pointer to [**[]FindInterconnectionEvents200Response**](FindInterconnectionEvents200Response.md) |  | [optional] 
**ProvisioningPercentage** | Pointer to **float32** | Only visible while device provisioning | [optional] 
**RootPassword** | Pointer to **string** | Root password is automatically generated when server is provisioned and it is removed after 24 hours | [optional] 
**ShortId** | Pointer to **string** |  | [optional] 
**SpotInstance** | Pointer to **bool** | Whether or not the device is a spot instance. | [optional] 
**SpotPriceMax** | Pointer to **float32** | The maximum price per hour you are willing to pay to keep this spot instance.  If you are outbid, the termination will be set allowing two minutes before shutdown. | [optional] 
**SshKeys** | Pointer to [**[]FindBatchById200ResponseDevicesInner**](FindBatchById200ResponseDevicesInner.md) |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**SwitchUuid** | Pointer to **string** | Switch short id. This can be used to determine if two devices are connected to the same switch, for example. | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**TerminationTime** | Pointer to **time.Time** | When the device will be terminated. This is commonly set in advance for ephemeral spot market instances but this field may also be set with on-demand and reservation instances to automatically delete the resource at a given time. The termination time can also be used to release a hardware reservation instance at a given time, keeping the reservation open for other uses.  On a spot market device, the termination time will be set automatically when outbid. | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**User** | Pointer to **string** |  | [optional] 
**Userdata** | Pointer to **string** |  | [optional] 
**Volumes** | Pointer to [**[]FindBatchById200ResponseDevicesInner**](FindBatchById200ResponseDevicesInner.md) |  | [optional] 

## Methods

### NewFindDeviceById200Response

`func NewFindDeviceById200Response() *FindDeviceById200Response`

NewFindDeviceById200Response instantiates a new FindDeviceById200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindDeviceById200ResponseWithDefaults

`func NewFindDeviceById200ResponseWithDefaults() *FindDeviceById200Response`

NewFindDeviceById200ResponseWithDefaults instantiates a new FindDeviceById200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlwaysPxe

`func (o *FindDeviceById200Response) GetAlwaysPxe() bool`

GetAlwaysPxe returns the AlwaysPxe field if non-nil, zero value otherwise.

### GetAlwaysPxeOk

`func (o *FindDeviceById200Response) GetAlwaysPxeOk() (*bool, bool)`

GetAlwaysPxeOk returns a tuple with the AlwaysPxe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlwaysPxe

`func (o *FindDeviceById200Response) SetAlwaysPxe(v bool)`

SetAlwaysPxe sets AlwaysPxe field to given value.

### HasAlwaysPxe

`func (o *FindDeviceById200Response) HasAlwaysPxe() bool`

HasAlwaysPxe returns a boolean if a field has been set.

### GetBillingCycle

`func (o *FindDeviceById200Response) GetBillingCycle() string`

GetBillingCycle returns the BillingCycle field if non-nil, zero value otherwise.

### GetBillingCycleOk

`func (o *FindDeviceById200Response) GetBillingCycleOk() (*string, bool)`

GetBillingCycleOk returns a tuple with the BillingCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCycle

`func (o *FindDeviceById200Response) SetBillingCycle(v string)`

SetBillingCycle sets BillingCycle field to given value.

### HasBillingCycle

`func (o *FindDeviceById200Response) HasBillingCycle() bool`

HasBillingCycle returns a boolean if a field has been set.

### GetBondingMode

`func (o *FindDeviceById200Response) GetBondingMode() int32`

GetBondingMode returns the BondingMode field if non-nil, zero value otherwise.

### GetBondingModeOk

`func (o *FindDeviceById200Response) GetBondingModeOk() (*int32, bool)`

GetBondingModeOk returns a tuple with the BondingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBondingMode

`func (o *FindDeviceById200Response) SetBondingMode(v int32)`

SetBondingMode sets BondingMode field to given value.

### HasBondingMode

`func (o *FindDeviceById200Response) HasBondingMode() bool`

HasBondingMode returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FindDeviceById200Response) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FindDeviceById200Response) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FindDeviceById200Response) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FindDeviceById200Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *FindDeviceById200Response) GetCreatedBy() FindDeviceById200ResponseCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *FindDeviceById200Response) GetCreatedByOk() (*FindDeviceById200ResponseCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *FindDeviceById200Response) SetCreatedBy(v FindDeviceById200ResponseCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *FindDeviceById200Response) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCustomdata

`func (o *FindDeviceById200Response) GetCustomdata() map[string]interface{}`

GetCustomdata returns the Customdata field if non-nil, zero value otherwise.

### GetCustomdataOk

`func (o *FindDeviceById200Response) GetCustomdataOk() (*map[string]interface{}, bool)`

GetCustomdataOk returns a tuple with the Customdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomdata

`func (o *FindDeviceById200Response) SetCustomdata(v map[string]interface{})`

SetCustomdata sets Customdata field to given value.

### HasCustomdata

`func (o *FindDeviceById200Response) HasCustomdata() bool`

HasCustomdata returns a boolean if a field has been set.

### GetDescription

`func (o *FindDeviceById200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FindDeviceById200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FindDeviceById200Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FindDeviceById200Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFacility

`func (o *FindDeviceById200Response) GetFacility() FindDeviceById200ResponseFacility`

GetFacility returns the Facility field if non-nil, zero value otherwise.

### GetFacilityOk

`func (o *FindDeviceById200Response) GetFacilityOk() (*FindDeviceById200ResponseFacility, bool)`

GetFacilityOk returns a tuple with the Facility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacility

`func (o *FindDeviceById200Response) SetFacility(v FindDeviceById200ResponseFacility)`

SetFacility sets Facility field to given value.

### HasFacility

`func (o *FindDeviceById200Response) HasFacility() bool`

HasFacility returns a boolean if a field has been set.

### GetHardwareReservation

`func (o *FindDeviceById200Response) GetHardwareReservation() FindBatchById200ResponseDevicesInner`

GetHardwareReservation returns the HardwareReservation field if non-nil, zero value otherwise.

### GetHardwareReservationOk

`func (o *FindDeviceById200Response) GetHardwareReservationOk() (*FindBatchById200ResponseDevicesInner, bool)`

GetHardwareReservationOk returns a tuple with the HardwareReservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareReservation

`func (o *FindDeviceById200Response) SetHardwareReservation(v FindBatchById200ResponseDevicesInner)`

SetHardwareReservation sets HardwareReservation field to given value.

### HasHardwareReservation

`func (o *FindDeviceById200Response) HasHardwareReservation() bool`

HasHardwareReservation returns a boolean if a field has been set.

### GetHostname

`func (o *FindDeviceById200Response) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *FindDeviceById200Response) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *FindDeviceById200Response) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *FindDeviceById200Response) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetHref

`func (o *FindDeviceById200Response) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *FindDeviceById200Response) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *FindDeviceById200Response) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *FindDeviceById200Response) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *FindDeviceById200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FindDeviceById200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FindDeviceById200Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FindDeviceById200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImageUrl

`func (o *FindDeviceById200Response) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *FindDeviceById200Response) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *FindDeviceById200Response) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *FindDeviceById200Response) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetIpAddresses

`func (o *FindDeviceById200Response) GetIpAddresses() []FindDeviceById200ResponseIpAddressesInner`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *FindDeviceById200Response) GetIpAddressesOk() (*[]FindDeviceById200ResponseIpAddressesInner, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *FindDeviceById200Response) SetIpAddresses(v []FindDeviceById200ResponseIpAddressesInner)`

SetIpAddresses sets IpAddresses field to given value.

### HasIpAddresses

`func (o *FindDeviceById200Response) HasIpAddresses() bool`

HasIpAddresses returns a boolean if a field has been set.

### GetIpxeScriptUrl

`func (o *FindDeviceById200Response) GetIpxeScriptUrl() string`

GetIpxeScriptUrl returns the IpxeScriptUrl field if non-nil, zero value otherwise.

### GetIpxeScriptUrlOk

`func (o *FindDeviceById200Response) GetIpxeScriptUrlOk() (*string, bool)`

GetIpxeScriptUrlOk returns a tuple with the IpxeScriptUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpxeScriptUrl

`func (o *FindDeviceById200Response) SetIpxeScriptUrl(v string)`

SetIpxeScriptUrl sets IpxeScriptUrl field to given value.

### HasIpxeScriptUrl

`func (o *FindDeviceById200Response) HasIpxeScriptUrl() bool`

HasIpxeScriptUrl returns a boolean if a field has been set.

### GetIqn

`func (o *FindDeviceById200Response) GetIqn() string`

GetIqn returns the Iqn field if non-nil, zero value otherwise.

### GetIqnOk

`func (o *FindDeviceById200Response) GetIqnOk() (*string, bool)`

GetIqnOk returns a tuple with the Iqn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIqn

`func (o *FindDeviceById200Response) SetIqn(v string)`

SetIqn sets Iqn field to given value.

### HasIqn

`func (o *FindDeviceById200Response) HasIqn() bool`

HasIqn returns a boolean if a field has been set.

### GetLocked

`func (o *FindDeviceById200Response) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *FindDeviceById200Response) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *FindDeviceById200Response) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *FindDeviceById200Response) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetMetro

`func (o *FindDeviceById200Response) GetMetro() FindDeviceById200ResponseFacilityMetro`

GetMetro returns the Metro field if non-nil, zero value otherwise.

### GetMetroOk

`func (o *FindDeviceById200Response) GetMetroOk() (*FindDeviceById200ResponseFacilityMetro, bool)`

GetMetroOk returns a tuple with the Metro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetro

`func (o *FindDeviceById200Response) SetMetro(v FindDeviceById200ResponseFacilityMetro)`

SetMetro sets Metro field to given value.

### HasMetro

`func (o *FindDeviceById200Response) HasMetro() bool`

HasMetro returns a boolean if a field has been set.

### GetNetworkPorts

`func (o *FindDeviceById200Response) GetNetworkPorts() []FindDeviceById200ResponseNetworkPortsInner`

GetNetworkPorts returns the NetworkPorts field if non-nil, zero value otherwise.

### GetNetworkPortsOk

`func (o *FindDeviceById200Response) GetNetworkPortsOk() (*[]FindDeviceById200ResponseNetworkPortsInner, bool)`

GetNetworkPortsOk returns a tuple with the NetworkPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPorts

`func (o *FindDeviceById200Response) SetNetworkPorts(v []FindDeviceById200ResponseNetworkPortsInner)`

SetNetworkPorts sets NetworkPorts field to given value.

### HasNetworkPorts

`func (o *FindDeviceById200Response) HasNetworkPorts() bool`

HasNetworkPorts returns a boolean if a field has been set.

### GetOperatingSystem

`func (o *FindDeviceById200Response) GetOperatingSystem() FindDeviceById200ResponseOperatingSystem`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *FindDeviceById200Response) GetOperatingSystemOk() (*FindDeviceById200ResponseOperatingSystem, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *FindDeviceById200Response) SetOperatingSystem(v FindDeviceById200ResponseOperatingSystem)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *FindDeviceById200Response) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### GetActions

`func (o *FindDeviceById200Response) GetActions() []FindDeviceById200ResponseActionsInner`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *FindDeviceById200Response) GetActionsOk() (*[]FindDeviceById200ResponseActionsInner, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *FindDeviceById200Response) SetActions(v []FindDeviceById200ResponseActionsInner)`

SetActions sets Actions field to given value.

### HasActions

`func (o *FindDeviceById200Response) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetPlan

`func (o *FindDeviceById200Response) GetPlan() FindDeviceById200ResponsePlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *FindDeviceById200Response) GetPlanOk() (*FindDeviceById200ResponsePlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *FindDeviceById200Response) SetPlan(v FindDeviceById200ResponsePlan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *FindDeviceById200Response) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetProject

`func (o *FindDeviceById200Response) GetProject() FindDeviceById200ResponseProject`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *FindDeviceById200Response) GetProjectOk() (*FindDeviceById200ResponseProject, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *FindDeviceById200Response) SetProject(v FindDeviceById200ResponseProject)`

SetProject sets Project field to given value.

### HasProject

`func (o *FindDeviceById200Response) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetProjectLite

`func (o *FindDeviceById200Response) GetProjectLite() FindDeviceById200ResponseProjectLite`

GetProjectLite returns the ProjectLite field if non-nil, zero value otherwise.

### GetProjectLiteOk

`func (o *FindDeviceById200Response) GetProjectLiteOk() (*FindDeviceById200ResponseProjectLite, bool)`

GetProjectLiteOk returns a tuple with the ProjectLite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectLite

`func (o *FindDeviceById200Response) SetProjectLite(v FindDeviceById200ResponseProjectLite)`

SetProjectLite sets ProjectLite field to given value.

### HasProjectLite

`func (o *FindDeviceById200Response) HasProjectLite() bool`

HasProjectLite returns a boolean if a field has been set.

### GetProvisioningEvents

`func (o *FindDeviceById200Response) GetProvisioningEvents() []FindInterconnectionEvents200Response`

GetProvisioningEvents returns the ProvisioningEvents field if non-nil, zero value otherwise.

### GetProvisioningEventsOk

`func (o *FindDeviceById200Response) GetProvisioningEventsOk() (*[]FindInterconnectionEvents200Response, bool)`

GetProvisioningEventsOk returns a tuple with the ProvisioningEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningEvents

`func (o *FindDeviceById200Response) SetProvisioningEvents(v []FindInterconnectionEvents200Response)`

SetProvisioningEvents sets ProvisioningEvents field to given value.

### HasProvisioningEvents

`func (o *FindDeviceById200Response) HasProvisioningEvents() bool`

HasProvisioningEvents returns a boolean if a field has been set.

### GetProvisioningPercentage

`func (o *FindDeviceById200Response) GetProvisioningPercentage() float32`

GetProvisioningPercentage returns the ProvisioningPercentage field if non-nil, zero value otherwise.

### GetProvisioningPercentageOk

`func (o *FindDeviceById200Response) GetProvisioningPercentageOk() (*float32, bool)`

GetProvisioningPercentageOk returns a tuple with the ProvisioningPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningPercentage

`func (o *FindDeviceById200Response) SetProvisioningPercentage(v float32)`

SetProvisioningPercentage sets ProvisioningPercentage field to given value.

### HasProvisioningPercentage

`func (o *FindDeviceById200Response) HasProvisioningPercentage() bool`

HasProvisioningPercentage returns a boolean if a field has been set.

### GetRootPassword

`func (o *FindDeviceById200Response) GetRootPassword() string`

GetRootPassword returns the RootPassword field if non-nil, zero value otherwise.

### GetRootPasswordOk

`func (o *FindDeviceById200Response) GetRootPasswordOk() (*string, bool)`

GetRootPasswordOk returns a tuple with the RootPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootPassword

`func (o *FindDeviceById200Response) SetRootPassword(v string)`

SetRootPassword sets RootPassword field to given value.

### HasRootPassword

`func (o *FindDeviceById200Response) HasRootPassword() bool`

HasRootPassword returns a boolean if a field has been set.

### GetShortId

`func (o *FindDeviceById200Response) GetShortId() string`

GetShortId returns the ShortId field if non-nil, zero value otherwise.

### GetShortIdOk

`func (o *FindDeviceById200Response) GetShortIdOk() (*string, bool)`

GetShortIdOk returns a tuple with the ShortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortId

`func (o *FindDeviceById200Response) SetShortId(v string)`

SetShortId sets ShortId field to given value.

### HasShortId

`func (o *FindDeviceById200Response) HasShortId() bool`

HasShortId returns a boolean if a field has been set.

### GetSpotInstance

`func (o *FindDeviceById200Response) GetSpotInstance() bool`

GetSpotInstance returns the SpotInstance field if non-nil, zero value otherwise.

### GetSpotInstanceOk

`func (o *FindDeviceById200Response) GetSpotInstanceOk() (*bool, bool)`

GetSpotInstanceOk returns a tuple with the SpotInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotInstance

`func (o *FindDeviceById200Response) SetSpotInstance(v bool)`

SetSpotInstance sets SpotInstance field to given value.

### HasSpotInstance

`func (o *FindDeviceById200Response) HasSpotInstance() bool`

HasSpotInstance returns a boolean if a field has been set.

### GetSpotPriceMax

`func (o *FindDeviceById200Response) GetSpotPriceMax() float32`

GetSpotPriceMax returns the SpotPriceMax field if non-nil, zero value otherwise.

### GetSpotPriceMaxOk

`func (o *FindDeviceById200Response) GetSpotPriceMaxOk() (*float32, bool)`

GetSpotPriceMaxOk returns a tuple with the SpotPriceMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotPriceMax

`func (o *FindDeviceById200Response) SetSpotPriceMax(v float32)`

SetSpotPriceMax sets SpotPriceMax field to given value.

### HasSpotPriceMax

`func (o *FindDeviceById200Response) HasSpotPriceMax() bool`

HasSpotPriceMax returns a boolean if a field has been set.

### GetSshKeys

`func (o *FindDeviceById200Response) GetSshKeys() []FindBatchById200ResponseDevicesInner`

GetSshKeys returns the SshKeys field if non-nil, zero value otherwise.

### GetSshKeysOk

`func (o *FindDeviceById200Response) GetSshKeysOk() (*[]FindBatchById200ResponseDevicesInner, bool)`

GetSshKeysOk returns a tuple with the SshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeys

`func (o *FindDeviceById200Response) SetSshKeys(v []FindBatchById200ResponseDevicesInner)`

SetSshKeys sets SshKeys field to given value.

### HasSshKeys

`func (o *FindDeviceById200Response) HasSshKeys() bool`

HasSshKeys returns a boolean if a field has been set.

### GetState

`func (o *FindDeviceById200Response) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *FindDeviceById200Response) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *FindDeviceById200Response) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *FindDeviceById200Response) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSwitchUuid

`func (o *FindDeviceById200Response) GetSwitchUuid() string`

GetSwitchUuid returns the SwitchUuid field if non-nil, zero value otherwise.

### GetSwitchUuidOk

`func (o *FindDeviceById200Response) GetSwitchUuidOk() (*string, bool)`

GetSwitchUuidOk returns a tuple with the SwitchUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchUuid

`func (o *FindDeviceById200Response) SetSwitchUuid(v string)`

SetSwitchUuid sets SwitchUuid field to given value.

### HasSwitchUuid

`func (o *FindDeviceById200Response) HasSwitchUuid() bool`

HasSwitchUuid returns a boolean if a field has been set.

### GetTags

`func (o *FindDeviceById200Response) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FindDeviceById200Response) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FindDeviceById200Response) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FindDeviceById200Response) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTerminationTime

`func (o *FindDeviceById200Response) GetTerminationTime() time.Time`

GetTerminationTime returns the TerminationTime field if non-nil, zero value otherwise.

### GetTerminationTimeOk

`func (o *FindDeviceById200Response) GetTerminationTimeOk() (*time.Time, bool)`

GetTerminationTimeOk returns a tuple with the TerminationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationTime

`func (o *FindDeviceById200Response) SetTerminationTime(v time.Time)`

SetTerminationTime sets TerminationTime field to given value.

### HasTerminationTime

`func (o *FindDeviceById200Response) HasTerminationTime() bool`

HasTerminationTime returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *FindDeviceById200Response) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FindDeviceById200Response) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FindDeviceById200Response) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FindDeviceById200Response) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUser

`func (o *FindDeviceById200Response) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *FindDeviceById200Response) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *FindDeviceById200Response) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *FindDeviceById200Response) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetUserdata

`func (o *FindDeviceById200Response) GetUserdata() string`

GetUserdata returns the Userdata field if non-nil, zero value otherwise.

### GetUserdataOk

`func (o *FindDeviceById200Response) GetUserdataOk() (*string, bool)`

GetUserdataOk returns a tuple with the Userdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserdata

`func (o *FindDeviceById200Response) SetUserdata(v string)`

SetUserdata sets Userdata field to given value.

### HasUserdata

`func (o *FindDeviceById200Response) HasUserdata() bool`

HasUserdata returns a boolean if a field has been set.

### GetVolumes

`func (o *FindDeviceById200Response) GetVolumes() []FindBatchById200ResponseDevicesInner`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *FindDeviceById200Response) GetVolumesOk() (*[]FindBatchById200ResponseDevicesInner, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *FindDeviceById200Response) SetVolumes(v []FindBatchById200ResponseDevicesInner)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *FindDeviceById200Response) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


