# CertificatesCertificateViolationUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Score** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**RemediationCode** | **string** |  | 

## Methods

### NewCertificatesCertificateViolationUpdateRequest

`func NewCertificatesCertificateViolationUpdateRequest(remediationCode string, ) *CertificatesCertificateViolationUpdateRequest`

NewCertificatesCertificateViolationUpdateRequest instantiates a new CertificatesCertificateViolationUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificatesCertificateViolationUpdateRequestWithDefaults

`func NewCertificatesCertificateViolationUpdateRequestWithDefaults() *CertificatesCertificateViolationUpdateRequest`

NewCertificatesCertificateViolationUpdateRequestWithDefaults instantiates a new CertificatesCertificateViolationUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScore

`func (o *CertificatesCertificateViolationUpdateRequest) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *CertificatesCertificateViolationUpdateRequest) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *CertificatesCertificateViolationUpdateRequest) SetScore(v int32)`

SetScore sets Score field to given value.

### HasScore

`func (o *CertificatesCertificateViolationUpdateRequest) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetDescription

`func (o *CertificatesCertificateViolationUpdateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CertificatesCertificateViolationUpdateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CertificatesCertificateViolationUpdateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CertificatesCertificateViolationUpdateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CertificatesCertificateViolationUpdateRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CertificatesCertificateViolationUpdateRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetRemediationCode

`func (o *CertificatesCertificateViolationUpdateRequest) GetRemediationCode() string`

GetRemediationCode returns the RemediationCode field if non-nil, zero value otherwise.

### GetRemediationCodeOk

`func (o *CertificatesCertificateViolationUpdateRequest) GetRemediationCodeOk() (*string, bool)`

GetRemediationCodeOk returns a tuple with the RemediationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediationCode

`func (o *CertificatesCertificateViolationUpdateRequest) SetRemediationCode(v string)`

SetRemediationCode sets RemediationCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


