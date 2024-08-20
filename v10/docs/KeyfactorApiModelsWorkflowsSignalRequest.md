# KeyfactorApiModelsWorkflowsSignalRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SignalKey** | Pointer to **string** | The signal key. This is expected to be in a format like \&quot;STEP_NAME.SIGNAL_NAME\&quot; | [optional] 
**Data** | Pointer to **map[string]map[string]interface{}** | Arbitrary data to associate with the signal. | [optional] 

## Methods

### NewKeyfactorApiModelsWorkflowsSignalRequest

`func NewKeyfactorApiModelsWorkflowsSignalRequest() *KeyfactorApiModelsWorkflowsSignalRequest`

NewKeyfactorApiModelsWorkflowsSignalRequest instantiates a new KeyfactorApiModelsWorkflowsSignalRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorApiModelsWorkflowsSignalRequestWithDefaults

`func NewKeyfactorApiModelsWorkflowsSignalRequestWithDefaults() *KeyfactorApiModelsWorkflowsSignalRequest`

NewKeyfactorApiModelsWorkflowsSignalRequestWithDefaults instantiates a new KeyfactorApiModelsWorkflowsSignalRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignalKey

`func (o *KeyfactorApiModelsWorkflowsSignalRequest) GetSignalKey() string`

GetSignalKey returns the SignalKey field if non-nil, zero value otherwise.

### GetSignalKeyOk

`func (o *KeyfactorApiModelsWorkflowsSignalRequest) GetSignalKeyOk() (*string, bool)`

GetSignalKeyOk returns a tuple with the SignalKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalKey

`func (o *KeyfactorApiModelsWorkflowsSignalRequest) SetSignalKey(v string)`

SetSignalKey sets SignalKey field to given value.

### HasSignalKey

`func (o *KeyfactorApiModelsWorkflowsSignalRequest) HasSignalKey() bool`

HasSignalKey returns a boolean if a field has been set.

### GetData

`func (o *KeyfactorApiModelsWorkflowsSignalRequest) GetData() map[string]map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *KeyfactorApiModelsWorkflowsSignalRequest) GetDataOk() (*map[string]map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *KeyfactorApiModelsWorkflowsSignalRequest) SetData(v map[string]map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *KeyfactorApiModelsWorkflowsSignalRequest) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


