# KeyfactorApiModelsOrchestratorsAgentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentId** | Pointer to **string** |  | [optional] 
**ClientMachine** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**AgentPlatform** | Pointer to **int32** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 
**LastSeen** | Pointer to **time.Time** |  | [optional] 
**Capabilities** | Pointer to **[]string** |  | [optional] 
**Blueprint** | Pointer to **string** |  | [optional] 
**Thumbprint** | Pointer to **string** |  | [optional] 
**LegacyThumbprint** | Pointer to **string** |  | [optional] 
**AuthCertificateReenrollment** | Pointer to **string** |  | [optional] 
**LastThumbprintUsed** | Pointer to **string** |  | [optional] 
**LastErrorCode** | Pointer to **int64** |  | [optional] 
**LastErrorMessage** | Pointer to **string** |  | [optional] 

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


