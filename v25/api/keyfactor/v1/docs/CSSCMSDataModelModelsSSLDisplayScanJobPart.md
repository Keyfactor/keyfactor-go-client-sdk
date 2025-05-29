# CSSCMSDataModelModelsSSLDisplayScanJobPart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScanJobPartId** | Pointer to **string** |  | [optional] 
**Agent** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to [**CSSCMSCoreEnumsSslScanJobStatus**](CSSCMSCoreEnumsSslScanJobStatus.md) |  | [optional] 
**StartTime** | Pointer to **NullableTime** |  | [optional] 
**EndTime** | Pointer to **NullableTime** |  | [optional] 
**EndpointCount** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewCSSCMSDataModelModelsSSLDisplayScanJobPart

`func NewCSSCMSDataModelModelsSSLDisplayScanJobPart() *CSSCMSDataModelModelsSSLDisplayScanJobPart`

NewCSSCMSDataModelModelsSSLDisplayScanJobPart instantiates a new CSSCMSDataModelModelsSSLDisplayScanJobPart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSSCMSDataModelModelsSSLDisplayScanJobPartWithDefaults

`func NewCSSCMSDataModelModelsSSLDisplayScanJobPartWithDefaults() *CSSCMSDataModelModelsSSLDisplayScanJobPart`

NewCSSCMSDataModelModelsSSLDisplayScanJobPartWithDefaults instantiates a new CSSCMSDataModelModelsSSLDisplayScanJobPart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScanJobPartId

`func (o *CSSCMSDataModelModelsSSLDisplayScanJobPart) GetScanJobPartId() string`

GetScanJobPartId returns the ScanJobPartId field if non-nil, zero value otherwise.

### GetScanJobPartIdOk

`func (o *CSSCMSDataModelModelsSSLDisplayScanJobPart) GetScanJobPartIdOk() (*string, bool)`

GetScanJobPartIdOk returns a tuple with the ScanJobPartId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanJobPartId

`func (o *CSSCMSDataModelModelsSSLDisplayScanJobPart) SetScanJobPartId(v string)`

SetScanJobPartId sets ScanJobPartId field to given value.

### HasScanJobPartId

`func (o *CSSCMSDataModelModelsSSLDisplayScanJobPart) HasScanJobPartId() bool`

HasScanJobPartId returns a boolean if a field has been set.

### GetAgent

`func (o *CSSCMSDataModelModelsSSLDisplayScanJobPart) GetAgent() string`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *CSSCMSDataModelModelsSSLDisplayScanJobPart) GetAgentOk() (*string, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *CSSCMSDataModelModelsSSLDisplayScanJobPart) SetAgent(v string)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *CSSCMSDataModelModelsSSLDisplayScanJobPart) HasAgent() bool`

HasAgent returns a boolean if a field has been set.

### SetAgentNil

`func (o *CSSCMSDataModelModelsSSLDisplayScanJobPart) SetAgentNil(b bool)`

 SetAgentNil sets the value for Agent to be an explicit nil

### UnsetAgent
`func (o *CSSCMSDataModelModelsSSLDisplayScanJobPart) UnsetAgent()`

UnsetAgent ensures that no value is present for Agent, not even an explicit nil
### GetStatus

`func (o *CSSCMSDataModelModelsSSLDisplayScanJobPart) GetStatus() CSSCMSCoreEnumsSslScanJobStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CSSCMSDataModelModelsSSLDisplayScanJobPart) GetStatusOk() (*CSSCMSCoreEnumsSslScanJobStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CSSCMSDataModelModelsSSLDisplayScanJobPart) SetStatus(v CSSCMSCoreEnumsSslScanJobStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CSSCMSDataModelModelsSSLDisplayScanJobPart) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStartTime

`func (o *CSSCMSDataModelModelsSSLDisplayScanJobPart) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *CSSCMSDataModelModelsSSLDisplayScanJobPart) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *CSSCMSDataModelModelsSSLDisplayScanJobPart) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *CSSCMSDataModelModelsSSLDisplayScanJobPart) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTimeNil

`func (o *CSSCMSDataModelModelsSSLDisplayScanJobPart) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *CSSCMSDataModelModelsSSLDisplayScanJobPart) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
### GetEndTime

`func (o *CSSCMSDataModelModelsSSLDisplayScanJobPart) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *CSSCMSDataModelModelsSSLDisplayScanJobPart) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *CSSCMSDataModelModelsSSLDisplayScanJobPart) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *CSSCMSDataModelModelsSSLDisplayScanJobPart) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### SetEndTimeNil

`func (o *CSSCMSDataModelModelsSSLDisplayScanJobPart) SetEndTimeNil(b bool)`

 SetEndTimeNil sets the value for EndTime to be an explicit nil

### UnsetEndTime
`func (o *CSSCMSDataModelModelsSSLDisplayScanJobPart) UnsetEndTime()`

UnsetEndTime ensures that no value is present for EndTime, not even an explicit nil
### GetEndpointCount

`func (o *CSSCMSDataModelModelsSSLDisplayScanJobPart) GetEndpointCount() int32`

GetEndpointCount returns the EndpointCount field if non-nil, zero value otherwise.

### GetEndpointCountOk

`func (o *CSSCMSDataModelModelsSSLDisplayScanJobPart) GetEndpointCountOk() (*int32, bool)`

GetEndpointCountOk returns a tuple with the EndpointCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointCount

`func (o *CSSCMSDataModelModelsSSLDisplayScanJobPart) SetEndpointCount(v int32)`

SetEndpointCount sets EndpointCount field to given value.

### HasEndpointCount

`func (o *CSSCMSDataModelModelsSSLDisplayScanJobPart) HasEndpointCount() bool`

HasEndpointCount returns a boolean if a field has been set.

### SetEndpointCountNil

`func (o *CSSCMSDataModelModelsSSLDisplayScanJobPart) SetEndpointCountNil(b bool)`

 SetEndpointCountNil sets the value for EndpointCount to be an explicit nil

### UnsetEndpointCount
`func (o *CSSCMSDataModelModelsSSLDisplayScanJobPart) UnsetEndpointCount()`

UnsetEndpointCount ensures that no value is present for EndpointCount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


