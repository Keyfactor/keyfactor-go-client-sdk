# KeyfactorWebKeyfactorApiModelsPAMProviderResponseLegacy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Area** | Pointer to **int32** |  | [optional] 
**ProviderType** | Pointer to [**CSSCMSDataModelModelsProviderType**](CSSCMSDataModelModelsProviderType.md) |  | [optional] 
**ProviderTypeParamValues** | Pointer to [**[]KeyfactorWebKeyfactorApiModelsPAMPamProviderTypeParamValueResponse**](KeyfactorWebKeyfactorApiModelsPAMPamProviderTypeParamValueResponse.md) |  | [optional] 
**SecuredAreaId** | Pointer to **NullableInt32** |  | [optional] 
**Remote** | Pointer to **bool** |  | [optional] 

## Methods

### NewKeyfactorWebKeyfactorApiModelsPAMProviderResponseLegacy

`func NewKeyfactorWebKeyfactorApiModelsPAMProviderResponseLegacy() *KeyfactorWebKeyfactorApiModelsPAMProviderResponseLegacy`

NewKeyfactorWebKeyfactorApiModelsPAMProviderResponseLegacy instantiates a new KeyfactorWebKeyfactorApiModelsPAMProviderResponseLegacy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorWebKeyfactorApiModelsPAMProviderResponseLegacyWithDefaults

`func NewKeyfactorWebKeyfactorApiModelsPAMProviderResponseLegacyWithDefaults() *KeyfactorWebKeyfactorApiModelsPAMProviderResponseLegacy`

NewKeyfactorWebKeyfactorApiModelsPAMProviderResponseLegacyWithDefaults instantiates a new KeyfactorWebKeyfactorApiModelsPAMProviderResponseLegacy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderResponseLegacy) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderResponseLegacy) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderResponseLegacy) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderResponseLegacy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderResponseLegacy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderResponseLegacy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderResponseLegacy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderResponseLegacy) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderResponseLegacy) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderResponseLegacy) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetArea

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderResponseLegacy) GetArea() int32`

GetArea returns the Area field if non-nil, zero value otherwise.

### GetAreaOk

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderResponseLegacy) GetAreaOk() (*int32, bool)`

GetAreaOk returns a tuple with the Area field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArea

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderResponseLegacy) SetArea(v int32)`

SetArea sets Area field to given value.

### HasArea

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderResponseLegacy) HasArea() bool`

HasArea returns a boolean if a field has been set.

### GetProviderType

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderResponseLegacy) GetProviderType() CSSCMSDataModelModelsProviderType`

GetProviderType returns the ProviderType field if non-nil, zero value otherwise.

### GetProviderTypeOk

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderResponseLegacy) GetProviderTypeOk() (*CSSCMSDataModelModelsProviderType, bool)`

GetProviderTypeOk returns a tuple with the ProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderType

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderResponseLegacy) SetProviderType(v CSSCMSDataModelModelsProviderType)`

SetProviderType sets ProviderType field to given value.

### HasProviderType

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderResponseLegacy) HasProviderType() bool`

HasProviderType returns a boolean if a field has been set.

### GetProviderTypeParamValues

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderResponseLegacy) GetProviderTypeParamValues() []KeyfactorWebKeyfactorApiModelsPAMPamProviderTypeParamValueResponse`

GetProviderTypeParamValues returns the ProviderTypeParamValues field if non-nil, zero value otherwise.

### GetProviderTypeParamValuesOk

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderResponseLegacy) GetProviderTypeParamValuesOk() (*[]KeyfactorWebKeyfactorApiModelsPAMPamProviderTypeParamValueResponse, bool)`

GetProviderTypeParamValuesOk returns a tuple with the ProviderTypeParamValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderTypeParamValues

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderResponseLegacy) SetProviderTypeParamValues(v []KeyfactorWebKeyfactorApiModelsPAMPamProviderTypeParamValueResponse)`

SetProviderTypeParamValues sets ProviderTypeParamValues field to given value.

### HasProviderTypeParamValues

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderResponseLegacy) HasProviderTypeParamValues() bool`

HasProviderTypeParamValues returns a boolean if a field has been set.

### SetProviderTypeParamValuesNil

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderResponseLegacy) SetProviderTypeParamValuesNil(b bool)`

 SetProviderTypeParamValuesNil sets the value for ProviderTypeParamValues to be an explicit nil

### UnsetProviderTypeParamValues
`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderResponseLegacy) UnsetProviderTypeParamValues()`

UnsetProviderTypeParamValues ensures that no value is present for ProviderTypeParamValues, not even an explicit nil
### GetSecuredAreaId

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderResponseLegacy) GetSecuredAreaId() int32`

GetSecuredAreaId returns the SecuredAreaId field if non-nil, zero value otherwise.

### GetSecuredAreaIdOk

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderResponseLegacy) GetSecuredAreaIdOk() (*int32, bool)`

GetSecuredAreaIdOk returns a tuple with the SecuredAreaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecuredAreaId

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderResponseLegacy) SetSecuredAreaId(v int32)`

SetSecuredAreaId sets SecuredAreaId field to given value.

### HasSecuredAreaId

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderResponseLegacy) HasSecuredAreaId() bool`

HasSecuredAreaId returns a boolean if a field has been set.

### SetSecuredAreaIdNil

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderResponseLegacy) SetSecuredAreaIdNil(b bool)`

 SetSecuredAreaIdNil sets the value for SecuredAreaId to be an explicit nil

### UnsetSecuredAreaId
`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderResponseLegacy) UnsetSecuredAreaId()`

UnsetSecuredAreaId ensures that no value is present for SecuredAreaId, not even an explicit nil
### GetRemote

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderResponseLegacy) GetRemote() bool`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderResponseLegacy) GetRemoteOk() (*bool, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderResponseLegacy) SetRemote(v bool)`

SetRemote sets Remote field to given value.

### HasRemote

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderResponseLegacy) HasRemote() bool`

HasRemote returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


