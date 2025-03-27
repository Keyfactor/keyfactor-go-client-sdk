# MonitoringRevocationMonitoringCreationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**EndpointType** | **string** |  | 
**Location** | **string** |  | 
**Email** | Pointer to [**MonitoringEmailRequest**](MonitoringEmailRequest.md) |  | [optional] 
**Dashboard** | [**MonitoringDashboardRequest**](MonitoringDashboardRequest.md) |  | 
**Schedule** | Pointer to [**KeyfactorCommonSchedulingKeyfactorSchedule**](KeyfactorCommonSchedulingKeyfactorSchedule.md) |  | [optional] 
**OCSPParameters** | Pointer to [**MonitoringOCSPParametersRequest**](MonitoringOCSPParametersRequest.md) |  | [optional] 
**UseWorkflows** | Pointer to **bool** |  | [optional] 

## Methods

### NewMonitoringRevocationMonitoringCreationRequest

`func NewMonitoringRevocationMonitoringCreationRequest(name string, endpointType string, location string, dashboard MonitoringDashboardRequest, ) *MonitoringRevocationMonitoringCreationRequest`

NewMonitoringRevocationMonitoringCreationRequest instantiates a new MonitoringRevocationMonitoringCreationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringRevocationMonitoringCreationRequestWithDefaults

`func NewMonitoringRevocationMonitoringCreationRequestWithDefaults() *MonitoringRevocationMonitoringCreationRequest`

NewMonitoringRevocationMonitoringCreationRequestWithDefaults instantiates a new MonitoringRevocationMonitoringCreationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MonitoringRevocationMonitoringCreationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MonitoringRevocationMonitoringCreationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MonitoringRevocationMonitoringCreationRequest) SetName(v string)`

SetName sets Name field to given value.


### GetEndpointType

`func (o *MonitoringRevocationMonitoringCreationRequest) GetEndpointType() string`

GetEndpointType returns the EndpointType field if non-nil, zero value otherwise.

### GetEndpointTypeOk

`func (o *MonitoringRevocationMonitoringCreationRequest) GetEndpointTypeOk() (*string, bool)`

GetEndpointTypeOk returns a tuple with the EndpointType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointType

`func (o *MonitoringRevocationMonitoringCreationRequest) SetEndpointType(v string)`

SetEndpointType sets EndpointType field to given value.


### GetLocation

`func (o *MonitoringRevocationMonitoringCreationRequest) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *MonitoringRevocationMonitoringCreationRequest) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *MonitoringRevocationMonitoringCreationRequest) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetEmail

`func (o *MonitoringRevocationMonitoringCreationRequest) GetEmail() MonitoringEmailRequest`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *MonitoringRevocationMonitoringCreationRequest) GetEmailOk() (*MonitoringEmailRequest, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *MonitoringRevocationMonitoringCreationRequest) SetEmail(v MonitoringEmailRequest)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *MonitoringRevocationMonitoringCreationRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetDashboard

`func (o *MonitoringRevocationMonitoringCreationRequest) GetDashboard() MonitoringDashboardRequest`

GetDashboard returns the Dashboard field if non-nil, zero value otherwise.

### GetDashboardOk

`func (o *MonitoringRevocationMonitoringCreationRequest) GetDashboardOk() (*MonitoringDashboardRequest, bool)`

GetDashboardOk returns a tuple with the Dashboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboard

`func (o *MonitoringRevocationMonitoringCreationRequest) SetDashboard(v MonitoringDashboardRequest)`

SetDashboard sets Dashboard field to given value.


### GetSchedule

`func (o *MonitoringRevocationMonitoringCreationRequest) GetSchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *MonitoringRevocationMonitoringCreationRequest) GetScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *MonitoringRevocationMonitoringCreationRequest) SetSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *MonitoringRevocationMonitoringCreationRequest) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetOCSPParameters

`func (o *MonitoringRevocationMonitoringCreationRequest) GetOCSPParameters() MonitoringOCSPParametersRequest`

GetOCSPParameters returns the OCSPParameters field if non-nil, zero value otherwise.

### GetOCSPParametersOk

`func (o *MonitoringRevocationMonitoringCreationRequest) GetOCSPParametersOk() (*MonitoringOCSPParametersRequest, bool)`

GetOCSPParametersOk returns a tuple with the OCSPParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOCSPParameters

`func (o *MonitoringRevocationMonitoringCreationRequest) SetOCSPParameters(v MonitoringOCSPParametersRequest)`

SetOCSPParameters sets OCSPParameters field to given value.

### HasOCSPParameters

`func (o *MonitoringRevocationMonitoringCreationRequest) HasOCSPParameters() bool`

HasOCSPParameters returns a boolean if a field has been set.

### GetUseWorkflows

`func (o *MonitoringRevocationMonitoringCreationRequest) GetUseWorkflows() bool`

GetUseWorkflows returns the UseWorkflows field if non-nil, zero value otherwise.

### GetUseWorkflowsOk

`func (o *MonitoringRevocationMonitoringCreationRequest) GetUseWorkflowsOk() (*bool, bool)`

GetUseWorkflowsOk returns a tuple with the UseWorkflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseWorkflows

`func (o *MonitoringRevocationMonitoringCreationRequest) SetUseWorkflows(v bool)`

SetUseWorkflows sets UseWorkflows field to given value.

### HasUseWorkflows

`func (o *MonitoringRevocationMonitoringCreationRequest) HasUseWorkflows() bool`

HasUseWorkflows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


