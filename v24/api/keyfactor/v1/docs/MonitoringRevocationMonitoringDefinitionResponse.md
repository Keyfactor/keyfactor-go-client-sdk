# MonitoringRevocationMonitoringDefinitionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**EndpointType** | Pointer to **NullableString** |  | [optional] 
**Location** | Pointer to **NullableString** |  | [optional] 
**Email** | Pointer to [**MonitoringEmailResponse**](MonitoringEmailResponse.md) |  | [optional] 
**Dashboard** | Pointer to [**MonitoringDashboardResponse**](MonitoringDashboardResponse.md) |  | [optional] 
**Schedule** | Pointer to [**KeyfactorCommonSchedulingKeyfactorSchedule**](KeyfactorCommonSchedulingKeyfactorSchedule.md) |  | [optional] 
**OCSPParameters** | Pointer to [**MonitoringOCSPParametersResponse**](MonitoringOCSPParametersResponse.md) |  | [optional] 
**UseWorkflows** | Pointer to **bool** |  | [optional] 
**WorkflowId** | Pointer to **NullableString** |  | [optional] 
**WorkflowName** | Pointer to **NullableString** |  | [optional] 
**WorkflowPublishedVersion** | Pointer to **NullableInt32** |  | [optional] 
**WorkflowEnabled** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewMonitoringRevocationMonitoringDefinitionResponse

`func NewMonitoringRevocationMonitoringDefinitionResponse() *MonitoringRevocationMonitoringDefinitionResponse`

NewMonitoringRevocationMonitoringDefinitionResponse instantiates a new MonitoringRevocationMonitoringDefinitionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringRevocationMonitoringDefinitionResponseWithDefaults

`func NewMonitoringRevocationMonitoringDefinitionResponseWithDefaults() *MonitoringRevocationMonitoringDefinitionResponse`

NewMonitoringRevocationMonitoringDefinitionResponseWithDefaults instantiates a new MonitoringRevocationMonitoringDefinitionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MonitoringRevocationMonitoringDefinitionResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MonitoringRevocationMonitoringDefinitionResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MonitoringRevocationMonitoringDefinitionResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MonitoringRevocationMonitoringDefinitionResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *MonitoringRevocationMonitoringDefinitionResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MonitoringRevocationMonitoringDefinitionResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MonitoringRevocationMonitoringDefinitionResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MonitoringRevocationMonitoringDefinitionResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MonitoringRevocationMonitoringDefinitionResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MonitoringRevocationMonitoringDefinitionResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetEndpointType

`func (o *MonitoringRevocationMonitoringDefinitionResponse) GetEndpointType() string`

GetEndpointType returns the EndpointType field if non-nil, zero value otherwise.

### GetEndpointTypeOk

`func (o *MonitoringRevocationMonitoringDefinitionResponse) GetEndpointTypeOk() (*string, bool)`

GetEndpointTypeOk returns a tuple with the EndpointType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointType

`func (o *MonitoringRevocationMonitoringDefinitionResponse) SetEndpointType(v string)`

SetEndpointType sets EndpointType field to given value.

### HasEndpointType

`func (o *MonitoringRevocationMonitoringDefinitionResponse) HasEndpointType() bool`

HasEndpointType returns a boolean if a field has been set.

### SetEndpointTypeNil

`func (o *MonitoringRevocationMonitoringDefinitionResponse) SetEndpointTypeNil(b bool)`

 SetEndpointTypeNil sets the value for EndpointType to be an explicit nil

### UnsetEndpointType
`func (o *MonitoringRevocationMonitoringDefinitionResponse) UnsetEndpointType()`

UnsetEndpointType ensures that no value is present for EndpointType, not even an explicit nil
### GetLocation

`func (o *MonitoringRevocationMonitoringDefinitionResponse) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *MonitoringRevocationMonitoringDefinitionResponse) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *MonitoringRevocationMonitoringDefinitionResponse) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *MonitoringRevocationMonitoringDefinitionResponse) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *MonitoringRevocationMonitoringDefinitionResponse) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *MonitoringRevocationMonitoringDefinitionResponse) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetEmail

`func (o *MonitoringRevocationMonitoringDefinitionResponse) GetEmail() MonitoringEmailResponse`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *MonitoringRevocationMonitoringDefinitionResponse) GetEmailOk() (*MonitoringEmailResponse, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *MonitoringRevocationMonitoringDefinitionResponse) SetEmail(v MonitoringEmailResponse)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *MonitoringRevocationMonitoringDefinitionResponse) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetDashboard

`func (o *MonitoringRevocationMonitoringDefinitionResponse) GetDashboard() MonitoringDashboardResponse`

GetDashboard returns the Dashboard field if non-nil, zero value otherwise.

### GetDashboardOk

`func (o *MonitoringRevocationMonitoringDefinitionResponse) GetDashboardOk() (*MonitoringDashboardResponse, bool)`

GetDashboardOk returns a tuple with the Dashboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboard

`func (o *MonitoringRevocationMonitoringDefinitionResponse) SetDashboard(v MonitoringDashboardResponse)`

SetDashboard sets Dashboard field to given value.

### HasDashboard

`func (o *MonitoringRevocationMonitoringDefinitionResponse) HasDashboard() bool`

HasDashboard returns a boolean if a field has been set.

### GetSchedule

`func (o *MonitoringRevocationMonitoringDefinitionResponse) GetSchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *MonitoringRevocationMonitoringDefinitionResponse) GetScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *MonitoringRevocationMonitoringDefinitionResponse) SetSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *MonitoringRevocationMonitoringDefinitionResponse) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetOCSPParameters

`func (o *MonitoringRevocationMonitoringDefinitionResponse) GetOCSPParameters() MonitoringOCSPParametersResponse`

GetOCSPParameters returns the OCSPParameters field if non-nil, zero value otherwise.

### GetOCSPParametersOk

`func (o *MonitoringRevocationMonitoringDefinitionResponse) GetOCSPParametersOk() (*MonitoringOCSPParametersResponse, bool)`

GetOCSPParametersOk returns a tuple with the OCSPParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOCSPParameters

`func (o *MonitoringRevocationMonitoringDefinitionResponse) SetOCSPParameters(v MonitoringOCSPParametersResponse)`

SetOCSPParameters sets OCSPParameters field to given value.

### HasOCSPParameters

`func (o *MonitoringRevocationMonitoringDefinitionResponse) HasOCSPParameters() bool`

HasOCSPParameters returns a boolean if a field has been set.

### GetUseWorkflows

`func (o *MonitoringRevocationMonitoringDefinitionResponse) GetUseWorkflows() bool`

GetUseWorkflows returns the UseWorkflows field if non-nil, zero value otherwise.

### GetUseWorkflowsOk

`func (o *MonitoringRevocationMonitoringDefinitionResponse) GetUseWorkflowsOk() (*bool, bool)`

GetUseWorkflowsOk returns a tuple with the UseWorkflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseWorkflows

`func (o *MonitoringRevocationMonitoringDefinitionResponse) SetUseWorkflows(v bool)`

SetUseWorkflows sets UseWorkflows field to given value.

### HasUseWorkflows

`func (o *MonitoringRevocationMonitoringDefinitionResponse) HasUseWorkflows() bool`

HasUseWorkflows returns a boolean if a field has been set.

### GetWorkflowId

`func (o *MonitoringRevocationMonitoringDefinitionResponse) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *MonitoringRevocationMonitoringDefinitionResponse) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *MonitoringRevocationMonitoringDefinitionResponse) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.

### HasWorkflowId

`func (o *MonitoringRevocationMonitoringDefinitionResponse) HasWorkflowId() bool`

HasWorkflowId returns a boolean if a field has been set.

### SetWorkflowIdNil

`func (o *MonitoringRevocationMonitoringDefinitionResponse) SetWorkflowIdNil(b bool)`

 SetWorkflowIdNil sets the value for WorkflowId to be an explicit nil

### UnsetWorkflowId
`func (o *MonitoringRevocationMonitoringDefinitionResponse) UnsetWorkflowId()`

UnsetWorkflowId ensures that no value is present for WorkflowId, not even an explicit nil
### GetWorkflowName

`func (o *MonitoringRevocationMonitoringDefinitionResponse) GetWorkflowName() string`

GetWorkflowName returns the WorkflowName field if non-nil, zero value otherwise.

### GetWorkflowNameOk

`func (o *MonitoringRevocationMonitoringDefinitionResponse) GetWorkflowNameOk() (*string, bool)`

GetWorkflowNameOk returns a tuple with the WorkflowName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowName

`func (o *MonitoringRevocationMonitoringDefinitionResponse) SetWorkflowName(v string)`

SetWorkflowName sets WorkflowName field to given value.

### HasWorkflowName

`func (o *MonitoringRevocationMonitoringDefinitionResponse) HasWorkflowName() bool`

HasWorkflowName returns a boolean if a field has been set.

### SetWorkflowNameNil

`func (o *MonitoringRevocationMonitoringDefinitionResponse) SetWorkflowNameNil(b bool)`

 SetWorkflowNameNil sets the value for WorkflowName to be an explicit nil

### UnsetWorkflowName
`func (o *MonitoringRevocationMonitoringDefinitionResponse) UnsetWorkflowName()`

UnsetWorkflowName ensures that no value is present for WorkflowName, not even an explicit nil
### GetWorkflowPublishedVersion

`func (o *MonitoringRevocationMonitoringDefinitionResponse) GetWorkflowPublishedVersion() int32`

GetWorkflowPublishedVersion returns the WorkflowPublishedVersion field if non-nil, zero value otherwise.

### GetWorkflowPublishedVersionOk

`func (o *MonitoringRevocationMonitoringDefinitionResponse) GetWorkflowPublishedVersionOk() (*int32, bool)`

GetWorkflowPublishedVersionOk returns a tuple with the WorkflowPublishedVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowPublishedVersion

`func (o *MonitoringRevocationMonitoringDefinitionResponse) SetWorkflowPublishedVersion(v int32)`

SetWorkflowPublishedVersion sets WorkflowPublishedVersion field to given value.

### HasWorkflowPublishedVersion

`func (o *MonitoringRevocationMonitoringDefinitionResponse) HasWorkflowPublishedVersion() bool`

HasWorkflowPublishedVersion returns a boolean if a field has been set.

### SetWorkflowPublishedVersionNil

`func (o *MonitoringRevocationMonitoringDefinitionResponse) SetWorkflowPublishedVersionNil(b bool)`

 SetWorkflowPublishedVersionNil sets the value for WorkflowPublishedVersion to be an explicit nil

### UnsetWorkflowPublishedVersion
`func (o *MonitoringRevocationMonitoringDefinitionResponse) UnsetWorkflowPublishedVersion()`

UnsetWorkflowPublishedVersion ensures that no value is present for WorkflowPublishedVersion, not even an explicit nil
### GetWorkflowEnabled

`func (o *MonitoringRevocationMonitoringDefinitionResponse) GetWorkflowEnabled() bool`

GetWorkflowEnabled returns the WorkflowEnabled field if non-nil, zero value otherwise.

### GetWorkflowEnabledOk

`func (o *MonitoringRevocationMonitoringDefinitionResponse) GetWorkflowEnabledOk() (*bool, bool)`

GetWorkflowEnabledOk returns a tuple with the WorkflowEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowEnabled

`func (o *MonitoringRevocationMonitoringDefinitionResponse) SetWorkflowEnabled(v bool)`

SetWorkflowEnabled sets WorkflowEnabled field to given value.

### HasWorkflowEnabled

`func (o *MonitoringRevocationMonitoringDefinitionResponse) HasWorkflowEnabled() bool`

HasWorkflowEnabled returns a boolean if a field has been set.

### SetWorkflowEnabledNil

`func (o *MonitoringRevocationMonitoringDefinitionResponse) SetWorkflowEnabledNil(b bool)`

 SetWorkflowEnabledNil sets the value for WorkflowEnabled to be an explicit nil

### UnsetWorkflowEnabled
`func (o *MonitoringRevocationMonitoringDefinitionResponse) UnsetWorkflowEnabled()`

UnsetWorkflowEnabled ensures that no value is present for WorkflowEnabled, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


