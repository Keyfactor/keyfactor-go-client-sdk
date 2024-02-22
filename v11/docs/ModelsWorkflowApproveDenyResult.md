# ModelsWorkflowApproveDenyResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Failures** | Pointer to [**[]ModelsWorkflowProcessedCertificateRequest**](ModelsWorkflowProcessedCertificateRequest.md) |  | [optional] 
**Denials** | Pointer to [**[]ModelsWorkflowProcessedCertificateRequest**](ModelsWorkflowProcessedCertificateRequest.md) |  | [optional] 
**Successes** | Pointer to [**[]ModelsWorkflowProcessedCertificateRequest**](ModelsWorkflowProcessedCertificateRequest.md) |  | [optional] 

## Methods

### NewModelsWorkflowApproveDenyResult

`func NewModelsWorkflowApproveDenyResult() *ModelsWorkflowApproveDenyResult`

NewModelsWorkflowApproveDenyResult instantiates a new ModelsWorkflowApproveDenyResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsWorkflowApproveDenyResultWithDefaults

`func NewModelsWorkflowApproveDenyResultWithDefaults() *ModelsWorkflowApproveDenyResult`

NewModelsWorkflowApproveDenyResultWithDefaults instantiates a new ModelsWorkflowApproveDenyResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailures

`func (o *ModelsWorkflowApproveDenyResult) GetFailures() []ModelsWorkflowProcessedCertificateRequest`

GetFailures returns the Failures field if non-nil, zero value otherwise.

### GetFailuresOk

`func (o *ModelsWorkflowApproveDenyResult) GetFailuresOk() (*[]ModelsWorkflowProcessedCertificateRequest, bool)`

GetFailuresOk returns a tuple with the Failures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailures

`func (o *ModelsWorkflowApproveDenyResult) SetFailures(v []ModelsWorkflowProcessedCertificateRequest)`

SetFailures sets Failures field to given value.

### HasFailures

`func (o *ModelsWorkflowApproveDenyResult) HasFailures() bool`

HasFailures returns a boolean if a field has been set.

### SetFailuresNil

`func (o *ModelsWorkflowApproveDenyResult) SetFailuresNil(b bool)`

 SetFailuresNil sets the value for Failures to be an explicit nil

### UnsetFailures
`func (o *ModelsWorkflowApproveDenyResult) UnsetFailures()`

UnsetFailures ensures that no value is present for Failures, not even an explicit nil
### GetDenials

`func (o *ModelsWorkflowApproveDenyResult) GetDenials() []ModelsWorkflowProcessedCertificateRequest`

GetDenials returns the Denials field if non-nil, zero value otherwise.

### GetDenialsOk

`func (o *ModelsWorkflowApproveDenyResult) GetDenialsOk() (*[]ModelsWorkflowProcessedCertificateRequest, bool)`

GetDenialsOk returns a tuple with the Denials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenials

`func (o *ModelsWorkflowApproveDenyResult) SetDenials(v []ModelsWorkflowProcessedCertificateRequest)`

SetDenials sets Denials field to given value.

### HasDenials

`func (o *ModelsWorkflowApproveDenyResult) HasDenials() bool`

HasDenials returns a boolean if a field has been set.

### SetDenialsNil

`func (o *ModelsWorkflowApproveDenyResult) SetDenialsNil(b bool)`

 SetDenialsNil sets the value for Denials to be an explicit nil

### UnsetDenials
`func (o *ModelsWorkflowApproveDenyResult) UnsetDenials()`

UnsetDenials ensures that no value is present for Denials, not even an explicit nil
### GetSuccesses

`func (o *ModelsWorkflowApproveDenyResult) GetSuccesses() []ModelsWorkflowProcessedCertificateRequest`

GetSuccesses returns the Successes field if non-nil, zero value otherwise.

### GetSuccessesOk

`func (o *ModelsWorkflowApproveDenyResult) GetSuccessesOk() (*[]ModelsWorkflowProcessedCertificateRequest, bool)`

GetSuccessesOk returns a tuple with the Successes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccesses

`func (o *ModelsWorkflowApproveDenyResult) SetSuccesses(v []ModelsWorkflowProcessedCertificateRequest)`

SetSuccesses sets Successes field to given value.

### HasSuccesses

`func (o *ModelsWorkflowApproveDenyResult) HasSuccesses() bool`

HasSuccesses returns a boolean if a field has been set.

### SetSuccessesNil

`func (o *ModelsWorkflowApproveDenyResult) SetSuccessesNil(b bool)`

 SetSuccessesNil sets the value for Successes to be an explicit nil

### UnsetSuccesses
`func (o *ModelsWorkflowApproveDenyResult) UnsetSuccesses()`

UnsetSuccesses ensures that no value is present for Successes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


