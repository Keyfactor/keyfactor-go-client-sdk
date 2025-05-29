# CSSCMSDataModelModelsReenrollmentStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **bool** |  | [optional] 
**AgentId** | Pointer to **NullableString** |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**JobProperties** | Pointer to **NullableString** |  | [optional] 
**CustomAliasAllowed** | Pointer to **int32** |  | [optional] 
**EntryParameters** | Pointer to [**[]CSSCMSDataModelModelsCertificateStoreTypesCertificateStoreTypeEntryParameter**](CSSCMSDataModelModelsCertificateStoreTypesCertificateStoreTypeEntryParameter.md) |  | [optional] 

## Methods

### NewCSSCMSDataModelModelsReenrollmentStatus

`func NewCSSCMSDataModelModelsReenrollmentStatus() *CSSCMSDataModelModelsReenrollmentStatus`

NewCSSCMSDataModelModelsReenrollmentStatus instantiates a new CSSCMSDataModelModelsReenrollmentStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSSCMSDataModelModelsReenrollmentStatusWithDefaults

`func NewCSSCMSDataModelModelsReenrollmentStatusWithDefaults() *CSSCMSDataModelModelsReenrollmentStatus`

NewCSSCMSDataModelModelsReenrollmentStatusWithDefaults instantiates a new CSSCMSDataModelModelsReenrollmentStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CSSCMSDataModelModelsReenrollmentStatus) GetData() bool`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CSSCMSDataModelModelsReenrollmentStatus) GetDataOk() (*bool, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CSSCMSDataModelModelsReenrollmentStatus) SetData(v bool)`

SetData sets Data field to given value.

### HasData

`func (o *CSSCMSDataModelModelsReenrollmentStatus) HasData() bool`

HasData returns a boolean if a field has been set.

### GetAgentId

`func (o *CSSCMSDataModelModelsReenrollmentStatus) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *CSSCMSDataModelModelsReenrollmentStatus) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *CSSCMSDataModelModelsReenrollmentStatus) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *CSSCMSDataModelModelsReenrollmentStatus) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.

### SetAgentIdNil

`func (o *CSSCMSDataModelModelsReenrollmentStatus) SetAgentIdNil(b bool)`

 SetAgentIdNil sets the value for AgentId to be an explicit nil

### UnsetAgentId
`func (o *CSSCMSDataModelModelsReenrollmentStatus) UnsetAgentId()`

UnsetAgentId ensures that no value is present for AgentId, not even an explicit nil
### GetMessage

`func (o *CSSCMSDataModelModelsReenrollmentStatus) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CSSCMSDataModelModelsReenrollmentStatus) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CSSCMSDataModelModelsReenrollmentStatus) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CSSCMSDataModelModelsReenrollmentStatus) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *CSSCMSDataModelModelsReenrollmentStatus) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *CSSCMSDataModelModelsReenrollmentStatus) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetJobProperties

`func (o *CSSCMSDataModelModelsReenrollmentStatus) GetJobProperties() string`

GetJobProperties returns the JobProperties field if non-nil, zero value otherwise.

### GetJobPropertiesOk

`func (o *CSSCMSDataModelModelsReenrollmentStatus) GetJobPropertiesOk() (*string, bool)`

GetJobPropertiesOk returns a tuple with the JobProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobProperties

`func (o *CSSCMSDataModelModelsReenrollmentStatus) SetJobProperties(v string)`

SetJobProperties sets JobProperties field to given value.

### HasJobProperties

`func (o *CSSCMSDataModelModelsReenrollmentStatus) HasJobProperties() bool`

HasJobProperties returns a boolean if a field has been set.

### SetJobPropertiesNil

`func (o *CSSCMSDataModelModelsReenrollmentStatus) SetJobPropertiesNil(b bool)`

 SetJobPropertiesNil sets the value for JobProperties to be an explicit nil

### UnsetJobProperties
`func (o *CSSCMSDataModelModelsReenrollmentStatus) UnsetJobProperties()`

UnsetJobProperties ensures that no value is present for JobProperties, not even an explicit nil
### GetCustomAliasAllowed

`func (o *CSSCMSDataModelModelsReenrollmentStatus) GetCustomAliasAllowed() int32`

GetCustomAliasAllowed returns the CustomAliasAllowed field if non-nil, zero value otherwise.

### GetCustomAliasAllowedOk

`func (o *CSSCMSDataModelModelsReenrollmentStatus) GetCustomAliasAllowedOk() (*int32, bool)`

GetCustomAliasAllowedOk returns a tuple with the CustomAliasAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAliasAllowed

`func (o *CSSCMSDataModelModelsReenrollmentStatus) SetCustomAliasAllowed(v int32)`

SetCustomAliasAllowed sets CustomAliasAllowed field to given value.

### HasCustomAliasAllowed

`func (o *CSSCMSDataModelModelsReenrollmentStatus) HasCustomAliasAllowed() bool`

HasCustomAliasAllowed returns a boolean if a field has been set.

### GetEntryParameters

`func (o *CSSCMSDataModelModelsReenrollmentStatus) GetEntryParameters() []CSSCMSDataModelModelsCertificateStoreTypesCertificateStoreTypeEntryParameter`

GetEntryParameters returns the EntryParameters field if non-nil, zero value otherwise.

### GetEntryParametersOk

`func (o *CSSCMSDataModelModelsReenrollmentStatus) GetEntryParametersOk() (*[]CSSCMSDataModelModelsCertificateStoreTypesCertificateStoreTypeEntryParameter, bool)`

GetEntryParametersOk returns a tuple with the EntryParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryParameters

`func (o *CSSCMSDataModelModelsReenrollmentStatus) SetEntryParameters(v []CSSCMSDataModelModelsCertificateStoreTypesCertificateStoreTypeEntryParameter)`

SetEntryParameters sets EntryParameters field to given value.

### HasEntryParameters

`func (o *CSSCMSDataModelModelsReenrollmentStatus) HasEntryParameters() bool`

HasEntryParameters returns a boolean if a field has been set.

### SetEntryParametersNil

`func (o *CSSCMSDataModelModelsReenrollmentStatus) SetEntryParametersNil(b bool)`

 SetEntryParametersNil sets the value for EntryParameters to be an explicit nil

### UnsetEntryParameters
`func (o *CSSCMSDataModelModelsReenrollmentStatus) UnsetEntryParameters()`

UnsetEntryParameters ensures that no value is present for EntryParameters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


