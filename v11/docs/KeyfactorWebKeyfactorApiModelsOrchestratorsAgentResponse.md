# KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentId** | Pointer to **string** |  | [optional] 
**ClientMachine** | Pointer to **NullableString** |  | [optional] 
**Username** | Pointer to **NullableString** |  | [optional] 
**AgentPlatform** | Pointer to **int32** |  | [optional] 
**Version** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 
**LastSeen** | Pointer to **time.Time** |  | [optional] 
**Capabilities** | Pointer to **[]string** |  | [optional] 
**Blueprint** | Pointer to **NullableString** |  | [optional] 
**Thumbprint** | Pointer to **NullableString** |  | [optional] 
**LegacyThumbprint** | Pointer to **NullableString** |  | [optional] 
**AuthCertificateReenrollment** | Pointer to **NullableString** |  | [optional] 
**LastThumbprintUsed** | Pointer to **NullableString** |  | [optional] 
**LastErrorCode** | Pointer to **NullableInt64** |  | [optional] 
**LastErrorMessage** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewKeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse

`func NewKeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse() *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse`

NewKeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse instantiates a new KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponseWithDefaults

`func NewKeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponseWithDefaults() *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse`

NewKeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponseWithDefaults instantiates a new KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentId

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.

### GetClientMachine

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) GetClientMachine() string`

GetClientMachine returns the ClientMachine field if non-nil, zero value otherwise.

### GetClientMachineOk

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) GetClientMachineOk() (*string, bool)`

GetClientMachineOk returns a tuple with the ClientMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientMachine

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) SetClientMachine(v string)`

SetClientMachine sets ClientMachine field to given value.

### HasClientMachine

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) HasClientMachine() bool`

HasClientMachine returns a boolean if a field has been set.

### SetClientMachineNil

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) SetClientMachineNil(b bool)`

 SetClientMachineNil sets the value for ClientMachine to be an explicit nil

### UnsetClientMachine
`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) UnsetClientMachine()`

UnsetClientMachine ensures that no value is present for ClientMachine, not even an explicit nil
### GetUsername

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetAgentPlatform

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) GetAgentPlatform() int32`

GetAgentPlatform returns the AgentPlatform field if non-nil, zero value otherwise.

### GetAgentPlatformOk

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) GetAgentPlatformOk() (*int32, bool)`

GetAgentPlatformOk returns a tuple with the AgentPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentPlatform

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) SetAgentPlatform(v int32)`

SetAgentPlatform sets AgentPlatform field to given value.

### HasAgentPlatform

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) HasAgentPlatform() bool`

HasAgentPlatform returns a boolean if a field has been set.

### GetVersion

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetStatus

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLastSeen

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) GetLastSeen() time.Time`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) GetLastSeenOk() (*time.Time, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) SetLastSeen(v time.Time)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetCapabilities

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) GetCapabilities() []string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) GetCapabilitiesOk() (*[]string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) SetCapabilities(v []string)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### SetCapabilitiesNil

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) SetCapabilitiesNil(b bool)`

 SetCapabilitiesNil sets the value for Capabilities to be an explicit nil

### UnsetCapabilities
`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) UnsetCapabilities()`

UnsetCapabilities ensures that no value is present for Capabilities, not even an explicit nil
### GetBlueprint

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) GetBlueprint() string`

GetBlueprint returns the Blueprint field if non-nil, zero value otherwise.

### GetBlueprintOk

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) GetBlueprintOk() (*string, bool)`

GetBlueprintOk returns a tuple with the Blueprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprint

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) SetBlueprint(v string)`

SetBlueprint sets Blueprint field to given value.

### HasBlueprint

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) HasBlueprint() bool`

HasBlueprint returns a boolean if a field has been set.

### SetBlueprintNil

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) SetBlueprintNil(b bool)`

 SetBlueprintNil sets the value for Blueprint to be an explicit nil

### UnsetBlueprint
`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) UnsetBlueprint()`

UnsetBlueprint ensures that no value is present for Blueprint, not even an explicit nil
### GetThumbprint

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) GetThumbprint() string`

GetThumbprint returns the Thumbprint field if non-nil, zero value otherwise.

### GetThumbprintOk

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) GetThumbprintOk() (*string, bool)`

GetThumbprintOk returns a tuple with the Thumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbprint

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) SetThumbprint(v string)`

SetThumbprint sets Thumbprint field to given value.

### HasThumbprint

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) HasThumbprint() bool`

HasThumbprint returns a boolean if a field has been set.

### SetThumbprintNil

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) SetThumbprintNil(b bool)`

 SetThumbprintNil sets the value for Thumbprint to be an explicit nil

### UnsetThumbprint
`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) UnsetThumbprint()`

UnsetThumbprint ensures that no value is present for Thumbprint, not even an explicit nil
### GetLegacyThumbprint

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) GetLegacyThumbprint() string`

GetLegacyThumbprint returns the LegacyThumbprint field if non-nil, zero value otherwise.

### GetLegacyThumbprintOk

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) GetLegacyThumbprintOk() (*string, bool)`

GetLegacyThumbprintOk returns a tuple with the LegacyThumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacyThumbprint

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) SetLegacyThumbprint(v string)`

SetLegacyThumbprint sets LegacyThumbprint field to given value.

### HasLegacyThumbprint

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) HasLegacyThumbprint() bool`

HasLegacyThumbprint returns a boolean if a field has been set.

### SetLegacyThumbprintNil

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) SetLegacyThumbprintNil(b bool)`

 SetLegacyThumbprintNil sets the value for LegacyThumbprint to be an explicit nil

### UnsetLegacyThumbprint
`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) UnsetLegacyThumbprint()`

UnsetLegacyThumbprint ensures that no value is present for LegacyThumbprint, not even an explicit nil
### GetAuthCertificateReenrollment

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) GetAuthCertificateReenrollment() string`

GetAuthCertificateReenrollment returns the AuthCertificateReenrollment field if non-nil, zero value otherwise.

### GetAuthCertificateReenrollmentOk

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) GetAuthCertificateReenrollmentOk() (*string, bool)`

GetAuthCertificateReenrollmentOk returns a tuple with the AuthCertificateReenrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCertificateReenrollment

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) SetAuthCertificateReenrollment(v string)`

SetAuthCertificateReenrollment sets AuthCertificateReenrollment field to given value.

### HasAuthCertificateReenrollment

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) HasAuthCertificateReenrollment() bool`

HasAuthCertificateReenrollment returns a boolean if a field has been set.

### SetAuthCertificateReenrollmentNil

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) SetAuthCertificateReenrollmentNil(b bool)`

 SetAuthCertificateReenrollmentNil sets the value for AuthCertificateReenrollment to be an explicit nil

### UnsetAuthCertificateReenrollment
`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) UnsetAuthCertificateReenrollment()`

UnsetAuthCertificateReenrollment ensures that no value is present for AuthCertificateReenrollment, not even an explicit nil
### GetLastThumbprintUsed

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) GetLastThumbprintUsed() string`

GetLastThumbprintUsed returns the LastThumbprintUsed field if non-nil, zero value otherwise.

### GetLastThumbprintUsedOk

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) GetLastThumbprintUsedOk() (*string, bool)`

GetLastThumbprintUsedOk returns a tuple with the LastThumbprintUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastThumbprintUsed

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) SetLastThumbprintUsed(v string)`

SetLastThumbprintUsed sets LastThumbprintUsed field to given value.

### HasLastThumbprintUsed

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) HasLastThumbprintUsed() bool`

HasLastThumbprintUsed returns a boolean if a field has been set.

### SetLastThumbprintUsedNil

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) SetLastThumbprintUsedNil(b bool)`

 SetLastThumbprintUsedNil sets the value for LastThumbprintUsed to be an explicit nil

### UnsetLastThumbprintUsed
`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) UnsetLastThumbprintUsed()`

UnsetLastThumbprintUsed ensures that no value is present for LastThumbprintUsed, not even an explicit nil
### GetLastErrorCode

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) GetLastErrorCode() int64`

GetLastErrorCode returns the LastErrorCode field if non-nil, zero value otherwise.

### GetLastErrorCodeOk

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) GetLastErrorCodeOk() (*int64, bool)`

GetLastErrorCodeOk returns a tuple with the LastErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastErrorCode

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) SetLastErrorCode(v int64)`

SetLastErrorCode sets LastErrorCode field to given value.

### HasLastErrorCode

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) HasLastErrorCode() bool`

HasLastErrorCode returns a boolean if a field has been set.

### SetLastErrorCodeNil

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) SetLastErrorCodeNil(b bool)`

 SetLastErrorCodeNil sets the value for LastErrorCode to be an explicit nil

### UnsetLastErrorCode
`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) UnsetLastErrorCode()`

UnsetLastErrorCode ensures that no value is present for LastErrorCode, not even an explicit nil
### GetLastErrorMessage

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) GetLastErrorMessage() string`

GetLastErrorMessage returns the LastErrorMessage field if non-nil, zero value otherwise.

### GetLastErrorMessageOk

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) GetLastErrorMessageOk() (*string, bool)`

GetLastErrorMessageOk returns a tuple with the LastErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastErrorMessage

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) SetLastErrorMessage(v string)`

SetLastErrorMessage sets LastErrorMessage field to given value.

### HasLastErrorMessage

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) HasLastErrorMessage() bool`

HasLastErrorMessage returns a boolean if a field has been set.

### SetLastErrorMessageNil

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) SetLastErrorMessageNil(b bool)`

 SetLastErrorMessageNil sets the value for LastErrorMessage to be an explicit nil

### UnsetLastErrorMessage
`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentResponse) UnsetLastErrorMessage()`

UnsetLastErrorMessage ensures that no value is present for LastErrorMessage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


