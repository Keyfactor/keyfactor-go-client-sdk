# IdentityProviderIdentityProviderCreateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**AuthenticationScheme** | Pointer to **NullableString** |  | [optional] 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**TypeId** | Pointer to **string** |  | [optional] 
**Parameters** | Pointer to [**[]IdentityProviderProviderTypeParameterValueResponse**](IdentityProviderProviderTypeParameterValueResponse.md) |  | [optional] 
**PermissionSetId** | Pointer to **string** |  | [optional] 
**AuthenticationEnabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewIdentityProviderIdentityProviderCreateResponse

`func NewIdentityProviderIdentityProviderCreateResponse() *IdentityProviderIdentityProviderCreateResponse`

NewIdentityProviderIdentityProviderCreateResponse instantiates a new IdentityProviderIdentityProviderCreateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityProviderIdentityProviderCreateResponseWithDefaults

`func NewIdentityProviderIdentityProviderCreateResponseWithDefaults() *IdentityProviderIdentityProviderCreateResponse`

NewIdentityProviderIdentityProviderCreateResponseWithDefaults instantiates a new IdentityProviderIdentityProviderCreateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IdentityProviderIdentityProviderCreateResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdentityProviderIdentityProviderCreateResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdentityProviderIdentityProviderCreateResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IdentityProviderIdentityProviderCreateResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAuthenticationScheme

`func (o *IdentityProviderIdentityProviderCreateResponse) GetAuthenticationScheme() string`

GetAuthenticationScheme returns the AuthenticationScheme field if non-nil, zero value otherwise.

### GetAuthenticationSchemeOk

`func (o *IdentityProviderIdentityProviderCreateResponse) GetAuthenticationSchemeOk() (*string, bool)`

GetAuthenticationSchemeOk returns a tuple with the AuthenticationScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationScheme

`func (o *IdentityProviderIdentityProviderCreateResponse) SetAuthenticationScheme(v string)`

SetAuthenticationScheme sets AuthenticationScheme field to given value.

### HasAuthenticationScheme

`func (o *IdentityProviderIdentityProviderCreateResponse) HasAuthenticationScheme() bool`

HasAuthenticationScheme returns a boolean if a field has been set.

### SetAuthenticationSchemeNil

`func (o *IdentityProviderIdentityProviderCreateResponse) SetAuthenticationSchemeNil(b bool)`

 SetAuthenticationSchemeNil sets the value for AuthenticationScheme to be an explicit nil

### UnsetAuthenticationScheme
`func (o *IdentityProviderIdentityProviderCreateResponse) UnsetAuthenticationScheme()`

UnsetAuthenticationScheme ensures that no value is present for AuthenticationScheme, not even an explicit nil
### GetDisplayName

`func (o *IdentityProviderIdentityProviderCreateResponse) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *IdentityProviderIdentityProviderCreateResponse) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *IdentityProviderIdentityProviderCreateResponse) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *IdentityProviderIdentityProviderCreateResponse) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *IdentityProviderIdentityProviderCreateResponse) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *IdentityProviderIdentityProviderCreateResponse) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetTypeId

`func (o *IdentityProviderIdentityProviderCreateResponse) GetTypeId() string`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *IdentityProviderIdentityProviderCreateResponse) GetTypeIdOk() (*string, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *IdentityProviderIdentityProviderCreateResponse) SetTypeId(v string)`

SetTypeId sets TypeId field to given value.

### HasTypeId

`func (o *IdentityProviderIdentityProviderCreateResponse) HasTypeId() bool`

HasTypeId returns a boolean if a field has been set.

### GetParameters

`func (o *IdentityProviderIdentityProviderCreateResponse) GetParameters() []IdentityProviderProviderTypeParameterValueResponse`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *IdentityProviderIdentityProviderCreateResponse) GetParametersOk() (*[]IdentityProviderProviderTypeParameterValueResponse, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *IdentityProviderIdentityProviderCreateResponse) SetParameters(v []IdentityProviderProviderTypeParameterValueResponse)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *IdentityProviderIdentityProviderCreateResponse) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### SetParametersNil

`func (o *IdentityProviderIdentityProviderCreateResponse) SetParametersNil(b bool)`

 SetParametersNil sets the value for Parameters to be an explicit nil

### UnsetParameters
`func (o *IdentityProviderIdentityProviderCreateResponse) UnsetParameters()`

UnsetParameters ensures that no value is present for Parameters, not even an explicit nil
### GetPermissionSetId

`func (o *IdentityProviderIdentityProviderCreateResponse) GetPermissionSetId() string`

GetPermissionSetId returns the PermissionSetId field if non-nil, zero value otherwise.

### GetPermissionSetIdOk

`func (o *IdentityProviderIdentityProviderCreateResponse) GetPermissionSetIdOk() (*string, bool)`

GetPermissionSetIdOk returns a tuple with the PermissionSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionSetId

`func (o *IdentityProviderIdentityProviderCreateResponse) SetPermissionSetId(v string)`

SetPermissionSetId sets PermissionSetId field to given value.

### HasPermissionSetId

`func (o *IdentityProviderIdentityProviderCreateResponse) HasPermissionSetId() bool`

HasPermissionSetId returns a boolean if a field has been set.

### GetAuthenticationEnabled

`func (o *IdentityProviderIdentityProviderCreateResponse) GetAuthenticationEnabled() bool`

GetAuthenticationEnabled returns the AuthenticationEnabled field if non-nil, zero value otherwise.

### GetAuthenticationEnabledOk

`func (o *IdentityProviderIdentityProviderCreateResponse) GetAuthenticationEnabledOk() (*bool, bool)`

GetAuthenticationEnabledOk returns a tuple with the AuthenticationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationEnabled

`func (o *IdentityProviderIdentityProviderCreateResponse) SetAuthenticationEnabled(v bool)`

SetAuthenticationEnabled sets AuthenticationEnabled field to given value.

### HasAuthenticationEnabled

`func (o *IdentityProviderIdentityProviderCreateResponse) HasAuthenticationEnabled() bool`

HasAuthenticationEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


