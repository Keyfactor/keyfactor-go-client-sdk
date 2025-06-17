# PAMProviderResponseLegacy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Area** | Pointer to **int32** |  | [optional] 
**ProviderType** | Pointer to [**CSSCMSDataModelModelsProviderType**](CSSCMSDataModelModelsProviderType.md) |  | [optional] 
**ProviderTypeParamValues** | Pointer to [**[]PAMPamProviderTypeParamValueResponse**](PAMPamProviderTypeParamValueResponse.md) |  | [optional] 
**SecuredAreaId** | Pointer to **NullableInt32** |  | [optional] 
**Remote** | Pointer to **bool** |  | [optional] 

## Methods

### NewPAMProviderResponseLegacy

`func NewPAMProviderResponseLegacy() *PAMProviderResponseLegacy`

NewPAMProviderResponseLegacy instantiates a new PAMProviderResponseLegacy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPAMProviderResponseLegacyWithDefaults

`func NewPAMProviderResponseLegacyWithDefaults() *PAMProviderResponseLegacy`

NewPAMProviderResponseLegacyWithDefaults instantiates a new PAMProviderResponseLegacy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PAMProviderResponseLegacy) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PAMProviderResponseLegacy) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PAMProviderResponseLegacy) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PAMProviderResponseLegacy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PAMProviderResponseLegacy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PAMProviderResponseLegacy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PAMProviderResponseLegacy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PAMProviderResponseLegacy) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PAMProviderResponseLegacy) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PAMProviderResponseLegacy) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetArea

`func (o *PAMProviderResponseLegacy) GetArea() int32`

GetArea returns the Area field if non-nil, zero value otherwise.

### GetAreaOk

`func (o *PAMProviderResponseLegacy) GetAreaOk() (*int32, bool)`

GetAreaOk returns a tuple with the Area field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArea

`func (o *PAMProviderResponseLegacy) SetArea(v int32)`

SetArea sets Area field to given value.

### HasArea

`func (o *PAMProviderResponseLegacy) HasArea() bool`

HasArea returns a boolean if a field has been set.

### GetProviderType

`func (o *PAMProviderResponseLegacy) GetProviderType() CSSCMSDataModelModelsProviderType`

GetProviderType returns the ProviderType field if non-nil, zero value otherwise.

### GetProviderTypeOk

`func (o *PAMProviderResponseLegacy) GetProviderTypeOk() (*CSSCMSDataModelModelsProviderType, bool)`

GetProviderTypeOk returns a tuple with the ProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderType

`func (o *PAMProviderResponseLegacy) SetProviderType(v CSSCMSDataModelModelsProviderType)`

SetProviderType sets ProviderType field to given value.

### HasProviderType

`func (o *PAMProviderResponseLegacy) HasProviderType() bool`

HasProviderType returns a boolean if a field has been set.

### GetProviderTypeParamValues

`func (o *PAMProviderResponseLegacy) GetProviderTypeParamValues() []PAMPamProviderTypeParamValueResponse`

GetProviderTypeParamValues returns the ProviderTypeParamValues field if non-nil, zero value otherwise.

### GetProviderTypeParamValuesOk

`func (o *PAMProviderResponseLegacy) GetProviderTypeParamValuesOk() (*[]PAMPamProviderTypeParamValueResponse, bool)`

GetProviderTypeParamValuesOk returns a tuple with the ProviderTypeParamValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderTypeParamValues

`func (o *PAMProviderResponseLegacy) SetProviderTypeParamValues(v []PAMPamProviderTypeParamValueResponse)`

SetProviderTypeParamValues sets ProviderTypeParamValues field to given value.

### HasProviderTypeParamValues

`func (o *PAMProviderResponseLegacy) HasProviderTypeParamValues() bool`

HasProviderTypeParamValues returns a boolean if a field has been set.

### SetProviderTypeParamValuesNil

`func (o *PAMProviderResponseLegacy) SetProviderTypeParamValuesNil(b bool)`

 SetProviderTypeParamValuesNil sets the value for ProviderTypeParamValues to be an explicit nil

### UnsetProviderTypeParamValues
`func (o *PAMProviderResponseLegacy) UnsetProviderTypeParamValues()`

UnsetProviderTypeParamValues ensures that no value is present for ProviderTypeParamValues, not even an explicit nil
### GetSecuredAreaId

`func (o *PAMProviderResponseLegacy) GetSecuredAreaId() int32`

GetSecuredAreaId returns the SecuredAreaId field if non-nil, zero value otherwise.

### GetSecuredAreaIdOk

`func (o *PAMProviderResponseLegacy) GetSecuredAreaIdOk() (*int32, bool)`

GetSecuredAreaIdOk returns a tuple with the SecuredAreaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecuredAreaId

`func (o *PAMProviderResponseLegacy) SetSecuredAreaId(v int32)`

SetSecuredAreaId sets SecuredAreaId field to given value.

### HasSecuredAreaId

`func (o *PAMProviderResponseLegacy) HasSecuredAreaId() bool`

HasSecuredAreaId returns a boolean if a field has been set.

### SetSecuredAreaIdNil

`func (o *PAMProviderResponseLegacy) SetSecuredAreaIdNil(b bool)`

 SetSecuredAreaIdNil sets the value for SecuredAreaId to be an explicit nil

### UnsetSecuredAreaId
`func (o *PAMProviderResponseLegacy) UnsetSecuredAreaId()`

UnsetSecuredAreaId ensures that no value is present for SecuredAreaId, not even an explicit nil
### GetRemote

`func (o *PAMProviderResponseLegacy) GetRemote() bool`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *PAMProviderResponseLegacy) GetRemoteOk() (*bool, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *PAMProviderResponseLegacy) SetRemote(v bool)`

SetRemote sets Remote field to given value.

### HasRemote

`func (o *PAMProviderResponseLegacy) HasRemote() bool`

HasRemote returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


