# PAMPamProviderTypeParamValueResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Value** | Pointer to **NullableString** |  | [optional] 
**InstanceId** | Pointer to **NullableInt32** |  | [optional] 
**InstanceGuid** | Pointer to **NullableString** |  | [optional] 
**ProviderTypeParam** | Pointer to [**PAMProviderTypeParameterResponse**](PAMProviderTypeParameterResponse.md) |  | [optional] 

## Methods

### NewPAMPamProviderTypeParamValueResponse

`func NewPAMPamProviderTypeParamValueResponse() *PAMPamProviderTypeParamValueResponse`

NewPAMPamProviderTypeParamValueResponse instantiates a new PAMPamProviderTypeParamValueResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPAMPamProviderTypeParamValueResponseWithDefaults

`func NewPAMPamProviderTypeParamValueResponseWithDefaults() *PAMPamProviderTypeParamValueResponse`

NewPAMPamProviderTypeParamValueResponseWithDefaults instantiates a new PAMPamProviderTypeParamValueResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PAMPamProviderTypeParamValueResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PAMPamProviderTypeParamValueResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PAMPamProviderTypeParamValueResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PAMPamProviderTypeParamValueResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetValue

`func (o *PAMPamProviderTypeParamValueResponse) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PAMPamProviderTypeParamValueResponse) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PAMPamProviderTypeParamValueResponse) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *PAMPamProviderTypeParamValueResponse) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *PAMPamProviderTypeParamValueResponse) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *PAMPamProviderTypeParamValueResponse) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetInstanceId

`func (o *PAMPamProviderTypeParamValueResponse) GetInstanceId() int32`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *PAMPamProviderTypeParamValueResponse) GetInstanceIdOk() (*int32, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *PAMPamProviderTypeParamValueResponse) SetInstanceId(v int32)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *PAMPamProviderTypeParamValueResponse) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### SetInstanceIdNil

`func (o *PAMPamProviderTypeParamValueResponse) SetInstanceIdNil(b bool)`

 SetInstanceIdNil sets the value for InstanceId to be an explicit nil

### UnsetInstanceId
`func (o *PAMPamProviderTypeParamValueResponse) UnsetInstanceId()`

UnsetInstanceId ensures that no value is present for InstanceId, not even an explicit nil
### GetInstanceGuid

`func (o *PAMPamProviderTypeParamValueResponse) GetInstanceGuid() string`

GetInstanceGuid returns the InstanceGuid field if non-nil, zero value otherwise.

### GetInstanceGuidOk

`func (o *PAMPamProviderTypeParamValueResponse) GetInstanceGuidOk() (*string, bool)`

GetInstanceGuidOk returns a tuple with the InstanceGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceGuid

`func (o *PAMPamProviderTypeParamValueResponse) SetInstanceGuid(v string)`

SetInstanceGuid sets InstanceGuid field to given value.

### HasInstanceGuid

`func (o *PAMPamProviderTypeParamValueResponse) HasInstanceGuid() bool`

HasInstanceGuid returns a boolean if a field has been set.

### SetInstanceGuidNil

`func (o *PAMPamProviderTypeParamValueResponse) SetInstanceGuidNil(b bool)`

 SetInstanceGuidNil sets the value for InstanceGuid to be an explicit nil

### UnsetInstanceGuid
`func (o *PAMPamProviderTypeParamValueResponse) UnsetInstanceGuid()`

UnsetInstanceGuid ensures that no value is present for InstanceGuid, not even an explicit nil
### GetProviderTypeParam

`func (o *PAMPamProviderTypeParamValueResponse) GetProviderTypeParam() PAMProviderTypeParameterResponse`

GetProviderTypeParam returns the ProviderTypeParam field if non-nil, zero value otherwise.

### GetProviderTypeParamOk

`func (o *PAMPamProviderTypeParamValueResponse) GetProviderTypeParamOk() (*PAMProviderTypeParameterResponse, bool)`

GetProviderTypeParamOk returns a tuple with the ProviderTypeParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderTypeParam

`func (o *PAMPamProviderTypeParamValueResponse) SetProviderTypeParam(v PAMProviderTypeParameterResponse)`

SetProviderTypeParam sets ProviderTypeParam field to given value.

### HasProviderTypeParam

`func (o *PAMPamProviderTypeParamValueResponse) HasProviderTypeParam() bool`

HasProviderTypeParam returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


