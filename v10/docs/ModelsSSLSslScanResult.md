# ModelsSSLSslScanResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndpointId** | Pointer to **string** |  | [optional] 
**ReverseDNS** | Pointer to **string** |  | [optional] 
**SNIName** | Pointer to **string** |  | [optional] 
**IpAddress** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int64** |  | [optional] 
**CertificateFound** | Pointer to **bool** |  | [optional] 
**AgentPoolName** | Pointer to **string** |  | [optional] 
**NetworkName** | Pointer to **string** |  | [optional] 
**MonitorStatus** | Pointer to **bool** |  | [optional] 
**CertificateCN** | Pointer to **string** |  | [optional] 
**Reviewed** | Pointer to **bool** |  | [optional] 

## Methods

### NewModelsSSLSslScanResult

`func NewModelsSSLSslScanResult() *ModelsSSLSslScanResult`

NewModelsSSLSslScanResult instantiates a new ModelsSSLSslScanResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsSSLSslScanResultWithDefaults

`func NewModelsSSLSslScanResultWithDefaults() *ModelsSSLSslScanResult`

NewModelsSSLSslScanResultWithDefaults instantiates a new ModelsSSLSslScanResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpointId

`func (o *ModelsSSLSslScanResult) GetEndpointId() string`

GetEndpointId returns the EndpointId field if non-nil, zero value otherwise.

### GetEndpointIdOk

`func (o *ModelsSSLSslScanResult) GetEndpointIdOk() (*string, bool)`

GetEndpointIdOk returns a tuple with the EndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointId

`func (o *ModelsSSLSslScanResult) SetEndpointId(v string)`

SetEndpointId sets EndpointId field to given value.

### HasEndpointId

`func (o *ModelsSSLSslScanResult) HasEndpointId() bool`

HasEndpointId returns a boolean if a field has been set.

### GetReverseDNS

`func (o *ModelsSSLSslScanResult) GetReverseDNS() string`

GetReverseDNS returns the ReverseDNS field if non-nil, zero value otherwise.

### GetReverseDNSOk

`func (o *ModelsSSLSslScanResult) GetReverseDNSOk() (*string, bool)`

GetReverseDNSOk returns a tuple with the ReverseDNS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReverseDNS

`func (o *ModelsSSLSslScanResult) SetReverseDNS(v string)`

SetReverseDNS sets ReverseDNS field to given value.

### HasReverseDNS

`func (o *ModelsSSLSslScanResult) HasReverseDNS() bool`

HasReverseDNS returns a boolean if a field has been set.

### GetSNIName

`func (o *ModelsSSLSslScanResult) GetSNIName() string`

GetSNIName returns the SNIName field if non-nil, zero value otherwise.

### GetSNINameOk

`func (o *ModelsSSLSslScanResult) GetSNINameOk() (*string, bool)`

GetSNINameOk returns a tuple with the SNIName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSNIName

`func (o *ModelsSSLSslScanResult) SetSNIName(v string)`

SetSNIName sets SNIName field to given value.

### HasSNIName

`func (o *ModelsSSLSslScanResult) HasSNIName() bool`

HasSNIName returns a boolean if a field has been set.

### GetIpAddress

`func (o *ModelsSSLSslScanResult) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *ModelsSSLSslScanResult) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *ModelsSSLSslScanResult) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *ModelsSSLSslScanResult) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetPort

`func (o *ModelsSSLSslScanResult) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ModelsSSLSslScanResult) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ModelsSSLSslScanResult) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *ModelsSSLSslScanResult) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetCertificateFound

`func (o *ModelsSSLSslScanResult) GetCertificateFound() bool`

GetCertificateFound returns the CertificateFound field if non-nil, zero value otherwise.

### GetCertificateFoundOk

`func (o *ModelsSSLSslScanResult) GetCertificateFoundOk() (*bool, bool)`

GetCertificateFoundOk returns a tuple with the CertificateFound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateFound

`func (o *ModelsSSLSslScanResult) SetCertificateFound(v bool)`

SetCertificateFound sets CertificateFound field to given value.

### HasCertificateFound

`func (o *ModelsSSLSslScanResult) HasCertificateFound() bool`

HasCertificateFound returns a boolean if a field has been set.

### GetAgentPoolName

`func (o *ModelsSSLSslScanResult) GetAgentPoolName() string`

GetAgentPoolName returns the AgentPoolName field if non-nil, zero value otherwise.

### GetAgentPoolNameOk

`func (o *ModelsSSLSslScanResult) GetAgentPoolNameOk() (*string, bool)`

GetAgentPoolNameOk returns a tuple with the AgentPoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentPoolName

`func (o *ModelsSSLSslScanResult) SetAgentPoolName(v string)`

SetAgentPoolName sets AgentPoolName field to given value.

### HasAgentPoolName

`func (o *ModelsSSLSslScanResult) HasAgentPoolName() bool`

HasAgentPoolName returns a boolean if a field has been set.

### GetNetworkName

`func (o *ModelsSSLSslScanResult) GetNetworkName() string`

GetNetworkName returns the NetworkName field if non-nil, zero value otherwise.

### GetNetworkNameOk

`func (o *ModelsSSLSslScanResult) GetNetworkNameOk() (*string, bool)`

GetNetworkNameOk returns a tuple with the NetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkName

`func (o *ModelsSSLSslScanResult) SetNetworkName(v string)`

SetNetworkName sets NetworkName field to given value.

### HasNetworkName

`func (o *ModelsSSLSslScanResult) HasNetworkName() bool`

HasNetworkName returns a boolean if a field has been set.

### GetMonitorStatus

`func (o *ModelsSSLSslScanResult) GetMonitorStatus() bool`

GetMonitorStatus returns the MonitorStatus field if non-nil, zero value otherwise.

### GetMonitorStatusOk

`func (o *ModelsSSLSslScanResult) GetMonitorStatusOk() (*bool, bool)`

GetMonitorStatusOk returns a tuple with the MonitorStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorStatus

`func (o *ModelsSSLSslScanResult) SetMonitorStatus(v bool)`

SetMonitorStatus sets MonitorStatus field to given value.

### HasMonitorStatus

`func (o *ModelsSSLSslScanResult) HasMonitorStatus() bool`

HasMonitorStatus returns a boolean if a field has been set.

### GetCertificateCN

`func (o *ModelsSSLSslScanResult) GetCertificateCN() string`

GetCertificateCN returns the CertificateCN field if non-nil, zero value otherwise.

### GetCertificateCNOk

`func (o *ModelsSSLSslScanResult) GetCertificateCNOk() (*string, bool)`

GetCertificateCNOk returns a tuple with the CertificateCN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateCN

`func (o *ModelsSSLSslScanResult) SetCertificateCN(v string)`

SetCertificateCN sets CertificateCN field to given value.

### HasCertificateCN

`func (o *ModelsSSLSslScanResult) HasCertificateCN() bool`

HasCertificateCN returns a boolean if a field has been set.

### GetReviewed

`func (o *ModelsSSLSslScanResult) GetReviewed() bool`

GetReviewed returns the Reviewed field if non-nil, zero value otherwise.

### GetReviewedOk

`func (o *ModelsSSLSslScanResult) GetReviewedOk() (*bool, bool)`

GetReviewedOk returns a tuple with the Reviewed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewed

`func (o *ModelsSSLSslScanResult) SetReviewed(v bool)`

SetReviewed sets Reviewed field to given value.

### HasReviewed

`func (o *ModelsSSLSslScanResult) HasReviewed() bool`

HasReviewed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


