# ModelsCertificateAuthoritiesCertificateAuthorityResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**LogicalName** | Pointer to **string** |  | [optional] 
**HostName** | Pointer to **string** |  | [optional] 
**Delegate** | Pointer to **bool** |  | [optional] 
**DelegateEnrollment** | Pointer to **bool** |  | [optional] 
**ForestRoot** | Pointer to **string** |  | [optional] [readonly] 
**ConfigurationTenant** | Pointer to **string** |  | [optional] 
**Remote** | Pointer to **bool** |  | [optional] 
**Agent** | Pointer to **string** |  | [optional] 
**Standalone** | Pointer to **bool** |  | [optional] 
**MonitorThresholds** | Pointer to **bool** |  | [optional] 
**IssuanceMax** | Pointer to **int32** |  | [optional] 
**IssuanceMin** | Pointer to **int32** |  | [optional] 
**DenialMax** | Pointer to **int32** |  | [optional] 
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
**CAType** | Pointer to **int32** |  | [optional] 
**AuthCertificate** | Pointer to [**ModelsCertificateAuthoritiesCertificateAuthorityAuthCertificate**](ModelsCertificateAuthoritiesCertificateAuthorityAuthCertificate.md) |  | [optional] 
**EnforceUniqueDN** | Pointer to **bool** |  | [optional] 
**LastScan** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewModelsCertificateAuthoritiesCertificateAuthorityResponse

`func NewModelsCertificateAuthoritiesCertificateAuthorityResponse() *ModelsCertificateAuthoritiesCertificateAuthorityResponse`

NewModelsCertificateAuthoritiesCertificateAuthorityResponse instantiates a new ModelsCertificateAuthoritiesCertificateAuthorityResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsCertificateAuthoritiesCertificateAuthorityResponseWithDefaults

`func NewModelsCertificateAuthoritiesCertificateAuthorityResponseWithDefaults() *ModelsCertificateAuthoritiesCertificateAuthorityResponse`

NewModelsCertificateAuthoritiesCertificateAuthorityResponseWithDefaults instantiates a new ModelsCertificateAuthoritiesCertificateAuthorityResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLogicalName

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) GetLogicalName() string`

GetLogicalName returns the LogicalName field if non-nil, zero value otherwise.

### GetLogicalNameOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) GetLogicalNameOk() (*string, bool)`

GetLogicalNameOk returns a tuple with the LogicalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalName

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) SetLogicalName(v string)`

SetLogicalName sets LogicalName field to given value.

### HasLogicalName

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) HasLogicalName() bool`

HasLogicalName returns a boolean if a field has been set.

### GetHostName

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetDelegate

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) GetDelegate() bool`

GetDelegate returns the Delegate field if non-nil, zero value otherwise.

### GetDelegateOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) GetDelegateOk() (*bool, bool)`

GetDelegateOk returns a tuple with the Delegate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegate

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) SetDelegate(v bool)`

SetDelegate sets Delegate field to given value.

### HasDelegate

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) HasDelegate() bool`

HasDelegate returns a boolean if a field has been set.

### GetDelegateEnrollment

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) GetDelegateEnrollment() bool`

GetDelegateEnrollment returns the DelegateEnrollment field if non-nil, zero value otherwise.

### GetDelegateEnrollmentOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) GetDelegateEnrollmentOk() (*bool, bool)`

GetDelegateEnrollmentOk returns a tuple with the DelegateEnrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegateEnrollment

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) SetDelegateEnrollment(v bool)`

SetDelegateEnrollment sets DelegateEnrollment field to given value.

### HasDelegateEnrollment

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) HasDelegateEnrollment() bool`

HasDelegateEnrollment returns a boolean if a field has been set.

### GetForestRoot

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) GetForestRoot() string`

GetForestRoot returns the ForestRoot field if non-nil, zero value otherwise.

### GetForestRootOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) GetForestRootOk() (*string, bool)`

GetForestRootOk returns a tuple with the ForestRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForestRoot

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) SetForestRoot(v string)`

SetForestRoot sets ForestRoot field to given value.

### HasForestRoot

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) HasForestRoot() bool`

HasForestRoot returns a boolean if a field has been set.

### GetConfigurationTenant

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) GetConfigurationTenant() string`

GetConfigurationTenant returns the ConfigurationTenant field if non-nil, zero value otherwise.

### GetConfigurationTenantOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) GetConfigurationTenantOk() (*string, bool)`

GetConfigurationTenantOk returns a tuple with the ConfigurationTenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationTenant

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) SetConfigurationTenant(v string)`

SetConfigurationTenant sets ConfigurationTenant field to given value.

### HasConfigurationTenant

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) HasConfigurationTenant() bool`

HasConfigurationTenant returns a boolean if a field has been set.

### GetRemote

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) GetRemote() bool`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) GetRemoteOk() (*bool, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) SetRemote(v bool)`

SetRemote sets Remote field to given value.

### HasRemote

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) HasRemote() bool`

HasRemote returns a boolean if a field has been set.

### GetAgent

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) GetAgent() string`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) GetAgentOk() (*string, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) SetAgent(v string)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) HasAgent() bool`

HasAgent returns a boolean if a field has been set.

### GetStandalone

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) GetStandalone() bool`

GetStandalone returns the Standalone field if non-nil, zero value otherwise.

### GetStandaloneOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) GetStandaloneOk() (*bool, bool)`

GetStandaloneOk returns a tuple with the Standalone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandalone

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) SetStandalone(v bool)`

SetStandalone sets Standalone field to given value.

### HasStandalone

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) HasStandalone() bool`

HasStandalone returns a boolean if a field has been set.

### GetMonitorThresholds

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) GetMonitorThresholds() bool`

GetMonitorThresholds returns the MonitorThresholds field if non-nil, zero value otherwise.

### GetMonitorThresholdsOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) GetMonitorThresholdsOk() (*bool, bool)`

GetMonitorThresholdsOk returns a tuple with the MonitorThresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorThresholds

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) SetMonitorThresholds(v bool)`

SetMonitorThresholds sets MonitorThresholds field to given value.

### HasMonitorThresholds

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) HasMonitorThresholds() bool`

HasMonitorThresholds returns a boolean if a field has been set.

### GetIssuanceMax

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) GetIssuanceMax() int32`

GetIssuanceMax returns the IssuanceMax field if non-nil, zero value otherwise.

### GetIssuanceMaxOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) GetIssuanceMaxOk() (*int32, bool)`

GetIssuanceMaxOk returns a tuple with the IssuanceMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuanceMax

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) SetIssuanceMax(v int32)`

SetIssuanceMax sets IssuanceMax field to given value.

### HasIssuanceMax

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) HasIssuanceMax() bool`

HasIssuanceMax returns a boolean if a field has been set.

### GetIssuanceMin

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) GetIssuanceMin() int32`

GetIssuanceMin returns the IssuanceMin field if non-nil, zero value otherwise.

### GetIssuanceMinOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) GetIssuanceMinOk() (*int32, bool)`

GetIssuanceMinOk returns a tuple with the IssuanceMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuanceMin

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) SetIssuanceMin(v int32)`

SetIssuanceMin sets IssuanceMin field to given value.

### HasIssuanceMin

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) HasIssuanceMin() bool`

HasIssuanceMin returns a boolean if a field has been set.

### GetDenialMax

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) GetDenialMax() int32`

GetDenialMax returns the DenialMax field if non-nil, zero value otherwise.

### GetDenialMaxOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) GetDenialMaxOk() (*int32, bool)`

GetDenialMaxOk returns a tuple with the DenialMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenialMax

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) SetDenialMax(v int32)`

SetDenialMax sets DenialMax field to given value.

### HasDenialMax

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) HasDenialMax() bool`

HasDenialMax returns a boolean if a field has been set.

### GetFailureMax

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) GetFailureMax() int32`

GetFailureMax returns the FailureMax field if non-nil, zero value otherwise.

### GetFailureMaxOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) GetFailureMaxOk() (*int32, bool)`

GetFailureMaxOk returns a tuple with the FailureMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureMax

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) SetFailureMax(v int32)`

SetFailureMax sets FailureMax field to given value.

### HasFailureMax

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) HasFailureMax() bool`

HasFailureMax returns a boolean if a field has been set.

### GetRFCEnforcement

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) GetRFCEnforcement() bool`

GetRFCEnforcement returns the RFCEnforcement field if non-nil, zero value otherwise.

### GetRFCEnforcementOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) GetRFCEnforcementOk() (*bool, bool)`

GetRFCEnforcementOk returns a tuple with the RFCEnforcement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRFCEnforcement

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) SetRFCEnforcement(v bool)`

SetRFCEnforcement sets RFCEnforcement field to given value.

### HasRFCEnforcement

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) HasRFCEnforcement() bool`

HasRFCEnforcement returns a boolean if a field has been set.

### GetProperties

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) GetProperties() string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) GetPropertiesOk() (*string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) SetProperties(v string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetAllowedEnrollmentTypes

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) GetAllowedEnrollmentTypes() int32`

GetAllowedEnrollmentTypes returns the AllowedEnrollmentTypes field if non-nil, zero value otherwise.

### GetAllowedEnrollmentTypesOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) GetAllowedEnrollmentTypesOk() (*int32, bool)`

GetAllowedEnrollmentTypesOk returns a tuple with the AllowedEnrollmentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedEnrollmentTypes

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) SetAllowedEnrollmentTypes(v int32)`

SetAllowedEnrollmentTypes sets AllowedEnrollmentTypes field to given value.

### HasAllowedEnrollmentTypes

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) HasAllowedEnrollmentTypes() bool`

HasAllowedEnrollmentTypes returns a boolean if a field has been set.

### GetKeyRetention

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) GetKeyRetention() int32`

GetKeyRetention returns the KeyRetention field if non-nil, zero value otherwise.

### GetKeyRetentionOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) GetKeyRetentionOk() (*int32, bool)`

GetKeyRetentionOk returns a tuple with the KeyRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyRetention

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) SetKeyRetention(v int32)`

SetKeyRetention sets KeyRetention field to given value.

### HasKeyRetention

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) HasKeyRetention() bool`

HasKeyRetention returns a boolean if a field has been set.

### GetKeyRetentionDays

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) GetKeyRetentionDays() int32`

GetKeyRetentionDays returns the KeyRetentionDays field if non-nil, zero value otherwise.

### GetKeyRetentionDaysOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) GetKeyRetentionDaysOk() (*int32, bool)`

GetKeyRetentionDaysOk returns a tuple with the KeyRetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyRetentionDays

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) SetKeyRetentionDays(v int32)`

SetKeyRetentionDays sets KeyRetentionDays field to given value.

### HasKeyRetentionDays

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) HasKeyRetentionDays() bool`

HasKeyRetentionDays returns a boolean if a field has been set.

### GetExplicitCredentials

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) GetExplicitCredentials() bool`

GetExplicitCredentials returns the ExplicitCredentials field if non-nil, zero value otherwise.

### GetExplicitCredentialsOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) GetExplicitCredentialsOk() (*bool, bool)`

GetExplicitCredentialsOk returns a tuple with the ExplicitCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicitCredentials

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) SetExplicitCredentials(v bool)`

SetExplicitCredentials sets ExplicitCredentials field to given value.

### HasExplicitCredentials

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) HasExplicitCredentials() bool`

HasExplicitCredentials returns a boolean if a field has been set.

### GetSubscriberTerms

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) GetSubscriberTerms() bool`

GetSubscriberTerms returns the SubscriberTerms field if non-nil, zero value otherwise.

### GetSubscriberTermsOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) GetSubscriberTermsOk() (*bool, bool)`

GetSubscriberTermsOk returns a tuple with the SubscriberTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriberTerms

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) SetSubscriberTerms(v bool)`

SetSubscriberTerms sets SubscriberTerms field to given value.

### HasSubscriberTerms

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) HasSubscriberTerms() bool`

HasSubscriberTerms returns a boolean if a field has been set.

### GetExplicitUser

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) GetExplicitUser() string`

GetExplicitUser returns the ExplicitUser field if non-nil, zero value otherwise.

### GetExplicitUserOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) GetExplicitUserOk() (*string, bool)`

GetExplicitUserOk returns a tuple with the ExplicitUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicitUser

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) SetExplicitUser(v string)`

SetExplicitUser sets ExplicitUser field to given value.

### HasExplicitUser

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) HasExplicitUser() bool`

HasExplicitUser returns a boolean if a field has been set.

### GetExplicitPassword

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) GetExplicitPassword() ModelsKeyfactorAPISecret`

GetExplicitPassword returns the ExplicitPassword field if non-nil, zero value otherwise.

### GetExplicitPasswordOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) GetExplicitPasswordOk() (*ModelsKeyfactorAPISecret, bool)`

GetExplicitPasswordOk returns a tuple with the ExplicitPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicitPassword

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) SetExplicitPassword(v ModelsKeyfactorAPISecret)`

SetExplicitPassword sets ExplicitPassword field to given value.

### HasExplicitPassword

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) HasExplicitPassword() bool`

HasExplicitPassword returns a boolean if a field has been set.

### GetUseAllowedRequesters

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) GetUseAllowedRequesters() bool`

GetUseAllowedRequesters returns the UseAllowedRequesters field if non-nil, zero value otherwise.

### GetUseAllowedRequestersOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) GetUseAllowedRequestersOk() (*bool, bool)`

GetUseAllowedRequestersOk returns a tuple with the UseAllowedRequesters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAllowedRequesters

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) SetUseAllowedRequesters(v bool)`

SetUseAllowedRequesters sets UseAllowedRequesters field to given value.

### HasUseAllowedRequesters

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) HasUseAllowedRequesters() bool`

HasUseAllowedRequesters returns a boolean if a field has been set.

### GetAllowedRequesters

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) GetAllowedRequesters() []string`

GetAllowedRequesters returns the AllowedRequesters field if non-nil, zero value otherwise.

### GetAllowedRequestersOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) GetAllowedRequestersOk() (*[]string, bool)`

GetAllowedRequestersOk returns a tuple with the AllowedRequesters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedRequesters

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) SetAllowedRequesters(v []string)`

SetAllowedRequesters sets AllowedRequesters field to given value.

### HasAllowedRequesters

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) HasAllowedRequesters() bool`

HasAllowedRequesters returns a boolean if a field has been set.

### GetFullScan

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) GetFullScan() KeyfactorCommonSchedulingKeyfactorSchedule`

GetFullScan returns the FullScan field if non-nil, zero value otherwise.

### GetFullScanOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) GetFullScanOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetFullScanOk returns a tuple with the FullScan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullScan

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) SetFullScan(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetFullScan sets FullScan field to given value.

### HasFullScan

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) HasFullScan() bool`

HasFullScan returns a boolean if a field has been set.

### GetIncrementalScan

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) GetIncrementalScan() KeyfactorCommonSchedulingKeyfactorSchedule`

GetIncrementalScan returns the IncrementalScan field if non-nil, zero value otherwise.

### GetIncrementalScanOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) GetIncrementalScanOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetIncrementalScanOk returns a tuple with the IncrementalScan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncrementalScan

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) SetIncrementalScan(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetIncrementalScan sets IncrementalScan field to given value.

### HasIncrementalScan

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) HasIncrementalScan() bool`

HasIncrementalScan returns a boolean if a field has been set.

### GetThresholdCheck

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) GetThresholdCheck() KeyfactorCommonSchedulingKeyfactorSchedule`

GetThresholdCheck returns the ThresholdCheck field if non-nil, zero value otherwise.

### GetThresholdCheckOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) GetThresholdCheckOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetThresholdCheckOk returns a tuple with the ThresholdCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdCheck

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) SetThresholdCheck(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetThresholdCheck sets ThresholdCheck field to given value.

### HasThresholdCheck

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) HasThresholdCheck() bool`

HasThresholdCheck returns a boolean if a field has been set.

### GetCAType

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) GetCAType() int32`

GetCAType returns the CAType field if non-nil, zero value otherwise.

### GetCATypeOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) GetCATypeOk() (*int32, bool)`

GetCATypeOk returns a tuple with the CAType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCAType

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) SetCAType(v int32)`

SetCAType sets CAType field to given value.

### HasCAType

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) HasCAType() bool`

HasCAType returns a boolean if a field has been set.

### GetAuthCertificate

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) GetAuthCertificate() ModelsCertificateAuthoritiesCertificateAuthorityAuthCertificate`

GetAuthCertificate returns the AuthCertificate field if non-nil, zero value otherwise.

### GetAuthCertificateOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) GetAuthCertificateOk() (*ModelsCertificateAuthoritiesCertificateAuthorityAuthCertificate, bool)`

GetAuthCertificateOk returns a tuple with the AuthCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCertificate

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) SetAuthCertificate(v ModelsCertificateAuthoritiesCertificateAuthorityAuthCertificate)`

SetAuthCertificate sets AuthCertificate field to given value.

### HasAuthCertificate

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) HasAuthCertificate() bool`

HasAuthCertificate returns a boolean if a field has been set.

### GetEnforceUniqueDN

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) GetEnforceUniqueDN() bool`

GetEnforceUniqueDN returns the EnforceUniqueDN field if non-nil, zero value otherwise.

### GetEnforceUniqueDNOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) GetEnforceUniqueDNOk() (*bool, bool)`

GetEnforceUniqueDNOk returns a tuple with the EnforceUniqueDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceUniqueDN

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) SetEnforceUniqueDN(v bool)`

SetEnforceUniqueDN sets EnforceUniqueDN field to given value.

### HasEnforceUniqueDN

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) HasEnforceUniqueDN() bool`

HasEnforceUniqueDN returns a boolean if a field has been set.

### GetLastScan

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) GetLastScan() string`

GetLastScan returns the LastScan field if non-nil, zero value otherwise.

### GetLastScanOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) GetLastScanOk() (*string, bool)`

GetLastScanOk returns a tuple with the LastScan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastScan

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) SetLastScan(v string)`

SetLastScan sets LastScan field to given value.

### HasLastScan

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityResponse) HasLastScan() bool`

HasLastScan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


