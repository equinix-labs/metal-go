# FindDeviceById200ResponsePlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailableIn** | Pointer to [**[]FindDeviceById200ResponsePlanAvailableInInner**](FindDeviceById200ResponsePlanAvailableInInner.md) | Shows which facilities the plan is available in, and the facility-based price if it is different from the default price. | [optional] 
**AvailableInMetros** | Pointer to [**[]FindDeviceById200ResponsePlanAvailableInMetrosInner**](FindDeviceById200ResponsePlanAvailableInMetrosInner.md) | Shows which metros the plan is available in, and the metro-based price if it is different from the default price. | [optional] 
**Class** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DeploymentTypes** | Pointer to **[]string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Legacy** | Pointer to **bool** |  | [optional] 
**Line** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Pricing** | Pointer to **map[string]interface{}** |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**Specs** | Pointer to [**FindDeviceById200ResponsePlanSpecs**](FindDeviceById200ResponsePlanSpecs.md) |  | [optional] 
**Type** | Pointer to **string** | The plan type | [optional] 

## Methods

### NewFindDeviceById200ResponsePlan

`func NewFindDeviceById200ResponsePlan() *FindDeviceById200ResponsePlan`

NewFindDeviceById200ResponsePlan instantiates a new FindDeviceById200ResponsePlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindDeviceById200ResponsePlanWithDefaults

`func NewFindDeviceById200ResponsePlanWithDefaults() *FindDeviceById200ResponsePlan`

NewFindDeviceById200ResponsePlanWithDefaults instantiates a new FindDeviceById200ResponsePlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailableIn

`func (o *FindDeviceById200ResponsePlan) GetAvailableIn() []FindDeviceById200ResponsePlanAvailableInInner`

GetAvailableIn returns the AvailableIn field if non-nil, zero value otherwise.

### GetAvailableInOk

`func (o *FindDeviceById200ResponsePlan) GetAvailableInOk() (*[]FindDeviceById200ResponsePlanAvailableInInner, bool)`

GetAvailableInOk returns a tuple with the AvailableIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableIn

`func (o *FindDeviceById200ResponsePlan) SetAvailableIn(v []FindDeviceById200ResponsePlanAvailableInInner)`

SetAvailableIn sets AvailableIn field to given value.

### HasAvailableIn

`func (o *FindDeviceById200ResponsePlan) HasAvailableIn() bool`

HasAvailableIn returns a boolean if a field has been set.

### GetAvailableInMetros

`func (o *FindDeviceById200ResponsePlan) GetAvailableInMetros() []FindDeviceById200ResponsePlanAvailableInMetrosInner`

GetAvailableInMetros returns the AvailableInMetros field if non-nil, zero value otherwise.

### GetAvailableInMetrosOk

`func (o *FindDeviceById200ResponsePlan) GetAvailableInMetrosOk() (*[]FindDeviceById200ResponsePlanAvailableInMetrosInner, bool)`

GetAvailableInMetrosOk returns a tuple with the AvailableInMetros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableInMetros

`func (o *FindDeviceById200ResponsePlan) SetAvailableInMetros(v []FindDeviceById200ResponsePlanAvailableInMetrosInner)`

SetAvailableInMetros sets AvailableInMetros field to given value.

### HasAvailableInMetros

`func (o *FindDeviceById200ResponsePlan) HasAvailableInMetros() bool`

HasAvailableInMetros returns a boolean if a field has been set.

### GetClass

`func (o *FindDeviceById200ResponsePlan) GetClass() string`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *FindDeviceById200ResponsePlan) GetClassOk() (*string, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *FindDeviceById200ResponsePlan) SetClass(v string)`

SetClass sets Class field to given value.

### HasClass

`func (o *FindDeviceById200ResponsePlan) HasClass() bool`

HasClass returns a boolean if a field has been set.

### GetDescription

`func (o *FindDeviceById200ResponsePlan) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FindDeviceById200ResponsePlan) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FindDeviceById200ResponsePlan) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FindDeviceById200ResponsePlan) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDeploymentTypes

`func (o *FindDeviceById200ResponsePlan) GetDeploymentTypes() []string`

GetDeploymentTypes returns the DeploymentTypes field if non-nil, zero value otherwise.

### GetDeploymentTypesOk

`func (o *FindDeviceById200ResponsePlan) GetDeploymentTypesOk() (*[]string, bool)`

GetDeploymentTypesOk returns a tuple with the DeploymentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentTypes

`func (o *FindDeviceById200ResponsePlan) SetDeploymentTypes(v []string)`

SetDeploymentTypes sets DeploymentTypes field to given value.

### HasDeploymentTypes

`func (o *FindDeviceById200ResponsePlan) HasDeploymentTypes() bool`

HasDeploymentTypes returns a boolean if a field has been set.

### GetId

`func (o *FindDeviceById200ResponsePlan) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FindDeviceById200ResponsePlan) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FindDeviceById200ResponsePlan) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FindDeviceById200ResponsePlan) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLegacy

`func (o *FindDeviceById200ResponsePlan) GetLegacy() bool`

GetLegacy returns the Legacy field if non-nil, zero value otherwise.

### GetLegacyOk

`func (o *FindDeviceById200ResponsePlan) GetLegacyOk() (*bool, bool)`

GetLegacyOk returns a tuple with the Legacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacy

`func (o *FindDeviceById200ResponsePlan) SetLegacy(v bool)`

SetLegacy sets Legacy field to given value.

### HasLegacy

`func (o *FindDeviceById200ResponsePlan) HasLegacy() bool`

HasLegacy returns a boolean if a field has been set.

### GetLine

`func (o *FindDeviceById200ResponsePlan) GetLine() string`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *FindDeviceById200ResponsePlan) GetLineOk() (*string, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *FindDeviceById200ResponsePlan) SetLine(v string)`

SetLine sets Line field to given value.

### HasLine

`func (o *FindDeviceById200ResponsePlan) HasLine() bool`

HasLine returns a boolean if a field has been set.

### GetName

`func (o *FindDeviceById200ResponsePlan) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FindDeviceById200ResponsePlan) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FindDeviceById200ResponsePlan) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FindDeviceById200ResponsePlan) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPricing

`func (o *FindDeviceById200ResponsePlan) GetPricing() map[string]interface{}`

GetPricing returns the Pricing field if non-nil, zero value otherwise.

### GetPricingOk

`func (o *FindDeviceById200ResponsePlan) GetPricingOk() (*map[string]interface{}, bool)`

GetPricingOk returns a tuple with the Pricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricing

`func (o *FindDeviceById200ResponsePlan) SetPricing(v map[string]interface{})`

SetPricing sets Pricing field to given value.

### HasPricing

`func (o *FindDeviceById200ResponsePlan) HasPricing() bool`

HasPricing returns a boolean if a field has been set.

### GetSlug

`func (o *FindDeviceById200ResponsePlan) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *FindDeviceById200ResponsePlan) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *FindDeviceById200ResponsePlan) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *FindDeviceById200ResponsePlan) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetSpecs

`func (o *FindDeviceById200ResponsePlan) GetSpecs() FindDeviceById200ResponsePlanSpecs`

GetSpecs returns the Specs field if non-nil, zero value otherwise.

### GetSpecsOk

`func (o *FindDeviceById200ResponsePlan) GetSpecsOk() (*FindDeviceById200ResponsePlanSpecs, bool)`

GetSpecsOk returns a tuple with the Specs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecs

`func (o *FindDeviceById200ResponsePlan) SetSpecs(v FindDeviceById200ResponsePlanSpecs)`

SetSpecs sets Specs field to given value.

### HasSpecs

`func (o *FindDeviceById200ResponsePlan) HasSpecs() bool`

HasSpecs returns a boolean if a field has been set.

### GetType

`func (o *FindDeviceById200ResponsePlan) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FindDeviceById200ResponsePlan) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FindDeviceById200ResponsePlan) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FindDeviceById200ResponsePlan) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


