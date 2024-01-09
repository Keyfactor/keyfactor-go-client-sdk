# CSSCMSDataModelModelsProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  | 
**Area** | Pointer to **int32** |  | [optional] 
**ProviderType** | [**CSSCMSDataModelModelsProviderType**](CSSCMSDataModelModelsProviderType.md) |  | 
**ProviderTypeParamValues** | Pointer to [**[]CSSCMSDataModelModelsPamProviderTypeParamValue**](CSSCMSDataModelModelsPamProviderTypeParamValue.md) |  | [optional] 
**SecuredAreaId** | Pointer to **int32** |  | [optional] 

## Methods

### NewCSSCMSDataModelModelsProvider

`func NewCSSCMSDataModelModelsProvider(name string, providerType CSSCMSDataModelModelsProviderType, ) *CSSCMSDataModelModelsProvider`

NewCSSCMSDataModelModelsProvider instantiates a new CSSCMSDataModelModelsProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSSCMSDataModelModelsProviderWithDefaults

`func NewCSSCMSDataModelModelsProviderWithDefaults() *CSSCMSDataModelModelsProvider`

NewCSSCMSDataModelModelsProviderWithDefaults instantiates a new CSSCMSDataModelModelsProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CSSCMSDataModelModelsProvider) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CSSCMSDataModelModelsProvider) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CSSCMSDataModelModelsProvider) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CSSCMSDataModelModelsProvider) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CSSCMSDataModelModelsProvider) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CSSCMSDataModelModelsProvider) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CSSCMSDataModelModelsProvider) SetName(v string)`

SetName sets Name field to given value.


### GetArea

`func (o *CSSCMSDataModelModelsProvider) GetArea() int32`

GetArea returns the Area field if non-nil, zero value otherwise.

### GetAreaOk

`func (o *CSSCMSDataModelModelsProvider) GetAreaOk() (*int32, bool)`

GetAreaOk returns a tuple with the Area field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArea

`func (o *CSSCMSDataModelModelsProvider) SetArea(v int32)`

SetArea sets Area field to given value.

### HasArea

`func (o *CSSCMSDataModelModelsProvider) HasArea() bool`

HasArea returns a boolean if a field has been set.

### GetProviderType

`func (o *CSSCMSDataModelModelsProvider) GetProviderType() CSSCMSDataModelModelsProviderType`

GetProviderType returns the ProviderType field if non-nil, zero value otherwise.

### GetProviderTypeOk

`func (o *CSSCMSDataModelModelsProvider) GetProviderTypeOk() (*CSSCMSDataModelModelsProviderType, bool)`

GetProviderTypeOk returns a tuple with the ProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderType

`func (o *CSSCMSDataModelModelsProvider) SetProviderType(v CSSCMSDataModelModelsProviderType)`

SetProviderType sets ProviderType field to given value.


### GetProviderTypeParamValues

`func (o *CSSCMSDataModelModelsProvider) GetProviderTypeParamValues() []CSSCMSDataModelModelsPamProviderTypeParamValue`

GetProviderTypeParamValues returns the ProviderTypeParamValues field if non-nil, zero value otherwise.

### GetProviderTypeParamValuesOk

`func (o *CSSCMSDataModelModelsProvider) GetProviderTypeParamValuesOk() (*[]CSSCMSDataModelModelsPamProviderTypeParamValue, bool)`

GetProviderTypeParamValuesOk returns a tuple with the ProviderTypeParamValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderTypeParamValues

`func (o *CSSCMSDataModelModelsProvider) SetProviderTypeParamValues(v []CSSCMSDataModelModelsPamProviderTypeParamValue)`

SetProviderTypeParamValues sets ProviderTypeParamValues field to given value.

### HasProviderTypeParamValues

`func (o *CSSCMSDataModelModelsProvider) HasProviderTypeParamValues() bool`

HasProviderTypeParamValues returns a boolean if a field has been set.

### GetSecuredAreaId

`func (o *CSSCMSDataModelModelsProvider) GetSecuredAreaId() int32`

GetSecuredAreaId returns the SecuredAreaId field if non-nil, zero value otherwise.

### GetSecuredAreaIdOk

`func (o *CSSCMSDataModelModelsProvider) GetSecuredAreaIdOk() (*int32, bool)`

GetSecuredAreaIdOk returns a tuple with the SecuredAreaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecuredAreaId

`func (o *CSSCMSDataModelModelsProvider) SetSecuredAreaId(v int32)`

SetSecuredAreaId sets SecuredAreaId field to given value.

### HasSecuredAreaId

`func (o *CSSCMSDataModelModelsProvider) HasSecuredAreaId() bool`

HasSecuredAreaId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


