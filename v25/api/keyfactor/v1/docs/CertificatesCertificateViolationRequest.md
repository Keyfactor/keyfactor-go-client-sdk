# CertificatesCertificateViolationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Score** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**RemediationCode** | **string** |  | 

## Methods

### NewCertificatesCertificateViolationRequest

`func NewCertificatesCertificateViolationRequest(name string, remediationCode string, ) *CertificatesCertificateViolationRequest`

NewCertificatesCertificateViolationRequest instantiates a new CertificatesCertificateViolationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificatesCertificateViolationRequestWithDefaults

`func NewCertificatesCertificateViolationRequestWithDefaults() *CertificatesCertificateViolationRequest`

NewCertificatesCertificateViolationRequestWithDefaults instantiates a new CertificatesCertificateViolationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CertificatesCertificateViolationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CertificatesCertificateViolationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CertificatesCertificateViolationRequest) SetName(v string)`

SetName sets Name field to given value.


### GetScore

`func (o *CertificatesCertificateViolationRequest) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *CertificatesCertificateViolationRequest) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *CertificatesCertificateViolationRequest) SetScore(v int32)`

SetScore sets Score field to given value.

### HasScore

`func (o *CertificatesCertificateViolationRequest) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetDescription

`func (o *CertificatesCertificateViolationRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CertificatesCertificateViolationRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CertificatesCertificateViolationRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CertificatesCertificateViolationRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CertificatesCertificateViolationRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CertificatesCertificateViolationRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetRemediationCode

`func (o *CertificatesCertificateViolationRequest) GetRemediationCode() string`

GetRemediationCode returns the RemediationCode field if non-nil, zero value otherwise.

### GetRemediationCodeOk

`func (o *CertificatesCertificateViolationRequest) GetRemediationCodeOk() (*string, bool)`

GetRemediationCodeOk returns a tuple with the RemediationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediationCode

`func (o *CertificatesCertificateViolationRequest) SetRemediationCode(v string)`

SetRemediationCode sets RemediationCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


