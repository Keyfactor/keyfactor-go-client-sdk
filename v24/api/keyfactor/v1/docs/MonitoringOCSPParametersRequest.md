# MonitoringOCSPParametersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertificateContents** | Pointer to **NullableString** |  | [optional] 
**CertificateAuthorityId** | Pointer to **NullableInt32** |  | [optional] 
**AuthorityName** | Pointer to **NullableString** |  | [optional] 
**AuthorityNameId** | Pointer to **NullableString** |  | [optional] 
**AuthorityKeyId** | Pointer to **NullableString** |  | [optional] 
**SampleSerialNumber** | Pointer to **NullableString** |  | [optional] 
**FileName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewMonitoringOCSPParametersRequest

`func NewMonitoringOCSPParametersRequest() *MonitoringOCSPParametersRequest`

NewMonitoringOCSPParametersRequest instantiates a new MonitoringOCSPParametersRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringOCSPParametersRequestWithDefaults

`func NewMonitoringOCSPParametersRequestWithDefaults() *MonitoringOCSPParametersRequest`

NewMonitoringOCSPParametersRequestWithDefaults instantiates a new MonitoringOCSPParametersRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificateContents

`func (o *MonitoringOCSPParametersRequest) GetCertificateContents() string`

GetCertificateContents returns the CertificateContents field if non-nil, zero value otherwise.

### GetCertificateContentsOk

`func (o *MonitoringOCSPParametersRequest) GetCertificateContentsOk() (*string, bool)`

GetCertificateContentsOk returns a tuple with the CertificateContents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateContents

`func (o *MonitoringOCSPParametersRequest) SetCertificateContents(v string)`

SetCertificateContents sets CertificateContents field to given value.

### HasCertificateContents

`func (o *MonitoringOCSPParametersRequest) HasCertificateContents() bool`

HasCertificateContents returns a boolean if a field has been set.

### SetCertificateContentsNil

`func (o *MonitoringOCSPParametersRequest) SetCertificateContentsNil(b bool)`

 SetCertificateContentsNil sets the value for CertificateContents to be an explicit nil

### UnsetCertificateContents
`func (o *MonitoringOCSPParametersRequest) UnsetCertificateContents()`

UnsetCertificateContents ensures that no value is present for CertificateContents, not even an explicit nil
### GetCertificateAuthorityId

`func (o *MonitoringOCSPParametersRequest) GetCertificateAuthorityId() int32`

GetCertificateAuthorityId returns the CertificateAuthorityId field if non-nil, zero value otherwise.

### GetCertificateAuthorityIdOk

`func (o *MonitoringOCSPParametersRequest) GetCertificateAuthorityIdOk() (*int32, bool)`

GetCertificateAuthorityIdOk returns a tuple with the CertificateAuthorityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateAuthorityId

`func (o *MonitoringOCSPParametersRequest) SetCertificateAuthorityId(v int32)`

SetCertificateAuthorityId sets CertificateAuthorityId field to given value.

### HasCertificateAuthorityId

`func (o *MonitoringOCSPParametersRequest) HasCertificateAuthorityId() bool`

HasCertificateAuthorityId returns a boolean if a field has been set.

### SetCertificateAuthorityIdNil

`func (o *MonitoringOCSPParametersRequest) SetCertificateAuthorityIdNil(b bool)`

 SetCertificateAuthorityIdNil sets the value for CertificateAuthorityId to be an explicit nil

### UnsetCertificateAuthorityId
`func (o *MonitoringOCSPParametersRequest) UnsetCertificateAuthorityId()`

UnsetCertificateAuthorityId ensures that no value is present for CertificateAuthorityId, not even an explicit nil
### GetAuthorityName

`func (o *MonitoringOCSPParametersRequest) GetAuthorityName() string`

GetAuthorityName returns the AuthorityName field if non-nil, zero value otherwise.

### GetAuthorityNameOk

`func (o *MonitoringOCSPParametersRequest) GetAuthorityNameOk() (*string, bool)`

GetAuthorityNameOk returns a tuple with the AuthorityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorityName

`func (o *MonitoringOCSPParametersRequest) SetAuthorityName(v string)`

SetAuthorityName sets AuthorityName field to given value.

### HasAuthorityName

`func (o *MonitoringOCSPParametersRequest) HasAuthorityName() bool`

HasAuthorityName returns a boolean if a field has been set.

### SetAuthorityNameNil

`func (o *MonitoringOCSPParametersRequest) SetAuthorityNameNil(b bool)`

 SetAuthorityNameNil sets the value for AuthorityName to be an explicit nil

### UnsetAuthorityName
`func (o *MonitoringOCSPParametersRequest) UnsetAuthorityName()`

UnsetAuthorityName ensures that no value is present for AuthorityName, not even an explicit nil
### GetAuthorityNameId

`func (o *MonitoringOCSPParametersRequest) GetAuthorityNameId() string`

GetAuthorityNameId returns the AuthorityNameId field if non-nil, zero value otherwise.

### GetAuthorityNameIdOk

`func (o *MonitoringOCSPParametersRequest) GetAuthorityNameIdOk() (*string, bool)`

GetAuthorityNameIdOk returns a tuple with the AuthorityNameId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorityNameId

`func (o *MonitoringOCSPParametersRequest) SetAuthorityNameId(v string)`

SetAuthorityNameId sets AuthorityNameId field to given value.

### HasAuthorityNameId

`func (o *MonitoringOCSPParametersRequest) HasAuthorityNameId() bool`

HasAuthorityNameId returns a boolean if a field has been set.

### SetAuthorityNameIdNil

`func (o *MonitoringOCSPParametersRequest) SetAuthorityNameIdNil(b bool)`

 SetAuthorityNameIdNil sets the value for AuthorityNameId to be an explicit nil

### UnsetAuthorityNameId
`func (o *MonitoringOCSPParametersRequest) UnsetAuthorityNameId()`

UnsetAuthorityNameId ensures that no value is present for AuthorityNameId, not even an explicit nil
### GetAuthorityKeyId

`func (o *MonitoringOCSPParametersRequest) GetAuthorityKeyId() string`

GetAuthorityKeyId returns the AuthorityKeyId field if non-nil, zero value otherwise.

### GetAuthorityKeyIdOk

`func (o *MonitoringOCSPParametersRequest) GetAuthorityKeyIdOk() (*string, bool)`

GetAuthorityKeyIdOk returns a tuple with the AuthorityKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorityKeyId

`func (o *MonitoringOCSPParametersRequest) SetAuthorityKeyId(v string)`

SetAuthorityKeyId sets AuthorityKeyId field to given value.

### HasAuthorityKeyId

`func (o *MonitoringOCSPParametersRequest) HasAuthorityKeyId() bool`

HasAuthorityKeyId returns a boolean if a field has been set.

### SetAuthorityKeyIdNil

`func (o *MonitoringOCSPParametersRequest) SetAuthorityKeyIdNil(b bool)`

 SetAuthorityKeyIdNil sets the value for AuthorityKeyId to be an explicit nil

### UnsetAuthorityKeyId
`func (o *MonitoringOCSPParametersRequest) UnsetAuthorityKeyId()`

UnsetAuthorityKeyId ensures that no value is present for AuthorityKeyId, not even an explicit nil
### GetSampleSerialNumber

`func (o *MonitoringOCSPParametersRequest) GetSampleSerialNumber() string`

GetSampleSerialNumber returns the SampleSerialNumber field if non-nil, zero value otherwise.

### GetSampleSerialNumberOk

`func (o *MonitoringOCSPParametersRequest) GetSampleSerialNumberOk() (*string, bool)`

GetSampleSerialNumberOk returns a tuple with the SampleSerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleSerialNumber

`func (o *MonitoringOCSPParametersRequest) SetSampleSerialNumber(v string)`

SetSampleSerialNumber sets SampleSerialNumber field to given value.

### HasSampleSerialNumber

`func (o *MonitoringOCSPParametersRequest) HasSampleSerialNumber() bool`

HasSampleSerialNumber returns a boolean if a field has been set.

### SetSampleSerialNumberNil

`func (o *MonitoringOCSPParametersRequest) SetSampleSerialNumberNil(b bool)`

 SetSampleSerialNumberNil sets the value for SampleSerialNumber to be an explicit nil

### UnsetSampleSerialNumber
`func (o *MonitoringOCSPParametersRequest) UnsetSampleSerialNumber()`

UnsetSampleSerialNumber ensures that no value is present for SampleSerialNumber, not even an explicit nil
### GetFileName

`func (o *MonitoringOCSPParametersRequest) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *MonitoringOCSPParametersRequest) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *MonitoringOCSPParametersRequest) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *MonitoringOCSPParametersRequest) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### SetFileNameNil

`func (o *MonitoringOCSPParametersRequest) SetFileNameNil(b bool)`

 SetFileNameNil sets the value for FileName to be an explicit nil

### UnsetFileName
`func (o *MonitoringOCSPParametersRequest) UnsetFileName()`

UnsetFileName ensures that no value is present for FileName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


