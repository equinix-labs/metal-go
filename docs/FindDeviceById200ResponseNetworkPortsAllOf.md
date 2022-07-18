# FindDeviceById200ResponseNetworkPortsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **map[string]interface{}** |  | [optional] 
**DisbondOperationSupported** | Pointer to **bool** | Indicates whether or not the bond can be broken on the port (when applicable). | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**VirtualNetworks** | Pointer to [**[]FindBatchById200ResponseDevicesInner**](FindBatchById200ResponseDevicesInner.md) |  | [optional] 

## Methods

### NewFindDeviceById200ResponseNetworkPortsAllOf

`func NewFindDeviceById200ResponseNetworkPortsAllOf() *FindDeviceById200ResponseNetworkPortsAllOf`

NewFindDeviceById200ResponseNetworkPortsAllOf instantiates a new FindDeviceById200ResponseNetworkPortsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindDeviceById200ResponseNetworkPortsAllOfWithDefaults

`func NewFindDeviceById200ResponseNetworkPortsAllOfWithDefaults() *FindDeviceById200ResponseNetworkPortsAllOf`

NewFindDeviceById200ResponseNetworkPortsAllOfWithDefaults instantiates a new FindDeviceById200ResponseNetworkPortsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *FindDeviceById200ResponseNetworkPortsAllOf) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *FindDeviceById200ResponseNetworkPortsAllOf) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *FindDeviceById200ResponseNetworkPortsAllOf) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *FindDeviceById200ResponseNetworkPortsAllOf) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDisbondOperationSupported

`func (o *FindDeviceById200ResponseNetworkPortsAllOf) GetDisbondOperationSupported() bool`

GetDisbondOperationSupported returns the DisbondOperationSupported field if non-nil, zero value otherwise.

### GetDisbondOperationSupportedOk

`func (o *FindDeviceById200ResponseNetworkPortsAllOf) GetDisbondOperationSupportedOk() (*bool, bool)`

GetDisbondOperationSupportedOk returns a tuple with the DisbondOperationSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisbondOperationSupported

`func (o *FindDeviceById200ResponseNetworkPortsAllOf) SetDisbondOperationSupported(v bool)`

SetDisbondOperationSupported sets DisbondOperationSupported field to given value.

### HasDisbondOperationSupported

`func (o *FindDeviceById200ResponseNetworkPortsAllOf) HasDisbondOperationSupported() bool`

HasDisbondOperationSupported returns a boolean if a field has been set.

### GetHref

`func (o *FindDeviceById200ResponseNetworkPortsAllOf) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *FindDeviceById200ResponseNetworkPortsAllOf) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *FindDeviceById200ResponseNetworkPortsAllOf) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *FindDeviceById200ResponseNetworkPortsAllOf) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *FindDeviceById200ResponseNetworkPortsAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FindDeviceById200ResponseNetworkPortsAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FindDeviceById200ResponseNetworkPortsAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FindDeviceById200ResponseNetworkPortsAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *FindDeviceById200ResponseNetworkPortsAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FindDeviceById200ResponseNetworkPortsAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FindDeviceById200ResponseNetworkPortsAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FindDeviceById200ResponseNetworkPortsAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *FindDeviceById200ResponseNetworkPortsAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FindDeviceById200ResponseNetworkPortsAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FindDeviceById200ResponseNetworkPortsAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FindDeviceById200ResponseNetworkPortsAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVirtualNetworks

`func (o *FindDeviceById200ResponseNetworkPortsAllOf) GetVirtualNetworks() []FindBatchById200ResponseDevicesInner`

GetVirtualNetworks returns the VirtualNetworks field if non-nil, zero value otherwise.

### GetVirtualNetworksOk

`func (o *FindDeviceById200ResponseNetworkPortsAllOf) GetVirtualNetworksOk() (*[]FindBatchById200ResponseDevicesInner, bool)`

GetVirtualNetworksOk returns a tuple with the VirtualNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualNetworks

`func (o *FindDeviceById200ResponseNetworkPortsAllOf) SetVirtualNetworks(v []FindBatchById200ResponseDevicesInner)`

SetVirtualNetworks sets VirtualNetworks field to given value.

### HasVirtualNetworks

`func (o *FindDeviceById200ResponseNetworkPortsAllOf) HasVirtualNetworks() bool`

HasVirtualNetworks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


