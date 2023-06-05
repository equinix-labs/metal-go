# DeviceCreateInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlwaysPxe** | Pointer to **bool** | When true, devices with a &#x60;custom_ipxe&#x60; OS will always boot to iPXE. The default setting of false ensures that iPXE will be used on only the first boot. | [optional] [default to false]
**BillingCycle** | Pointer to **string** | The billing cycle of the device. | [optional] 
**Customdata** | Pointer to **map[string]interface{}** | Customdata is an arbitrary JSON value that can be accessed via the metadata service. | [optional] [default to {}]
**Description** | Pointer to **string** | Any description of the device or how it will be used. This may be used to inform other API consumers with project access. | [optional] 
**Features** | Pointer to **[]string** | The features attribute allows you to optionally specify what features your server should have.  In the API shorthand syntax, all features listed are &#x60;required&#x60;:  &#x60;&#x60;&#x60; { \&quot;features\&quot;: [\&quot;tpm\&quot;] } &#x60;&#x60;&#x60;  Alternatively, if you do not require a certain feature, but would prefer to be assigned a server with that feature if there are any available, you may specify that feature with a &#x60;preferred&#x60; value. The request will not fail if we have no servers with that feature in our inventory. The API offers an alternative syntax for mixing preferred and required features:  &#x60;&#x60;&#x60; { \&quot;features\&quot;: { \&quot;tpm\&quot;: \&quot;required\&quot;, \&quot;raid\&quot;: \&quot;preferred\&quot; } } &#x60;&#x60;&#x60;  The request will only fail if there are no available servers matching the required &#x60;tpm&#x60; criteria. | [optional] 
**HardwareReservationId** | Pointer to **string** | The Hardware Reservation UUID to provision. Alternatively, &#x60;next-available&#x60; can be specified to select from any of the available hardware reservations. An error will be returned if the requested reservation option is not available.  See [Reserved Hardware](https://metal.equinix.com/developers/docs/deploy/reserved/) for more details. | [optional] [default to ""]
**Hostname** | Pointer to **string** | The hostname to use within the operating system. The same hostname may be used on multiple devices within a project. | [optional] 
**IpAddresses** | Pointer to [**[]IPAddress**](IPAddress.md) | The &#x60;ip_addresses attribute will allow you to specify the addresses you want created with your device.  The default value configures public IPv4, public IPv6, and private IPv4.  Private IPv4 address is required. When specifying &#x60;ip_addresses&#x60;, one of the array items must enable private IPv4.  Some operating systems require public IPv4 address. In those cases you will receive an error message if public IPv4 is not enabled.  For example, to only configure your server with a private IPv4 address, you can send &#x60;{ \&quot;ip_addresses\&quot;: [{ \&quot;address_family\&quot;: 4, \&quot;public\&quot;: false }] }&#x60;.  It is possible to request a subnet size larger than a &#x60;/30&#x60; by assigning addresses using the UUID(s) of ip_reservations in your project.  For example, &#x60;{ \&quot;ip_addresses\&quot;: [..., {\&quot;address_family\&quot;: 4, \&quot;public\&quot;: true, \&quot;ip_reservations\&quot;: [\&quot;uuid1\&quot;, \&quot;uuid2\&quot;]}] }&#x60;  To access a server without public IPs, you can use our Out-of-Band console access (SOS) or proxy through another server in the project with public IPs enabled. | [optional] [default to [{address_family=4, public=true}, {address_family=4, public=false}, {address_family=6, public=true}]]
**IpxeScriptUrl** | Pointer to **string** | When set, the device will chainload an iPXE Script at boot fetched from the supplied URL.  See [Custom iPXE](https://metal.equinix.com/developers/docs/operating-systems/custom-ipxe/) for more details. | [optional] 
**Locked** | Pointer to **bool** | Whether the device should be locked, preventing accidental deletion. | [optional] [default to false]
**NoSshKeys** | Pointer to **bool** | Overrides default behaviour of attaching all of the organization members ssh keys and project ssh keys to device if no specific keys specified | [optional] [default to false]
**OperatingSystem** | **string** | The slug of the operating system to provision. Check the Equinix Metal operating system documentation for rules that may be imposed per operating system, including restrictions on IP address options and device plans. | 
**Plan** | **string** | The slug of the device plan to provision. | 
**PrivateIpv4SubnetSize** | Pointer to **int32** | Deprecated. Use ip_addresses. Subnet range for addresses allocated to this device. | [optional] [default to 28]
**ProjectSshKeys** | Pointer to **[]string** | A list of UUIDs identifying the device parent project that should be authorized to access this device (typically via /root/.ssh/authorized_keys). These keys will also appear in the device metadata.  If no SSH keys are specified (&#x60;user_ssh_keys&#x60;, &#x60;project_ssh_keys&#x60;, and &#x60;ssh_keys&#x60; are all empty lists or omitted), all parent project keys, parent project members keys and organization members keys will be included. This behaviour can be changed with &#39;no_ssh_keys&#39; option to omit any SSH key being added.  | [optional] 
**PublicIpv4SubnetSize** | Pointer to **int32** | Deprecated. Use ip_addresses. Subnet range for addresses allocated to this device. Your project must have addresses available for a non-default request. | [optional] [default to 31]
**SpotInstance** | Pointer to **bool** | Create a spot instance. Spot instances are created with a maximum bid price. If the bid price is not met, the spot instance will be terminated as indicated by the &#x60;termination_time&#x60; field. | [optional] 
**SpotPriceMax** | Pointer to **float32** | The maximum amount to bid for a spot instance. | [optional] 
**SshKeys** | Pointer to [**[]SSHKeyInput**](SSHKeyInput.md) | A list of new or existing project ssh_keys that should be authorized to access this device (typically via /root/.ssh/authorized_keys). These keys will also appear in the device metadata.  These keys are added in addition to any keys defined by   &#x60;project_ssh_keys&#x60; and &#x60;user_ssh_keys&#x60;.  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**TerminationTime** | Pointer to **time.Time** |  | [optional] 
**UserSshKeys** | Pointer to **[]string** | A list of UUIDs identifying the users that should be authorized to access this device (typically via /root/.ssh/authorized_keys).  These keys will also appear in the device metadata.  The users must be members of the project or organization.  If no SSH keys are specified (&#x60;user_ssh_keys&#x60;, &#x60;project_ssh_keys&#x60;, and &#x60;ssh_keys&#x60; are all empty lists or omitted), all parent project keys, parent project members keys and organization members keys will be included. This behaviour can be changed with &#39;no_ssh_keys&#39; option to omit any SSH key being added.  | [optional] 
**Userdata** | Pointer to **string** | The userdata presented in the metadata service for this device.  Userdata is fetched and interpreted by the operating system installed on the device. Acceptable formats are determined by the operating system, with the exception of a special iPXE enabling syntax which is handled before the operating system starts.  See [Server User Data](https://metal.equinix.com/developers/docs/servers/user-data/) and [Provisioning with Custom iPXE](https://metal.equinix.com/developers/docs/operating-systems/custom-ipxe/#provisioning-with-custom-ipxe) for more details. | [optional] 

## Methods

### NewDeviceCreateInput

`func NewDeviceCreateInput(operatingSystem string, plan string, ) *DeviceCreateInput`

NewDeviceCreateInput instantiates a new DeviceCreateInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceCreateInputWithDefaults

`func NewDeviceCreateInputWithDefaults() *DeviceCreateInput`

NewDeviceCreateInputWithDefaults instantiates a new DeviceCreateInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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

### GetIpAddresses

`func (o *DeviceCreateInput) GetIpAddresses() []IPAddress`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *DeviceCreateInput) GetIpAddressesOk() (*[]IPAddress, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *DeviceCreateInput) SetIpAddresses(v []IPAddress)`

SetIpAddresses sets IpAddresses field to given value.

### HasIpAddresses

`func (o *DeviceCreateInput) HasIpAddresses() bool`

HasIpAddresses returns a boolean if a field has been set.

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

### GetNoSshKeys

`func (o *DeviceCreateInput) GetNoSshKeys() bool`

GetNoSshKeys returns the NoSshKeys field if non-nil, zero value otherwise.

### GetNoSshKeysOk

`func (o *DeviceCreateInput) GetNoSshKeysOk() (*bool, bool)`

GetNoSshKeysOk returns a tuple with the NoSshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoSshKeys

`func (o *DeviceCreateInput) SetNoSshKeys(v bool)`

SetNoSshKeys sets NoSshKeys field to given value.

### HasNoSshKeys

`func (o *DeviceCreateInput) HasNoSshKeys() bool`

HasNoSshKeys returns a boolean if a field has been set.

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


### GetPrivateIpv4SubnetSize

`func (o *DeviceCreateInput) GetPrivateIpv4SubnetSize() int32`

GetPrivateIpv4SubnetSize returns the PrivateIpv4SubnetSize field if non-nil, zero value otherwise.

### GetPrivateIpv4SubnetSizeOk

`func (o *DeviceCreateInput) GetPrivateIpv4SubnetSizeOk() (*int32, bool)`

GetPrivateIpv4SubnetSizeOk returns a tuple with the PrivateIpv4SubnetSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIpv4SubnetSize

`func (o *DeviceCreateInput) SetPrivateIpv4SubnetSize(v int32)`

SetPrivateIpv4SubnetSize sets PrivateIpv4SubnetSize field to given value.

### HasPrivateIpv4SubnetSize

`func (o *DeviceCreateInput) HasPrivateIpv4SubnetSize() bool`

HasPrivateIpv4SubnetSize returns a boolean if a field has been set.

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

### GetPublicIpv4SubnetSize

`func (o *DeviceCreateInput) GetPublicIpv4SubnetSize() int32`

GetPublicIpv4SubnetSize returns the PublicIpv4SubnetSize field if non-nil, zero value otherwise.

### GetPublicIpv4SubnetSizeOk

`func (o *DeviceCreateInput) GetPublicIpv4SubnetSizeOk() (*int32, bool)`

GetPublicIpv4SubnetSizeOk returns a tuple with the PublicIpv4SubnetSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpv4SubnetSize

`func (o *DeviceCreateInput) SetPublicIpv4SubnetSize(v int32)`

SetPublicIpv4SubnetSize sets PublicIpv4SubnetSize field to given value.

### HasPublicIpv4SubnetSize

`func (o *DeviceCreateInput) HasPublicIpv4SubnetSize() bool`

HasPublicIpv4SubnetSize returns a boolean if a field has been set.

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

### GetSshKeys

`func (o *DeviceCreateInput) GetSshKeys() []SSHKeyInput`

GetSshKeys returns the SshKeys field if non-nil, zero value otherwise.

### GetSshKeysOk

`func (o *DeviceCreateInput) GetSshKeysOk() (*[]SSHKeyInput, bool)`

GetSshKeysOk returns a tuple with the SshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeys

`func (o *DeviceCreateInput) SetSshKeys(v []SSHKeyInput)`

SetSshKeys sets SshKeys field to given value.

### HasSshKeys

`func (o *DeviceCreateInput) HasSshKeys() bool`

HasSshKeys returns a boolean if a field has been set.

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


