# PAMProviderCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Remote** | Pointer to **bool** |  | [optional] 
**Area** | Pointer to **int32** |  | [optional] 
**ProviderType** | [**PAMProviderCreateRequestProviderType**](PAMProviderCreateRequestProviderType.md) |  | 
**ProviderTypeParamValues** | Pointer to [**[]PAMProviderCreateRequestTypeParamValue**](PAMProviderCreateRequestTypeParamValue.md) |  | [optional] 
**SecuredAreaId** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewPAMProviderCreateRequest

`func NewPAMProviderCreateRequest(name string, providerType PAMProviderCreateRequestProviderType, ) *PAMProviderCreateRequest`

NewPAMProviderCreateRequest instantiates a new PAMProviderCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPAMProviderCreateRequestWithDefaults

`func NewPAMProviderCreateRequestWithDefaults() *PAMProviderCreateRequest`

NewPAMProviderCreateRequestWithDefaults instantiates a new PAMProviderCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PAMProviderCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PAMProviderCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PAMProviderCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetRemote

`func (o *PAMProviderCreateRequest) GetRemote() bool`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *PAMProviderCreateRequest) GetRemoteOk() (*bool, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *PAMProviderCreateRequest) SetRemote(v bool)`

SetRemote sets Remote field to given value.

### HasRemote

`func (o *PAMProviderCreateRequest) HasRemote() bool`

HasRemote returns a boolean if a field has been set.

### GetArea

`func (o *PAMProviderCreateRequest) GetArea() int32`

GetArea returns the Area field if non-nil, zero value otherwise.

### GetAreaOk

`func (o *PAMProviderCreateRequest) GetAreaOk() (*int32, bool)`

GetAreaOk returns a tuple with the Area field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArea

`func (o *PAMProviderCreateRequest) SetArea(v int32)`

SetArea sets Area field to given value.

### HasArea

`func (o *PAMProviderCreateRequest) HasArea() bool`

HasArea returns a boolean if a field has been set.

### GetProviderType

`func (o *PAMProviderCreateRequest) GetProviderType() PAMProviderCreateRequestProviderType`

GetProviderType returns the ProviderType field if non-nil, zero value otherwise.

### GetProviderTypeOk

`func (o *PAMProviderCreateRequest) GetProviderTypeOk() (*PAMProviderCreateRequestProviderType, bool)`

GetProviderTypeOk returns a tuple with the ProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderType

`func (o *PAMProviderCreateRequest) SetProviderType(v PAMProviderCreateRequestProviderType)`

SetProviderType sets ProviderType field to given value.


### GetProviderTypeParamValues

`func (o *PAMProviderCreateRequest) GetProviderTypeParamValues() []PAMProviderCreateRequestTypeParamValue`

GetProviderTypeParamValues returns the ProviderTypeParamValues field if non-nil, zero value otherwise.

### GetProviderTypeParamValuesOk

`func (o *PAMProviderCreateRequest) GetProviderTypeParamValuesOk() (*[]PAMProviderCreateRequestTypeParamValue, bool)`

GetProviderTypeParamValuesOk returns a tuple with the ProviderTypeParamValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderTypeParamValues

`func (o *PAMProviderCreateRequest) SetProviderTypeParamValues(v []PAMProviderCreateRequestTypeParamValue)`

SetProviderTypeParamValues sets ProviderTypeParamValues field to given value.

### HasProviderTypeParamValues

`func (o *PAMProviderCreateRequest) HasProviderTypeParamValues() bool`

HasProviderTypeParamValues returns a boolean if a field has been set.

### SetProviderTypeParamValuesNil

`func (o *PAMProviderCreateRequest) SetProviderTypeParamValuesNil(b bool)`

 SetProviderTypeParamValuesNil sets the value for ProviderTypeParamValues to be an explicit nil

### UnsetProviderTypeParamValues
`func (o *PAMProviderCreateRequest) UnsetProviderTypeParamValues()`

UnsetProviderTypeParamValues ensures that no value is present for ProviderTypeParamValues, not even an explicit nil
### GetSecuredAreaId

`func (o *PAMProviderCreateRequest) GetSecuredAreaId() int32`

GetSecuredAreaId returns the SecuredAreaId field if non-nil, zero value otherwise.

### GetSecuredAreaIdOk

`func (o *PAMProviderCreateRequest) GetSecuredAreaIdOk() (*int32, bool)`

GetSecuredAreaIdOk returns a tuple with the SecuredAreaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecuredAreaId

`func (o *PAMProviderCreateRequest) SetSecuredAreaId(v int32)`

SetSecuredAreaId sets SecuredAreaId field to given value.

### HasSecuredAreaId

`func (o *PAMProviderCreateRequest) HasSecuredAreaId() bool`

HasSecuredAreaId returns a boolean if a field has been set.

### SetSecuredAreaIdNil

`func (o *PAMProviderCreateRequest) SetSecuredAreaIdNil(b bool)`

 SetSecuredAreaIdNil sets the value for SecuredAreaId to be an explicit nil

### UnsetSecuredAreaId
`func (o *PAMProviderCreateRequest) UnsetSecuredAreaId()`

UnsetSecuredAreaId ensures that no value is present for SecuredAreaId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


