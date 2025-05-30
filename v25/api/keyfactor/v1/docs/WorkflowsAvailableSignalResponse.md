# WorkflowsAvailableSignalResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SignalName** | Pointer to **NullableString** | The name of the signal. | [optional] 
**StepSignalId** | Pointer to **string** | The signal Id. | [optional] 
**SignalReceived** | Pointer to **bool** | Whether or not the signal has been received. | [optional] 

## Methods

### NewWorkflowsAvailableSignalResponse

`func NewWorkflowsAvailableSignalResponse() *WorkflowsAvailableSignalResponse`

NewWorkflowsAvailableSignalResponse instantiates a new WorkflowsAvailableSignalResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowsAvailableSignalResponseWithDefaults

`func NewWorkflowsAvailableSignalResponseWithDefaults() *WorkflowsAvailableSignalResponse`

NewWorkflowsAvailableSignalResponseWithDefaults instantiates a new WorkflowsAvailableSignalResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignalName

`func (o *WorkflowsAvailableSignalResponse) GetSignalName() string`

GetSignalName returns the SignalName field if non-nil, zero value otherwise.

### GetSignalNameOk

`func (o *WorkflowsAvailableSignalResponse) GetSignalNameOk() (*string, bool)`

GetSignalNameOk returns a tuple with the SignalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalName

`func (o *WorkflowsAvailableSignalResponse) SetSignalName(v string)`

SetSignalName sets SignalName field to given value.

### HasSignalName

`func (o *WorkflowsAvailableSignalResponse) HasSignalName() bool`

HasSignalName returns a boolean if a field has been set.

### SetSignalNameNil

`func (o *WorkflowsAvailableSignalResponse) SetSignalNameNil(b bool)`

 SetSignalNameNil sets the value for SignalName to be an explicit nil

### UnsetSignalName
`func (o *WorkflowsAvailableSignalResponse) UnsetSignalName()`

UnsetSignalName ensures that no value is present for SignalName, not even an explicit nil
### GetStepSignalId

`func (o *WorkflowsAvailableSignalResponse) GetStepSignalId() string`

GetStepSignalId returns the StepSignalId field if non-nil, zero value otherwise.

### GetStepSignalIdOk

`func (o *WorkflowsAvailableSignalResponse) GetStepSignalIdOk() (*string, bool)`

GetStepSignalIdOk returns a tuple with the StepSignalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepSignalId

`func (o *WorkflowsAvailableSignalResponse) SetStepSignalId(v string)`

SetStepSignalId sets StepSignalId field to given value.

### HasStepSignalId

`func (o *WorkflowsAvailableSignalResponse) HasStepSignalId() bool`

HasStepSignalId returns a boolean if a field has been set.

### GetSignalReceived

`func (o *WorkflowsAvailableSignalResponse) GetSignalReceived() bool`

GetSignalReceived returns the SignalReceived field if non-nil, zero value otherwise.

### GetSignalReceivedOk

`func (o *WorkflowsAvailableSignalResponse) GetSignalReceivedOk() (*bool, bool)`

GetSignalReceivedOk returns a tuple with the SignalReceived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalReceived

`func (o *WorkflowsAvailableSignalResponse) SetSignalReceived(v bool)`

SetSignalReceived sets SignalReceived field to given value.

### HasSignalReceived

`func (o *WorkflowsAvailableSignalResponse) HasSignalReceived() bool`

HasSignalReceived returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


