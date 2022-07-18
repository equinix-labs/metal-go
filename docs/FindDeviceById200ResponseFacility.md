# FindDeviceById200ResponseFacility

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

### NewFindDeviceById200ResponseFacility

`func NewFindDeviceById200ResponseFacility() *FindDeviceById200ResponseFacility`

NewFindDeviceById200ResponseFacility instantiates a new FindDeviceById200ResponseFacility object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindDeviceById200ResponseFacilityWithDefaults

`func NewFindDeviceById200ResponseFacilityWithDefaults() *FindDeviceById200ResponseFacility`

NewFindDeviceById200ResponseFacilityWithDefaults instantiates a new FindDeviceById200ResponseFacility object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *FindDeviceById200ResponseFacility) GetAddress() FindDeviceById200ResponseFacilityAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *FindDeviceById200ResponseFacility) GetAddressOk() (*FindDeviceById200ResponseFacilityAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *FindDeviceById200ResponseFacility) SetAddress(v FindDeviceById200ResponseFacilityAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *FindDeviceById200ResponseFacility) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetCode

`func (o *FindDeviceById200ResponseFacility) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *FindDeviceById200ResponseFacility) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *FindDeviceById200ResponseFacility) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *FindDeviceById200ResponseFacility) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetFeatures

`func (o *FindDeviceById200ResponseFacility) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *FindDeviceById200ResponseFacility) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *FindDeviceById200ResponseFacility) SetFeatures(v []string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *FindDeviceById200ResponseFacility) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetId

`func (o *FindDeviceById200ResponseFacility) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FindDeviceById200ResponseFacility) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FindDeviceById200ResponseFacility) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FindDeviceById200ResponseFacility) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIpRanges

`func (o *FindDeviceById200ResponseFacility) GetIpRanges() []string`

GetIpRanges returns the IpRanges field if non-nil, zero value otherwise.

### GetIpRangesOk

`func (o *FindDeviceById200ResponseFacility) GetIpRangesOk() (*[]string, bool)`

GetIpRangesOk returns a tuple with the IpRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRanges

`func (o *FindDeviceById200ResponseFacility) SetIpRanges(v []string)`

SetIpRanges sets IpRanges field to given value.

### HasIpRanges

`func (o *FindDeviceById200ResponseFacility) HasIpRanges() bool`

HasIpRanges returns a boolean if a field has been set.

### GetMetro

`func (o *FindDeviceById200ResponseFacility) GetMetro() FindDeviceById200ResponseFacilityMetro`

GetMetro returns the Metro field if non-nil, zero value otherwise.

### GetMetroOk

`func (o *FindDeviceById200ResponseFacility) GetMetroOk() (*FindDeviceById200ResponseFacilityMetro, bool)`

GetMetroOk returns a tuple with the Metro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetro

`func (o *FindDeviceById200ResponseFacility) SetMetro(v FindDeviceById200ResponseFacilityMetro)`

SetMetro sets Metro field to given value.

### HasMetro

`func (o *FindDeviceById200ResponseFacility) HasMetro() bool`

HasMetro returns a boolean if a field has been set.

### GetName

`func (o *FindDeviceById200ResponseFacility) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FindDeviceById200ResponseFacility) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FindDeviceById200ResponseFacility) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FindDeviceById200ResponseFacility) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


