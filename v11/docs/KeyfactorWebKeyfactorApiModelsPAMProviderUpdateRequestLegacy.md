# KeyfactorWebKeyfactorApiModelsPAMProviderUpdateRequestLegacy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**Remote** | Pointer to **bool** |  | [optional] 
**Area** | Pointer to **int32** |  | [optional] 
**ProviderType** | [**KeyfactorWebKeyfactorApiModelsPAMProviderCreateRequestProviderType**](KeyfactorWebKeyfactorApiModelsPAMProviderCreateRequestProviderType.md) |  | 
**ProviderTypeParamValues** | Pointer to [**[]KeyfactorWebKeyfactorApiModelsPAMProviderCreateRequestTypeParamValue**](KeyfactorWebKeyfactorApiModelsPAMProviderCreateRequestTypeParamValue.md) |  | [optional] 
**SecuredAreaId** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewKeyfactorWebKeyfactorApiModelsPAMProviderUpdateRequestLegacy

`func NewKeyfactorWebKeyfactorApiModelsPAMProviderUpdateRequestLegacy(id int32, name string, providerType KeyfactorWebKeyfactorApiModelsPAMProviderCreateRequestProviderType, ) *KeyfactorWebKeyfactorApiModelsPAMProviderUpdateRequestLegacy`

NewKeyfactorWebKeyfactorApiModelsPAMProviderUpdateRequestLegacy instantiates a new KeyfactorWebKeyfactorApiModelsPAMProviderUpdateRequestLegacy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorWebKeyfactorApiModelsPAMProviderUpdateRequestLegacyWithDefaults

`func NewKeyfactorWebKeyfactorApiModelsPAMProviderUpdateRequestLegacyWithDefaults() *KeyfactorWebKeyfactorApiModelsPAMProviderUpdateRequestLegacy`

NewKeyfactorWebKeyfactorApiModelsPAMProviderUpdateRequestLegacyWithDefaults instantiates a new KeyfactorWebKeyfactorApiModelsPAMProviderUpdateRequestLegacy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderUpdateRequestLegacy) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderUpdateRequestLegacy) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderUpdateRequestLegacy) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderUpdateRequestLegacy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderUpdateRequestLegacy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderUpdateRequestLegacy) SetName(v string)`

SetName sets Name field to given value.


### GetRemote

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderUpdateRequestLegacy) GetRemote() bool`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderUpdateRequestLegacy) GetRemoteOk() (*bool, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderUpdateRequestLegacy) SetRemote(v bool)`

SetRemote sets Remote field to given value.

### HasRemote

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderUpdateRequestLegacy) HasRemote() bool`

HasRemote returns a boolean if a field has been set.

### GetArea

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderUpdateRequestLegacy) GetArea() int32`

GetArea returns the Area field if non-nil, zero value otherwise.

### GetAreaOk

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderUpdateRequestLegacy) GetAreaOk() (*int32, bool)`

GetAreaOk returns a tuple with the Area field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArea

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderUpdateRequestLegacy) SetArea(v int32)`

SetArea sets Area field to given value.

### HasArea

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderUpdateRequestLegacy) HasArea() bool`

HasArea returns a boolean if a field has been set.

### GetProviderType

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderUpdateRequestLegacy) GetProviderType() KeyfactorWebKeyfactorApiModelsPAMProviderCreateRequestProviderType`

GetProviderType returns the ProviderType field if non-nil, zero value otherwise.

### GetProviderTypeOk

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderUpdateRequestLegacy) GetProviderTypeOk() (*KeyfactorWebKeyfactorApiModelsPAMProviderCreateRequestProviderType, bool)`

GetProviderTypeOk returns a tuple with the ProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderType

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderUpdateRequestLegacy) SetProviderType(v KeyfactorWebKeyfactorApiModelsPAMProviderCreateRequestProviderType)`

SetProviderType sets ProviderType field to given value.


### GetProviderTypeParamValues

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderUpdateRequestLegacy) GetProviderTypeParamValues() []KeyfactorWebKeyfactorApiModelsPAMProviderCreateRequestTypeParamValue`

GetProviderTypeParamValues returns the ProviderTypeParamValues field if non-nil, zero value otherwise.

### GetProviderTypeParamValuesOk

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderUpdateRequestLegacy) GetProviderTypeParamValuesOk() (*[]KeyfactorWebKeyfactorApiModelsPAMProviderCreateRequestTypeParamValue, bool)`

GetProviderTypeParamValuesOk returns a tuple with the ProviderTypeParamValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderTypeParamValues

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderUpdateRequestLegacy) SetProviderTypeParamValues(v []KeyfactorWebKeyfactorApiModelsPAMProviderCreateRequestTypeParamValue)`

SetProviderTypeParamValues sets ProviderTypeParamValues field to given value.

### HasProviderTypeParamValues

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderUpdateRequestLegacy) HasProviderTypeParamValues() bool`

HasProviderTypeParamValues returns a boolean if a field has been set.

### SetProviderTypeParamValuesNil

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderUpdateRequestLegacy) SetProviderTypeParamValuesNil(b bool)`

 SetProviderTypeParamValuesNil sets the value for ProviderTypeParamValues to be an explicit nil

### UnsetProviderTypeParamValues
`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderUpdateRequestLegacy) UnsetProviderTypeParamValues()`

UnsetProviderTypeParamValues ensures that no value is present for ProviderTypeParamValues, not even an explicit nil
### GetSecuredAreaId

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderUpdateRequestLegacy) GetSecuredAreaId() int32`

GetSecuredAreaId returns the SecuredAreaId field if non-nil, zero value otherwise.

### GetSecuredAreaIdOk

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderUpdateRequestLegacy) GetSecuredAreaIdOk() (*int32, bool)`

GetSecuredAreaIdOk returns a tuple with the SecuredAreaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecuredAreaId

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderUpdateRequestLegacy) SetSecuredAreaId(v int32)`

SetSecuredAreaId sets SecuredAreaId field to given value.

### HasSecuredAreaId

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderUpdateRequestLegacy) HasSecuredAreaId() bool`

HasSecuredAreaId returns a boolean if a field has been set.

### SetSecuredAreaIdNil

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderUpdateRequestLegacy) SetSecuredAreaIdNil(b bool)`

 SetSecuredAreaIdNil sets the value for SecuredAreaId to be an explicit nil

### UnsetSecuredAreaId
`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderUpdateRequestLegacy) UnsetSecuredAreaId()`

UnsetSecuredAreaId ensures that no value is present for SecuredAreaId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


