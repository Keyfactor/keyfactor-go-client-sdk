# KeyfactorWebKeyfactorApiModelsWorkflowsSignalRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SignalKey** | Pointer to **NullableString** | The signal key. This is expected to be in a format like \&quot;STEP_NAME.SIGNAL_NAME\&quot; | [optional] 
**Data** | Pointer to **map[string]interface{}** | Arbitrary data to associate with the signal. | [optional] 

## Methods

### NewKeyfactorWebKeyfactorApiModelsWorkflowsSignalRequest

`func NewKeyfactorWebKeyfactorApiModelsWorkflowsSignalRequest() *KeyfactorWebKeyfactorApiModelsWorkflowsSignalRequest`

NewKeyfactorWebKeyfactorApiModelsWorkflowsSignalRequest instantiates a new KeyfactorWebKeyfactorApiModelsWorkflowsSignalRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorWebKeyfactorApiModelsWorkflowsSignalRequestWithDefaults

`func NewKeyfactorWebKeyfactorApiModelsWorkflowsSignalRequestWithDefaults() *KeyfactorWebKeyfactorApiModelsWorkflowsSignalRequest`

NewKeyfactorWebKeyfactorApiModelsWorkflowsSignalRequestWithDefaults instantiates a new KeyfactorWebKeyfactorApiModelsWorkflowsSignalRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignalKey

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsSignalRequest) GetSignalKey() string`

GetSignalKey returns the SignalKey field if non-nil, zero value otherwise.

### GetSignalKeyOk

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsSignalRequest) GetSignalKeyOk() (*string, bool)`

GetSignalKeyOk returns a tuple with the SignalKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalKey

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsSignalRequest) SetSignalKey(v string)`

SetSignalKey sets SignalKey field to given value.

### HasSignalKey

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsSignalRequest) HasSignalKey() bool`

HasSignalKey returns a boolean if a field has been set.

### SetSignalKeyNil

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsSignalRequest) SetSignalKeyNil(b bool)`

 SetSignalKeyNil sets the value for SignalKey to be an explicit nil

### UnsetSignalKey
`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsSignalRequest) UnsetSignalKey()`

UnsetSignalKey ensures that no value is present for SignalKey, not even an explicit nil
### GetData

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsSignalRequest) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsSignalRequest) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsSignalRequest) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsSignalRequest) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsSignalRequest) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsSignalRequest) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


