# CSSCMSDataModelModelsSSLEndpointHistoryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HistoryId** | Pointer to **string** |  | [optional] 
**EndpointId** | Pointer to **string** |  | [optional] 
**AuditId** | Pointer to **int64** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Status** | Pointer to [**KeyfactorOrchestratorsCommonEnumsSslEndpointStatus**](KeyfactorOrchestratorsCommonEnumsSslEndpointStatus.md) |  | [optional] 
**JobType** | Pointer to [**KeyfactorOrchestratorsCommonEnumsSslJobType**](KeyfactorOrchestratorsCommonEnumsSslJobType.md) |  | [optional] 
**ProbeType** | Pointer to [**KeyfactorOrchestratorsCommonEnumsSslProbeType**](KeyfactorOrchestratorsCommonEnumsSslProbeType.md) |  | [optional] 
**ReverseDNS** | Pointer to **NullableString** |  | [optional] 
**HistoryCertificates** | Pointer to [**[]CSSCMSDataModelModelsSSLEndpointHistoryResponseCertificateModel**](CSSCMSDataModelModelsSSLEndpointHistoryResponseCertificateModel.md) |  | [optional] 

## Methods

### NewCSSCMSDataModelModelsSSLEndpointHistoryResponse

`func NewCSSCMSDataModelModelsSSLEndpointHistoryResponse() *CSSCMSDataModelModelsSSLEndpointHistoryResponse`

NewCSSCMSDataModelModelsSSLEndpointHistoryResponse instantiates a new CSSCMSDataModelModelsSSLEndpointHistoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSSCMSDataModelModelsSSLEndpointHistoryResponseWithDefaults

`func NewCSSCMSDataModelModelsSSLEndpointHistoryResponseWithDefaults() *CSSCMSDataModelModelsSSLEndpointHistoryResponse`

NewCSSCMSDataModelModelsSSLEndpointHistoryResponseWithDefaults instantiates a new CSSCMSDataModelModelsSSLEndpointHistoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHistoryId

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponse) GetHistoryId() string`

GetHistoryId returns the HistoryId field if non-nil, zero value otherwise.

### GetHistoryIdOk

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponse) GetHistoryIdOk() (*string, bool)`

GetHistoryIdOk returns a tuple with the HistoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryId

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponse) SetHistoryId(v string)`

SetHistoryId sets HistoryId field to given value.

### HasHistoryId

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponse) HasHistoryId() bool`

HasHistoryId returns a boolean if a field has been set.

### GetEndpointId

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponse) GetEndpointId() string`

GetEndpointId returns the EndpointId field if non-nil, zero value otherwise.

### GetEndpointIdOk

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponse) GetEndpointIdOk() (*string, bool)`

GetEndpointIdOk returns a tuple with the EndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointId

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponse) SetEndpointId(v string)`

SetEndpointId sets EndpointId field to given value.

### HasEndpointId

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponse) HasEndpointId() bool`

HasEndpointId returns a boolean if a field has been set.

### GetAuditId

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponse) GetAuditId() int64`

GetAuditId returns the AuditId field if non-nil, zero value otherwise.

### GetAuditIdOk

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponse) GetAuditIdOk() (*int64, bool)`

GetAuditIdOk returns a tuple with the AuditId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditId

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponse) SetAuditId(v int64)`

SetAuditId sets AuditId field to given value.

### HasAuditId

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponse) HasAuditId() bool`

HasAuditId returns a boolean if a field has been set.

### GetTimestamp

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponse) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponse) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponse) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponse) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetStatus

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponse) GetStatus() KeyfactorOrchestratorsCommonEnumsSslEndpointStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponse) GetStatusOk() (*KeyfactorOrchestratorsCommonEnumsSslEndpointStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponse) SetStatus(v KeyfactorOrchestratorsCommonEnumsSslEndpointStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetJobType

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponse) GetJobType() KeyfactorOrchestratorsCommonEnumsSslJobType`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponse) GetJobTypeOk() (*KeyfactorOrchestratorsCommonEnumsSslJobType, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponse) SetJobType(v KeyfactorOrchestratorsCommonEnumsSslJobType)`

SetJobType sets JobType field to given value.

### HasJobType

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponse) HasJobType() bool`

HasJobType returns a boolean if a field has been set.

### GetProbeType

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponse) GetProbeType() KeyfactorOrchestratorsCommonEnumsSslProbeType`

GetProbeType returns the ProbeType field if non-nil, zero value otherwise.

### GetProbeTypeOk

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponse) GetProbeTypeOk() (*KeyfactorOrchestratorsCommonEnumsSslProbeType, bool)`

GetProbeTypeOk returns a tuple with the ProbeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbeType

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponse) SetProbeType(v KeyfactorOrchestratorsCommonEnumsSslProbeType)`

SetProbeType sets ProbeType field to given value.

### HasProbeType

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponse) HasProbeType() bool`

HasProbeType returns a boolean if a field has been set.

### GetReverseDNS

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponse) GetReverseDNS() string`

GetReverseDNS returns the ReverseDNS field if non-nil, zero value otherwise.

### GetReverseDNSOk

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponse) GetReverseDNSOk() (*string, bool)`

GetReverseDNSOk returns a tuple with the ReverseDNS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReverseDNS

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponse) SetReverseDNS(v string)`

SetReverseDNS sets ReverseDNS field to given value.

### HasReverseDNS

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponse) HasReverseDNS() bool`

HasReverseDNS returns a boolean if a field has been set.

### SetReverseDNSNil

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponse) SetReverseDNSNil(b bool)`

 SetReverseDNSNil sets the value for ReverseDNS to be an explicit nil

### UnsetReverseDNS
`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponse) UnsetReverseDNS()`

UnsetReverseDNS ensures that no value is present for ReverseDNS, not even an explicit nil
### GetHistoryCertificates

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponse) GetHistoryCertificates() []CSSCMSDataModelModelsSSLEndpointHistoryResponseCertificateModel`

GetHistoryCertificates returns the HistoryCertificates field if non-nil, zero value otherwise.

### GetHistoryCertificatesOk

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponse) GetHistoryCertificatesOk() (*[]CSSCMSDataModelModelsSSLEndpointHistoryResponseCertificateModel, bool)`

GetHistoryCertificatesOk returns a tuple with the HistoryCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryCertificates

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponse) SetHistoryCertificates(v []CSSCMSDataModelModelsSSLEndpointHistoryResponseCertificateModel)`

SetHistoryCertificates sets HistoryCertificates field to given value.

### HasHistoryCertificates

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponse) HasHistoryCertificates() bool`

HasHistoryCertificates returns a boolean if a field has been set.

### SetHistoryCertificatesNil

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponse) SetHistoryCertificatesNil(b bool)`

 SetHistoryCertificatesNil sets the value for HistoryCertificates to be an explicit nil

### UnsetHistoryCertificates
`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponse) UnsetHistoryCertificates()`

UnsetHistoryCertificates ensures that no value is present for HistoryCertificates, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


