# FindDeviceById200ResponseOperatingSystem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Distro** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Licensed** | Pointer to **bool** | Licenced OS is priced according to pricing property | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Preinstallable** | Pointer to **bool** | Servers can be already preinstalled with OS in order to shorten provision time. | [optional] 
**Pricing** | Pointer to **map[string]interface{}** | This object contains price per time unit and optional multiplier value if licence price depends on hardware plan or components (e.g. number of cores) | [optional] 
**ProvisionableOn** | Pointer to **[]string** |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewFindDeviceById200ResponseOperatingSystem

`func NewFindDeviceById200ResponseOperatingSystem() *FindDeviceById200ResponseOperatingSystem`

NewFindDeviceById200ResponseOperatingSystem instantiates a new FindDeviceById200ResponseOperatingSystem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindDeviceById200ResponseOperatingSystemWithDefaults

`func NewFindDeviceById200ResponseOperatingSystemWithDefaults() *FindDeviceById200ResponseOperatingSystem`

NewFindDeviceById200ResponseOperatingSystemWithDefaults instantiates a new FindDeviceById200ResponseOperatingSystem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDistro

`func (o *FindDeviceById200ResponseOperatingSystem) GetDistro() string`

GetDistro returns the Distro field if non-nil, zero value otherwise.

### GetDistroOk

`func (o *FindDeviceById200ResponseOperatingSystem) GetDistroOk() (*string, bool)`

GetDistroOk returns a tuple with the Distro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistro

`func (o *FindDeviceById200ResponseOperatingSystem) SetDistro(v string)`

SetDistro sets Distro field to given value.

### HasDistro

`func (o *FindDeviceById200ResponseOperatingSystem) HasDistro() bool`

HasDistro returns a boolean if a field has been set.

### GetId

`func (o *FindDeviceById200ResponseOperatingSystem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FindDeviceById200ResponseOperatingSystem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FindDeviceById200ResponseOperatingSystem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FindDeviceById200ResponseOperatingSystem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLicensed

`func (o *FindDeviceById200ResponseOperatingSystem) GetLicensed() bool`

GetLicensed returns the Licensed field if non-nil, zero value otherwise.

### GetLicensedOk

`func (o *FindDeviceById200ResponseOperatingSystem) GetLicensedOk() (*bool, bool)`

GetLicensedOk returns a tuple with the Licensed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensed

`func (o *FindDeviceById200ResponseOperatingSystem) SetLicensed(v bool)`

SetLicensed sets Licensed field to given value.

### HasLicensed

`func (o *FindDeviceById200ResponseOperatingSystem) HasLicensed() bool`

HasLicensed returns a boolean if a field has been set.

### GetName

`func (o *FindDeviceById200ResponseOperatingSystem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FindDeviceById200ResponseOperatingSystem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FindDeviceById200ResponseOperatingSystem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FindDeviceById200ResponseOperatingSystem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPreinstallable

`func (o *FindDeviceById200ResponseOperatingSystem) GetPreinstallable() bool`

GetPreinstallable returns the Preinstallable field if non-nil, zero value otherwise.

### GetPreinstallableOk

`func (o *FindDeviceById200ResponseOperatingSystem) GetPreinstallableOk() (*bool, bool)`

GetPreinstallableOk returns a tuple with the Preinstallable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreinstallable

`func (o *FindDeviceById200ResponseOperatingSystem) SetPreinstallable(v bool)`

SetPreinstallable sets Preinstallable field to given value.

### HasPreinstallable

`func (o *FindDeviceById200ResponseOperatingSystem) HasPreinstallable() bool`

HasPreinstallable returns a boolean if a field has been set.

### GetPricing

`func (o *FindDeviceById200ResponseOperatingSystem) GetPricing() map[string]interface{}`

GetPricing returns the Pricing field if non-nil, zero value otherwise.

### GetPricingOk

`func (o *FindDeviceById200ResponseOperatingSystem) GetPricingOk() (*map[string]interface{}, bool)`

GetPricingOk returns a tuple with the Pricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricing

`func (o *FindDeviceById200ResponseOperatingSystem) SetPricing(v map[string]interface{})`

SetPricing sets Pricing field to given value.

### HasPricing

`func (o *FindDeviceById200ResponseOperatingSystem) HasPricing() bool`

HasPricing returns a boolean if a field has been set.

### GetProvisionableOn

`func (o *FindDeviceById200ResponseOperatingSystem) GetProvisionableOn() []string`

GetProvisionableOn returns the ProvisionableOn field if non-nil, zero value otherwise.

### GetProvisionableOnOk

`func (o *FindDeviceById200ResponseOperatingSystem) GetProvisionableOnOk() (*[]string, bool)`

GetProvisionableOnOk returns a tuple with the ProvisionableOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionableOn

`func (o *FindDeviceById200ResponseOperatingSystem) SetProvisionableOn(v []string)`

SetProvisionableOn sets ProvisionableOn field to given value.

### HasProvisionableOn

`func (o *FindDeviceById200ResponseOperatingSystem) HasProvisionableOn() bool`

HasProvisionableOn returns a boolean if a field has been set.

### GetSlug

`func (o *FindDeviceById200ResponseOperatingSystem) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *FindDeviceById200ResponseOperatingSystem) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *FindDeviceById200ResponseOperatingSystem) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *FindDeviceById200ResponseOperatingSystem) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetVersion

`func (o *FindDeviceById200ResponseOperatingSystem) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *FindDeviceById200ResponseOperatingSystem) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *FindDeviceById200ResponseOperatingSystem) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *FindDeviceById200ResponseOperatingSystem) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


