# ModelsProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  | 
**Area** | Pointer to **int32** |  | [optional] 
**ProviderType** | [**ModelsProviderType**](ModelsProviderType.md) |  | 
**ProviderTypeParamValues** | Pointer to [**[]ModelsPamProviderTypeParamValue**](ModelsPamProviderTypeParamValue.md) |  | [optional] 
**SecuredAreaId** | Pointer to **NullableInt32** |  | [optional] 
**Remote** | Pointer to **bool** |  | [optional] 

## Methods

### NewModelsProvider

`func NewModelsProvider(name string, providerType ModelsProviderType, ) *ModelsProvider`

NewModelsProvider instantiates a new ModelsProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsProviderWithDefaults

`func NewModelsProviderWithDefaults() *ModelsProvider`

NewModelsProviderWithDefaults instantiates a new ModelsProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelsProvider) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsProvider) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsProvider) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsProvider) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ModelsProvider) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelsProvider) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelsProvider) SetName(v string)`

SetName sets Name field to given value.


### GetArea

`func (o *ModelsProvider) GetArea() int32`

GetArea returns the Area field if non-nil, zero value otherwise.

### GetAreaOk

`func (o *ModelsProvider) GetAreaOk() (*int32, bool)`

GetAreaOk returns a tuple with the Area field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArea

`func (o *ModelsProvider) SetArea(v int32)`

SetArea sets Area field to given value.

### HasArea

`func (o *ModelsProvider) HasArea() bool`

HasArea returns a boolean if a field has been set.

### GetProviderType

`func (o *ModelsProvider) GetProviderType() ModelsProviderType`

GetProviderType returns the ProviderType field if non-nil, zero value otherwise.

### GetProviderTypeOk

`func (o *ModelsProvider) GetProviderTypeOk() (*ModelsProviderType, bool)`

GetProviderTypeOk returns a tuple with the ProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderType

`func (o *ModelsProvider) SetProviderType(v ModelsProviderType)`

SetProviderType sets ProviderType field to given value.


### GetProviderTypeParamValues

`func (o *ModelsProvider) GetProviderTypeParamValues() []ModelsPamProviderTypeParamValue`

GetProviderTypeParamValues returns the ProviderTypeParamValues field if non-nil, zero value otherwise.

### GetProviderTypeParamValuesOk

`func (o *ModelsProvider) GetProviderTypeParamValuesOk() (*[]ModelsPamProviderTypeParamValue, bool)`

GetProviderTypeParamValuesOk returns a tuple with the ProviderTypeParamValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderTypeParamValues

`func (o *ModelsProvider) SetProviderTypeParamValues(v []ModelsPamProviderTypeParamValue)`

SetProviderTypeParamValues sets ProviderTypeParamValues field to given value.

### HasProviderTypeParamValues

`func (o *ModelsProvider) HasProviderTypeParamValues() bool`

HasProviderTypeParamValues returns a boolean if a field has been set.

### SetProviderTypeParamValuesNil

`func (o *ModelsProvider) SetProviderTypeParamValuesNil(b bool)`

 SetProviderTypeParamValuesNil sets the value for ProviderTypeParamValues to be an explicit nil

### UnsetProviderTypeParamValues
`func (o *ModelsProvider) UnsetProviderTypeParamValues()`

UnsetProviderTypeParamValues ensures that no value is present for ProviderTypeParamValues, not even an explicit nil
### GetSecuredAreaId

`func (o *ModelsProvider) GetSecuredAreaId() int32`

GetSecuredAreaId returns the SecuredAreaId field if non-nil, zero value otherwise.

### GetSecuredAreaIdOk

`func (o *ModelsProvider) GetSecuredAreaIdOk() (*int32, bool)`

GetSecuredAreaIdOk returns a tuple with the SecuredAreaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecuredAreaId

`func (o *ModelsProvider) SetSecuredAreaId(v int32)`

SetSecuredAreaId sets SecuredAreaId field to given value.

### HasSecuredAreaId

`func (o *ModelsProvider) HasSecuredAreaId() bool`

HasSecuredAreaId returns a boolean if a field has been set.

### SetSecuredAreaIdNil

`func (o *ModelsProvider) SetSecuredAreaIdNil(b bool)`

 SetSecuredAreaIdNil sets the value for SecuredAreaId to be an explicit nil

### UnsetSecuredAreaId
`func (o *ModelsProvider) UnsetSecuredAreaId()`

UnsetSecuredAreaId ensures that no value is present for SecuredAreaId, not even an explicit nil
### GetRemote

`func (o *ModelsProvider) GetRemote() bool`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *ModelsProvider) GetRemoteOk() (*bool, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *ModelsProvider) SetRemote(v bool)`

SetRemote sets Remote field to given value.

### HasRemote

`func (o *ModelsProvider) HasRemote() bool`

HasRemote returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


