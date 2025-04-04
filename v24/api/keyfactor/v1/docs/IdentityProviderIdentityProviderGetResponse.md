# IdentityProviderIdentityProviderGetResponse

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

### NewIdentityProviderIdentityProviderGetResponse

`func NewIdentityProviderIdentityProviderGetResponse() *IdentityProviderIdentityProviderGetResponse`

NewIdentityProviderIdentityProviderGetResponse instantiates a new IdentityProviderIdentityProviderGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityProviderIdentityProviderGetResponseWithDefaults

`func NewIdentityProviderIdentityProviderGetResponseWithDefaults() *IdentityProviderIdentityProviderGetResponse`

NewIdentityProviderIdentityProviderGetResponseWithDefaults instantiates a new IdentityProviderIdentityProviderGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IdentityProviderIdentityProviderGetResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdentityProviderIdentityProviderGetResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdentityProviderIdentityProviderGetResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IdentityProviderIdentityProviderGetResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAuthenticationScheme

`func (o *IdentityProviderIdentityProviderGetResponse) GetAuthenticationScheme() string`

GetAuthenticationScheme returns the AuthenticationScheme field if non-nil, zero value otherwise.

### GetAuthenticationSchemeOk

`func (o *IdentityProviderIdentityProviderGetResponse) GetAuthenticationSchemeOk() (*string, bool)`

GetAuthenticationSchemeOk returns a tuple with the AuthenticationScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationScheme

`func (o *IdentityProviderIdentityProviderGetResponse) SetAuthenticationScheme(v string)`

SetAuthenticationScheme sets AuthenticationScheme field to given value.

### HasAuthenticationScheme

`func (o *IdentityProviderIdentityProviderGetResponse) HasAuthenticationScheme() bool`

HasAuthenticationScheme returns a boolean if a field has been set.

### SetAuthenticationSchemeNil

`func (o *IdentityProviderIdentityProviderGetResponse) SetAuthenticationSchemeNil(b bool)`

 SetAuthenticationSchemeNil sets the value for AuthenticationScheme to be an explicit nil

### UnsetAuthenticationScheme
`func (o *IdentityProviderIdentityProviderGetResponse) UnsetAuthenticationScheme()`

UnsetAuthenticationScheme ensures that no value is present for AuthenticationScheme, not even an explicit nil
### GetDisplayName

`func (o *IdentityProviderIdentityProviderGetResponse) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *IdentityProviderIdentityProviderGetResponse) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *IdentityProviderIdentityProviderGetResponse) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *IdentityProviderIdentityProviderGetResponse) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *IdentityProviderIdentityProviderGetResponse) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *IdentityProviderIdentityProviderGetResponse) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetTypeId

`func (o *IdentityProviderIdentityProviderGetResponse) GetTypeId() string`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *IdentityProviderIdentityProviderGetResponse) GetTypeIdOk() (*string, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *IdentityProviderIdentityProviderGetResponse) SetTypeId(v string)`

SetTypeId sets TypeId field to given value.

### HasTypeId

`func (o *IdentityProviderIdentityProviderGetResponse) HasTypeId() bool`

HasTypeId returns a boolean if a field has been set.

### GetParameters

`func (o *IdentityProviderIdentityProviderGetResponse) GetParameters() []IdentityProviderProviderTypeParameterValueResponse`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *IdentityProviderIdentityProviderGetResponse) GetParametersOk() (*[]IdentityProviderProviderTypeParameterValueResponse, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *IdentityProviderIdentityProviderGetResponse) SetParameters(v []IdentityProviderProviderTypeParameterValueResponse)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *IdentityProviderIdentityProviderGetResponse) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### SetParametersNil

`func (o *IdentityProviderIdentityProviderGetResponse) SetParametersNil(b bool)`

 SetParametersNil sets the value for Parameters to be an explicit nil

### UnsetParameters
`func (o *IdentityProviderIdentityProviderGetResponse) UnsetParameters()`

UnsetParameters ensures that no value is present for Parameters, not even an explicit nil
### GetPermissionSetId

`func (o *IdentityProviderIdentityProviderGetResponse) GetPermissionSetId() string`

GetPermissionSetId returns the PermissionSetId field if non-nil, zero value otherwise.

### GetPermissionSetIdOk

`func (o *IdentityProviderIdentityProviderGetResponse) GetPermissionSetIdOk() (*string, bool)`

GetPermissionSetIdOk returns a tuple with the PermissionSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionSetId

`func (o *IdentityProviderIdentityProviderGetResponse) SetPermissionSetId(v string)`

SetPermissionSetId sets PermissionSetId field to given value.

### HasPermissionSetId

`func (o *IdentityProviderIdentityProviderGetResponse) HasPermissionSetId() bool`

HasPermissionSetId returns a boolean if a field has been set.

### GetAuthenticationEnabled

`func (o *IdentityProviderIdentityProviderGetResponse) GetAuthenticationEnabled() bool`

GetAuthenticationEnabled returns the AuthenticationEnabled field if non-nil, zero value otherwise.

### GetAuthenticationEnabledOk

`func (o *IdentityProviderIdentityProviderGetResponse) GetAuthenticationEnabledOk() (*bool, bool)`

GetAuthenticationEnabledOk returns a tuple with the AuthenticationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationEnabled

`func (o *IdentityProviderIdentityProviderGetResponse) SetAuthenticationEnabled(v bool)`

SetAuthenticationEnabled sets AuthenticationEnabled field to given value.

### HasAuthenticationEnabled

`func (o *IdentityProviderIdentityProviderGetResponse) HasAuthenticationEnabled() bool`

HasAuthenticationEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


