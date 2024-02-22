# KeyfactorWebKeyfactorApiModelsPAMProviderCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Remote** | Pointer to **bool** |  | [optional] 
**Area** | Pointer to **int32** |  | [optional] 
**ProviderType** | [**KeyfactorWebKeyfactorApiModelsPAMProviderCreateRequestProviderType**](KeyfactorWebKeyfactorApiModelsPAMProviderCreateRequestProviderType.md) |  | 
**ProviderTypeParamValues** | Pointer to [**[]KeyfactorWebKeyfactorApiModelsPAMProviderCreateRequestTypeParamValue**](KeyfactorWebKeyfactorApiModelsPAMProviderCreateRequestTypeParamValue.md) |  | [optional] 
**SecuredAreaId** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewKeyfactorWebKeyfactorApiModelsPAMProviderCreateRequest

`func NewKeyfactorWebKeyfactorApiModelsPAMProviderCreateRequest(name string, providerType KeyfactorWebKeyfactorApiModelsPAMProviderCreateRequestProviderType, ) *KeyfactorWebKeyfactorApiModelsPAMProviderCreateRequest`

NewKeyfactorWebKeyfactorApiModelsPAMProviderCreateRequest instantiates a new KeyfactorWebKeyfactorApiModelsPAMProviderCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorWebKeyfactorApiModelsPAMProviderCreateRequestWithDefaults

`func NewKeyfactorWebKeyfactorApiModelsPAMProviderCreateRequestWithDefaults() *KeyfactorWebKeyfactorApiModelsPAMProviderCreateRequest`

NewKeyfactorWebKeyfactorApiModelsPAMProviderCreateRequestWithDefaults instantiates a new KeyfactorWebKeyfactorApiModelsPAMProviderCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetRemote

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderCreateRequest) GetRemote() bool`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderCreateRequest) GetRemoteOk() (*bool, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderCreateRequest) SetRemote(v bool)`

SetRemote sets Remote field to given value.

### HasRemote

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderCreateRequest) HasRemote() bool`

HasRemote returns a boolean if a field has been set.

### GetArea

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderCreateRequest) GetArea() int32`

GetArea returns the Area field if non-nil, zero value otherwise.

### GetAreaOk

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderCreateRequest) GetAreaOk() (*int32, bool)`

GetAreaOk returns a tuple with the Area field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArea

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderCreateRequest) SetArea(v int32)`

SetArea sets Area field to given value.

### HasArea

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderCreateRequest) HasArea() bool`

HasArea returns a boolean if a field has been set.

### GetProviderType

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderCreateRequest) GetProviderType() KeyfactorWebKeyfactorApiModelsPAMProviderCreateRequestProviderType`

GetProviderType returns the ProviderType field if non-nil, zero value otherwise.

### GetProviderTypeOk

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderCreateRequest) GetProviderTypeOk() (*KeyfactorWebKeyfactorApiModelsPAMProviderCreateRequestProviderType, bool)`

GetProviderTypeOk returns a tuple with the ProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderType

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderCreateRequest) SetProviderType(v KeyfactorWebKeyfactorApiModelsPAMProviderCreateRequestProviderType)`

SetProviderType sets ProviderType field to given value.


### GetProviderTypeParamValues

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderCreateRequest) GetProviderTypeParamValues() []KeyfactorWebKeyfactorApiModelsPAMProviderCreateRequestTypeParamValue`

GetProviderTypeParamValues returns the ProviderTypeParamValues field if non-nil, zero value otherwise.

### GetProviderTypeParamValuesOk

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderCreateRequest) GetProviderTypeParamValuesOk() (*[]KeyfactorWebKeyfactorApiModelsPAMProviderCreateRequestTypeParamValue, bool)`

GetProviderTypeParamValuesOk returns a tuple with the ProviderTypeParamValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderTypeParamValues

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderCreateRequest) SetProviderTypeParamValues(v []KeyfactorWebKeyfactorApiModelsPAMProviderCreateRequestTypeParamValue)`

SetProviderTypeParamValues sets ProviderTypeParamValues field to given value.

### HasProviderTypeParamValues

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderCreateRequest) HasProviderTypeParamValues() bool`

HasProviderTypeParamValues returns a boolean if a field has been set.

### SetProviderTypeParamValuesNil

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderCreateRequest) SetProviderTypeParamValuesNil(b bool)`

 SetProviderTypeParamValuesNil sets the value for ProviderTypeParamValues to be an explicit nil

### UnsetProviderTypeParamValues
`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderCreateRequest) UnsetProviderTypeParamValues()`

UnsetProviderTypeParamValues ensures that no value is present for ProviderTypeParamValues, not even an explicit nil
### GetSecuredAreaId

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderCreateRequest) GetSecuredAreaId() int32`

GetSecuredAreaId returns the SecuredAreaId field if non-nil, zero value otherwise.

### GetSecuredAreaIdOk

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderCreateRequest) GetSecuredAreaIdOk() (*int32, bool)`

GetSecuredAreaIdOk returns a tuple with the SecuredAreaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecuredAreaId

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderCreateRequest) SetSecuredAreaId(v int32)`

SetSecuredAreaId sets SecuredAreaId field to given value.

### HasSecuredAreaId

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderCreateRequest) HasSecuredAreaId() bool`

HasSecuredAreaId returns a boolean if a field has been set.

### SetSecuredAreaIdNil

`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderCreateRequest) SetSecuredAreaIdNil(b bool)`

 SetSecuredAreaIdNil sets the value for SecuredAreaId to be an explicit nil

### UnsetSecuredAreaId
`func (o *KeyfactorWebKeyfactorApiModelsPAMProviderCreateRequest) UnsetSecuredAreaId()`

UnsetSecuredAreaId ensures that no value is present for SecuredAreaId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


