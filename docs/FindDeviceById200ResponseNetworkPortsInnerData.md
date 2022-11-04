# FindDeviceById200ResponseNetworkPortsInnerData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mac** | Pointer to **string** | MAC address is set for NetworkPort ports | [optional] 
**Bonded** | Pointer to **bool** | Bonded is true for NetworkPort ports in a bond and NetworkBondPort ports that are active | [optional] 

## Methods

### NewFindDeviceById200ResponseNetworkPortsInnerData

`func NewFindDeviceById200ResponseNetworkPortsInnerData() *FindDeviceById200ResponseNetworkPortsInnerData`

NewFindDeviceById200ResponseNetworkPortsInnerData instantiates a new FindDeviceById200ResponseNetworkPortsInnerData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindDeviceById200ResponseNetworkPortsInnerDataWithDefaults

`func NewFindDeviceById200ResponseNetworkPortsInnerDataWithDefaults() *FindDeviceById200ResponseNetworkPortsInnerData`

NewFindDeviceById200ResponseNetworkPortsInnerDataWithDefaults instantiates a new FindDeviceById200ResponseNetworkPortsInnerData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMac

`func (o *FindDeviceById200ResponseNetworkPortsInnerData) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *FindDeviceById200ResponseNetworkPortsInnerData) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *FindDeviceById200ResponseNetworkPortsInnerData) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *FindDeviceById200ResponseNetworkPortsInnerData) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetBonded

`func (o *FindDeviceById200ResponseNetworkPortsInnerData) GetBonded() bool`

GetBonded returns the Bonded field if non-nil, zero value otherwise.

### GetBondedOk

`func (o *FindDeviceById200ResponseNetworkPortsInnerData) GetBondedOk() (*bool, bool)`

GetBondedOk returns a tuple with the Bonded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBonded

`func (o *FindDeviceById200ResponseNetworkPortsInnerData) SetBonded(v bool)`

SetBonded sets Bonded field to given value.

### HasBonded

`func (o *FindDeviceById200ResponseNetworkPortsInnerData) HasBonded() bool`

HasBonded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


