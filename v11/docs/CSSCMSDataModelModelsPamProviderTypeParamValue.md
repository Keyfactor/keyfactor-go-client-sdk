# CSSCMSDataModelModelsPamProviderTypeParamValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Value** | Pointer to **NullableString** |  | [optional] 
**InstanceId** | Pointer to **NullableInt32** |  | [optional] 
**InstanceGuid** | Pointer to **NullableString** |  | [optional] 
**Provider** | Pointer to [**CSSCMSDataModelModelsProvider**](CSSCMSDataModelModelsProvider.md) |  | [optional] 
**ProviderTypeParam** | Pointer to [**CSSCMSDataModelModelsProviderTypeParam**](CSSCMSDataModelModelsProviderTypeParam.md) |  | [optional] 

## Methods

### NewCSSCMSDataModelModelsPamProviderTypeParamValue

`func NewCSSCMSDataModelModelsPamProviderTypeParamValue() *CSSCMSDataModelModelsPamProviderTypeParamValue`

NewCSSCMSDataModelModelsPamProviderTypeParamValue instantiates a new CSSCMSDataModelModelsPamProviderTypeParamValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSSCMSDataModelModelsPamProviderTypeParamValueWithDefaults

`func NewCSSCMSDataModelModelsPamProviderTypeParamValueWithDefaults() *CSSCMSDataModelModelsPamProviderTypeParamValue`

NewCSSCMSDataModelModelsPamProviderTypeParamValueWithDefaults instantiates a new CSSCMSDataModelModelsPamProviderTypeParamValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CSSCMSDataModelModelsPamProviderTypeParamValue) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CSSCMSDataModelModelsPamProviderTypeParamValue) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CSSCMSDataModelModelsPamProviderTypeParamValue) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CSSCMSDataModelModelsPamProviderTypeParamValue) HasId() bool`

HasId returns a boolean if a field has been set.

### GetValue

`func (o *CSSCMSDataModelModelsPamProviderTypeParamValue) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CSSCMSDataModelModelsPamProviderTypeParamValue) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CSSCMSDataModelModelsPamProviderTypeParamValue) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *CSSCMSDataModelModelsPamProviderTypeParamValue) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *CSSCMSDataModelModelsPamProviderTypeParamValue) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *CSSCMSDataModelModelsPamProviderTypeParamValue) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetInstanceId

`func (o *CSSCMSDataModelModelsPamProviderTypeParamValue) GetInstanceId() int32`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *CSSCMSDataModelModelsPamProviderTypeParamValue) GetInstanceIdOk() (*int32, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *CSSCMSDataModelModelsPamProviderTypeParamValue) SetInstanceId(v int32)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *CSSCMSDataModelModelsPamProviderTypeParamValue) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### SetInstanceIdNil

`func (o *CSSCMSDataModelModelsPamProviderTypeParamValue) SetInstanceIdNil(b bool)`

 SetInstanceIdNil sets the value for InstanceId to be an explicit nil

### UnsetInstanceId
`func (o *CSSCMSDataModelModelsPamProviderTypeParamValue) UnsetInstanceId()`

UnsetInstanceId ensures that no value is present for InstanceId, not even an explicit nil
### GetInstanceGuid

`func (o *CSSCMSDataModelModelsPamProviderTypeParamValue) GetInstanceGuid() string`

GetInstanceGuid returns the InstanceGuid field if non-nil, zero value otherwise.

### GetInstanceGuidOk

`func (o *CSSCMSDataModelModelsPamProviderTypeParamValue) GetInstanceGuidOk() (*string, bool)`

GetInstanceGuidOk returns a tuple with the InstanceGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceGuid

`func (o *CSSCMSDataModelModelsPamProviderTypeParamValue) SetInstanceGuid(v string)`

SetInstanceGuid sets InstanceGuid field to given value.

### HasInstanceGuid

`func (o *CSSCMSDataModelModelsPamProviderTypeParamValue) HasInstanceGuid() bool`

HasInstanceGuid returns a boolean if a field has been set.

### SetInstanceGuidNil

`func (o *CSSCMSDataModelModelsPamProviderTypeParamValue) SetInstanceGuidNil(b bool)`

 SetInstanceGuidNil sets the value for InstanceGuid to be an explicit nil

### UnsetInstanceGuid
`func (o *CSSCMSDataModelModelsPamProviderTypeParamValue) UnsetInstanceGuid()`

UnsetInstanceGuid ensures that no value is present for InstanceGuid, not even an explicit nil
### GetProvider

`func (o *CSSCMSDataModelModelsPamProviderTypeParamValue) GetProvider() CSSCMSDataModelModelsProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CSSCMSDataModelModelsPamProviderTypeParamValue) GetProviderOk() (*CSSCMSDataModelModelsProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CSSCMSDataModelModelsPamProviderTypeParamValue) SetProvider(v CSSCMSDataModelModelsProvider)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *CSSCMSDataModelModelsPamProviderTypeParamValue) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetProviderTypeParam

`func (o *CSSCMSDataModelModelsPamProviderTypeParamValue) GetProviderTypeParam() CSSCMSDataModelModelsProviderTypeParam`

GetProviderTypeParam returns the ProviderTypeParam field if non-nil, zero value otherwise.

### GetProviderTypeParamOk

`func (o *CSSCMSDataModelModelsPamProviderTypeParamValue) GetProviderTypeParamOk() (*CSSCMSDataModelModelsProviderTypeParam, bool)`

GetProviderTypeParamOk returns a tuple with the ProviderTypeParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderTypeParam

`func (o *CSSCMSDataModelModelsPamProviderTypeParamValue) SetProviderTypeParam(v CSSCMSDataModelModelsProviderTypeParam)`

SetProviderTypeParam sets ProviderTypeParam field to given value.

### HasProviderTypeParam

`func (o *CSSCMSDataModelModelsPamProviderTypeParamValue) HasProviderTypeParam() bool`

HasProviderTypeParam returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


