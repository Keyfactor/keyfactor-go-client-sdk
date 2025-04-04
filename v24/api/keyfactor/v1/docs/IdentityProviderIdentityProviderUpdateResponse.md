# IdentityProviderIdentityProviderUpdateResponse

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

### NewIdentityProviderIdentityProviderUpdateResponse

`func NewIdentityProviderIdentityProviderUpdateResponse() *IdentityProviderIdentityProviderUpdateResponse`

NewIdentityProviderIdentityProviderUpdateResponse instantiates a new IdentityProviderIdentityProviderUpdateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityProviderIdentityProviderUpdateResponseWithDefaults

`func NewIdentityProviderIdentityProviderUpdateResponseWithDefaults() *IdentityProviderIdentityProviderUpdateResponse`

NewIdentityProviderIdentityProviderUpdateResponseWithDefaults instantiates a new IdentityProviderIdentityProviderUpdateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IdentityProviderIdentityProviderUpdateResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdentityProviderIdentityProviderUpdateResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdentityProviderIdentityProviderUpdateResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IdentityProviderIdentityProviderUpdateResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAuthenticationScheme

`func (o *IdentityProviderIdentityProviderUpdateResponse) GetAuthenticationScheme() string`

GetAuthenticationScheme returns the AuthenticationScheme field if non-nil, zero value otherwise.

### GetAuthenticationSchemeOk

`func (o *IdentityProviderIdentityProviderUpdateResponse) GetAuthenticationSchemeOk() (*string, bool)`

GetAuthenticationSchemeOk returns a tuple with the AuthenticationScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationScheme

`func (o *IdentityProviderIdentityProviderUpdateResponse) SetAuthenticationScheme(v string)`

SetAuthenticationScheme sets AuthenticationScheme field to given value.

### HasAuthenticationScheme

`func (o *IdentityProviderIdentityProviderUpdateResponse) HasAuthenticationScheme() bool`

HasAuthenticationScheme returns a boolean if a field has been set.

### SetAuthenticationSchemeNil

`func (o *IdentityProviderIdentityProviderUpdateResponse) SetAuthenticationSchemeNil(b bool)`

 SetAuthenticationSchemeNil sets the value for AuthenticationScheme to be an explicit nil

### UnsetAuthenticationScheme
`func (o *IdentityProviderIdentityProviderUpdateResponse) UnsetAuthenticationScheme()`

UnsetAuthenticationScheme ensures that no value is present for AuthenticationScheme, not even an explicit nil
### GetDisplayName

`func (o *IdentityProviderIdentityProviderUpdateResponse) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *IdentityProviderIdentityProviderUpdateResponse) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *IdentityProviderIdentityProviderUpdateResponse) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *IdentityProviderIdentityProviderUpdateResponse) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *IdentityProviderIdentityProviderUpdateResponse) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *IdentityProviderIdentityProviderUpdateResponse) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetTypeId

`func (o *IdentityProviderIdentityProviderUpdateResponse) GetTypeId() string`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *IdentityProviderIdentityProviderUpdateResponse) GetTypeIdOk() (*string, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *IdentityProviderIdentityProviderUpdateResponse) SetTypeId(v string)`

SetTypeId sets TypeId field to given value.

### HasTypeId

`func (o *IdentityProviderIdentityProviderUpdateResponse) HasTypeId() bool`

HasTypeId returns a boolean if a field has been set.

### GetParameters

`func (o *IdentityProviderIdentityProviderUpdateResponse) GetParameters() []IdentityProviderProviderTypeParameterValueResponse`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *IdentityProviderIdentityProviderUpdateResponse) GetParametersOk() (*[]IdentityProviderProviderTypeParameterValueResponse, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *IdentityProviderIdentityProviderUpdateResponse) SetParameters(v []IdentityProviderProviderTypeParameterValueResponse)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *IdentityProviderIdentityProviderUpdateResponse) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### SetParametersNil

`func (o *IdentityProviderIdentityProviderUpdateResponse) SetParametersNil(b bool)`

 SetParametersNil sets the value for Parameters to be an explicit nil

### UnsetParameters
`func (o *IdentityProviderIdentityProviderUpdateResponse) UnsetParameters()`

UnsetParameters ensures that no value is present for Parameters, not even an explicit nil
### GetPermissionSetId

`func (o *IdentityProviderIdentityProviderUpdateResponse) GetPermissionSetId() string`

GetPermissionSetId returns the PermissionSetId field if non-nil, zero value otherwise.

### GetPermissionSetIdOk

`func (o *IdentityProviderIdentityProviderUpdateResponse) GetPermissionSetIdOk() (*string, bool)`

GetPermissionSetIdOk returns a tuple with the PermissionSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionSetId

`func (o *IdentityProviderIdentityProviderUpdateResponse) SetPermissionSetId(v string)`

SetPermissionSetId sets PermissionSetId field to given value.

### HasPermissionSetId

`func (o *IdentityProviderIdentityProviderUpdateResponse) HasPermissionSetId() bool`

HasPermissionSetId returns a boolean if a field has been set.

### GetAuthenticationEnabled

`func (o *IdentityProviderIdentityProviderUpdateResponse) GetAuthenticationEnabled() bool`

GetAuthenticationEnabled returns the AuthenticationEnabled field if non-nil, zero value otherwise.

### GetAuthenticationEnabledOk

`func (o *IdentityProviderIdentityProviderUpdateResponse) GetAuthenticationEnabledOk() (*bool, bool)`

GetAuthenticationEnabledOk returns a tuple with the AuthenticationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationEnabled

`func (o *IdentityProviderIdentityProviderUpdateResponse) SetAuthenticationEnabled(v bool)`

SetAuthenticationEnabled sets AuthenticationEnabled field to given value.

### HasAuthenticationEnabled

`func (o *IdentityProviderIdentityProviderUpdateResponse) HasAuthenticationEnabled() bool`

HasAuthenticationEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


