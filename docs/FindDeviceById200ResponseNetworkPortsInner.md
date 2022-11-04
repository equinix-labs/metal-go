# FindDeviceById200ResponseNetworkPortsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bond** | Pointer to [**FindDeviceById200ResponseNetworkPortsInnerBond**](FindDeviceById200ResponseNetworkPortsInnerBond.md) |  | [optional] 
**Data** | Pointer to [**FindDeviceById200ResponseNetworkPortsInnerData**](FindDeviceById200ResponseNetworkPortsInnerData.md) |  | [optional] 
**DisbondOperationSupported** | Pointer to **bool** | Indicates whether or not the bond can be broken on the port (when applicable). | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** | Type is either \&quot;NetworkBondPort\&quot; for bond ports or \&quot;NetworkPort\&quot; for bondable ethernet ports | [optional] 
**NetworkType** | Pointer to **string** | Composite network type of the bond | [optional] 
**NativeVirtualNetwork** | Pointer to [**FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork**](FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork.md) |  | [optional] 
**VirtualNetworks** | Pointer to [**[]FindBatchById200ResponseDevicesInner**](FindBatchById200ResponseDevicesInner.md) |  | [optional] 

## Methods

### NewFindDeviceById200ResponseNetworkPortsInner

`func NewFindDeviceById200ResponseNetworkPortsInner() *FindDeviceById200ResponseNetworkPortsInner`

NewFindDeviceById200ResponseNetworkPortsInner instantiates a new FindDeviceById200ResponseNetworkPortsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindDeviceById200ResponseNetworkPortsInnerWithDefaults

`func NewFindDeviceById200ResponseNetworkPortsInnerWithDefaults() *FindDeviceById200ResponseNetworkPortsInner`

NewFindDeviceById200ResponseNetworkPortsInnerWithDefaults instantiates a new FindDeviceById200ResponseNetworkPortsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBond

`func (o *FindDeviceById200ResponseNetworkPortsInner) GetBond() FindDeviceById200ResponseNetworkPortsInnerBond`

GetBond returns the Bond field if non-nil, zero value otherwise.

### GetBondOk

`func (o *FindDeviceById200ResponseNetworkPortsInner) GetBondOk() (*FindDeviceById200ResponseNetworkPortsInnerBond, bool)`

GetBondOk returns a tuple with the Bond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBond

`func (o *FindDeviceById200ResponseNetworkPortsInner) SetBond(v FindDeviceById200ResponseNetworkPortsInnerBond)`

SetBond sets Bond field to given value.

### HasBond

`func (o *FindDeviceById200ResponseNetworkPortsInner) HasBond() bool`

HasBond returns a boolean if a field has been set.

### GetData

`func (o *FindDeviceById200ResponseNetworkPortsInner) GetData() FindDeviceById200ResponseNetworkPortsInnerData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *FindDeviceById200ResponseNetworkPortsInner) GetDataOk() (*FindDeviceById200ResponseNetworkPortsInnerData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *FindDeviceById200ResponseNetworkPortsInner) SetData(v FindDeviceById200ResponseNetworkPortsInnerData)`

SetData sets Data field to given value.

### HasData

`func (o *FindDeviceById200ResponseNetworkPortsInner) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDisbondOperationSupported

`func (o *FindDeviceById200ResponseNetworkPortsInner) GetDisbondOperationSupported() bool`

GetDisbondOperationSupported returns the DisbondOperationSupported field if non-nil, zero value otherwise.

### GetDisbondOperationSupportedOk

`func (o *FindDeviceById200ResponseNetworkPortsInner) GetDisbondOperationSupportedOk() (*bool, bool)`

GetDisbondOperationSupportedOk returns a tuple with the DisbondOperationSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisbondOperationSupported

`func (o *FindDeviceById200ResponseNetworkPortsInner) SetDisbondOperationSupported(v bool)`

SetDisbondOperationSupported sets DisbondOperationSupported field to given value.

### HasDisbondOperationSupported

`func (o *FindDeviceById200ResponseNetworkPortsInner) HasDisbondOperationSupported() bool`

HasDisbondOperationSupported returns a boolean if a field has been set.

### GetHref

`func (o *FindDeviceById200ResponseNetworkPortsInner) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *FindDeviceById200ResponseNetworkPortsInner) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *FindDeviceById200ResponseNetworkPortsInner) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *FindDeviceById200ResponseNetworkPortsInner) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *FindDeviceById200ResponseNetworkPortsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FindDeviceById200ResponseNetworkPortsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FindDeviceById200ResponseNetworkPortsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FindDeviceById200ResponseNetworkPortsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *FindDeviceById200ResponseNetworkPortsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FindDeviceById200ResponseNetworkPortsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FindDeviceById200ResponseNetworkPortsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FindDeviceById200ResponseNetworkPortsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *FindDeviceById200ResponseNetworkPortsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FindDeviceById200ResponseNetworkPortsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FindDeviceById200ResponseNetworkPortsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FindDeviceById200ResponseNetworkPortsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetNetworkType

`func (o *FindDeviceById200ResponseNetworkPortsInner) GetNetworkType() string`

GetNetworkType returns the NetworkType field if non-nil, zero value otherwise.

### GetNetworkTypeOk

`func (o *FindDeviceById200ResponseNetworkPortsInner) GetNetworkTypeOk() (*string, bool)`

GetNetworkTypeOk returns a tuple with the NetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkType

`func (o *FindDeviceById200ResponseNetworkPortsInner) SetNetworkType(v string)`

SetNetworkType sets NetworkType field to given value.

### HasNetworkType

`func (o *FindDeviceById200ResponseNetworkPortsInner) HasNetworkType() bool`

HasNetworkType returns a boolean if a field has been set.

### GetNativeVirtualNetwork

`func (o *FindDeviceById200ResponseNetworkPortsInner) GetNativeVirtualNetwork() FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork`

GetNativeVirtualNetwork returns the NativeVirtualNetwork field if non-nil, zero value otherwise.

### GetNativeVirtualNetworkOk

`func (o *FindDeviceById200ResponseNetworkPortsInner) GetNativeVirtualNetworkOk() (*FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork, bool)`

GetNativeVirtualNetworkOk returns a tuple with the NativeVirtualNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeVirtualNetwork

`func (o *FindDeviceById200ResponseNetworkPortsInner) SetNativeVirtualNetwork(v FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork)`

SetNativeVirtualNetwork sets NativeVirtualNetwork field to given value.

### HasNativeVirtualNetwork

`func (o *FindDeviceById200ResponseNetworkPortsInner) HasNativeVirtualNetwork() bool`

HasNativeVirtualNetwork returns a boolean if a field has been set.

### GetVirtualNetworks

`func (o *FindDeviceById200ResponseNetworkPortsInner) GetVirtualNetworks() []FindBatchById200ResponseDevicesInner`

GetVirtualNetworks returns the VirtualNetworks field if non-nil, zero value otherwise.

### GetVirtualNetworksOk

`func (o *FindDeviceById200ResponseNetworkPortsInner) GetVirtualNetworksOk() (*[]FindBatchById200ResponseDevicesInner, bool)`

GetVirtualNetworksOk returns a tuple with the VirtualNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualNetworks

`func (o *FindDeviceById200ResponseNetworkPortsInner) SetVirtualNetworks(v []FindBatchById200ResponseDevicesInner)`

SetVirtualNetworks sets VirtualNetworks field to given value.

### HasVirtualNetworks

`func (o *FindDeviceById200ResponseNetworkPortsInner) HasVirtualNetworks() bool`

HasVirtualNetworks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


