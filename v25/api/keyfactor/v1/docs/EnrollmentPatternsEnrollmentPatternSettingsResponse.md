# EnrollmentPatternsEnrollmentPatternSettingsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Regexes** | Pointer to [**[]EnrollmentPatternsEnrollmentPatternRegexesResponse**](EnrollmentPatternsEnrollmentPatternRegexesResponse.md) | The regular expressions to use for validation during enrollment. | [optional] 
**Defaults** | Pointer to [**[]EnrollmentPatternsEnrollmentPatternDefaultResponse**](EnrollmentPatternsEnrollmentPatternDefaultResponse.md) | The default values to use during enrollment. | [optional] 
**Policies** | Pointer to [**EnrollmentPatternsEnrollmentPatternPolicyResponse**](EnrollmentPatternsEnrollmentPatternPolicyResponse.md) |  | [optional] 

## Methods

### NewEnrollmentPatternsEnrollmentPatternSettingsResponse

`func NewEnrollmentPatternsEnrollmentPatternSettingsResponse() *EnrollmentPatternsEnrollmentPatternSettingsResponse`

NewEnrollmentPatternsEnrollmentPatternSettingsResponse instantiates a new EnrollmentPatternsEnrollmentPatternSettingsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnrollmentPatternsEnrollmentPatternSettingsResponseWithDefaults

`func NewEnrollmentPatternsEnrollmentPatternSettingsResponseWithDefaults() *EnrollmentPatternsEnrollmentPatternSettingsResponse`

NewEnrollmentPatternsEnrollmentPatternSettingsResponseWithDefaults instantiates a new EnrollmentPatternsEnrollmentPatternSettingsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegexes

`func (o *EnrollmentPatternsEnrollmentPatternSettingsResponse) GetRegexes() []EnrollmentPatternsEnrollmentPatternRegexesResponse`

GetRegexes returns the Regexes field if non-nil, zero value otherwise.

### GetRegexesOk

`func (o *EnrollmentPatternsEnrollmentPatternSettingsResponse) GetRegexesOk() (*[]EnrollmentPatternsEnrollmentPatternRegexesResponse, bool)`

GetRegexesOk returns a tuple with the Regexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegexes

`func (o *EnrollmentPatternsEnrollmentPatternSettingsResponse) SetRegexes(v []EnrollmentPatternsEnrollmentPatternRegexesResponse)`

SetRegexes sets Regexes field to given value.

### HasRegexes

`func (o *EnrollmentPatternsEnrollmentPatternSettingsResponse) HasRegexes() bool`

HasRegexes returns a boolean if a field has been set.

### SetRegexesNil

`func (o *EnrollmentPatternsEnrollmentPatternSettingsResponse) SetRegexesNil(b bool)`

 SetRegexesNil sets the value for Regexes to be an explicit nil

### UnsetRegexes
`func (o *EnrollmentPatternsEnrollmentPatternSettingsResponse) UnsetRegexes()`

UnsetRegexes ensures that no value is present for Regexes, not even an explicit nil
### GetDefaults

`func (o *EnrollmentPatternsEnrollmentPatternSettingsResponse) GetDefaults() []EnrollmentPatternsEnrollmentPatternDefaultResponse`

GetDefaults returns the Defaults field if non-nil, zero value otherwise.

### GetDefaultsOk

`func (o *EnrollmentPatternsEnrollmentPatternSettingsResponse) GetDefaultsOk() (*[]EnrollmentPatternsEnrollmentPatternDefaultResponse, bool)`

GetDefaultsOk returns a tuple with the Defaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaults

`func (o *EnrollmentPatternsEnrollmentPatternSettingsResponse) SetDefaults(v []EnrollmentPatternsEnrollmentPatternDefaultResponse)`

SetDefaults sets Defaults field to given value.

### HasDefaults

`func (o *EnrollmentPatternsEnrollmentPatternSettingsResponse) HasDefaults() bool`

HasDefaults returns a boolean if a field has been set.

### SetDefaultsNil

`func (o *EnrollmentPatternsEnrollmentPatternSettingsResponse) SetDefaultsNil(b bool)`

 SetDefaultsNil sets the value for Defaults to be an explicit nil

### UnsetDefaults
`func (o *EnrollmentPatternsEnrollmentPatternSettingsResponse) UnsetDefaults()`

UnsetDefaults ensures that no value is present for Defaults, not even an explicit nil
### GetPolicies

`func (o *EnrollmentPatternsEnrollmentPatternSettingsResponse) GetPolicies() EnrollmentPatternsEnrollmentPatternPolicyResponse`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *EnrollmentPatternsEnrollmentPatternSettingsResponse) GetPoliciesOk() (*EnrollmentPatternsEnrollmentPatternPolicyResponse, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *EnrollmentPatternsEnrollmentPatternSettingsResponse) SetPolicies(v EnrollmentPatternsEnrollmentPatternPolicyResponse)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *EnrollmentPatternsEnrollmentPatternSettingsResponse) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


