# KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**LogicalName** | Pointer to **NullableString** |  | [optional] 
**HostName** | Pointer to **NullableString** |  | [optional] 
**Delegate** | Pointer to **bool** |  | [optional] 
**DelegateEnrollment** | Pointer to **bool** |  | [optional] 
**ForestRoot** | Pointer to **NullableString** |  | [optional] 
**ConfigurationTenant** | Pointer to **NullableString** |  | [optional] 
**Remote** | Pointer to **bool** |  | [optional] 
**Agent** | Pointer to **NullableString** |  | [optional] 
**Standalone** | Pointer to **bool** |  | [optional] 
**MonitorThresholds** | Pointer to **bool** |  | [optional] 
**IssuanceMax** | Pointer to **NullableInt32** |  | [optional] 
**IssuanceMin** | Pointer to **NullableInt32** |  | [optional] 
**FailureMax** | Pointer to **NullableInt32** |  | [optional] 
**RfcEnforcement** | Pointer to **bool** |  | [optional] 
**Properties** | Pointer to **NullableString** |  | [optional] 
**AllowedEnrollmentTypes** | Pointer to **int32** |  | [optional] 
**KeyRetention** | Pointer to **int32** |  | [optional] 
**KeyRetentionDays** | Pointer to **NullableInt32** |  | [optional] 
**ExplicitCredentials** | Pointer to **bool** |  | [optional] 
**SubscriberTerms** | Pointer to **bool** |  | [optional] 
**ExplicitUser** | Pointer to **NullableString** |  | [optional] 
**ExplicitPassword** | Pointer to [**ModelsKeyfactorAPISecret**](ModelsKeyfactorAPISecret.md) |  | [optional] 
**UseAllowedRequesters** | Pointer to **bool** |  | [optional] 
**AllowedRequesters** | Pointer to **[]string** |  | [optional] 
**FullScan** | Pointer to [**KeyfactorCommonSchedulingKeyfactorSchedule**](KeyfactorCommonSchedulingKeyfactorSchedule.md) |  | [optional] 
**IncrementalScan** | Pointer to [**KeyfactorCommonSchedulingKeyfactorSchedule**](KeyfactorCommonSchedulingKeyfactorSchedule.md) |  | [optional] 
**ThresholdCheck** | Pointer to [**KeyfactorCommonSchedulingKeyfactorSchedule**](KeyfactorCommonSchedulingKeyfactorSchedule.md) |  | [optional] 
**AuthCertificatePassword** | Pointer to [**ModelsKeyfactorAPISecret**](ModelsKeyfactorAPISecret.md) |  | [optional] 
**AuthCertificate** | Pointer to [**ModelsKeyfactorAPISecret**](ModelsKeyfactorAPISecret.md) |  | [optional] 
**CaType** | Pointer to **int32** |  | [optional] 
**EnforceUniqueDN** | Pointer to **bool** |  | [optional] 
**AllowOneClickRenewals** | Pointer to **bool** |  | [optional] 
**NewEndEntityOnRenewAndReissue** | Pointer to **bool** |  | [optional] 

## Methods

### NewKeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest

`func NewKeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest() *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest`

NewKeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest instantiates a new KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequestWithDefaults

`func NewKeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequestWithDefaults() *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest`

NewKeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequestWithDefaults instantiates a new KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLogicalName

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) GetLogicalName() string`

GetLogicalName returns the LogicalName field if non-nil, zero value otherwise.

### GetLogicalNameOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) GetLogicalNameOk() (*string, bool)`

GetLogicalNameOk returns a tuple with the LogicalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalName

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) SetLogicalName(v string)`

SetLogicalName sets LogicalName field to given value.

### HasLogicalName

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) HasLogicalName() bool`

HasLogicalName returns a boolean if a field has been set.

### SetLogicalNameNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) SetLogicalNameNil(b bool)`

 SetLogicalNameNil sets the value for LogicalName to be an explicit nil

### UnsetLogicalName
`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) UnsetLogicalName()`

UnsetLogicalName ensures that no value is present for LogicalName, not even an explicit nil
### GetHostName

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### SetHostNameNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) SetHostNameNil(b bool)`

 SetHostNameNil sets the value for HostName to be an explicit nil

### UnsetHostName
`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) UnsetHostName()`

UnsetHostName ensures that no value is present for HostName, not even an explicit nil
### GetDelegate

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) GetDelegate() bool`

GetDelegate returns the Delegate field if non-nil, zero value otherwise.

### GetDelegateOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) GetDelegateOk() (*bool, bool)`

GetDelegateOk returns a tuple with the Delegate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegate

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) SetDelegate(v bool)`

SetDelegate sets Delegate field to given value.

### HasDelegate

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) HasDelegate() bool`

HasDelegate returns a boolean if a field has been set.

### GetDelegateEnrollment

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) GetDelegateEnrollment() bool`

GetDelegateEnrollment returns the DelegateEnrollment field if non-nil, zero value otherwise.

### GetDelegateEnrollmentOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) GetDelegateEnrollmentOk() (*bool, bool)`

GetDelegateEnrollmentOk returns a tuple with the DelegateEnrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegateEnrollment

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) SetDelegateEnrollment(v bool)`

SetDelegateEnrollment sets DelegateEnrollment field to given value.

### HasDelegateEnrollment

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) HasDelegateEnrollment() bool`

HasDelegateEnrollment returns a boolean if a field has been set.

### GetForestRoot

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) GetForestRoot() string`

GetForestRoot returns the ForestRoot field if non-nil, zero value otherwise.

### GetForestRootOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) GetForestRootOk() (*string, bool)`

GetForestRootOk returns a tuple with the ForestRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForestRoot

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) SetForestRoot(v string)`

SetForestRoot sets ForestRoot field to given value.

### HasForestRoot

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) HasForestRoot() bool`

HasForestRoot returns a boolean if a field has been set.

### SetForestRootNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) SetForestRootNil(b bool)`

 SetForestRootNil sets the value for ForestRoot to be an explicit nil

### UnsetForestRoot
`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) UnsetForestRoot()`

UnsetForestRoot ensures that no value is present for ForestRoot, not even an explicit nil
### GetConfigurationTenant

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) GetConfigurationTenant() string`

GetConfigurationTenant returns the ConfigurationTenant field if non-nil, zero value otherwise.

### GetConfigurationTenantOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) GetConfigurationTenantOk() (*string, bool)`

GetConfigurationTenantOk returns a tuple with the ConfigurationTenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationTenant

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) SetConfigurationTenant(v string)`

SetConfigurationTenant sets ConfigurationTenant field to given value.

### HasConfigurationTenant

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) HasConfigurationTenant() bool`

HasConfigurationTenant returns a boolean if a field has been set.

### SetConfigurationTenantNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) SetConfigurationTenantNil(b bool)`

 SetConfigurationTenantNil sets the value for ConfigurationTenant to be an explicit nil

### UnsetConfigurationTenant
`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) UnsetConfigurationTenant()`

UnsetConfigurationTenant ensures that no value is present for ConfigurationTenant, not even an explicit nil
### GetRemote

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) GetRemote() bool`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) GetRemoteOk() (*bool, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) SetRemote(v bool)`

SetRemote sets Remote field to given value.

### HasRemote

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) HasRemote() bool`

HasRemote returns a boolean if a field has been set.

### GetAgent

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) GetAgent() string`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) GetAgentOk() (*string, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) SetAgent(v string)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) HasAgent() bool`

HasAgent returns a boolean if a field has been set.

### SetAgentNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) SetAgentNil(b bool)`

 SetAgentNil sets the value for Agent to be an explicit nil

### UnsetAgent
`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) UnsetAgent()`

UnsetAgent ensures that no value is present for Agent, not even an explicit nil
### GetStandalone

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) GetStandalone() bool`

GetStandalone returns the Standalone field if non-nil, zero value otherwise.

### GetStandaloneOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) GetStandaloneOk() (*bool, bool)`

GetStandaloneOk returns a tuple with the Standalone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandalone

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) SetStandalone(v bool)`

SetStandalone sets Standalone field to given value.

### HasStandalone

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) HasStandalone() bool`

HasStandalone returns a boolean if a field has been set.

### GetMonitorThresholds

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) GetMonitorThresholds() bool`

GetMonitorThresholds returns the MonitorThresholds field if non-nil, zero value otherwise.

### GetMonitorThresholdsOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) GetMonitorThresholdsOk() (*bool, bool)`

GetMonitorThresholdsOk returns a tuple with the MonitorThresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorThresholds

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) SetMonitorThresholds(v bool)`

SetMonitorThresholds sets MonitorThresholds field to given value.

### HasMonitorThresholds

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) HasMonitorThresholds() bool`

HasMonitorThresholds returns a boolean if a field has been set.

### GetIssuanceMax

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) GetIssuanceMax() int32`

GetIssuanceMax returns the IssuanceMax field if non-nil, zero value otherwise.

### GetIssuanceMaxOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) GetIssuanceMaxOk() (*int32, bool)`

GetIssuanceMaxOk returns a tuple with the IssuanceMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuanceMax

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) SetIssuanceMax(v int32)`

SetIssuanceMax sets IssuanceMax field to given value.

### HasIssuanceMax

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) HasIssuanceMax() bool`

HasIssuanceMax returns a boolean if a field has been set.

### SetIssuanceMaxNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) SetIssuanceMaxNil(b bool)`

 SetIssuanceMaxNil sets the value for IssuanceMax to be an explicit nil

### UnsetIssuanceMax
`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) UnsetIssuanceMax()`

UnsetIssuanceMax ensures that no value is present for IssuanceMax, not even an explicit nil
### GetIssuanceMin

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) GetIssuanceMin() int32`

GetIssuanceMin returns the IssuanceMin field if non-nil, zero value otherwise.

### GetIssuanceMinOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) GetIssuanceMinOk() (*int32, bool)`

GetIssuanceMinOk returns a tuple with the IssuanceMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuanceMin

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) SetIssuanceMin(v int32)`

SetIssuanceMin sets IssuanceMin field to given value.

### HasIssuanceMin

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) HasIssuanceMin() bool`

HasIssuanceMin returns a boolean if a field has been set.

### SetIssuanceMinNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) SetIssuanceMinNil(b bool)`

 SetIssuanceMinNil sets the value for IssuanceMin to be an explicit nil

### UnsetIssuanceMin
`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) UnsetIssuanceMin()`

UnsetIssuanceMin ensures that no value is present for IssuanceMin, not even an explicit nil
### GetFailureMax

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) GetFailureMax() int32`

GetFailureMax returns the FailureMax field if non-nil, zero value otherwise.

### GetFailureMaxOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) GetFailureMaxOk() (*int32, bool)`

GetFailureMaxOk returns a tuple with the FailureMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureMax

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) SetFailureMax(v int32)`

SetFailureMax sets FailureMax field to given value.

### HasFailureMax

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) HasFailureMax() bool`

HasFailureMax returns a boolean if a field has been set.

### SetFailureMaxNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) SetFailureMaxNil(b bool)`

 SetFailureMaxNil sets the value for FailureMax to be an explicit nil

### UnsetFailureMax
`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) UnsetFailureMax()`

UnsetFailureMax ensures that no value is present for FailureMax, not even an explicit nil
### GetRfcEnforcement

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) GetRfcEnforcement() bool`

GetRfcEnforcement returns the RfcEnforcement field if non-nil, zero value otherwise.

### GetRfcEnforcementOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) GetRfcEnforcementOk() (*bool, bool)`

GetRfcEnforcementOk returns a tuple with the RfcEnforcement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRfcEnforcement

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) SetRfcEnforcement(v bool)`

SetRfcEnforcement sets RfcEnforcement field to given value.

### HasRfcEnforcement

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) HasRfcEnforcement() bool`

HasRfcEnforcement returns a boolean if a field has been set.

### GetProperties

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) GetProperties() string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) GetPropertiesOk() (*string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) SetProperties(v string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetAllowedEnrollmentTypes

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) GetAllowedEnrollmentTypes() int32`

GetAllowedEnrollmentTypes returns the AllowedEnrollmentTypes field if non-nil, zero value otherwise.

### GetAllowedEnrollmentTypesOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) GetAllowedEnrollmentTypesOk() (*int32, bool)`

GetAllowedEnrollmentTypesOk returns a tuple with the AllowedEnrollmentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedEnrollmentTypes

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) SetAllowedEnrollmentTypes(v int32)`

SetAllowedEnrollmentTypes sets AllowedEnrollmentTypes field to given value.

### HasAllowedEnrollmentTypes

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) HasAllowedEnrollmentTypes() bool`

HasAllowedEnrollmentTypes returns a boolean if a field has been set.

### GetKeyRetention

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) GetKeyRetention() int32`

GetKeyRetention returns the KeyRetention field if non-nil, zero value otherwise.

### GetKeyRetentionOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) GetKeyRetentionOk() (*int32, bool)`

GetKeyRetentionOk returns a tuple with the KeyRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyRetention

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) SetKeyRetention(v int32)`

SetKeyRetention sets KeyRetention field to given value.

### HasKeyRetention

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) HasKeyRetention() bool`

HasKeyRetention returns a boolean if a field has been set.

### GetKeyRetentionDays

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) GetKeyRetentionDays() int32`

GetKeyRetentionDays returns the KeyRetentionDays field if non-nil, zero value otherwise.

### GetKeyRetentionDaysOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) GetKeyRetentionDaysOk() (*int32, bool)`

GetKeyRetentionDaysOk returns a tuple with the KeyRetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyRetentionDays

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) SetKeyRetentionDays(v int32)`

SetKeyRetentionDays sets KeyRetentionDays field to given value.

### HasKeyRetentionDays

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) HasKeyRetentionDays() bool`

HasKeyRetentionDays returns a boolean if a field has been set.

### SetKeyRetentionDaysNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) SetKeyRetentionDaysNil(b bool)`

 SetKeyRetentionDaysNil sets the value for KeyRetentionDays to be an explicit nil

### UnsetKeyRetentionDays
`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) UnsetKeyRetentionDays()`

UnsetKeyRetentionDays ensures that no value is present for KeyRetentionDays, not even an explicit nil
### GetExplicitCredentials

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) GetExplicitCredentials() bool`

GetExplicitCredentials returns the ExplicitCredentials field if non-nil, zero value otherwise.

### GetExplicitCredentialsOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) GetExplicitCredentialsOk() (*bool, bool)`

GetExplicitCredentialsOk returns a tuple with the ExplicitCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicitCredentials

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) SetExplicitCredentials(v bool)`

SetExplicitCredentials sets ExplicitCredentials field to given value.

### HasExplicitCredentials

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) HasExplicitCredentials() bool`

HasExplicitCredentials returns a boolean if a field has been set.

### GetSubscriberTerms

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) GetSubscriberTerms() bool`

GetSubscriberTerms returns the SubscriberTerms field if non-nil, zero value otherwise.

### GetSubscriberTermsOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) GetSubscriberTermsOk() (*bool, bool)`

GetSubscriberTermsOk returns a tuple with the SubscriberTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriberTerms

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) SetSubscriberTerms(v bool)`

SetSubscriberTerms sets SubscriberTerms field to given value.

### HasSubscriberTerms

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) HasSubscriberTerms() bool`

HasSubscriberTerms returns a boolean if a field has been set.

### GetExplicitUser

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) GetExplicitUser() string`

GetExplicitUser returns the ExplicitUser field if non-nil, zero value otherwise.

### GetExplicitUserOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) GetExplicitUserOk() (*string, bool)`

GetExplicitUserOk returns a tuple with the ExplicitUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicitUser

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) SetExplicitUser(v string)`

SetExplicitUser sets ExplicitUser field to given value.

### HasExplicitUser

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) HasExplicitUser() bool`

HasExplicitUser returns a boolean if a field has been set.

### SetExplicitUserNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) SetExplicitUserNil(b bool)`

 SetExplicitUserNil sets the value for ExplicitUser to be an explicit nil

### UnsetExplicitUser
`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) UnsetExplicitUser()`

UnsetExplicitUser ensures that no value is present for ExplicitUser, not even an explicit nil
### GetExplicitPassword

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) GetExplicitPassword() ModelsKeyfactorAPISecret`

GetExplicitPassword returns the ExplicitPassword field if non-nil, zero value otherwise.

### GetExplicitPasswordOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) GetExplicitPasswordOk() (*ModelsKeyfactorAPISecret, bool)`

GetExplicitPasswordOk returns a tuple with the ExplicitPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicitPassword

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) SetExplicitPassword(v ModelsKeyfactorAPISecret)`

SetExplicitPassword sets ExplicitPassword field to given value.

### HasExplicitPassword

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) HasExplicitPassword() bool`

HasExplicitPassword returns a boolean if a field has been set.

### GetUseAllowedRequesters

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) GetUseAllowedRequesters() bool`

GetUseAllowedRequesters returns the UseAllowedRequesters field if non-nil, zero value otherwise.

### GetUseAllowedRequestersOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) GetUseAllowedRequestersOk() (*bool, bool)`

GetUseAllowedRequestersOk returns a tuple with the UseAllowedRequesters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAllowedRequesters

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) SetUseAllowedRequesters(v bool)`

SetUseAllowedRequesters sets UseAllowedRequesters field to given value.

### HasUseAllowedRequesters

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) HasUseAllowedRequesters() bool`

HasUseAllowedRequesters returns a boolean if a field has been set.

### GetAllowedRequesters

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) GetAllowedRequesters() []string`

GetAllowedRequesters returns the AllowedRequesters field if non-nil, zero value otherwise.

### GetAllowedRequestersOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) GetAllowedRequestersOk() (*[]string, bool)`

GetAllowedRequestersOk returns a tuple with the AllowedRequesters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedRequesters

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) SetAllowedRequesters(v []string)`

SetAllowedRequesters sets AllowedRequesters field to given value.

### HasAllowedRequesters

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) HasAllowedRequesters() bool`

HasAllowedRequesters returns a boolean if a field has been set.

### SetAllowedRequestersNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) SetAllowedRequestersNil(b bool)`

 SetAllowedRequestersNil sets the value for AllowedRequesters to be an explicit nil

### UnsetAllowedRequesters
`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) UnsetAllowedRequesters()`

UnsetAllowedRequesters ensures that no value is present for AllowedRequesters, not even an explicit nil
### GetFullScan

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) GetFullScan() KeyfactorCommonSchedulingKeyfactorSchedule`

GetFullScan returns the FullScan field if non-nil, zero value otherwise.

### GetFullScanOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) GetFullScanOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetFullScanOk returns a tuple with the FullScan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullScan

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) SetFullScan(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetFullScan sets FullScan field to given value.

### HasFullScan

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) HasFullScan() bool`

HasFullScan returns a boolean if a field has been set.

### GetIncrementalScan

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) GetIncrementalScan() KeyfactorCommonSchedulingKeyfactorSchedule`

GetIncrementalScan returns the IncrementalScan field if non-nil, zero value otherwise.

### GetIncrementalScanOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) GetIncrementalScanOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetIncrementalScanOk returns a tuple with the IncrementalScan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncrementalScan

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) SetIncrementalScan(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetIncrementalScan sets IncrementalScan field to given value.

### HasIncrementalScan

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) HasIncrementalScan() bool`

HasIncrementalScan returns a boolean if a field has been set.

### GetThresholdCheck

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) GetThresholdCheck() KeyfactorCommonSchedulingKeyfactorSchedule`

GetThresholdCheck returns the ThresholdCheck field if non-nil, zero value otherwise.

### GetThresholdCheckOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) GetThresholdCheckOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetThresholdCheckOk returns a tuple with the ThresholdCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdCheck

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) SetThresholdCheck(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetThresholdCheck sets ThresholdCheck field to given value.

### HasThresholdCheck

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) HasThresholdCheck() bool`

HasThresholdCheck returns a boolean if a field has been set.

### GetAuthCertificatePassword

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) GetAuthCertificatePassword() ModelsKeyfactorAPISecret`

GetAuthCertificatePassword returns the AuthCertificatePassword field if non-nil, zero value otherwise.

### GetAuthCertificatePasswordOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) GetAuthCertificatePasswordOk() (*ModelsKeyfactorAPISecret, bool)`

GetAuthCertificatePasswordOk returns a tuple with the AuthCertificatePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCertificatePassword

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) SetAuthCertificatePassword(v ModelsKeyfactorAPISecret)`

SetAuthCertificatePassword sets AuthCertificatePassword field to given value.

### HasAuthCertificatePassword

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) HasAuthCertificatePassword() bool`

HasAuthCertificatePassword returns a boolean if a field has been set.

### GetAuthCertificate

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) GetAuthCertificate() ModelsKeyfactorAPISecret`

GetAuthCertificate returns the AuthCertificate field if non-nil, zero value otherwise.

### GetAuthCertificateOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) GetAuthCertificateOk() (*ModelsKeyfactorAPISecret, bool)`

GetAuthCertificateOk returns a tuple with the AuthCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCertificate

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) SetAuthCertificate(v ModelsKeyfactorAPISecret)`

SetAuthCertificate sets AuthCertificate field to given value.

### HasAuthCertificate

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) HasAuthCertificate() bool`

HasAuthCertificate returns a boolean if a field has been set.

### GetCaType

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) GetCaType() int32`

GetCaType returns the CaType field if non-nil, zero value otherwise.

### GetCaTypeOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) GetCaTypeOk() (*int32, bool)`

GetCaTypeOk returns a tuple with the CaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaType

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) SetCaType(v int32)`

SetCaType sets CaType field to given value.

### HasCaType

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) HasCaType() bool`

HasCaType returns a boolean if a field has been set.

### GetEnforceUniqueDN

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) GetEnforceUniqueDN() bool`

GetEnforceUniqueDN returns the EnforceUniqueDN field if non-nil, zero value otherwise.

### GetEnforceUniqueDNOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) GetEnforceUniqueDNOk() (*bool, bool)`

GetEnforceUniqueDNOk returns a tuple with the EnforceUniqueDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceUniqueDN

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) SetEnforceUniqueDN(v bool)`

SetEnforceUniqueDN sets EnforceUniqueDN field to given value.

### HasEnforceUniqueDN

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) HasEnforceUniqueDN() bool`

HasEnforceUniqueDN returns a boolean if a field has been set.

### GetAllowOneClickRenewals

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) GetAllowOneClickRenewals() bool`

GetAllowOneClickRenewals returns the AllowOneClickRenewals field if non-nil, zero value otherwise.

### GetAllowOneClickRenewalsOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) GetAllowOneClickRenewalsOk() (*bool, bool)`

GetAllowOneClickRenewalsOk returns a tuple with the AllowOneClickRenewals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowOneClickRenewals

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) SetAllowOneClickRenewals(v bool)`

SetAllowOneClickRenewals sets AllowOneClickRenewals field to given value.

### HasAllowOneClickRenewals

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) HasAllowOneClickRenewals() bool`

HasAllowOneClickRenewals returns a boolean if a field has been set.

### GetNewEndEntityOnRenewAndReissue

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) GetNewEndEntityOnRenewAndReissue() bool`

GetNewEndEntityOnRenewAndReissue returns the NewEndEntityOnRenewAndReissue field if non-nil, zero value otherwise.

### GetNewEndEntityOnRenewAndReissueOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) GetNewEndEntityOnRenewAndReissueOk() (*bool, bool)`

GetNewEndEntityOnRenewAndReissueOk returns a tuple with the NewEndEntityOnRenewAndReissue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewEndEntityOnRenewAndReissue

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) SetNewEndEntityOnRenewAndReissue(v bool)`

SetNewEndEntityOnRenewAndReissue sets NewEndEntityOnRenewAndReissue field to given value.

### HasNewEndEntityOnRenewAndReissue

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest) HasNewEndEntityOnRenewAndReissue() bool`

HasNewEndEntityOnRenewAndReissue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


