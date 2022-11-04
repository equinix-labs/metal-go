# CreateDeviceBatchRequestBatchesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostnames** | Pointer to **[]string** |  | [optional] 
**Quantity** | Pointer to **int32** | The number of devices to create in this batch. The hostname may contain an &#x60;{{index}}&#x60; placeholder, which will be replaced with the index of the device in the batch. For example, if the hostname is &#x60;device-{{index}}&#x60;, the first device in the batch will have the hostname &#x60;device-01&#x60;, the second device will have the hostname &#x60;device-02&#x60;, and so on. | [optional] 
**Metro** | **string** | Metro code or ID of where the instance should be provisioned in.  Either metro or facility must be provided. | 
**AlwaysPxe** | Pointer to **bool** | When true, devices with a &#x60;custom_ipxe&#x60; OS will always boot to iPXE. The default setting of false ensures that iPXE will be used on only the first boot. | [optional] [default to false]
**BillingCycle** | Pointer to **string** | The billing cycle of the device. | [optional] 
**Customdata** | Pointer to **interface{}** | Customdata is an arbitrary JSON value that can be accessed via the metadata service. | [optional] [default to {}]
**Description** | Pointer to **string** | Any description of the device or how it will be used. This may be used to inform other API consumers with project access. | [optional] 
**Features** | Pointer to **[]string** | The features attribute allows you to optionally specify what features your server should have.  In the API shorthand syntax, all features listed are &#x60;required&#x60;:  &#x60;&#x60;&#x60; { \&quot;features\&quot;: [\&quot;tpm\&quot;] } &#x60;&#x60;&#x60;  Alternatively, if you do not require a certain feature, but would prefer to be assigned a server with that feature if there are any available, you may specify that feature with a &#x60;preferred&#x60; value. The request will not fail if we have no servers with that feature in our inventory. The API offers an alternative syntax for mixing preferred and required features:  &#x60;&#x60;&#x60; { \&quot;features\&quot;: { \&quot;tpm\&quot;: \&quot;required\&quot;, \&quot;raid\&quot;: \&quot;preferred\&quot; } } &#x60;&#x60;&#x60;  The request will only fail if there are no available servers matching the required &#x60;tpm&#x60; criteria. | [optional] 
**HardwareReservationId** | Pointer to **string** | The Hardware Reservation UUID to provision. Alternatively, &#x60;next-available&#x60; can be specified to select from any of the available hardware reservations. An error will be returned if the requested reservation option is not available.  See [Reserved Hardware](https://metal.equinix.com/developers/docs/deploy/reserved/) for more details. | [optional] [default to ""]
**Hostname** | Pointer to **string** | The hostname to use within the operating system. The same hostname may be used on multiple devices within a project. | [optional] 
**IpAddresses** | Pointer to [**[]CreateDeviceRequestOneOfAllOf1IpAddressesInner**](CreateDeviceRequestOneOfAllOf1IpAddressesInner.md) | The &#x60;ip_addresses attribute will allow you to specify the addresses you want created with your device.  The default value configures public IPv4, public IPv6, and private IPv4.  Private IPv4 address is required. When specifying &#x60;ip_addresses&#x60;, one of the array items must enable private IPv4.  Some operating systems require public IPv4 address. In those cases you will receive an error message if public IPv4 is not enabled.  For example, to only configure your server with a private IPv4 address, you can send &#x60;{ \&quot;ip_addresses\&quot;: [{ \&quot;address_family\&quot;: 4, \&quot;public\&quot;: false }] }&#x60;.  It is possible to request a subnet size larger than a &#x60;/30&#x60; by assigning addresses using the UUID(s) of ip_reservations in your project.  For example, &#x60;{ \&quot;ip_addresses\&quot;: [..., {\&quot;address_family\&quot;: 4, \&quot;public\&quot;: true, \&quot;ip_reservations\&quot;: [\&quot;uuid1\&quot;, \&quot;uuid2\&quot;]}] }&#x60;  To access a server without public IPs, you can use our Out-of-Band console access (SOS) or proxy through another server in the project with public IPs enabled. | [optional] [default to [{"address_family":4,"public":true},{"address_family":4,"public":false},{"address_family":6,"public":true}]]
**IpxeScriptUrl** | Pointer to **string** | When set, the device will chainload an iPXE Script at boot fetched from the supplied URL.  See [Custom iPXE](https://metal.equinix.com/developers/docs/operating-systems/custom-ipxe/) for more details. | [optional] 
**Locked** | Pointer to **bool** | Whether the device should be locked, preventing accidental deletion. | [optional] [default to false]
**NoSshKeys** | Pointer to **bool** | Overrides default behaviour of attaching all of the organization members ssh keys and project ssh keys to device if no specific keys specified | [optional] [default to false]
**OperatingSystem** | **string** | The slug of the operating system to provision. Check the Equinix Metal operating system documentation for rules that may be imposed per operating system, including restrictions on IP address options and device plans. | 
**Plan** | **string** | The slug of the device plan to provision. | 
**PrivateIpv4SubnetSize** | Pointer to **float32** | Deprecated. Use ip_addresses. Subnet range for addresses allocated to this device. | [optional] [default to 28]
**ProjectSshKeys** | Pointer to **[]string** | A list of UUIDs identifying the device parent project that should be authorized to access this device (typically via /root/.ssh/authorized_keys). These keys will also appear in the device metadata.  If no SSH keys are specified (&#x60;user_ssh_keys&#x60;, &#x60;project_ssh_keys&#x60;, and &#x60;ssh_keys&#x60; are all empty lists or omitted), all parent project keys, parent project members keys and organization members keys will be included. This behaviour can be changed with &#39;no_ssh_keys&#39; option to omit any SSH key being added.  | [optional] 
**PublicIpv4SubnetSize** | Pointer to **float32** | Deprecated. Use ip_addresses. Subnet range for addresses allocated to this device. Your project must have addresses available for a non-default request. | [optional] [default to 31]
**SpotInstance** | Pointer to **bool** | Create a spot instance. Spot instances are created with a maximum bid price. If the bid price is not met, the spot instance will be terminated as indicated by the &#x60;termination_time&#x60; field. | [optional] 
**SpotPriceMax** | Pointer to **float32** | The maximum amount to bid for a spot instance. | [optional] 
**SshKeys** | Pointer to [**[]CreateDeviceRequestOneOfAllOf1SshKeysInner**](CreateDeviceRequestOneOfAllOf1SshKeysInner.md) | A list of new or existing project ssh_keys that should be authorized to access this device (typically via /root/.ssh/authorized_keys). These keys will also appear in the device metadata.  These keys are added in addition to any keys defined by   &#x60;project_ssh_keys&#x60; and &#x60;user_ssh_keys&#x60;.  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**TerminationTime** | Pointer to **time.Time** |  | [optional] 
**UserSshKeys** | Pointer to **[]string** | A list of UUIDs identifying the users that should be authorized to access this device (typically via /root/.ssh/authorized_keys).  These keys will also appear in the device metadata.  The users must be members of the project or organization.  If no SSH keys are specified (&#x60;user_ssh_keys&#x60;, &#x60;project_ssh_keys&#x60;, and &#x60;ssh_keys&#x60; are all empty lists or omitted), all parent project keys, parent project members keys and organization members keys will be included. This behaviour can be changed with &#39;no_ssh_keys&#39; option to omit any SSH key being added.  | [optional] 
**Userdata** | Pointer to **string** | The userdata presented in the metadata service for this device.  Userdata is fetched and interpreted by the operating system installed on the device. Acceptable formats are determined by the operating system, with the exception of a special iPXE enabling syntax which is handled before the operating system starts.  See [Server User Data](https://metal.equinix.com/developers/docs/servers/user-data/) and [Provisioning with Custom iPXE](https://metal.equinix.com/developers/docs/operating-systems/custom-ipxe/#provisioning-with-custom-ipxe) for more details. | [optional] 
**Facility** | [**CreateDeviceRequestOneOf1AllOfFacility**](CreateDeviceRequestOneOf1AllOfFacility.md) |  | 

## Methods

### NewCreateDeviceBatchRequestBatchesInner

`func NewCreateDeviceBatchRequestBatchesInner(metro string, operatingSystem string, plan string, facility CreateDeviceRequestOneOf1AllOfFacility, ) *CreateDeviceBatchRequestBatchesInner`

NewCreateDeviceBatchRequestBatchesInner instantiates a new CreateDeviceBatchRequestBatchesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDeviceBatchRequestBatchesInnerWithDefaults

`func NewCreateDeviceBatchRequestBatchesInnerWithDefaults() *CreateDeviceBatchRequestBatchesInner`

NewCreateDeviceBatchRequestBatchesInnerWithDefaults instantiates a new CreateDeviceBatchRequestBatchesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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

### GetQuantity

`func (o *CreateDeviceBatchRequestBatchesInner) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CreateDeviceBatchRequestBatchesInner) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CreateDeviceBatchRequestBatchesInner) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *CreateDeviceBatchRequestBatchesInner) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

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

`func (o *CreateDeviceBatchRequestBatchesInner) GetCustomdata() interface{}`

GetCustomdata returns the Customdata field if non-nil, zero value otherwise.

### GetCustomdataOk

`func (o *CreateDeviceBatchRequestBatchesInner) GetCustomdataOk() (*interface{}, bool)`

GetCustomdataOk returns a tuple with the Customdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomdata

`func (o *CreateDeviceBatchRequestBatchesInner) SetCustomdata(v interface{})`

SetCustomdata sets Customdata field to given value.

### HasCustomdata

`func (o *CreateDeviceBatchRequestBatchesInner) HasCustomdata() bool`

HasCustomdata returns a boolean if a field has been set.

### SetCustomdataNil

`func (o *CreateDeviceBatchRequestBatchesInner) SetCustomdataNil(b bool)`

 SetCustomdataNil sets the value for Customdata to be an explicit nil

### UnsetCustomdata
`func (o *CreateDeviceBatchRequestBatchesInner) UnsetCustomdata()`

UnsetCustomdata ensures that no value is present for Customdata, not even an explicit nil
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

### GetHardwareReservationId

`func (o *CreateDeviceBatchRequestBatchesInner) GetHardwareReservationId() string`

GetHardwareReservationId returns the HardwareReservationId field if non-nil, zero value otherwise.

### GetHardwareReservationIdOk

`func (o *CreateDeviceBatchRequestBatchesInner) GetHardwareReservationIdOk() (*string, bool)`

GetHardwareReservationIdOk returns a tuple with the HardwareReservationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareReservationId

`func (o *CreateDeviceBatchRequestBatchesInner) SetHardwareReservationId(v string)`

SetHardwareReservationId sets HardwareReservationId field to given value.

### HasHardwareReservationId

`func (o *CreateDeviceBatchRequestBatchesInner) HasHardwareReservationId() bool`

HasHardwareReservationId returns a boolean if a field has been set.

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

### GetIpAddresses

`func (o *CreateDeviceBatchRequestBatchesInner) GetIpAddresses() []CreateDeviceRequestOneOfAllOf1IpAddressesInner`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *CreateDeviceBatchRequestBatchesInner) GetIpAddressesOk() (*[]CreateDeviceRequestOneOfAllOf1IpAddressesInner, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *CreateDeviceBatchRequestBatchesInner) SetIpAddresses(v []CreateDeviceRequestOneOfAllOf1IpAddressesInner)`

SetIpAddresses sets IpAddresses field to given value.

### HasIpAddresses

`func (o *CreateDeviceBatchRequestBatchesInner) HasIpAddresses() bool`

HasIpAddresses returns a boolean if a field has been set.

### GetIpxeScriptUrl

`func (o *CreateDeviceBatchRequestBatchesInner) GetIpxeScriptUrl() string`

GetIpxeScriptUrl returns the IpxeScriptUrl field if non-nil, zero value otherwise.

### GetIpxeScriptUrlOk

`func (o *CreateDeviceBatchRequestBatchesInner) GetIpxeScriptUrlOk() (*string, bool)`

GetIpxeScriptUrlOk returns a tuple with the IpxeScriptUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpxeScriptUrl

`func (o *CreateDeviceBatchRequestBatchesInner) SetIpxeScriptUrl(v string)`

SetIpxeScriptUrl sets IpxeScriptUrl field to given value.

### HasIpxeScriptUrl

`func (o *CreateDeviceBatchRequestBatchesInner) HasIpxeScriptUrl() bool`

HasIpxeScriptUrl returns a boolean if a field has been set.

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


### GetPrivateIpv4SubnetSize

`func (o *CreateDeviceBatchRequestBatchesInner) GetPrivateIpv4SubnetSize() float32`

GetPrivateIpv4SubnetSize returns the PrivateIpv4SubnetSize field if non-nil, zero value otherwise.

### GetPrivateIpv4SubnetSizeOk

`func (o *CreateDeviceBatchRequestBatchesInner) GetPrivateIpv4SubnetSizeOk() (*float32, bool)`

GetPrivateIpv4SubnetSizeOk returns a tuple with the PrivateIpv4SubnetSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIpv4SubnetSize

`func (o *CreateDeviceBatchRequestBatchesInner) SetPrivateIpv4SubnetSize(v float32)`

SetPrivateIpv4SubnetSize sets PrivateIpv4SubnetSize field to given value.

### HasPrivateIpv4SubnetSize

`func (o *CreateDeviceBatchRequestBatchesInner) HasPrivateIpv4SubnetSize() bool`

HasPrivateIpv4SubnetSize returns a boolean if a field has been set.

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

### GetPublicIpv4SubnetSize

`func (o *CreateDeviceBatchRequestBatchesInner) GetPublicIpv4SubnetSize() float32`

GetPublicIpv4SubnetSize returns the PublicIpv4SubnetSize field if non-nil, zero value otherwise.

### GetPublicIpv4SubnetSizeOk

`func (o *CreateDeviceBatchRequestBatchesInner) GetPublicIpv4SubnetSizeOk() (*float32, bool)`

GetPublicIpv4SubnetSizeOk returns a tuple with the PublicIpv4SubnetSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpv4SubnetSize

`func (o *CreateDeviceBatchRequestBatchesInner) SetPublicIpv4SubnetSize(v float32)`

SetPublicIpv4SubnetSize sets PublicIpv4SubnetSize field to given value.

### HasPublicIpv4SubnetSize

`func (o *CreateDeviceBatchRequestBatchesInner) HasPublicIpv4SubnetSize() bool`

HasPublicIpv4SubnetSize returns a boolean if a field has been set.

### GetSpotInstance

`func (o *CreateDeviceBatchRequestBatchesInner) GetSpotInstance() bool`

GetSpotInstance returns the SpotInstance field if non-nil, zero value otherwise.

### GetSpotInstanceOk

`func (o *CreateDeviceBatchRequestBatchesInner) GetSpotInstanceOk() (*bool, bool)`

GetSpotInstanceOk returns a tuple with the SpotInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotInstance

`func (o *CreateDeviceBatchRequestBatchesInner) SetSpotInstance(v bool)`

SetSpotInstance sets SpotInstance field to given value.

### HasSpotInstance

`func (o *CreateDeviceBatchRequestBatchesInner) HasSpotInstance() bool`

HasSpotInstance returns a boolean if a field has been set.

### GetSpotPriceMax

`func (o *CreateDeviceBatchRequestBatchesInner) GetSpotPriceMax() float32`

GetSpotPriceMax returns the SpotPriceMax field if non-nil, zero value otherwise.

### GetSpotPriceMaxOk

`func (o *CreateDeviceBatchRequestBatchesInner) GetSpotPriceMaxOk() (*float32, bool)`

GetSpotPriceMaxOk returns a tuple with the SpotPriceMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotPriceMax

`func (o *CreateDeviceBatchRequestBatchesInner) SetSpotPriceMax(v float32)`

SetSpotPriceMax sets SpotPriceMax field to given value.

### HasSpotPriceMax

`func (o *CreateDeviceBatchRequestBatchesInner) HasSpotPriceMax() bool`

HasSpotPriceMax returns a boolean if a field has been set.

### GetSshKeys

`func (o *CreateDeviceBatchRequestBatchesInner) GetSshKeys() []CreateDeviceRequestOneOfAllOf1SshKeysInner`

GetSshKeys returns the SshKeys field if non-nil, zero value otherwise.

### GetSshKeysOk

`func (o *CreateDeviceBatchRequestBatchesInner) GetSshKeysOk() (*[]CreateDeviceRequestOneOfAllOf1SshKeysInner, bool)`

GetSshKeysOk returns a tuple with the SshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeys

`func (o *CreateDeviceBatchRequestBatchesInner) SetSshKeys(v []CreateDeviceRequestOneOfAllOf1SshKeysInner)`

SetSshKeys sets SshKeys field to given value.

### HasSshKeys

`func (o *CreateDeviceBatchRequestBatchesInner) HasSshKeys() bool`

HasSshKeys returns a boolean if a field has been set.

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

### GetFacility

`func (o *CreateDeviceBatchRequestBatchesInner) GetFacility() CreateDeviceRequestOneOf1AllOfFacility`

GetFacility returns the Facility field if non-nil, zero value otherwise.

### GetFacilityOk

`func (o *CreateDeviceBatchRequestBatchesInner) GetFacilityOk() (*CreateDeviceRequestOneOf1AllOfFacility, bool)`

GetFacilityOk returns a tuple with the Facility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacility

`func (o *CreateDeviceBatchRequestBatchesInner) SetFacility(v CreateDeviceRequestOneOf1AllOfFacility)`

SetFacility sets Facility field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


