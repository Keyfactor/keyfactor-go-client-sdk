# CertificatesCertificateValidationRuleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**Violations** | Pointer to [**[]CertificatesCertificateViolationResponse**](CertificatesCertificateViolationResponse.md) |  | [optional] 

## Methods

### NewCertificatesCertificateValidationRuleResponse

`func NewCertificatesCertificateValidationRuleResponse() *CertificatesCertificateValidationRuleResponse`

NewCertificatesCertificateValidationRuleResponse instantiates a new CertificatesCertificateValidationRuleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificatesCertificateValidationRuleResponseWithDefaults

`func NewCertificatesCertificateValidationRuleResponseWithDefaults() *CertificatesCertificateValidationRuleResponse`

NewCertificatesCertificateValidationRuleResponseWithDefaults instantiates a new CertificatesCertificateValidationRuleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CertificatesCertificateValidationRuleResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CertificatesCertificateValidationRuleResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CertificatesCertificateValidationRuleResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CertificatesCertificateValidationRuleResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CertificatesCertificateValidationRuleResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CertificatesCertificateValidationRuleResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetViolations

`func (o *CertificatesCertificateValidationRuleResponse) GetViolations() []CertificatesCertificateViolationResponse`

GetViolations returns the Violations field if non-nil, zero value otherwise.

### GetViolationsOk

`func (o *CertificatesCertificateValidationRuleResponse) GetViolationsOk() (*[]CertificatesCertificateViolationResponse, bool)`

GetViolationsOk returns a tuple with the Violations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViolations

`func (o *CertificatesCertificateValidationRuleResponse) SetViolations(v []CertificatesCertificateViolationResponse)`

SetViolations sets Violations field to given value.

### HasViolations

`func (o *CertificatesCertificateValidationRuleResponse) HasViolations() bool`

HasViolations returns a boolean if a field has been set.

### SetViolationsNil

`func (o *CertificatesCertificateValidationRuleResponse) SetViolationsNil(b bool)`

 SetViolationsNil sets the value for Violations to be an explicit nil

### UnsetViolations
`func (o *CertificatesCertificateValidationRuleResponse) UnsetViolations()`

UnsetViolations ensures that no value is present for Violations, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


