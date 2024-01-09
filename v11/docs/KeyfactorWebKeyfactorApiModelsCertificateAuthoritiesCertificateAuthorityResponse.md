# KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**LogicalName** | Pointer to **NullableString** |  | [optional] 
**HostName** | Pointer to **NullableString** |  | [optional] 
**Delegate** | Pointer to **bool** |  | [optional] 
**DelegateEnrollment** | Pointer to **bool** |  | [optional] 
**ForestRoot** | Pointer to **NullableString** |  | [optional] [readonly] 
**ConfigurationTenant** | Pointer to **NullableString** |  | [optional] 
**Remote** | Pointer to **bool** |  | [optional] 
**Agent** | Pointer to **NullableString** |  | [optional] 
**Standalone** | Pointer to **bool** |  | [optional] 
**MonitorThresholds** | Pointer to **bool** |  | [optional] 
**IssuanceMax** | Pointer to **NullableInt32** |  | [optional] 
**IssuanceMin** | Pointer to **NullableInt32** |  | [optional] 
**DenialMax** | Pointer to **NullableInt32** |  | [optional] 
**FailureMax** | Pointer to **NullableInt32** |  | [optional] 
**RfcEnforcement** | Pointer to **bool** |  | [optional] 
**Properties** | Pointer to **NullableString** |  | [optional] 
**AllowedEnrollmentTypes** | Pointer to [**CSSCMSCoreEnumsEnrollmentType**](CSSCMSCoreEnumsEnrollmentType.md) |  | [optional] 
**KeyRetention** | Pointer to [**CSSCMSCoreEnumsKeyRetentionPolicy**](CSSCMSCoreEnumsKeyRetentionPolicy.md) |  | [optional] 
**KeyRetentionDays** | Pointer to **NullableInt32** |  | [optional] 
**ExplicitCredentials** | Pointer to **bool** |  | [optional] 
**SubscriberTerms** | Pointer to **bool** |  | [optional] 
**ExplicitUser** | Pointer to **NullableString** |  | [optional] 
**ExplicitPassword** | Pointer to [**CSSCMSDataModelModelsKeyfactorAPISecret**](CSSCMSDataModelModelsKeyfactorAPISecret.md) |  | [optional] 
**UseAllowedRequesters** | Pointer to **bool** |  | [optional] 
**AllowedRequesters** | Pointer to **[]string** |  | [optional] 
**FullScan** | Pointer to [**KeyfactorCommonSchedulingKeyfactorSchedule**](KeyfactorCommonSchedulingKeyfactorSchedule.md) |  | [optional] 
**IncrementalScan** | Pointer to [**KeyfactorCommonSchedulingKeyfactorSchedule**](KeyfactorCommonSchedulingKeyfactorSchedule.md) |  | [optional] 
**ThresholdCheck** | Pointer to [**KeyfactorCommonSchedulingKeyfactorSchedule**](KeyfactorCommonSchedulingKeyfactorSchedule.md) |  | [optional] 
**CaType** | Pointer to [**CSSCMSCoreEnumsCertificateAuthorityType**](CSSCMSCoreEnumsCertificateAuthorityType.md) |  | [optional] 
**AuthCertificate** | Pointer to [**CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityAuthCertificate**](CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityAuthCertificate.md) |  | [optional] 
**EnforceUniqueDN** | Pointer to **bool** |  | [optional] 
**AllowOneClickRenewals** | Pointer to **bool** |  | [optional] 
**NewEndEntityOnRenewAndReissue** | Pointer to **bool** |  | [optional] 
**CaSyncScheduledTasks** | Pointer to [**[]CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityScheduledTask**](CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityScheduledTask.md) |  | [optional] 
**LastScan** | Pointer to **NullableString** |  | [optional] [readonly] 

## Methods

### NewKeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse

`func NewKeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse() *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse`

NewKeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse instantiates a new KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponseWithDefaults

`func NewKeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponseWithDefaults() *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse`

NewKeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponseWithDefaults instantiates a new KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLogicalName

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) GetLogicalName() string`

GetLogicalName returns the LogicalName field if non-nil, zero value otherwise.

### GetLogicalNameOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) GetLogicalNameOk() (*string, bool)`

GetLogicalNameOk returns a tuple with the LogicalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalName

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) SetLogicalName(v string)`

SetLogicalName sets LogicalName field to given value.

### HasLogicalName

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) HasLogicalName() bool`

HasLogicalName returns a boolean if a field has been set.

### SetLogicalNameNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) SetLogicalNameNil(b bool)`

 SetLogicalNameNil sets the value for LogicalName to be an explicit nil

### UnsetLogicalName
`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) UnsetLogicalName()`

UnsetLogicalName ensures that no value is present for LogicalName, not even an explicit nil
### GetHostName

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### SetHostNameNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) SetHostNameNil(b bool)`

 SetHostNameNil sets the value for HostName to be an explicit nil

### UnsetHostName
`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) UnsetHostName()`

UnsetHostName ensures that no value is present for HostName, not even an explicit nil
### GetDelegate

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) GetDelegate() bool`

GetDelegate returns the Delegate field if non-nil, zero value otherwise.

### GetDelegateOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) GetDelegateOk() (*bool, bool)`

GetDelegateOk returns a tuple with the Delegate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegate

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) SetDelegate(v bool)`

SetDelegate sets Delegate field to given value.

### HasDelegate

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) HasDelegate() bool`

HasDelegate returns a boolean if a field has been set.

### GetDelegateEnrollment

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) GetDelegateEnrollment() bool`

GetDelegateEnrollment returns the DelegateEnrollment field if non-nil, zero value otherwise.

### GetDelegateEnrollmentOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) GetDelegateEnrollmentOk() (*bool, bool)`

GetDelegateEnrollmentOk returns a tuple with the DelegateEnrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegateEnrollment

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) SetDelegateEnrollment(v bool)`

SetDelegateEnrollment sets DelegateEnrollment field to given value.

### HasDelegateEnrollment

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) HasDelegateEnrollment() bool`

HasDelegateEnrollment returns a boolean if a field has been set.

### GetForestRoot

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) GetForestRoot() string`

GetForestRoot returns the ForestRoot field if non-nil, zero value otherwise.

### GetForestRootOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) GetForestRootOk() (*string, bool)`

GetForestRootOk returns a tuple with the ForestRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForestRoot

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) SetForestRoot(v string)`

SetForestRoot sets ForestRoot field to given value.

### HasForestRoot

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) HasForestRoot() bool`

HasForestRoot returns a boolean if a field has been set.

### SetForestRootNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) SetForestRootNil(b bool)`

 SetForestRootNil sets the value for ForestRoot to be an explicit nil

### UnsetForestRoot
`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) UnsetForestRoot()`

UnsetForestRoot ensures that no value is present for ForestRoot, not even an explicit nil
### GetConfigurationTenant

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) GetConfigurationTenant() string`

GetConfigurationTenant returns the ConfigurationTenant field if non-nil, zero value otherwise.

### GetConfigurationTenantOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) GetConfigurationTenantOk() (*string, bool)`

GetConfigurationTenantOk returns a tuple with the ConfigurationTenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationTenant

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) SetConfigurationTenant(v string)`

SetConfigurationTenant sets ConfigurationTenant field to given value.

### HasConfigurationTenant

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) HasConfigurationTenant() bool`

HasConfigurationTenant returns a boolean if a field has been set.

### SetConfigurationTenantNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) SetConfigurationTenantNil(b bool)`

 SetConfigurationTenantNil sets the value for ConfigurationTenant to be an explicit nil

### UnsetConfigurationTenant
`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) UnsetConfigurationTenant()`

UnsetConfigurationTenant ensures that no value is present for ConfigurationTenant, not even an explicit nil
### GetRemote

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) GetRemote() bool`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) GetRemoteOk() (*bool, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) SetRemote(v bool)`

SetRemote sets Remote field to given value.

### HasRemote

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) HasRemote() bool`

HasRemote returns a boolean if a field has been set.

### GetAgent

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) GetAgent() string`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) GetAgentOk() (*string, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) SetAgent(v string)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) HasAgent() bool`

HasAgent returns a boolean if a field has been set.

### SetAgentNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) SetAgentNil(b bool)`

 SetAgentNil sets the value for Agent to be an explicit nil

### UnsetAgent
`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) UnsetAgent()`

UnsetAgent ensures that no value is present for Agent, not even an explicit nil
### GetStandalone

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) GetStandalone() bool`

GetStandalone returns the Standalone field if non-nil, zero value otherwise.

### GetStandaloneOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) GetStandaloneOk() (*bool, bool)`

GetStandaloneOk returns a tuple with the Standalone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandalone

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) SetStandalone(v bool)`

SetStandalone sets Standalone field to given value.

### HasStandalone

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) HasStandalone() bool`

HasStandalone returns a boolean if a field has been set.

### GetMonitorThresholds

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) GetMonitorThresholds() bool`

GetMonitorThresholds returns the MonitorThresholds field if non-nil, zero value otherwise.

### GetMonitorThresholdsOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) GetMonitorThresholdsOk() (*bool, bool)`

GetMonitorThresholdsOk returns a tuple with the MonitorThresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorThresholds

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) SetMonitorThresholds(v bool)`

SetMonitorThresholds sets MonitorThresholds field to given value.

### HasMonitorThresholds

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) HasMonitorThresholds() bool`

HasMonitorThresholds returns a boolean if a field has been set.

### GetIssuanceMax

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) GetIssuanceMax() int32`

GetIssuanceMax returns the IssuanceMax field if non-nil, zero value otherwise.

### GetIssuanceMaxOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) GetIssuanceMaxOk() (*int32, bool)`

GetIssuanceMaxOk returns a tuple with the IssuanceMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuanceMax

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) SetIssuanceMax(v int32)`

SetIssuanceMax sets IssuanceMax field to given value.

### HasIssuanceMax

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) HasIssuanceMax() bool`

HasIssuanceMax returns a boolean if a field has been set.

### SetIssuanceMaxNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) SetIssuanceMaxNil(b bool)`

 SetIssuanceMaxNil sets the value for IssuanceMax to be an explicit nil

### UnsetIssuanceMax
`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) UnsetIssuanceMax()`

UnsetIssuanceMax ensures that no value is present for IssuanceMax, not even an explicit nil
### GetIssuanceMin

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) GetIssuanceMin() int32`

GetIssuanceMin returns the IssuanceMin field if non-nil, zero value otherwise.

### GetIssuanceMinOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) GetIssuanceMinOk() (*int32, bool)`

GetIssuanceMinOk returns a tuple with the IssuanceMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuanceMin

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) SetIssuanceMin(v int32)`

SetIssuanceMin sets IssuanceMin field to given value.

### HasIssuanceMin

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) HasIssuanceMin() bool`

HasIssuanceMin returns a boolean if a field has been set.

### SetIssuanceMinNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) SetIssuanceMinNil(b bool)`

 SetIssuanceMinNil sets the value for IssuanceMin to be an explicit nil

### UnsetIssuanceMin
`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) UnsetIssuanceMin()`

UnsetIssuanceMin ensures that no value is present for IssuanceMin, not even an explicit nil
### GetDenialMax

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) GetDenialMax() int32`

GetDenialMax returns the DenialMax field if non-nil, zero value otherwise.

### GetDenialMaxOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) GetDenialMaxOk() (*int32, bool)`

GetDenialMaxOk returns a tuple with the DenialMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenialMax

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) SetDenialMax(v int32)`

SetDenialMax sets DenialMax field to given value.

### HasDenialMax

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) HasDenialMax() bool`

HasDenialMax returns a boolean if a field has been set.

### SetDenialMaxNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) SetDenialMaxNil(b bool)`

 SetDenialMaxNil sets the value for DenialMax to be an explicit nil

### UnsetDenialMax
`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) UnsetDenialMax()`

UnsetDenialMax ensures that no value is present for DenialMax, not even an explicit nil
### GetFailureMax

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) GetFailureMax() int32`

GetFailureMax returns the FailureMax field if non-nil, zero value otherwise.

### GetFailureMaxOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) GetFailureMaxOk() (*int32, bool)`

GetFailureMaxOk returns a tuple with the FailureMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureMax

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) SetFailureMax(v int32)`

SetFailureMax sets FailureMax field to given value.

### HasFailureMax

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) HasFailureMax() bool`

HasFailureMax returns a boolean if a field has been set.

### SetFailureMaxNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) SetFailureMaxNil(b bool)`

 SetFailureMaxNil sets the value for FailureMax to be an explicit nil

### UnsetFailureMax
`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) UnsetFailureMax()`

UnsetFailureMax ensures that no value is present for FailureMax, not even an explicit nil
### GetRfcEnforcement

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) GetRfcEnforcement() bool`

GetRfcEnforcement returns the RfcEnforcement field if non-nil, zero value otherwise.

### GetRfcEnforcementOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) GetRfcEnforcementOk() (*bool, bool)`

GetRfcEnforcementOk returns a tuple with the RfcEnforcement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRfcEnforcement

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) SetRfcEnforcement(v bool)`

SetRfcEnforcement sets RfcEnforcement field to given value.

### HasRfcEnforcement

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) HasRfcEnforcement() bool`

HasRfcEnforcement returns a boolean if a field has been set.

### GetProperties

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) GetProperties() string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) GetPropertiesOk() (*string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) SetProperties(v string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetAllowedEnrollmentTypes

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) GetAllowedEnrollmentTypes() CSSCMSCoreEnumsEnrollmentType`

GetAllowedEnrollmentTypes returns the AllowedEnrollmentTypes field if non-nil, zero value otherwise.

### GetAllowedEnrollmentTypesOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) GetAllowedEnrollmentTypesOk() (*CSSCMSCoreEnumsEnrollmentType, bool)`

GetAllowedEnrollmentTypesOk returns a tuple with the AllowedEnrollmentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedEnrollmentTypes

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) SetAllowedEnrollmentTypes(v CSSCMSCoreEnumsEnrollmentType)`

SetAllowedEnrollmentTypes sets AllowedEnrollmentTypes field to given value.

### HasAllowedEnrollmentTypes

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) HasAllowedEnrollmentTypes() bool`

HasAllowedEnrollmentTypes returns a boolean if a field has been set.

### GetKeyRetention

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) GetKeyRetention() CSSCMSCoreEnumsKeyRetentionPolicy`

GetKeyRetention returns the KeyRetention field if non-nil, zero value otherwise.

### GetKeyRetentionOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) GetKeyRetentionOk() (*CSSCMSCoreEnumsKeyRetentionPolicy, bool)`

GetKeyRetentionOk returns a tuple with the KeyRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyRetention

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) SetKeyRetention(v CSSCMSCoreEnumsKeyRetentionPolicy)`

SetKeyRetention sets KeyRetention field to given value.

### HasKeyRetention

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) HasKeyRetention() bool`

HasKeyRetention returns a boolean if a field has been set.

### GetKeyRetentionDays

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) GetKeyRetentionDays() int32`

GetKeyRetentionDays returns the KeyRetentionDays field if non-nil, zero value otherwise.

### GetKeyRetentionDaysOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) GetKeyRetentionDaysOk() (*int32, bool)`

GetKeyRetentionDaysOk returns a tuple with the KeyRetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyRetentionDays

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) SetKeyRetentionDays(v int32)`

SetKeyRetentionDays sets KeyRetentionDays field to given value.

### HasKeyRetentionDays

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) HasKeyRetentionDays() bool`

HasKeyRetentionDays returns a boolean if a field has been set.

### SetKeyRetentionDaysNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) SetKeyRetentionDaysNil(b bool)`

 SetKeyRetentionDaysNil sets the value for KeyRetentionDays to be an explicit nil

### UnsetKeyRetentionDays
`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) UnsetKeyRetentionDays()`

UnsetKeyRetentionDays ensures that no value is present for KeyRetentionDays, not even an explicit nil
### GetExplicitCredentials

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) GetExplicitCredentials() bool`

GetExplicitCredentials returns the ExplicitCredentials field if non-nil, zero value otherwise.

### GetExplicitCredentialsOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) GetExplicitCredentialsOk() (*bool, bool)`

GetExplicitCredentialsOk returns a tuple with the ExplicitCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicitCredentials

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) SetExplicitCredentials(v bool)`

SetExplicitCredentials sets ExplicitCredentials field to given value.

### HasExplicitCredentials

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) HasExplicitCredentials() bool`

HasExplicitCredentials returns a boolean if a field has been set.

### GetSubscriberTerms

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) GetSubscriberTerms() bool`

GetSubscriberTerms returns the SubscriberTerms field if non-nil, zero value otherwise.

### GetSubscriberTermsOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) GetSubscriberTermsOk() (*bool, bool)`

GetSubscriberTermsOk returns a tuple with the SubscriberTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriberTerms

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) SetSubscriberTerms(v bool)`

SetSubscriberTerms sets SubscriberTerms field to given value.

### HasSubscriberTerms

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) HasSubscriberTerms() bool`

HasSubscriberTerms returns a boolean if a field has been set.

### GetExplicitUser

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) GetExplicitUser() string`

GetExplicitUser returns the ExplicitUser field if non-nil, zero value otherwise.

### GetExplicitUserOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) GetExplicitUserOk() (*string, bool)`

GetExplicitUserOk returns a tuple with the ExplicitUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicitUser

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) SetExplicitUser(v string)`

SetExplicitUser sets ExplicitUser field to given value.

### HasExplicitUser

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) HasExplicitUser() bool`

HasExplicitUser returns a boolean if a field has been set.

### SetExplicitUserNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) SetExplicitUserNil(b bool)`

 SetExplicitUserNil sets the value for ExplicitUser to be an explicit nil

### UnsetExplicitUser
`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) UnsetExplicitUser()`

UnsetExplicitUser ensures that no value is present for ExplicitUser, not even an explicit nil
### GetExplicitPassword

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) GetExplicitPassword() CSSCMSDataModelModelsKeyfactorAPISecret`

GetExplicitPassword returns the ExplicitPassword field if non-nil, zero value otherwise.

### GetExplicitPasswordOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) GetExplicitPasswordOk() (*CSSCMSDataModelModelsKeyfactorAPISecret, bool)`

GetExplicitPasswordOk returns a tuple with the ExplicitPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicitPassword

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) SetExplicitPassword(v CSSCMSDataModelModelsKeyfactorAPISecret)`

SetExplicitPassword sets ExplicitPassword field to given value.

### HasExplicitPassword

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) HasExplicitPassword() bool`

HasExplicitPassword returns a boolean if a field has been set.

### GetUseAllowedRequesters

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) GetUseAllowedRequesters() bool`

GetUseAllowedRequesters returns the UseAllowedRequesters field if non-nil, zero value otherwise.

### GetUseAllowedRequestersOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) GetUseAllowedRequestersOk() (*bool, bool)`

GetUseAllowedRequestersOk returns a tuple with the UseAllowedRequesters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAllowedRequesters

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) SetUseAllowedRequesters(v bool)`

SetUseAllowedRequesters sets UseAllowedRequesters field to given value.

### HasUseAllowedRequesters

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) HasUseAllowedRequesters() bool`

HasUseAllowedRequesters returns a boolean if a field has been set.

### GetAllowedRequesters

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) GetAllowedRequesters() []string`

GetAllowedRequesters returns the AllowedRequesters field if non-nil, zero value otherwise.

### GetAllowedRequestersOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) GetAllowedRequestersOk() (*[]string, bool)`

GetAllowedRequestersOk returns a tuple with the AllowedRequesters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedRequesters

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) SetAllowedRequesters(v []string)`

SetAllowedRequesters sets AllowedRequesters field to given value.

### HasAllowedRequesters

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) HasAllowedRequesters() bool`

HasAllowedRequesters returns a boolean if a field has been set.

### SetAllowedRequestersNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) SetAllowedRequestersNil(b bool)`

 SetAllowedRequestersNil sets the value for AllowedRequesters to be an explicit nil

### UnsetAllowedRequesters
`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) UnsetAllowedRequesters()`

UnsetAllowedRequesters ensures that no value is present for AllowedRequesters, not even an explicit nil
### GetFullScan

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) GetFullScan() KeyfactorCommonSchedulingKeyfactorSchedule`

GetFullScan returns the FullScan field if non-nil, zero value otherwise.

### GetFullScanOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) GetFullScanOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetFullScanOk returns a tuple with the FullScan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullScan

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) SetFullScan(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetFullScan sets FullScan field to given value.

### HasFullScan

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) HasFullScan() bool`

HasFullScan returns a boolean if a field has been set.

### GetIncrementalScan

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) GetIncrementalScan() KeyfactorCommonSchedulingKeyfactorSchedule`

GetIncrementalScan returns the IncrementalScan field if non-nil, zero value otherwise.

### GetIncrementalScanOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) GetIncrementalScanOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetIncrementalScanOk returns a tuple with the IncrementalScan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncrementalScan

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) SetIncrementalScan(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetIncrementalScan sets IncrementalScan field to given value.

### HasIncrementalScan

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) HasIncrementalScan() bool`

HasIncrementalScan returns a boolean if a field has been set.

### GetThresholdCheck

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) GetThresholdCheck() KeyfactorCommonSchedulingKeyfactorSchedule`

GetThresholdCheck returns the ThresholdCheck field if non-nil, zero value otherwise.

### GetThresholdCheckOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) GetThresholdCheckOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetThresholdCheckOk returns a tuple with the ThresholdCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdCheck

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) SetThresholdCheck(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetThresholdCheck sets ThresholdCheck field to given value.

### HasThresholdCheck

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) HasThresholdCheck() bool`

HasThresholdCheck returns a boolean if a field has been set.

### GetCaType

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) GetCaType() CSSCMSCoreEnumsCertificateAuthorityType`

GetCaType returns the CaType field if non-nil, zero value otherwise.

### GetCaTypeOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) GetCaTypeOk() (*CSSCMSCoreEnumsCertificateAuthorityType, bool)`

GetCaTypeOk returns a tuple with the CaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaType

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) SetCaType(v CSSCMSCoreEnumsCertificateAuthorityType)`

SetCaType sets CaType field to given value.

### HasCaType

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) HasCaType() bool`

HasCaType returns a boolean if a field has been set.

### GetAuthCertificate

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) GetAuthCertificate() CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityAuthCertificate`

GetAuthCertificate returns the AuthCertificate field if non-nil, zero value otherwise.

### GetAuthCertificateOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) GetAuthCertificateOk() (*CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityAuthCertificate, bool)`

GetAuthCertificateOk returns a tuple with the AuthCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCertificate

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) SetAuthCertificate(v CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityAuthCertificate)`

SetAuthCertificate sets AuthCertificate field to given value.

### HasAuthCertificate

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) HasAuthCertificate() bool`

HasAuthCertificate returns a boolean if a field has been set.

### GetEnforceUniqueDN

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) GetEnforceUniqueDN() bool`

GetEnforceUniqueDN returns the EnforceUniqueDN field if non-nil, zero value otherwise.

### GetEnforceUniqueDNOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) GetEnforceUniqueDNOk() (*bool, bool)`

GetEnforceUniqueDNOk returns a tuple with the EnforceUniqueDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceUniqueDN

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) SetEnforceUniqueDN(v bool)`

SetEnforceUniqueDN sets EnforceUniqueDN field to given value.

### HasEnforceUniqueDN

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) HasEnforceUniqueDN() bool`

HasEnforceUniqueDN returns a boolean if a field has been set.

### GetAllowOneClickRenewals

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) GetAllowOneClickRenewals() bool`

GetAllowOneClickRenewals returns the AllowOneClickRenewals field if non-nil, zero value otherwise.

### GetAllowOneClickRenewalsOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) GetAllowOneClickRenewalsOk() (*bool, bool)`

GetAllowOneClickRenewalsOk returns a tuple with the AllowOneClickRenewals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowOneClickRenewals

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) SetAllowOneClickRenewals(v bool)`

SetAllowOneClickRenewals sets AllowOneClickRenewals field to given value.

### HasAllowOneClickRenewals

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) HasAllowOneClickRenewals() bool`

HasAllowOneClickRenewals returns a boolean if a field has been set.

### GetNewEndEntityOnRenewAndReissue

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) GetNewEndEntityOnRenewAndReissue() bool`

GetNewEndEntityOnRenewAndReissue returns the NewEndEntityOnRenewAndReissue field if non-nil, zero value otherwise.

### GetNewEndEntityOnRenewAndReissueOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) GetNewEndEntityOnRenewAndReissueOk() (*bool, bool)`

GetNewEndEntityOnRenewAndReissueOk returns a tuple with the NewEndEntityOnRenewAndReissue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewEndEntityOnRenewAndReissue

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) SetNewEndEntityOnRenewAndReissue(v bool)`

SetNewEndEntityOnRenewAndReissue sets NewEndEntityOnRenewAndReissue field to given value.

### HasNewEndEntityOnRenewAndReissue

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) HasNewEndEntityOnRenewAndReissue() bool`

HasNewEndEntityOnRenewAndReissue returns a boolean if a field has been set.

### GetCaSyncScheduledTasks

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) GetCaSyncScheduledTasks() []CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityScheduledTask`

GetCaSyncScheduledTasks returns the CaSyncScheduledTasks field if non-nil, zero value otherwise.

### GetCaSyncScheduledTasksOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) GetCaSyncScheduledTasksOk() (*[]CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityScheduledTask, bool)`

GetCaSyncScheduledTasksOk returns a tuple with the CaSyncScheduledTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaSyncScheduledTasks

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) SetCaSyncScheduledTasks(v []CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityScheduledTask)`

SetCaSyncScheduledTasks sets CaSyncScheduledTasks field to given value.

### HasCaSyncScheduledTasks

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) HasCaSyncScheduledTasks() bool`

HasCaSyncScheduledTasks returns a boolean if a field has been set.

### SetCaSyncScheduledTasksNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) SetCaSyncScheduledTasksNil(b bool)`

 SetCaSyncScheduledTasksNil sets the value for CaSyncScheduledTasks to be an explicit nil

### UnsetCaSyncScheduledTasks
`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) UnsetCaSyncScheduledTasks()`

UnsetCaSyncScheduledTasks ensures that no value is present for CaSyncScheduledTasks, not even an explicit nil
### GetLastScan

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) GetLastScan() string`

GetLastScan returns the LastScan field if non-nil, zero value otherwise.

### GetLastScanOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) GetLastScanOk() (*string, bool)`

GetLastScanOk returns a tuple with the LastScan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastScan

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) SetLastScan(v string)`

SetLastScan sets LastScan field to given value.

### HasLastScan

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) HasLastScan() bool`

HasLastScan returns a boolean if a field has been set.

### SetLastScanNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) SetLastScanNil(b bool)`

 SetLastScanNil sets the value for LastScan to be an explicit nil

### UnsetLastScan
`func (o *KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse) UnsetLastScan()`

UnsetLastScan ensures that no value is present for LastScan, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


