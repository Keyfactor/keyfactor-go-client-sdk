# CertificateAuthoritiesCertificateAuthorityResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**LogicalName** | Pointer to **NullableString** |  | [optional] 
**HostName** | Pointer to **NullableString** |  | [optional] 
**Delegate** | Pointer to **bool** |  | [optional] 
**UseCAConnector** | Pointer to **bool** |  | [optional] 
**ConnectorPool** | Pointer to **NullableString** |  | [optional] 
**DelegateEnrollment** | Pointer to **bool** |  | [optional] 
**ForestRoot** | Pointer to **NullableString** |  | [optional] [readonly] 
**ConfigurationTenant** | Pointer to **NullableString** |  | [optional] 
**Remote** | Pointer to **bool** |  | [optional] 
**Agent** | Pointer to **NullableString** |  | [optional] 
**AgentName** | Pointer to **NullableString** |  | [optional] 
**AgentUsername** | Pointer to **NullableString** |  | [optional] 
**Standalone** | Pointer to **bool** |  | [optional] 
**MonitorThresholds** | Pointer to **bool** |  | [optional] 
**IssuanceMax** | Pointer to **NullableInt32** |  | [optional] 
**IssuanceMin** | Pointer to **NullableInt32** |  | [optional] 
**DenialMax** | Pointer to **NullableInt32** |  | [optional] 
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
**CAType** | Pointer to [**CSSCMSCoreEnumsCertificateAuthorityType**](CSSCMSCoreEnumsCertificateAuthorityType.md) |  | [optional] 
**AuthCertificate** | Pointer to [**CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityAuthCertificate**](CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityAuthCertificate.md) |  | [optional] 
**EnforceUniqueDN** | Pointer to **bool** |  | [optional] 
**AllowOneClickRenewals** | Pointer to **bool** |  | [optional] 
**NewEndEntityOnRenewAndReissue** | Pointer to **bool** |  | [optional] 
**TokenURL** | Pointer to **NullableString** |  | [optional] 
**ClientId** | Pointer to **NullableString** |  | [optional] 
**Scope** | Pointer to **NullableString** |  | [optional] 
**Audience** | Pointer to **NullableString** |  | [optional] 
**ClientSecret** | Pointer to [**CSSCMSDataModelModelsKeyfactorAPISecret**](CSSCMSDataModelModelsKeyfactorAPISecret.md) |  | [optional] 
**LastScan** | Pointer to **NullableString** |  | [optional] [readonly] 

## Methods

### NewCertificateAuthoritiesCertificateAuthorityResponse

`func NewCertificateAuthoritiesCertificateAuthorityResponse() *CertificateAuthoritiesCertificateAuthorityResponse`

NewCertificateAuthoritiesCertificateAuthorityResponse instantiates a new CertificateAuthoritiesCertificateAuthorityResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateAuthoritiesCertificateAuthorityResponseWithDefaults

`func NewCertificateAuthoritiesCertificateAuthorityResponseWithDefaults() *CertificateAuthoritiesCertificateAuthorityResponse`

NewCertificateAuthoritiesCertificateAuthorityResponseWithDefaults instantiates a new CertificateAuthoritiesCertificateAuthorityResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLogicalName

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetLogicalName() string`

GetLogicalName returns the LogicalName field if non-nil, zero value otherwise.

### GetLogicalNameOk

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetLogicalNameOk() (*string, bool)`

GetLogicalNameOk returns a tuple with the LogicalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalName

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) SetLogicalName(v string)`

SetLogicalName sets LogicalName field to given value.

### HasLogicalName

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) HasLogicalName() bool`

HasLogicalName returns a boolean if a field has been set.

### SetLogicalNameNil

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) SetLogicalNameNil(b bool)`

 SetLogicalNameNil sets the value for LogicalName to be an explicit nil

### UnsetLogicalName
`func (o *CertificateAuthoritiesCertificateAuthorityResponse) UnsetLogicalName()`

UnsetLogicalName ensures that no value is present for LogicalName, not even an explicit nil
### GetHostName

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### SetHostNameNil

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) SetHostNameNil(b bool)`

 SetHostNameNil sets the value for HostName to be an explicit nil

### UnsetHostName
`func (o *CertificateAuthoritiesCertificateAuthorityResponse) UnsetHostName()`

UnsetHostName ensures that no value is present for HostName, not even an explicit nil
### GetDelegate

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetDelegate() bool`

GetDelegate returns the Delegate field if non-nil, zero value otherwise.

### GetDelegateOk

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetDelegateOk() (*bool, bool)`

GetDelegateOk returns a tuple with the Delegate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegate

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) SetDelegate(v bool)`

SetDelegate sets Delegate field to given value.

### HasDelegate

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) HasDelegate() bool`

HasDelegate returns a boolean if a field has been set.

### GetUseCAConnector

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetUseCAConnector() bool`

GetUseCAConnector returns the UseCAConnector field if non-nil, zero value otherwise.

### GetUseCAConnectorOk

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetUseCAConnectorOk() (*bool, bool)`

GetUseCAConnectorOk returns a tuple with the UseCAConnector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCAConnector

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) SetUseCAConnector(v bool)`

SetUseCAConnector sets UseCAConnector field to given value.

### HasUseCAConnector

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) HasUseCAConnector() bool`

HasUseCAConnector returns a boolean if a field has been set.

### GetConnectorPool

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetConnectorPool() string`

GetConnectorPool returns the ConnectorPool field if non-nil, zero value otherwise.

### GetConnectorPoolOk

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetConnectorPoolOk() (*string, bool)`

GetConnectorPoolOk returns a tuple with the ConnectorPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorPool

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) SetConnectorPool(v string)`

SetConnectorPool sets ConnectorPool field to given value.

### HasConnectorPool

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) HasConnectorPool() bool`

HasConnectorPool returns a boolean if a field has been set.

### SetConnectorPoolNil

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) SetConnectorPoolNil(b bool)`

 SetConnectorPoolNil sets the value for ConnectorPool to be an explicit nil

### UnsetConnectorPool
`func (o *CertificateAuthoritiesCertificateAuthorityResponse) UnsetConnectorPool()`

UnsetConnectorPool ensures that no value is present for ConnectorPool, not even an explicit nil
### GetDelegateEnrollment

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetDelegateEnrollment() bool`

GetDelegateEnrollment returns the DelegateEnrollment field if non-nil, zero value otherwise.

### GetDelegateEnrollmentOk

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetDelegateEnrollmentOk() (*bool, bool)`

GetDelegateEnrollmentOk returns a tuple with the DelegateEnrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegateEnrollment

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) SetDelegateEnrollment(v bool)`

SetDelegateEnrollment sets DelegateEnrollment field to given value.

### HasDelegateEnrollment

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) HasDelegateEnrollment() bool`

HasDelegateEnrollment returns a boolean if a field has been set.

### GetForestRoot

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetForestRoot() string`

GetForestRoot returns the ForestRoot field if non-nil, zero value otherwise.

### GetForestRootOk

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetForestRootOk() (*string, bool)`

GetForestRootOk returns a tuple with the ForestRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForestRoot

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) SetForestRoot(v string)`

SetForestRoot sets ForestRoot field to given value.

### HasForestRoot

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) HasForestRoot() bool`

HasForestRoot returns a boolean if a field has been set.

### SetForestRootNil

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) SetForestRootNil(b bool)`

 SetForestRootNil sets the value for ForestRoot to be an explicit nil

### UnsetForestRoot
`func (o *CertificateAuthoritiesCertificateAuthorityResponse) UnsetForestRoot()`

UnsetForestRoot ensures that no value is present for ForestRoot, not even an explicit nil
### GetConfigurationTenant

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetConfigurationTenant() string`

GetConfigurationTenant returns the ConfigurationTenant field if non-nil, zero value otherwise.

### GetConfigurationTenantOk

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetConfigurationTenantOk() (*string, bool)`

GetConfigurationTenantOk returns a tuple with the ConfigurationTenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationTenant

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) SetConfigurationTenant(v string)`

SetConfigurationTenant sets ConfigurationTenant field to given value.

### HasConfigurationTenant

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) HasConfigurationTenant() bool`

HasConfigurationTenant returns a boolean if a field has been set.

### SetConfigurationTenantNil

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) SetConfigurationTenantNil(b bool)`

 SetConfigurationTenantNil sets the value for ConfigurationTenant to be an explicit nil

### UnsetConfigurationTenant
`func (o *CertificateAuthoritiesCertificateAuthorityResponse) UnsetConfigurationTenant()`

UnsetConfigurationTenant ensures that no value is present for ConfigurationTenant, not even an explicit nil
### GetRemote

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetRemote() bool`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetRemoteOk() (*bool, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) SetRemote(v bool)`

SetRemote sets Remote field to given value.

### HasRemote

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) HasRemote() bool`

HasRemote returns a boolean if a field has been set.

### GetAgent

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetAgent() string`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetAgentOk() (*string, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) SetAgent(v string)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) HasAgent() bool`

HasAgent returns a boolean if a field has been set.

### SetAgentNil

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) SetAgentNil(b bool)`

 SetAgentNil sets the value for Agent to be an explicit nil

### UnsetAgent
`func (o *CertificateAuthoritiesCertificateAuthorityResponse) UnsetAgent()`

UnsetAgent ensures that no value is present for Agent, not even an explicit nil
### GetAgentName

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetAgentName() string`

GetAgentName returns the AgentName field if non-nil, zero value otherwise.

### GetAgentNameOk

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetAgentNameOk() (*string, bool)`

GetAgentNameOk returns a tuple with the AgentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentName

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) SetAgentName(v string)`

SetAgentName sets AgentName field to given value.

### HasAgentName

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) HasAgentName() bool`

HasAgentName returns a boolean if a field has been set.

### SetAgentNameNil

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) SetAgentNameNil(b bool)`

 SetAgentNameNil sets the value for AgentName to be an explicit nil

### UnsetAgentName
`func (o *CertificateAuthoritiesCertificateAuthorityResponse) UnsetAgentName()`

UnsetAgentName ensures that no value is present for AgentName, not even an explicit nil
### GetAgentUsername

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetAgentUsername() string`

GetAgentUsername returns the AgentUsername field if non-nil, zero value otherwise.

### GetAgentUsernameOk

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetAgentUsernameOk() (*string, bool)`

GetAgentUsernameOk returns a tuple with the AgentUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentUsername

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) SetAgentUsername(v string)`

SetAgentUsername sets AgentUsername field to given value.

### HasAgentUsername

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) HasAgentUsername() bool`

HasAgentUsername returns a boolean if a field has been set.

### SetAgentUsernameNil

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) SetAgentUsernameNil(b bool)`

 SetAgentUsernameNil sets the value for AgentUsername to be an explicit nil

### UnsetAgentUsername
`func (o *CertificateAuthoritiesCertificateAuthorityResponse) UnsetAgentUsername()`

UnsetAgentUsername ensures that no value is present for AgentUsername, not even an explicit nil
### GetStandalone

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetStandalone() bool`

GetStandalone returns the Standalone field if non-nil, zero value otherwise.

### GetStandaloneOk

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetStandaloneOk() (*bool, bool)`

GetStandaloneOk returns a tuple with the Standalone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandalone

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) SetStandalone(v bool)`

SetStandalone sets Standalone field to given value.

### HasStandalone

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) HasStandalone() bool`

HasStandalone returns a boolean if a field has been set.

### GetMonitorThresholds

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetMonitorThresholds() bool`

GetMonitorThresholds returns the MonitorThresholds field if non-nil, zero value otherwise.

### GetMonitorThresholdsOk

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetMonitorThresholdsOk() (*bool, bool)`

GetMonitorThresholdsOk returns a tuple with the MonitorThresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorThresholds

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) SetMonitorThresholds(v bool)`

SetMonitorThresholds sets MonitorThresholds field to given value.

### HasMonitorThresholds

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) HasMonitorThresholds() bool`

HasMonitorThresholds returns a boolean if a field has been set.

### GetIssuanceMax

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetIssuanceMax() int32`

GetIssuanceMax returns the IssuanceMax field if non-nil, zero value otherwise.

### GetIssuanceMaxOk

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetIssuanceMaxOk() (*int32, bool)`

GetIssuanceMaxOk returns a tuple with the IssuanceMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuanceMax

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) SetIssuanceMax(v int32)`

SetIssuanceMax sets IssuanceMax field to given value.

### HasIssuanceMax

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) HasIssuanceMax() bool`

HasIssuanceMax returns a boolean if a field has been set.

### SetIssuanceMaxNil

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) SetIssuanceMaxNil(b bool)`

 SetIssuanceMaxNil sets the value for IssuanceMax to be an explicit nil

### UnsetIssuanceMax
`func (o *CertificateAuthoritiesCertificateAuthorityResponse) UnsetIssuanceMax()`

UnsetIssuanceMax ensures that no value is present for IssuanceMax, not even an explicit nil
### GetIssuanceMin

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetIssuanceMin() int32`

GetIssuanceMin returns the IssuanceMin field if non-nil, zero value otherwise.

### GetIssuanceMinOk

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetIssuanceMinOk() (*int32, bool)`

GetIssuanceMinOk returns a tuple with the IssuanceMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuanceMin

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) SetIssuanceMin(v int32)`

SetIssuanceMin sets IssuanceMin field to given value.

### HasIssuanceMin

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) HasIssuanceMin() bool`

HasIssuanceMin returns a boolean if a field has been set.

### SetIssuanceMinNil

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) SetIssuanceMinNil(b bool)`

 SetIssuanceMinNil sets the value for IssuanceMin to be an explicit nil

### UnsetIssuanceMin
`func (o *CertificateAuthoritiesCertificateAuthorityResponse) UnsetIssuanceMin()`

UnsetIssuanceMin ensures that no value is present for IssuanceMin, not even an explicit nil
### GetDenialMax

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetDenialMax() int32`

GetDenialMax returns the DenialMax field if non-nil, zero value otherwise.

### GetDenialMaxOk

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetDenialMaxOk() (*int32, bool)`

GetDenialMaxOk returns a tuple with the DenialMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenialMax

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) SetDenialMax(v int32)`

SetDenialMax sets DenialMax field to given value.

### HasDenialMax

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) HasDenialMax() bool`

HasDenialMax returns a boolean if a field has been set.

### SetDenialMaxNil

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) SetDenialMaxNil(b bool)`

 SetDenialMaxNil sets the value for DenialMax to be an explicit nil

### UnsetDenialMax
`func (o *CertificateAuthoritiesCertificateAuthorityResponse) UnsetDenialMax()`

UnsetDenialMax ensures that no value is present for DenialMax, not even an explicit nil
### GetFailureMax

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetFailureMax() int32`

GetFailureMax returns the FailureMax field if non-nil, zero value otherwise.

### GetFailureMaxOk

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetFailureMaxOk() (*int32, bool)`

GetFailureMaxOk returns a tuple with the FailureMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureMax

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) SetFailureMax(v int32)`

SetFailureMax sets FailureMax field to given value.

### HasFailureMax

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) HasFailureMax() bool`

HasFailureMax returns a boolean if a field has been set.

### SetFailureMaxNil

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) SetFailureMaxNil(b bool)`

 SetFailureMaxNil sets the value for FailureMax to be an explicit nil

### UnsetFailureMax
`func (o *CertificateAuthoritiesCertificateAuthorityResponse) UnsetFailureMax()`

UnsetFailureMax ensures that no value is present for FailureMax, not even an explicit nil
### GetRFCEnforcement

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetRFCEnforcement() bool`

GetRFCEnforcement returns the RFCEnforcement field if non-nil, zero value otherwise.

### GetRFCEnforcementOk

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetRFCEnforcementOk() (*bool, bool)`

GetRFCEnforcementOk returns a tuple with the RFCEnforcement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRFCEnforcement

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) SetRFCEnforcement(v bool)`

SetRFCEnforcement sets RFCEnforcement field to given value.

### HasRFCEnforcement

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) HasRFCEnforcement() bool`

HasRFCEnforcement returns a boolean if a field has been set.

### GetProperties

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetProperties() string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetPropertiesOk() (*string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) SetProperties(v string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *CertificateAuthoritiesCertificateAuthorityResponse) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetAllowedEnrollmentTypes

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetAllowedEnrollmentTypes() CSSCMSCoreEnumsEnrollmentType`

GetAllowedEnrollmentTypes returns the AllowedEnrollmentTypes field if non-nil, zero value otherwise.

### GetAllowedEnrollmentTypesOk

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetAllowedEnrollmentTypesOk() (*CSSCMSCoreEnumsEnrollmentType, bool)`

GetAllowedEnrollmentTypesOk returns a tuple with the AllowedEnrollmentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedEnrollmentTypes

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) SetAllowedEnrollmentTypes(v CSSCMSCoreEnumsEnrollmentType)`

SetAllowedEnrollmentTypes sets AllowedEnrollmentTypes field to given value.

### HasAllowedEnrollmentTypes

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) HasAllowedEnrollmentTypes() bool`

HasAllowedEnrollmentTypes returns a boolean if a field has been set.

### GetKeyRetention

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetKeyRetention() CSSCMSCoreEnumsKeyRetentionPolicy`

GetKeyRetention returns the KeyRetention field if non-nil, zero value otherwise.

### GetKeyRetentionOk

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetKeyRetentionOk() (*CSSCMSCoreEnumsKeyRetentionPolicy, bool)`

GetKeyRetentionOk returns a tuple with the KeyRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyRetention

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) SetKeyRetention(v CSSCMSCoreEnumsKeyRetentionPolicy)`

SetKeyRetention sets KeyRetention field to given value.

### HasKeyRetention

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) HasKeyRetention() bool`

HasKeyRetention returns a boolean if a field has been set.

### GetKeyRetentionDays

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetKeyRetentionDays() int32`

GetKeyRetentionDays returns the KeyRetentionDays field if non-nil, zero value otherwise.

### GetKeyRetentionDaysOk

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetKeyRetentionDaysOk() (*int32, bool)`

GetKeyRetentionDaysOk returns a tuple with the KeyRetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyRetentionDays

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) SetKeyRetentionDays(v int32)`

SetKeyRetentionDays sets KeyRetentionDays field to given value.

### HasKeyRetentionDays

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) HasKeyRetentionDays() bool`

HasKeyRetentionDays returns a boolean if a field has been set.

### SetKeyRetentionDaysNil

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) SetKeyRetentionDaysNil(b bool)`

 SetKeyRetentionDaysNil sets the value for KeyRetentionDays to be an explicit nil

### UnsetKeyRetentionDays
`func (o *CertificateAuthoritiesCertificateAuthorityResponse) UnsetKeyRetentionDays()`

UnsetKeyRetentionDays ensures that no value is present for KeyRetentionDays, not even an explicit nil
### GetExplicitCredentials

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetExplicitCredentials() bool`

GetExplicitCredentials returns the ExplicitCredentials field if non-nil, zero value otherwise.

### GetExplicitCredentialsOk

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetExplicitCredentialsOk() (*bool, bool)`

GetExplicitCredentialsOk returns a tuple with the ExplicitCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicitCredentials

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) SetExplicitCredentials(v bool)`

SetExplicitCredentials sets ExplicitCredentials field to given value.

### HasExplicitCredentials

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) HasExplicitCredentials() bool`

HasExplicitCredentials returns a boolean if a field has been set.

### GetSubscriberTerms

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetSubscriberTerms() bool`

GetSubscriberTerms returns the SubscriberTerms field if non-nil, zero value otherwise.

### GetSubscriberTermsOk

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetSubscriberTermsOk() (*bool, bool)`

GetSubscriberTermsOk returns a tuple with the SubscriberTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriberTerms

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) SetSubscriberTerms(v bool)`

SetSubscriberTerms sets SubscriberTerms field to given value.

### HasSubscriberTerms

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) HasSubscriberTerms() bool`

HasSubscriberTerms returns a boolean if a field has been set.

### GetExplicitUser

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetExplicitUser() string`

GetExplicitUser returns the ExplicitUser field if non-nil, zero value otherwise.

### GetExplicitUserOk

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetExplicitUserOk() (*string, bool)`

GetExplicitUserOk returns a tuple with the ExplicitUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicitUser

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) SetExplicitUser(v string)`

SetExplicitUser sets ExplicitUser field to given value.

### HasExplicitUser

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) HasExplicitUser() bool`

HasExplicitUser returns a boolean if a field has been set.

### SetExplicitUserNil

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) SetExplicitUserNil(b bool)`

 SetExplicitUserNil sets the value for ExplicitUser to be an explicit nil

### UnsetExplicitUser
`func (o *CertificateAuthoritiesCertificateAuthorityResponse) UnsetExplicitUser()`

UnsetExplicitUser ensures that no value is present for ExplicitUser, not even an explicit nil
### GetExplicitPassword

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetExplicitPassword() CSSCMSDataModelModelsKeyfactorAPISecret`

GetExplicitPassword returns the ExplicitPassword field if non-nil, zero value otherwise.

### GetExplicitPasswordOk

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetExplicitPasswordOk() (*CSSCMSDataModelModelsKeyfactorAPISecret, bool)`

GetExplicitPasswordOk returns a tuple with the ExplicitPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicitPassword

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) SetExplicitPassword(v CSSCMSDataModelModelsKeyfactorAPISecret)`

SetExplicitPassword sets ExplicitPassword field to given value.

### HasExplicitPassword

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) HasExplicitPassword() bool`

HasExplicitPassword returns a boolean if a field has been set.

### GetUseAllowedRequesters

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetUseAllowedRequesters() bool`

GetUseAllowedRequesters returns the UseAllowedRequesters field if non-nil, zero value otherwise.

### GetUseAllowedRequestersOk

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetUseAllowedRequestersOk() (*bool, bool)`

GetUseAllowedRequestersOk returns a tuple with the UseAllowedRequesters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAllowedRequesters

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) SetUseAllowedRequesters(v bool)`

SetUseAllowedRequesters sets UseAllowedRequesters field to given value.

### HasUseAllowedRequesters

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) HasUseAllowedRequesters() bool`

HasUseAllowedRequesters returns a boolean if a field has been set.

### GetAllowedRequesters

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetAllowedRequesters() []string`

GetAllowedRequesters returns the AllowedRequesters field if non-nil, zero value otherwise.

### GetAllowedRequestersOk

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetAllowedRequestersOk() (*[]string, bool)`

GetAllowedRequestersOk returns a tuple with the AllowedRequesters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedRequesters

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) SetAllowedRequesters(v []string)`

SetAllowedRequesters sets AllowedRequesters field to given value.

### HasAllowedRequesters

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) HasAllowedRequesters() bool`

HasAllowedRequesters returns a boolean if a field has been set.

### SetAllowedRequestersNil

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) SetAllowedRequestersNil(b bool)`

 SetAllowedRequestersNil sets the value for AllowedRequesters to be an explicit nil

### UnsetAllowedRequesters
`func (o *CertificateAuthoritiesCertificateAuthorityResponse) UnsetAllowedRequesters()`

UnsetAllowedRequesters ensures that no value is present for AllowedRequesters, not even an explicit nil
### GetFullScan

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetFullScan() KeyfactorCommonSchedulingKeyfactorSchedule`

GetFullScan returns the FullScan field if non-nil, zero value otherwise.

### GetFullScanOk

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetFullScanOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetFullScanOk returns a tuple with the FullScan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullScan

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) SetFullScan(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetFullScan sets FullScan field to given value.

### HasFullScan

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) HasFullScan() bool`

HasFullScan returns a boolean if a field has been set.

### GetIncrementalScan

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetIncrementalScan() KeyfactorCommonSchedulingKeyfactorSchedule`

GetIncrementalScan returns the IncrementalScan field if non-nil, zero value otherwise.

### GetIncrementalScanOk

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetIncrementalScanOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetIncrementalScanOk returns a tuple with the IncrementalScan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncrementalScan

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) SetIncrementalScan(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetIncrementalScan sets IncrementalScan field to given value.

### HasIncrementalScan

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) HasIncrementalScan() bool`

HasIncrementalScan returns a boolean if a field has been set.

### GetThresholdCheck

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetThresholdCheck() KeyfactorCommonSchedulingKeyfactorSchedule`

GetThresholdCheck returns the ThresholdCheck field if non-nil, zero value otherwise.

### GetThresholdCheckOk

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetThresholdCheckOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetThresholdCheckOk returns a tuple with the ThresholdCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdCheck

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) SetThresholdCheck(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetThresholdCheck sets ThresholdCheck field to given value.

### HasThresholdCheck

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) HasThresholdCheck() bool`

HasThresholdCheck returns a boolean if a field has been set.

### GetCAType

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetCAType() CSSCMSCoreEnumsCertificateAuthorityType`

GetCAType returns the CAType field if non-nil, zero value otherwise.

### GetCATypeOk

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetCATypeOk() (*CSSCMSCoreEnumsCertificateAuthorityType, bool)`

GetCATypeOk returns a tuple with the CAType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCAType

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) SetCAType(v CSSCMSCoreEnumsCertificateAuthorityType)`

SetCAType sets CAType field to given value.

### HasCAType

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) HasCAType() bool`

HasCAType returns a boolean if a field has been set.

### GetAuthCertificate

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetAuthCertificate() CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityAuthCertificate`

GetAuthCertificate returns the AuthCertificate field if non-nil, zero value otherwise.

### GetAuthCertificateOk

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetAuthCertificateOk() (*CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityAuthCertificate, bool)`

GetAuthCertificateOk returns a tuple with the AuthCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCertificate

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) SetAuthCertificate(v CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityAuthCertificate)`

SetAuthCertificate sets AuthCertificate field to given value.

### HasAuthCertificate

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) HasAuthCertificate() bool`

HasAuthCertificate returns a boolean if a field has been set.

### GetEnforceUniqueDN

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetEnforceUniqueDN() bool`

GetEnforceUniqueDN returns the EnforceUniqueDN field if non-nil, zero value otherwise.

### GetEnforceUniqueDNOk

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetEnforceUniqueDNOk() (*bool, bool)`

GetEnforceUniqueDNOk returns a tuple with the EnforceUniqueDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceUniqueDN

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) SetEnforceUniqueDN(v bool)`

SetEnforceUniqueDN sets EnforceUniqueDN field to given value.

### HasEnforceUniqueDN

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) HasEnforceUniqueDN() bool`

HasEnforceUniqueDN returns a boolean if a field has been set.

### GetAllowOneClickRenewals

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetAllowOneClickRenewals() bool`

GetAllowOneClickRenewals returns the AllowOneClickRenewals field if non-nil, zero value otherwise.

### GetAllowOneClickRenewalsOk

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetAllowOneClickRenewalsOk() (*bool, bool)`

GetAllowOneClickRenewalsOk returns a tuple with the AllowOneClickRenewals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowOneClickRenewals

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) SetAllowOneClickRenewals(v bool)`

SetAllowOneClickRenewals sets AllowOneClickRenewals field to given value.

### HasAllowOneClickRenewals

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) HasAllowOneClickRenewals() bool`

HasAllowOneClickRenewals returns a boolean if a field has been set.

### GetNewEndEntityOnRenewAndReissue

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetNewEndEntityOnRenewAndReissue() bool`

GetNewEndEntityOnRenewAndReissue returns the NewEndEntityOnRenewAndReissue field if non-nil, zero value otherwise.

### GetNewEndEntityOnRenewAndReissueOk

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetNewEndEntityOnRenewAndReissueOk() (*bool, bool)`

GetNewEndEntityOnRenewAndReissueOk returns a tuple with the NewEndEntityOnRenewAndReissue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewEndEntityOnRenewAndReissue

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) SetNewEndEntityOnRenewAndReissue(v bool)`

SetNewEndEntityOnRenewAndReissue sets NewEndEntityOnRenewAndReissue field to given value.

### HasNewEndEntityOnRenewAndReissue

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) HasNewEndEntityOnRenewAndReissue() bool`

HasNewEndEntityOnRenewAndReissue returns a boolean if a field has been set.

### GetTokenURL

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetTokenURL() string`

GetTokenURL returns the TokenURL field if non-nil, zero value otherwise.

### GetTokenURLOk

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetTokenURLOk() (*string, bool)`

GetTokenURLOk returns a tuple with the TokenURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenURL

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) SetTokenURL(v string)`

SetTokenURL sets TokenURL field to given value.

### HasTokenURL

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) HasTokenURL() bool`

HasTokenURL returns a boolean if a field has been set.

### SetTokenURLNil

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) SetTokenURLNil(b bool)`

 SetTokenURLNil sets the value for TokenURL to be an explicit nil

### UnsetTokenURL
`func (o *CertificateAuthoritiesCertificateAuthorityResponse) UnsetTokenURL()`

UnsetTokenURL ensures that no value is present for TokenURL, not even an explicit nil
### GetClientId

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### SetClientIdNil

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) SetClientIdNil(b bool)`

 SetClientIdNil sets the value for ClientId to be an explicit nil

### UnsetClientId
`func (o *CertificateAuthoritiesCertificateAuthorityResponse) UnsetClientId()`

UnsetClientId ensures that no value is present for ClientId, not even an explicit nil
### GetScope

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) HasScope() bool`

HasScope returns a boolean if a field has been set.

### SetScopeNil

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) SetScopeNil(b bool)`

 SetScopeNil sets the value for Scope to be an explicit nil

### UnsetScope
`func (o *CertificateAuthoritiesCertificateAuthorityResponse) UnsetScope()`

UnsetScope ensures that no value is present for Scope, not even an explicit nil
### GetAudience

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) SetAudience(v string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### SetAudienceNil

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) SetAudienceNil(b bool)`

 SetAudienceNil sets the value for Audience to be an explicit nil

### UnsetAudience
`func (o *CertificateAuthoritiesCertificateAuthorityResponse) UnsetAudience()`

UnsetAudience ensures that no value is present for Audience, not even an explicit nil
### GetClientSecret

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetClientSecret() CSSCMSDataModelModelsKeyfactorAPISecret`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetClientSecretOk() (*CSSCMSDataModelModelsKeyfactorAPISecret, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) SetClientSecret(v CSSCMSDataModelModelsKeyfactorAPISecret)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetLastScan

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetLastScan() string`

GetLastScan returns the LastScan field if non-nil, zero value otherwise.

### GetLastScanOk

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) GetLastScanOk() (*string, bool)`

GetLastScanOk returns a tuple with the LastScan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastScan

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) SetLastScan(v string)`

SetLastScan sets LastScan field to given value.

### HasLastScan

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) HasLastScan() bool`

HasLastScan returns a boolean if a field has been set.

### SetLastScanNil

`func (o *CertificateAuthoritiesCertificateAuthorityResponse) SetLastScanNil(b bool)`

 SetLastScanNil sets the value for LastScan to be an explicit nil

### UnsetLastScan
`func (o *CertificateAuthoritiesCertificateAuthorityResponse) UnsetLastScan()`

UnsetLastScan ensures that no value is present for LastScan, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


