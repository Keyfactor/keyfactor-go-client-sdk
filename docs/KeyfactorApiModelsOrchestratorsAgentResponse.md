# KeyfactorApiModelsOrchestratorsAgentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentId** | Pointer to **string** | A string indicating the GUID of the orchestrator. | [optional] 
**ClientMachine** | Pointer to **string** | A string indicating the client machine on which the orchestrator is installed. | [optional] 
**Username** | Pointer to **string** | A string indicating the Active Directory user or service account the orchestrator is using to connect to Keyfactor Command. | [optional] 
**AgentPlatform** | Pointer to **int32** | An integer indicating the platform for the orchestrator. - 0 &#x3D; Unknown - 1 &#x3D; Keyfactor Windows Orchestrator - 2 &#x3D; Keyfactor Java Agent - 3 &#x3D; Keyfactor Mac Auto-Enrollment Agent - 4 &#x3D; Keyfactor Android Agent - 5 &#x3D; Keyfactor Native Agent - 6 &#x3D; Keyfactor Bash Orchestrator - 7 &#x3D; Keyfactor Universal Orchestrator  | [optional] 
**Version** | Pointer to **string** | A string indicating the version of the orchestrator. | [optional] 
**Status** | Pointer to **int32** | An integer indicating the orchestrator status: - 1 &#x3D; New - 2 &#x3D; Approved - 3 &#x3D; Disapproved  | [optional] 
**LastSeen** | Pointer to **time.Time** | The time, in UTC, at which the orchestrator last contacted Keyfactor Command. | [optional] 
**Capabilities** | Pointer to **[]string** | An array of strings indicating the capabilities reported by the orchestrator. These may be built-in or custom capabilities. | [optional] 
**Blueprint** | Pointer to **string** | A string indicating the name of the blueprint associated with the orchestrator. | [optional] 
**Thumbprint** | Pointer to **string** | A string indicating the thumbprint of the certificate that Keyfactor Command is expecting the orchestrator to use for client certificate authentication. | [optional] 
**LegacyThumbprint** | Pointer to **string** | A string indicating the thumbprint of the certificate previously used by the orchestrator for client certificate authentication before a certificate renewal operation took place (rotating the current thumbprint into the legacy thumbprint). The legacy thumbprint is cleared once the orchestrator successfully registers with the new thumbprint. | [optional] 
**AuthCertificateReenrollment** | Pointer to **string** | An integer indicating the value of the orchestrator certificate reenrollment request or require status. Possible values: - 0 &#x3D; None—Unset the value so that the orchestrator will not request a new client authentication certificate (based on this value). - 1 &#x3D; Requested—The orchestrator will request a new client authentication certificate when it next registers for a session. Orchestrator activity will be allowed to continue as usual. - 2 &#x3D; Required—The orchestrator will request a new client authentication certificate when it next registers for a session. A new session will not be granted and orchestrator activity will not be allowed to continue until the orchestrator acquires a new certificate.  | [optional] 
**LastThumbprintUsed** | Pointer to **string** | A string indicating the thumbprint of the certificate that the orchestrator most recently used for client certificate authentication. In most cases, this will match the Thumbprint. | [optional] 
**LastErrorCode** | Pointer to **int64** | An integer indicating the last error code, if any, reported from the orchestrator when trying to register a session. This code is cleared on successful session registration. | [optional] 
**LastErrorMessage** | Pointer to **string** | A string indicating the last error message, if any, reported from the orchestrator when trying to register a session. This message is cleared on successful session registration. | [optional] 

## Methods

### NewKeyfactorApiModelsOrchestratorsAgentResponse

`func NewKeyfactorApiModelsOrchestratorsAgentResponse() *KeyfactorApiModelsOrchestratorsAgentResponse`

NewKeyfactorApiModelsOrchestratorsAgentResponse instantiates a new KeyfactorApiModelsOrchestratorsAgentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorApiModelsOrchestratorsAgentResponseWithDefaults

`func NewKeyfactorApiModelsOrchestratorsAgentResponseWithDefaults() *KeyfactorApiModelsOrchestratorsAgentResponse`

NewKeyfactorApiModelsOrchestratorsAgentResponseWithDefaults instantiates a new KeyfactorApiModelsOrchestratorsAgentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentId

`func (o *KeyfactorApiModelsOrchestratorsAgentResponse) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *KeyfactorApiModelsOrchestratorsAgentResponse) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *KeyfactorApiModelsOrchestratorsAgentResponse) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *KeyfactorApiModelsOrchestratorsAgentResponse) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.

### GetClientMachine

`func (o *KeyfactorApiModelsOrchestratorsAgentResponse) GetClientMachine() string`

GetClientMachine returns the ClientMachine field if non-nil, zero value otherwise.

### GetClientMachineOk

`func (o *KeyfactorApiModelsOrchestratorsAgentResponse) GetClientMachineOk() (*string, bool)`

GetClientMachineOk returns a tuple with the ClientMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientMachine

`func (o *KeyfactorApiModelsOrchestratorsAgentResponse) SetClientMachine(v string)`

SetClientMachine sets ClientMachine field to given value.

### HasClientMachine

`func (o *KeyfactorApiModelsOrchestratorsAgentResponse) HasClientMachine() bool`

HasClientMachine returns a boolean if a field has been set.

### GetUsername

`func (o *KeyfactorApiModelsOrchestratorsAgentResponse) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *KeyfactorApiModelsOrchestratorsAgentResponse) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *KeyfactorApiModelsOrchestratorsAgentResponse) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *KeyfactorApiModelsOrchestratorsAgentResponse) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetAgentPlatform

`func (o *KeyfactorApiModelsOrchestratorsAgentResponse) GetAgentPlatform() int32`

GetAgentPlatform returns the AgentPlatform field if non-nil, zero value otherwise.

### GetAgentPlatformOk

`func (o *KeyfactorApiModelsOrchestratorsAgentResponse) GetAgentPlatformOk() (*int32, bool)`

GetAgentPlatformOk returns a tuple with the AgentPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentPlatform

`func (o *KeyfactorApiModelsOrchestratorsAgentResponse) SetAgentPlatform(v int32)`

SetAgentPlatform sets AgentPlatform field to given value.

### HasAgentPlatform

`func (o *KeyfactorApiModelsOrchestratorsAgentResponse) HasAgentPlatform() bool`

HasAgentPlatform returns a boolean if a field has been set.

### GetVersion

`func (o *KeyfactorApiModelsOrchestratorsAgentResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *KeyfactorApiModelsOrchestratorsAgentResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *KeyfactorApiModelsOrchestratorsAgentResponse) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *KeyfactorApiModelsOrchestratorsAgentResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetStatus

`func (o *KeyfactorApiModelsOrchestratorsAgentResponse) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KeyfactorApiModelsOrchestratorsAgentResponse) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KeyfactorApiModelsOrchestratorsAgentResponse) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KeyfactorApiModelsOrchestratorsAgentResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLastSeen

`func (o *KeyfactorApiModelsOrchestratorsAgentResponse) GetLastSeen() time.Time`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *KeyfactorApiModelsOrchestratorsAgentResponse) GetLastSeenOk() (*time.Time, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *KeyfactorApiModelsOrchestratorsAgentResponse) SetLastSeen(v time.Time)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *KeyfactorApiModelsOrchestratorsAgentResponse) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetCapabilities

`func (o *KeyfactorApiModelsOrchestratorsAgentResponse) GetCapabilities() []string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *KeyfactorApiModelsOrchestratorsAgentResponse) GetCapabilitiesOk() (*[]string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *KeyfactorApiModelsOrchestratorsAgentResponse) SetCapabilities(v []string)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *KeyfactorApiModelsOrchestratorsAgentResponse) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetBlueprint

`func (o *KeyfactorApiModelsOrchestratorsAgentResponse) GetBlueprint() string`

GetBlueprint returns the Blueprint field if non-nil, zero value otherwise.

### GetBlueprintOk

`func (o *KeyfactorApiModelsOrchestratorsAgentResponse) GetBlueprintOk() (*string, bool)`

GetBlueprintOk returns a tuple with the Blueprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprint

`func (o *KeyfactorApiModelsOrchestratorsAgentResponse) SetBlueprint(v string)`

SetBlueprint sets Blueprint field to given value.

### HasBlueprint

`func (o *KeyfactorApiModelsOrchestratorsAgentResponse) HasBlueprint() bool`

HasBlueprint returns a boolean if a field has been set.

### GetThumbprint

`func (o *KeyfactorApiModelsOrchestratorsAgentResponse) GetThumbprint() string`

GetThumbprint returns the Thumbprint field if non-nil, zero value otherwise.

### GetThumbprintOk

`func (o *KeyfactorApiModelsOrchestratorsAgentResponse) GetThumbprintOk() (*string, bool)`

GetThumbprintOk returns a tuple with the Thumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbprint

`func (o *KeyfactorApiModelsOrchestratorsAgentResponse) SetThumbprint(v string)`

SetThumbprint sets Thumbprint field to given value.

### HasThumbprint

`func (o *KeyfactorApiModelsOrchestratorsAgentResponse) HasThumbprint() bool`

HasThumbprint returns a boolean if a field has been set.

### GetLegacyThumbprint

`func (o *KeyfactorApiModelsOrchestratorsAgentResponse) GetLegacyThumbprint() string`

GetLegacyThumbprint returns the LegacyThumbprint field if non-nil, zero value otherwise.

### GetLegacyThumbprintOk

`func (o *KeyfactorApiModelsOrchestratorsAgentResponse) GetLegacyThumbprintOk() (*string, bool)`

GetLegacyThumbprintOk returns a tuple with the LegacyThumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacyThumbprint

`func (o *KeyfactorApiModelsOrchestratorsAgentResponse) SetLegacyThumbprint(v string)`

SetLegacyThumbprint sets LegacyThumbprint field to given value.

### HasLegacyThumbprint

`func (o *KeyfactorApiModelsOrchestratorsAgentResponse) HasLegacyThumbprint() bool`

HasLegacyThumbprint returns a boolean if a field has been set.

### GetAuthCertificateReenrollment

`func (o *KeyfactorApiModelsOrchestratorsAgentResponse) GetAuthCertificateReenrollment() string`

GetAuthCertificateReenrollment returns the AuthCertificateReenrollment field if non-nil, zero value otherwise.

### GetAuthCertificateReenrollmentOk

`func (o *KeyfactorApiModelsOrchestratorsAgentResponse) GetAuthCertificateReenrollmentOk() (*string, bool)`

GetAuthCertificateReenrollmentOk returns a tuple with the AuthCertificateReenrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCertificateReenrollment

`func (o *KeyfactorApiModelsOrchestratorsAgentResponse) SetAuthCertificateReenrollment(v string)`

SetAuthCertificateReenrollment sets AuthCertificateReenrollment field to given value.

### HasAuthCertificateReenrollment

`func (o *KeyfactorApiModelsOrchestratorsAgentResponse) HasAuthCertificateReenrollment() bool`

HasAuthCertificateReenrollment returns a boolean if a field has been set.

### GetLastThumbprintUsed

`func (o *KeyfactorApiModelsOrchestratorsAgentResponse) GetLastThumbprintUsed() string`

GetLastThumbprintUsed returns the LastThumbprintUsed field if non-nil, zero value otherwise.

### GetLastThumbprintUsedOk

`func (o *KeyfactorApiModelsOrchestratorsAgentResponse) GetLastThumbprintUsedOk() (*string, bool)`

GetLastThumbprintUsedOk returns a tuple with the LastThumbprintUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastThumbprintUsed

`func (o *KeyfactorApiModelsOrchestratorsAgentResponse) SetLastThumbprintUsed(v string)`

SetLastThumbprintUsed sets LastThumbprintUsed field to given value.

### HasLastThumbprintUsed

`func (o *KeyfactorApiModelsOrchestratorsAgentResponse) HasLastThumbprintUsed() bool`

HasLastThumbprintUsed returns a boolean if a field has been set.

### GetLastErrorCode

`func (o *KeyfactorApiModelsOrchestratorsAgentResponse) GetLastErrorCode() int64`

GetLastErrorCode returns the LastErrorCode field if non-nil, zero value otherwise.

### GetLastErrorCodeOk

`func (o *KeyfactorApiModelsOrchestratorsAgentResponse) GetLastErrorCodeOk() (*int64, bool)`

GetLastErrorCodeOk returns a tuple with the LastErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastErrorCode

`func (o *KeyfactorApiModelsOrchestratorsAgentResponse) SetLastErrorCode(v int64)`

SetLastErrorCode sets LastErrorCode field to given value.

### HasLastErrorCode

`func (o *KeyfactorApiModelsOrchestratorsAgentResponse) HasLastErrorCode() bool`

HasLastErrorCode returns a boolean if a field has been set.

### GetLastErrorMessage

`func (o *KeyfactorApiModelsOrchestratorsAgentResponse) GetLastErrorMessage() string`

GetLastErrorMessage returns the LastErrorMessage field if non-nil, zero value otherwise.

### GetLastErrorMessageOk

`func (o *KeyfactorApiModelsOrchestratorsAgentResponse) GetLastErrorMessageOk() (*string, bool)`

GetLastErrorMessageOk returns a tuple with the LastErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastErrorMessage

`func (o *KeyfactorApiModelsOrchestratorsAgentResponse) SetLastErrorMessage(v string)`

SetLastErrorMessage sets LastErrorMessage field to given value.

### HasLastErrorMessage

`func (o *KeyfactorApiModelsOrchestratorsAgentResponse) HasLastErrorMessage() bool`

HasLastErrorMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


