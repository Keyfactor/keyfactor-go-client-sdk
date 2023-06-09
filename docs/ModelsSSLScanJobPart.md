# ModelsSSLScanJobPart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScanJobPartId** | Pointer to **string** |  | [optional] 
**LogicalScanJobId** | Pointer to **string** |  | [optional] 
**AgentJobId** | Pointer to **string** |  | [optional] 
**EstimatedEndpointCount** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 
**StatTotalEndpointCount** | Pointer to **int32** |  | [optional] 
**StatTimedOutConnectingCount** | Pointer to **int32** |  | [optional] 
**StatConnectionRefusedCount** | Pointer to **int32** |  | [optional] 
**StatTimedOutDownloadingCount** | Pointer to **int32** |  | [optional] 
**StatExceptionDownloadingCount** | Pointer to **int32** |  | [optional] 
**StatNotSslCount** | Pointer to **int32** |  | [optional] 
**StatBadSslHandshakeCount** | Pointer to **int32** |  | [optional] 
**StatCertificateFoundCount** | Pointer to **int32** |  | [optional] 
**StatNoCertificateCount** | Pointer to **int32** |  | [optional] 
**ScanJobPartDefinitions** | Pointer to [**[]ModelsSSLScanJobPartDefinition**](ModelsSSLScanJobPartDefinition.md) |  | [optional] 
**StartTime** | Pointer to **time.Time** |  | [optional] 
**EndTime** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewModelsSSLScanJobPart

`func NewModelsSSLScanJobPart() *ModelsSSLScanJobPart`

NewModelsSSLScanJobPart instantiates a new ModelsSSLScanJobPart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsSSLScanJobPartWithDefaults

`func NewModelsSSLScanJobPartWithDefaults() *ModelsSSLScanJobPart`

NewModelsSSLScanJobPartWithDefaults instantiates a new ModelsSSLScanJobPart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScanJobPartId

`func (o *ModelsSSLScanJobPart) GetScanJobPartId() string`

GetScanJobPartId returns the ScanJobPartId field if non-nil, zero value otherwise.

### GetScanJobPartIdOk

`func (o *ModelsSSLScanJobPart) GetScanJobPartIdOk() (*string, bool)`

GetScanJobPartIdOk returns a tuple with the ScanJobPartId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanJobPartId

`func (o *ModelsSSLScanJobPart) SetScanJobPartId(v string)`

SetScanJobPartId sets ScanJobPartId field to given value.

### HasScanJobPartId

`func (o *ModelsSSLScanJobPart) HasScanJobPartId() bool`

HasScanJobPartId returns a boolean if a field has been set.

### GetLogicalScanJobId

`func (o *ModelsSSLScanJobPart) GetLogicalScanJobId() string`

GetLogicalScanJobId returns the LogicalScanJobId field if non-nil, zero value otherwise.

### GetLogicalScanJobIdOk

`func (o *ModelsSSLScanJobPart) GetLogicalScanJobIdOk() (*string, bool)`

GetLogicalScanJobIdOk returns a tuple with the LogicalScanJobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalScanJobId

`func (o *ModelsSSLScanJobPart) SetLogicalScanJobId(v string)`

SetLogicalScanJobId sets LogicalScanJobId field to given value.

### HasLogicalScanJobId

`func (o *ModelsSSLScanJobPart) HasLogicalScanJobId() bool`

HasLogicalScanJobId returns a boolean if a field has been set.

### GetAgentJobId

`func (o *ModelsSSLScanJobPart) GetAgentJobId() string`

GetAgentJobId returns the AgentJobId field if non-nil, zero value otherwise.

### GetAgentJobIdOk

`func (o *ModelsSSLScanJobPart) GetAgentJobIdOk() (*string, bool)`

GetAgentJobIdOk returns a tuple with the AgentJobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentJobId

`func (o *ModelsSSLScanJobPart) SetAgentJobId(v string)`

SetAgentJobId sets AgentJobId field to given value.

### HasAgentJobId

`func (o *ModelsSSLScanJobPart) HasAgentJobId() bool`

HasAgentJobId returns a boolean if a field has been set.

### GetEstimatedEndpointCount

`func (o *ModelsSSLScanJobPart) GetEstimatedEndpointCount() int32`

GetEstimatedEndpointCount returns the EstimatedEndpointCount field if non-nil, zero value otherwise.

### GetEstimatedEndpointCountOk

`func (o *ModelsSSLScanJobPart) GetEstimatedEndpointCountOk() (*int32, bool)`

GetEstimatedEndpointCountOk returns a tuple with the EstimatedEndpointCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedEndpointCount

`func (o *ModelsSSLScanJobPart) SetEstimatedEndpointCount(v int32)`

SetEstimatedEndpointCount sets EstimatedEndpointCount field to given value.

### HasEstimatedEndpointCount

`func (o *ModelsSSLScanJobPart) HasEstimatedEndpointCount() bool`

HasEstimatedEndpointCount returns a boolean if a field has been set.

### GetStatus

`func (o *ModelsSSLScanJobPart) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModelsSSLScanJobPart) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModelsSSLScanJobPart) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ModelsSSLScanJobPart) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatTotalEndpointCount

`func (o *ModelsSSLScanJobPart) GetStatTotalEndpointCount() int32`

GetStatTotalEndpointCount returns the StatTotalEndpointCount field if non-nil, zero value otherwise.

### GetStatTotalEndpointCountOk

`func (o *ModelsSSLScanJobPart) GetStatTotalEndpointCountOk() (*int32, bool)`

GetStatTotalEndpointCountOk returns a tuple with the StatTotalEndpointCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatTotalEndpointCount

`func (o *ModelsSSLScanJobPart) SetStatTotalEndpointCount(v int32)`

SetStatTotalEndpointCount sets StatTotalEndpointCount field to given value.

### HasStatTotalEndpointCount

`func (o *ModelsSSLScanJobPart) HasStatTotalEndpointCount() bool`

HasStatTotalEndpointCount returns a boolean if a field has been set.

### GetStatTimedOutConnectingCount

`func (o *ModelsSSLScanJobPart) GetStatTimedOutConnectingCount() int32`

GetStatTimedOutConnectingCount returns the StatTimedOutConnectingCount field if non-nil, zero value otherwise.

### GetStatTimedOutConnectingCountOk

`func (o *ModelsSSLScanJobPart) GetStatTimedOutConnectingCountOk() (*int32, bool)`

GetStatTimedOutConnectingCountOk returns a tuple with the StatTimedOutConnectingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatTimedOutConnectingCount

`func (o *ModelsSSLScanJobPart) SetStatTimedOutConnectingCount(v int32)`

SetStatTimedOutConnectingCount sets StatTimedOutConnectingCount field to given value.

### HasStatTimedOutConnectingCount

`func (o *ModelsSSLScanJobPart) HasStatTimedOutConnectingCount() bool`

HasStatTimedOutConnectingCount returns a boolean if a field has been set.

### GetStatConnectionRefusedCount

`func (o *ModelsSSLScanJobPart) GetStatConnectionRefusedCount() int32`

GetStatConnectionRefusedCount returns the StatConnectionRefusedCount field if non-nil, zero value otherwise.

### GetStatConnectionRefusedCountOk

`func (o *ModelsSSLScanJobPart) GetStatConnectionRefusedCountOk() (*int32, bool)`

GetStatConnectionRefusedCountOk returns a tuple with the StatConnectionRefusedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatConnectionRefusedCount

`func (o *ModelsSSLScanJobPart) SetStatConnectionRefusedCount(v int32)`

SetStatConnectionRefusedCount sets StatConnectionRefusedCount field to given value.

### HasStatConnectionRefusedCount

`func (o *ModelsSSLScanJobPart) HasStatConnectionRefusedCount() bool`

HasStatConnectionRefusedCount returns a boolean if a field has been set.

### GetStatTimedOutDownloadingCount

`func (o *ModelsSSLScanJobPart) GetStatTimedOutDownloadingCount() int32`

GetStatTimedOutDownloadingCount returns the StatTimedOutDownloadingCount field if non-nil, zero value otherwise.

### GetStatTimedOutDownloadingCountOk

`func (o *ModelsSSLScanJobPart) GetStatTimedOutDownloadingCountOk() (*int32, bool)`

GetStatTimedOutDownloadingCountOk returns a tuple with the StatTimedOutDownloadingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatTimedOutDownloadingCount

`func (o *ModelsSSLScanJobPart) SetStatTimedOutDownloadingCount(v int32)`

SetStatTimedOutDownloadingCount sets StatTimedOutDownloadingCount field to given value.

### HasStatTimedOutDownloadingCount

`func (o *ModelsSSLScanJobPart) HasStatTimedOutDownloadingCount() bool`

HasStatTimedOutDownloadingCount returns a boolean if a field has been set.

### GetStatExceptionDownloadingCount

`func (o *ModelsSSLScanJobPart) GetStatExceptionDownloadingCount() int32`

GetStatExceptionDownloadingCount returns the StatExceptionDownloadingCount field if non-nil, zero value otherwise.

### GetStatExceptionDownloadingCountOk

`func (o *ModelsSSLScanJobPart) GetStatExceptionDownloadingCountOk() (*int32, bool)`

GetStatExceptionDownloadingCountOk returns a tuple with the StatExceptionDownloadingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatExceptionDownloadingCount

`func (o *ModelsSSLScanJobPart) SetStatExceptionDownloadingCount(v int32)`

SetStatExceptionDownloadingCount sets StatExceptionDownloadingCount field to given value.

### HasStatExceptionDownloadingCount

`func (o *ModelsSSLScanJobPart) HasStatExceptionDownloadingCount() bool`

HasStatExceptionDownloadingCount returns a boolean if a field has been set.

### GetStatNotSslCount

`func (o *ModelsSSLScanJobPart) GetStatNotSslCount() int32`

GetStatNotSslCount returns the StatNotSslCount field if non-nil, zero value otherwise.

### GetStatNotSslCountOk

`func (o *ModelsSSLScanJobPart) GetStatNotSslCountOk() (*int32, bool)`

GetStatNotSslCountOk returns a tuple with the StatNotSslCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatNotSslCount

`func (o *ModelsSSLScanJobPart) SetStatNotSslCount(v int32)`

SetStatNotSslCount sets StatNotSslCount field to given value.

### HasStatNotSslCount

`func (o *ModelsSSLScanJobPart) HasStatNotSslCount() bool`

HasStatNotSslCount returns a boolean if a field has been set.

### GetStatBadSslHandshakeCount

`func (o *ModelsSSLScanJobPart) GetStatBadSslHandshakeCount() int32`

GetStatBadSslHandshakeCount returns the StatBadSslHandshakeCount field if non-nil, zero value otherwise.

### GetStatBadSslHandshakeCountOk

`func (o *ModelsSSLScanJobPart) GetStatBadSslHandshakeCountOk() (*int32, bool)`

GetStatBadSslHandshakeCountOk returns a tuple with the StatBadSslHandshakeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatBadSslHandshakeCount

`func (o *ModelsSSLScanJobPart) SetStatBadSslHandshakeCount(v int32)`

SetStatBadSslHandshakeCount sets StatBadSslHandshakeCount field to given value.

### HasStatBadSslHandshakeCount

`func (o *ModelsSSLScanJobPart) HasStatBadSslHandshakeCount() bool`

HasStatBadSslHandshakeCount returns a boolean if a field has been set.

### GetStatCertificateFoundCount

`func (o *ModelsSSLScanJobPart) GetStatCertificateFoundCount() int32`

GetStatCertificateFoundCount returns the StatCertificateFoundCount field if non-nil, zero value otherwise.

### GetStatCertificateFoundCountOk

`func (o *ModelsSSLScanJobPart) GetStatCertificateFoundCountOk() (*int32, bool)`

GetStatCertificateFoundCountOk returns a tuple with the StatCertificateFoundCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatCertificateFoundCount

`func (o *ModelsSSLScanJobPart) SetStatCertificateFoundCount(v int32)`

SetStatCertificateFoundCount sets StatCertificateFoundCount field to given value.

### HasStatCertificateFoundCount

`func (o *ModelsSSLScanJobPart) HasStatCertificateFoundCount() bool`

HasStatCertificateFoundCount returns a boolean if a field has been set.

### GetStatNoCertificateCount

`func (o *ModelsSSLScanJobPart) GetStatNoCertificateCount() int32`

GetStatNoCertificateCount returns the StatNoCertificateCount field if non-nil, zero value otherwise.

### GetStatNoCertificateCountOk

`func (o *ModelsSSLScanJobPart) GetStatNoCertificateCountOk() (*int32, bool)`

GetStatNoCertificateCountOk returns a tuple with the StatNoCertificateCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatNoCertificateCount

`func (o *ModelsSSLScanJobPart) SetStatNoCertificateCount(v int32)`

SetStatNoCertificateCount sets StatNoCertificateCount field to given value.

### HasStatNoCertificateCount

`func (o *ModelsSSLScanJobPart) HasStatNoCertificateCount() bool`

HasStatNoCertificateCount returns a boolean if a field has been set.

### GetScanJobPartDefinitions

`func (o *ModelsSSLScanJobPart) GetScanJobPartDefinitions() []ModelsSSLScanJobPartDefinition`

GetScanJobPartDefinitions returns the ScanJobPartDefinitions field if non-nil, zero value otherwise.

### GetScanJobPartDefinitionsOk

`func (o *ModelsSSLScanJobPart) GetScanJobPartDefinitionsOk() (*[]ModelsSSLScanJobPartDefinition, bool)`

GetScanJobPartDefinitionsOk returns a tuple with the ScanJobPartDefinitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanJobPartDefinitions

`func (o *ModelsSSLScanJobPart) SetScanJobPartDefinitions(v []ModelsSSLScanJobPartDefinition)`

SetScanJobPartDefinitions sets ScanJobPartDefinitions field to given value.

### HasScanJobPartDefinitions

`func (o *ModelsSSLScanJobPart) HasScanJobPartDefinitions() bool`

HasScanJobPartDefinitions returns a boolean if a field has been set.

### GetStartTime

`func (o *ModelsSSLScanJobPart) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ModelsSSLScanJobPart) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ModelsSSLScanJobPart) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *ModelsSSLScanJobPart) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *ModelsSSLScanJobPart) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *ModelsSSLScanJobPart) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *ModelsSSLScanJobPart) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *ModelsSSLScanJobPart) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


