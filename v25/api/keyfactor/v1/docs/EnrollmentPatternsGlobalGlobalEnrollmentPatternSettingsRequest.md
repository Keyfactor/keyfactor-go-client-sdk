# EnrollmentPatternsGlobalGlobalEnrollmentPatternSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Regexes** | [**[]EnrollmentPatternsEnrollmentPatternRegexesRequest**](EnrollmentPatternsEnrollmentPatternRegexesRequest.md) | The regular expressions to use for validation during enrollment. | 
**Defaults** | [**[]EnrollmentPatternsEnrollmentPatternDefaultRequest**](EnrollmentPatternsEnrollmentPatternDefaultRequest.md) | The default values to use during enrollment. | 
**Policies** | [**EnrollmentPatternsGlobalGlobalEnrollmentPatternPolicyRequest**](EnrollmentPatternsGlobalGlobalEnrollmentPatternPolicyRequest.md) |  | 

## Methods

### NewEnrollmentPatternsGlobalGlobalEnrollmentPatternSettingsRequest

`func NewEnrollmentPatternsGlobalGlobalEnrollmentPatternSettingsRequest(regexes []EnrollmentPatternsEnrollmentPatternRegexesRequest, defaults []EnrollmentPatternsEnrollmentPatternDefaultRequest, policies EnrollmentPatternsGlobalGlobalEnrollmentPatternPolicyRequest, ) *EnrollmentPatternsGlobalGlobalEnrollmentPatternSettingsRequest`

NewEnrollmentPatternsGlobalGlobalEnrollmentPatternSettingsRequest instantiates a new EnrollmentPatternsGlobalGlobalEnrollmentPatternSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnrollmentPatternsGlobalGlobalEnrollmentPatternSettingsRequestWithDefaults

`func NewEnrollmentPatternsGlobalGlobalEnrollmentPatternSettingsRequestWithDefaults() *EnrollmentPatternsGlobalGlobalEnrollmentPatternSettingsRequest`

NewEnrollmentPatternsGlobalGlobalEnrollmentPatternSettingsRequestWithDefaults instantiates a new EnrollmentPatternsGlobalGlobalEnrollmentPatternSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegexes

`func (o *EnrollmentPatternsGlobalGlobalEnrollmentPatternSettingsRequest) GetRegexes() []EnrollmentPatternsEnrollmentPatternRegexesRequest`

GetRegexes returns the Regexes field if non-nil, zero value otherwise.

### GetRegexesOk

`func (o *EnrollmentPatternsGlobalGlobalEnrollmentPatternSettingsRequest) GetRegexesOk() (*[]EnrollmentPatternsEnrollmentPatternRegexesRequest, bool)`

GetRegexesOk returns a tuple with the Regexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegexes

`func (o *EnrollmentPatternsGlobalGlobalEnrollmentPatternSettingsRequest) SetRegexes(v []EnrollmentPatternsEnrollmentPatternRegexesRequest)`

SetRegexes sets Regexes field to given value.


### GetDefaults

`func (o *EnrollmentPatternsGlobalGlobalEnrollmentPatternSettingsRequest) GetDefaults() []EnrollmentPatternsEnrollmentPatternDefaultRequest`

GetDefaults returns the Defaults field if non-nil, zero value otherwise.

### GetDefaultsOk

`func (o *EnrollmentPatternsGlobalGlobalEnrollmentPatternSettingsRequest) GetDefaultsOk() (*[]EnrollmentPatternsEnrollmentPatternDefaultRequest, bool)`

GetDefaultsOk returns a tuple with the Defaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaults

`func (o *EnrollmentPatternsGlobalGlobalEnrollmentPatternSettingsRequest) SetDefaults(v []EnrollmentPatternsEnrollmentPatternDefaultRequest)`

SetDefaults sets Defaults field to given value.


### GetPolicies

`func (o *EnrollmentPatternsGlobalGlobalEnrollmentPatternSettingsRequest) GetPolicies() EnrollmentPatternsGlobalGlobalEnrollmentPatternPolicyRequest`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *EnrollmentPatternsGlobalGlobalEnrollmentPatternSettingsRequest) GetPoliciesOk() (*EnrollmentPatternsGlobalGlobalEnrollmentPatternPolicyRequest, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *EnrollmentPatternsGlobalGlobalEnrollmentPatternSettingsRequest) SetPolicies(v EnrollmentPatternsGlobalGlobalEnrollmentPatternPolicyRequest)`

SetPolicies sets Policies field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


