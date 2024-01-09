# CSSCMSDataModelModelsKeyfactorSecret

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **interface{}** |  | [optional] 
**SecretTypeGuid** | Pointer to **string** |  | [optional] 
**InstanceId** | Pointer to **NullableInt32** |  | [optional] 
**InstanceGuid** | Pointer to **NullableString** |  | [optional] 
**ProviderTypeParameterValues** | Pointer to [**[]CSSCMSDataModelModelsPamProviderTypeParamValue**](CSSCMSDataModelModelsPamProviderTypeParamValue.md) |  | [optional] 
**ProviderId** | Pointer to **NullableInt32** |  | [optional] 
**IsManaged** | Pointer to **bool** |  | [optional] [readonly] 

## Methods

### NewCSSCMSDataModelModelsKeyfactorSecret

`func NewCSSCMSDataModelModelsKeyfactorSecret() *CSSCMSDataModelModelsKeyfactorSecret`

NewCSSCMSDataModelModelsKeyfactorSecret instantiates a new CSSCMSDataModelModelsKeyfactorSecret object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSSCMSDataModelModelsKeyfactorSecretWithDefaults

`func NewCSSCMSDataModelModelsKeyfactorSecretWithDefaults() *CSSCMSDataModelModelsKeyfactorSecret`

NewCSSCMSDataModelModelsKeyfactorSecretWithDefaults instantiates a new CSSCMSDataModelModelsKeyfactorSecret object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *CSSCMSDataModelModelsKeyfactorSecret) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CSSCMSDataModelModelsKeyfactorSecret) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CSSCMSDataModelModelsKeyfactorSecret) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *CSSCMSDataModelModelsKeyfactorSecret) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *CSSCMSDataModelModelsKeyfactorSecret) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *CSSCMSDataModelModelsKeyfactorSecret) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetSecretTypeGuid

`func (o *CSSCMSDataModelModelsKeyfactorSecret) GetSecretTypeGuid() string`

GetSecretTypeGuid returns the SecretTypeGuid field if non-nil, zero value otherwise.

### GetSecretTypeGuidOk

`func (o *CSSCMSDataModelModelsKeyfactorSecret) GetSecretTypeGuidOk() (*string, bool)`

GetSecretTypeGuidOk returns a tuple with the SecretTypeGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretTypeGuid

`func (o *CSSCMSDataModelModelsKeyfactorSecret) SetSecretTypeGuid(v string)`

SetSecretTypeGuid sets SecretTypeGuid field to given value.

### HasSecretTypeGuid

`func (o *CSSCMSDataModelModelsKeyfactorSecret) HasSecretTypeGuid() bool`

HasSecretTypeGuid returns a boolean if a field has been set.

### GetInstanceId

`func (o *CSSCMSDataModelModelsKeyfactorSecret) GetInstanceId() int32`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *CSSCMSDataModelModelsKeyfactorSecret) GetInstanceIdOk() (*int32, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *CSSCMSDataModelModelsKeyfactorSecret) SetInstanceId(v int32)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *CSSCMSDataModelModelsKeyfactorSecret) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### SetInstanceIdNil

`func (o *CSSCMSDataModelModelsKeyfactorSecret) SetInstanceIdNil(b bool)`

 SetInstanceIdNil sets the value for InstanceId to be an explicit nil

### UnsetInstanceId
`func (o *CSSCMSDataModelModelsKeyfactorSecret) UnsetInstanceId()`

UnsetInstanceId ensures that no value is present for InstanceId, not even an explicit nil
### GetInstanceGuid

`func (o *CSSCMSDataModelModelsKeyfactorSecret) GetInstanceGuid() string`

GetInstanceGuid returns the InstanceGuid field if non-nil, zero value otherwise.

### GetInstanceGuidOk

`func (o *CSSCMSDataModelModelsKeyfactorSecret) GetInstanceGuidOk() (*string, bool)`

GetInstanceGuidOk returns a tuple with the InstanceGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceGuid

`func (o *CSSCMSDataModelModelsKeyfactorSecret) SetInstanceGuid(v string)`

SetInstanceGuid sets InstanceGuid field to given value.

### HasInstanceGuid

`func (o *CSSCMSDataModelModelsKeyfactorSecret) HasInstanceGuid() bool`

HasInstanceGuid returns a boolean if a field has been set.

### SetInstanceGuidNil

`func (o *CSSCMSDataModelModelsKeyfactorSecret) SetInstanceGuidNil(b bool)`

 SetInstanceGuidNil sets the value for InstanceGuid to be an explicit nil

### UnsetInstanceGuid
`func (o *CSSCMSDataModelModelsKeyfactorSecret) UnsetInstanceGuid()`

UnsetInstanceGuid ensures that no value is present for InstanceGuid, not even an explicit nil
### GetProviderTypeParameterValues

`func (o *CSSCMSDataModelModelsKeyfactorSecret) GetProviderTypeParameterValues() []CSSCMSDataModelModelsPamProviderTypeParamValue`

GetProviderTypeParameterValues returns the ProviderTypeParameterValues field if non-nil, zero value otherwise.

### GetProviderTypeParameterValuesOk

`func (o *CSSCMSDataModelModelsKeyfactorSecret) GetProviderTypeParameterValuesOk() (*[]CSSCMSDataModelModelsPamProviderTypeParamValue, bool)`

GetProviderTypeParameterValuesOk returns a tuple with the ProviderTypeParameterValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderTypeParameterValues

`func (o *CSSCMSDataModelModelsKeyfactorSecret) SetProviderTypeParameterValues(v []CSSCMSDataModelModelsPamProviderTypeParamValue)`

SetProviderTypeParameterValues sets ProviderTypeParameterValues field to given value.

### HasProviderTypeParameterValues

`func (o *CSSCMSDataModelModelsKeyfactorSecret) HasProviderTypeParameterValues() bool`

HasProviderTypeParameterValues returns a boolean if a field has been set.

### SetProviderTypeParameterValuesNil

`func (o *CSSCMSDataModelModelsKeyfactorSecret) SetProviderTypeParameterValuesNil(b bool)`

 SetProviderTypeParameterValuesNil sets the value for ProviderTypeParameterValues to be an explicit nil

### UnsetProviderTypeParameterValues
`func (o *CSSCMSDataModelModelsKeyfactorSecret) UnsetProviderTypeParameterValues()`

UnsetProviderTypeParameterValues ensures that no value is present for ProviderTypeParameterValues, not even an explicit nil
### GetProviderId

`func (o *CSSCMSDataModelModelsKeyfactorSecret) GetProviderId() int32`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *CSSCMSDataModelModelsKeyfactorSecret) GetProviderIdOk() (*int32, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *CSSCMSDataModelModelsKeyfactorSecret) SetProviderId(v int32)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *CSSCMSDataModelModelsKeyfactorSecret) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### SetProviderIdNil

`func (o *CSSCMSDataModelModelsKeyfactorSecret) SetProviderIdNil(b bool)`

 SetProviderIdNil sets the value for ProviderId to be an explicit nil

### UnsetProviderId
`func (o *CSSCMSDataModelModelsKeyfactorSecret) UnsetProviderId()`

UnsetProviderId ensures that no value is present for ProviderId, not even an explicit nil
### GetIsManaged

`func (o *CSSCMSDataModelModelsKeyfactorSecret) GetIsManaged() bool`

GetIsManaged returns the IsManaged field if non-nil, zero value otherwise.

### GetIsManagedOk

`func (o *CSSCMSDataModelModelsKeyfactorSecret) GetIsManagedOk() (*bool, bool)`

GetIsManagedOk returns a tuple with the IsManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsManaged

`func (o *CSSCMSDataModelModelsKeyfactorSecret) SetIsManaged(v bool)`

SetIsManaged sets IsManaged field to given value.

### HasIsManaged

`func (o *CSSCMSDataModelModelsKeyfactorSecret) HasIsManaged() bool`

HasIsManaged returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


