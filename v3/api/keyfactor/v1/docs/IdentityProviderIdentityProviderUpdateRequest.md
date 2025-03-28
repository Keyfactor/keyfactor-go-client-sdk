# IdentityProviderIdentityProviderUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationScheme** | **string** |  | 
**DisplayName** | **string** |  | 
**Parameters** | Pointer to [**IdentityProviderProviderTypeParameterRequest**](IdentityProviderProviderTypeParameterRequest.md) |  | [optional] 
**PermissionSetId** | Pointer to **NullableString** |  | [optional] 
**AuthenticationEnabled** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewIdentityProviderIdentityProviderUpdateRequest

`func NewIdentityProviderIdentityProviderUpdateRequest(authenticationScheme string, displayName string, ) *IdentityProviderIdentityProviderUpdateRequest`

NewIdentityProviderIdentityProviderUpdateRequest instantiates a new IdentityProviderIdentityProviderUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityProviderIdentityProviderUpdateRequestWithDefaults

`func NewIdentityProviderIdentityProviderUpdateRequestWithDefaults() *IdentityProviderIdentityProviderUpdateRequest`

NewIdentityProviderIdentityProviderUpdateRequestWithDefaults instantiates a new IdentityProviderIdentityProviderUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationScheme

`func (o *IdentityProviderIdentityProviderUpdateRequest) GetAuthenticationScheme() string`

GetAuthenticationScheme returns the AuthenticationScheme field if non-nil, zero value otherwise.

### GetAuthenticationSchemeOk

`func (o *IdentityProviderIdentityProviderUpdateRequest) GetAuthenticationSchemeOk() (*string, bool)`

GetAuthenticationSchemeOk returns a tuple with the AuthenticationScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationScheme

`func (o *IdentityProviderIdentityProviderUpdateRequest) SetAuthenticationScheme(v string)`

SetAuthenticationScheme sets AuthenticationScheme field to given value.


### GetDisplayName

`func (o *IdentityProviderIdentityProviderUpdateRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *IdentityProviderIdentityProviderUpdateRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *IdentityProviderIdentityProviderUpdateRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetParameters

`func (o *IdentityProviderIdentityProviderUpdateRequest) GetParameters() IdentityProviderProviderTypeParameterRequest`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *IdentityProviderIdentityProviderUpdateRequest) GetParametersOk() (*IdentityProviderProviderTypeParameterRequest, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *IdentityProviderIdentityProviderUpdateRequest) SetParameters(v IdentityProviderProviderTypeParameterRequest)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *IdentityProviderIdentityProviderUpdateRequest) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetPermissionSetId

`func (o *IdentityProviderIdentityProviderUpdateRequest) GetPermissionSetId() string`

GetPermissionSetId returns the PermissionSetId field if non-nil, zero value otherwise.

### GetPermissionSetIdOk

`func (o *IdentityProviderIdentityProviderUpdateRequest) GetPermissionSetIdOk() (*string, bool)`

GetPermissionSetIdOk returns a tuple with the PermissionSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionSetId

`func (o *IdentityProviderIdentityProviderUpdateRequest) SetPermissionSetId(v string)`

SetPermissionSetId sets PermissionSetId field to given value.

### HasPermissionSetId

`func (o *IdentityProviderIdentityProviderUpdateRequest) HasPermissionSetId() bool`

HasPermissionSetId returns a boolean if a field has been set.

### SetPermissionSetIdNil

`func (o *IdentityProviderIdentityProviderUpdateRequest) SetPermissionSetIdNil(b bool)`

 SetPermissionSetIdNil sets the value for PermissionSetId to be an explicit nil

### UnsetPermissionSetId
`func (o *IdentityProviderIdentityProviderUpdateRequest) UnsetPermissionSetId()`

UnsetPermissionSetId ensures that no value is present for PermissionSetId, not even an explicit nil
### GetAuthenticationEnabled

`func (o *IdentityProviderIdentityProviderUpdateRequest) GetAuthenticationEnabled() bool`

GetAuthenticationEnabled returns the AuthenticationEnabled field if non-nil, zero value otherwise.

### GetAuthenticationEnabledOk

`func (o *IdentityProviderIdentityProviderUpdateRequest) GetAuthenticationEnabledOk() (*bool, bool)`

GetAuthenticationEnabledOk returns a tuple with the AuthenticationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationEnabled

`func (o *IdentityProviderIdentityProviderUpdateRequest) SetAuthenticationEnabled(v bool)`

SetAuthenticationEnabled sets AuthenticationEnabled field to given value.

### HasAuthenticationEnabled

`func (o *IdentityProviderIdentityProviderUpdateRequest) HasAuthenticationEnabled() bool`

HasAuthenticationEnabled returns a boolean if a field has been set.

### SetAuthenticationEnabledNil

`func (o *IdentityProviderIdentityProviderUpdateRequest) SetAuthenticationEnabledNil(b bool)`

 SetAuthenticationEnabledNil sets the value for AuthenticationEnabled to be an explicit nil

### UnsetAuthenticationEnabled
`func (o *IdentityProviderIdentityProviderUpdateRequest) UnsetAuthenticationEnabled()`

UnsetAuthenticationEnabled ensures that no value is present for AuthenticationEnabled, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


