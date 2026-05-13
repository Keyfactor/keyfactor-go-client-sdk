# CSSCMSDataModelModelsSSLSslScanResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndpointId** | Pointer to **string** |  | [optional] 
**ReverseDNS** | Pointer to **NullableString** |  | [optional] 
**SNIName** | Pointer to **NullableString** |  | [optional] 
**IpAddress** | Pointer to **NullableString** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**CertificateFound** | Pointer to **bool** |  | [optional] 
**AgentPoolName** | Pointer to **NullableString** |  | [optional] 
**NetworkName** | Pointer to **NullableString** |  | [optional] 
**MonitorStatus** | Pointer to **bool** |  | [optional] 
**CertificateCN** | Pointer to **NullableString** |  | [optional] 
**Reviewed** | Pointer to **bool** |  | [optional] 
**ExpirationDate** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewCSSCMSDataModelModelsSSLSslScanResult

`func NewCSSCMSDataModelModelsSSLSslScanResult() *CSSCMSDataModelModelsSSLSslScanResult`

NewCSSCMSDataModelModelsSSLSslScanResult instantiates a new CSSCMSDataModelModelsSSLSslScanResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSSCMSDataModelModelsSSLSslScanResultWithDefaults

`func NewCSSCMSDataModelModelsSSLSslScanResultWithDefaults() *CSSCMSDataModelModelsSSLSslScanResult`

NewCSSCMSDataModelModelsSSLSslScanResultWithDefaults instantiates a new CSSCMSDataModelModelsSSLSslScanResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpointId

`func (o *CSSCMSDataModelModelsSSLSslScanResult) GetEndpointId() string`

GetEndpointId returns the EndpointId field if non-nil, zero value otherwise.

### GetEndpointIdOk

`func (o *CSSCMSDataModelModelsSSLSslScanResult) GetEndpointIdOk() (*string, bool)`

GetEndpointIdOk returns a tuple with the EndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointId

`func (o *CSSCMSDataModelModelsSSLSslScanResult) SetEndpointId(v string)`

SetEndpointId sets EndpointId field to given value.

### HasEndpointId

`func (o *CSSCMSDataModelModelsSSLSslScanResult) HasEndpointId() bool`

HasEndpointId returns a boolean if a field has been set.

### GetReverseDNS

`func (o *CSSCMSDataModelModelsSSLSslScanResult) GetReverseDNS() string`

GetReverseDNS returns the ReverseDNS field if non-nil, zero value otherwise.

### GetReverseDNSOk

`func (o *CSSCMSDataModelModelsSSLSslScanResult) GetReverseDNSOk() (*string, bool)`

GetReverseDNSOk returns a tuple with the ReverseDNS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReverseDNS

`func (o *CSSCMSDataModelModelsSSLSslScanResult) SetReverseDNS(v string)`

SetReverseDNS sets ReverseDNS field to given value.

### HasReverseDNS

`func (o *CSSCMSDataModelModelsSSLSslScanResult) HasReverseDNS() bool`

HasReverseDNS returns a boolean if a field has been set.

### SetReverseDNSNil

`func (o *CSSCMSDataModelModelsSSLSslScanResult) SetReverseDNSNil(b bool)`

 SetReverseDNSNil sets the value for ReverseDNS to be an explicit nil

### UnsetReverseDNS
`func (o *CSSCMSDataModelModelsSSLSslScanResult) UnsetReverseDNS()`

UnsetReverseDNS ensures that no value is present for ReverseDNS, not even an explicit nil
### GetSNIName

`func (o *CSSCMSDataModelModelsSSLSslScanResult) GetSNIName() string`

GetSNIName returns the SNIName field if non-nil, zero value otherwise.

### GetSNINameOk

`func (o *CSSCMSDataModelModelsSSLSslScanResult) GetSNINameOk() (*string, bool)`

GetSNINameOk returns a tuple with the SNIName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSNIName

`func (o *CSSCMSDataModelModelsSSLSslScanResult) SetSNIName(v string)`

SetSNIName sets SNIName field to given value.

### HasSNIName

`func (o *CSSCMSDataModelModelsSSLSslScanResult) HasSNIName() bool`

HasSNIName returns a boolean if a field has been set.

### SetSNINameNil

`func (o *CSSCMSDataModelModelsSSLSslScanResult) SetSNINameNil(b bool)`

 SetSNINameNil sets the value for SNIName to be an explicit nil

### UnsetSNIName
`func (o *CSSCMSDataModelModelsSSLSslScanResult) UnsetSNIName()`

UnsetSNIName ensures that no value is present for SNIName, not even an explicit nil
### GetIpAddress

`func (o *CSSCMSDataModelModelsSSLSslScanResult) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *CSSCMSDataModelModelsSSLSslScanResult) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *CSSCMSDataModelModelsSSLSslScanResult) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *CSSCMSDataModelModelsSSLSslScanResult) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### SetIpAddressNil

`func (o *CSSCMSDataModelModelsSSLSslScanResult) SetIpAddressNil(b bool)`

 SetIpAddressNil sets the value for IpAddress to be an explicit nil

### UnsetIpAddress
`func (o *CSSCMSDataModelModelsSSLSslScanResult) UnsetIpAddress()`

UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil
### GetPort

`func (o *CSSCMSDataModelModelsSSLSslScanResult) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *CSSCMSDataModelModelsSSLSslScanResult) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *CSSCMSDataModelModelsSSLSslScanResult) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *CSSCMSDataModelModelsSSLSslScanResult) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetCertificateFound

`func (o *CSSCMSDataModelModelsSSLSslScanResult) GetCertificateFound() bool`

GetCertificateFound returns the CertificateFound field if non-nil, zero value otherwise.

### GetCertificateFoundOk

`func (o *CSSCMSDataModelModelsSSLSslScanResult) GetCertificateFoundOk() (*bool, bool)`

GetCertificateFoundOk returns a tuple with the CertificateFound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateFound

`func (o *CSSCMSDataModelModelsSSLSslScanResult) SetCertificateFound(v bool)`

SetCertificateFound sets CertificateFound field to given value.

### HasCertificateFound

`func (o *CSSCMSDataModelModelsSSLSslScanResult) HasCertificateFound() bool`

HasCertificateFound returns a boolean if a field has been set.

### GetAgentPoolName

`func (o *CSSCMSDataModelModelsSSLSslScanResult) GetAgentPoolName() string`

GetAgentPoolName returns the AgentPoolName field if non-nil, zero value otherwise.

### GetAgentPoolNameOk

`func (o *CSSCMSDataModelModelsSSLSslScanResult) GetAgentPoolNameOk() (*string, bool)`

GetAgentPoolNameOk returns a tuple with the AgentPoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentPoolName

`func (o *CSSCMSDataModelModelsSSLSslScanResult) SetAgentPoolName(v string)`

SetAgentPoolName sets AgentPoolName field to given value.

### HasAgentPoolName

`func (o *CSSCMSDataModelModelsSSLSslScanResult) HasAgentPoolName() bool`

HasAgentPoolName returns a boolean if a field has been set.

### SetAgentPoolNameNil

`func (o *CSSCMSDataModelModelsSSLSslScanResult) SetAgentPoolNameNil(b bool)`

 SetAgentPoolNameNil sets the value for AgentPoolName to be an explicit nil

### UnsetAgentPoolName
`func (o *CSSCMSDataModelModelsSSLSslScanResult) UnsetAgentPoolName()`

UnsetAgentPoolName ensures that no value is present for AgentPoolName, not even an explicit nil
### GetNetworkName

`func (o *CSSCMSDataModelModelsSSLSslScanResult) GetNetworkName() string`

GetNetworkName returns the NetworkName field if non-nil, zero value otherwise.

### GetNetworkNameOk

`func (o *CSSCMSDataModelModelsSSLSslScanResult) GetNetworkNameOk() (*string, bool)`

GetNetworkNameOk returns a tuple with the NetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkName

`func (o *CSSCMSDataModelModelsSSLSslScanResult) SetNetworkName(v string)`

SetNetworkName sets NetworkName field to given value.

### HasNetworkName

`func (o *CSSCMSDataModelModelsSSLSslScanResult) HasNetworkName() bool`

HasNetworkName returns a boolean if a field has been set.

### SetNetworkNameNil

`func (o *CSSCMSDataModelModelsSSLSslScanResult) SetNetworkNameNil(b bool)`

 SetNetworkNameNil sets the value for NetworkName to be an explicit nil

### UnsetNetworkName
`func (o *CSSCMSDataModelModelsSSLSslScanResult) UnsetNetworkName()`

UnsetNetworkName ensures that no value is present for NetworkName, not even an explicit nil
### GetMonitorStatus

`func (o *CSSCMSDataModelModelsSSLSslScanResult) GetMonitorStatus() bool`

GetMonitorStatus returns the MonitorStatus field if non-nil, zero value otherwise.

### GetMonitorStatusOk

`func (o *CSSCMSDataModelModelsSSLSslScanResult) GetMonitorStatusOk() (*bool, bool)`

GetMonitorStatusOk returns a tuple with the MonitorStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorStatus

`func (o *CSSCMSDataModelModelsSSLSslScanResult) SetMonitorStatus(v bool)`

SetMonitorStatus sets MonitorStatus field to given value.

### HasMonitorStatus

`func (o *CSSCMSDataModelModelsSSLSslScanResult) HasMonitorStatus() bool`

HasMonitorStatus returns a boolean if a field has been set.

### GetCertificateCN

`func (o *CSSCMSDataModelModelsSSLSslScanResult) GetCertificateCN() string`

GetCertificateCN returns the CertificateCN field if non-nil, zero value otherwise.

### GetCertificateCNOk

`func (o *CSSCMSDataModelModelsSSLSslScanResult) GetCertificateCNOk() (*string, bool)`

GetCertificateCNOk returns a tuple with the CertificateCN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateCN

`func (o *CSSCMSDataModelModelsSSLSslScanResult) SetCertificateCN(v string)`

SetCertificateCN sets CertificateCN field to given value.

### HasCertificateCN

`func (o *CSSCMSDataModelModelsSSLSslScanResult) HasCertificateCN() bool`

HasCertificateCN returns a boolean if a field has been set.

### SetCertificateCNNil

`func (o *CSSCMSDataModelModelsSSLSslScanResult) SetCertificateCNNil(b bool)`

 SetCertificateCNNil sets the value for CertificateCN to be an explicit nil

### UnsetCertificateCN
`func (o *CSSCMSDataModelModelsSSLSslScanResult) UnsetCertificateCN()`

UnsetCertificateCN ensures that no value is present for CertificateCN, not even an explicit nil
### GetReviewed

`func (o *CSSCMSDataModelModelsSSLSslScanResult) GetReviewed() bool`

GetReviewed returns the Reviewed field if non-nil, zero value otherwise.

### GetReviewedOk

`func (o *CSSCMSDataModelModelsSSLSslScanResult) GetReviewedOk() (*bool, bool)`

GetReviewedOk returns a tuple with the Reviewed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewed

`func (o *CSSCMSDataModelModelsSSLSslScanResult) SetReviewed(v bool)`

SetReviewed sets Reviewed field to given value.

### HasReviewed

`func (o *CSSCMSDataModelModelsSSLSslScanResult) HasReviewed() bool`

HasReviewed returns a boolean if a field has been set.

### GetExpirationDate

`func (o *CSSCMSDataModelModelsSSLSslScanResult) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *CSSCMSDataModelModelsSSLSslScanResult) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *CSSCMSDataModelModelsSSLSslScanResult) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *CSSCMSDataModelModelsSSLSslScanResult) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### SetExpirationDateNil

`func (o *CSSCMSDataModelModelsSSLSslScanResult) SetExpirationDateNil(b bool)`

 SetExpirationDateNil sets the value for ExpirationDate to be an explicit nil

### UnsetExpirationDate
`func (o *CSSCMSDataModelModelsSSLSslScanResult) UnsetExpirationDate()`

UnsetExpirationDate ensures that no value is present for ExpirationDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


