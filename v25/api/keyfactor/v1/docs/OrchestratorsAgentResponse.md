# OrchestratorsAgentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentId** | Pointer to **string** |  | [optional] 
**ClientMachine** | Pointer to **NullableString** |  | [optional] 
**Username** | Pointer to **NullableString** |  | [optional] 
**AgentPlatform** | Pointer to [**CSSCMSCoreEnumsAgentPlatformType**](CSSCMSCoreEnumsAgentPlatformType.md) |  | [optional] 
**Version** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to [**CSSCMSCoreEnumsAgentStatusType**](CSSCMSCoreEnumsAgentStatusType.md) |  | [optional] 
**LastSeen** | Pointer to **time.Time** |  | [optional] 
**Capabilities** | Pointer to **[]string** |  | [optional] 
**Blueprint** | Pointer to **NullableString** |  | [optional] 
**Thumbprint** | Pointer to **NullableString** |  | [optional] 
**LegacyThumbprint** | Pointer to **NullableString** |  | [optional] 
**AuthCertificateReenrollment** | Pointer to **NullableString** |  | [optional] 
**LastThumbprintUsed** | Pointer to **NullableString** |  | [optional] 
**LastErrorCode** | Pointer to **NullableInt64** |  | [optional] 
**LastErrorMessage** | Pointer to **NullableString** |  | [optional] 
**JobTypes** | Pointer to [**[]OrchestratorsJobTypesResponse**](OrchestratorsJobTypesResponse.md) |  | [optional] 

## Methods

### NewOrchestratorsAgentResponse

`func NewOrchestratorsAgentResponse() *OrchestratorsAgentResponse`

NewOrchestratorsAgentResponse instantiates a new OrchestratorsAgentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrchestratorsAgentResponseWithDefaults

`func NewOrchestratorsAgentResponseWithDefaults() *OrchestratorsAgentResponse`

NewOrchestratorsAgentResponseWithDefaults instantiates a new OrchestratorsAgentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentId

`func (o *OrchestratorsAgentResponse) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *OrchestratorsAgentResponse) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *OrchestratorsAgentResponse) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *OrchestratorsAgentResponse) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.

### GetClientMachine

`func (o *OrchestratorsAgentResponse) GetClientMachine() string`

GetClientMachine returns the ClientMachine field if non-nil, zero value otherwise.

### GetClientMachineOk

`func (o *OrchestratorsAgentResponse) GetClientMachineOk() (*string, bool)`

GetClientMachineOk returns a tuple with the ClientMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientMachine

`func (o *OrchestratorsAgentResponse) SetClientMachine(v string)`

SetClientMachine sets ClientMachine field to given value.

### HasClientMachine

`func (o *OrchestratorsAgentResponse) HasClientMachine() bool`

HasClientMachine returns a boolean if a field has been set.

### SetClientMachineNil

`func (o *OrchestratorsAgentResponse) SetClientMachineNil(b bool)`

 SetClientMachineNil sets the value for ClientMachine to be an explicit nil

### UnsetClientMachine
`func (o *OrchestratorsAgentResponse) UnsetClientMachine()`

UnsetClientMachine ensures that no value is present for ClientMachine, not even an explicit nil
### GetUsername

`func (o *OrchestratorsAgentResponse) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *OrchestratorsAgentResponse) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *OrchestratorsAgentResponse) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *OrchestratorsAgentResponse) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *OrchestratorsAgentResponse) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *OrchestratorsAgentResponse) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetAgentPlatform

`func (o *OrchestratorsAgentResponse) GetAgentPlatform() CSSCMSCoreEnumsAgentPlatformType`

GetAgentPlatform returns the AgentPlatform field if non-nil, zero value otherwise.

### GetAgentPlatformOk

`func (o *OrchestratorsAgentResponse) GetAgentPlatformOk() (*CSSCMSCoreEnumsAgentPlatformType, bool)`

GetAgentPlatformOk returns a tuple with the AgentPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentPlatform

`func (o *OrchestratorsAgentResponse) SetAgentPlatform(v CSSCMSCoreEnumsAgentPlatformType)`

SetAgentPlatform sets AgentPlatform field to given value.

### HasAgentPlatform

`func (o *OrchestratorsAgentResponse) HasAgentPlatform() bool`

HasAgentPlatform returns a boolean if a field has been set.

### GetVersion

`func (o *OrchestratorsAgentResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *OrchestratorsAgentResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *OrchestratorsAgentResponse) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *OrchestratorsAgentResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *OrchestratorsAgentResponse) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *OrchestratorsAgentResponse) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetStatus

`func (o *OrchestratorsAgentResponse) GetStatus() CSSCMSCoreEnumsAgentStatusType`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrchestratorsAgentResponse) GetStatusOk() (*CSSCMSCoreEnumsAgentStatusType, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrchestratorsAgentResponse) SetStatus(v CSSCMSCoreEnumsAgentStatusType)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OrchestratorsAgentResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLastSeen

`func (o *OrchestratorsAgentResponse) GetLastSeen() time.Time`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *OrchestratorsAgentResponse) GetLastSeenOk() (*time.Time, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *OrchestratorsAgentResponse) SetLastSeen(v time.Time)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *OrchestratorsAgentResponse) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetCapabilities

`func (o *OrchestratorsAgentResponse) GetCapabilities() []string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *OrchestratorsAgentResponse) GetCapabilitiesOk() (*[]string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *OrchestratorsAgentResponse) SetCapabilities(v []string)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *OrchestratorsAgentResponse) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### SetCapabilitiesNil

`func (o *OrchestratorsAgentResponse) SetCapabilitiesNil(b bool)`

 SetCapabilitiesNil sets the value for Capabilities to be an explicit nil

### UnsetCapabilities
`func (o *OrchestratorsAgentResponse) UnsetCapabilities()`

UnsetCapabilities ensures that no value is present for Capabilities, not even an explicit nil
### GetBlueprint

`func (o *OrchestratorsAgentResponse) GetBlueprint() string`

GetBlueprint returns the Blueprint field if non-nil, zero value otherwise.

### GetBlueprintOk

`func (o *OrchestratorsAgentResponse) GetBlueprintOk() (*string, bool)`

GetBlueprintOk returns a tuple with the Blueprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprint

`func (o *OrchestratorsAgentResponse) SetBlueprint(v string)`

SetBlueprint sets Blueprint field to given value.

### HasBlueprint

`func (o *OrchestratorsAgentResponse) HasBlueprint() bool`

HasBlueprint returns a boolean if a field has been set.

### SetBlueprintNil

`func (o *OrchestratorsAgentResponse) SetBlueprintNil(b bool)`

 SetBlueprintNil sets the value for Blueprint to be an explicit nil

### UnsetBlueprint
`func (o *OrchestratorsAgentResponse) UnsetBlueprint()`

UnsetBlueprint ensures that no value is present for Blueprint, not even an explicit nil
### GetThumbprint

`func (o *OrchestratorsAgentResponse) GetThumbprint() string`

GetThumbprint returns the Thumbprint field if non-nil, zero value otherwise.

### GetThumbprintOk

`func (o *OrchestratorsAgentResponse) GetThumbprintOk() (*string, bool)`

GetThumbprintOk returns a tuple with the Thumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbprint

`func (o *OrchestratorsAgentResponse) SetThumbprint(v string)`

SetThumbprint sets Thumbprint field to given value.

### HasThumbprint

`func (o *OrchestratorsAgentResponse) HasThumbprint() bool`

HasThumbprint returns a boolean if a field has been set.

### SetThumbprintNil

`func (o *OrchestratorsAgentResponse) SetThumbprintNil(b bool)`

 SetThumbprintNil sets the value for Thumbprint to be an explicit nil

### UnsetThumbprint
`func (o *OrchestratorsAgentResponse) UnsetThumbprint()`

UnsetThumbprint ensures that no value is present for Thumbprint, not even an explicit nil
### GetLegacyThumbprint

`func (o *OrchestratorsAgentResponse) GetLegacyThumbprint() string`

GetLegacyThumbprint returns the LegacyThumbprint field if non-nil, zero value otherwise.

### GetLegacyThumbprintOk

`func (o *OrchestratorsAgentResponse) GetLegacyThumbprintOk() (*string, bool)`

GetLegacyThumbprintOk returns a tuple with the LegacyThumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacyThumbprint

`func (o *OrchestratorsAgentResponse) SetLegacyThumbprint(v string)`

SetLegacyThumbprint sets LegacyThumbprint field to given value.

### HasLegacyThumbprint

`func (o *OrchestratorsAgentResponse) HasLegacyThumbprint() bool`

HasLegacyThumbprint returns a boolean if a field has been set.

### SetLegacyThumbprintNil

`func (o *OrchestratorsAgentResponse) SetLegacyThumbprintNil(b bool)`

 SetLegacyThumbprintNil sets the value for LegacyThumbprint to be an explicit nil

### UnsetLegacyThumbprint
`func (o *OrchestratorsAgentResponse) UnsetLegacyThumbprint()`

UnsetLegacyThumbprint ensures that no value is present for LegacyThumbprint, not even an explicit nil
### GetAuthCertificateReenrollment

`func (o *OrchestratorsAgentResponse) GetAuthCertificateReenrollment() string`

GetAuthCertificateReenrollment returns the AuthCertificateReenrollment field if non-nil, zero value otherwise.

### GetAuthCertificateReenrollmentOk

`func (o *OrchestratorsAgentResponse) GetAuthCertificateReenrollmentOk() (*string, bool)`

GetAuthCertificateReenrollmentOk returns a tuple with the AuthCertificateReenrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCertificateReenrollment

`func (o *OrchestratorsAgentResponse) SetAuthCertificateReenrollment(v string)`

SetAuthCertificateReenrollment sets AuthCertificateReenrollment field to given value.

### HasAuthCertificateReenrollment

`func (o *OrchestratorsAgentResponse) HasAuthCertificateReenrollment() bool`

HasAuthCertificateReenrollment returns a boolean if a field has been set.

### SetAuthCertificateReenrollmentNil

`func (o *OrchestratorsAgentResponse) SetAuthCertificateReenrollmentNil(b bool)`

 SetAuthCertificateReenrollmentNil sets the value for AuthCertificateReenrollment to be an explicit nil

### UnsetAuthCertificateReenrollment
`func (o *OrchestratorsAgentResponse) UnsetAuthCertificateReenrollment()`

UnsetAuthCertificateReenrollment ensures that no value is present for AuthCertificateReenrollment, not even an explicit nil
### GetLastThumbprintUsed

`func (o *OrchestratorsAgentResponse) GetLastThumbprintUsed() string`

GetLastThumbprintUsed returns the LastThumbprintUsed field if non-nil, zero value otherwise.

### GetLastThumbprintUsedOk

`func (o *OrchestratorsAgentResponse) GetLastThumbprintUsedOk() (*string, bool)`

GetLastThumbprintUsedOk returns a tuple with the LastThumbprintUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastThumbprintUsed

`func (o *OrchestratorsAgentResponse) SetLastThumbprintUsed(v string)`

SetLastThumbprintUsed sets LastThumbprintUsed field to given value.

### HasLastThumbprintUsed

`func (o *OrchestratorsAgentResponse) HasLastThumbprintUsed() bool`

HasLastThumbprintUsed returns a boolean if a field has been set.

### SetLastThumbprintUsedNil

`func (o *OrchestratorsAgentResponse) SetLastThumbprintUsedNil(b bool)`

 SetLastThumbprintUsedNil sets the value for LastThumbprintUsed to be an explicit nil

### UnsetLastThumbprintUsed
`func (o *OrchestratorsAgentResponse) UnsetLastThumbprintUsed()`

UnsetLastThumbprintUsed ensures that no value is present for LastThumbprintUsed, not even an explicit nil
### GetLastErrorCode

`func (o *OrchestratorsAgentResponse) GetLastErrorCode() int64`

GetLastErrorCode returns the LastErrorCode field if non-nil, zero value otherwise.

### GetLastErrorCodeOk

`func (o *OrchestratorsAgentResponse) GetLastErrorCodeOk() (*int64, bool)`

GetLastErrorCodeOk returns a tuple with the LastErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastErrorCode

`func (o *OrchestratorsAgentResponse) SetLastErrorCode(v int64)`

SetLastErrorCode sets LastErrorCode field to given value.

### HasLastErrorCode

`func (o *OrchestratorsAgentResponse) HasLastErrorCode() bool`

HasLastErrorCode returns a boolean if a field has been set.

### SetLastErrorCodeNil

`func (o *OrchestratorsAgentResponse) SetLastErrorCodeNil(b bool)`

 SetLastErrorCodeNil sets the value for LastErrorCode to be an explicit nil

### UnsetLastErrorCode
`func (o *OrchestratorsAgentResponse) UnsetLastErrorCode()`

UnsetLastErrorCode ensures that no value is present for LastErrorCode, not even an explicit nil
### GetLastErrorMessage

`func (o *OrchestratorsAgentResponse) GetLastErrorMessage() string`

GetLastErrorMessage returns the LastErrorMessage field if non-nil, zero value otherwise.

### GetLastErrorMessageOk

`func (o *OrchestratorsAgentResponse) GetLastErrorMessageOk() (*string, bool)`

GetLastErrorMessageOk returns a tuple with the LastErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastErrorMessage

`func (o *OrchestratorsAgentResponse) SetLastErrorMessage(v string)`

SetLastErrorMessage sets LastErrorMessage field to given value.

### HasLastErrorMessage

`func (o *OrchestratorsAgentResponse) HasLastErrorMessage() bool`

HasLastErrorMessage returns a boolean if a field has been set.

### SetLastErrorMessageNil

`func (o *OrchestratorsAgentResponse) SetLastErrorMessageNil(b bool)`

 SetLastErrorMessageNil sets the value for LastErrorMessage to be an explicit nil

### UnsetLastErrorMessage
`func (o *OrchestratorsAgentResponse) UnsetLastErrorMessage()`

UnsetLastErrorMessage ensures that no value is present for LastErrorMessage, not even an explicit nil
### GetJobTypes

`func (o *OrchestratorsAgentResponse) GetJobTypes() []OrchestratorsJobTypesResponse`

GetJobTypes returns the JobTypes field if non-nil, zero value otherwise.

### GetJobTypesOk

`func (o *OrchestratorsAgentResponse) GetJobTypesOk() (*[]OrchestratorsJobTypesResponse, bool)`

GetJobTypesOk returns a tuple with the JobTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTypes

`func (o *OrchestratorsAgentResponse) SetJobTypes(v []OrchestratorsJobTypesResponse)`

SetJobTypes sets JobTypes field to given value.

### HasJobTypes

`func (o *OrchestratorsAgentResponse) HasJobTypes() bool`

HasJobTypes returns a boolean if a field has been set.

### SetJobTypesNil

`func (o *OrchestratorsAgentResponse) SetJobTypesNil(b bool)`

 SetJobTypesNil sets the value for JobTypes to be an explicit nil

### UnsetJobTypes
`func (o *OrchestratorsAgentResponse) UnsetJobTypes()`

UnsetJobTypes ensures that no value is present for JobTypes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


