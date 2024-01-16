# ModelsProviderType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**ProviderTypeParams** | Pointer to [**[]ModelsProviderTypeParam**](ModelsProviderTypeParam.md) |  | [optional] 

## Methods

### NewModelsProviderType

`func NewModelsProviderType() *ModelsProviderType`

NewModelsProviderType instantiates a new ModelsProviderType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsProviderTypeWithDefaults

`func NewModelsProviderTypeWithDefaults() *ModelsProviderType`

NewModelsProviderTypeWithDefaults instantiates a new ModelsProviderType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelsProviderType) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsProviderType) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsProviderType) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsProviderType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ModelsProviderType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelsProviderType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelsProviderType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelsProviderType) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ModelsProviderType) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ModelsProviderType) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetProviderTypeParams

`func (o *ModelsProviderType) GetProviderTypeParams() []ModelsProviderTypeParam`

GetProviderTypeParams returns the ProviderTypeParams field if non-nil, zero value otherwise.

### GetProviderTypeParamsOk

`func (o *ModelsProviderType) GetProviderTypeParamsOk() (*[]ModelsProviderTypeParam, bool)`

GetProviderTypeParamsOk returns a tuple with the ProviderTypeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderTypeParams

`func (o *ModelsProviderType) SetProviderTypeParams(v []ModelsProviderTypeParam)`

SetProviderTypeParams sets ProviderTypeParams field to given value.

### HasProviderTypeParams

`func (o *ModelsProviderType) HasProviderTypeParams() bool`

HasProviderTypeParams returns a boolean if a field has been set.

### SetProviderTypeParamsNil

`func (o *ModelsProviderType) SetProviderTypeParamsNil(b bool)`

 SetProviderTypeParamsNil sets the value for ProviderTypeParams to be an explicit nil

### UnsetProviderTypeParams
`func (o *ModelsProviderType) UnsetProviderTypeParams()`

UnsetProviderTypeParams ensures that no value is present for ProviderTypeParams, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


