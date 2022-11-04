# GetInterconnection200ResponseServiceTokensInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresAt** | Pointer to **time.Time** | The expiration date and time of the Fabric service token. Once a service token is expired, it is no longer redeemable. | [optional] 
**Id** | Pointer to **string** | The UUID that can be used on the Fabric Portal to redeem either an A-Side or Z-Side Service Token. For Fabric VCs (Metal Billed), this UUID will represent an A-Side Service Token, which will allow interconnections to be made from Equinix Metal to other Service Providers on Fabric. For Fabric VCs (Fabric Billed), this UUID will represent a Z-Side Service Token, which will allow interconnections to be made to connect an owned Fabric Port or  Virtual Device to Equinix Metal. | [optional] 
**MaxAllowedSpeed** | Pointer to **int32** | The maximum speed that can be selected on the Fabric Portal when configuring a interconnection with either  an A-Side or Z-Side Service Token. For Fabric VCs (Metal Billed), this is what the billing is based off of, and can be one of the following options, &#39;50mbps&#39;, &#39;200mbps&#39;, &#39;500mbps&#39;, &#39;1gbps&#39;, &#39;2gbps&#39;, &#39;5gbps&#39; or &#39;10gbps&#39;. For Fabric VCs (Fabric Billed), this will default to 10Gbps. | [optional] 
**Role** | Pointer to **string** | Either primary or secondary, depending on which interconnection the service token is associated to. | [optional] 
**ServiceTokenType** | Pointer to **string** | Either &#39;a_side&#39; or &#39;z_side&#39;, depending on which type of Fabric VC was requested. | [optional] 
**State** | Pointer to **string** | The state of the service token that corresponds with the service token state on Fabric. An &#39;inactive&#39; state refers to a token that has not been redeemed yet on the Fabric side, an &#39;active&#39; state refers to a token that has already been redeemed, and an &#39;expired&#39; state refers to a token that has reached its expiry time. | [optional] 

## Methods

### NewGetInterconnection200ResponseServiceTokensInner

`func NewGetInterconnection200ResponseServiceTokensInner() *GetInterconnection200ResponseServiceTokensInner`

NewGetInterconnection200ResponseServiceTokensInner instantiates a new GetInterconnection200ResponseServiceTokensInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInterconnection200ResponseServiceTokensInnerWithDefaults

`func NewGetInterconnection200ResponseServiceTokensInnerWithDefaults() *GetInterconnection200ResponseServiceTokensInner`

NewGetInterconnection200ResponseServiceTokensInnerWithDefaults instantiates a new GetInterconnection200ResponseServiceTokensInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiresAt

`func (o *GetInterconnection200ResponseServiceTokensInner) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *GetInterconnection200ResponseServiceTokensInner) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *GetInterconnection200ResponseServiceTokensInner) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *GetInterconnection200ResponseServiceTokensInner) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetId

`func (o *GetInterconnection200ResponseServiceTokensInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetInterconnection200ResponseServiceTokensInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetInterconnection200ResponseServiceTokensInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetInterconnection200ResponseServiceTokensInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMaxAllowedSpeed

`func (o *GetInterconnection200ResponseServiceTokensInner) GetMaxAllowedSpeed() int32`

GetMaxAllowedSpeed returns the MaxAllowedSpeed field if non-nil, zero value otherwise.

### GetMaxAllowedSpeedOk

`func (o *GetInterconnection200ResponseServiceTokensInner) GetMaxAllowedSpeedOk() (*int32, bool)`

GetMaxAllowedSpeedOk returns a tuple with the MaxAllowedSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAllowedSpeed

`func (o *GetInterconnection200ResponseServiceTokensInner) SetMaxAllowedSpeed(v int32)`

SetMaxAllowedSpeed sets MaxAllowedSpeed field to given value.

### HasMaxAllowedSpeed

`func (o *GetInterconnection200ResponseServiceTokensInner) HasMaxAllowedSpeed() bool`

HasMaxAllowedSpeed returns a boolean if a field has been set.

### GetRole

`func (o *GetInterconnection200ResponseServiceTokensInner) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *GetInterconnection200ResponseServiceTokensInner) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *GetInterconnection200ResponseServiceTokensInner) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *GetInterconnection200ResponseServiceTokensInner) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetServiceTokenType

`func (o *GetInterconnection200ResponseServiceTokensInner) GetServiceTokenType() string`

GetServiceTokenType returns the ServiceTokenType field if non-nil, zero value otherwise.

### GetServiceTokenTypeOk

`func (o *GetInterconnection200ResponseServiceTokensInner) GetServiceTokenTypeOk() (*string, bool)`

GetServiceTokenTypeOk returns a tuple with the ServiceTokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTokenType

`func (o *GetInterconnection200ResponseServiceTokensInner) SetServiceTokenType(v string)`

SetServiceTokenType sets ServiceTokenType field to given value.

### HasServiceTokenType

`func (o *GetInterconnection200ResponseServiceTokensInner) HasServiceTokenType() bool`

HasServiceTokenType returns a boolean if a field has been set.

### GetState

`func (o *GetInterconnection200ResponseServiceTokensInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetInterconnection200ResponseServiceTokensInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetInterconnection200ResponseServiceTokensInner) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GetInterconnection200ResponseServiceTokensInner) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


