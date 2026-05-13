# IdentityProviderProviderTypeParameterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timeout** | Pointer to **NullableInt32** |  | [optional] 
**OIDCAudience** | Pointer to **NullableString** |  | [optional] 
**OIDCScope** | Pointer to **NullableString** |  | [optional] 
**RequestHeaders** | Pointer to [**[]SharedRequestHeaderRequest**](SharedRequestHeaderRequest.md) |  | [optional] 
**RequestURLParameters** | Pointer to [**[]IdentityProviderRequestURLParameterRequest**](IdentityProviderRequestURLParameterRequest.md) |  | [optional] 
**NameClaimType** | **string** |  | 
**RoleClaimType** | **string** |  | 
**UniqueClaimType** | **string** |  | 
**FallbackUniqueClaimType** | **string** |  | 
**ClientId** | **string** |  | 
**ClientSecret** | [**CSSCMSDataModelModelsKeyfactorAPISecret**](CSSCMSDataModelModelsKeyfactorAPISecret.md) |  | 
**AuthorizationEndpoint** | **string** |  | 
**TokenEndpoint** | **string** |  | 
**JSONWebKeySetUri** | **string** |  | 
**Authority** | **string** |  | 
**UserInfoEndpoint** | Pointer to **NullableString** |  | [optional] 
**Auth0ApiUrl** | Pointer to **NullableString** |  | [optional] 
**SignOutUrl** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewIdentityProviderProviderTypeParameterRequest

`func NewIdentityProviderProviderTypeParameterRequest(nameClaimType string, roleClaimType string, uniqueClaimType string, fallbackUniqueClaimType string, clientId string, clientSecret CSSCMSDataModelModelsKeyfactorAPISecret, authorizationEndpoint string, tokenEndpoint string, jSONWebKeySetUri string, authority string, ) *IdentityProviderProviderTypeParameterRequest`

NewIdentityProviderProviderTypeParameterRequest instantiates a new IdentityProviderProviderTypeParameterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityProviderProviderTypeParameterRequestWithDefaults

`func NewIdentityProviderProviderTypeParameterRequestWithDefaults() *IdentityProviderProviderTypeParameterRequest`

NewIdentityProviderProviderTypeParameterRequestWithDefaults instantiates a new IdentityProviderProviderTypeParameterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeout

`func (o *IdentityProviderProviderTypeParameterRequest) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *IdentityProviderProviderTypeParameterRequest) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *IdentityProviderProviderTypeParameterRequest) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *IdentityProviderProviderTypeParameterRequest) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *IdentityProviderProviderTypeParameterRequest) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *IdentityProviderProviderTypeParameterRequest) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetOIDCAudience

`func (o *IdentityProviderProviderTypeParameterRequest) GetOIDCAudience() string`

GetOIDCAudience returns the OIDCAudience field if non-nil, zero value otherwise.

### GetOIDCAudienceOk

`func (o *IdentityProviderProviderTypeParameterRequest) GetOIDCAudienceOk() (*string, bool)`

GetOIDCAudienceOk returns a tuple with the OIDCAudience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOIDCAudience

`func (o *IdentityProviderProviderTypeParameterRequest) SetOIDCAudience(v string)`

SetOIDCAudience sets OIDCAudience field to given value.

### HasOIDCAudience

`func (o *IdentityProviderProviderTypeParameterRequest) HasOIDCAudience() bool`

HasOIDCAudience returns a boolean if a field has been set.

### SetOIDCAudienceNil

`func (o *IdentityProviderProviderTypeParameterRequest) SetOIDCAudienceNil(b bool)`

 SetOIDCAudienceNil sets the value for OIDCAudience to be an explicit nil

### UnsetOIDCAudience
`func (o *IdentityProviderProviderTypeParameterRequest) UnsetOIDCAudience()`

UnsetOIDCAudience ensures that no value is present for OIDCAudience, not even an explicit nil
### GetOIDCScope

`func (o *IdentityProviderProviderTypeParameterRequest) GetOIDCScope() string`

GetOIDCScope returns the OIDCScope field if non-nil, zero value otherwise.

### GetOIDCScopeOk

`func (o *IdentityProviderProviderTypeParameterRequest) GetOIDCScopeOk() (*string, bool)`

GetOIDCScopeOk returns a tuple with the OIDCScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOIDCScope

`func (o *IdentityProviderProviderTypeParameterRequest) SetOIDCScope(v string)`

SetOIDCScope sets OIDCScope field to given value.

### HasOIDCScope

`func (o *IdentityProviderProviderTypeParameterRequest) HasOIDCScope() bool`

HasOIDCScope returns a boolean if a field has been set.

### SetOIDCScopeNil

`func (o *IdentityProviderProviderTypeParameterRequest) SetOIDCScopeNil(b bool)`

 SetOIDCScopeNil sets the value for OIDCScope to be an explicit nil

### UnsetOIDCScope
`func (o *IdentityProviderProviderTypeParameterRequest) UnsetOIDCScope()`

UnsetOIDCScope ensures that no value is present for OIDCScope, not even an explicit nil
### GetRequestHeaders

`func (o *IdentityProviderProviderTypeParameterRequest) GetRequestHeaders() []SharedRequestHeaderRequest`

GetRequestHeaders returns the RequestHeaders field if non-nil, zero value otherwise.

### GetRequestHeadersOk

`func (o *IdentityProviderProviderTypeParameterRequest) GetRequestHeadersOk() (*[]SharedRequestHeaderRequest, bool)`

GetRequestHeadersOk returns a tuple with the RequestHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestHeaders

`func (o *IdentityProviderProviderTypeParameterRequest) SetRequestHeaders(v []SharedRequestHeaderRequest)`

SetRequestHeaders sets RequestHeaders field to given value.

### HasRequestHeaders

`func (o *IdentityProviderProviderTypeParameterRequest) HasRequestHeaders() bool`

HasRequestHeaders returns a boolean if a field has been set.

### SetRequestHeadersNil

`func (o *IdentityProviderProviderTypeParameterRequest) SetRequestHeadersNil(b bool)`

 SetRequestHeadersNil sets the value for RequestHeaders to be an explicit nil

### UnsetRequestHeaders
`func (o *IdentityProviderProviderTypeParameterRequest) UnsetRequestHeaders()`

UnsetRequestHeaders ensures that no value is present for RequestHeaders, not even an explicit nil
### GetRequestURLParameters

`func (o *IdentityProviderProviderTypeParameterRequest) GetRequestURLParameters() []IdentityProviderRequestURLParameterRequest`

GetRequestURLParameters returns the RequestURLParameters field if non-nil, zero value otherwise.

### GetRequestURLParametersOk

`func (o *IdentityProviderProviderTypeParameterRequest) GetRequestURLParametersOk() (*[]IdentityProviderRequestURLParameterRequest, bool)`

GetRequestURLParametersOk returns a tuple with the RequestURLParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestURLParameters

`func (o *IdentityProviderProviderTypeParameterRequest) SetRequestURLParameters(v []IdentityProviderRequestURLParameterRequest)`

SetRequestURLParameters sets RequestURLParameters field to given value.

### HasRequestURLParameters

`func (o *IdentityProviderProviderTypeParameterRequest) HasRequestURLParameters() bool`

HasRequestURLParameters returns a boolean if a field has been set.

### SetRequestURLParametersNil

`func (o *IdentityProviderProviderTypeParameterRequest) SetRequestURLParametersNil(b bool)`

 SetRequestURLParametersNil sets the value for RequestURLParameters to be an explicit nil

### UnsetRequestURLParameters
`func (o *IdentityProviderProviderTypeParameterRequest) UnsetRequestURLParameters()`

UnsetRequestURLParameters ensures that no value is present for RequestURLParameters, not even an explicit nil
### GetNameClaimType

`func (o *IdentityProviderProviderTypeParameterRequest) GetNameClaimType() string`

GetNameClaimType returns the NameClaimType field if non-nil, zero value otherwise.

### GetNameClaimTypeOk

`func (o *IdentityProviderProviderTypeParameterRequest) GetNameClaimTypeOk() (*string, bool)`

GetNameClaimTypeOk returns a tuple with the NameClaimType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameClaimType

`func (o *IdentityProviderProviderTypeParameterRequest) SetNameClaimType(v string)`

SetNameClaimType sets NameClaimType field to given value.


### GetRoleClaimType

`func (o *IdentityProviderProviderTypeParameterRequest) GetRoleClaimType() string`

GetRoleClaimType returns the RoleClaimType field if non-nil, zero value otherwise.

### GetRoleClaimTypeOk

`func (o *IdentityProviderProviderTypeParameterRequest) GetRoleClaimTypeOk() (*string, bool)`

GetRoleClaimTypeOk returns a tuple with the RoleClaimType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleClaimType

`func (o *IdentityProviderProviderTypeParameterRequest) SetRoleClaimType(v string)`

SetRoleClaimType sets RoleClaimType field to given value.


### GetUniqueClaimType

`func (o *IdentityProviderProviderTypeParameterRequest) GetUniqueClaimType() string`

GetUniqueClaimType returns the UniqueClaimType field if non-nil, zero value otherwise.

### GetUniqueClaimTypeOk

`func (o *IdentityProviderProviderTypeParameterRequest) GetUniqueClaimTypeOk() (*string, bool)`

GetUniqueClaimTypeOk returns a tuple with the UniqueClaimType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueClaimType

`func (o *IdentityProviderProviderTypeParameterRequest) SetUniqueClaimType(v string)`

SetUniqueClaimType sets UniqueClaimType field to given value.


### GetFallbackUniqueClaimType

`func (o *IdentityProviderProviderTypeParameterRequest) GetFallbackUniqueClaimType() string`

GetFallbackUniqueClaimType returns the FallbackUniqueClaimType field if non-nil, zero value otherwise.

### GetFallbackUniqueClaimTypeOk

`func (o *IdentityProviderProviderTypeParameterRequest) GetFallbackUniqueClaimTypeOk() (*string, bool)`

GetFallbackUniqueClaimTypeOk returns a tuple with the FallbackUniqueClaimType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackUniqueClaimType

`func (o *IdentityProviderProviderTypeParameterRequest) SetFallbackUniqueClaimType(v string)`

SetFallbackUniqueClaimType sets FallbackUniqueClaimType field to given value.


### GetClientId

`func (o *IdentityProviderProviderTypeParameterRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *IdentityProviderProviderTypeParameterRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *IdentityProviderProviderTypeParameterRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *IdentityProviderProviderTypeParameterRequest) GetClientSecret() CSSCMSDataModelModelsKeyfactorAPISecret`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *IdentityProviderProviderTypeParameterRequest) GetClientSecretOk() (*CSSCMSDataModelModelsKeyfactorAPISecret, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *IdentityProviderProviderTypeParameterRequest) SetClientSecret(v CSSCMSDataModelModelsKeyfactorAPISecret)`

SetClientSecret sets ClientSecret field to given value.


### GetAuthorizationEndpoint

`func (o *IdentityProviderProviderTypeParameterRequest) GetAuthorizationEndpoint() string`

GetAuthorizationEndpoint returns the AuthorizationEndpoint field if non-nil, zero value otherwise.

### GetAuthorizationEndpointOk

`func (o *IdentityProviderProviderTypeParameterRequest) GetAuthorizationEndpointOk() (*string, bool)`

GetAuthorizationEndpointOk returns a tuple with the AuthorizationEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationEndpoint

`func (o *IdentityProviderProviderTypeParameterRequest) SetAuthorizationEndpoint(v string)`

SetAuthorizationEndpoint sets AuthorizationEndpoint field to given value.


### GetTokenEndpoint

`func (o *IdentityProviderProviderTypeParameterRequest) GetTokenEndpoint() string`

GetTokenEndpoint returns the TokenEndpoint field if non-nil, zero value otherwise.

### GetTokenEndpointOk

`func (o *IdentityProviderProviderTypeParameterRequest) GetTokenEndpointOk() (*string, bool)`

GetTokenEndpointOk returns a tuple with the TokenEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpoint

`func (o *IdentityProviderProviderTypeParameterRequest) SetTokenEndpoint(v string)`

SetTokenEndpoint sets TokenEndpoint field to given value.


### GetJSONWebKeySetUri

`func (o *IdentityProviderProviderTypeParameterRequest) GetJSONWebKeySetUri() string`

GetJSONWebKeySetUri returns the JSONWebKeySetUri field if non-nil, zero value otherwise.

### GetJSONWebKeySetUriOk

`func (o *IdentityProviderProviderTypeParameterRequest) GetJSONWebKeySetUriOk() (*string, bool)`

GetJSONWebKeySetUriOk returns a tuple with the JSONWebKeySetUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJSONWebKeySetUri

`func (o *IdentityProviderProviderTypeParameterRequest) SetJSONWebKeySetUri(v string)`

SetJSONWebKeySetUri sets JSONWebKeySetUri field to given value.


### GetAuthority

`func (o *IdentityProviderProviderTypeParameterRequest) GetAuthority() string`

GetAuthority returns the Authority field if non-nil, zero value otherwise.

### GetAuthorityOk

`func (o *IdentityProviderProviderTypeParameterRequest) GetAuthorityOk() (*string, bool)`

GetAuthorityOk returns a tuple with the Authority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthority

`func (o *IdentityProviderProviderTypeParameterRequest) SetAuthority(v string)`

SetAuthority sets Authority field to given value.


### GetUserInfoEndpoint

`func (o *IdentityProviderProviderTypeParameterRequest) GetUserInfoEndpoint() string`

GetUserInfoEndpoint returns the UserInfoEndpoint field if non-nil, zero value otherwise.

### GetUserInfoEndpointOk

`func (o *IdentityProviderProviderTypeParameterRequest) GetUserInfoEndpointOk() (*string, bool)`

GetUserInfoEndpointOk returns a tuple with the UserInfoEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInfoEndpoint

`func (o *IdentityProviderProviderTypeParameterRequest) SetUserInfoEndpoint(v string)`

SetUserInfoEndpoint sets UserInfoEndpoint field to given value.

### HasUserInfoEndpoint

`func (o *IdentityProviderProviderTypeParameterRequest) HasUserInfoEndpoint() bool`

HasUserInfoEndpoint returns a boolean if a field has been set.

### SetUserInfoEndpointNil

`func (o *IdentityProviderProviderTypeParameterRequest) SetUserInfoEndpointNil(b bool)`

 SetUserInfoEndpointNil sets the value for UserInfoEndpoint to be an explicit nil

### UnsetUserInfoEndpoint
`func (o *IdentityProviderProviderTypeParameterRequest) UnsetUserInfoEndpoint()`

UnsetUserInfoEndpoint ensures that no value is present for UserInfoEndpoint, not even an explicit nil
### GetAuth0ApiUrl

`func (o *IdentityProviderProviderTypeParameterRequest) GetAuth0ApiUrl() string`

GetAuth0ApiUrl returns the Auth0ApiUrl field if non-nil, zero value otherwise.

### GetAuth0ApiUrlOk

`func (o *IdentityProviderProviderTypeParameterRequest) GetAuth0ApiUrlOk() (*string, bool)`

GetAuth0ApiUrlOk returns a tuple with the Auth0ApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth0ApiUrl

`func (o *IdentityProviderProviderTypeParameterRequest) SetAuth0ApiUrl(v string)`

SetAuth0ApiUrl sets Auth0ApiUrl field to given value.

### HasAuth0ApiUrl

`func (o *IdentityProviderProviderTypeParameterRequest) HasAuth0ApiUrl() bool`

HasAuth0ApiUrl returns a boolean if a field has been set.

### SetAuth0ApiUrlNil

`func (o *IdentityProviderProviderTypeParameterRequest) SetAuth0ApiUrlNil(b bool)`

 SetAuth0ApiUrlNil sets the value for Auth0ApiUrl to be an explicit nil

### UnsetAuth0ApiUrl
`func (o *IdentityProviderProviderTypeParameterRequest) UnsetAuth0ApiUrl()`

UnsetAuth0ApiUrl ensures that no value is present for Auth0ApiUrl, not even an explicit nil
### GetSignOutUrl

`func (o *IdentityProviderProviderTypeParameterRequest) GetSignOutUrl() string`

GetSignOutUrl returns the SignOutUrl field if non-nil, zero value otherwise.

### GetSignOutUrlOk

`func (o *IdentityProviderProviderTypeParameterRequest) GetSignOutUrlOk() (*string, bool)`

GetSignOutUrlOk returns a tuple with the SignOutUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignOutUrl

`func (o *IdentityProviderProviderTypeParameterRequest) SetSignOutUrl(v string)`

SetSignOutUrl sets SignOutUrl field to given value.

### HasSignOutUrl

`func (o *IdentityProviderProviderTypeParameterRequest) HasSignOutUrl() bool`

HasSignOutUrl returns a boolean if a field has been set.

### SetSignOutUrlNil

`func (o *IdentityProviderProviderTypeParameterRequest) SetSignOutUrlNil(b bool)`

 SetSignOutUrlNil sets the value for SignOutUrl to be an explicit nil

### UnsetSignOutUrl
`func (o *IdentityProviderProviderTypeParameterRequest) UnsetSignOutUrl()`

UnsetSignOutUrl ensures that no value is present for SignOutUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


