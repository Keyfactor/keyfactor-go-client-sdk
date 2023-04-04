# ModelsCertificateAuthoritiesCertificateAuthorityRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**LogicalName** | Pointer to **string** |  | [optional] 
**HostName** | Pointer to **string** |  | [optional] 
**Delegate** | Pointer to **bool** |  | [optional] 
**DelegateEnrollment** | Pointer to **bool** |  | [optional] 
**ForestRoot** | Pointer to **string** |  | [optional] 
**ConfigurationTenant** | Pointer to **string** |  | [optional] 
**Remote** | Pointer to **bool** |  | [optional] 
**Agent** | Pointer to **string** |  | [optional] 
**Standalone** | Pointer to **bool** |  | [optional] 
**MonitorThresholds** | Pointer to **bool** |  | [optional] 
**IssuanceMax** | Pointer to **int32** |  | [optional] 
**IssuanceMin** | Pointer to **int32** |  | [optional] 
**FailureMax** | Pointer to **int32** |  | [optional] 
**RFCEnforcement** | Pointer to **bool** |  | [optional] 
**Properties** | Pointer to **string** |  | [optional] 
**AllowedEnrollmentTypes** | Pointer to **int32** |  | [optional] 
**KeyRetention** | Pointer to **int32** |  | [optional] 
**KeyRetentionDays** | Pointer to **int32** |  | [optional] 
**ExplicitCredentials** | Pointer to **bool** |  | [optional] 
**SubscriberTerms** | Pointer to **bool** |  | [optional] 
**ExplicitUser** | Pointer to **string** |  | [optional] 
**ExplicitPassword** | Pointer to [**ModelsKeyfactorAPISecret**](ModelsKeyfactorAPISecret.md) |  | [optional] 
**UseAllowedRequesters** | Pointer to **bool** |  | [optional] 
**AllowedRequesters** | Pointer to **[]string** |  | [optional] 
**FullScan** | Pointer to [**KeyfactorCommonSchedulingKeyfactorSchedule**](KeyfactorCommonSchedulingKeyfactorSchedule.md) |  | [optional] 
**IncrementalScan** | Pointer to [**KeyfactorCommonSchedulingKeyfactorSchedule**](KeyfactorCommonSchedulingKeyfactorSchedule.md) |  | [optional] 
**ThresholdCheck** | Pointer to [**KeyfactorCommonSchedulingKeyfactorSchedule**](KeyfactorCommonSchedulingKeyfactorSchedule.md) |  | [optional] 
**AuthCertificatePassword** | Pointer to [**ModelsKeyfactorAPISecret**](ModelsKeyfactorAPISecret.md) |  | [optional] 
**AuthCertificate** | Pointer to [**ModelsKeyfactorAPISecret**](ModelsKeyfactorAPISecret.md) |  | [optional] 
**CAType** | Pointer to **int32** |  | [optional] 
**EnforceUniqueDN** | Pointer to **bool** |  | [optional] 

## Methods

### NewModelsCertificateAuthoritiesCertificateAuthorityRequest

`func NewModelsCertificateAuthoritiesCertificateAuthorityRequest() *ModelsCertificateAuthoritiesCertificateAuthorityRequest`

NewModelsCertificateAuthoritiesCertificateAuthorityRequest instantiates a new ModelsCertificateAuthoritiesCertificateAuthorityRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsCertificateAuthoritiesCertificateAuthorityRequestWithDefaults

`func NewModelsCertificateAuthoritiesCertificateAuthorityRequestWithDefaults() *ModelsCertificateAuthoritiesCertificateAuthorityRequest`

NewModelsCertificateAuthoritiesCertificateAuthorityRequestWithDefaults instantiates a new ModelsCertificateAuthoritiesCertificateAuthorityRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLogicalName

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) GetLogicalName() string`

GetLogicalName returns the LogicalName field if non-nil, zero value otherwise.

### GetLogicalNameOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) GetLogicalNameOk() (*string, bool)`

GetLogicalNameOk returns a tuple with the LogicalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalName

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) SetLogicalName(v string)`

SetLogicalName sets LogicalName field to given value.

### HasLogicalName

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) HasLogicalName() bool`

HasLogicalName returns a boolean if a field has been set.

### GetHostName

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetDelegate

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) GetDelegate() bool`

GetDelegate returns the Delegate field if non-nil, zero value otherwise.

### GetDelegateOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) GetDelegateOk() (*bool, bool)`

GetDelegateOk returns a tuple with the Delegate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegate

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) SetDelegate(v bool)`

SetDelegate sets Delegate field to given value.

### HasDelegate

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) HasDelegate() bool`

HasDelegate returns a boolean if a field has been set.

### GetDelegateEnrollment

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) GetDelegateEnrollment() bool`

GetDelegateEnrollment returns the DelegateEnrollment field if non-nil, zero value otherwise.

### GetDelegateEnrollmentOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) GetDelegateEnrollmentOk() (*bool, bool)`

GetDelegateEnrollmentOk returns a tuple with the DelegateEnrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegateEnrollment

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) SetDelegateEnrollment(v bool)`

SetDelegateEnrollment sets DelegateEnrollment field to given value.

### HasDelegateEnrollment

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) HasDelegateEnrollment() bool`

HasDelegateEnrollment returns a boolean if a field has been set.

### GetForestRoot

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) GetForestRoot() string`

GetForestRoot returns the ForestRoot field if non-nil, zero value otherwise.

### GetForestRootOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) GetForestRootOk() (*string, bool)`

GetForestRootOk returns a tuple with the ForestRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForestRoot

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) SetForestRoot(v string)`

SetForestRoot sets ForestRoot field to given value.

### HasForestRoot

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) HasForestRoot() bool`

HasForestRoot returns a boolean if a field has been set.

### GetConfigurationTenant

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) GetConfigurationTenant() string`

GetConfigurationTenant returns the ConfigurationTenant field if non-nil, zero value otherwise.

### GetConfigurationTenantOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) GetConfigurationTenantOk() (*string, bool)`

GetConfigurationTenantOk returns a tuple with the ConfigurationTenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationTenant

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) SetConfigurationTenant(v string)`

SetConfigurationTenant sets ConfigurationTenant field to given value.

### HasConfigurationTenant

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) HasConfigurationTenant() bool`

HasConfigurationTenant returns a boolean if a field has been set.

### GetRemote

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) GetRemote() bool`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) GetRemoteOk() (*bool, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) SetRemote(v bool)`

SetRemote sets Remote field to given value.

### HasRemote

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) HasRemote() bool`

HasRemote returns a boolean if a field has been set.

### GetAgent

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) GetAgent() string`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) GetAgentOk() (*string, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) SetAgent(v string)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) HasAgent() bool`

HasAgent returns a boolean if a field has been set.

### GetStandalone

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) GetStandalone() bool`

GetStandalone returns the Standalone field if non-nil, zero value otherwise.

### GetStandaloneOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) GetStandaloneOk() (*bool, bool)`

GetStandaloneOk returns a tuple with the Standalone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandalone

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) SetStandalone(v bool)`

SetStandalone sets Standalone field to given value.

### HasStandalone

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) HasStandalone() bool`

HasStandalone returns a boolean if a field has been set.

### GetMonitorThresholds

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) GetMonitorThresholds() bool`

GetMonitorThresholds returns the MonitorThresholds field if non-nil, zero value otherwise.

### GetMonitorThresholdsOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) GetMonitorThresholdsOk() (*bool, bool)`

GetMonitorThresholdsOk returns a tuple with the MonitorThresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorThresholds

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) SetMonitorThresholds(v bool)`

SetMonitorThresholds sets MonitorThresholds field to given value.

### HasMonitorThresholds

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) HasMonitorThresholds() bool`

HasMonitorThresholds returns a boolean if a field has been set.

### GetIssuanceMax

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) GetIssuanceMax() int32`

GetIssuanceMax returns the IssuanceMax field if non-nil, zero value otherwise.

### GetIssuanceMaxOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) GetIssuanceMaxOk() (*int32, bool)`

GetIssuanceMaxOk returns a tuple with the IssuanceMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuanceMax

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) SetIssuanceMax(v int32)`

SetIssuanceMax sets IssuanceMax field to given value.

### HasIssuanceMax

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) HasIssuanceMax() bool`

HasIssuanceMax returns a boolean if a field has been set.

### GetIssuanceMin

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) GetIssuanceMin() int32`

GetIssuanceMin returns the IssuanceMin field if non-nil, zero value otherwise.

### GetIssuanceMinOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) GetIssuanceMinOk() (*int32, bool)`

GetIssuanceMinOk returns a tuple with the IssuanceMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuanceMin

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) SetIssuanceMin(v int32)`

SetIssuanceMin sets IssuanceMin field to given value.

### HasIssuanceMin

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) HasIssuanceMin() bool`

HasIssuanceMin returns a boolean if a field has been set.

### GetFailureMax

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) GetFailureMax() int32`

GetFailureMax returns the FailureMax field if non-nil, zero value otherwise.

### GetFailureMaxOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) GetFailureMaxOk() (*int32, bool)`

GetFailureMaxOk returns a tuple with the FailureMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureMax

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) SetFailureMax(v int32)`

SetFailureMax sets FailureMax field to given value.

### HasFailureMax

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) HasFailureMax() bool`

HasFailureMax returns a boolean if a field has been set.

### GetRFCEnforcement

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) GetRFCEnforcement() bool`

GetRFCEnforcement returns the RFCEnforcement field if non-nil, zero value otherwise.

### GetRFCEnforcementOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) GetRFCEnforcementOk() (*bool, bool)`

GetRFCEnforcementOk returns a tuple with the RFCEnforcement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRFCEnforcement

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) SetRFCEnforcement(v bool)`

SetRFCEnforcement sets RFCEnforcement field to given value.

### HasRFCEnforcement

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) HasRFCEnforcement() bool`

HasRFCEnforcement returns a boolean if a field has been set.

### GetProperties

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) GetProperties() string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) GetPropertiesOk() (*string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) SetProperties(v string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetAllowedEnrollmentTypes

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) GetAllowedEnrollmentTypes() int32`

GetAllowedEnrollmentTypes returns the AllowedEnrollmentTypes field if non-nil, zero value otherwise.

### GetAllowedEnrollmentTypesOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) GetAllowedEnrollmentTypesOk() (*int32, bool)`

GetAllowedEnrollmentTypesOk returns a tuple with the AllowedEnrollmentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedEnrollmentTypes

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) SetAllowedEnrollmentTypes(v int32)`

SetAllowedEnrollmentTypes sets AllowedEnrollmentTypes field to given value.

### HasAllowedEnrollmentTypes

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) HasAllowedEnrollmentTypes() bool`

HasAllowedEnrollmentTypes returns a boolean if a field has been set.

### GetKeyRetention

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) GetKeyRetention() int32`

GetKeyRetention returns the KeyRetention field if non-nil, zero value otherwise.

### GetKeyRetentionOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) GetKeyRetentionOk() (*int32, bool)`

GetKeyRetentionOk returns a tuple with the KeyRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyRetention

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) SetKeyRetention(v int32)`

SetKeyRetention sets KeyRetention field to given value.

### HasKeyRetention

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) HasKeyRetention() bool`

HasKeyRetention returns a boolean if a field has been set.

### GetKeyRetentionDays

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) GetKeyRetentionDays() int32`

GetKeyRetentionDays returns the KeyRetentionDays field if non-nil, zero value otherwise.

### GetKeyRetentionDaysOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) GetKeyRetentionDaysOk() (*int32, bool)`

GetKeyRetentionDaysOk returns a tuple with the KeyRetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyRetentionDays

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) SetKeyRetentionDays(v int32)`

SetKeyRetentionDays sets KeyRetentionDays field to given value.

### HasKeyRetentionDays

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) HasKeyRetentionDays() bool`

HasKeyRetentionDays returns a boolean if a field has been set.

### GetExplicitCredentials

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) GetExplicitCredentials() bool`

GetExplicitCredentials returns the ExplicitCredentials field if non-nil, zero value otherwise.

### GetExplicitCredentialsOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) GetExplicitCredentialsOk() (*bool, bool)`

GetExplicitCredentialsOk returns a tuple with the ExplicitCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicitCredentials

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) SetExplicitCredentials(v bool)`

SetExplicitCredentials sets ExplicitCredentials field to given value.

### HasExplicitCredentials

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) HasExplicitCredentials() bool`

HasExplicitCredentials returns a boolean if a field has been set.

### GetSubscriberTerms

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) GetSubscriberTerms() bool`

GetSubscriberTerms returns the SubscriberTerms field if non-nil, zero value otherwise.

### GetSubscriberTermsOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) GetSubscriberTermsOk() (*bool, bool)`

GetSubscriberTermsOk returns a tuple with the SubscriberTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriberTerms

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) SetSubscriberTerms(v bool)`

SetSubscriberTerms sets SubscriberTerms field to given value.

### HasSubscriberTerms

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) HasSubscriberTerms() bool`

HasSubscriberTerms returns a boolean if a field has been set.

### GetExplicitUser

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) GetExplicitUser() string`

GetExplicitUser returns the ExplicitUser field if non-nil, zero value otherwise.

### GetExplicitUserOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) GetExplicitUserOk() (*string, bool)`

GetExplicitUserOk returns a tuple with the ExplicitUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicitUser

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) SetExplicitUser(v string)`

SetExplicitUser sets ExplicitUser field to given value.

### HasExplicitUser

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) HasExplicitUser() bool`

HasExplicitUser returns a boolean if a field has been set.

### GetExplicitPassword

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) GetExplicitPassword() ModelsKeyfactorAPISecret`

GetExplicitPassword returns the ExplicitPassword field if non-nil, zero value otherwise.

### GetExplicitPasswordOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) GetExplicitPasswordOk() (*ModelsKeyfactorAPISecret, bool)`

GetExplicitPasswordOk returns a tuple with the ExplicitPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicitPassword

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) SetExplicitPassword(v ModelsKeyfactorAPISecret)`

SetExplicitPassword sets ExplicitPassword field to given value.

### HasExplicitPassword

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) HasExplicitPassword() bool`

HasExplicitPassword returns a boolean if a field has been set.

### GetUseAllowedRequesters

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) GetUseAllowedRequesters() bool`

GetUseAllowedRequesters returns the UseAllowedRequesters field if non-nil, zero value otherwise.

### GetUseAllowedRequestersOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) GetUseAllowedRequestersOk() (*bool, bool)`

GetUseAllowedRequestersOk returns a tuple with the UseAllowedRequesters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAllowedRequesters

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) SetUseAllowedRequesters(v bool)`

SetUseAllowedRequesters sets UseAllowedRequesters field to given value.

### HasUseAllowedRequesters

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) HasUseAllowedRequesters() bool`

HasUseAllowedRequesters returns a boolean if a field has been set.

### GetAllowedRequesters

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) GetAllowedRequesters() []string`

GetAllowedRequesters returns the AllowedRequesters field if non-nil, zero value otherwise.

### GetAllowedRequestersOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) GetAllowedRequestersOk() (*[]string, bool)`

GetAllowedRequestersOk returns a tuple with the AllowedRequesters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedRequesters

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) SetAllowedRequesters(v []string)`

SetAllowedRequesters sets AllowedRequesters field to given value.

### HasAllowedRequesters

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) HasAllowedRequesters() bool`

HasAllowedRequesters returns a boolean if a field has been set.

### GetFullScan

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) GetFullScan() KeyfactorCommonSchedulingKeyfactorSchedule`

GetFullScan returns the FullScan field if non-nil, zero value otherwise.

### GetFullScanOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) GetFullScanOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetFullScanOk returns a tuple with the FullScan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullScan

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) SetFullScan(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetFullScan sets FullScan field to given value.

### HasFullScan

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) HasFullScan() bool`

HasFullScan returns a boolean if a field has been set.

### GetIncrementalScan

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) GetIncrementalScan() KeyfactorCommonSchedulingKeyfactorSchedule`

GetIncrementalScan returns the IncrementalScan field if non-nil, zero value otherwise.

### GetIncrementalScanOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) GetIncrementalScanOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetIncrementalScanOk returns a tuple with the IncrementalScan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncrementalScan

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) SetIncrementalScan(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetIncrementalScan sets IncrementalScan field to given value.

### HasIncrementalScan

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) HasIncrementalScan() bool`

HasIncrementalScan returns a boolean if a field has been set.

### GetThresholdCheck

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) GetThresholdCheck() KeyfactorCommonSchedulingKeyfactorSchedule`

GetThresholdCheck returns the ThresholdCheck field if non-nil, zero value otherwise.

### GetThresholdCheckOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) GetThresholdCheckOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetThresholdCheckOk returns a tuple with the ThresholdCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdCheck

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) SetThresholdCheck(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetThresholdCheck sets ThresholdCheck field to given value.

### HasThresholdCheck

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) HasThresholdCheck() bool`

HasThresholdCheck returns a boolean if a field has been set.

### GetAuthCertificatePassword

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) GetAuthCertificatePassword() ModelsKeyfactorAPISecret`

GetAuthCertificatePassword returns the AuthCertificatePassword field if non-nil, zero value otherwise.

### GetAuthCertificatePasswordOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) GetAuthCertificatePasswordOk() (*ModelsKeyfactorAPISecret, bool)`

GetAuthCertificatePasswordOk returns a tuple with the AuthCertificatePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCertificatePassword

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) SetAuthCertificatePassword(v ModelsKeyfactorAPISecret)`

SetAuthCertificatePassword sets AuthCertificatePassword field to given value.

### HasAuthCertificatePassword

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) HasAuthCertificatePassword() bool`

HasAuthCertificatePassword returns a boolean if a field has been set.

### GetAuthCertificate

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) GetAuthCertificate() ModelsKeyfactorAPISecret`

GetAuthCertificate returns the AuthCertificate field if non-nil, zero value otherwise.

### GetAuthCertificateOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) GetAuthCertificateOk() (*ModelsKeyfactorAPISecret, bool)`

GetAuthCertificateOk returns a tuple with the AuthCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCertificate

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) SetAuthCertificate(v ModelsKeyfactorAPISecret)`

SetAuthCertificate sets AuthCertificate field to given value.

### HasAuthCertificate

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) HasAuthCertificate() bool`

HasAuthCertificate returns a boolean if a field has been set.

### GetCAType

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) GetCAType() int32`

GetCAType returns the CAType field if non-nil, zero value otherwise.

### GetCATypeOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) GetCATypeOk() (*int32, bool)`

GetCATypeOk returns a tuple with the CAType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCAType

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) SetCAType(v int32)`

SetCAType sets CAType field to given value.

### HasCAType

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) HasCAType() bool`

HasCAType returns a boolean if a field has been set.

### GetEnforceUniqueDN

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) GetEnforceUniqueDN() bool`

GetEnforceUniqueDN returns the EnforceUniqueDN field if non-nil, zero value otherwise.

### GetEnforceUniqueDNOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) GetEnforceUniqueDNOk() (*bool, bool)`

GetEnforceUniqueDNOk returns a tuple with the EnforceUniqueDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceUniqueDN

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) SetEnforceUniqueDN(v bool)`

SetEnforceUniqueDN sets EnforceUniqueDN field to given value.

### HasEnforceUniqueDN

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityRequest) HasEnforceUniqueDN() bool`

HasEnforceUniqueDN returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


