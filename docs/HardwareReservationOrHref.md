# HardwareReservationOrHref

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | **string** |  | 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**CustomRate** | Pointer to **float32** | Amount that will be charged for every billing_cycle. | [optional] 
**Device** | Pointer to [**Device**](Device.md) |  | [optional] 
**Facility** | Pointer to [**Facility**](Facility.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**NeedOfService** | Pointer to **bool** | Whether this Device requires assistance from Equinix Metal. | [optional] 
**Plan** | Pointer to [**Plan**](Plan.md) |  | [optional] 
**Project** | Pointer to [**Project**](Project.md) |  | [optional] 
**Provisionable** | Pointer to **bool** | Whether the reserved server is provisionable or not. Spare devices can&#39;t be provisioned unless they are activated first. | [optional] 
**ShortId** | Pointer to **string** | Short version of the ID. | [optional] 
**Spare** | Pointer to **bool** | Whether the Hardware Reservation is a spare. Spare Hardware Reservations are used when a Hardware Reservations requires service from Equinix Metal | [optional] 
**SwitchUuid** | Pointer to **string** | Switch short id. This can be used to determine if two devices are connected to the same switch, for example. | [optional] 
**TerminationTime** | Pointer to **time.Time** | Expiration date for the reservation. | [optional] 

## Methods

### NewHardwareReservationOrHref

`func NewHardwareReservationOrHref(href string, ) *HardwareReservationOrHref`

NewHardwareReservationOrHref instantiates a new HardwareReservationOrHref object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHardwareReservationOrHrefWithDefaults

`func NewHardwareReservationOrHrefWithDefaults() *HardwareReservationOrHref`

NewHardwareReservationOrHrefWithDefaults instantiates a new HardwareReservationOrHref object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *HardwareReservationOrHref) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *HardwareReservationOrHref) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *HardwareReservationOrHref) SetHref(v string)`

SetHref sets Href field to given value.


### GetCreatedAt

`func (o *HardwareReservationOrHref) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *HardwareReservationOrHref) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *HardwareReservationOrHref) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *HardwareReservationOrHref) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCustomRate

`func (o *HardwareReservationOrHref) GetCustomRate() float32`

GetCustomRate returns the CustomRate field if non-nil, zero value otherwise.

### GetCustomRateOk

`func (o *HardwareReservationOrHref) GetCustomRateOk() (*float32, bool)`

GetCustomRateOk returns a tuple with the CustomRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRate

`func (o *HardwareReservationOrHref) SetCustomRate(v float32)`

SetCustomRate sets CustomRate field to given value.

### HasCustomRate

`func (o *HardwareReservationOrHref) HasCustomRate() bool`

HasCustomRate returns a boolean if a field has been set.

### GetDevice

`func (o *HardwareReservationOrHref) GetDevice() Device`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *HardwareReservationOrHref) GetDeviceOk() (*Device, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *HardwareReservationOrHref) SetDevice(v Device)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *HardwareReservationOrHref) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetFacility

`func (o *HardwareReservationOrHref) GetFacility() Facility`

GetFacility returns the Facility field if non-nil, zero value otherwise.

### GetFacilityOk

`func (o *HardwareReservationOrHref) GetFacilityOk() (*Facility, bool)`

GetFacilityOk returns a tuple with the Facility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacility

`func (o *HardwareReservationOrHref) SetFacility(v Facility)`

SetFacility sets Facility field to given value.

### HasFacility

`func (o *HardwareReservationOrHref) HasFacility() bool`

HasFacility returns a boolean if a field has been set.

### GetId

`func (o *HardwareReservationOrHref) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HardwareReservationOrHref) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HardwareReservationOrHref) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HardwareReservationOrHref) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNeedOfService

`func (o *HardwareReservationOrHref) GetNeedOfService() bool`

GetNeedOfService returns the NeedOfService field if non-nil, zero value otherwise.

### GetNeedOfServiceOk

`func (o *HardwareReservationOrHref) GetNeedOfServiceOk() (*bool, bool)`

GetNeedOfServiceOk returns a tuple with the NeedOfService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedOfService

`func (o *HardwareReservationOrHref) SetNeedOfService(v bool)`

SetNeedOfService sets NeedOfService field to given value.

### HasNeedOfService

`func (o *HardwareReservationOrHref) HasNeedOfService() bool`

HasNeedOfService returns a boolean if a field has been set.

### GetPlan

`func (o *HardwareReservationOrHref) GetPlan() Plan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *HardwareReservationOrHref) GetPlanOk() (*Plan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *HardwareReservationOrHref) SetPlan(v Plan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *HardwareReservationOrHref) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetProject

`func (o *HardwareReservationOrHref) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *HardwareReservationOrHref) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *HardwareReservationOrHref) SetProject(v Project)`

SetProject sets Project field to given value.

### HasProject

`func (o *HardwareReservationOrHref) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetProvisionable

`func (o *HardwareReservationOrHref) GetProvisionable() bool`

GetProvisionable returns the Provisionable field if non-nil, zero value otherwise.

### GetProvisionableOk

`func (o *HardwareReservationOrHref) GetProvisionableOk() (*bool, bool)`

GetProvisionableOk returns a tuple with the Provisionable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionable

`func (o *HardwareReservationOrHref) SetProvisionable(v bool)`

SetProvisionable sets Provisionable field to given value.

### HasProvisionable

`func (o *HardwareReservationOrHref) HasProvisionable() bool`

HasProvisionable returns a boolean if a field has been set.

### GetShortId

`func (o *HardwareReservationOrHref) GetShortId() string`

GetShortId returns the ShortId field if non-nil, zero value otherwise.

### GetShortIdOk

`func (o *HardwareReservationOrHref) GetShortIdOk() (*string, bool)`

GetShortIdOk returns a tuple with the ShortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortId

`func (o *HardwareReservationOrHref) SetShortId(v string)`

SetShortId sets ShortId field to given value.

### HasShortId

`func (o *HardwareReservationOrHref) HasShortId() bool`

HasShortId returns a boolean if a field has been set.

### GetSpare

`func (o *HardwareReservationOrHref) GetSpare() bool`

GetSpare returns the Spare field if non-nil, zero value otherwise.

### GetSpareOk

`func (o *HardwareReservationOrHref) GetSpareOk() (*bool, bool)`

GetSpareOk returns a tuple with the Spare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpare

`func (o *HardwareReservationOrHref) SetSpare(v bool)`

SetSpare sets Spare field to given value.

### HasSpare

`func (o *HardwareReservationOrHref) HasSpare() bool`

HasSpare returns a boolean if a field has been set.

### GetSwitchUuid

`func (o *HardwareReservationOrHref) GetSwitchUuid() string`

GetSwitchUuid returns the SwitchUuid field if non-nil, zero value otherwise.

### GetSwitchUuidOk

`func (o *HardwareReservationOrHref) GetSwitchUuidOk() (*string, bool)`

GetSwitchUuidOk returns a tuple with the SwitchUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchUuid

`func (o *HardwareReservationOrHref) SetSwitchUuid(v string)`

SetSwitchUuid sets SwitchUuid field to given value.

### HasSwitchUuid

`func (o *HardwareReservationOrHref) HasSwitchUuid() bool`

HasSwitchUuid returns a boolean if a field has been set.

### GetTerminationTime

`func (o *HardwareReservationOrHref) GetTerminationTime() time.Time`

GetTerminationTime returns the TerminationTime field if non-nil, zero value otherwise.

### GetTerminationTimeOk

`func (o *HardwareReservationOrHref) GetTerminationTimeOk() (*time.Time, bool)`

GetTerminationTimeOk returns a tuple with the TerminationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationTime

`func (o *HardwareReservationOrHref) SetTerminationTime(v time.Time)`

SetTerminationTime sets TerminationTime field to given value.

### HasTerminationTime

`func (o *HardwareReservationOrHref) HasTerminationTime() bool`

HasTerminationTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


