# PAMProviderUpdateRequestLegacy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**Remote** | Pointer to **bool** |  | [optional] 
**Area** | Pointer to **int32** |  | [optional] 
**ProviderType** | [**PAMProviderCreateRequestProviderType**](PAMProviderCreateRequestProviderType.md) |  | 
**ProviderTypeParamValues** | Pointer to [**[]PAMProviderCreateRequestTypeParamValue**](PAMProviderCreateRequestTypeParamValue.md) |  | [optional] 
**SecuredAreaId** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewPAMProviderUpdateRequestLegacy

`func NewPAMProviderUpdateRequestLegacy(id int32, name string, providerType PAMProviderCreateRequestProviderType, ) *PAMProviderUpdateRequestLegacy`

NewPAMProviderUpdateRequestLegacy instantiates a new PAMProviderUpdateRequestLegacy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPAMProviderUpdateRequestLegacyWithDefaults

`func NewPAMProviderUpdateRequestLegacyWithDefaults() *PAMProviderUpdateRequestLegacy`

NewPAMProviderUpdateRequestLegacyWithDefaults instantiates a new PAMProviderUpdateRequestLegacy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PAMProviderUpdateRequestLegacy) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PAMProviderUpdateRequestLegacy) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PAMProviderUpdateRequestLegacy) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *PAMProviderUpdateRequestLegacy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PAMProviderUpdateRequestLegacy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PAMProviderUpdateRequestLegacy) SetName(v string)`

SetName sets Name field to given value.


### GetRemote

`func (o *PAMProviderUpdateRequestLegacy) GetRemote() bool`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *PAMProviderUpdateRequestLegacy) GetRemoteOk() (*bool, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *PAMProviderUpdateRequestLegacy) SetRemote(v bool)`

SetRemote sets Remote field to given value.

### HasRemote

`func (o *PAMProviderUpdateRequestLegacy) HasRemote() bool`

HasRemote returns a boolean if a field has been set.

### GetArea

`func (o *PAMProviderUpdateRequestLegacy) GetArea() int32`

GetArea returns the Area field if non-nil, zero value otherwise.

### GetAreaOk

`func (o *PAMProviderUpdateRequestLegacy) GetAreaOk() (*int32, bool)`

GetAreaOk returns a tuple with the Area field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArea

`func (o *PAMProviderUpdateRequestLegacy) SetArea(v int32)`

SetArea sets Area field to given value.

### HasArea

`func (o *PAMProviderUpdateRequestLegacy) HasArea() bool`

HasArea returns a boolean if a field has been set.

### GetProviderType

`func (o *PAMProviderUpdateRequestLegacy) GetProviderType() PAMProviderCreateRequestProviderType`

GetProviderType returns the ProviderType field if non-nil, zero value otherwise.

### GetProviderTypeOk

`func (o *PAMProviderUpdateRequestLegacy) GetProviderTypeOk() (*PAMProviderCreateRequestProviderType, bool)`

GetProviderTypeOk returns a tuple with the ProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderType

`func (o *PAMProviderUpdateRequestLegacy) SetProviderType(v PAMProviderCreateRequestProviderType)`

SetProviderType sets ProviderType field to given value.


### GetProviderTypeParamValues

`func (o *PAMProviderUpdateRequestLegacy) GetProviderTypeParamValues() []PAMProviderCreateRequestTypeParamValue`

GetProviderTypeParamValues returns the ProviderTypeParamValues field if non-nil, zero value otherwise.

### GetProviderTypeParamValuesOk

`func (o *PAMProviderUpdateRequestLegacy) GetProviderTypeParamValuesOk() (*[]PAMProviderCreateRequestTypeParamValue, bool)`

GetProviderTypeParamValuesOk returns a tuple with the ProviderTypeParamValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderTypeParamValues

`func (o *PAMProviderUpdateRequestLegacy) SetProviderTypeParamValues(v []PAMProviderCreateRequestTypeParamValue)`

SetProviderTypeParamValues sets ProviderTypeParamValues field to given value.

### HasProviderTypeParamValues

`func (o *PAMProviderUpdateRequestLegacy) HasProviderTypeParamValues() bool`

HasProviderTypeParamValues returns a boolean if a field has been set.

### SetProviderTypeParamValuesNil

`func (o *PAMProviderUpdateRequestLegacy) SetProviderTypeParamValuesNil(b bool)`

 SetProviderTypeParamValuesNil sets the value for ProviderTypeParamValues to be an explicit nil

### UnsetProviderTypeParamValues
`func (o *PAMProviderUpdateRequestLegacy) UnsetProviderTypeParamValues()`

UnsetProviderTypeParamValues ensures that no value is present for ProviderTypeParamValues, not even an explicit nil
### GetSecuredAreaId

`func (o *PAMProviderUpdateRequestLegacy) GetSecuredAreaId() int32`

GetSecuredAreaId returns the SecuredAreaId field if non-nil, zero value otherwise.

### GetSecuredAreaIdOk

`func (o *PAMProviderUpdateRequestLegacy) GetSecuredAreaIdOk() (*int32, bool)`

GetSecuredAreaIdOk returns a tuple with the SecuredAreaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecuredAreaId

`func (o *PAMProviderUpdateRequestLegacy) SetSecuredAreaId(v int32)`

SetSecuredAreaId sets SecuredAreaId field to given value.

### HasSecuredAreaId

`func (o *PAMProviderUpdateRequestLegacy) HasSecuredAreaId() bool`

HasSecuredAreaId returns a boolean if a field has been set.

### SetSecuredAreaIdNil

`func (o *PAMProviderUpdateRequestLegacy) SetSecuredAreaIdNil(b bool)`

 SetSecuredAreaIdNil sets the value for SecuredAreaId to be an explicit nil

### UnsetSecuredAreaId
`func (o *PAMProviderUpdateRequestLegacy) UnsetSecuredAreaId()`

UnsetSecuredAreaId ensures that no value is present for SecuredAreaId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


