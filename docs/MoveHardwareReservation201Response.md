# MoveHardwareReservation201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**CustomRate** | Pointer to **float32** | Amount that will be charged for every billing_cycle. | [optional] 
**Device** | Pointer to [**FindDeviceById200Response**](FindDeviceById200Response.md) |  | [optional] 
**Facility** | Pointer to [**FindDeviceById200ResponseFacility**](FindDeviceById200ResponseFacility.md) |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**NeedOfService** | Pointer to **bool** | Whether this Device requires assistance from Metal Equinix. | [optional] 
**Plan** | Pointer to [**FindDeviceById200ResponsePlan**](FindDeviceById200ResponsePlan.md) |  | [optional] 
**Project** | Pointer to [**GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1VrfProject**](GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1VrfProject.md) |  | [optional] 
**Provisionable** | Pointer to **bool** | Whether the reserved server is provisionable or not. Spare devices can&#39;t be provisioned unless they are activated first. | [optional] 
**ShortId** | Pointer to **string** | Short version of the ID. | [optional] 
**Spare** | Pointer to **bool** | Whether the Hardware Reservation is a spare. Spare Hardware Reservations are used when a Hardware Reservations requires service from Metal Equinix | [optional] 
**SwitchUuid** | Pointer to **string** | Switch short id. This can be used to determine if two devices are connected to the same switch, for example. | [optional] 

## Methods

### NewMoveHardwareReservation201Response

`func NewMoveHardwareReservation201Response() *MoveHardwareReservation201Response`

NewMoveHardwareReservation201Response instantiates a new MoveHardwareReservation201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoveHardwareReservation201ResponseWithDefaults

`func NewMoveHardwareReservation201ResponseWithDefaults() *MoveHardwareReservation201Response`

NewMoveHardwareReservation201ResponseWithDefaults instantiates a new MoveHardwareReservation201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *MoveHardwareReservation201Response) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MoveHardwareReservation201Response) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MoveHardwareReservation201Response) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MoveHardwareReservation201Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCustomRate

`func (o *MoveHardwareReservation201Response) GetCustomRate() float32`

GetCustomRate returns the CustomRate field if non-nil, zero value otherwise.

### GetCustomRateOk

`func (o *MoveHardwareReservation201Response) GetCustomRateOk() (*float32, bool)`

GetCustomRateOk returns a tuple with the CustomRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRate

`func (o *MoveHardwareReservation201Response) SetCustomRate(v float32)`

SetCustomRate sets CustomRate field to given value.

### HasCustomRate

`func (o *MoveHardwareReservation201Response) HasCustomRate() bool`

HasCustomRate returns a boolean if a field has been set.

### GetDevice

`func (o *MoveHardwareReservation201Response) GetDevice() FindDeviceById200Response`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *MoveHardwareReservation201Response) GetDeviceOk() (*FindDeviceById200Response, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *MoveHardwareReservation201Response) SetDevice(v FindDeviceById200Response)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *MoveHardwareReservation201Response) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetFacility

`func (o *MoveHardwareReservation201Response) GetFacility() FindDeviceById200ResponseFacility`

GetFacility returns the Facility field if non-nil, zero value otherwise.

### GetFacilityOk

`func (o *MoveHardwareReservation201Response) GetFacilityOk() (*FindDeviceById200ResponseFacility, bool)`

GetFacilityOk returns a tuple with the Facility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacility

`func (o *MoveHardwareReservation201Response) SetFacility(v FindDeviceById200ResponseFacility)`

SetFacility sets Facility field to given value.

### HasFacility

`func (o *MoveHardwareReservation201Response) HasFacility() bool`

HasFacility returns a boolean if a field has been set.

### GetHref

`func (o *MoveHardwareReservation201Response) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *MoveHardwareReservation201Response) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *MoveHardwareReservation201Response) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *MoveHardwareReservation201Response) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *MoveHardwareReservation201Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MoveHardwareReservation201Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MoveHardwareReservation201Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MoveHardwareReservation201Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNeedOfService

`func (o *MoveHardwareReservation201Response) GetNeedOfService() bool`

GetNeedOfService returns the NeedOfService field if non-nil, zero value otherwise.

### GetNeedOfServiceOk

`func (o *MoveHardwareReservation201Response) GetNeedOfServiceOk() (*bool, bool)`

GetNeedOfServiceOk returns a tuple with the NeedOfService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedOfService

`func (o *MoveHardwareReservation201Response) SetNeedOfService(v bool)`

SetNeedOfService sets NeedOfService field to given value.

### HasNeedOfService

`func (o *MoveHardwareReservation201Response) HasNeedOfService() bool`

HasNeedOfService returns a boolean if a field has been set.

### GetPlan

`func (o *MoveHardwareReservation201Response) GetPlan() FindDeviceById200ResponsePlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *MoveHardwareReservation201Response) GetPlanOk() (*FindDeviceById200ResponsePlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *MoveHardwareReservation201Response) SetPlan(v FindDeviceById200ResponsePlan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *MoveHardwareReservation201Response) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetProject

`func (o *MoveHardwareReservation201Response) GetProject() GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1VrfProject`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *MoveHardwareReservation201Response) GetProjectOk() (*GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1VrfProject, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *MoveHardwareReservation201Response) SetProject(v GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1VrfProject)`

SetProject sets Project field to given value.

### HasProject

`func (o *MoveHardwareReservation201Response) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetProvisionable

`func (o *MoveHardwareReservation201Response) GetProvisionable() bool`

GetProvisionable returns the Provisionable field if non-nil, zero value otherwise.

### GetProvisionableOk

`func (o *MoveHardwareReservation201Response) GetProvisionableOk() (*bool, bool)`

GetProvisionableOk returns a tuple with the Provisionable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionable

`func (o *MoveHardwareReservation201Response) SetProvisionable(v bool)`

SetProvisionable sets Provisionable field to given value.

### HasProvisionable

`func (o *MoveHardwareReservation201Response) HasProvisionable() bool`

HasProvisionable returns a boolean if a field has been set.

### GetShortId

`func (o *MoveHardwareReservation201Response) GetShortId() string`

GetShortId returns the ShortId field if non-nil, zero value otherwise.

### GetShortIdOk

`func (o *MoveHardwareReservation201Response) GetShortIdOk() (*string, bool)`

GetShortIdOk returns a tuple with the ShortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortId

`func (o *MoveHardwareReservation201Response) SetShortId(v string)`

SetShortId sets ShortId field to given value.

### HasShortId

`func (o *MoveHardwareReservation201Response) HasShortId() bool`

HasShortId returns a boolean if a field has been set.

### GetSpare

`func (o *MoveHardwareReservation201Response) GetSpare() bool`

GetSpare returns the Spare field if non-nil, zero value otherwise.

### GetSpareOk

`func (o *MoveHardwareReservation201Response) GetSpareOk() (*bool, bool)`

GetSpareOk returns a tuple with the Spare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpare

`func (o *MoveHardwareReservation201Response) SetSpare(v bool)`

SetSpare sets Spare field to given value.

### HasSpare

`func (o *MoveHardwareReservation201Response) HasSpare() bool`

HasSpare returns a boolean if a field has been set.

### GetSwitchUuid

`func (o *MoveHardwareReservation201Response) GetSwitchUuid() string`

GetSwitchUuid returns the SwitchUuid field if non-nil, zero value otherwise.

### GetSwitchUuidOk

`func (o *MoveHardwareReservation201Response) GetSwitchUuidOk() (*string, bool)`

GetSwitchUuidOk returns a tuple with the SwitchUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchUuid

`func (o *MoveHardwareReservation201Response) SetSwitchUuid(v string)`

SetSwitchUuid sets SwitchUuid field to given value.

### HasSwitchUuid

`func (o *MoveHardwareReservation201Response) HasSwitchUuid() bool`

HasSwitchUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


