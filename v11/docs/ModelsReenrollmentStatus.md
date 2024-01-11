# ModelsReenrollmentStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **bool** |  | [optional] 
**AgentId** | Pointer to **NullableString** |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**JobProperties** | Pointer to **NullableString** |  | [optional] 
**CustomAliasAllowed** | Pointer to **int32** |  | [optional] 
**EntryParameters** | Pointer to [**[]ModelsCertificateStoreTypesCertificateStoreTypeEntryParameter**](ModelsCertificateStoreTypesCertificateStoreTypeEntryParameter.md) |  | [optional] 

## Methods

### NewModelsReenrollmentStatus

`func NewModelsReenrollmentStatus() *ModelsReenrollmentStatus`

NewModelsReenrollmentStatus instantiates a new ModelsReenrollmentStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsReenrollmentStatusWithDefaults

`func NewModelsReenrollmentStatusWithDefaults() *ModelsReenrollmentStatus`

NewModelsReenrollmentStatusWithDefaults instantiates a new ModelsReenrollmentStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ModelsReenrollmentStatus) GetData() bool`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ModelsReenrollmentStatus) GetDataOk() (*bool, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ModelsReenrollmentStatus) SetData(v bool)`

SetData sets Data field to given value.

### HasData

`func (o *ModelsReenrollmentStatus) HasData() bool`

HasData returns a boolean if a field has been set.

### GetAgentId

`func (o *ModelsReenrollmentStatus) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *ModelsReenrollmentStatus) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *ModelsReenrollmentStatus) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *ModelsReenrollmentStatus) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.

### SetAgentIdNil

`func (o *ModelsReenrollmentStatus) SetAgentIdNil(b bool)`

 SetAgentIdNil sets the value for AgentId to be an explicit nil

### UnsetAgentId
`func (o *ModelsReenrollmentStatus) UnsetAgentId()`

UnsetAgentId ensures that no value is present for AgentId, not even an explicit nil
### GetMessage

`func (o *ModelsReenrollmentStatus) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ModelsReenrollmentStatus) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ModelsReenrollmentStatus) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ModelsReenrollmentStatus) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *ModelsReenrollmentStatus) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *ModelsReenrollmentStatus) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetJobProperties

`func (o *ModelsReenrollmentStatus) GetJobProperties() string`

GetJobProperties returns the JobProperties field if non-nil, zero value otherwise.

### GetJobPropertiesOk

`func (o *ModelsReenrollmentStatus) GetJobPropertiesOk() (*string, bool)`

GetJobPropertiesOk returns a tuple with the JobProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobProperties

`func (o *ModelsReenrollmentStatus) SetJobProperties(v string)`

SetJobProperties sets JobProperties field to given value.

### HasJobProperties

`func (o *ModelsReenrollmentStatus) HasJobProperties() bool`

HasJobProperties returns a boolean if a field has been set.

### SetJobPropertiesNil

`func (o *ModelsReenrollmentStatus) SetJobPropertiesNil(b bool)`

 SetJobPropertiesNil sets the value for JobProperties to be an explicit nil

### UnsetJobProperties
`func (o *ModelsReenrollmentStatus) UnsetJobProperties()`

UnsetJobProperties ensures that no value is present for JobProperties, not even an explicit nil
### GetCustomAliasAllowed

`func (o *ModelsReenrollmentStatus) GetCustomAliasAllowed() int32`

GetCustomAliasAllowed returns the CustomAliasAllowed field if non-nil, zero value otherwise.

### GetCustomAliasAllowedOk

`func (o *ModelsReenrollmentStatus) GetCustomAliasAllowedOk() (*int32, bool)`

GetCustomAliasAllowedOk returns a tuple with the CustomAliasAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAliasAllowed

`func (o *ModelsReenrollmentStatus) SetCustomAliasAllowed(v int32)`

SetCustomAliasAllowed sets CustomAliasAllowed field to given value.

### HasCustomAliasAllowed

`func (o *ModelsReenrollmentStatus) HasCustomAliasAllowed() bool`

HasCustomAliasAllowed returns a boolean if a field has been set.

### GetEntryParameters

`func (o *ModelsReenrollmentStatus) GetEntryParameters() []ModelsCertificateStoreTypesCertificateStoreTypeEntryParameter`

GetEntryParameters returns the EntryParameters field if non-nil, zero value otherwise.

### GetEntryParametersOk

`func (o *ModelsReenrollmentStatus) GetEntryParametersOk() (*[]ModelsCertificateStoreTypesCertificateStoreTypeEntryParameter, bool)`

GetEntryParametersOk returns a tuple with the EntryParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryParameters

`func (o *ModelsReenrollmentStatus) SetEntryParameters(v []ModelsCertificateStoreTypesCertificateStoreTypeEntryParameter)`

SetEntryParameters sets EntryParameters field to given value.

### HasEntryParameters

`func (o *ModelsReenrollmentStatus) HasEntryParameters() bool`

HasEntryParameters returns a boolean if a field has been set.

### SetEntryParametersNil

`func (o *ModelsReenrollmentStatus) SetEntryParametersNil(b bool)`

 SetEntryParametersNil sets the value for EntryParameters to be an explicit nil

### UnsetEntryParameters
`func (o *ModelsReenrollmentStatus) UnsetEntryParameters()`

UnsetEntryParameters ensures that no value is present for EntryParameters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


