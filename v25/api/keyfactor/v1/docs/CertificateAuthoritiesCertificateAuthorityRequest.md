# CertificateAuthoritiesCertificateAuthorityRequest

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
**UseCAConnector** | Pointer to **bool** |  | [optional] 
**ConnectorPool** | Pointer to **NullableString** |  | [optional] 
**Remote** | Pointer to **bool** |  | [optional] 
**Agent** | Pointer to **NullableString** |  | [optional] 
**Standalone** | Pointer to **bool** |  | [optional] 
**MonitorThresholds** | Pointer to **bool** |  | [optional] 
**IssuanceMax** | Pointer to **NullableInt32** |  | [optional] 
**IssuanceMin** | Pointer to **NullableInt32** |  | [optional] 
**FailureMax** | Pointer to **NullableInt32** |  | [optional] 
**RFCEnforcement** | Pointer to **bool** |  | [optional] 
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
**AuthCertificatePassword** | Pointer to **map[string]interface{}** | Overrides default KeyfactorAPISecret with inline SecretValue | [optional] 
**AuthCertificate** | Pointer to **map[string]interface{}** | Overrides default KeyfactorAPISecret with inline SecretValue | [optional] 
**CAType** | Pointer to [**CSSCMSCoreEnumsCertificateAuthorityType**](CSSCMSCoreEnumsCertificateAuthorityType.md) |  | [optional] 
**EnforceUniqueDN** | Pointer to **bool** |  | [optional] 
**AllowOneClickRenewals** | Pointer to **bool** |  | [optional] 
**NewEndEntityOnRenewAndReissue** | Pointer to **bool** |  | [optional] 
**TokenURL** | Pointer to **NullableString** |  | [optional] 
**ClientId** | Pointer to **NullableString** |  | [optional] 
**ClientSecret** | Pointer to [**CSSCMSDataModelModelsKeyfactorAPISecret**](CSSCMSDataModelModelsKeyfactorAPISecret.md) |  | [optional] 
**Scope** | Pointer to **NullableString** |  | [optional] 
**Audience** | Pointer to **NullableString** |  | [optional] 
**CertificateCleanupEnabled** | Pointer to **NullableBool** |  | [optional] 
**TimeAfterExpiration** | Pointer to **NullableInt32** |  | [optional] 
**TimeAfterExpirationUnits** | Pointer to [**CSSCMSDataModelEnumsCertificateCleanupTimeUnits**](CSSCMSDataModelEnumsCertificateCleanupTimeUnits.md) |  | [optional] 
**DeleteWithArchivedKey** | Pointer to **NullableBool** |  | [optional] 
**UseForEnrollment** | Pointer to **bool** |  | [optional] 

## Methods

### NewCertificateAuthoritiesCertificateAuthorityRequest

`func NewCertificateAuthoritiesCertificateAuthorityRequest() *CertificateAuthoritiesCertificateAuthorityRequest`

NewCertificateAuthoritiesCertificateAuthorityRequest instantiates a new CertificateAuthoritiesCertificateAuthorityRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateAuthoritiesCertificateAuthorityRequestWithDefaults

`func NewCertificateAuthoritiesCertificateAuthorityRequestWithDefaults() *CertificateAuthoritiesCertificateAuthorityRequest`

NewCertificateAuthoritiesCertificateAuthorityRequestWithDefaults instantiates a new CertificateAuthoritiesCertificateAuthorityRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLogicalName

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetLogicalName() string`

GetLogicalName returns the LogicalName field if non-nil, zero value otherwise.

### GetLogicalNameOk

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetLogicalNameOk() (*string, bool)`

GetLogicalNameOk returns a tuple with the LogicalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalName

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) SetLogicalName(v string)`

SetLogicalName sets LogicalName field to given value.

### HasLogicalName

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) HasLogicalName() bool`

HasLogicalName returns a boolean if a field has been set.

### SetLogicalNameNil

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) SetLogicalNameNil(b bool)`

 SetLogicalNameNil sets the value for LogicalName to be an explicit nil

### UnsetLogicalName
`func (o *CertificateAuthoritiesCertificateAuthorityRequest) UnsetLogicalName()`

UnsetLogicalName ensures that no value is present for LogicalName, not even an explicit nil
### GetHostName

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### SetHostNameNil

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) SetHostNameNil(b bool)`

 SetHostNameNil sets the value for HostName to be an explicit nil

### UnsetHostName
`func (o *CertificateAuthoritiesCertificateAuthorityRequest) UnsetHostName()`

UnsetHostName ensures that no value is present for HostName, not even an explicit nil
### GetDelegate

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetDelegate() bool`

GetDelegate returns the Delegate field if non-nil, zero value otherwise.

### GetDelegateOk

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetDelegateOk() (*bool, bool)`

GetDelegateOk returns a tuple with the Delegate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegate

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) SetDelegate(v bool)`

SetDelegate sets Delegate field to given value.

### HasDelegate

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) HasDelegate() bool`

HasDelegate returns a boolean if a field has been set.

### GetDelegateEnrollment

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetDelegateEnrollment() bool`

GetDelegateEnrollment returns the DelegateEnrollment field if non-nil, zero value otherwise.

### GetDelegateEnrollmentOk

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetDelegateEnrollmentOk() (*bool, bool)`

GetDelegateEnrollmentOk returns a tuple with the DelegateEnrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegateEnrollment

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) SetDelegateEnrollment(v bool)`

SetDelegateEnrollment sets DelegateEnrollment field to given value.

### HasDelegateEnrollment

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) HasDelegateEnrollment() bool`

HasDelegateEnrollment returns a boolean if a field has been set.

### GetForestRoot

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetForestRoot() string`

GetForestRoot returns the ForestRoot field if non-nil, zero value otherwise.

### GetForestRootOk

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetForestRootOk() (*string, bool)`

GetForestRootOk returns a tuple with the ForestRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForestRoot

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) SetForestRoot(v string)`

SetForestRoot sets ForestRoot field to given value.

### HasForestRoot

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) HasForestRoot() bool`

HasForestRoot returns a boolean if a field has been set.

### SetForestRootNil

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) SetForestRootNil(b bool)`

 SetForestRootNil sets the value for ForestRoot to be an explicit nil

### UnsetForestRoot
`func (o *CertificateAuthoritiesCertificateAuthorityRequest) UnsetForestRoot()`

UnsetForestRoot ensures that no value is present for ForestRoot, not even an explicit nil
### GetConfigurationTenant

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetConfigurationTenant() string`

GetConfigurationTenant returns the ConfigurationTenant field if non-nil, zero value otherwise.

### GetConfigurationTenantOk

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetConfigurationTenantOk() (*string, bool)`

GetConfigurationTenantOk returns a tuple with the ConfigurationTenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationTenant

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) SetConfigurationTenant(v string)`

SetConfigurationTenant sets ConfigurationTenant field to given value.

### HasConfigurationTenant

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) HasConfigurationTenant() bool`

HasConfigurationTenant returns a boolean if a field has been set.

### SetConfigurationTenantNil

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) SetConfigurationTenantNil(b bool)`

 SetConfigurationTenantNil sets the value for ConfigurationTenant to be an explicit nil

### UnsetConfigurationTenant
`func (o *CertificateAuthoritiesCertificateAuthorityRequest) UnsetConfigurationTenant()`

UnsetConfigurationTenant ensures that no value is present for ConfigurationTenant, not even an explicit nil
### GetUseCAConnector

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetUseCAConnector() bool`

GetUseCAConnector returns the UseCAConnector field if non-nil, zero value otherwise.

### GetUseCAConnectorOk

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetUseCAConnectorOk() (*bool, bool)`

GetUseCAConnectorOk returns a tuple with the UseCAConnector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCAConnector

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) SetUseCAConnector(v bool)`

SetUseCAConnector sets UseCAConnector field to given value.

### HasUseCAConnector

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) HasUseCAConnector() bool`

HasUseCAConnector returns a boolean if a field has been set.

### GetConnectorPool

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetConnectorPool() string`

GetConnectorPool returns the ConnectorPool field if non-nil, zero value otherwise.

### GetConnectorPoolOk

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetConnectorPoolOk() (*string, bool)`

GetConnectorPoolOk returns a tuple with the ConnectorPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorPool

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) SetConnectorPool(v string)`

SetConnectorPool sets ConnectorPool field to given value.

### HasConnectorPool

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) HasConnectorPool() bool`

HasConnectorPool returns a boolean if a field has been set.

### SetConnectorPoolNil

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) SetConnectorPoolNil(b bool)`

 SetConnectorPoolNil sets the value for ConnectorPool to be an explicit nil

### UnsetConnectorPool
`func (o *CertificateAuthoritiesCertificateAuthorityRequest) UnsetConnectorPool()`

UnsetConnectorPool ensures that no value is present for ConnectorPool, not even an explicit nil
### GetRemote

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetRemote() bool`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetRemoteOk() (*bool, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) SetRemote(v bool)`

SetRemote sets Remote field to given value.

### HasRemote

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) HasRemote() bool`

HasRemote returns a boolean if a field has been set.

### GetAgent

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetAgent() string`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetAgentOk() (*string, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) SetAgent(v string)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) HasAgent() bool`

HasAgent returns a boolean if a field has been set.

### SetAgentNil

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) SetAgentNil(b bool)`

 SetAgentNil sets the value for Agent to be an explicit nil

### UnsetAgent
`func (o *CertificateAuthoritiesCertificateAuthorityRequest) UnsetAgent()`

UnsetAgent ensures that no value is present for Agent, not even an explicit nil
### GetStandalone

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetStandalone() bool`

GetStandalone returns the Standalone field if non-nil, zero value otherwise.

### GetStandaloneOk

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetStandaloneOk() (*bool, bool)`

GetStandaloneOk returns a tuple with the Standalone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandalone

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) SetStandalone(v bool)`

SetStandalone sets Standalone field to given value.

### HasStandalone

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) HasStandalone() bool`

HasStandalone returns a boolean if a field has been set.

### GetMonitorThresholds

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetMonitorThresholds() bool`

GetMonitorThresholds returns the MonitorThresholds field if non-nil, zero value otherwise.

### GetMonitorThresholdsOk

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetMonitorThresholdsOk() (*bool, bool)`

GetMonitorThresholdsOk returns a tuple with the MonitorThresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorThresholds

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) SetMonitorThresholds(v bool)`

SetMonitorThresholds sets MonitorThresholds field to given value.

### HasMonitorThresholds

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) HasMonitorThresholds() bool`

HasMonitorThresholds returns a boolean if a field has been set.

### GetIssuanceMax

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetIssuanceMax() int32`

GetIssuanceMax returns the IssuanceMax field if non-nil, zero value otherwise.

### GetIssuanceMaxOk

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetIssuanceMaxOk() (*int32, bool)`

GetIssuanceMaxOk returns a tuple with the IssuanceMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuanceMax

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) SetIssuanceMax(v int32)`

SetIssuanceMax sets IssuanceMax field to given value.

### HasIssuanceMax

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) HasIssuanceMax() bool`

HasIssuanceMax returns a boolean if a field has been set.

### SetIssuanceMaxNil

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) SetIssuanceMaxNil(b bool)`

 SetIssuanceMaxNil sets the value for IssuanceMax to be an explicit nil

### UnsetIssuanceMax
`func (o *CertificateAuthoritiesCertificateAuthorityRequest) UnsetIssuanceMax()`

UnsetIssuanceMax ensures that no value is present for IssuanceMax, not even an explicit nil
### GetIssuanceMin

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetIssuanceMin() int32`

GetIssuanceMin returns the IssuanceMin field if non-nil, zero value otherwise.

### GetIssuanceMinOk

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetIssuanceMinOk() (*int32, bool)`

GetIssuanceMinOk returns a tuple with the IssuanceMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuanceMin

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) SetIssuanceMin(v int32)`

SetIssuanceMin sets IssuanceMin field to given value.

### HasIssuanceMin

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) HasIssuanceMin() bool`

HasIssuanceMin returns a boolean if a field has been set.

### SetIssuanceMinNil

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) SetIssuanceMinNil(b bool)`

 SetIssuanceMinNil sets the value for IssuanceMin to be an explicit nil

### UnsetIssuanceMin
`func (o *CertificateAuthoritiesCertificateAuthorityRequest) UnsetIssuanceMin()`

UnsetIssuanceMin ensures that no value is present for IssuanceMin, not even an explicit nil
### GetFailureMax

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetFailureMax() int32`

GetFailureMax returns the FailureMax field if non-nil, zero value otherwise.

### GetFailureMaxOk

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetFailureMaxOk() (*int32, bool)`

GetFailureMaxOk returns a tuple with the FailureMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureMax

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) SetFailureMax(v int32)`

SetFailureMax sets FailureMax field to given value.

### HasFailureMax

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) HasFailureMax() bool`

HasFailureMax returns a boolean if a field has been set.

### SetFailureMaxNil

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) SetFailureMaxNil(b bool)`

 SetFailureMaxNil sets the value for FailureMax to be an explicit nil

### UnsetFailureMax
`func (o *CertificateAuthoritiesCertificateAuthorityRequest) UnsetFailureMax()`

UnsetFailureMax ensures that no value is present for FailureMax, not even an explicit nil
### GetRFCEnforcement

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetRFCEnforcement() bool`

GetRFCEnforcement returns the RFCEnforcement field if non-nil, zero value otherwise.

### GetRFCEnforcementOk

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetRFCEnforcementOk() (*bool, bool)`

GetRFCEnforcementOk returns a tuple with the RFCEnforcement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRFCEnforcement

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) SetRFCEnforcement(v bool)`

SetRFCEnforcement sets RFCEnforcement field to given value.

### HasRFCEnforcement

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) HasRFCEnforcement() bool`

HasRFCEnforcement returns a boolean if a field has been set.

### GetProperties

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetProperties() string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetPropertiesOk() (*string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) SetProperties(v string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *CertificateAuthoritiesCertificateAuthorityRequest) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetAllowedEnrollmentTypes

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetAllowedEnrollmentTypes() CSSCMSCoreEnumsEnrollmentType`

GetAllowedEnrollmentTypes returns the AllowedEnrollmentTypes field if non-nil, zero value otherwise.

### GetAllowedEnrollmentTypesOk

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetAllowedEnrollmentTypesOk() (*CSSCMSCoreEnumsEnrollmentType, bool)`

GetAllowedEnrollmentTypesOk returns a tuple with the AllowedEnrollmentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedEnrollmentTypes

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) SetAllowedEnrollmentTypes(v CSSCMSCoreEnumsEnrollmentType)`

SetAllowedEnrollmentTypes sets AllowedEnrollmentTypes field to given value.

### HasAllowedEnrollmentTypes

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) HasAllowedEnrollmentTypes() bool`

HasAllowedEnrollmentTypes returns a boolean if a field has been set.

### GetKeyRetention

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetKeyRetention() CSSCMSCoreEnumsKeyRetentionPolicy`

GetKeyRetention returns the KeyRetention field if non-nil, zero value otherwise.

### GetKeyRetentionOk

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetKeyRetentionOk() (*CSSCMSCoreEnumsKeyRetentionPolicy, bool)`

GetKeyRetentionOk returns a tuple with the KeyRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyRetention

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) SetKeyRetention(v CSSCMSCoreEnumsKeyRetentionPolicy)`

SetKeyRetention sets KeyRetention field to given value.

### HasKeyRetention

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) HasKeyRetention() bool`

HasKeyRetention returns a boolean if a field has been set.

### GetKeyRetentionDays

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetKeyRetentionDays() int32`

GetKeyRetentionDays returns the KeyRetentionDays field if non-nil, zero value otherwise.

### GetKeyRetentionDaysOk

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetKeyRetentionDaysOk() (*int32, bool)`

GetKeyRetentionDaysOk returns a tuple with the KeyRetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyRetentionDays

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) SetKeyRetentionDays(v int32)`

SetKeyRetentionDays sets KeyRetentionDays field to given value.

### HasKeyRetentionDays

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) HasKeyRetentionDays() bool`

HasKeyRetentionDays returns a boolean if a field has been set.

### SetKeyRetentionDaysNil

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) SetKeyRetentionDaysNil(b bool)`

 SetKeyRetentionDaysNil sets the value for KeyRetentionDays to be an explicit nil

### UnsetKeyRetentionDays
`func (o *CertificateAuthoritiesCertificateAuthorityRequest) UnsetKeyRetentionDays()`

UnsetKeyRetentionDays ensures that no value is present for KeyRetentionDays, not even an explicit nil
### GetExplicitCredentials

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetExplicitCredentials() bool`

GetExplicitCredentials returns the ExplicitCredentials field if non-nil, zero value otherwise.

### GetExplicitCredentialsOk

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetExplicitCredentialsOk() (*bool, bool)`

GetExplicitCredentialsOk returns a tuple with the ExplicitCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicitCredentials

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) SetExplicitCredentials(v bool)`

SetExplicitCredentials sets ExplicitCredentials field to given value.

### HasExplicitCredentials

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) HasExplicitCredentials() bool`

HasExplicitCredentials returns a boolean if a field has been set.

### GetSubscriberTerms

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetSubscriberTerms() bool`

GetSubscriberTerms returns the SubscriberTerms field if non-nil, zero value otherwise.

### GetSubscriberTermsOk

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetSubscriberTermsOk() (*bool, bool)`

GetSubscriberTermsOk returns a tuple with the SubscriberTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriberTerms

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) SetSubscriberTerms(v bool)`

SetSubscriberTerms sets SubscriberTerms field to given value.

### HasSubscriberTerms

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) HasSubscriberTerms() bool`

HasSubscriberTerms returns a boolean if a field has been set.

### GetExplicitUser

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetExplicitUser() string`

GetExplicitUser returns the ExplicitUser field if non-nil, zero value otherwise.

### GetExplicitUserOk

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetExplicitUserOk() (*string, bool)`

GetExplicitUserOk returns a tuple with the ExplicitUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicitUser

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) SetExplicitUser(v string)`

SetExplicitUser sets ExplicitUser field to given value.

### HasExplicitUser

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) HasExplicitUser() bool`

HasExplicitUser returns a boolean if a field has been set.

### SetExplicitUserNil

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) SetExplicitUserNil(b bool)`

 SetExplicitUserNil sets the value for ExplicitUser to be an explicit nil

### UnsetExplicitUser
`func (o *CertificateAuthoritiesCertificateAuthorityRequest) UnsetExplicitUser()`

UnsetExplicitUser ensures that no value is present for ExplicitUser, not even an explicit nil
### GetExplicitPassword

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetExplicitPassword() CSSCMSDataModelModelsKeyfactorAPISecret`

GetExplicitPassword returns the ExplicitPassword field if non-nil, zero value otherwise.

### GetExplicitPasswordOk

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetExplicitPasswordOk() (*CSSCMSDataModelModelsKeyfactorAPISecret, bool)`

GetExplicitPasswordOk returns a tuple with the ExplicitPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicitPassword

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) SetExplicitPassword(v CSSCMSDataModelModelsKeyfactorAPISecret)`

SetExplicitPassword sets ExplicitPassword field to given value.

### HasExplicitPassword

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) HasExplicitPassword() bool`

HasExplicitPassword returns a boolean if a field has been set.

### GetUseAllowedRequesters

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetUseAllowedRequesters() bool`

GetUseAllowedRequesters returns the UseAllowedRequesters field if non-nil, zero value otherwise.

### GetUseAllowedRequestersOk

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetUseAllowedRequestersOk() (*bool, bool)`

GetUseAllowedRequestersOk returns a tuple with the UseAllowedRequesters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAllowedRequesters

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) SetUseAllowedRequesters(v bool)`

SetUseAllowedRequesters sets UseAllowedRequesters field to given value.

### HasUseAllowedRequesters

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) HasUseAllowedRequesters() bool`

HasUseAllowedRequesters returns a boolean if a field has been set.

### GetAllowedRequesters

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetAllowedRequesters() []string`

GetAllowedRequesters returns the AllowedRequesters field if non-nil, zero value otherwise.

### GetAllowedRequestersOk

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetAllowedRequestersOk() (*[]string, bool)`

GetAllowedRequestersOk returns a tuple with the AllowedRequesters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedRequesters

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) SetAllowedRequesters(v []string)`

SetAllowedRequesters sets AllowedRequesters field to given value.

### HasAllowedRequesters

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) HasAllowedRequesters() bool`

HasAllowedRequesters returns a boolean if a field has been set.

### SetAllowedRequestersNil

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) SetAllowedRequestersNil(b bool)`

 SetAllowedRequestersNil sets the value for AllowedRequesters to be an explicit nil

### UnsetAllowedRequesters
`func (o *CertificateAuthoritiesCertificateAuthorityRequest) UnsetAllowedRequesters()`

UnsetAllowedRequesters ensures that no value is present for AllowedRequesters, not even an explicit nil
### GetFullScan

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetFullScan() KeyfactorCommonSchedulingKeyfactorSchedule`

GetFullScan returns the FullScan field if non-nil, zero value otherwise.

### GetFullScanOk

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetFullScanOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetFullScanOk returns a tuple with the FullScan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullScan

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) SetFullScan(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetFullScan sets FullScan field to given value.

### HasFullScan

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) HasFullScan() bool`

HasFullScan returns a boolean if a field has been set.

### GetIncrementalScan

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetIncrementalScan() KeyfactorCommonSchedulingKeyfactorSchedule`

GetIncrementalScan returns the IncrementalScan field if non-nil, zero value otherwise.

### GetIncrementalScanOk

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetIncrementalScanOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetIncrementalScanOk returns a tuple with the IncrementalScan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncrementalScan

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) SetIncrementalScan(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetIncrementalScan sets IncrementalScan field to given value.

### HasIncrementalScan

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) HasIncrementalScan() bool`

HasIncrementalScan returns a boolean if a field has been set.

### GetThresholdCheck

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetThresholdCheck() KeyfactorCommonSchedulingKeyfactorSchedule`

GetThresholdCheck returns the ThresholdCheck field if non-nil, zero value otherwise.

### GetThresholdCheckOk

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetThresholdCheckOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetThresholdCheckOk returns a tuple with the ThresholdCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdCheck

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) SetThresholdCheck(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetThresholdCheck sets ThresholdCheck field to given value.

### HasThresholdCheck

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) HasThresholdCheck() bool`

HasThresholdCheck returns a boolean if a field has been set.

### GetAuthCertificatePassword

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetAuthCertificatePassword() map[string]interface{}`

GetAuthCertificatePassword returns the AuthCertificatePassword field if non-nil, zero value otherwise.

### GetAuthCertificatePasswordOk

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetAuthCertificatePasswordOk() (*map[string]interface{}, bool)`

GetAuthCertificatePasswordOk returns a tuple with the AuthCertificatePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCertificatePassword

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) SetAuthCertificatePassword(v map[string]interface{})`

SetAuthCertificatePassword sets AuthCertificatePassword field to given value.

### HasAuthCertificatePassword

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) HasAuthCertificatePassword() bool`

HasAuthCertificatePassword returns a boolean if a field has been set.

### GetAuthCertificate

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetAuthCertificate() map[string]interface{}`

GetAuthCertificate returns the AuthCertificate field if non-nil, zero value otherwise.

### GetAuthCertificateOk

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetAuthCertificateOk() (*map[string]interface{}, bool)`

GetAuthCertificateOk returns a tuple with the AuthCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCertificate

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) SetAuthCertificate(v map[string]interface{})`

SetAuthCertificate sets AuthCertificate field to given value.

### HasAuthCertificate

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) HasAuthCertificate() bool`

HasAuthCertificate returns a boolean if a field has been set.

### GetCAType

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetCAType() CSSCMSCoreEnumsCertificateAuthorityType`

GetCAType returns the CAType field if non-nil, zero value otherwise.

### GetCATypeOk

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetCATypeOk() (*CSSCMSCoreEnumsCertificateAuthorityType, bool)`

GetCATypeOk returns a tuple with the CAType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCAType

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) SetCAType(v CSSCMSCoreEnumsCertificateAuthorityType)`

SetCAType sets CAType field to given value.

### HasCAType

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) HasCAType() bool`

HasCAType returns a boolean if a field has been set.

### GetEnforceUniqueDN

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetEnforceUniqueDN() bool`

GetEnforceUniqueDN returns the EnforceUniqueDN field if non-nil, zero value otherwise.

### GetEnforceUniqueDNOk

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetEnforceUniqueDNOk() (*bool, bool)`

GetEnforceUniqueDNOk returns a tuple with the EnforceUniqueDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceUniqueDN

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) SetEnforceUniqueDN(v bool)`

SetEnforceUniqueDN sets EnforceUniqueDN field to given value.

### HasEnforceUniqueDN

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) HasEnforceUniqueDN() bool`

HasEnforceUniqueDN returns a boolean if a field has been set.

### GetAllowOneClickRenewals

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetAllowOneClickRenewals() bool`

GetAllowOneClickRenewals returns the AllowOneClickRenewals field if non-nil, zero value otherwise.

### GetAllowOneClickRenewalsOk

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetAllowOneClickRenewalsOk() (*bool, bool)`

GetAllowOneClickRenewalsOk returns a tuple with the AllowOneClickRenewals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowOneClickRenewals

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) SetAllowOneClickRenewals(v bool)`

SetAllowOneClickRenewals sets AllowOneClickRenewals field to given value.

### HasAllowOneClickRenewals

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) HasAllowOneClickRenewals() bool`

HasAllowOneClickRenewals returns a boolean if a field has been set.

### GetNewEndEntityOnRenewAndReissue

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetNewEndEntityOnRenewAndReissue() bool`

GetNewEndEntityOnRenewAndReissue returns the NewEndEntityOnRenewAndReissue field if non-nil, zero value otherwise.

### GetNewEndEntityOnRenewAndReissueOk

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetNewEndEntityOnRenewAndReissueOk() (*bool, bool)`

GetNewEndEntityOnRenewAndReissueOk returns a tuple with the NewEndEntityOnRenewAndReissue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewEndEntityOnRenewAndReissue

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) SetNewEndEntityOnRenewAndReissue(v bool)`

SetNewEndEntityOnRenewAndReissue sets NewEndEntityOnRenewAndReissue field to given value.

### HasNewEndEntityOnRenewAndReissue

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) HasNewEndEntityOnRenewAndReissue() bool`

HasNewEndEntityOnRenewAndReissue returns a boolean if a field has been set.

### GetTokenURL

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetTokenURL() string`

GetTokenURL returns the TokenURL field if non-nil, zero value otherwise.

### GetTokenURLOk

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetTokenURLOk() (*string, bool)`

GetTokenURLOk returns a tuple with the TokenURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenURL

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) SetTokenURL(v string)`

SetTokenURL sets TokenURL field to given value.

### HasTokenURL

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) HasTokenURL() bool`

HasTokenURL returns a boolean if a field has been set.

### SetTokenURLNil

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) SetTokenURLNil(b bool)`

 SetTokenURLNil sets the value for TokenURL to be an explicit nil

### UnsetTokenURL
`func (o *CertificateAuthoritiesCertificateAuthorityRequest) UnsetTokenURL()`

UnsetTokenURL ensures that no value is present for TokenURL, not even an explicit nil
### GetClientId

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### SetClientIdNil

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) SetClientIdNil(b bool)`

 SetClientIdNil sets the value for ClientId to be an explicit nil

### UnsetClientId
`func (o *CertificateAuthoritiesCertificateAuthorityRequest) UnsetClientId()`

UnsetClientId ensures that no value is present for ClientId, not even an explicit nil
### GetClientSecret

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetClientSecret() CSSCMSDataModelModelsKeyfactorAPISecret`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetClientSecretOk() (*CSSCMSDataModelModelsKeyfactorAPISecret, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) SetClientSecret(v CSSCMSDataModelModelsKeyfactorAPISecret)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetScope

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) HasScope() bool`

HasScope returns a boolean if a field has been set.

### SetScopeNil

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) SetScopeNil(b bool)`

 SetScopeNil sets the value for Scope to be an explicit nil

### UnsetScope
`func (o *CertificateAuthoritiesCertificateAuthorityRequest) UnsetScope()`

UnsetScope ensures that no value is present for Scope, not even an explicit nil
### GetAudience

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) SetAudience(v string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### SetAudienceNil

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) SetAudienceNil(b bool)`

 SetAudienceNil sets the value for Audience to be an explicit nil

### UnsetAudience
`func (o *CertificateAuthoritiesCertificateAuthorityRequest) UnsetAudience()`

UnsetAudience ensures that no value is present for Audience, not even an explicit nil
### GetCertificateCleanupEnabled

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetCertificateCleanupEnabled() bool`

GetCertificateCleanupEnabled returns the CertificateCleanupEnabled field if non-nil, zero value otherwise.

### GetCertificateCleanupEnabledOk

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetCertificateCleanupEnabledOk() (*bool, bool)`

GetCertificateCleanupEnabledOk returns a tuple with the CertificateCleanupEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateCleanupEnabled

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) SetCertificateCleanupEnabled(v bool)`

SetCertificateCleanupEnabled sets CertificateCleanupEnabled field to given value.

### HasCertificateCleanupEnabled

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) HasCertificateCleanupEnabled() bool`

HasCertificateCleanupEnabled returns a boolean if a field has been set.

### SetCertificateCleanupEnabledNil

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) SetCertificateCleanupEnabledNil(b bool)`

 SetCertificateCleanupEnabledNil sets the value for CertificateCleanupEnabled to be an explicit nil

### UnsetCertificateCleanupEnabled
`func (o *CertificateAuthoritiesCertificateAuthorityRequest) UnsetCertificateCleanupEnabled()`

UnsetCertificateCleanupEnabled ensures that no value is present for CertificateCleanupEnabled, not even an explicit nil
### GetTimeAfterExpiration

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetTimeAfterExpiration() int32`

GetTimeAfterExpiration returns the TimeAfterExpiration field if non-nil, zero value otherwise.

### GetTimeAfterExpirationOk

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetTimeAfterExpirationOk() (*int32, bool)`

GetTimeAfterExpirationOk returns a tuple with the TimeAfterExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeAfterExpiration

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) SetTimeAfterExpiration(v int32)`

SetTimeAfterExpiration sets TimeAfterExpiration field to given value.

### HasTimeAfterExpiration

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) HasTimeAfterExpiration() bool`

HasTimeAfterExpiration returns a boolean if a field has been set.

### SetTimeAfterExpirationNil

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) SetTimeAfterExpirationNil(b bool)`

 SetTimeAfterExpirationNil sets the value for TimeAfterExpiration to be an explicit nil

### UnsetTimeAfterExpiration
`func (o *CertificateAuthoritiesCertificateAuthorityRequest) UnsetTimeAfterExpiration()`

UnsetTimeAfterExpiration ensures that no value is present for TimeAfterExpiration, not even an explicit nil
### GetTimeAfterExpirationUnits

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetTimeAfterExpirationUnits() CSSCMSDataModelEnumsCertificateCleanupTimeUnits`

GetTimeAfterExpirationUnits returns the TimeAfterExpirationUnits field if non-nil, zero value otherwise.

### GetTimeAfterExpirationUnitsOk

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetTimeAfterExpirationUnitsOk() (*CSSCMSDataModelEnumsCertificateCleanupTimeUnits, bool)`

GetTimeAfterExpirationUnitsOk returns a tuple with the TimeAfterExpirationUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeAfterExpirationUnits

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) SetTimeAfterExpirationUnits(v CSSCMSDataModelEnumsCertificateCleanupTimeUnits)`

SetTimeAfterExpirationUnits sets TimeAfterExpirationUnits field to given value.

### HasTimeAfterExpirationUnits

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) HasTimeAfterExpirationUnits() bool`

HasTimeAfterExpirationUnits returns a boolean if a field has been set.

### GetDeleteWithArchivedKey

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetDeleteWithArchivedKey() bool`

GetDeleteWithArchivedKey returns the DeleteWithArchivedKey field if non-nil, zero value otherwise.

### GetDeleteWithArchivedKeyOk

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetDeleteWithArchivedKeyOk() (*bool, bool)`

GetDeleteWithArchivedKeyOk returns a tuple with the DeleteWithArchivedKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteWithArchivedKey

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) SetDeleteWithArchivedKey(v bool)`

SetDeleteWithArchivedKey sets DeleteWithArchivedKey field to given value.

### HasDeleteWithArchivedKey

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) HasDeleteWithArchivedKey() bool`

HasDeleteWithArchivedKey returns a boolean if a field has been set.

### SetDeleteWithArchivedKeyNil

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) SetDeleteWithArchivedKeyNil(b bool)`

 SetDeleteWithArchivedKeyNil sets the value for DeleteWithArchivedKey to be an explicit nil

### UnsetDeleteWithArchivedKey
`func (o *CertificateAuthoritiesCertificateAuthorityRequest) UnsetDeleteWithArchivedKey()`

UnsetDeleteWithArchivedKey ensures that no value is present for DeleteWithArchivedKey, not even an explicit nil
### GetUseForEnrollment

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetUseForEnrollment() bool`

GetUseForEnrollment returns the UseForEnrollment field if non-nil, zero value otherwise.

### GetUseForEnrollmentOk

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) GetUseForEnrollmentOk() (*bool, bool)`

GetUseForEnrollmentOk returns a tuple with the UseForEnrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseForEnrollment

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) SetUseForEnrollment(v bool)`

SetUseForEnrollment sets UseForEnrollment field to given value.

### HasUseForEnrollment

`func (o *CertificateAuthoritiesCertificateAuthorityRequest) HasUseForEnrollment() bool`

HasUseForEnrollment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


