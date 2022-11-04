# FindIPAddressById200ResponseOneOfFacility

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to [**FindDeviceById200ResponseFacilityAddress**](FindDeviceById200ResponseFacilityAddress.md) |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Features** | Pointer to **[]string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IpRanges** | Pointer to **[]string** | IP ranges registered in facility. Can be used for GeoIP location | [optional] 
**Metro** | Pointer to [**FindDeviceById200ResponseFacilityMetro**](FindDeviceById200ResponseFacilityMetro.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewFindIPAddressById200ResponseOneOfFacility

`func NewFindIPAddressById200ResponseOneOfFacility() *FindIPAddressById200ResponseOneOfFacility`

NewFindIPAddressById200ResponseOneOfFacility instantiates a new FindIPAddressById200ResponseOneOfFacility object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindIPAddressById200ResponseOneOfFacilityWithDefaults

`func NewFindIPAddressById200ResponseOneOfFacilityWithDefaults() *FindIPAddressById200ResponseOneOfFacility`

NewFindIPAddressById200ResponseOneOfFacilityWithDefaults instantiates a new FindIPAddressById200ResponseOneOfFacility object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *FindIPAddressById200ResponseOneOfFacility) GetAddress() FindDeviceById200ResponseFacilityAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *FindIPAddressById200ResponseOneOfFacility) GetAddressOk() (*FindDeviceById200ResponseFacilityAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *FindIPAddressById200ResponseOneOfFacility) SetAddress(v FindDeviceById200ResponseFacilityAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *FindIPAddressById200ResponseOneOfFacility) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetCode

`func (o *FindIPAddressById200ResponseOneOfFacility) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *FindIPAddressById200ResponseOneOfFacility) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *FindIPAddressById200ResponseOneOfFacility) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *FindIPAddressById200ResponseOneOfFacility) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetFeatures

`func (o *FindIPAddressById200ResponseOneOfFacility) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *FindIPAddressById200ResponseOneOfFacility) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *FindIPAddressById200ResponseOneOfFacility) SetFeatures(v []string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *FindIPAddressById200ResponseOneOfFacility) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetId

`func (o *FindIPAddressById200ResponseOneOfFacility) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FindIPAddressById200ResponseOneOfFacility) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FindIPAddressById200ResponseOneOfFacility) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FindIPAddressById200ResponseOneOfFacility) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIpRanges

`func (o *FindIPAddressById200ResponseOneOfFacility) GetIpRanges() []string`

GetIpRanges returns the IpRanges field if non-nil, zero value otherwise.

### GetIpRangesOk

`func (o *FindIPAddressById200ResponseOneOfFacility) GetIpRangesOk() (*[]string, bool)`

GetIpRangesOk returns a tuple with the IpRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRanges

`func (o *FindIPAddressById200ResponseOneOfFacility) SetIpRanges(v []string)`

SetIpRanges sets IpRanges field to given value.

### HasIpRanges

`func (o *FindIPAddressById200ResponseOneOfFacility) HasIpRanges() bool`

HasIpRanges returns a boolean if a field has been set.

### GetMetro

`func (o *FindIPAddressById200ResponseOneOfFacility) GetMetro() FindDeviceById200ResponseFacilityMetro`

GetMetro returns the Metro field if non-nil, zero value otherwise.

### GetMetroOk

`func (o *FindIPAddressById200ResponseOneOfFacility) GetMetroOk() (*FindDeviceById200ResponseFacilityMetro, bool)`

GetMetroOk returns a tuple with the Metro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetro

`func (o *FindIPAddressById200ResponseOneOfFacility) SetMetro(v FindDeviceById200ResponseFacilityMetro)`

SetMetro sets Metro field to given value.

### HasMetro

`func (o *FindIPAddressById200ResponseOneOfFacility) HasMetro() bool`

HasMetro returns a boolean if a field has been set.

### GetName

`func (o *FindIPAddressById200ResponseOneOfFacility) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FindIPAddressById200ResponseOneOfFacility) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FindIPAddressById200ResponseOneOfFacility) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FindIPAddressById200ResponseOneOfFacility) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


