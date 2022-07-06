# DeviceNetworkPorts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Data** | Pointer to **map[string]interface{}** |  | [optional] 
**DisbondOperationSupported** | Pointer to **bool** | Indicates whether or not the bond can be broken on the port (when applicable). | [optional] 
**VirtualNetworks** | Pointer to [**[]Href**](Href.md) |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 

## Methods

### NewDeviceNetworkPorts

`func NewDeviceNetworkPorts() *DeviceNetworkPorts`

NewDeviceNetworkPorts instantiates a new DeviceNetworkPorts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceNetworkPortsWithDefaults

`func NewDeviceNetworkPortsWithDefaults() *DeviceNetworkPorts`

NewDeviceNetworkPortsWithDefaults instantiates a new DeviceNetworkPorts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeviceNetworkPorts) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeviceNetworkPorts) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeviceNetworkPorts) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DeviceNetworkPorts) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *DeviceNetworkPorts) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DeviceNetworkPorts) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DeviceNetworkPorts) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DeviceNetworkPorts) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *DeviceNetworkPorts) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeviceNetworkPorts) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeviceNetworkPorts) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeviceNetworkPorts) HasName() bool`

HasName returns a boolean if a field has been set.

### GetData

`func (o *DeviceNetworkPorts) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DeviceNetworkPorts) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DeviceNetworkPorts) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *DeviceNetworkPorts) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDisbondOperationSupported

`func (o *DeviceNetworkPorts) GetDisbondOperationSupported() bool`

GetDisbondOperationSupported returns the DisbondOperationSupported field if non-nil, zero value otherwise.

### GetDisbondOperationSupportedOk

`func (o *DeviceNetworkPorts) GetDisbondOperationSupportedOk() (*bool, bool)`

GetDisbondOperationSupportedOk returns a tuple with the DisbondOperationSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisbondOperationSupported

`func (o *DeviceNetworkPorts) SetDisbondOperationSupported(v bool)`

SetDisbondOperationSupported sets DisbondOperationSupported field to given value.

### HasDisbondOperationSupported

`func (o *DeviceNetworkPorts) HasDisbondOperationSupported() bool`

HasDisbondOperationSupported returns a boolean if a field has been set.

### GetVirtualNetworks

`func (o *DeviceNetworkPorts) GetVirtualNetworks() []Href`

GetVirtualNetworks returns the VirtualNetworks field if non-nil, zero value otherwise.

### GetVirtualNetworksOk

`func (o *DeviceNetworkPorts) GetVirtualNetworksOk() (*[]Href, bool)`

GetVirtualNetworksOk returns a tuple with the VirtualNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualNetworks

`func (o *DeviceNetworkPorts) SetVirtualNetworks(v []Href)`

SetVirtualNetworks sets VirtualNetworks field to given value.

### HasVirtualNetworks

`func (o *DeviceNetworkPorts) HasVirtualNetworks() bool`

HasVirtualNetworks returns a boolean if a field has been set.

### GetHref

`func (o *DeviceNetworkPorts) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *DeviceNetworkPorts) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *DeviceNetworkPorts) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *DeviceNetworkPorts) HasHref() bool`

HasHref returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


