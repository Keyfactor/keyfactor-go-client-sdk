# KeyfactorApiModelsWorkflowsSignalConfigurationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SignalName** | Pointer to **string** | The name of the signal. | [optional] 
**RoleIds** | Pointer to **[]int64** | The roles that are allowed to send the signal. | [optional] 

## Methods

### NewKeyfactorApiModelsWorkflowsSignalConfigurationRequest

`func NewKeyfactorApiModelsWorkflowsSignalConfigurationRequest() *KeyfactorApiModelsWorkflowsSignalConfigurationRequest`

NewKeyfactorApiModelsWorkflowsSignalConfigurationRequest instantiates a new KeyfactorApiModelsWorkflowsSignalConfigurationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorApiModelsWorkflowsSignalConfigurationRequestWithDefaults

`func NewKeyfactorApiModelsWorkflowsSignalConfigurationRequestWithDefaults() *KeyfactorApiModelsWorkflowsSignalConfigurationRequest`

NewKeyfactorApiModelsWorkflowsSignalConfigurationRequestWithDefaults instantiates a new KeyfactorApiModelsWorkflowsSignalConfigurationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignalName

`func (o *KeyfactorApiModelsWorkflowsSignalConfigurationRequest) GetSignalName() string`

GetSignalName returns the SignalName field if non-nil, zero value otherwise.

### GetSignalNameOk

`func (o *KeyfactorApiModelsWorkflowsSignalConfigurationRequest) GetSignalNameOk() (*string, bool)`

GetSignalNameOk returns a tuple with the SignalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalName

`func (o *KeyfactorApiModelsWorkflowsSignalConfigurationRequest) SetSignalName(v string)`

SetSignalName sets SignalName field to given value.

### HasSignalName

`func (o *KeyfactorApiModelsWorkflowsSignalConfigurationRequest) HasSignalName() bool`

HasSignalName returns a boolean if a field has been set.

### GetRoleIds

`func (o *KeyfactorApiModelsWorkflowsSignalConfigurationRequest) GetRoleIds() []int64`

GetRoleIds returns the RoleIds field if non-nil, zero value otherwise.

### GetRoleIdsOk

`func (o *KeyfactorApiModelsWorkflowsSignalConfigurationRequest) GetRoleIdsOk() (*[]int64, bool)`

GetRoleIdsOk returns a tuple with the RoleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleIds

`func (o *KeyfactorApiModelsWorkflowsSignalConfigurationRequest) SetRoleIds(v []int64)`

SetRoleIds sets RoleIds field to given value.

### HasRoleIds

`func (o *KeyfactorApiModelsWorkflowsSignalConfigurationRequest) HasRoleIds() bool`

HasRoleIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


