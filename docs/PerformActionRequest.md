# PerformActionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Action to perform. See Device.actions for possible actions. | 
**ForceDelete** | Pointer to **bool** | May be required to perform actions under certain conditions | [optional] 
**DeprovisionFast** | Pointer to **bool** | When type is &#x60;reinstall&#x60;, enabling fast deprovisioning will bypass full disk wiping. | [optional] 
**PreserveData** | Pointer to **bool** | When type is &#x60;reinstall&#x60;, preserve the existing data on all disks except the operating-system disk. | [optional] 
**OperatingSystem** | Pointer to **string** | When type is &#x60;reinstall&#x60;, use this &#x60;operating_system&#x60; (defaults to the current &#x60;operating system&#x60;) | [optional] 
**IpxeScriptUrl** | Pointer to **string** | When type is &#x60;reinstall&#x60;, use this &#x60;ipxe_script_url&#x60; (&#x60;operating_system&#x60; must be &#x60;custom_ipxe&#x60;, defaults to the current &#x60;ipxe_script_url&#x60;) | [optional] 

## Methods

### NewPerformActionRequest

`func NewPerformActionRequest(type_ string, ) *PerformActionRequest`

NewPerformActionRequest instantiates a new PerformActionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPerformActionRequestWithDefaults

`func NewPerformActionRequestWithDefaults() *PerformActionRequest`

NewPerformActionRequestWithDefaults instantiates a new PerformActionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PerformActionRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PerformActionRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PerformActionRequest) SetType(v string)`

SetType sets Type field to given value.


### GetForceDelete

`func (o *PerformActionRequest) GetForceDelete() bool`

GetForceDelete returns the ForceDelete field if non-nil, zero value otherwise.

### GetForceDeleteOk

`func (o *PerformActionRequest) GetForceDeleteOk() (*bool, bool)`

GetForceDeleteOk returns a tuple with the ForceDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceDelete

`func (o *PerformActionRequest) SetForceDelete(v bool)`

SetForceDelete sets ForceDelete field to given value.

### HasForceDelete

`func (o *PerformActionRequest) HasForceDelete() bool`

HasForceDelete returns a boolean if a field has been set.

### GetDeprovisionFast

`func (o *PerformActionRequest) GetDeprovisionFast() bool`

GetDeprovisionFast returns the DeprovisionFast field if non-nil, zero value otherwise.

### GetDeprovisionFastOk

`func (o *PerformActionRequest) GetDeprovisionFastOk() (*bool, bool)`

GetDeprovisionFastOk returns a tuple with the DeprovisionFast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprovisionFast

`func (o *PerformActionRequest) SetDeprovisionFast(v bool)`

SetDeprovisionFast sets DeprovisionFast field to given value.

### HasDeprovisionFast

`func (o *PerformActionRequest) HasDeprovisionFast() bool`

HasDeprovisionFast returns a boolean if a field has been set.

### GetPreserveData

`func (o *PerformActionRequest) GetPreserveData() bool`

GetPreserveData returns the PreserveData field if non-nil, zero value otherwise.

### GetPreserveDataOk

`func (o *PerformActionRequest) GetPreserveDataOk() (*bool, bool)`

GetPreserveDataOk returns a tuple with the PreserveData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveData

`func (o *PerformActionRequest) SetPreserveData(v bool)`

SetPreserveData sets PreserveData field to given value.

### HasPreserveData

`func (o *PerformActionRequest) HasPreserveData() bool`

HasPreserveData returns a boolean if a field has been set.

### GetOperatingSystem

`func (o *PerformActionRequest) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *PerformActionRequest) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *PerformActionRequest) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *PerformActionRequest) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### GetIpxeScriptUrl

`func (o *PerformActionRequest) GetIpxeScriptUrl() string`

GetIpxeScriptUrl returns the IpxeScriptUrl field if non-nil, zero value otherwise.

### GetIpxeScriptUrlOk

`func (o *PerformActionRequest) GetIpxeScriptUrlOk() (*string, bool)`

GetIpxeScriptUrlOk returns a tuple with the IpxeScriptUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpxeScriptUrl

`func (o *PerformActionRequest) SetIpxeScriptUrl(v string)`

SetIpxeScriptUrl sets IpxeScriptUrl field to given value.

### HasIpxeScriptUrl

`func (o *PerformActionRequest) HasIpxeScriptUrl() bool`

HasIpxeScriptUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


