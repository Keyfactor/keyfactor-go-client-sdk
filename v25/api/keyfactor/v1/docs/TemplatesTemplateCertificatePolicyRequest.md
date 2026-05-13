# TemplatesTemplateCertificatePolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicyID** | **string** |  | 
**PolicyType** | [**KeyfactorPlatformExtensionsEnumsCertificatePolicyType**](KeyfactorPlatformExtensionsEnumsCertificatePolicyType.md) |  | 
**Value** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTemplatesTemplateCertificatePolicyRequest

`func NewTemplatesTemplateCertificatePolicyRequest(policyID string, policyType KeyfactorPlatformExtensionsEnumsCertificatePolicyType, ) *TemplatesTemplateCertificatePolicyRequest`

NewTemplatesTemplateCertificatePolicyRequest instantiates a new TemplatesTemplateCertificatePolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplatesTemplateCertificatePolicyRequestWithDefaults

`func NewTemplatesTemplateCertificatePolicyRequestWithDefaults() *TemplatesTemplateCertificatePolicyRequest`

NewTemplatesTemplateCertificatePolicyRequestWithDefaults instantiates a new TemplatesTemplateCertificatePolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicyID

`func (o *TemplatesTemplateCertificatePolicyRequest) GetPolicyID() string`

GetPolicyID returns the PolicyID field if non-nil, zero value otherwise.

### GetPolicyIDOk

`func (o *TemplatesTemplateCertificatePolicyRequest) GetPolicyIDOk() (*string, bool)`

GetPolicyIDOk returns a tuple with the PolicyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyID

`func (o *TemplatesTemplateCertificatePolicyRequest) SetPolicyID(v string)`

SetPolicyID sets PolicyID field to given value.


### GetPolicyType

`func (o *TemplatesTemplateCertificatePolicyRequest) GetPolicyType() KeyfactorPlatformExtensionsEnumsCertificatePolicyType`

GetPolicyType returns the PolicyType field if non-nil, zero value otherwise.

### GetPolicyTypeOk

`func (o *TemplatesTemplateCertificatePolicyRequest) GetPolicyTypeOk() (*KeyfactorPlatformExtensionsEnumsCertificatePolicyType, bool)`

GetPolicyTypeOk returns a tuple with the PolicyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyType

`func (o *TemplatesTemplateCertificatePolicyRequest) SetPolicyType(v KeyfactorPlatformExtensionsEnumsCertificatePolicyType)`

SetPolicyType sets PolicyType field to given value.


### GetValue

`func (o *TemplatesTemplateCertificatePolicyRequest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TemplatesTemplateCertificatePolicyRequest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TemplatesTemplateCertificatePolicyRequest) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *TemplatesTemplateCertificatePolicyRequest) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *TemplatesTemplateCertificatePolicyRequest) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *TemplatesTemplateCertificatePolicyRequest) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


