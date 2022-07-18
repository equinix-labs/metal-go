# FindBgpConfigByProject200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asn** | Pointer to **int32** | Autonomous System Number. ASN is required with Global BGP. With Local BGP the private ASN, 65000, is assigned. | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**DeploymentType** | Pointer to **string** | In a Local BGP deployment, a customer uses an internal ASN to control routes within a single Equinix Metal datacenter. This means that the routes are never advertised to the global Internet. Global BGP, on the other hand, requires a customer to have a registered ASN and IP space.  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**MaxPrefix** | Pointer to **int32** | The maximum number of route filters allowed per server | [optional] 
**Md5** | Pointer to **NullableString** | (Optional) Password for BGP session in plaintext (not a checksum) | [optional] 
**Project** | Pointer to [**FindBatchById200ResponseDevicesInner**](FindBatchById200ResponseDevicesInner.md) |  | [optional] 
**Ranges** | Pointer to [**[]FindBgpConfigByProject200ResponseRangesInner**](FindBgpConfigByProject200ResponseRangesInner.md) | The IP block ranges associated to the ASN (Populated in Global BGP only) | [optional] 
**RequestedAt** | Pointer to **time.Time** |  | [optional] 
**RouteObject** | Pointer to **string** | Specifies AS-MACRO (aka AS-SET) to use when building client route filters | [optional] 
**Sessions** | Pointer to [**[]FindBgpSessionById200Response**](FindBgpSessionById200Response.md) | The direct connections between neighboring routers that want to exchange routing information. | [optional] 
**Status** | Pointer to **string** | Status of the BGP Config. Status \&quot;requested\&quot; is valid only with the \&quot;global\&quot; deployment_type. | [optional] 

## Methods

### NewFindBgpConfigByProject200Response

`func NewFindBgpConfigByProject200Response() *FindBgpConfigByProject200Response`

NewFindBgpConfigByProject200Response instantiates a new FindBgpConfigByProject200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindBgpConfigByProject200ResponseWithDefaults

`func NewFindBgpConfigByProject200ResponseWithDefaults() *FindBgpConfigByProject200Response`

NewFindBgpConfigByProject200ResponseWithDefaults instantiates a new FindBgpConfigByProject200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsn

`func (o *FindBgpConfigByProject200Response) GetAsn() int32`

GetAsn returns the Asn field if non-nil, zero value otherwise.

### GetAsnOk

`func (o *FindBgpConfigByProject200Response) GetAsnOk() (*int32, bool)`

GetAsnOk returns a tuple with the Asn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsn

`func (o *FindBgpConfigByProject200Response) SetAsn(v int32)`

SetAsn sets Asn field to given value.

### HasAsn

`func (o *FindBgpConfigByProject200Response) HasAsn() bool`

HasAsn returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FindBgpConfigByProject200Response) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FindBgpConfigByProject200Response) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FindBgpConfigByProject200Response) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FindBgpConfigByProject200Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeploymentType

`func (o *FindBgpConfigByProject200Response) GetDeploymentType() string`

GetDeploymentType returns the DeploymentType field if non-nil, zero value otherwise.

### GetDeploymentTypeOk

`func (o *FindBgpConfigByProject200Response) GetDeploymentTypeOk() (*string, bool)`

GetDeploymentTypeOk returns a tuple with the DeploymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentType

`func (o *FindBgpConfigByProject200Response) SetDeploymentType(v string)`

SetDeploymentType sets DeploymentType field to given value.

### HasDeploymentType

`func (o *FindBgpConfigByProject200Response) HasDeploymentType() bool`

HasDeploymentType returns a boolean if a field has been set.

### GetHref

`func (o *FindBgpConfigByProject200Response) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *FindBgpConfigByProject200Response) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *FindBgpConfigByProject200Response) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *FindBgpConfigByProject200Response) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *FindBgpConfigByProject200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FindBgpConfigByProject200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FindBgpConfigByProject200Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FindBgpConfigByProject200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMaxPrefix

`func (o *FindBgpConfigByProject200Response) GetMaxPrefix() int32`

GetMaxPrefix returns the MaxPrefix field if non-nil, zero value otherwise.

### GetMaxPrefixOk

`func (o *FindBgpConfigByProject200Response) GetMaxPrefixOk() (*int32, bool)`

GetMaxPrefixOk returns a tuple with the MaxPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPrefix

`func (o *FindBgpConfigByProject200Response) SetMaxPrefix(v int32)`

SetMaxPrefix sets MaxPrefix field to given value.

### HasMaxPrefix

`func (o *FindBgpConfigByProject200Response) HasMaxPrefix() bool`

HasMaxPrefix returns a boolean if a field has been set.

### GetMd5

`func (o *FindBgpConfigByProject200Response) GetMd5() string`

GetMd5 returns the Md5 field if non-nil, zero value otherwise.

### GetMd5Ok

`func (o *FindBgpConfigByProject200Response) GetMd5Ok() (*string, bool)`

GetMd5Ok returns a tuple with the Md5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5

`func (o *FindBgpConfigByProject200Response) SetMd5(v string)`

SetMd5 sets Md5 field to given value.

### HasMd5

`func (o *FindBgpConfigByProject200Response) HasMd5() bool`

HasMd5 returns a boolean if a field has been set.

### SetMd5Nil

`func (o *FindBgpConfigByProject200Response) SetMd5Nil(b bool)`

 SetMd5Nil sets the value for Md5 to be an explicit nil

### UnsetMd5
`func (o *FindBgpConfigByProject200Response) UnsetMd5()`

UnsetMd5 ensures that no value is present for Md5, not even an explicit nil
### GetProject

`func (o *FindBgpConfigByProject200Response) GetProject() FindBatchById200ResponseDevicesInner`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *FindBgpConfigByProject200Response) GetProjectOk() (*FindBatchById200ResponseDevicesInner, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *FindBgpConfigByProject200Response) SetProject(v FindBatchById200ResponseDevicesInner)`

SetProject sets Project field to given value.

### HasProject

`func (o *FindBgpConfigByProject200Response) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetRanges

`func (o *FindBgpConfigByProject200Response) GetRanges() []FindBgpConfigByProject200ResponseRangesInner`

GetRanges returns the Ranges field if non-nil, zero value otherwise.

### GetRangesOk

`func (o *FindBgpConfigByProject200Response) GetRangesOk() (*[]FindBgpConfigByProject200ResponseRangesInner, bool)`

GetRangesOk returns a tuple with the Ranges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanges

`func (o *FindBgpConfigByProject200Response) SetRanges(v []FindBgpConfigByProject200ResponseRangesInner)`

SetRanges sets Ranges field to given value.

### HasRanges

`func (o *FindBgpConfigByProject200Response) HasRanges() bool`

HasRanges returns a boolean if a field has been set.

### GetRequestedAt

`func (o *FindBgpConfigByProject200Response) GetRequestedAt() time.Time`

GetRequestedAt returns the RequestedAt field if non-nil, zero value otherwise.

### GetRequestedAtOk

`func (o *FindBgpConfigByProject200Response) GetRequestedAtOk() (*time.Time, bool)`

GetRequestedAtOk returns a tuple with the RequestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedAt

`func (o *FindBgpConfigByProject200Response) SetRequestedAt(v time.Time)`

SetRequestedAt sets RequestedAt field to given value.

### HasRequestedAt

`func (o *FindBgpConfigByProject200Response) HasRequestedAt() bool`

HasRequestedAt returns a boolean if a field has been set.

### GetRouteObject

`func (o *FindBgpConfigByProject200Response) GetRouteObject() string`

GetRouteObject returns the RouteObject field if non-nil, zero value otherwise.

### GetRouteObjectOk

`func (o *FindBgpConfigByProject200Response) GetRouteObjectOk() (*string, bool)`

GetRouteObjectOk returns a tuple with the RouteObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteObject

`func (o *FindBgpConfigByProject200Response) SetRouteObject(v string)`

SetRouteObject sets RouteObject field to given value.

### HasRouteObject

`func (o *FindBgpConfigByProject200Response) HasRouteObject() bool`

HasRouteObject returns a boolean if a field has been set.

### GetSessions

`func (o *FindBgpConfigByProject200Response) GetSessions() []FindBgpSessionById200Response`

GetSessions returns the Sessions field if non-nil, zero value otherwise.

### GetSessionsOk

`func (o *FindBgpConfigByProject200Response) GetSessionsOk() (*[]FindBgpSessionById200Response, bool)`

GetSessionsOk returns a tuple with the Sessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessions

`func (o *FindBgpConfigByProject200Response) SetSessions(v []FindBgpSessionById200Response)`

SetSessions sets Sessions field to given value.

### HasSessions

`func (o *FindBgpConfigByProject200Response) HasSessions() bool`

HasSessions returns a boolean if a field has been set.

### GetStatus

`func (o *FindBgpConfigByProject200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FindBgpConfigByProject200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FindBgpConfigByProject200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FindBgpConfigByProject200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

