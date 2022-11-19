# CreateMetalGatewayRequestOneOf1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpReservationId** | **string** | The UUID an a VRF IP Reservation that belongs to the same project as the one in which the Metal Gateway is to be created. Additionally, the VRF IP Reservation and the Virtual Network must reside in the same Metro. | 
**VirtualNetworkId** | **string** | THe UUID of a Metro Virtual Network that belongs to the same project as the one in which the Metal Gateway is to be created. Additionally, the Virtual Network and the VRF IP Reservation must reside in the same metro. | 

## Methods

### NewCreateMetalGatewayRequestOneOf1

`func NewCreateMetalGatewayRequestOneOf1(ipReservationId string, virtualNetworkId string, ) *CreateMetalGatewayRequestOneOf1`

NewCreateMetalGatewayRequestOneOf1 instantiates a new CreateMetalGatewayRequestOneOf1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMetalGatewayRequestOneOf1WithDefaults

`func NewCreateMetalGatewayRequestOneOf1WithDefaults() *CreateMetalGatewayRequestOneOf1`

NewCreateMetalGatewayRequestOneOf1WithDefaults instantiates a new CreateMetalGatewayRequestOneOf1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpReservationId

`func (o *CreateMetalGatewayRequestOneOf1) GetIpReservationId() string`

GetIpReservationId returns the IpReservationId field if non-nil, zero value otherwise.

### GetIpReservationIdOk

`func (o *CreateMetalGatewayRequestOneOf1) GetIpReservationIdOk() (*string, bool)`

GetIpReservationIdOk returns a tuple with the IpReservationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpReservationId

`func (o *CreateMetalGatewayRequestOneOf1) SetIpReservationId(v string)`

SetIpReservationId sets IpReservationId field to given value.


### GetVirtualNetworkId

`func (o *CreateMetalGatewayRequestOneOf1) GetVirtualNetworkId() string`

GetVirtualNetworkId returns the VirtualNetworkId field if non-nil, zero value otherwise.

### GetVirtualNetworkIdOk

`func (o *CreateMetalGatewayRequestOneOf1) GetVirtualNetworkIdOk() (*string, bool)`

GetVirtualNetworkIdOk returns a tuple with the VirtualNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualNetworkId

`func (o *CreateMetalGatewayRequestOneOf1) SetVirtualNetworkId(v string)`

SetVirtualNetworkId sets VirtualNetworkId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


