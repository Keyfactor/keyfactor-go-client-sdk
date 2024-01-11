# KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**AuthenticationScheme** | Pointer to **NullableString** |  | [optional] 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**TypeId** | Pointer to **string** |  | [optional] 
**Parameters** | Pointer to [**[]KeyfactorWebKeyfactorApiModelsIdentityProviderProviderTypeParameterValueResponse**](KeyfactorWebKeyfactorApiModelsIdentityProviderProviderTypeParameterValueResponse.md) |  | [optional] 

## Methods

### NewKeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderResponse

`func NewKeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderResponse() *KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderResponse`

NewKeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderResponse instantiates a new KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderResponseWithDefaults

`func NewKeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderResponseWithDefaults() *KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderResponse`

NewKeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderResponseWithDefaults instantiates a new KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAuthenticationScheme

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderResponse) GetAuthenticationScheme() string`

GetAuthenticationScheme returns the AuthenticationScheme field if non-nil, zero value otherwise.

### GetAuthenticationSchemeOk

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderResponse) GetAuthenticationSchemeOk() (*string, bool)`

GetAuthenticationSchemeOk returns a tuple with the AuthenticationScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationScheme

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderResponse) SetAuthenticationScheme(v string)`

SetAuthenticationScheme sets AuthenticationScheme field to given value.

### HasAuthenticationScheme

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderResponse) HasAuthenticationScheme() bool`

HasAuthenticationScheme returns a boolean if a field has been set.

### SetAuthenticationSchemeNil

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderResponse) SetAuthenticationSchemeNil(b bool)`

 SetAuthenticationSchemeNil sets the value for AuthenticationScheme to be an explicit nil

### UnsetAuthenticationScheme
`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderResponse) UnsetAuthenticationScheme()`

UnsetAuthenticationScheme ensures that no value is present for AuthenticationScheme, not even an explicit nil
### GetDisplayName

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderResponse) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderResponse) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderResponse) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderResponse) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderResponse) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderResponse) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetTypeId

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderResponse) GetTypeId() string`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderResponse) GetTypeIdOk() (*string, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderResponse) SetTypeId(v string)`

SetTypeId sets TypeId field to given value.

### HasTypeId

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderResponse) HasTypeId() bool`

HasTypeId returns a boolean if a field has been set.

### GetParameters

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderResponse) GetParameters() []KeyfactorWebKeyfactorApiModelsIdentityProviderProviderTypeParameterValueResponse`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderResponse) GetParametersOk() (*[]KeyfactorWebKeyfactorApiModelsIdentityProviderProviderTypeParameterValueResponse, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderResponse) SetParameters(v []KeyfactorWebKeyfactorApiModelsIdentityProviderProviderTypeParameterValueResponse)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderResponse) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### SetParametersNil

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderResponse) SetParametersNil(b bool)`

 SetParametersNil sets the value for Parameters to be an explicit nil

### UnsetParameters
`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderResponse) UnsetParameters()`

UnsetParameters ensures that no value is present for Parameters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


