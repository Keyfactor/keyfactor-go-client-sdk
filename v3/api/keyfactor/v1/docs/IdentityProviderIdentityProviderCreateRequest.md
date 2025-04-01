# IdentityProviderIdentityProviderCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationScheme** | **string** |  | 
**DisplayName** | **string** |  | 
**ProviderType** | **string** |  | 
**PermissionSetId** | Pointer to **NullableString** |  | [optional] 
**Parameters** | Pointer to [**IdentityProviderProviderTypeParameterRequest**](IdentityProviderProviderTypeParameterRequest.md) |  | [optional] 

## Methods

### NewIdentityProviderIdentityProviderCreateRequest

`func NewIdentityProviderIdentityProviderCreateRequest(authenticationScheme string, displayName string, providerType string, ) *IdentityProviderIdentityProviderCreateRequest`

NewIdentityProviderIdentityProviderCreateRequest instantiates a new IdentityProviderIdentityProviderCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityProviderIdentityProviderCreateRequestWithDefaults

`func NewIdentityProviderIdentityProviderCreateRequestWithDefaults() *IdentityProviderIdentityProviderCreateRequest`

NewIdentityProviderIdentityProviderCreateRequestWithDefaults instantiates a new IdentityProviderIdentityProviderCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationScheme

`func (o *IdentityProviderIdentityProviderCreateRequest) GetAuthenticationScheme() string`

GetAuthenticationScheme returns the AuthenticationScheme field if non-nil, zero value otherwise.

### GetAuthenticationSchemeOk

`func (o *IdentityProviderIdentityProviderCreateRequest) GetAuthenticationSchemeOk() (*string, bool)`

GetAuthenticationSchemeOk returns a tuple with the AuthenticationScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationScheme

`func (o *IdentityProviderIdentityProviderCreateRequest) SetAuthenticationScheme(v string)`

SetAuthenticationScheme sets AuthenticationScheme field to given value.


### GetDisplayName

`func (o *IdentityProviderIdentityProviderCreateRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *IdentityProviderIdentityProviderCreateRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *IdentityProviderIdentityProviderCreateRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetProviderType

`func (o *IdentityProviderIdentityProviderCreateRequest) GetProviderType() string`

GetProviderType returns the ProviderType field if non-nil, zero value otherwise.

### GetProviderTypeOk

`func (o *IdentityProviderIdentityProviderCreateRequest) GetProviderTypeOk() (*string, bool)`

GetProviderTypeOk returns a tuple with the ProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderType

`func (o *IdentityProviderIdentityProviderCreateRequest) SetProviderType(v string)`

SetProviderType sets ProviderType field to given value.


### GetPermissionSetId

`func (o *IdentityProviderIdentityProviderCreateRequest) GetPermissionSetId() string`

GetPermissionSetId returns the PermissionSetId field if non-nil, zero value otherwise.

### GetPermissionSetIdOk

`func (o *IdentityProviderIdentityProviderCreateRequest) GetPermissionSetIdOk() (*string, bool)`

GetPermissionSetIdOk returns a tuple with the PermissionSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionSetId

`func (o *IdentityProviderIdentityProviderCreateRequest) SetPermissionSetId(v string)`

SetPermissionSetId sets PermissionSetId field to given value.

### HasPermissionSetId

`func (o *IdentityProviderIdentityProviderCreateRequest) HasPermissionSetId() bool`

HasPermissionSetId returns a boolean if a field has been set.

### SetPermissionSetIdNil

`func (o *IdentityProviderIdentityProviderCreateRequest) SetPermissionSetIdNil(b bool)`

 SetPermissionSetIdNil sets the value for PermissionSetId to be an explicit nil

### UnsetPermissionSetId
`func (o *IdentityProviderIdentityProviderCreateRequest) UnsetPermissionSetId()`

UnsetPermissionSetId ensures that no value is present for PermissionSetId, not even an explicit nil
### GetParameters

`func (o *IdentityProviderIdentityProviderCreateRequest) GetParameters() IdentityProviderProviderTypeParameterRequest`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *IdentityProviderIdentityProviderCreateRequest) GetParametersOk() (*IdentityProviderProviderTypeParameterRequest, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *IdentityProviderIdentityProviderCreateRequest) SetParameters(v IdentityProviderProviderTypeParameterRequest)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *IdentityProviderIdentityProviderCreateRequest) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


