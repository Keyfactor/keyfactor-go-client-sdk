# ModelsSSLEndpointHistoryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HistoryId** | Pointer to **string** |  | [optional] 
**EndpointId** | Pointer to **string** |  | [optional] 
**AuditId** | Pointer to **int64** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Status** | Pointer to **int64** |  | [optional] 
**JobType** | Pointer to **int64** |  | [optional] 
**ProbeType** | Pointer to **int64** |  | [optional] 
**ReverseDNS** | Pointer to **string** |  | [optional] 
**HistoryCertificates** | Pointer to [**[]ModelsSSLEndpointHistoryResponseCertificateModel**](ModelsSSLEndpointHistoryResponseCertificateModel.md) |  | [optional] 

## Methods

### NewModelsSSLEndpointHistoryResponse

`func NewModelsSSLEndpointHistoryResponse() *ModelsSSLEndpointHistoryResponse`

NewModelsSSLEndpointHistoryResponse instantiates a new ModelsSSLEndpointHistoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsSSLEndpointHistoryResponseWithDefaults

`func NewModelsSSLEndpointHistoryResponseWithDefaults() *ModelsSSLEndpointHistoryResponse`

NewModelsSSLEndpointHistoryResponseWithDefaults instantiates a new ModelsSSLEndpointHistoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHistoryId

`func (o *ModelsSSLEndpointHistoryResponse) GetHistoryId() string`

GetHistoryId returns the HistoryId field if non-nil, zero value otherwise.

### GetHistoryIdOk

`func (o *ModelsSSLEndpointHistoryResponse) GetHistoryIdOk() (*string, bool)`

GetHistoryIdOk returns a tuple with the HistoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryId

`func (o *ModelsSSLEndpointHistoryResponse) SetHistoryId(v string)`

SetHistoryId sets HistoryId field to given value.

### HasHistoryId

`func (o *ModelsSSLEndpointHistoryResponse) HasHistoryId() bool`

HasHistoryId returns a boolean if a field has been set.

### GetEndpointId

`func (o *ModelsSSLEndpointHistoryResponse) GetEndpointId() string`

GetEndpointId returns the EndpointId field if non-nil, zero value otherwise.

### GetEndpointIdOk

`func (o *ModelsSSLEndpointHistoryResponse) GetEndpointIdOk() (*string, bool)`

GetEndpointIdOk returns a tuple with the EndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointId

`func (o *ModelsSSLEndpointHistoryResponse) SetEndpointId(v string)`

SetEndpointId sets EndpointId field to given value.

### HasEndpointId

`func (o *ModelsSSLEndpointHistoryResponse) HasEndpointId() bool`

HasEndpointId returns a boolean if a field has been set.

### GetAuditId

`func (o *ModelsSSLEndpointHistoryResponse) GetAuditId() int64`

GetAuditId returns the AuditId field if non-nil, zero value otherwise.

### GetAuditIdOk

`func (o *ModelsSSLEndpointHistoryResponse) GetAuditIdOk() (*int64, bool)`

GetAuditIdOk returns a tuple with the AuditId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditId

`func (o *ModelsSSLEndpointHistoryResponse) SetAuditId(v int64)`

SetAuditId sets AuditId field to given value.

### HasAuditId

`func (o *ModelsSSLEndpointHistoryResponse) HasAuditId() bool`

HasAuditId returns a boolean if a field has been set.

### GetTimestamp

`func (o *ModelsSSLEndpointHistoryResponse) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ModelsSSLEndpointHistoryResponse) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ModelsSSLEndpointHistoryResponse) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ModelsSSLEndpointHistoryResponse) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetStatus

`func (o *ModelsSSLEndpointHistoryResponse) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModelsSSLEndpointHistoryResponse) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModelsSSLEndpointHistoryResponse) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ModelsSSLEndpointHistoryResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetJobType

`func (o *ModelsSSLEndpointHistoryResponse) GetJobType() int64`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *ModelsSSLEndpointHistoryResponse) GetJobTypeOk() (*int64, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *ModelsSSLEndpointHistoryResponse) SetJobType(v int64)`

SetJobType sets JobType field to given value.

### HasJobType

`func (o *ModelsSSLEndpointHistoryResponse) HasJobType() bool`

HasJobType returns a boolean if a field has been set.

### GetProbeType

`func (o *ModelsSSLEndpointHistoryResponse) GetProbeType() int64`

GetProbeType returns the ProbeType field if non-nil, zero value otherwise.

### GetProbeTypeOk

`func (o *ModelsSSLEndpointHistoryResponse) GetProbeTypeOk() (*int64, bool)`

GetProbeTypeOk returns a tuple with the ProbeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbeType

`func (o *ModelsSSLEndpointHistoryResponse) SetProbeType(v int64)`

SetProbeType sets ProbeType field to given value.

### HasProbeType

`func (o *ModelsSSLEndpointHistoryResponse) HasProbeType() bool`

HasProbeType returns a boolean if a field has been set.

### GetReverseDNS

`func (o *ModelsSSLEndpointHistoryResponse) GetReverseDNS() string`

GetReverseDNS returns the ReverseDNS field if non-nil, zero value otherwise.

### GetReverseDNSOk

`func (o *ModelsSSLEndpointHistoryResponse) GetReverseDNSOk() (*string, bool)`

GetReverseDNSOk returns a tuple with the ReverseDNS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReverseDNS

`func (o *ModelsSSLEndpointHistoryResponse) SetReverseDNS(v string)`

SetReverseDNS sets ReverseDNS field to given value.

### HasReverseDNS

`func (o *ModelsSSLEndpointHistoryResponse) HasReverseDNS() bool`

HasReverseDNS returns a boolean if a field has been set.

### GetHistoryCertificates

`func (o *ModelsSSLEndpointHistoryResponse) GetHistoryCertificates() []ModelsSSLEndpointHistoryResponseCertificateModel`

GetHistoryCertificates returns the HistoryCertificates field if non-nil, zero value otherwise.

### GetHistoryCertificatesOk

`func (o *ModelsSSLEndpointHistoryResponse) GetHistoryCertificatesOk() (*[]ModelsSSLEndpointHistoryResponseCertificateModel, bool)`

GetHistoryCertificatesOk returns a tuple with the HistoryCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryCertificates

`func (o *ModelsSSLEndpointHistoryResponse) SetHistoryCertificates(v []ModelsSSLEndpointHistoryResponseCertificateModel)`

SetHistoryCertificates sets HistoryCertificates field to given value.

### HasHistoryCertificates

`func (o *ModelsSSLEndpointHistoryResponse) HasHistoryCertificates() bool`

HasHistoryCertificates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


