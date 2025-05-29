# WorkflowsConditionConfigurationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The Id of the condition. | [optional] 
**Value** | Pointer to **NullableString** | The value to compare to. This value will be compared to a true value. | [optional] 

## Methods

### NewWorkflowsConditionConfigurationResponse

`func NewWorkflowsConditionConfigurationResponse() *WorkflowsConditionConfigurationResponse`

NewWorkflowsConditionConfigurationResponse instantiates a new WorkflowsConditionConfigurationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowsConditionConfigurationResponseWithDefaults

`func NewWorkflowsConditionConfigurationResponseWithDefaults() *WorkflowsConditionConfigurationResponse`

NewWorkflowsConditionConfigurationResponseWithDefaults instantiates a new WorkflowsConditionConfigurationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkflowsConditionConfigurationResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowsConditionConfigurationResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowsConditionConfigurationResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WorkflowsConditionConfigurationResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetValue

`func (o *WorkflowsConditionConfigurationResponse) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *WorkflowsConditionConfigurationResponse) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *WorkflowsConditionConfigurationResponse) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *WorkflowsConditionConfigurationResponse) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *WorkflowsConditionConfigurationResponse) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *WorkflowsConditionConfigurationResponse) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


