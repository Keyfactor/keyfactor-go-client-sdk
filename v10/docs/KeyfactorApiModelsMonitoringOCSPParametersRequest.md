# KeyfactorApiModelsMonitoringOCSPParametersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertificateContents** | Pointer to **string** |  | [optional] 
**CertificateAuthorityId** | Pointer to **int64** |  | [optional] 
**AuthorityName** | Pointer to **string** |  | [optional] 
**AuthorityNameId** | Pointer to **string** |  | [optional] 
**AuthorityKeyId** | Pointer to **string** |  | [optional] 
**SampleSerialNumber** | Pointer to **string** |  | [optional] 

## Methods

### NewKeyfactorApiModelsMonitoringOCSPParametersRequest

`func NewKeyfactorApiModelsMonitoringOCSPParametersRequest() *KeyfactorApiModelsMonitoringOCSPParametersRequest`

NewKeyfactorApiModelsMonitoringOCSPParametersRequest instantiates a new KeyfactorApiModelsMonitoringOCSPParametersRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorApiModelsMonitoringOCSPParametersRequestWithDefaults

`func NewKeyfactorApiModelsMonitoringOCSPParametersRequestWithDefaults() *KeyfactorApiModelsMonitoringOCSPParametersRequest`

NewKeyfactorApiModelsMonitoringOCSPParametersRequestWithDefaults instantiates a new KeyfactorApiModelsMonitoringOCSPParametersRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificateContents

`func (o *KeyfactorApiModelsMonitoringOCSPParametersRequest) GetCertificateContents() string`

GetCertificateContents returns the CertificateContents field if non-nil, zero value otherwise.

### GetCertificateContentsOk

`func (o *KeyfactorApiModelsMonitoringOCSPParametersRequest) GetCertificateContentsOk() (*string, bool)`

GetCertificateContentsOk returns a tuple with the CertificateContents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateContents

`func (o *KeyfactorApiModelsMonitoringOCSPParametersRequest) SetCertificateContents(v string)`

SetCertificateContents sets CertificateContents field to given value.

### HasCertificateContents

`func (o *KeyfactorApiModelsMonitoringOCSPParametersRequest) HasCertificateContents() bool`

HasCertificateContents returns a boolean if a field has been set.

### GetCertificateAuthorityId

`func (o *KeyfactorApiModelsMonitoringOCSPParametersRequest) GetCertificateAuthorityId() int64`

GetCertificateAuthorityId returns the CertificateAuthorityId field if non-nil, zero value otherwise.

### GetCertificateAuthorityIdOk

`func (o *KeyfactorApiModelsMonitoringOCSPParametersRequest) GetCertificateAuthorityIdOk() (*int64, bool)`

GetCertificateAuthorityIdOk returns a tuple with the CertificateAuthorityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateAuthorityId

`func (o *KeyfactorApiModelsMonitoringOCSPParametersRequest) SetCertificateAuthorityId(v int64)`

SetCertificateAuthorityId sets CertificateAuthorityId field to given value.

### HasCertificateAuthorityId

`func (o *KeyfactorApiModelsMonitoringOCSPParametersRequest) HasCertificateAuthorityId() bool`

HasCertificateAuthorityId returns a boolean if a field has been set.

### GetAuthorityName

`func (o *KeyfactorApiModelsMonitoringOCSPParametersRequest) GetAuthorityName() string`

GetAuthorityName returns the AuthorityName field if non-nil, zero value otherwise.

### GetAuthorityNameOk

`func (o *KeyfactorApiModelsMonitoringOCSPParametersRequest) GetAuthorityNameOk() (*string, bool)`

GetAuthorityNameOk returns a tuple with the AuthorityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorityName

`func (o *KeyfactorApiModelsMonitoringOCSPParametersRequest) SetAuthorityName(v string)`

SetAuthorityName sets AuthorityName field to given value.

### HasAuthorityName

`func (o *KeyfactorApiModelsMonitoringOCSPParametersRequest) HasAuthorityName() bool`

HasAuthorityName returns a boolean if a field has been set.

### GetAuthorityNameId

`func (o *KeyfactorApiModelsMonitoringOCSPParametersRequest) GetAuthorityNameId() string`

GetAuthorityNameId returns the AuthorityNameId field if non-nil, zero value otherwise.

### GetAuthorityNameIdOk

`func (o *KeyfactorApiModelsMonitoringOCSPParametersRequest) GetAuthorityNameIdOk() (*string, bool)`

GetAuthorityNameIdOk returns a tuple with the AuthorityNameId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorityNameId

`func (o *KeyfactorApiModelsMonitoringOCSPParametersRequest) SetAuthorityNameId(v string)`

SetAuthorityNameId sets AuthorityNameId field to given value.

### HasAuthorityNameId

`func (o *KeyfactorApiModelsMonitoringOCSPParametersRequest) HasAuthorityNameId() bool`

HasAuthorityNameId returns a boolean if a field has been set.

### GetAuthorityKeyId

`func (o *KeyfactorApiModelsMonitoringOCSPParametersRequest) GetAuthorityKeyId() string`

GetAuthorityKeyId returns the AuthorityKeyId field if non-nil, zero value otherwise.

### GetAuthorityKeyIdOk

`func (o *KeyfactorApiModelsMonitoringOCSPParametersRequest) GetAuthorityKeyIdOk() (*string, bool)`

GetAuthorityKeyIdOk returns a tuple with the AuthorityKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorityKeyId

`func (o *KeyfactorApiModelsMonitoringOCSPParametersRequest) SetAuthorityKeyId(v string)`

SetAuthorityKeyId sets AuthorityKeyId field to given value.

### HasAuthorityKeyId

`func (o *KeyfactorApiModelsMonitoringOCSPParametersRequest) HasAuthorityKeyId() bool`

HasAuthorityKeyId returns a boolean if a field has been set.

### GetSampleSerialNumber

`func (o *KeyfactorApiModelsMonitoringOCSPParametersRequest) GetSampleSerialNumber() string`

GetSampleSerialNumber returns the SampleSerialNumber field if non-nil, zero value otherwise.

### GetSampleSerialNumberOk

`func (o *KeyfactorApiModelsMonitoringOCSPParametersRequest) GetSampleSerialNumberOk() (*string, bool)`

GetSampleSerialNumberOk returns a tuple with the SampleSerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleSerialNumber

`func (o *KeyfactorApiModelsMonitoringOCSPParametersRequest) SetSampleSerialNumber(v string)`

SetSampleSerialNumber sets SampleSerialNumber field to given value.

### HasSampleSerialNumber

`func (o *KeyfactorApiModelsMonitoringOCSPParametersRequest) HasSampleSerialNumber() bool`

HasSampleSerialNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


