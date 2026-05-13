# CertificatesCertificateRiskIntelligenceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**TotalScore** | Pointer to **int32** |  | [optional] 
**ValidationResults** | Pointer to [**[]CertificatesCertificateValidationRuleResponse**](CertificatesCertificateValidationRuleResponse.md) |  | [optional] 

## Methods

### NewCertificatesCertificateRiskIntelligenceResponse

`func NewCertificatesCertificateRiskIntelligenceResponse() *CertificatesCertificateRiskIntelligenceResponse`

NewCertificatesCertificateRiskIntelligenceResponse instantiates a new CertificatesCertificateRiskIntelligenceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificatesCertificateRiskIntelligenceResponseWithDefaults

`func NewCertificatesCertificateRiskIntelligenceResponseWithDefaults() *CertificatesCertificateRiskIntelligenceResponse`

NewCertificatesCertificateRiskIntelligenceResponseWithDefaults instantiates a new CertificatesCertificateRiskIntelligenceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CertificatesCertificateRiskIntelligenceResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CertificatesCertificateRiskIntelligenceResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CertificatesCertificateRiskIntelligenceResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CertificatesCertificateRiskIntelligenceResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTotalScore

`func (o *CertificatesCertificateRiskIntelligenceResponse) GetTotalScore() int32`

GetTotalScore returns the TotalScore field if non-nil, zero value otherwise.

### GetTotalScoreOk

`func (o *CertificatesCertificateRiskIntelligenceResponse) GetTotalScoreOk() (*int32, bool)`

GetTotalScoreOk returns a tuple with the TotalScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalScore

`func (o *CertificatesCertificateRiskIntelligenceResponse) SetTotalScore(v int32)`

SetTotalScore sets TotalScore field to given value.

### HasTotalScore

`func (o *CertificatesCertificateRiskIntelligenceResponse) HasTotalScore() bool`

HasTotalScore returns a boolean if a field has been set.

### GetValidationResults

`func (o *CertificatesCertificateRiskIntelligenceResponse) GetValidationResults() []CertificatesCertificateValidationRuleResponse`

GetValidationResults returns the ValidationResults field if non-nil, zero value otherwise.

### GetValidationResultsOk

`func (o *CertificatesCertificateRiskIntelligenceResponse) GetValidationResultsOk() (*[]CertificatesCertificateValidationRuleResponse, bool)`

GetValidationResultsOk returns a tuple with the ValidationResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationResults

`func (o *CertificatesCertificateRiskIntelligenceResponse) SetValidationResults(v []CertificatesCertificateValidationRuleResponse)`

SetValidationResults sets ValidationResults field to given value.

### HasValidationResults

`func (o *CertificatesCertificateRiskIntelligenceResponse) HasValidationResults() bool`

HasValidationResults returns a boolean if a field has been set.

### SetValidationResultsNil

`func (o *CertificatesCertificateRiskIntelligenceResponse) SetValidationResultsNil(b bool)`

 SetValidationResultsNil sets the value for ValidationResults to be an explicit nil

### UnsetValidationResults
`func (o *CertificatesCertificateRiskIntelligenceResponse) UnsetValidationResults()`

UnsetValidationResults ensures that no value is present for ValidationResults, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


