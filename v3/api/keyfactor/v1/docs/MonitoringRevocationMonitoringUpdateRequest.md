# MonitoringRevocationMonitoringUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  | 
**EndpointType** | **string** |  | 
**Location** | **string** |  | 
**Email** | Pointer to [**MonitoringEmailRequest**](MonitoringEmailRequest.md) |  | [optional] 
**Dashboard** | [**MonitoringDashboardRequest**](MonitoringDashboardRequest.md) |  | 
**Schedule** | Pointer to [**KeyfactorCommonSchedulingKeyfactorSchedule**](KeyfactorCommonSchedulingKeyfactorSchedule.md) |  | [optional] 
**OCSPParameters** | Pointer to [**MonitoringOCSPParametersRequest**](MonitoringOCSPParametersRequest.md) |  | [optional] 
**UseWorkflows** | Pointer to **bool** |  | [optional] 

## Methods

### NewMonitoringRevocationMonitoringUpdateRequest

`func NewMonitoringRevocationMonitoringUpdateRequest(name string, endpointType string, location string, dashboard MonitoringDashboardRequest, ) *MonitoringRevocationMonitoringUpdateRequest`

NewMonitoringRevocationMonitoringUpdateRequest instantiates a new MonitoringRevocationMonitoringUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringRevocationMonitoringUpdateRequestWithDefaults

`func NewMonitoringRevocationMonitoringUpdateRequestWithDefaults() *MonitoringRevocationMonitoringUpdateRequest`

NewMonitoringRevocationMonitoringUpdateRequestWithDefaults instantiates a new MonitoringRevocationMonitoringUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MonitoringRevocationMonitoringUpdateRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MonitoringRevocationMonitoringUpdateRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MonitoringRevocationMonitoringUpdateRequest) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MonitoringRevocationMonitoringUpdateRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *MonitoringRevocationMonitoringUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MonitoringRevocationMonitoringUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MonitoringRevocationMonitoringUpdateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetEndpointType

`func (o *MonitoringRevocationMonitoringUpdateRequest) GetEndpointType() string`

GetEndpointType returns the EndpointType field if non-nil, zero value otherwise.

### GetEndpointTypeOk

`func (o *MonitoringRevocationMonitoringUpdateRequest) GetEndpointTypeOk() (*string, bool)`

GetEndpointTypeOk returns a tuple with the EndpointType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointType

`func (o *MonitoringRevocationMonitoringUpdateRequest) SetEndpointType(v string)`

SetEndpointType sets EndpointType field to given value.


### GetLocation

`func (o *MonitoringRevocationMonitoringUpdateRequest) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *MonitoringRevocationMonitoringUpdateRequest) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *MonitoringRevocationMonitoringUpdateRequest) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetEmail

`func (o *MonitoringRevocationMonitoringUpdateRequest) GetEmail() MonitoringEmailRequest`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *MonitoringRevocationMonitoringUpdateRequest) GetEmailOk() (*MonitoringEmailRequest, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *MonitoringRevocationMonitoringUpdateRequest) SetEmail(v MonitoringEmailRequest)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *MonitoringRevocationMonitoringUpdateRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetDashboard

`func (o *MonitoringRevocationMonitoringUpdateRequest) GetDashboard() MonitoringDashboardRequest`

GetDashboard returns the Dashboard field if non-nil, zero value otherwise.

### GetDashboardOk

`func (o *MonitoringRevocationMonitoringUpdateRequest) GetDashboardOk() (*MonitoringDashboardRequest, bool)`

GetDashboardOk returns a tuple with the Dashboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboard

`func (o *MonitoringRevocationMonitoringUpdateRequest) SetDashboard(v MonitoringDashboardRequest)`

SetDashboard sets Dashboard field to given value.


### GetSchedule

`func (o *MonitoringRevocationMonitoringUpdateRequest) GetSchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *MonitoringRevocationMonitoringUpdateRequest) GetScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *MonitoringRevocationMonitoringUpdateRequest) SetSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *MonitoringRevocationMonitoringUpdateRequest) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetOCSPParameters

`func (o *MonitoringRevocationMonitoringUpdateRequest) GetOCSPParameters() MonitoringOCSPParametersRequest`

GetOCSPParameters returns the OCSPParameters field if non-nil, zero value otherwise.

### GetOCSPParametersOk

`func (o *MonitoringRevocationMonitoringUpdateRequest) GetOCSPParametersOk() (*MonitoringOCSPParametersRequest, bool)`

GetOCSPParametersOk returns a tuple with the OCSPParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOCSPParameters

`func (o *MonitoringRevocationMonitoringUpdateRequest) SetOCSPParameters(v MonitoringOCSPParametersRequest)`

SetOCSPParameters sets OCSPParameters field to given value.

### HasOCSPParameters

`func (o *MonitoringRevocationMonitoringUpdateRequest) HasOCSPParameters() bool`

HasOCSPParameters returns a boolean if a field has been set.

### GetUseWorkflows

`func (o *MonitoringRevocationMonitoringUpdateRequest) GetUseWorkflows() bool`

GetUseWorkflows returns the UseWorkflows field if non-nil, zero value otherwise.

### GetUseWorkflowsOk

`func (o *MonitoringRevocationMonitoringUpdateRequest) GetUseWorkflowsOk() (*bool, bool)`

GetUseWorkflowsOk returns a tuple with the UseWorkflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseWorkflows

`func (o *MonitoringRevocationMonitoringUpdateRequest) SetUseWorkflows(v bool)`

SetUseWorkflows sets UseWorkflows field to given value.

### HasUseWorkflows

`func (o *MonitoringRevocationMonitoringUpdateRequest) HasUseWorkflows() bool`

HasUseWorkflows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


