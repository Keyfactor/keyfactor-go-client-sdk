# KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderParameterUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OidcAudience** | Pointer to **NullableString** |  | [optional] 
**NameClaimType** | Pointer to **NullableString** |  | [optional] 
**RoleClaimType** | Pointer to **NullableString** |  | [optional] 
**UniqueClaimType** | Pointer to **NullableString** |  | [optional] 
**FallbackUniqueClaimType** | Pointer to **NullableString** |  | [optional] 
**ClientId** | Pointer to **NullableString** |  | [optional] 
**ClientSecret** | Pointer to [**CSSCMSDataModelModelsKeyfactorAPISecret**](CSSCMSDataModelModelsKeyfactorAPISecret.md) |  | [optional] 
**Timeout** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewKeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderParameterUpdateRequest

`func NewKeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderParameterUpdateRequest() *KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderParameterUpdateRequest`

NewKeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderParameterUpdateRequest instantiates a new KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderParameterUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderParameterUpdateRequestWithDefaults

`func NewKeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderParameterUpdateRequestWithDefaults() *KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderParameterUpdateRequest`

NewKeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderParameterUpdateRequestWithDefaults instantiates a new KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderParameterUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOidcAudience

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderParameterUpdateRequest) GetOidcAudience() string`

GetOidcAudience returns the OidcAudience field if non-nil, zero value otherwise.

### GetOidcAudienceOk

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderParameterUpdateRequest) GetOidcAudienceOk() (*string, bool)`

GetOidcAudienceOk returns a tuple with the OidcAudience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcAudience

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderParameterUpdateRequest) SetOidcAudience(v string)`

SetOidcAudience sets OidcAudience field to given value.

### HasOidcAudience

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderParameterUpdateRequest) HasOidcAudience() bool`

HasOidcAudience returns a boolean if a field has been set.

### SetOidcAudienceNil

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderParameterUpdateRequest) SetOidcAudienceNil(b bool)`

 SetOidcAudienceNil sets the value for OidcAudience to be an explicit nil

### UnsetOidcAudience
`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderParameterUpdateRequest) UnsetOidcAudience()`

UnsetOidcAudience ensures that no value is present for OidcAudience, not even an explicit nil
### GetNameClaimType

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderParameterUpdateRequest) GetNameClaimType() string`

GetNameClaimType returns the NameClaimType field if non-nil, zero value otherwise.

### GetNameClaimTypeOk

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderParameterUpdateRequest) GetNameClaimTypeOk() (*string, bool)`

GetNameClaimTypeOk returns a tuple with the NameClaimType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameClaimType

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderParameterUpdateRequest) SetNameClaimType(v string)`

SetNameClaimType sets NameClaimType field to given value.

### HasNameClaimType

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderParameterUpdateRequest) HasNameClaimType() bool`

HasNameClaimType returns a boolean if a field has been set.

### SetNameClaimTypeNil

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderParameterUpdateRequest) SetNameClaimTypeNil(b bool)`

 SetNameClaimTypeNil sets the value for NameClaimType to be an explicit nil

### UnsetNameClaimType
`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderParameterUpdateRequest) UnsetNameClaimType()`

UnsetNameClaimType ensures that no value is present for NameClaimType, not even an explicit nil
### GetRoleClaimType

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderParameterUpdateRequest) GetRoleClaimType() string`

GetRoleClaimType returns the RoleClaimType field if non-nil, zero value otherwise.

### GetRoleClaimTypeOk

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderParameterUpdateRequest) GetRoleClaimTypeOk() (*string, bool)`

GetRoleClaimTypeOk returns a tuple with the RoleClaimType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleClaimType

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderParameterUpdateRequest) SetRoleClaimType(v string)`

SetRoleClaimType sets RoleClaimType field to given value.

### HasRoleClaimType

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderParameterUpdateRequest) HasRoleClaimType() bool`

HasRoleClaimType returns a boolean if a field has been set.

### SetRoleClaimTypeNil

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderParameterUpdateRequest) SetRoleClaimTypeNil(b bool)`

 SetRoleClaimTypeNil sets the value for RoleClaimType to be an explicit nil

### UnsetRoleClaimType
`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderParameterUpdateRequest) UnsetRoleClaimType()`

UnsetRoleClaimType ensures that no value is present for RoleClaimType, not even an explicit nil
### GetUniqueClaimType

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderParameterUpdateRequest) GetUniqueClaimType() string`

GetUniqueClaimType returns the UniqueClaimType field if non-nil, zero value otherwise.

### GetUniqueClaimTypeOk

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderParameterUpdateRequest) GetUniqueClaimTypeOk() (*string, bool)`

GetUniqueClaimTypeOk returns a tuple with the UniqueClaimType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueClaimType

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderParameterUpdateRequest) SetUniqueClaimType(v string)`

SetUniqueClaimType sets UniqueClaimType field to given value.

### HasUniqueClaimType

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderParameterUpdateRequest) HasUniqueClaimType() bool`

HasUniqueClaimType returns a boolean if a field has been set.

### SetUniqueClaimTypeNil

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderParameterUpdateRequest) SetUniqueClaimTypeNil(b bool)`

 SetUniqueClaimTypeNil sets the value for UniqueClaimType to be an explicit nil

### UnsetUniqueClaimType
`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderParameterUpdateRequest) UnsetUniqueClaimType()`

UnsetUniqueClaimType ensures that no value is present for UniqueClaimType, not even an explicit nil
### GetFallbackUniqueClaimType

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderParameterUpdateRequest) GetFallbackUniqueClaimType() string`

GetFallbackUniqueClaimType returns the FallbackUniqueClaimType field if non-nil, zero value otherwise.

### GetFallbackUniqueClaimTypeOk

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderParameterUpdateRequest) GetFallbackUniqueClaimTypeOk() (*string, bool)`

GetFallbackUniqueClaimTypeOk returns a tuple with the FallbackUniqueClaimType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackUniqueClaimType

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderParameterUpdateRequest) SetFallbackUniqueClaimType(v string)`

SetFallbackUniqueClaimType sets FallbackUniqueClaimType field to given value.

### HasFallbackUniqueClaimType

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderParameterUpdateRequest) HasFallbackUniqueClaimType() bool`

HasFallbackUniqueClaimType returns a boolean if a field has been set.

### SetFallbackUniqueClaimTypeNil

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderParameterUpdateRequest) SetFallbackUniqueClaimTypeNil(b bool)`

 SetFallbackUniqueClaimTypeNil sets the value for FallbackUniqueClaimType to be an explicit nil

### UnsetFallbackUniqueClaimType
`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderParameterUpdateRequest) UnsetFallbackUniqueClaimType()`

UnsetFallbackUniqueClaimType ensures that no value is present for FallbackUniqueClaimType, not even an explicit nil
### GetClientId

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderParameterUpdateRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderParameterUpdateRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderParameterUpdateRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderParameterUpdateRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### SetClientIdNil

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderParameterUpdateRequest) SetClientIdNil(b bool)`

 SetClientIdNil sets the value for ClientId to be an explicit nil

### UnsetClientId
`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderParameterUpdateRequest) UnsetClientId()`

UnsetClientId ensures that no value is present for ClientId, not even an explicit nil
### GetClientSecret

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderParameterUpdateRequest) GetClientSecret() CSSCMSDataModelModelsKeyfactorAPISecret`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderParameterUpdateRequest) GetClientSecretOk() (*CSSCMSDataModelModelsKeyfactorAPISecret, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderParameterUpdateRequest) SetClientSecret(v CSSCMSDataModelModelsKeyfactorAPISecret)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderParameterUpdateRequest) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetTimeout

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderParameterUpdateRequest) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderParameterUpdateRequest) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderParameterUpdateRequest) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderParameterUpdateRequest) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderParameterUpdateRequest) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderParameterUpdateRequest) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


