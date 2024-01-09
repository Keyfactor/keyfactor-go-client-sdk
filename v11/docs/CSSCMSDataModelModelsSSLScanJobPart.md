# CSSCMSDataModelModelsSSLScanJobPart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScanJobPartId** | Pointer to **string** |  | [optional] 
**LogicalScanJobId** | Pointer to **string** |  | [optional] 
**AgentJobId** | Pointer to **string** |  | [optional] 
**EstimatedEndpointCount** | Pointer to **NullableInt32** |  | [optional] 
**Status** | Pointer to [**CSSCMSCoreEnumsSslScanJobStatus**](CSSCMSCoreEnumsSslScanJobStatus.md) |  | [optional] 
**StatTotalEndpointCount** | Pointer to **NullableInt32** |  | [optional] 
**StatTimedOutConnectingCount** | Pointer to **NullableInt32** |  | [optional] 
**StatConnectionRefusedCount** | Pointer to **NullableInt32** |  | [optional] 
**StatTimedOutDownloadingCount** | Pointer to **NullableInt32** |  | [optional] 
**StatExceptionDownloadingCount** | Pointer to **NullableInt32** |  | [optional] 
**StatNotSslCount** | Pointer to **NullableInt32** |  | [optional] 
**StatBadSslHandshakeCount** | Pointer to **NullableInt32** |  | [optional] 
**StatCertificateFoundCount** | Pointer to **NullableInt32** |  | [optional] 
**StatNoCertificateCount** | Pointer to **NullableInt32** |  | [optional] 
**ScanJobPartDefinitions** | Pointer to [**[]CSSCMSDataModelModelsSSLScanJobPartDefinition**](CSSCMSDataModelModelsSSLScanJobPartDefinition.md) |  | [optional] 
**StartTime** | Pointer to **NullableTime** |  | [optional] 
**EndTime** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewCSSCMSDataModelModelsSSLScanJobPart

`func NewCSSCMSDataModelModelsSSLScanJobPart() *CSSCMSDataModelModelsSSLScanJobPart`

NewCSSCMSDataModelModelsSSLScanJobPart instantiates a new CSSCMSDataModelModelsSSLScanJobPart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSSCMSDataModelModelsSSLScanJobPartWithDefaults

`func NewCSSCMSDataModelModelsSSLScanJobPartWithDefaults() *CSSCMSDataModelModelsSSLScanJobPart`

NewCSSCMSDataModelModelsSSLScanJobPartWithDefaults instantiates a new CSSCMSDataModelModelsSSLScanJobPart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScanJobPartId

`func (o *CSSCMSDataModelModelsSSLScanJobPart) GetScanJobPartId() string`

GetScanJobPartId returns the ScanJobPartId field if non-nil, zero value otherwise.

### GetScanJobPartIdOk

`func (o *CSSCMSDataModelModelsSSLScanJobPart) GetScanJobPartIdOk() (*string, bool)`

GetScanJobPartIdOk returns a tuple with the ScanJobPartId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanJobPartId

`func (o *CSSCMSDataModelModelsSSLScanJobPart) SetScanJobPartId(v string)`

SetScanJobPartId sets ScanJobPartId field to given value.

### HasScanJobPartId

`func (o *CSSCMSDataModelModelsSSLScanJobPart) HasScanJobPartId() bool`

HasScanJobPartId returns a boolean if a field has been set.

### GetLogicalScanJobId

`func (o *CSSCMSDataModelModelsSSLScanJobPart) GetLogicalScanJobId() string`

GetLogicalScanJobId returns the LogicalScanJobId field if non-nil, zero value otherwise.

### GetLogicalScanJobIdOk

`func (o *CSSCMSDataModelModelsSSLScanJobPart) GetLogicalScanJobIdOk() (*string, bool)`

GetLogicalScanJobIdOk returns a tuple with the LogicalScanJobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalScanJobId

`func (o *CSSCMSDataModelModelsSSLScanJobPart) SetLogicalScanJobId(v string)`

SetLogicalScanJobId sets LogicalScanJobId field to given value.

### HasLogicalScanJobId

`func (o *CSSCMSDataModelModelsSSLScanJobPart) HasLogicalScanJobId() bool`

HasLogicalScanJobId returns a boolean if a field has been set.

### GetAgentJobId

`func (o *CSSCMSDataModelModelsSSLScanJobPart) GetAgentJobId() string`

GetAgentJobId returns the AgentJobId field if non-nil, zero value otherwise.

### GetAgentJobIdOk

`func (o *CSSCMSDataModelModelsSSLScanJobPart) GetAgentJobIdOk() (*string, bool)`

GetAgentJobIdOk returns a tuple with the AgentJobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentJobId

`func (o *CSSCMSDataModelModelsSSLScanJobPart) SetAgentJobId(v string)`

SetAgentJobId sets AgentJobId field to given value.

### HasAgentJobId

`func (o *CSSCMSDataModelModelsSSLScanJobPart) HasAgentJobId() bool`

HasAgentJobId returns a boolean if a field has been set.

### GetEstimatedEndpointCount

`func (o *CSSCMSDataModelModelsSSLScanJobPart) GetEstimatedEndpointCount() int32`

GetEstimatedEndpointCount returns the EstimatedEndpointCount field if non-nil, zero value otherwise.

### GetEstimatedEndpointCountOk

`func (o *CSSCMSDataModelModelsSSLScanJobPart) GetEstimatedEndpointCountOk() (*int32, bool)`

GetEstimatedEndpointCountOk returns a tuple with the EstimatedEndpointCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedEndpointCount

`func (o *CSSCMSDataModelModelsSSLScanJobPart) SetEstimatedEndpointCount(v int32)`

SetEstimatedEndpointCount sets EstimatedEndpointCount field to given value.

### HasEstimatedEndpointCount

`func (o *CSSCMSDataModelModelsSSLScanJobPart) HasEstimatedEndpointCount() bool`

HasEstimatedEndpointCount returns a boolean if a field has been set.

### SetEstimatedEndpointCountNil

`func (o *CSSCMSDataModelModelsSSLScanJobPart) SetEstimatedEndpointCountNil(b bool)`

 SetEstimatedEndpointCountNil sets the value for EstimatedEndpointCount to be an explicit nil

### UnsetEstimatedEndpointCount
`func (o *CSSCMSDataModelModelsSSLScanJobPart) UnsetEstimatedEndpointCount()`

UnsetEstimatedEndpointCount ensures that no value is present for EstimatedEndpointCount, not even an explicit nil
### GetStatus

`func (o *CSSCMSDataModelModelsSSLScanJobPart) GetStatus() CSSCMSCoreEnumsSslScanJobStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CSSCMSDataModelModelsSSLScanJobPart) GetStatusOk() (*CSSCMSCoreEnumsSslScanJobStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CSSCMSDataModelModelsSSLScanJobPart) SetStatus(v CSSCMSCoreEnumsSslScanJobStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CSSCMSDataModelModelsSSLScanJobPart) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatTotalEndpointCount

`func (o *CSSCMSDataModelModelsSSLScanJobPart) GetStatTotalEndpointCount() int32`

GetStatTotalEndpointCount returns the StatTotalEndpointCount field if non-nil, zero value otherwise.

### GetStatTotalEndpointCountOk

`func (o *CSSCMSDataModelModelsSSLScanJobPart) GetStatTotalEndpointCountOk() (*int32, bool)`

GetStatTotalEndpointCountOk returns a tuple with the StatTotalEndpointCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatTotalEndpointCount

`func (o *CSSCMSDataModelModelsSSLScanJobPart) SetStatTotalEndpointCount(v int32)`

SetStatTotalEndpointCount sets StatTotalEndpointCount field to given value.

### HasStatTotalEndpointCount

`func (o *CSSCMSDataModelModelsSSLScanJobPart) HasStatTotalEndpointCount() bool`

HasStatTotalEndpointCount returns a boolean if a field has been set.

### SetStatTotalEndpointCountNil

`func (o *CSSCMSDataModelModelsSSLScanJobPart) SetStatTotalEndpointCountNil(b bool)`

 SetStatTotalEndpointCountNil sets the value for StatTotalEndpointCount to be an explicit nil

### UnsetStatTotalEndpointCount
`func (o *CSSCMSDataModelModelsSSLScanJobPart) UnsetStatTotalEndpointCount()`

UnsetStatTotalEndpointCount ensures that no value is present for StatTotalEndpointCount, not even an explicit nil
### GetStatTimedOutConnectingCount

`func (o *CSSCMSDataModelModelsSSLScanJobPart) GetStatTimedOutConnectingCount() int32`

GetStatTimedOutConnectingCount returns the StatTimedOutConnectingCount field if non-nil, zero value otherwise.

### GetStatTimedOutConnectingCountOk

`func (o *CSSCMSDataModelModelsSSLScanJobPart) GetStatTimedOutConnectingCountOk() (*int32, bool)`

GetStatTimedOutConnectingCountOk returns a tuple with the StatTimedOutConnectingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatTimedOutConnectingCount

`func (o *CSSCMSDataModelModelsSSLScanJobPart) SetStatTimedOutConnectingCount(v int32)`

SetStatTimedOutConnectingCount sets StatTimedOutConnectingCount field to given value.

### HasStatTimedOutConnectingCount

`func (o *CSSCMSDataModelModelsSSLScanJobPart) HasStatTimedOutConnectingCount() bool`

HasStatTimedOutConnectingCount returns a boolean if a field has been set.

### SetStatTimedOutConnectingCountNil

`func (o *CSSCMSDataModelModelsSSLScanJobPart) SetStatTimedOutConnectingCountNil(b bool)`

 SetStatTimedOutConnectingCountNil sets the value for StatTimedOutConnectingCount to be an explicit nil

### UnsetStatTimedOutConnectingCount
`func (o *CSSCMSDataModelModelsSSLScanJobPart) UnsetStatTimedOutConnectingCount()`

UnsetStatTimedOutConnectingCount ensures that no value is present for StatTimedOutConnectingCount, not even an explicit nil
### GetStatConnectionRefusedCount

`func (o *CSSCMSDataModelModelsSSLScanJobPart) GetStatConnectionRefusedCount() int32`

GetStatConnectionRefusedCount returns the StatConnectionRefusedCount field if non-nil, zero value otherwise.

### GetStatConnectionRefusedCountOk

`func (o *CSSCMSDataModelModelsSSLScanJobPart) GetStatConnectionRefusedCountOk() (*int32, bool)`

GetStatConnectionRefusedCountOk returns a tuple with the StatConnectionRefusedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatConnectionRefusedCount

`func (o *CSSCMSDataModelModelsSSLScanJobPart) SetStatConnectionRefusedCount(v int32)`

SetStatConnectionRefusedCount sets StatConnectionRefusedCount field to given value.

### HasStatConnectionRefusedCount

`func (o *CSSCMSDataModelModelsSSLScanJobPart) HasStatConnectionRefusedCount() bool`

HasStatConnectionRefusedCount returns a boolean if a field has been set.

### SetStatConnectionRefusedCountNil

`func (o *CSSCMSDataModelModelsSSLScanJobPart) SetStatConnectionRefusedCountNil(b bool)`

 SetStatConnectionRefusedCountNil sets the value for StatConnectionRefusedCount to be an explicit nil

### UnsetStatConnectionRefusedCount
`func (o *CSSCMSDataModelModelsSSLScanJobPart) UnsetStatConnectionRefusedCount()`

UnsetStatConnectionRefusedCount ensures that no value is present for StatConnectionRefusedCount, not even an explicit nil
### GetStatTimedOutDownloadingCount

`func (o *CSSCMSDataModelModelsSSLScanJobPart) GetStatTimedOutDownloadingCount() int32`

GetStatTimedOutDownloadingCount returns the StatTimedOutDownloadingCount field if non-nil, zero value otherwise.

### GetStatTimedOutDownloadingCountOk

`func (o *CSSCMSDataModelModelsSSLScanJobPart) GetStatTimedOutDownloadingCountOk() (*int32, bool)`

GetStatTimedOutDownloadingCountOk returns a tuple with the StatTimedOutDownloadingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatTimedOutDownloadingCount

`func (o *CSSCMSDataModelModelsSSLScanJobPart) SetStatTimedOutDownloadingCount(v int32)`

SetStatTimedOutDownloadingCount sets StatTimedOutDownloadingCount field to given value.

### HasStatTimedOutDownloadingCount

`func (o *CSSCMSDataModelModelsSSLScanJobPart) HasStatTimedOutDownloadingCount() bool`

HasStatTimedOutDownloadingCount returns a boolean if a field has been set.

### SetStatTimedOutDownloadingCountNil

`func (o *CSSCMSDataModelModelsSSLScanJobPart) SetStatTimedOutDownloadingCountNil(b bool)`

 SetStatTimedOutDownloadingCountNil sets the value for StatTimedOutDownloadingCount to be an explicit nil

### UnsetStatTimedOutDownloadingCount
`func (o *CSSCMSDataModelModelsSSLScanJobPart) UnsetStatTimedOutDownloadingCount()`

UnsetStatTimedOutDownloadingCount ensures that no value is present for StatTimedOutDownloadingCount, not even an explicit nil
### GetStatExceptionDownloadingCount

`func (o *CSSCMSDataModelModelsSSLScanJobPart) GetStatExceptionDownloadingCount() int32`

GetStatExceptionDownloadingCount returns the StatExceptionDownloadingCount field if non-nil, zero value otherwise.

### GetStatExceptionDownloadingCountOk

`func (o *CSSCMSDataModelModelsSSLScanJobPart) GetStatExceptionDownloadingCountOk() (*int32, bool)`

GetStatExceptionDownloadingCountOk returns a tuple with the StatExceptionDownloadingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatExceptionDownloadingCount

`func (o *CSSCMSDataModelModelsSSLScanJobPart) SetStatExceptionDownloadingCount(v int32)`

SetStatExceptionDownloadingCount sets StatExceptionDownloadingCount field to given value.

### HasStatExceptionDownloadingCount

`func (o *CSSCMSDataModelModelsSSLScanJobPart) HasStatExceptionDownloadingCount() bool`

HasStatExceptionDownloadingCount returns a boolean if a field has been set.

### SetStatExceptionDownloadingCountNil

`func (o *CSSCMSDataModelModelsSSLScanJobPart) SetStatExceptionDownloadingCountNil(b bool)`

 SetStatExceptionDownloadingCountNil sets the value for StatExceptionDownloadingCount to be an explicit nil

### UnsetStatExceptionDownloadingCount
`func (o *CSSCMSDataModelModelsSSLScanJobPart) UnsetStatExceptionDownloadingCount()`

UnsetStatExceptionDownloadingCount ensures that no value is present for StatExceptionDownloadingCount, not even an explicit nil
### GetStatNotSslCount

`func (o *CSSCMSDataModelModelsSSLScanJobPart) GetStatNotSslCount() int32`

GetStatNotSslCount returns the StatNotSslCount field if non-nil, zero value otherwise.

### GetStatNotSslCountOk

`func (o *CSSCMSDataModelModelsSSLScanJobPart) GetStatNotSslCountOk() (*int32, bool)`

GetStatNotSslCountOk returns a tuple with the StatNotSslCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatNotSslCount

`func (o *CSSCMSDataModelModelsSSLScanJobPart) SetStatNotSslCount(v int32)`

SetStatNotSslCount sets StatNotSslCount field to given value.

### HasStatNotSslCount

`func (o *CSSCMSDataModelModelsSSLScanJobPart) HasStatNotSslCount() bool`

HasStatNotSslCount returns a boolean if a field has been set.

### SetStatNotSslCountNil

`func (o *CSSCMSDataModelModelsSSLScanJobPart) SetStatNotSslCountNil(b bool)`

 SetStatNotSslCountNil sets the value for StatNotSslCount to be an explicit nil

### UnsetStatNotSslCount
`func (o *CSSCMSDataModelModelsSSLScanJobPart) UnsetStatNotSslCount()`

UnsetStatNotSslCount ensures that no value is present for StatNotSslCount, not even an explicit nil
### GetStatBadSslHandshakeCount

`func (o *CSSCMSDataModelModelsSSLScanJobPart) GetStatBadSslHandshakeCount() int32`

GetStatBadSslHandshakeCount returns the StatBadSslHandshakeCount field if non-nil, zero value otherwise.

### GetStatBadSslHandshakeCountOk

`func (o *CSSCMSDataModelModelsSSLScanJobPart) GetStatBadSslHandshakeCountOk() (*int32, bool)`

GetStatBadSslHandshakeCountOk returns a tuple with the StatBadSslHandshakeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatBadSslHandshakeCount

`func (o *CSSCMSDataModelModelsSSLScanJobPart) SetStatBadSslHandshakeCount(v int32)`

SetStatBadSslHandshakeCount sets StatBadSslHandshakeCount field to given value.

### HasStatBadSslHandshakeCount

`func (o *CSSCMSDataModelModelsSSLScanJobPart) HasStatBadSslHandshakeCount() bool`

HasStatBadSslHandshakeCount returns a boolean if a field has been set.

### SetStatBadSslHandshakeCountNil

`func (o *CSSCMSDataModelModelsSSLScanJobPart) SetStatBadSslHandshakeCountNil(b bool)`

 SetStatBadSslHandshakeCountNil sets the value for StatBadSslHandshakeCount to be an explicit nil

### UnsetStatBadSslHandshakeCount
`func (o *CSSCMSDataModelModelsSSLScanJobPart) UnsetStatBadSslHandshakeCount()`

UnsetStatBadSslHandshakeCount ensures that no value is present for StatBadSslHandshakeCount, not even an explicit nil
### GetStatCertificateFoundCount

`func (o *CSSCMSDataModelModelsSSLScanJobPart) GetStatCertificateFoundCount() int32`

GetStatCertificateFoundCount returns the StatCertificateFoundCount field if non-nil, zero value otherwise.

### GetStatCertificateFoundCountOk

`func (o *CSSCMSDataModelModelsSSLScanJobPart) GetStatCertificateFoundCountOk() (*int32, bool)`

GetStatCertificateFoundCountOk returns a tuple with the StatCertificateFoundCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatCertificateFoundCount

`func (o *CSSCMSDataModelModelsSSLScanJobPart) SetStatCertificateFoundCount(v int32)`

SetStatCertificateFoundCount sets StatCertificateFoundCount field to given value.

### HasStatCertificateFoundCount

`func (o *CSSCMSDataModelModelsSSLScanJobPart) HasStatCertificateFoundCount() bool`

HasStatCertificateFoundCount returns a boolean if a field has been set.

### SetStatCertificateFoundCountNil

`func (o *CSSCMSDataModelModelsSSLScanJobPart) SetStatCertificateFoundCountNil(b bool)`

 SetStatCertificateFoundCountNil sets the value for StatCertificateFoundCount to be an explicit nil

### UnsetStatCertificateFoundCount
`func (o *CSSCMSDataModelModelsSSLScanJobPart) UnsetStatCertificateFoundCount()`

UnsetStatCertificateFoundCount ensures that no value is present for StatCertificateFoundCount, not even an explicit nil
### GetStatNoCertificateCount

`func (o *CSSCMSDataModelModelsSSLScanJobPart) GetStatNoCertificateCount() int32`

GetStatNoCertificateCount returns the StatNoCertificateCount field if non-nil, zero value otherwise.

### GetStatNoCertificateCountOk

`func (o *CSSCMSDataModelModelsSSLScanJobPart) GetStatNoCertificateCountOk() (*int32, bool)`

GetStatNoCertificateCountOk returns a tuple with the StatNoCertificateCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatNoCertificateCount

`func (o *CSSCMSDataModelModelsSSLScanJobPart) SetStatNoCertificateCount(v int32)`

SetStatNoCertificateCount sets StatNoCertificateCount field to given value.

### HasStatNoCertificateCount

`func (o *CSSCMSDataModelModelsSSLScanJobPart) HasStatNoCertificateCount() bool`

HasStatNoCertificateCount returns a boolean if a field has been set.

### SetStatNoCertificateCountNil

`func (o *CSSCMSDataModelModelsSSLScanJobPart) SetStatNoCertificateCountNil(b bool)`

 SetStatNoCertificateCountNil sets the value for StatNoCertificateCount to be an explicit nil

### UnsetStatNoCertificateCount
`func (o *CSSCMSDataModelModelsSSLScanJobPart) UnsetStatNoCertificateCount()`

UnsetStatNoCertificateCount ensures that no value is present for StatNoCertificateCount, not even an explicit nil
### GetScanJobPartDefinitions

`func (o *CSSCMSDataModelModelsSSLScanJobPart) GetScanJobPartDefinitions() []CSSCMSDataModelModelsSSLScanJobPartDefinition`

GetScanJobPartDefinitions returns the ScanJobPartDefinitions field if non-nil, zero value otherwise.

### GetScanJobPartDefinitionsOk

`func (o *CSSCMSDataModelModelsSSLScanJobPart) GetScanJobPartDefinitionsOk() (*[]CSSCMSDataModelModelsSSLScanJobPartDefinition, bool)`

GetScanJobPartDefinitionsOk returns a tuple with the ScanJobPartDefinitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanJobPartDefinitions

`func (o *CSSCMSDataModelModelsSSLScanJobPart) SetScanJobPartDefinitions(v []CSSCMSDataModelModelsSSLScanJobPartDefinition)`

SetScanJobPartDefinitions sets ScanJobPartDefinitions field to given value.

### HasScanJobPartDefinitions

`func (o *CSSCMSDataModelModelsSSLScanJobPart) HasScanJobPartDefinitions() bool`

HasScanJobPartDefinitions returns a boolean if a field has been set.

### SetScanJobPartDefinitionsNil

`func (o *CSSCMSDataModelModelsSSLScanJobPart) SetScanJobPartDefinitionsNil(b bool)`

 SetScanJobPartDefinitionsNil sets the value for ScanJobPartDefinitions to be an explicit nil

### UnsetScanJobPartDefinitions
`func (o *CSSCMSDataModelModelsSSLScanJobPart) UnsetScanJobPartDefinitions()`

UnsetScanJobPartDefinitions ensures that no value is present for ScanJobPartDefinitions, not even an explicit nil
### GetStartTime

`func (o *CSSCMSDataModelModelsSSLScanJobPart) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *CSSCMSDataModelModelsSSLScanJobPart) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *CSSCMSDataModelModelsSSLScanJobPart) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *CSSCMSDataModelModelsSSLScanJobPart) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTimeNil

`func (o *CSSCMSDataModelModelsSSLScanJobPart) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *CSSCMSDataModelModelsSSLScanJobPart) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
### GetEndTime

`func (o *CSSCMSDataModelModelsSSLScanJobPart) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *CSSCMSDataModelModelsSSLScanJobPart) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *CSSCMSDataModelModelsSSLScanJobPart) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *CSSCMSDataModelModelsSSLScanJobPart) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### SetEndTimeNil

`func (o *CSSCMSDataModelModelsSSLScanJobPart) SetEndTimeNil(b bool)`

 SetEndTimeNil sets the value for EndTime to be an explicit nil

### UnsetEndTime
`func (o *CSSCMSDataModelModelsSSLScanJobPart) UnsetEndTime()`

UnsetEndTime ensures that no value is present for EndTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


