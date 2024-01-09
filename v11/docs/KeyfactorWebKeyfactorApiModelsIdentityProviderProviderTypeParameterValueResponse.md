# KeyfactorWebKeyfactorApiModelsIdentityProviderProviderTypeParameterValueResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**Required** | Pointer to **bool** |  | [optional] 
**DataType** | Pointer to [**CSSCMSDataModelEnumsIdentityProviderDataType**](CSSCMSDataModelEnumsIdentityProviderDataType.md) |  | [optional] 
**Value** | Pointer to **NullableString** |  | [optional] 
**SecretValue** | Pointer to [**CSSCMSDataModelModelsKeyfactorAPISecret**](CSSCMSDataModelModelsKeyfactorAPISecret.md) |  | [optional] 

## Methods

### NewKeyfactorWebKeyfactorApiModelsIdentityProviderProviderTypeParameterValueResponse

`func NewKeyfactorWebKeyfactorApiModelsIdentityProviderProviderTypeParameterValueResponse() *KeyfactorWebKeyfactorApiModelsIdentityProviderProviderTypeParameterValueResponse`

NewKeyfactorWebKeyfactorApiModelsIdentityProviderProviderTypeParameterValueResponse instantiates a new KeyfactorWebKeyfactorApiModelsIdentityProviderProviderTypeParameterValueResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorWebKeyfactorApiModelsIdentityProviderProviderTypeParameterValueResponseWithDefaults

`func NewKeyfactorWebKeyfactorApiModelsIdentityProviderProviderTypeParameterValueResponseWithDefaults() *KeyfactorWebKeyfactorApiModelsIdentityProviderProviderTypeParameterValueResponse`

NewKeyfactorWebKeyfactorApiModelsIdentityProviderProviderTypeParameterValueResponseWithDefaults instantiates a new KeyfactorWebKeyfactorApiModelsIdentityProviderProviderTypeParameterValueResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderProviderTypeParameterValueResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderProviderTypeParameterValueResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderProviderTypeParameterValueResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderProviderTypeParameterValueResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderProviderTypeParameterValueResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderProviderTypeParameterValueResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderProviderTypeParameterValueResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderProviderTypeParameterValueResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderProviderTypeParameterValueResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderProviderTypeParameterValueResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDisplayName

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderProviderTypeParameterValueResponse) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderProviderTypeParameterValueResponse) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderProviderTypeParameterValueResponse) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderProviderTypeParameterValueResponse) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderProviderTypeParameterValueResponse) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderProviderTypeParameterValueResponse) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetRequired

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderProviderTypeParameterValueResponse) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderProviderTypeParameterValueResponse) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderProviderTypeParameterValueResponse) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderProviderTypeParameterValueResponse) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetDataType

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderProviderTypeParameterValueResponse) GetDataType() CSSCMSDataModelEnumsIdentityProviderDataType`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderProviderTypeParameterValueResponse) GetDataTypeOk() (*CSSCMSDataModelEnumsIdentityProviderDataType, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderProviderTypeParameterValueResponse) SetDataType(v CSSCMSDataModelEnumsIdentityProviderDataType)`

SetDataType sets DataType field to given value.

### HasDataType

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderProviderTypeParameterValueResponse) HasDataType() bool`

HasDataType returns a boolean if a field has been set.

### GetValue

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderProviderTypeParameterValueResponse) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderProviderTypeParameterValueResponse) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderProviderTypeParameterValueResponse) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderProviderTypeParameterValueResponse) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderProviderTypeParameterValueResponse) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderProviderTypeParameterValueResponse) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetSecretValue

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderProviderTypeParameterValueResponse) GetSecretValue() CSSCMSDataModelModelsKeyfactorAPISecret`

GetSecretValue returns the SecretValue field if non-nil, zero value otherwise.

### GetSecretValueOk

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderProviderTypeParameterValueResponse) GetSecretValueOk() (*CSSCMSDataModelModelsKeyfactorAPISecret, bool)`

GetSecretValueOk returns a tuple with the SecretValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretValue

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderProviderTypeParameterValueResponse) SetSecretValue(v CSSCMSDataModelModelsKeyfactorAPISecret)`

SetSecretValue sets SecretValue field to given value.

### HasSecretValue

`func (o *KeyfactorWebKeyfactorApiModelsIdentityProviderProviderTypeParameterValueResponse) HasSecretValue() bool`

HasSecretValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


